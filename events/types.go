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

// Badge represents a coloured badge
type Badge int

// A collection of constants used to identify coloured badges.
const (
	BlueBadge Badge = iota
	RedBadge
	GreenBadge
)

// AccountBadgeEarned indicates that an account has earned a badge of a given colour.
// Accounts can earn more than one badge of the same colour.
type AccountBadgeEarned struct {
	AccountID   string
	BadgeColour Badge
}

// AccountBadgeRevoked indicates that a badge has been revoked for a given account.
type AccountBadgeRevoked struct {
	AccountID   string
	BadgeColour Badge
}
