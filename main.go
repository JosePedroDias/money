package main

import (
	"encoding/json"
	"fmt"
	//"github.com/josepedrodias/money/core"
	provider "github.com/josepedrodias/money/provider/montepio"
)

func main() {
	filePath := "/home/jdias/Downloads/dinheiro/ExtractosDO_019100068022_20140901_20140930.XLS"
	movimentos, err := provider.ImportCSV(filePath)
	if err != nil {
		panic(err)
	}

	//mov1 := movimentos.Front().Value.(core.Movimento)

	jsonBytes, err := json.Marshal(movimentos)
	fmt.Printf("%s\n", jsonBytes)

	//fmt.Printf("%+v\n", mov1) // %v %+v %#v
}
