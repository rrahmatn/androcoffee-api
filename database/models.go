package database

type User struct {
	Id       int64  `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"type:varchar(300)" json:"name"`
	Email    string `gorm:"type:varchar(300)" json:"email"`
	Password string `gorm:"type:text" json:"password"`
}
