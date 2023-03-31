package events

// ===========================================
// Please don't edit this file.
// That would be cheating. :)
// ===========================================

// GetEvents that describe the time-series state of accounts in our system.
func GetEvents() []interface{} {
	const (
		user1 = "730f2a64-ad93-4f84-b90f-349591066be7"
		user2 = "04957cb1-7a75-44af-95e2-1214b90a9088"
		user3 = "3f0920c6-6354-4343-8148-1e9aa89ec711"
	)

	return []interface{}{
		AccountCreated{
			AccountID: user1,
		},
		AccountCreated{
			AccountID: user2,
		},
		AccountCreated{
			AccountID: user3,
		},
		AccountBadgeEarned{
			AccountID:   user1,
			BadgeColour: BlueBadge,
		},
		AccountBadgeEarned{
			AccountID:   user1,
			BadgeColour: BlueBadge,
		},
		AccountBadgeEarned{
			AccountID:   user1,
			BadgeColour: BlueBadge,
		},
		AccountBadgeEarned{
			AccountID:   user2,
			BadgeColour: BlueBadge,
		},
		AccountInformationUpdated{
			AccountID:    user2,
			AccountName:  "Ryanne Gold",
			AccountEmail: "r.gold98@yahoo.co.uk",
		},
		AccountBadgeEarned{
			AccountID:   user2,
			BadgeColour: BlueBadge,
		},
		AccountBadgeEarned{
			AccountID:   user2,
			BadgeColour: RedBadge,
		},
		AccountBadgeEarned{
			AccountID:   user1,
			BadgeColour: BlueBadge,
		},
		AccountInformationUpdated{
			AccountID:    user1,
			AccountName:  "Amy Smith",
			AccountEmail: "asmith@gmail.com",
		},
		AccountBadgeEarned{
			AccountID:   user1,
			BadgeColour: GreenBadge,
		},
		AccountBadgeRevoked{
			AccountID:   user1,
			BadgeColour: BlueBadge,
		},
		AccountBadgeEarned{
			AccountID:   user2,
			BadgeColour: BlueBadge,
		},
		AccountBadgeEarned{
			AccountID:   user2,
			BadgeColour: BlueBadge,
		},
		AccountBadgeEarned{
			AccountID:   user1,
			BadgeColour: RedBadge,
		},
		AccountBadgeEarned{
			AccountID:   user1,
			BadgeColour: GreenBadge,
		},
		AccountBadgeEarned{
			AccountID:   user2,
			BadgeColour: BlueBadge,
		},
		AccountBadgeEarned{
			AccountID:   user2,
			BadgeColour: BlueBadge,
		},
		AccountBadgeEarned{
			AccountID:   user2,
			BadgeColour: RedBadge,
		},
		AccountBadgeEarned{
			AccountID:   user2,
			BadgeColour: RedBadge,
		},
		AccountBadgeEarned{
			AccountID:   user2,
			BadgeColour: RedBadge,
		},
		AccountBadgeRevoked{
			AccountID:   user2,
			BadgeColour: RedBadge,
		},
		AccountBadgeEarned{
			AccountID:   user2,
			BadgeColour: RedBadge,
		},
		AccountBadgeEarned{
			AccountID:   user3,
			BadgeColour: BlueBadge,
		},
		AccountBadgeEarned{
			AccountID:   user3,
			BadgeColour: BlueBadge,
		},
		AccountBadgeEarned{
			AccountID:   user3,
			BadgeColour: BlueBadge,
		},
		AccountBadgeEarned{
			AccountID:   user3,
			BadgeColour: BlueBadge,
		},
		AccountBadgeEarned{
			AccountID:   user3,
			BadgeColour: BlueBadge,
		},
		AccountBadgeEarned{
			AccountID:   user3,
			BadgeColour: BlueBadge,
		},
		AccountInformationUpdated{
			AccountID:    user3,
			AccountName:  "Sam Lang",
			AccountEmail: "samalang@hey.com",
		},
		AccountBadgeEarned{
			AccountID:   user3,
			BadgeColour: BlueBadge,
		},
		AccountBadgeEarned{
			AccountID:   user3,
			BadgeColour: BlueBadge,
		},
		AccountInformationUpdated{
			AccountID:    user3,
			AccountEmail: "samlang@hey.com",
		},
		AccountBadgeEarned{
			AccountID:   user3,
			BadgeColour: BlueBadge,
		},
		AccountBadgeEarned{
			AccountID:   user3,
			BadgeColour: BlueBadge,
		},
		AccountBadgeEarned{
			AccountID:   user3,
			BadgeColour: RedBadge,
		},
		AccountBadgeEarned{
			AccountID:   user3,
			BadgeColour: RedBadge,
		},
		AccountBadgeEarned{
			AccountID:   user3,
			BadgeColour: RedBadge,
		},
		AccountBadgeEarned{
			AccountID:   user3,
			BadgeColour: RedBadge,
		},
		AccountBadgeEarned{
			AccountID:   user3,
			BadgeColour: RedBadge,
		},
		AccountBadgeEarned{
			AccountID:   user3,
			BadgeColour: GreenBadge,
		},
		AccountInformationUpdated{
			AccountID:    user2,
			AccountName:  "Ryan Gold",
			AccountEmail: "r.gold98@yahoo.co.uk",
		},
	}
}
