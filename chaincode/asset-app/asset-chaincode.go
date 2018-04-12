// SPDX-License-Identifier: Apache-2.0

/*
  Sample Chaincode based on Demonstrated Scenario

 This code is based on code written by the Hyperledger Fabric community.
  Original code can be found here: https://github.com/hyperledger/fabric-samples/blob/release/chaincode/fabcar/fabcar.go
*/

package main

/* Imports
* 4 utility libraries for handling bytes, reading and writing JSON,
formatting, and string manipulation
* 2 specific Hyperledger Fabric specific libraries for Smart Contracts
*/
import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

/* Define asset structure, with 4 properties.
Structure tags are used by encoding/json library
*/
type Asset struct {
	Item     string `json:"item"`
	Holder   string `json:"holder"`
	Location string `json:"location"`
	Cost     string `json:"cost"`
}

/*
 * The Init method *
 called when the Smart Contract "asset-chaincode" is instantiated by the network
 * Best practice is to have any Ledger initialization in separate function
 -- see initLedger()
*/
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
 * The Invoke method *
 called when an application requests to run the Smart Contract "asset-chaincode"
 The app also specifies the specific smart contract function to call with args
*/
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger
	if function == "queryAsset" {
		return s.queryAsset(APIstub, args)
	} else if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "recordAsset" {
		return s.recordAsset(APIstub, args)
	} else if function == "queryAllAsset" {
		return s.queryAllAsset(APIstub)
	} else if function == "changeAssetHolder" {
		return s.changeAssetHolder(APIstub, args)
	}
	return shim.Error("Invalid Smart Contract function name.")
}

/*
 * The queryAsset method *
Used to view the records of one particular asset
It takes one argument -- the key for the asset in question
*/
func (s *SmartContract) queryAsset(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	assetAsBytes, _ := APIstub.GetState(args[0])
	if assetAsBytes == nil {
		return shim.Error("Could not locate asset")
	}
	return shim.Success(assetAsBytes)
}

/*
 * The initLedger method *
Will add test data (10 asset catches)to our network
*/
func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	asset := []Asset{
		Asset{Cost: "$85,000", Location: "Washington, D.C.", Item: "Porche 911", Holder: "Roger Raceman"},
		Asset{Cost: "$6,000,000", Location: "San Diego, CA", Item: "Leer Jet G5", Holder: "Jerry Pits"},
		Asset{Cost: "$575,000", Location: "Bali, Indonesia" Item: "Beach House", Holder: "Rick Jortz"},
		Asset{Cost: "$85,000", Location: "Washington, D.C.", Item: "Porche 911", Holder: "Roger Raceman"},
		Asset{Cost: "$3,750,000", Location: "Park City, UT", Item: "Park City Mountainside Home", Holder: "Sloan Slacks"},
		Asset{Cost: "$33,333", Location: "New Orleans,LA", Item: "Dinasaur Egg", Holder: "Madam Mern"},
		Asset{Cost: "$2,000,000", Location: "Elon Musk's Warehouse", Item: "JetPack 3000", Holder: "Perry Pants"},
		Asset{Cost: "$50", Location: "Blockbuster in Omaha, NE", Item: "Unopened Heavyweights DVD", Holder: "Fiona Fiddle"},
		Asset{Cost: "$75,000", Location: "New York, NY", Item: "Rare Red Ruby Ring", Holder: "Brick Block"},
		Asset{Cost: "$25,00", Location: "Zurich, Switerland", Item: "Gold Bars", Holder: "Eddie Eth"},
		Asset{Cost: "$65,000,000", Location: "Montengero", Item: "The Dark Knight Yacht", Holder: "Tim Cook"},
	}

	i := 0
	for i < len(asset) {
		fmt.Println("i is ", i)
		assetAsBytes, _ := json.Marshal(asset[i])
		APIstub.PutState(strconv.Itoa(i+1), assetAsBytes)
		fmt.Println("Added", asset[i])
		i = i + 1
	}

	return shim.Success(nil)
}

/*
 * The recordAsset method *
Fisherman like Sarah would use to record each of her asset catches.
This method takes in five arguments (attributes to be saved in the ledger).
*/
func (s *SmartContract) recordAsset(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	var asset = Asset{Cost: args[1], Location: args[2], Item: args[3], Holder: args[4]}

	assetAsBytes, _ := json.Marshal(asset)
	err := APIstub.PutState(args[0], assetAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to record asset: %s", args[0]))
	}

	return shim.Success(nil)
}

/*
 * The queryAllasset method *
allows for assessing all the records added to the ledger(all asset catches)
This method does not take any arguments. Returns JSON string containing results.
*/
func (s *SmartContract) queryAllAsset(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "0"
	endKey := "999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add comma before array members,suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllAsset:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

/*
 * The changeAssetHolder method *
The data in the world state can be updated with who has possession.
This function takes in 2 arguments, asset id and new holder name.
*/
func (s *SmartContract) changeAssetHolder(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	assetAsBytes, _ := APIstub.GetState(args[0])
	if assetAsBytes == nil {
		return shim.Error("Could not locate asset")
	}
	asset := Asset{}

	json.Unmarshal(assetAsBytes, &asset)
	// Normally check that the specified argument is a valid holder of asset
	// we are skipping this check for this example
	asset.Holder = args[1]

	assetAsBytes, _ = json.Marshal(asset)
	err := APIstub.PutState(args[0], assetAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to change asset holder: %s", args[0]))
	}

	return shim.Success(nil)
}

/*
 * main function *
calls the Start function
The main function starts the chaincode in the container during instantiation.
*/
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
