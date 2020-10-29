package model

type Category struct {
	ID    uint
	Name  *string
	Posts []*Post `gorm:"many2many:category_posts;"`
}
