package model

type Role struct {
	ID    uint
	Name  *string
	Users []*User `gorm:"many2many:user_roles;"`
}
