package models

// import (
// 	"github.com/go-playground/validator/v10"
// )

type ValidationError struct {
	Message string              `json:"message"`
	Details []ValidationDetails `json:"details"`
}

type ValidationDetails struct {
	Field      string `json:"field"`
	Value      string `json:"value"`
	Constraint string `json:"constraint"`
}
