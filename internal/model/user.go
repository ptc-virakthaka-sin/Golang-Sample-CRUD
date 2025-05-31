package model

type User struct {
	Id       uint   `gorm:"column:id; not null; primaryKey"`
	Username string `gorm:"column:username; type:varchar(255)"`
	Password string `gorm:"column:password; type:varchar(255)"`
	Role     string `gorm:"column:role; type:varchar(50)"`
}

func (User) TableName() string {
	return "users"
}
