package users

import "time"

type Core struct {
	ID        int64
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Business interface {
	GetAllData(limit, offset int) (data []Core, err error)
}

type Data interface {
	SelectData(param string) (data []Core, err error)
}
