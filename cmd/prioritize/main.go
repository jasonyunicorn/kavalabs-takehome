package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/jasonyunicorn/kavalabs/takehome/pkg/models"
	"github.com/jasonyunicorn/kavalabs/takehome/pkg/utils"
)

const MaxMempoolSize int = 5000

var (
	inputFilename  = "./transactions.txt"
	outputFilename = "./prioritized-transactions.txt"
)

func main() {
	txCh := make(chan *models.Transaction)

	go func() {
		defer close(txCh)

		readTransactions(inputFilename, txCh)
	}()

	txs := handleTransactions(txCh)

	writeTransactions(outputFilename, txs)
}

// readTransactions opens the given file name (relative path), parses each line as a transaction
// and sends each transaction to a non-buffered channel
func readTransactions(filename string, txCh chan *models.Transaction) {
	f, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("%v", err)
		panic("unable to open file")
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		txStr := sc.Text()

		tx, err := models.ParseTransaction(fmt.Sprintf("%s\n", txStr))
		if err != nil {
			log.Printf("encountered an error while parsing transaction string %s, err = %v", txStr, err)
			continue
		}

		txCh <- tx
	}
}

// handleTransactions inserts each transaction sent over the transactions channel to be prioritzed
// then returns a slice of transactions once the transactions channel is closed
func handleTransactions(txCh chan *models.Transaction) []models.Transaction {
	mp := utils.NewMempool(MaxMempoolSize)

	for tx := range txCh {
		mp.Insert(tx)
	}

	return mp.ListTransactions()
}

// writeTransactions opens the given file name (relative path) and writes each transaction to the file
// NOTE: the file must not already exist, otherwise the program panics
func writeTransactions(filename string, txs []models.Transaction) {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_EXCL|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("%v", err)
		panic("unable to create file")
	}
	defer f.Close()

	wr := bufio.NewWriter(f)
	defer wr.Flush()

	for _, tx := range txs {
		txStr := tx.ToString()

		_, err = wr.WriteString(txStr)
		if err != nil {
			log.Printf("encountered an error while writing transaction: %s", txStr)
			continue
		}
	}
}
