package users

import "time"

type Core struct {
	ID        int
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Business interface {
	GetAllData(limit, offset int) (data []Core, err error)
	CreateData(data Core) (row int, err error)
	LoginUser(authData AuthRequestData) (token, name string, err error)
	GetUserById(id int) (data Core, err error)
}

type Data interface {
	SelectData(param string) (data []Core, err error)
	InsertData(data Core) (row int, err error)
	LoginUserDB(authData AuthRequestData) (token, name string, err error)
	SelectDataById(id int) (data Core, err error)
}
