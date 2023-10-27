package repository

import (
	i "github.com/tazafrosoul/hms-project/common/interfaces"
	s "github.com/tazafrosoul/hms-project/common/structs"
)

// TODO implement redis as a cache and postgress as main DB system
// TODO on the Init stage (within NewRepo/Repo.Run/Repo.Init) Hotreload Postgres most queried data to redis as a cache
type Repo struct {
	i.DB
}

func NewRepo(db i.DB) *Repo {
	return &Repo{
		DB: db,
	}
}

func (mr *Repo) Create(newuser s.User) {

}
