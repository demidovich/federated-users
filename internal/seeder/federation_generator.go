package seeder

import (
	"math/rand"

	"github.com/google/uuid"
)

type federationGenerator struct {
	uuids []string
}

func newFederationGenerator(count int) federationGenerator {
	uuids := make([]string, 0, count)
	for range count - 1 {
		uuids = append(uuids, uuid.NewString())
	}

	return federationGenerator{
		uuids: uuids,
	}
}

func (f federationGenerator) randUuid() string {
	i := rand.Intn(len(f.uuids) - 1)
	return f.uuids[i]
}
