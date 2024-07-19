package exception

type ErrorCode string

const (
	// 400: Bad Request
	InvalidRequest ErrorCode = "InvalidRequest"
	UnknownData    ErrorCode = "UnknownData"
	NotAllowed     ErrorCode = "NotAllowed"
	Duplicate      ErrorCode = "Duplicate"
	Deleted        ErrorCode = "Deleted"
	// 401: Unauthorized
	AuthFailed    ErrorCode = "AuthFailed"
	ExpiredToken  ErrorCode = "ExpiredToken"
	TokenRequired ErrorCode = "TokenRequired"
	// 403: Forbidden
	PermissionDenied    ErrorCode = "PermissionDenied"
	ModuleNotConfigured ErrorCode = "ModuleNotConfigured"
	// 500: Internal Error
	ApiInternalError ErrorCode = "ApiInternalError"
)
