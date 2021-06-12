module github.com/webtech-fmi/phonebook/backend/go/profile-service

go 1.16

replace (
	github.com/webtech-fmi/phonebook/backend/go/domain => ../domain
	github.com/webtech-fmi/phonebook/backend/go/infrastructure => ../infrastructure
)

require (
	github.com/go-ozzo/ozzo-dbx v1.5.0
	github.com/go-ozzo/ozzo-routing v2.1.4+incompatible
	github.com/go-ozzo/ozzo-validation/v4 v4.3.0
	github.com/golang/gddo v0.0.0-20210115222349-20d68f94ee1f // indirect
	github.com/google/uuid v1.2.0
	github.com/webtech-fmi/phonebook/backend/go/domain v0.0.0-00010101000000-000000000000
	github.com/webtech-fmi/phonebook/backend/go/infrastructure v0.0.0-00010101000000-000000000000
	github.com/spf13/cobra v1.1.3
	github.com/spf13/viper v1.7.1
)
