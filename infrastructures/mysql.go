package infrastructures

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySQL(
	userName string,
	password string,
	dbHostname string,
	dbPort string,
	dbName string,
) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", userName, password, dbHostname, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
