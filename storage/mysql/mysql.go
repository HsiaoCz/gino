package mysql

type MysqlStorage struct {
}

func NewMysqlStorage() *MysqlStorage {
	return &MysqlStorage{}
}

func (m *MysqlStorage)InitStore()error{
	return nil
}