package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"recipe_api/src/model"
)

var DB *gorm.DB

func ConnectionString(user string, pass string, host string, port string, db string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, pass, host, port, db)
}

func Migrate() error {
	return DB.AutoMigrate(&model.User{}, &model.Role{}, &model.Post{})
}

func Open() error {
	var err error
	DB, err = gorm.Open(
		mysql.Open(
			ConnectionString(
				os.Getenv("DB_USER"),
				os.Getenv("DB_PASS"),
				os.Getenv("DB_HOST"),
				os.Getenv("DB_PORT"),
				os.Getenv("DB_NAME"),
			),
		),
		&gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
		},
	)
	return err
}
