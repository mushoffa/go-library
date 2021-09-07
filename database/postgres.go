package database

import (
	"errors"
	"fmt"
	"time"

	"github.com/mushoffa/go-library/config"

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
// @Updated
type Postgres struct {
	postgres *gorm.DB
}

// @Created 06/09/2021
// @Updated
func NewPostgres(c config.Config) (Database, error) {
	postgres := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		c.Postgres.PostgresqlHost,
		c.Postgres.PostgresqlPort,
		c.Postgres.PostgresqlDbName,
		c.Postgres.PostgresqlUser,
		c.Postgres.PostgresqlPassword,
	)

	db, err := gorm.Open("postgres", postgres)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to open database: %s", err.Error()))
	}

	db.LogMode(true)
	// db.AutoMigrate(&dbmodels.User{})
	db.DB().SetMaxOpenConns(maxOpenConns)
	db.DB().SetConnMaxLifetime(connMaxLifetime * time.Second)
	db.DB().SetMaxIdleConns(maxIdleConns)
	db.DB().SetConnMaxIdleTime(connMaxIdleTime * time.Second)

	if err = db.DB().Ping(); err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to connect database: %s", err.Error()))
	}

	return &Postgres{db}, nil
}


// @Created 06/09/2021
// @Updated
func (db *Postgres) Create(model *interface{}) error {

	if err := db.postgres.Create(model).Error; err != nil {
		return err
	}

	return nil
}

// @Created 07/09/2021
// @Updated
func (db *Postgres) FindByID(queryField, queryID string, data *interface{}) (error) {

	query := fmt.Sprintf("%s = ?", queryField)

	if err := db.postgres.Model(data).Where(query, queryID).Take(data).Error; err != nil {
		return err
	}

	return nil
}

// @Created 07/09/2021
// @Updated
func (db *Postgres) UpdateByID(queryField, queryID, updateField, updateID string, data *interface{}) (error) {

	query := fmt.Sprintf("%s = ?", queryField)
	update := fmt.Sprintf("%s = ?", updateField)

	if err := db.postgres.Model(data).Where(query, queryID).Update(update, updateID).Error; err != nil {
		return err
	}

	return nil
}