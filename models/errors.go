package models

//
// Generic errors
//

type ErrorGeneric int64

const (
	ErrorGenericUnknown ErrorGeneric = iota
	ErrorNotImplemented
)

// TODO: i18n here
func (e ErrorGeneric) Error() string {
	switch e {
	case ErrorNotImplemented:
		return "This method is not implemented yet."
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
// User operation related errors
//

type ErrorUser int64

const (
	ErrorUserUnknown ErrorUser = iota
	ErrorUserNotSpecified
	ErrorUserNotFound
	ErrorUserNotDeleted
	ErrorUserAccountStillExists
)

// TODO: i18n here
func (e ErrorUser) Error() string {
	switch e {
	case ErrorUserNotSpecified:
		return "No user was specified."
	case ErrorUserNotFound:
		return "The specified user was not found."
	case ErrorUserNotDeleted:
		return "Failed to delete user."
	case ErrorUserAccountStillExists:
		return "The associated account with this user still exists."
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
