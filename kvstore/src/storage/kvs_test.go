package storage

import (
	"fmt"
	"logging"
	"strconv"
	"sync"
	"testing"
)

func TestGetRecord(t *testing.T) {
	// arrange
	logging.CreateLogger([]string{})
	defer logging.Close()

	store := Start()
	defer store.Close()
	dummyData(store.kvs)

	// act
	record, err := ExecuteCommand(store.kvs, "k1", nil, CmdGetRecord)

	// assert
	if err != nil {
		t.Errorf("GetRecord is nil error %v", err)
	}

	if record == nil {
		t.Errorf("Expected record on kvs with key %s", "k1")
	}
}

func TestAddRecordsParallel(t *testing.T) {
	wg := sync.WaitGroup{}
	iterations := 200
	wg.Add(iterations)

	logging.CreateLogger([]string{})
	defer logging.Close()

	store := Start()
	defer store.Close()

	for i := 0; i < iterations; i++ {
		go func(idx int) {
			resChan := make(chan Result)
			req := Request{
				CmdAddRecord, strconv.Itoa(idx),
				&StoreValue{CType: "application/json", CValue: []byte(string(idx))},
				resChan}
			store.Requests <- req
			<-resChan
			wg.Done()
		}(i)
	}
	wg.Wait()

	records := len(store.kvs)

	if records != 200 {
		t.Errorf("Expected initial kvs %d keys but got %d", iterations, records)
	}
}

func TestAddRecordsParallel2(t *testing.T) {
	// arrange
	wg := sync.WaitGroup{}
	wg.Add(2)

	logging.CreateLogger([]string{})
	defer logging.Close()

	store := Start()
	defer store.Close()

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			resChan := make(chan Result)
			req := Request{
				CmdAddRecord,
				strconv.Itoa(i),
				&StoreValue{CType: "application/json", CValue: []byte(string(i))},
				resChan}
			store.Requests <- req
			<-resChan
		}
	}()
	go func() {
		defer wg.Done()
		for i := 10; i < 30; i++ {
			resChan := make(chan Result)
			req := Request{
				CmdAddRecord,
				strconv.Itoa(i),
				&StoreValue{CType: "application/json", CValue: []byte(string(i))},
				resChan}
			store.Requests <- req
			<-resChan
		}
	}()
	wg.Wait()

	records := len(store.kvs)

	if records != 30 {
		t.Errorf("Expected initial kvs 30 keys but got %d", records)
	}
}

func TestDeleteRecordsParallel(t *testing.T) {
	// arrange
	wg := sync.WaitGroup{}
	wg.Add(2)

	logging.CreateLogger([]string{})
	defer logging.Close()

	store := Start()
	defer store.Close()

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			resChan := make(chan Result)
			req := Request{
				CmdAddRecord,
				strconv.Itoa(i),
				&StoreValue{CType: "application/json", CValue: []byte(string(i))},
				resChan}
			store.Requests <- req
			<-resChan
		}
	}()
	go func() {
		defer wg.Done()
		for i := 10; i < 30; i++ {
			resChan := make(chan Result)
			req := Request{
				CmdAddRecord,
				strconv.Itoa(i),
				&StoreValue{CType: "application/json", CValue: []byte(string(i))},
				resChan}
			store.Requests <- req
			<-resChan
		}
	}()
	wg.Wait()

	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			resChan := make(chan Result)
			req := Request{
				CmdDeleteRecord,
				strconv.Itoa(i),
				&StoreValue{CType: "application/json", CValue: []byte(string(i))},
				resChan}
			store.Requests <- req
			<-resChan
		}
	}()
	go func() {
		defer wg.Done()
		for i := 10; i < 30; i++ {
			resChan := make(chan Result)
			req := Request{
				CmdDeleteRecord,
				strconv.Itoa(i),
				&StoreValue{CType: "application/json", CValue: []byte(string(i))},
				resChan}
			store.Requests <- req
			<-resChan
		}
	}()
	wg.Wait()
}

// dummydata injection
func dummyData(kvs map[string]*StoreValue) {
	kvs["k1"] = &StoreValue{CType: "application/json", CValue: []byte("k1")}
	kvs["k2"] = &StoreValue{CType: "application/json", CValue: []byte("999")}
	kvs["k3"] = &StoreValue{CType: "application/json", CValue: []byte(fmt.Sprintf("%v", []string{"t1", "t2", "t3"}))}
	kvs["k4"] = &StoreValue{CType: "application/json", CValue: []byte(fmt.Sprintf("%v", struct{}{}))}
	kvs["k5"] = &StoreValue{CType: "application/json", CValue: []byte(fmt.Sprintf("%v", map[int]interface{}{1: "test1", 2: "test2"}))}
}
