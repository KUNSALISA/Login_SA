package models

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"-"` // ไม่แสดงรหัสผ่าน
}

// SetPassword รับรหัสผ่านเป็นข้อความธรรมดาและแฮชมัน
func (u *User) SetPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return nil
}

// CheckPassword ตรวจสอบรหัสผ่านที่แฮชกับรหัสผ่านที่ผู้ใช้ป้อน
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
