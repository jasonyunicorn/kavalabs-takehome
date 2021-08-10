package utils

import (
	"sort"

	"github.com/jasonyunicorn/kavalabs/takehome/pkg/models"
)

// Mempool is a data structure used to prioritize transactions by fee
type Mempool struct {
	// maximum size of transactions to store
	maxSize int

	// slice of (prioritized) transactions
	transactions []models.Transaction
}

// NewMempool is a constructor function for instantiating a new Mempool
func NewMempool(maxSize int) *Mempool {
	return &Mempool{
		maxSize:      maxSize,
		transactions: make([]models.Transaction, 0, maxSize),
	}
}

// ListTransaction returns a slice of the Mempool's transactions
func (m *Mempool) ListTransactions() []models.Transaction {
	return m.transactions
}

// Insert appends the given transaction to the slice of transactions, sorts by fee and
// truncates the slice to the maximum mempool size if necessary
func (m *Mempool) Insert(tx *models.Transaction) error {
	if tx == nil {
		return models.ErrNilTx
	}

	newTxs := models.ByFee(append(m.transactions, *tx))
	sort.Sort(newTxs)

	if len(newTxs) > m.maxSize {
		newTxs = newTxs[0:m.maxSize]
	}

	m.transactions = newTxs

	return nil
}
