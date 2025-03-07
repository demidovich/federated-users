package seeder

import (
	"encoding/csv"
	"federated/internal/person"
	"fmt"
	"math/rand"
	"os"

	"github.com/google/uuid"
)

const csvFileName = "./database/seed/person.csv"

func GenerateRows(rowsCount int) (string, error) {

	csvFile := csvFile(csvFileName)
	defer csvFile.Close()

	csvWriter := csv.NewWriter(csvFile)
	csvWriter.Write([]string{"uuid", "federation_uuid", "attrs"})

	attrs := newAttrsGenerator(person.Attrs{})
	federations := newFederationGenerator(100)

	for range rowsCount - 1 {
		row := make([]string, 0, 3)
		row = append(row, uuid.NewString())
		row = append(row, federations.randUuid())

		rowAttrsCount := rand.Intn(90) + 10
		row = append(row, attrs.generateJson(rowAttrsCount))

		err := csvWriter.Write(row)
		if err != nil {
			fmt.Println(err)
		}
	}

	csvWriter.Flush()
	if err := csvWriter.Error(); err != nil {
		fmt.Println("Error flushing CSV writer:", err)
	}

	return csvFileName, nil
}

func csvFile(fileName string) *os.File {
	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	return f
}
