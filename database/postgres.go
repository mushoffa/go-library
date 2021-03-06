package database

import (
	"errors"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

)

// @Created 06/09/2021
const (
	maxOpenConns    = 60
	connMaxLifetime = 120
	maxIdleConns    = 30
	connMaxIdleTime = 20
)


// @Created 06/09/2021
// @Updated 17/09/2021
type Postgres struct {
	Gorm *gorm.DB
}

// @Created 06/09/2021
// @Updated 17/09/2021
func NewPostgres(host, port, dbName, user, password string, sslmode bool) (Database, error) {
	postgres := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		host,
		port,
		dbName,
		user,
		password,
	)

	db, err := gorm.Open("postgres", postgres)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to open database: %s", err.Error()))
	}

	db.LogMode(true)
	db.DB().SetMaxOpenConns(maxOpenConns)
	db.DB().SetConnMaxLifetime(connMaxLifetime * time.Second)
	db.DB().SetMaxIdleConns(maxIdleConns)
	db.DB().SetConnMaxIdleTime(connMaxIdleTime * time.Second)

	if err = db.DB().Ping(); err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to connect database: %s", err.Error()))
	}

	return &Postgres{db}, nil
}

// @Created 16/10/2021
// @Updated
func (db *Postgres) GetInstance() *gorm.DB {
	// if(db.Gorm != nil) {
	// 	return db.Gorm
	// }

	// return nil
	return db.Gorm
}

// @Created 07/09/2021
// @Updated
func (db *Postgres) FindByID(queryField, queryID string, data interface{}) (error) {

	query := fmt.Sprintf("%s = ?", queryField)

	if err := db.Gorm.Model(data).Where(query, queryID).Take(data).Error; err != nil {
		return err
	}

	return nil
}

// @Created 16/10/2021
// @Updated
func (db *Postgres) FindAll(model interface{}) error {
	if err := db.Gorm.Find(model).Error; err != nil {
		return err
	}
	return nil
}

// @Created 17/09/2021
func (db *Postgres) AutoMigrate(models ...interface{}) {
	
	for _, model := range models {
		db.Gorm.AutoMigrate(model)
		// go func() {
		// 	db.Gorm.AutoMigrate(&model)
		// }()
	}
}


// @Created 06/09/2021
// @Updated
func (db *Postgres) Create(model interface{}) error {

	if err := db.Gorm.Create(model).Error; err != nil {
		return err
	}

	return nil
}

// @Created 07/09/2021
// @Updated 16/10/2021
func (db *Postgres) UpdateByID(queryField, queryID, updateField string, updateID, data interface{}) (error) {

	query := fmt.Sprintf("%s = ?", queryField)

	if err := db.Gorm.Model(data).Where(query, queryID).Update(updateField, updateID).Error; err != nil {
		return err
	}

	return nil
}