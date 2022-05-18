package database

// @Created 06/09/2021
// @Updated
type Database interface {
	Reader
	Writer
}

// @Created 06/09/2021
// @Updated 18/05/2022
type Reader interface {
	GetInstance() interface{}
	FindByID(string, string, interface{}) (error)
	FindAll(interface{}) error
}

// @Created 06/09/2021
// @Updated 16/10/2021
type Writer interface {
	AutoMigrate(...interface{})
	Create(interface{}) error
	UpdateByID(string, string, string, interface{}, interface{}) (error)
}