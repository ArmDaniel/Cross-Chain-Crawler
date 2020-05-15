package types

const (
	// ModuleName is the name of the module
	ModuleName = "dcrawl"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName

  // RouterKey to be used for routing msgs
  RouterKey = ModuleName

  // QuerierRoute to be used for querierer msgs
	QuerierRoute = ModuleName
  //prefixes added that act like unique identifiers for the types of hashes used
	ScavengePrefix = "sk-"
	CommitPrefix   = "ck-"
)
