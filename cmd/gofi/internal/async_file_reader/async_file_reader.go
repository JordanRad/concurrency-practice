package asyncfilereader

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type AsyncDataCollector struct {
	mu           sync.Mutex
	transactions []string
	successCount int
}
type AsyncFileReader struct {
	data *AsyncDataCollector
}

func newCollector() *AsyncDataCollector {
	return &AsyncDataCollector{
		transactions: make([]string, 0),
		successCount: 0,
	}
}
func NewAsyncFileReader() *AsyncFileReader {
	collector := newCollector()
	return &AsyncFileReader{collector}
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

func (fr *AsyncFileReader) insert(transaction *Transaction) error {
	rn := rand.Intn(10)
	time.Sleep(time.Millisecond * 2)
	fr.data.mu.Lock()
	defer fr.data.mu.Unlock()

	if rn%4 == 0 {

		fr.data.transactions = append(fr.data.transactions, transaction.ID)
		return errors.New("error inserting a transaction")
	}

	fr.data.successCount++
	return nil
}

func getRoutinesCount(fileRowsNumber uint) int {
	switch {
	case fileRowsNumber >= 100_000:
		return 10
	case fileRowsNumber >= 50_000:
		return 5
	default:
		return 1
	}
}
func (fr *AsyncFileReader) InsertIntoDB(fileContent []string) {
	start := time.Now()
	fmt.Println("Start: ", start.Format("15:04:05"))

	for _, row := range fileContent {
		t := toTransactionEntity(row)
		go fr.insert(t)
	}

	end := time.Now()
	fmt.Println("End: ", end.Format("15:04:05"))

	total := (end.UnixMilli() - start.UnixMilli()) / 1000
	fmt.Printf("Time Elapsed: %v \n", total)
	fmt.Printf("Transaction success number: %v \n", fr.data.successCount)
	fmt.Printf("Transaction failed number: %v", len(fr.data.transactions))
}
