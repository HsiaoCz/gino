package redis

type RedisStorage struct {
}

func NewRedisStorage() *RedisStorage {
	return &RedisStorage{}
}

func (r *RedisStorage) InitStore() error {
	return nil
}
