package exceptions

import "errors"

var (
	DuplicateId = errors.New("book duplicate id")
)