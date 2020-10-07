package model

type Role struct {
	ID    uint `gorm:"primarykey"`
	Name  *string
	Users []*User `gorm:"many2many:user_roles;"`
}
