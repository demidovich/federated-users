package seeder

import (
	"encoding/json"
	"federated/internal/person"
	"fmt"
	"math/rand"
	"reflect"
	"strings"

	"github.com/go-faker/faker/v4"
)

type attrsGenerator struct {
	names []string
}

func newAttrsGenerator(a person.Attrs) attrsGenerator {
	return attrsGenerator{
		names: attrNames(a),
	}
}

func (a attrsGenerator) generateJson(attrsCount int) string {
	if attrsCount > len(a.names) {
		attrsCount = len(a.names)
	}

	attrs := make(map[string]any, attrsCount)
	for i := range attrsCount {
		name := a.names[i]
		value := a.generatedValue(name)
		if value != "" {
			attrs[name] = value
		}
	}

	jsonBytes, err := json.Marshal(attrs)
	if err != nil {
		return "{}"
	}

	return string(jsonBytes)
}

func (a attrsGenerator) generatedValue(name string) string {

	if strings.HasPrefix(name, "FirstName") {
		return faker.FirstName()
	}

	if strings.HasPrefix(name, "MiddleName") {
		return faker.FirstName()
	}

	if strings.HasPrefix(name, "LastName") {
		return faker.LastName()
	}

	if strings.HasPrefix(name, "Age") {
		return fmt.Sprintf("%d", rand.Int63n(40)+30)
	}

	if strings.HasPrefix(name, "Latitude") {
		return fmt.Sprintf("%f", faker.Latitude())
	}

	if strings.HasPrefix(name, "Longitude") {
		return fmt.Sprintf("%f", faker.Longitude())
	}

	if strings.HasPrefix(name, "Country") {
		return faker.GetCountryInfo().Name
	}

	if strings.HasPrefix(name, "Currency") {
		return faker.Currency()
	}

	if strings.HasPrefix(name, "Lang") {
		return faker.GetCountryInfo().Abbr
	}

	if strings.HasPrefix(name, "Address") {
		return faker.GetRealAddress().Address
	}

	if strings.HasPrefix(name, "Tag") {
		return faker.Word()
	}

	if strings.HasPrefix(name, "Social") {
		return fmt.Sprintf("%d", rand.Int63n(100000000)+100000000)
	}

	if strings.HasPrefix(name, "Attribute") {
		return faker.Word()
	}

	return ""
}

func attrNames(a person.Attrs) []string {
	fields := reflect.VisibleFields(reflect.TypeOf(a))

	result := []string{}
	for _, field := range fields {
		result = append(result, field.Name)
	}

	return result
}
