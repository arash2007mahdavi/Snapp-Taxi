package loggers

type Category string
type SubCategory string
type ExtraKey string

const (
	General         Category = "General"
	Internal        Category = "Internal"
	Postgres        Category = "Postgres"
	Redis           Category = "Redis"
	Validation      Category = "Validation"
	RequestResponse Category = "RequestResponse"
	Admin           Category = "Admin"
	User            Category = "User"
	Driver          Category = "Driver"
)

const (
	// General
	Startup         SubCategory = "Startup"
	ExternalService SubCategory = "ExternalService"

	// Postgres
	Select   SubCategory = "Select"
	Rollback SubCategory = "Rollback"
	Update   SubCategory = "Update"
	Delete   SubCategory = "Delete"
	Insert   SubCategory = "Insert"

	// Internal
	Api                 SubCategory = "Api"
	HashPassword        SubCategory = "HashPassword"
	DefaultRoleNotFound SubCategory = "DefaultRoleNotFound"

	// Validation
	MobileValidation   SubCategory = "MobileValidation"
	PasswordValidation SubCategory = "PasswordValidation"

	// Admin
	Add       SubCategory = "Add"
	Check     SubCategory = "Check"
	NewAdmin  SubCategory = "New Admin"
	SeeUsers  SubCategory = "See Users"
	SeeAdmins SubCategory = "See Admins"

	// User
	NewUser SubCategory = "New User"
)

const (
	AppName      ExtraKey = "AppName"
	LoggerName   ExtraKey = "LoggerName"
	ClientIp     ExtraKey = "ClientIp"
	HostIp       ExtraKey = "HostIp"
	Method       ExtraKey = "Method"
	StatusCode   ExtraKey = "StatusCode"
	BodySize     ExtraKey = "BodySize"
	Path         ExtraKey = "Path"
	Latency      ExtraKey = "Latency"
	Body         ExtraKey = "Body"
	ErrorMessage ExtraKey = "ErrorMessage"
	UserID       ExtraKey = "User ID"
	AdminID      ExtraKey = "Admin ID"
	Fullname     ExtraKey = "Fullname"
	Password     ExtraKey = "Password"
	Wallet       ExtraKey = "Wallet"
)
