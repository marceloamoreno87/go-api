package notification

import "strings"

type ErrorsInterface interface {
	AddError(message string, context string)
	Messages() string
	HasErrors() bool
	CheckRequiredField(field, fieldName, errorKey string)
	CheckIsContains(field string, array []string, fieldName, errorKey string)
}

type NotificationError struct {
	Message string `json:"message"`
	Context string `json:"context"`
}

type Errors struct {
	Errors []NotificationError `json:"errors"`
}

func New() *Errors {
	return &Errors{}
}

func (e *Errors) AddError(message string, context string) {
	e.Errors = append(e.Errors, NotificationError{Message: message, Context: context})
}

func (e *Errors) Messages() (messages string) {
	var errs []string
	for _, err := range e.Errors {
		errs = append(errs, "["+err.Context+"]: "+err.Message)
	}

	return strings.Join(errs, ", ")
}

func (e *Errors) HasErrors() bool {
	return len(e.Errors) > 0
}

func (e *Errors) CheckRequiredField(field, fieldName, errorKey string) {
	if field == "" {
		e.AddError(fieldName+" is required", errorKey)
	}
}

func (e *Errors) CheckIsContains(field string, array []string, fieldName, errorKey string) {
	found := false
	for _, value := range array {
		if value == field {
			found = true
			break
		}
	}
	if !found {
		e.AddError(fieldName+" is invalid", errorKey)
	}
}
