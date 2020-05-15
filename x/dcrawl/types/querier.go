package types

import "strings"
// Query endpoints supported by the dcrawl querier
const (
  QueryListScavenges = "list"
	QueryGetScavenge   = "get"
	QueryCommit        = "commit"
)



// QueryResList Queries Result Payload for a query
type QueryResScavenges []string

// implement fmt.Stringer
func (n QueryResScavenges) String() string {
	return strings.Join(n[:], "\n")
}
