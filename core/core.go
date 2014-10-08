package core

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

type Movimento struct {
	DataMovim   time.Time `json:"dataMovim"`
	DataOper    time.Time `json:"dataOper"`
	Descricao   string    `json:"descricao"`
	Importancia int64     `json:"importancia"`
	MoedaI      string    `json:"moedaI"`
	SaldoContab int64     `json:"saldoContab"`
	MoedaSC     string    `json:"moedaSC"`
}

/**
money can have thousand dots and cents separated by ,
"-20.4" -> -2040
*/
func Money2Int64(s string) (int64, error) {
	s = strings.Replace(s, ".", "", -1) // remove thousand separators

	i := strings.LastIndex(s, ",") // find cents separator

	if i != -1 { // it exists - wrap around it, adding remaining decimals if lacking
		l := len(s)
		s = s[:i] + s[i+1:] + strings.Repeat("0", 2-(l-i-1))
	}

	return strconv.ParseInt(s, 10, 64)
}

// REFERENCE TIME Mon Jan 2 15:04:05 -0700 MST 2006
const dateLayout string = "2006-01-02"

func YMD2Time(s string) (time.Time, error) {
	return time.Parse(dateLayout, s)
}

func MovimentosFromJSON(jsonFilePath string) ([]Movimento, error) {
	//var movimentos []Movimento
	movimentos := make([]Movimento, 0)
	f, err := os.Open(jsonFilePath)
	if err != nil {
		return movimentos, err
	}
	jsonS, err := ioutil.ReadAll(f)
	if err != nil {
		return movimentos, err
	}
	err = json.Unmarshal(jsonS, &movimentos)
	if err != nil {
		return movimentos, err
	}
	return movimentos, nil
}

func MovimentosToJSON(movimentos []Movimento, jsonFilePath string) error {
	jsonBytes, err := json.Marshal(movimentos)
	// jsonBytes, err := json.MarshalIndent(movimentos, "", "\t") // prefix, indent
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(jsonFilePath, jsonBytes, 0644)
	if err != nil {
		return err
	}
	return nil
}
