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
	if err := setup("images"); err != nil {
		log.Fatal(err)
	}
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
