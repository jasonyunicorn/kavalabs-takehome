package models

import (
	"fmt"
	"strconv"
	"strings"
)

// Transaction is a struct which holds information about a transaction
type Transaction struct {
	TxHash    string
	Gas       int
	FeePerGas float64
	Signature string
}

// ToString converts the transaction into a formatted string
func (tx *Transaction) ToString() string {
	return fmt.Sprintf("TxHash=%s Gas=%d FeePerGas=%f Signature=%s",
		tx.TxHash,
		tx.Gas,
		tx.FeePerGas,
		tx.Signature,
	)
}

// ParseTransaction parses a string of text following a specific format into a transaction
func ParseTransaction(txString string) (*Transaction, error) {
	data := strings.Split(txString, " ")

	if len(data) != 4 {
		return nil, ErrInvalidTxDataFormat
	}

	// parse TxHash
	txHashData := strings.Split(data[0], "=")
	if len(txHashData) != 2 {
		return nil, ErrInvalidTxHashDataFormat
	}

	// parse Gas
	gasData := strings.Split(data[1], "=")
	if len(gasData) != 2 {
		return nil, ErrInvalidGasDataFormat
	}

	gas, err := strconv.Atoi(gasData[1])
	if err != nil {
		return nil, ErrGasDataNotAnInt
	}

	// parse FeePerGas
	feePerGasData := strings.Split(data[2], "=")
	if len(feePerGasData) != 2 {
		return nil, ErrInvalidFeePerGasDataFormat
	}

	feePerGas, err := strconv.ParseFloat(feePerGasData[1], 64)
	if err != nil {
		return nil, ErrFeePerGasNotAFloat
	}

	// parse Signature
	signatureData := strings.Split(data[3], "=")
	if len(signatureData) != 2 {
		return nil, ErrInvalidSignatureStringFormat
	}

	return &Transaction{
		TxHash:    txHashData[1],
		Gas:       gas,
		FeePerGas: feePerGas,
		Signature: signatureData[1],
	}, nil
}

// ByFee implements the sort.Interface based on the Transaction fee paid
type ByFee []Transaction

// Len implements the Len() function of th sort.Interface for ByFee
func (a ByFee) Len() int {
	return len(a)
}

// Len implements the Less() function of th sort.Interface for ByFee
func (a ByFee) Less(i, j int) bool {
	return float64(a[i].Gas)*a[i].FeePerGas > float64(a[j].Gas)*a[j].FeePerGas
}

// Len implements the Swap() function of th sort.Interface for ByFee
func (a ByFee) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
