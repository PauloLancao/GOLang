package storage

import (
	"logging"
	"response"
)

// Request struct
type Request struct {
	Cmd      string
	Key      string
	Value    *StoreValue
	Response chan<- Result
}

// Result struct
type Result struct {
	Value *StoreValue
	Err   error
}

// StoreValue struct
type StoreValue struct {
	CType  string
	CValue []byte
}

// Store func
type Store struct {
	kvs      map[string]*StoreValue
	Requests chan Request
}

// Cmd commands for store
const (
	CmdGetRecord    string = "getrecord"
	CmdAddRecord    string = "addrecord"
	CmdUpdateRecord string = "updaterecord"
	CmdDeleteRecord string = "deleterecord"
)

// Start func
func Start() Store {
	store := Store{make(map[string]*StoreValue), make(chan Request)}
	go store.servingRequests()
	return store
}

// Close func
func (s *Store) Close() {
	close(s.Requests)
}

func (s *Store) servingRequests() {
	for req := range s.Requests {
		res, err := ExecuteCommand(s.kvs, req.Key, req.Value, req.Cmd)
		go deliver(req, Result{res, err})
	}

	logging.Msg(logging.UUID(),
		"KVS",
		"ServingRequests",
		"Leaving Serving Requests...")
}

// Deliver result to request channel
func deliver(request Request, result Result) {
	request.Response <- result
	close(request.Response)
}

// ExecuteCommand func
func ExecuteCommand(kvs map[string]*StoreValue, k string, v *StoreValue, cmd string) (*StoreValue, error) {

	if cmd == CmdAddRecord || cmd == CmdUpdateRecord {
		return addOrUpdateRecord(kvs, k, v)
	} else if cmd == CmdGetRecord {
		return getRecord(kvs, k, v)
	} else if cmd == CmdDeleteRecord {
		return deleteRecord(kvs, k, v)
	}

	return nil, response.New(response.ErrMissingCommand, "Store::ExecuteCommand")
}

// GetRecord from kvs
func getRecord(kvs map[string]*StoreValue, k string, v *StoreValue) (*StoreValue, error) {
	value, found := kvs[k]

	if !found {
		return nil, response.New(response.ErrNotFound, "KVS::GetRecord")
	}

	return value, nil
}

// addOrUpdateRecord to kvs
func addOrUpdateRecord(kvs map[string]*StoreValue, k string, v *StoreValue) (*StoreValue, error) {
	kvs[k] = v

	return v, nil
}

// DeleteRecord from kvs
func deleteRecord(kvs map[string]*StoreValue, k string, v *StoreValue) (*StoreValue, error) {
	todelete, found := kvs[k]

	if !found {
		return nil, response.New(response.ErrUnknown, "KVS::DeleteRecord")
	}

	delete(kvs, k)

	return todelete, nil
}
