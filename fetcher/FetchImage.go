package fetcher

import (
	"os"
	"net/http"
	"log"
	"strings"
	"io"
	"models"
	"fmt"
	"sync"
)

func fetchImage(post models.Post, wg *sync.WaitGroup)  {

	fmt.Println("Fetching image "+post.Url)

	filename := strings.Split(post.Url, "/")
	extension := strings.Split(filename[len(filename)-1], ".")
	extensionNoQuery := strings.Split(extension[1], "?")
	pwd, _ := os.Getwd()
	filePath := pwd + "/../images/" + post.Id + "." + extensionNoQuery[0]
	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		defer wg.Done()
		fmt.Println("Image " + post.Url + " already exists, skipping")
		return;
	}
	response, e := http.Get(post.Url)
	if e != nil {
		log.Fatal(e)
	}

	defer response.Body.Close()

	//open a file for writing
	file, err := os.Create(filePath)

	if err != nil {
		log.Fatal(err)
	}
	// Use io.Copy to just dump the response body to the file. This supports huge files
	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
	defer wg.Done()
}
