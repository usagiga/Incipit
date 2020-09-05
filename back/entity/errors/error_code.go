package errors

import "fmt"

type PrimaryErrorCode int

func (c *PrimaryErrorCode) String() string {
	return fmt.Sprintf("%04d", c)
}

const (
	// Validation
	AdminUserValidation PrimaryErrorCode = 0101
	LinkValidation      PrimaryErrorCode = 0102

	// Model
	AdminModel     PrimaryErrorCode = 0201
	AdminAuthModel PrimaryErrorCode = 0202
	HashModel      PrimaryErrorCode = 0203
	InstallerModel PrimaryErrorCode = 0204
	LinkModel      PrimaryErrorCode = 0205

	// Handler
	AdminUserHandler PrimaryErrorCode = 0301
	AdminAuthHandler PrimaryErrorCode = 0302
	LinkHandler      PrimaryErrorCode = 0303
	InstallerHandler      PrimaryErrorCode = 0304
)

type SecondaryErrorCode int

func (c *SecondaryErrorCode) String() string {
	return fmt.Sprintf("%04d", c)
}

const (
	// AdminUserValidation
	AdminUserValidation_NameIsTooShort         SecondaryErrorCode = 0101
	AdminUserValidation_NameIsTooLong          SecondaryErrorCode = 0102
	AdminUserValidation_NameHasUnavailableChar SecondaryErrorCode = 0103

	AdminUserValidation_ScreenNameIsTooShort         SecondaryErrorCode = 0201
	AdminUserValidation_ScreenNameIsTooLong          SecondaryErrorCode = 0202
	AdminUserValidation_ScreenNameHasUnavailableChar SecondaryErrorCode = 0203

	AdminUserValidation_PasswordIsTooShort         SecondaryErrorCode = 0301
	AdminUserValidation_PasswordIsTooLong          SecondaryErrorCode = 0302
	AdminUserValidation_PasswordHasUnavailableChar SecondaryErrorCode = 0303

	// LinkValidation
	LinkValidation_URLIsIncipit SecondaryErrorCode = 0101
	LinkValidation_URLIsInvalid SecondaryErrorCode = 0102

	// AdminModel
	AdminModel_FailedAdd            SecondaryErrorCode = 0101
	AdminModel_FailedFind           SecondaryErrorCode = 0201
	AdminModel_FindingUserNotFound  SecondaryErrorCode = 0202
	AdminModel_FailedUpdate         SecondaryErrorCode = 0301
	AdminModel_UpdatingUserNotFound SecondaryErrorCode = 0302
	AdminModel_FailedDelete         SecondaryErrorCode = 0401

	// AdminAuthModel
	AdminAuthModel_FailedToFindUser   SecondaryErrorCode = 0101
	AdminAuthModel_UnmatchPassword    SecondaryErrorCode = 0102
	AdminAuthModel_ExpiredToken       SecondaryErrorCode = 0103
	AdminAuthModel_FailedToStoreToken SecondaryErrorCode = 0104

	// HashModel
	HashModel_FailedGenerate SecondaryErrorCode = 0101

	HashModel_FailedCompare SecondaryErrorCode = 0201

	// InstallerModel
	// None

	// LinkModel
	LinkModel_FailedAdd            SecondaryErrorCode = 0101
	LinkModel_FailedFind           SecondaryErrorCode = 0201
	LinkModel_FindingLinkNotFound  SecondaryErrorCode = 0202
	LinkModel_FailedUpdate         SecondaryErrorCode = 0301
	LinkModel_UpdatingLinkNotFound SecondaryErrorCode = 0302
	LinkModel_FailedDelete         SecondaryErrorCode = 0401

	// AdminUserHandler
	AdminUserHandler_FailedBindJson SecondaryErrorCode = 0101

	// AdminAuthHandler
	AdminAuthHandler_FailedBindJson SecondaryErrorCode = 0101

	// LinkHandler
	LinkHandler_FailedBindQuery SecondaryErrorCode = 0101

	// InstallHandler
	InstallHandler_FailedBindJson SecondaryErrorCode = 0101
)
