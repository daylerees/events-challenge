# Projection Challenge

UW uses event-sourcing to record changes that happen to its customers and their accounts. Normally, we use read events from Apache Kafka, but instead, this test contains a simplified (in-memory) version of an event sequence which still presents a practical challenge for what UW backend engineers may face in their work.

## Context

The file `events/types.go` contains definitions for a number of events that describe account activity. For example:

- `AccountCreated` - A new user account has been created.
- `AccountInformationUpdated` - An update has been made to the username or/and email address for an account. Empty values should not be changed.
- `AccountBadgeEarned` - An account has earned a coloured badge. An account can earn more than one badge of the same colour.
- `AccountBadgeRevoked` - A coloured badge has been revoked for a given user account.

Calling the `GetEvents()` function from the events package will retrieve a time-sequenced slice of account events.

```go
evts := events.GetEvents()
```

You can assume that these events are ordered by the time they were raised. So the last event in the slice is the most recently emmited event.

AccountIDs (UUID) are immutable and cannot be updated with the `AccountInformationUpdated` event.

An account 'status' can be calculated for a given account according to the number of coloured badges they have been granted. See the list below for more information.

- **Great** - Default status.
- **Amazing** - 3+ Blue Badges.
- **Ultimate** - 6+ Blue Badges and 3+ Red Badges.
- **Champion** - 10+ Blue 5+ Red 1+ Green.

## Challenges

1. Create an in-memory projection for accounts, including account ID, name, email, number of badges earned, and the 'status' of the account. There is no need to use an external data storage for this challenge. Consider writing some tests to prove the correctness of your projection.
2. Write a simple API to retrieve account information (id, name, email, badge count and account type) as JSON by account ID.
