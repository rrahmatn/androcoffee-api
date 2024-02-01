package database

type User struct {
	Id           int64  `gorm:"primaryKey" json:"id" `
	Name         string `gorm:"type:varchar(300)" json:"name" validate:"required"`
	Email        string `gorm:"type:varchar(300)" json:"email" validate:"email,required"`
	Password     string `gorm:"type:text" json:"password" validate:"required"`
	ConfPassword string `json:"confPassword" validate:"required"`
}
