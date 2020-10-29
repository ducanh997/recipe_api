package model

type Post struct {
	ID         uint
	Title      *string
	Content    *string
	UserID     *uint
	User       *User
	Categories []*Category `gorm:"many2many:category_posts;"`
}
