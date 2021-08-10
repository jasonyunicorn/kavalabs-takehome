package utils

import (
	"testing"

	"github.com/jasonyunicorn/kavalabs/takehome/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestMempool_NewMempool(t *testing.T) {
	testMaxMempoolSize := 10

	t.Run("NewMempool returns a non nil pointer to a new Mempool", func(t *testing.T) {
		m := NewMempool(testMaxMempoolSize)

		assert.NotNil(t, m)
		assert.Equal(t, m.maxSize, testMaxMempoolSize)
		assert.Equal(t, cap(m.transactions), testMaxMempoolSize)
		assert.Equal(t, len(m.transactions), 0)
	})
}

func TestMempool_ListTransactions(t *testing.T) {
	testMaxMempoolSize := 3
	testTx := &models.Transaction{TxHash: "txHash", Gas: 1234, FeePerGas: 0.1234, Signature: "signature"}

	tests := []struct {
		name    string
		setupFn func() *Mempool
		testFn  func(*testing.T, []models.Transaction)
	}{
		{
			name:    "new mempool returns empty slice",
			setupFn: func() *Mempool { return NewMempool(testMaxMempoolSize) },
			testFn: func(t *testing.T, txs []models.Transaction) {
				assert.NotNil(t, txs)
				assert.Empty(t, txs)
			},
		},
		{
			name: "pre-populated mempool returns non empty slice",
			setupFn: func() *Mempool {
				m := NewMempool(testMaxMempoolSize)
				m.Insert(testTx)

				return m
			},
			testFn: func(t *testing.T, txs []models.Transaction) {
				assert.NotNil(t, txs)
				assert.NotEmpty(t, txs)
				assert.Len(t, txs, 1)
				assert.Equal(t, testTx, &txs[0])
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := tt.setupFn()
			txs := m.ListTransactions()
			tt.testFn(t, txs)
		})
	}
}

func TestMempool_Insert(t *testing.T) {
	testMaxMempoolSize := 3
	testTx1 := &models.Transaction{TxHash: "txHash1", Gas: 1, FeePerGas: .1, Signature: "signature1"}
	testTx2 := &models.Transaction{TxHash: "txHash2", Gas: 2, FeePerGas: .2, Signature: "signature2"}
	testTx3 := &models.Transaction{TxHash: "txHash3", Gas: 3, FeePerGas: .3, Signature: "signature3"}
	testTx4 := &models.Transaction{TxHash: "txHash4", Gas: 4, FeePerGas: .4, Signature: "signature4"}

	tests := []struct {
		name   string
		testFn func(*testing.T, *Mempool)
	}{
		{
			name: "insert nil transaction returns error",
			testFn: func(t *testing.T, m *Mempool) {
				gotErr := m.Insert(nil)

				assert.Error(t, gotErr)
			},
		},
		{
			name: "insert one transaction into new mempool",
			testFn: func(t *testing.T, m *Mempool) {
				gotErr := m.Insert(testTx1)

				assert.Nil(t, gotErr)
				assert.Len(t, m.transactions, 1)
				assert.Equal(t, testTx1, &m.transactions[0])
				assert.Equal(t, "txHash1", m.transactions[0].TxHash)
			},
		},
		{
			name: "insert multiple transactions, less than max mempool size",
			testFn: func(t *testing.T, m *Mempool) {
				gotErr1 := m.Insert(testTx1)
				gotErr2 := m.Insert(testTx2)

				assert.Nil(t, gotErr1)
				assert.Nil(t, gotErr2)
				assert.Len(t, m.transactions, 2)
				assert.Equal(t, "txHash2", m.transactions[0].TxHash)
				assert.Equal(t, "txHash1", m.transactions[1].TxHash)
			},
		},
		{
			name: "insert multiple transactions, greater than max mempool size",
			testFn: func(t *testing.T, m *Mempool) {
				gotErr1 := m.Insert(testTx1)
				gotErr2 := m.Insert(testTx2)
				gotErr3 := m.Insert(testTx3)
				gotErr4 := m.Insert(testTx4)

				assert.Nil(t, gotErr1)
				assert.Nil(t, gotErr2)
				assert.Nil(t, gotErr3)
				assert.Nil(t, gotErr4)
				assert.Len(t, m.transactions, 3)
				assert.Equal(t, "txHash4", m.transactions[0].TxHash)
				assert.Equal(t, "txHash3", m.transactions[1].TxHash)
				assert.Equal(t, "txHash2", m.transactions[2].TxHash)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewMempool(testMaxMempoolSize)
			tt.testFn(t, m)
		})
	}
}
