package main

import (
	"bufio"
	"io"
	"net/http"
	"os"
	"strings"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	wget("https://example.com", "here.txt")
}

const (
	bufSize = 1024 * 8
)

func wget(url, fileName string) error {
	resp, err := getResponse(url)
	if err != nil {
		return err
	}

	// если имя файла пустое то создается файл с названием из ссылки от последнего слэша до конца url
	if fileName == "" {
		urlSplit := strings.Split(url, "/")
		fileName = urlSplit[len(urlSplit)-1]
	}
	err = writeToFile(fileName, resp)
	if err != nil {
		return err
	}

	return nil
}

func getResponse(url string) (*http.Response, error) {
	tr := new(http.Transport)
	client := &http.Client{Transport: tr}

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func writeToFile(fileName string, resp *http.Response) error {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	bufferedWriter := bufio.NewWriterSize(file, bufSize)

	_, err = io.Copy(bufferedWriter, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
