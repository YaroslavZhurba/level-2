package wget

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func downloadUrl(dest string, url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	err = os.MkdirAll(dest, os.ModePerm)
	if err != nil {
		return "", err
	}

	var file *os.File
	if strings.Contains(response.Header.Get("content-type"), "text/html") {
		file, err = os.Create(dest + "/index.html")
	} else {
		file, err = os.Create(dest + "/file.txt")
	}
	if err != nil {
		return "", err
	}

	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return "", err
	}
	return file.Name(), nil
}



func RunWget(dest string, url string) {
	filename, err := downloadUrl(dest, url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Filemane: %s\n", filename)
}
