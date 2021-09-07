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
// @Updated
type Reader interface {

}

// @Created 06/09/2021
// @Updated
type Writer interface {
	Create(*interface{}) error
	Save()
}