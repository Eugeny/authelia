package handlers

import (
	"github.com/authelia/authelia/v4/internal/authentication"
	"github.com/authelia/authelia/v4/internal/authorization"
)

// newAnonymousUserDetails returns an empty extended user details value. The embedded *UserDetails is initialized as
// consumers dereference the promoted fields without a nil check.
func newAnonymousUserDetails() authentication.UserDetailsExtended {
	return authentication.UserDetailsExtended{UserDetails: &authentication.UserDetails{}}
}

func friendlyMethod(m string) (fm string) {
	switch m {
	case "":
		return "unknown"
	default:
		return m
	}
}

func friendlyUsername(username string) (fusername string) {
	switch username {
	case "":
		return anonymous
	default:
		return username
	}
}

func isAuthzResult(level authentication.Level, required authorization.Level, ruleHasSubject bool) AuthzResult {
	switch {
	case required == authorization.Bypass:
		return AuthzResultAuthorized
	case required == authorization.Denied && (level != authentication.NotAuthenticated || !ruleHasSubject):
		// If the user is not anonymous, it means that we went through all the rules related to that user identity and
		// can safely conclude their access is actually forbidden. If a user is anonymous however this is not actually
		// possible without some more advanced logic.
		return AuthzResultForbidden
	case required == authorization.OneFactor && level >= authentication.OneFactor,
		required == authorization.TwoFactor && level >= authentication.TwoFactor:
		return AuthzResultAuthorized
	default:
		return AuthzResultUnauthorized
	}
}
