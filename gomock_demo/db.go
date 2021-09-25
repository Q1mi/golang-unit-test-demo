package gomock_demo

import "database/sql"

// db.go

// DB 数据接口
type DB interface {
	Get(key string)(int, error)
	Add(key string, value int) error
}


type MySQL struct {
	sql.DB
}

func (m *MySQL) Get(key string)(int, error)  {
	// ...
	return 0, nil
}

func (m *MySQL) Add(key string, value int)error  {
	// ...
	return nil
}

// GetFromDB 根据key从DB查询数据的函数
func GetFromDB(db DB, key string) int {
	if v, err := db.Get(key);err == nil{
		return v
	}
	return -1
}

