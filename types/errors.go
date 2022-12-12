package types

import (
	"errors"
	"strings"
)

//
// Generic errors
//

type ErrorGeneric int64

const (
	ErrorUnknown ErrorGeneric = iota
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
// String related errors
//

type ErrorString int64

const (
	ErrorStringUnknown ErrorString = iota
	ErrorStringEmpty
	ErrorStringShort
	ErrorStringLong
	ErrorStringInvalidChar
	ErrorStringUniqueChar
	ErrorStringSymbolChar
	ErrorStringNonSymbolChar
)

// TODO: i18n here
func (e ErrorString) Error() string {
	switch e {
	case ErrorStringEmpty:
		return "String cannot be empty."
	case ErrorStringShort:
		return "String is too short."
	case ErrorStringLong:
		return "String is too long."
	case ErrorStringInvalidChar:
		return "String contains invalid character(s)."
	case ErrorStringUniqueChar:
		return "String does not contain enough unique characters."
	case ErrorStringSymbolChar:
		return "String does not contain enough symbol characters."
	case ErrorStringNonSymbolChar:
		return "String does not contain enough non-symbol characters."
	default:
		return ErrorUnknown.Error()
	}
}

func (e ErrorString) PrefixedError(s string) error {
	return errors.New(s + strings.TrimPrefix(e.Error(), "String"))
}

//
// Account operation related errors
//

type ErrorAccount int64

const (
	ErrorAccountUnknown ErrorAccount = iota
	ErrorAccountEmailExists
	ErrorAccountEmailEmpty
	ErrorAccountEmailDomainEmpty
	ErrorAccountEmailTLDEmpty
	ErrorAccountUsernameExists
	ErrorAccountUsernameEmpty
	ErrorAccountUsernameShort
	ErrorAccountUsernameLong
	ErrorAccountUsernameInvalid
	ErrorAccountUsernameUniqueChar
	ErrorAccountUsernameSymbolChar
	ErrorAccountUsernameNonSymbolChar
	ErrorAccountNotSpecified
	ErrorAccountNotFound
	ErrorAccountNotDeleted
	ErrorAccountPasswordMatch
	ErrorAccountPasswordRequirements
	ErrorAccountPasswordEmpty
	ErrorAccountPasswordSaltEmpty
)

// TODO: i18n here
func (e ErrorAccount) Error() string {
	switch e {
	case ErrorAccountEmailExists:
		return "An account with that email already exists."
	case ErrorAccountEmailEmpty:
		return "Email is required."
	case ErrorAccountEmailDomainEmpty:
		return "mail: domain length is 0"
	case ErrorAccountEmailTLDEmpty:
		return "mail: domain does not contain a TLD"
	case ErrorAccountUsernameExists:
		return "An account with that username already exists."
	case ErrorAccountUsernameEmpty:
		return "Username is required."
	case ErrorAccountUsernameShort:
		return "Username is too short."
	case ErrorAccountUsernameLong:
		return "Username is too long."
	case ErrorAccountUsernameInvalid:
		return "Username contains invalid character(s)."
	case ErrorAccountUsernameUniqueChar:
		return "Username does not contain enough unique characters."
	case ErrorAccountUsernameSymbolChar:
		return "Username does not contain enough symbol characters."
	case ErrorAccountUsernameNonSymbolChar:
		return "Username does not contain enough non-symbol characters."
	case ErrorAccountNotSpecified:
		return "No account was specified."
	case ErrorAccountNotFound:
		return "The specified account was not found."
	case ErrorAccountNotDeleted:
		return "Failed to delete account."
	case ErrorAccountPasswordMatch:
		return "Incorrect password or account."
	case ErrorAccountPasswordRequirements:
		return "Password does not match requirements."
	case ErrorAccountPasswordEmpty:
		return "Password is required."
	case ErrorAccountPasswordSaltEmpty:
		return "Password salt is required."
	default:
		return ErrorUnknown.Error()
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
	ErrorUserBirthNotSpecified
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
	case ErrorUserBirthNotSpecified:
		return "User's age was specified without specifying date of birth."
	default:
		return ErrorUnknown.Error()
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
		return ErrorUnknown.Error()
	}
}
