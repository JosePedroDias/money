package montepio

import (
	"container/list"
	"encoding/csv"
	iconv "github.com/djimenez/iconv-go"
	"github.com/josepedrodias/money/core"
	"os"
	"strings"
)

func trim(s string) string {
	return strings.Trim(s, " \t")
}

func ImportCSV(filePath string) (*list.List, error) {
	movimentos := list.New()

	// open file
	f, err := os.Open(filePath)
	if err != nil {
		return movimentos, err
	}

	// convert encoding from iso-8859-1 to utf-8
	r0, err := iconv.NewReader(f, "iso-8859-1", "utf-8")
	if err != nil {
		return movimentos, err
	}

	// configure CSV reader
	r := csv.NewReader(r0)
	r.Comma = '\t'

	// reads first header row
	_, err = r.Read()
	if err != nil {
		return movimentos, err
	}
	/*
		DATA MOVIMENTO
		DATA OPERAÇÃO
		DESCRIÇÃO
		IMPORTÂNCIA
		MOEDA
		SALDO CONTABILÍSTICO
		MOEDA
	*/

	for {
		record, err := r.Read()
		if record == nil {
			break
		}
		if err != nil {
			return movimentos, err
		}

		/*
			2014-09-04
			2014-09-04
			SP-80054855619 Lisboagas
			-52,91
			EUR
			3.554,38
			EUR
		*/

		mov := core.Movimento{
			Descricao: trim(record[2]),
			MoedaI:    trim(record[4]),
			MoedaSC:   trim(record[6]),
		}

		mov.DataMovim, err = core.YMD2Time(record[0])
		if err != nil {
			return movimentos, err
		}

		mov.DataOper, err = core.YMD2Time(record[1])
		if err != nil {
			return movimentos, err
		}

		mov.Importancia, err = core.Money2Int64(record[3])
		if err != nil {
			return movimentos, err
		}

		mov.SaldoContab, err = core.Money2Int64(record[5])
		if err != nil {
			return movimentos, err
		}

		movimentos.PushBack(mov)
	}

	return movimentos, nil
}
