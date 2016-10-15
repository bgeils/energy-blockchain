package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

var logger = shim.NewLogger("fabric-boilerplate")

//==============================================================================================================================
//	 Structure Definitions
//==============================================================================================================================
//	SimpleChaincode - A blank struct for use with Shim (An IBM Blockchain included go file used for get/put state
//					  and other IBM Blockchain functions)
//==============================================================================================================================
type SimpleChaincode struct {
}

type ECertResponse struct {
	OK string `json:"OK"`
}

type User struct {
	UserId    string `json:"userId"` //Same username as on certificate in CA
	Salt      string `json:"salt"`
	Hash      string `json:"hash"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	// TODO do you want to match any orders  /transactions to the user?
	Things       []string `json:"things"` //Array of thing IDs
	Address      string   `json:"address"`
	PhoneNumber  string   `json:"phoneNumber"`
	EmailAddress string   `json:"emailAddress"`
}

type Order struct {
	Id        string `json:"id"`
	KwhAmount string `json:"kwhAmount"`
	PriceKwh  string `json:"priceKwh"`
	TimeStart int    `json:"timeStart"`
	Duration  int    `json:"duration"`
	SellerId  string `json:"sellerId"` // link to userID
	SoldBool  bool   `json:"soldBool"`
}

type Transaction struct {
	Id      string `json:"id"`
	OrderId string `json:"orderId"`
	Seller  string `json:"sellerId"`
	Buyer   string `json:"buyerId"`
}

// FIXME to improve query efficiency you might want to create a TimeSlot struct under which you store the Orders for that timeslot

//=================================================================================================================================
//  Index collections - In order to create new IDs dynamically and in progressive sorting
//  Example:
//    signaturesAsBytes, err := stub.GetState(signaturesIndexStr)
//    if err != nil { return nil, errors.New("Failed to get Signatures Index") }
//    fmt.Println("Signature index retrieved")
//
//    // Unmarshal the signatures index
//    var signaturesIndex []string
//    json.Unmarshal(signaturesAsBytes, &signaturesIndex)
//    fmt.Println("Signature index unmarshalled")
//
//    // Create new id for the signature
//    var newSignatureId string
//    newSignatureId = "sg" + strconv.Itoa(len(signaturesIndex) + 1)
//
//    // append the new signature to the index
//    signaturesIndex = append(signaturesIndex, newSignatureId)
//    jsonAsBytes, _ := json.Marshal(signaturesIndex)
//    err = stub.PutState(signaturesIndexStr, jsonAsBytes)
//    if err != nil { return nil, errors.New("Error storing new signaturesIndex into ledger") }
//=================================================================================================================================
var usersIndexStr = "_users"
var ordersIndexStr = "_orders"
var transactionsIndexStr = "_transactions"

var indexes = []string{usersIndexStr, transactionsIndexStr, ordersIndexStr}

//==============================================================================================================================
//	Invoke - Called on chaincode invoke. Takes a function name passed and calls that function. Passes the
//  		 initial arguments passed are passed on to the called function.
//==============================================================================================================================

func (t *SimpleChaincode) Invoke(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	logger.Infof("Invoke is running " + function)

	if function == "init" {
		return t.Init(stub, "init", args)
	} else if function == "reset_indexes" {
		return t.reset_indexes(stub, args)
	} else if function == "add_user" {
		return t.add_user(stub, args)
	} else if function == "add_order" {
		return t.add_order(stub, args)
	} else if function == "add_transaction" {
		return t.add_transaction(stub, args)
	}

	return nil, errors.New("Received unknown invoke function name")
}

//=================================================================================================================================
//	Query - Called on chaincode query. Takes a function name passed and calls that function. Passes the
//  		initial arguments passed are passed on to the called function.
//=================================================================================================================================
func (t *SimpleChaincode) Query(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	logger.Infof("Query is running " + function)

	if function == "get_user" {
		return t.get_user(stub, args[1])
	} else if function == "get_order" {
		return t.get_order(stub, args)
	} else if function == "get_all_orders" {
		return t.get_all_orders(stub, args)
	} else if function == "authenticate" {
		return t.authenticate(stub, args)
	} else if function == "get_transaction" {
		return t.get_transaction(stub, args)
	} else if function == "get_all_transactions" {
		return t.get_all_transactions(stub, args)
	} 

	// TODO get orders by timeslot

	return nil, errors.New("Received unknown query function name")
}

//=================================================================================================================================
//  Main - main - Starts up the chaincode
//=================================================================================================================================

func main() {

	// LogDebug, LogInfo, LogNotice, LogWarning, LogError, LogCritical (Default: LogDebug)
	logger.SetLevel(shim.LogInfo)

	logLevel, _ := shim.LogLevel(os.Getenv("SHIM_LOGGING_LEVEL"))
	shim.SetLoggingLevel(logLevel)

	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting SimpleChaincode: %s", err)
	}
}

//==============================================================================================================================
//  Init Function - Called when the user deploys the chaincode
//==============================================================================================================================

func (t *SimpleChaincode) Init(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	return nil, nil
}

//==============================================================================================================================
//  Utility Functions
//==============================================================================================================================

// "create":  true -> create new ID, false -> append the id
func append_id(stub *shim.ChaincodeStub, indexStr string, id string, create bool) ([]byte, error) {

	indexAsBytes, err := stub.GetState(indexStr)
	if err != nil {
		return nil, errors.New("Failed to get " + indexStr)
	}

	// Unmarshal the index
	var tmpIndex []string
	json.Unmarshal(indexAsBytes, &tmpIndex)

	// Create new id
	var newId = id
	if create {
		newId += strconv.Itoa(len(tmpIndex) + 1)
	}

	// append the new id to the index
	tmpIndex = append(tmpIndex, newId)

	jsonAsBytes, _ := json.Marshal(tmpIndex)
	err = stub.PutState(indexStr, jsonAsBytes)
	if err != nil {
		return nil, errors.New("Error storing new " + indexStr + " into ledger")
	}

	return []byte(newId), nil

}

//==============================================================================================================================
//  Invoke Functions
//==============================================================================================================================
func (t *SimpleChaincode) reset_indexes(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	for _, i := range indexes {
		// Marshal the index
		var emptyIndex []string

		empty, err := json.Marshal(emptyIndex)
		if err != nil {
			return nil, errors.New("Error marshalling")
		}
		err = stub.PutState(i, empty)

		if err != nil {
			return nil, errors.New("Error deleting index")
		}
		logger.Infof("Delete with success from ledger: " + i)
	}
	return nil, nil
}

func (t *SimpleChaincode) add_user(stub *shim.ChaincodeStub, args []string) ([]byte, error) {

	//Args
	//			0				1
	//		  index		user JSON object (as string)

	id, err := append_id(stub, usersIndexStr, args[0], false)
	if err != nil {
		return nil, errors.New("Error creating new id for user " + args[0])
	}

	err = stub.PutState(string(id), []byte(args[1]))
	if err != nil {
		return nil, errors.New("Error putting user data on ledger")
	}

	return nil, nil
}

func (t *SimpleChaincode) add_order(stub *shim.ChaincodeStub, args []string) ([]byte, error) {

	// args
	// 		0			1
	//	   index	   order JSON object (as string)

	id, err := append_id(stub, ordersIndexStr, args[0], false)
	if err != nil {
		return nil, errors.New("Error creating new id for order " + args[0])
	}

	err = stub.PutState(string(id), []byte(args[1]))
	if err != nil {
		return nil, errors.New("Error putting order data on ledger")
	}

	return nil, nil

}

func (t *SimpleChaincode) add_transaction(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	// steps
	// 1.  get order
	// 2a. create new transaction
	// 2b. add transaction to ledger 
	// 3a. set flag as sold
	// 3b. put order back on ledger

	// args
	// 		0			1				2
	//	 id	   orderId	buyerId

	id, err := append_id(stub, transactionsIndexStr, args[0], false)
	if err != nil {
		return nil, errors.New("Error creating new id for transaction " + args[0])
	}

	err = stub.PutState(string(id), []byte(args[1]))
	if err != nil {
		return nil, errors.New("Error putting transaction data on ledger")
	}

	return nil, nil

}

//==============================================================================================================================
//		Query Functions
//==============================================================================================================================

func (t *SimpleChaincode) get_user(stub *shim.ChaincodeStub, userID string) ([]byte, error) {

	bytes, err := stub.GetState(userID)

	if err != nil {
		return nil, errors.New("Could not retrieve information for this user")
	}

	return bytes, nil

}

func (t *SimpleChaincode) get_order(stub *shim.ChaincodeStub, args []string) ([]byte, error) {

	//Args
	//			0
	//		orderID

	bytes, err := stub.GetState(args[0])

	if err != nil {
		return nil, errors.New("Error getting order from ledger")
	}

	return bytes, nil

}

func (t *SimpleChaincode) get_transaction(stub *shim.ChaincodeStub, args []string) ([]byte, error) {

	//Args
	//			0
	//		orderID

	bytes, err := stub.GetState(args[0])

	if err != nil {
		return nil, errors.New("Error getting order from ledger")
	}

	return bytes, nil

}

// get all orders
func (t *SimpleChaincode) get_all_orders(stub *shim.ChaincodeStub, args []string) ([]byte, error) {

	indexAsBytes, err := stub.GetState(ordersIndexStr)
	if err != nil {
		return nil, errors.New("Failed to get " + ordersIndexStr)
	}
	// TODO replace thing with order / transaction
	// Unmarshal the index
	var ordersIndex []string
	json.Unmarshal(indexAsBytes, &ordersIndex)

	var orders []Order
	for _, order := range ordersIndex {

		bytes, err := stub.GetState(order)
		if err != nil {
			return nil, errors.New("Unable to get order with ID: " + order)
		}

		var t Order
		json.Unmarshal(bytes, &t)
		orders = append(orders, t)
	}

	ordersAsJsonBytes, _ := json.Marshal(orders)
	if err != nil {
		return nil, errors.New("Could not convert orders to JSON ")
	}

	return ordersAsJsonBytes, nil
}

//get all transactions
func (t *SimpleChaincode) get_all_transactions(stub *shim.ChaincodeStub, args []string) ([]byte, error) {

	indexAsBytes, err := stub.GetState(transactionsIndexStr)
	if err != nil {
		return nil, errors.New("Failed to get " + transactionsIndexStr)
	}
	// TODO replace thing with order / transaction
	// Unmarshal the index
	var transactionsIndex []string
	json.Unmarshal(indexAsBytes, &transactionsIndex)

	var transactions []Transaction
	for _, transaction := range transactionsIndex {

		bytes, err := stub.GetState(transaction)
		if err != nil {
			return nil, errors.New("Unable to get order with ID: " + transaction)
		}

		var t Transaction
		json.Unmarshal(bytes, &t)
		transactions = append(transactions, t)
	}

	transactionsAsJsonBytes, _ := json.Marshal(transactions)
	if err != nil {
		return nil, errors.New("Could not convert transactions to JSON ")
	}

	return transactionsAsJsonBytes, nil
}

func (t *SimpleChaincode) authenticate(stub *shim.ChaincodeStub, args []string) ([]byte, error) {

	// Args
	//	0		1
	//	userId	password

	var u User

	username := args[0]

	user, err := t.get_user(stub, username)

	// If user can not be found in ledgerstore, return authenticated false
	if err != nil {
		return []byte(`{ "authenticated": false }`), nil
	}

	//Check if the user is an employee, if not return error message
	err = json.Unmarshal(user, &u)
	if err != nil {
		return []byte(`{ "authenticated": false}`), nil
	}

	// Marshal the user object
	userAsBytes, err := json.Marshal(u)
	if err != nil {
		return []byte(`{ "authenticated": false}`), nil
	}

	// Return authenticated true, and include the user object
	str := `{ "authenticated": true, "user": ` + string(userAsBytes) + `  }`

	return []byte(str), nil
}
