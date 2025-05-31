package model

type Token struct {
	Id           uint   `gorm:"column:id; not null; primaryKey"`
	AccessToken  string `gorm:"column:access_token; type:varchar(255)"`
	RefreshToken string `gorm:"column:refresh_token; type:varchar(50)"`
	Username     string `gorm:"column:username; type:varchar(50)"`
	Expired      int64  `gorm:"column:expired; type:varchar(50)"`
}

func (Token) TableName() string {
	return "tokens"
}
