package vocabulary

type LockReason string

const (
	ResetPassword = LockReason("reset_password")
	VerifyEmail   = LockReason("verify_email")
)
