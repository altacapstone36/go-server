package models


type Level struct {
	ID int `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

func (*Level) TableName() string {
	return "roles"
}
