package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	run()
}

func run() {
	if err := setup("images"); err != nil {
		log.Fatal(err)
	}
	dataFile := "data/imageurls.txt"
	urls, err := loadUrls(dataFile)
	if err != nil {
		log.Fatal(err)
	}

	errs := make([]error, 0)

	printInterval := len(urls)/50 + 1

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

func setup(folder string) error {
	if err := cleanup(folder); err != nil {
		return err
	}
	if err := createFolder(folder); err != nil {
		return err
	}
	return nil
}

// we will run this before every run so that our directory gets cleaned up
func cleanup(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	d.Close()
	return os.RemoveAll(dir)
}

func createFolder(folder string) error {
	return os.Mkdir(folder, 0755)
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

/*
time go run main.go
Downloading: [==================================================]
Done!
go run main.go  11.63s user 11.41s system 5% cpu 7:39.06 total
*/
