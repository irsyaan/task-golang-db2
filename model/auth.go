package model

type Auth struct {
	AuthID    int64 `gorm:"primaryKey;autoIncrement;<-:false"`
	AccountID int64
	Username  string `json:"username"`
	Password  string `json:"password"`
}
