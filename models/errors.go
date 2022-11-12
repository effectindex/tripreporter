package models

//
// Generic errors
//

type ErrorGeneric int64

const (
	ErrorGenericUnknown ErrorGeneric = iota
)

// TODO: i18n here
func (e ErrorGeneric) Error() string {
	switch e {
	default:
		return "An unknown error occurred."
	}
}

//
// Account operation related errors
//

type ErrorAccount int64

const (
	ErrorAccountUnknown ErrorAccount = iota
	ErrorAccountEmailExists
	ErrorAccountUsernameExists
	ErrorAccountNotSpecified
	ErrorAccountNotFound
	ErrorAccountNotDeleted
	ErrorAccountPasswordMatch
)

// TODO: i18n here
func (e ErrorAccount) Error() string {
	switch e {
	case ErrorAccountEmailExists:
		return "An account with that email already exists."
	case ErrorAccountUsernameExists:
		return "An account with that username already exists."
	case ErrorAccountNotSpecified:
		return "No account was specified."
	case ErrorAccountNotFound:
		return "The specified account was not found."
	case ErrorAccountNotDeleted:
		return "Failed to delete account."
	case ErrorAccountPasswordMatch:
		return "Incorrect password or account."
	default:
		return ErrorGenericUnknown.Error()
	}
}

//
// Context related errors
//

type ErrorContext int64

const (
	ErrorContextUnknown ErrorContext = iota
	ErrorContextNilLogger
	ErrorContextNilDatabase
)

// TODO: i18n here
func (e ErrorContext) Error() string {
	switch e {
	case ErrorContextNilLogger:
		return "ctx.Logger is nil"
	case ErrorContextNilDatabase:
		return "ctx.Database is nil"
	default:
		return ErrorGenericUnknown.Error()
	}
}
