package main

import (
	"errors"
	"fmt"
)

var (
	ErrPoolInv    = errors.New("invalid pool")
	ErrConnClose  = errors.New("connection closed")
	ErrConnInv    = errors.New("invalid connection")
	ErrNotExists  = errors.New("user id was not found")
	ErrExists     = errors.New("email was already exists")
	ErrConnFailed = errors.New("connection failed")
	ErrQuery      = errors.New("execute query error")
	ErrBeginTx    = errors.New("begin transaction error")
	ErrScan       = errors.New("scan error")
	ErrCommit     = errors.New("commit error")
	ErrExec       = errors.New("query exec error")
	ErrRollback   = errors.New("rollback error")

	ErrHashPass        = errors.New("hash password error")
	ErrPasswordIsEmpty = errors.New("password is empty")
	ErrPasswordIsShort = errors.New("password is short")
)

type ErrValidation struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ErrDatabase struct {
	Err  error
	Type string
	At   string
}

type CustomError struct {
	Err           error
	CustomMessage string
}

func (e ErrDatabase) Error() string {
	return fmt.Sprintf("database error: %s, type: %s, at: %s", e.Err, e.Type, e.At)
}

func (e CustomError) Error() string {
	return fmt.Sprintf("error: %s,  message: %s", e.Err, e.CustomMessage)
}