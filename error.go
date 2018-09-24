package main

type ErrNotFound struct{}

func (e ErrNotFound) Error() string {
	return "target record is not found"
}
