package errors

import "fmt"

type PrimaryErrorCode int

func (c *PrimaryErrorCode) String() string {
	return fmt.Sprintf("%04d", c)
}

const (
	// Validation
	AdminUserValidation PrimaryErrorCode = 101
	LinkValidation      PrimaryErrorCode = 102

	// Model
	AdminModel     PrimaryErrorCode = 201
	AdminAuthModel PrimaryErrorCode = 202
	HashModel      PrimaryErrorCode = 203
	InstallerModel PrimaryErrorCode = 204
	LinkModel      PrimaryErrorCode = 205

	// Handler
	AdminUserHandler PrimaryErrorCode = 301
	AdminAuthHandler PrimaryErrorCode = 302
	LinkHandler      PrimaryErrorCode = 303
	InstallerHandler      PrimaryErrorCode = 304
)

type SecondaryErrorCode int

func (c *SecondaryErrorCode) String() string {
	return fmt.Sprintf("%04d", c)
}

const (
	// AdminUserValidation
	AdminUserValidation_NameIsTooShort         SecondaryErrorCode = 101
	AdminUserValidation_NameIsTooLong          SecondaryErrorCode = 102
	AdminUserValidation_NameHasUnavailableChar SecondaryErrorCode = 103

	AdminUserValidation_ScreenNameIsTooShort         SecondaryErrorCode = 201
	AdminUserValidation_ScreenNameIsTooLong          SecondaryErrorCode = 202
	AdminUserValidation_ScreenNameHasUnavailableChar SecondaryErrorCode = 203

	AdminUserValidation_PasswordIsTooShort         SecondaryErrorCode = 301
	AdminUserValidation_PasswordIsTooLong          SecondaryErrorCode = 302
	AdminUserValidation_PasswordHasUnavailableChar SecondaryErrorCode = 303

	// LinkValidation
	LinkValidation_URLIsIncipit SecondaryErrorCode = 101
	LinkValidation_URLIsInvalid SecondaryErrorCode = 102

	// AdminModel
	AdminModel_FailedAdd            SecondaryErrorCode = 101
	AdminModel_FailedFind           SecondaryErrorCode = 201
	AdminModel_FindingUserNotFound  SecondaryErrorCode = 202
	AdminModel_FailedUpdate         SecondaryErrorCode = 301
	AdminModel_UpdatingUserNotFound SecondaryErrorCode = 302
	AdminModel_FailedDelete         SecondaryErrorCode = 401

	// AdminAuthModel
	AdminAuthModel_FailedToFindUser   SecondaryErrorCode = 101
	AdminAuthModel_UnmatchPassword    SecondaryErrorCode = 102
	AdminAuthModel_ExpiredToken       SecondaryErrorCode = 103
	AdminAuthModel_FailedToStoreToken SecondaryErrorCode = 104

	// HashModel
	HashModel_FailedGenerate SecondaryErrorCode = 101

	HashModel_FailedCompare SecondaryErrorCode = 201

	// InstallerModel
	// None

	// LinkModel
	LinkModel_FailedAdd            SecondaryErrorCode = 101
	LinkModel_FailedFind           SecondaryErrorCode = 201
	LinkModel_FindingLinkNotFound  SecondaryErrorCode = 202
	LinkModel_FailedUpdate         SecondaryErrorCode = 301
	LinkModel_UpdatingLinkNotFound SecondaryErrorCode = 302
	LinkModel_FailedDelete         SecondaryErrorCode = 401

	// AdminUserHandler
	AdminUserHandler_FailedBindJson SecondaryErrorCode = 101

	// AdminAuthHandler
	AdminAuthHandler_FailedBindJson SecondaryErrorCode = 101

	// LinkHandler
	LinkHandler_FailedBindQuery SecondaryErrorCode = 101

	// InstallHandler
	InstallHandler_FailedBindJson SecondaryErrorCode = 101
)
