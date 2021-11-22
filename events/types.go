package events

// AccountCreated indicates that a new user account has been created.
type AccountCreated struct {
	AccountID string
}

// AccountInformationUpdated details updates to account information.
type AccountInformationUpdated struct {
	AccountID    string
	AccountName  string
	AccountEmail string
}

// A collection of constants used to identify coloured badges.
const (
	BlueBadge = iota
	RedBadge
	GreenBadge
)

// AccountEarnedBadge indicates that an account has earned a badge of a given colour.
// Accounts can earn more than one badge of the same colour.
type AccountEarnedBadge struct {
	AccountID   string
	BadgeColour int
}

// AccountBadgeRevoked indicates that a badge has been revoked for a given account.
type AccountBadgeRevoked struct {
	AccountID   string
	BadgeColour int
}