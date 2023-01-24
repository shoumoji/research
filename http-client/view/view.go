package view

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"

	"github.com/shoumoji/research/http3-client/models"
)

func OutputJSON(results ...*models.Result) error {
	for _, r := range results {
		jsonData, err := json.Marshal(&r)
		if err != nil {
			return err
		}

		fmt.Println(string(jsonData))
	}
	return nil
}

func OutputCSV(shouldWriteHeader bool, count int64, results ...*models.Result) error {
	var records [][]string

	if shouldWriteHeader {
		records = make([][]string, count+1)

		for _, result := range results {
			records[0] = append(records[0], result.Protocol)
			for it, time := range result.TimeMicroSeconds {
				records[it+1] = append(records[it+1], fmt.Sprint(time))
			}
		}
	} else {
		records = make([][]string, count)

		for _, result := range results {
			for it, time := range result.TimeMicroSeconds {
				records[it] = append(records[it], fmt.Sprint(time))
			}
		}
	}

	csvWriter := csv.NewWriter(os.Stdout)
	if err := csvWriter.WriteAll(records); err != nil {
		return err
	}
	defer csvWriter.Flush()

	return nil
}
