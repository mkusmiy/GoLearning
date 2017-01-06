package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func CountLines(r io.Reader) (int, error) {
	sc := bufio.NewScanner(r)
	var lines int
	for sc.Scan() {
		lines++
	}
	return lines, sc.Err()
}

func CountFile(path string) int {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	lines, err := CountLines(f)
	if err != nil {
		log.Fatal(err)
	}
	_, filename := filepath.Split(path)
	fmt.Printf("%-40s%6d\n", filename, lines)
	return lines
}

func CountDir(dir string) {
	d, err := os.Open(dir)
	if err != nil {
		log.Fatal(err)
	}
	defer d.Close()
	entries, err := d.Readdir(-1)
	if err != nil {
		log.Fatal(err)
	}
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		if strings.HasSuffix(e.Name(), ".txt") {
			CountFile(filepath.Join(dir, e.Name()))
		}
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	message := "Hello World!!!"
	w.Write([]byte(message))
}

var Path string

func main() {
	args := os.Args[1:]
	for _, arg := range args {
		CountDir(arg)
	}

	http.HandleFunc("/hello", hello)

	Path = os.Args[1]
	_ = Path

	http.HandleFunc("/books/", func(w http.ResponseWriter, r *http.Request) {
		book := strings.SplitN(r.URL.Path, "/", 3)[2]

		data := analyze(book)

		w.Header().Set("Content-Type", "application/json' charset=utf-8")
		json.NewEncoder(w).Encode(data)
	})

	http.ListenAndServe(":8080", nil)
}

func analyze(book string) (b bookName) {

	length := CountFile(filepath.Join(Path, book))

	b = bookName{
		Name:   book,
		Length: length,
	}

	return
}

type bookName struct {
	Name   string `json:"title"`
	Length int    `json:"lines"`
}
