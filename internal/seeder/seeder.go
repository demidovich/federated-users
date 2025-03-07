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

func GenerateRows(rowsCount, minAttrs, maxAttrs int) (string, error) {
	file, writer, err := newCsvWriter(csvFileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// writer.Write([]string{"uuid", "federation_uuid", "attrs"})

	attrs := newAttrsGenerator(person.Attrs{})
	federations := newFederationGenerator(100)

	for range rowsCount - 1 {
		federationUuid := federations.randUuid()

		row := make([]string, 0, 3)
		row = append(row, uuid.NewString())
		row = append(row, federationUuid)

		attrsCount := rand.Intn(maxAttrs-minAttrs) + minAttrs
		row = append(row, attrs.generateJson(federationUuid, attrsCount))

		err := writer.Write(row)
		if err != nil {
			fmt.Println(err)
		}
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		fmt.Println("Error flushing CSV writer:", err)
	}

	return csvFileName, nil
}

func newCsvWriter(fileName string) (*os.File, *csv.Writer, error) {
	file, err := os.Create(fileName)
	if err != nil {
		return nil, nil, err
	}

	return file, csv.NewWriter(file), nil
}
