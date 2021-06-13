package vocabulary

type LockReason string

const (
	ResetPassword = LockReason("reset_password")
	VerifyEmail   = LockReason("verify_email")
)

type CredentialsType string

const (
	CredentialsPassword = CredentialsType("password")
	CredentialsLock     = CredentialsType("lock")
)
