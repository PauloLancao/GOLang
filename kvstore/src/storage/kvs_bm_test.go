package storage

import (
	"logging"
	"strconv"
	"testing"
)

func BenchmarkAddRecordParallel(b *testing.B) {
	// arrange
	logging.CreateLogger([]string{})
	defer logging.Close()

	store := Start()
	defer store.Close()

	// act
	for i := 0; i < b.N; i++ {
		resChan := make(chan Result)
		req := Request{
			CmdAddRecord,
			strconv.Itoa(i),
			&StoreValue{CType: "application/json", CValue: []byte(string(i))},
			resChan}
		store.Requests <- req
		<-resChan
	}
}
