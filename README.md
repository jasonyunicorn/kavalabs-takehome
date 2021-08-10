## Mempool Takehome Assignment

This is my solution for my takehome assignment for Kava Labs.  It is a simple implementation of a [Mempool](https://www.blocknative.com/blog/mempool-intro).

### Dependencies

go v1.16+

### Installation

To install this application you'll need to clone the repository and run `go mod tidy`

### Build and Run

To compile and run the application, you'll need to run the following:

```bash
cd cmd/prioritize
go build .
./prioritize
```

**NOTE: The output file (priortize-transactions.txt) must be deleted first before you can successfully run the application.  Otherwise, the program will panic.**
