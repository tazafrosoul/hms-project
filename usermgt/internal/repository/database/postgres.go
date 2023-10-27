package database

import s "github.com/tazafrosoul/hms-project/common/structs"

type PostgresDB struct {
}

func NewPostgresDB() *PostgresDB {
	return &PostgresDB{}
}

func (pdb *PostgresDB) Create(newuser s.User) {

}
