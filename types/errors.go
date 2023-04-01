package types

import (
	"errors"
	"fmt"
	"strings"
)

// TODO: Refactor to simplify reused errors such as NotSpecified and NotFound

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
		return "String is required."
	case ErrorStringShort:
		return "String is too short."
	case ErrorStringLong:
		return "String is too long."
	case ErrorStringInvalidChar:
		return "String contains invalid characters."
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

// PrefixedError will return e with any "String" prefix removed and s prefixed instead. TODO: #107
func PrefixedError(e error, s string) error {
	if e == nil {
		return nil
	}

	return errors.New(s + strings.TrimPrefix(e.Error(), "String"))
}

func (e ErrorString) ContextError(ctx ...any) error {
	if len(ctx) == 0 {
		return e
	}

	// Remove the period suffix, we will re-add it after.
	// Make the error singular if context is 1.
	err := strings.TrimSuffix(e.Error(), ".")
	if strings.HasSuffix(err, "s") && len(ctx) == 1 {
		strings.TrimSuffix(err, "s")
	}

	// Convert context elems to a []string
	var ctxStr []string
	notNum := false
	for _, c := range ctx {
		// If we find anything that isn't an int, we will want to use a different separator
		if _, ok := c.(int); !ok {
			notNum = true
		}

		// Special case for using map[string]bool as a de-duplicated []string
		if s, ok := c.(map[string]bool); ok {
			for v := range s {
				ctxStr = append(ctxStr, v)
			}
			continue
		}

		ctxStr = append(ctxStr, fmt.Sprintf("%v", c))
	}

	sep := " / "
	pre := " ("
	suf := ")"

	if notNum {
		sep = ", "
		pre = ": "
		suf = ""
	}

	// Return the context string with our original error
	return errors.New(fmt.Sprintf("%s%s%s%s.", err, pre, strings.Join(ctxStr, sep), suf))
}

//
// Account operation related errors
//

type ErrorAccount int64

const (
	ErrorAccountUnknown ErrorAccount = iota
	ErrorAccountEmailExists
	ErrorAccountEmailDomainEmpty
	ErrorAccountEmailTLDEmpty
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
	case ErrorAccountEmailDomainEmpty:
		return "mail: domain length is 0"
	case ErrorAccountEmailTLDEmpty:
		return "mail: domain does not contain a TLD"
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
// Session operation related errors
//

type ErrorSession int64

const (
	ErrorSessionUnknown ErrorSession = iota
	ErrorSessionNotSpecified
	ErrorSessionIndexNotFound
	ErrorSessionKeyNotFound
	ErrorSessionRefreshNotFound
	ErrorSessionClaimNotValid
)

// TODO: i18n here
func (e ErrorSession) Error() string {
	switch e {
	case ErrorSessionNotSpecified:
		return "No account for the session was specified."
	case ErrorSessionIndexNotFound:
		return "No session with this index was found."
	case ErrorSessionKeyNotFound:
		return "No session with this key was found."
	case ErrorSessionRefreshNotFound:
		return "No session with this refresh token was found."
	case ErrorSessionClaimNotValid:
		return "No valid account is associated with this claim."
	default:
		return ErrorUnknown.Error()
	}
}

//
// Report related errors
//

type ErrorReport int64

const (
	ErrorReportUnknown ErrorReport = iota
	ErrorReportNotFound
	ErrorReportNotSpecified
)

func (e ErrorReport) Error() string {
	switch e {
	case ErrorReportNotFound:
		return "The specified report was not found."
	case ErrorReportNotSpecified:
		return "No report was specified."
	default:
		return ErrorUnknown.Error()
	}
}

//
// Drug related errors
//

type ErrorDrug int64

const (
	ErrorDrugUnknown ErrorDrug = iota
	ErrorDrugNotFound
	ErrorDrugNotSpecified
)

func (e ErrorDrug) Error() string {
	switch e {
	case ErrorDrugNotFound:
		return "The specified drug was not found."
	case ErrorDrugNotSpecified:
		return "No drug was specified."
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
	ErrorContextNilCache
	ErrorContextCastFailed
)

// TODO: i18n here
func (e ErrorContext) Error() string {
	switch e {
	case ErrorContextNilLogger:
		return "ctx.Logger is nil"
	case ErrorContextNilDatabase:
		return "ctx.Database is nil"
	case ErrorContextNilCache:
		return "ctx.Cache is nil"
	case ErrorContextCastFailed:
		return "Context failed to cast!"
	default:
		return ErrorUnknown.Error()
	}
}
