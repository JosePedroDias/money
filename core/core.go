package core

import (
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
