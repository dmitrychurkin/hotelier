package auth

import (
	"fmt"
	"net/http"

	"github.com/vektah/gqlparser/gqlerror"
)

// TODO: refactor to custom error codes
var (
	userNotFound = &gqlerror.Error{
		Message: "We don't have records with associated credentials",
		Extensions: map[string]interface{}{
			"code": http.StatusUnprocessableEntity,
		},
	}

	passwordsMismatch = &gqlerror.Error{
		Message: "Passwords mismatch",
		Extensions: map[string]interface{}{
			"code": http.StatusConflict,
		},
	}

	invalidinputFields = &gqlerror.Error{
		Message: "Some fields are invalid",
		Extensions: map[string]interface{}{
			"code": http.StatusConflict,
		},
	}

	allowSendEmailAfter = func(m int) *gqlerror.Error {
		return &gqlerror.Error{
			Message: fmt.Sprintf("Email already send. Try again after %d", m),
			Extensions: map[string]interface{}{
				"code": http.StatusForbidden,
			},
		}
	}

	invalidPasswordResetToken = &gqlerror.Error{
		Message: "Password reset token expired",
		Extensions: map[string]interface{}{
			"code": http.StatusForbidden,
		},
	}
)
