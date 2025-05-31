package constant

const (
	Body    = "body"
	Query   = "query"
	Params  = "params"
	Header  = "header"
	Cookie  = "cookie"
	Success = "success"
	Bearer  = "Bearer "
	Error   = "error"
)

// variables scoped in the request context
const (
	ContextValidatedBody  = "validated_body"
	ContextValidatedQuery = "validated_query"
	ContextUser           = "user"
)

const (
	AddressTypeBirth   = "BIRTH"
	AddressTypeLiving  = "LIVING"
	AddressTypeWorking = "WORKING"
)

const (
	AccountTypeSavings       = "101"
	AccountTypeMobileSavings = "102"
)

const (
	AccountStatusEnabled  = "ENAB"
	AccountStatusDisabled = "DISA"
	AccountStatusDeleted  = "DELE"
	AccountStatusProForma = "FORM"
	AccountStatusPending  = "PEND"
)

const (
	CustomerStatusEnabled   = "ENAB"
	CustomerStatusDisabled  = "DISA"
	CustomerStatusSuspended = "SUSP"
	CustomerStatusDeleted   = "DELE"
)

const (
	BranchCodeHQ = "00"
)

const (
	ConsentTypeCodePDPA        = "101"
	ConsentTypeCodeTermAndCond = "102"
	ConsentTypeCodeFacePass    = "103"
	ConsentTypeMobileSavings   = "104"
)

const (
	ConsentStatusDraft     = "DRAFT"
	ConsentStatusPublish   = "PUBLISHED"
	ConsentStatusUnpublish = "UNPUBLISHED"
)

const (
	IdTypeNID      = "NID"
	IdTypePassport = "PASSPORT"
)

const (
	KYCTypeFace = "FACE"
	KYCTypeID   = "ID"

	KYCStatusPending  = "PEND"
	KYCStatusApproved = "APPR"
	KYCStatusRejected = "REJE"
)

const (
	BuySide  = "BUY"
	SellSide = "SELL"
)

const (
	CurrencyUSD = "USD"
	CurrencyKHR = "KHR"
)

const (
	TransactionTypeTransfer   = "TRANSFER"
	TransactionTypeDeposit    = "DEPOSIT"
	TransactionTypeWithdrawal = "WITHDRAWAL"
	TransactionTypeBill       = "BILL"
)

const (
	TransactionSubTypeFundIn  = "FUND_IN"
	TransactionSubTypeFundOut = "FUND_OUT"
)

const (
	WithdrawalTypeSelf  = "SELF"
	WithdrawalTypeOther = "OTHER"
)

const (
	MinimumTransferUSD = "0.01"
	MinimumTransferKHR = "100"
)

const (
	FeeTypeFixedRate   = "FIXE"
	FeeTypePercentRate = "PERC"
	FeeTypeTierRate    = "TIER"
)

const (
	StoreKeyAccessToken    = "access_token"
	StoreKeyRefreshToken   = "refresh_token"
	StoreKeyUserPermission = "user_permission"
)

const (
	PaymentStatusSuccess = "SUCCESS"
	PaymentStatusFail    = "FAIL"
	PaymentStatusPending = "PENDING"
)

const HeaderUserID = "PTC-User-ID"

const (
	SystemDefaultUserId = "1"
	SystemDefaultEmpId  = "0000000000"
)
