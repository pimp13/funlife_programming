package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("No arguments provided")
	}

	word := os.Args[1]
	url := fmt.Sprintf("https://www.google.com/search?sca_esv=09a4dc9e07e27686&sxsrf=ANbL-n5TvPz9gQvQlMwtdM8MjnlqHaOWTw:1769170109522&udm=2&fbs=ADc_l-aN0CWEZBOHjofHoaMMDiKpaEWjvZ2Py1XXV8d8KvlI3vWUtYx0DZdicpfE1faGYenqWn-q4MFiFFtvJjTKeAVxBf9XF8ByrMpEedseJb6C24e7QdJQdIE3TPpl5mEwf0EBCRLNeEhQy5amsEYIcoTye3rrZrd3IP3OYha6_rH_GIVOU8GK5eecabclKqwVhxmAmgM5&q=%s", word)

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatalf("failed to get send request status code %d", res.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	secoundImg := doc.Find("img").Eq(1)
	if secoundImg.Length() > 0 {
		src, ok := secoundImg.Attr("src")
		if ok {
			fmt.Printf("%s \n", src)
		} else {
			fmt.Println("Second image has no src attribute")
		}
	} else {
		fmt.Println("Less than 2 images found on page")
	}

}
