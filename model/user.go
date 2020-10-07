package model

type User struct {
	ID        uint `gorm:"primarykey"`
	Username  *string
	Email     *string
	AvatarURL *string
	Roles     []*Role `gorm:"many2many:user_roles;"`
	Posts     []*Post
}
