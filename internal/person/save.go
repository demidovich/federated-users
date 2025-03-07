package person

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type SaveCommand struct {
	FederationUuid string
	Attrs          Attrs
}

type saveUsecase struct {
	db *sqlx.DB
}

func (u saveUsecase) HandleCreate(command SaveCommand) (Person, error) {
	item := Person{}
	item.Uuid = uuid.NewString()
	item.FederationUuid = command.FederationUuid
	item.Attrs = &command.Attrs

	attrsJson, err := command.Attrs.ToJson()
	if err != nil {
		return Person{}, err
	}

	_, err = u.db.Query(
		`insert into person (uuid, federation_uuid, attrs) values ($1, $2, $3)`,
		item.Uuid,
		item.FederationUuid,
		attrsJson,
	)

	return item, err
}

func (u saveUsecase) HandleUpdate(item *Person, command SaveCommand) error {
	attrsJson, err := command.Attrs.ToJson()
	if err != nil {
		return err
	}

	item.FederationUuid = command.FederationUuid
	item.Attrs = &command.Attrs

	_, err = u.db.Query(
		`update person set federation_uuid = $1, arrts = $2 where uuid = $3`,
		item.FederationUuid,
		attrsJson,
		item.Uuid,
	)

	return err
}
