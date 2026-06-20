package main
import (
	"fmt"
	"os"
	"net/http"
	"regexp"
)
func check(url string) int {
    resp, err := http.Get(url)
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

	
	re := regexp.MustCompile(`https?://[^\s)]+`)
	links := re.FindAllStringSubmatch(content,-1)

	for _,link := range links {
		url := link[1]
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
