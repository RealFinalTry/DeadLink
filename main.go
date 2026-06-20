package main
import (
	"fmt"
	"os"
	"net/http"
	"regexp"
	"time"
	"strings"
)
func check(url string) int {
	client := http.Client{
		Timeout : 10 * time.Second,
	}
	if !strings.HasPrefix(url, "http") {
    url = "https://" + url
    }
	resp, err := client.Get(url)
    if err != nil {
        return 0 
    }
    defer resp.Body.Close()
    return resp.StatusCode
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

	for _,link := range links {
		url := link
                code := check(url)
		if code == 200 {
                	alive++
                	fmt.Println("is alive :",url , "\n-----")
        	} else if code == 403 {
                	forbidden++
                	fmt.Println("error : 403 forbidden",url , "\n-----")
        	} else if code == 404 {
                	dead++
                	fmt.Println("is dead :",url,"\n-----")
        	} else {
			dead++
                	fmt.Println("error :", url,"\n-----")
        	}
	}
	total := alive + dead + forbidden
	fmt.Printf("Total: %d | Alive: %d | Dead: %d | Forbidden: %d\n", total, alive, dead, forbidden)
}
