package asyncfilereader

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type AsyncDataCollector struct {
	mu           sync.Mutex
	transactions []Transaction
}
type AsyncFileReader struct {
	data *AsyncDataCollector
}

func NewAsyncFileReader() *AsyncFileReader {
	return &AsyncFileReader{&AsyncDataCollector{}}
}

func (fr *AsyncFileReader) ReadCSV(path string) *CSVFile {
	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Errorf("file cannot be read: %w", err)
	}
	fmt.Printf("Filesize: %v bytes \n", len(file))
	reader := bytes.NewReader(file)
	fileScanner := bufio.NewScanner(reader)
	fileScanner.Split(bufio.ScanLines)

	rows := make([]string, 0)
	for fileScanner.Scan() {
		row := fileScanner.Text()
		rows = append(rows, row)
	}

	csvFile := &CSVFile{
		Header: rows[0],
		Rows:   rows[1:],
	}
	return csvFile
}

func toTransactionEntity(row string) *Transaction {
	rowEntries := strings.Split(row, ",")
	amount, _ := strconv.ParseFloat(rowEntries[3], 32)

	return &Transaction{
		ID:        rowEntries[0],
		From:      rowEntries[1],
		To:        rowEntries[2],
		Amount:    float32(amount),
		CreatedAt: rowEntries[4],
	}
}

func (fr *AsyncFileReader) ReadContent(fileContent []string) {
	start := time.Now().UnixMilli()
	//fmt.Println("Start: ", start.Format("15:04:05"))
	ids := make([]*Transaction, 0, 0)
	for _, row := range fileContent {
		tr := toTransactionEntity(row)
		ids = append(ids, tr)

	}
	fmt.Println(ids[0].Amount)
	end := time.Now().UnixMilli()

	//fmt.Println("End: ", end.Format("15:04:05"))
	fmt.Print("Time Elapsed: ", (end - start))
}
