package exception

import (
	"encoding/json"
	"net/http"
)

func mapStatus(code ErrorCode) int {
	switch code {
	case InvalidRequest:
	case UnknownData:
	case NotAllowed:
	case Duplicate:
	case Deleted:
		return http.StatusBadRequest
	case AuthFailed:
	case ExpiredToken:
	case TokenRequired:
		return http.StatusUnauthorized
	case PermissionDenied:
	case ModuleNotConfigured:
		return http.StatusForbidden
	}
	return http.StatusInternalServerError
}

func Throw(w http.ResponseWriter, code ErrorCode, err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(mapStatus(code))
	response := map[string]interface{}{
		"status":  code,
		"message": err.Error(),
	}
	er := json.NewEncoder(w).Encode(response)
	if er != nil {
		return
	}
	return
}

func ThrowInvalidRequest(w http.ResponseWriter, err error) {
	Throw(w, InvalidRequest, err)
}

func ThrowUnknownData(w http.ResponseWriter, err error) {
	Throw(w, UnknownData, err)
}

func ThrowNotAllowed(w http.ResponseWriter, err error) {
	Throw(w, NotAllowed, err)
}

func ThrowDuplicate(w http.ResponseWriter, err error) {
	Throw(w, Duplicate, err)
}

func ThrowDeleted(w http.ResponseWriter, err error) {
	Throw(w, Deleted, err)
}

func ThrowAuthFailed(w http.ResponseWriter, err error) {
	Throw(w, AuthFailed, err)
}

func ThrowExpiredToken(w http.ResponseWriter, err error) {
	Throw(w, ExpiredToken, err)
}

func ThrowTokenRequired(w http.ResponseWriter, err error) {
	Throw(w, TokenRequired, err)
}

func ThrowPermissionDenied(w http.ResponseWriter, err error) {
	Throw(w, PermissionDenied, err)
}

func ThrowModuleNotConfigured(w http.ResponseWriter, err error) {
	Throw(w, ModuleNotConfigured, err)
}

func ThrowInternalServerError(w http.ResponseWriter, err error) {
	Throw(w, ApiInternalError, err)
}
