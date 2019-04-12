package model

// User model ใช้บันทึกข้อมูลลง DB
type User struct {
	UID      int    `db:"uid"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}
