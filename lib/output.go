package jc

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"sort"

	"github.com/jmespath/go-jmespath"
	"github.com/olekukonko/tablewriter"
	"github.com/yukithm/json2csv"
)

func outputData(outputFormat string, query string, res interface{}) {

	jsondata, err := json.Marshal(res)
	if err != nil {
		log.Fatalf("Could not marshal system users, err='%s'\n", err)
	}

	var output []byte
	var data interface{}
	err = json.Unmarshal(jsondata, &data)
	if err != nil {
		log.Fatalf("error unmarshalling data: %s", err)
	}
	if query != "" {
		result, err := jmespath.Search(query, data)
		if err != nil {
			log.Fatalf("Error querying data: %s\n", err)
		}
		output, _ = json.MarshalIndent(result, "", "  ")

	} else {
		output, _ = json.MarshalIndent(data, "", "  ")
	}

	if outputFormat == "table" {

		var tableData []map[string]string

		var csvOutput []map[string]interface{}
		_ = json.Unmarshal(output, &csvOutput)
		for i, record := range csvOutput {
			for key, value := range record {
				csvOutput[i][key] = fmt.Sprintf("%v", value)
			}
		}

		output, _ = json.MarshalIndent(csvOutput, "", "  ")
		err = json.Unmarshal(output, &tableData)
		header, records := makeTableData(tableData)

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader(header)
		table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		table.SetCenterSeparator("|")
		table.AppendBulk(records)
		table.Render()

		// Format the output as CSV or TEXT
	} else if outputFormat == "csv" {
		var csvOutput []map[string]interface{}
		err := json.Unmarshal(output, &csvOutput)

		if len(csvOutput) == 0 {
			return
		}

		// Single column records should not print headers
		if len(csvOutput[0]) == 1 {
			for _, records := range csvOutput {
				for _, value := range records {
					fmt.Println(value)
				}
			}
			return
		}

		results, err := json2csv.JSON2CSV(csvOutput)
		if err != nil {
			log.Fatal(err)
		}

		err = printCSV(os.Stdout, results, json2csv.DotNotationStyle, false)
		if err != nil {
			log.Fatal(err)
		}

		// Format the output as JSON
	} else {
		fmt.Println(string(output))
	}

}

func makeHeader(mymap map[string]string) []string {
	keys := make([]string, len(mymap))

	i := 0
	for k := range mymap {
		keys[i] = k
		i++
	}
	return keys
}

func makeTableData(mymap []map[string]string) (header []string, records [][]string) {
	header = makeHeader(mymap[0])
	sort.Strings(header)
	for _, record := range mymap {
		var row []string
		sum := 0
		for i := 0; i < len(header); i++ {
			sum += i
			row = append(row, record[header[i]])
		}

		records = append(records, row)
	}

	return header, records
}

func printCSV(w io.Writer, results []json2csv.KeyValue, headerStyle json2csv.KeyStyle, transpose bool) error {
	csv := json2csv.NewCSVWriter(w)
	csv.HeaderStyle = headerStyle
	csv.Transpose = transpose
	if err := csv.WriteCSV(results); err != nil {
		return err
	}
	return nil
}
