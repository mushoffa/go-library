package database

import (
	"errors"
)

// @Created 06/09/2021
// @Updated
type Database interface {
	Reader
	Writer
}

// @Created 06/09/2021
// @Updated 07/09/2021
type Reader interface {
	FindByID(string, string, *interface{}) (error)
}

// @Created 06/09/2021
// @Updated 07/09/2021
type Writer interface {
	Create(*interface{}) error
	UpdateByID(string, string, *interface{}) (error)
}