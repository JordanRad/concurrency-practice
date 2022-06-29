package main

import (
	"fmt"
	"os"

	asyncfr "github.com/JordanRad/concurrency-practice/cmd/gofi/internal/async_file_reader"
)

func main() {
	fmt.Println("Hello from Gofi")
	fileReader := asyncfr.NewAsyncFileReader()

	filePath := os.Args[1]

	file := fileReader.ReadCSV(filePath)
	fmt.Println("File rows: ", len(file.Rows))

	fileReader.ReadContent(file.Rows)
}
