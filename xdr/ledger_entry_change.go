package xdr

import "fmt"

// EntryType is a helper to get at the entry type for a change.
func (change *LedgerEntryChange) EntryType() LedgerEntryType {
	return change.LedgerKey().Type
}

// LedgerKey returns the key for the ledger entry that was changed
// in `change`.
func (change *LedgerEntryChange) LedgerKey() LedgerKey {
	switch change.Type {
	case LedgerEntryChangeTypeCreated:
		change := change.MustCreated()
		return change.LedgerKey()
	case LedgerEntryChangeTypeRemoved:
		return change.MustRemoved()
	case LedgerEntryChangeTypeUpdated:
		change := change.MustUpdated()
		return change.LedgerKey()
	case LedgerEntryChangeTypeState:
		change := change.MustState()
		return change.LedgerKey()
	default:
		panic(fmt.Errorf("Unknown change type: %v", change.Type))
	}
}
