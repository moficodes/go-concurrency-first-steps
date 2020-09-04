package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	run()
}

func run() {
	dataFile := "data/imageurls.txt"
	urls, err := loadUrls(dataFile)
	if err != nil {
		log.Fatal(err)
	}

	errs := make([]error, 0)

	printInterval := len(urls) / 50

	fmt.Print("Downloading: [")
	for i, url := range urls {
		if i%printInterval == 0 {
			fmt.Print("=")
		}
		if err := downloadFile(url); err != nil {
			errs = append(errs, err)
		}
	}
	fmt.Println("]")

	fmt.Println("Done!")
	for _, err := range errs {
		log.Println(err)
	}
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

func downloadFile(URL string) error {
	res, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	urlParts := strings.Split(URL, "/")
	fileName := urlParts[len(urlParts)-1]
	folderName := "images/"
	file, err := os.Create(folderName + fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, res.Body)
	if err != nil {
		return err
	}

	return nil
}
