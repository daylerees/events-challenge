package events

// ===========================================
// Please don't edit this file.
// That would be cheating. :)
// ===========================================

// GetEvents that describe the time-series state of accounts in our system.
func GetEvents() []interface{} {
	evts := make([]interface{}, 0)

	const user1 = "730f2a64-ad93-4f84-b90f-349591066be7"
	const user2 = "04957cb1-7a75-44af-95e2-1214b90a9088"
	const user3 = "3f0920c6-6354-4343-8148-1e9aa89ec711"

	evts = append(evts, AccountCreated{
		AccountID: user1,
	})
	evts = append(evts, AccountCreated{
		AccountID: user2,
	})
	evts = append(evts, AccountCreated{
		AccountID: user3,
	})
	evts = append(evts, AccountEarnedBadge{
		AccountID:   user1,
		BadgeColour: BlueBadge,
	})
	evts = append(evts, AccountEarnedBadge{
		AccountID:   user1,
		BadgeColour: BlueBadge,
	})
	evts = append(evts, AccountEarnedBadge{
		AccountID:   user1,
		BadgeColour: BlueBadge,
	})
	evts = append(evts, AccountEarnedBadge{
		AccountID:   user2,
		BadgeColour: BlueBadge,
	})
	evts = append(evts, AccountInformationUpdated{
		AccountID:   user2,
		AccountName: "Ryanne Gold",
		AccountEmail: "r.gold98@yahoo.co.uk",
	})
	evts = append(evts, AccountEarnedBadge{
		AccountID:   user2,
		BadgeColour: BlueBadge,
	})
	evts = append(evts, AccountEarnedBadge{
		AccountID:   user2,
		BadgeColour: RedBadge,
	})
	evts = append(evts, AccountEarnedBadge{
		AccountID:   user1,
		BadgeColour: BlueBadge,
	})
	evts = append(evts, AccountInformationUpdated{
		AccountID:   user1,
		AccountName: "Amy Smith",
		AccountEmail: "asmith@gmail.com",
	})
	evts = append(evts, AccountEarnedBadge{
		AccountID:   user1,
		BadgeColour: GreenBadge,
	})
	evts = append(evts, AccountBadgeRevoked{
		AccountID:   user1,
		BadgeColour: BlueBadge,
	})
	evts = append(evts, AccountEarnedBadge{
		AccountID:   user2,
		BadgeColour: BlueBadge,
	})
	evts = append(evts, AccountEarnedBadge{
		AccountID:   user2,
		BadgeColour: BlueBadge,
	})
	evts = append(evts, AccountEarnedBadge{
		AccountID:   user1,
		BadgeColour: RedBadge,
	})
	evts = append(evts, AccountEarnedBadge{
		AccountID:   user1,
		BadgeColour: GreenBadge,
	})
	evts = append(evts, AccountEarnedBadge{
		AccountID:   user2,
		BadgeColour: BlueBadge,
	})
	evts = append(evts, AccountEarnedBadge{
		AccountID:   user2,
		BadgeColour: BlueBadge,
	})
	evts = append(evts, AccountEarnedBadge{
		AccountID:   user2,
		BadgeColour: RedBadge,
	})
	evts = append(evts, AccountEarnedBadge{
		AccountID:   user2,
		BadgeColour: RedBadge,
	})
	evts = append(evts, AccountEarnedBadge{
		AccountID:   user2,
		BadgeColour: RedBadge,
	})
	evts = append(evts, AccountBadgeRevoked{
		AccountID:   user2,
		BadgeColour: RedBadge,
	})
	evts = append(evts, AccountEarnedBadge{
		AccountID:   user2,
		BadgeColour: RedBadge,
	})
	evts = append(evts, AccountEarnedBadge{
		AccountID:   user3,
		BadgeColour: BlueBadge,
	})
	evts = append(evts, AccountEarnedBadge{
		AccountID:   user3,
		BadgeColour: BlueBadge,
	})
	evts = append(evts, AccountEarnedBadge{
		AccountID:   user3,
		BadgeColour: BlueBadge,
	})
	evts = append(evts, AccountEarnedBadge{
		AccountID:   user3,
		BadgeColour: BlueBadge,
	})
	evts = append(evts, AccountEarnedBadge{
		AccountID:   user3,
		BadgeColour: BlueBadge,
	})
	evts = append(evts, AccountEarnedBadge{
		AccountID:   user3,
		BadgeColour: BlueBadge,
	})
	evts = append(evts, AccountInformationUpdated{
		AccountID:   user3,
		AccountName: "Sam Lang",
		AccountEmail: "samalang@hey.com",
	})
	evts = append(evts, AccountEarnedBadge{
		AccountID:   user3,
		BadgeColour: BlueBadge,
	})
	evts = append(evts, AccountEarnedBadge{
		AccountID:   user3,
		BadgeColour: BlueBadge,
	})
	evts = append(evts, AccountInformationUpdated{
		AccountID:   user3,
		AccountEmail: "samlang@hey.com",
	})
	evts = append(evts, AccountEarnedBadge{
		AccountID:   user3,
		BadgeColour: BlueBadge,
	})
	evts = append(evts, AccountEarnedBadge{
		AccountID:   user3,
		BadgeColour: BlueBadge,
	})
	evts = append(evts, AccountEarnedBadge{
		AccountID:   user3,
		BadgeColour: RedBadge,
	})
	evts = append(evts, AccountEarnedBadge{
		AccountID:   user3,
		BadgeColour: RedBadge,
	})
	evts = append(evts, AccountEarnedBadge{
		AccountID:   user3,
		BadgeColour: RedBadge,
	})
	evts = append(evts, AccountEarnedBadge{
		AccountID:   user3,
		BadgeColour: RedBadge,
	})
	evts = append(evts, AccountEarnedBadge{
		AccountID:   user3,
		BadgeColour: RedBadge,
	})
	evts = append(evts, AccountEarnedBadge{
		AccountID:   user3,
		BadgeColour: GreenBadge,
	})
	evts = append(evts, AccountInformationUpdated{
		AccountID:   user2,
		AccountName: "Ryan Gold",
		AccountEmail: "r.gold98@yahoo.co.uk",
	})
	return evts
}
