package ierror

import (
	"net/http"
)

const (
	ErrCodeGeneralError                = "GENERAL_ERROR"
	ErrCodeRequestError                = "REQUEST_ERROR"
	ErrCodeValidationError             = "VALIDATION_ERROR"
	ErrCodeAuthenticationError         = "AUTHENTICATION_ERROR"
	ErrCodeAuthorizationError          = "AUTHORIZATION_ERROR"
	ErrCodeTokenError                  = "TOKEN_ERROR"
	ErrCodeDatabaseError               = "DATABASE_ERROR"
	ErrCodeRedisError                  = "REDIS_ERROR"
	ErrCodeDataNotFound                = "DATA_NOT_FOUND"
	ErrCodeDataAlreadyExist            = "DATA_ALREADY_EXIST"
	ErrCodeInsufficientFunds           = "INSUFFICIENT_FUNDS"
	ErrCodeIncorrectPin                = "INCORRECT_PIN"
	ErrCodeAccountNotFound             = "ACCOUNT_NOT_FOUND"
	ErrCodePhoneNumberNotFound         = "PHONE_NUMBER_NOT_FOUND"
	ErrCodePepsRisk                    = "PEPS_RISK"
	ErrCodePhoneNumberNotMatch         = "PHONE_NUMBER_NOT_MATCH"
	ErrCodeAccountBlocked              = "ACCOUNT_BLOCKED"
	ErrCodeWeakPin                     = "WEAK_PIN"
	ErrCodeIncorrectAmount             = "INCORRECT_AMOUNT"
	ErrCodeIncorrectDigits             = "INCORRECT_DIGITS"
	ErrCodeTransferExceedLimit         = "TRANSFER_EXCEED_LIMIT"
	ErrCodeAccountNotEnable            = "ACCOUNT_NOT_ENABLE"
	ErrCodeCustomerNotEnable           = "CUSTOMER_NOT_ENABLE"
	ErrCodeGLError                     = "GL_ERROR"
	ErrCodeDepositHistoryNotFound      = "DEPOSIT_HISTORY_NOT_FOUND"
	ErrCodeWithdrawalHistoryNotFound   = "WITHDRAWAL_HISTORY_NOT_FOUND"
	ErrCodeBillPaymentHistoryNotFound  = "BILL_PAYMENT_HISTORY_NOT_FOUND"
	ErrCodeUnableDeleteData            = "UNABLE_DELETE_DATA"
	ErrCodeInvalidCustomerId           = "INVALID_CUSTOMER_ID"
	ErrCodeInvalidVersion              = "INVALID_VERSION"
	ErrCodeInvalidExpiryDate           = "INVALID_EXPIRY_DATE"
	ErrCodeInvalidIssueDate            = "INVALID_ISSUE_DATE"
	ErrCodeInvalidDateOfBirth          = "INVALID_DATE_OF_BIRTH"
	ErrCodeAmountProviderNotMatch      = "AMOUNT_PROVIDER_NOT_MATCH"
	ErrCodeFirstNameOrLastNameNotExist = "FIRST_NAME_OR_LAST_NAME_NOT_EXIST"
	ErrCodeInvalidUsernameOrPassword   = "INVALID_USERNAME_OR_PASSWORD"
	ErrCodeUserNotActive               = "USER_NOT_ACTIVE"
)

const (
	ErrSourceNameMApi = "mapi"
)

var (
	ErrInvalidRequestBody = NewClientError(http.StatusBadRequest, ErrCodeRequestError, "invalid request body")
	ErrInvalidQueryString = NewClientError(http.StatusBadRequest, ErrCodeRequestError, "invalid query string")
)
