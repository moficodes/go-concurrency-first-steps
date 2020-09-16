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

type fileContent struct {
	body *io.ReadCloser
	url  string
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

	urlChan := make(chan string)
	downloadChan := make(chan fileContent)

	go func(urlChan chan<- string) {
		for _, url := range urls {
			urlChan <- url
		}
	}(urlChan)

	var wg sync.WaitGroup
	wg.Add(2 * len(urls))
	go func() {
		for {
			select {
			case url := <-urlChan:
				go download(downloadChan, url, &wg)
			case content := <-downloadChan:
				go write(content, &wg)
			}
		}
	}()

	fmt.Println("Downloading: ")
	wg.Wait()
	fmt.Println("Done!")
}

func download(out chan<- fileContent, URL string, wg *sync.WaitGroup) {
	defer wg.Done()
	out <- downloadFile(URL)
}

func write(content fileContent, wg *sync.WaitGroup) {
	defer wg.Done()
	writeFile(content)
}

func checkFolder(folder string, count int) {

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

func downloadFile(URL string) fileContent {
	res, err := http.Get(URL)
	if err != nil {
		log.Println(err)
	}
	return fileContent{body: &res.Body, url: URL}
}

func writeFile(content fileContent) error {
	defer (*content.body).Close()
	urlParts := strings.Split(content.url, "/")
	fileName := urlParts[len(urlParts)-1]
	folderName := "images/"
	file, err := os.Create(folderName + fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, *content.body)
	if err != nil {
		return err
	}
	return nil
}

/*
time go run main.go
Downloading:
Done!
go run main.go  10.94s user 10.74s system 5% cpu 6:54.86 total
*/
