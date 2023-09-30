package registry

import (
	"encoding/csv"
	"os"
)

func FindRecord(filePath string, name string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	csvReader := csv.NewReader(file)
	csvReader.FieldsPerRecord = -1
	read := true
	for read {
		record, err := csvReader.Read()
		if err != nil {
			panic(err)
		}
		if record == nil {
			break
		}
		if record[2] == name {
			return record
		}
	}
	return nil
}
