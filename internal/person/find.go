package person

import "github.com/jmoiron/sqlx"

type findUsecase struct {
	db *sqlx.DB
}

func (u findUsecase) Handle(uuid string) (Person, error) {
	item := Person{}
	err := u.db.Get(&item, `select * from person where uuid = $1`, uuid)

	return item, err
}
