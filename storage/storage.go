package storage

import (
	"github.com/HsiaoCz/gino/storage/mysql"
	"github.com/HsiaoCz/gino/storage/redis"
)

type Storage struct {
	Ms *mysql.MysqlStorage
	Rs *redis.RedisStorage
}

func NewStorage() *Storage {
	return &Storage{
		Ms: mysql.NewMysqlStorage(),
		Rs: redis.NewRedisStorage(),
	}
}
