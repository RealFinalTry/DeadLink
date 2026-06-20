package main
import (
        "fmt"
        "os"
        "net/http"
        "regexp"
        "time"
        "strings"
)
type Result struct {
        URL string
        Code int
}
func check(url string, ch chan Result) {
        client := http.Client{
                Timeout : 10 * time.Second,
        }
        if !strings.HasPrefix(url, "http") {
                url = "https://" + url
        }
        resp, err := client.Get(url)
        if err != nil {
                ch <- Result{URL : url , Code : 0 }
                return
        }
        defer resp.Body.Close()
        ch <- Result{URL: url, Code: resp.StatusCode}
}
func main() {
        var filename string
        var forbidden int
        var alive int
        var dead int
        fmt.Println("enter file name : \nexample : example.txt")
        fmt.Scanln(&filename)
        fmt.Printf("---(%s)---\n",filename)
        data,err := os.ReadFile(filename)
        if err != nil {
                fmt.Println("error :", err)
                return
        }
        content := string(data)


        urlPattern := `(?:https?://)?(?:www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b(?:[-a-zA-Z0-9()@:%_\+.~#?&/=]*)`
        re := regexp.MustCompile(urlPattern)
        links := re.FindAllString(content,-1)
        cd := make(chan Result)
        for _,link := range links {
                url := link
                go check(url,cd)
        }
        for i := 0; i < len(links); i++ {
                result := <- cd
                if result.Code == 200 {
                        alive++
                        fmt.Println("is alive :",result.URL , "\n-----")
                } else if result.Code == 403 {
                        forbidden++
                        fmt.Println("error : 403 forbidden",result.URL , "\n-----")
                } else if result.Code == 404 {
                        dead++
                        fmt.Println("is dead :",result.URL,"\n-----")
                } else {
                        dead++
                        fmt.Println("error :", result.URL,"\n-----")
                }
        }
        total := alive + dead + forbidden
        fmt.Printf("Total: %d | Alive: %d | Dead: %d | Forbidden: %d\n", total, alive, dead, forbidden)
}
