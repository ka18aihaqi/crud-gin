package models

type User struct {
    ID       int64  `gorm:"primaryKey" json:"id"`
    Email    string `gorm:"size:300;unique" json:"email"`
    Username string `gorm:"size:300;unique" json:"username"`
    Name     string `gorm:"size:300" json:"name"`
    Password string `gorm:"size:300" json:"password"`
}
