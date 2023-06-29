package database

import s "hms-project/common/structs"

type PostgresDB struct {
}

func NewPostgresDB() *PostgresDB {
	return &PostgresDB{}
}

func (pdb *PostgresDB) Create(newuser s.User) {

}
