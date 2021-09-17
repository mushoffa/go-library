package database

import (
	"log"
	"testing"

	"github.com/mushoffa/go-library/config"
	"github.com/mushoffa/go-library/database"
)

type testStruct struct {
	ID int64 `gorm:"column:id;autoIncrement:true"`
	Data string `gorm:"column:data"`
	Data2 string `gorm:"column:data2"`
}

func (db *testStruct) TableName() string {
	return "test5"
}

type testStruct1 struct {
	Data string
}

func (db *testStruct1) TableName() string {
	return "public.test2"
}

// @Created 17/09/2021
func TestNewPostgres_AutoMigrate(t *testing.T) {

	db := testInitPostgresDB()
	if db == nil {
		t.Errorf("Can't connect to database")
	}

	db.AutoMigrate(&testStruct{})
}

// @Created 17/09/2021
func TestNewPostgres_Create(t *testing.T) {
	db := testInitPostgresDB()
	if db == nil {
		t.Errorf("Can't connect to database")
	}

	test := testStruct{
		Data: "test create6",
		Data2: "secondary data6",
	}
	err := db.Create(&test)
	if err != nil {
		t.Errorf("Error insert to database: %v", err)
	}
}

// @Created 17/09/2021
func TestNewPostgres_FindByID(t *testing.T) {
	db := testInitPostgresDB()
	if db == nil {
		t.Errorf("Can't connect to database")
	}

	test := testStruct{}
	err := db.FindByID("id", "1",&test)
	if err != nil {
		t.Errorf("Error query to database: %v", err)
	}

	log.Println(test)
}

// @Created 17/09/2021
func TestNewPostgres_UpdateByID(t *testing.T) {
	db := testInitPostgresDB()
	if db == nil {
		t.Errorf("Can't connect to database")
	}

	test := testStruct{}
	err := db.UpdateByID("id", "1", "data", "update data1",&test)
	if err != nil {
		t.Errorf("Error update to database: %v", err)
	}

	log.Println(test)
}

func testInitPostgresConfig() config.Config {
	postgresConfig := config.PostgresConfig {
		PostgresqlHost: "localhost",
		PostgresqlPort: "5432",
		PostgresqlDbName: "postgres-dev",
		PostgresqlUser: "admin",
		PostgresqlPassword: "password",
		PostgresqlSSLMode: false,
	}

	cfg := config.Config {Postgres: postgresConfig}

	return cfg
} 

func testInitPostgresDB() (database.Database ) {

	cfg := testInitPostgresConfig()

	db, err := database.NewPostgres(
		cfg.Postgres.PostgresqlHost,
		cfg.Postgres.PostgresqlPort,
		cfg.Postgres.PostgresqlDbName,
		cfg.Postgres.PostgresqlUser,
		cfg.Postgres.PostgresqlPassword,
		cfg.Postgres.PostgresqlSSLMode,
	)
	if err != nil {
		log.Println("Error: ", err)
		return nil
	}

	return db
}