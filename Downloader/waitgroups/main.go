package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

var wg sync.WaitGroup

//func init() {
//	runtime.GOMAXPROCS(1)
//}

func main() {
	run()
}

func run() {
	dataFile := "data/imageurls.txt"
	urls, err := loadUrls(dataFile)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Downloading: ")
	wg.Add(len(urls))
	for _, url := range urls {
		go downloadFile(url)
	}

	wg.Wait()
	fmt.Println("Done!")
}

func loadUrls(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)

	if err != nil {
		return nil, err
	}

	data := string(b)

	return strings.Split(data, "\n"), nil
}

func downloadFile(URL string) {
	defer wg.Done()
	res, err := http.Get(URL)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	urlParts := strings.Split(URL, "/")
	fileName := urlParts[len(urlParts)-1]
	folderName := "images/"
	file, err := os.Create(folderName + fileName)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	_, err = io.Copy(file, res.Body)
	if err != nil {
		log.Println(err)
	}
}
