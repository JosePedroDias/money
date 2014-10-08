package main

import (
	"fmt"
	"github.com/josepedrodias/money/core"
	//provider "github.com/josepedrodias/money/provider/montepio"
)

func main() {
	/*// import from montepio CSV
	movimentos, err := provider.ImportCSV("/home/jdias/Downloads/dinheiro/ExtractosDO_019100068022_20140901_20140930.XLS")
	if err != nil {
		panic(err)
	}*/

	/*// save to JSON
	err = core.MovimentosToJSON(movimentos, "/tmp/movimentos1.json")
	if err != nil {
		panic(err)
	}*/

	//load from JSON
	movimentos, err := core.MovimentosFromJSON("/tmp/movimentos1.json")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", movimentos[0]) // %v %+v %#v
}
