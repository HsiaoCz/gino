package storage

type InitStorage interface {
	InitStore() error
}

type istore struct {
	stores []InitStorage
}

func NewIstore() *istore {
	return &istore{
		stores: make([]InitStorage, 0),
	}
}

func (i *istore) StartConnStorage(stores ...InitStorage) (err error) {
	i.stores = append(i.stores, stores...)
	for _, store := range i.stores {
		if err = store.InitStore(); err != nil {
			return
		}
	}
	return
}
