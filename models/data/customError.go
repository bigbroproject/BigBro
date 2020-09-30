package data

import (
	"encoding/json"
	"errors"
)

type CustomErr struct {
	error
}

func NewCustomErr(msg string) CustomErr{
	return CustomErr{errors.New(msg)}
}

func (cE CustomErr) MarshalJSON() ([]byte, error) {
	return json.Marshal(cE.Error())
}
