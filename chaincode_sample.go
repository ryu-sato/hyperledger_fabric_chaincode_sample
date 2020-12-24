package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

var logger = shim.NewLogger("hyperledger_fabric_chaincode_sample")

type SmartContract struct {
}

// chaincode の初期化
func (t *SmartContract) Init(stub shim.ChaincodeStubInterface) peer.Response {
	logger.Info("Chaincode initialized.")
	return shim.Success(nil) // 何もしない
}

// トランザクションの実行
// args[0]: トランザクション名
// args[1..*]: コマンドに渡す引数
func (t *SmartContract) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	cmdMap := map[string]interface{}{
		"helloChaincode": HelloChaincode,
	}
	cmdName, cmdArgs := stub.GetFunctionAndParameters()

	if cmdInterface, exist := cmdMap[cmdName]; exist {
		cmd := reflect.ValueOf(cmdInterface)
		out := cmd.Call([]reflect.Value{reflect.ValueOf(stub), reflect.ValueOf(cmdArgs)})
		if ret, ok := out[0].Interface().(peer.Response); ok {
			return ret
		}
		return shim.Error(fmt.Sprintf("Invoke: Wrong response type (given %T, expected peer.Response)", out[0].Interface()))
	}

	keys := reflect.ValueOf(cmdMap).MapKeys()
	return shim.Error(fmt.Sprintf("Invoke: Invalid command (given %s, expexted %v", cmdName, keys))
}

func HelloChaincode(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	bytes, _ := json.Marshal("hello chaincode!");
	return shim.Success(bytes);
}

func main() {
	logLevel := shim.LogInfo
	if os.Getenv("SHIM_LOGGING_LEVEL") != "" {
		logLevel, _ = shim.LogLevel(os.Getenv("SHIM_LOGGING_LEVEL"))
	}
	logger.SetLevel(logLevel)
	shim.SetLoggingLevel(logLevel)

	err := shim.Start(new(SmartContract))
	if err != nil {
		logger.Errorf("Error hello chaincode: %s\n", err)
	}
}

