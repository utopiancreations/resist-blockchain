package types

import "cosmossdk.io/collections"

// VoteKey is the prefix to retrieve all Vote
var VoteKey = collections.NewPrefix("vote/value/")
