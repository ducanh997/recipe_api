package model

type User struct {
	ID        uint `gorm:"primarykey"`
	Username  *string
	Age       *int
	Email     *string
	AvatarURL *string
	Roles     []*Role `gorm:"many2many:user_roles;"`
	Posts     []*Post
}
