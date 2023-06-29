package database

import s "hms-project/common/structs"

type RedisDB struct {
}

func NewRedisDB() *RedisDB {
	return &RedisDB{}
}

func (rdb *RedisDB) Create(newuser s.User) {

}
