package models

import "errors"

var (
	ErrInvalidTxDataFormat          = errors.New("invalid transaction data format")
	ErrInvalidTxHashDataFormat      = errors.New("invalid TxHash data format")
	ErrInvalidGasDataFormat         = errors.New("invalid Gas string format")
	ErrGasDataNotAnInt              = errors.New("gas data not an integer")
	ErrInvalidFeePerGasDataFormat   = errors.New("invalid FeePerGas data format")
	ErrFeePerGasNotAFloat           = errors.New("feePerGas data not a float")
	ErrInvalidSignatureStringFormat = errors.New("invalid Signature data format")
	ErrNilTx                        = errors.New("transaction was nil")
)
