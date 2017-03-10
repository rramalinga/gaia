package zackmodels

import "github.com/aporeto-inc/elemental"

func init() {

	elemental.RegisterIdentity(ReportIdentity)
	elemental.RegisterIdentity(RootIdentity)
}

// ModelVersion returns the current version of the model
func ModelVersion() float64 { return 1.0 }

// IdentifiableForIdentity returns a new instance of the Identifiable for the given identity name.
func IdentifiableForIdentity(identity string) elemental.Identifiable {

	switch identity {
	case ReportIdentity.Name:
		return NewReport()
	case RootIdentity.Name:
		return NewRoot()
	default:
		return nil
	}
}

// ContentIdentifiableForIdentity returns a new instance of a ContentIdentifiable for the given identity name.
func ContentIdentifiableForIdentity(identity string) elemental.ContentIdentifiable {

	switch identity {
	case ReportIdentity.Name:
		return &ReportsList{}
	default:
		return nil
	}
}

// AllIdentities returns all existing identities.
func AllIdentities() []elemental.Identity {

	return []elemental.Identity{
		ReportIdentity,
		RootIdentity,
	}
}

var aliasesMap = map[string]elemental.Identity{}

// IdentityFromAlias returns the Identity associated to the given alias.
func IdentityFromAlias(alias string) elemental.Identity {

	return aliasesMap[alias]
}

// AliasesForIdentity returns all the aliases for the given identity.
func AliasesForIdentity(identity elemental.Identity) []string {

	switch identity {
	case ReportIdentity:
		return []string{}
	case RootIdentity:
		return []string{}
	}

	return nil
}