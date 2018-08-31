// Package xdr is generated from:
//
//  xdr/raw/Stellar-SCP.x
//  xdr/raw/Stellar-ledger-entries-account-KYC.x
//  xdr/raw/Stellar-ledger-entries-account-limits.x
//  xdr/raw/Stellar-ledger-entries-account-type-limits.x
//  xdr/raw/Stellar-ledger-entries-account.x
//  xdr/raw/Stellar-ledger-entries-asset-pair.x
//  xdr/raw/Stellar-ledger-entries-asset.x
//  xdr/raw/Stellar-ledger-entries-balance.x
//  xdr/raw/Stellar-ledger-entries-contract.x
//  xdr/raw/Stellar-ledger-entries-external-system-id-pool-entry.x
//  xdr/raw/Stellar-ledger-entries-external-system-id.x
//  xdr/raw/Stellar-ledger-entries-fee.x
//  xdr/raw/Stellar-ledger-entries-key-value.x
//  xdr/raw/Stellar-ledger-entries-limits-v2.x
//  xdr/raw/Stellar-ledger-entries-offer.x
//  xdr/raw/Stellar-ledger-entries-pending-statistics.x
//  xdr/raw/Stellar-ledger-entries-reference.x
//  xdr/raw/Stellar-ledger-entries-reviewable-request.x
//  xdr/raw/Stellar-ledger-entries-sale-ante.x
//  xdr/raw/Stellar-ledger-entries-sale.x
//  xdr/raw/Stellar-ledger-entries-statistics-v2.x
//  xdr/raw/Stellar-ledger-entries-statistics.x
//  xdr/raw/Stellar-ledger-entries.x
//  xdr/raw/Stellar-ledger.x
//  xdr/raw/Stellar-operation-bind-external-system-id.x
//  xdr/raw/Stellar-operation-check-sale-state.x
//  xdr/raw/Stellar-operation-create-AML-alert-request.x
//  xdr/raw/Stellar-operation-create-KYC-request.x
//  xdr/raw/Stellar-operation-create-account.x
//  xdr/raw/Stellar-operation-create-issuance-request.x
//  xdr/raw/Stellar-operation-create-manage-limits-request.x
//  xdr/raw/Stellar-operation-create-preissuance-request.x
//  xdr/raw/Stellar-operation-create-sale-creation-request.x
//  xdr/raw/Stellar-operation-create-withdrawal-request.x
//  xdr/raw/Stellar-operation-direct-debit.x
//  xdr/raw/Stellar-operation-manage-account.x
//  xdr/raw/Stellar-operation-manage-asset-pair.x
//  xdr/raw/Stellar-operation-manage-asset.x
//  xdr/raw/Stellar-operation-manage-balance.x
//  xdr/raw/Stellar-operation-manage-contract-request.x
//  xdr/raw/Stellar-operation-manage-contract.x
//  xdr/raw/Stellar-operation-manage-external-system-id-pool-entry.x
//  xdr/raw/Stellar-operation-manage-invoice-request.x
//  xdr/raw/Stellar-operation-manage-key-value.x
//  xdr/raw/Stellar-operation-manage-limits.x
//  xdr/raw/Stellar-operation-manage-offer.x
//  xdr/raw/Stellar-operation-manage-sale.x
//  xdr/raw/Stellar-operation-payment-v2.x
//  xdr/raw/Stellar-operation-payment.x
//  xdr/raw/Stellar-operation-review-request.x
//  xdr/raw/Stellar-operation-set-fees.x
//  xdr/raw/Stellar-operation-set-options.x
//  xdr/raw/Stellar-overlay.x
//  xdr/raw/Stellar-reviewable-request-AML-alert.x
//  xdr/raw/Stellar-reviewable-request-asset.x
//  xdr/raw/Stellar-reviewable-request-contract.x
//  xdr/raw/Stellar-reviewable-request-invoice.x
//  xdr/raw/Stellar-reviewable-request-issuance.x
//  xdr/raw/Stellar-reviewable-request-limits-update.x
//  xdr/raw/Stellar-reviewable-request-sale.x
//  xdr/raw/Stellar-reviewable-request-update-KYC.x
//  xdr/raw/Stellar-reviewable-request-update-promotion.x
//  xdr/raw/Stellar-reviewable-request-update-sale-details.x
//  xdr/raw/Stellar-reviewable-request-update-sale-end-time.x
//  xdr/raw/Stellar-reviewable-request-withdrawal.x
//  xdr/raw/Stellar-transaction.x
//  xdr/raw/Stellar-types.x
//
// DO NOT EDIT or your changes may be overwritten
package xdr

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/nullstyle/go-xdr/xdr3"
)

// Unmarshal reads an xdr element from `r` into `v`.
func Unmarshal(r io.Reader, v interface{}) (int, error) {
	// delegate to xdr package's Unmarshal
	return xdr.Unmarshal(r, v)
}

// Marshal writes an xdr element `v` into `w`.
func Marshal(w io.Writer, v interface{}) (int, error) {
	// delegate to xdr package's Marshal
	return xdr.Marshal(w, v)
}

// Value is an XDR Typedef defines as:
//
//   typedef opaque Value<>;
//
type Value []byte

// ScpBallot is an XDR Struct defines as:
//
//   struct SCPBallot
//    {
//        uint32 counter; // n
//        Value value;    // x
//    };
//
type ScpBallot struct {
	Counter Uint32 `json:"counter,omitempty"`
	Value   Value  `json:"value,omitempty"`
}

// ScpStatementType is an XDR Enum defines as:
//
//   enum SCPStatementType
//    {
//        PREPARE = 0,
//        CONFIRM = 1,
//        EXTERNALIZE = 2,
//        NOMINATE = 3
//    };
//
type ScpStatementType int32

const (
	ScpStatementTypePrepare     ScpStatementType = 0
	ScpStatementTypeConfirm     ScpStatementType = 1
	ScpStatementTypeExternalize ScpStatementType = 2
	ScpStatementTypeNominate    ScpStatementType = 3
)

var ScpStatementTypeAll = []ScpStatementType{
	ScpStatementTypePrepare,
	ScpStatementTypeConfirm,
	ScpStatementTypeExternalize,
	ScpStatementTypeNominate,
}

var scpStatementTypeMap = map[int32]string{
	0: "ScpStatementTypePrepare",
	1: "ScpStatementTypeConfirm",
	2: "ScpStatementTypeExternalize",
	3: "ScpStatementTypeNominate",
}

var scpStatementTypeShortMap = map[int32]string{
	0: "prepare",
	1: "confirm",
	2: "externalize",
	3: "nominate",
}

var scpStatementTypeRevMap = map[string]int32{
	"ScpStatementTypePrepare":     0,
	"ScpStatementTypeConfirm":     1,
	"ScpStatementTypeExternalize": 2,
	"ScpStatementTypeNominate":    3,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ScpStatementType
func (e ScpStatementType) ValidEnum(v int32) bool {
	_, ok := scpStatementTypeMap[v]
	return ok
}
func (e ScpStatementType) isFlag() bool {
	for i := len(ScpStatementTypeAll) - 1; i >= 0; i-- {
		expected := ScpStatementType(2) << uint64(len(ScpStatementTypeAll)-1) >> uint64(len(ScpStatementTypeAll)-i)
		if expected != ScpStatementTypeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ScpStatementType) String() string {
	name, _ := scpStatementTypeMap[int32(e)]
	return name
}

func (e ScpStatementType) ShortString() string {
	name, _ := scpStatementTypeShortMap[int32(e)]
	return name
}

func (e ScpStatementType) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range ScpStatementTypeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ScpStatementType) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ScpStatementType(t.Value)
	return nil
}

// ScpNomination is an XDR Struct defines as:
//
//   struct SCPNomination
//    {
//        Hash quorumSetHash; // D
//        Value votes<>;      // X
//        Value accepted<>;   // Y
//    };
//
type ScpNomination struct {
	QuorumSetHash Hash    `json:"quorumSetHash,omitempty"`
	Votes         []Value `json:"votes,omitempty"`
	Accepted      []Value `json:"accepted,omitempty"`
}

// ScpStatementPrepare is an XDR NestedStruct defines as:
//
//   struct
//            {
//                Hash quorumSetHash;       // D
//                SCPBallot ballot;         // b
//                SCPBallot* prepared;      // p
//                SCPBallot* preparedPrime; // p'
//                uint32 nC;                // c.n
//                uint32 nH;                // h.n
//            }
//
type ScpStatementPrepare struct {
	QuorumSetHash Hash       `json:"quorumSetHash,omitempty"`
	Ballot        ScpBallot  `json:"ballot,omitempty"`
	Prepared      *ScpBallot `json:"prepared,omitempty"`
	PreparedPrime *ScpBallot `json:"preparedPrime,omitempty"`
	NC            Uint32     `json:"nC,omitempty"`
	NH            Uint32     `json:"nH,omitempty"`
}

// ScpStatementConfirm is an XDR NestedStruct defines as:
//
//   struct
//            {
//                SCPBallot ballot;   // b
//                uint32 nPrepared;   // p.n
//                uint32 nCommit;     // c.n
//                uint32 nH;          // h.n
//                Hash quorumSetHash; // D
//            }
//
type ScpStatementConfirm struct {
	Ballot        ScpBallot `json:"ballot,omitempty"`
	NPrepared     Uint32    `json:"nPrepared,omitempty"`
	NCommit       Uint32    `json:"nCommit,omitempty"`
	NH            Uint32    `json:"nH,omitempty"`
	QuorumSetHash Hash      `json:"quorumSetHash,omitempty"`
}

// ScpStatementExternalize is an XDR NestedStruct defines as:
//
//   struct
//            {
//                SCPBallot commit;         // c
//                uint32 nH;                // h.n
//                Hash commitQuorumSetHash; // D used before EXTERNALIZE
//            }
//
type ScpStatementExternalize struct {
	Commit              ScpBallot `json:"commit,omitempty"`
	NH                  Uint32    `json:"nH,omitempty"`
	CommitQuorumSetHash Hash      `json:"commitQuorumSetHash,omitempty"`
}

// ScpStatementPledges is an XDR NestedUnion defines as:
//
//   union switch (SCPStatementType type)
//        {
//        case PREPARE:
//            struct
//            {
//                Hash quorumSetHash;       // D
//                SCPBallot ballot;         // b
//                SCPBallot* prepared;      // p
//                SCPBallot* preparedPrime; // p'
//                uint32 nC;                // c.n
//                uint32 nH;                // h.n
//            } prepare;
//        case CONFIRM:
//            struct
//            {
//                SCPBallot ballot;   // b
//                uint32 nPrepared;   // p.n
//                uint32 nCommit;     // c.n
//                uint32 nH;          // h.n
//                Hash quorumSetHash; // D
//            } confirm;
//        case EXTERNALIZE:
//            struct
//            {
//                SCPBallot commit;         // c
//                uint32 nH;                // h.n
//                Hash commitQuorumSetHash; // D used before EXTERNALIZE
//            } externalize;
//        case NOMINATE:
//            SCPNomination nominate;
//        }
//
type ScpStatementPledges struct {
	Type        ScpStatementType         `json:"type,omitempty"`
	Prepare     *ScpStatementPrepare     `json:"prepare,omitempty"`
	Confirm     *ScpStatementConfirm     `json:"confirm,omitempty"`
	Externalize *ScpStatementExternalize `json:"externalize,omitempty"`
	Nominate    *ScpNomination           `json:"nominate,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ScpStatementPledges) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ScpStatementPledges
func (u ScpStatementPledges) ArmForSwitch(sw int32) (string, bool) {
	switch ScpStatementType(sw) {
	case ScpStatementTypePrepare:
		return "Prepare", true
	case ScpStatementTypeConfirm:
		return "Confirm", true
	case ScpStatementTypeExternalize:
		return "Externalize", true
	case ScpStatementTypeNominate:
		return "Nominate", true
	}
	return "-", false
}

// NewScpStatementPledges creates a new  ScpStatementPledges.
func NewScpStatementPledges(aType ScpStatementType, value interface{}) (result ScpStatementPledges, err error) {
	result.Type = aType
	switch ScpStatementType(aType) {
	case ScpStatementTypePrepare:
		tv, ok := value.(ScpStatementPrepare)
		if !ok {
			err = fmt.Errorf("invalid value, must be ScpStatementPrepare")
			return
		}
		result.Prepare = &tv
	case ScpStatementTypeConfirm:
		tv, ok := value.(ScpStatementConfirm)
		if !ok {
			err = fmt.Errorf("invalid value, must be ScpStatementConfirm")
			return
		}
		result.Confirm = &tv
	case ScpStatementTypeExternalize:
		tv, ok := value.(ScpStatementExternalize)
		if !ok {
			err = fmt.Errorf("invalid value, must be ScpStatementExternalize")
			return
		}
		result.Externalize = &tv
	case ScpStatementTypeNominate:
		tv, ok := value.(ScpNomination)
		if !ok {
			err = fmt.Errorf("invalid value, must be ScpNomination")
			return
		}
		result.Nominate = &tv
	}
	return
}

// MustPrepare retrieves the Prepare value from the union,
// panicing if the value is not set.
func (u ScpStatementPledges) MustPrepare() ScpStatementPrepare {
	val, ok := u.GetPrepare()

	if !ok {
		panic("arm Prepare is not set")
	}

	return val
}

// GetPrepare retrieves the Prepare value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ScpStatementPledges) GetPrepare() (result ScpStatementPrepare, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Prepare" {
		result = *u.Prepare
		ok = true
	}

	return
}

// MustConfirm retrieves the Confirm value from the union,
// panicing if the value is not set.
func (u ScpStatementPledges) MustConfirm() ScpStatementConfirm {
	val, ok := u.GetConfirm()

	if !ok {
		panic("arm Confirm is not set")
	}

	return val
}

// GetConfirm retrieves the Confirm value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ScpStatementPledges) GetConfirm() (result ScpStatementConfirm, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Confirm" {
		result = *u.Confirm
		ok = true
	}

	return
}

// MustExternalize retrieves the Externalize value from the union,
// panicing if the value is not set.
func (u ScpStatementPledges) MustExternalize() ScpStatementExternalize {
	val, ok := u.GetExternalize()

	if !ok {
		panic("arm Externalize is not set")
	}

	return val
}

// GetExternalize retrieves the Externalize value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ScpStatementPledges) GetExternalize() (result ScpStatementExternalize, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Externalize" {
		result = *u.Externalize
		ok = true
	}

	return
}

// MustNominate retrieves the Nominate value from the union,
// panicing if the value is not set.
func (u ScpStatementPledges) MustNominate() ScpNomination {
	val, ok := u.GetNominate()

	if !ok {
		panic("arm Nominate is not set")
	}

	return val
}

// GetNominate retrieves the Nominate value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ScpStatementPledges) GetNominate() (result ScpNomination, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Nominate" {
		result = *u.Nominate
		ok = true
	}

	return
}

// ScpStatement is an XDR Struct defines as:
//
//   struct SCPStatement
//    {
//        NodeID nodeID;    // v
//        uint64 slotIndex; // i
//
//        union switch (SCPStatementType type)
//        {
//        case PREPARE:
//            struct
//            {
//                Hash quorumSetHash;       // D
//                SCPBallot ballot;         // b
//                SCPBallot* prepared;      // p
//                SCPBallot* preparedPrime; // p'
//                uint32 nC;                // c.n
//                uint32 nH;                // h.n
//            } prepare;
//        case CONFIRM:
//            struct
//            {
//                SCPBallot ballot;   // b
//                uint32 nPrepared;   // p.n
//                uint32 nCommit;     // c.n
//                uint32 nH;          // h.n
//                Hash quorumSetHash; // D
//            } confirm;
//        case EXTERNALIZE:
//            struct
//            {
//                SCPBallot commit;         // c
//                uint32 nH;                // h.n
//                Hash commitQuorumSetHash; // D used before EXTERNALIZE
//            } externalize;
//        case NOMINATE:
//            SCPNomination nominate;
//        }
//        pledges;
//    };
//
type ScpStatement struct {
	NodeId    NodeId              `json:"nodeID,omitempty"`
	SlotIndex Uint64              `json:"slotIndex,omitempty"`
	Pledges   ScpStatementPledges `json:"pledges,omitempty"`
}

// ScpEnvelope is an XDR Struct defines as:
//
//   struct SCPEnvelope
//    {
//        SCPStatement statement;
//        Signature signature;
//    };
//
type ScpEnvelope struct {
	Statement ScpStatement `json:"statement,omitempty"`
	Signature Signature    `json:"signature,omitempty"`
}

// ScpQuorumSet is an XDR Struct defines as:
//
//   struct SCPQuorumSet
//    {
//        uint32 threshold;
//        PublicKey validators<>;
//        SCPQuorumSet innerSets<>;
//    };
//
type ScpQuorumSet struct {
	Threshold  Uint32         `json:"threshold,omitempty"`
	Validators []PublicKey    `json:"validators,omitempty"`
	InnerSets  []ScpQuorumSet `json:"innerSets,omitempty"`
}

// AccountKycEntryExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type AccountKycEntryExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u AccountKycEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of AccountKycEntryExt
func (u AccountKycEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewAccountKycEntryExt creates a new  AccountKycEntryExt.
func NewAccountKycEntryExt(v LedgerVersion, value interface{}) (result AccountKycEntryExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// AccountKycEntry is an XDR Struct defines as:
//
//   struct AccountKYCEntry
//    {
//        AccountID accountID;
//        longstring KYCData;
//
//        // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type AccountKycEntry struct {
	AccountId AccountId          `json:"accountID,omitempty"`
	KycData   Longstring         `json:"KYCData,omitempty"`
	Ext       AccountKycEntryExt `json:"ext,omitempty"`
}

// AccountLimitsEntryExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type AccountLimitsEntryExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u AccountLimitsEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of AccountLimitsEntryExt
func (u AccountLimitsEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewAccountLimitsEntryExt creates a new  AccountLimitsEntryExt.
func NewAccountLimitsEntryExt(v LedgerVersion, value interface{}) (result AccountLimitsEntryExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// AccountLimitsEntry is an XDR Struct defines as:
//
//   struct AccountLimitsEntry
//    {
//        AccountID accountID;
//        Limits limits;
//
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type AccountLimitsEntry struct {
	AccountId AccountId             `json:"accountID,omitempty"`
	Limits    Limits                `json:"limits,omitempty"`
	Ext       AccountLimitsEntryExt `json:"ext,omitempty"`
}

// AccountTypeLimitsEntryExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type AccountTypeLimitsEntryExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u AccountTypeLimitsEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of AccountTypeLimitsEntryExt
func (u AccountTypeLimitsEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewAccountTypeLimitsEntryExt creates a new  AccountTypeLimitsEntryExt.
func NewAccountTypeLimitsEntryExt(v LedgerVersion, value interface{}) (result AccountTypeLimitsEntryExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// AccountTypeLimitsEntry is an XDR Struct defines as:
//
//   struct AccountTypeLimitsEntry
//    {
//    	AccountType accountType;
//        Limits limits;
//
//    	// reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type AccountTypeLimitsEntry struct {
	AccountType AccountType               `json:"accountType,omitempty"`
	Limits      Limits                    `json:"limits,omitempty"`
	Ext         AccountTypeLimitsEntryExt `json:"ext,omitempty"`
}

// SignerType is an XDR Enum defines as:
//
//   enum SignerType
//    {
//    	READER = 1,                  // can only read data from API and Horizon
//    	NOT_VERIFIED_ACC_MANAGER = 2,// can manage not verified account and block/unblock general
//    	GENERAL_ACC_MANAGER = 4,     // allowed to create account, block/unblock, change limits for particular general account
//    	DIRECT_DEBIT_OPERATOR = 8, // allowed to perform direct debit operation
//    	ASSET_MANAGER = 16, // allowed to create assets/asset pairs and update policies, set fees
//    	ASSET_RATE_MANAGER = 32, // allowed to set physical asset price
//    	BALANCE_MANAGER = 64, // allowed to create balances, spend assets from balances
//    	ISSUANCE_MANAGER = 128, // allowed to make preissuance request
//    	INVOICE_MANAGER = 256, // allowed to create payment requests to other accounts
//    	PAYMENT_OPERATOR = 512, // allowed to review payment requests
//    	LIMITS_MANAGER = 1024, // allowed to change limits
//    	ACCOUNT_MANAGER = 2048, // allowed to add/delete signers and trust
//    	COMMISSION_BALANCE_MANAGER  = 4096,// allowed to spend from commission balances
//    	OPERATIONAL_BALANCE_MANAGER = 8192, // allowed to spend from operational balances
//    	EVENTS_CHECKER = 16384, // allow to check and trigger events
//    	EXCHANGE_ACC_MANAGER = 32768, // can manage exchange account
//    	SYNDICATE_ACC_MANAGER = 65536, // can manage syndicate account
//    	USER_ASSET_MANAGER = 131072, // can review sale, asset creation/update requests
//    	USER_ISSUANCE_MANAGER = 262144, // can review pre-issuance/issuance requests
//    	WITHDRAW_MANAGER = 524288, // can review withdraw requests
//    	FEES_MANAGER = 1048576, // can set fee
//    	TX_SENDER = 2097152, // can send tx
//    	AML_ALERT_MANAGER = 4194304, // can manage AML alert request
//    	AML_ALERT_REVIEWER = 8388608, // can review aml alert requests
//    	KYC_ACC_MANAGER = 16777216, // can manage kyc
//    	KYC_SUPER_ADMIN = 33554432,
//    	EXTERNAL_SYSTEM_ACCOUNT_ID_POOL_MANAGER = 67108864,
//        KEY_VALUE_MANAGER = 134217728, // can manage keyValue
//        SUPER_ISSUANCE_MANAGER = 268435456,
//        CONTRACT_MANAGER = 536870912
//    };
//
type SignerType int32

const (
	SignerTypeReader                             SignerType = 1
	SignerTypeNotVerifiedAccManager              SignerType = 2
	SignerTypeGeneralAccManager                  SignerType = 4
	SignerTypeDirectDebitOperator                SignerType = 8
	SignerTypeAssetManager                       SignerType = 16
	SignerTypeAssetRateManager                   SignerType = 32
	SignerTypeBalanceManager                     SignerType = 64
	SignerTypeIssuanceManager                    SignerType = 128
	SignerTypeInvoiceManager                     SignerType = 256
	SignerTypePaymentOperator                    SignerType = 512
	SignerTypeLimitsManager                      SignerType = 1024
	SignerTypeAccountManager                     SignerType = 2048
	SignerTypeCommissionBalanceManager           SignerType = 4096
	SignerTypeOperationalBalanceManager          SignerType = 8192
	SignerTypeEventsChecker                      SignerType = 16384
	SignerTypeExchangeAccManager                 SignerType = 32768
	SignerTypeSyndicateAccManager                SignerType = 65536
	SignerTypeUserAssetManager                   SignerType = 131072
	SignerTypeUserIssuanceManager                SignerType = 262144
	SignerTypeWithdrawManager                    SignerType = 524288
	SignerTypeFeesManager                        SignerType = 1048576
	SignerTypeTxSender                           SignerType = 2097152
	SignerTypeAmlAlertManager                    SignerType = 4194304
	SignerTypeAmlAlertReviewer                   SignerType = 8388608
	SignerTypeKycAccManager                      SignerType = 16777216
	SignerTypeKycSuperAdmin                      SignerType = 33554432
	SignerTypeExternalSystemAccountIdPoolManager SignerType = 67108864
	SignerTypeKeyValueManager                    SignerType = 134217728
	SignerTypeSuperIssuanceManager               SignerType = 268435456
	SignerTypeContractManager                    SignerType = 536870912
)

var SignerTypeAll = []SignerType{
	SignerTypeReader,
	SignerTypeNotVerifiedAccManager,
	SignerTypeGeneralAccManager,
	SignerTypeDirectDebitOperator,
	SignerTypeAssetManager,
	SignerTypeAssetRateManager,
	SignerTypeBalanceManager,
	SignerTypeIssuanceManager,
	SignerTypeInvoiceManager,
	SignerTypePaymentOperator,
	SignerTypeLimitsManager,
	SignerTypeAccountManager,
	SignerTypeCommissionBalanceManager,
	SignerTypeOperationalBalanceManager,
	SignerTypeEventsChecker,
	SignerTypeExchangeAccManager,
	SignerTypeSyndicateAccManager,
	SignerTypeUserAssetManager,
	SignerTypeUserIssuanceManager,
	SignerTypeWithdrawManager,
	SignerTypeFeesManager,
	SignerTypeTxSender,
	SignerTypeAmlAlertManager,
	SignerTypeAmlAlertReviewer,
	SignerTypeKycAccManager,
	SignerTypeKycSuperAdmin,
	SignerTypeExternalSystemAccountIdPoolManager,
	SignerTypeKeyValueManager,
	SignerTypeSuperIssuanceManager,
	SignerTypeContractManager,
}

var signerTypeMap = map[int32]string{
	1:         "SignerTypeReader",
	2:         "SignerTypeNotVerifiedAccManager",
	4:         "SignerTypeGeneralAccManager",
	8:         "SignerTypeDirectDebitOperator",
	16:        "SignerTypeAssetManager",
	32:        "SignerTypeAssetRateManager",
	64:        "SignerTypeBalanceManager",
	128:       "SignerTypeIssuanceManager",
	256:       "SignerTypeInvoiceManager",
	512:       "SignerTypePaymentOperator",
	1024:      "SignerTypeLimitsManager",
	2048:      "SignerTypeAccountManager",
	4096:      "SignerTypeCommissionBalanceManager",
	8192:      "SignerTypeOperationalBalanceManager",
	16384:     "SignerTypeEventsChecker",
	32768:     "SignerTypeExchangeAccManager",
	65536:     "SignerTypeSyndicateAccManager",
	131072:    "SignerTypeUserAssetManager",
	262144:    "SignerTypeUserIssuanceManager",
	524288:    "SignerTypeWithdrawManager",
	1048576:   "SignerTypeFeesManager",
	2097152:   "SignerTypeTxSender",
	4194304:   "SignerTypeAmlAlertManager",
	8388608:   "SignerTypeAmlAlertReviewer",
	16777216:  "SignerTypeKycAccManager",
	33554432:  "SignerTypeKycSuperAdmin",
	67108864:  "SignerTypeExternalSystemAccountIdPoolManager",
	134217728: "SignerTypeKeyValueManager",
	268435456: "SignerTypeSuperIssuanceManager",
	536870912: "SignerTypeContractManager",
}

var signerTypeShortMap = map[int32]string{
	1:         "reader",
	2:         "not_verified_acc_manager",
	4:         "general_acc_manager",
	8:         "direct_debit_operator",
	16:        "asset_manager",
	32:        "asset_rate_manager",
	64:        "balance_manager",
	128:       "issuance_manager",
	256:       "invoice_manager",
	512:       "payment_operator",
	1024:      "limits_manager",
	2048:      "account_manager",
	4096:      "commission_balance_manager",
	8192:      "operational_balance_manager",
	16384:     "events_checker",
	32768:     "exchange_acc_manager",
	65536:     "syndicate_acc_manager",
	131072:    "user_asset_manager",
	262144:    "user_issuance_manager",
	524288:    "withdraw_manager",
	1048576:   "fees_manager",
	2097152:   "tx_sender",
	4194304:   "aml_alert_manager",
	8388608:   "aml_alert_reviewer",
	16777216:  "kyc_acc_manager",
	33554432:  "kyc_super_admin",
	67108864:  "external_system_account_id_pool_manager",
	134217728: "key_value_manager",
	268435456: "super_issuance_manager",
	536870912: "contract_manager",
}

var signerTypeRevMap = map[string]int32{
	"SignerTypeReader":                             1,
	"SignerTypeNotVerifiedAccManager":              2,
	"SignerTypeGeneralAccManager":                  4,
	"SignerTypeDirectDebitOperator":                8,
	"SignerTypeAssetManager":                       16,
	"SignerTypeAssetRateManager":                   32,
	"SignerTypeBalanceManager":                     64,
	"SignerTypeIssuanceManager":                    128,
	"SignerTypeInvoiceManager":                     256,
	"SignerTypePaymentOperator":                    512,
	"SignerTypeLimitsManager":                      1024,
	"SignerTypeAccountManager":                     2048,
	"SignerTypeCommissionBalanceManager":           4096,
	"SignerTypeOperationalBalanceManager":          8192,
	"SignerTypeEventsChecker":                      16384,
	"SignerTypeExchangeAccManager":                 32768,
	"SignerTypeSyndicateAccManager":                65536,
	"SignerTypeUserAssetManager":                   131072,
	"SignerTypeUserIssuanceManager":                262144,
	"SignerTypeWithdrawManager":                    524288,
	"SignerTypeFeesManager":                        1048576,
	"SignerTypeTxSender":                           2097152,
	"SignerTypeAmlAlertManager":                    4194304,
	"SignerTypeAmlAlertReviewer":                   8388608,
	"SignerTypeKycAccManager":                      16777216,
	"SignerTypeKycSuperAdmin":                      33554432,
	"SignerTypeExternalSystemAccountIdPoolManager": 67108864,
	"SignerTypeKeyValueManager":                    134217728,
	"SignerTypeSuperIssuanceManager":               268435456,
	"SignerTypeContractManager":                    536870912,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for SignerType
func (e SignerType) ValidEnum(v int32) bool {
	_, ok := signerTypeMap[v]
	return ok
}
func (e SignerType) isFlag() bool {
	for i := len(SignerTypeAll) - 1; i >= 0; i-- {
		expected := SignerType(2) << uint64(len(SignerTypeAll)-1) >> uint64(len(SignerTypeAll)-i)
		if expected != SignerTypeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e SignerType) String() string {
	name, _ := signerTypeMap[int32(e)]
	return name
}

func (e SignerType) ShortString() string {
	name, _ := signerTypeShortMap[int32(e)]
	return name
}

func (e SignerType) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range SignerTypeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *SignerType) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = SignerType(t.Value)
	return nil
}

// SignerExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type SignerExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u SignerExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of SignerExt
func (u SignerExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewSignerExt creates a new  SignerExt.
func NewSignerExt(v LedgerVersion, value interface{}) (result SignerExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// Signer is an XDR Struct defines as:
//
//   struct Signer
//    {
//        AccountID pubKey;
//        uint32 weight; // really only need 1byte
//    	uint32 signerType;
//    	uint32 identity;
//    	string256 name;
//
//    	 // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type Signer struct {
	PubKey     AccountId `json:"pubKey,omitempty"`
	Weight     Uint32    `json:"weight,omitempty"`
	SignerType Uint32    `json:"signerType,omitempty"`
	Identity   Uint32    `json:"identity,omitempty"`
	Name       String256 `json:"name,omitempty"`
	Ext        SignerExt `json:"ext,omitempty"`
}

// TrustEntryExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type TrustEntryExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u TrustEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of TrustEntryExt
func (u TrustEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewTrustEntryExt creates a new  TrustEntryExt.
func NewTrustEntryExt(v LedgerVersion, value interface{}) (result TrustEntryExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// TrustEntry is an XDR Struct defines as:
//
//   struct TrustEntry
//    {
//        AccountID allowedAccount;
//        BalanceID balanceToUse;
//
//    	 // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type TrustEntry struct {
	AllowedAccount AccountId     `json:"allowedAccount,omitempty"`
	BalanceToUse   BalanceId     `json:"balanceToUse,omitempty"`
	Ext            TrustEntryExt `json:"ext,omitempty"`
}

// LimitsExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type LimitsExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LimitsExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LimitsExt
func (u LimitsExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewLimitsExt creates a new  LimitsExt.
func NewLimitsExt(v LedgerVersion, value interface{}) (result LimitsExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// Limits is an XDR Struct defines as:
//
//   struct Limits
//    {
//        int64 dailyOut;
//    	int64 weeklyOut;
//    	int64 monthlyOut;
//        int64 annualOut;
//
//    	 // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//
//    };
//
type Limits struct {
	DailyOut   Int64     `json:"dailyOut,omitempty"`
	WeeklyOut  Int64     `json:"weeklyOut,omitempty"`
	MonthlyOut Int64     `json:"monthlyOut,omitempty"`
	AnnualOut  Int64     `json:"annualOut,omitempty"`
	Ext        LimitsExt `json:"ext,omitempty"`
}

// AccountPolicies is an XDR Enum defines as:
//
//   enum AccountPolicies
//    {
//    	NO_PERMISSIONS = 0,
//    	ALLOW_TO_CREATE_USER_VIA_API = 1
//    };
//
type AccountPolicies int32

const (
	AccountPoliciesNoPermissions           AccountPolicies = 0
	AccountPoliciesAllowToCreateUserViaApi AccountPolicies = 1
)

var AccountPoliciesAll = []AccountPolicies{
	AccountPoliciesNoPermissions,
	AccountPoliciesAllowToCreateUserViaApi,
}

var accountPoliciesMap = map[int32]string{
	0: "AccountPoliciesNoPermissions",
	1: "AccountPoliciesAllowToCreateUserViaApi",
}

var accountPoliciesShortMap = map[int32]string{
	0: "no_permissions",
	1: "allow_to_create_user_via_api",
}

var accountPoliciesRevMap = map[string]int32{
	"AccountPoliciesNoPermissions":           0,
	"AccountPoliciesAllowToCreateUserViaApi": 1,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for AccountPolicies
func (e AccountPolicies) ValidEnum(v int32) bool {
	_, ok := accountPoliciesMap[v]
	return ok
}
func (e AccountPolicies) isFlag() bool {
	for i := len(AccountPoliciesAll) - 1; i >= 0; i-- {
		expected := AccountPolicies(2) << uint64(len(AccountPoliciesAll)-1) >> uint64(len(AccountPoliciesAll)-i)
		if expected != AccountPoliciesAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e AccountPolicies) String() string {
	name, _ := accountPoliciesMap[int32(e)]
	return name
}

func (e AccountPolicies) ShortString() string {
	name, _ := accountPoliciesShortMap[int32(e)]
	return name
}

func (e AccountPolicies) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range AccountPoliciesAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *AccountPolicies) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = AccountPolicies(t.Value)
	return nil
}

// AccountType is an XDR Enum defines as:
//
//   enum AccountType
//    {
//    	OPERATIONAL = 1,       // operational account of the system
//    	GENERAL = 2,           // general account can perform payments, setoptions, be source account for tx, etc.
//    	COMMISSION = 3,        // commission account
//    	MASTER = 4,            // master account
//        NOT_VERIFIED = 5,
//    	SYNDICATE = 6, // can create asset
//    	EXCHANGE = 7,
//    	ACCREDITED_INVESTOR = 8,
//    	INSTITUTIONAL_INVESTOR = 9,
//    	VERIFIED = 10
//    };
//
type AccountType int32

const (
	AccountTypeOperational           AccountType = 1
	AccountTypeGeneral               AccountType = 2
	AccountTypeCommission            AccountType = 3
	AccountTypeMaster                AccountType = 4
	AccountTypeNotVerified           AccountType = 5
	AccountTypeSyndicate             AccountType = 6
	AccountTypeExchange              AccountType = 7
	AccountTypeAccreditedInvestor    AccountType = 8
	AccountTypeInstitutionalInvestor AccountType = 9
	AccountTypeVerified              AccountType = 10
)

var AccountTypeAll = []AccountType{
	AccountTypeOperational,
	AccountTypeGeneral,
	AccountTypeCommission,
	AccountTypeMaster,
	AccountTypeNotVerified,
	AccountTypeSyndicate,
	AccountTypeExchange,
	AccountTypeAccreditedInvestor,
	AccountTypeInstitutionalInvestor,
	AccountTypeVerified,
}

var accountTypeMap = map[int32]string{
	1:  "AccountTypeOperational",
	2:  "AccountTypeGeneral",
	3:  "AccountTypeCommission",
	4:  "AccountTypeMaster",
	5:  "AccountTypeNotVerified",
	6:  "AccountTypeSyndicate",
	7:  "AccountTypeExchange",
	8:  "AccountTypeAccreditedInvestor",
	9:  "AccountTypeInstitutionalInvestor",
	10: "AccountTypeVerified",
}

var accountTypeShortMap = map[int32]string{
	1:  "operational",
	2:  "general",
	3:  "commission",
	4:  "master",
	5:  "not_verified",
	6:  "syndicate",
	7:  "exchange",
	8:  "accredited_investor",
	9:  "institutional_investor",
	10: "verified",
}

var accountTypeRevMap = map[string]int32{
	"AccountTypeOperational":           1,
	"AccountTypeGeneral":               2,
	"AccountTypeCommission":            3,
	"AccountTypeMaster":                4,
	"AccountTypeNotVerified":           5,
	"AccountTypeSyndicate":             6,
	"AccountTypeExchange":              7,
	"AccountTypeAccreditedInvestor":    8,
	"AccountTypeInstitutionalInvestor": 9,
	"AccountTypeVerified":              10,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for AccountType
func (e AccountType) ValidEnum(v int32) bool {
	_, ok := accountTypeMap[v]
	return ok
}
func (e AccountType) isFlag() bool {
	for i := len(AccountTypeAll) - 1; i >= 0; i-- {
		expected := AccountType(2) << uint64(len(AccountTypeAll)-1) >> uint64(len(AccountTypeAll)-i)
		if expected != AccountTypeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e AccountType) String() string {
	name, _ := accountTypeMap[int32(e)]
	return name
}

func (e AccountType) ShortString() string {
	name, _ := accountTypeShortMap[int32(e)]
	return name
}

func (e AccountType) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range AccountTypeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *AccountType) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = AccountType(t.Value)
	return nil
}

// BlockReasons is an XDR Enum defines as:
//
//   enum BlockReasons
//    {
//    	RECOVERY_REQUEST = 1,
//    	KYC_UPDATE = 2,
//    	SUSPICIOUS_BEHAVIOR = 4,
//    	TOO_MANY_KYC_UPDATE_REQUESTS = 8,
//    	WITHDRAWAL = 16
//    };
//
type BlockReasons int32

const (
	BlockReasonsRecoveryRequest          BlockReasons = 1
	BlockReasonsKycUpdate                BlockReasons = 2
	BlockReasonsSuspiciousBehavior       BlockReasons = 4
	BlockReasonsTooManyKycUpdateRequests BlockReasons = 8
	BlockReasonsWithdrawal               BlockReasons = 16
)

var BlockReasonsAll = []BlockReasons{
	BlockReasonsRecoveryRequest,
	BlockReasonsKycUpdate,
	BlockReasonsSuspiciousBehavior,
	BlockReasonsTooManyKycUpdateRequests,
	BlockReasonsWithdrawal,
}

var blockReasonsMap = map[int32]string{
	1:  "BlockReasonsRecoveryRequest",
	2:  "BlockReasonsKycUpdate",
	4:  "BlockReasonsSuspiciousBehavior",
	8:  "BlockReasonsTooManyKycUpdateRequests",
	16: "BlockReasonsWithdrawal",
}

var blockReasonsShortMap = map[int32]string{
	1:  "recovery_request",
	2:  "kyc_update",
	4:  "suspicious_behavior",
	8:  "too_many_kyc_update_requests",
	16: "withdrawal",
}

var blockReasonsRevMap = map[string]int32{
	"BlockReasonsRecoveryRequest":          1,
	"BlockReasonsKycUpdate":                2,
	"BlockReasonsSuspiciousBehavior":       4,
	"BlockReasonsTooManyKycUpdateRequests": 8,
	"BlockReasonsWithdrawal":               16,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for BlockReasons
func (e BlockReasons) ValidEnum(v int32) bool {
	_, ok := blockReasonsMap[v]
	return ok
}
func (e BlockReasons) isFlag() bool {
	for i := len(BlockReasonsAll) - 1; i >= 0; i-- {
		expected := BlockReasons(2) << uint64(len(BlockReasonsAll)-1) >> uint64(len(BlockReasonsAll)-i)
		if expected != BlockReasonsAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e BlockReasons) String() string {
	name, _ := blockReasonsMap[int32(e)]
	return name
}

func (e BlockReasons) ShortString() string {
	name, _ := blockReasonsShortMap[int32(e)]
	return name
}

func (e BlockReasons) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range BlockReasonsAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *BlockReasons) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = BlockReasons(t.Value)
	return nil
}

// AccountEntryExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//    	case USE_KYC_LEVEL:
//    		uint32 kycLevel;
//        }
//
type AccountEntryExt struct {
	V        LedgerVersion `json:"v,omitempty"`
	KycLevel *Uint32       `json:"kycLevel,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u AccountEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of AccountEntryExt
func (u AccountEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	case LedgerVersionUseKycLevel:
		return "KycLevel", true
	}
	return "-", false
}

// NewAccountEntryExt creates a new  AccountEntryExt.
func NewAccountEntryExt(v LedgerVersion, value interface{}) (result AccountEntryExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	case LedgerVersionUseKycLevel:
		tv, ok := value.(Uint32)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint32")
			return
		}
		result.KycLevel = &tv
	}
	return
}

// MustKycLevel retrieves the KycLevel value from the union,
// panicing if the value is not set.
func (u AccountEntryExt) MustKycLevel() Uint32 {
	val, ok := u.GetKycLevel()

	if !ok {
		panic("arm KycLevel is not set")
	}

	return val
}

// GetKycLevel retrieves the KycLevel value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u AccountEntryExt) GetKycLevel() (result Uint32, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.V))

	if armName == "KycLevel" {
		result = *u.KycLevel
		ok = true
	}

	return
}

// AccountEntry is an XDR Struct defines as:
//
//   struct AccountEntry
//    {
//        AccountID accountID;      // master public key for this account
//        AccountID recoveryID;
//
//        // fields used for signatures
//        // thresholds stores unsigned bytes: [weight of master|low|medium|high]
//        Thresholds thresholds;
//
//        Signer signers<>; // possible signers for this account
//        Limits* limits;
//
//    	uint32 blockReasons;
//        AccountType accountType; // type of the account
//
//        // Referral marketing
//        AccountID* referrer;     // parent account
//
//    	int32 policies;
//
//        // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//    	case USE_KYC_LEVEL:
//    		uint32 kycLevel;
//        }
//
//        ext;
//    };
//
type AccountEntry struct {
	AccountId    AccountId       `json:"accountID,omitempty"`
	RecoveryId   AccountId       `json:"recoveryID,omitempty"`
	Thresholds   Thresholds      `json:"thresholds,omitempty"`
	Signers      []Signer        `json:"signers,omitempty"`
	Limits       *Limits         `json:"limits,omitempty"`
	BlockReasons Uint32          `json:"blockReasons,omitempty"`
	AccountType  AccountType     `json:"accountType,omitempty"`
	Referrer     *AccountId      `json:"referrer,omitempty"`
	Policies     Int32           `json:"policies,omitempty"`
	Ext          AccountEntryExt `json:"ext,omitempty"`
}

// AssetPairPolicy is an XDR Enum defines as:
//
//   enum AssetPairPolicy
//    {
//    	TRADEABLE_SECONDARY_MARKET = 1, // if not set pair can not be traided on secondary market
//    	PHYSICAL_PRICE_RESTRICTION = 2, // if set, then prices for new offers must be greater then physical price with correction
//    	CURRENT_PRICE_RESTRICTION = 4 // if set, then price for new offers must be in interval of (1 +- maxPriceStep)*currentPrice
//    };
//
type AssetPairPolicy int32

const (
	AssetPairPolicyTradeableSecondaryMarket AssetPairPolicy = 1
	AssetPairPolicyPhysicalPriceRestriction AssetPairPolicy = 2
	AssetPairPolicyCurrentPriceRestriction  AssetPairPolicy = 4
)

var AssetPairPolicyAll = []AssetPairPolicy{
	AssetPairPolicyTradeableSecondaryMarket,
	AssetPairPolicyPhysicalPriceRestriction,
	AssetPairPolicyCurrentPriceRestriction,
}

var assetPairPolicyMap = map[int32]string{
	1: "AssetPairPolicyTradeableSecondaryMarket",
	2: "AssetPairPolicyPhysicalPriceRestriction",
	4: "AssetPairPolicyCurrentPriceRestriction",
}

var assetPairPolicyShortMap = map[int32]string{
	1: "tradeable_secondary_market",
	2: "physical_price_restriction",
	4: "current_price_restriction",
}

var assetPairPolicyRevMap = map[string]int32{
	"AssetPairPolicyTradeableSecondaryMarket": 1,
	"AssetPairPolicyPhysicalPriceRestriction": 2,
	"AssetPairPolicyCurrentPriceRestriction":  4,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for AssetPairPolicy
func (e AssetPairPolicy) ValidEnum(v int32) bool {
	_, ok := assetPairPolicyMap[v]
	return ok
}
func (e AssetPairPolicy) isFlag() bool {
	for i := len(AssetPairPolicyAll) - 1; i >= 0; i-- {
		expected := AssetPairPolicy(2) << uint64(len(AssetPairPolicyAll)-1) >> uint64(len(AssetPairPolicyAll)-i)
		if expected != AssetPairPolicyAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e AssetPairPolicy) String() string {
	name, _ := assetPairPolicyMap[int32(e)]
	return name
}

func (e AssetPairPolicy) ShortString() string {
	name, _ := assetPairPolicyShortMap[int32(e)]
	return name
}

func (e AssetPairPolicy) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range AssetPairPolicyAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *AssetPairPolicy) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = AssetPairPolicy(t.Value)
	return nil
}

// AssetPairEntryExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type AssetPairEntryExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u AssetPairEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of AssetPairEntryExt
func (u AssetPairEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewAssetPairEntryExt creates a new  AssetPairEntryExt.
func NewAssetPairEntryExt(v LedgerVersion, value interface{}) (result AssetPairEntryExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// AssetPairEntry is an XDR Struct defines as:
//
//   struct AssetPairEntry
//    {
//        AssetCode base;
//    	AssetCode quote;
//
//        int64 currentPrice;
//        int64 physicalPrice;
//
//    	int64 physicalPriceCorrection; // correction of physical price in percents. If physical price is set and restriction by physical price set, mininal price for offer for this pair will be physicalPrice * physicalPriceCorrection
//    	int64 maxPriceStep; // max price step in percent. User is allowed to set offer with price < (1 - maxPriceStep)*currentPrice and > (1 + maxPriceStep)*currentPrice
//
//
//    	int32 policies;
//
//        // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type AssetPairEntry struct {
	Base                    AssetCode         `json:"base,omitempty"`
	Quote                   AssetCode         `json:"quote,omitempty"`
	CurrentPrice            Int64             `json:"currentPrice,omitempty"`
	PhysicalPrice           Int64             `json:"physicalPrice,omitempty"`
	PhysicalPriceCorrection Int64             `json:"physicalPriceCorrection,omitempty"`
	MaxPriceStep            Int64             `json:"maxPriceStep,omitempty"`
	Policies                Int32             `json:"policies,omitempty"`
	Ext                     AssetPairEntryExt `json:"ext,omitempty"`
}

// AssetPolicy is an XDR Enum defines as:
//
//   enum AssetPolicy
//    {
//    	TRANSFERABLE = 1,
//    	BASE_ASSET = 2,
//    	STATS_QUOTE_ASSET = 4,
//    	WITHDRAWABLE = 8,
//    	TWO_STEP_WITHDRAWAL = 16,
//    	REQUIRES_KYC = 32,
//    	ISSUANCE_MANUAL_REVIEW_REQUIRED = 64,
//    	REQUIRES_VERIFICATION = 128
//    };
//
type AssetPolicy int32

const (
	AssetPolicyTransferable                 AssetPolicy = 1
	AssetPolicyBaseAsset                    AssetPolicy = 2
	AssetPolicyStatsQuoteAsset              AssetPolicy = 4
	AssetPolicyWithdrawable                 AssetPolicy = 8
	AssetPolicyTwoStepWithdrawal            AssetPolicy = 16
	AssetPolicyRequiresKyc                  AssetPolicy = 32
	AssetPolicyIssuanceManualReviewRequired AssetPolicy = 64
	AssetPolicyRequiresVerification         AssetPolicy = 128
)

var AssetPolicyAll = []AssetPolicy{
	AssetPolicyTransferable,
	AssetPolicyBaseAsset,
	AssetPolicyStatsQuoteAsset,
	AssetPolicyWithdrawable,
	AssetPolicyTwoStepWithdrawal,
	AssetPolicyRequiresKyc,
	AssetPolicyIssuanceManualReviewRequired,
	AssetPolicyRequiresVerification,
}

var assetPolicyMap = map[int32]string{
	1:   "AssetPolicyTransferable",
	2:   "AssetPolicyBaseAsset",
	4:   "AssetPolicyStatsQuoteAsset",
	8:   "AssetPolicyWithdrawable",
	16:  "AssetPolicyTwoStepWithdrawal",
	32:  "AssetPolicyRequiresKyc",
	64:  "AssetPolicyIssuanceManualReviewRequired",
	128: "AssetPolicyRequiresVerification",
}

var assetPolicyShortMap = map[int32]string{
	1:   "transferable",
	2:   "base_asset",
	4:   "stats_quote_asset",
	8:   "withdrawable",
	16:  "two_step_withdrawal",
	32:  "requires_kyc",
	64:  "issuance_manual_review_required",
	128: "requires_verification",
}

var assetPolicyRevMap = map[string]int32{
	"AssetPolicyTransferable":                 1,
	"AssetPolicyBaseAsset":                    2,
	"AssetPolicyStatsQuoteAsset":              4,
	"AssetPolicyWithdrawable":                 8,
	"AssetPolicyTwoStepWithdrawal":            16,
	"AssetPolicyRequiresKyc":                  32,
	"AssetPolicyIssuanceManualReviewRequired": 64,
	"AssetPolicyRequiresVerification":         128,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for AssetPolicy
func (e AssetPolicy) ValidEnum(v int32) bool {
	_, ok := assetPolicyMap[v]
	return ok
}
func (e AssetPolicy) isFlag() bool {
	for i := len(AssetPolicyAll) - 1; i >= 0; i-- {
		expected := AssetPolicy(2) << uint64(len(AssetPolicyAll)-1) >> uint64(len(AssetPolicyAll)-i)
		if expected != AssetPolicyAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e AssetPolicy) String() string {
	name, _ := assetPolicyMap[int32(e)]
	return name
}

func (e AssetPolicy) ShortString() string {
	name, _ := assetPolicyShortMap[int32(e)]
	return name
}

func (e AssetPolicy) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range AssetPolicyAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *AssetPolicy) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = AssetPolicy(t.Value)
	return nil
}

// AssetSystemPolicies is an XDR Enum defines as:
//
//   enum AssetSystemPolicies
//    {
//    	TWO_STEP_WITHDRAW = 1
//    };
//
type AssetSystemPolicies int32

const (
	AssetSystemPoliciesTwoStepWithdraw AssetSystemPolicies = 1
)

var AssetSystemPoliciesAll = []AssetSystemPolicies{
	AssetSystemPoliciesTwoStepWithdraw,
}

var assetSystemPoliciesMap = map[int32]string{
	1: "AssetSystemPoliciesTwoStepWithdraw",
}

var assetSystemPoliciesShortMap = map[int32]string{
	1: "two_step_withdraw",
}

var assetSystemPoliciesRevMap = map[string]int32{
	"AssetSystemPoliciesTwoStepWithdraw": 1,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for AssetSystemPolicies
func (e AssetSystemPolicies) ValidEnum(v int32) bool {
	_, ok := assetSystemPoliciesMap[v]
	return ok
}
func (e AssetSystemPolicies) isFlag() bool {
	for i := len(AssetSystemPoliciesAll) - 1; i >= 0; i-- {
		expected := AssetSystemPolicies(2) << uint64(len(AssetSystemPoliciesAll)-1) >> uint64(len(AssetSystemPoliciesAll)-i)
		if expected != AssetSystemPoliciesAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e AssetSystemPolicies) String() string {
	name, _ := assetSystemPoliciesMap[int32(e)]
	return name
}

func (e AssetSystemPolicies) ShortString() string {
	name, _ := assetSystemPoliciesShortMap[int32(e)]
	return name
}

func (e AssetSystemPolicies) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range AssetSystemPoliciesAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *AssetSystemPolicies) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = AssetSystemPolicies(t.Value)
	return nil
}

// AssetEntryExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type AssetEntryExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u AssetEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of AssetEntryExt
func (u AssetEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewAssetEntryExt creates a new  AssetEntryExt.
func NewAssetEntryExt(v LedgerVersion, value interface{}) (result AssetEntryExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// AssetEntry is an XDR Struct defines as:
//
//   struct AssetEntry
//    {
//        AssetCode code;
//    	AccountID owner;
//    	AccountID preissuedAssetSigner; // signer of pre issuance tokens
//    	longstring details;
//    	uint64 maxIssuanceAmount; // max number of tokens to be issued
//    	uint64 availableForIssueance; // pre issued tokens available for issuance
//    	uint64 issued; // number of issued tokens
//    	uint64 pendingIssuance; // number of tokens locked for entries like token sale. lockedIssuance + issued can not be > maxIssuanceAmount
//        uint32 policies;
//
//        // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type AssetEntry struct {
	Code                  AssetCode     `json:"code,omitempty"`
	Owner                 AccountId     `json:"owner,omitempty"`
	PreissuedAssetSigner  AccountId     `json:"preissuedAssetSigner,omitempty"`
	Details               Longstring    `json:"details,omitempty"`
	MaxIssuanceAmount     Uint64        `json:"maxIssuanceAmount,omitempty"`
	AvailableForIssueance Uint64        `json:"availableForIssueance,omitempty"`
	Issued                Uint64        `json:"issued,omitempty"`
	PendingIssuance       Uint64        `json:"pendingIssuance,omitempty"`
	Policies              Uint32        `json:"policies,omitempty"`
	Ext                   AssetEntryExt `json:"ext,omitempty"`
}

// BalanceEntryExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type BalanceEntryExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u BalanceEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of BalanceEntryExt
func (u BalanceEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewBalanceEntryExt creates a new  BalanceEntryExt.
func NewBalanceEntryExt(v LedgerVersion, value interface{}) (result BalanceEntryExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// BalanceEntry is an XDR Struct defines as:
//
//   struct BalanceEntry
//    {
//        BalanceID balanceID;
//        AssetCode asset;
//        AccountID accountID;
//        uint64 amount;
//        uint64 locked;
//
//        // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type BalanceEntry struct {
	BalanceId BalanceId       `json:"balanceID,omitempty"`
	Asset     AssetCode       `json:"asset,omitempty"`
	AccountId AccountId       `json:"accountID,omitempty"`
	Amount    Uint64          `json:"amount,omitempty"`
	Locked    Uint64          `json:"locked,omitempty"`
	Ext       BalanceEntryExt `json:"ext,omitempty"`
}

// ContractState is an XDR Enum defines as:
//
//   enum ContractState
//    {
//        NO_CONFIRMATIONS = 0,
//        CUSTOMER_CONFIRMED = 1,
//        CONTRACTOR_CONFIRMED = 2,
//        DISPUTING = 4,
//        REVERTING_RESOLVE = 8,
//        NOT_REVERTING_RESOLVE = 16
//    };
//
type ContractState int32

const (
	ContractStateNoConfirmations     ContractState = 0
	ContractStateCustomerConfirmed   ContractState = 1
	ContractStateContractorConfirmed ContractState = 2
	ContractStateDisputing           ContractState = 4
	ContractStateRevertingResolve    ContractState = 8
	ContractStateNotRevertingResolve ContractState = 16
)

var ContractStateAll = []ContractState{
	ContractStateNoConfirmations,
	ContractStateCustomerConfirmed,
	ContractStateContractorConfirmed,
	ContractStateDisputing,
	ContractStateRevertingResolve,
	ContractStateNotRevertingResolve,
}

var contractStateMap = map[int32]string{
	0:  "ContractStateNoConfirmations",
	1:  "ContractStateCustomerConfirmed",
	2:  "ContractStateContractorConfirmed",
	4:  "ContractStateDisputing",
	8:  "ContractStateRevertingResolve",
	16: "ContractStateNotRevertingResolve",
}

var contractStateShortMap = map[int32]string{
	0:  "no_confirmations",
	1:  "customer_confirmed",
	2:  "contractor_confirmed",
	4:  "disputing",
	8:  "reverting_resolve",
	16: "not_reverting_resolve",
}

var contractStateRevMap = map[string]int32{
	"ContractStateNoConfirmations":     0,
	"ContractStateCustomerConfirmed":   1,
	"ContractStateContractorConfirmed": 2,
	"ContractStateDisputing":           4,
	"ContractStateRevertingResolve":    8,
	"ContractStateNotRevertingResolve": 16,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ContractState
func (e ContractState) ValidEnum(v int32) bool {
	_, ok := contractStateMap[v]
	return ok
}
func (e ContractState) isFlag() bool {
	for i := len(ContractStateAll) - 1; i >= 0; i-- {
		expected := ContractState(2) << uint64(len(ContractStateAll)-1) >> uint64(len(ContractStateAll)-i)
		if expected != ContractStateAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ContractState) String() string {
	name, _ := contractStateMap[int32(e)]
	return name
}

func (e ContractState) ShortString() string {
	name, _ := contractStateShortMap[int32(e)]
	return name
}

func (e ContractState) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range ContractStateAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ContractState) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ContractState(t.Value)
	return nil
}

// ContractEntryExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        case ADD_CUSTOMER_DETAILS_TO_CONTRACT:
//            longstring customerDetails;
//        }
//
type ContractEntryExt struct {
	V               LedgerVersion `json:"v,omitempty"`
	CustomerDetails *Longstring   `json:"customerDetails,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ContractEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ContractEntryExt
func (u ContractEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	case LedgerVersionAddCustomerDetailsToContract:
		return "CustomerDetails", true
	}
	return "-", false
}

// NewContractEntryExt creates a new  ContractEntryExt.
func NewContractEntryExt(v LedgerVersion, value interface{}) (result ContractEntryExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	case LedgerVersionAddCustomerDetailsToContract:
		tv, ok := value.(Longstring)
		if !ok {
			err = fmt.Errorf("invalid value, must be Longstring")
			return
		}
		result.CustomerDetails = &tv
	}
	return
}

// MustCustomerDetails retrieves the CustomerDetails value from the union,
// panicing if the value is not set.
func (u ContractEntryExt) MustCustomerDetails() Longstring {
	val, ok := u.GetCustomerDetails()

	if !ok {
		panic("arm CustomerDetails is not set")
	}

	return val
}

// GetCustomerDetails retrieves the CustomerDetails value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ContractEntryExt) GetCustomerDetails() (result Longstring, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.V))

	if armName == "CustomerDetails" {
		result = *u.CustomerDetails
		ok = true
	}

	return
}

// ContractEntry is an XDR Struct defines as:
//
//   struct ContractEntry
//    {
//        uint64 contractID;
//
//        AccountID contractor;
//        AccountID customer;
//        AccountID escrow;
//
//        uint64 startTime;
//        uint64 endTime;
//        uint64 invoiceRequestsIDs<>;
//        longstring initialDetails;
//
//        uint32 state;
//
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        case ADD_CUSTOMER_DETAILS_TO_CONTRACT:
//            longstring customerDetails;
//        }
//        ext;
//    };
//
type ContractEntry struct {
	ContractId         Uint64           `json:"contractID,omitempty"`
	Contractor         AccountId        `json:"contractor,omitempty"`
	Customer           AccountId        `json:"customer,omitempty"`
	Escrow             AccountId        `json:"escrow,omitempty"`
	StartTime          Uint64           `json:"startTime,omitempty"`
	EndTime            Uint64           `json:"endTime,omitempty"`
	InvoiceRequestsIDs []Uint64         `json:"invoiceRequestsIDs,omitempty"`
	InitialDetails     Longstring       `json:"initialDetails,omitempty"`
	State              Uint32           `json:"state,omitempty"`
	Ext                ContractEntryExt `json:"ext,omitempty"`
}

// ExternalSystemAccountIdPoolEntryExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//           void;
//        }
//
type ExternalSystemAccountIdPoolEntryExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ExternalSystemAccountIdPoolEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ExternalSystemAccountIdPoolEntryExt
func (u ExternalSystemAccountIdPoolEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewExternalSystemAccountIdPoolEntryExt creates a new  ExternalSystemAccountIdPoolEntryExt.
func NewExternalSystemAccountIdPoolEntryExt(v LedgerVersion, value interface{}) (result ExternalSystemAccountIdPoolEntryExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ExternalSystemAccountIdPoolEntry is an XDR Struct defines as:
//
//   struct ExternalSystemAccountIDPoolEntry
//    {
//        uint64 poolEntryID;
//        int32 externalSystemType;
//        longstring data;
//        AccountID* accountID;
//        uint64 expiresAt;
//        uint64 bindedAt;
//        uint64 parent;
//        bool isDeleted;
//
//
//        // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//           void;
//        }
//        ext;
//    };
//
type ExternalSystemAccountIdPoolEntry struct {
	PoolEntryId        Uint64                              `json:"poolEntryID,omitempty"`
	ExternalSystemType Int32                               `json:"externalSystemType,omitempty"`
	Data               Longstring                          `json:"data,omitempty"`
	AccountId          *AccountId                          `json:"accountID,omitempty"`
	ExpiresAt          Uint64                              `json:"expiresAt,omitempty"`
	BindedAt           Uint64                              `json:"bindedAt,omitempty"`
	Parent             Uint64                              `json:"parent,omitempty"`
	IsDeleted          bool                                `json:"isDeleted,omitempty"`
	Ext                ExternalSystemAccountIdPoolEntryExt `json:"ext,omitempty"`
}

// ExternalSystemAccountIdExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type ExternalSystemAccountIdExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ExternalSystemAccountIdExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ExternalSystemAccountIdExt
func (u ExternalSystemAccountIdExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewExternalSystemAccountIdExt creates a new  ExternalSystemAccountIdExt.
func NewExternalSystemAccountIdExt(v LedgerVersion, value interface{}) (result ExternalSystemAccountIdExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ExternalSystemAccountId is an XDR Struct defines as:
//
//   struct ExternalSystemAccountID
//    {
//        AccountID accountID;
//        int32 externalSystemType;
//    	longstring data;
//
//    	 // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type ExternalSystemAccountId struct {
	AccountId          AccountId                  `json:"accountID,omitempty"`
	ExternalSystemType Int32                      `json:"externalSystemType,omitempty"`
	Data               Longstring                 `json:"data,omitempty"`
	Ext                ExternalSystemAccountIdExt `json:"ext,omitempty"`
}

// FeeType is an XDR Enum defines as:
//
//   enum FeeType
//    {
//        PAYMENT_FEE = 0,
//    	OFFER_FEE = 1,
//        WITHDRAWAL_FEE = 2,
//        ISSUANCE_FEE = 3,
//        INVEST_FEE = 4 // fee to be taken while creating sale participation
//    };
//
type FeeType int32

const (
	FeeTypePaymentFee    FeeType = 0
	FeeTypeOfferFee      FeeType = 1
	FeeTypeWithdrawalFee FeeType = 2
	FeeTypeIssuanceFee   FeeType = 3
	FeeTypeInvestFee     FeeType = 4
)

var FeeTypeAll = []FeeType{
	FeeTypePaymentFee,
	FeeTypeOfferFee,
	FeeTypeWithdrawalFee,
	FeeTypeIssuanceFee,
	FeeTypeInvestFee,
}

var feeTypeMap = map[int32]string{
	0: "FeeTypePaymentFee",
	1: "FeeTypeOfferFee",
	2: "FeeTypeWithdrawalFee",
	3: "FeeTypeIssuanceFee",
	4: "FeeTypeInvestFee",
}

var feeTypeShortMap = map[int32]string{
	0: "payment_fee",
	1: "offer_fee",
	2: "withdrawal_fee",
	3: "issuance_fee",
	4: "invest_fee",
}

var feeTypeRevMap = map[string]int32{
	"FeeTypePaymentFee":    0,
	"FeeTypeOfferFee":      1,
	"FeeTypeWithdrawalFee": 2,
	"FeeTypeIssuanceFee":   3,
	"FeeTypeInvestFee":     4,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for FeeType
func (e FeeType) ValidEnum(v int32) bool {
	_, ok := feeTypeMap[v]
	return ok
}
func (e FeeType) isFlag() bool {
	for i := len(FeeTypeAll) - 1; i >= 0; i-- {
		expected := FeeType(2) << uint64(len(FeeTypeAll)-1) >> uint64(len(FeeTypeAll)-i)
		if expected != FeeTypeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e FeeType) String() string {
	name, _ := feeTypeMap[int32(e)]
	return name
}

func (e FeeType) ShortString() string {
	name, _ := feeTypeShortMap[int32(e)]
	return name
}

func (e FeeType) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range FeeTypeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *FeeType) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = FeeType(t.Value)
	return nil
}

// EmissionFeeType is an XDR Enum defines as:
//
//   enum EmissionFeeType
//    {
//    	PRIMARY_MARKET = 1,
//    	SECONDARY_MARKET = 2
//    };
//
type EmissionFeeType int32

const (
	EmissionFeeTypePrimaryMarket   EmissionFeeType = 1
	EmissionFeeTypeSecondaryMarket EmissionFeeType = 2
)

var EmissionFeeTypeAll = []EmissionFeeType{
	EmissionFeeTypePrimaryMarket,
	EmissionFeeTypeSecondaryMarket,
}

var emissionFeeTypeMap = map[int32]string{
	1: "EmissionFeeTypePrimaryMarket",
	2: "EmissionFeeTypeSecondaryMarket",
}

var emissionFeeTypeShortMap = map[int32]string{
	1: "primary_market",
	2: "secondary_market",
}

var emissionFeeTypeRevMap = map[string]int32{
	"EmissionFeeTypePrimaryMarket":   1,
	"EmissionFeeTypeSecondaryMarket": 2,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for EmissionFeeType
func (e EmissionFeeType) ValidEnum(v int32) bool {
	_, ok := emissionFeeTypeMap[v]
	return ok
}
func (e EmissionFeeType) isFlag() bool {
	for i := len(EmissionFeeTypeAll) - 1; i >= 0; i-- {
		expected := EmissionFeeType(2) << uint64(len(EmissionFeeTypeAll)-1) >> uint64(len(EmissionFeeTypeAll)-i)
		if expected != EmissionFeeTypeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e EmissionFeeType) String() string {
	name, _ := emissionFeeTypeMap[int32(e)]
	return name
}

func (e EmissionFeeType) ShortString() string {
	name, _ := emissionFeeTypeShortMap[int32(e)]
	return name
}

func (e EmissionFeeType) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range EmissionFeeTypeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *EmissionFeeType) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = EmissionFeeType(t.Value)
	return nil
}

// PaymentFeeType is an XDR Enum defines as:
//
//   enum PaymentFeeType
//    {
//        OUTGOING = 1,
//        INCOMING = 2
//    };
//
type PaymentFeeType int32

const (
	PaymentFeeTypeOutgoing PaymentFeeType = 1
	PaymentFeeTypeIncoming PaymentFeeType = 2
)

var PaymentFeeTypeAll = []PaymentFeeType{
	PaymentFeeTypeOutgoing,
	PaymentFeeTypeIncoming,
}

var paymentFeeTypeMap = map[int32]string{
	1: "PaymentFeeTypeOutgoing",
	2: "PaymentFeeTypeIncoming",
}

var paymentFeeTypeShortMap = map[int32]string{
	1: "outgoing",
	2: "incoming",
}

var paymentFeeTypeRevMap = map[string]int32{
	"PaymentFeeTypeOutgoing": 1,
	"PaymentFeeTypeIncoming": 2,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for PaymentFeeType
func (e PaymentFeeType) ValidEnum(v int32) bool {
	_, ok := paymentFeeTypeMap[v]
	return ok
}
func (e PaymentFeeType) isFlag() bool {
	for i := len(PaymentFeeTypeAll) - 1; i >= 0; i-- {
		expected := PaymentFeeType(2) << uint64(len(PaymentFeeTypeAll)-1) >> uint64(len(PaymentFeeTypeAll)-i)
		if expected != PaymentFeeTypeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e PaymentFeeType) String() string {
	name, _ := paymentFeeTypeMap[int32(e)]
	return name
}

func (e PaymentFeeType) ShortString() string {
	name, _ := paymentFeeTypeShortMap[int32(e)]
	return name
}

func (e PaymentFeeType) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range PaymentFeeTypeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *PaymentFeeType) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = PaymentFeeType(t.Value)
	return nil
}

// FeeEntryExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        case CROSS_ASSET_FEE:
//            AssetCode feeAsset;
//        }
//
type FeeEntryExt struct {
	V        LedgerVersion `json:"v,omitempty"`
	FeeAsset *AssetCode    `json:"feeAsset,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u FeeEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of FeeEntryExt
func (u FeeEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	case LedgerVersionCrossAssetFee:
		return "FeeAsset", true
	}
	return "-", false
}

// NewFeeEntryExt creates a new  FeeEntryExt.
func NewFeeEntryExt(v LedgerVersion, value interface{}) (result FeeEntryExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	case LedgerVersionCrossAssetFee:
		tv, ok := value.(AssetCode)
		if !ok {
			err = fmt.Errorf("invalid value, must be AssetCode")
			return
		}
		result.FeeAsset = &tv
	}
	return
}

// MustFeeAsset retrieves the FeeAsset value from the union,
// panicing if the value is not set.
func (u FeeEntryExt) MustFeeAsset() AssetCode {
	val, ok := u.GetFeeAsset()

	if !ok {
		panic("arm FeeAsset is not set")
	}

	return val
}

// GetFeeAsset retrieves the FeeAsset value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u FeeEntryExt) GetFeeAsset() (result AssetCode, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.V))

	if armName == "FeeAsset" {
		result = *u.FeeAsset
		ok = true
	}

	return
}

// FeeEntry is an XDR Struct defines as:
//
//   struct FeeEntry
//    {
//        FeeType feeType;
//        AssetCode asset;
//        int64 fixedFee; // fee paid for operation
//    	int64 percentFee; // percent of transfer amount to be charged
//
//        AccountID* accountID;
//        AccountType* accountType;
//        int64 subtype; // for example, different withdrawals — bars or coins
//
//        int64 lowerBound;
//        int64 upperBound;
//
//        Hash hash;
//
//    	// reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        case CROSS_ASSET_FEE:
//            AssetCode feeAsset;
//        }
//        ext;
//    };
//
type FeeEntry struct {
	FeeType     FeeType      `json:"feeType,omitempty"`
	Asset       AssetCode    `json:"asset,omitempty"`
	FixedFee    Int64        `json:"fixedFee,omitempty"`
	PercentFee  Int64        `json:"percentFee,omitempty"`
	AccountId   *AccountId   `json:"accountID,omitempty"`
	AccountType *AccountType `json:"accountType,omitempty"`
	Subtype     Int64        `json:"subtype,omitempty"`
	LowerBound  Int64        `json:"lowerBound,omitempty"`
	UpperBound  Int64        `json:"upperBound,omitempty"`
	Hash        Hash         `json:"hash,omitempty"`
	Ext         FeeEntryExt  `json:"ext,omitempty"`
}

// KeyValueEntryType is an XDR Enum defines as:
//
//   enum KeyValueEntryType
//        {
//            UINT32 = 1,
//            STRING = 2
//        };
//
type KeyValueEntryType int32

const (
	KeyValueEntryTypeUint32 KeyValueEntryType = 1
	KeyValueEntryTypeString KeyValueEntryType = 2
)

var KeyValueEntryTypeAll = []KeyValueEntryType{
	KeyValueEntryTypeUint32,
	KeyValueEntryTypeString,
}

var keyValueEntryTypeMap = map[int32]string{
	1: "KeyValueEntryTypeUint32",
	2: "KeyValueEntryTypeString",
}

var keyValueEntryTypeShortMap = map[int32]string{
	1: "uint32",
	2: "string",
}

var keyValueEntryTypeRevMap = map[string]int32{
	"KeyValueEntryTypeUint32": 1,
	"KeyValueEntryTypeString": 2,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for KeyValueEntryType
func (e KeyValueEntryType) ValidEnum(v int32) bool {
	_, ok := keyValueEntryTypeMap[v]
	return ok
}
func (e KeyValueEntryType) isFlag() bool {
	for i := len(KeyValueEntryTypeAll) - 1; i >= 0; i-- {
		expected := KeyValueEntryType(2) << uint64(len(KeyValueEntryTypeAll)-1) >> uint64(len(KeyValueEntryTypeAll)-i)
		if expected != KeyValueEntryTypeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e KeyValueEntryType) String() string {
	name, _ := keyValueEntryTypeMap[int32(e)]
	return name
}

func (e KeyValueEntryType) ShortString() string {
	name, _ := keyValueEntryTypeShortMap[int32(e)]
	return name
}

func (e KeyValueEntryType) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range KeyValueEntryTypeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *KeyValueEntryType) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = KeyValueEntryType(t.Value)
	return nil
}

// KeyValueEntryValue is an XDR NestedUnion defines as:
//
//   union switch (KeyValueEntryType type)
//            {
//                 case UINT32:
//                    uint32 ui32Value;
//                 case STRING:
//                    string stringValue<>;
//            }
//
type KeyValueEntryValue struct {
	Type        KeyValueEntryType `json:"type,omitempty"`
	Ui32Value   *Uint32           `json:"ui32Value,omitempty"`
	StringValue *string           `json:"stringValue,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u KeyValueEntryValue) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of KeyValueEntryValue
func (u KeyValueEntryValue) ArmForSwitch(sw int32) (string, bool) {
	switch KeyValueEntryType(sw) {
	case KeyValueEntryTypeUint32:
		return "Ui32Value", true
	case KeyValueEntryTypeString:
		return "StringValue", true
	}
	return "-", false
}

// NewKeyValueEntryValue creates a new  KeyValueEntryValue.
func NewKeyValueEntryValue(aType KeyValueEntryType, value interface{}) (result KeyValueEntryValue, err error) {
	result.Type = aType
	switch KeyValueEntryType(aType) {
	case KeyValueEntryTypeUint32:
		tv, ok := value.(Uint32)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint32")
			return
		}
		result.Ui32Value = &tv
	case KeyValueEntryTypeString:
		tv, ok := value.(string)
		if !ok {
			err = fmt.Errorf("invalid value, must be string")
			return
		}
		result.StringValue = &tv
	}
	return
}

// MustUi32Value retrieves the Ui32Value value from the union,
// panicing if the value is not set.
func (u KeyValueEntryValue) MustUi32Value() Uint32 {
	val, ok := u.GetUi32Value()

	if !ok {
		panic("arm Ui32Value is not set")
	}

	return val
}

// GetUi32Value retrieves the Ui32Value value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u KeyValueEntryValue) GetUi32Value() (result Uint32, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Ui32Value" {
		result = *u.Ui32Value
		ok = true
	}

	return
}

// MustStringValue retrieves the StringValue value from the union,
// panicing if the value is not set.
func (u KeyValueEntryValue) MustStringValue() string {
	val, ok := u.GetStringValue()

	if !ok {
		panic("arm StringValue is not set")
	}

	return val
}

// GetStringValue retrieves the StringValue value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u KeyValueEntryValue) GetStringValue() (result string, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "StringValue" {
		result = *u.StringValue
		ok = true
	}

	return
}

// KeyValueEntryExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//                case EMPTY_VERSION:
//                    void;
//            }
//
type KeyValueEntryExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u KeyValueEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of KeyValueEntryExt
func (u KeyValueEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewKeyValueEntryExt creates a new  KeyValueEntryExt.
func NewKeyValueEntryExt(v LedgerVersion, value interface{}) (result KeyValueEntryExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// KeyValueEntry is an XDR Struct defines as:
//
//   struct KeyValueEntry
//        {
//            longstring key;
//
//            union switch (KeyValueEntryType type)
//            {
//                 case UINT32:
//                    uint32 ui32Value;
//                 case STRING:
//                    string stringValue<>;
//            }
//            value;
//
//            // reserved for future use
//            union switch (LedgerVersion v)
//            {
//                case EMPTY_VERSION:
//                    void;
//            }
//            ext;
//        };
//
type KeyValueEntry struct {
	Key   Longstring         `json:"key,omitempty"`
	Value KeyValueEntryValue `json:"value,omitempty"`
	Ext   KeyValueEntryExt   `json:"ext,omitempty"`
}

// StatsOpType is an XDR Enum defines as:
//
//   enum StatsOpType
//    {
//        PAYMENT_OUT = 1,
//        WITHDRAW = 2,
//        SPEND = 3,
//        DEPOSIT = 4
//    };
//
type StatsOpType int32

const (
	StatsOpTypePaymentOut StatsOpType = 1
	StatsOpTypeWithdraw   StatsOpType = 2
	StatsOpTypeSpend      StatsOpType = 3
	StatsOpTypeDeposit    StatsOpType = 4
)

var StatsOpTypeAll = []StatsOpType{
	StatsOpTypePaymentOut,
	StatsOpTypeWithdraw,
	StatsOpTypeSpend,
	StatsOpTypeDeposit,
}

var statsOpTypeMap = map[int32]string{
	1: "StatsOpTypePaymentOut",
	2: "StatsOpTypeWithdraw",
	3: "StatsOpTypeSpend",
	4: "StatsOpTypeDeposit",
}

var statsOpTypeShortMap = map[int32]string{
	1: "payment_out",
	2: "withdraw",
	3: "spend",
	4: "deposit",
}

var statsOpTypeRevMap = map[string]int32{
	"StatsOpTypePaymentOut": 1,
	"StatsOpTypeWithdraw":   2,
	"StatsOpTypeSpend":      3,
	"StatsOpTypeDeposit":    4,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for StatsOpType
func (e StatsOpType) ValidEnum(v int32) bool {
	_, ok := statsOpTypeMap[v]
	return ok
}
func (e StatsOpType) isFlag() bool {
	for i := len(StatsOpTypeAll) - 1; i >= 0; i-- {
		expected := StatsOpType(2) << uint64(len(StatsOpTypeAll)-1) >> uint64(len(StatsOpTypeAll)-i)
		if expected != StatsOpTypeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e StatsOpType) String() string {
	name, _ := statsOpTypeMap[int32(e)]
	return name
}

func (e StatsOpType) ShortString() string {
	name, _ := statsOpTypeShortMap[int32(e)]
	return name
}

func (e StatsOpType) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range StatsOpTypeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *StatsOpType) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = StatsOpType(t.Value)
	return nil
}

// LimitsV2EntryExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type LimitsV2EntryExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LimitsV2EntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LimitsV2EntryExt
func (u LimitsV2EntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewLimitsV2EntryExt creates a new  LimitsV2EntryExt.
func NewLimitsV2EntryExt(v LedgerVersion, value interface{}) (result LimitsV2EntryExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// LimitsV2Entry is an XDR Struct defines as:
//
//   struct LimitsV2Entry
//    {
//        uint64      id;
//        AccountType *accountType;
//        AccountID   *accountID;
//        StatsOpType statsOpType;
//        AssetCode   assetCode;
//        bool        isConvertNeeded;
//
//        uint64 dailyOut;
//        uint64 weeklyOut;
//        uint64 monthlyOut;
//        uint64 annualOut;
//
//         // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type LimitsV2Entry struct {
	Id              Uint64           `json:"id,omitempty"`
	AccountType     *AccountType     `json:"accountType,omitempty"`
	AccountId       *AccountId       `json:"accountID,omitempty"`
	StatsOpType     StatsOpType      `json:"statsOpType,omitempty"`
	AssetCode       AssetCode        `json:"assetCode,omitempty"`
	IsConvertNeeded bool             `json:"isConvertNeeded,omitempty"`
	DailyOut        Uint64           `json:"dailyOut,omitempty"`
	WeeklyOut       Uint64           `json:"weeklyOut,omitempty"`
	MonthlyOut      Uint64           `json:"monthlyOut,omitempty"`
	AnnualOut       Uint64           `json:"annualOut,omitempty"`
	Ext             LimitsV2EntryExt `json:"ext,omitempty"`
}

// OfferEntryExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type OfferEntryExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u OfferEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of OfferEntryExt
func (u OfferEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewOfferEntryExt creates a new  OfferEntryExt.
func NewOfferEntryExt(v LedgerVersion, value interface{}) (result OfferEntryExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// OfferEntry is an XDR Struct defines as:
//
//   struct OfferEntry
//    {
//        uint64 offerID;
//    	uint64 orderBookID;
//    	AccountID ownerID;
//    	bool isBuy;
//        AssetCode base; // A
//        AssetCode quote;  // B
//    	BalanceID baseBalance;
//    	BalanceID quoteBalance;
//        int64 baseAmount;
//    	int64 quoteAmount;
//    	uint64 createdAt;
//    	int64 fee;
//
//        int64 percentFee;
//
//    	// price of A in terms of B
//        int64 price;
//
//        // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type OfferEntry struct {
	OfferId      Uint64        `json:"offerID,omitempty"`
	OrderBookId  Uint64        `json:"orderBookID,omitempty"`
	OwnerId      AccountId     `json:"ownerID,omitempty"`
	IsBuy        bool          `json:"isBuy,omitempty"`
	Base         AssetCode     `json:"base,omitempty"`
	Quote        AssetCode     `json:"quote,omitempty"`
	BaseBalance  BalanceId     `json:"baseBalance,omitempty"`
	QuoteBalance BalanceId     `json:"quoteBalance,omitempty"`
	BaseAmount   Int64         `json:"baseAmount,omitempty"`
	QuoteAmount  Int64         `json:"quoteAmount,omitempty"`
	CreatedAt    Uint64        `json:"createdAt,omitempty"`
	Fee          Int64         `json:"fee,omitempty"`
	PercentFee   Int64         `json:"percentFee,omitempty"`
	Price        Int64         `json:"price,omitempty"`
	Ext          OfferEntryExt `json:"ext,omitempty"`
}

// PendingStatisticsEntryExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type PendingStatisticsEntryExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u PendingStatisticsEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of PendingStatisticsEntryExt
func (u PendingStatisticsEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewPendingStatisticsEntryExt creates a new  PendingStatisticsEntryExt.
func NewPendingStatisticsEntryExt(v LedgerVersion, value interface{}) (result PendingStatisticsEntryExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// PendingStatisticsEntry is an XDR Struct defines as:
//
//   struct PendingStatisticsEntry
//    {
//        uint64 statisticsID;
//        uint64 requestID;
//        uint64 amount;
//
//        // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type PendingStatisticsEntry struct {
	StatisticsId Uint64                    `json:"statisticsID,omitempty"`
	RequestId    Uint64                    `json:"requestID,omitempty"`
	Amount       Uint64                    `json:"amount,omitempty"`
	Ext          PendingStatisticsEntryExt `json:"ext,omitempty"`
}

// ReferenceEntryExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type ReferenceEntryExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ReferenceEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ReferenceEntryExt
func (u ReferenceEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewReferenceEntryExt creates a new  ReferenceEntryExt.
func NewReferenceEntryExt(v LedgerVersion, value interface{}) (result ReferenceEntryExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ReferenceEntry is an XDR Struct defines as:
//
//   struct ReferenceEntry
//    {
//    	AccountID sender;
//        string64 reference;
//
//    	// reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type ReferenceEntry struct {
	Sender    AccountId         `json:"sender,omitempty"`
	Reference String64          `json:"reference,omitempty"`
	Ext       ReferenceEntryExt `json:"ext,omitempty"`
}

// ReviewableRequestType is an XDR Enum defines as:
//
//   enum ReviewableRequestType
//    {
//        ASSET_CREATE = 0,
//    	ASSET_UPDATE = 1,
//    	PRE_ISSUANCE_CREATE = 2,
//    	ISSUANCE_CREATE = 3,
//    	WITHDRAW = 4,
//    	SALE = 5,
//    	LIMITS_UPDATE = 6,
//    	TWO_STEP_WITHDRAWAL = 7,
//        AML_ALERT = 8,
//    	UPDATE_KYC = 9,
//    	UPDATE_SALE_DETAILS = 10,
//    	UPDATE_PROMOTION = 11,
//    	UPDATE_SALE_END_TIME = 12,
//    	NONE = 13, // use this request type in ReviewRequestOp extended result if additional info is not required
//    	INVOICE = 14,
//    	CONTRACT = 15
//
//    };
//
type ReviewableRequestType int32

const (
	ReviewableRequestTypeAssetCreate       ReviewableRequestType = 0
	ReviewableRequestTypeAssetUpdate       ReviewableRequestType = 1
	ReviewableRequestTypePreIssuanceCreate ReviewableRequestType = 2
	ReviewableRequestTypeIssuanceCreate    ReviewableRequestType = 3
	ReviewableRequestTypeWithdraw          ReviewableRequestType = 4
	ReviewableRequestTypeSale              ReviewableRequestType = 5
	ReviewableRequestTypeLimitsUpdate      ReviewableRequestType = 6
	ReviewableRequestTypeTwoStepWithdrawal ReviewableRequestType = 7
	ReviewableRequestTypeAmlAlert          ReviewableRequestType = 8
	ReviewableRequestTypeUpdateKyc         ReviewableRequestType = 9
	ReviewableRequestTypeUpdateSaleDetails ReviewableRequestType = 10
	ReviewableRequestTypeUpdatePromotion   ReviewableRequestType = 11
	ReviewableRequestTypeUpdateSaleEndTime ReviewableRequestType = 12
	ReviewableRequestTypeNone              ReviewableRequestType = 13
	ReviewableRequestTypeInvoice           ReviewableRequestType = 14
	ReviewableRequestTypeContract          ReviewableRequestType = 15
)

var ReviewableRequestTypeAll = []ReviewableRequestType{
	ReviewableRequestTypeAssetCreate,
	ReviewableRequestTypeAssetUpdate,
	ReviewableRequestTypePreIssuanceCreate,
	ReviewableRequestTypeIssuanceCreate,
	ReviewableRequestTypeWithdraw,
	ReviewableRequestTypeSale,
	ReviewableRequestTypeLimitsUpdate,
	ReviewableRequestTypeTwoStepWithdrawal,
	ReviewableRequestTypeAmlAlert,
	ReviewableRequestTypeUpdateKyc,
	ReviewableRequestTypeUpdateSaleDetails,
	ReviewableRequestTypeUpdatePromotion,
	ReviewableRequestTypeUpdateSaleEndTime,
	ReviewableRequestTypeNone,
	ReviewableRequestTypeInvoice,
	ReviewableRequestTypeContract,
}

var reviewableRequestTypeMap = map[int32]string{
	0:  "ReviewableRequestTypeAssetCreate",
	1:  "ReviewableRequestTypeAssetUpdate",
	2:  "ReviewableRequestTypePreIssuanceCreate",
	3:  "ReviewableRequestTypeIssuanceCreate",
	4:  "ReviewableRequestTypeWithdraw",
	5:  "ReviewableRequestTypeSale",
	6:  "ReviewableRequestTypeLimitsUpdate",
	7:  "ReviewableRequestTypeTwoStepWithdrawal",
	8:  "ReviewableRequestTypeAmlAlert",
	9:  "ReviewableRequestTypeUpdateKyc",
	10: "ReviewableRequestTypeUpdateSaleDetails",
	11: "ReviewableRequestTypeUpdatePromotion",
	12: "ReviewableRequestTypeUpdateSaleEndTime",
	13: "ReviewableRequestTypeNone",
	14: "ReviewableRequestTypeInvoice",
	15: "ReviewableRequestTypeContract",
}

var reviewableRequestTypeShortMap = map[int32]string{
	0:  "asset_create",
	1:  "asset_update",
	2:  "pre_issuance_create",
	3:  "issuance_create",
	4:  "withdraw",
	5:  "sale",
	6:  "limits_update",
	7:  "two_step_withdrawal",
	8:  "aml_alert",
	9:  "update_kyc",
	10: "update_sale_details",
	11: "update_promotion",
	12: "update_sale_end_time",
	13: "none",
	14: "invoice",
	15: "contract",
}

var reviewableRequestTypeRevMap = map[string]int32{
	"ReviewableRequestTypeAssetCreate":       0,
	"ReviewableRequestTypeAssetUpdate":       1,
	"ReviewableRequestTypePreIssuanceCreate": 2,
	"ReviewableRequestTypeIssuanceCreate":    3,
	"ReviewableRequestTypeWithdraw":          4,
	"ReviewableRequestTypeSale":              5,
	"ReviewableRequestTypeLimitsUpdate":      6,
	"ReviewableRequestTypeTwoStepWithdrawal": 7,
	"ReviewableRequestTypeAmlAlert":          8,
	"ReviewableRequestTypeUpdateKyc":         9,
	"ReviewableRequestTypeUpdateSaleDetails": 10,
	"ReviewableRequestTypeUpdatePromotion":   11,
	"ReviewableRequestTypeUpdateSaleEndTime": 12,
	"ReviewableRequestTypeNone":              13,
	"ReviewableRequestTypeInvoice":           14,
	"ReviewableRequestTypeContract":          15,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ReviewableRequestType
func (e ReviewableRequestType) ValidEnum(v int32) bool {
	_, ok := reviewableRequestTypeMap[v]
	return ok
}
func (e ReviewableRequestType) isFlag() bool {
	for i := len(ReviewableRequestTypeAll) - 1; i >= 0; i-- {
		expected := ReviewableRequestType(2) << uint64(len(ReviewableRequestTypeAll)-1) >> uint64(len(ReviewableRequestTypeAll)-i)
		if expected != ReviewableRequestTypeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ReviewableRequestType) String() string {
	name, _ := reviewableRequestTypeMap[int32(e)]
	return name
}

func (e ReviewableRequestType) ShortString() string {
	name, _ := reviewableRequestTypeShortMap[int32(e)]
	return name
}

func (e ReviewableRequestType) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range ReviewableRequestTypeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ReviewableRequestType) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ReviewableRequestType(t.Value)
	return nil
}

// TasksExtExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type TasksExtExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u TasksExtExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of TasksExtExt
func (u TasksExtExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewTasksExtExt creates a new  TasksExtExt.
func NewTasksExtExt(v LedgerVersion, value interface{}) (result TasksExtExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// TasksExt is an XDR Struct defines as:
//
//   struct TasksExt {
//        // Tasks are represented by a bitmask
//        uint32 allTasks;
//        uint32 pendingTasks;
//
//        // External details vector consists of comments written by request reviewers
//        longstring externalDetails<>;
//
//        // Reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type TasksExt struct {
	AllTasks        Uint32       `json:"allTasks,omitempty"`
	PendingTasks    Uint32       `json:"pendingTasks,omitempty"`
	ExternalDetails []Longstring `json:"externalDetails,omitempty"`
	Ext             TasksExtExt  `json:"ext,omitempty"`
}

// ReviewableRequestEntryBody is an XDR NestedUnion defines as:
//
//   union switch (ReviewableRequestType type) {
//    		case ASSET_CREATE:
//    			AssetCreationRequest assetCreationRequest;
//    		case ASSET_UPDATE:
//    			AssetUpdateRequest assetUpdateRequest;
//    		case PRE_ISSUANCE_CREATE:
//    			PreIssuanceRequest preIssuanceRequest;
//    		case ISSUANCE_CREATE:
//    			IssuanceRequest issuanceRequest;
//    		case WITHDRAW:
//    			WithdrawalRequest withdrawalRequest;
//    		case SALE:
//    			SaleCreationRequest saleCreationRequest;
//            case LIMITS_UPDATE:
//                LimitsUpdateRequest limitsUpdateRequest;
//    		case TWO_STEP_WITHDRAWAL:
//    			WithdrawalRequest twoStepWithdrawalRequest;
//            case AML_ALERT:
//                AMLAlertRequest amlAlertRequest;
//            case UPDATE_KYC:
//                UpdateKYCRequest updateKYCRequest;
//            case UPDATE_SALE_DETAILS:
//                UpdateSaleDetailsRequest updateSaleDetailsRequest;
//            case UPDATE_PROMOTION:
//                PromotionUpdateRequest promotionUpdateRequest;
//            case INVOICE:
//                InvoiceRequest invoiceRequest;
//            case UPDATE_SALE_END_TIME:
//                UpdateSaleEndTimeRequest updateSaleEndTimeRequest;
//            case CONTRACT:
//                ContractRequest contractRequest;
//    	}
//
type ReviewableRequestEntryBody struct {
	Type                     ReviewableRequestType     `json:"type,omitempty"`
	AssetCreationRequest     *AssetCreationRequest     `json:"assetCreationRequest,omitempty"`
	AssetUpdateRequest       *AssetUpdateRequest       `json:"assetUpdateRequest,omitempty"`
	PreIssuanceRequest       *PreIssuanceRequest       `json:"preIssuanceRequest,omitempty"`
	IssuanceRequest          *IssuanceRequest          `json:"issuanceRequest,omitempty"`
	WithdrawalRequest        *WithdrawalRequest        `json:"withdrawalRequest,omitempty"`
	SaleCreationRequest      *SaleCreationRequest      `json:"saleCreationRequest,omitempty"`
	LimitsUpdateRequest      *LimitsUpdateRequest      `json:"limitsUpdateRequest,omitempty"`
	TwoStepWithdrawalRequest *WithdrawalRequest        `json:"twoStepWithdrawalRequest,omitempty"`
	AmlAlertRequest          *AmlAlertRequest          `json:"amlAlertRequest,omitempty"`
	UpdateKycRequest         *UpdateKycRequest         `json:"updateKYCRequest,omitempty"`
	UpdateSaleDetailsRequest *UpdateSaleDetailsRequest `json:"updateSaleDetailsRequest,omitempty"`
	PromotionUpdateRequest   *PromotionUpdateRequest   `json:"promotionUpdateRequest,omitempty"`
	InvoiceRequest           *InvoiceRequest           `json:"invoiceRequest,omitempty"`
	UpdateSaleEndTimeRequest *UpdateSaleEndTimeRequest `json:"updateSaleEndTimeRequest,omitempty"`
	ContractRequest          *ContractRequest          `json:"contractRequest,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ReviewableRequestEntryBody) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ReviewableRequestEntryBody
func (u ReviewableRequestEntryBody) ArmForSwitch(sw int32) (string, bool) {
	switch ReviewableRequestType(sw) {
	case ReviewableRequestTypeAssetCreate:
		return "AssetCreationRequest", true
	case ReviewableRequestTypeAssetUpdate:
		return "AssetUpdateRequest", true
	case ReviewableRequestTypePreIssuanceCreate:
		return "PreIssuanceRequest", true
	case ReviewableRequestTypeIssuanceCreate:
		return "IssuanceRequest", true
	case ReviewableRequestTypeWithdraw:
		return "WithdrawalRequest", true
	case ReviewableRequestTypeSale:
		return "SaleCreationRequest", true
	case ReviewableRequestTypeLimitsUpdate:
		return "LimitsUpdateRequest", true
	case ReviewableRequestTypeTwoStepWithdrawal:
		return "TwoStepWithdrawalRequest", true
	case ReviewableRequestTypeAmlAlert:
		return "AmlAlertRequest", true
	case ReviewableRequestTypeUpdateKyc:
		return "UpdateKycRequest", true
	case ReviewableRequestTypeUpdateSaleDetails:
		return "UpdateSaleDetailsRequest", true
	case ReviewableRequestTypeUpdatePromotion:
		return "PromotionUpdateRequest", true
	case ReviewableRequestTypeInvoice:
		return "InvoiceRequest", true
	case ReviewableRequestTypeUpdateSaleEndTime:
		return "UpdateSaleEndTimeRequest", true
	case ReviewableRequestTypeContract:
		return "ContractRequest", true
	}
	return "-", false
}

// NewReviewableRequestEntryBody creates a new  ReviewableRequestEntryBody.
func NewReviewableRequestEntryBody(aType ReviewableRequestType, value interface{}) (result ReviewableRequestEntryBody, err error) {
	result.Type = aType
	switch ReviewableRequestType(aType) {
	case ReviewableRequestTypeAssetCreate:
		tv, ok := value.(AssetCreationRequest)
		if !ok {
			err = fmt.Errorf("invalid value, must be AssetCreationRequest")
			return
		}
		result.AssetCreationRequest = &tv
	case ReviewableRequestTypeAssetUpdate:
		tv, ok := value.(AssetUpdateRequest)
		if !ok {
			err = fmt.Errorf("invalid value, must be AssetUpdateRequest")
			return
		}
		result.AssetUpdateRequest = &tv
	case ReviewableRequestTypePreIssuanceCreate:
		tv, ok := value.(PreIssuanceRequest)
		if !ok {
			err = fmt.Errorf("invalid value, must be PreIssuanceRequest")
			return
		}
		result.PreIssuanceRequest = &tv
	case ReviewableRequestTypeIssuanceCreate:
		tv, ok := value.(IssuanceRequest)
		if !ok {
			err = fmt.Errorf("invalid value, must be IssuanceRequest")
			return
		}
		result.IssuanceRequest = &tv
	case ReviewableRequestTypeWithdraw:
		tv, ok := value.(WithdrawalRequest)
		if !ok {
			err = fmt.Errorf("invalid value, must be WithdrawalRequest")
			return
		}
		result.WithdrawalRequest = &tv
	case ReviewableRequestTypeSale:
		tv, ok := value.(SaleCreationRequest)
		if !ok {
			err = fmt.Errorf("invalid value, must be SaleCreationRequest")
			return
		}
		result.SaleCreationRequest = &tv
	case ReviewableRequestTypeLimitsUpdate:
		tv, ok := value.(LimitsUpdateRequest)
		if !ok {
			err = fmt.Errorf("invalid value, must be LimitsUpdateRequest")
			return
		}
		result.LimitsUpdateRequest = &tv
	case ReviewableRequestTypeTwoStepWithdrawal:
		tv, ok := value.(WithdrawalRequest)
		if !ok {
			err = fmt.Errorf("invalid value, must be WithdrawalRequest")
			return
		}
		result.TwoStepWithdrawalRequest = &tv
	case ReviewableRequestTypeAmlAlert:
		tv, ok := value.(AmlAlertRequest)
		if !ok {
			err = fmt.Errorf("invalid value, must be AmlAlertRequest")
			return
		}
		result.AmlAlertRequest = &tv
	case ReviewableRequestTypeUpdateKyc:
		tv, ok := value.(UpdateKycRequest)
		if !ok {
			err = fmt.Errorf("invalid value, must be UpdateKycRequest")
			return
		}
		result.UpdateKycRequest = &tv
	case ReviewableRequestTypeUpdateSaleDetails:
		tv, ok := value.(UpdateSaleDetailsRequest)
		if !ok {
			err = fmt.Errorf("invalid value, must be UpdateSaleDetailsRequest")
			return
		}
		result.UpdateSaleDetailsRequest = &tv
	case ReviewableRequestTypeUpdatePromotion:
		tv, ok := value.(PromotionUpdateRequest)
		if !ok {
			err = fmt.Errorf("invalid value, must be PromotionUpdateRequest")
			return
		}
		result.PromotionUpdateRequest = &tv
	case ReviewableRequestTypeInvoice:
		tv, ok := value.(InvoiceRequest)
		if !ok {
			err = fmt.Errorf("invalid value, must be InvoiceRequest")
			return
		}
		result.InvoiceRequest = &tv
	case ReviewableRequestTypeUpdateSaleEndTime:
		tv, ok := value.(UpdateSaleEndTimeRequest)
		if !ok {
			err = fmt.Errorf("invalid value, must be UpdateSaleEndTimeRequest")
			return
		}
		result.UpdateSaleEndTimeRequest = &tv
	case ReviewableRequestTypeContract:
		tv, ok := value.(ContractRequest)
		if !ok {
			err = fmt.Errorf("invalid value, must be ContractRequest")
			return
		}
		result.ContractRequest = &tv
	}
	return
}

// MustAssetCreationRequest retrieves the AssetCreationRequest value from the union,
// panicing if the value is not set.
func (u ReviewableRequestEntryBody) MustAssetCreationRequest() AssetCreationRequest {
	val, ok := u.GetAssetCreationRequest()

	if !ok {
		panic("arm AssetCreationRequest is not set")
	}

	return val
}

// GetAssetCreationRequest retrieves the AssetCreationRequest value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ReviewableRequestEntryBody) GetAssetCreationRequest() (result AssetCreationRequest, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "AssetCreationRequest" {
		result = *u.AssetCreationRequest
		ok = true
	}

	return
}

// MustAssetUpdateRequest retrieves the AssetUpdateRequest value from the union,
// panicing if the value is not set.
func (u ReviewableRequestEntryBody) MustAssetUpdateRequest() AssetUpdateRequest {
	val, ok := u.GetAssetUpdateRequest()

	if !ok {
		panic("arm AssetUpdateRequest is not set")
	}

	return val
}

// GetAssetUpdateRequest retrieves the AssetUpdateRequest value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ReviewableRequestEntryBody) GetAssetUpdateRequest() (result AssetUpdateRequest, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "AssetUpdateRequest" {
		result = *u.AssetUpdateRequest
		ok = true
	}

	return
}

// MustPreIssuanceRequest retrieves the PreIssuanceRequest value from the union,
// panicing if the value is not set.
func (u ReviewableRequestEntryBody) MustPreIssuanceRequest() PreIssuanceRequest {
	val, ok := u.GetPreIssuanceRequest()

	if !ok {
		panic("arm PreIssuanceRequest is not set")
	}

	return val
}

// GetPreIssuanceRequest retrieves the PreIssuanceRequest value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ReviewableRequestEntryBody) GetPreIssuanceRequest() (result PreIssuanceRequest, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "PreIssuanceRequest" {
		result = *u.PreIssuanceRequest
		ok = true
	}

	return
}

// MustIssuanceRequest retrieves the IssuanceRequest value from the union,
// panicing if the value is not set.
func (u ReviewableRequestEntryBody) MustIssuanceRequest() IssuanceRequest {
	val, ok := u.GetIssuanceRequest()

	if !ok {
		panic("arm IssuanceRequest is not set")
	}

	return val
}

// GetIssuanceRequest retrieves the IssuanceRequest value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ReviewableRequestEntryBody) GetIssuanceRequest() (result IssuanceRequest, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "IssuanceRequest" {
		result = *u.IssuanceRequest
		ok = true
	}

	return
}

// MustWithdrawalRequest retrieves the WithdrawalRequest value from the union,
// panicing if the value is not set.
func (u ReviewableRequestEntryBody) MustWithdrawalRequest() WithdrawalRequest {
	val, ok := u.GetWithdrawalRequest()

	if !ok {
		panic("arm WithdrawalRequest is not set")
	}

	return val
}

// GetWithdrawalRequest retrieves the WithdrawalRequest value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ReviewableRequestEntryBody) GetWithdrawalRequest() (result WithdrawalRequest, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "WithdrawalRequest" {
		result = *u.WithdrawalRequest
		ok = true
	}

	return
}

// MustSaleCreationRequest retrieves the SaleCreationRequest value from the union,
// panicing if the value is not set.
func (u ReviewableRequestEntryBody) MustSaleCreationRequest() SaleCreationRequest {
	val, ok := u.GetSaleCreationRequest()

	if !ok {
		panic("arm SaleCreationRequest is not set")
	}

	return val
}

// GetSaleCreationRequest retrieves the SaleCreationRequest value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ReviewableRequestEntryBody) GetSaleCreationRequest() (result SaleCreationRequest, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "SaleCreationRequest" {
		result = *u.SaleCreationRequest
		ok = true
	}

	return
}

// MustLimitsUpdateRequest retrieves the LimitsUpdateRequest value from the union,
// panicing if the value is not set.
func (u ReviewableRequestEntryBody) MustLimitsUpdateRequest() LimitsUpdateRequest {
	val, ok := u.GetLimitsUpdateRequest()

	if !ok {
		panic("arm LimitsUpdateRequest is not set")
	}

	return val
}

// GetLimitsUpdateRequest retrieves the LimitsUpdateRequest value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ReviewableRequestEntryBody) GetLimitsUpdateRequest() (result LimitsUpdateRequest, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "LimitsUpdateRequest" {
		result = *u.LimitsUpdateRequest
		ok = true
	}

	return
}

// MustTwoStepWithdrawalRequest retrieves the TwoStepWithdrawalRequest value from the union,
// panicing if the value is not set.
func (u ReviewableRequestEntryBody) MustTwoStepWithdrawalRequest() WithdrawalRequest {
	val, ok := u.GetTwoStepWithdrawalRequest()

	if !ok {
		panic("arm TwoStepWithdrawalRequest is not set")
	}

	return val
}

// GetTwoStepWithdrawalRequest retrieves the TwoStepWithdrawalRequest value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ReviewableRequestEntryBody) GetTwoStepWithdrawalRequest() (result WithdrawalRequest, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "TwoStepWithdrawalRequest" {
		result = *u.TwoStepWithdrawalRequest
		ok = true
	}

	return
}

// MustAmlAlertRequest retrieves the AmlAlertRequest value from the union,
// panicing if the value is not set.
func (u ReviewableRequestEntryBody) MustAmlAlertRequest() AmlAlertRequest {
	val, ok := u.GetAmlAlertRequest()

	if !ok {
		panic("arm AmlAlertRequest is not set")
	}

	return val
}

// GetAmlAlertRequest retrieves the AmlAlertRequest value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ReviewableRequestEntryBody) GetAmlAlertRequest() (result AmlAlertRequest, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "AmlAlertRequest" {
		result = *u.AmlAlertRequest
		ok = true
	}

	return
}

// MustUpdateKycRequest retrieves the UpdateKycRequest value from the union,
// panicing if the value is not set.
func (u ReviewableRequestEntryBody) MustUpdateKycRequest() UpdateKycRequest {
	val, ok := u.GetUpdateKycRequest()

	if !ok {
		panic("arm UpdateKycRequest is not set")
	}

	return val
}

// GetUpdateKycRequest retrieves the UpdateKycRequest value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ReviewableRequestEntryBody) GetUpdateKycRequest() (result UpdateKycRequest, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "UpdateKycRequest" {
		result = *u.UpdateKycRequest
		ok = true
	}

	return
}

// MustUpdateSaleDetailsRequest retrieves the UpdateSaleDetailsRequest value from the union,
// panicing if the value is not set.
func (u ReviewableRequestEntryBody) MustUpdateSaleDetailsRequest() UpdateSaleDetailsRequest {
	val, ok := u.GetUpdateSaleDetailsRequest()

	if !ok {
		panic("arm UpdateSaleDetailsRequest is not set")
	}

	return val
}

// GetUpdateSaleDetailsRequest retrieves the UpdateSaleDetailsRequest value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ReviewableRequestEntryBody) GetUpdateSaleDetailsRequest() (result UpdateSaleDetailsRequest, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "UpdateSaleDetailsRequest" {
		result = *u.UpdateSaleDetailsRequest
		ok = true
	}

	return
}

// MustPromotionUpdateRequest retrieves the PromotionUpdateRequest value from the union,
// panicing if the value is not set.
func (u ReviewableRequestEntryBody) MustPromotionUpdateRequest() PromotionUpdateRequest {
	val, ok := u.GetPromotionUpdateRequest()

	if !ok {
		panic("arm PromotionUpdateRequest is not set")
	}

	return val
}

// GetPromotionUpdateRequest retrieves the PromotionUpdateRequest value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ReviewableRequestEntryBody) GetPromotionUpdateRequest() (result PromotionUpdateRequest, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "PromotionUpdateRequest" {
		result = *u.PromotionUpdateRequest
		ok = true
	}

	return
}

// MustInvoiceRequest retrieves the InvoiceRequest value from the union,
// panicing if the value is not set.
func (u ReviewableRequestEntryBody) MustInvoiceRequest() InvoiceRequest {
	val, ok := u.GetInvoiceRequest()

	if !ok {
		panic("arm InvoiceRequest is not set")
	}

	return val
}

// GetInvoiceRequest retrieves the InvoiceRequest value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ReviewableRequestEntryBody) GetInvoiceRequest() (result InvoiceRequest, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "InvoiceRequest" {
		result = *u.InvoiceRequest
		ok = true
	}

	return
}

// MustUpdateSaleEndTimeRequest retrieves the UpdateSaleEndTimeRequest value from the union,
// panicing if the value is not set.
func (u ReviewableRequestEntryBody) MustUpdateSaleEndTimeRequest() UpdateSaleEndTimeRequest {
	val, ok := u.GetUpdateSaleEndTimeRequest()

	if !ok {
		panic("arm UpdateSaleEndTimeRequest is not set")
	}

	return val
}

// GetUpdateSaleEndTimeRequest retrieves the UpdateSaleEndTimeRequest value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ReviewableRequestEntryBody) GetUpdateSaleEndTimeRequest() (result UpdateSaleEndTimeRequest, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "UpdateSaleEndTimeRequest" {
		result = *u.UpdateSaleEndTimeRequest
		ok = true
	}

	return
}

// MustContractRequest retrieves the ContractRequest value from the union,
// panicing if the value is not set.
func (u ReviewableRequestEntryBody) MustContractRequest() ContractRequest {
	val, ok := u.GetContractRequest()

	if !ok {
		panic("arm ContractRequest is not set")
	}

	return val
}

// GetContractRequest retrieves the ContractRequest value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ReviewableRequestEntryBody) GetContractRequest() (result ContractRequest, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ContractRequest" {
		result = *u.ContractRequest
		ok = true
	}

	return
}

// ReviewableRequestEntryExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        case ADD_TASKS_TO_REVIEWABLE_REQUEST:
//            TasksExt tasksExt;
//        }
//
type ReviewableRequestEntryExt struct {
	V        LedgerVersion `json:"v,omitempty"`
	TasksExt *TasksExt     `json:"tasksExt,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ReviewableRequestEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ReviewableRequestEntryExt
func (u ReviewableRequestEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	case LedgerVersionAddTasksToReviewableRequest:
		return "TasksExt", true
	}
	return "-", false
}

// NewReviewableRequestEntryExt creates a new  ReviewableRequestEntryExt.
func NewReviewableRequestEntryExt(v LedgerVersion, value interface{}) (result ReviewableRequestEntryExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	case LedgerVersionAddTasksToReviewableRequest:
		tv, ok := value.(TasksExt)
		if !ok {
			err = fmt.Errorf("invalid value, must be TasksExt")
			return
		}
		result.TasksExt = &tv
	}
	return
}

// MustTasksExt retrieves the TasksExt value from the union,
// panicing if the value is not set.
func (u ReviewableRequestEntryExt) MustTasksExt() TasksExt {
	val, ok := u.GetTasksExt()

	if !ok {
		panic("arm TasksExt is not set")
	}

	return val
}

// GetTasksExt retrieves the TasksExt value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ReviewableRequestEntryExt) GetTasksExt() (result TasksExt, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.V))

	if armName == "TasksExt" {
		result = *u.TasksExt
		ok = true
	}

	return
}

// ReviewableRequestEntry is an XDR Struct defines as:
//
//   struct ReviewableRequestEntry {
//    	uint64 requestID;
//    	Hash hash; // hash of the request body
//    	AccountID requestor;
//    	longstring rejectReason;
//    	AccountID reviewer;
//    	string64* reference; // reference for request which will act as an unique key for the request (will reject request with the same reference from same requestor)
//    	int64 createdAt; // when request was created
//
//    	union switch (ReviewableRequestType type) {
//    		case ASSET_CREATE:
//    			AssetCreationRequest assetCreationRequest;
//    		case ASSET_UPDATE:
//    			AssetUpdateRequest assetUpdateRequest;
//    		case PRE_ISSUANCE_CREATE:
//    			PreIssuanceRequest preIssuanceRequest;
//    		case ISSUANCE_CREATE:
//    			IssuanceRequest issuanceRequest;
//    		case WITHDRAW:
//    			WithdrawalRequest withdrawalRequest;
//    		case SALE:
//    			SaleCreationRequest saleCreationRequest;
//            case LIMITS_UPDATE:
//                LimitsUpdateRequest limitsUpdateRequest;
//    		case TWO_STEP_WITHDRAWAL:
//    			WithdrawalRequest twoStepWithdrawalRequest;
//            case AML_ALERT:
//                AMLAlertRequest amlAlertRequest;
//            case UPDATE_KYC:
//                UpdateKYCRequest updateKYCRequest;
//            case UPDATE_SALE_DETAILS:
//                UpdateSaleDetailsRequest updateSaleDetailsRequest;
//            case UPDATE_PROMOTION:
//                PromotionUpdateRequest promotionUpdateRequest;
//            case INVOICE:
//                InvoiceRequest invoiceRequest;
//            case UPDATE_SALE_END_TIME:
//                UpdateSaleEndTimeRequest updateSaleEndTimeRequest;
//            case CONTRACT:
//                ContractRequest contractRequest;
//    	} body;
//
//    	// reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        case ADD_TASKS_TO_REVIEWABLE_REQUEST:
//            TasksExt tasksExt;
//        }
//        ext;
//    };
//
type ReviewableRequestEntry struct {
	RequestId    Uint64                     `json:"requestID,omitempty"`
	Hash         Hash                       `json:"hash,omitempty"`
	Requestor    AccountId                  `json:"requestor,omitempty"`
	RejectReason Longstring                 `json:"rejectReason,omitempty"`
	Reviewer     AccountId                  `json:"reviewer,omitempty"`
	Reference    *String64                  `json:"reference,omitempty"`
	CreatedAt    Int64                      `json:"createdAt,omitempty"`
	Body         ReviewableRequestEntryBody `json:"body,omitempty"`
	Ext          ReviewableRequestEntryExt  `json:"ext,omitempty"`
}

// SaleAnteEntryExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type SaleAnteEntryExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u SaleAnteEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of SaleAnteEntryExt
func (u SaleAnteEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewSaleAnteEntryExt creates a new  SaleAnteEntryExt.
func NewSaleAnteEntryExt(v LedgerVersion, value interface{}) (result SaleAnteEntryExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// SaleAnteEntry is an XDR Struct defines as:
//
//   struct SaleAnteEntry
//    {
//        uint64 saleID;
//        BalanceID participantBalanceID;
//        uint64 amount; // amount to be locked from participant balance
//
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type SaleAnteEntry struct {
	SaleId               Uint64           `json:"saleID,omitempty"`
	ParticipantBalanceId BalanceId        `json:"participantBalanceID,omitempty"`
	Amount               Uint64           `json:"amount,omitempty"`
	Ext                  SaleAnteEntryExt `json:"ext,omitempty"`
}

// SaleState is an XDR Enum defines as:
//
//   enum SaleState {
//    	NONE = 0, // default state
//    	VOTING = 1, // not allowed to invest
//    	PROMOTION = 2 // not allowed to invest, but allowed to change all the details
//    };
//
type SaleState int32

const (
	SaleStateNone      SaleState = 0
	SaleStateVoting    SaleState = 1
	SaleStatePromotion SaleState = 2
)

var SaleStateAll = []SaleState{
	SaleStateNone,
	SaleStateVoting,
	SaleStatePromotion,
}

var saleStateMap = map[int32]string{
	0: "SaleStateNone",
	1: "SaleStateVoting",
	2: "SaleStatePromotion",
}

var saleStateShortMap = map[int32]string{
	0: "none",
	1: "voting",
	2: "promotion",
}

var saleStateRevMap = map[string]int32{
	"SaleStateNone":      0,
	"SaleStateVoting":    1,
	"SaleStatePromotion": 2,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for SaleState
func (e SaleState) ValidEnum(v int32) bool {
	_, ok := saleStateMap[v]
	return ok
}
func (e SaleState) isFlag() bool {
	for i := len(SaleStateAll) - 1; i >= 0; i-- {
		expected := SaleState(2) << uint64(len(SaleStateAll)-1) >> uint64(len(SaleStateAll)-i)
		if expected != SaleStateAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e SaleState) String() string {
	name, _ := saleStateMap[int32(e)]
	return name
}

func (e SaleState) ShortString() string {
	name, _ := saleStateShortMap[int32(e)]
	return name
}

func (e SaleState) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range SaleStateAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *SaleState) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = SaleState(t.Value)
	return nil
}

// SaleType is an XDR Enum defines as:
//
//   enum SaleType {
//    	BASIC_SALE = 1, // sale creator specifies price for each quote asset
//    	CROWD_FUNDING = 2, // sale creator does not specify price,
//    	                  // price is defined on sale close based on amount of base asset to be sold and amount of quote assets collected
//        FIXED_PRICE=3
//    };
//
type SaleType int32

const (
	SaleTypeBasicSale    SaleType = 1
	SaleTypeCrowdFunding SaleType = 2
	SaleTypeFixedPrice   SaleType = 3
)

var SaleTypeAll = []SaleType{
	SaleTypeBasicSale,
	SaleTypeCrowdFunding,
	SaleTypeFixedPrice,
}

var saleTypeMap = map[int32]string{
	1: "SaleTypeBasicSale",
	2: "SaleTypeCrowdFunding",
	3: "SaleTypeFixedPrice",
}

var saleTypeShortMap = map[int32]string{
	1: "basic_sale",
	2: "crowd_funding",
	3: "fixed_price",
}

var saleTypeRevMap = map[string]int32{
	"SaleTypeBasicSale":    1,
	"SaleTypeCrowdFunding": 2,
	"SaleTypeFixedPrice":   3,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for SaleType
func (e SaleType) ValidEnum(v int32) bool {
	_, ok := saleTypeMap[v]
	return ok
}
func (e SaleType) isFlag() bool {
	for i := len(SaleTypeAll) - 1; i >= 0; i-- {
		expected := SaleType(2) << uint64(len(SaleTypeAll)-1) >> uint64(len(SaleTypeAll)-i)
		if expected != SaleTypeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e SaleType) String() string {
	name, _ := saleTypeMap[int32(e)]
	return name
}

func (e SaleType) ShortString() string {
	name, _ := saleTypeShortMap[int32(e)]
	return name
}

func (e SaleType) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range SaleTypeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *SaleType) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = SaleType(t.Value)
	return nil
}

// FixedPriceSaleExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type FixedPriceSaleExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u FixedPriceSaleExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of FixedPriceSaleExt
func (u FixedPriceSaleExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewFixedPriceSaleExt creates a new  FixedPriceSaleExt.
func NewFixedPriceSaleExt(v LedgerVersion, value interface{}) (result FixedPriceSaleExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// FixedPriceSale is an XDR Struct defines as:
//
//   struct FixedPriceSale {
//    	union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type FixedPriceSale struct {
	Ext FixedPriceSaleExt `json:"ext,omitempty"`
}

// CrowdFundingSaleExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type CrowdFundingSaleExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CrowdFundingSaleExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CrowdFundingSaleExt
func (u CrowdFundingSaleExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewCrowdFundingSaleExt creates a new  CrowdFundingSaleExt.
func NewCrowdFundingSaleExt(v LedgerVersion, value interface{}) (result CrowdFundingSaleExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// CrowdFundingSale is an XDR Struct defines as:
//
//   struct CrowdFundingSale {
//    	union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type CrowdFundingSale struct {
	Ext CrowdFundingSaleExt `json:"ext,omitempty"`
}

// BasicSaleExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type BasicSaleExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u BasicSaleExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of BasicSaleExt
func (u BasicSaleExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewBasicSaleExt creates a new  BasicSaleExt.
func NewBasicSaleExt(v LedgerVersion, value interface{}) (result BasicSaleExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// BasicSale is an XDR Struct defines as:
//
//   struct BasicSale {
//    	union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type BasicSale struct {
	Ext BasicSaleExt `json:"ext,omitempty"`
}

// SaleTypeExtTypedSale is an XDR NestedUnion defines as:
//
//   union switch (SaleType saleType)
//        {
//    	case BASIC_SALE:
//    		BasicSale basicSale;
//        case CROWD_FUNDING:
//            CrowdFundingSale crowdFundingSale;
//        case FIXED_PRICE:
//            FixedPriceSale fixedPriceSale;
//        }
//
type SaleTypeExtTypedSale struct {
	SaleType         SaleType          `json:"saleType,omitempty"`
	BasicSale        *BasicSale        `json:"basicSale,omitempty"`
	CrowdFundingSale *CrowdFundingSale `json:"crowdFundingSale,omitempty"`
	FixedPriceSale   *FixedPriceSale   `json:"fixedPriceSale,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u SaleTypeExtTypedSale) SwitchFieldName() string {
	return "SaleType"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of SaleTypeExtTypedSale
func (u SaleTypeExtTypedSale) ArmForSwitch(sw int32) (string, bool) {
	switch SaleType(sw) {
	case SaleTypeBasicSale:
		return "BasicSale", true
	case SaleTypeCrowdFunding:
		return "CrowdFundingSale", true
	case SaleTypeFixedPrice:
		return "FixedPriceSale", true
	}
	return "-", false
}

// NewSaleTypeExtTypedSale creates a new  SaleTypeExtTypedSale.
func NewSaleTypeExtTypedSale(saleType SaleType, value interface{}) (result SaleTypeExtTypedSale, err error) {
	result.SaleType = saleType
	switch SaleType(saleType) {
	case SaleTypeBasicSale:
		tv, ok := value.(BasicSale)
		if !ok {
			err = fmt.Errorf("invalid value, must be BasicSale")
			return
		}
		result.BasicSale = &tv
	case SaleTypeCrowdFunding:
		tv, ok := value.(CrowdFundingSale)
		if !ok {
			err = fmt.Errorf("invalid value, must be CrowdFundingSale")
			return
		}
		result.CrowdFundingSale = &tv
	case SaleTypeFixedPrice:
		tv, ok := value.(FixedPriceSale)
		if !ok {
			err = fmt.Errorf("invalid value, must be FixedPriceSale")
			return
		}
		result.FixedPriceSale = &tv
	}
	return
}

// MustBasicSale retrieves the BasicSale value from the union,
// panicing if the value is not set.
func (u SaleTypeExtTypedSale) MustBasicSale() BasicSale {
	val, ok := u.GetBasicSale()

	if !ok {
		panic("arm BasicSale is not set")
	}

	return val
}

// GetBasicSale retrieves the BasicSale value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u SaleTypeExtTypedSale) GetBasicSale() (result BasicSale, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.SaleType))

	if armName == "BasicSale" {
		result = *u.BasicSale
		ok = true
	}

	return
}

// MustCrowdFundingSale retrieves the CrowdFundingSale value from the union,
// panicing if the value is not set.
func (u SaleTypeExtTypedSale) MustCrowdFundingSale() CrowdFundingSale {
	val, ok := u.GetCrowdFundingSale()

	if !ok {
		panic("arm CrowdFundingSale is not set")
	}

	return val
}

// GetCrowdFundingSale retrieves the CrowdFundingSale value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u SaleTypeExtTypedSale) GetCrowdFundingSale() (result CrowdFundingSale, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.SaleType))

	if armName == "CrowdFundingSale" {
		result = *u.CrowdFundingSale
		ok = true
	}

	return
}

// MustFixedPriceSale retrieves the FixedPriceSale value from the union,
// panicing if the value is not set.
func (u SaleTypeExtTypedSale) MustFixedPriceSale() FixedPriceSale {
	val, ok := u.GetFixedPriceSale()

	if !ok {
		panic("arm FixedPriceSale is not set")
	}

	return val
}

// GetFixedPriceSale retrieves the FixedPriceSale value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u SaleTypeExtTypedSale) GetFixedPriceSale() (result FixedPriceSale, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.SaleType))

	if armName == "FixedPriceSale" {
		result = *u.FixedPriceSale
		ok = true
	}

	return
}

// SaleTypeExt is an XDR Struct defines as:
//
//   struct SaleTypeExt {
//    	union switch (SaleType saleType)
//        {
//    	case BASIC_SALE:
//    		BasicSale basicSale;
//        case CROWD_FUNDING:
//            CrowdFundingSale crowdFundingSale;
//        case FIXED_PRICE:
//            FixedPriceSale fixedPriceSale;
//        }
//        typedSale;
//    };
//
type SaleTypeExt struct {
	TypedSale SaleTypeExtTypedSale `json:"typedSale,omitempty"`
}

// StatableSaleExt is an XDR Struct defines as:
//
//   struct StatableSaleExt {
//    	SaleTypeExt saleTypeExt;
//    	SaleState state;
//    };
//
type StatableSaleExt struct {
	SaleTypeExt SaleTypeExt `json:"saleTypeExt,omitempty"`
	State       SaleState   `json:"state,omitempty"`
}

// SaleQuoteAssetExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type SaleQuoteAssetExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u SaleQuoteAssetExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of SaleQuoteAssetExt
func (u SaleQuoteAssetExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewSaleQuoteAssetExt creates a new  SaleQuoteAssetExt.
func NewSaleQuoteAssetExt(v LedgerVersion, value interface{}) (result SaleQuoteAssetExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// SaleQuoteAsset is an XDR Struct defines as:
//
//   struct SaleQuoteAsset {
//    	AssetCode quoteAsset; // asset in which participation will be accepted
//    	uint64 price; // price for 1 baseAsset in terms of quote asset
//    	BalanceID quoteBalance;
//    	uint64 currentCap; // current capitalization
//    	union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type SaleQuoteAsset struct {
	QuoteAsset   AssetCode         `json:"quoteAsset,omitempty"`
	Price        Uint64            `json:"price,omitempty"`
	QuoteBalance BalanceId         `json:"quoteBalance,omitempty"`
	CurrentCap   Uint64            `json:"currentCap,omitempty"`
	Ext          SaleQuoteAssetExt `json:"ext,omitempty"`
}

// SaleEntryExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//    	case TYPED_SALE:
//    		SaleTypeExt saleTypeExt;
//    	case STATABLE_SALES:
//    		StatableSaleExt statableSaleExt;
//        }
//
type SaleEntryExt struct {
	V               LedgerVersion    `json:"v,omitempty"`
	SaleTypeExt     *SaleTypeExt     `json:"saleTypeExt,omitempty"`
	StatableSaleExt *StatableSaleExt `json:"statableSaleExt,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u SaleEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of SaleEntryExt
func (u SaleEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	case LedgerVersionTypedSale:
		return "SaleTypeExt", true
	case LedgerVersionStatableSales:
		return "StatableSaleExt", true
	}
	return "-", false
}

// NewSaleEntryExt creates a new  SaleEntryExt.
func NewSaleEntryExt(v LedgerVersion, value interface{}) (result SaleEntryExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	case LedgerVersionTypedSale:
		tv, ok := value.(SaleTypeExt)
		if !ok {
			err = fmt.Errorf("invalid value, must be SaleTypeExt")
			return
		}
		result.SaleTypeExt = &tv
	case LedgerVersionStatableSales:
		tv, ok := value.(StatableSaleExt)
		if !ok {
			err = fmt.Errorf("invalid value, must be StatableSaleExt")
			return
		}
		result.StatableSaleExt = &tv
	}
	return
}

// MustSaleTypeExt retrieves the SaleTypeExt value from the union,
// panicing if the value is not set.
func (u SaleEntryExt) MustSaleTypeExt() SaleTypeExt {
	val, ok := u.GetSaleTypeExt()

	if !ok {
		panic("arm SaleTypeExt is not set")
	}

	return val
}

// GetSaleTypeExt retrieves the SaleTypeExt value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u SaleEntryExt) GetSaleTypeExt() (result SaleTypeExt, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.V))

	if armName == "SaleTypeExt" {
		result = *u.SaleTypeExt
		ok = true
	}

	return
}

// MustStatableSaleExt retrieves the StatableSaleExt value from the union,
// panicing if the value is not set.
func (u SaleEntryExt) MustStatableSaleExt() StatableSaleExt {
	val, ok := u.GetStatableSaleExt()

	if !ok {
		panic("arm StatableSaleExt is not set")
	}

	return val
}

// GetStatableSaleExt retrieves the StatableSaleExt value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u SaleEntryExt) GetStatableSaleExt() (result StatableSaleExt, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.V))

	if armName == "StatableSaleExt" {
		result = *u.StatableSaleExt
		ok = true
	}

	return
}

// SaleEntry is an XDR Struct defines as:
//
//   struct SaleEntry
//    {
//    	uint64 saleID;
//    	AccountID ownerID;
//        AssetCode baseAsset; // asset for which sale will be performed
//    	uint64 startTime; // start time of the sale
//    	uint64 endTime; // close time of the sale
//    	AssetCode defaultQuoteAsset; // asset for soft and hard cap
//    	uint64 softCap; // minimum amount of quote asset to be received at which sale will be considered a successful
//    	uint64 hardCap; // max amount of quote asset to be received
//    	uint64 currentCapInBase;
//    	uint64 maxAmountToBeSold;
//    	longstring details; // sale specific details
//    	SaleQuoteAsset quoteAssets<100>;
//
//    	BalanceID baseBalance;
//
//    	union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//    	case TYPED_SALE:
//    		SaleTypeExt saleTypeExt;
//    	case STATABLE_SALES:
//    		StatableSaleExt statableSaleExt;
//        }
//        ext;
//    };
//
type SaleEntry struct {
	SaleId            Uint64           `json:"saleID,omitempty"`
	OwnerId           AccountId        `json:"ownerID,omitempty"`
	BaseAsset         AssetCode        `json:"baseAsset,omitempty"`
	StartTime         Uint64           `json:"startTime,omitempty"`
	EndTime           Uint64           `json:"endTime,omitempty"`
	DefaultQuoteAsset AssetCode        `json:"defaultQuoteAsset,omitempty"`
	SoftCap           Uint64           `json:"softCap,omitempty"`
	HardCap           Uint64           `json:"hardCap,omitempty"`
	CurrentCapInBase  Uint64           `json:"currentCapInBase,omitempty"`
	MaxAmountToBeSold Uint64           `json:"maxAmountToBeSold,omitempty"`
	Details           Longstring       `json:"details,omitempty"`
	QuoteAssets       []SaleQuoteAsset `json:"quoteAssets,omitempty" xdrmaxsize:"100"`
	BaseBalance       BalanceId        `json:"baseBalance,omitempty"`
	Ext               SaleEntryExt     `json:"ext,omitempty"`
}

// StatisticsV2EntryExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type StatisticsV2EntryExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u StatisticsV2EntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of StatisticsV2EntryExt
func (u StatisticsV2EntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewStatisticsV2EntryExt creates a new  StatisticsV2EntryExt.
func NewStatisticsV2EntryExt(v LedgerVersion, value interface{}) (result StatisticsV2EntryExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// StatisticsV2Entry is an XDR Struct defines as:
//
//   struct StatisticsV2Entry
//    {
//        uint64      id;
//    	AccountID   accountID;
//    	StatsOpType statsOpType;
//        AssetCode   assetCode;
//        bool        isConvertNeeded;
//
//    	uint64 dailyOutcome;
//    	uint64 weeklyOutcome;
//    	uint64 monthlyOutcome;
//    	uint64 annualOutcome;
//
//    	int64 updatedAt;
//
//        // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type StatisticsV2Entry struct {
	Id              Uint64               `json:"id,omitempty"`
	AccountId       AccountId            `json:"accountID,omitempty"`
	StatsOpType     StatsOpType          `json:"statsOpType,omitempty"`
	AssetCode       AssetCode            `json:"assetCode,omitempty"`
	IsConvertNeeded bool                 `json:"isConvertNeeded,omitempty"`
	DailyOutcome    Uint64               `json:"dailyOutcome,omitempty"`
	WeeklyOutcome   Uint64               `json:"weeklyOutcome,omitempty"`
	MonthlyOutcome  Uint64               `json:"monthlyOutcome,omitempty"`
	AnnualOutcome   Uint64               `json:"annualOutcome,omitempty"`
	UpdatedAt       Int64                `json:"updatedAt,omitempty"`
	Ext             StatisticsV2EntryExt `json:"ext,omitempty"`
}

// StatisticsEntryExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type StatisticsEntryExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u StatisticsEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of StatisticsEntryExt
func (u StatisticsEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewStatisticsEntryExt creates a new  StatisticsEntryExt.
func NewStatisticsEntryExt(v LedgerVersion, value interface{}) (result StatisticsEntryExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// StatisticsEntry is an XDR Struct defines as:
//
//   struct StatisticsEntry
//    {
//    	AccountID accountID;
//
//    	uint64 dailyOutcome;
//    	uint64 weeklyOutcome;
//    	uint64 monthlyOutcome;
//    	uint64 annualOutcome;
//
//    	int64 updatedAt;
//
//        // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type StatisticsEntry struct {
	AccountId      AccountId          `json:"accountID,omitempty"`
	DailyOutcome   Uint64             `json:"dailyOutcome,omitempty"`
	WeeklyOutcome  Uint64             `json:"weeklyOutcome,omitempty"`
	MonthlyOutcome Uint64             `json:"monthlyOutcome,omitempty"`
	AnnualOutcome  Uint64             `json:"annualOutcome,omitempty"`
	UpdatedAt      Int64              `json:"updatedAt,omitempty"`
	Ext            StatisticsEntryExt `json:"ext,omitempty"`
}

// ThresholdIndexes is an XDR Enum defines as:
//
//   enum ThresholdIndexes
//    {
//        MASTER_WEIGHT = 0,
//        LOW = 1,
//        MED = 2,
//        HIGH = 3
//    };
//
type ThresholdIndexes int32

const (
	ThresholdIndexesMasterWeight ThresholdIndexes = 0
	ThresholdIndexesLow          ThresholdIndexes = 1
	ThresholdIndexesMed          ThresholdIndexes = 2
	ThresholdIndexesHigh         ThresholdIndexes = 3
)

var ThresholdIndexesAll = []ThresholdIndexes{
	ThresholdIndexesMasterWeight,
	ThresholdIndexesLow,
	ThresholdIndexesMed,
	ThresholdIndexesHigh,
}

var thresholdIndexesMap = map[int32]string{
	0: "ThresholdIndexesMasterWeight",
	1: "ThresholdIndexesLow",
	2: "ThresholdIndexesMed",
	3: "ThresholdIndexesHigh",
}

var thresholdIndexesShortMap = map[int32]string{
	0: "master_weight",
	1: "low",
	2: "med",
	3: "high",
}

var thresholdIndexesRevMap = map[string]int32{
	"ThresholdIndexesMasterWeight": 0,
	"ThresholdIndexesLow":          1,
	"ThresholdIndexesMed":          2,
	"ThresholdIndexesHigh":         3,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ThresholdIndexes
func (e ThresholdIndexes) ValidEnum(v int32) bool {
	_, ok := thresholdIndexesMap[v]
	return ok
}
func (e ThresholdIndexes) isFlag() bool {
	for i := len(ThresholdIndexesAll) - 1; i >= 0; i-- {
		expected := ThresholdIndexes(2) << uint64(len(ThresholdIndexesAll)-1) >> uint64(len(ThresholdIndexesAll)-i)
		if expected != ThresholdIndexesAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ThresholdIndexes) String() string {
	name, _ := thresholdIndexesMap[int32(e)]
	return name
}

func (e ThresholdIndexes) ShortString() string {
	name, _ := thresholdIndexesShortMap[int32(e)]
	return name
}

func (e ThresholdIndexes) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range ThresholdIndexesAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ThresholdIndexes) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ThresholdIndexes(t.Value)
	return nil
}

// LedgerEntryType is an XDR Enum defines as:
//
//   enum LedgerEntryType
//    {
//        ACCOUNT = 0,
//        FEE = 2,
//        BALANCE = 4,
//        PAYMENT_REQUEST = 5,
//        ASSET = 6,
//        REFERENCE_ENTRY = 7,
//        ACCOUNT_TYPE_LIMITS = 8,
//        STATISTICS = 9,
//        TRUST = 10,
//        ACCOUNT_LIMITS = 11,
//    	ASSET_PAIR = 12,
//    	OFFER_ENTRY = 13,
//    	REVIEWABLE_REQUEST = 15,
//    	EXTERNAL_SYSTEM_ACCOUNT_ID = 16,
//    	SALE = 17,
//    	ACCOUNT_KYC = 18,
//    	EXTERNAL_SYSTEM_ACCOUNT_ID_POOL_ENTRY = 19,
//        KEY_VALUE = 20,
//        SALE_ANTE = 21,
//        LIMITS_V2 = 22,
//        STATISTICS_V2 = 23,
//        PENDING_STATISTICS = 24,
//        CONTRACT = 25
//    };
//
type LedgerEntryType int32

const (
	LedgerEntryTypeAccount                          LedgerEntryType = 0
	LedgerEntryTypeFee                              LedgerEntryType = 2
	LedgerEntryTypeBalance                          LedgerEntryType = 4
	LedgerEntryTypePaymentRequest                   LedgerEntryType = 5
	LedgerEntryTypeAsset                            LedgerEntryType = 6
	LedgerEntryTypeReferenceEntry                   LedgerEntryType = 7
	LedgerEntryTypeAccountTypeLimits                LedgerEntryType = 8
	LedgerEntryTypeStatistics                       LedgerEntryType = 9
	LedgerEntryTypeTrust                            LedgerEntryType = 10
	LedgerEntryTypeAccountLimits                    LedgerEntryType = 11
	LedgerEntryTypeAssetPair                        LedgerEntryType = 12
	LedgerEntryTypeOfferEntry                       LedgerEntryType = 13
	LedgerEntryTypeReviewableRequest                LedgerEntryType = 15
	LedgerEntryTypeExternalSystemAccountId          LedgerEntryType = 16
	LedgerEntryTypeSale                             LedgerEntryType = 17
	LedgerEntryTypeAccountKyc                       LedgerEntryType = 18
	LedgerEntryTypeExternalSystemAccountIdPoolEntry LedgerEntryType = 19
	LedgerEntryTypeKeyValue                         LedgerEntryType = 20
	LedgerEntryTypeSaleAnte                         LedgerEntryType = 21
	LedgerEntryTypeLimitsV2                         LedgerEntryType = 22
	LedgerEntryTypeStatisticsV2                     LedgerEntryType = 23
	LedgerEntryTypePendingStatistics                LedgerEntryType = 24
	LedgerEntryTypeContract                         LedgerEntryType = 25
)

var LedgerEntryTypeAll = []LedgerEntryType{
	LedgerEntryTypeAccount,
	LedgerEntryTypeFee,
	LedgerEntryTypeBalance,
	LedgerEntryTypePaymentRequest,
	LedgerEntryTypeAsset,
	LedgerEntryTypeReferenceEntry,
	LedgerEntryTypeAccountTypeLimits,
	LedgerEntryTypeStatistics,
	LedgerEntryTypeTrust,
	LedgerEntryTypeAccountLimits,
	LedgerEntryTypeAssetPair,
	LedgerEntryTypeOfferEntry,
	LedgerEntryTypeReviewableRequest,
	LedgerEntryTypeExternalSystemAccountId,
	LedgerEntryTypeSale,
	LedgerEntryTypeAccountKyc,
	LedgerEntryTypeExternalSystemAccountIdPoolEntry,
	LedgerEntryTypeKeyValue,
	LedgerEntryTypeSaleAnte,
	LedgerEntryTypeLimitsV2,
	LedgerEntryTypeStatisticsV2,
	LedgerEntryTypePendingStatistics,
	LedgerEntryTypeContract,
}

var ledgerEntryTypeMap = map[int32]string{
	0:  "LedgerEntryTypeAccount",
	2:  "LedgerEntryTypeFee",
	4:  "LedgerEntryTypeBalance",
	5:  "LedgerEntryTypePaymentRequest",
	6:  "LedgerEntryTypeAsset",
	7:  "LedgerEntryTypeReferenceEntry",
	8:  "LedgerEntryTypeAccountTypeLimits",
	9:  "LedgerEntryTypeStatistics",
	10: "LedgerEntryTypeTrust",
	11: "LedgerEntryTypeAccountLimits",
	12: "LedgerEntryTypeAssetPair",
	13: "LedgerEntryTypeOfferEntry",
	15: "LedgerEntryTypeReviewableRequest",
	16: "LedgerEntryTypeExternalSystemAccountId",
	17: "LedgerEntryTypeSale",
	18: "LedgerEntryTypeAccountKyc",
	19: "LedgerEntryTypeExternalSystemAccountIdPoolEntry",
	20: "LedgerEntryTypeKeyValue",
	21: "LedgerEntryTypeSaleAnte",
	22: "LedgerEntryTypeLimitsV2",
	23: "LedgerEntryTypeStatisticsV2",
	24: "LedgerEntryTypePendingStatistics",
	25: "LedgerEntryTypeContract",
}

var ledgerEntryTypeShortMap = map[int32]string{
	0:  "account",
	2:  "fee",
	4:  "balance",
	5:  "payment_request",
	6:  "asset",
	7:  "reference_entry",
	8:  "account_type_limits",
	9:  "statistics",
	10: "trust",
	11: "account_limits",
	12: "asset_pair",
	13: "offer_entry",
	15: "reviewable_request",
	16: "external_system_account_id",
	17: "sale",
	18: "account_kyc",
	19: "external_system_account_id_pool_entry",
	20: "key_value",
	21: "sale_ante",
	22: "limits_v2",
	23: "statistics_v2",
	24: "pending_statistics",
	25: "contract",
}

var ledgerEntryTypeRevMap = map[string]int32{
	"LedgerEntryTypeAccount":                          0,
	"LedgerEntryTypeFee":                              2,
	"LedgerEntryTypeBalance":                          4,
	"LedgerEntryTypePaymentRequest":                   5,
	"LedgerEntryTypeAsset":                            6,
	"LedgerEntryTypeReferenceEntry":                   7,
	"LedgerEntryTypeAccountTypeLimits":                8,
	"LedgerEntryTypeStatistics":                       9,
	"LedgerEntryTypeTrust":                            10,
	"LedgerEntryTypeAccountLimits":                    11,
	"LedgerEntryTypeAssetPair":                        12,
	"LedgerEntryTypeOfferEntry":                       13,
	"LedgerEntryTypeReviewableRequest":                15,
	"LedgerEntryTypeExternalSystemAccountId":          16,
	"LedgerEntryTypeSale":                             17,
	"LedgerEntryTypeAccountKyc":                       18,
	"LedgerEntryTypeExternalSystemAccountIdPoolEntry": 19,
	"LedgerEntryTypeKeyValue":                         20,
	"LedgerEntryTypeSaleAnte":                         21,
	"LedgerEntryTypeLimitsV2":                         22,
	"LedgerEntryTypeStatisticsV2":                     23,
	"LedgerEntryTypePendingStatistics":                24,
	"LedgerEntryTypeContract":                         25,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for LedgerEntryType
func (e LedgerEntryType) ValidEnum(v int32) bool {
	_, ok := ledgerEntryTypeMap[v]
	return ok
}
func (e LedgerEntryType) isFlag() bool {
	for i := len(LedgerEntryTypeAll) - 1; i >= 0; i-- {
		expected := LedgerEntryType(2) << uint64(len(LedgerEntryTypeAll)-1) >> uint64(len(LedgerEntryTypeAll)-i)
		if expected != LedgerEntryTypeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e LedgerEntryType) String() string {
	name, _ := ledgerEntryTypeMap[int32(e)]
	return name
}

func (e LedgerEntryType) ShortString() string {
	name, _ := ledgerEntryTypeShortMap[int32(e)]
	return name
}

func (e LedgerEntryType) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range LedgerEntryTypeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *LedgerEntryType) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = LedgerEntryType(t.Value)
	return nil
}

// LedgerEntryData is an XDR NestedUnion defines as:
//
//   union switch (LedgerEntryType type)
//        {
//        case ACCOUNT:
//            AccountEntry account;
//        case FEE:
//            FeeEntry feeState;
//        case BALANCE:
//            BalanceEntry balance;
//        case ASSET:
//            AssetEntry asset;
//        case REFERENCE_ENTRY:
//            ReferenceEntry reference;
//        case ACCOUNT_TYPE_LIMITS:
//            AccountTypeLimitsEntry accountTypeLimits;
//        case STATISTICS:
//            StatisticsEntry stats;
//        case TRUST:
//            TrustEntry trust;
//        case ACCOUNT_LIMITS:
//            AccountLimitsEntry accountLimits;
//    	case ASSET_PAIR:
//    		AssetPairEntry assetPair;
//    	case OFFER_ENTRY:
//    		OfferEntry offer;
//    	case REVIEWABLE_REQUEST:
//    		ReviewableRequestEntry reviewableRequest;
//    	case EXTERNAL_SYSTEM_ACCOUNT_ID:
//    		ExternalSystemAccountID externalSystemAccountID;
//    	case SALE:
//    		SaleEntry sale;
//    	case KEY_VALUE:
//    	    KeyValueEntry keyValue;
//    	case ACCOUNT_KYC:
//            AccountKYCEntry accountKYC;
//        case EXTERNAL_SYSTEM_ACCOUNT_ID_POOL_ENTRY:
//            ExternalSystemAccountIDPoolEntry externalSystemAccountIDPoolEntry;
//        case SALE_ANTE:
//            SaleAnteEntry saleAnte;
//        case LIMITS_V2:
//            LimitsV2Entry limitsV2;
//        case STATISTICS_V2:
//            StatisticsV2Entry statisticsV2;
//        case PENDING_STATISTICS:
//            PendingStatisticsEntry pendingStatistics;
//        case CONTRACT:
//            ContractEntry contract;
//        }
//
type LedgerEntryData struct {
	Type                             LedgerEntryType                   `json:"type,omitempty"`
	Account                          *AccountEntry                     `json:"account,omitempty"`
	FeeState                         *FeeEntry                         `json:"feeState,omitempty"`
	Balance                          *BalanceEntry                     `json:"balance,omitempty"`
	Asset                            *AssetEntry                       `json:"asset,omitempty"`
	Reference                        *ReferenceEntry                   `json:"reference,omitempty"`
	AccountTypeLimits                *AccountTypeLimitsEntry           `json:"accountTypeLimits,omitempty"`
	Stats                            *StatisticsEntry                  `json:"stats,omitempty"`
	Trust                            *TrustEntry                       `json:"trust,omitempty"`
	AccountLimits                    *AccountLimitsEntry               `json:"accountLimits,omitempty"`
	AssetPair                        *AssetPairEntry                   `json:"assetPair,omitempty"`
	Offer                            *OfferEntry                       `json:"offer,omitempty"`
	ReviewableRequest                *ReviewableRequestEntry           `json:"reviewableRequest,omitempty"`
	ExternalSystemAccountId          *ExternalSystemAccountId          `json:"externalSystemAccountID,omitempty"`
	Sale                             *SaleEntry                        `json:"sale,omitempty"`
	KeyValue                         *KeyValueEntry                    `json:"keyValue,omitempty"`
	AccountKyc                       *AccountKycEntry                  `json:"accountKYC,omitempty"`
	ExternalSystemAccountIdPoolEntry *ExternalSystemAccountIdPoolEntry `json:"externalSystemAccountIDPoolEntry,omitempty"`
	SaleAnte                         *SaleAnteEntry                    `json:"saleAnte,omitempty"`
	LimitsV2                         *LimitsV2Entry                    `json:"limitsV2,omitempty"`
	StatisticsV2                     *StatisticsV2Entry                `json:"statisticsV2,omitempty"`
	PendingStatistics                *PendingStatisticsEntry           `json:"pendingStatistics,omitempty"`
	Contract                         *ContractEntry                    `json:"contract,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerEntryData) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerEntryData
func (u LedgerEntryData) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerEntryType(sw) {
	case LedgerEntryTypeAccount:
		return "Account", true
	case LedgerEntryTypeFee:
		return "FeeState", true
	case LedgerEntryTypeBalance:
		return "Balance", true
	case LedgerEntryTypeAsset:
		return "Asset", true
	case LedgerEntryTypeReferenceEntry:
		return "Reference", true
	case LedgerEntryTypeAccountTypeLimits:
		return "AccountTypeLimits", true
	case LedgerEntryTypeStatistics:
		return "Stats", true
	case LedgerEntryTypeTrust:
		return "Trust", true
	case LedgerEntryTypeAccountLimits:
		return "AccountLimits", true
	case LedgerEntryTypeAssetPair:
		return "AssetPair", true
	case LedgerEntryTypeOfferEntry:
		return "Offer", true
	case LedgerEntryTypeReviewableRequest:
		return "ReviewableRequest", true
	case LedgerEntryTypeExternalSystemAccountId:
		return "ExternalSystemAccountId", true
	case LedgerEntryTypeSale:
		return "Sale", true
	case LedgerEntryTypeKeyValue:
		return "KeyValue", true
	case LedgerEntryTypeAccountKyc:
		return "AccountKyc", true
	case LedgerEntryTypeExternalSystemAccountIdPoolEntry:
		return "ExternalSystemAccountIdPoolEntry", true
	case LedgerEntryTypeSaleAnte:
		return "SaleAnte", true
	case LedgerEntryTypeLimitsV2:
		return "LimitsV2", true
	case LedgerEntryTypeStatisticsV2:
		return "StatisticsV2", true
	case LedgerEntryTypePendingStatistics:
		return "PendingStatistics", true
	case LedgerEntryTypeContract:
		return "Contract", true
	}
	return "-", false
}

// NewLedgerEntryData creates a new  LedgerEntryData.
func NewLedgerEntryData(aType LedgerEntryType, value interface{}) (result LedgerEntryData, err error) {
	result.Type = aType
	switch LedgerEntryType(aType) {
	case LedgerEntryTypeAccount:
		tv, ok := value.(AccountEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be AccountEntry")
			return
		}
		result.Account = &tv
	case LedgerEntryTypeFee:
		tv, ok := value.(FeeEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be FeeEntry")
			return
		}
		result.FeeState = &tv
	case LedgerEntryTypeBalance:
		tv, ok := value.(BalanceEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be BalanceEntry")
			return
		}
		result.Balance = &tv
	case LedgerEntryTypeAsset:
		tv, ok := value.(AssetEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be AssetEntry")
			return
		}
		result.Asset = &tv
	case LedgerEntryTypeReferenceEntry:
		tv, ok := value.(ReferenceEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be ReferenceEntry")
			return
		}
		result.Reference = &tv
	case LedgerEntryTypeAccountTypeLimits:
		tv, ok := value.(AccountTypeLimitsEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be AccountTypeLimitsEntry")
			return
		}
		result.AccountTypeLimits = &tv
	case LedgerEntryTypeStatistics:
		tv, ok := value.(StatisticsEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be StatisticsEntry")
			return
		}
		result.Stats = &tv
	case LedgerEntryTypeTrust:
		tv, ok := value.(TrustEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be TrustEntry")
			return
		}
		result.Trust = &tv
	case LedgerEntryTypeAccountLimits:
		tv, ok := value.(AccountLimitsEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be AccountLimitsEntry")
			return
		}
		result.AccountLimits = &tv
	case LedgerEntryTypeAssetPair:
		tv, ok := value.(AssetPairEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be AssetPairEntry")
			return
		}
		result.AssetPair = &tv
	case LedgerEntryTypeOfferEntry:
		tv, ok := value.(OfferEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be OfferEntry")
			return
		}
		result.Offer = &tv
	case LedgerEntryTypeReviewableRequest:
		tv, ok := value.(ReviewableRequestEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be ReviewableRequestEntry")
			return
		}
		result.ReviewableRequest = &tv
	case LedgerEntryTypeExternalSystemAccountId:
		tv, ok := value.(ExternalSystemAccountId)
		if !ok {
			err = fmt.Errorf("invalid value, must be ExternalSystemAccountId")
			return
		}
		result.ExternalSystemAccountId = &tv
	case LedgerEntryTypeSale:
		tv, ok := value.(SaleEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be SaleEntry")
			return
		}
		result.Sale = &tv
	case LedgerEntryTypeKeyValue:
		tv, ok := value.(KeyValueEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be KeyValueEntry")
			return
		}
		result.KeyValue = &tv
	case LedgerEntryTypeAccountKyc:
		tv, ok := value.(AccountKycEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be AccountKycEntry")
			return
		}
		result.AccountKyc = &tv
	case LedgerEntryTypeExternalSystemAccountIdPoolEntry:
		tv, ok := value.(ExternalSystemAccountIdPoolEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be ExternalSystemAccountIdPoolEntry")
			return
		}
		result.ExternalSystemAccountIdPoolEntry = &tv
	case LedgerEntryTypeSaleAnte:
		tv, ok := value.(SaleAnteEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be SaleAnteEntry")
			return
		}
		result.SaleAnte = &tv
	case LedgerEntryTypeLimitsV2:
		tv, ok := value.(LimitsV2Entry)
		if !ok {
			err = fmt.Errorf("invalid value, must be LimitsV2Entry")
			return
		}
		result.LimitsV2 = &tv
	case LedgerEntryTypeStatisticsV2:
		tv, ok := value.(StatisticsV2Entry)
		if !ok {
			err = fmt.Errorf("invalid value, must be StatisticsV2Entry")
			return
		}
		result.StatisticsV2 = &tv
	case LedgerEntryTypePendingStatistics:
		tv, ok := value.(PendingStatisticsEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be PendingStatisticsEntry")
			return
		}
		result.PendingStatistics = &tv
	case LedgerEntryTypeContract:
		tv, ok := value.(ContractEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be ContractEntry")
			return
		}
		result.Contract = &tv
	}
	return
}

// MustAccount retrieves the Account value from the union,
// panicing if the value is not set.
func (u LedgerEntryData) MustAccount() AccountEntry {
	val, ok := u.GetAccount()

	if !ok {
		panic("arm Account is not set")
	}

	return val
}

// GetAccount retrieves the Account value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryData) GetAccount() (result AccountEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Account" {
		result = *u.Account
		ok = true
	}

	return
}

// MustFeeState retrieves the FeeState value from the union,
// panicing if the value is not set.
func (u LedgerEntryData) MustFeeState() FeeEntry {
	val, ok := u.GetFeeState()

	if !ok {
		panic("arm FeeState is not set")
	}

	return val
}

// GetFeeState retrieves the FeeState value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryData) GetFeeState() (result FeeEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "FeeState" {
		result = *u.FeeState
		ok = true
	}

	return
}

// MustBalance retrieves the Balance value from the union,
// panicing if the value is not set.
func (u LedgerEntryData) MustBalance() BalanceEntry {
	val, ok := u.GetBalance()

	if !ok {
		panic("arm Balance is not set")
	}

	return val
}

// GetBalance retrieves the Balance value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryData) GetBalance() (result BalanceEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Balance" {
		result = *u.Balance
		ok = true
	}

	return
}

// MustAsset retrieves the Asset value from the union,
// panicing if the value is not set.
func (u LedgerEntryData) MustAsset() AssetEntry {
	val, ok := u.GetAsset()

	if !ok {
		panic("arm Asset is not set")
	}

	return val
}

// GetAsset retrieves the Asset value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryData) GetAsset() (result AssetEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Asset" {
		result = *u.Asset
		ok = true
	}

	return
}

// MustReference retrieves the Reference value from the union,
// panicing if the value is not set.
func (u LedgerEntryData) MustReference() ReferenceEntry {
	val, ok := u.GetReference()

	if !ok {
		panic("arm Reference is not set")
	}

	return val
}

// GetReference retrieves the Reference value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryData) GetReference() (result ReferenceEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Reference" {
		result = *u.Reference
		ok = true
	}

	return
}

// MustAccountTypeLimits retrieves the AccountTypeLimits value from the union,
// panicing if the value is not set.
func (u LedgerEntryData) MustAccountTypeLimits() AccountTypeLimitsEntry {
	val, ok := u.GetAccountTypeLimits()

	if !ok {
		panic("arm AccountTypeLimits is not set")
	}

	return val
}

// GetAccountTypeLimits retrieves the AccountTypeLimits value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryData) GetAccountTypeLimits() (result AccountTypeLimitsEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "AccountTypeLimits" {
		result = *u.AccountTypeLimits
		ok = true
	}

	return
}

// MustStats retrieves the Stats value from the union,
// panicing if the value is not set.
func (u LedgerEntryData) MustStats() StatisticsEntry {
	val, ok := u.GetStats()

	if !ok {
		panic("arm Stats is not set")
	}

	return val
}

// GetStats retrieves the Stats value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryData) GetStats() (result StatisticsEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Stats" {
		result = *u.Stats
		ok = true
	}

	return
}

// MustTrust retrieves the Trust value from the union,
// panicing if the value is not set.
func (u LedgerEntryData) MustTrust() TrustEntry {
	val, ok := u.GetTrust()

	if !ok {
		panic("arm Trust is not set")
	}

	return val
}

// GetTrust retrieves the Trust value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryData) GetTrust() (result TrustEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Trust" {
		result = *u.Trust
		ok = true
	}

	return
}

// MustAccountLimits retrieves the AccountLimits value from the union,
// panicing if the value is not set.
func (u LedgerEntryData) MustAccountLimits() AccountLimitsEntry {
	val, ok := u.GetAccountLimits()

	if !ok {
		panic("arm AccountLimits is not set")
	}

	return val
}

// GetAccountLimits retrieves the AccountLimits value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryData) GetAccountLimits() (result AccountLimitsEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "AccountLimits" {
		result = *u.AccountLimits
		ok = true
	}

	return
}

// MustAssetPair retrieves the AssetPair value from the union,
// panicing if the value is not set.
func (u LedgerEntryData) MustAssetPair() AssetPairEntry {
	val, ok := u.GetAssetPair()

	if !ok {
		panic("arm AssetPair is not set")
	}

	return val
}

// GetAssetPair retrieves the AssetPair value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryData) GetAssetPair() (result AssetPairEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "AssetPair" {
		result = *u.AssetPair
		ok = true
	}

	return
}

// MustOffer retrieves the Offer value from the union,
// panicing if the value is not set.
func (u LedgerEntryData) MustOffer() OfferEntry {
	val, ok := u.GetOffer()

	if !ok {
		panic("arm Offer is not set")
	}

	return val
}

// GetOffer retrieves the Offer value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryData) GetOffer() (result OfferEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Offer" {
		result = *u.Offer
		ok = true
	}

	return
}

// MustReviewableRequest retrieves the ReviewableRequest value from the union,
// panicing if the value is not set.
func (u LedgerEntryData) MustReviewableRequest() ReviewableRequestEntry {
	val, ok := u.GetReviewableRequest()

	if !ok {
		panic("arm ReviewableRequest is not set")
	}

	return val
}

// GetReviewableRequest retrieves the ReviewableRequest value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryData) GetReviewableRequest() (result ReviewableRequestEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ReviewableRequest" {
		result = *u.ReviewableRequest
		ok = true
	}

	return
}

// MustExternalSystemAccountId retrieves the ExternalSystemAccountId value from the union,
// panicing if the value is not set.
func (u LedgerEntryData) MustExternalSystemAccountId() ExternalSystemAccountId {
	val, ok := u.GetExternalSystemAccountId()

	if !ok {
		panic("arm ExternalSystemAccountId is not set")
	}

	return val
}

// GetExternalSystemAccountId retrieves the ExternalSystemAccountId value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryData) GetExternalSystemAccountId() (result ExternalSystemAccountId, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ExternalSystemAccountId" {
		result = *u.ExternalSystemAccountId
		ok = true
	}

	return
}

// MustSale retrieves the Sale value from the union,
// panicing if the value is not set.
func (u LedgerEntryData) MustSale() SaleEntry {
	val, ok := u.GetSale()

	if !ok {
		panic("arm Sale is not set")
	}

	return val
}

// GetSale retrieves the Sale value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryData) GetSale() (result SaleEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Sale" {
		result = *u.Sale
		ok = true
	}

	return
}

// MustKeyValue retrieves the KeyValue value from the union,
// panicing if the value is not set.
func (u LedgerEntryData) MustKeyValue() KeyValueEntry {
	val, ok := u.GetKeyValue()

	if !ok {
		panic("arm KeyValue is not set")
	}

	return val
}

// GetKeyValue retrieves the KeyValue value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryData) GetKeyValue() (result KeyValueEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "KeyValue" {
		result = *u.KeyValue
		ok = true
	}

	return
}

// MustAccountKyc retrieves the AccountKyc value from the union,
// panicing if the value is not set.
func (u LedgerEntryData) MustAccountKyc() AccountKycEntry {
	val, ok := u.GetAccountKyc()

	if !ok {
		panic("arm AccountKyc is not set")
	}

	return val
}

// GetAccountKyc retrieves the AccountKyc value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryData) GetAccountKyc() (result AccountKycEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "AccountKyc" {
		result = *u.AccountKyc
		ok = true
	}

	return
}

// MustExternalSystemAccountIdPoolEntry retrieves the ExternalSystemAccountIdPoolEntry value from the union,
// panicing if the value is not set.
func (u LedgerEntryData) MustExternalSystemAccountIdPoolEntry() ExternalSystemAccountIdPoolEntry {
	val, ok := u.GetExternalSystemAccountIdPoolEntry()

	if !ok {
		panic("arm ExternalSystemAccountIdPoolEntry is not set")
	}

	return val
}

// GetExternalSystemAccountIdPoolEntry retrieves the ExternalSystemAccountIdPoolEntry value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryData) GetExternalSystemAccountIdPoolEntry() (result ExternalSystemAccountIdPoolEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ExternalSystemAccountIdPoolEntry" {
		result = *u.ExternalSystemAccountIdPoolEntry
		ok = true
	}

	return
}

// MustSaleAnte retrieves the SaleAnte value from the union,
// panicing if the value is not set.
func (u LedgerEntryData) MustSaleAnte() SaleAnteEntry {
	val, ok := u.GetSaleAnte()

	if !ok {
		panic("arm SaleAnte is not set")
	}

	return val
}

// GetSaleAnte retrieves the SaleAnte value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryData) GetSaleAnte() (result SaleAnteEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "SaleAnte" {
		result = *u.SaleAnte
		ok = true
	}

	return
}

// MustLimitsV2 retrieves the LimitsV2 value from the union,
// panicing if the value is not set.
func (u LedgerEntryData) MustLimitsV2() LimitsV2Entry {
	val, ok := u.GetLimitsV2()

	if !ok {
		panic("arm LimitsV2 is not set")
	}

	return val
}

// GetLimitsV2 retrieves the LimitsV2 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryData) GetLimitsV2() (result LimitsV2Entry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "LimitsV2" {
		result = *u.LimitsV2
		ok = true
	}

	return
}

// MustStatisticsV2 retrieves the StatisticsV2 value from the union,
// panicing if the value is not set.
func (u LedgerEntryData) MustStatisticsV2() StatisticsV2Entry {
	val, ok := u.GetStatisticsV2()

	if !ok {
		panic("arm StatisticsV2 is not set")
	}

	return val
}

// GetStatisticsV2 retrieves the StatisticsV2 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryData) GetStatisticsV2() (result StatisticsV2Entry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "StatisticsV2" {
		result = *u.StatisticsV2
		ok = true
	}

	return
}

// MustPendingStatistics retrieves the PendingStatistics value from the union,
// panicing if the value is not set.
func (u LedgerEntryData) MustPendingStatistics() PendingStatisticsEntry {
	val, ok := u.GetPendingStatistics()

	if !ok {
		panic("arm PendingStatistics is not set")
	}

	return val
}

// GetPendingStatistics retrieves the PendingStatistics value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryData) GetPendingStatistics() (result PendingStatisticsEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "PendingStatistics" {
		result = *u.PendingStatistics
		ok = true
	}

	return
}

// MustContract retrieves the Contract value from the union,
// panicing if the value is not set.
func (u LedgerEntryData) MustContract() ContractEntry {
	val, ok := u.GetContract()

	if !ok {
		panic("arm Contract is not set")
	}

	return val
}

// GetContract retrieves the Contract value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryData) GetContract() (result ContractEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Contract" {
		result = *u.Contract
		ok = true
	}

	return
}

// LedgerEntryExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type LedgerEntryExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerEntryExt
func (u LedgerEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewLedgerEntryExt creates a new  LedgerEntryExt.
func NewLedgerEntryExt(v LedgerVersion, value interface{}) (result LedgerEntryExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// LedgerEntry is an XDR Struct defines as:
//
//   struct LedgerEntry
//    {
//        uint32 lastModifiedLedgerSeq; // ledger the LedgerEntry was last changed
//
//        union switch (LedgerEntryType type)
//        {
//        case ACCOUNT:
//            AccountEntry account;
//        case FEE:
//            FeeEntry feeState;
//        case BALANCE:
//            BalanceEntry balance;
//        case ASSET:
//            AssetEntry asset;
//        case REFERENCE_ENTRY:
//            ReferenceEntry reference;
//        case ACCOUNT_TYPE_LIMITS:
//            AccountTypeLimitsEntry accountTypeLimits;
//        case STATISTICS:
//            StatisticsEntry stats;
//        case TRUST:
//            TrustEntry trust;
//        case ACCOUNT_LIMITS:
//            AccountLimitsEntry accountLimits;
//    	case ASSET_PAIR:
//    		AssetPairEntry assetPair;
//    	case OFFER_ENTRY:
//    		OfferEntry offer;
//    	case REVIEWABLE_REQUEST:
//    		ReviewableRequestEntry reviewableRequest;
//    	case EXTERNAL_SYSTEM_ACCOUNT_ID:
//    		ExternalSystemAccountID externalSystemAccountID;
//    	case SALE:
//    		SaleEntry sale;
//    	case KEY_VALUE:
//    	    KeyValueEntry keyValue;
//    	case ACCOUNT_KYC:
//            AccountKYCEntry accountKYC;
//        case EXTERNAL_SYSTEM_ACCOUNT_ID_POOL_ENTRY:
//            ExternalSystemAccountIDPoolEntry externalSystemAccountIDPoolEntry;
//        case SALE_ANTE:
//            SaleAnteEntry saleAnte;
//        case LIMITS_V2:
//            LimitsV2Entry limitsV2;
//        case STATISTICS_V2:
//            StatisticsV2Entry statisticsV2;
//        case PENDING_STATISTICS:
//            PendingStatisticsEntry pendingStatistics;
//        case CONTRACT:
//            ContractEntry contract;
//        }
//        data;
//
//        // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type LedgerEntry struct {
	LastModifiedLedgerSeq Uint32          `json:"lastModifiedLedgerSeq,omitempty"`
	Data                  LedgerEntryData `json:"data,omitempty"`
	Ext                   LedgerEntryExt  `json:"ext,omitempty"`
}

// EnvelopeType is an XDR Enum defines as:
//
//   enum EnvelopeType
//    {
//        SCP = 1,
//        TX = 2,
//        AUTH = 3
//    };
//
type EnvelopeType int32

const (
	EnvelopeTypeScp  EnvelopeType = 1
	EnvelopeTypeTx   EnvelopeType = 2
	EnvelopeTypeAuth EnvelopeType = 3
)

var EnvelopeTypeAll = []EnvelopeType{
	EnvelopeTypeScp,
	EnvelopeTypeTx,
	EnvelopeTypeAuth,
}

var envelopeTypeMap = map[int32]string{
	1: "EnvelopeTypeScp",
	2: "EnvelopeTypeTx",
	3: "EnvelopeTypeAuth",
}

var envelopeTypeShortMap = map[int32]string{
	1: "scp",
	2: "tx",
	3: "auth",
}

var envelopeTypeRevMap = map[string]int32{
	"EnvelopeTypeScp":  1,
	"EnvelopeTypeTx":   2,
	"EnvelopeTypeAuth": 3,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for EnvelopeType
func (e EnvelopeType) ValidEnum(v int32) bool {
	_, ok := envelopeTypeMap[v]
	return ok
}
func (e EnvelopeType) isFlag() bool {
	for i := len(EnvelopeTypeAll) - 1; i >= 0; i-- {
		expected := EnvelopeType(2) << uint64(len(EnvelopeTypeAll)-1) >> uint64(len(EnvelopeTypeAll)-i)
		if expected != EnvelopeTypeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e EnvelopeType) String() string {
	name, _ := envelopeTypeMap[int32(e)]
	return name
}

func (e EnvelopeType) ShortString() string {
	name, _ := envelopeTypeShortMap[int32(e)]
	return name
}

func (e EnvelopeType) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range EnvelopeTypeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *EnvelopeType) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = EnvelopeType(t.Value)
	return nil
}

// ExternalSystemIdGeneratorType is an XDR Enum defines as:
//
//   enum ExternalSystemIDGeneratorType {
//    	BITCOIN_BASIC = 1,
//    	ETHEREUM_BASIC = 2
//    };
//
type ExternalSystemIdGeneratorType int32

const (
	ExternalSystemIdGeneratorTypeBitcoinBasic  ExternalSystemIdGeneratorType = 1
	ExternalSystemIdGeneratorTypeEthereumBasic ExternalSystemIdGeneratorType = 2
)

var ExternalSystemIdGeneratorTypeAll = []ExternalSystemIdGeneratorType{
	ExternalSystemIdGeneratorTypeBitcoinBasic,
	ExternalSystemIdGeneratorTypeEthereumBasic,
}

var externalSystemIdGeneratorTypeMap = map[int32]string{
	1: "ExternalSystemIdGeneratorTypeBitcoinBasic",
	2: "ExternalSystemIdGeneratorTypeEthereumBasic",
}

var externalSystemIdGeneratorTypeShortMap = map[int32]string{
	1: "bitcoin_basic",
	2: "ethereum_basic",
}

var externalSystemIdGeneratorTypeRevMap = map[string]int32{
	"ExternalSystemIdGeneratorTypeBitcoinBasic":  1,
	"ExternalSystemIdGeneratorTypeEthereumBasic": 2,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ExternalSystemIdGeneratorType
func (e ExternalSystemIdGeneratorType) ValidEnum(v int32) bool {
	_, ok := externalSystemIdGeneratorTypeMap[v]
	return ok
}
func (e ExternalSystemIdGeneratorType) isFlag() bool {
	for i := len(ExternalSystemIdGeneratorTypeAll) - 1; i >= 0; i-- {
		expected := ExternalSystemIdGeneratorType(2) << uint64(len(ExternalSystemIdGeneratorTypeAll)-1) >> uint64(len(ExternalSystemIdGeneratorTypeAll)-i)
		if expected != ExternalSystemIdGeneratorTypeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ExternalSystemIdGeneratorType) String() string {
	name, _ := externalSystemIdGeneratorTypeMap[int32(e)]
	return name
}

func (e ExternalSystemIdGeneratorType) ShortString() string {
	name, _ := externalSystemIdGeneratorTypeShortMap[int32(e)]
	return name
}

func (e ExternalSystemIdGeneratorType) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range ExternalSystemIdGeneratorTypeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ExternalSystemIdGeneratorType) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ExternalSystemIdGeneratorType(t.Value)
	return nil
}

// UpgradeType is an XDR Typedef defines as:
//
//   typedef opaque UpgradeType<128>;
//
type UpgradeType []byte

// StellarValueExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type StellarValueExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u StellarValueExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of StellarValueExt
func (u StellarValueExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewStellarValueExt creates a new  StellarValueExt.
func NewStellarValueExt(v LedgerVersion, value interface{}) (result StellarValueExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// StellarValue is an XDR Struct defines as:
//
//   struct StellarValue
//    {
//        Hash txSetHash;   // transaction set to apply to previous ledger
//        uint64 closeTime; // network close time
//
//        // upgrades to apply to the previous ledger (usually empty)
//        // this is a vector of encoded 'LedgerUpgrade' so that nodes can drop
//        // unknown steps during consensus if needed.
//        // see notes below on 'LedgerUpgrade' for more detail
//        // max size is dictated by number of upgrade types (+ room for future)
//        UpgradeType upgrades<6>;
//
//        // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type StellarValue struct {
	TxSetHash Hash            `json:"txSetHash,omitempty"`
	CloseTime Uint64          `json:"closeTime,omitempty"`
	Upgrades  []UpgradeType   `json:"upgrades,omitempty" xdrmaxsize:"6"`
	Ext       StellarValueExt `json:"ext,omitempty"`
}

// IdGenerator is an XDR Struct defines as:
//
//   struct IdGenerator {
//    	LedgerEntryType entryType; // type of the entry, for which ids will be generated
//    	uint64 idPool; // last used entry specific ID, used for generating entry of specified type
//    };
//
type IdGenerator struct {
	EntryType LedgerEntryType `json:"entryType,omitempty"`
	IdPool    Uint64          `json:"idPool,omitempty"`
}

// LedgerHeaderExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type LedgerHeaderExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerHeaderExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerHeaderExt
func (u LedgerHeaderExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewLedgerHeaderExt creates a new  LedgerHeaderExt.
func NewLedgerHeaderExt(v LedgerVersion, value interface{}) (result LedgerHeaderExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// LedgerHeader is an XDR Struct defines as:
//
//   struct LedgerHeader
//    {
//        uint32 ledgerVersion;    // the protocol version of the ledger
//        Hash previousLedgerHash; // hash of the previous ledger header
//        StellarValue scpValue;   // what consensus agreed to
//        Hash txSetResultHash;    // the TransactionResultSet that led to this ledger
//        Hash bucketListHash;     // hash of the ledger state
//
//        uint32 ledgerSeq; // sequence number of this ledger
//
//        IdGenerator idGenerators<>; // generators of ids
//
//        uint32 baseFee;     // base fee per operation in stroops
//        uint32 baseReserve; // account base reserve in stroops
//
//        uint32 maxTxSetSize; // maximum size a transaction set can be
//
//        ExternalSystemIDGeneratorType externalSystemIDGenerators<>;
//        int64 txExpirationPeriod;
//
//        Hash skipList[4]; // hashes of ledgers in the past. allows you to jump back
//                          // in time without walking the chain back ledger by ledger
//                          // each slot contains the oldest ledger that is mod of
//                          // either 50  5000  50000 or 500000 depending on index
//                          // skipList[0] mod(50), skipList[1] mod(5000), etc
//
//        // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type LedgerHeader struct {
	LedgerVersion              Uint32                          `json:"ledgerVersion,omitempty"`
	PreviousLedgerHash         Hash                            `json:"previousLedgerHash,omitempty"`
	ScpValue                   StellarValue                    `json:"scpValue,omitempty"`
	TxSetResultHash            Hash                            `json:"txSetResultHash,omitempty"`
	BucketListHash             Hash                            `json:"bucketListHash,omitempty"`
	LedgerSeq                  Uint32                          `json:"ledgerSeq,omitempty"`
	IdGenerators               []IdGenerator                   `json:"idGenerators,omitempty"`
	BaseFee                    Uint32                          `json:"baseFee,omitempty"`
	BaseReserve                Uint32                          `json:"baseReserve,omitempty"`
	MaxTxSetSize               Uint32                          `json:"maxTxSetSize,omitempty"`
	ExternalSystemIdGenerators []ExternalSystemIdGeneratorType `json:"externalSystemIDGenerators,omitempty"`
	TxExpirationPeriod         Int64                           `json:"txExpirationPeriod,omitempty"`
	SkipList                   [4]Hash                         `json:"skipList,omitempty"`
	Ext                        LedgerHeaderExt                 `json:"ext,omitempty"`
}

// LedgerUpgradeType is an XDR Enum defines as:
//
//   enum LedgerUpgradeType
//    {
//        VERSION = 1,
//        MAX_TX_SET_SIZE = 2,
//        TX_EXPIRATION_PERIOD = 3,
//    	EXTERNAL_SYSTEM_ID_GENERATOR = 4
//    };
//
type LedgerUpgradeType int32

const (
	LedgerUpgradeTypeVersion                   LedgerUpgradeType = 1
	LedgerUpgradeTypeMaxTxSetSize              LedgerUpgradeType = 2
	LedgerUpgradeTypeTxExpirationPeriod        LedgerUpgradeType = 3
	LedgerUpgradeTypeExternalSystemIdGenerator LedgerUpgradeType = 4
)

var LedgerUpgradeTypeAll = []LedgerUpgradeType{
	LedgerUpgradeTypeVersion,
	LedgerUpgradeTypeMaxTxSetSize,
	LedgerUpgradeTypeTxExpirationPeriod,
	LedgerUpgradeTypeExternalSystemIdGenerator,
}

var ledgerUpgradeTypeMap = map[int32]string{
	1: "LedgerUpgradeTypeVersion",
	2: "LedgerUpgradeTypeMaxTxSetSize",
	3: "LedgerUpgradeTypeTxExpirationPeriod",
	4: "LedgerUpgradeTypeExternalSystemIdGenerator",
}

var ledgerUpgradeTypeShortMap = map[int32]string{
	1: "version",
	2: "max_tx_set_size",
	3: "tx_expiration_period",
	4: "external_system_id_generator",
}

var ledgerUpgradeTypeRevMap = map[string]int32{
	"LedgerUpgradeTypeVersion":                   1,
	"LedgerUpgradeTypeMaxTxSetSize":              2,
	"LedgerUpgradeTypeTxExpirationPeriod":        3,
	"LedgerUpgradeTypeExternalSystemIdGenerator": 4,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for LedgerUpgradeType
func (e LedgerUpgradeType) ValidEnum(v int32) bool {
	_, ok := ledgerUpgradeTypeMap[v]
	return ok
}
func (e LedgerUpgradeType) isFlag() bool {
	for i := len(LedgerUpgradeTypeAll) - 1; i >= 0; i-- {
		expected := LedgerUpgradeType(2) << uint64(len(LedgerUpgradeTypeAll)-1) >> uint64(len(LedgerUpgradeTypeAll)-i)
		if expected != LedgerUpgradeTypeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e LedgerUpgradeType) String() string {
	name, _ := ledgerUpgradeTypeMap[int32(e)]
	return name
}

func (e LedgerUpgradeType) ShortString() string {
	name, _ := ledgerUpgradeTypeShortMap[int32(e)]
	return name
}

func (e LedgerUpgradeType) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range LedgerUpgradeTypeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *LedgerUpgradeType) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = LedgerUpgradeType(t.Value)
	return nil
}

// LedgerUpgrade is an XDR Union defines as:
//
//   union LedgerUpgrade switch (LedgerUpgradeType type)
//    {
//    case VERSION:
//        uint32 newLedgerVersion; // update ledgerVersion
//    case MAX_TX_SET_SIZE:
//        uint32 newMaxTxSetSize; // update maxTxSetSize
//    case EXTERNAL_SYSTEM_ID_GENERATOR:
//        ExternalSystemIDGeneratorType newExternalSystemIDGenerators<>;
//    case TX_EXPIRATION_PERIOD:
//        int64 newTxExpirationPeriod;
//    };
//
type LedgerUpgrade struct {
	Type                          LedgerUpgradeType                `json:"type,omitempty"`
	NewLedgerVersion              *Uint32                          `json:"newLedgerVersion,omitempty"`
	NewMaxTxSetSize               *Uint32                          `json:"newMaxTxSetSize,omitempty"`
	NewExternalSystemIdGenerators *[]ExternalSystemIdGeneratorType `json:"newExternalSystemIDGenerators,omitempty"`
	NewTxExpirationPeriod         *Int64                           `json:"newTxExpirationPeriod,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerUpgrade) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerUpgrade
func (u LedgerUpgrade) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerUpgradeType(sw) {
	case LedgerUpgradeTypeVersion:
		return "NewLedgerVersion", true
	case LedgerUpgradeTypeMaxTxSetSize:
		return "NewMaxTxSetSize", true
	case LedgerUpgradeTypeExternalSystemIdGenerator:
		return "NewExternalSystemIdGenerators", true
	case LedgerUpgradeTypeTxExpirationPeriod:
		return "NewTxExpirationPeriod", true
	}
	return "-", false
}

// NewLedgerUpgrade creates a new  LedgerUpgrade.
func NewLedgerUpgrade(aType LedgerUpgradeType, value interface{}) (result LedgerUpgrade, err error) {
	result.Type = aType
	switch LedgerUpgradeType(aType) {
	case LedgerUpgradeTypeVersion:
		tv, ok := value.(Uint32)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint32")
			return
		}
		result.NewLedgerVersion = &tv
	case LedgerUpgradeTypeMaxTxSetSize:
		tv, ok := value.(Uint32)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint32")
			return
		}
		result.NewMaxTxSetSize = &tv
	case LedgerUpgradeTypeExternalSystemIdGenerator:
		tv, ok := value.([]ExternalSystemIdGeneratorType)
		if !ok {
			err = fmt.Errorf("invalid value, must be []ExternalSystemIdGeneratorType")
			return
		}
		result.NewExternalSystemIdGenerators = &tv
	case LedgerUpgradeTypeTxExpirationPeriod:
		tv, ok := value.(Int64)
		if !ok {
			err = fmt.Errorf("invalid value, must be Int64")
			return
		}
		result.NewTxExpirationPeriod = &tv
	}
	return
}

// MustNewLedgerVersion retrieves the NewLedgerVersion value from the union,
// panicing if the value is not set.
func (u LedgerUpgrade) MustNewLedgerVersion() Uint32 {
	val, ok := u.GetNewLedgerVersion()

	if !ok {
		panic("arm NewLedgerVersion is not set")
	}

	return val
}

// GetNewLedgerVersion retrieves the NewLedgerVersion value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerUpgrade) GetNewLedgerVersion() (result Uint32, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "NewLedgerVersion" {
		result = *u.NewLedgerVersion
		ok = true
	}

	return
}

// MustNewMaxTxSetSize retrieves the NewMaxTxSetSize value from the union,
// panicing if the value is not set.
func (u LedgerUpgrade) MustNewMaxTxSetSize() Uint32 {
	val, ok := u.GetNewMaxTxSetSize()

	if !ok {
		panic("arm NewMaxTxSetSize is not set")
	}

	return val
}

// GetNewMaxTxSetSize retrieves the NewMaxTxSetSize value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerUpgrade) GetNewMaxTxSetSize() (result Uint32, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "NewMaxTxSetSize" {
		result = *u.NewMaxTxSetSize
		ok = true
	}

	return
}

// MustNewExternalSystemIdGenerators retrieves the NewExternalSystemIdGenerators value from the union,
// panicing if the value is not set.
func (u LedgerUpgrade) MustNewExternalSystemIdGenerators() []ExternalSystemIdGeneratorType {
	val, ok := u.GetNewExternalSystemIdGenerators()

	if !ok {
		panic("arm NewExternalSystemIdGenerators is not set")
	}

	return val
}

// GetNewExternalSystemIdGenerators retrieves the NewExternalSystemIdGenerators value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerUpgrade) GetNewExternalSystemIdGenerators() (result []ExternalSystemIdGeneratorType, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "NewExternalSystemIdGenerators" {
		result = *u.NewExternalSystemIdGenerators
		ok = true
	}

	return
}

// MustNewTxExpirationPeriod retrieves the NewTxExpirationPeriod value from the union,
// panicing if the value is not set.
func (u LedgerUpgrade) MustNewTxExpirationPeriod() Int64 {
	val, ok := u.GetNewTxExpirationPeriod()

	if !ok {
		panic("arm NewTxExpirationPeriod is not set")
	}

	return val
}

// GetNewTxExpirationPeriod retrieves the NewTxExpirationPeriod value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerUpgrade) GetNewTxExpirationPeriod() (result Int64, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "NewTxExpirationPeriod" {
		result = *u.NewTxExpirationPeriod
		ok = true
	}

	return
}

// LedgerKeyAccountExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//
type LedgerKeyAccountExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerKeyAccountExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerKeyAccountExt
func (u LedgerKeyAccountExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewLedgerKeyAccountExt creates a new  LedgerKeyAccountExt.
func NewLedgerKeyAccountExt(v LedgerVersion, value interface{}) (result LedgerKeyAccountExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// LedgerKeyAccount is an XDR NestedStruct defines as:
//
//   struct
//        {
//            AccountID accountID;
//    		union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//    		ext;
//        }
//
type LedgerKeyAccount struct {
	AccountId AccountId           `json:"accountID,omitempty"`
	Ext       LedgerKeyAccountExt `json:"ext,omitempty"`
}

// LedgerKeyFeeStateExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//
type LedgerKeyFeeStateExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerKeyFeeStateExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerKeyFeeStateExt
func (u LedgerKeyFeeStateExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewLedgerKeyFeeStateExt creates a new  LedgerKeyFeeStateExt.
func NewLedgerKeyFeeStateExt(v LedgerVersion, value interface{}) (result LedgerKeyFeeStateExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// LedgerKeyFeeState is an XDR NestedStruct defines as:
//
//   struct {
//            Hash hash;
//    		int64 lowerBound;
//    		int64 upperBound;
//    		 union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//    		ext;
//        }
//
type LedgerKeyFeeState struct {
	Hash       Hash                 `json:"hash,omitempty"`
	LowerBound Int64                `json:"lowerBound,omitempty"`
	UpperBound Int64                `json:"upperBound,omitempty"`
	Ext        LedgerKeyFeeStateExt `json:"ext,omitempty"`
}

// LedgerKeyBalanceExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//
type LedgerKeyBalanceExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerKeyBalanceExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerKeyBalanceExt
func (u LedgerKeyBalanceExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewLedgerKeyBalanceExt creates a new  LedgerKeyBalanceExt.
func NewLedgerKeyBalanceExt(v LedgerVersion, value interface{}) (result LedgerKeyBalanceExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// LedgerKeyBalance is an XDR NestedStruct defines as:
//
//   struct
//        {
//    		BalanceID balanceID;
//    		union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//    		ext;
//        }
//
type LedgerKeyBalance struct {
	BalanceId BalanceId           `json:"balanceID,omitempty"`
	Ext       LedgerKeyBalanceExt `json:"ext,omitempty"`
}

// LedgerKeyAssetExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//
type LedgerKeyAssetExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerKeyAssetExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerKeyAssetExt
func (u LedgerKeyAssetExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewLedgerKeyAssetExt creates a new  LedgerKeyAssetExt.
func NewLedgerKeyAssetExt(v LedgerVersion, value interface{}) (result LedgerKeyAssetExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// LedgerKeyAsset is an XDR NestedStruct defines as:
//
//   struct
//        {
//    		AssetCode code;
//    		union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//    		ext;
//        }
//
type LedgerKeyAsset struct {
	Code AssetCode         `json:"code,omitempty"`
	Ext  LedgerKeyAssetExt `json:"ext,omitempty"`
}

// LedgerKeyReferenceExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//
type LedgerKeyReferenceExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerKeyReferenceExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerKeyReferenceExt
func (u LedgerKeyReferenceExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewLedgerKeyReferenceExt creates a new  LedgerKeyReferenceExt.
func NewLedgerKeyReferenceExt(v LedgerVersion, value interface{}) (result LedgerKeyReferenceExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// LedgerKeyReference is an XDR NestedStruct defines as:
//
//   struct
//        {
//    		AccountID sender;
//    		string64 reference;
//    		union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//    		ext;
//        }
//
type LedgerKeyReference struct {
	Sender    AccountId             `json:"sender,omitempty"`
	Reference String64              `json:"reference,omitempty"`
	Ext       LedgerKeyReferenceExt `json:"ext,omitempty"`
}

// LedgerKeyAccountTypeLimitsExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//
type LedgerKeyAccountTypeLimitsExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerKeyAccountTypeLimitsExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerKeyAccountTypeLimitsExt
func (u LedgerKeyAccountTypeLimitsExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewLedgerKeyAccountTypeLimitsExt creates a new  LedgerKeyAccountTypeLimitsExt.
func NewLedgerKeyAccountTypeLimitsExt(v LedgerVersion, value interface{}) (result LedgerKeyAccountTypeLimitsExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// LedgerKeyAccountTypeLimits is an XDR NestedStruct defines as:
//
//   struct {
//            AccountType accountType;
//    		union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//    		ext;
//        }
//
type LedgerKeyAccountTypeLimits struct {
	AccountType AccountType                   `json:"accountType,omitempty"`
	Ext         LedgerKeyAccountTypeLimitsExt `json:"ext,omitempty"`
}

// LedgerKeyStatsExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//
type LedgerKeyStatsExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerKeyStatsExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerKeyStatsExt
func (u LedgerKeyStatsExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewLedgerKeyStatsExt creates a new  LedgerKeyStatsExt.
func NewLedgerKeyStatsExt(v LedgerVersion, value interface{}) (result LedgerKeyStatsExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// LedgerKeyStats is an XDR NestedStruct defines as:
//
//   struct {
//            AccountID accountID;
//    		union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//    		ext;
//        }
//
type LedgerKeyStats struct {
	AccountId AccountId         `json:"accountID,omitempty"`
	Ext       LedgerKeyStatsExt `json:"ext,omitempty"`
}

// LedgerKeyTrustExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//
type LedgerKeyTrustExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerKeyTrustExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerKeyTrustExt
func (u LedgerKeyTrustExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewLedgerKeyTrustExt creates a new  LedgerKeyTrustExt.
func NewLedgerKeyTrustExt(v LedgerVersion, value interface{}) (result LedgerKeyTrustExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// LedgerKeyTrust is an XDR NestedStruct defines as:
//
//   struct {
//            AccountID allowedAccount;
//            BalanceID balanceToUse;
//    		union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//    		ext;
//        }
//
type LedgerKeyTrust struct {
	AllowedAccount AccountId         `json:"allowedAccount,omitempty"`
	BalanceToUse   BalanceId         `json:"balanceToUse,omitempty"`
	Ext            LedgerKeyTrustExt `json:"ext,omitempty"`
}

// LedgerKeyAccountLimitsExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//
type LedgerKeyAccountLimitsExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerKeyAccountLimitsExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerKeyAccountLimitsExt
func (u LedgerKeyAccountLimitsExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewLedgerKeyAccountLimitsExt creates a new  LedgerKeyAccountLimitsExt.
func NewLedgerKeyAccountLimitsExt(v LedgerVersion, value interface{}) (result LedgerKeyAccountLimitsExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// LedgerKeyAccountLimits is an XDR NestedStruct defines as:
//
//   struct {
//            AccountID accountID;
//    		union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//    		ext;
//        }
//
type LedgerKeyAccountLimits struct {
	AccountId AccountId                 `json:"accountID,omitempty"`
	Ext       LedgerKeyAccountLimitsExt `json:"ext,omitempty"`
}

// LedgerKeyAssetPairExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//
type LedgerKeyAssetPairExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerKeyAssetPairExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerKeyAssetPairExt
func (u LedgerKeyAssetPairExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewLedgerKeyAssetPairExt creates a new  LedgerKeyAssetPairExt.
func NewLedgerKeyAssetPairExt(v LedgerVersion, value interface{}) (result LedgerKeyAssetPairExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// LedgerKeyAssetPair is an XDR NestedStruct defines as:
//
//   struct {
//             AssetCode base;
//    		 AssetCode quote;
//    		 union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//    		ext;
//        }
//
type LedgerKeyAssetPair struct {
	Base  AssetCode             `json:"base,omitempty"`
	Quote AssetCode             `json:"quote,omitempty"`
	Ext   LedgerKeyAssetPairExt `json:"ext,omitempty"`
}

// LedgerKeyOffer is an XDR NestedStruct defines as:
//
//   struct {
//    		uint64 offerID;
//    		AccountID ownerID;
//    	}
//
type LedgerKeyOffer struct {
	OfferId Uint64    `json:"offerID,omitempty"`
	OwnerId AccountId `json:"ownerID,omitempty"`
}

// LedgerKeyReviewableRequestExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//
type LedgerKeyReviewableRequestExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerKeyReviewableRequestExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerKeyReviewableRequestExt
func (u LedgerKeyReviewableRequestExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewLedgerKeyReviewableRequestExt creates a new  LedgerKeyReviewableRequestExt.
func NewLedgerKeyReviewableRequestExt(v LedgerVersion, value interface{}) (result LedgerKeyReviewableRequestExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// LedgerKeyReviewableRequest is an XDR NestedStruct defines as:
//
//   struct {
//            uint64 requestID;
//    		union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//    		ext;
//        }
//
type LedgerKeyReviewableRequest struct {
	RequestId Uint64                        `json:"requestID,omitempty"`
	Ext       LedgerKeyReviewableRequestExt `json:"ext,omitempty"`
}

// LedgerKeyExternalSystemAccountIdExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//
type LedgerKeyExternalSystemAccountIdExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerKeyExternalSystemAccountIdExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerKeyExternalSystemAccountIdExt
func (u LedgerKeyExternalSystemAccountIdExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewLedgerKeyExternalSystemAccountIdExt creates a new  LedgerKeyExternalSystemAccountIdExt.
func NewLedgerKeyExternalSystemAccountIdExt(v LedgerVersion, value interface{}) (result LedgerKeyExternalSystemAccountIdExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// LedgerKeyExternalSystemAccountId is an XDR NestedStruct defines as:
//
//   struct {
//    		AccountID accountID;
//    		int32 externalSystemType;
//    		union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//    		ext;
//    	}
//
type LedgerKeyExternalSystemAccountId struct {
	AccountId          AccountId                           `json:"accountID,omitempty"`
	ExternalSystemType Int32                               `json:"externalSystemType,omitempty"`
	Ext                LedgerKeyExternalSystemAccountIdExt `json:"ext,omitempty"`
}

// LedgerKeySaleExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//
type LedgerKeySaleExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerKeySaleExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerKeySaleExt
func (u LedgerKeySaleExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewLedgerKeySaleExt creates a new  LedgerKeySaleExt.
func NewLedgerKeySaleExt(v LedgerVersion, value interface{}) (result LedgerKeySaleExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// LedgerKeySale is an XDR NestedStruct defines as:
//
//   struct {
//    		uint64 saleID;
//    		union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//    		ext;
//    	}
//
type LedgerKeySale struct {
	SaleId Uint64           `json:"saleID,omitempty"`
	Ext    LedgerKeySaleExt `json:"ext,omitempty"`
}

// LedgerKeyKeyValueExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            	case EMPTY_VERSION:
//            		void;
//            }
//
type LedgerKeyKeyValueExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerKeyKeyValueExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerKeyKeyValueExt
func (u LedgerKeyKeyValueExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewLedgerKeyKeyValueExt creates a new  LedgerKeyKeyValueExt.
func NewLedgerKeyKeyValueExt(v LedgerVersion, value interface{}) (result LedgerKeyKeyValueExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// LedgerKeyKeyValue is an XDR NestedStruct defines as:
//
//   struct {
//            string256 key;
//            union switch (LedgerVersion v)
//            {
//            	case EMPTY_VERSION:
//            		void;
//            }
//            ext;
//        }
//
type LedgerKeyKeyValue struct {
	Key String256            `json:"key,omitempty"`
	Ext LedgerKeyKeyValueExt `json:"ext,omitempty"`
}

// LedgerKeyAccountKycExt is an XDR NestedUnion defines as:
//
//   union switch(LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//
type LedgerKeyAccountKycExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerKeyAccountKycExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerKeyAccountKycExt
func (u LedgerKeyAccountKycExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewLedgerKeyAccountKycExt creates a new  LedgerKeyAccountKycExt.
func NewLedgerKeyAccountKycExt(v LedgerVersion, value interface{}) (result LedgerKeyAccountKycExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// LedgerKeyAccountKyc is an XDR NestedStruct defines as:
//
//   struct {
//            AccountID accountID;
//            union switch(LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//        }
//
type LedgerKeyAccountKyc struct {
	AccountId AccountId              `json:"accountID,omitempty"`
	Ext       LedgerKeyAccountKycExt `json:"ext,omitempty"`
}

// LedgerKeyExternalSystemAccountIdPoolEntryExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//
type LedgerKeyExternalSystemAccountIdPoolEntryExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerKeyExternalSystemAccountIdPoolEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerKeyExternalSystemAccountIdPoolEntryExt
func (u LedgerKeyExternalSystemAccountIdPoolEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewLedgerKeyExternalSystemAccountIdPoolEntryExt creates a new  LedgerKeyExternalSystemAccountIdPoolEntryExt.
func NewLedgerKeyExternalSystemAccountIdPoolEntryExt(v LedgerVersion, value interface{}) (result LedgerKeyExternalSystemAccountIdPoolEntryExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// LedgerKeyExternalSystemAccountIdPoolEntry is an XDR NestedStruct defines as:
//
//   struct {
//    		uint64 poolEntryID;
//    		union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//    		ext;
//    	}
//
type LedgerKeyExternalSystemAccountIdPoolEntry struct {
	PoolEntryId Uint64                                       `json:"poolEntryID,omitempty"`
	Ext         LedgerKeyExternalSystemAccountIdPoolEntryExt `json:"ext,omitempty"`
}

// LedgerKeySaleAnteExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//
type LedgerKeySaleAnteExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerKeySaleAnteExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerKeySaleAnteExt
func (u LedgerKeySaleAnteExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewLedgerKeySaleAnteExt creates a new  LedgerKeySaleAnteExt.
func NewLedgerKeySaleAnteExt(v LedgerVersion, value interface{}) (result LedgerKeySaleAnteExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// LedgerKeySaleAnte is an XDR NestedStruct defines as:
//
//   struct {
//            uint64 saleID;
//            BalanceID participantBalanceID;
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            } ext;
//        }
//
type LedgerKeySaleAnte struct {
	SaleId               Uint64               `json:"saleID,omitempty"`
	ParticipantBalanceId BalanceId            `json:"participantBalanceID,omitempty"`
	Ext                  LedgerKeySaleAnteExt `json:"ext,omitempty"`
}

// LedgerKeyLimitsV2Ext is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//
type LedgerKeyLimitsV2Ext struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerKeyLimitsV2Ext) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerKeyLimitsV2Ext
func (u LedgerKeyLimitsV2Ext) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewLedgerKeyLimitsV2Ext creates a new  LedgerKeyLimitsV2Ext.
func NewLedgerKeyLimitsV2Ext(v LedgerVersion, value interface{}) (result LedgerKeyLimitsV2Ext, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// LedgerKeyLimitsV2 is an XDR NestedStruct defines as:
//
//   struct {
//            uint64 id;
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            } ext;
//        }
//
type LedgerKeyLimitsV2 struct {
	Id  Uint64               `json:"id,omitempty"`
	Ext LedgerKeyLimitsV2Ext `json:"ext,omitempty"`
}

// LedgerKeyStatisticsV2Ext is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//
type LedgerKeyStatisticsV2Ext struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerKeyStatisticsV2Ext) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerKeyStatisticsV2Ext
func (u LedgerKeyStatisticsV2Ext) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewLedgerKeyStatisticsV2Ext creates a new  LedgerKeyStatisticsV2Ext.
func NewLedgerKeyStatisticsV2Ext(v LedgerVersion, value interface{}) (result LedgerKeyStatisticsV2Ext, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// LedgerKeyStatisticsV2 is an XDR NestedStruct defines as:
//
//   struct {
//            uint64 id;
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//        }
//
type LedgerKeyStatisticsV2 struct {
	Id  Uint64                   `json:"id,omitempty"`
	Ext LedgerKeyStatisticsV2Ext `json:"ext,omitempty"`
}

// LedgerKeyPendingStatisticsExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//
type LedgerKeyPendingStatisticsExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerKeyPendingStatisticsExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerKeyPendingStatisticsExt
func (u LedgerKeyPendingStatisticsExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewLedgerKeyPendingStatisticsExt creates a new  LedgerKeyPendingStatisticsExt.
func NewLedgerKeyPendingStatisticsExt(v LedgerVersion, value interface{}) (result LedgerKeyPendingStatisticsExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// LedgerKeyPendingStatistics is an XDR NestedStruct defines as:
//
//   struct {
//            uint64 statisticsID;
//            uint64 requestID;
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//        }
//
type LedgerKeyPendingStatistics struct {
	StatisticsId Uint64                        `json:"statisticsID,omitempty"`
	RequestId    Uint64                        `json:"requestID,omitempty"`
	Ext          LedgerKeyPendingStatisticsExt `json:"ext,omitempty"`
}

// LedgerKeyContractExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//
type LedgerKeyContractExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerKeyContractExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerKeyContractExt
func (u LedgerKeyContractExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewLedgerKeyContractExt creates a new  LedgerKeyContractExt.
func NewLedgerKeyContractExt(v LedgerVersion, value interface{}) (result LedgerKeyContractExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// LedgerKeyContract is an XDR NestedStruct defines as:
//
//   struct {
//            uint64 contractID;
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//        }
//
type LedgerKeyContract struct {
	ContractId Uint64               `json:"contractID,omitempty"`
	Ext        LedgerKeyContractExt `json:"ext,omitempty"`
}

// LedgerKey is an XDR Union defines as:
//
//   union LedgerKey switch (LedgerEntryType type)
//    {
//    case ACCOUNT:
//        struct
//        {
//            AccountID accountID;
//    		union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//    		ext;
//        } account;
//    case FEE:
//        struct {
//            Hash hash;
//    		int64 lowerBound;
//    		int64 upperBound;
//    		 union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//    		ext;
//        } feeState;
//    case BALANCE:
//        struct
//        {
//    		BalanceID balanceID;
//    		union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//    		ext;
//        } balance;
//    case ASSET:
//        struct
//        {
//    		AssetCode code;
//    		union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//    		ext;
//        } asset;
//    case REFERENCE_ENTRY:
//        struct
//        {
//    		AccountID sender;
//    		string64 reference;
//    		union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//    		ext;
//        } reference;
//    case ACCOUNT_TYPE_LIMITS:
//        struct {
//            AccountType accountType;
//    		union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//    		ext;
//        } accountTypeLimits;
//    case STATISTICS:
//        struct {
//            AccountID accountID;
//    		union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//    		ext;
//        } stats;
//    case TRUST:
//        struct {
//            AccountID allowedAccount;
//            BalanceID balanceToUse;
//    		union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//    		ext;
//        } trust;
//    case ACCOUNT_LIMITS:
//        struct {
//            AccountID accountID;
//    		union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//    		ext;
//        } accountLimits;
//    case ASSET_PAIR:
//    	struct {
//             AssetCode base;
//    		 AssetCode quote;
//    		 union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//    		ext;
//        } assetPair;
//    case OFFER_ENTRY:
//    	struct {
//    		uint64 offerID;
//    		AccountID ownerID;
//    	} offer;
//    case REVIEWABLE_REQUEST:
//        struct {
//            uint64 requestID;
//    		union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//    		ext;
//        } reviewableRequest;
//    case EXTERNAL_SYSTEM_ACCOUNT_ID:
//    	struct {
//    		AccountID accountID;
//    		int32 externalSystemType;
//    		union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//    		ext;
//    	} externalSystemAccountID;
//    case SALE:
//    	struct {
//    		uint64 saleID;
//    		union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//    		ext;
//    	} sale;
//    case KEY_VALUE:
//        struct {
//            string256 key;
//            union switch (LedgerVersion v)
//            {
//            	case EMPTY_VERSION:
//            		void;
//            }
//            ext;
//        } keyValue;
//    case ACCOUNT_KYC:
//        struct {
//            AccountID accountID;
//            union switch(LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//        } accountKYC;
//    case EXTERNAL_SYSTEM_ACCOUNT_ID_POOL_ENTRY:
//        struct {
//    		uint64 poolEntryID;
//    		union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//    		ext;
//    	} externalSystemAccountIDPoolEntry;
//    case SALE_ANTE:
//        struct {
//            uint64 saleID;
//            BalanceID participantBalanceID;
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            } ext;
//        } saleAnte;
//    case LIMITS_V2:
//        struct {
//            uint64 id;
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            } ext;
//        } limitsV2;
//    case STATISTICS_V2:
//        struct {
//            uint64 id;
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//        } statisticsV2;
//    case PENDING_STATISTICS:
//        struct {
//            uint64 statisticsID;
//            uint64 requestID;
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//        } pendingStatistics;
//    case CONTRACT:
//        struct {
//            uint64 contractID;
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//        } contract;
//    };
//
type LedgerKey struct {
	Type                             LedgerEntryType                            `json:"type,omitempty"`
	Account                          *LedgerKeyAccount                          `json:"account,omitempty"`
	FeeState                         *LedgerKeyFeeState                         `json:"feeState,omitempty"`
	Balance                          *LedgerKeyBalance                          `json:"balance,omitempty"`
	Asset                            *LedgerKeyAsset                            `json:"asset,omitempty"`
	Reference                        *LedgerKeyReference                        `json:"reference,omitempty"`
	AccountTypeLimits                *LedgerKeyAccountTypeLimits                `json:"accountTypeLimits,omitempty"`
	Stats                            *LedgerKeyStats                            `json:"stats,omitempty"`
	Trust                            *LedgerKeyTrust                            `json:"trust,omitempty"`
	AccountLimits                    *LedgerKeyAccountLimits                    `json:"accountLimits,omitempty"`
	AssetPair                        *LedgerKeyAssetPair                        `json:"assetPair,omitempty"`
	Offer                            *LedgerKeyOffer                            `json:"offer,omitempty"`
	ReviewableRequest                *LedgerKeyReviewableRequest                `json:"reviewableRequest,omitempty"`
	ExternalSystemAccountId          *LedgerKeyExternalSystemAccountId          `json:"externalSystemAccountID,omitempty"`
	Sale                             *LedgerKeySale                             `json:"sale,omitempty"`
	KeyValue                         *LedgerKeyKeyValue                         `json:"keyValue,omitempty"`
	AccountKyc                       *LedgerKeyAccountKyc                       `json:"accountKYC,omitempty"`
	ExternalSystemAccountIdPoolEntry *LedgerKeyExternalSystemAccountIdPoolEntry `json:"externalSystemAccountIDPoolEntry,omitempty"`
	SaleAnte                         *LedgerKeySaleAnte                         `json:"saleAnte,omitempty"`
	LimitsV2                         *LedgerKeyLimitsV2                         `json:"limitsV2,omitempty"`
	StatisticsV2                     *LedgerKeyStatisticsV2                     `json:"statisticsV2,omitempty"`
	PendingStatistics                *LedgerKeyPendingStatistics                `json:"pendingStatistics,omitempty"`
	Contract                         *LedgerKeyContract                         `json:"contract,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerKey) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerKey
func (u LedgerKey) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerEntryType(sw) {
	case LedgerEntryTypeAccount:
		return "Account", true
	case LedgerEntryTypeFee:
		return "FeeState", true
	case LedgerEntryTypeBalance:
		return "Balance", true
	case LedgerEntryTypeAsset:
		return "Asset", true
	case LedgerEntryTypeReferenceEntry:
		return "Reference", true
	case LedgerEntryTypeAccountTypeLimits:
		return "AccountTypeLimits", true
	case LedgerEntryTypeStatistics:
		return "Stats", true
	case LedgerEntryTypeTrust:
		return "Trust", true
	case LedgerEntryTypeAccountLimits:
		return "AccountLimits", true
	case LedgerEntryTypeAssetPair:
		return "AssetPair", true
	case LedgerEntryTypeOfferEntry:
		return "Offer", true
	case LedgerEntryTypeReviewableRequest:
		return "ReviewableRequest", true
	case LedgerEntryTypeExternalSystemAccountId:
		return "ExternalSystemAccountId", true
	case LedgerEntryTypeSale:
		return "Sale", true
	case LedgerEntryTypeKeyValue:
		return "KeyValue", true
	case LedgerEntryTypeAccountKyc:
		return "AccountKyc", true
	case LedgerEntryTypeExternalSystemAccountIdPoolEntry:
		return "ExternalSystemAccountIdPoolEntry", true
	case LedgerEntryTypeSaleAnte:
		return "SaleAnte", true
	case LedgerEntryTypeLimitsV2:
		return "LimitsV2", true
	case LedgerEntryTypeStatisticsV2:
		return "StatisticsV2", true
	case LedgerEntryTypePendingStatistics:
		return "PendingStatistics", true
	case LedgerEntryTypeContract:
		return "Contract", true
	}
	return "-", false
}

// NewLedgerKey creates a new  LedgerKey.
func NewLedgerKey(aType LedgerEntryType, value interface{}) (result LedgerKey, err error) {
	result.Type = aType
	switch LedgerEntryType(aType) {
	case LedgerEntryTypeAccount:
		tv, ok := value.(LedgerKeyAccount)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKeyAccount")
			return
		}
		result.Account = &tv
	case LedgerEntryTypeFee:
		tv, ok := value.(LedgerKeyFeeState)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKeyFeeState")
			return
		}
		result.FeeState = &tv
	case LedgerEntryTypeBalance:
		tv, ok := value.(LedgerKeyBalance)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKeyBalance")
			return
		}
		result.Balance = &tv
	case LedgerEntryTypeAsset:
		tv, ok := value.(LedgerKeyAsset)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKeyAsset")
			return
		}
		result.Asset = &tv
	case LedgerEntryTypeReferenceEntry:
		tv, ok := value.(LedgerKeyReference)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKeyReference")
			return
		}
		result.Reference = &tv
	case LedgerEntryTypeAccountTypeLimits:
		tv, ok := value.(LedgerKeyAccountTypeLimits)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKeyAccountTypeLimits")
			return
		}
		result.AccountTypeLimits = &tv
	case LedgerEntryTypeStatistics:
		tv, ok := value.(LedgerKeyStats)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKeyStats")
			return
		}
		result.Stats = &tv
	case LedgerEntryTypeTrust:
		tv, ok := value.(LedgerKeyTrust)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKeyTrust")
			return
		}
		result.Trust = &tv
	case LedgerEntryTypeAccountLimits:
		tv, ok := value.(LedgerKeyAccountLimits)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKeyAccountLimits")
			return
		}
		result.AccountLimits = &tv
	case LedgerEntryTypeAssetPair:
		tv, ok := value.(LedgerKeyAssetPair)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKeyAssetPair")
			return
		}
		result.AssetPair = &tv
	case LedgerEntryTypeOfferEntry:
		tv, ok := value.(LedgerKeyOffer)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKeyOffer")
			return
		}
		result.Offer = &tv
	case LedgerEntryTypeReviewableRequest:
		tv, ok := value.(LedgerKeyReviewableRequest)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKeyReviewableRequest")
			return
		}
		result.ReviewableRequest = &tv
	case LedgerEntryTypeExternalSystemAccountId:
		tv, ok := value.(LedgerKeyExternalSystemAccountId)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKeyExternalSystemAccountId")
			return
		}
		result.ExternalSystemAccountId = &tv
	case LedgerEntryTypeSale:
		tv, ok := value.(LedgerKeySale)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKeySale")
			return
		}
		result.Sale = &tv
	case LedgerEntryTypeKeyValue:
		tv, ok := value.(LedgerKeyKeyValue)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKeyKeyValue")
			return
		}
		result.KeyValue = &tv
	case LedgerEntryTypeAccountKyc:
		tv, ok := value.(LedgerKeyAccountKyc)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKeyAccountKyc")
			return
		}
		result.AccountKyc = &tv
	case LedgerEntryTypeExternalSystemAccountIdPoolEntry:
		tv, ok := value.(LedgerKeyExternalSystemAccountIdPoolEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKeyExternalSystemAccountIdPoolEntry")
			return
		}
		result.ExternalSystemAccountIdPoolEntry = &tv
	case LedgerEntryTypeSaleAnte:
		tv, ok := value.(LedgerKeySaleAnte)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKeySaleAnte")
			return
		}
		result.SaleAnte = &tv
	case LedgerEntryTypeLimitsV2:
		tv, ok := value.(LedgerKeyLimitsV2)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKeyLimitsV2")
			return
		}
		result.LimitsV2 = &tv
	case LedgerEntryTypeStatisticsV2:
		tv, ok := value.(LedgerKeyStatisticsV2)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKeyStatisticsV2")
			return
		}
		result.StatisticsV2 = &tv
	case LedgerEntryTypePendingStatistics:
		tv, ok := value.(LedgerKeyPendingStatistics)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKeyPendingStatistics")
			return
		}
		result.PendingStatistics = &tv
	case LedgerEntryTypeContract:
		tv, ok := value.(LedgerKeyContract)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKeyContract")
			return
		}
		result.Contract = &tv
	}
	return
}

// MustAccount retrieves the Account value from the union,
// panicing if the value is not set.
func (u LedgerKey) MustAccount() LedgerKeyAccount {
	val, ok := u.GetAccount()

	if !ok {
		panic("arm Account is not set")
	}

	return val
}

// GetAccount retrieves the Account value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerKey) GetAccount() (result LedgerKeyAccount, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Account" {
		result = *u.Account
		ok = true
	}

	return
}

// MustFeeState retrieves the FeeState value from the union,
// panicing if the value is not set.
func (u LedgerKey) MustFeeState() LedgerKeyFeeState {
	val, ok := u.GetFeeState()

	if !ok {
		panic("arm FeeState is not set")
	}

	return val
}

// GetFeeState retrieves the FeeState value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerKey) GetFeeState() (result LedgerKeyFeeState, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "FeeState" {
		result = *u.FeeState
		ok = true
	}

	return
}

// MustBalance retrieves the Balance value from the union,
// panicing if the value is not set.
func (u LedgerKey) MustBalance() LedgerKeyBalance {
	val, ok := u.GetBalance()

	if !ok {
		panic("arm Balance is not set")
	}

	return val
}

// GetBalance retrieves the Balance value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerKey) GetBalance() (result LedgerKeyBalance, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Balance" {
		result = *u.Balance
		ok = true
	}

	return
}

// MustAsset retrieves the Asset value from the union,
// panicing if the value is not set.
func (u LedgerKey) MustAsset() LedgerKeyAsset {
	val, ok := u.GetAsset()

	if !ok {
		panic("arm Asset is not set")
	}

	return val
}

// GetAsset retrieves the Asset value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerKey) GetAsset() (result LedgerKeyAsset, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Asset" {
		result = *u.Asset
		ok = true
	}

	return
}

// MustReference retrieves the Reference value from the union,
// panicing if the value is not set.
func (u LedgerKey) MustReference() LedgerKeyReference {
	val, ok := u.GetReference()

	if !ok {
		panic("arm Reference is not set")
	}

	return val
}

// GetReference retrieves the Reference value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerKey) GetReference() (result LedgerKeyReference, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Reference" {
		result = *u.Reference
		ok = true
	}

	return
}

// MustAccountTypeLimits retrieves the AccountTypeLimits value from the union,
// panicing if the value is not set.
func (u LedgerKey) MustAccountTypeLimits() LedgerKeyAccountTypeLimits {
	val, ok := u.GetAccountTypeLimits()

	if !ok {
		panic("arm AccountTypeLimits is not set")
	}

	return val
}

// GetAccountTypeLimits retrieves the AccountTypeLimits value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerKey) GetAccountTypeLimits() (result LedgerKeyAccountTypeLimits, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "AccountTypeLimits" {
		result = *u.AccountTypeLimits
		ok = true
	}

	return
}

// MustStats retrieves the Stats value from the union,
// panicing if the value is not set.
func (u LedgerKey) MustStats() LedgerKeyStats {
	val, ok := u.GetStats()

	if !ok {
		panic("arm Stats is not set")
	}

	return val
}

// GetStats retrieves the Stats value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerKey) GetStats() (result LedgerKeyStats, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Stats" {
		result = *u.Stats
		ok = true
	}

	return
}

// MustTrust retrieves the Trust value from the union,
// panicing if the value is not set.
func (u LedgerKey) MustTrust() LedgerKeyTrust {
	val, ok := u.GetTrust()

	if !ok {
		panic("arm Trust is not set")
	}

	return val
}

// GetTrust retrieves the Trust value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerKey) GetTrust() (result LedgerKeyTrust, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Trust" {
		result = *u.Trust
		ok = true
	}

	return
}

// MustAccountLimits retrieves the AccountLimits value from the union,
// panicing if the value is not set.
func (u LedgerKey) MustAccountLimits() LedgerKeyAccountLimits {
	val, ok := u.GetAccountLimits()

	if !ok {
		panic("arm AccountLimits is not set")
	}

	return val
}

// GetAccountLimits retrieves the AccountLimits value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerKey) GetAccountLimits() (result LedgerKeyAccountLimits, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "AccountLimits" {
		result = *u.AccountLimits
		ok = true
	}

	return
}

// MustAssetPair retrieves the AssetPair value from the union,
// panicing if the value is not set.
func (u LedgerKey) MustAssetPair() LedgerKeyAssetPair {
	val, ok := u.GetAssetPair()

	if !ok {
		panic("arm AssetPair is not set")
	}

	return val
}

// GetAssetPair retrieves the AssetPair value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerKey) GetAssetPair() (result LedgerKeyAssetPair, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "AssetPair" {
		result = *u.AssetPair
		ok = true
	}

	return
}

// MustOffer retrieves the Offer value from the union,
// panicing if the value is not set.
func (u LedgerKey) MustOffer() LedgerKeyOffer {
	val, ok := u.GetOffer()

	if !ok {
		panic("arm Offer is not set")
	}

	return val
}

// GetOffer retrieves the Offer value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerKey) GetOffer() (result LedgerKeyOffer, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Offer" {
		result = *u.Offer
		ok = true
	}

	return
}

// MustReviewableRequest retrieves the ReviewableRequest value from the union,
// panicing if the value is not set.
func (u LedgerKey) MustReviewableRequest() LedgerKeyReviewableRequest {
	val, ok := u.GetReviewableRequest()

	if !ok {
		panic("arm ReviewableRequest is not set")
	}

	return val
}

// GetReviewableRequest retrieves the ReviewableRequest value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerKey) GetReviewableRequest() (result LedgerKeyReviewableRequest, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ReviewableRequest" {
		result = *u.ReviewableRequest
		ok = true
	}

	return
}

// MustExternalSystemAccountId retrieves the ExternalSystemAccountId value from the union,
// panicing if the value is not set.
func (u LedgerKey) MustExternalSystemAccountId() LedgerKeyExternalSystemAccountId {
	val, ok := u.GetExternalSystemAccountId()

	if !ok {
		panic("arm ExternalSystemAccountId is not set")
	}

	return val
}

// GetExternalSystemAccountId retrieves the ExternalSystemAccountId value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerKey) GetExternalSystemAccountId() (result LedgerKeyExternalSystemAccountId, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ExternalSystemAccountId" {
		result = *u.ExternalSystemAccountId
		ok = true
	}

	return
}

// MustSale retrieves the Sale value from the union,
// panicing if the value is not set.
func (u LedgerKey) MustSale() LedgerKeySale {
	val, ok := u.GetSale()

	if !ok {
		panic("arm Sale is not set")
	}

	return val
}

// GetSale retrieves the Sale value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerKey) GetSale() (result LedgerKeySale, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Sale" {
		result = *u.Sale
		ok = true
	}

	return
}

// MustKeyValue retrieves the KeyValue value from the union,
// panicing if the value is not set.
func (u LedgerKey) MustKeyValue() LedgerKeyKeyValue {
	val, ok := u.GetKeyValue()

	if !ok {
		panic("arm KeyValue is not set")
	}

	return val
}

// GetKeyValue retrieves the KeyValue value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerKey) GetKeyValue() (result LedgerKeyKeyValue, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "KeyValue" {
		result = *u.KeyValue
		ok = true
	}

	return
}

// MustAccountKyc retrieves the AccountKyc value from the union,
// panicing if the value is not set.
func (u LedgerKey) MustAccountKyc() LedgerKeyAccountKyc {
	val, ok := u.GetAccountKyc()

	if !ok {
		panic("arm AccountKyc is not set")
	}

	return val
}

// GetAccountKyc retrieves the AccountKyc value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerKey) GetAccountKyc() (result LedgerKeyAccountKyc, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "AccountKyc" {
		result = *u.AccountKyc
		ok = true
	}

	return
}

// MustExternalSystemAccountIdPoolEntry retrieves the ExternalSystemAccountIdPoolEntry value from the union,
// panicing if the value is not set.
func (u LedgerKey) MustExternalSystemAccountIdPoolEntry() LedgerKeyExternalSystemAccountIdPoolEntry {
	val, ok := u.GetExternalSystemAccountIdPoolEntry()

	if !ok {
		panic("arm ExternalSystemAccountIdPoolEntry is not set")
	}

	return val
}

// GetExternalSystemAccountIdPoolEntry retrieves the ExternalSystemAccountIdPoolEntry value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerKey) GetExternalSystemAccountIdPoolEntry() (result LedgerKeyExternalSystemAccountIdPoolEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ExternalSystemAccountIdPoolEntry" {
		result = *u.ExternalSystemAccountIdPoolEntry
		ok = true
	}

	return
}

// MustSaleAnte retrieves the SaleAnte value from the union,
// panicing if the value is not set.
func (u LedgerKey) MustSaleAnte() LedgerKeySaleAnte {
	val, ok := u.GetSaleAnte()

	if !ok {
		panic("arm SaleAnte is not set")
	}

	return val
}

// GetSaleAnte retrieves the SaleAnte value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerKey) GetSaleAnte() (result LedgerKeySaleAnte, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "SaleAnte" {
		result = *u.SaleAnte
		ok = true
	}

	return
}

// MustLimitsV2 retrieves the LimitsV2 value from the union,
// panicing if the value is not set.
func (u LedgerKey) MustLimitsV2() LedgerKeyLimitsV2 {
	val, ok := u.GetLimitsV2()

	if !ok {
		panic("arm LimitsV2 is not set")
	}

	return val
}

// GetLimitsV2 retrieves the LimitsV2 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerKey) GetLimitsV2() (result LedgerKeyLimitsV2, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "LimitsV2" {
		result = *u.LimitsV2
		ok = true
	}

	return
}

// MustStatisticsV2 retrieves the StatisticsV2 value from the union,
// panicing if the value is not set.
func (u LedgerKey) MustStatisticsV2() LedgerKeyStatisticsV2 {
	val, ok := u.GetStatisticsV2()

	if !ok {
		panic("arm StatisticsV2 is not set")
	}

	return val
}

// GetStatisticsV2 retrieves the StatisticsV2 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerKey) GetStatisticsV2() (result LedgerKeyStatisticsV2, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "StatisticsV2" {
		result = *u.StatisticsV2
		ok = true
	}

	return
}

// MustPendingStatistics retrieves the PendingStatistics value from the union,
// panicing if the value is not set.
func (u LedgerKey) MustPendingStatistics() LedgerKeyPendingStatistics {
	val, ok := u.GetPendingStatistics()

	if !ok {
		panic("arm PendingStatistics is not set")
	}

	return val
}

// GetPendingStatistics retrieves the PendingStatistics value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerKey) GetPendingStatistics() (result LedgerKeyPendingStatistics, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "PendingStatistics" {
		result = *u.PendingStatistics
		ok = true
	}

	return
}

// MustContract retrieves the Contract value from the union,
// panicing if the value is not set.
func (u LedgerKey) MustContract() LedgerKeyContract {
	val, ok := u.GetContract()

	if !ok {
		panic("arm Contract is not set")
	}

	return val
}

// GetContract retrieves the Contract value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerKey) GetContract() (result LedgerKeyContract, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Contract" {
		result = *u.Contract
		ok = true
	}

	return
}

// BucketEntryType is an XDR Enum defines as:
//
//   enum BucketEntryType
//    {
//        LIVEENTRY = 0,
//        DEADENTRY = 1
//    };
//
type BucketEntryType int32

const (
	BucketEntryTypeLiveentry BucketEntryType = 0
	BucketEntryTypeDeadentry BucketEntryType = 1
)

var BucketEntryTypeAll = []BucketEntryType{
	BucketEntryTypeLiveentry,
	BucketEntryTypeDeadentry,
}

var bucketEntryTypeMap = map[int32]string{
	0: "BucketEntryTypeLiveentry",
	1: "BucketEntryTypeDeadentry",
}

var bucketEntryTypeShortMap = map[int32]string{
	0: "liveentry",
	1: "deadentry",
}

var bucketEntryTypeRevMap = map[string]int32{
	"BucketEntryTypeLiveentry": 0,
	"BucketEntryTypeDeadentry": 1,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for BucketEntryType
func (e BucketEntryType) ValidEnum(v int32) bool {
	_, ok := bucketEntryTypeMap[v]
	return ok
}
func (e BucketEntryType) isFlag() bool {
	for i := len(BucketEntryTypeAll) - 1; i >= 0; i-- {
		expected := BucketEntryType(2) << uint64(len(BucketEntryTypeAll)-1) >> uint64(len(BucketEntryTypeAll)-i)
		if expected != BucketEntryTypeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e BucketEntryType) String() string {
	name, _ := bucketEntryTypeMap[int32(e)]
	return name
}

func (e BucketEntryType) ShortString() string {
	name, _ := bucketEntryTypeShortMap[int32(e)]
	return name
}

func (e BucketEntryType) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range BucketEntryTypeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *BucketEntryType) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = BucketEntryType(t.Value)
	return nil
}

// BucketEntry is an XDR Union defines as:
//
//   union BucketEntry switch (BucketEntryType type)
//    {
//    case LIVEENTRY:
//        LedgerEntry liveEntry;
//
//    case DEADENTRY:
//        LedgerKey deadEntry;
//    };
//
type BucketEntry struct {
	Type      BucketEntryType `json:"type,omitempty"`
	LiveEntry *LedgerEntry    `json:"liveEntry,omitempty"`
	DeadEntry *LedgerKey      `json:"deadEntry,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u BucketEntry) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of BucketEntry
func (u BucketEntry) ArmForSwitch(sw int32) (string, bool) {
	switch BucketEntryType(sw) {
	case BucketEntryTypeLiveentry:
		return "LiveEntry", true
	case BucketEntryTypeDeadentry:
		return "DeadEntry", true
	}
	return "-", false
}

// NewBucketEntry creates a new  BucketEntry.
func NewBucketEntry(aType BucketEntryType, value interface{}) (result BucketEntry, err error) {
	result.Type = aType
	switch BucketEntryType(aType) {
	case BucketEntryTypeLiveentry:
		tv, ok := value.(LedgerEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerEntry")
			return
		}
		result.LiveEntry = &tv
	case BucketEntryTypeDeadentry:
		tv, ok := value.(LedgerKey)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKey")
			return
		}
		result.DeadEntry = &tv
	}
	return
}

// MustLiveEntry retrieves the LiveEntry value from the union,
// panicing if the value is not set.
func (u BucketEntry) MustLiveEntry() LedgerEntry {
	val, ok := u.GetLiveEntry()

	if !ok {
		panic("arm LiveEntry is not set")
	}

	return val
}

// GetLiveEntry retrieves the LiveEntry value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u BucketEntry) GetLiveEntry() (result LedgerEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "LiveEntry" {
		result = *u.LiveEntry
		ok = true
	}

	return
}

// MustDeadEntry retrieves the DeadEntry value from the union,
// panicing if the value is not set.
func (u BucketEntry) MustDeadEntry() LedgerKey {
	val, ok := u.GetDeadEntry()

	if !ok {
		panic("arm DeadEntry is not set")
	}

	return val
}

// GetDeadEntry retrieves the DeadEntry value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u BucketEntry) GetDeadEntry() (result LedgerKey, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "DeadEntry" {
		result = *u.DeadEntry
		ok = true
	}

	return
}

// TransactionSet is an XDR Struct defines as:
//
//   struct TransactionSet
//    {
//        Hash previousLedgerHash;
//        TransactionEnvelope txs<>;
//    };
//
type TransactionSet struct {
	PreviousLedgerHash Hash                  `json:"previousLedgerHash,omitempty"`
	Txs                []TransactionEnvelope `json:"txs,omitempty"`
}

// TransactionResultPair is an XDR Struct defines as:
//
//   struct TransactionResultPair
//    {
//        Hash transactionHash;
//        TransactionResult result; // result for the transaction
//    };
//
type TransactionResultPair struct {
	TransactionHash Hash              `json:"transactionHash,omitempty"`
	Result          TransactionResult `json:"result,omitempty"`
}

// TransactionResultSet is an XDR Struct defines as:
//
//   struct TransactionResultSet
//    {
//        TransactionResultPair results<>;
//    };
//
type TransactionResultSet struct {
	Results []TransactionResultPair `json:"results,omitempty"`
}

// TransactionHistoryEntryExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type TransactionHistoryEntryExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u TransactionHistoryEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of TransactionHistoryEntryExt
func (u TransactionHistoryEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewTransactionHistoryEntryExt creates a new  TransactionHistoryEntryExt.
func NewTransactionHistoryEntryExt(v LedgerVersion, value interface{}) (result TransactionHistoryEntryExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// TransactionHistoryEntry is an XDR Struct defines as:
//
//   struct TransactionHistoryEntry
//    {
//        uint32 ledgerSeq;
//        TransactionSet txSet;
//
//        // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type TransactionHistoryEntry struct {
	LedgerSeq Uint32                     `json:"ledgerSeq,omitempty"`
	TxSet     TransactionSet             `json:"txSet,omitempty"`
	Ext       TransactionHistoryEntryExt `json:"ext,omitempty"`
}

// TransactionHistoryResultEntryExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type TransactionHistoryResultEntryExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u TransactionHistoryResultEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of TransactionHistoryResultEntryExt
func (u TransactionHistoryResultEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewTransactionHistoryResultEntryExt creates a new  TransactionHistoryResultEntryExt.
func NewTransactionHistoryResultEntryExt(v LedgerVersion, value interface{}) (result TransactionHistoryResultEntryExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// TransactionHistoryResultEntry is an XDR Struct defines as:
//
//   struct TransactionHistoryResultEntry
//    {
//        uint32 ledgerSeq;
//        TransactionResultSet txResultSet;
//
//        // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type TransactionHistoryResultEntry struct {
	LedgerSeq   Uint32                           `json:"ledgerSeq,omitempty"`
	TxResultSet TransactionResultSet             `json:"txResultSet,omitempty"`
	Ext         TransactionHistoryResultEntryExt `json:"ext,omitempty"`
}

// LedgerHeaderHistoryEntryExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type LedgerHeaderHistoryEntryExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerHeaderHistoryEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerHeaderHistoryEntryExt
func (u LedgerHeaderHistoryEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewLedgerHeaderHistoryEntryExt creates a new  LedgerHeaderHistoryEntryExt.
func NewLedgerHeaderHistoryEntryExt(v LedgerVersion, value interface{}) (result LedgerHeaderHistoryEntryExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// LedgerHeaderHistoryEntry is an XDR Struct defines as:
//
//   struct LedgerHeaderHistoryEntry
//    {
//        Hash hash;
//        LedgerHeader header;
//
//        // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type LedgerHeaderHistoryEntry struct {
	Hash   Hash                        `json:"hash,omitempty"`
	Header LedgerHeader                `json:"header,omitempty"`
	Ext    LedgerHeaderHistoryEntryExt `json:"ext,omitempty"`
}

// LedgerScpMessages is an XDR Struct defines as:
//
//   struct LedgerSCPMessages
//    {
//        uint32 ledgerSeq;
//        SCPEnvelope messages<>;
//    };
//
type LedgerScpMessages struct {
	LedgerSeq Uint32        `json:"ledgerSeq,omitempty"`
	Messages  []ScpEnvelope `json:"messages,omitempty"`
}

// ScpHistoryEntryV0 is an XDR Struct defines as:
//
//   struct SCPHistoryEntryV0
//    {
//        SCPQuorumSet quorumSets<>; // additional quorum sets used by ledgerMessages
//        LedgerSCPMessages ledgerMessages;
//    };
//
type ScpHistoryEntryV0 struct {
	QuorumSets     []ScpQuorumSet    `json:"quorumSets,omitempty"`
	LedgerMessages LedgerScpMessages `json:"ledgerMessages,omitempty"`
}

// ScpHistoryEntry is an XDR Union defines as:
//
//   union SCPHistoryEntry switch (LedgerVersion v)
//    {
//    case EMPTY_VERSION:
//        SCPHistoryEntryV0 v0;
//    };
//
type ScpHistoryEntry struct {
	V  LedgerVersion      `json:"v,omitempty"`
	V0 *ScpHistoryEntryV0 `json:"v0,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ScpHistoryEntry) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ScpHistoryEntry
func (u ScpHistoryEntry) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "V0", true
	}
	return "-", false
}

// NewScpHistoryEntry creates a new  ScpHistoryEntry.
func NewScpHistoryEntry(v LedgerVersion, value interface{}) (result ScpHistoryEntry, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		tv, ok := value.(ScpHistoryEntryV0)
		if !ok {
			err = fmt.Errorf("invalid value, must be ScpHistoryEntryV0")
			return
		}
		result.V0 = &tv
	}
	return
}

// MustV0 retrieves the V0 value from the union,
// panicing if the value is not set.
func (u ScpHistoryEntry) MustV0() ScpHistoryEntryV0 {
	val, ok := u.GetV0()

	if !ok {
		panic("arm V0 is not set")
	}

	return val
}

// GetV0 retrieves the V0 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ScpHistoryEntry) GetV0() (result ScpHistoryEntryV0, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.V))

	if armName == "V0" {
		result = *u.V0
		ok = true
	}

	return
}

// LedgerEntryChangeType is an XDR Enum defines as:
//
//   enum LedgerEntryChangeType
//    {
//        CREATED = 0, // entry was added to the ledger
//        UPDATED = 1, // entry was modified in the ledger
//        REMOVED = 2, // entry was removed from the ledger
//        STATE = 3    // value of the entry
//    };
//
type LedgerEntryChangeType int32

const (
	LedgerEntryChangeTypeCreated LedgerEntryChangeType = 0
	LedgerEntryChangeTypeUpdated LedgerEntryChangeType = 1
	LedgerEntryChangeTypeRemoved LedgerEntryChangeType = 2
	LedgerEntryChangeTypeState   LedgerEntryChangeType = 3
)

var LedgerEntryChangeTypeAll = []LedgerEntryChangeType{
	LedgerEntryChangeTypeCreated,
	LedgerEntryChangeTypeUpdated,
	LedgerEntryChangeTypeRemoved,
	LedgerEntryChangeTypeState,
}

var ledgerEntryChangeTypeMap = map[int32]string{
	0: "LedgerEntryChangeTypeCreated",
	1: "LedgerEntryChangeTypeUpdated",
	2: "LedgerEntryChangeTypeRemoved",
	3: "LedgerEntryChangeTypeState",
}

var ledgerEntryChangeTypeShortMap = map[int32]string{
	0: "created",
	1: "updated",
	2: "removed",
	3: "state",
}

var ledgerEntryChangeTypeRevMap = map[string]int32{
	"LedgerEntryChangeTypeCreated": 0,
	"LedgerEntryChangeTypeUpdated": 1,
	"LedgerEntryChangeTypeRemoved": 2,
	"LedgerEntryChangeTypeState":   3,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for LedgerEntryChangeType
func (e LedgerEntryChangeType) ValidEnum(v int32) bool {
	_, ok := ledgerEntryChangeTypeMap[v]
	return ok
}
func (e LedgerEntryChangeType) isFlag() bool {
	for i := len(LedgerEntryChangeTypeAll) - 1; i >= 0; i-- {
		expected := LedgerEntryChangeType(2) << uint64(len(LedgerEntryChangeTypeAll)-1) >> uint64(len(LedgerEntryChangeTypeAll)-i)
		if expected != LedgerEntryChangeTypeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e LedgerEntryChangeType) String() string {
	name, _ := ledgerEntryChangeTypeMap[int32(e)]
	return name
}

func (e LedgerEntryChangeType) ShortString() string {
	name, _ := ledgerEntryChangeTypeShortMap[int32(e)]
	return name
}

func (e LedgerEntryChangeType) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range LedgerEntryChangeTypeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *LedgerEntryChangeType) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = LedgerEntryChangeType(t.Value)
	return nil
}

// LedgerEntryChange is an XDR Union defines as:
//
//   union LedgerEntryChange switch (LedgerEntryChangeType type)
//    {
//    case CREATED:
//        LedgerEntry created;
//    case UPDATED:
//        LedgerEntry updated;
//    case REMOVED:
//        LedgerKey removed;
//    case STATE:
//        LedgerEntry state;
//    };
//
type LedgerEntryChange struct {
	Type    LedgerEntryChangeType `json:"type,omitempty"`
	Created *LedgerEntry          `json:"created,omitempty"`
	Updated *LedgerEntry          `json:"updated,omitempty"`
	Removed *LedgerKey            `json:"removed,omitempty"`
	State   *LedgerEntry          `json:"state,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerEntryChange) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerEntryChange
func (u LedgerEntryChange) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerEntryChangeType(sw) {
	case LedgerEntryChangeTypeCreated:
		return "Created", true
	case LedgerEntryChangeTypeUpdated:
		return "Updated", true
	case LedgerEntryChangeTypeRemoved:
		return "Removed", true
	case LedgerEntryChangeTypeState:
		return "State", true
	}
	return "-", false
}

// NewLedgerEntryChange creates a new  LedgerEntryChange.
func NewLedgerEntryChange(aType LedgerEntryChangeType, value interface{}) (result LedgerEntryChange, err error) {
	result.Type = aType
	switch LedgerEntryChangeType(aType) {
	case LedgerEntryChangeTypeCreated:
		tv, ok := value.(LedgerEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerEntry")
			return
		}
		result.Created = &tv
	case LedgerEntryChangeTypeUpdated:
		tv, ok := value.(LedgerEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerEntry")
			return
		}
		result.Updated = &tv
	case LedgerEntryChangeTypeRemoved:
		tv, ok := value.(LedgerKey)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKey")
			return
		}
		result.Removed = &tv
	case LedgerEntryChangeTypeState:
		tv, ok := value.(LedgerEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerEntry")
			return
		}
		result.State = &tv
	}
	return
}

// MustCreated retrieves the Created value from the union,
// panicing if the value is not set.
func (u LedgerEntryChange) MustCreated() LedgerEntry {
	val, ok := u.GetCreated()

	if !ok {
		panic("arm Created is not set")
	}

	return val
}

// GetCreated retrieves the Created value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryChange) GetCreated() (result LedgerEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Created" {
		result = *u.Created
		ok = true
	}

	return
}

// MustUpdated retrieves the Updated value from the union,
// panicing if the value is not set.
func (u LedgerEntryChange) MustUpdated() LedgerEntry {
	val, ok := u.GetUpdated()

	if !ok {
		panic("arm Updated is not set")
	}

	return val
}

// GetUpdated retrieves the Updated value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryChange) GetUpdated() (result LedgerEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Updated" {
		result = *u.Updated
		ok = true
	}

	return
}

// MustRemoved retrieves the Removed value from the union,
// panicing if the value is not set.
func (u LedgerEntryChange) MustRemoved() LedgerKey {
	val, ok := u.GetRemoved()

	if !ok {
		panic("arm Removed is not set")
	}

	return val
}

// GetRemoved retrieves the Removed value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryChange) GetRemoved() (result LedgerKey, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Removed" {
		result = *u.Removed
		ok = true
	}

	return
}

// MustState retrieves the State value from the union,
// panicing if the value is not set.
func (u LedgerEntryChange) MustState() LedgerEntry {
	val, ok := u.GetState()

	if !ok {
		panic("arm State is not set")
	}

	return val
}

// GetState retrieves the State value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryChange) GetState() (result LedgerEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "State" {
		result = *u.State
		ok = true
	}

	return
}

// LedgerEntryChanges is an XDR Typedef defines as:
//
//   typedef LedgerEntryChange LedgerEntryChanges<>;
//
type LedgerEntryChanges []LedgerEntryChange

// OperationMeta is an XDR Struct defines as:
//
//   struct OperationMeta
//    {
//        LedgerEntryChanges changes;
//    };
//
type OperationMeta struct {
	Changes LedgerEntryChanges `json:"changes,omitempty"`
}

// TransactionMeta is an XDR Union defines as:
//
//   union TransactionMeta switch (LedgerVersion v)
//    {
//    case EMPTY_VERSION:
//        OperationMeta operations<>;
//    };
//
type TransactionMeta struct {
	V          LedgerVersion    `json:"v,omitempty"`
	Operations *[]OperationMeta `json:"operations,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u TransactionMeta) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of TransactionMeta
func (u TransactionMeta) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "Operations", true
	}
	return "-", false
}

// NewTransactionMeta creates a new  TransactionMeta.
func NewTransactionMeta(v LedgerVersion, value interface{}) (result TransactionMeta, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		tv, ok := value.([]OperationMeta)
		if !ok {
			err = fmt.Errorf("invalid value, must be []OperationMeta")
			return
		}
		result.Operations = &tv
	}
	return
}

// MustOperations retrieves the Operations value from the union,
// panicing if the value is not set.
func (u TransactionMeta) MustOperations() []OperationMeta {
	val, ok := u.GetOperations()

	if !ok {
		panic("arm Operations is not set")
	}

	return val
}

// GetOperations retrieves the Operations value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u TransactionMeta) GetOperations() (result []OperationMeta, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.V))

	if armName == "Operations" {
		result = *u.Operations
		ok = true
	}

	return
}

// BindExternalSystemAccountIdOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type BindExternalSystemAccountIdOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u BindExternalSystemAccountIdOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of BindExternalSystemAccountIdOpExt
func (u BindExternalSystemAccountIdOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewBindExternalSystemAccountIdOpExt creates a new  BindExternalSystemAccountIdOpExt.
func NewBindExternalSystemAccountIdOpExt(v LedgerVersion, value interface{}) (result BindExternalSystemAccountIdOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// BindExternalSystemAccountIdOp is an XDR Struct defines as:
//
//   struct BindExternalSystemAccountIdOp
//    {
//        int32 externalSystemType;
//
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type BindExternalSystemAccountIdOp struct {
	ExternalSystemType Int32                            `json:"externalSystemType,omitempty"`
	Ext                BindExternalSystemAccountIdOpExt `json:"ext,omitempty"`
}

// BindExternalSystemAccountIdResultCode is an XDR Enum defines as:
//
//   enum BindExternalSystemAccountIdResultCode
//    {
//        // codes considered as "success" for the operation
//        SUCCESS = 0,
//
//        // codes considered as "failure" for the operation
//        MALFORMED = -1,
//        NO_AVAILABLE_ID = -2,
//        AUTO_GENERATED_TYPE_NOT_ALLOWED = -3
//    };
//
type BindExternalSystemAccountIdResultCode int32

const (
	BindExternalSystemAccountIdResultCodeSuccess                     BindExternalSystemAccountIdResultCode = 0
	BindExternalSystemAccountIdResultCodeMalformed                   BindExternalSystemAccountIdResultCode = -1
	BindExternalSystemAccountIdResultCodeNoAvailableId               BindExternalSystemAccountIdResultCode = -2
	BindExternalSystemAccountIdResultCodeAutoGeneratedTypeNotAllowed BindExternalSystemAccountIdResultCode = -3
)

var BindExternalSystemAccountIdResultCodeAll = []BindExternalSystemAccountIdResultCode{
	BindExternalSystemAccountIdResultCodeSuccess,
	BindExternalSystemAccountIdResultCodeMalformed,
	BindExternalSystemAccountIdResultCodeNoAvailableId,
	BindExternalSystemAccountIdResultCodeAutoGeneratedTypeNotAllowed,
}

var bindExternalSystemAccountIdResultCodeMap = map[int32]string{
	0:  "BindExternalSystemAccountIdResultCodeSuccess",
	-1: "BindExternalSystemAccountIdResultCodeMalformed",
	-2: "BindExternalSystemAccountIdResultCodeNoAvailableId",
	-3: "BindExternalSystemAccountIdResultCodeAutoGeneratedTypeNotAllowed",
}

var bindExternalSystemAccountIdResultCodeShortMap = map[int32]string{
	0:  "success",
	-1: "malformed",
	-2: "no_available_id",
	-3: "auto_generated_type_not_allowed",
}

var bindExternalSystemAccountIdResultCodeRevMap = map[string]int32{
	"BindExternalSystemAccountIdResultCodeSuccess":                     0,
	"BindExternalSystemAccountIdResultCodeMalformed":                   -1,
	"BindExternalSystemAccountIdResultCodeNoAvailableId":               -2,
	"BindExternalSystemAccountIdResultCodeAutoGeneratedTypeNotAllowed": -3,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for BindExternalSystemAccountIdResultCode
func (e BindExternalSystemAccountIdResultCode) ValidEnum(v int32) bool {
	_, ok := bindExternalSystemAccountIdResultCodeMap[v]
	return ok
}
func (e BindExternalSystemAccountIdResultCode) isFlag() bool {
	for i := len(BindExternalSystemAccountIdResultCodeAll) - 1; i >= 0; i-- {
		expected := BindExternalSystemAccountIdResultCode(2) << uint64(len(BindExternalSystemAccountIdResultCodeAll)-1) >> uint64(len(BindExternalSystemAccountIdResultCodeAll)-i)
		if expected != BindExternalSystemAccountIdResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e BindExternalSystemAccountIdResultCode) String() string {
	name, _ := bindExternalSystemAccountIdResultCodeMap[int32(e)]
	return name
}

func (e BindExternalSystemAccountIdResultCode) ShortString() string {
	name, _ := bindExternalSystemAccountIdResultCodeShortMap[int32(e)]
	return name
}

func (e BindExternalSystemAccountIdResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range BindExternalSystemAccountIdResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *BindExternalSystemAccountIdResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = BindExternalSystemAccountIdResultCode(t.Value)
	return nil
}

// BindExternalSystemAccountIdSuccessExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type BindExternalSystemAccountIdSuccessExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u BindExternalSystemAccountIdSuccessExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of BindExternalSystemAccountIdSuccessExt
func (u BindExternalSystemAccountIdSuccessExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewBindExternalSystemAccountIdSuccessExt creates a new  BindExternalSystemAccountIdSuccessExt.
func NewBindExternalSystemAccountIdSuccessExt(v LedgerVersion, value interface{}) (result BindExternalSystemAccountIdSuccessExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// BindExternalSystemAccountIdSuccess is an XDR Struct defines as:
//
//   struct BindExternalSystemAccountIdSuccess {
//        longstring data;
//        // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type BindExternalSystemAccountIdSuccess struct {
	Data Longstring                            `json:"data,omitempty"`
	Ext  BindExternalSystemAccountIdSuccessExt `json:"ext,omitempty"`
}

// BindExternalSystemAccountIdResult is an XDR Union defines as:
//
//   union BindExternalSystemAccountIdResult switch (BindExternalSystemAccountIdResultCode code)
//    {
//    case SUCCESS:
//        BindExternalSystemAccountIdSuccess success;
//    default:
//        void;
//    };
//
type BindExternalSystemAccountIdResult struct {
	Code    BindExternalSystemAccountIdResultCode `json:"code,omitempty"`
	Success *BindExternalSystemAccountIdSuccess   `json:"success,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u BindExternalSystemAccountIdResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of BindExternalSystemAccountIdResult
func (u BindExternalSystemAccountIdResult) ArmForSwitch(sw int32) (string, bool) {
	switch BindExternalSystemAccountIdResultCode(sw) {
	case BindExternalSystemAccountIdResultCodeSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewBindExternalSystemAccountIdResult creates a new  BindExternalSystemAccountIdResult.
func NewBindExternalSystemAccountIdResult(code BindExternalSystemAccountIdResultCode, value interface{}) (result BindExternalSystemAccountIdResult, err error) {
	result.Code = code
	switch BindExternalSystemAccountIdResultCode(code) {
	case BindExternalSystemAccountIdResultCodeSuccess:
		tv, ok := value.(BindExternalSystemAccountIdSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be BindExternalSystemAccountIdSuccess")
			return
		}
		result.Success = &tv
	default:
		// void
	}
	return
}

// MustSuccess retrieves the Success value from the union,
// panicing if the value is not set.
func (u BindExternalSystemAccountIdResult) MustSuccess() BindExternalSystemAccountIdSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u BindExternalSystemAccountIdResult) GetSuccess() (result BindExternalSystemAccountIdSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// CheckSaleStateOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type CheckSaleStateOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CheckSaleStateOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CheckSaleStateOpExt
func (u CheckSaleStateOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewCheckSaleStateOpExt creates a new  CheckSaleStateOpExt.
func NewCheckSaleStateOpExt(v LedgerVersion, value interface{}) (result CheckSaleStateOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// CheckSaleStateOp is an XDR Struct defines as:
//
//   struct CheckSaleStateOp
//    {
//    	uint64 saleID;
//    	 // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type CheckSaleStateOp struct {
	SaleId Uint64              `json:"saleID,omitempty"`
	Ext    CheckSaleStateOpExt `json:"ext,omitempty"`
}

// CheckSaleStateResultCode is an XDR Enum defines as:
//
//   enum CheckSaleStateResultCode
//    {
//        // codes considered as "success" for the operation
//        SUCCESS = 0, // sale was processed
//
//        // codes considered as "failure" for the operation
//        NOT_FOUND = -1, // sale was not found
//    	NOT_READY = -2 // sale is not ready to be closed or canceled
//    };
//
type CheckSaleStateResultCode int32

const (
	CheckSaleStateResultCodeSuccess  CheckSaleStateResultCode = 0
	CheckSaleStateResultCodeNotFound CheckSaleStateResultCode = -1
	CheckSaleStateResultCodeNotReady CheckSaleStateResultCode = -2
)

var CheckSaleStateResultCodeAll = []CheckSaleStateResultCode{
	CheckSaleStateResultCodeSuccess,
	CheckSaleStateResultCodeNotFound,
	CheckSaleStateResultCodeNotReady,
}

var checkSaleStateResultCodeMap = map[int32]string{
	0:  "CheckSaleStateResultCodeSuccess",
	-1: "CheckSaleStateResultCodeNotFound",
	-2: "CheckSaleStateResultCodeNotReady",
}

var checkSaleStateResultCodeShortMap = map[int32]string{
	0:  "success",
	-1: "not_found",
	-2: "not_ready",
}

var checkSaleStateResultCodeRevMap = map[string]int32{
	"CheckSaleStateResultCodeSuccess":  0,
	"CheckSaleStateResultCodeNotFound": -1,
	"CheckSaleStateResultCodeNotReady": -2,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for CheckSaleStateResultCode
func (e CheckSaleStateResultCode) ValidEnum(v int32) bool {
	_, ok := checkSaleStateResultCodeMap[v]
	return ok
}
func (e CheckSaleStateResultCode) isFlag() bool {
	for i := len(CheckSaleStateResultCodeAll) - 1; i >= 0; i-- {
		expected := CheckSaleStateResultCode(2) << uint64(len(CheckSaleStateResultCodeAll)-1) >> uint64(len(CheckSaleStateResultCodeAll)-i)
		if expected != CheckSaleStateResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e CheckSaleStateResultCode) String() string {
	name, _ := checkSaleStateResultCodeMap[int32(e)]
	return name
}

func (e CheckSaleStateResultCode) ShortString() string {
	name, _ := checkSaleStateResultCodeShortMap[int32(e)]
	return name
}

func (e CheckSaleStateResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range CheckSaleStateResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *CheckSaleStateResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = CheckSaleStateResultCode(t.Value)
	return nil
}

// CheckSaleStateEffect is an XDR Enum defines as:
//
//   enum CheckSaleStateEffect {
//    	CANCELED = 1, // sale did not managed to go over soft cap in time
//    	CLOSED = 2, // sale met hard cap or (end time and soft cap)
//    	UPDATED = 3 // on check sale was modified and modifications must be saved
//    };
//
type CheckSaleStateEffect int32

const (
	CheckSaleStateEffectCanceled CheckSaleStateEffect = 1
	CheckSaleStateEffectClosed   CheckSaleStateEffect = 2
	CheckSaleStateEffectUpdated  CheckSaleStateEffect = 3
)

var CheckSaleStateEffectAll = []CheckSaleStateEffect{
	CheckSaleStateEffectCanceled,
	CheckSaleStateEffectClosed,
	CheckSaleStateEffectUpdated,
}

var checkSaleStateEffectMap = map[int32]string{
	1: "CheckSaleStateEffectCanceled",
	2: "CheckSaleStateEffectClosed",
	3: "CheckSaleStateEffectUpdated",
}

var checkSaleStateEffectShortMap = map[int32]string{
	1: "canceled",
	2: "closed",
	3: "updated",
}

var checkSaleStateEffectRevMap = map[string]int32{
	"CheckSaleStateEffectCanceled": 1,
	"CheckSaleStateEffectClosed":   2,
	"CheckSaleStateEffectUpdated":  3,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for CheckSaleStateEffect
func (e CheckSaleStateEffect) ValidEnum(v int32) bool {
	_, ok := checkSaleStateEffectMap[v]
	return ok
}
func (e CheckSaleStateEffect) isFlag() bool {
	for i := len(CheckSaleStateEffectAll) - 1; i >= 0; i-- {
		expected := CheckSaleStateEffect(2) << uint64(len(CheckSaleStateEffectAll)-1) >> uint64(len(CheckSaleStateEffectAll)-i)
		if expected != CheckSaleStateEffectAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e CheckSaleStateEffect) String() string {
	name, _ := checkSaleStateEffectMap[int32(e)]
	return name
}

func (e CheckSaleStateEffect) ShortString() string {
	name, _ := checkSaleStateEffectShortMap[int32(e)]
	return name
}

func (e CheckSaleStateEffect) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range CheckSaleStateEffectAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *CheckSaleStateEffect) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = CheckSaleStateEffect(t.Value)
	return nil
}

// SaleCanceledExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type SaleCanceledExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u SaleCanceledExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of SaleCanceledExt
func (u SaleCanceledExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewSaleCanceledExt creates a new  SaleCanceledExt.
func NewSaleCanceledExt(v LedgerVersion, value interface{}) (result SaleCanceledExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// SaleCanceled is an XDR Struct defines as:
//
//   struct SaleCanceled {
//    	 // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type SaleCanceled struct {
	Ext SaleCanceledExt `json:"ext,omitempty"`
}

// SaleUpdatedExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type SaleUpdatedExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u SaleUpdatedExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of SaleUpdatedExt
func (u SaleUpdatedExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewSaleUpdatedExt creates a new  SaleUpdatedExt.
func NewSaleUpdatedExt(v LedgerVersion, value interface{}) (result SaleUpdatedExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// SaleUpdated is an XDR Struct defines as:
//
//   struct SaleUpdated {
//    	 // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type SaleUpdated struct {
	Ext SaleUpdatedExt `json:"ext,omitempty"`
}

// CheckSubSaleClosedResultExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type CheckSubSaleClosedResultExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CheckSubSaleClosedResultExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CheckSubSaleClosedResultExt
func (u CheckSubSaleClosedResultExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewCheckSubSaleClosedResultExt creates a new  CheckSubSaleClosedResultExt.
func NewCheckSubSaleClosedResultExt(v LedgerVersion, value interface{}) (result CheckSubSaleClosedResultExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// CheckSubSaleClosedResult is an XDR Struct defines as:
//
//   struct CheckSubSaleClosedResult {
//    	BalanceID saleBaseBalance;
//    	BalanceID saleQuoteBalance;
//    	ManageOfferSuccessResult saleDetails;
//    	 // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type CheckSubSaleClosedResult struct {
	SaleBaseBalance  BalanceId                   `json:"saleBaseBalance,omitempty"`
	SaleQuoteBalance BalanceId                   `json:"saleQuoteBalance,omitempty"`
	SaleDetails      ManageOfferSuccessResult    `json:"saleDetails,omitempty"`
	Ext              CheckSubSaleClosedResultExt `json:"ext,omitempty"`
}

// CheckSaleClosedResultExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type CheckSaleClosedResultExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CheckSaleClosedResultExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CheckSaleClosedResultExt
func (u CheckSaleClosedResultExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewCheckSaleClosedResultExt creates a new  CheckSaleClosedResultExt.
func NewCheckSaleClosedResultExt(v LedgerVersion, value interface{}) (result CheckSaleClosedResultExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// CheckSaleClosedResult is an XDR Struct defines as:
//
//   struct CheckSaleClosedResult {
//    	AccountID saleOwner;
//    	CheckSubSaleClosedResult results<>;
//    	 // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type CheckSaleClosedResult struct {
	SaleOwner AccountId                  `json:"saleOwner,omitempty"`
	Results   []CheckSubSaleClosedResult `json:"results,omitempty"`
	Ext       CheckSaleClosedResultExt   `json:"ext,omitempty"`
}

// CheckSaleStateSuccessEffect is an XDR NestedUnion defines as:
//
//   union switch (CheckSaleStateEffect effect)
//        {
//        case CANCELED:
//            SaleCanceled saleCanceled;
//    	case CLOSED:
//    		CheckSaleClosedResult saleClosed;
//    	case UPDATED:
//    		SaleUpdated saleUpdated;
//        }
//
type CheckSaleStateSuccessEffect struct {
	Effect       CheckSaleStateEffect   `json:"effect,omitempty"`
	SaleCanceled *SaleCanceled          `json:"saleCanceled,omitempty"`
	SaleClosed   *CheckSaleClosedResult `json:"saleClosed,omitempty"`
	SaleUpdated  *SaleUpdated           `json:"saleUpdated,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CheckSaleStateSuccessEffect) SwitchFieldName() string {
	return "Effect"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CheckSaleStateSuccessEffect
func (u CheckSaleStateSuccessEffect) ArmForSwitch(sw int32) (string, bool) {
	switch CheckSaleStateEffect(sw) {
	case CheckSaleStateEffectCanceled:
		return "SaleCanceled", true
	case CheckSaleStateEffectClosed:
		return "SaleClosed", true
	case CheckSaleStateEffectUpdated:
		return "SaleUpdated", true
	}
	return "-", false
}

// NewCheckSaleStateSuccessEffect creates a new  CheckSaleStateSuccessEffect.
func NewCheckSaleStateSuccessEffect(effect CheckSaleStateEffect, value interface{}) (result CheckSaleStateSuccessEffect, err error) {
	result.Effect = effect
	switch CheckSaleStateEffect(effect) {
	case CheckSaleStateEffectCanceled:
		tv, ok := value.(SaleCanceled)
		if !ok {
			err = fmt.Errorf("invalid value, must be SaleCanceled")
			return
		}
		result.SaleCanceled = &tv
	case CheckSaleStateEffectClosed:
		tv, ok := value.(CheckSaleClosedResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be CheckSaleClosedResult")
			return
		}
		result.SaleClosed = &tv
	case CheckSaleStateEffectUpdated:
		tv, ok := value.(SaleUpdated)
		if !ok {
			err = fmt.Errorf("invalid value, must be SaleUpdated")
			return
		}
		result.SaleUpdated = &tv
	}
	return
}

// MustSaleCanceled retrieves the SaleCanceled value from the union,
// panicing if the value is not set.
func (u CheckSaleStateSuccessEffect) MustSaleCanceled() SaleCanceled {
	val, ok := u.GetSaleCanceled()

	if !ok {
		panic("arm SaleCanceled is not set")
	}

	return val
}

// GetSaleCanceled retrieves the SaleCanceled value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u CheckSaleStateSuccessEffect) GetSaleCanceled() (result SaleCanceled, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Effect))

	if armName == "SaleCanceled" {
		result = *u.SaleCanceled
		ok = true
	}

	return
}

// MustSaleClosed retrieves the SaleClosed value from the union,
// panicing if the value is not set.
func (u CheckSaleStateSuccessEffect) MustSaleClosed() CheckSaleClosedResult {
	val, ok := u.GetSaleClosed()

	if !ok {
		panic("arm SaleClosed is not set")
	}

	return val
}

// GetSaleClosed retrieves the SaleClosed value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u CheckSaleStateSuccessEffect) GetSaleClosed() (result CheckSaleClosedResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Effect))

	if armName == "SaleClosed" {
		result = *u.SaleClosed
		ok = true
	}

	return
}

// MustSaleUpdated retrieves the SaleUpdated value from the union,
// panicing if the value is not set.
func (u CheckSaleStateSuccessEffect) MustSaleUpdated() SaleUpdated {
	val, ok := u.GetSaleUpdated()

	if !ok {
		panic("arm SaleUpdated is not set")
	}

	return val
}

// GetSaleUpdated retrieves the SaleUpdated value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u CheckSaleStateSuccessEffect) GetSaleUpdated() (result SaleUpdated, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Effect))

	if armName == "SaleUpdated" {
		result = *u.SaleUpdated
		ok = true
	}

	return
}

// CheckSaleStateSuccessExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type CheckSaleStateSuccessExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CheckSaleStateSuccessExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CheckSaleStateSuccessExt
func (u CheckSaleStateSuccessExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewCheckSaleStateSuccessExt creates a new  CheckSaleStateSuccessExt.
func NewCheckSaleStateSuccessExt(v LedgerVersion, value interface{}) (result CheckSaleStateSuccessExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// CheckSaleStateSuccess is an XDR Struct defines as:
//
//   struct CheckSaleStateSuccess
//    {
//    	uint64 saleID;
//    	union switch (CheckSaleStateEffect effect)
//        {
//        case CANCELED:
//            SaleCanceled saleCanceled;
//    	case CLOSED:
//    		CheckSaleClosedResult saleClosed;
//    	case UPDATED:
//    		SaleUpdated saleUpdated;
//        }
//        effect;
//    	 // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type CheckSaleStateSuccess struct {
	SaleId Uint64                      `json:"saleID,omitempty"`
	Effect CheckSaleStateSuccessEffect `json:"effect,omitempty"`
	Ext    CheckSaleStateSuccessExt    `json:"ext,omitempty"`
}

// CheckSaleStateResult is an XDR Union defines as:
//
//   union CheckSaleStateResult switch (CheckSaleStateResultCode code)
//    {
//    case SUCCESS:
//        CheckSaleStateSuccess success;
//    default:
//        void;
//    };
//
type CheckSaleStateResult struct {
	Code    CheckSaleStateResultCode `json:"code,omitempty"`
	Success *CheckSaleStateSuccess   `json:"success,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CheckSaleStateResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CheckSaleStateResult
func (u CheckSaleStateResult) ArmForSwitch(sw int32) (string, bool) {
	switch CheckSaleStateResultCode(sw) {
	case CheckSaleStateResultCodeSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewCheckSaleStateResult creates a new  CheckSaleStateResult.
func NewCheckSaleStateResult(code CheckSaleStateResultCode, value interface{}) (result CheckSaleStateResult, err error) {
	result.Code = code
	switch CheckSaleStateResultCode(code) {
	case CheckSaleStateResultCodeSuccess:
		tv, ok := value.(CheckSaleStateSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be CheckSaleStateSuccess")
			return
		}
		result.Success = &tv
	default:
		// void
	}
	return
}

// MustSuccess retrieves the Success value from the union,
// panicing if the value is not set.
func (u CheckSaleStateResult) MustSuccess() CheckSaleStateSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u CheckSaleStateResult) GetSuccess() (result CheckSaleStateSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// CreateAmlAlertRequestOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type CreateAmlAlertRequestOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreateAmlAlertRequestOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreateAmlAlertRequestOpExt
func (u CreateAmlAlertRequestOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewCreateAmlAlertRequestOpExt creates a new  CreateAmlAlertRequestOpExt.
func NewCreateAmlAlertRequestOpExt(v LedgerVersion, value interface{}) (result CreateAmlAlertRequestOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// CreateAmlAlertRequestOp is an XDR Struct defines as:
//
//   struct CreateAMLAlertRequestOp
//    {
//        string64 reference;
//        AMLAlertRequest amlAlertRequest;
//
//    	union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//
//    };
//
type CreateAmlAlertRequestOp struct {
	Reference       String64                   `json:"reference,omitempty"`
	AmlAlertRequest AmlAlertRequest            `json:"amlAlertRequest,omitempty"`
	Ext             CreateAmlAlertRequestOpExt `json:"ext,omitempty"`
}

// CreateAmlAlertRequestResultCode is an XDR Enum defines as:
//
//   enum CreateAMLAlertRequestResultCode
//    {
//        // codes considered as "success" for the operation
//        SUCCESS = 0,
//        BALANCE_NOT_EXIST = 1, // balance doesn't exist
//        INVALID_REASON = 2, //invalid reason for request
//        UNDERFUNDED = 3, //when couldn't lock balance
//    	REFERENCE_DUPLICATION = 4, // reference already exists
//    	INVALID_AMOUNT = 5 // amount must be positive
//
//
//    };
//
type CreateAmlAlertRequestResultCode int32

const (
	CreateAmlAlertRequestResultCodeSuccess              CreateAmlAlertRequestResultCode = 0
	CreateAmlAlertRequestResultCodeBalanceNotExist      CreateAmlAlertRequestResultCode = 1
	CreateAmlAlertRequestResultCodeInvalidReason        CreateAmlAlertRequestResultCode = 2
	CreateAmlAlertRequestResultCodeUnderfunded          CreateAmlAlertRequestResultCode = 3
	CreateAmlAlertRequestResultCodeReferenceDuplication CreateAmlAlertRequestResultCode = 4
	CreateAmlAlertRequestResultCodeInvalidAmount        CreateAmlAlertRequestResultCode = 5
)

var CreateAmlAlertRequestResultCodeAll = []CreateAmlAlertRequestResultCode{
	CreateAmlAlertRequestResultCodeSuccess,
	CreateAmlAlertRequestResultCodeBalanceNotExist,
	CreateAmlAlertRequestResultCodeInvalidReason,
	CreateAmlAlertRequestResultCodeUnderfunded,
	CreateAmlAlertRequestResultCodeReferenceDuplication,
	CreateAmlAlertRequestResultCodeInvalidAmount,
}

var createAmlAlertRequestResultCodeMap = map[int32]string{
	0: "CreateAmlAlertRequestResultCodeSuccess",
	1: "CreateAmlAlertRequestResultCodeBalanceNotExist",
	2: "CreateAmlAlertRequestResultCodeInvalidReason",
	3: "CreateAmlAlertRequestResultCodeUnderfunded",
	4: "CreateAmlAlertRequestResultCodeReferenceDuplication",
	5: "CreateAmlAlertRequestResultCodeInvalidAmount",
}

var createAmlAlertRequestResultCodeShortMap = map[int32]string{
	0: "success",
	1: "balance_not_exist",
	2: "invalid_reason",
	3: "underfunded",
	4: "reference_duplication",
	5: "invalid_amount",
}

var createAmlAlertRequestResultCodeRevMap = map[string]int32{
	"CreateAmlAlertRequestResultCodeSuccess":              0,
	"CreateAmlAlertRequestResultCodeBalanceNotExist":      1,
	"CreateAmlAlertRequestResultCodeInvalidReason":        2,
	"CreateAmlAlertRequestResultCodeUnderfunded":          3,
	"CreateAmlAlertRequestResultCodeReferenceDuplication": 4,
	"CreateAmlAlertRequestResultCodeInvalidAmount":        5,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for CreateAmlAlertRequestResultCode
func (e CreateAmlAlertRequestResultCode) ValidEnum(v int32) bool {
	_, ok := createAmlAlertRequestResultCodeMap[v]
	return ok
}
func (e CreateAmlAlertRequestResultCode) isFlag() bool {
	for i := len(CreateAmlAlertRequestResultCodeAll) - 1; i >= 0; i-- {
		expected := CreateAmlAlertRequestResultCode(2) << uint64(len(CreateAmlAlertRequestResultCodeAll)-1) >> uint64(len(CreateAmlAlertRequestResultCodeAll)-i)
		if expected != CreateAmlAlertRequestResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e CreateAmlAlertRequestResultCode) String() string {
	name, _ := createAmlAlertRequestResultCodeMap[int32(e)]
	return name
}

func (e CreateAmlAlertRequestResultCode) ShortString() string {
	name, _ := createAmlAlertRequestResultCodeShortMap[int32(e)]
	return name
}

func (e CreateAmlAlertRequestResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range CreateAmlAlertRequestResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *CreateAmlAlertRequestResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = CreateAmlAlertRequestResultCode(t.Value)
	return nil
}

// CreateAmlAlertRequestSuccessExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type CreateAmlAlertRequestSuccessExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreateAmlAlertRequestSuccessExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreateAmlAlertRequestSuccessExt
func (u CreateAmlAlertRequestSuccessExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewCreateAmlAlertRequestSuccessExt creates a new  CreateAmlAlertRequestSuccessExt.
func NewCreateAmlAlertRequestSuccessExt(v LedgerVersion, value interface{}) (result CreateAmlAlertRequestSuccessExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// CreateAmlAlertRequestSuccess is an XDR Struct defines as:
//
//   struct CreateAMLAlertRequestSuccess {
//    	uint64 requestID;
//
//    	union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type CreateAmlAlertRequestSuccess struct {
	RequestId Uint64                          `json:"requestID,omitempty"`
	Ext       CreateAmlAlertRequestSuccessExt `json:"ext,omitempty"`
}

// CreateAmlAlertRequestResult is an XDR Union defines as:
//
//   union CreateAMLAlertRequestResult switch (CreateAMLAlertRequestResultCode code)
//    {
//        case SUCCESS:
//            CreateAMLAlertRequestSuccess success;
//        default:
//            void;
//    };
//
type CreateAmlAlertRequestResult struct {
	Code    CreateAmlAlertRequestResultCode `json:"code,omitempty"`
	Success *CreateAmlAlertRequestSuccess   `json:"success,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreateAmlAlertRequestResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreateAmlAlertRequestResult
func (u CreateAmlAlertRequestResult) ArmForSwitch(sw int32) (string, bool) {
	switch CreateAmlAlertRequestResultCode(sw) {
	case CreateAmlAlertRequestResultCodeSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewCreateAmlAlertRequestResult creates a new  CreateAmlAlertRequestResult.
func NewCreateAmlAlertRequestResult(code CreateAmlAlertRequestResultCode, value interface{}) (result CreateAmlAlertRequestResult, err error) {
	result.Code = code
	switch CreateAmlAlertRequestResultCode(code) {
	case CreateAmlAlertRequestResultCodeSuccess:
		tv, ok := value.(CreateAmlAlertRequestSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateAmlAlertRequestSuccess")
			return
		}
		result.Success = &tv
	default:
		// void
	}
	return
}

// MustSuccess retrieves the Success value from the union,
// panicing if the value is not set.
func (u CreateAmlAlertRequestResult) MustSuccess() CreateAmlAlertRequestSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u CreateAmlAlertRequestResult) GetSuccess() (result CreateAmlAlertRequestSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// UpdateKycRequestDataExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type UpdateKycRequestDataExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u UpdateKycRequestDataExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of UpdateKycRequestDataExt
func (u UpdateKycRequestDataExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewUpdateKycRequestDataExt creates a new  UpdateKycRequestDataExt.
func NewUpdateKycRequestDataExt(v LedgerVersion, value interface{}) (result UpdateKycRequestDataExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// UpdateKycRequestData is an XDR Struct defines as:
//
//   struct UpdateKYCRequestData {
//        AccountID accountToUpdateKYC;
//    	AccountType accountTypeToSet;
//    	uint32 kycLevelToSet;
//        longstring kycData;
//    	uint32* allTasks;
//
//    	// Reserved for future use
//    	union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type UpdateKycRequestData struct {
	AccountToUpdateKyc AccountId               `json:"accountToUpdateKYC,omitempty"`
	AccountTypeToSet   AccountType             `json:"accountTypeToSet,omitempty"`
	KycLevelToSet      Uint32                  `json:"kycLevelToSet,omitempty"`
	KycData            Longstring              `json:"kycData,omitempty"`
	AllTasks           *Uint32                 `json:"allTasks,omitempty"`
	Ext                UpdateKycRequestDataExt `json:"ext,omitempty"`
}

// CreateUpdateKycRequestOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type CreateUpdateKycRequestOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreateUpdateKycRequestOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreateUpdateKycRequestOpExt
func (u CreateUpdateKycRequestOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewCreateUpdateKycRequestOpExt creates a new  CreateUpdateKycRequestOpExt.
func NewCreateUpdateKycRequestOpExt(v LedgerVersion, value interface{}) (result CreateUpdateKycRequestOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// CreateUpdateKycRequestOp is an XDR Struct defines as:
//
//   struct CreateUpdateKYCRequestOp {
//        uint64 requestID;
//        UpdateKYCRequestData updateKYCRequestData;
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type CreateUpdateKycRequestOp struct {
	RequestId            Uint64                      `json:"requestID,omitempty"`
	UpdateKycRequestData UpdateKycRequestData        `json:"updateKYCRequestData,omitempty"`
	Ext                  CreateUpdateKycRequestOpExt `json:"ext,omitempty"`
}

// CreateUpdateKycRequestResultCode is an XDR Enum defines as:
//
//   enum CreateUpdateKYCRequestResultCode
//    {
//        // codes considered as "success" for the operation
//        SUCCESS = 0,
//
//        // codes considered as "failure" for the operation
//        ACC_TO_UPDATE_DOES_NOT_EXIST = -1, // account to update does not exist
//        REQUEST_ALREADY_EXISTS = -2,
//    	SAME_ACC_TYPE_TO_SET = -3,
//    	REQUEST_DOES_NOT_EXIST = -4,
//    	PENDING_REQUEST_UPDATE_NOT_ALLOWED = -5,
//    	NOT_ALLOWED_TO_UPDATE_REQUEST = -6, // master account can update request only through review request operation
//    	INVALID_UPDATE_KYC_REQUEST_DATA = -7,
//    	INVALID_KYC_DATA = -8,
//    	KYC_RULE_NOT_FOUND = -9
//    };
//
type CreateUpdateKycRequestResultCode int32

const (
	CreateUpdateKycRequestResultCodeSuccess                        CreateUpdateKycRequestResultCode = 0
	CreateUpdateKycRequestResultCodeAccToUpdateDoesNotExist        CreateUpdateKycRequestResultCode = -1
	CreateUpdateKycRequestResultCodeRequestAlreadyExists           CreateUpdateKycRequestResultCode = -2
	CreateUpdateKycRequestResultCodeSameAccTypeToSet               CreateUpdateKycRequestResultCode = -3
	CreateUpdateKycRequestResultCodeRequestDoesNotExist            CreateUpdateKycRequestResultCode = -4
	CreateUpdateKycRequestResultCodePendingRequestUpdateNotAllowed CreateUpdateKycRequestResultCode = -5
	CreateUpdateKycRequestResultCodeNotAllowedToUpdateRequest      CreateUpdateKycRequestResultCode = -6
	CreateUpdateKycRequestResultCodeInvalidUpdateKycRequestData    CreateUpdateKycRequestResultCode = -7
	CreateUpdateKycRequestResultCodeInvalidKycData                 CreateUpdateKycRequestResultCode = -8
	CreateUpdateKycRequestResultCodeKycRuleNotFound                CreateUpdateKycRequestResultCode = -9
)

var CreateUpdateKycRequestResultCodeAll = []CreateUpdateKycRequestResultCode{
	CreateUpdateKycRequestResultCodeSuccess,
	CreateUpdateKycRequestResultCodeAccToUpdateDoesNotExist,
	CreateUpdateKycRequestResultCodeRequestAlreadyExists,
	CreateUpdateKycRequestResultCodeSameAccTypeToSet,
	CreateUpdateKycRequestResultCodeRequestDoesNotExist,
	CreateUpdateKycRequestResultCodePendingRequestUpdateNotAllowed,
	CreateUpdateKycRequestResultCodeNotAllowedToUpdateRequest,
	CreateUpdateKycRequestResultCodeInvalidUpdateKycRequestData,
	CreateUpdateKycRequestResultCodeInvalidKycData,
	CreateUpdateKycRequestResultCodeKycRuleNotFound,
}

var createUpdateKycRequestResultCodeMap = map[int32]string{
	0:  "CreateUpdateKycRequestResultCodeSuccess",
	-1: "CreateUpdateKycRequestResultCodeAccToUpdateDoesNotExist",
	-2: "CreateUpdateKycRequestResultCodeRequestAlreadyExists",
	-3: "CreateUpdateKycRequestResultCodeSameAccTypeToSet",
	-4: "CreateUpdateKycRequestResultCodeRequestDoesNotExist",
	-5: "CreateUpdateKycRequestResultCodePendingRequestUpdateNotAllowed",
	-6: "CreateUpdateKycRequestResultCodeNotAllowedToUpdateRequest",
	-7: "CreateUpdateKycRequestResultCodeInvalidUpdateKycRequestData",
	-8: "CreateUpdateKycRequestResultCodeInvalidKycData",
	-9: "CreateUpdateKycRequestResultCodeKycRuleNotFound",
}

var createUpdateKycRequestResultCodeShortMap = map[int32]string{
	0:  "success",
	-1: "acc_to_update_does_not_exist",
	-2: "request_already_exists",
	-3: "same_acc_type_to_set",
	-4: "request_does_not_exist",
	-5: "pending_request_update_not_allowed",
	-6: "not_allowed_to_update_request",
	-7: "invalid_update_kyc_request_data",
	-8: "invalid_kyc_data",
	-9: "kyc_rule_not_found",
}

var createUpdateKycRequestResultCodeRevMap = map[string]int32{
	"CreateUpdateKycRequestResultCodeSuccess":                        0,
	"CreateUpdateKycRequestResultCodeAccToUpdateDoesNotExist":        -1,
	"CreateUpdateKycRequestResultCodeRequestAlreadyExists":           -2,
	"CreateUpdateKycRequestResultCodeSameAccTypeToSet":               -3,
	"CreateUpdateKycRequestResultCodeRequestDoesNotExist":            -4,
	"CreateUpdateKycRequestResultCodePendingRequestUpdateNotAllowed": -5,
	"CreateUpdateKycRequestResultCodeNotAllowedToUpdateRequest":      -6,
	"CreateUpdateKycRequestResultCodeInvalidUpdateKycRequestData":    -7,
	"CreateUpdateKycRequestResultCodeInvalidKycData":                 -8,
	"CreateUpdateKycRequestResultCodeKycRuleNotFound":                -9,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for CreateUpdateKycRequestResultCode
func (e CreateUpdateKycRequestResultCode) ValidEnum(v int32) bool {
	_, ok := createUpdateKycRequestResultCodeMap[v]
	return ok
}
func (e CreateUpdateKycRequestResultCode) isFlag() bool {
	for i := len(CreateUpdateKycRequestResultCodeAll) - 1; i >= 0; i-- {
		expected := CreateUpdateKycRequestResultCode(2) << uint64(len(CreateUpdateKycRequestResultCodeAll)-1) >> uint64(len(CreateUpdateKycRequestResultCodeAll)-i)
		if expected != CreateUpdateKycRequestResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e CreateUpdateKycRequestResultCode) String() string {
	name, _ := createUpdateKycRequestResultCodeMap[int32(e)]
	return name
}

func (e CreateUpdateKycRequestResultCode) ShortString() string {
	name, _ := createUpdateKycRequestResultCodeShortMap[int32(e)]
	return name
}

func (e CreateUpdateKycRequestResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range CreateUpdateKycRequestResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *CreateUpdateKycRequestResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = CreateUpdateKycRequestResultCode(t.Value)
	return nil
}

// CreateUpdateKycRequestResultSuccessExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//
type CreateUpdateKycRequestResultSuccessExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreateUpdateKycRequestResultSuccessExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreateUpdateKycRequestResultSuccessExt
func (u CreateUpdateKycRequestResultSuccessExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewCreateUpdateKycRequestResultSuccessExt creates a new  CreateUpdateKycRequestResultSuccessExt.
func NewCreateUpdateKycRequestResultSuccessExt(v LedgerVersion, value interface{}) (result CreateUpdateKycRequestResultSuccessExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// CreateUpdateKycRequestResultSuccess is an XDR NestedStruct defines as:
//
//   struct {
//    		uint64 requestID;
//    		bool fulfilled;
//    		// Reserved for future use
//    		union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//    		ext;
//    	}
//
type CreateUpdateKycRequestResultSuccess struct {
	RequestId Uint64                                 `json:"requestID,omitempty"`
	Fulfilled bool                                   `json:"fulfilled,omitempty"`
	Ext       CreateUpdateKycRequestResultSuccessExt `json:"ext,omitempty"`
}

// CreateUpdateKycRequestResult is an XDR Union defines as:
//
//   union CreateUpdateKYCRequestResult switch (CreateUpdateKYCRequestResultCode code)
//    {
//    case SUCCESS:
//        struct {
//    		uint64 requestID;
//    		bool fulfilled;
//    		// Reserved for future use
//    		union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//    		ext;
//    	} success;
//    default:
//        void;
//    };
//
type CreateUpdateKycRequestResult struct {
	Code    CreateUpdateKycRequestResultCode     `json:"code,omitempty"`
	Success *CreateUpdateKycRequestResultSuccess `json:"success,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreateUpdateKycRequestResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreateUpdateKycRequestResult
func (u CreateUpdateKycRequestResult) ArmForSwitch(sw int32) (string, bool) {
	switch CreateUpdateKycRequestResultCode(sw) {
	case CreateUpdateKycRequestResultCodeSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewCreateUpdateKycRequestResult creates a new  CreateUpdateKycRequestResult.
func NewCreateUpdateKycRequestResult(code CreateUpdateKycRequestResultCode, value interface{}) (result CreateUpdateKycRequestResult, err error) {
	result.Code = code
	switch CreateUpdateKycRequestResultCode(code) {
	case CreateUpdateKycRequestResultCodeSuccess:
		tv, ok := value.(CreateUpdateKycRequestResultSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateUpdateKycRequestResultSuccess")
			return
		}
		result.Success = &tv
	default:
		// void
	}
	return
}

// MustSuccess retrieves the Success value from the union,
// panicing if the value is not set.
func (u CreateUpdateKycRequestResult) MustSuccess() CreateUpdateKycRequestResultSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u CreateUpdateKycRequestResult) GetSuccess() (result CreateUpdateKycRequestResultSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// CreateAccountOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//    	case PASS_EXTERNAL_SYS_ACC_ID_IN_CREATE_ACC:
//    		ExternalSystemAccountID externalSystemIDs<>;
//        }
//
type CreateAccountOpExt struct {
	V                 LedgerVersion              `json:"v,omitempty"`
	ExternalSystemIDs *[]ExternalSystemAccountId `json:"externalSystemIDs,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreateAccountOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreateAccountOpExt
func (u CreateAccountOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	case LedgerVersionPassExternalSysAccIdInCreateAcc:
		return "ExternalSystemIDs", true
	}
	return "-", false
}

// NewCreateAccountOpExt creates a new  CreateAccountOpExt.
func NewCreateAccountOpExt(v LedgerVersion, value interface{}) (result CreateAccountOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	case LedgerVersionPassExternalSysAccIdInCreateAcc:
		tv, ok := value.([]ExternalSystemAccountId)
		if !ok {
			err = fmt.Errorf("invalid value, must be []ExternalSystemAccountId")
			return
		}
		result.ExternalSystemIDs = &tv
	}
	return
}

// MustExternalSystemIDs retrieves the ExternalSystemIDs value from the union,
// panicing if the value is not set.
func (u CreateAccountOpExt) MustExternalSystemIDs() []ExternalSystemAccountId {
	val, ok := u.GetExternalSystemIDs()

	if !ok {
		panic("arm ExternalSystemIDs is not set")
	}

	return val
}

// GetExternalSystemIDs retrieves the ExternalSystemIDs value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u CreateAccountOpExt) GetExternalSystemIDs() (result []ExternalSystemAccountId, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.V))

	if armName == "ExternalSystemIDs" {
		result = *u.ExternalSystemIDs
		ok = true
	}

	return
}

// CreateAccountOp is an XDR Struct defines as:
//
//   struct CreateAccountOp
//    {
//        AccountID destination; // account to create
//        AccountID recoveryKey; // recovery signer's public key
//        AccountID* referrer;     // parent account
//    	AccountType accountType;
//    	uint32 policies;
//
//    	 // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//    	case PASS_EXTERNAL_SYS_ACC_ID_IN_CREATE_ACC:
//    		ExternalSystemAccountID externalSystemIDs<>;
//        }
//        ext;
//    };
//
type CreateAccountOp struct {
	Destination AccountId          `json:"destination,omitempty"`
	RecoveryKey AccountId          `json:"recoveryKey,omitempty"`
	Referrer    *AccountId         `json:"referrer,omitempty"`
	AccountType AccountType        `json:"accountType,omitempty"`
	Policies    Uint32             `json:"policies,omitempty"`
	Ext         CreateAccountOpExt `json:"ext,omitempty"`
}

// CreateAccountResultCode is an XDR Enum defines as:
//
//   enum CreateAccountResultCode
//    {
//        // codes considered as "success" for the operation
//        SUCCESS = 0, // account was created
//
//        // codes considered as "failure" for the operation
//        MALFORMED = -1,       // invalid destination
//    	ACCOUNT_TYPE_MISMATCHED = -2, // account already exist and change of account type is not allowed
//    	TYPE_NOT_ALLOWED = -3, // master or commission account types are not allowed
//        NAME_DUPLICATION = -4,
//        REFERRER_NOT_FOUND = -5,
//    	INVALID_ACCOUNT_VERSION = -6, // if account version is higher than ledger version
//    	NOT_VERIFIED_CANNOT_HAVE_POLICIES = -7,
//    	EXTERNAL_SYS_ACC_NOT_ALLOWED = -8, // op contains external system account ID which should be generated on core level
//    	EXTERNAL_SYS_ID_EXISTS = -9 // external system account ID already exists
//    };
//
type CreateAccountResultCode int32

const (
	CreateAccountResultCodeSuccess                       CreateAccountResultCode = 0
	CreateAccountResultCodeMalformed                     CreateAccountResultCode = -1
	CreateAccountResultCodeAccountTypeMismatched         CreateAccountResultCode = -2
	CreateAccountResultCodeTypeNotAllowed                CreateAccountResultCode = -3
	CreateAccountResultCodeNameDuplication               CreateAccountResultCode = -4
	CreateAccountResultCodeReferrerNotFound              CreateAccountResultCode = -5
	CreateAccountResultCodeInvalidAccountVersion         CreateAccountResultCode = -6
	CreateAccountResultCodeNotVerifiedCannotHavePolicies CreateAccountResultCode = -7
	CreateAccountResultCodeExternalSysAccNotAllowed      CreateAccountResultCode = -8
	CreateAccountResultCodeExternalSysIdExists           CreateAccountResultCode = -9
)

var CreateAccountResultCodeAll = []CreateAccountResultCode{
	CreateAccountResultCodeSuccess,
	CreateAccountResultCodeMalformed,
	CreateAccountResultCodeAccountTypeMismatched,
	CreateAccountResultCodeTypeNotAllowed,
	CreateAccountResultCodeNameDuplication,
	CreateAccountResultCodeReferrerNotFound,
	CreateAccountResultCodeInvalidAccountVersion,
	CreateAccountResultCodeNotVerifiedCannotHavePolicies,
	CreateAccountResultCodeExternalSysAccNotAllowed,
	CreateAccountResultCodeExternalSysIdExists,
}

var createAccountResultCodeMap = map[int32]string{
	0:  "CreateAccountResultCodeSuccess",
	-1: "CreateAccountResultCodeMalformed",
	-2: "CreateAccountResultCodeAccountTypeMismatched",
	-3: "CreateAccountResultCodeTypeNotAllowed",
	-4: "CreateAccountResultCodeNameDuplication",
	-5: "CreateAccountResultCodeReferrerNotFound",
	-6: "CreateAccountResultCodeInvalidAccountVersion",
	-7: "CreateAccountResultCodeNotVerifiedCannotHavePolicies",
	-8: "CreateAccountResultCodeExternalSysAccNotAllowed",
	-9: "CreateAccountResultCodeExternalSysIdExists",
}

var createAccountResultCodeShortMap = map[int32]string{
	0:  "success",
	-1: "malformed",
	-2: "account_type_mismatched",
	-3: "type_not_allowed",
	-4: "name_duplication",
	-5: "referrer_not_found",
	-6: "invalid_account_version",
	-7: "not_verified_cannot_have_policies",
	-8: "external_sys_acc_not_allowed",
	-9: "external_sys_id_exists",
}

var createAccountResultCodeRevMap = map[string]int32{
	"CreateAccountResultCodeSuccess":                       0,
	"CreateAccountResultCodeMalformed":                     -1,
	"CreateAccountResultCodeAccountTypeMismatched":         -2,
	"CreateAccountResultCodeTypeNotAllowed":                -3,
	"CreateAccountResultCodeNameDuplication":               -4,
	"CreateAccountResultCodeReferrerNotFound":              -5,
	"CreateAccountResultCodeInvalidAccountVersion":         -6,
	"CreateAccountResultCodeNotVerifiedCannotHavePolicies": -7,
	"CreateAccountResultCodeExternalSysAccNotAllowed":      -8,
	"CreateAccountResultCodeExternalSysIdExists":           -9,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for CreateAccountResultCode
func (e CreateAccountResultCode) ValidEnum(v int32) bool {
	_, ok := createAccountResultCodeMap[v]
	return ok
}
func (e CreateAccountResultCode) isFlag() bool {
	for i := len(CreateAccountResultCodeAll) - 1; i >= 0; i-- {
		expected := CreateAccountResultCode(2) << uint64(len(CreateAccountResultCodeAll)-1) >> uint64(len(CreateAccountResultCodeAll)-i)
		if expected != CreateAccountResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e CreateAccountResultCode) String() string {
	name, _ := createAccountResultCodeMap[int32(e)]
	return name
}

func (e CreateAccountResultCode) ShortString() string {
	name, _ := createAccountResultCodeShortMap[int32(e)]
	return name
}

func (e CreateAccountResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range CreateAccountResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *CreateAccountResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = CreateAccountResultCode(t.Value)
	return nil
}

// CreateAccountSuccessExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type CreateAccountSuccessExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreateAccountSuccessExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreateAccountSuccessExt
func (u CreateAccountSuccessExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewCreateAccountSuccessExt creates a new  CreateAccountSuccessExt.
func NewCreateAccountSuccessExt(v LedgerVersion, value interface{}) (result CreateAccountSuccessExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// CreateAccountSuccess is an XDR Struct defines as:
//
//   struct CreateAccountSuccess
//    {
//    	ExternalSystemAccountID externalSystemIDs<>;
//    	 // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type CreateAccountSuccess struct {
	ExternalSystemIDs []ExternalSystemAccountId `json:"externalSystemIDs,omitempty"`
	Ext               CreateAccountSuccessExt   `json:"ext,omitempty"`
}

// CreateAccountResult is an XDR Union defines as:
//
//   union CreateAccountResult switch (CreateAccountResultCode code)
//    {
//    case SUCCESS:
//        CreateAccountSuccess success;
//    default:
//        void;
//    };
//
type CreateAccountResult struct {
	Code    CreateAccountResultCode `json:"code,omitempty"`
	Success *CreateAccountSuccess   `json:"success,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreateAccountResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreateAccountResult
func (u CreateAccountResult) ArmForSwitch(sw int32) (string, bool) {
	switch CreateAccountResultCode(sw) {
	case CreateAccountResultCodeSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewCreateAccountResult creates a new  CreateAccountResult.
func NewCreateAccountResult(code CreateAccountResultCode, value interface{}) (result CreateAccountResult, err error) {
	result.Code = code
	switch CreateAccountResultCode(code) {
	case CreateAccountResultCodeSuccess:
		tv, ok := value.(CreateAccountSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateAccountSuccess")
			return
		}
		result.Success = &tv
	default:
		// void
	}
	return
}

// MustSuccess retrieves the Success value from the union,
// panicing if the value is not set.
func (u CreateAccountResult) MustSuccess() CreateAccountSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u CreateAccountResult) GetSuccess() (result CreateAccountSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// CreateIssuanceRequestOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        case ADD_TASKS_TO_REVIEWABLE_REQUEST:
//            uint32* allTasks;
//        }
//
type CreateIssuanceRequestOpExt struct {
	V        LedgerVersion `json:"v,omitempty"`
	AllTasks **Uint32      `json:"allTasks,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreateIssuanceRequestOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreateIssuanceRequestOpExt
func (u CreateIssuanceRequestOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	case LedgerVersionAddTasksToReviewableRequest:
		return "AllTasks", true
	}
	return "-", false
}

// NewCreateIssuanceRequestOpExt creates a new  CreateIssuanceRequestOpExt.
func NewCreateIssuanceRequestOpExt(v LedgerVersion, value interface{}) (result CreateIssuanceRequestOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	case LedgerVersionAddTasksToReviewableRequest:
		tv, ok := value.(*Uint32)
		if !ok {
			err = fmt.Errorf("invalid value, must be *Uint32")
			return
		}
		result.AllTasks = &tv
	}
	return
}

// MustAllTasks retrieves the AllTasks value from the union,
// panicing if the value is not set.
func (u CreateIssuanceRequestOpExt) MustAllTasks() *Uint32 {
	val, ok := u.GetAllTasks()

	if !ok {
		panic("arm AllTasks is not set")
	}

	return val
}

// GetAllTasks retrieves the AllTasks value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u CreateIssuanceRequestOpExt) GetAllTasks() (result *Uint32, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.V))

	if armName == "AllTasks" {
		result = *u.AllTasks
		ok = true
	}

	return
}

// CreateIssuanceRequestOp is an XDR Struct defines as:
//
//   struct CreateIssuanceRequestOp
//    {
//    	IssuanceRequest request;
//    	string64 reference;
//    	// reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        case ADD_TASKS_TO_REVIEWABLE_REQUEST:
//            uint32* allTasks;
//        }
//        ext;
//    };
//
type CreateIssuanceRequestOp struct {
	Request   IssuanceRequest            `json:"request,omitempty"`
	Reference String64                   `json:"reference,omitempty"`
	Ext       CreateIssuanceRequestOpExt `json:"ext,omitempty"`
}

// CreateIssuanceRequestResultCode is an XDR Enum defines as:
//
//   enum CreateIssuanceRequestResultCode
//    {
//        // codes considered as "success" for the operation
//        SUCCESS = 0,
//
//        // codes considered as "failure" for the operation
//        ASSET_NOT_FOUND = -1,
//    	INVALID_AMOUNT = -2,
//    	REFERENCE_DUPLICATION = -3,
//    	NO_COUNTERPARTY = -4,
//    	NOT_AUTHORIZED = -5,
//    	EXCEEDS_MAX_ISSUANCE_AMOUNT = -6,
//    	RECEIVER_FULL_LINE = -7,
//    	INVALID_EXTERNAL_DETAILS = -8, // external details size exceeds max allowed
//    	FEE_EXCEEDS_AMOUNT = -9, // fee more than amount to issue
//        REQUIRES_KYC = -10, // asset requires receiver to have KYC
//        REQUIRES_VERIFICATION = -11, //asset requires receiver to be verified
//        ISSUANCE_TASKS_NOT_FOUND = -12, // issuance tasks have not been provided by the source and don't exist in 'KeyValue' table
//        SYSTEM_TASKS_NOT_ALLOWED = -13
//    };
//
type CreateIssuanceRequestResultCode int32

const (
	CreateIssuanceRequestResultCodeSuccess                  CreateIssuanceRequestResultCode = 0
	CreateIssuanceRequestResultCodeAssetNotFound            CreateIssuanceRequestResultCode = -1
	CreateIssuanceRequestResultCodeInvalidAmount            CreateIssuanceRequestResultCode = -2
	CreateIssuanceRequestResultCodeReferenceDuplication     CreateIssuanceRequestResultCode = -3
	CreateIssuanceRequestResultCodeNoCounterparty           CreateIssuanceRequestResultCode = -4
	CreateIssuanceRequestResultCodeNotAuthorized            CreateIssuanceRequestResultCode = -5
	CreateIssuanceRequestResultCodeExceedsMaxIssuanceAmount CreateIssuanceRequestResultCode = -6
	CreateIssuanceRequestResultCodeReceiverFullLine         CreateIssuanceRequestResultCode = -7
	CreateIssuanceRequestResultCodeInvalidExternalDetails   CreateIssuanceRequestResultCode = -8
	CreateIssuanceRequestResultCodeFeeExceedsAmount         CreateIssuanceRequestResultCode = -9
	CreateIssuanceRequestResultCodeRequiresKyc              CreateIssuanceRequestResultCode = -10
	CreateIssuanceRequestResultCodeRequiresVerification     CreateIssuanceRequestResultCode = -11
	CreateIssuanceRequestResultCodeIssuanceTasksNotFound    CreateIssuanceRequestResultCode = -12
	CreateIssuanceRequestResultCodeSystemTasksNotAllowed    CreateIssuanceRequestResultCode = -13
)

var CreateIssuanceRequestResultCodeAll = []CreateIssuanceRequestResultCode{
	CreateIssuanceRequestResultCodeSuccess,
	CreateIssuanceRequestResultCodeAssetNotFound,
	CreateIssuanceRequestResultCodeInvalidAmount,
	CreateIssuanceRequestResultCodeReferenceDuplication,
	CreateIssuanceRequestResultCodeNoCounterparty,
	CreateIssuanceRequestResultCodeNotAuthorized,
	CreateIssuanceRequestResultCodeExceedsMaxIssuanceAmount,
	CreateIssuanceRequestResultCodeReceiverFullLine,
	CreateIssuanceRequestResultCodeInvalidExternalDetails,
	CreateIssuanceRequestResultCodeFeeExceedsAmount,
	CreateIssuanceRequestResultCodeRequiresKyc,
	CreateIssuanceRequestResultCodeRequiresVerification,
	CreateIssuanceRequestResultCodeIssuanceTasksNotFound,
	CreateIssuanceRequestResultCodeSystemTasksNotAllowed,
}

var createIssuanceRequestResultCodeMap = map[int32]string{
	0:   "CreateIssuanceRequestResultCodeSuccess",
	-1:  "CreateIssuanceRequestResultCodeAssetNotFound",
	-2:  "CreateIssuanceRequestResultCodeInvalidAmount",
	-3:  "CreateIssuanceRequestResultCodeReferenceDuplication",
	-4:  "CreateIssuanceRequestResultCodeNoCounterparty",
	-5:  "CreateIssuanceRequestResultCodeNotAuthorized",
	-6:  "CreateIssuanceRequestResultCodeExceedsMaxIssuanceAmount",
	-7:  "CreateIssuanceRequestResultCodeReceiverFullLine",
	-8:  "CreateIssuanceRequestResultCodeInvalidExternalDetails",
	-9:  "CreateIssuanceRequestResultCodeFeeExceedsAmount",
	-10: "CreateIssuanceRequestResultCodeRequiresKyc",
	-11: "CreateIssuanceRequestResultCodeRequiresVerification",
	-12: "CreateIssuanceRequestResultCodeIssuanceTasksNotFound",
	-13: "CreateIssuanceRequestResultCodeSystemTasksNotAllowed",
}

var createIssuanceRequestResultCodeShortMap = map[int32]string{
	0:   "success",
	-1:  "asset_not_found",
	-2:  "invalid_amount",
	-3:  "reference_duplication",
	-4:  "no_counterparty",
	-5:  "not_authorized",
	-6:  "exceeds_max_issuance_amount",
	-7:  "receiver_full_line",
	-8:  "invalid_external_details",
	-9:  "fee_exceeds_amount",
	-10: "requires_kyc",
	-11: "requires_verification",
	-12: "issuance_tasks_not_found",
	-13: "system_tasks_not_allowed",
}

var createIssuanceRequestResultCodeRevMap = map[string]int32{
	"CreateIssuanceRequestResultCodeSuccess":                  0,
	"CreateIssuanceRequestResultCodeAssetNotFound":            -1,
	"CreateIssuanceRequestResultCodeInvalidAmount":            -2,
	"CreateIssuanceRequestResultCodeReferenceDuplication":     -3,
	"CreateIssuanceRequestResultCodeNoCounterparty":           -4,
	"CreateIssuanceRequestResultCodeNotAuthorized":            -5,
	"CreateIssuanceRequestResultCodeExceedsMaxIssuanceAmount": -6,
	"CreateIssuanceRequestResultCodeReceiverFullLine":         -7,
	"CreateIssuanceRequestResultCodeInvalidExternalDetails":   -8,
	"CreateIssuanceRequestResultCodeFeeExceedsAmount":         -9,
	"CreateIssuanceRequestResultCodeRequiresKyc":              -10,
	"CreateIssuanceRequestResultCodeRequiresVerification":     -11,
	"CreateIssuanceRequestResultCodeIssuanceTasksNotFound":    -12,
	"CreateIssuanceRequestResultCodeSystemTasksNotAllowed":    -13,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for CreateIssuanceRequestResultCode
func (e CreateIssuanceRequestResultCode) ValidEnum(v int32) bool {
	_, ok := createIssuanceRequestResultCodeMap[v]
	return ok
}
func (e CreateIssuanceRequestResultCode) isFlag() bool {
	for i := len(CreateIssuanceRequestResultCodeAll) - 1; i >= 0; i-- {
		expected := CreateIssuanceRequestResultCode(2) << uint64(len(CreateIssuanceRequestResultCodeAll)-1) >> uint64(len(CreateIssuanceRequestResultCodeAll)-i)
		if expected != CreateIssuanceRequestResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e CreateIssuanceRequestResultCode) String() string {
	name, _ := createIssuanceRequestResultCodeMap[int32(e)]
	return name
}

func (e CreateIssuanceRequestResultCode) ShortString() string {
	name, _ := createIssuanceRequestResultCodeShortMap[int32(e)]
	return name
}

func (e CreateIssuanceRequestResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range CreateIssuanceRequestResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *CreateIssuanceRequestResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = CreateIssuanceRequestResultCode(t.Value)
	return nil
}

// CreateIssuanceRequestSuccessExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//    	{
//    	case EMPTY_VERSION:
//    		void;
//    	}
//
type CreateIssuanceRequestSuccessExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreateIssuanceRequestSuccessExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreateIssuanceRequestSuccessExt
func (u CreateIssuanceRequestSuccessExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewCreateIssuanceRequestSuccessExt creates a new  CreateIssuanceRequestSuccessExt.
func NewCreateIssuanceRequestSuccessExt(v LedgerVersion, value interface{}) (result CreateIssuanceRequestSuccessExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// CreateIssuanceRequestSuccess is an XDR Struct defines as:
//
//   struct CreateIssuanceRequestSuccess {
//    	uint64 requestID;
//    	AccountID receiver;
//    	bool fulfilled;
//    	Fee fee;
//    	union switch (LedgerVersion v)
//    	{
//    	case EMPTY_VERSION:
//    		void;
//    	}
//    	ext;
//    };
//
type CreateIssuanceRequestSuccess struct {
	RequestId Uint64                          `json:"requestID,omitempty"`
	Receiver  AccountId                       `json:"receiver,omitempty"`
	Fulfilled bool                            `json:"fulfilled,omitempty"`
	Fee       Fee                             `json:"fee,omitempty"`
	Ext       CreateIssuanceRequestSuccessExt `json:"ext,omitempty"`
}

// CreateIssuanceRequestResult is an XDR Union defines as:
//
//   union CreateIssuanceRequestResult switch (CreateIssuanceRequestResultCode code)
//    {
//    case SUCCESS:
//        CreateIssuanceRequestSuccess success;
//    default:
//        void;
//    };
//
type CreateIssuanceRequestResult struct {
	Code    CreateIssuanceRequestResultCode `json:"code,omitempty"`
	Success *CreateIssuanceRequestSuccess   `json:"success,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreateIssuanceRequestResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreateIssuanceRequestResult
func (u CreateIssuanceRequestResult) ArmForSwitch(sw int32) (string, bool) {
	switch CreateIssuanceRequestResultCode(sw) {
	case CreateIssuanceRequestResultCodeSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewCreateIssuanceRequestResult creates a new  CreateIssuanceRequestResult.
func NewCreateIssuanceRequestResult(code CreateIssuanceRequestResultCode, value interface{}) (result CreateIssuanceRequestResult, err error) {
	result.Code = code
	switch CreateIssuanceRequestResultCode(code) {
	case CreateIssuanceRequestResultCodeSuccess:
		tv, ok := value.(CreateIssuanceRequestSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateIssuanceRequestSuccess")
			return
		}
		result.Success = &tv
	default:
		// void
	}
	return
}

// MustSuccess retrieves the Success value from the union,
// panicing if the value is not set.
func (u CreateIssuanceRequestResult) MustSuccess() CreateIssuanceRequestSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u CreateIssuanceRequestResult) GetSuccess() (result CreateIssuanceRequestSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// CreateManageLimitsRequestOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//    	{
//    	case EMPTY_VERSION:
//    		void;
//        case ALLOW_TO_UPDATE_AND_REJECT_LIMITS_UPDATE_REQUESTS:
//            uint64 requestID;
//    	}
//
type CreateManageLimitsRequestOpExt struct {
	V         LedgerVersion `json:"v,omitempty"`
	RequestId *Uint64       `json:"requestID,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreateManageLimitsRequestOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreateManageLimitsRequestOpExt
func (u CreateManageLimitsRequestOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	case LedgerVersionAllowToUpdateAndRejectLimitsUpdateRequests:
		return "RequestId", true
	}
	return "-", false
}

// NewCreateManageLimitsRequestOpExt creates a new  CreateManageLimitsRequestOpExt.
func NewCreateManageLimitsRequestOpExt(v LedgerVersion, value interface{}) (result CreateManageLimitsRequestOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	case LedgerVersionAllowToUpdateAndRejectLimitsUpdateRequests:
		tv, ok := value.(Uint64)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint64")
			return
		}
		result.RequestId = &tv
	}
	return
}

// MustRequestId retrieves the RequestId value from the union,
// panicing if the value is not set.
func (u CreateManageLimitsRequestOpExt) MustRequestId() Uint64 {
	val, ok := u.GetRequestId()

	if !ok {
		panic("arm RequestId is not set")
	}

	return val
}

// GetRequestId retrieves the RequestId value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u CreateManageLimitsRequestOpExt) GetRequestId() (result Uint64, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.V))

	if armName == "RequestId" {
		result = *u.RequestId
		ok = true
	}

	return
}

// CreateManageLimitsRequestOp is an XDR Struct defines as:
//
//   struct CreateManageLimitsRequestOp
//    {
//        LimitsUpdateRequest manageLimitsRequest;
//
//    	// reserved for future use
//    	union switch (LedgerVersion v)
//    	{
//    	case EMPTY_VERSION:
//    		void;
//        case ALLOW_TO_UPDATE_AND_REJECT_LIMITS_UPDATE_REQUESTS:
//            uint64 requestID;
//    	}
//    	ext;
//
//    };
//
type CreateManageLimitsRequestOp struct {
	ManageLimitsRequest LimitsUpdateRequest            `json:"manageLimitsRequest,omitempty"`
	Ext                 CreateManageLimitsRequestOpExt `json:"ext,omitempty"`
}

// CreateManageLimitsRequestResultCode is an XDR Enum defines as:
//
//   enum CreateManageLimitsRequestResultCode
//    {
//        // codes considered as "success" for the operation
//        SUCCESS = 0,
//
//        // codes considered as "failure" for the operation
//    	MANAGE_LIMITS_REQUEST_REFERENCE_DUPLICATION = -1,
//        MANAGE_LIMITS_REQUEST_NOT_FOUND = -2,
//        INVALID_DETAILS = -3, // details must be valid json
//        INVALID_MANAGE_LIMITS_REQUEST_VERSION = -4 // a version of the request is higher than ledger version
//    };
//
type CreateManageLimitsRequestResultCode int32

const (
	CreateManageLimitsRequestResultCodeSuccess                                 CreateManageLimitsRequestResultCode = 0
	CreateManageLimitsRequestResultCodeManageLimitsRequestReferenceDuplication CreateManageLimitsRequestResultCode = -1
	CreateManageLimitsRequestResultCodeManageLimitsRequestNotFound             CreateManageLimitsRequestResultCode = -2
	CreateManageLimitsRequestResultCodeInvalidDetails                          CreateManageLimitsRequestResultCode = -3
	CreateManageLimitsRequestResultCodeInvalidManageLimitsRequestVersion       CreateManageLimitsRequestResultCode = -4
)

var CreateManageLimitsRequestResultCodeAll = []CreateManageLimitsRequestResultCode{
	CreateManageLimitsRequestResultCodeSuccess,
	CreateManageLimitsRequestResultCodeManageLimitsRequestReferenceDuplication,
	CreateManageLimitsRequestResultCodeManageLimitsRequestNotFound,
	CreateManageLimitsRequestResultCodeInvalidDetails,
	CreateManageLimitsRequestResultCodeInvalidManageLimitsRequestVersion,
}

var createManageLimitsRequestResultCodeMap = map[int32]string{
	0:  "CreateManageLimitsRequestResultCodeSuccess",
	-1: "CreateManageLimitsRequestResultCodeManageLimitsRequestReferenceDuplication",
	-2: "CreateManageLimitsRequestResultCodeManageLimitsRequestNotFound",
	-3: "CreateManageLimitsRequestResultCodeInvalidDetails",
	-4: "CreateManageLimitsRequestResultCodeInvalidManageLimitsRequestVersion",
}

var createManageLimitsRequestResultCodeShortMap = map[int32]string{
	0:  "success",
	-1: "manage_limits_request_reference_duplication",
	-2: "manage_limits_request_not_found",
	-3: "invalid_details",
	-4: "invalid_manage_limits_request_version",
}

var createManageLimitsRequestResultCodeRevMap = map[string]int32{
	"CreateManageLimitsRequestResultCodeSuccess":                                 0,
	"CreateManageLimitsRequestResultCodeManageLimitsRequestReferenceDuplication": -1,
	"CreateManageLimitsRequestResultCodeManageLimitsRequestNotFound":             -2,
	"CreateManageLimitsRequestResultCodeInvalidDetails":                          -3,
	"CreateManageLimitsRequestResultCodeInvalidManageLimitsRequestVersion":       -4,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for CreateManageLimitsRequestResultCode
func (e CreateManageLimitsRequestResultCode) ValidEnum(v int32) bool {
	_, ok := createManageLimitsRequestResultCodeMap[v]
	return ok
}
func (e CreateManageLimitsRequestResultCode) isFlag() bool {
	for i := len(CreateManageLimitsRequestResultCodeAll) - 1; i >= 0; i-- {
		expected := CreateManageLimitsRequestResultCode(2) << uint64(len(CreateManageLimitsRequestResultCodeAll)-1) >> uint64(len(CreateManageLimitsRequestResultCodeAll)-i)
		if expected != CreateManageLimitsRequestResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e CreateManageLimitsRequestResultCode) String() string {
	name, _ := createManageLimitsRequestResultCodeMap[int32(e)]
	return name
}

func (e CreateManageLimitsRequestResultCode) ShortString() string {
	name, _ := createManageLimitsRequestResultCodeShortMap[int32(e)]
	return name
}

func (e CreateManageLimitsRequestResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range CreateManageLimitsRequestResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *CreateManageLimitsRequestResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = CreateManageLimitsRequestResultCode(t.Value)
	return nil
}

// CreateManageLimitsRequestResultSuccessExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//
type CreateManageLimitsRequestResultSuccessExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreateManageLimitsRequestResultSuccessExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreateManageLimitsRequestResultSuccessExt
func (u CreateManageLimitsRequestResultSuccessExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewCreateManageLimitsRequestResultSuccessExt creates a new  CreateManageLimitsRequestResultSuccessExt.
func NewCreateManageLimitsRequestResultSuccessExt(v LedgerVersion, value interface{}) (result CreateManageLimitsRequestResultSuccessExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// CreateManageLimitsRequestResultSuccess is an XDR NestedStruct defines as:
//
//   struct {
//            uint64 manageLimitsRequestID;
//    		// reserved for future use
//    		union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//    		ext;
//    	}
//
type CreateManageLimitsRequestResultSuccess struct {
	ManageLimitsRequestId Uint64                                    `json:"manageLimitsRequestID,omitempty"`
	Ext                   CreateManageLimitsRequestResultSuccessExt `json:"ext,omitempty"`
}

// CreateManageLimitsRequestResult is an XDR Union defines as:
//
//   union CreateManageLimitsRequestResult switch (CreateManageLimitsRequestResultCode code)
//    {
//    case SUCCESS:
//        struct {
//            uint64 manageLimitsRequestID;
//    		// reserved for future use
//    		union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//    		ext;
//    	} success;
//    default:
//        void;
//    };
//
type CreateManageLimitsRequestResult struct {
	Code    CreateManageLimitsRequestResultCode     `json:"code,omitempty"`
	Success *CreateManageLimitsRequestResultSuccess `json:"success,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreateManageLimitsRequestResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreateManageLimitsRequestResult
func (u CreateManageLimitsRequestResult) ArmForSwitch(sw int32) (string, bool) {
	switch CreateManageLimitsRequestResultCode(sw) {
	case CreateManageLimitsRequestResultCodeSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewCreateManageLimitsRequestResult creates a new  CreateManageLimitsRequestResult.
func NewCreateManageLimitsRequestResult(code CreateManageLimitsRequestResultCode, value interface{}) (result CreateManageLimitsRequestResult, err error) {
	result.Code = code
	switch CreateManageLimitsRequestResultCode(code) {
	case CreateManageLimitsRequestResultCodeSuccess:
		tv, ok := value.(CreateManageLimitsRequestResultSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateManageLimitsRequestResultSuccess")
			return
		}
		result.Success = &tv
	default:
		// void
	}
	return
}

// MustSuccess retrieves the Success value from the union,
// panicing if the value is not set.
func (u CreateManageLimitsRequestResult) MustSuccess() CreateManageLimitsRequestResultSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u CreateManageLimitsRequestResult) GetSuccess() (result CreateManageLimitsRequestResultSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// CreatePreIssuanceRequestOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type CreatePreIssuanceRequestOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreatePreIssuanceRequestOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreatePreIssuanceRequestOpExt
func (u CreatePreIssuanceRequestOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewCreatePreIssuanceRequestOpExt creates a new  CreatePreIssuanceRequestOpExt.
func NewCreatePreIssuanceRequestOpExt(v LedgerVersion, value interface{}) (result CreatePreIssuanceRequestOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// CreatePreIssuanceRequestOp is an XDR Struct defines as:
//
//   struct CreatePreIssuanceRequestOp
//    {
//        PreIssuanceRequest request;
//    	// reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type CreatePreIssuanceRequestOp struct {
	Request PreIssuanceRequest            `json:"request,omitempty"`
	Ext     CreatePreIssuanceRequestOpExt `json:"ext,omitempty"`
}

// CreatePreIssuanceRequestResultCode is an XDR Enum defines as:
//
//   enum CreatePreIssuanceRequestResultCode
//    {
//        // codes considered as "success" for the operation
//        SUCCESS = 0,
//
//        // codes considered as "failure" for the operation
//        ASSET_NOT_FOUND = -1,
//        REFERENCE_DUPLICATION = -2,    // reference is already used
//        NOT_AUTHORIZED_UPLOAD = -3, // tries to pre issue asset for not owned asset
//        INVALID_SIGNATURE = -4,
//        EXCEEDED_MAX_AMOUNT = -5,
//    	INVALID_AMOUNT = -6,
//    	INVALID_REFERENCE = -7
//    };
//
type CreatePreIssuanceRequestResultCode int32

const (
	CreatePreIssuanceRequestResultCodeSuccess              CreatePreIssuanceRequestResultCode = 0
	CreatePreIssuanceRequestResultCodeAssetNotFound        CreatePreIssuanceRequestResultCode = -1
	CreatePreIssuanceRequestResultCodeReferenceDuplication CreatePreIssuanceRequestResultCode = -2
	CreatePreIssuanceRequestResultCodeNotAuthorizedUpload  CreatePreIssuanceRequestResultCode = -3
	CreatePreIssuanceRequestResultCodeInvalidSignature     CreatePreIssuanceRequestResultCode = -4
	CreatePreIssuanceRequestResultCodeExceededMaxAmount    CreatePreIssuanceRequestResultCode = -5
	CreatePreIssuanceRequestResultCodeInvalidAmount        CreatePreIssuanceRequestResultCode = -6
	CreatePreIssuanceRequestResultCodeInvalidReference     CreatePreIssuanceRequestResultCode = -7
)

var CreatePreIssuanceRequestResultCodeAll = []CreatePreIssuanceRequestResultCode{
	CreatePreIssuanceRequestResultCodeSuccess,
	CreatePreIssuanceRequestResultCodeAssetNotFound,
	CreatePreIssuanceRequestResultCodeReferenceDuplication,
	CreatePreIssuanceRequestResultCodeNotAuthorizedUpload,
	CreatePreIssuanceRequestResultCodeInvalidSignature,
	CreatePreIssuanceRequestResultCodeExceededMaxAmount,
	CreatePreIssuanceRequestResultCodeInvalidAmount,
	CreatePreIssuanceRequestResultCodeInvalidReference,
}

var createPreIssuanceRequestResultCodeMap = map[int32]string{
	0:  "CreatePreIssuanceRequestResultCodeSuccess",
	-1: "CreatePreIssuanceRequestResultCodeAssetNotFound",
	-2: "CreatePreIssuanceRequestResultCodeReferenceDuplication",
	-3: "CreatePreIssuanceRequestResultCodeNotAuthorizedUpload",
	-4: "CreatePreIssuanceRequestResultCodeInvalidSignature",
	-5: "CreatePreIssuanceRequestResultCodeExceededMaxAmount",
	-6: "CreatePreIssuanceRequestResultCodeInvalidAmount",
	-7: "CreatePreIssuanceRequestResultCodeInvalidReference",
}

var createPreIssuanceRequestResultCodeShortMap = map[int32]string{
	0:  "success",
	-1: "asset_not_found",
	-2: "reference_duplication",
	-3: "not_authorized_upload",
	-4: "invalid_signature",
	-5: "exceeded_max_amount",
	-6: "invalid_amount",
	-7: "invalid_reference",
}

var createPreIssuanceRequestResultCodeRevMap = map[string]int32{
	"CreatePreIssuanceRequestResultCodeSuccess":              0,
	"CreatePreIssuanceRequestResultCodeAssetNotFound":        -1,
	"CreatePreIssuanceRequestResultCodeReferenceDuplication": -2,
	"CreatePreIssuanceRequestResultCodeNotAuthorizedUpload":  -3,
	"CreatePreIssuanceRequestResultCodeInvalidSignature":     -4,
	"CreatePreIssuanceRequestResultCodeExceededMaxAmount":    -5,
	"CreatePreIssuanceRequestResultCodeInvalidAmount":        -6,
	"CreatePreIssuanceRequestResultCodeInvalidReference":     -7,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for CreatePreIssuanceRequestResultCode
func (e CreatePreIssuanceRequestResultCode) ValidEnum(v int32) bool {
	_, ok := createPreIssuanceRequestResultCodeMap[v]
	return ok
}
func (e CreatePreIssuanceRequestResultCode) isFlag() bool {
	for i := len(CreatePreIssuanceRequestResultCodeAll) - 1; i >= 0; i-- {
		expected := CreatePreIssuanceRequestResultCode(2) << uint64(len(CreatePreIssuanceRequestResultCodeAll)-1) >> uint64(len(CreatePreIssuanceRequestResultCodeAll)-i)
		if expected != CreatePreIssuanceRequestResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e CreatePreIssuanceRequestResultCode) String() string {
	name, _ := createPreIssuanceRequestResultCodeMap[int32(e)]
	return name
}

func (e CreatePreIssuanceRequestResultCode) ShortString() string {
	name, _ := createPreIssuanceRequestResultCodeShortMap[int32(e)]
	return name
}

func (e CreatePreIssuanceRequestResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range CreatePreIssuanceRequestResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *CreatePreIssuanceRequestResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = CreatePreIssuanceRequestResultCode(t.Value)
	return nil
}

// CreatePreIssuanceRequestResultSuccessExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//
type CreatePreIssuanceRequestResultSuccessExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreatePreIssuanceRequestResultSuccessExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreatePreIssuanceRequestResultSuccessExt
func (u CreatePreIssuanceRequestResultSuccessExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewCreatePreIssuanceRequestResultSuccessExt creates a new  CreatePreIssuanceRequestResultSuccessExt.
func NewCreatePreIssuanceRequestResultSuccessExt(v LedgerVersion, value interface{}) (result CreatePreIssuanceRequestResultSuccessExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// CreatePreIssuanceRequestResultSuccess is an XDR NestedStruct defines as:
//
//   struct {
//    		uint64 requestID;
//    		bool fulfilled;
//    		// reserved for future use
//    		union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//    		ext;
//    	}
//
type CreatePreIssuanceRequestResultSuccess struct {
	RequestId Uint64                                   `json:"requestID,omitempty"`
	Fulfilled bool                                     `json:"fulfilled,omitempty"`
	Ext       CreatePreIssuanceRequestResultSuccessExt `json:"ext,omitempty"`
}

// CreatePreIssuanceRequestResult is an XDR Union defines as:
//
//   union CreatePreIssuanceRequestResult switch (CreatePreIssuanceRequestResultCode code)
//    {
//    case SUCCESS:
//        struct {
//    		uint64 requestID;
//    		bool fulfilled;
//    		// reserved for future use
//    		union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//    		ext;
//    	} success;
//    default:
//        void;
//    };
//
type CreatePreIssuanceRequestResult struct {
	Code    CreatePreIssuanceRequestResultCode     `json:"code,omitempty"`
	Success *CreatePreIssuanceRequestResultSuccess `json:"success,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreatePreIssuanceRequestResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreatePreIssuanceRequestResult
func (u CreatePreIssuanceRequestResult) ArmForSwitch(sw int32) (string, bool) {
	switch CreatePreIssuanceRequestResultCode(sw) {
	case CreatePreIssuanceRequestResultCodeSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewCreatePreIssuanceRequestResult creates a new  CreatePreIssuanceRequestResult.
func NewCreatePreIssuanceRequestResult(code CreatePreIssuanceRequestResultCode, value interface{}) (result CreatePreIssuanceRequestResult, err error) {
	result.Code = code
	switch CreatePreIssuanceRequestResultCode(code) {
	case CreatePreIssuanceRequestResultCodeSuccess:
		tv, ok := value.(CreatePreIssuanceRequestResultSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreatePreIssuanceRequestResultSuccess")
			return
		}
		result.Success = &tv
	default:
		// void
	}
	return
}

// MustSuccess retrieves the Success value from the union,
// panicing if the value is not set.
func (u CreatePreIssuanceRequestResult) MustSuccess() CreatePreIssuanceRequestResultSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u CreatePreIssuanceRequestResult) GetSuccess() (result CreatePreIssuanceRequestResultSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// CreateSaleCreationRequestOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type CreateSaleCreationRequestOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreateSaleCreationRequestOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreateSaleCreationRequestOpExt
func (u CreateSaleCreationRequestOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewCreateSaleCreationRequestOpExt creates a new  CreateSaleCreationRequestOpExt.
func NewCreateSaleCreationRequestOpExt(v LedgerVersion, value interface{}) (result CreateSaleCreationRequestOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// CreateSaleCreationRequestOp is an XDR Struct defines as:
//
//   struct CreateSaleCreationRequestOp
//    {
//    	uint64 requestID;
//        SaleCreationRequest request;
//
//    	union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//
//    };
//
type CreateSaleCreationRequestOp struct {
	RequestId Uint64                         `json:"requestID,omitempty"`
	Request   SaleCreationRequest            `json:"request,omitempty"`
	Ext       CreateSaleCreationRequestOpExt `json:"ext,omitempty"`
}

// CreateSaleCreationRequestResultCode is an XDR Enum defines as:
//
//   enum CreateSaleCreationRequestResultCode
//    {
//        // codes considered as "success" for the operation
//        SUCCESS = 0,
//
//        // codes considered as "failure" for the operation
//    	REQUEST_NOT_FOUND = -1, // trying to update reviewable request which does not exists
//    	BASE_ASSET_OR_ASSET_REQUEST_NOT_FOUND = -2, // failed to find asset or asset request for sale
//    	QUOTE_ASSET_NOT_FOUND = -3, // failed to find quote asset
//    	START_END_INVALID = -4, // sale ends before start
//    	INVALID_END = -5, // end date is in the past
//    	INVALID_PRICE = -6, // price can not be 0
//    	INVALID_CAP = -7, // hard cap is < soft cap
//    	INSUFFICIENT_MAX_ISSUANCE = -8, // max number of tokens is less then number of tokens required for soft cap
//    	INVALID_ASSET_PAIR = -9, // one of the assets has invalid code or base asset is equal to quote asset
//    	REQUEST_OR_SALE_ALREADY_EXISTS = -10,
//    	INSUFFICIENT_PREISSUED = -11, // amount of pre issued tokens is insufficient for hard cap
//    	INVALID_DETAILS = -12, // details must be a valid json
//    	VERSION_IS_NOT_SUPPORTED_YET = -13 // version specified in request is not supported yet
//    };
//
type CreateSaleCreationRequestResultCode int32

const (
	CreateSaleCreationRequestResultCodeSuccess                         CreateSaleCreationRequestResultCode = 0
	CreateSaleCreationRequestResultCodeRequestNotFound                 CreateSaleCreationRequestResultCode = -1
	CreateSaleCreationRequestResultCodeBaseAssetOrAssetRequestNotFound CreateSaleCreationRequestResultCode = -2
	CreateSaleCreationRequestResultCodeQuoteAssetNotFound              CreateSaleCreationRequestResultCode = -3
	CreateSaleCreationRequestResultCodeStartEndInvalid                 CreateSaleCreationRequestResultCode = -4
	CreateSaleCreationRequestResultCodeInvalidEnd                      CreateSaleCreationRequestResultCode = -5
	CreateSaleCreationRequestResultCodeInvalidPrice                    CreateSaleCreationRequestResultCode = -6
	CreateSaleCreationRequestResultCodeInvalidCap                      CreateSaleCreationRequestResultCode = -7
	CreateSaleCreationRequestResultCodeInsufficientMaxIssuance         CreateSaleCreationRequestResultCode = -8
	CreateSaleCreationRequestResultCodeInvalidAssetPair                CreateSaleCreationRequestResultCode = -9
	CreateSaleCreationRequestResultCodeRequestOrSaleAlreadyExists      CreateSaleCreationRequestResultCode = -10
	CreateSaleCreationRequestResultCodeInsufficientPreissued           CreateSaleCreationRequestResultCode = -11
	CreateSaleCreationRequestResultCodeInvalidDetails                  CreateSaleCreationRequestResultCode = -12
	CreateSaleCreationRequestResultCodeVersionIsNotSupportedYet        CreateSaleCreationRequestResultCode = -13
)

var CreateSaleCreationRequestResultCodeAll = []CreateSaleCreationRequestResultCode{
	CreateSaleCreationRequestResultCodeSuccess,
	CreateSaleCreationRequestResultCodeRequestNotFound,
	CreateSaleCreationRequestResultCodeBaseAssetOrAssetRequestNotFound,
	CreateSaleCreationRequestResultCodeQuoteAssetNotFound,
	CreateSaleCreationRequestResultCodeStartEndInvalid,
	CreateSaleCreationRequestResultCodeInvalidEnd,
	CreateSaleCreationRequestResultCodeInvalidPrice,
	CreateSaleCreationRequestResultCodeInvalidCap,
	CreateSaleCreationRequestResultCodeInsufficientMaxIssuance,
	CreateSaleCreationRequestResultCodeInvalidAssetPair,
	CreateSaleCreationRequestResultCodeRequestOrSaleAlreadyExists,
	CreateSaleCreationRequestResultCodeInsufficientPreissued,
	CreateSaleCreationRequestResultCodeInvalidDetails,
	CreateSaleCreationRequestResultCodeVersionIsNotSupportedYet,
}

var createSaleCreationRequestResultCodeMap = map[int32]string{
	0:   "CreateSaleCreationRequestResultCodeSuccess",
	-1:  "CreateSaleCreationRequestResultCodeRequestNotFound",
	-2:  "CreateSaleCreationRequestResultCodeBaseAssetOrAssetRequestNotFound",
	-3:  "CreateSaleCreationRequestResultCodeQuoteAssetNotFound",
	-4:  "CreateSaleCreationRequestResultCodeStartEndInvalid",
	-5:  "CreateSaleCreationRequestResultCodeInvalidEnd",
	-6:  "CreateSaleCreationRequestResultCodeInvalidPrice",
	-7:  "CreateSaleCreationRequestResultCodeInvalidCap",
	-8:  "CreateSaleCreationRequestResultCodeInsufficientMaxIssuance",
	-9:  "CreateSaleCreationRequestResultCodeInvalidAssetPair",
	-10: "CreateSaleCreationRequestResultCodeRequestOrSaleAlreadyExists",
	-11: "CreateSaleCreationRequestResultCodeInsufficientPreissued",
	-12: "CreateSaleCreationRequestResultCodeInvalidDetails",
	-13: "CreateSaleCreationRequestResultCodeVersionIsNotSupportedYet",
}

var createSaleCreationRequestResultCodeShortMap = map[int32]string{
	0:   "success",
	-1:  "request_not_found",
	-2:  "base_asset_or_asset_request_not_found",
	-3:  "quote_asset_not_found",
	-4:  "start_end_invalid",
	-5:  "invalid_end",
	-6:  "invalid_price",
	-7:  "invalid_cap",
	-8:  "insufficient_max_issuance",
	-9:  "invalid_asset_pair",
	-10: "request_or_sale_already_exists",
	-11: "insufficient_preissued",
	-12: "invalid_details",
	-13: "version_is_not_supported_yet",
}

var createSaleCreationRequestResultCodeRevMap = map[string]int32{
	"CreateSaleCreationRequestResultCodeSuccess":                         0,
	"CreateSaleCreationRequestResultCodeRequestNotFound":                 -1,
	"CreateSaleCreationRequestResultCodeBaseAssetOrAssetRequestNotFound": -2,
	"CreateSaleCreationRequestResultCodeQuoteAssetNotFound":              -3,
	"CreateSaleCreationRequestResultCodeStartEndInvalid":                 -4,
	"CreateSaleCreationRequestResultCodeInvalidEnd":                      -5,
	"CreateSaleCreationRequestResultCodeInvalidPrice":                    -6,
	"CreateSaleCreationRequestResultCodeInvalidCap":                      -7,
	"CreateSaleCreationRequestResultCodeInsufficientMaxIssuance":         -8,
	"CreateSaleCreationRequestResultCodeInvalidAssetPair":                -9,
	"CreateSaleCreationRequestResultCodeRequestOrSaleAlreadyExists":      -10,
	"CreateSaleCreationRequestResultCodeInsufficientPreissued":           -11,
	"CreateSaleCreationRequestResultCodeInvalidDetails":                  -12,
	"CreateSaleCreationRequestResultCodeVersionIsNotSupportedYet":        -13,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for CreateSaleCreationRequestResultCode
func (e CreateSaleCreationRequestResultCode) ValidEnum(v int32) bool {
	_, ok := createSaleCreationRequestResultCodeMap[v]
	return ok
}
func (e CreateSaleCreationRequestResultCode) isFlag() bool {
	for i := len(CreateSaleCreationRequestResultCodeAll) - 1; i >= 0; i-- {
		expected := CreateSaleCreationRequestResultCode(2) << uint64(len(CreateSaleCreationRequestResultCodeAll)-1) >> uint64(len(CreateSaleCreationRequestResultCodeAll)-i)
		if expected != CreateSaleCreationRequestResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e CreateSaleCreationRequestResultCode) String() string {
	name, _ := createSaleCreationRequestResultCodeMap[int32(e)]
	return name
}

func (e CreateSaleCreationRequestResultCode) ShortString() string {
	name, _ := createSaleCreationRequestResultCodeShortMap[int32(e)]
	return name
}

func (e CreateSaleCreationRequestResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range CreateSaleCreationRequestResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *CreateSaleCreationRequestResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = CreateSaleCreationRequestResultCode(t.Value)
	return nil
}

// CreateSaleCreationSuccessExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type CreateSaleCreationSuccessExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreateSaleCreationSuccessExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreateSaleCreationSuccessExt
func (u CreateSaleCreationSuccessExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewCreateSaleCreationSuccessExt creates a new  CreateSaleCreationSuccessExt.
func NewCreateSaleCreationSuccessExt(v LedgerVersion, value interface{}) (result CreateSaleCreationSuccessExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// CreateSaleCreationSuccess is an XDR Struct defines as:
//
//   struct CreateSaleCreationSuccess {
//    	uint64 requestID;
//
//    	union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type CreateSaleCreationSuccess struct {
	RequestId Uint64                       `json:"requestID,omitempty"`
	Ext       CreateSaleCreationSuccessExt `json:"ext,omitempty"`
}

// CreateSaleCreationRequestResult is an XDR Union defines as:
//
//   union CreateSaleCreationRequestResult switch (CreateSaleCreationRequestResultCode code)
//    {
//        case SUCCESS:
//            CreateSaleCreationSuccess success;
//        default:
//            void;
//    };
//
type CreateSaleCreationRequestResult struct {
	Code    CreateSaleCreationRequestResultCode `json:"code,omitempty"`
	Success *CreateSaleCreationSuccess          `json:"success,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreateSaleCreationRequestResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreateSaleCreationRequestResult
func (u CreateSaleCreationRequestResult) ArmForSwitch(sw int32) (string, bool) {
	switch CreateSaleCreationRequestResultCode(sw) {
	case CreateSaleCreationRequestResultCodeSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewCreateSaleCreationRequestResult creates a new  CreateSaleCreationRequestResult.
func NewCreateSaleCreationRequestResult(code CreateSaleCreationRequestResultCode, value interface{}) (result CreateSaleCreationRequestResult, err error) {
	result.Code = code
	switch CreateSaleCreationRequestResultCode(code) {
	case CreateSaleCreationRequestResultCodeSuccess:
		tv, ok := value.(CreateSaleCreationSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateSaleCreationSuccess")
			return
		}
		result.Success = &tv
	default:
		// void
	}
	return
}

// MustSuccess retrieves the Success value from the union,
// panicing if the value is not set.
func (u CreateSaleCreationRequestResult) MustSuccess() CreateSaleCreationSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u CreateSaleCreationRequestResult) GetSuccess() (result CreateSaleCreationSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// CreateWithdrawalRequestOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type CreateWithdrawalRequestOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreateWithdrawalRequestOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreateWithdrawalRequestOpExt
func (u CreateWithdrawalRequestOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewCreateWithdrawalRequestOpExt creates a new  CreateWithdrawalRequestOpExt.
func NewCreateWithdrawalRequestOpExt(v LedgerVersion, value interface{}) (result CreateWithdrawalRequestOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// CreateWithdrawalRequestOp is an XDR Struct defines as:
//
//   struct CreateWithdrawalRequestOp
//    {
//        WithdrawalRequest request;
//
//    	union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//
//    };
//
type CreateWithdrawalRequestOp struct {
	Request WithdrawalRequest            `json:"request,omitempty"`
	Ext     CreateWithdrawalRequestOpExt `json:"ext,omitempty"`
}

// CreateWithdrawalRequestResultCode is an XDR Enum defines as:
//
//   enum CreateWithdrawalRequestResultCode
//    {
//        // codes considered as "success" for the operation
//        SUCCESS = 0,
//
//        // codes considered as "failure" for the operation
//    	INVALID_AMOUNT = -1, // amount is 0
//        INVALID_EXTERNAL_DETAILS = -2, // external details size exceeds max allowed
//    	BALANCE_NOT_FOUND = -3, // balance not found
//    	ASSET_IS_NOT_WITHDRAWABLE = -4, // asset is not withdrawable
//    	CONVERSION_PRICE_IS_NOT_AVAILABLE = -5, // failed to find conversion price - conversion is not allowed
//    	FEE_MISMATCHED = -6, // expected fee does not match calculated fee
//    	CONVERSION_OVERFLOW = -7, // overflow during converting source asset to dest asset
//    	CONVERTED_AMOUNT_MISMATCHED = -8, // expected converted amount passed by user, does not match calculated
//    	BALANCE_LOCK_OVERFLOW = -9, // overflow while tried to lock amount
//    	UNDERFUNDED = -10, // insufficient balance to perform operation
//    	INVALID_UNIVERSAL_AMOUNT = -11, // non-zero universal amount
//    	STATS_OVERFLOW = -12, // statistics overflowed by the operation
//        LIMITS_EXCEEDED = -13, // withdraw exceeds limits for source account
//    	INVALID_PRE_CONFIRMATION_DETAILS = -14 // it's not allowed to pass pre confirmation details
//    };
//
type CreateWithdrawalRequestResultCode int32

const (
	CreateWithdrawalRequestResultCodeSuccess                       CreateWithdrawalRequestResultCode = 0
	CreateWithdrawalRequestResultCodeInvalidAmount                 CreateWithdrawalRequestResultCode = -1
	CreateWithdrawalRequestResultCodeInvalidExternalDetails        CreateWithdrawalRequestResultCode = -2
	CreateWithdrawalRequestResultCodeBalanceNotFound               CreateWithdrawalRequestResultCode = -3
	CreateWithdrawalRequestResultCodeAssetIsNotWithdrawable        CreateWithdrawalRequestResultCode = -4
	CreateWithdrawalRequestResultCodeConversionPriceIsNotAvailable CreateWithdrawalRequestResultCode = -5
	CreateWithdrawalRequestResultCodeFeeMismatched                 CreateWithdrawalRequestResultCode = -6
	CreateWithdrawalRequestResultCodeConversionOverflow            CreateWithdrawalRequestResultCode = -7
	CreateWithdrawalRequestResultCodeConvertedAmountMismatched     CreateWithdrawalRequestResultCode = -8
	CreateWithdrawalRequestResultCodeBalanceLockOverflow           CreateWithdrawalRequestResultCode = -9
	CreateWithdrawalRequestResultCodeUnderfunded                   CreateWithdrawalRequestResultCode = -10
	CreateWithdrawalRequestResultCodeInvalidUniversalAmount        CreateWithdrawalRequestResultCode = -11
	CreateWithdrawalRequestResultCodeStatsOverflow                 CreateWithdrawalRequestResultCode = -12
	CreateWithdrawalRequestResultCodeLimitsExceeded                CreateWithdrawalRequestResultCode = -13
	CreateWithdrawalRequestResultCodeInvalidPreConfirmationDetails CreateWithdrawalRequestResultCode = -14
)

var CreateWithdrawalRequestResultCodeAll = []CreateWithdrawalRequestResultCode{
	CreateWithdrawalRequestResultCodeSuccess,
	CreateWithdrawalRequestResultCodeInvalidAmount,
	CreateWithdrawalRequestResultCodeInvalidExternalDetails,
	CreateWithdrawalRequestResultCodeBalanceNotFound,
	CreateWithdrawalRequestResultCodeAssetIsNotWithdrawable,
	CreateWithdrawalRequestResultCodeConversionPriceIsNotAvailable,
	CreateWithdrawalRequestResultCodeFeeMismatched,
	CreateWithdrawalRequestResultCodeConversionOverflow,
	CreateWithdrawalRequestResultCodeConvertedAmountMismatched,
	CreateWithdrawalRequestResultCodeBalanceLockOverflow,
	CreateWithdrawalRequestResultCodeUnderfunded,
	CreateWithdrawalRequestResultCodeInvalidUniversalAmount,
	CreateWithdrawalRequestResultCodeStatsOverflow,
	CreateWithdrawalRequestResultCodeLimitsExceeded,
	CreateWithdrawalRequestResultCodeInvalidPreConfirmationDetails,
}

var createWithdrawalRequestResultCodeMap = map[int32]string{
	0:   "CreateWithdrawalRequestResultCodeSuccess",
	-1:  "CreateWithdrawalRequestResultCodeInvalidAmount",
	-2:  "CreateWithdrawalRequestResultCodeInvalidExternalDetails",
	-3:  "CreateWithdrawalRequestResultCodeBalanceNotFound",
	-4:  "CreateWithdrawalRequestResultCodeAssetIsNotWithdrawable",
	-5:  "CreateWithdrawalRequestResultCodeConversionPriceIsNotAvailable",
	-6:  "CreateWithdrawalRequestResultCodeFeeMismatched",
	-7:  "CreateWithdrawalRequestResultCodeConversionOverflow",
	-8:  "CreateWithdrawalRequestResultCodeConvertedAmountMismatched",
	-9:  "CreateWithdrawalRequestResultCodeBalanceLockOverflow",
	-10: "CreateWithdrawalRequestResultCodeUnderfunded",
	-11: "CreateWithdrawalRequestResultCodeInvalidUniversalAmount",
	-12: "CreateWithdrawalRequestResultCodeStatsOverflow",
	-13: "CreateWithdrawalRequestResultCodeLimitsExceeded",
	-14: "CreateWithdrawalRequestResultCodeInvalidPreConfirmationDetails",
}

var createWithdrawalRequestResultCodeShortMap = map[int32]string{
	0:   "success",
	-1:  "invalid_amount",
	-2:  "invalid_external_details",
	-3:  "balance_not_found",
	-4:  "asset_is_not_withdrawable",
	-5:  "conversion_price_is_not_available",
	-6:  "fee_mismatched",
	-7:  "conversion_overflow",
	-8:  "converted_amount_mismatched",
	-9:  "balance_lock_overflow",
	-10: "underfunded",
	-11: "invalid_universal_amount",
	-12: "stats_overflow",
	-13: "limits_exceeded",
	-14: "invalid_pre_confirmation_details",
}

var createWithdrawalRequestResultCodeRevMap = map[string]int32{
	"CreateWithdrawalRequestResultCodeSuccess":                       0,
	"CreateWithdrawalRequestResultCodeInvalidAmount":                 -1,
	"CreateWithdrawalRequestResultCodeInvalidExternalDetails":        -2,
	"CreateWithdrawalRequestResultCodeBalanceNotFound":               -3,
	"CreateWithdrawalRequestResultCodeAssetIsNotWithdrawable":        -4,
	"CreateWithdrawalRequestResultCodeConversionPriceIsNotAvailable": -5,
	"CreateWithdrawalRequestResultCodeFeeMismatched":                 -6,
	"CreateWithdrawalRequestResultCodeConversionOverflow":            -7,
	"CreateWithdrawalRequestResultCodeConvertedAmountMismatched":     -8,
	"CreateWithdrawalRequestResultCodeBalanceLockOverflow":           -9,
	"CreateWithdrawalRequestResultCodeUnderfunded":                   -10,
	"CreateWithdrawalRequestResultCodeInvalidUniversalAmount":        -11,
	"CreateWithdrawalRequestResultCodeStatsOverflow":                 -12,
	"CreateWithdrawalRequestResultCodeLimitsExceeded":                -13,
	"CreateWithdrawalRequestResultCodeInvalidPreConfirmationDetails": -14,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for CreateWithdrawalRequestResultCode
func (e CreateWithdrawalRequestResultCode) ValidEnum(v int32) bool {
	_, ok := createWithdrawalRequestResultCodeMap[v]
	return ok
}
func (e CreateWithdrawalRequestResultCode) isFlag() bool {
	for i := len(CreateWithdrawalRequestResultCodeAll) - 1; i >= 0; i-- {
		expected := CreateWithdrawalRequestResultCode(2) << uint64(len(CreateWithdrawalRequestResultCodeAll)-1) >> uint64(len(CreateWithdrawalRequestResultCodeAll)-i)
		if expected != CreateWithdrawalRequestResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e CreateWithdrawalRequestResultCode) String() string {
	name, _ := createWithdrawalRequestResultCodeMap[int32(e)]
	return name
}

func (e CreateWithdrawalRequestResultCode) ShortString() string {
	name, _ := createWithdrawalRequestResultCodeShortMap[int32(e)]
	return name
}

func (e CreateWithdrawalRequestResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range CreateWithdrawalRequestResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *CreateWithdrawalRequestResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = CreateWithdrawalRequestResultCode(t.Value)
	return nil
}

// CreateWithdrawalSuccessExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type CreateWithdrawalSuccessExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreateWithdrawalSuccessExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreateWithdrawalSuccessExt
func (u CreateWithdrawalSuccessExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewCreateWithdrawalSuccessExt creates a new  CreateWithdrawalSuccessExt.
func NewCreateWithdrawalSuccessExt(v LedgerVersion, value interface{}) (result CreateWithdrawalSuccessExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// CreateWithdrawalSuccess is an XDR Struct defines as:
//
//   struct CreateWithdrawalSuccess {
//    	uint64 requestID;
//
//    	union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type CreateWithdrawalSuccess struct {
	RequestId Uint64                     `json:"requestID,omitempty"`
	Ext       CreateWithdrawalSuccessExt `json:"ext,omitempty"`
}

// CreateWithdrawalRequestResult is an XDR Union defines as:
//
//   union CreateWithdrawalRequestResult switch (CreateWithdrawalRequestResultCode code)
//    {
//        case SUCCESS:
//            CreateWithdrawalSuccess success;
//        default:
//            void;
//    };
//
type CreateWithdrawalRequestResult struct {
	Code    CreateWithdrawalRequestResultCode `json:"code,omitempty"`
	Success *CreateWithdrawalSuccess          `json:"success,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreateWithdrawalRequestResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreateWithdrawalRequestResult
func (u CreateWithdrawalRequestResult) ArmForSwitch(sw int32) (string, bool) {
	switch CreateWithdrawalRequestResultCode(sw) {
	case CreateWithdrawalRequestResultCodeSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewCreateWithdrawalRequestResult creates a new  CreateWithdrawalRequestResult.
func NewCreateWithdrawalRequestResult(code CreateWithdrawalRequestResultCode, value interface{}) (result CreateWithdrawalRequestResult, err error) {
	result.Code = code
	switch CreateWithdrawalRequestResultCode(code) {
	case CreateWithdrawalRequestResultCodeSuccess:
		tv, ok := value.(CreateWithdrawalSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateWithdrawalSuccess")
			return
		}
		result.Success = &tv
	default:
		// void
	}
	return
}

// MustSuccess retrieves the Success value from the union,
// panicing if the value is not set.
func (u CreateWithdrawalRequestResult) MustSuccess() CreateWithdrawalSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u CreateWithdrawalRequestResult) GetSuccess() (result CreateWithdrawalSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// DirectDebitOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type DirectDebitOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u DirectDebitOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of DirectDebitOpExt
func (u DirectDebitOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewDirectDebitOpExt creates a new  DirectDebitOpExt.
func NewDirectDebitOpExt(v LedgerVersion, value interface{}) (result DirectDebitOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// DirectDebitOp is an XDR Struct defines as:
//
//   struct DirectDebitOp
//    {
//        AccountID from;
//        PaymentOp paymentOp;
//    	 // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type DirectDebitOp struct {
	From      AccountId        `json:"from,omitempty"`
	PaymentOp PaymentOp        `json:"paymentOp,omitempty"`
	Ext       DirectDebitOpExt `json:"ext,omitempty"`
}

// DirectDebitResultCode is an XDR Enum defines as:
//
//   enum DirectDebitResultCode
//    {
//        // codes considered as "success" for the operation
//        SUCCESS = 0, // payment successfuly completed
//
//        // codes considered as "failure" for the operation
//        MALFORMED = -1,       // bad input
//        UNDERFUNDED = -2,     // not enough funds in source account
//        LINE_FULL = -3,       // destination would go above their limit
//    	FEE_MISMATCHED = -4,   // fee is not equal to expected fee
//        BALANCE_NOT_FOUND = -5, // destination balance not found
//        BALANCE_ACCOUNT_MISMATCHED = -6,
//        BALANCE_ASSETS_MISMATCHED = -7,
//    	SRC_BALANCE_NOT_FOUND = -8, // source balance not found
//        REFERENCE_DUPLICATION = -9,
//        STATS_OVERFLOW = -10,
//        LIMITS_EXCEEDED = -11,
//        NOT_ALLOWED_BY_ASSET_POLICY = -12,
//        NO_TRUST = -13
//    };
//
type DirectDebitResultCode int32

const (
	DirectDebitResultCodeSuccess                  DirectDebitResultCode = 0
	DirectDebitResultCodeMalformed                DirectDebitResultCode = -1
	DirectDebitResultCodeUnderfunded              DirectDebitResultCode = -2
	DirectDebitResultCodeLineFull                 DirectDebitResultCode = -3
	DirectDebitResultCodeFeeMismatched            DirectDebitResultCode = -4
	DirectDebitResultCodeBalanceNotFound          DirectDebitResultCode = -5
	DirectDebitResultCodeBalanceAccountMismatched DirectDebitResultCode = -6
	DirectDebitResultCodeBalanceAssetsMismatched  DirectDebitResultCode = -7
	DirectDebitResultCodeSrcBalanceNotFound       DirectDebitResultCode = -8
	DirectDebitResultCodeReferenceDuplication     DirectDebitResultCode = -9
	DirectDebitResultCodeStatsOverflow            DirectDebitResultCode = -10
	DirectDebitResultCodeLimitsExceeded           DirectDebitResultCode = -11
	DirectDebitResultCodeNotAllowedByAssetPolicy  DirectDebitResultCode = -12
	DirectDebitResultCodeNoTrust                  DirectDebitResultCode = -13
)

var DirectDebitResultCodeAll = []DirectDebitResultCode{
	DirectDebitResultCodeSuccess,
	DirectDebitResultCodeMalformed,
	DirectDebitResultCodeUnderfunded,
	DirectDebitResultCodeLineFull,
	DirectDebitResultCodeFeeMismatched,
	DirectDebitResultCodeBalanceNotFound,
	DirectDebitResultCodeBalanceAccountMismatched,
	DirectDebitResultCodeBalanceAssetsMismatched,
	DirectDebitResultCodeSrcBalanceNotFound,
	DirectDebitResultCodeReferenceDuplication,
	DirectDebitResultCodeStatsOverflow,
	DirectDebitResultCodeLimitsExceeded,
	DirectDebitResultCodeNotAllowedByAssetPolicy,
	DirectDebitResultCodeNoTrust,
}

var directDebitResultCodeMap = map[int32]string{
	0:   "DirectDebitResultCodeSuccess",
	-1:  "DirectDebitResultCodeMalformed",
	-2:  "DirectDebitResultCodeUnderfunded",
	-3:  "DirectDebitResultCodeLineFull",
	-4:  "DirectDebitResultCodeFeeMismatched",
	-5:  "DirectDebitResultCodeBalanceNotFound",
	-6:  "DirectDebitResultCodeBalanceAccountMismatched",
	-7:  "DirectDebitResultCodeBalanceAssetsMismatched",
	-8:  "DirectDebitResultCodeSrcBalanceNotFound",
	-9:  "DirectDebitResultCodeReferenceDuplication",
	-10: "DirectDebitResultCodeStatsOverflow",
	-11: "DirectDebitResultCodeLimitsExceeded",
	-12: "DirectDebitResultCodeNotAllowedByAssetPolicy",
	-13: "DirectDebitResultCodeNoTrust",
}

var directDebitResultCodeShortMap = map[int32]string{
	0:   "success",
	-1:  "malformed",
	-2:  "underfunded",
	-3:  "line_full",
	-4:  "fee_mismatched",
	-5:  "balance_not_found",
	-6:  "balance_account_mismatched",
	-7:  "balance_assets_mismatched",
	-8:  "src_balance_not_found",
	-9:  "reference_duplication",
	-10: "stats_overflow",
	-11: "limits_exceeded",
	-12: "not_allowed_by_asset_policy",
	-13: "no_trust",
}

var directDebitResultCodeRevMap = map[string]int32{
	"DirectDebitResultCodeSuccess":                  0,
	"DirectDebitResultCodeMalformed":                -1,
	"DirectDebitResultCodeUnderfunded":              -2,
	"DirectDebitResultCodeLineFull":                 -3,
	"DirectDebitResultCodeFeeMismatched":            -4,
	"DirectDebitResultCodeBalanceNotFound":          -5,
	"DirectDebitResultCodeBalanceAccountMismatched": -6,
	"DirectDebitResultCodeBalanceAssetsMismatched":  -7,
	"DirectDebitResultCodeSrcBalanceNotFound":       -8,
	"DirectDebitResultCodeReferenceDuplication":     -9,
	"DirectDebitResultCodeStatsOverflow":            -10,
	"DirectDebitResultCodeLimitsExceeded":           -11,
	"DirectDebitResultCodeNotAllowedByAssetPolicy":  -12,
	"DirectDebitResultCodeNoTrust":                  -13,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for DirectDebitResultCode
func (e DirectDebitResultCode) ValidEnum(v int32) bool {
	_, ok := directDebitResultCodeMap[v]
	return ok
}
func (e DirectDebitResultCode) isFlag() bool {
	for i := len(DirectDebitResultCodeAll) - 1; i >= 0; i-- {
		expected := DirectDebitResultCode(2) << uint64(len(DirectDebitResultCodeAll)-1) >> uint64(len(DirectDebitResultCodeAll)-i)
		if expected != DirectDebitResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e DirectDebitResultCode) String() string {
	name, _ := directDebitResultCodeMap[int32(e)]
	return name
}

func (e DirectDebitResultCode) ShortString() string {
	name, _ := directDebitResultCodeShortMap[int32(e)]
	return name
}

func (e DirectDebitResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range DirectDebitResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *DirectDebitResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = DirectDebitResultCode(t.Value)
	return nil
}

// DirectDebitSuccessExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type DirectDebitSuccessExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u DirectDebitSuccessExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of DirectDebitSuccessExt
func (u DirectDebitSuccessExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewDirectDebitSuccessExt creates a new  DirectDebitSuccessExt.
func NewDirectDebitSuccessExt(v LedgerVersion, value interface{}) (result DirectDebitSuccessExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// DirectDebitSuccess is an XDR Struct defines as:
//
//   struct DirectDebitSuccess {
//    	PaymentResponse paymentResponse;
//    	// reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type DirectDebitSuccess struct {
	PaymentResponse PaymentResponse       `json:"paymentResponse,omitempty"`
	Ext             DirectDebitSuccessExt `json:"ext,omitempty"`
}

// DirectDebitResult is an XDR Union defines as:
//
//   union DirectDebitResult switch (DirectDebitResultCode code)
//    {
//    case SUCCESS:
//        DirectDebitSuccess success;
//    default:
//        void;
//    };
//
type DirectDebitResult struct {
	Code    DirectDebitResultCode `json:"code,omitempty"`
	Success *DirectDebitSuccess   `json:"success,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u DirectDebitResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of DirectDebitResult
func (u DirectDebitResult) ArmForSwitch(sw int32) (string, bool) {
	switch DirectDebitResultCode(sw) {
	case DirectDebitResultCodeSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewDirectDebitResult creates a new  DirectDebitResult.
func NewDirectDebitResult(code DirectDebitResultCode, value interface{}) (result DirectDebitResult, err error) {
	result.Code = code
	switch DirectDebitResultCode(code) {
	case DirectDebitResultCodeSuccess:
		tv, ok := value.(DirectDebitSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be DirectDebitSuccess")
			return
		}
		result.Success = &tv
	default:
		// void
	}
	return
}

// MustSuccess retrieves the Success value from the union,
// panicing if the value is not set.
func (u DirectDebitResult) MustSuccess() DirectDebitSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u DirectDebitResult) GetSuccess() (result DirectDebitSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// ManageAccountOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type ManageAccountOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageAccountOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageAccountOpExt
func (u ManageAccountOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewManageAccountOpExt creates a new  ManageAccountOpExt.
func NewManageAccountOpExt(v LedgerVersion, value interface{}) (result ManageAccountOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ManageAccountOp is an XDR Struct defines as:
//
//   struct ManageAccountOp
//    {
//        AccountID account; // account to manage
//        AccountType accountType;
//        uint32 blockReasonsToAdd;
//        uint32 blockReasonsToRemove;
//    	 // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type ManageAccountOp struct {
	Account              AccountId          `json:"account,omitempty"`
	AccountType          AccountType        `json:"accountType,omitempty"`
	BlockReasonsToAdd    Uint32             `json:"blockReasonsToAdd,omitempty"`
	BlockReasonsToRemove Uint32             `json:"blockReasonsToRemove,omitempty"`
	Ext                  ManageAccountOpExt `json:"ext,omitempty"`
}

// ManageAccountResultCode is an XDR Enum defines as:
//
//   enum ManageAccountResultCode
//    {
//        // codes considered as "success" for the operation
//        SUCCESS = 0, // account was created
//
//        // codes considered as "failure" for the operation
//        NOT_FOUND = -1,         // account does not exists
//        MALFORMED = -2,
//    	NOT_ALLOWED = -3,         // manage account operation is not allowed on this account
//        TYPE_MISMATCH = -4
//    };
//
type ManageAccountResultCode int32

const (
	ManageAccountResultCodeSuccess      ManageAccountResultCode = 0
	ManageAccountResultCodeNotFound     ManageAccountResultCode = -1
	ManageAccountResultCodeMalformed    ManageAccountResultCode = -2
	ManageAccountResultCodeNotAllowed   ManageAccountResultCode = -3
	ManageAccountResultCodeTypeMismatch ManageAccountResultCode = -4
)

var ManageAccountResultCodeAll = []ManageAccountResultCode{
	ManageAccountResultCodeSuccess,
	ManageAccountResultCodeNotFound,
	ManageAccountResultCodeMalformed,
	ManageAccountResultCodeNotAllowed,
	ManageAccountResultCodeTypeMismatch,
}

var manageAccountResultCodeMap = map[int32]string{
	0:  "ManageAccountResultCodeSuccess",
	-1: "ManageAccountResultCodeNotFound",
	-2: "ManageAccountResultCodeMalformed",
	-3: "ManageAccountResultCodeNotAllowed",
	-4: "ManageAccountResultCodeTypeMismatch",
}

var manageAccountResultCodeShortMap = map[int32]string{
	0:  "success",
	-1: "not_found",
	-2: "malformed",
	-3: "not_allowed",
	-4: "type_mismatch",
}

var manageAccountResultCodeRevMap = map[string]int32{
	"ManageAccountResultCodeSuccess":      0,
	"ManageAccountResultCodeNotFound":     -1,
	"ManageAccountResultCodeMalformed":    -2,
	"ManageAccountResultCodeNotAllowed":   -3,
	"ManageAccountResultCodeTypeMismatch": -4,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ManageAccountResultCode
func (e ManageAccountResultCode) ValidEnum(v int32) bool {
	_, ok := manageAccountResultCodeMap[v]
	return ok
}
func (e ManageAccountResultCode) isFlag() bool {
	for i := len(ManageAccountResultCodeAll) - 1; i >= 0; i-- {
		expected := ManageAccountResultCode(2) << uint64(len(ManageAccountResultCodeAll)-1) >> uint64(len(ManageAccountResultCodeAll)-i)
		if expected != ManageAccountResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ManageAccountResultCode) String() string {
	name, _ := manageAccountResultCodeMap[int32(e)]
	return name
}

func (e ManageAccountResultCode) ShortString() string {
	name, _ := manageAccountResultCodeShortMap[int32(e)]
	return name
}

func (e ManageAccountResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range ManageAccountResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ManageAccountResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ManageAccountResultCode(t.Value)
	return nil
}

// ManageAccountSuccessExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type ManageAccountSuccessExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageAccountSuccessExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageAccountSuccessExt
func (u ManageAccountSuccessExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewManageAccountSuccessExt creates a new  ManageAccountSuccessExt.
func NewManageAccountSuccessExt(v LedgerVersion, value interface{}) (result ManageAccountSuccessExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ManageAccountSuccess is an XDR Struct defines as:
//
//   struct ManageAccountSuccess {
//    	uint32 blockReasons;
//     // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type ManageAccountSuccess struct {
	BlockReasons Uint32                  `json:"blockReasons,omitempty"`
	Ext          ManageAccountSuccessExt `json:"ext,omitempty"`
}

// ManageAccountResult is an XDR Union defines as:
//
//   union ManageAccountResult switch (ManageAccountResultCode code)
//    {
//    case SUCCESS:
//        ManageAccountSuccess success;
//    default:
//        void;
//    };
//
type ManageAccountResult struct {
	Code    ManageAccountResultCode `json:"code,omitempty"`
	Success *ManageAccountSuccess   `json:"success,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageAccountResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageAccountResult
func (u ManageAccountResult) ArmForSwitch(sw int32) (string, bool) {
	switch ManageAccountResultCode(sw) {
	case ManageAccountResultCodeSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewManageAccountResult creates a new  ManageAccountResult.
func NewManageAccountResult(code ManageAccountResultCode, value interface{}) (result ManageAccountResult, err error) {
	result.Code = code
	switch ManageAccountResultCode(code) {
	case ManageAccountResultCodeSuccess:
		tv, ok := value.(ManageAccountSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageAccountSuccess")
			return
		}
		result.Success = &tv
	default:
		// void
	}
	return
}

// MustSuccess retrieves the Success value from the union,
// panicing if the value is not set.
func (u ManageAccountResult) MustSuccess() ManageAccountSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageAccountResult) GetSuccess() (result ManageAccountSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// ManageAssetPairAction is an XDR Enum defines as:
//
//   enum ManageAssetPairAction
//    {
//        CREATE = 0,
//        UPDATE_PRICE = 1,
//        UPDATE_POLICIES = 2
//    };
//
type ManageAssetPairAction int32

const (
	ManageAssetPairActionCreate         ManageAssetPairAction = 0
	ManageAssetPairActionUpdatePrice    ManageAssetPairAction = 1
	ManageAssetPairActionUpdatePolicies ManageAssetPairAction = 2
)

var ManageAssetPairActionAll = []ManageAssetPairAction{
	ManageAssetPairActionCreate,
	ManageAssetPairActionUpdatePrice,
	ManageAssetPairActionUpdatePolicies,
}

var manageAssetPairActionMap = map[int32]string{
	0: "ManageAssetPairActionCreate",
	1: "ManageAssetPairActionUpdatePrice",
	2: "ManageAssetPairActionUpdatePolicies",
}

var manageAssetPairActionShortMap = map[int32]string{
	0: "create",
	1: "update_price",
	2: "update_policies",
}

var manageAssetPairActionRevMap = map[string]int32{
	"ManageAssetPairActionCreate":         0,
	"ManageAssetPairActionUpdatePrice":    1,
	"ManageAssetPairActionUpdatePolicies": 2,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ManageAssetPairAction
func (e ManageAssetPairAction) ValidEnum(v int32) bool {
	_, ok := manageAssetPairActionMap[v]
	return ok
}
func (e ManageAssetPairAction) isFlag() bool {
	for i := len(ManageAssetPairActionAll) - 1; i >= 0; i-- {
		expected := ManageAssetPairAction(2) << uint64(len(ManageAssetPairActionAll)-1) >> uint64(len(ManageAssetPairActionAll)-i)
		if expected != ManageAssetPairActionAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ManageAssetPairAction) String() string {
	name, _ := manageAssetPairActionMap[int32(e)]
	return name
}

func (e ManageAssetPairAction) ShortString() string {
	name, _ := manageAssetPairActionShortMap[int32(e)]
	return name
}

func (e ManageAssetPairAction) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range ManageAssetPairActionAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ManageAssetPairAction) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ManageAssetPairAction(t.Value)
	return nil
}

// ManageAssetPairOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type ManageAssetPairOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageAssetPairOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageAssetPairOpExt
func (u ManageAssetPairOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewManageAssetPairOpExt creates a new  ManageAssetPairOpExt.
func NewManageAssetPairOpExt(v LedgerVersion, value interface{}) (result ManageAssetPairOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ManageAssetPairOp is an XDR Struct defines as:
//
//   struct ManageAssetPairOp
//    {
//        ManageAssetPairAction action;
//    	AssetCode base;
//    	AssetCode quote;
//
//        int64 physicalPrice;
//
//    	int64 physicalPriceCorrection; // correction of physical price in percents. If physical price is set and restriction by physical price set, mininal price for offer for this pair will be physicalPrice * physicalPriceCorrection
//    	int64 maxPriceStep;
//
//    	int32 policies;
//
//    	 // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type ManageAssetPairOp struct {
	Action                  ManageAssetPairAction `json:"action,omitempty"`
	Base                    AssetCode             `json:"base,omitempty"`
	Quote                   AssetCode             `json:"quote,omitempty"`
	PhysicalPrice           Int64                 `json:"physicalPrice,omitempty"`
	PhysicalPriceCorrection Int64                 `json:"physicalPriceCorrection,omitempty"`
	MaxPriceStep            Int64                 `json:"maxPriceStep,omitempty"`
	Policies                Int32                 `json:"policies,omitempty"`
	Ext                     ManageAssetPairOpExt  `json:"ext,omitempty"`
}

// ManageAssetPairResultCode is an XDR Enum defines as:
//
//   enum ManageAssetPairResultCode
//    {
//        // codes considered as "success" for the operation
//        SUCCESS = 0,
//
//        // codes considered as "failure" for the operation
//    	NOT_FOUND = -1,           // failed to find asset with such code
//    	ALREADY_EXISTS = -2,
//        MALFORMED = -3,
//    	INVALID_ASSET = -4,
//    	INVALID_ACTION = -5,
//    	INVALID_POLICIES = -6,
//    	ASSET_NOT_FOUND = -7
//    };
//
type ManageAssetPairResultCode int32

const (
	ManageAssetPairResultCodeSuccess         ManageAssetPairResultCode = 0
	ManageAssetPairResultCodeNotFound        ManageAssetPairResultCode = -1
	ManageAssetPairResultCodeAlreadyExists   ManageAssetPairResultCode = -2
	ManageAssetPairResultCodeMalformed       ManageAssetPairResultCode = -3
	ManageAssetPairResultCodeInvalidAsset    ManageAssetPairResultCode = -4
	ManageAssetPairResultCodeInvalidAction   ManageAssetPairResultCode = -5
	ManageAssetPairResultCodeInvalidPolicies ManageAssetPairResultCode = -6
	ManageAssetPairResultCodeAssetNotFound   ManageAssetPairResultCode = -7
)

var ManageAssetPairResultCodeAll = []ManageAssetPairResultCode{
	ManageAssetPairResultCodeSuccess,
	ManageAssetPairResultCodeNotFound,
	ManageAssetPairResultCodeAlreadyExists,
	ManageAssetPairResultCodeMalformed,
	ManageAssetPairResultCodeInvalidAsset,
	ManageAssetPairResultCodeInvalidAction,
	ManageAssetPairResultCodeInvalidPolicies,
	ManageAssetPairResultCodeAssetNotFound,
}

var manageAssetPairResultCodeMap = map[int32]string{
	0:  "ManageAssetPairResultCodeSuccess",
	-1: "ManageAssetPairResultCodeNotFound",
	-2: "ManageAssetPairResultCodeAlreadyExists",
	-3: "ManageAssetPairResultCodeMalformed",
	-4: "ManageAssetPairResultCodeInvalidAsset",
	-5: "ManageAssetPairResultCodeInvalidAction",
	-6: "ManageAssetPairResultCodeInvalidPolicies",
	-7: "ManageAssetPairResultCodeAssetNotFound",
}

var manageAssetPairResultCodeShortMap = map[int32]string{
	0:  "success",
	-1: "not_found",
	-2: "already_exists",
	-3: "malformed",
	-4: "invalid_asset",
	-5: "invalid_action",
	-6: "invalid_policies",
	-7: "asset_not_found",
}

var manageAssetPairResultCodeRevMap = map[string]int32{
	"ManageAssetPairResultCodeSuccess":         0,
	"ManageAssetPairResultCodeNotFound":        -1,
	"ManageAssetPairResultCodeAlreadyExists":   -2,
	"ManageAssetPairResultCodeMalformed":       -3,
	"ManageAssetPairResultCodeInvalidAsset":    -4,
	"ManageAssetPairResultCodeInvalidAction":   -5,
	"ManageAssetPairResultCodeInvalidPolicies": -6,
	"ManageAssetPairResultCodeAssetNotFound":   -7,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ManageAssetPairResultCode
func (e ManageAssetPairResultCode) ValidEnum(v int32) bool {
	_, ok := manageAssetPairResultCodeMap[v]
	return ok
}
func (e ManageAssetPairResultCode) isFlag() bool {
	for i := len(ManageAssetPairResultCodeAll) - 1; i >= 0; i-- {
		expected := ManageAssetPairResultCode(2) << uint64(len(ManageAssetPairResultCodeAll)-1) >> uint64(len(ManageAssetPairResultCodeAll)-i)
		if expected != ManageAssetPairResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ManageAssetPairResultCode) String() string {
	name, _ := manageAssetPairResultCodeMap[int32(e)]
	return name
}

func (e ManageAssetPairResultCode) ShortString() string {
	name, _ := manageAssetPairResultCodeShortMap[int32(e)]
	return name
}

func (e ManageAssetPairResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range ManageAssetPairResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ManageAssetPairResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ManageAssetPairResultCode(t.Value)
	return nil
}

// ManageAssetPairSuccessExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type ManageAssetPairSuccessExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageAssetPairSuccessExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageAssetPairSuccessExt
func (u ManageAssetPairSuccessExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewManageAssetPairSuccessExt creates a new  ManageAssetPairSuccessExt.
func NewManageAssetPairSuccessExt(v LedgerVersion, value interface{}) (result ManageAssetPairSuccessExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ManageAssetPairSuccess is an XDR Struct defines as:
//
//   struct ManageAssetPairSuccess
//    {
//    	int64 currentPrice;
//    	// reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type ManageAssetPairSuccess struct {
	CurrentPrice Int64                     `json:"currentPrice,omitempty"`
	Ext          ManageAssetPairSuccessExt `json:"ext,omitempty"`
}

// ManageAssetPairResult is an XDR Union defines as:
//
//   union ManageAssetPairResult switch (ManageAssetPairResultCode code)
//    {
//    case SUCCESS:
//        ManageAssetPairSuccess success;
//    default:
//        void;
//    };
//
type ManageAssetPairResult struct {
	Code    ManageAssetPairResultCode `json:"code,omitempty"`
	Success *ManageAssetPairSuccess   `json:"success,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageAssetPairResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageAssetPairResult
func (u ManageAssetPairResult) ArmForSwitch(sw int32) (string, bool) {
	switch ManageAssetPairResultCode(sw) {
	case ManageAssetPairResultCodeSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewManageAssetPairResult creates a new  ManageAssetPairResult.
func NewManageAssetPairResult(code ManageAssetPairResultCode, value interface{}) (result ManageAssetPairResult, err error) {
	result.Code = code
	switch ManageAssetPairResultCode(code) {
	case ManageAssetPairResultCodeSuccess:
		tv, ok := value.(ManageAssetPairSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageAssetPairSuccess")
			return
		}
		result.Success = &tv
	default:
		// void
	}
	return
}

// MustSuccess retrieves the Success value from the union,
// panicing if the value is not set.
func (u ManageAssetPairResult) MustSuccess() ManageAssetPairSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageAssetPairResult) GetSuccess() (result ManageAssetPairSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// ManageAssetAction is an XDR Enum defines as:
//
//   enum ManageAssetAction
//    {
//        CREATE_ASSET_CREATION_REQUEST = 0,
//        CREATE_ASSET_UPDATE_REQUEST = 1,
//    	CANCEL_ASSET_REQUEST = 2,
//    	CHANGE_PREISSUED_ASSET_SIGNER = 3,
//    	UPDATE_MAX_ISSUANCE = 4
//    };
//
type ManageAssetAction int32

const (
	ManageAssetActionCreateAssetCreationRequest ManageAssetAction = 0
	ManageAssetActionCreateAssetUpdateRequest   ManageAssetAction = 1
	ManageAssetActionCancelAssetRequest         ManageAssetAction = 2
	ManageAssetActionChangePreissuedAssetSigner ManageAssetAction = 3
	ManageAssetActionUpdateMaxIssuance          ManageAssetAction = 4
)

var ManageAssetActionAll = []ManageAssetAction{
	ManageAssetActionCreateAssetCreationRequest,
	ManageAssetActionCreateAssetUpdateRequest,
	ManageAssetActionCancelAssetRequest,
	ManageAssetActionChangePreissuedAssetSigner,
	ManageAssetActionUpdateMaxIssuance,
}

var manageAssetActionMap = map[int32]string{
	0: "ManageAssetActionCreateAssetCreationRequest",
	1: "ManageAssetActionCreateAssetUpdateRequest",
	2: "ManageAssetActionCancelAssetRequest",
	3: "ManageAssetActionChangePreissuedAssetSigner",
	4: "ManageAssetActionUpdateMaxIssuance",
}

var manageAssetActionShortMap = map[int32]string{
	0: "create_asset_creation_request",
	1: "create_asset_update_request",
	2: "cancel_asset_request",
	3: "change_preissued_asset_signer",
	4: "update_max_issuance",
}

var manageAssetActionRevMap = map[string]int32{
	"ManageAssetActionCreateAssetCreationRequest": 0,
	"ManageAssetActionCreateAssetUpdateRequest":   1,
	"ManageAssetActionCancelAssetRequest":         2,
	"ManageAssetActionChangePreissuedAssetSigner": 3,
	"ManageAssetActionUpdateMaxIssuance":          4,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ManageAssetAction
func (e ManageAssetAction) ValidEnum(v int32) bool {
	_, ok := manageAssetActionMap[v]
	return ok
}
func (e ManageAssetAction) isFlag() bool {
	for i := len(ManageAssetActionAll) - 1; i >= 0; i-- {
		expected := ManageAssetAction(2) << uint64(len(ManageAssetActionAll)-1) >> uint64(len(ManageAssetActionAll)-i)
		if expected != ManageAssetActionAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ManageAssetAction) String() string {
	name, _ := manageAssetActionMap[int32(e)]
	return name
}

func (e ManageAssetAction) ShortString() string {
	name, _ := manageAssetActionShortMap[int32(e)]
	return name
}

func (e ManageAssetAction) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range ManageAssetActionAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ManageAssetAction) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ManageAssetAction(t.Value)
	return nil
}

// CancelAssetRequestExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type CancelAssetRequestExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CancelAssetRequestExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CancelAssetRequestExt
func (u CancelAssetRequestExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewCancelAssetRequestExt creates a new  CancelAssetRequestExt.
func NewCancelAssetRequestExt(v LedgerVersion, value interface{}) (result CancelAssetRequestExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// CancelAssetRequest is an XDR Struct defines as:
//
//   struct CancelAssetRequest {
//
//    	// reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type CancelAssetRequest struct {
	Ext CancelAssetRequestExt `json:"ext,omitempty"`
}

// UpdateMaxIssuanceExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type UpdateMaxIssuanceExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u UpdateMaxIssuanceExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of UpdateMaxIssuanceExt
func (u UpdateMaxIssuanceExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewUpdateMaxIssuanceExt creates a new  UpdateMaxIssuanceExt.
func NewUpdateMaxIssuanceExt(v LedgerVersion, value interface{}) (result UpdateMaxIssuanceExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// UpdateMaxIssuance is an XDR Struct defines as:
//
//   struct UpdateMaxIssuance {
//
//    	AssetCode assetCode;
//    	uint64 maxIssuanceAmount;
//    	// reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type UpdateMaxIssuance struct {
	AssetCode         AssetCode            `json:"assetCode,omitempty"`
	MaxIssuanceAmount Uint64               `json:"maxIssuanceAmount,omitempty"`
	Ext               UpdateMaxIssuanceExt `json:"ext,omitempty"`
}

// ManageAssetOpRequest is an XDR NestedUnion defines as:
//
//   union switch (ManageAssetAction action)
//    	{
//    	case CREATE_ASSET_CREATION_REQUEST:
//    		AssetCreationRequest createAsset;
//    	case CREATE_ASSET_UPDATE_REQUEST:
//    		AssetUpdateRequest updateAsset;
//    	case CANCEL_ASSET_REQUEST:
//    		CancelAssetRequest cancelRequest;
//    	case CHANGE_PREISSUED_ASSET_SIGNER:
//    		AssetChangePreissuedSigner changePreissuedSigner;
//        case UPDATE_MAX_ISSUANCE:
//            UpdateMaxIssuance updateMaxIssuance;
//    	}
//
type ManageAssetOpRequest struct {
	Action                ManageAssetAction           `json:"action,omitempty"`
	CreateAsset           *AssetCreationRequest       `json:"createAsset,omitempty"`
	UpdateAsset           *AssetUpdateRequest         `json:"updateAsset,omitempty"`
	CancelRequest         *CancelAssetRequest         `json:"cancelRequest,omitempty"`
	ChangePreissuedSigner *AssetChangePreissuedSigner `json:"changePreissuedSigner,omitempty"`
	UpdateMaxIssuance     *UpdateMaxIssuance          `json:"updateMaxIssuance,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageAssetOpRequest) SwitchFieldName() string {
	return "Action"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageAssetOpRequest
func (u ManageAssetOpRequest) ArmForSwitch(sw int32) (string, bool) {
	switch ManageAssetAction(sw) {
	case ManageAssetActionCreateAssetCreationRequest:
		return "CreateAsset", true
	case ManageAssetActionCreateAssetUpdateRequest:
		return "UpdateAsset", true
	case ManageAssetActionCancelAssetRequest:
		return "CancelRequest", true
	case ManageAssetActionChangePreissuedAssetSigner:
		return "ChangePreissuedSigner", true
	case ManageAssetActionUpdateMaxIssuance:
		return "UpdateMaxIssuance", true
	}
	return "-", false
}

// NewManageAssetOpRequest creates a new  ManageAssetOpRequest.
func NewManageAssetOpRequest(action ManageAssetAction, value interface{}) (result ManageAssetOpRequest, err error) {
	result.Action = action
	switch ManageAssetAction(action) {
	case ManageAssetActionCreateAssetCreationRequest:
		tv, ok := value.(AssetCreationRequest)
		if !ok {
			err = fmt.Errorf("invalid value, must be AssetCreationRequest")
			return
		}
		result.CreateAsset = &tv
	case ManageAssetActionCreateAssetUpdateRequest:
		tv, ok := value.(AssetUpdateRequest)
		if !ok {
			err = fmt.Errorf("invalid value, must be AssetUpdateRequest")
			return
		}
		result.UpdateAsset = &tv
	case ManageAssetActionCancelAssetRequest:
		tv, ok := value.(CancelAssetRequest)
		if !ok {
			err = fmt.Errorf("invalid value, must be CancelAssetRequest")
			return
		}
		result.CancelRequest = &tv
	case ManageAssetActionChangePreissuedAssetSigner:
		tv, ok := value.(AssetChangePreissuedSigner)
		if !ok {
			err = fmt.Errorf("invalid value, must be AssetChangePreissuedSigner")
			return
		}
		result.ChangePreissuedSigner = &tv
	case ManageAssetActionUpdateMaxIssuance:
		tv, ok := value.(UpdateMaxIssuance)
		if !ok {
			err = fmt.Errorf("invalid value, must be UpdateMaxIssuance")
			return
		}
		result.UpdateMaxIssuance = &tv
	}
	return
}

// MustCreateAsset retrieves the CreateAsset value from the union,
// panicing if the value is not set.
func (u ManageAssetOpRequest) MustCreateAsset() AssetCreationRequest {
	val, ok := u.GetCreateAsset()

	if !ok {
		panic("arm CreateAsset is not set")
	}

	return val
}

// GetCreateAsset retrieves the CreateAsset value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageAssetOpRequest) GetCreateAsset() (result AssetCreationRequest, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Action))

	if armName == "CreateAsset" {
		result = *u.CreateAsset
		ok = true
	}

	return
}

// MustUpdateAsset retrieves the UpdateAsset value from the union,
// panicing if the value is not set.
func (u ManageAssetOpRequest) MustUpdateAsset() AssetUpdateRequest {
	val, ok := u.GetUpdateAsset()

	if !ok {
		panic("arm UpdateAsset is not set")
	}

	return val
}

// GetUpdateAsset retrieves the UpdateAsset value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageAssetOpRequest) GetUpdateAsset() (result AssetUpdateRequest, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Action))

	if armName == "UpdateAsset" {
		result = *u.UpdateAsset
		ok = true
	}

	return
}

// MustCancelRequest retrieves the CancelRequest value from the union,
// panicing if the value is not set.
func (u ManageAssetOpRequest) MustCancelRequest() CancelAssetRequest {
	val, ok := u.GetCancelRequest()

	if !ok {
		panic("arm CancelRequest is not set")
	}

	return val
}

// GetCancelRequest retrieves the CancelRequest value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageAssetOpRequest) GetCancelRequest() (result CancelAssetRequest, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Action))

	if armName == "CancelRequest" {
		result = *u.CancelRequest
		ok = true
	}

	return
}

// MustChangePreissuedSigner retrieves the ChangePreissuedSigner value from the union,
// panicing if the value is not set.
func (u ManageAssetOpRequest) MustChangePreissuedSigner() AssetChangePreissuedSigner {
	val, ok := u.GetChangePreissuedSigner()

	if !ok {
		panic("arm ChangePreissuedSigner is not set")
	}

	return val
}

// GetChangePreissuedSigner retrieves the ChangePreissuedSigner value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageAssetOpRequest) GetChangePreissuedSigner() (result AssetChangePreissuedSigner, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Action))

	if armName == "ChangePreissuedSigner" {
		result = *u.ChangePreissuedSigner
		ok = true
	}

	return
}

// MustUpdateMaxIssuance retrieves the UpdateMaxIssuance value from the union,
// panicing if the value is not set.
func (u ManageAssetOpRequest) MustUpdateMaxIssuance() UpdateMaxIssuance {
	val, ok := u.GetUpdateMaxIssuance()

	if !ok {
		panic("arm UpdateMaxIssuance is not set")
	}

	return val
}

// GetUpdateMaxIssuance retrieves the UpdateMaxIssuance value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageAssetOpRequest) GetUpdateMaxIssuance() (result UpdateMaxIssuance, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Action))

	if armName == "UpdateMaxIssuance" {
		result = *u.UpdateMaxIssuance
		ok = true
	}

	return
}

// ManageAssetOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type ManageAssetOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageAssetOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageAssetOpExt
func (u ManageAssetOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewManageAssetOpExt creates a new  ManageAssetOpExt.
func NewManageAssetOpExt(v LedgerVersion, value interface{}) (result ManageAssetOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ManageAssetOp is an XDR Struct defines as:
//
//   struct ManageAssetOp
//    {
//    	uint64 requestID; // 0 to create, non zero will try to update
//        union switch (ManageAssetAction action)
//    	{
//    	case CREATE_ASSET_CREATION_REQUEST:
//    		AssetCreationRequest createAsset;
//    	case CREATE_ASSET_UPDATE_REQUEST:
//    		AssetUpdateRequest updateAsset;
//    	case CANCEL_ASSET_REQUEST:
//    		CancelAssetRequest cancelRequest;
//    	case CHANGE_PREISSUED_ASSET_SIGNER:
//    		AssetChangePreissuedSigner changePreissuedSigner;
//        case UPDATE_MAX_ISSUANCE:
//            UpdateMaxIssuance updateMaxIssuance;
//    	} request;
//
//    	// reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type ManageAssetOp struct {
	RequestId Uint64               `json:"requestID,omitempty"`
	Request   ManageAssetOpRequest `json:"request,omitempty"`
	Ext       ManageAssetOpExt     `json:"ext,omitempty"`
}

// ManageAssetResultCode is an XDR Enum defines as:
//
//   enum ManageAssetResultCode
//    {
//        // codes considered as "success" for the operation
//        SUCCESS = 0,                       // request was successfully created/updated/canceled
//
//        // codes considered as "failure" for the operation
//    	REQUEST_NOT_FOUND = -1,           // failed to find asset request with such id
//    	ASSET_ALREADY_EXISTS = -3,			   // asset with such code already exist
//        INVALID_MAX_ISSUANCE_AMOUNT = -4, // max issuance amount is 0
//    	INVALID_CODE = -5,                // asset code is invalid (empty or contains space)
//    	INVALID_POLICIES = -7,            // asset policies (has flag which does not belong to AssetPolicies enum)
//    	ASSET_NOT_FOUND = -8,             // asset does not exists
//    	REQUEST_ALREADY_EXISTS = -9,      // request for creation of unique entry already exists
//    	STATS_ASSET_ALREADY_EXISTS = -10, // statistics quote asset already exists
//    	INITIAL_PREISSUED_EXCEEDS_MAX_ISSUANCE = -11, // initial pre issued amount exceeds max issuance amount
//    	INVALID_DETAILS = -12 // details must be a valid json
//    };
//
type ManageAssetResultCode int32

const (
	ManageAssetResultCodeSuccess                            ManageAssetResultCode = 0
	ManageAssetResultCodeRequestNotFound                    ManageAssetResultCode = -1
	ManageAssetResultCodeAssetAlreadyExists                 ManageAssetResultCode = -3
	ManageAssetResultCodeInvalidMaxIssuanceAmount           ManageAssetResultCode = -4
	ManageAssetResultCodeInvalidCode                        ManageAssetResultCode = -5
	ManageAssetResultCodeInvalidPolicies                    ManageAssetResultCode = -7
	ManageAssetResultCodeAssetNotFound                      ManageAssetResultCode = -8
	ManageAssetResultCodeRequestAlreadyExists               ManageAssetResultCode = -9
	ManageAssetResultCodeStatsAssetAlreadyExists            ManageAssetResultCode = -10
	ManageAssetResultCodeInitialPreissuedExceedsMaxIssuance ManageAssetResultCode = -11
	ManageAssetResultCodeInvalidDetails                     ManageAssetResultCode = -12
)

var ManageAssetResultCodeAll = []ManageAssetResultCode{
	ManageAssetResultCodeSuccess,
	ManageAssetResultCodeRequestNotFound,
	ManageAssetResultCodeAssetAlreadyExists,
	ManageAssetResultCodeInvalidMaxIssuanceAmount,
	ManageAssetResultCodeInvalidCode,
	ManageAssetResultCodeInvalidPolicies,
	ManageAssetResultCodeAssetNotFound,
	ManageAssetResultCodeRequestAlreadyExists,
	ManageAssetResultCodeStatsAssetAlreadyExists,
	ManageAssetResultCodeInitialPreissuedExceedsMaxIssuance,
	ManageAssetResultCodeInvalidDetails,
}

var manageAssetResultCodeMap = map[int32]string{
	0:   "ManageAssetResultCodeSuccess",
	-1:  "ManageAssetResultCodeRequestNotFound",
	-3:  "ManageAssetResultCodeAssetAlreadyExists",
	-4:  "ManageAssetResultCodeInvalidMaxIssuanceAmount",
	-5:  "ManageAssetResultCodeInvalidCode",
	-7:  "ManageAssetResultCodeInvalidPolicies",
	-8:  "ManageAssetResultCodeAssetNotFound",
	-9:  "ManageAssetResultCodeRequestAlreadyExists",
	-10: "ManageAssetResultCodeStatsAssetAlreadyExists",
	-11: "ManageAssetResultCodeInitialPreissuedExceedsMaxIssuance",
	-12: "ManageAssetResultCodeInvalidDetails",
}

var manageAssetResultCodeShortMap = map[int32]string{
	0:   "success",
	-1:  "request_not_found",
	-3:  "asset_already_exists",
	-4:  "invalid_max_issuance_amount",
	-5:  "invalid_code",
	-7:  "invalid_policies",
	-8:  "asset_not_found",
	-9:  "request_already_exists",
	-10: "stats_asset_already_exists",
	-11: "initial_preissued_exceeds_max_issuance",
	-12: "invalid_details",
}

var manageAssetResultCodeRevMap = map[string]int32{
	"ManageAssetResultCodeSuccess":                            0,
	"ManageAssetResultCodeRequestNotFound":                    -1,
	"ManageAssetResultCodeAssetAlreadyExists":                 -3,
	"ManageAssetResultCodeInvalidMaxIssuanceAmount":           -4,
	"ManageAssetResultCodeInvalidCode":                        -5,
	"ManageAssetResultCodeInvalidPolicies":                    -7,
	"ManageAssetResultCodeAssetNotFound":                      -8,
	"ManageAssetResultCodeRequestAlreadyExists":               -9,
	"ManageAssetResultCodeStatsAssetAlreadyExists":            -10,
	"ManageAssetResultCodeInitialPreissuedExceedsMaxIssuance": -11,
	"ManageAssetResultCodeInvalidDetails":                     -12,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ManageAssetResultCode
func (e ManageAssetResultCode) ValidEnum(v int32) bool {
	_, ok := manageAssetResultCodeMap[v]
	return ok
}
func (e ManageAssetResultCode) isFlag() bool {
	for i := len(ManageAssetResultCodeAll) - 1; i >= 0; i-- {
		expected := ManageAssetResultCode(2) << uint64(len(ManageAssetResultCodeAll)-1) >> uint64(len(ManageAssetResultCodeAll)-i)
		if expected != ManageAssetResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ManageAssetResultCode) String() string {
	name, _ := manageAssetResultCodeMap[int32(e)]
	return name
}

func (e ManageAssetResultCode) ShortString() string {
	name, _ := manageAssetResultCodeShortMap[int32(e)]
	return name
}

func (e ManageAssetResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range ManageAssetResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ManageAssetResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ManageAssetResultCode(t.Value)
	return nil
}

// ManageAssetSuccessExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type ManageAssetSuccessExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageAssetSuccessExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageAssetSuccessExt
func (u ManageAssetSuccessExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewManageAssetSuccessExt creates a new  ManageAssetSuccessExt.
func NewManageAssetSuccessExt(v LedgerVersion, value interface{}) (result ManageAssetSuccessExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ManageAssetSuccess is an XDR Struct defines as:
//
//   struct ManageAssetSuccess
//    {
//    	uint64 requestID;
//    	bool fulfilled;
//        // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type ManageAssetSuccess struct {
	RequestId Uint64                `json:"requestID,omitempty"`
	Fulfilled bool                  `json:"fulfilled,omitempty"`
	Ext       ManageAssetSuccessExt `json:"ext,omitempty"`
}

// ManageAssetResult is an XDR Union defines as:
//
//   union ManageAssetResult switch (ManageAssetResultCode code)
//    {
//    case SUCCESS:
//        ManageAssetSuccess success;
//    default:
//        void;
//    };
//
type ManageAssetResult struct {
	Code    ManageAssetResultCode `json:"code,omitempty"`
	Success *ManageAssetSuccess   `json:"success,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageAssetResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageAssetResult
func (u ManageAssetResult) ArmForSwitch(sw int32) (string, bool) {
	switch ManageAssetResultCode(sw) {
	case ManageAssetResultCodeSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewManageAssetResult creates a new  ManageAssetResult.
func NewManageAssetResult(code ManageAssetResultCode, value interface{}) (result ManageAssetResult, err error) {
	result.Code = code
	switch ManageAssetResultCode(code) {
	case ManageAssetResultCodeSuccess:
		tv, ok := value.(ManageAssetSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageAssetSuccess")
			return
		}
		result.Success = &tv
	default:
		// void
	}
	return
}

// MustSuccess retrieves the Success value from the union,
// panicing if the value is not set.
func (u ManageAssetResult) MustSuccess() ManageAssetSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageAssetResult) GetSuccess() (result ManageAssetSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// ManageBalanceAction is an XDR Enum defines as:
//
//   enum ManageBalanceAction
//    {
//        CREATE = 0,
//        DELETE_BALANCE = 1,
//    	CREATE_UNIQUE = 2 // ensures that balance will not be created if one for such asset and account exists
//    };
//
type ManageBalanceAction int32

const (
	ManageBalanceActionCreate        ManageBalanceAction = 0
	ManageBalanceActionDeleteBalance ManageBalanceAction = 1
	ManageBalanceActionCreateUnique  ManageBalanceAction = 2
)

var ManageBalanceActionAll = []ManageBalanceAction{
	ManageBalanceActionCreate,
	ManageBalanceActionDeleteBalance,
	ManageBalanceActionCreateUnique,
}

var manageBalanceActionMap = map[int32]string{
	0: "ManageBalanceActionCreate",
	1: "ManageBalanceActionDeleteBalance",
	2: "ManageBalanceActionCreateUnique",
}

var manageBalanceActionShortMap = map[int32]string{
	0: "create",
	1: "delete_balance",
	2: "create_unique",
}

var manageBalanceActionRevMap = map[string]int32{
	"ManageBalanceActionCreate":        0,
	"ManageBalanceActionDeleteBalance": 1,
	"ManageBalanceActionCreateUnique":  2,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ManageBalanceAction
func (e ManageBalanceAction) ValidEnum(v int32) bool {
	_, ok := manageBalanceActionMap[v]
	return ok
}
func (e ManageBalanceAction) isFlag() bool {
	for i := len(ManageBalanceActionAll) - 1; i >= 0; i-- {
		expected := ManageBalanceAction(2) << uint64(len(ManageBalanceActionAll)-1) >> uint64(len(ManageBalanceActionAll)-i)
		if expected != ManageBalanceActionAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ManageBalanceAction) String() string {
	name, _ := manageBalanceActionMap[int32(e)]
	return name
}

func (e ManageBalanceAction) ShortString() string {
	name, _ := manageBalanceActionShortMap[int32(e)]
	return name
}

func (e ManageBalanceAction) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range ManageBalanceActionAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ManageBalanceAction) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ManageBalanceAction(t.Value)
	return nil
}

// ManageBalanceOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type ManageBalanceOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageBalanceOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageBalanceOpExt
func (u ManageBalanceOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewManageBalanceOpExt creates a new  ManageBalanceOpExt.
func NewManageBalanceOpExt(v LedgerVersion, value interface{}) (result ManageBalanceOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ManageBalanceOp is an XDR Struct defines as:
//
//   struct ManageBalanceOp
//    {
//        ManageBalanceAction action;
//        AccountID destination;
//        AssetCode asset;
//    	union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type ManageBalanceOp struct {
	Action      ManageBalanceAction `json:"action,omitempty"`
	Destination AccountId           `json:"destination,omitempty"`
	Asset       AssetCode           `json:"asset,omitempty"`
	Ext         ManageBalanceOpExt  `json:"ext,omitempty"`
}

// ManageBalanceResultCode is an XDR Enum defines as:
//
//   enum ManageBalanceResultCode
//    {
//        // codes considered as "success" for the operation
//        SUCCESS = 0,
//
//        // codes considered as "failure" for the operation
//        MALFORMED = -1,       // invalid destination
//        NOT_FOUND = -2,
//        DESTINATION_NOT_FOUND = -3,
//        ASSET_NOT_FOUND = -4,
//        INVALID_ASSET = -5,
//    	BALANCE_ALREADY_EXISTS = -6,
//    	VERSION_IS_NOT_SUPPORTED_YET = -7 // version specified in request is not supported yet
//    };
//
type ManageBalanceResultCode int32

const (
	ManageBalanceResultCodeSuccess                  ManageBalanceResultCode = 0
	ManageBalanceResultCodeMalformed                ManageBalanceResultCode = -1
	ManageBalanceResultCodeNotFound                 ManageBalanceResultCode = -2
	ManageBalanceResultCodeDestinationNotFound      ManageBalanceResultCode = -3
	ManageBalanceResultCodeAssetNotFound            ManageBalanceResultCode = -4
	ManageBalanceResultCodeInvalidAsset             ManageBalanceResultCode = -5
	ManageBalanceResultCodeBalanceAlreadyExists     ManageBalanceResultCode = -6
	ManageBalanceResultCodeVersionIsNotSupportedYet ManageBalanceResultCode = -7
)

var ManageBalanceResultCodeAll = []ManageBalanceResultCode{
	ManageBalanceResultCodeSuccess,
	ManageBalanceResultCodeMalformed,
	ManageBalanceResultCodeNotFound,
	ManageBalanceResultCodeDestinationNotFound,
	ManageBalanceResultCodeAssetNotFound,
	ManageBalanceResultCodeInvalidAsset,
	ManageBalanceResultCodeBalanceAlreadyExists,
	ManageBalanceResultCodeVersionIsNotSupportedYet,
}

var manageBalanceResultCodeMap = map[int32]string{
	0:  "ManageBalanceResultCodeSuccess",
	-1: "ManageBalanceResultCodeMalformed",
	-2: "ManageBalanceResultCodeNotFound",
	-3: "ManageBalanceResultCodeDestinationNotFound",
	-4: "ManageBalanceResultCodeAssetNotFound",
	-5: "ManageBalanceResultCodeInvalidAsset",
	-6: "ManageBalanceResultCodeBalanceAlreadyExists",
	-7: "ManageBalanceResultCodeVersionIsNotSupportedYet",
}

var manageBalanceResultCodeShortMap = map[int32]string{
	0:  "success",
	-1: "malformed",
	-2: "not_found",
	-3: "destination_not_found",
	-4: "asset_not_found",
	-5: "invalid_asset",
	-6: "balance_already_exists",
	-7: "version_is_not_supported_yet",
}

var manageBalanceResultCodeRevMap = map[string]int32{
	"ManageBalanceResultCodeSuccess":                  0,
	"ManageBalanceResultCodeMalformed":                -1,
	"ManageBalanceResultCodeNotFound":                 -2,
	"ManageBalanceResultCodeDestinationNotFound":      -3,
	"ManageBalanceResultCodeAssetNotFound":            -4,
	"ManageBalanceResultCodeInvalidAsset":             -5,
	"ManageBalanceResultCodeBalanceAlreadyExists":     -6,
	"ManageBalanceResultCodeVersionIsNotSupportedYet": -7,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ManageBalanceResultCode
func (e ManageBalanceResultCode) ValidEnum(v int32) bool {
	_, ok := manageBalanceResultCodeMap[v]
	return ok
}
func (e ManageBalanceResultCode) isFlag() bool {
	for i := len(ManageBalanceResultCodeAll) - 1; i >= 0; i-- {
		expected := ManageBalanceResultCode(2) << uint64(len(ManageBalanceResultCodeAll)-1) >> uint64(len(ManageBalanceResultCodeAll)-i)
		if expected != ManageBalanceResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ManageBalanceResultCode) String() string {
	name, _ := manageBalanceResultCodeMap[int32(e)]
	return name
}

func (e ManageBalanceResultCode) ShortString() string {
	name, _ := manageBalanceResultCodeShortMap[int32(e)]
	return name
}

func (e ManageBalanceResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range ManageBalanceResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ManageBalanceResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ManageBalanceResultCode(t.Value)
	return nil
}

// ManageBalanceSuccessExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type ManageBalanceSuccessExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageBalanceSuccessExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageBalanceSuccessExt
func (u ManageBalanceSuccessExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewManageBalanceSuccessExt creates a new  ManageBalanceSuccessExt.
func NewManageBalanceSuccessExt(v LedgerVersion, value interface{}) (result ManageBalanceSuccessExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ManageBalanceSuccess is an XDR Struct defines as:
//
//   struct ManageBalanceSuccess {
//    	BalanceID balanceID;
//    	// reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type ManageBalanceSuccess struct {
	BalanceId BalanceId               `json:"balanceID,omitempty"`
	Ext       ManageBalanceSuccessExt `json:"ext,omitempty"`
}

// ManageBalanceResult is an XDR Union defines as:
//
//   union ManageBalanceResult switch (ManageBalanceResultCode code)
//    {
//    case SUCCESS:
//        ManageBalanceSuccess success;
//    default:
//        void;
//    };
//
type ManageBalanceResult struct {
	Code    ManageBalanceResultCode `json:"code,omitempty"`
	Success *ManageBalanceSuccess   `json:"success,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageBalanceResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageBalanceResult
func (u ManageBalanceResult) ArmForSwitch(sw int32) (string, bool) {
	switch ManageBalanceResultCode(sw) {
	case ManageBalanceResultCodeSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewManageBalanceResult creates a new  ManageBalanceResult.
func NewManageBalanceResult(code ManageBalanceResultCode, value interface{}) (result ManageBalanceResult, err error) {
	result.Code = code
	switch ManageBalanceResultCode(code) {
	case ManageBalanceResultCodeSuccess:
		tv, ok := value.(ManageBalanceSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageBalanceSuccess")
			return
		}
		result.Success = &tv
	default:
		// void
	}
	return
}

// MustSuccess retrieves the Success value from the union,
// panicing if the value is not set.
func (u ManageBalanceResult) MustSuccess() ManageBalanceSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageBalanceResult) GetSuccess() (result ManageBalanceSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// ManageContractRequestAction is an XDR Enum defines as:
//
//   enum ManageContractRequestAction
//    {
//        CREATE = 0,
//        REMOVE = 1
//    };
//
type ManageContractRequestAction int32

const (
	ManageContractRequestActionCreate ManageContractRequestAction = 0
	ManageContractRequestActionRemove ManageContractRequestAction = 1
)

var ManageContractRequestActionAll = []ManageContractRequestAction{
	ManageContractRequestActionCreate,
	ManageContractRequestActionRemove,
}

var manageContractRequestActionMap = map[int32]string{
	0: "ManageContractRequestActionCreate",
	1: "ManageContractRequestActionRemove",
}

var manageContractRequestActionShortMap = map[int32]string{
	0: "create",
	1: "remove",
}

var manageContractRequestActionRevMap = map[string]int32{
	"ManageContractRequestActionCreate": 0,
	"ManageContractRequestActionRemove": 1,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ManageContractRequestAction
func (e ManageContractRequestAction) ValidEnum(v int32) bool {
	_, ok := manageContractRequestActionMap[v]
	return ok
}
func (e ManageContractRequestAction) isFlag() bool {
	for i := len(ManageContractRequestActionAll) - 1; i >= 0; i-- {
		expected := ManageContractRequestAction(2) << uint64(len(ManageContractRequestActionAll)-1) >> uint64(len(ManageContractRequestActionAll)-i)
		if expected != ManageContractRequestActionAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ManageContractRequestAction) String() string {
	name, _ := manageContractRequestActionMap[int32(e)]
	return name
}

func (e ManageContractRequestAction) ShortString() string {
	name, _ := manageContractRequestActionShortMap[int32(e)]
	return name
}

func (e ManageContractRequestAction) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range ManageContractRequestActionAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ManageContractRequestAction) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ManageContractRequestAction(t.Value)
	return nil
}

// ManageContractRequestOpDetails is an XDR NestedUnion defines as:
//
//   union switch (ManageContractRequestAction action){
//        case CREATE:
//            ContractRequest contractRequest;
//        case REMOVE:
//            uint64 requestID;
//        }
//
type ManageContractRequestOpDetails struct {
	Action          ManageContractRequestAction `json:"action,omitempty"`
	ContractRequest *ContractRequest            `json:"contractRequest,omitempty"`
	RequestId       *Uint64                     `json:"requestID,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageContractRequestOpDetails) SwitchFieldName() string {
	return "Action"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageContractRequestOpDetails
func (u ManageContractRequestOpDetails) ArmForSwitch(sw int32) (string, bool) {
	switch ManageContractRequestAction(sw) {
	case ManageContractRequestActionCreate:
		return "ContractRequest", true
	case ManageContractRequestActionRemove:
		return "RequestId", true
	}
	return "-", false
}

// NewManageContractRequestOpDetails creates a new  ManageContractRequestOpDetails.
func NewManageContractRequestOpDetails(action ManageContractRequestAction, value interface{}) (result ManageContractRequestOpDetails, err error) {
	result.Action = action
	switch ManageContractRequestAction(action) {
	case ManageContractRequestActionCreate:
		tv, ok := value.(ContractRequest)
		if !ok {
			err = fmt.Errorf("invalid value, must be ContractRequest")
			return
		}
		result.ContractRequest = &tv
	case ManageContractRequestActionRemove:
		tv, ok := value.(Uint64)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint64")
			return
		}
		result.RequestId = &tv
	}
	return
}

// MustContractRequest retrieves the ContractRequest value from the union,
// panicing if the value is not set.
func (u ManageContractRequestOpDetails) MustContractRequest() ContractRequest {
	val, ok := u.GetContractRequest()

	if !ok {
		panic("arm ContractRequest is not set")
	}

	return val
}

// GetContractRequest retrieves the ContractRequest value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageContractRequestOpDetails) GetContractRequest() (result ContractRequest, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Action))

	if armName == "ContractRequest" {
		result = *u.ContractRequest
		ok = true
	}

	return
}

// MustRequestId retrieves the RequestId value from the union,
// panicing if the value is not set.
func (u ManageContractRequestOpDetails) MustRequestId() Uint64 {
	val, ok := u.GetRequestId()

	if !ok {
		panic("arm RequestId is not set")
	}

	return val
}

// GetRequestId retrieves the RequestId value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageContractRequestOpDetails) GetRequestId() (result Uint64, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Action))

	if armName == "RequestId" {
		result = *u.RequestId
		ok = true
	}

	return
}

// ManageContractRequestOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type ManageContractRequestOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageContractRequestOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageContractRequestOpExt
func (u ManageContractRequestOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewManageContractRequestOpExt creates a new  ManageContractRequestOpExt.
func NewManageContractRequestOpExt(v LedgerVersion, value interface{}) (result ManageContractRequestOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ManageContractRequestOp is an XDR Struct defines as:
//
//   struct ManageContractRequestOp
//    {
//        union switch (ManageContractRequestAction action){
//        case CREATE:
//            ContractRequest contractRequest;
//        case REMOVE:
//            uint64 requestID;
//        } details;
//
//    	// reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type ManageContractRequestOp struct {
	Details ManageContractRequestOpDetails `json:"details,omitempty"`
	Ext     ManageContractRequestOpExt     `json:"ext,omitempty"`
}

// ManageContractRequestResultCode is an XDR Enum defines as:
//
//   enum ManageContractRequestResultCode
//    {
//        // codes considered as "success" for the operation
//        SUCCESS = 0,
//
//        // codes considered as "failure" for the operation
//        MALFORMED = -1,
//        NOT_FOUND = -2, // not found contract request, when try to remove
//        TOO_MANY_CONTRACTS = -3,
//        NOT_ALLOWED_TO_REMOVE = -4, // only contract creator can remove contract
//        DETAILS_TOO_LONG = -5
//    };
//
type ManageContractRequestResultCode int32

const (
	ManageContractRequestResultCodeSuccess            ManageContractRequestResultCode = 0
	ManageContractRequestResultCodeMalformed          ManageContractRequestResultCode = -1
	ManageContractRequestResultCodeNotFound           ManageContractRequestResultCode = -2
	ManageContractRequestResultCodeTooManyContracts   ManageContractRequestResultCode = -3
	ManageContractRequestResultCodeNotAllowedToRemove ManageContractRequestResultCode = -4
	ManageContractRequestResultCodeDetailsTooLong     ManageContractRequestResultCode = -5
)

var ManageContractRequestResultCodeAll = []ManageContractRequestResultCode{
	ManageContractRequestResultCodeSuccess,
	ManageContractRequestResultCodeMalformed,
	ManageContractRequestResultCodeNotFound,
	ManageContractRequestResultCodeTooManyContracts,
	ManageContractRequestResultCodeNotAllowedToRemove,
	ManageContractRequestResultCodeDetailsTooLong,
}

var manageContractRequestResultCodeMap = map[int32]string{
	0:  "ManageContractRequestResultCodeSuccess",
	-1: "ManageContractRequestResultCodeMalformed",
	-2: "ManageContractRequestResultCodeNotFound",
	-3: "ManageContractRequestResultCodeTooManyContracts",
	-4: "ManageContractRequestResultCodeNotAllowedToRemove",
	-5: "ManageContractRequestResultCodeDetailsTooLong",
}

var manageContractRequestResultCodeShortMap = map[int32]string{
	0:  "success",
	-1: "malformed",
	-2: "not_found",
	-3: "too_many_contracts",
	-4: "not_allowed_to_remove",
	-5: "details_too_long",
}

var manageContractRequestResultCodeRevMap = map[string]int32{
	"ManageContractRequestResultCodeSuccess":            0,
	"ManageContractRequestResultCodeMalformed":          -1,
	"ManageContractRequestResultCodeNotFound":           -2,
	"ManageContractRequestResultCodeTooManyContracts":   -3,
	"ManageContractRequestResultCodeNotAllowedToRemove": -4,
	"ManageContractRequestResultCodeDetailsTooLong":     -5,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ManageContractRequestResultCode
func (e ManageContractRequestResultCode) ValidEnum(v int32) bool {
	_, ok := manageContractRequestResultCodeMap[v]
	return ok
}
func (e ManageContractRequestResultCode) isFlag() bool {
	for i := len(ManageContractRequestResultCodeAll) - 1; i >= 0; i-- {
		expected := ManageContractRequestResultCode(2) << uint64(len(ManageContractRequestResultCodeAll)-1) >> uint64(len(ManageContractRequestResultCodeAll)-i)
		if expected != ManageContractRequestResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ManageContractRequestResultCode) String() string {
	name, _ := manageContractRequestResultCodeMap[int32(e)]
	return name
}

func (e ManageContractRequestResultCode) ShortString() string {
	name, _ := manageContractRequestResultCodeShortMap[int32(e)]
	return name
}

func (e ManageContractRequestResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range ManageContractRequestResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ManageContractRequestResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ManageContractRequestResultCode(t.Value)
	return nil
}

// CreateContractRequestResponseExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type CreateContractRequestResponseExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreateContractRequestResponseExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreateContractRequestResponseExt
func (u CreateContractRequestResponseExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewCreateContractRequestResponseExt creates a new  CreateContractRequestResponseExt.
func NewCreateContractRequestResponseExt(v LedgerVersion, value interface{}) (result CreateContractRequestResponseExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// CreateContractRequestResponse is an XDR Struct defines as:
//
//   struct CreateContractRequestResponse
//    {
//    	uint64 requestID;
//
//    	union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type CreateContractRequestResponse struct {
	RequestId Uint64                           `json:"requestID,omitempty"`
	Ext       CreateContractRequestResponseExt `json:"ext,omitempty"`
}

// ManageContractRequestResultSuccessDetails is an XDR NestedUnion defines as:
//
//   union switch (ManageContractRequestAction action)
//            {
//            case CREATE:
//                CreateContractRequestResponse response;
//            case REMOVE:
//                void;
//            }
//
type ManageContractRequestResultSuccessDetails struct {
	Action   ManageContractRequestAction    `json:"action,omitempty"`
	Response *CreateContractRequestResponse `json:"response,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageContractRequestResultSuccessDetails) SwitchFieldName() string {
	return "Action"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageContractRequestResultSuccessDetails
func (u ManageContractRequestResultSuccessDetails) ArmForSwitch(sw int32) (string, bool) {
	switch ManageContractRequestAction(sw) {
	case ManageContractRequestActionCreate:
		return "Response", true
	case ManageContractRequestActionRemove:
		return "", true
	}
	return "-", false
}

// NewManageContractRequestResultSuccessDetails creates a new  ManageContractRequestResultSuccessDetails.
func NewManageContractRequestResultSuccessDetails(action ManageContractRequestAction, value interface{}) (result ManageContractRequestResultSuccessDetails, err error) {
	result.Action = action
	switch ManageContractRequestAction(action) {
	case ManageContractRequestActionCreate:
		tv, ok := value.(CreateContractRequestResponse)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateContractRequestResponse")
			return
		}
		result.Response = &tv
	case ManageContractRequestActionRemove:
		// void
	}
	return
}

// MustResponse retrieves the Response value from the union,
// panicing if the value is not set.
func (u ManageContractRequestResultSuccessDetails) MustResponse() CreateContractRequestResponse {
	val, ok := u.GetResponse()

	if !ok {
		panic("arm Response is not set")
	}

	return val
}

// GetResponse retrieves the Response value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageContractRequestResultSuccessDetails) GetResponse() (result CreateContractRequestResponse, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Action))

	if armName == "Response" {
		result = *u.Response
		ok = true
	}

	return
}

// ManageContractRequestResultSuccessExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//
type ManageContractRequestResultSuccessExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageContractRequestResultSuccessExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageContractRequestResultSuccessExt
func (u ManageContractRequestResultSuccessExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewManageContractRequestResultSuccessExt creates a new  ManageContractRequestResultSuccessExt.
func NewManageContractRequestResultSuccessExt(v LedgerVersion, value interface{}) (result ManageContractRequestResultSuccessExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ManageContractRequestResultSuccess is an XDR NestedStruct defines as:
//
//   struct
//        {
//            union switch (ManageContractRequestAction action)
//            {
//            case CREATE:
//                CreateContractRequestResponse response;
//            case REMOVE:
//                void;
//            } details;
//
//            // reserved for future use
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            } ext;
//        }
//
type ManageContractRequestResultSuccess struct {
	Details ManageContractRequestResultSuccessDetails `json:"details,omitempty"`
	Ext     ManageContractRequestResultSuccessExt     `json:"ext,omitempty"`
}

// ManageContractRequestResult is an XDR Union defines as:
//
//   union ManageContractRequestResult switch (ManageContractRequestResultCode code)
//    {
//    case SUCCESS:
//        struct
//        {
//            union switch (ManageContractRequestAction action)
//            {
//            case CREATE:
//                CreateContractRequestResponse response;
//            case REMOVE:
//                void;
//            } details;
//
//            // reserved for future use
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            } ext;
//        } success;
//    default:
//        void;
//    };
//
type ManageContractRequestResult struct {
	Code    ManageContractRequestResultCode     `json:"code,omitempty"`
	Success *ManageContractRequestResultSuccess `json:"success,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageContractRequestResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageContractRequestResult
func (u ManageContractRequestResult) ArmForSwitch(sw int32) (string, bool) {
	switch ManageContractRequestResultCode(sw) {
	case ManageContractRequestResultCodeSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewManageContractRequestResult creates a new  ManageContractRequestResult.
func NewManageContractRequestResult(code ManageContractRequestResultCode, value interface{}) (result ManageContractRequestResult, err error) {
	result.Code = code
	switch ManageContractRequestResultCode(code) {
	case ManageContractRequestResultCodeSuccess:
		tv, ok := value.(ManageContractRequestResultSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageContractRequestResultSuccess")
			return
		}
		result.Success = &tv
	default:
		// void
	}
	return
}

// MustSuccess retrieves the Success value from the union,
// panicing if the value is not set.
func (u ManageContractRequestResult) MustSuccess() ManageContractRequestResultSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageContractRequestResult) GetSuccess() (result ManageContractRequestResultSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// ManageContractAction is an XDR Enum defines as:
//
//   enum ManageContractAction
//    {
//        ADD_DETAILS = 0,
//        CONFIRM_COMPLETED = 1,
//        START_DISPUTE = 2,
//        RESOLVE_DISPUTE = 3
//    };
//
type ManageContractAction int32

const (
	ManageContractActionAddDetails       ManageContractAction = 0
	ManageContractActionConfirmCompleted ManageContractAction = 1
	ManageContractActionStartDispute     ManageContractAction = 2
	ManageContractActionResolveDispute   ManageContractAction = 3
)

var ManageContractActionAll = []ManageContractAction{
	ManageContractActionAddDetails,
	ManageContractActionConfirmCompleted,
	ManageContractActionStartDispute,
	ManageContractActionResolveDispute,
}

var manageContractActionMap = map[int32]string{
	0: "ManageContractActionAddDetails",
	1: "ManageContractActionConfirmCompleted",
	2: "ManageContractActionStartDispute",
	3: "ManageContractActionResolveDispute",
}

var manageContractActionShortMap = map[int32]string{
	0: "add_details",
	1: "confirm_completed",
	2: "start_dispute",
	3: "resolve_dispute",
}

var manageContractActionRevMap = map[string]int32{
	"ManageContractActionAddDetails":       0,
	"ManageContractActionConfirmCompleted": 1,
	"ManageContractActionStartDispute":     2,
	"ManageContractActionResolveDispute":   3,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ManageContractAction
func (e ManageContractAction) ValidEnum(v int32) bool {
	_, ok := manageContractActionMap[v]
	return ok
}
func (e ManageContractAction) isFlag() bool {
	for i := len(ManageContractActionAll) - 1; i >= 0; i-- {
		expected := ManageContractAction(2) << uint64(len(ManageContractActionAll)-1) >> uint64(len(ManageContractActionAll)-i)
		if expected != ManageContractActionAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ManageContractAction) String() string {
	name, _ := manageContractActionMap[int32(e)]
	return name
}

func (e ManageContractAction) ShortString() string {
	name, _ := manageContractActionShortMap[int32(e)]
	return name
}

func (e ManageContractAction) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range ManageContractActionAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ManageContractAction) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ManageContractAction(t.Value)
	return nil
}

// ManageContractOpData is an XDR NestedUnion defines as:
//
//   union switch (ManageContractAction action)
//        {
//        case ADD_DETAILS:
//            longstring details;
//        case CONFIRM_COMPLETED:
//            void;
//        case START_DISPUTE:
//            longstring disputeReason;
//        case RESOLVE_DISPUTE:
//            bool isRevert;
//        }
//
type ManageContractOpData struct {
	Action        ManageContractAction `json:"action,omitempty"`
	Details       *Longstring          `json:"details,omitempty"`
	DisputeReason *Longstring          `json:"disputeReason,omitempty"`
	IsRevert      *bool                `json:"isRevert,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageContractOpData) SwitchFieldName() string {
	return "Action"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageContractOpData
func (u ManageContractOpData) ArmForSwitch(sw int32) (string, bool) {
	switch ManageContractAction(sw) {
	case ManageContractActionAddDetails:
		return "Details", true
	case ManageContractActionConfirmCompleted:
		return "", true
	case ManageContractActionStartDispute:
		return "DisputeReason", true
	case ManageContractActionResolveDispute:
		return "IsRevert", true
	}
	return "-", false
}

// NewManageContractOpData creates a new  ManageContractOpData.
func NewManageContractOpData(action ManageContractAction, value interface{}) (result ManageContractOpData, err error) {
	result.Action = action
	switch ManageContractAction(action) {
	case ManageContractActionAddDetails:
		tv, ok := value.(Longstring)
		if !ok {
			err = fmt.Errorf("invalid value, must be Longstring")
			return
		}
		result.Details = &tv
	case ManageContractActionConfirmCompleted:
		// void
	case ManageContractActionStartDispute:
		tv, ok := value.(Longstring)
		if !ok {
			err = fmt.Errorf("invalid value, must be Longstring")
			return
		}
		result.DisputeReason = &tv
	case ManageContractActionResolveDispute:
		tv, ok := value.(bool)
		if !ok {
			err = fmt.Errorf("invalid value, must be bool")
			return
		}
		result.IsRevert = &tv
	}
	return
}

// MustDetails retrieves the Details value from the union,
// panicing if the value is not set.
func (u ManageContractOpData) MustDetails() Longstring {
	val, ok := u.GetDetails()

	if !ok {
		panic("arm Details is not set")
	}

	return val
}

// GetDetails retrieves the Details value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageContractOpData) GetDetails() (result Longstring, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Action))

	if armName == "Details" {
		result = *u.Details
		ok = true
	}

	return
}

// MustDisputeReason retrieves the DisputeReason value from the union,
// panicing if the value is not set.
func (u ManageContractOpData) MustDisputeReason() Longstring {
	val, ok := u.GetDisputeReason()

	if !ok {
		panic("arm DisputeReason is not set")
	}

	return val
}

// GetDisputeReason retrieves the DisputeReason value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageContractOpData) GetDisputeReason() (result Longstring, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Action))

	if armName == "DisputeReason" {
		result = *u.DisputeReason
		ok = true
	}

	return
}

// MustIsRevert retrieves the IsRevert value from the union,
// panicing if the value is not set.
func (u ManageContractOpData) MustIsRevert() bool {
	val, ok := u.GetIsRevert()

	if !ok {
		panic("arm IsRevert is not set")
	}

	return val
}

// GetIsRevert retrieves the IsRevert value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageContractOpData) GetIsRevert() (result bool, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Action))

	if armName == "IsRevert" {
		result = *u.IsRevert
		ok = true
	}

	return
}

// ManageContractOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type ManageContractOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageContractOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageContractOpExt
func (u ManageContractOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewManageContractOpExt creates a new  ManageContractOpExt.
func NewManageContractOpExt(v LedgerVersion, value interface{}) (result ManageContractOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ManageContractOp is an XDR Struct defines as:
//
//   struct ManageContractOp
//    {
//        uint64 contractID;
//
//        union switch (ManageContractAction action)
//        {
//        case ADD_DETAILS:
//            longstring details;
//        case CONFIRM_COMPLETED:
//            void;
//        case START_DISPUTE:
//            longstring disputeReason;
//        case RESOLVE_DISPUTE:
//            bool isRevert;
//        }
//        data;
//
//    	// reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type ManageContractOp struct {
	ContractId Uint64               `json:"contractID,omitempty"`
	Data       ManageContractOpData `json:"data,omitempty"`
	Ext        ManageContractOpExt  `json:"ext,omitempty"`
}

// ManageContractResultCode is an XDR Enum defines as:
//
//   enum ManageContractResultCode
//    {
//        // codes considered as "success" for the operation
//        SUCCESS = 0,
//
//        // codes considered as "failure" for the operation
//        MALFORMED = -1,
//        NOT_FOUND = -2, // not found contract
//        NOT_ALLOWED = -3, // only contractor or customer can add details
//        DETAILS_TOO_LONG = -4,
//        DISPUTE_REASON_TOO_LONG = -5,
//        ALREADY_CONFIRMED = -6,
//        INVOICE_NOT_APPROVED = -7, // all contract invoices must be approved
//        DISPUTE_ALREADY_STARTED = -8,
//        CUSTOMER_BALANCE_OVERFLOW = -9
//    };
//
type ManageContractResultCode int32

const (
	ManageContractResultCodeSuccess                 ManageContractResultCode = 0
	ManageContractResultCodeMalformed               ManageContractResultCode = -1
	ManageContractResultCodeNotFound                ManageContractResultCode = -2
	ManageContractResultCodeNotAllowed              ManageContractResultCode = -3
	ManageContractResultCodeDetailsTooLong          ManageContractResultCode = -4
	ManageContractResultCodeDisputeReasonTooLong    ManageContractResultCode = -5
	ManageContractResultCodeAlreadyConfirmed        ManageContractResultCode = -6
	ManageContractResultCodeInvoiceNotApproved      ManageContractResultCode = -7
	ManageContractResultCodeDisputeAlreadyStarted   ManageContractResultCode = -8
	ManageContractResultCodeCustomerBalanceOverflow ManageContractResultCode = -9
)

var ManageContractResultCodeAll = []ManageContractResultCode{
	ManageContractResultCodeSuccess,
	ManageContractResultCodeMalformed,
	ManageContractResultCodeNotFound,
	ManageContractResultCodeNotAllowed,
	ManageContractResultCodeDetailsTooLong,
	ManageContractResultCodeDisputeReasonTooLong,
	ManageContractResultCodeAlreadyConfirmed,
	ManageContractResultCodeInvoiceNotApproved,
	ManageContractResultCodeDisputeAlreadyStarted,
	ManageContractResultCodeCustomerBalanceOverflow,
}

var manageContractResultCodeMap = map[int32]string{
	0:  "ManageContractResultCodeSuccess",
	-1: "ManageContractResultCodeMalformed",
	-2: "ManageContractResultCodeNotFound",
	-3: "ManageContractResultCodeNotAllowed",
	-4: "ManageContractResultCodeDetailsTooLong",
	-5: "ManageContractResultCodeDisputeReasonTooLong",
	-6: "ManageContractResultCodeAlreadyConfirmed",
	-7: "ManageContractResultCodeInvoiceNotApproved",
	-8: "ManageContractResultCodeDisputeAlreadyStarted",
	-9: "ManageContractResultCodeCustomerBalanceOverflow",
}

var manageContractResultCodeShortMap = map[int32]string{
	0:  "success",
	-1: "malformed",
	-2: "not_found",
	-3: "not_allowed",
	-4: "details_too_long",
	-5: "dispute_reason_too_long",
	-6: "already_confirmed",
	-7: "invoice_not_approved",
	-8: "dispute_already_started",
	-9: "customer_balance_overflow",
}

var manageContractResultCodeRevMap = map[string]int32{
	"ManageContractResultCodeSuccess":                 0,
	"ManageContractResultCodeMalformed":               -1,
	"ManageContractResultCodeNotFound":                -2,
	"ManageContractResultCodeNotAllowed":              -3,
	"ManageContractResultCodeDetailsTooLong":          -4,
	"ManageContractResultCodeDisputeReasonTooLong":    -5,
	"ManageContractResultCodeAlreadyConfirmed":        -6,
	"ManageContractResultCodeInvoiceNotApproved":      -7,
	"ManageContractResultCodeDisputeAlreadyStarted":   -8,
	"ManageContractResultCodeCustomerBalanceOverflow": -9,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ManageContractResultCode
func (e ManageContractResultCode) ValidEnum(v int32) bool {
	_, ok := manageContractResultCodeMap[v]
	return ok
}
func (e ManageContractResultCode) isFlag() bool {
	for i := len(ManageContractResultCodeAll) - 1; i >= 0; i-- {
		expected := ManageContractResultCode(2) << uint64(len(ManageContractResultCodeAll)-1) >> uint64(len(ManageContractResultCodeAll)-i)
		if expected != ManageContractResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ManageContractResultCode) String() string {
	name, _ := manageContractResultCodeMap[int32(e)]
	return name
}

func (e ManageContractResultCode) ShortString() string {
	name, _ := manageContractResultCodeShortMap[int32(e)]
	return name
}

func (e ManageContractResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range ManageContractResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ManageContractResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ManageContractResultCode(t.Value)
	return nil
}

// ManageContractResponseData is an XDR NestedUnion defines as:
//
//   union switch (ManageContractAction action)
//        {
//        case CONFIRM_COMPLETED:
//            bool isCompleted;
//        default:
//            void;
//        }
//
type ManageContractResponseData struct {
	Action      ManageContractAction `json:"action,omitempty"`
	IsCompleted *bool                `json:"isCompleted,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageContractResponseData) SwitchFieldName() string {
	return "Action"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageContractResponseData
func (u ManageContractResponseData) ArmForSwitch(sw int32) (string, bool) {
	switch ManageContractAction(sw) {
	case ManageContractActionConfirmCompleted:
		return "IsCompleted", true
	default:
		return "", true
	}
}

// NewManageContractResponseData creates a new  ManageContractResponseData.
func NewManageContractResponseData(action ManageContractAction, value interface{}) (result ManageContractResponseData, err error) {
	result.Action = action
	switch ManageContractAction(action) {
	case ManageContractActionConfirmCompleted:
		tv, ok := value.(bool)
		if !ok {
			err = fmt.Errorf("invalid value, must be bool")
			return
		}
		result.IsCompleted = &tv
	default:
		// void
	}
	return
}

// MustIsCompleted retrieves the IsCompleted value from the union,
// panicing if the value is not set.
func (u ManageContractResponseData) MustIsCompleted() bool {
	val, ok := u.GetIsCompleted()

	if !ok {
		panic("arm IsCompleted is not set")
	}

	return val
}

// GetIsCompleted retrieves the IsCompleted value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageContractResponseData) GetIsCompleted() (result bool, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Action))

	if armName == "IsCompleted" {
		result = *u.IsCompleted
		ok = true
	}

	return
}

// ManageContractResponseExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type ManageContractResponseExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageContractResponseExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageContractResponseExt
func (u ManageContractResponseExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewManageContractResponseExt creates a new  ManageContractResponseExt.
func NewManageContractResponseExt(v LedgerVersion, value interface{}) (result ManageContractResponseExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ManageContractResponse is an XDR Struct defines as:
//
//   struct ManageContractResponse
//    {
//        union switch (ManageContractAction action)
//        {
//        case CONFIRM_COMPLETED:
//            bool isCompleted;
//        default:
//            void;
//        }
//        data;
//
//        // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type ManageContractResponse struct {
	Data ManageContractResponseData `json:"data,omitempty"`
	Ext  ManageContractResponseExt  `json:"ext,omitempty"`
}

// ManageContractResult is an XDR Union defines as:
//
//   union ManageContractResult switch (ManageContractResultCode code)
//    {
//    case SUCCESS:
//        ManageContractResponse response;
//    default:
//        void;
//    };
//
type ManageContractResult struct {
	Code     ManageContractResultCode `json:"code,omitempty"`
	Response *ManageContractResponse  `json:"response,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageContractResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageContractResult
func (u ManageContractResult) ArmForSwitch(sw int32) (string, bool) {
	switch ManageContractResultCode(sw) {
	case ManageContractResultCodeSuccess:
		return "Response", true
	default:
		return "", true
	}
}

// NewManageContractResult creates a new  ManageContractResult.
func NewManageContractResult(code ManageContractResultCode, value interface{}) (result ManageContractResult, err error) {
	result.Code = code
	switch ManageContractResultCode(code) {
	case ManageContractResultCodeSuccess:
		tv, ok := value.(ManageContractResponse)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageContractResponse")
			return
		}
		result.Response = &tv
	default:
		// void
	}
	return
}

// MustResponse retrieves the Response value from the union,
// panicing if the value is not set.
func (u ManageContractResult) MustResponse() ManageContractResponse {
	val, ok := u.GetResponse()

	if !ok {
		panic("arm Response is not set")
	}

	return val
}

// GetResponse retrieves the Response value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageContractResult) GetResponse() (result ManageContractResponse, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Response" {
		result = *u.Response
		ok = true
	}

	return
}

// ManageExternalSystemAccountIdPoolEntryAction is an XDR Enum defines as:
//
//   enum ManageExternalSystemAccountIdPoolEntryAction
//    {
//        CREATE = 0,
//        REMOVE = 1
//    };
//
type ManageExternalSystemAccountIdPoolEntryAction int32

const (
	ManageExternalSystemAccountIdPoolEntryActionCreate ManageExternalSystemAccountIdPoolEntryAction = 0
	ManageExternalSystemAccountIdPoolEntryActionRemove ManageExternalSystemAccountIdPoolEntryAction = 1
)

var ManageExternalSystemAccountIdPoolEntryActionAll = []ManageExternalSystemAccountIdPoolEntryAction{
	ManageExternalSystemAccountIdPoolEntryActionCreate,
	ManageExternalSystemAccountIdPoolEntryActionRemove,
}

var manageExternalSystemAccountIdPoolEntryActionMap = map[int32]string{
	0: "ManageExternalSystemAccountIdPoolEntryActionCreate",
	1: "ManageExternalSystemAccountIdPoolEntryActionRemove",
}

var manageExternalSystemAccountIdPoolEntryActionShortMap = map[int32]string{
	0: "create",
	1: "remove",
}

var manageExternalSystemAccountIdPoolEntryActionRevMap = map[string]int32{
	"ManageExternalSystemAccountIdPoolEntryActionCreate": 0,
	"ManageExternalSystemAccountIdPoolEntryActionRemove": 1,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ManageExternalSystemAccountIdPoolEntryAction
func (e ManageExternalSystemAccountIdPoolEntryAction) ValidEnum(v int32) bool {
	_, ok := manageExternalSystemAccountIdPoolEntryActionMap[v]
	return ok
}
func (e ManageExternalSystemAccountIdPoolEntryAction) isFlag() bool {
	for i := len(ManageExternalSystemAccountIdPoolEntryActionAll) - 1; i >= 0; i-- {
		expected := ManageExternalSystemAccountIdPoolEntryAction(2) << uint64(len(ManageExternalSystemAccountIdPoolEntryActionAll)-1) >> uint64(len(ManageExternalSystemAccountIdPoolEntryActionAll)-i)
		if expected != ManageExternalSystemAccountIdPoolEntryActionAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ManageExternalSystemAccountIdPoolEntryAction) String() string {
	name, _ := manageExternalSystemAccountIdPoolEntryActionMap[int32(e)]
	return name
}

func (e ManageExternalSystemAccountIdPoolEntryAction) ShortString() string {
	name, _ := manageExternalSystemAccountIdPoolEntryActionShortMap[int32(e)]
	return name
}

func (e ManageExternalSystemAccountIdPoolEntryAction) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range ManageExternalSystemAccountIdPoolEntryActionAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ManageExternalSystemAccountIdPoolEntryAction) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ManageExternalSystemAccountIdPoolEntryAction(t.Value)
	return nil
}

// CreateExternalSystemAccountIdPoolEntryActionInputExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type CreateExternalSystemAccountIdPoolEntryActionInputExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreateExternalSystemAccountIdPoolEntryActionInputExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreateExternalSystemAccountIdPoolEntryActionInputExt
func (u CreateExternalSystemAccountIdPoolEntryActionInputExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewCreateExternalSystemAccountIdPoolEntryActionInputExt creates a new  CreateExternalSystemAccountIdPoolEntryActionInputExt.
func NewCreateExternalSystemAccountIdPoolEntryActionInputExt(v LedgerVersion, value interface{}) (result CreateExternalSystemAccountIdPoolEntryActionInputExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// CreateExternalSystemAccountIdPoolEntryActionInput is an XDR Struct defines as:
//
//   struct CreateExternalSystemAccountIdPoolEntryActionInput
//    {
//        int32 externalSystemType;
//        longstring data;
//        uint64 parent;
//
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type CreateExternalSystemAccountIdPoolEntryActionInput struct {
	ExternalSystemType Int32                                                `json:"externalSystemType,omitempty"`
	Data               Longstring                                           `json:"data,omitempty"`
	Parent             Uint64                                               `json:"parent,omitempty"`
	Ext                CreateExternalSystemAccountIdPoolEntryActionInputExt `json:"ext,omitempty"`
}

// DeleteExternalSystemAccountIdPoolEntryActionInputExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type DeleteExternalSystemAccountIdPoolEntryActionInputExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u DeleteExternalSystemAccountIdPoolEntryActionInputExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of DeleteExternalSystemAccountIdPoolEntryActionInputExt
func (u DeleteExternalSystemAccountIdPoolEntryActionInputExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewDeleteExternalSystemAccountIdPoolEntryActionInputExt creates a new  DeleteExternalSystemAccountIdPoolEntryActionInputExt.
func NewDeleteExternalSystemAccountIdPoolEntryActionInputExt(v LedgerVersion, value interface{}) (result DeleteExternalSystemAccountIdPoolEntryActionInputExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// DeleteExternalSystemAccountIdPoolEntryActionInput is an XDR Struct defines as:
//
//   struct DeleteExternalSystemAccountIdPoolEntryActionInput
//    {
//        uint64 poolEntryID;
//
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type DeleteExternalSystemAccountIdPoolEntryActionInput struct {
	PoolEntryId Uint64                                               `json:"poolEntryID,omitempty"`
	Ext         DeleteExternalSystemAccountIdPoolEntryActionInputExt `json:"ext,omitempty"`
}

// ManageExternalSystemAccountIdPoolEntryOpActionInput is an XDR NestedUnion defines as:
//
//   union switch (ManageExternalSystemAccountIdPoolEntryAction action)
//        {
//        case CREATE:
//            CreateExternalSystemAccountIdPoolEntryActionInput createExternalSystemAccountIdPoolEntryActionInput;
//        case REMOVE:
//            DeleteExternalSystemAccountIdPoolEntryActionInput deleteExternalSystemAccountIdPoolEntryActionInput;
//        }
//
type ManageExternalSystemAccountIdPoolEntryOpActionInput struct {
	Action                                            ManageExternalSystemAccountIdPoolEntryAction       `json:"action,omitempty"`
	CreateExternalSystemAccountIdPoolEntryActionInput *CreateExternalSystemAccountIdPoolEntryActionInput `json:"createExternalSystemAccountIdPoolEntryActionInput,omitempty"`
	DeleteExternalSystemAccountIdPoolEntryActionInput *DeleteExternalSystemAccountIdPoolEntryActionInput `json:"deleteExternalSystemAccountIdPoolEntryActionInput,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageExternalSystemAccountIdPoolEntryOpActionInput) SwitchFieldName() string {
	return "Action"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageExternalSystemAccountIdPoolEntryOpActionInput
func (u ManageExternalSystemAccountIdPoolEntryOpActionInput) ArmForSwitch(sw int32) (string, bool) {
	switch ManageExternalSystemAccountIdPoolEntryAction(sw) {
	case ManageExternalSystemAccountIdPoolEntryActionCreate:
		return "CreateExternalSystemAccountIdPoolEntryActionInput", true
	case ManageExternalSystemAccountIdPoolEntryActionRemove:
		return "DeleteExternalSystemAccountIdPoolEntryActionInput", true
	}
	return "-", false
}

// NewManageExternalSystemAccountIdPoolEntryOpActionInput creates a new  ManageExternalSystemAccountIdPoolEntryOpActionInput.
func NewManageExternalSystemAccountIdPoolEntryOpActionInput(action ManageExternalSystemAccountIdPoolEntryAction, value interface{}) (result ManageExternalSystemAccountIdPoolEntryOpActionInput, err error) {
	result.Action = action
	switch ManageExternalSystemAccountIdPoolEntryAction(action) {
	case ManageExternalSystemAccountIdPoolEntryActionCreate:
		tv, ok := value.(CreateExternalSystemAccountIdPoolEntryActionInput)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateExternalSystemAccountIdPoolEntryActionInput")
			return
		}
		result.CreateExternalSystemAccountIdPoolEntryActionInput = &tv
	case ManageExternalSystemAccountIdPoolEntryActionRemove:
		tv, ok := value.(DeleteExternalSystemAccountIdPoolEntryActionInput)
		if !ok {
			err = fmt.Errorf("invalid value, must be DeleteExternalSystemAccountIdPoolEntryActionInput")
			return
		}
		result.DeleteExternalSystemAccountIdPoolEntryActionInput = &tv
	}
	return
}

// MustCreateExternalSystemAccountIdPoolEntryActionInput retrieves the CreateExternalSystemAccountIdPoolEntryActionInput value from the union,
// panicing if the value is not set.
func (u ManageExternalSystemAccountIdPoolEntryOpActionInput) MustCreateExternalSystemAccountIdPoolEntryActionInput() CreateExternalSystemAccountIdPoolEntryActionInput {
	val, ok := u.GetCreateExternalSystemAccountIdPoolEntryActionInput()

	if !ok {
		panic("arm CreateExternalSystemAccountIdPoolEntryActionInput is not set")
	}

	return val
}

// GetCreateExternalSystemAccountIdPoolEntryActionInput retrieves the CreateExternalSystemAccountIdPoolEntryActionInput value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageExternalSystemAccountIdPoolEntryOpActionInput) GetCreateExternalSystemAccountIdPoolEntryActionInput() (result CreateExternalSystemAccountIdPoolEntryActionInput, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Action))

	if armName == "CreateExternalSystemAccountIdPoolEntryActionInput" {
		result = *u.CreateExternalSystemAccountIdPoolEntryActionInput
		ok = true
	}

	return
}

// MustDeleteExternalSystemAccountIdPoolEntryActionInput retrieves the DeleteExternalSystemAccountIdPoolEntryActionInput value from the union,
// panicing if the value is not set.
func (u ManageExternalSystemAccountIdPoolEntryOpActionInput) MustDeleteExternalSystemAccountIdPoolEntryActionInput() DeleteExternalSystemAccountIdPoolEntryActionInput {
	val, ok := u.GetDeleteExternalSystemAccountIdPoolEntryActionInput()

	if !ok {
		panic("arm DeleteExternalSystemAccountIdPoolEntryActionInput is not set")
	}

	return val
}

// GetDeleteExternalSystemAccountIdPoolEntryActionInput retrieves the DeleteExternalSystemAccountIdPoolEntryActionInput value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageExternalSystemAccountIdPoolEntryOpActionInput) GetDeleteExternalSystemAccountIdPoolEntryActionInput() (result DeleteExternalSystemAccountIdPoolEntryActionInput, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Action))

	if armName == "DeleteExternalSystemAccountIdPoolEntryActionInput" {
		result = *u.DeleteExternalSystemAccountIdPoolEntryActionInput
		ok = true
	}

	return
}

// ManageExternalSystemAccountIdPoolEntryOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type ManageExternalSystemAccountIdPoolEntryOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageExternalSystemAccountIdPoolEntryOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageExternalSystemAccountIdPoolEntryOpExt
func (u ManageExternalSystemAccountIdPoolEntryOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewManageExternalSystemAccountIdPoolEntryOpExt creates a new  ManageExternalSystemAccountIdPoolEntryOpExt.
func NewManageExternalSystemAccountIdPoolEntryOpExt(v LedgerVersion, value interface{}) (result ManageExternalSystemAccountIdPoolEntryOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ManageExternalSystemAccountIdPoolEntryOp is an XDR Struct defines as:
//
//   struct ManageExternalSystemAccountIdPoolEntryOp
//    {
//        union switch (ManageExternalSystemAccountIdPoolEntryAction action)
//        {
//        case CREATE:
//            CreateExternalSystemAccountIdPoolEntryActionInput createExternalSystemAccountIdPoolEntryActionInput;
//        case REMOVE:
//            DeleteExternalSystemAccountIdPoolEntryActionInput deleteExternalSystemAccountIdPoolEntryActionInput;
//        } actionInput;
//
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type ManageExternalSystemAccountIdPoolEntryOp struct {
	ActionInput ManageExternalSystemAccountIdPoolEntryOpActionInput `json:"actionInput,omitempty"`
	Ext         ManageExternalSystemAccountIdPoolEntryOpExt         `json:"ext,omitempty"`
}

// ManageExternalSystemAccountIdPoolEntryResultCode is an XDR Enum defines as:
//
//   enum ManageExternalSystemAccountIdPoolEntryResultCode
//    {
//        // codes considered as "success" for the operation
//        SUCCESS = 0,
//
//        // codes considered as "failure" for the operation
//        MALFORMED = -1,
//        ALREADY_EXISTS = -2,
//        AUTO_GENERATED_TYPE_NOT_ALLOWED = -3,
//        NOT_FOUND = -4
//    };
//
type ManageExternalSystemAccountIdPoolEntryResultCode int32

const (
	ManageExternalSystemAccountIdPoolEntryResultCodeSuccess                     ManageExternalSystemAccountIdPoolEntryResultCode = 0
	ManageExternalSystemAccountIdPoolEntryResultCodeMalformed                   ManageExternalSystemAccountIdPoolEntryResultCode = -1
	ManageExternalSystemAccountIdPoolEntryResultCodeAlreadyExists               ManageExternalSystemAccountIdPoolEntryResultCode = -2
	ManageExternalSystemAccountIdPoolEntryResultCodeAutoGeneratedTypeNotAllowed ManageExternalSystemAccountIdPoolEntryResultCode = -3
	ManageExternalSystemAccountIdPoolEntryResultCodeNotFound                    ManageExternalSystemAccountIdPoolEntryResultCode = -4
)

var ManageExternalSystemAccountIdPoolEntryResultCodeAll = []ManageExternalSystemAccountIdPoolEntryResultCode{
	ManageExternalSystemAccountIdPoolEntryResultCodeSuccess,
	ManageExternalSystemAccountIdPoolEntryResultCodeMalformed,
	ManageExternalSystemAccountIdPoolEntryResultCodeAlreadyExists,
	ManageExternalSystemAccountIdPoolEntryResultCodeAutoGeneratedTypeNotAllowed,
	ManageExternalSystemAccountIdPoolEntryResultCodeNotFound,
}

var manageExternalSystemAccountIdPoolEntryResultCodeMap = map[int32]string{
	0:  "ManageExternalSystemAccountIdPoolEntryResultCodeSuccess",
	-1: "ManageExternalSystemAccountIdPoolEntryResultCodeMalformed",
	-2: "ManageExternalSystemAccountIdPoolEntryResultCodeAlreadyExists",
	-3: "ManageExternalSystemAccountIdPoolEntryResultCodeAutoGeneratedTypeNotAllowed",
	-4: "ManageExternalSystemAccountIdPoolEntryResultCodeNotFound",
}

var manageExternalSystemAccountIdPoolEntryResultCodeShortMap = map[int32]string{
	0:  "success",
	-1: "malformed",
	-2: "already_exists",
	-3: "auto_generated_type_not_allowed",
	-4: "not_found",
}

var manageExternalSystemAccountIdPoolEntryResultCodeRevMap = map[string]int32{
	"ManageExternalSystemAccountIdPoolEntryResultCodeSuccess":                     0,
	"ManageExternalSystemAccountIdPoolEntryResultCodeMalformed":                   -1,
	"ManageExternalSystemAccountIdPoolEntryResultCodeAlreadyExists":               -2,
	"ManageExternalSystemAccountIdPoolEntryResultCodeAutoGeneratedTypeNotAllowed": -3,
	"ManageExternalSystemAccountIdPoolEntryResultCodeNotFound":                    -4,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ManageExternalSystemAccountIdPoolEntryResultCode
func (e ManageExternalSystemAccountIdPoolEntryResultCode) ValidEnum(v int32) bool {
	_, ok := manageExternalSystemAccountIdPoolEntryResultCodeMap[v]
	return ok
}
func (e ManageExternalSystemAccountIdPoolEntryResultCode) isFlag() bool {
	for i := len(ManageExternalSystemAccountIdPoolEntryResultCodeAll) - 1; i >= 0; i-- {
		expected := ManageExternalSystemAccountIdPoolEntryResultCode(2) << uint64(len(ManageExternalSystemAccountIdPoolEntryResultCodeAll)-1) >> uint64(len(ManageExternalSystemAccountIdPoolEntryResultCodeAll)-i)
		if expected != ManageExternalSystemAccountIdPoolEntryResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ManageExternalSystemAccountIdPoolEntryResultCode) String() string {
	name, _ := manageExternalSystemAccountIdPoolEntryResultCodeMap[int32(e)]
	return name
}

func (e ManageExternalSystemAccountIdPoolEntryResultCode) ShortString() string {
	name, _ := manageExternalSystemAccountIdPoolEntryResultCodeShortMap[int32(e)]
	return name
}

func (e ManageExternalSystemAccountIdPoolEntryResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range ManageExternalSystemAccountIdPoolEntryResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ManageExternalSystemAccountIdPoolEntryResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ManageExternalSystemAccountIdPoolEntryResultCode(t.Value)
	return nil
}

// ManageExternalSystemAccountIdPoolEntrySuccessExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type ManageExternalSystemAccountIdPoolEntrySuccessExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageExternalSystemAccountIdPoolEntrySuccessExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageExternalSystemAccountIdPoolEntrySuccessExt
func (u ManageExternalSystemAccountIdPoolEntrySuccessExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewManageExternalSystemAccountIdPoolEntrySuccessExt creates a new  ManageExternalSystemAccountIdPoolEntrySuccessExt.
func NewManageExternalSystemAccountIdPoolEntrySuccessExt(v LedgerVersion, value interface{}) (result ManageExternalSystemAccountIdPoolEntrySuccessExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ManageExternalSystemAccountIdPoolEntrySuccess is an XDR Struct defines as:
//
//   struct ManageExternalSystemAccountIdPoolEntrySuccess {
//    	uint64 poolEntryID;
//    	// reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type ManageExternalSystemAccountIdPoolEntrySuccess struct {
	PoolEntryId Uint64                                           `json:"poolEntryID,omitempty"`
	Ext         ManageExternalSystemAccountIdPoolEntrySuccessExt `json:"ext,omitempty"`
}

// ManageExternalSystemAccountIdPoolEntryResult is an XDR Union defines as:
//
//   union ManageExternalSystemAccountIdPoolEntryResult switch (ManageExternalSystemAccountIdPoolEntryResultCode code)
//    {
//    case SUCCESS:
//        ManageExternalSystemAccountIdPoolEntrySuccess success;
//    default:
//        void;
//    };
//
type ManageExternalSystemAccountIdPoolEntryResult struct {
	Code    ManageExternalSystemAccountIdPoolEntryResultCode `json:"code,omitempty"`
	Success *ManageExternalSystemAccountIdPoolEntrySuccess   `json:"success,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageExternalSystemAccountIdPoolEntryResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageExternalSystemAccountIdPoolEntryResult
func (u ManageExternalSystemAccountIdPoolEntryResult) ArmForSwitch(sw int32) (string, bool) {
	switch ManageExternalSystemAccountIdPoolEntryResultCode(sw) {
	case ManageExternalSystemAccountIdPoolEntryResultCodeSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewManageExternalSystemAccountIdPoolEntryResult creates a new  ManageExternalSystemAccountIdPoolEntryResult.
func NewManageExternalSystemAccountIdPoolEntryResult(code ManageExternalSystemAccountIdPoolEntryResultCode, value interface{}) (result ManageExternalSystemAccountIdPoolEntryResult, err error) {
	result.Code = code
	switch ManageExternalSystemAccountIdPoolEntryResultCode(code) {
	case ManageExternalSystemAccountIdPoolEntryResultCodeSuccess:
		tv, ok := value.(ManageExternalSystemAccountIdPoolEntrySuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageExternalSystemAccountIdPoolEntrySuccess")
			return
		}
		result.Success = &tv
	default:
		// void
	}
	return
}

// MustSuccess retrieves the Success value from the union,
// panicing if the value is not set.
func (u ManageExternalSystemAccountIdPoolEntryResult) MustSuccess() ManageExternalSystemAccountIdPoolEntrySuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageExternalSystemAccountIdPoolEntryResult) GetSuccess() (result ManageExternalSystemAccountIdPoolEntrySuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// ManageInvoiceRequestAction is an XDR Enum defines as:
//
//   enum ManageInvoiceRequestAction
//    {
//        CREATE = 0,
//        REMOVE = 1
//    };
//
type ManageInvoiceRequestAction int32

const (
	ManageInvoiceRequestActionCreate ManageInvoiceRequestAction = 0
	ManageInvoiceRequestActionRemove ManageInvoiceRequestAction = 1
)

var ManageInvoiceRequestActionAll = []ManageInvoiceRequestAction{
	ManageInvoiceRequestActionCreate,
	ManageInvoiceRequestActionRemove,
}

var manageInvoiceRequestActionMap = map[int32]string{
	0: "ManageInvoiceRequestActionCreate",
	1: "ManageInvoiceRequestActionRemove",
}

var manageInvoiceRequestActionShortMap = map[int32]string{
	0: "create",
	1: "remove",
}

var manageInvoiceRequestActionRevMap = map[string]int32{
	"ManageInvoiceRequestActionCreate": 0,
	"ManageInvoiceRequestActionRemove": 1,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ManageInvoiceRequestAction
func (e ManageInvoiceRequestAction) ValidEnum(v int32) bool {
	_, ok := manageInvoiceRequestActionMap[v]
	return ok
}
func (e ManageInvoiceRequestAction) isFlag() bool {
	for i := len(ManageInvoiceRequestActionAll) - 1; i >= 0; i-- {
		expected := ManageInvoiceRequestAction(2) << uint64(len(ManageInvoiceRequestActionAll)-1) >> uint64(len(ManageInvoiceRequestActionAll)-i)
		if expected != ManageInvoiceRequestActionAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ManageInvoiceRequestAction) String() string {
	name, _ := manageInvoiceRequestActionMap[int32(e)]
	return name
}

func (e ManageInvoiceRequestAction) ShortString() string {
	name, _ := manageInvoiceRequestActionShortMap[int32(e)]
	return name
}

func (e ManageInvoiceRequestAction) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range ManageInvoiceRequestActionAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ManageInvoiceRequestAction) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ManageInvoiceRequestAction(t.Value)
	return nil
}

// InvoiceCreationRequestExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type InvoiceCreationRequestExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u InvoiceCreationRequestExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of InvoiceCreationRequestExt
func (u InvoiceCreationRequestExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewInvoiceCreationRequestExt creates a new  InvoiceCreationRequestExt.
func NewInvoiceCreationRequestExt(v LedgerVersion, value interface{}) (result InvoiceCreationRequestExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// InvoiceCreationRequest is an XDR Struct defines as:
//
//   struct InvoiceCreationRequest
//    {
//        AssetCode asset;
//        AccountID sender;
//        uint64 amount; // not allowed to set 0
//
//        uint64 *contractID;
//        longstring details;
//
//        // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type InvoiceCreationRequest struct {
	Asset      AssetCode                 `json:"asset,omitempty"`
	Sender     AccountId                 `json:"sender,omitempty"`
	Amount     Uint64                    `json:"amount,omitempty"`
	ContractId *Uint64                   `json:"contractID,omitempty"`
	Details    Longstring                `json:"details,omitempty"`
	Ext        InvoiceCreationRequestExt `json:"ext,omitempty"`
}

// ManageInvoiceRequestOpDetails is an XDR NestedUnion defines as:
//
//   union switch (ManageInvoiceRequestAction action){
//        case CREATE:
//            InvoiceCreationRequest invoiceRequest;
//        case REMOVE:
//            uint64 requestID;
//        }
//
type ManageInvoiceRequestOpDetails struct {
	Action         ManageInvoiceRequestAction `json:"action,omitempty"`
	InvoiceRequest *InvoiceCreationRequest    `json:"invoiceRequest,omitempty"`
	RequestId      *Uint64                    `json:"requestID,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageInvoiceRequestOpDetails) SwitchFieldName() string {
	return "Action"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageInvoiceRequestOpDetails
func (u ManageInvoiceRequestOpDetails) ArmForSwitch(sw int32) (string, bool) {
	switch ManageInvoiceRequestAction(sw) {
	case ManageInvoiceRequestActionCreate:
		return "InvoiceRequest", true
	case ManageInvoiceRequestActionRemove:
		return "RequestId", true
	}
	return "-", false
}

// NewManageInvoiceRequestOpDetails creates a new  ManageInvoiceRequestOpDetails.
func NewManageInvoiceRequestOpDetails(action ManageInvoiceRequestAction, value interface{}) (result ManageInvoiceRequestOpDetails, err error) {
	result.Action = action
	switch ManageInvoiceRequestAction(action) {
	case ManageInvoiceRequestActionCreate:
		tv, ok := value.(InvoiceCreationRequest)
		if !ok {
			err = fmt.Errorf("invalid value, must be InvoiceCreationRequest")
			return
		}
		result.InvoiceRequest = &tv
	case ManageInvoiceRequestActionRemove:
		tv, ok := value.(Uint64)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint64")
			return
		}
		result.RequestId = &tv
	}
	return
}

// MustInvoiceRequest retrieves the InvoiceRequest value from the union,
// panicing if the value is not set.
func (u ManageInvoiceRequestOpDetails) MustInvoiceRequest() InvoiceCreationRequest {
	val, ok := u.GetInvoiceRequest()

	if !ok {
		panic("arm InvoiceRequest is not set")
	}

	return val
}

// GetInvoiceRequest retrieves the InvoiceRequest value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageInvoiceRequestOpDetails) GetInvoiceRequest() (result InvoiceCreationRequest, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Action))

	if armName == "InvoiceRequest" {
		result = *u.InvoiceRequest
		ok = true
	}

	return
}

// MustRequestId retrieves the RequestId value from the union,
// panicing if the value is not set.
func (u ManageInvoiceRequestOpDetails) MustRequestId() Uint64 {
	val, ok := u.GetRequestId()

	if !ok {
		panic("arm RequestId is not set")
	}

	return val
}

// GetRequestId retrieves the RequestId value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageInvoiceRequestOpDetails) GetRequestId() (result Uint64, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Action))

	if armName == "RequestId" {
		result = *u.RequestId
		ok = true
	}

	return
}

// ManageInvoiceRequestOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type ManageInvoiceRequestOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageInvoiceRequestOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageInvoiceRequestOpExt
func (u ManageInvoiceRequestOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewManageInvoiceRequestOpExt creates a new  ManageInvoiceRequestOpExt.
func NewManageInvoiceRequestOpExt(v LedgerVersion, value interface{}) (result ManageInvoiceRequestOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ManageInvoiceRequestOp is an XDR Struct defines as:
//
//   struct ManageInvoiceRequestOp
//    {
//        union switch (ManageInvoiceRequestAction action){
//        case CREATE:
//            InvoiceCreationRequest invoiceRequest;
//        case REMOVE:
//            uint64 requestID;
//        } details;
//
//    	// reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type ManageInvoiceRequestOp struct {
	Details ManageInvoiceRequestOpDetails `json:"details,omitempty"`
	Ext     ManageInvoiceRequestOpExt     `json:"ext,omitempty"`
}

// ManageInvoiceRequestResultCode is an XDR Enum defines as:
//
//   enum ManageInvoiceRequestResultCode
//    {
//        // codes considered as "success" for the operation
//        SUCCESS = 0,
//
//        // codes considered as "failure" for the operation
//        MALFORMED = -1,
//        BALANCE_NOT_FOUND = -2, // sender balance not found
//        NOT_FOUND = -3, // not found invoice request, when try to remove
//        TOO_MANY_INVOICES = -4,
//        DETAILS_TOO_LONG = -5,
//        NOT_ALLOWED_TO_REMOVE = -6, // only invoice creator can remove invoice
//        CONTRACT_NOT_FOUND = -7,
//        ONLY_CONTRACTOR_CAN_ATTACH_INVOICE_TO_CONTRACT = -8,
//        SENDER_ACCOUNT_MISMATCHED = -9,
//        INVOICE_IS_APPROVED = -10 // not allowed to remove approved invoice
//    };
//
type ManageInvoiceRequestResultCode int32

const (
	ManageInvoiceRequestResultCodeSuccess                                  ManageInvoiceRequestResultCode = 0
	ManageInvoiceRequestResultCodeMalformed                                ManageInvoiceRequestResultCode = -1
	ManageInvoiceRequestResultCodeBalanceNotFound                          ManageInvoiceRequestResultCode = -2
	ManageInvoiceRequestResultCodeNotFound                                 ManageInvoiceRequestResultCode = -3
	ManageInvoiceRequestResultCodeTooManyInvoices                          ManageInvoiceRequestResultCode = -4
	ManageInvoiceRequestResultCodeDetailsTooLong                           ManageInvoiceRequestResultCode = -5
	ManageInvoiceRequestResultCodeNotAllowedToRemove                       ManageInvoiceRequestResultCode = -6
	ManageInvoiceRequestResultCodeContractNotFound                         ManageInvoiceRequestResultCode = -7
	ManageInvoiceRequestResultCodeOnlyContractorCanAttachInvoiceToContract ManageInvoiceRequestResultCode = -8
	ManageInvoiceRequestResultCodeSenderAccountMismatched                  ManageInvoiceRequestResultCode = -9
	ManageInvoiceRequestResultCodeInvoiceIsApproved                        ManageInvoiceRequestResultCode = -10
)

var ManageInvoiceRequestResultCodeAll = []ManageInvoiceRequestResultCode{
	ManageInvoiceRequestResultCodeSuccess,
	ManageInvoiceRequestResultCodeMalformed,
	ManageInvoiceRequestResultCodeBalanceNotFound,
	ManageInvoiceRequestResultCodeNotFound,
	ManageInvoiceRequestResultCodeTooManyInvoices,
	ManageInvoiceRequestResultCodeDetailsTooLong,
	ManageInvoiceRequestResultCodeNotAllowedToRemove,
	ManageInvoiceRequestResultCodeContractNotFound,
	ManageInvoiceRequestResultCodeOnlyContractorCanAttachInvoiceToContract,
	ManageInvoiceRequestResultCodeSenderAccountMismatched,
	ManageInvoiceRequestResultCodeInvoiceIsApproved,
}

var manageInvoiceRequestResultCodeMap = map[int32]string{
	0:   "ManageInvoiceRequestResultCodeSuccess",
	-1:  "ManageInvoiceRequestResultCodeMalformed",
	-2:  "ManageInvoiceRequestResultCodeBalanceNotFound",
	-3:  "ManageInvoiceRequestResultCodeNotFound",
	-4:  "ManageInvoiceRequestResultCodeTooManyInvoices",
	-5:  "ManageInvoiceRequestResultCodeDetailsTooLong",
	-6:  "ManageInvoiceRequestResultCodeNotAllowedToRemove",
	-7:  "ManageInvoiceRequestResultCodeContractNotFound",
	-8:  "ManageInvoiceRequestResultCodeOnlyContractorCanAttachInvoiceToContract",
	-9:  "ManageInvoiceRequestResultCodeSenderAccountMismatched",
	-10: "ManageInvoiceRequestResultCodeInvoiceIsApproved",
}

var manageInvoiceRequestResultCodeShortMap = map[int32]string{
	0:   "success",
	-1:  "malformed",
	-2:  "balance_not_found",
	-3:  "not_found",
	-4:  "too_many_invoices",
	-5:  "details_too_long",
	-6:  "not_allowed_to_remove",
	-7:  "contract_not_found",
	-8:  "only_contractor_can_attach_invoice_to_contract",
	-9:  "sender_account_mismatched",
	-10: "invoice_is_approved",
}

var manageInvoiceRequestResultCodeRevMap = map[string]int32{
	"ManageInvoiceRequestResultCodeSuccess":                                  0,
	"ManageInvoiceRequestResultCodeMalformed":                                -1,
	"ManageInvoiceRequestResultCodeBalanceNotFound":                          -2,
	"ManageInvoiceRequestResultCodeNotFound":                                 -3,
	"ManageInvoiceRequestResultCodeTooManyInvoices":                          -4,
	"ManageInvoiceRequestResultCodeDetailsTooLong":                           -5,
	"ManageInvoiceRequestResultCodeNotAllowedToRemove":                       -6,
	"ManageInvoiceRequestResultCodeContractNotFound":                         -7,
	"ManageInvoiceRequestResultCodeOnlyContractorCanAttachInvoiceToContract": -8,
	"ManageInvoiceRequestResultCodeSenderAccountMismatched":                  -9,
	"ManageInvoiceRequestResultCodeInvoiceIsApproved":                        -10,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ManageInvoiceRequestResultCode
func (e ManageInvoiceRequestResultCode) ValidEnum(v int32) bool {
	_, ok := manageInvoiceRequestResultCodeMap[v]
	return ok
}
func (e ManageInvoiceRequestResultCode) isFlag() bool {
	for i := len(ManageInvoiceRequestResultCodeAll) - 1; i >= 0; i-- {
		expected := ManageInvoiceRequestResultCode(2) << uint64(len(ManageInvoiceRequestResultCodeAll)-1) >> uint64(len(ManageInvoiceRequestResultCodeAll)-i)
		if expected != ManageInvoiceRequestResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ManageInvoiceRequestResultCode) String() string {
	name, _ := manageInvoiceRequestResultCodeMap[int32(e)]
	return name
}

func (e ManageInvoiceRequestResultCode) ShortString() string {
	name, _ := manageInvoiceRequestResultCodeShortMap[int32(e)]
	return name
}

func (e ManageInvoiceRequestResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range ManageInvoiceRequestResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ManageInvoiceRequestResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ManageInvoiceRequestResultCode(t.Value)
	return nil
}

// CreateInvoiceRequestResponseExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type CreateInvoiceRequestResponseExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u CreateInvoiceRequestResponseExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of CreateInvoiceRequestResponseExt
func (u CreateInvoiceRequestResponseExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewCreateInvoiceRequestResponseExt creates a new  CreateInvoiceRequestResponseExt.
func NewCreateInvoiceRequestResponseExt(v LedgerVersion, value interface{}) (result CreateInvoiceRequestResponseExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// CreateInvoiceRequestResponse is an XDR Struct defines as:
//
//   struct CreateInvoiceRequestResponse
//    {
//    	BalanceID receiverBalance;
//    	BalanceID senderBalance;
//
//    	uint64 requestID;
//
//    	union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type CreateInvoiceRequestResponse struct {
	ReceiverBalance BalanceId                       `json:"receiverBalance,omitempty"`
	SenderBalance   BalanceId                       `json:"senderBalance,omitempty"`
	RequestId       Uint64                          `json:"requestID,omitempty"`
	Ext             CreateInvoiceRequestResponseExt `json:"ext,omitempty"`
}

// ManageInvoiceRequestResultSuccessDetails is an XDR NestedUnion defines as:
//
//   union switch (ManageInvoiceRequestAction action)
//            {
//            case CREATE:
//                CreateInvoiceRequestResponse response;
//            case REMOVE:
//                void;
//            }
//
type ManageInvoiceRequestResultSuccessDetails struct {
	Action   ManageInvoiceRequestAction    `json:"action,omitempty"`
	Response *CreateInvoiceRequestResponse `json:"response,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageInvoiceRequestResultSuccessDetails) SwitchFieldName() string {
	return "Action"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageInvoiceRequestResultSuccessDetails
func (u ManageInvoiceRequestResultSuccessDetails) ArmForSwitch(sw int32) (string, bool) {
	switch ManageInvoiceRequestAction(sw) {
	case ManageInvoiceRequestActionCreate:
		return "Response", true
	case ManageInvoiceRequestActionRemove:
		return "", true
	}
	return "-", false
}

// NewManageInvoiceRequestResultSuccessDetails creates a new  ManageInvoiceRequestResultSuccessDetails.
func NewManageInvoiceRequestResultSuccessDetails(action ManageInvoiceRequestAction, value interface{}) (result ManageInvoiceRequestResultSuccessDetails, err error) {
	result.Action = action
	switch ManageInvoiceRequestAction(action) {
	case ManageInvoiceRequestActionCreate:
		tv, ok := value.(CreateInvoiceRequestResponse)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateInvoiceRequestResponse")
			return
		}
		result.Response = &tv
	case ManageInvoiceRequestActionRemove:
		// void
	}
	return
}

// MustResponse retrieves the Response value from the union,
// panicing if the value is not set.
func (u ManageInvoiceRequestResultSuccessDetails) MustResponse() CreateInvoiceRequestResponse {
	val, ok := u.GetResponse()

	if !ok {
		panic("arm Response is not set")
	}

	return val
}

// GetResponse retrieves the Response value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageInvoiceRequestResultSuccessDetails) GetResponse() (result CreateInvoiceRequestResponse, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Action))

	if armName == "Response" {
		result = *u.Response
		ok = true
	}

	return
}

// ManageInvoiceRequestResultSuccessExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//
type ManageInvoiceRequestResultSuccessExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageInvoiceRequestResultSuccessExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageInvoiceRequestResultSuccessExt
func (u ManageInvoiceRequestResultSuccessExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewManageInvoiceRequestResultSuccessExt creates a new  ManageInvoiceRequestResultSuccessExt.
func NewManageInvoiceRequestResultSuccessExt(v LedgerVersion, value interface{}) (result ManageInvoiceRequestResultSuccessExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ManageInvoiceRequestResultSuccess is an XDR NestedStruct defines as:
//
//   struct
//        {
//            union switch (ManageInvoiceRequestAction action)
//            {
//            case CREATE:
//                CreateInvoiceRequestResponse response;
//            case REMOVE:
//                void;
//            } details;
//
//            // reserved for future use
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            } ext;
//        }
//
type ManageInvoiceRequestResultSuccess struct {
	Details ManageInvoiceRequestResultSuccessDetails `json:"details,omitempty"`
	Ext     ManageInvoiceRequestResultSuccessExt     `json:"ext,omitempty"`
}

// ManageInvoiceRequestResult is an XDR Union defines as:
//
//   union ManageInvoiceRequestResult switch (ManageInvoiceRequestResultCode code)
//    {
//    case SUCCESS:
//        struct
//        {
//            union switch (ManageInvoiceRequestAction action)
//            {
//            case CREATE:
//                CreateInvoiceRequestResponse response;
//            case REMOVE:
//                void;
//            } details;
//
//            // reserved for future use
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            } ext;
//        } success;
//    default:
//        void;
//    };
//
type ManageInvoiceRequestResult struct {
	Code    ManageInvoiceRequestResultCode     `json:"code,omitempty"`
	Success *ManageInvoiceRequestResultSuccess `json:"success,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageInvoiceRequestResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageInvoiceRequestResult
func (u ManageInvoiceRequestResult) ArmForSwitch(sw int32) (string, bool) {
	switch ManageInvoiceRequestResultCode(sw) {
	case ManageInvoiceRequestResultCodeSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewManageInvoiceRequestResult creates a new  ManageInvoiceRequestResult.
func NewManageInvoiceRequestResult(code ManageInvoiceRequestResultCode, value interface{}) (result ManageInvoiceRequestResult, err error) {
	result.Code = code
	switch ManageInvoiceRequestResultCode(code) {
	case ManageInvoiceRequestResultCodeSuccess:
		tv, ok := value.(ManageInvoiceRequestResultSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageInvoiceRequestResultSuccess")
			return
		}
		result.Success = &tv
	default:
		// void
	}
	return
}

// MustSuccess retrieves the Success value from the union,
// panicing if the value is not set.
func (u ManageInvoiceRequestResult) MustSuccess() ManageInvoiceRequestResultSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageInvoiceRequestResult) GetSuccess() (result ManageInvoiceRequestResultSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// ManageKvAction is an XDR Enum defines as:
//
//   enum ManageKVAction
//        {
//            PUT = 1,
//            REMOVE = 2
//        };
//
type ManageKvAction int32

const (
	ManageKvActionPut    ManageKvAction = 1
	ManageKvActionRemove ManageKvAction = 2
)

var ManageKvActionAll = []ManageKvAction{
	ManageKvActionPut,
	ManageKvActionRemove,
}

var manageKvActionMap = map[int32]string{
	1: "ManageKvActionPut",
	2: "ManageKvActionRemove",
}

var manageKvActionShortMap = map[int32]string{
	1: "put",
	2: "remove",
}

var manageKvActionRevMap = map[string]int32{
	"ManageKvActionPut":    1,
	"ManageKvActionRemove": 2,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ManageKvAction
func (e ManageKvAction) ValidEnum(v int32) bool {
	_, ok := manageKvActionMap[v]
	return ok
}
func (e ManageKvAction) isFlag() bool {
	for i := len(ManageKvActionAll) - 1; i >= 0; i-- {
		expected := ManageKvAction(2) << uint64(len(ManageKvActionAll)-1) >> uint64(len(ManageKvActionAll)-i)
		if expected != ManageKvActionAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ManageKvAction) String() string {
	name, _ := manageKvActionMap[int32(e)]
	return name
}

func (e ManageKvAction) ShortString() string {
	name, _ := manageKvActionShortMap[int32(e)]
	return name
}

func (e ManageKvAction) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range ManageKvActionAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ManageKvAction) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ManageKvAction(t.Value)
	return nil
}

// ManageKeyValueOpAction is an XDR NestedUnion defines as:
//
//   union switch(ManageKVAction action)
//            {
//                case PUT:
//                    KeyValueEntry value;
//                case REMOVE:
//                    void;
//            }
//
type ManageKeyValueOpAction struct {
	Action ManageKvAction `json:"action,omitempty"`
	Value  *KeyValueEntry `json:"value,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageKeyValueOpAction) SwitchFieldName() string {
	return "Action"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageKeyValueOpAction
func (u ManageKeyValueOpAction) ArmForSwitch(sw int32) (string, bool) {
	switch ManageKvAction(sw) {
	case ManageKvActionPut:
		return "Value", true
	case ManageKvActionRemove:
		return "", true
	}
	return "-", false
}

// NewManageKeyValueOpAction creates a new  ManageKeyValueOpAction.
func NewManageKeyValueOpAction(action ManageKvAction, value interface{}) (result ManageKeyValueOpAction, err error) {
	result.Action = action
	switch ManageKvAction(action) {
	case ManageKvActionPut:
		tv, ok := value.(KeyValueEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be KeyValueEntry")
			return
		}
		result.Value = &tv
	case ManageKvActionRemove:
		// void
	}
	return
}

// MustValue retrieves the Value value from the union,
// panicing if the value is not set.
func (u ManageKeyValueOpAction) MustValue() KeyValueEntry {
	val, ok := u.GetValue()

	if !ok {
		panic("arm Value is not set")
	}

	return val
}

// GetValue retrieves the Value value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageKeyValueOpAction) GetValue() (result KeyValueEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Action))

	if armName == "Value" {
		result = *u.Value
		ok = true
	}

	return
}

// ManageKeyValueOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//
type ManageKeyValueOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageKeyValueOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageKeyValueOpExt
func (u ManageKeyValueOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewManageKeyValueOpExt creates a new  ManageKeyValueOpExt.
func NewManageKeyValueOpExt(v LedgerVersion, value interface{}) (result ManageKeyValueOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ManageKeyValueOp is an XDR Struct defines as:
//
//   struct ManageKeyValueOp
//        {
//            string256 key;
//            union switch(ManageKVAction action)
//            {
//                case PUT:
//                    KeyValueEntry value;
//                case REMOVE:
//                    void;
//            }
//            action;
//
//            // reserved for future use
//            union switch (LedgerVersion v)
//            {
//            case EMPTY_VERSION:
//                void;
//            }
//            ext;
//        };
//
type ManageKeyValueOp struct {
	Key    String256              `json:"key,omitempty"`
	Action ManageKeyValueOpAction `json:"action,omitempty"`
	Ext    ManageKeyValueOpExt    `json:"ext,omitempty"`
}

// ManageKeyValueSuccessExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//            {
//                case EMPTY_VERSION:
//                    void;
//            }
//
type ManageKeyValueSuccessExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageKeyValueSuccessExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageKeyValueSuccessExt
func (u ManageKeyValueSuccessExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewManageKeyValueSuccessExt creates a new  ManageKeyValueSuccessExt.
func NewManageKeyValueSuccessExt(v LedgerVersion, value interface{}) (result ManageKeyValueSuccessExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ManageKeyValueSuccess is an XDR Struct defines as:
//
//   struct ManageKeyValueSuccess
//        {
//            // reserved for future use
//            union switch (LedgerVersion v)
//            {
//                case EMPTY_VERSION:
//                    void;
//            }
//            ext;
//        };
//
type ManageKeyValueSuccess struct {
	Ext ManageKeyValueSuccessExt `json:"ext,omitempty"`
}

// ManageKeyValueResultCode is an XDR Enum defines as:
//
//   enum ManageKeyValueResultCode
//        {
//            SUCCESS = 1,
//            NOT_FOUND = -1,
//            INVALID_TYPE = -2
//        };
//
type ManageKeyValueResultCode int32

const (
	ManageKeyValueResultCodeSuccess     ManageKeyValueResultCode = 1
	ManageKeyValueResultCodeNotFound    ManageKeyValueResultCode = -1
	ManageKeyValueResultCodeInvalidType ManageKeyValueResultCode = -2
)

var ManageKeyValueResultCodeAll = []ManageKeyValueResultCode{
	ManageKeyValueResultCodeSuccess,
	ManageKeyValueResultCodeNotFound,
	ManageKeyValueResultCodeInvalidType,
}

var manageKeyValueResultCodeMap = map[int32]string{
	1:  "ManageKeyValueResultCodeSuccess",
	-1: "ManageKeyValueResultCodeNotFound",
	-2: "ManageKeyValueResultCodeInvalidType",
}

var manageKeyValueResultCodeShortMap = map[int32]string{
	1:  "success",
	-1: "not_found",
	-2: "invalid_type",
}

var manageKeyValueResultCodeRevMap = map[string]int32{
	"ManageKeyValueResultCodeSuccess":     1,
	"ManageKeyValueResultCodeNotFound":    -1,
	"ManageKeyValueResultCodeInvalidType": -2,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ManageKeyValueResultCode
func (e ManageKeyValueResultCode) ValidEnum(v int32) bool {
	_, ok := manageKeyValueResultCodeMap[v]
	return ok
}
func (e ManageKeyValueResultCode) isFlag() bool {
	for i := len(ManageKeyValueResultCodeAll) - 1; i >= 0; i-- {
		expected := ManageKeyValueResultCode(2) << uint64(len(ManageKeyValueResultCodeAll)-1) >> uint64(len(ManageKeyValueResultCodeAll)-i)
		if expected != ManageKeyValueResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ManageKeyValueResultCode) String() string {
	name, _ := manageKeyValueResultCodeMap[int32(e)]
	return name
}

func (e ManageKeyValueResultCode) ShortString() string {
	name, _ := manageKeyValueResultCodeShortMap[int32(e)]
	return name
}

func (e ManageKeyValueResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range ManageKeyValueResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ManageKeyValueResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ManageKeyValueResultCode(t.Value)
	return nil
}

// ManageKeyValueResult is an XDR Union defines as:
//
//   union ManageKeyValueResult switch (ManageKeyValueResultCode code)
//        {
//            case SUCCESS:
//                ManageKeyValueSuccess success;
//            default:
//                void;
//        };
//
type ManageKeyValueResult struct {
	Code    ManageKeyValueResultCode `json:"code,omitempty"`
	Success *ManageKeyValueSuccess   `json:"success,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageKeyValueResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageKeyValueResult
func (u ManageKeyValueResult) ArmForSwitch(sw int32) (string, bool) {
	switch ManageKeyValueResultCode(sw) {
	case ManageKeyValueResultCodeSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewManageKeyValueResult creates a new  ManageKeyValueResult.
func NewManageKeyValueResult(code ManageKeyValueResultCode, value interface{}) (result ManageKeyValueResult, err error) {
	result.Code = code
	switch ManageKeyValueResultCode(code) {
	case ManageKeyValueResultCodeSuccess:
		tv, ok := value.(ManageKeyValueSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageKeyValueSuccess")
			return
		}
		result.Success = &tv
	default:
		// void
	}
	return
}

// MustSuccess retrieves the Success value from the union,
// panicing if the value is not set.
func (u ManageKeyValueResult) MustSuccess() ManageKeyValueSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageKeyValueResult) GetSuccess() (result ManageKeyValueSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// ManageLimitsAction is an XDR Enum defines as:
//
//   enum ManageLimitsAction
//    {
//        CREATE = 0,
//        REMOVE = 1
//    };
//
type ManageLimitsAction int32

const (
	ManageLimitsActionCreate ManageLimitsAction = 0
	ManageLimitsActionRemove ManageLimitsAction = 1
)

var ManageLimitsActionAll = []ManageLimitsAction{
	ManageLimitsActionCreate,
	ManageLimitsActionRemove,
}

var manageLimitsActionMap = map[int32]string{
	0: "ManageLimitsActionCreate",
	1: "ManageLimitsActionRemove",
}

var manageLimitsActionShortMap = map[int32]string{
	0: "create",
	1: "remove",
}

var manageLimitsActionRevMap = map[string]int32{
	"ManageLimitsActionCreate": 0,
	"ManageLimitsActionRemove": 1,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ManageLimitsAction
func (e ManageLimitsAction) ValidEnum(v int32) bool {
	_, ok := manageLimitsActionMap[v]
	return ok
}
func (e ManageLimitsAction) isFlag() bool {
	for i := len(ManageLimitsActionAll) - 1; i >= 0; i-- {
		expected := ManageLimitsAction(2) << uint64(len(ManageLimitsActionAll)-1) >> uint64(len(ManageLimitsActionAll)-i)
		if expected != ManageLimitsActionAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ManageLimitsAction) String() string {
	name, _ := manageLimitsActionMap[int32(e)]
	return name
}

func (e ManageLimitsAction) ShortString() string {
	name, _ := manageLimitsActionShortMap[int32(e)]
	return name
}

func (e ManageLimitsAction) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range ManageLimitsActionAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ManageLimitsAction) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ManageLimitsAction(t.Value)
	return nil
}

// LimitsCreateDetails is an XDR Struct defines as:
//
//   struct LimitsCreateDetails
//    {
//        AccountType *accountType;
//        AccountID   *accountID;
//        StatsOpType statsOpType;
//        AssetCode   assetCode;
//        bool        isConvertNeeded;
//
//        uint64 dailyOut;
//        uint64 weeklyOut;
//        uint64 monthlyOut;
//        uint64 annualOut;
//    };
//
type LimitsCreateDetails struct {
	AccountType     *AccountType `json:"accountType,omitempty"`
	AccountId       *AccountId   `json:"accountID,omitempty"`
	StatsOpType     StatsOpType  `json:"statsOpType,omitempty"`
	AssetCode       AssetCode    `json:"assetCode,omitempty"`
	IsConvertNeeded bool         `json:"isConvertNeeded,omitempty"`
	DailyOut        Uint64       `json:"dailyOut,omitempty"`
	WeeklyOut       Uint64       `json:"weeklyOut,omitempty"`
	MonthlyOut      Uint64       `json:"monthlyOut,omitempty"`
	AnnualOut       Uint64       `json:"annualOut,omitempty"`
}

// ManageLimitsOpDetails is an XDR NestedUnion defines as:
//
//   union switch (ManageLimitsAction action)
//        {
//        case CREATE:
//            LimitsCreateDetails limitsCreateDetails;
//        case REMOVE:
//            uint64 id;
//        }
//
type ManageLimitsOpDetails struct {
	Action              ManageLimitsAction   `json:"action,omitempty"`
	LimitsCreateDetails *LimitsCreateDetails `json:"limitsCreateDetails,omitempty"`
	Id                  *Uint64              `json:"id,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageLimitsOpDetails) SwitchFieldName() string {
	return "Action"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageLimitsOpDetails
func (u ManageLimitsOpDetails) ArmForSwitch(sw int32) (string, bool) {
	switch ManageLimitsAction(sw) {
	case ManageLimitsActionCreate:
		return "LimitsCreateDetails", true
	case ManageLimitsActionRemove:
		return "Id", true
	}
	return "-", false
}

// NewManageLimitsOpDetails creates a new  ManageLimitsOpDetails.
func NewManageLimitsOpDetails(action ManageLimitsAction, value interface{}) (result ManageLimitsOpDetails, err error) {
	result.Action = action
	switch ManageLimitsAction(action) {
	case ManageLimitsActionCreate:
		tv, ok := value.(LimitsCreateDetails)
		if !ok {
			err = fmt.Errorf("invalid value, must be LimitsCreateDetails")
			return
		}
		result.LimitsCreateDetails = &tv
	case ManageLimitsActionRemove:
		tv, ok := value.(Uint64)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint64")
			return
		}
		result.Id = &tv
	}
	return
}

// MustLimitsCreateDetails retrieves the LimitsCreateDetails value from the union,
// panicing if the value is not set.
func (u ManageLimitsOpDetails) MustLimitsCreateDetails() LimitsCreateDetails {
	val, ok := u.GetLimitsCreateDetails()

	if !ok {
		panic("arm LimitsCreateDetails is not set")
	}

	return val
}

// GetLimitsCreateDetails retrieves the LimitsCreateDetails value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageLimitsOpDetails) GetLimitsCreateDetails() (result LimitsCreateDetails, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Action))

	if armName == "LimitsCreateDetails" {
		result = *u.LimitsCreateDetails
		ok = true
	}

	return
}

// MustId retrieves the Id value from the union,
// panicing if the value is not set.
func (u ManageLimitsOpDetails) MustId() Uint64 {
	val, ok := u.GetId()

	if !ok {
		panic("arm Id is not set")
	}

	return val
}

// GetId retrieves the Id value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageLimitsOpDetails) GetId() (result Uint64, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Action))

	if armName == "Id" {
		result = *u.Id
		ok = true
	}

	return
}

// ManageLimitsOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type ManageLimitsOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageLimitsOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageLimitsOpExt
func (u ManageLimitsOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewManageLimitsOpExt creates a new  ManageLimitsOpExt.
func NewManageLimitsOpExt(v LedgerVersion, value interface{}) (result ManageLimitsOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ManageLimitsOp is an XDR Struct defines as:
//
//   struct ManageLimitsOp
//    {
//        union switch (ManageLimitsAction action)
//        {
//        case CREATE:
//            LimitsCreateDetails limitsCreateDetails;
//        case REMOVE:
//            uint64 id;
//        } details;
//
//         // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type ManageLimitsOp struct {
	Details ManageLimitsOpDetails `json:"details,omitempty"`
	Ext     ManageLimitsOpExt     `json:"ext,omitempty"`
}

// ManageLimitsResultCode is an XDR Enum defines as:
//
//   enum ManageLimitsResultCode
//    {
//        // codes considered as "success" for the operation
//        SUCCESS = 0,
//
//        // codes considered as "failure" for the operation
//        MALFORMED = -1,
//        NOT_FOUND = -2,
//        ALREADY_EXISTS = -3,
//        CANNOT_CREATE_FOR_ACC_ID_AND_ACC_TYPE = -4, // limits cannot be created for account ID and account type simultaneously
//        INVALID_LIMITS = -5
//    };
//
type ManageLimitsResultCode int32

const (
	ManageLimitsResultCodeSuccess                        ManageLimitsResultCode = 0
	ManageLimitsResultCodeMalformed                      ManageLimitsResultCode = -1
	ManageLimitsResultCodeNotFound                       ManageLimitsResultCode = -2
	ManageLimitsResultCodeAlreadyExists                  ManageLimitsResultCode = -3
	ManageLimitsResultCodeCannotCreateForAccIdAndAccType ManageLimitsResultCode = -4
	ManageLimitsResultCodeInvalidLimits                  ManageLimitsResultCode = -5
)

var ManageLimitsResultCodeAll = []ManageLimitsResultCode{
	ManageLimitsResultCodeSuccess,
	ManageLimitsResultCodeMalformed,
	ManageLimitsResultCodeNotFound,
	ManageLimitsResultCodeAlreadyExists,
	ManageLimitsResultCodeCannotCreateForAccIdAndAccType,
	ManageLimitsResultCodeInvalidLimits,
}

var manageLimitsResultCodeMap = map[int32]string{
	0:  "ManageLimitsResultCodeSuccess",
	-1: "ManageLimitsResultCodeMalformed",
	-2: "ManageLimitsResultCodeNotFound",
	-3: "ManageLimitsResultCodeAlreadyExists",
	-4: "ManageLimitsResultCodeCannotCreateForAccIdAndAccType",
	-5: "ManageLimitsResultCodeInvalidLimits",
}

var manageLimitsResultCodeShortMap = map[int32]string{
	0:  "success",
	-1: "malformed",
	-2: "not_found",
	-3: "already_exists",
	-4: "cannot_create_for_acc_id_and_acc_type",
	-5: "invalid_limits",
}

var manageLimitsResultCodeRevMap = map[string]int32{
	"ManageLimitsResultCodeSuccess":                        0,
	"ManageLimitsResultCodeMalformed":                      -1,
	"ManageLimitsResultCodeNotFound":                       -2,
	"ManageLimitsResultCodeAlreadyExists":                  -3,
	"ManageLimitsResultCodeCannotCreateForAccIdAndAccType": -4,
	"ManageLimitsResultCodeInvalidLimits":                  -5,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ManageLimitsResultCode
func (e ManageLimitsResultCode) ValidEnum(v int32) bool {
	_, ok := manageLimitsResultCodeMap[v]
	return ok
}
func (e ManageLimitsResultCode) isFlag() bool {
	for i := len(ManageLimitsResultCodeAll) - 1; i >= 0; i-- {
		expected := ManageLimitsResultCode(2) << uint64(len(ManageLimitsResultCodeAll)-1) >> uint64(len(ManageLimitsResultCodeAll)-i)
		if expected != ManageLimitsResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ManageLimitsResultCode) String() string {
	name, _ := manageLimitsResultCodeMap[int32(e)]
	return name
}

func (e ManageLimitsResultCode) ShortString() string {
	name, _ := manageLimitsResultCodeShortMap[int32(e)]
	return name
}

func (e ManageLimitsResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range ManageLimitsResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ManageLimitsResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ManageLimitsResultCode(t.Value)
	return nil
}

// ManageLimitsResultSuccessDetails is an XDR NestedUnion defines as:
//
//   union switch (ManageLimitsAction action)
//            {
//            case CREATE:
//                uint64 id;
//            case REMOVE:
//                void;
//            }
//
type ManageLimitsResultSuccessDetails struct {
	Action ManageLimitsAction `json:"action,omitempty"`
	Id     *Uint64            `json:"id,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageLimitsResultSuccessDetails) SwitchFieldName() string {
	return "Action"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageLimitsResultSuccessDetails
func (u ManageLimitsResultSuccessDetails) ArmForSwitch(sw int32) (string, bool) {
	switch ManageLimitsAction(sw) {
	case ManageLimitsActionCreate:
		return "Id", true
	case ManageLimitsActionRemove:
		return "", true
	}
	return "-", false
}

// NewManageLimitsResultSuccessDetails creates a new  ManageLimitsResultSuccessDetails.
func NewManageLimitsResultSuccessDetails(action ManageLimitsAction, value interface{}) (result ManageLimitsResultSuccessDetails, err error) {
	result.Action = action
	switch ManageLimitsAction(action) {
	case ManageLimitsActionCreate:
		tv, ok := value.(Uint64)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint64")
			return
		}
		result.Id = &tv
	case ManageLimitsActionRemove:
		// void
	}
	return
}

// MustId retrieves the Id value from the union,
// panicing if the value is not set.
func (u ManageLimitsResultSuccessDetails) MustId() Uint64 {
	val, ok := u.GetId()

	if !ok {
		panic("arm Id is not set")
	}

	return val
}

// GetId retrieves the Id value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageLimitsResultSuccessDetails) GetId() (result Uint64, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Action))

	if armName == "Id" {
		result = *u.Id
		ok = true
	}

	return
}

// ManageLimitsResultSuccessExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//
type ManageLimitsResultSuccessExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageLimitsResultSuccessExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageLimitsResultSuccessExt
func (u ManageLimitsResultSuccessExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewManageLimitsResultSuccessExt creates a new  ManageLimitsResultSuccessExt.
func NewManageLimitsResultSuccessExt(v LedgerVersion, value interface{}) (result ManageLimitsResultSuccessExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ManageLimitsResultSuccess is an XDR NestedStruct defines as:
//
//   struct {
//            union switch (ManageLimitsAction action)
//            {
//            case CREATE:
//                uint64 id;
//            case REMOVE:
//                void;
//            } details;
//
//    		// reserved for future use
//    		union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//    		ext;
//    	}
//
type ManageLimitsResultSuccess struct {
	Details ManageLimitsResultSuccessDetails `json:"details,omitempty"`
	Ext     ManageLimitsResultSuccessExt     `json:"ext,omitempty"`
}

// ManageLimitsResult is an XDR Union defines as:
//
//   union ManageLimitsResult switch (ManageLimitsResultCode code)
//    {
//    case SUCCESS:
//        struct {
//            union switch (ManageLimitsAction action)
//            {
//            case CREATE:
//                uint64 id;
//            case REMOVE:
//                void;
//            } details;
//
//    		// reserved for future use
//    		union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//    		ext;
//    	} success;
//    default:
//        void;
//    };
//
type ManageLimitsResult struct {
	Code    ManageLimitsResultCode     `json:"code,omitempty"`
	Success *ManageLimitsResultSuccess `json:"success,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageLimitsResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageLimitsResult
func (u ManageLimitsResult) ArmForSwitch(sw int32) (string, bool) {
	switch ManageLimitsResultCode(sw) {
	case ManageLimitsResultCodeSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewManageLimitsResult creates a new  ManageLimitsResult.
func NewManageLimitsResult(code ManageLimitsResultCode, value interface{}) (result ManageLimitsResult, err error) {
	result.Code = code
	switch ManageLimitsResultCode(code) {
	case ManageLimitsResultCodeSuccess:
		tv, ok := value.(ManageLimitsResultSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageLimitsResultSuccess")
			return
		}
		result.Success = &tv
	default:
		// void
	}
	return
}

// MustSuccess retrieves the Success value from the union,
// panicing if the value is not set.
func (u ManageLimitsResult) MustSuccess() ManageLimitsResultSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageLimitsResult) GetSuccess() (result ManageLimitsResultSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// ManageOfferOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type ManageOfferOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageOfferOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageOfferOpExt
func (u ManageOfferOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewManageOfferOpExt creates a new  ManageOfferOpExt.
func NewManageOfferOpExt(v LedgerVersion, value interface{}) (result ManageOfferOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ManageOfferOp is an XDR Struct defines as:
//
//   struct ManageOfferOp
//    {
//        BalanceID baseBalance; // balance for base asset
//    	BalanceID quoteBalance; // balance for quote asset
//    	bool isBuy;
//        int64 amount; // if set to 0, delete the offer
//        int64 price;  // price of base asset in terms of quote
//
//        int64 fee;
//
//        // 0=create a new offer, otherwise edit an existing offer
//        uint64 offerID;
//    	uint64 orderBookID;
//    	// reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type ManageOfferOp struct {
	BaseBalance  BalanceId        `json:"baseBalance,omitempty"`
	QuoteBalance BalanceId        `json:"quoteBalance,omitempty"`
	IsBuy        bool             `json:"isBuy,omitempty"`
	Amount       Int64            `json:"amount,omitempty"`
	Price        Int64            `json:"price,omitempty"`
	Fee          Int64            `json:"fee,omitempty"`
	OfferId      Uint64           `json:"offerID,omitempty"`
	OrderBookId  Uint64           `json:"orderBookID,omitempty"`
	Ext          ManageOfferOpExt `json:"ext,omitempty"`
}

// ManageOfferResultCode is an XDR Enum defines as:
//
//   enum ManageOfferResultCode
//    {
//        // codes considered as "success" for the operation
//        SUCCESS = 0,
//
//        // codes considered as "failure" for the operation
//        MALFORMED = -1,     // generated offer would be invalid
//        PAIR_NOT_TRADED = -2, // it's not allowed to trage with this pair
//        BALANCE_NOT_FOUND = -3,  // does not own balance for buying or selling
//        UNDERFUNDED = -4,    // doesn't hold what it's trying to sell
//        CROSS_SELF = -5,     // would cross an offer from the same user
//    	OFFER_OVERFLOW = -6,
//    	ASSET_PAIR_NOT_TRADABLE = -7,
//    	PHYSICAL_PRICE_RESTRICTION = -8, // offer price violates physical price restriction
//    	CURRENT_PRICE_RESTRICTION = -9,
//        NOT_FOUND = -10, // offerID does not match an existing offer
//        INVALID_PERCENT_FEE = -11,
//    	INSUFFICIENT_PRICE = -12,
//    	ORDER_BOOK_DOES_NOT_EXISTS = -13, // specified order book does not exists
//    	SALE_IS_NOT_STARTED_YET = -14, // sale is not started yet
//    	SALE_ALREADY_ENDED = -15, // sale has already ended
//    	ORDER_VIOLATES_HARD_CAP = -16, // currentcap + order will exceed hard cap
//    	CANT_PARTICIPATE_OWN_SALE = -17, // it's not allowed to participate in own sale
//    	ASSET_MISMATCHED = -18, // sale assets does not match assets for specified balances
//    	PRICE_DOES_NOT_MATCH = -19, // price does not match sale price
//    	PRICE_IS_INVALID = -20, // price must be positive
//    	UPDATE_IS_NOT_ALLOWED = -21, // update of the offer is not allowed
//    	INVALID_AMOUNT = -22, // amount must be positive
//    	SALE_IS_NOT_ACTIVE = -23,
//    	REQUIRES_KYC = -24, // source must have KYC in order to participate
//    	SOURCE_UNDERFUNDED = -25,
//    	SOURCE_BALANCE_LOCK_OVERFLOW = -26,
//    	REQUIRES_VERIFICATION = -27 // source must be verified in order to participate
//    };
//
type ManageOfferResultCode int32

const (
	ManageOfferResultCodeSuccess                   ManageOfferResultCode = 0
	ManageOfferResultCodeMalformed                 ManageOfferResultCode = -1
	ManageOfferResultCodePairNotTraded             ManageOfferResultCode = -2
	ManageOfferResultCodeBalanceNotFound           ManageOfferResultCode = -3
	ManageOfferResultCodeUnderfunded               ManageOfferResultCode = -4
	ManageOfferResultCodeCrossSelf                 ManageOfferResultCode = -5
	ManageOfferResultCodeOfferOverflow             ManageOfferResultCode = -6
	ManageOfferResultCodeAssetPairNotTradable      ManageOfferResultCode = -7
	ManageOfferResultCodePhysicalPriceRestriction  ManageOfferResultCode = -8
	ManageOfferResultCodeCurrentPriceRestriction   ManageOfferResultCode = -9
	ManageOfferResultCodeNotFound                  ManageOfferResultCode = -10
	ManageOfferResultCodeInvalidPercentFee         ManageOfferResultCode = -11
	ManageOfferResultCodeInsufficientPrice         ManageOfferResultCode = -12
	ManageOfferResultCodeOrderBookDoesNotExists    ManageOfferResultCode = -13
	ManageOfferResultCodeSaleIsNotStartedYet       ManageOfferResultCode = -14
	ManageOfferResultCodeSaleAlreadyEnded          ManageOfferResultCode = -15
	ManageOfferResultCodeOrderViolatesHardCap      ManageOfferResultCode = -16
	ManageOfferResultCodeCantParticipateOwnSale    ManageOfferResultCode = -17
	ManageOfferResultCodeAssetMismatched           ManageOfferResultCode = -18
	ManageOfferResultCodePriceDoesNotMatch         ManageOfferResultCode = -19
	ManageOfferResultCodePriceIsInvalid            ManageOfferResultCode = -20
	ManageOfferResultCodeUpdateIsNotAllowed        ManageOfferResultCode = -21
	ManageOfferResultCodeInvalidAmount             ManageOfferResultCode = -22
	ManageOfferResultCodeSaleIsNotActive           ManageOfferResultCode = -23
	ManageOfferResultCodeRequiresKyc               ManageOfferResultCode = -24
	ManageOfferResultCodeSourceUnderfunded         ManageOfferResultCode = -25
	ManageOfferResultCodeSourceBalanceLockOverflow ManageOfferResultCode = -26
	ManageOfferResultCodeRequiresVerification      ManageOfferResultCode = -27
)

var ManageOfferResultCodeAll = []ManageOfferResultCode{
	ManageOfferResultCodeSuccess,
	ManageOfferResultCodeMalformed,
	ManageOfferResultCodePairNotTraded,
	ManageOfferResultCodeBalanceNotFound,
	ManageOfferResultCodeUnderfunded,
	ManageOfferResultCodeCrossSelf,
	ManageOfferResultCodeOfferOverflow,
	ManageOfferResultCodeAssetPairNotTradable,
	ManageOfferResultCodePhysicalPriceRestriction,
	ManageOfferResultCodeCurrentPriceRestriction,
	ManageOfferResultCodeNotFound,
	ManageOfferResultCodeInvalidPercentFee,
	ManageOfferResultCodeInsufficientPrice,
	ManageOfferResultCodeOrderBookDoesNotExists,
	ManageOfferResultCodeSaleIsNotStartedYet,
	ManageOfferResultCodeSaleAlreadyEnded,
	ManageOfferResultCodeOrderViolatesHardCap,
	ManageOfferResultCodeCantParticipateOwnSale,
	ManageOfferResultCodeAssetMismatched,
	ManageOfferResultCodePriceDoesNotMatch,
	ManageOfferResultCodePriceIsInvalid,
	ManageOfferResultCodeUpdateIsNotAllowed,
	ManageOfferResultCodeInvalidAmount,
	ManageOfferResultCodeSaleIsNotActive,
	ManageOfferResultCodeRequiresKyc,
	ManageOfferResultCodeSourceUnderfunded,
	ManageOfferResultCodeSourceBalanceLockOverflow,
	ManageOfferResultCodeRequiresVerification,
}

var manageOfferResultCodeMap = map[int32]string{
	0:   "ManageOfferResultCodeSuccess",
	-1:  "ManageOfferResultCodeMalformed",
	-2:  "ManageOfferResultCodePairNotTraded",
	-3:  "ManageOfferResultCodeBalanceNotFound",
	-4:  "ManageOfferResultCodeUnderfunded",
	-5:  "ManageOfferResultCodeCrossSelf",
	-6:  "ManageOfferResultCodeOfferOverflow",
	-7:  "ManageOfferResultCodeAssetPairNotTradable",
	-8:  "ManageOfferResultCodePhysicalPriceRestriction",
	-9:  "ManageOfferResultCodeCurrentPriceRestriction",
	-10: "ManageOfferResultCodeNotFound",
	-11: "ManageOfferResultCodeInvalidPercentFee",
	-12: "ManageOfferResultCodeInsufficientPrice",
	-13: "ManageOfferResultCodeOrderBookDoesNotExists",
	-14: "ManageOfferResultCodeSaleIsNotStartedYet",
	-15: "ManageOfferResultCodeSaleAlreadyEnded",
	-16: "ManageOfferResultCodeOrderViolatesHardCap",
	-17: "ManageOfferResultCodeCantParticipateOwnSale",
	-18: "ManageOfferResultCodeAssetMismatched",
	-19: "ManageOfferResultCodePriceDoesNotMatch",
	-20: "ManageOfferResultCodePriceIsInvalid",
	-21: "ManageOfferResultCodeUpdateIsNotAllowed",
	-22: "ManageOfferResultCodeInvalidAmount",
	-23: "ManageOfferResultCodeSaleIsNotActive",
	-24: "ManageOfferResultCodeRequiresKyc",
	-25: "ManageOfferResultCodeSourceUnderfunded",
	-26: "ManageOfferResultCodeSourceBalanceLockOverflow",
	-27: "ManageOfferResultCodeRequiresVerification",
}

var manageOfferResultCodeShortMap = map[int32]string{
	0:   "success",
	-1:  "malformed",
	-2:  "pair_not_traded",
	-3:  "balance_not_found",
	-4:  "underfunded",
	-5:  "cross_self",
	-6:  "offer_overflow",
	-7:  "asset_pair_not_tradable",
	-8:  "physical_price_restriction",
	-9:  "current_price_restriction",
	-10: "not_found",
	-11: "invalid_percent_fee",
	-12: "insufficient_price",
	-13: "order_book_does_not_exists",
	-14: "sale_is_not_started_yet",
	-15: "sale_already_ended",
	-16: "order_violates_hard_cap",
	-17: "cant_participate_own_sale",
	-18: "asset_mismatched",
	-19: "price_does_not_match",
	-20: "price_is_invalid",
	-21: "update_is_not_allowed",
	-22: "invalid_amount",
	-23: "sale_is_not_active",
	-24: "requires_kyc",
	-25: "source_underfunded",
	-26: "source_balance_lock_overflow",
	-27: "requires_verification",
}

var manageOfferResultCodeRevMap = map[string]int32{
	"ManageOfferResultCodeSuccess":                   0,
	"ManageOfferResultCodeMalformed":                 -1,
	"ManageOfferResultCodePairNotTraded":             -2,
	"ManageOfferResultCodeBalanceNotFound":           -3,
	"ManageOfferResultCodeUnderfunded":               -4,
	"ManageOfferResultCodeCrossSelf":                 -5,
	"ManageOfferResultCodeOfferOverflow":             -6,
	"ManageOfferResultCodeAssetPairNotTradable":      -7,
	"ManageOfferResultCodePhysicalPriceRestriction":  -8,
	"ManageOfferResultCodeCurrentPriceRestriction":   -9,
	"ManageOfferResultCodeNotFound":                  -10,
	"ManageOfferResultCodeInvalidPercentFee":         -11,
	"ManageOfferResultCodeInsufficientPrice":         -12,
	"ManageOfferResultCodeOrderBookDoesNotExists":    -13,
	"ManageOfferResultCodeSaleIsNotStartedYet":       -14,
	"ManageOfferResultCodeSaleAlreadyEnded":          -15,
	"ManageOfferResultCodeOrderViolatesHardCap":      -16,
	"ManageOfferResultCodeCantParticipateOwnSale":    -17,
	"ManageOfferResultCodeAssetMismatched":           -18,
	"ManageOfferResultCodePriceDoesNotMatch":         -19,
	"ManageOfferResultCodePriceIsInvalid":            -20,
	"ManageOfferResultCodeUpdateIsNotAllowed":        -21,
	"ManageOfferResultCodeInvalidAmount":             -22,
	"ManageOfferResultCodeSaleIsNotActive":           -23,
	"ManageOfferResultCodeRequiresKyc":               -24,
	"ManageOfferResultCodeSourceUnderfunded":         -25,
	"ManageOfferResultCodeSourceBalanceLockOverflow": -26,
	"ManageOfferResultCodeRequiresVerification":      -27,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ManageOfferResultCode
func (e ManageOfferResultCode) ValidEnum(v int32) bool {
	_, ok := manageOfferResultCodeMap[v]
	return ok
}
func (e ManageOfferResultCode) isFlag() bool {
	for i := len(ManageOfferResultCodeAll) - 1; i >= 0; i-- {
		expected := ManageOfferResultCode(2) << uint64(len(ManageOfferResultCodeAll)-1) >> uint64(len(ManageOfferResultCodeAll)-i)
		if expected != ManageOfferResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ManageOfferResultCode) String() string {
	name, _ := manageOfferResultCodeMap[int32(e)]
	return name
}

func (e ManageOfferResultCode) ShortString() string {
	name, _ := manageOfferResultCodeShortMap[int32(e)]
	return name
}

func (e ManageOfferResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range ManageOfferResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ManageOfferResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ManageOfferResultCode(t.Value)
	return nil
}

// ManageOfferEffect is an XDR Enum defines as:
//
//   enum ManageOfferEffect
//    {
//        CREATED = 0,
//        UPDATED = 1,
//        DELETED = 2
//    };
//
type ManageOfferEffect int32

const (
	ManageOfferEffectCreated ManageOfferEffect = 0
	ManageOfferEffectUpdated ManageOfferEffect = 1
	ManageOfferEffectDeleted ManageOfferEffect = 2
)

var ManageOfferEffectAll = []ManageOfferEffect{
	ManageOfferEffectCreated,
	ManageOfferEffectUpdated,
	ManageOfferEffectDeleted,
}

var manageOfferEffectMap = map[int32]string{
	0: "ManageOfferEffectCreated",
	1: "ManageOfferEffectUpdated",
	2: "ManageOfferEffectDeleted",
}

var manageOfferEffectShortMap = map[int32]string{
	0: "created",
	1: "updated",
	2: "deleted",
}

var manageOfferEffectRevMap = map[string]int32{
	"ManageOfferEffectCreated": 0,
	"ManageOfferEffectUpdated": 1,
	"ManageOfferEffectDeleted": 2,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ManageOfferEffect
func (e ManageOfferEffect) ValidEnum(v int32) bool {
	_, ok := manageOfferEffectMap[v]
	return ok
}
func (e ManageOfferEffect) isFlag() bool {
	for i := len(ManageOfferEffectAll) - 1; i >= 0; i-- {
		expected := ManageOfferEffect(2) << uint64(len(ManageOfferEffectAll)-1) >> uint64(len(ManageOfferEffectAll)-i)
		if expected != ManageOfferEffectAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ManageOfferEffect) String() string {
	name, _ := manageOfferEffectMap[int32(e)]
	return name
}

func (e ManageOfferEffect) ShortString() string {
	name, _ := manageOfferEffectShortMap[int32(e)]
	return name
}

func (e ManageOfferEffect) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range ManageOfferEffectAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ManageOfferEffect) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ManageOfferEffect(t.Value)
	return nil
}

// ClaimOfferAtomExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type ClaimOfferAtomExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ClaimOfferAtomExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ClaimOfferAtomExt
func (u ClaimOfferAtomExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewClaimOfferAtomExt creates a new  ClaimOfferAtomExt.
func NewClaimOfferAtomExt(v LedgerVersion, value interface{}) (result ClaimOfferAtomExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ClaimOfferAtom is an XDR Struct defines as:
//
//   struct ClaimOfferAtom
//    {
//        // emitted to identify the offer
//        AccountID bAccountID; // Account that owns the offer
//        uint64 offerID;
//    	int64 baseAmount;
//    	int64 quoteAmount;
//    	int64 bFeePaid;
//    	int64 aFeePaid;
//    	BalanceID baseBalance;
//    	BalanceID quoteBalance;
//
//    	int64 currentPrice;
//
//    	union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type ClaimOfferAtom struct {
	BAccountId   AccountId         `json:"bAccountID,omitempty"`
	OfferId      Uint64            `json:"offerID,omitempty"`
	BaseAmount   Int64             `json:"baseAmount,omitempty"`
	QuoteAmount  Int64             `json:"quoteAmount,omitempty"`
	BFeePaid     Int64             `json:"bFeePaid,omitempty"`
	AFeePaid     Int64             `json:"aFeePaid,omitempty"`
	BaseBalance  BalanceId         `json:"baseBalance,omitempty"`
	QuoteBalance BalanceId         `json:"quoteBalance,omitempty"`
	CurrentPrice Int64             `json:"currentPrice,omitempty"`
	Ext          ClaimOfferAtomExt `json:"ext,omitempty"`
}

// ManageOfferSuccessResultOffer is an XDR NestedUnion defines as:
//
//   union switch (ManageOfferEffect effect)
//        {
//        case CREATED:
//        case UPDATED:
//            OfferEntry offer;
//        default:
//            void;
//        }
//
type ManageOfferSuccessResultOffer struct {
	Effect ManageOfferEffect `json:"effect,omitempty"`
	Offer  *OfferEntry       `json:"offer,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageOfferSuccessResultOffer) SwitchFieldName() string {
	return "Effect"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageOfferSuccessResultOffer
func (u ManageOfferSuccessResultOffer) ArmForSwitch(sw int32) (string, bool) {
	switch ManageOfferEffect(sw) {
	case ManageOfferEffectCreated:
		return "Offer", true
	case ManageOfferEffectUpdated:
		return "Offer", true
	default:
		return "", true
	}
}

// NewManageOfferSuccessResultOffer creates a new  ManageOfferSuccessResultOffer.
func NewManageOfferSuccessResultOffer(effect ManageOfferEffect, value interface{}) (result ManageOfferSuccessResultOffer, err error) {
	result.Effect = effect
	switch ManageOfferEffect(effect) {
	case ManageOfferEffectCreated:
		tv, ok := value.(OfferEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be OfferEntry")
			return
		}
		result.Offer = &tv
	case ManageOfferEffectUpdated:
		tv, ok := value.(OfferEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be OfferEntry")
			return
		}
		result.Offer = &tv
	default:
		// void
	}
	return
}

// MustOffer retrieves the Offer value from the union,
// panicing if the value is not set.
func (u ManageOfferSuccessResultOffer) MustOffer() OfferEntry {
	val, ok := u.GetOffer()

	if !ok {
		panic("arm Offer is not set")
	}

	return val
}

// GetOffer retrieves the Offer value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageOfferSuccessResultOffer) GetOffer() (result OfferEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Effect))

	if armName == "Offer" {
		result = *u.Offer
		ok = true
	}

	return
}

// ManageOfferSuccessResultExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type ManageOfferSuccessResultExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageOfferSuccessResultExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageOfferSuccessResultExt
func (u ManageOfferSuccessResultExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewManageOfferSuccessResultExt creates a new  ManageOfferSuccessResultExt.
func NewManageOfferSuccessResultExt(v LedgerVersion, value interface{}) (result ManageOfferSuccessResultExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ManageOfferSuccessResult is an XDR Struct defines as:
//
//   struct ManageOfferSuccessResult
//    {
//
//        // offers that got claimed while creating this offer
//        ClaimOfferAtom offersClaimed<>;
//    	AssetCode baseAsset;
//    	AssetCode quoteAsset;
//
//        union switch (ManageOfferEffect effect)
//        {
//        case CREATED:
//        case UPDATED:
//            OfferEntry offer;
//        default:
//            void;
//        }
//        offer;
//
//    	union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type ManageOfferSuccessResult struct {
	OffersClaimed []ClaimOfferAtom              `json:"offersClaimed,omitempty"`
	BaseAsset     AssetCode                     `json:"baseAsset,omitempty"`
	QuoteAsset    AssetCode                     `json:"quoteAsset,omitempty"`
	Offer         ManageOfferSuccessResultOffer `json:"offer,omitempty"`
	Ext           ManageOfferSuccessResultExt   `json:"ext,omitempty"`
}

// ManageOfferResultPhysicalPriceRestrictionExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//
type ManageOfferResultPhysicalPriceRestrictionExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageOfferResultPhysicalPriceRestrictionExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageOfferResultPhysicalPriceRestrictionExt
func (u ManageOfferResultPhysicalPriceRestrictionExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewManageOfferResultPhysicalPriceRestrictionExt creates a new  ManageOfferResultPhysicalPriceRestrictionExt.
func NewManageOfferResultPhysicalPriceRestrictionExt(v LedgerVersion, value interface{}) (result ManageOfferResultPhysicalPriceRestrictionExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ManageOfferResultPhysicalPriceRestriction is an XDR NestedStruct defines as:
//
//   struct {
//    		int64 physicalPrice;
//    		union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//    		ext;
//    	}
//
type ManageOfferResultPhysicalPriceRestriction struct {
	PhysicalPrice Int64                                        `json:"physicalPrice,omitempty"`
	Ext           ManageOfferResultPhysicalPriceRestrictionExt `json:"ext,omitempty"`
}

// ManageOfferResultCurrentPriceRestrictionExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//
type ManageOfferResultCurrentPriceRestrictionExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageOfferResultCurrentPriceRestrictionExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageOfferResultCurrentPriceRestrictionExt
func (u ManageOfferResultCurrentPriceRestrictionExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewManageOfferResultCurrentPriceRestrictionExt creates a new  ManageOfferResultCurrentPriceRestrictionExt.
func NewManageOfferResultCurrentPriceRestrictionExt(v LedgerVersion, value interface{}) (result ManageOfferResultCurrentPriceRestrictionExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ManageOfferResultCurrentPriceRestriction is an XDR NestedStruct defines as:
//
//   struct {
//    		int64 currentPrice;
//    		union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//    		ext;
//    	}
//
type ManageOfferResultCurrentPriceRestriction struct {
	CurrentPrice Int64                                       `json:"currentPrice,omitempty"`
	Ext          ManageOfferResultCurrentPriceRestrictionExt `json:"ext,omitempty"`
}

// ManageOfferResult is an XDR Union defines as:
//
//   union ManageOfferResult switch (ManageOfferResultCode code)
//    {
//    case SUCCESS:
//        ManageOfferSuccessResult success;
//    case PHYSICAL_PRICE_RESTRICTION:
//    	struct {
//    		int64 physicalPrice;
//    		union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//    		ext;
//    	} physicalPriceRestriction;
//    case CURRENT_PRICE_RESTRICTION:
//    	struct {
//    		int64 currentPrice;
//    		union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//    		ext;
//    	} currentPriceRestriction;
//
//    default:
//        void;
//    };
//
type ManageOfferResult struct {
	Code                     ManageOfferResultCode                      `json:"code,omitempty"`
	Success                  *ManageOfferSuccessResult                  `json:"success,omitempty"`
	PhysicalPriceRestriction *ManageOfferResultPhysicalPriceRestriction `json:"physicalPriceRestriction,omitempty"`
	CurrentPriceRestriction  *ManageOfferResultCurrentPriceRestriction  `json:"currentPriceRestriction,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageOfferResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageOfferResult
func (u ManageOfferResult) ArmForSwitch(sw int32) (string, bool) {
	switch ManageOfferResultCode(sw) {
	case ManageOfferResultCodeSuccess:
		return "Success", true
	case ManageOfferResultCodePhysicalPriceRestriction:
		return "PhysicalPriceRestriction", true
	case ManageOfferResultCodeCurrentPriceRestriction:
		return "CurrentPriceRestriction", true
	default:
		return "", true
	}
}

// NewManageOfferResult creates a new  ManageOfferResult.
func NewManageOfferResult(code ManageOfferResultCode, value interface{}) (result ManageOfferResult, err error) {
	result.Code = code
	switch ManageOfferResultCode(code) {
	case ManageOfferResultCodeSuccess:
		tv, ok := value.(ManageOfferSuccessResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageOfferSuccessResult")
			return
		}
		result.Success = &tv
	case ManageOfferResultCodePhysicalPriceRestriction:
		tv, ok := value.(ManageOfferResultPhysicalPriceRestriction)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageOfferResultPhysicalPriceRestriction")
			return
		}
		result.PhysicalPriceRestriction = &tv
	case ManageOfferResultCodeCurrentPriceRestriction:
		tv, ok := value.(ManageOfferResultCurrentPriceRestriction)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageOfferResultCurrentPriceRestriction")
			return
		}
		result.CurrentPriceRestriction = &tv
	default:
		// void
	}
	return
}

// MustSuccess retrieves the Success value from the union,
// panicing if the value is not set.
func (u ManageOfferResult) MustSuccess() ManageOfferSuccessResult {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageOfferResult) GetSuccess() (result ManageOfferSuccessResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// MustPhysicalPriceRestriction retrieves the PhysicalPriceRestriction value from the union,
// panicing if the value is not set.
func (u ManageOfferResult) MustPhysicalPriceRestriction() ManageOfferResultPhysicalPriceRestriction {
	val, ok := u.GetPhysicalPriceRestriction()

	if !ok {
		panic("arm PhysicalPriceRestriction is not set")
	}

	return val
}

// GetPhysicalPriceRestriction retrieves the PhysicalPriceRestriction value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageOfferResult) GetPhysicalPriceRestriction() (result ManageOfferResultPhysicalPriceRestriction, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "PhysicalPriceRestriction" {
		result = *u.PhysicalPriceRestriction
		ok = true
	}

	return
}

// MustCurrentPriceRestriction retrieves the CurrentPriceRestriction value from the union,
// panicing if the value is not set.
func (u ManageOfferResult) MustCurrentPriceRestriction() ManageOfferResultCurrentPriceRestriction {
	val, ok := u.GetCurrentPriceRestriction()

	if !ok {
		panic("arm CurrentPriceRestriction is not set")
	}

	return val
}

// GetCurrentPriceRestriction retrieves the CurrentPriceRestriction value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageOfferResult) GetCurrentPriceRestriction() (result ManageOfferResultCurrentPriceRestriction, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "CurrentPriceRestriction" {
		result = *u.CurrentPriceRestriction
		ok = true
	}

	return
}

// ManageSaleAction is an XDR Enum defines as:
//
//   enum ManageSaleAction
//    {
//        CREATE_UPDATE_DETAILS_REQUEST = 1,
//        CANCEL = 2,
//    	SET_STATE = 3,
//    	CREATE_PROMOTION_UPDATE_REQUEST = 4,
//    	CREATE_UPDATE_END_TIME_REQUEST = 5
//    };
//
type ManageSaleAction int32

const (
	ManageSaleActionCreateUpdateDetailsRequest   ManageSaleAction = 1
	ManageSaleActionCancel                       ManageSaleAction = 2
	ManageSaleActionSetState                     ManageSaleAction = 3
	ManageSaleActionCreatePromotionUpdateRequest ManageSaleAction = 4
	ManageSaleActionCreateUpdateEndTimeRequest   ManageSaleAction = 5
)

var ManageSaleActionAll = []ManageSaleAction{
	ManageSaleActionCreateUpdateDetailsRequest,
	ManageSaleActionCancel,
	ManageSaleActionSetState,
	ManageSaleActionCreatePromotionUpdateRequest,
	ManageSaleActionCreateUpdateEndTimeRequest,
}

var manageSaleActionMap = map[int32]string{
	1: "ManageSaleActionCreateUpdateDetailsRequest",
	2: "ManageSaleActionCancel",
	3: "ManageSaleActionSetState",
	4: "ManageSaleActionCreatePromotionUpdateRequest",
	5: "ManageSaleActionCreateUpdateEndTimeRequest",
}

var manageSaleActionShortMap = map[int32]string{
	1: "create_update_details_request",
	2: "cancel",
	3: "set_state",
	4: "create_promotion_update_request",
	5: "create_update_end_time_request",
}

var manageSaleActionRevMap = map[string]int32{
	"ManageSaleActionCreateUpdateDetailsRequest":   1,
	"ManageSaleActionCancel":                       2,
	"ManageSaleActionSetState":                     3,
	"ManageSaleActionCreatePromotionUpdateRequest": 4,
	"ManageSaleActionCreateUpdateEndTimeRequest":   5,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ManageSaleAction
func (e ManageSaleAction) ValidEnum(v int32) bool {
	_, ok := manageSaleActionMap[v]
	return ok
}
func (e ManageSaleAction) isFlag() bool {
	for i := len(ManageSaleActionAll) - 1; i >= 0; i-- {
		expected := ManageSaleAction(2) << uint64(len(ManageSaleActionAll)-1) >> uint64(len(ManageSaleActionAll)-i)
		if expected != ManageSaleActionAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ManageSaleAction) String() string {
	name, _ := manageSaleActionMap[int32(e)]
	return name
}

func (e ManageSaleAction) ShortString() string {
	name, _ := manageSaleActionShortMap[int32(e)]
	return name
}

func (e ManageSaleAction) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range ManageSaleActionAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ManageSaleAction) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ManageSaleAction(t.Value)
	return nil
}

// UpdateSaleDetailsDataExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type UpdateSaleDetailsDataExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u UpdateSaleDetailsDataExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of UpdateSaleDetailsDataExt
func (u UpdateSaleDetailsDataExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewUpdateSaleDetailsDataExt creates a new  UpdateSaleDetailsDataExt.
func NewUpdateSaleDetailsDataExt(v LedgerVersion, value interface{}) (result UpdateSaleDetailsDataExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// UpdateSaleDetailsData is an XDR Struct defines as:
//
//   struct UpdateSaleDetailsData {
//        uint64 requestID; // if requestID is 0 - create request, else - update
//        longstring newDetails;
//
//        // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        } ext;
//    };
//
type UpdateSaleDetailsData struct {
	RequestId  Uint64                   `json:"requestID,omitempty"`
	NewDetails Longstring               `json:"newDetails,omitempty"`
	Ext        UpdateSaleDetailsDataExt `json:"ext,omitempty"`
}

// PromotionUpdateDataExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type PromotionUpdateDataExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u PromotionUpdateDataExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of PromotionUpdateDataExt
func (u PromotionUpdateDataExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewPromotionUpdateDataExt creates a new  PromotionUpdateDataExt.
func NewPromotionUpdateDataExt(v LedgerVersion, value interface{}) (result PromotionUpdateDataExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// PromotionUpdateData is an XDR Struct defines as:
//
//   struct PromotionUpdateData {
//        uint64 requestID; // if requestID is 0 - create request, else - update
//        SaleCreationRequest newPromotionData;
//
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        } ext;
//    };
//
type PromotionUpdateData struct {
	RequestId        Uint64                 `json:"requestID,omitempty"`
	NewPromotionData SaleCreationRequest    `json:"newPromotionData,omitempty"`
	Ext              PromotionUpdateDataExt `json:"ext,omitempty"`
}

// UpdateSaleEndTimeDataExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type UpdateSaleEndTimeDataExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u UpdateSaleEndTimeDataExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of UpdateSaleEndTimeDataExt
func (u UpdateSaleEndTimeDataExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewUpdateSaleEndTimeDataExt creates a new  UpdateSaleEndTimeDataExt.
func NewUpdateSaleEndTimeDataExt(v LedgerVersion, value interface{}) (result UpdateSaleEndTimeDataExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// UpdateSaleEndTimeData is an XDR Struct defines as:
//
//   struct UpdateSaleEndTimeData {
//        uint64 requestID; // if requestID is 0 - create request, else - update
//        uint64 newEndTime;
//
//        // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        } ext;
//    };
//
type UpdateSaleEndTimeData struct {
	RequestId  Uint64                   `json:"requestID,omitempty"`
	NewEndTime Uint64                   `json:"newEndTime,omitempty"`
	Ext        UpdateSaleEndTimeDataExt `json:"ext,omitempty"`
}

// ManageSaleOpData is an XDR NestedUnion defines as:
//
//   union switch (ManageSaleAction action) {
//        case CREATE_UPDATE_DETAILS_REQUEST:
//            UpdateSaleDetailsData updateSaleDetailsData;
//        case CANCEL:
//            void;
//    	case SET_STATE:
//    		SaleState saleState;
//        case CREATE_PROMOTION_UPDATE_REQUEST:
//            PromotionUpdateData promotionUpdateData;
//        case CREATE_UPDATE_END_TIME_REQUEST:
//            UpdateSaleEndTimeData updateSaleEndTimeData;
//        }
//
type ManageSaleOpData struct {
	Action                ManageSaleAction       `json:"action,omitempty"`
	UpdateSaleDetailsData *UpdateSaleDetailsData `json:"updateSaleDetailsData,omitempty"`
	SaleState             *SaleState             `json:"saleState,omitempty"`
	PromotionUpdateData   *PromotionUpdateData   `json:"promotionUpdateData,omitempty"`
	UpdateSaleEndTimeData *UpdateSaleEndTimeData `json:"updateSaleEndTimeData,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageSaleOpData) SwitchFieldName() string {
	return "Action"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageSaleOpData
func (u ManageSaleOpData) ArmForSwitch(sw int32) (string, bool) {
	switch ManageSaleAction(sw) {
	case ManageSaleActionCreateUpdateDetailsRequest:
		return "UpdateSaleDetailsData", true
	case ManageSaleActionCancel:
		return "", true
	case ManageSaleActionSetState:
		return "SaleState", true
	case ManageSaleActionCreatePromotionUpdateRequest:
		return "PromotionUpdateData", true
	case ManageSaleActionCreateUpdateEndTimeRequest:
		return "UpdateSaleEndTimeData", true
	}
	return "-", false
}

// NewManageSaleOpData creates a new  ManageSaleOpData.
func NewManageSaleOpData(action ManageSaleAction, value interface{}) (result ManageSaleOpData, err error) {
	result.Action = action
	switch ManageSaleAction(action) {
	case ManageSaleActionCreateUpdateDetailsRequest:
		tv, ok := value.(UpdateSaleDetailsData)
		if !ok {
			err = fmt.Errorf("invalid value, must be UpdateSaleDetailsData")
			return
		}
		result.UpdateSaleDetailsData = &tv
	case ManageSaleActionCancel:
		// void
	case ManageSaleActionSetState:
		tv, ok := value.(SaleState)
		if !ok {
			err = fmt.Errorf("invalid value, must be SaleState")
			return
		}
		result.SaleState = &tv
	case ManageSaleActionCreatePromotionUpdateRequest:
		tv, ok := value.(PromotionUpdateData)
		if !ok {
			err = fmt.Errorf("invalid value, must be PromotionUpdateData")
			return
		}
		result.PromotionUpdateData = &tv
	case ManageSaleActionCreateUpdateEndTimeRequest:
		tv, ok := value.(UpdateSaleEndTimeData)
		if !ok {
			err = fmt.Errorf("invalid value, must be UpdateSaleEndTimeData")
			return
		}
		result.UpdateSaleEndTimeData = &tv
	}
	return
}

// MustUpdateSaleDetailsData retrieves the UpdateSaleDetailsData value from the union,
// panicing if the value is not set.
func (u ManageSaleOpData) MustUpdateSaleDetailsData() UpdateSaleDetailsData {
	val, ok := u.GetUpdateSaleDetailsData()

	if !ok {
		panic("arm UpdateSaleDetailsData is not set")
	}

	return val
}

// GetUpdateSaleDetailsData retrieves the UpdateSaleDetailsData value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageSaleOpData) GetUpdateSaleDetailsData() (result UpdateSaleDetailsData, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Action))

	if armName == "UpdateSaleDetailsData" {
		result = *u.UpdateSaleDetailsData
		ok = true
	}

	return
}

// MustSaleState retrieves the SaleState value from the union,
// panicing if the value is not set.
func (u ManageSaleOpData) MustSaleState() SaleState {
	val, ok := u.GetSaleState()

	if !ok {
		panic("arm SaleState is not set")
	}

	return val
}

// GetSaleState retrieves the SaleState value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageSaleOpData) GetSaleState() (result SaleState, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Action))

	if armName == "SaleState" {
		result = *u.SaleState
		ok = true
	}

	return
}

// MustPromotionUpdateData retrieves the PromotionUpdateData value from the union,
// panicing if the value is not set.
func (u ManageSaleOpData) MustPromotionUpdateData() PromotionUpdateData {
	val, ok := u.GetPromotionUpdateData()

	if !ok {
		panic("arm PromotionUpdateData is not set")
	}

	return val
}

// GetPromotionUpdateData retrieves the PromotionUpdateData value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageSaleOpData) GetPromotionUpdateData() (result PromotionUpdateData, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Action))

	if armName == "PromotionUpdateData" {
		result = *u.PromotionUpdateData
		ok = true
	}

	return
}

// MustUpdateSaleEndTimeData retrieves the UpdateSaleEndTimeData value from the union,
// panicing if the value is not set.
func (u ManageSaleOpData) MustUpdateSaleEndTimeData() UpdateSaleEndTimeData {
	val, ok := u.GetUpdateSaleEndTimeData()

	if !ok {
		panic("arm UpdateSaleEndTimeData is not set")
	}

	return val
}

// GetUpdateSaleEndTimeData retrieves the UpdateSaleEndTimeData value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageSaleOpData) GetUpdateSaleEndTimeData() (result UpdateSaleEndTimeData, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Action))

	if armName == "UpdateSaleEndTimeData" {
		result = *u.UpdateSaleEndTimeData
		ok = true
	}

	return
}

// ManageSaleOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type ManageSaleOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageSaleOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageSaleOpExt
func (u ManageSaleOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewManageSaleOpExt creates a new  ManageSaleOpExt.
func NewManageSaleOpExt(v LedgerVersion, value interface{}) (result ManageSaleOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ManageSaleOp is an XDR Struct defines as:
//
//   struct ManageSaleOp
//    {
//        uint64 saleID;
//
//        union switch (ManageSaleAction action) {
//        case CREATE_UPDATE_DETAILS_REQUEST:
//            UpdateSaleDetailsData updateSaleDetailsData;
//        case CANCEL:
//            void;
//    	case SET_STATE:
//    		SaleState saleState;
//        case CREATE_PROMOTION_UPDATE_REQUEST:
//            PromotionUpdateData promotionUpdateData;
//        case CREATE_UPDATE_END_TIME_REQUEST:
//            UpdateSaleEndTimeData updateSaleEndTimeData;
//        } data;
//
//        // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        } ext;
//    };
//
type ManageSaleOp struct {
	SaleId Uint64           `json:"saleID,omitempty"`
	Data   ManageSaleOpData `json:"data,omitempty"`
	Ext    ManageSaleOpExt  `json:"ext,omitempty"`
}

// ManageSaleResultCode is an XDR Enum defines as:
//
//   enum ManageSaleResultCode
//    {
//        SUCCESS = 0,
//
//        SALE_NOT_FOUND = -1, // sale not found
//
//        // errors related to action "CREATE_UPDATE_DETAILS_REQUEST"
//        INVALID_NEW_DETAILS = -2, // newDetails field is invalid JSON
//        UPDATE_DETAILS_REQUEST_ALREADY_EXISTS = -3,
//        UPDATE_DETAILS_REQUEST_NOT_FOUND = -4,
//
//        // errors related to action "SET_STATE"
//        NOT_ALLOWED = -5, // it's not allowed to set state for non master account
//
//        // errors related to action "CREATE_PROMOTION_UPDATE_REQUEST"
//        PROMOTION_UPDATE_REQUEST_INVALID_ASSET_PAIR = -6, // one of the assets has invalid code or base asset is equal to quote asset
//        PROMOTION_UPDATE_REQUEST_INVALID_PRICE = -7, // price cannot be 0
//        PROMOTION_UPDATE_REQUEST_START_END_INVALID = -8, // sale ends before start
//        PROMOTION_UPDATE_REQUEST_INVALID_CAP = -9, // hard cap is < soft cap
//        PROMOTION_UPDATE_REQUEST_INVALID_DETAILS = -10, // details field is invalid JSON
//        INVALID_SALE_STATE = -11, // sale state must be "PROMOTION"
//        PROMOTION_UPDATE_REQUEST_ALREADY_EXISTS = -12,
//        PROMOTION_UPDATE_REQUEST_NOT_FOUND = -13,
//
//        // errors related to action "CREATE_UPDATE_END_TIME_REQUEST"
//        INVALID_NEW_END_TIME = -14, // new end time is before start time or current ledger close time
//        UPDATE_END_TIME_REQUEST_ALREADY_EXISTS = -15,
//        UPDATE_END_TIME_REQUEST_NOT_FOUND = -16
//    };
//
type ManageSaleResultCode int32

const (
	ManageSaleResultCodeSuccess                                ManageSaleResultCode = 0
	ManageSaleResultCodeSaleNotFound                           ManageSaleResultCode = -1
	ManageSaleResultCodeInvalidNewDetails                      ManageSaleResultCode = -2
	ManageSaleResultCodeUpdateDetailsRequestAlreadyExists      ManageSaleResultCode = -3
	ManageSaleResultCodeUpdateDetailsRequestNotFound           ManageSaleResultCode = -4
	ManageSaleResultCodeNotAllowed                             ManageSaleResultCode = -5
	ManageSaleResultCodePromotionUpdateRequestInvalidAssetPair ManageSaleResultCode = -6
	ManageSaleResultCodePromotionUpdateRequestInvalidPrice     ManageSaleResultCode = -7
	ManageSaleResultCodePromotionUpdateRequestStartEndInvalid  ManageSaleResultCode = -8
	ManageSaleResultCodePromotionUpdateRequestInvalidCap       ManageSaleResultCode = -9
	ManageSaleResultCodePromotionUpdateRequestInvalidDetails   ManageSaleResultCode = -10
	ManageSaleResultCodeInvalidSaleState                       ManageSaleResultCode = -11
	ManageSaleResultCodePromotionUpdateRequestAlreadyExists    ManageSaleResultCode = -12
	ManageSaleResultCodePromotionUpdateRequestNotFound         ManageSaleResultCode = -13
	ManageSaleResultCodeInvalidNewEndTime                      ManageSaleResultCode = -14
	ManageSaleResultCodeUpdateEndTimeRequestAlreadyExists      ManageSaleResultCode = -15
	ManageSaleResultCodeUpdateEndTimeRequestNotFound           ManageSaleResultCode = -16
)

var ManageSaleResultCodeAll = []ManageSaleResultCode{
	ManageSaleResultCodeSuccess,
	ManageSaleResultCodeSaleNotFound,
	ManageSaleResultCodeInvalidNewDetails,
	ManageSaleResultCodeUpdateDetailsRequestAlreadyExists,
	ManageSaleResultCodeUpdateDetailsRequestNotFound,
	ManageSaleResultCodeNotAllowed,
	ManageSaleResultCodePromotionUpdateRequestInvalidAssetPair,
	ManageSaleResultCodePromotionUpdateRequestInvalidPrice,
	ManageSaleResultCodePromotionUpdateRequestStartEndInvalid,
	ManageSaleResultCodePromotionUpdateRequestInvalidCap,
	ManageSaleResultCodePromotionUpdateRequestInvalidDetails,
	ManageSaleResultCodeInvalidSaleState,
	ManageSaleResultCodePromotionUpdateRequestAlreadyExists,
	ManageSaleResultCodePromotionUpdateRequestNotFound,
	ManageSaleResultCodeInvalidNewEndTime,
	ManageSaleResultCodeUpdateEndTimeRequestAlreadyExists,
	ManageSaleResultCodeUpdateEndTimeRequestNotFound,
}

var manageSaleResultCodeMap = map[int32]string{
	0:   "ManageSaleResultCodeSuccess",
	-1:  "ManageSaleResultCodeSaleNotFound",
	-2:  "ManageSaleResultCodeInvalidNewDetails",
	-3:  "ManageSaleResultCodeUpdateDetailsRequestAlreadyExists",
	-4:  "ManageSaleResultCodeUpdateDetailsRequestNotFound",
	-5:  "ManageSaleResultCodeNotAllowed",
	-6:  "ManageSaleResultCodePromotionUpdateRequestInvalidAssetPair",
	-7:  "ManageSaleResultCodePromotionUpdateRequestInvalidPrice",
	-8:  "ManageSaleResultCodePromotionUpdateRequestStartEndInvalid",
	-9:  "ManageSaleResultCodePromotionUpdateRequestInvalidCap",
	-10: "ManageSaleResultCodePromotionUpdateRequestInvalidDetails",
	-11: "ManageSaleResultCodeInvalidSaleState",
	-12: "ManageSaleResultCodePromotionUpdateRequestAlreadyExists",
	-13: "ManageSaleResultCodePromotionUpdateRequestNotFound",
	-14: "ManageSaleResultCodeInvalidNewEndTime",
	-15: "ManageSaleResultCodeUpdateEndTimeRequestAlreadyExists",
	-16: "ManageSaleResultCodeUpdateEndTimeRequestNotFound",
}

var manageSaleResultCodeShortMap = map[int32]string{
	0:   "success",
	-1:  "sale_not_found",
	-2:  "invalid_new_details",
	-3:  "update_details_request_already_exists",
	-4:  "update_details_request_not_found",
	-5:  "not_allowed",
	-6:  "promotion_update_request_invalid_asset_pair",
	-7:  "promotion_update_request_invalid_price",
	-8:  "promotion_update_request_start_end_invalid",
	-9:  "promotion_update_request_invalid_cap",
	-10: "promotion_update_request_invalid_details",
	-11: "invalid_sale_state",
	-12: "promotion_update_request_already_exists",
	-13: "promotion_update_request_not_found",
	-14: "invalid_new_end_time",
	-15: "update_end_time_request_already_exists",
	-16: "update_end_time_request_not_found",
}

var manageSaleResultCodeRevMap = map[string]int32{
	"ManageSaleResultCodeSuccess":                                0,
	"ManageSaleResultCodeSaleNotFound":                           -1,
	"ManageSaleResultCodeInvalidNewDetails":                      -2,
	"ManageSaleResultCodeUpdateDetailsRequestAlreadyExists":      -3,
	"ManageSaleResultCodeUpdateDetailsRequestNotFound":           -4,
	"ManageSaleResultCodeNotAllowed":                             -5,
	"ManageSaleResultCodePromotionUpdateRequestInvalidAssetPair": -6,
	"ManageSaleResultCodePromotionUpdateRequestInvalidPrice":     -7,
	"ManageSaleResultCodePromotionUpdateRequestStartEndInvalid":  -8,
	"ManageSaleResultCodePromotionUpdateRequestInvalidCap":       -9,
	"ManageSaleResultCodePromotionUpdateRequestInvalidDetails":   -10,
	"ManageSaleResultCodeInvalidSaleState":                       -11,
	"ManageSaleResultCodePromotionUpdateRequestAlreadyExists":    -12,
	"ManageSaleResultCodePromotionUpdateRequestNotFound":         -13,
	"ManageSaleResultCodeInvalidNewEndTime":                      -14,
	"ManageSaleResultCodeUpdateEndTimeRequestAlreadyExists":      -15,
	"ManageSaleResultCodeUpdateEndTimeRequestNotFound":           -16,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ManageSaleResultCode
func (e ManageSaleResultCode) ValidEnum(v int32) bool {
	_, ok := manageSaleResultCodeMap[v]
	return ok
}
func (e ManageSaleResultCode) isFlag() bool {
	for i := len(ManageSaleResultCodeAll) - 1; i >= 0; i-- {
		expected := ManageSaleResultCode(2) << uint64(len(ManageSaleResultCodeAll)-1) >> uint64(len(ManageSaleResultCodeAll)-i)
		if expected != ManageSaleResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ManageSaleResultCode) String() string {
	name, _ := manageSaleResultCodeMap[int32(e)]
	return name
}

func (e ManageSaleResultCode) ShortString() string {
	name, _ := manageSaleResultCodeShortMap[int32(e)]
	return name
}

func (e ManageSaleResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range ManageSaleResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ManageSaleResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ManageSaleResultCode(t.Value)
	return nil
}

// ManageSaleResultSuccessResponse is an XDR NestedUnion defines as:
//
//   union switch (ManageSaleAction action) {
//        case CREATE_UPDATE_DETAILS_REQUEST:
//            uint64 requestID;
//        case CANCEL:
//            void;
//    	case SET_STATE:
//    		void;
//        case CREATE_PROMOTION_UPDATE_REQUEST:
//            uint64 promotionUpdateRequestID;
//    	case CREATE_UPDATE_END_TIME_REQUEST:
//    	    uint64 updateEndTimeRequestID;
//        }
//
type ManageSaleResultSuccessResponse struct {
	Action                   ManageSaleAction `json:"action,omitempty"`
	RequestId                *Uint64          `json:"requestID,omitempty"`
	PromotionUpdateRequestId *Uint64          `json:"promotionUpdateRequestID,omitempty"`
	UpdateEndTimeRequestId   *Uint64          `json:"updateEndTimeRequestID,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageSaleResultSuccessResponse) SwitchFieldName() string {
	return "Action"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageSaleResultSuccessResponse
func (u ManageSaleResultSuccessResponse) ArmForSwitch(sw int32) (string, bool) {
	switch ManageSaleAction(sw) {
	case ManageSaleActionCreateUpdateDetailsRequest:
		return "RequestId", true
	case ManageSaleActionCancel:
		return "", true
	case ManageSaleActionSetState:
		return "", true
	case ManageSaleActionCreatePromotionUpdateRequest:
		return "PromotionUpdateRequestId", true
	case ManageSaleActionCreateUpdateEndTimeRequest:
		return "UpdateEndTimeRequestId", true
	}
	return "-", false
}

// NewManageSaleResultSuccessResponse creates a new  ManageSaleResultSuccessResponse.
func NewManageSaleResultSuccessResponse(action ManageSaleAction, value interface{}) (result ManageSaleResultSuccessResponse, err error) {
	result.Action = action
	switch ManageSaleAction(action) {
	case ManageSaleActionCreateUpdateDetailsRequest:
		tv, ok := value.(Uint64)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint64")
			return
		}
		result.RequestId = &tv
	case ManageSaleActionCancel:
		// void
	case ManageSaleActionSetState:
		// void
	case ManageSaleActionCreatePromotionUpdateRequest:
		tv, ok := value.(Uint64)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint64")
			return
		}
		result.PromotionUpdateRequestId = &tv
	case ManageSaleActionCreateUpdateEndTimeRequest:
		tv, ok := value.(Uint64)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint64")
			return
		}
		result.UpdateEndTimeRequestId = &tv
	}
	return
}

// MustRequestId retrieves the RequestId value from the union,
// panicing if the value is not set.
func (u ManageSaleResultSuccessResponse) MustRequestId() Uint64 {
	val, ok := u.GetRequestId()

	if !ok {
		panic("arm RequestId is not set")
	}

	return val
}

// GetRequestId retrieves the RequestId value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageSaleResultSuccessResponse) GetRequestId() (result Uint64, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Action))

	if armName == "RequestId" {
		result = *u.RequestId
		ok = true
	}

	return
}

// MustPromotionUpdateRequestId retrieves the PromotionUpdateRequestId value from the union,
// panicing if the value is not set.
func (u ManageSaleResultSuccessResponse) MustPromotionUpdateRequestId() Uint64 {
	val, ok := u.GetPromotionUpdateRequestId()

	if !ok {
		panic("arm PromotionUpdateRequestId is not set")
	}

	return val
}

// GetPromotionUpdateRequestId retrieves the PromotionUpdateRequestId value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageSaleResultSuccessResponse) GetPromotionUpdateRequestId() (result Uint64, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Action))

	if armName == "PromotionUpdateRequestId" {
		result = *u.PromotionUpdateRequestId
		ok = true
	}

	return
}

// MustUpdateEndTimeRequestId retrieves the UpdateEndTimeRequestId value from the union,
// panicing if the value is not set.
func (u ManageSaleResultSuccessResponse) MustUpdateEndTimeRequestId() Uint64 {
	val, ok := u.GetUpdateEndTimeRequestId()

	if !ok {
		panic("arm UpdateEndTimeRequestId is not set")
	}

	return val
}

// GetUpdateEndTimeRequestId retrieves the UpdateEndTimeRequestId value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageSaleResultSuccessResponse) GetUpdateEndTimeRequestId() (result Uint64, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Action))

	if armName == "UpdateEndTimeRequestId" {
		result = *u.UpdateEndTimeRequestId
		ok = true
	}

	return
}

// ManageSaleResultSuccessExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        case ALLOW_TO_UPDATE_VOTING_SALES_AS_PROMOTION:
//            bool fulfilled; // can be used for any reviewable request type created with manage sale operation
//        }
//
type ManageSaleResultSuccessExt struct {
	V         LedgerVersion `json:"v,omitempty"`
	Fulfilled *bool         `json:"fulfilled,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageSaleResultSuccessExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageSaleResultSuccessExt
func (u ManageSaleResultSuccessExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	case LedgerVersionAllowToUpdateVotingSalesAsPromotion:
		return "Fulfilled", true
	}
	return "-", false
}

// NewManageSaleResultSuccessExt creates a new  ManageSaleResultSuccessExt.
func NewManageSaleResultSuccessExt(v LedgerVersion, value interface{}) (result ManageSaleResultSuccessExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	case LedgerVersionAllowToUpdateVotingSalesAsPromotion:
		tv, ok := value.(bool)
		if !ok {
			err = fmt.Errorf("invalid value, must be bool")
			return
		}
		result.Fulfilled = &tv
	}
	return
}

// MustFulfilled retrieves the Fulfilled value from the union,
// panicing if the value is not set.
func (u ManageSaleResultSuccessExt) MustFulfilled() bool {
	val, ok := u.GetFulfilled()

	if !ok {
		panic("arm Fulfilled is not set")
	}

	return val
}

// GetFulfilled retrieves the Fulfilled value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageSaleResultSuccessExt) GetFulfilled() (result bool, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.V))

	if armName == "Fulfilled" {
		result = *u.Fulfilled
		ok = true
	}

	return
}

// ManageSaleResultSuccess is an XDR Struct defines as:
//
//   struct ManageSaleResultSuccess
//    {
//        union switch (ManageSaleAction action) {
//        case CREATE_UPDATE_DETAILS_REQUEST:
//            uint64 requestID;
//        case CANCEL:
//            void;
//    	case SET_STATE:
//    		void;
//        case CREATE_PROMOTION_UPDATE_REQUEST:
//            uint64 promotionUpdateRequestID;
//    	case CREATE_UPDATE_END_TIME_REQUEST:
//    	    uint64 updateEndTimeRequestID;
//        } response;
//
//        // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        case ALLOW_TO_UPDATE_VOTING_SALES_AS_PROMOTION:
//            bool fulfilled; // can be used for any reviewable request type created with manage sale operation
//        }
//        ext;
//    };
//
type ManageSaleResultSuccess struct {
	Response ManageSaleResultSuccessResponse `json:"response,omitempty"`
	Ext      ManageSaleResultSuccessExt      `json:"ext,omitempty"`
}

// ManageSaleResult is an XDR Union defines as:
//
//   union ManageSaleResult switch (ManageSaleResultCode code)
//    {
//    case SUCCESS:
//        ManageSaleResultSuccess success;
//    default:
//        void;
//    };
//
type ManageSaleResult struct {
	Code    ManageSaleResultCode     `json:"code,omitempty"`
	Success *ManageSaleResultSuccess `json:"success,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageSaleResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageSaleResult
func (u ManageSaleResult) ArmForSwitch(sw int32) (string, bool) {
	switch ManageSaleResultCode(sw) {
	case ManageSaleResultCodeSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewManageSaleResult creates a new  ManageSaleResult.
func NewManageSaleResult(code ManageSaleResultCode, value interface{}) (result ManageSaleResult, err error) {
	result.Code = code
	switch ManageSaleResultCode(code) {
	case ManageSaleResultCodeSuccess:
		tv, ok := value.(ManageSaleResultSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageSaleResultSuccess")
			return
		}
		result.Success = &tv
	default:
		// void
	}
	return
}

// MustSuccess retrieves the Success value from the union,
// panicing if the value is not set.
func (u ManageSaleResult) MustSuccess() ManageSaleResultSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageSaleResult) GetSuccess() (result ManageSaleResultSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// FeeDataV2Ext is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type FeeDataV2Ext struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u FeeDataV2Ext) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of FeeDataV2Ext
func (u FeeDataV2Ext) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewFeeDataV2Ext creates a new  FeeDataV2Ext.
func NewFeeDataV2Ext(v LedgerVersion, value interface{}) (result FeeDataV2Ext, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// FeeDataV2 is an XDR Struct defines as:
//
//   struct FeeDataV2 {
//        uint64 maxPaymentFee;
//        uint64 fixedFee;
//
//        // Cross asset fees
//        AssetCode feeAsset;
//
//    	// reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type FeeDataV2 struct {
	MaxPaymentFee Uint64       `json:"maxPaymentFee,omitempty"`
	FixedFee      Uint64       `json:"fixedFee,omitempty"`
	FeeAsset      AssetCode    `json:"feeAsset,omitempty"`
	Ext           FeeDataV2Ext `json:"ext,omitempty"`
}

// PaymentFeeDataV2Ext is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type PaymentFeeDataV2Ext struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u PaymentFeeDataV2Ext) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of PaymentFeeDataV2Ext
func (u PaymentFeeDataV2Ext) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewPaymentFeeDataV2Ext creates a new  PaymentFeeDataV2Ext.
func NewPaymentFeeDataV2Ext(v LedgerVersion, value interface{}) (result PaymentFeeDataV2Ext, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// PaymentFeeDataV2 is an XDR Struct defines as:
//
//   struct PaymentFeeDataV2 {
//        FeeDataV2 sourceFee;
//        FeeDataV2 destinationFee;
//        bool sourcePaysForDest; // if true - source account pays fee, else destination
//
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type PaymentFeeDataV2 struct {
	SourceFee         FeeDataV2           `json:"sourceFee,omitempty"`
	DestinationFee    FeeDataV2           `json:"destinationFee,omitempty"`
	SourcePaysForDest bool                `json:"sourcePaysForDest,omitempty"`
	Ext               PaymentFeeDataV2Ext `json:"ext,omitempty"`
}

// PaymentDestinationType is an XDR Enum defines as:
//
//   enum PaymentDestinationType {
//        ACCOUNT = 0,
//        BALANCE = 1
//    };
//
type PaymentDestinationType int32

const (
	PaymentDestinationTypeAccount PaymentDestinationType = 0
	PaymentDestinationTypeBalance PaymentDestinationType = 1
)

var PaymentDestinationTypeAll = []PaymentDestinationType{
	PaymentDestinationTypeAccount,
	PaymentDestinationTypeBalance,
}

var paymentDestinationTypeMap = map[int32]string{
	0: "PaymentDestinationTypeAccount",
	1: "PaymentDestinationTypeBalance",
}

var paymentDestinationTypeShortMap = map[int32]string{
	0: "account",
	1: "balance",
}

var paymentDestinationTypeRevMap = map[string]int32{
	"PaymentDestinationTypeAccount": 0,
	"PaymentDestinationTypeBalance": 1,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for PaymentDestinationType
func (e PaymentDestinationType) ValidEnum(v int32) bool {
	_, ok := paymentDestinationTypeMap[v]
	return ok
}
func (e PaymentDestinationType) isFlag() bool {
	for i := len(PaymentDestinationTypeAll) - 1; i >= 0; i-- {
		expected := PaymentDestinationType(2) << uint64(len(PaymentDestinationTypeAll)-1) >> uint64(len(PaymentDestinationTypeAll)-i)
		if expected != PaymentDestinationTypeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e PaymentDestinationType) String() string {
	name, _ := paymentDestinationTypeMap[int32(e)]
	return name
}

func (e PaymentDestinationType) ShortString() string {
	name, _ := paymentDestinationTypeShortMap[int32(e)]
	return name
}

func (e PaymentDestinationType) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range PaymentDestinationTypeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *PaymentDestinationType) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = PaymentDestinationType(t.Value)
	return nil
}

// PaymentOpV2Destination is an XDR NestedUnion defines as:
//
//   union switch (PaymentDestinationType type) {
//            case ACCOUNT:
//                AccountID accountID;
//            case BALANCE:
//                BalanceID balanceID;
//        }
//
type PaymentOpV2Destination struct {
	Type      PaymentDestinationType `json:"type,omitempty"`
	AccountId *AccountId             `json:"accountID,omitempty"`
	BalanceId *BalanceId             `json:"balanceID,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u PaymentOpV2Destination) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of PaymentOpV2Destination
func (u PaymentOpV2Destination) ArmForSwitch(sw int32) (string, bool) {
	switch PaymentDestinationType(sw) {
	case PaymentDestinationTypeAccount:
		return "AccountId", true
	case PaymentDestinationTypeBalance:
		return "BalanceId", true
	}
	return "-", false
}

// NewPaymentOpV2Destination creates a new  PaymentOpV2Destination.
func NewPaymentOpV2Destination(aType PaymentDestinationType, value interface{}) (result PaymentOpV2Destination, err error) {
	result.Type = aType
	switch PaymentDestinationType(aType) {
	case PaymentDestinationTypeAccount:
		tv, ok := value.(AccountId)
		if !ok {
			err = fmt.Errorf("invalid value, must be AccountId")
			return
		}
		result.AccountId = &tv
	case PaymentDestinationTypeBalance:
		tv, ok := value.(BalanceId)
		if !ok {
			err = fmt.Errorf("invalid value, must be BalanceId")
			return
		}
		result.BalanceId = &tv
	}
	return
}

// MustAccountId retrieves the AccountId value from the union,
// panicing if the value is not set.
func (u PaymentOpV2Destination) MustAccountId() AccountId {
	val, ok := u.GetAccountId()

	if !ok {
		panic("arm AccountId is not set")
	}

	return val
}

// GetAccountId retrieves the AccountId value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u PaymentOpV2Destination) GetAccountId() (result AccountId, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "AccountId" {
		result = *u.AccountId
		ok = true
	}

	return
}

// MustBalanceId retrieves the BalanceId value from the union,
// panicing if the value is not set.
func (u PaymentOpV2Destination) MustBalanceId() BalanceId {
	val, ok := u.GetBalanceId()

	if !ok {
		panic("arm BalanceId is not set")
	}

	return val
}

// GetBalanceId retrieves the BalanceId value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u PaymentOpV2Destination) GetBalanceId() (result BalanceId, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "BalanceId" {
		result = *u.BalanceId
		ok = true
	}

	return
}

// PaymentOpV2Ext is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type PaymentOpV2Ext struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u PaymentOpV2Ext) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of PaymentOpV2Ext
func (u PaymentOpV2Ext) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewPaymentOpV2Ext creates a new  PaymentOpV2Ext.
func NewPaymentOpV2Ext(v LedgerVersion, value interface{}) (result PaymentOpV2Ext, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// PaymentOpV2 is an XDR Struct defines as:
//
//   struct PaymentOpV2
//    {
//        BalanceID sourceBalanceID;
//
//        union switch (PaymentDestinationType type) {
//            case ACCOUNT:
//                AccountID accountID;
//            case BALANCE:
//                BalanceID balanceID;
//        } destination;
//
//        uint64 amount;
//
//        PaymentFeeDataV2 feeData;
//
//        longstring subject;
//        longstring reference;
//
//        // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type PaymentOpV2 struct {
	SourceBalanceId BalanceId              `json:"sourceBalanceID,omitempty"`
	Destination     PaymentOpV2Destination `json:"destination,omitempty"`
	Amount          Uint64                 `json:"amount,omitempty"`
	FeeData         PaymentFeeDataV2       `json:"feeData,omitempty"`
	Subject         Longstring             `json:"subject,omitempty"`
	Reference       Longstring             `json:"reference,omitempty"`
	Ext             PaymentOpV2Ext         `json:"ext,omitempty"`
}

// PaymentV2ResultCode is an XDR Enum defines as:
//
//   enum PaymentV2ResultCode
//    {
//        // codes considered as "success" for the operation
//        SUCCESS = 0, // payment successfully completed
//
//        // codes considered as "failure" for the operation
//        MALFORMED = -1, // bad input
//        UNDERFUNDED = -2, // not enough funds in source account
//        LINE_FULL = -3, // destination would go above their limit
//    	DESTINATION_BALANCE_NOT_FOUND = -4,
//        BALANCE_ASSETS_MISMATCHED = -5,
//    	SRC_BALANCE_NOT_FOUND = -6, // source balance not found
//        REFERENCE_DUPLICATION = -7,
//        STATS_OVERFLOW = -8,
//        LIMITS_EXCEEDED = -9,
//        NOT_ALLOWED_BY_ASSET_POLICY = -10,
//        INVALID_DESTINATION_FEE = -11,
//        INVALID_DESTINATION_FEE_ASSET = -12, // destination fee asset must be the same as source balance asset
//        FEE_ASSET_MISMATCHED = -13,
//        INSUFFICIENT_FEE_AMOUNT = -14,
//        BALANCE_TO_CHARGE_FEE_FROM_NOT_FOUND = -15,
//        PAYMENT_AMOUNT_IS_LESS_THAN_DEST_FEE = -16,
//        DESTINATION_ACCOUNT_NOT_FOUND = -17
//
//         // !!! Add new result code to review invoice op too !!!
//    };
//
type PaymentV2ResultCode int32

const (
	PaymentV2ResultCodeSuccess                        PaymentV2ResultCode = 0
	PaymentV2ResultCodeMalformed                      PaymentV2ResultCode = -1
	PaymentV2ResultCodeUnderfunded                    PaymentV2ResultCode = -2
	PaymentV2ResultCodeLineFull                       PaymentV2ResultCode = -3
	PaymentV2ResultCodeDestinationBalanceNotFound     PaymentV2ResultCode = -4
	PaymentV2ResultCodeBalanceAssetsMismatched        PaymentV2ResultCode = -5
	PaymentV2ResultCodeSrcBalanceNotFound             PaymentV2ResultCode = -6
	PaymentV2ResultCodeReferenceDuplication           PaymentV2ResultCode = -7
	PaymentV2ResultCodeStatsOverflow                  PaymentV2ResultCode = -8
	PaymentV2ResultCodeLimitsExceeded                 PaymentV2ResultCode = -9
	PaymentV2ResultCodeNotAllowedByAssetPolicy        PaymentV2ResultCode = -10
	PaymentV2ResultCodeInvalidDestinationFee          PaymentV2ResultCode = -11
	PaymentV2ResultCodeInvalidDestinationFeeAsset     PaymentV2ResultCode = -12
	PaymentV2ResultCodeFeeAssetMismatched             PaymentV2ResultCode = -13
	PaymentV2ResultCodeInsufficientFeeAmount          PaymentV2ResultCode = -14
	PaymentV2ResultCodeBalanceToChargeFeeFromNotFound PaymentV2ResultCode = -15
	PaymentV2ResultCodePaymentAmountIsLessThanDestFee PaymentV2ResultCode = -16
	PaymentV2ResultCodeDestinationAccountNotFound     PaymentV2ResultCode = -17
)

var PaymentV2ResultCodeAll = []PaymentV2ResultCode{
	PaymentV2ResultCodeSuccess,
	PaymentV2ResultCodeMalformed,
	PaymentV2ResultCodeUnderfunded,
	PaymentV2ResultCodeLineFull,
	PaymentV2ResultCodeDestinationBalanceNotFound,
	PaymentV2ResultCodeBalanceAssetsMismatched,
	PaymentV2ResultCodeSrcBalanceNotFound,
	PaymentV2ResultCodeReferenceDuplication,
	PaymentV2ResultCodeStatsOverflow,
	PaymentV2ResultCodeLimitsExceeded,
	PaymentV2ResultCodeNotAllowedByAssetPolicy,
	PaymentV2ResultCodeInvalidDestinationFee,
	PaymentV2ResultCodeInvalidDestinationFeeAsset,
	PaymentV2ResultCodeFeeAssetMismatched,
	PaymentV2ResultCodeInsufficientFeeAmount,
	PaymentV2ResultCodeBalanceToChargeFeeFromNotFound,
	PaymentV2ResultCodePaymentAmountIsLessThanDestFee,
	PaymentV2ResultCodeDestinationAccountNotFound,
}

var paymentV2ResultCodeMap = map[int32]string{
	0:   "PaymentV2ResultCodeSuccess",
	-1:  "PaymentV2ResultCodeMalformed",
	-2:  "PaymentV2ResultCodeUnderfunded",
	-3:  "PaymentV2ResultCodeLineFull",
	-4:  "PaymentV2ResultCodeDestinationBalanceNotFound",
	-5:  "PaymentV2ResultCodeBalanceAssetsMismatched",
	-6:  "PaymentV2ResultCodeSrcBalanceNotFound",
	-7:  "PaymentV2ResultCodeReferenceDuplication",
	-8:  "PaymentV2ResultCodeStatsOverflow",
	-9:  "PaymentV2ResultCodeLimitsExceeded",
	-10: "PaymentV2ResultCodeNotAllowedByAssetPolicy",
	-11: "PaymentV2ResultCodeInvalidDestinationFee",
	-12: "PaymentV2ResultCodeInvalidDestinationFeeAsset",
	-13: "PaymentV2ResultCodeFeeAssetMismatched",
	-14: "PaymentV2ResultCodeInsufficientFeeAmount",
	-15: "PaymentV2ResultCodeBalanceToChargeFeeFromNotFound",
	-16: "PaymentV2ResultCodePaymentAmountIsLessThanDestFee",
	-17: "PaymentV2ResultCodeDestinationAccountNotFound",
}

var paymentV2ResultCodeShortMap = map[int32]string{
	0:   "success",
	-1:  "malformed",
	-2:  "underfunded",
	-3:  "line_full",
	-4:  "destination_balance_not_found",
	-5:  "balance_assets_mismatched",
	-6:  "src_balance_not_found",
	-7:  "reference_duplication",
	-8:  "stats_overflow",
	-9:  "limits_exceeded",
	-10: "not_allowed_by_asset_policy",
	-11: "invalid_destination_fee",
	-12: "invalid_destination_fee_asset",
	-13: "fee_asset_mismatched",
	-14: "insufficient_fee_amount",
	-15: "balance_to_charge_fee_from_not_found",
	-16: "payment_amount_is_less_than_dest_fee",
	-17: "destination_account_not_found",
}

var paymentV2ResultCodeRevMap = map[string]int32{
	"PaymentV2ResultCodeSuccess":                        0,
	"PaymentV2ResultCodeMalformed":                      -1,
	"PaymentV2ResultCodeUnderfunded":                    -2,
	"PaymentV2ResultCodeLineFull":                       -3,
	"PaymentV2ResultCodeDestinationBalanceNotFound":     -4,
	"PaymentV2ResultCodeBalanceAssetsMismatched":        -5,
	"PaymentV2ResultCodeSrcBalanceNotFound":             -6,
	"PaymentV2ResultCodeReferenceDuplication":           -7,
	"PaymentV2ResultCodeStatsOverflow":                  -8,
	"PaymentV2ResultCodeLimitsExceeded":                 -9,
	"PaymentV2ResultCodeNotAllowedByAssetPolicy":        -10,
	"PaymentV2ResultCodeInvalidDestinationFee":          -11,
	"PaymentV2ResultCodeInvalidDestinationFeeAsset":     -12,
	"PaymentV2ResultCodeFeeAssetMismatched":             -13,
	"PaymentV2ResultCodeInsufficientFeeAmount":          -14,
	"PaymentV2ResultCodeBalanceToChargeFeeFromNotFound": -15,
	"PaymentV2ResultCodePaymentAmountIsLessThanDestFee": -16,
	"PaymentV2ResultCodeDestinationAccountNotFound":     -17,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for PaymentV2ResultCode
func (e PaymentV2ResultCode) ValidEnum(v int32) bool {
	_, ok := paymentV2ResultCodeMap[v]
	return ok
}
func (e PaymentV2ResultCode) isFlag() bool {
	for i := len(PaymentV2ResultCodeAll) - 1; i >= 0; i-- {
		expected := PaymentV2ResultCode(2) << uint64(len(PaymentV2ResultCodeAll)-1) >> uint64(len(PaymentV2ResultCodeAll)-i)
		if expected != PaymentV2ResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e PaymentV2ResultCode) String() string {
	name, _ := paymentV2ResultCodeMap[int32(e)]
	return name
}

func (e PaymentV2ResultCode) ShortString() string {
	name, _ := paymentV2ResultCodeShortMap[int32(e)]
	return name
}

func (e PaymentV2ResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range PaymentV2ResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *PaymentV2ResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = PaymentV2ResultCode(t.Value)
	return nil
}

// PaymentV2ResponseExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type PaymentV2ResponseExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u PaymentV2ResponseExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of PaymentV2ResponseExt
func (u PaymentV2ResponseExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewPaymentV2ResponseExt creates a new  PaymentV2ResponseExt.
func NewPaymentV2ResponseExt(v LedgerVersion, value interface{}) (result PaymentV2ResponseExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// PaymentV2Response is an XDR Struct defines as:
//
//   struct PaymentV2Response {
//        AccountID destination;
//        BalanceID destinationBalanceID;
//
//        AssetCode asset;
//        uint64 sourceSentUniversal;
//        uint64 paymentID;
//
//        uint64 actualSourcePaymentFee;
//        uint64 actualDestinationPaymentFee;
//
//        // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type PaymentV2Response struct {
	Destination                 AccountId            `json:"destination,omitempty"`
	DestinationBalanceId        BalanceId            `json:"destinationBalanceID,omitempty"`
	Asset                       AssetCode            `json:"asset,omitempty"`
	SourceSentUniversal         Uint64               `json:"sourceSentUniversal,omitempty"`
	PaymentId                   Uint64               `json:"paymentID,omitempty"`
	ActualSourcePaymentFee      Uint64               `json:"actualSourcePaymentFee,omitempty"`
	ActualDestinationPaymentFee Uint64               `json:"actualDestinationPaymentFee,omitempty"`
	Ext                         PaymentV2ResponseExt `json:"ext,omitempty"`
}

// PaymentV2Result is an XDR Union defines as:
//
//   union PaymentV2Result switch (PaymentV2ResultCode code)
//    {
//    case SUCCESS:
//        PaymentV2Response paymentV2Response;
//    default:
//        void;
//    };
//
type PaymentV2Result struct {
	Code              PaymentV2ResultCode `json:"code,omitempty"`
	PaymentV2Response *PaymentV2Response  `json:"paymentV2Response,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u PaymentV2Result) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of PaymentV2Result
func (u PaymentV2Result) ArmForSwitch(sw int32) (string, bool) {
	switch PaymentV2ResultCode(sw) {
	case PaymentV2ResultCodeSuccess:
		return "PaymentV2Response", true
	default:
		return "", true
	}
}

// NewPaymentV2Result creates a new  PaymentV2Result.
func NewPaymentV2Result(code PaymentV2ResultCode, value interface{}) (result PaymentV2Result, err error) {
	result.Code = code
	switch PaymentV2ResultCode(code) {
	case PaymentV2ResultCodeSuccess:
		tv, ok := value.(PaymentV2Response)
		if !ok {
			err = fmt.Errorf("invalid value, must be PaymentV2Response")
			return
		}
		result.PaymentV2Response = &tv
	default:
		// void
	}
	return
}

// MustPaymentV2Response retrieves the PaymentV2Response value from the union,
// panicing if the value is not set.
func (u PaymentV2Result) MustPaymentV2Response() PaymentV2Response {
	val, ok := u.GetPaymentV2Response()

	if !ok {
		panic("arm PaymentV2Response is not set")
	}

	return val
}

// GetPaymentV2Response retrieves the PaymentV2Response value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u PaymentV2Result) GetPaymentV2Response() (result PaymentV2Response, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "PaymentV2Response" {
		result = *u.PaymentV2Response
		ok = true
	}

	return
}

// InvoiceReferenceExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type InvoiceReferenceExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u InvoiceReferenceExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of InvoiceReferenceExt
func (u InvoiceReferenceExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewInvoiceReferenceExt creates a new  InvoiceReferenceExt.
func NewInvoiceReferenceExt(v LedgerVersion, value interface{}) (result InvoiceReferenceExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// InvoiceReference is an XDR Struct defines as:
//
//   struct InvoiceReference {
//        uint64 invoiceID;
//        bool accept;
//    	// reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type InvoiceReference struct {
	InvoiceId Uint64              `json:"invoiceID,omitempty"`
	Accept    bool                `json:"accept,omitempty"`
	Ext       InvoiceReferenceExt `json:"ext,omitempty"`
}

// FeeDataExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type FeeDataExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u FeeDataExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of FeeDataExt
func (u FeeDataExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewFeeDataExt creates a new  FeeDataExt.
func NewFeeDataExt(v LedgerVersion, value interface{}) (result FeeDataExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// FeeData is an XDR Struct defines as:
//
//   struct FeeData {
//        int64 paymentFee;
//        int64 fixedFee;
//    	// reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type FeeData struct {
	PaymentFee Int64      `json:"paymentFee,omitempty"`
	FixedFee   Int64      `json:"fixedFee,omitempty"`
	Ext        FeeDataExt `json:"ext,omitempty"`
}

// PaymentFeeDataExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type PaymentFeeDataExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u PaymentFeeDataExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of PaymentFeeDataExt
func (u PaymentFeeDataExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewPaymentFeeDataExt creates a new  PaymentFeeDataExt.
func NewPaymentFeeDataExt(v LedgerVersion, value interface{}) (result PaymentFeeDataExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// PaymentFeeData is an XDR Struct defines as:
//
//   struct PaymentFeeData {
//        FeeData sourceFee;
//        FeeData destinationFee;
//        bool sourcePaysForDest;    // if true source account pays fee, else destination
//    	// reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type PaymentFeeData struct {
	SourceFee         FeeData           `json:"sourceFee,omitempty"`
	DestinationFee    FeeData           `json:"destinationFee,omitempty"`
	SourcePaysForDest bool              `json:"sourcePaysForDest,omitempty"`
	Ext               PaymentFeeDataExt `json:"ext,omitempty"`
}

// PaymentOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type PaymentOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u PaymentOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of PaymentOpExt
func (u PaymentOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewPaymentOpExt creates a new  PaymentOpExt.
func NewPaymentOpExt(v LedgerVersion, value interface{}) (result PaymentOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// PaymentOp is an XDR Struct defines as:
//
//   struct PaymentOp
//    {
//        BalanceID sourceBalanceID;
//        BalanceID destinationBalanceID;
//        int64 amount;          // amount they end up with
//
//        PaymentFeeData feeData;
//
//        string256 subject;
//        string64 reference;
//
//        InvoiceReference* invoiceReference;
//
//    	// reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type PaymentOp struct {
	SourceBalanceId      BalanceId         `json:"sourceBalanceID,omitempty"`
	DestinationBalanceId BalanceId         `json:"destinationBalanceID,omitempty"`
	Amount               Int64             `json:"amount,omitempty"`
	FeeData              PaymentFeeData    `json:"feeData,omitempty"`
	Subject              String256         `json:"subject,omitempty"`
	Reference            String64          `json:"reference,omitempty"`
	InvoiceReference     *InvoiceReference `json:"invoiceReference,omitempty"`
	Ext                  PaymentOpExt      `json:"ext,omitempty"`
}

// PaymentResultCode is an XDR Enum defines as:
//
//   enum PaymentResultCode
//    {
//        // codes considered as "success" for the operation
//        SUCCESS = 0, // payment successfuly completed
//
//        // codes considered as "failure" for the operation
//        MALFORMED = -1,       // bad input
//        UNDERFUNDED = -2,     // not enough funds in source account
//        LINE_FULL = -3,       // destination would go above their limit
//    	FEE_MISMATCHED = -4,   // fee is not equal to expected fee
//        BALANCE_NOT_FOUND = -5, // destination balance not found
//        BALANCE_ACCOUNT_MISMATCHED = -6,
//        BALANCE_ASSETS_MISMATCHED = -7,
//    	SRC_BALANCE_NOT_FOUND = -8, // source balance not found
//        REFERENCE_DUPLICATION = -9,
//        STATS_OVERFLOW = -10,
//        LIMITS_EXCEEDED = -11,
//        NOT_ALLOWED_BY_ASSET_POLICY = -12,
//        INVOICE_NOT_FOUND = -13,
//        INVOICE_WRONG_AMOUNT = -14,
//        INVOICE_BALANCE_MISMATCH = -15,
//        INVOICE_ACCOUNT_MISMATCH = -16,
//        INVOICE_ALREADY_PAID = -17,
//        PAYMENT_V1_NO_LONGER_SUPPORTED = -18
//    };
//
type PaymentResultCode int32

const (
	PaymentResultCodeSuccess                    PaymentResultCode = 0
	PaymentResultCodeMalformed                  PaymentResultCode = -1
	PaymentResultCodeUnderfunded                PaymentResultCode = -2
	PaymentResultCodeLineFull                   PaymentResultCode = -3
	PaymentResultCodeFeeMismatched              PaymentResultCode = -4
	PaymentResultCodeBalanceNotFound            PaymentResultCode = -5
	PaymentResultCodeBalanceAccountMismatched   PaymentResultCode = -6
	PaymentResultCodeBalanceAssetsMismatched    PaymentResultCode = -7
	PaymentResultCodeSrcBalanceNotFound         PaymentResultCode = -8
	PaymentResultCodeReferenceDuplication       PaymentResultCode = -9
	PaymentResultCodeStatsOverflow              PaymentResultCode = -10
	PaymentResultCodeLimitsExceeded             PaymentResultCode = -11
	PaymentResultCodeNotAllowedByAssetPolicy    PaymentResultCode = -12
	PaymentResultCodeInvoiceNotFound            PaymentResultCode = -13
	PaymentResultCodeInvoiceWrongAmount         PaymentResultCode = -14
	PaymentResultCodeInvoiceBalanceMismatch     PaymentResultCode = -15
	PaymentResultCodeInvoiceAccountMismatch     PaymentResultCode = -16
	PaymentResultCodeInvoiceAlreadyPaid         PaymentResultCode = -17
	PaymentResultCodePaymentV1NoLongerSupported PaymentResultCode = -18
)

var PaymentResultCodeAll = []PaymentResultCode{
	PaymentResultCodeSuccess,
	PaymentResultCodeMalformed,
	PaymentResultCodeUnderfunded,
	PaymentResultCodeLineFull,
	PaymentResultCodeFeeMismatched,
	PaymentResultCodeBalanceNotFound,
	PaymentResultCodeBalanceAccountMismatched,
	PaymentResultCodeBalanceAssetsMismatched,
	PaymentResultCodeSrcBalanceNotFound,
	PaymentResultCodeReferenceDuplication,
	PaymentResultCodeStatsOverflow,
	PaymentResultCodeLimitsExceeded,
	PaymentResultCodeNotAllowedByAssetPolicy,
	PaymentResultCodeInvoiceNotFound,
	PaymentResultCodeInvoiceWrongAmount,
	PaymentResultCodeInvoiceBalanceMismatch,
	PaymentResultCodeInvoiceAccountMismatch,
	PaymentResultCodeInvoiceAlreadyPaid,
	PaymentResultCodePaymentV1NoLongerSupported,
}

var paymentResultCodeMap = map[int32]string{
	0:   "PaymentResultCodeSuccess",
	-1:  "PaymentResultCodeMalformed",
	-2:  "PaymentResultCodeUnderfunded",
	-3:  "PaymentResultCodeLineFull",
	-4:  "PaymentResultCodeFeeMismatched",
	-5:  "PaymentResultCodeBalanceNotFound",
	-6:  "PaymentResultCodeBalanceAccountMismatched",
	-7:  "PaymentResultCodeBalanceAssetsMismatched",
	-8:  "PaymentResultCodeSrcBalanceNotFound",
	-9:  "PaymentResultCodeReferenceDuplication",
	-10: "PaymentResultCodeStatsOverflow",
	-11: "PaymentResultCodeLimitsExceeded",
	-12: "PaymentResultCodeNotAllowedByAssetPolicy",
	-13: "PaymentResultCodeInvoiceNotFound",
	-14: "PaymentResultCodeInvoiceWrongAmount",
	-15: "PaymentResultCodeInvoiceBalanceMismatch",
	-16: "PaymentResultCodeInvoiceAccountMismatch",
	-17: "PaymentResultCodeInvoiceAlreadyPaid",
	-18: "PaymentResultCodePaymentV1NoLongerSupported",
}

var paymentResultCodeShortMap = map[int32]string{
	0:   "success",
	-1:  "malformed",
	-2:  "underfunded",
	-3:  "line_full",
	-4:  "fee_mismatched",
	-5:  "balance_not_found",
	-6:  "balance_account_mismatched",
	-7:  "balance_assets_mismatched",
	-8:  "src_balance_not_found",
	-9:  "reference_duplication",
	-10: "stats_overflow",
	-11: "limits_exceeded",
	-12: "not_allowed_by_asset_policy",
	-13: "invoice_not_found",
	-14: "invoice_wrong_amount",
	-15: "invoice_balance_mismatch",
	-16: "invoice_account_mismatch",
	-17: "invoice_already_paid",
	-18: "payment_v1_no_longer_supported",
}

var paymentResultCodeRevMap = map[string]int32{
	"PaymentResultCodeSuccess":                    0,
	"PaymentResultCodeMalformed":                  -1,
	"PaymentResultCodeUnderfunded":                -2,
	"PaymentResultCodeLineFull":                   -3,
	"PaymentResultCodeFeeMismatched":              -4,
	"PaymentResultCodeBalanceNotFound":            -5,
	"PaymentResultCodeBalanceAccountMismatched":   -6,
	"PaymentResultCodeBalanceAssetsMismatched":    -7,
	"PaymentResultCodeSrcBalanceNotFound":         -8,
	"PaymentResultCodeReferenceDuplication":       -9,
	"PaymentResultCodeStatsOverflow":              -10,
	"PaymentResultCodeLimitsExceeded":             -11,
	"PaymentResultCodeNotAllowedByAssetPolicy":    -12,
	"PaymentResultCodeInvoiceNotFound":            -13,
	"PaymentResultCodeInvoiceWrongAmount":         -14,
	"PaymentResultCodeInvoiceBalanceMismatch":     -15,
	"PaymentResultCodeInvoiceAccountMismatch":     -16,
	"PaymentResultCodeInvoiceAlreadyPaid":         -17,
	"PaymentResultCodePaymentV1NoLongerSupported": -18,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for PaymentResultCode
func (e PaymentResultCode) ValidEnum(v int32) bool {
	_, ok := paymentResultCodeMap[v]
	return ok
}
func (e PaymentResultCode) isFlag() bool {
	for i := len(PaymentResultCodeAll) - 1; i >= 0; i-- {
		expected := PaymentResultCode(2) << uint64(len(PaymentResultCodeAll)-1) >> uint64(len(PaymentResultCodeAll)-i)
		if expected != PaymentResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e PaymentResultCode) String() string {
	name, _ := paymentResultCodeMap[int32(e)]
	return name
}

func (e PaymentResultCode) ShortString() string {
	name, _ := paymentResultCodeShortMap[int32(e)]
	return name
}

func (e PaymentResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range PaymentResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *PaymentResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = PaymentResultCode(t.Value)
	return nil
}

// PaymentResponseExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type PaymentResponseExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u PaymentResponseExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of PaymentResponseExt
func (u PaymentResponseExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewPaymentResponseExt creates a new  PaymentResponseExt.
func NewPaymentResponseExt(v LedgerVersion, value interface{}) (result PaymentResponseExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// PaymentResponse is an XDR Struct defines as:
//
//   struct PaymentResponse {
//        AccountID destination;
//        uint64 paymentID;
//        AssetCode asset;
//        // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type PaymentResponse struct {
	Destination AccountId          `json:"destination,omitempty"`
	PaymentId   Uint64             `json:"paymentID,omitempty"`
	Asset       AssetCode          `json:"asset,omitempty"`
	Ext         PaymentResponseExt `json:"ext,omitempty"`
}

// PaymentResult is an XDR Union defines as:
//
//   union PaymentResult switch (PaymentResultCode code)
//    {
//    case SUCCESS:
//        PaymentResponse paymentResponse;
//    default:
//        void;
//    };
//
type PaymentResult struct {
	Code            PaymentResultCode `json:"code,omitempty"`
	PaymentResponse *PaymentResponse  `json:"paymentResponse,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u PaymentResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of PaymentResult
func (u PaymentResult) ArmForSwitch(sw int32) (string, bool) {
	switch PaymentResultCode(sw) {
	case PaymentResultCodeSuccess:
		return "PaymentResponse", true
	default:
		return "", true
	}
}

// NewPaymentResult creates a new  PaymentResult.
func NewPaymentResult(code PaymentResultCode, value interface{}) (result PaymentResult, err error) {
	result.Code = code
	switch PaymentResultCode(code) {
	case PaymentResultCodeSuccess:
		tv, ok := value.(PaymentResponse)
		if !ok {
			err = fmt.Errorf("invalid value, must be PaymentResponse")
			return
		}
		result.PaymentResponse = &tv
	default:
		// void
	}
	return
}

// MustPaymentResponse retrieves the PaymentResponse value from the union,
// panicing if the value is not set.
func (u PaymentResult) MustPaymentResponse() PaymentResponse {
	val, ok := u.GetPaymentResponse()

	if !ok {
		panic("arm PaymentResponse is not set")
	}

	return val
}

// GetPaymentResponse retrieves the PaymentResponse value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u PaymentResult) GetPaymentResponse() (result PaymentResponse, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "PaymentResponse" {
		result = *u.PaymentResponse
		ok = true
	}

	return
}

// ReviewRequestOpAction is an XDR Enum defines as:
//
//   enum ReviewRequestOpAction {
//    	APPROVE = 1,
//    	REJECT = 2,
//    	PERMANENT_REJECT = 3
//    };
//
type ReviewRequestOpAction int32

const (
	ReviewRequestOpActionApprove         ReviewRequestOpAction = 1
	ReviewRequestOpActionReject          ReviewRequestOpAction = 2
	ReviewRequestOpActionPermanentReject ReviewRequestOpAction = 3
)

var ReviewRequestOpActionAll = []ReviewRequestOpAction{
	ReviewRequestOpActionApprove,
	ReviewRequestOpActionReject,
	ReviewRequestOpActionPermanentReject,
}

var reviewRequestOpActionMap = map[int32]string{
	1: "ReviewRequestOpActionApprove",
	2: "ReviewRequestOpActionReject",
	3: "ReviewRequestOpActionPermanentReject",
}

var reviewRequestOpActionShortMap = map[int32]string{
	1: "approve",
	2: "reject",
	3: "permanent_reject",
}

var reviewRequestOpActionRevMap = map[string]int32{
	"ReviewRequestOpActionApprove":         1,
	"ReviewRequestOpActionReject":          2,
	"ReviewRequestOpActionPermanentReject": 3,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ReviewRequestOpAction
func (e ReviewRequestOpAction) ValidEnum(v int32) bool {
	_, ok := reviewRequestOpActionMap[v]
	return ok
}
func (e ReviewRequestOpAction) isFlag() bool {
	for i := len(ReviewRequestOpActionAll) - 1; i >= 0; i-- {
		expected := ReviewRequestOpAction(2) << uint64(len(ReviewRequestOpActionAll)-1) >> uint64(len(ReviewRequestOpActionAll)-i)
		if expected != ReviewRequestOpActionAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ReviewRequestOpAction) String() string {
	name, _ := reviewRequestOpActionMap[int32(e)]
	return name
}

func (e ReviewRequestOpAction) ShortString() string {
	name, _ := reviewRequestOpActionShortMap[int32(e)]
	return name
}

func (e ReviewRequestOpAction) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range ReviewRequestOpActionAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ReviewRequestOpAction) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ReviewRequestOpAction(t.Value)
	return nil
}

// LimitsUpdateDetailsExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type LimitsUpdateDetailsExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LimitsUpdateDetailsExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LimitsUpdateDetailsExt
func (u LimitsUpdateDetailsExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewLimitsUpdateDetailsExt creates a new  LimitsUpdateDetailsExt.
func NewLimitsUpdateDetailsExt(v LedgerVersion, value interface{}) (result LimitsUpdateDetailsExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// LimitsUpdateDetails is an XDR Struct defines as:
//
//   struct LimitsUpdateDetails {
//        LimitsV2Entry newLimitsV2;
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type LimitsUpdateDetails struct {
	NewLimitsV2 LimitsV2Entry          `json:"newLimitsV2,omitempty"`
	Ext         LimitsUpdateDetailsExt `json:"ext,omitempty"`
}

// WithdrawalDetailsExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type WithdrawalDetailsExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u WithdrawalDetailsExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of WithdrawalDetailsExt
func (u WithdrawalDetailsExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewWithdrawalDetailsExt creates a new  WithdrawalDetailsExt.
func NewWithdrawalDetailsExt(v LedgerVersion, value interface{}) (result WithdrawalDetailsExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// WithdrawalDetails is an XDR Struct defines as:
//
//   struct WithdrawalDetails {
//    	string externalDetails<>;
//    	// reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type WithdrawalDetails struct {
	ExternalDetails string               `json:"externalDetails,omitempty"`
	Ext             WithdrawalDetailsExt `json:"ext,omitempty"`
}

// AmlAlertDetailsExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type AmlAlertDetailsExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u AmlAlertDetailsExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of AmlAlertDetailsExt
func (u AmlAlertDetailsExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewAmlAlertDetailsExt creates a new  AmlAlertDetailsExt.
func NewAmlAlertDetailsExt(v LedgerVersion, value interface{}) (result AmlAlertDetailsExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// AmlAlertDetails is an XDR Struct defines as:
//
//   struct AMLAlertDetails {
//    	string comment<>;
//    	// reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type AmlAlertDetails struct {
	Comment string             `json:"comment,omitempty"`
	Ext     AmlAlertDetailsExt `json:"ext,omitempty"`
}

// UpdateKycDetailsExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type UpdateKycDetailsExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u UpdateKycDetailsExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of UpdateKycDetailsExt
func (u UpdateKycDetailsExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewUpdateKycDetailsExt creates a new  UpdateKycDetailsExt.
func NewUpdateKycDetailsExt(v LedgerVersion, value interface{}) (result UpdateKycDetailsExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// UpdateKycDetails is an XDR Struct defines as:
//
//   struct UpdateKYCDetails {
//        uint32 tasksToAdd;
//        uint32 tasksToRemove;
//        string externalDetails<>;
//        // Reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type UpdateKycDetails struct {
	TasksToAdd      Uint32              `json:"tasksToAdd,omitempty"`
	TasksToRemove   Uint32              `json:"tasksToRemove,omitempty"`
	ExternalDetails string              `json:"externalDetails,omitempty"`
	Ext             UpdateKycDetailsExt `json:"ext,omitempty"`
}

// ContractDetailsExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type ContractDetailsExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ContractDetailsExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ContractDetailsExt
func (u ContractDetailsExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewContractDetailsExt creates a new  ContractDetailsExt.
func NewContractDetailsExt(v LedgerVersion, value interface{}) (result ContractDetailsExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ContractDetails is an XDR Struct defines as:
//
//   struct ContractDetails {
//        longstring details;
//
//        // Reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type ContractDetails struct {
	Details Longstring         `json:"details,omitempty"`
	Ext     ContractDetailsExt `json:"ext,omitempty"`
}

// BillPayDetailsExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type BillPayDetailsExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u BillPayDetailsExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of BillPayDetailsExt
func (u BillPayDetailsExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewBillPayDetailsExt creates a new  BillPayDetailsExt.
func NewBillPayDetailsExt(v LedgerVersion, value interface{}) (result BillPayDetailsExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// BillPayDetails is an XDR Struct defines as:
//
//   struct BillPayDetails {
//        PaymentOpV2 paymentDetails;
//
//        // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type BillPayDetails struct {
	PaymentDetails PaymentOpV2       `json:"paymentDetails,omitempty"`
	Ext            BillPayDetailsExt `json:"ext,omitempty"`
}

// ReviewDetailsExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type ReviewDetailsExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ReviewDetailsExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ReviewDetailsExt
func (u ReviewDetailsExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewReviewDetailsExt creates a new  ReviewDetailsExt.
func NewReviewDetailsExt(v LedgerVersion, value interface{}) (result ReviewDetailsExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ReviewDetails is an XDR Struct defines as:
//
//   struct ReviewDetails {
//        uint32 tasksToAdd;
//        uint32 tasksToRemove;
//        string externalDetails<>;
//        // Reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type ReviewDetails struct {
	TasksToAdd      Uint32           `json:"tasksToAdd,omitempty"`
	TasksToRemove   Uint32           `json:"tasksToRemove,omitempty"`
	ExternalDetails string           `json:"externalDetails,omitempty"`
	Ext             ReviewDetailsExt `json:"ext,omitempty"`
}

// SaleExtendedExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type SaleExtendedExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u SaleExtendedExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of SaleExtendedExt
func (u SaleExtendedExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewSaleExtendedExt creates a new  SaleExtendedExt.
func NewSaleExtendedExt(v LedgerVersion, value interface{}) (result SaleExtendedExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// SaleExtended is an XDR Struct defines as:
//
//   struct SaleExtended {
//        uint64 saleID;
//
//        // Reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type SaleExtended struct {
	SaleId Uint64          `json:"saleID,omitempty"`
	Ext    SaleExtendedExt `json:"ext,omitempty"`
}

// ExtendedResultTypeExt is an XDR NestedUnion defines as:
//
//   union switch(ReviewableRequestType requestType) {
//        case SALE:
//            SaleExtended saleExtended;
//        case NONE:
//            void;
//        }
//
type ExtendedResultTypeExt struct {
	RequestType  ReviewableRequestType `json:"requestType,omitempty"`
	SaleExtended *SaleExtended         `json:"saleExtended,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ExtendedResultTypeExt) SwitchFieldName() string {
	return "RequestType"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ExtendedResultTypeExt
func (u ExtendedResultTypeExt) ArmForSwitch(sw int32) (string, bool) {
	switch ReviewableRequestType(sw) {
	case ReviewableRequestTypeSale:
		return "SaleExtended", true
	case ReviewableRequestTypeNone:
		return "", true
	}
	return "-", false
}

// NewExtendedResultTypeExt creates a new  ExtendedResultTypeExt.
func NewExtendedResultTypeExt(requestType ReviewableRequestType, value interface{}) (result ExtendedResultTypeExt, err error) {
	result.RequestType = requestType
	switch ReviewableRequestType(requestType) {
	case ReviewableRequestTypeSale:
		tv, ok := value.(SaleExtended)
		if !ok {
			err = fmt.Errorf("invalid value, must be SaleExtended")
			return
		}
		result.SaleExtended = &tv
	case ReviewableRequestTypeNone:
		// void
	}
	return
}

// MustSaleExtended retrieves the SaleExtended value from the union,
// panicing if the value is not set.
func (u ExtendedResultTypeExt) MustSaleExtended() SaleExtended {
	val, ok := u.GetSaleExtended()

	if !ok {
		panic("arm SaleExtended is not set")
	}

	return val
}

// GetSaleExtended retrieves the SaleExtended value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ExtendedResultTypeExt) GetSaleExtended() (result SaleExtended, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.RequestType))

	if armName == "SaleExtended" {
		result = *u.SaleExtended
		ok = true
	}

	return
}

// ExtendedResultExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//       {
//       case EMPTY_VERSION:
//           void;
//       }
//
type ExtendedResultExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ExtendedResultExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ExtendedResultExt
func (u ExtendedResultExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewExtendedResultExt creates a new  ExtendedResultExt.
func NewExtendedResultExt(v LedgerVersion, value interface{}) (result ExtendedResultExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ExtendedResult is an XDR Struct defines as:
//
//   struct ExtendedResult {
//        bool fulfilled;
//
//        union switch(ReviewableRequestType requestType) {
//        case SALE:
//            SaleExtended saleExtended;
//        case NONE:
//            void;
//        } typeExt;
//
//       // Reserved for future use
//       union switch (LedgerVersion v)
//       {
//       case EMPTY_VERSION:
//           void;
//       }
//       ext;
//    };
//
type ExtendedResult struct {
	Fulfilled bool                  `json:"fulfilled,omitempty"`
	TypeExt   ExtendedResultTypeExt `json:"typeExt,omitempty"`
	Ext       ExtendedResultExt     `json:"ext,omitempty"`
}

// ReviewRequestOpRequestDetails is an XDR NestedUnion defines as:
//
//   union switch(ReviewableRequestType requestType) {
//    	case WITHDRAW:
//    		WithdrawalDetails withdrawal;
//        case LIMITS_UPDATE:
//            LimitsUpdateDetails limitsUpdate;
//    	case TWO_STEP_WITHDRAWAL:
//    		WithdrawalDetails twoStepWithdrawal;
//        case AML_ALERT:
//            AMLAlertDetails amlAlertDetails;
//        case UPDATE_KYC:
//            UpdateKYCDetails updateKYC;
//        case INVOICE:
//            BillPayDetails billPay;
//        case CONTRACT:
//            ContractDetails contract;
//    	default:
//    		void;
//    	}
//
type ReviewRequestOpRequestDetails struct {
	RequestType       ReviewableRequestType `json:"requestType,omitempty"`
	Withdrawal        *WithdrawalDetails    `json:"withdrawal,omitempty"`
	LimitsUpdate      *LimitsUpdateDetails  `json:"limitsUpdate,omitempty"`
	TwoStepWithdrawal *WithdrawalDetails    `json:"twoStepWithdrawal,omitempty"`
	AmlAlertDetails   *AmlAlertDetails      `json:"amlAlertDetails,omitempty"`
	UpdateKyc         *UpdateKycDetails     `json:"updateKYC,omitempty"`
	BillPay           *BillPayDetails       `json:"billPay,omitempty"`
	Contract          *ContractDetails      `json:"contract,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ReviewRequestOpRequestDetails) SwitchFieldName() string {
	return "RequestType"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ReviewRequestOpRequestDetails
func (u ReviewRequestOpRequestDetails) ArmForSwitch(sw int32) (string, bool) {
	switch ReviewableRequestType(sw) {
	case ReviewableRequestTypeWithdraw:
		return "Withdrawal", true
	case ReviewableRequestTypeLimitsUpdate:
		return "LimitsUpdate", true
	case ReviewableRequestTypeTwoStepWithdrawal:
		return "TwoStepWithdrawal", true
	case ReviewableRequestTypeAmlAlert:
		return "AmlAlertDetails", true
	case ReviewableRequestTypeUpdateKyc:
		return "UpdateKyc", true
	case ReviewableRequestTypeInvoice:
		return "BillPay", true
	case ReviewableRequestTypeContract:
		return "Contract", true
	default:
		return "", true
	}
}

// NewReviewRequestOpRequestDetails creates a new  ReviewRequestOpRequestDetails.
func NewReviewRequestOpRequestDetails(requestType ReviewableRequestType, value interface{}) (result ReviewRequestOpRequestDetails, err error) {
	result.RequestType = requestType
	switch ReviewableRequestType(requestType) {
	case ReviewableRequestTypeWithdraw:
		tv, ok := value.(WithdrawalDetails)
		if !ok {
			err = fmt.Errorf("invalid value, must be WithdrawalDetails")
			return
		}
		result.Withdrawal = &tv
	case ReviewableRequestTypeLimitsUpdate:
		tv, ok := value.(LimitsUpdateDetails)
		if !ok {
			err = fmt.Errorf("invalid value, must be LimitsUpdateDetails")
			return
		}
		result.LimitsUpdate = &tv
	case ReviewableRequestTypeTwoStepWithdrawal:
		tv, ok := value.(WithdrawalDetails)
		if !ok {
			err = fmt.Errorf("invalid value, must be WithdrawalDetails")
			return
		}
		result.TwoStepWithdrawal = &tv
	case ReviewableRequestTypeAmlAlert:
		tv, ok := value.(AmlAlertDetails)
		if !ok {
			err = fmt.Errorf("invalid value, must be AmlAlertDetails")
			return
		}
		result.AmlAlertDetails = &tv
	case ReviewableRequestTypeUpdateKyc:
		tv, ok := value.(UpdateKycDetails)
		if !ok {
			err = fmt.Errorf("invalid value, must be UpdateKycDetails")
			return
		}
		result.UpdateKyc = &tv
	case ReviewableRequestTypeInvoice:
		tv, ok := value.(BillPayDetails)
		if !ok {
			err = fmt.Errorf("invalid value, must be BillPayDetails")
			return
		}
		result.BillPay = &tv
	case ReviewableRequestTypeContract:
		tv, ok := value.(ContractDetails)
		if !ok {
			err = fmt.Errorf("invalid value, must be ContractDetails")
			return
		}
		result.Contract = &tv
	default:
		// void
	}
	return
}

// MustWithdrawal retrieves the Withdrawal value from the union,
// panicing if the value is not set.
func (u ReviewRequestOpRequestDetails) MustWithdrawal() WithdrawalDetails {
	val, ok := u.GetWithdrawal()

	if !ok {
		panic("arm Withdrawal is not set")
	}

	return val
}

// GetWithdrawal retrieves the Withdrawal value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ReviewRequestOpRequestDetails) GetWithdrawal() (result WithdrawalDetails, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.RequestType))

	if armName == "Withdrawal" {
		result = *u.Withdrawal
		ok = true
	}

	return
}

// MustLimitsUpdate retrieves the LimitsUpdate value from the union,
// panicing if the value is not set.
func (u ReviewRequestOpRequestDetails) MustLimitsUpdate() LimitsUpdateDetails {
	val, ok := u.GetLimitsUpdate()

	if !ok {
		panic("arm LimitsUpdate is not set")
	}

	return val
}

// GetLimitsUpdate retrieves the LimitsUpdate value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ReviewRequestOpRequestDetails) GetLimitsUpdate() (result LimitsUpdateDetails, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.RequestType))

	if armName == "LimitsUpdate" {
		result = *u.LimitsUpdate
		ok = true
	}

	return
}

// MustTwoStepWithdrawal retrieves the TwoStepWithdrawal value from the union,
// panicing if the value is not set.
func (u ReviewRequestOpRequestDetails) MustTwoStepWithdrawal() WithdrawalDetails {
	val, ok := u.GetTwoStepWithdrawal()

	if !ok {
		panic("arm TwoStepWithdrawal is not set")
	}

	return val
}

// GetTwoStepWithdrawal retrieves the TwoStepWithdrawal value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ReviewRequestOpRequestDetails) GetTwoStepWithdrawal() (result WithdrawalDetails, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.RequestType))

	if armName == "TwoStepWithdrawal" {
		result = *u.TwoStepWithdrawal
		ok = true
	}

	return
}

// MustAmlAlertDetails retrieves the AmlAlertDetails value from the union,
// panicing if the value is not set.
func (u ReviewRequestOpRequestDetails) MustAmlAlertDetails() AmlAlertDetails {
	val, ok := u.GetAmlAlertDetails()

	if !ok {
		panic("arm AmlAlertDetails is not set")
	}

	return val
}

// GetAmlAlertDetails retrieves the AmlAlertDetails value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ReviewRequestOpRequestDetails) GetAmlAlertDetails() (result AmlAlertDetails, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.RequestType))

	if armName == "AmlAlertDetails" {
		result = *u.AmlAlertDetails
		ok = true
	}

	return
}

// MustUpdateKyc retrieves the UpdateKyc value from the union,
// panicing if the value is not set.
func (u ReviewRequestOpRequestDetails) MustUpdateKyc() UpdateKycDetails {
	val, ok := u.GetUpdateKyc()

	if !ok {
		panic("arm UpdateKyc is not set")
	}

	return val
}

// GetUpdateKyc retrieves the UpdateKyc value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ReviewRequestOpRequestDetails) GetUpdateKyc() (result UpdateKycDetails, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.RequestType))

	if armName == "UpdateKyc" {
		result = *u.UpdateKyc
		ok = true
	}

	return
}

// MustBillPay retrieves the BillPay value from the union,
// panicing if the value is not set.
func (u ReviewRequestOpRequestDetails) MustBillPay() BillPayDetails {
	val, ok := u.GetBillPay()

	if !ok {
		panic("arm BillPay is not set")
	}

	return val
}

// GetBillPay retrieves the BillPay value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ReviewRequestOpRequestDetails) GetBillPay() (result BillPayDetails, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.RequestType))

	if armName == "BillPay" {
		result = *u.BillPay
		ok = true
	}

	return
}

// MustContract retrieves the Contract value from the union,
// panicing if the value is not set.
func (u ReviewRequestOpRequestDetails) MustContract() ContractDetails {
	val, ok := u.GetContract()

	if !ok {
		panic("arm Contract is not set")
	}

	return val
}

// GetContract retrieves the Contract value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ReviewRequestOpRequestDetails) GetContract() (result ContractDetails, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.RequestType))

	if armName == "Contract" {
		result = *u.Contract
		ok = true
	}

	return
}

// ReviewRequestOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        case ADD_TASKS_TO_REVIEWABLE_REQUEST:
//            ReviewDetails reviewDetails;
//        }
//
type ReviewRequestOpExt struct {
	V             LedgerVersion  `json:"v,omitempty"`
	ReviewDetails *ReviewDetails `json:"reviewDetails,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ReviewRequestOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ReviewRequestOpExt
func (u ReviewRequestOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	case LedgerVersionAddTasksToReviewableRequest:
		return "ReviewDetails", true
	}
	return "-", false
}

// NewReviewRequestOpExt creates a new  ReviewRequestOpExt.
func NewReviewRequestOpExt(v LedgerVersion, value interface{}) (result ReviewRequestOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	case LedgerVersionAddTasksToReviewableRequest:
		tv, ok := value.(ReviewDetails)
		if !ok {
			err = fmt.Errorf("invalid value, must be ReviewDetails")
			return
		}
		result.ReviewDetails = &tv
	}
	return
}

// MustReviewDetails retrieves the ReviewDetails value from the union,
// panicing if the value is not set.
func (u ReviewRequestOpExt) MustReviewDetails() ReviewDetails {
	val, ok := u.GetReviewDetails()

	if !ok {
		panic("arm ReviewDetails is not set")
	}

	return val
}

// GetReviewDetails retrieves the ReviewDetails value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ReviewRequestOpExt) GetReviewDetails() (result ReviewDetails, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.V))

	if armName == "ReviewDetails" {
		result = *u.ReviewDetails
		ok = true
	}

	return
}

// ReviewRequestOp is an XDR Struct defines as:
//
//   struct ReviewRequestOp
//    {
//    	uint64 requestID;
//    	Hash requestHash;
//    	union switch(ReviewableRequestType requestType) {
//    	case WITHDRAW:
//    		WithdrawalDetails withdrawal;
//        case LIMITS_UPDATE:
//            LimitsUpdateDetails limitsUpdate;
//    	case TWO_STEP_WITHDRAWAL:
//    		WithdrawalDetails twoStepWithdrawal;
//        case AML_ALERT:
//            AMLAlertDetails amlAlertDetails;
//        case UPDATE_KYC:
//            UpdateKYCDetails updateKYC;
//        case INVOICE:
//            BillPayDetails billPay;
//        case CONTRACT:
//            ContractDetails contract;
//    	default:
//    		void;
//    	} requestDetails;
//    	ReviewRequestOpAction action;
//    	longstring reason;
//    	// reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        case ADD_TASKS_TO_REVIEWABLE_REQUEST:
//            ReviewDetails reviewDetails;
//        }
//        ext;
//    };
//
type ReviewRequestOp struct {
	RequestId      Uint64                        `json:"requestID,omitempty"`
	RequestHash    Hash                          `json:"requestHash,omitempty"`
	RequestDetails ReviewRequestOpRequestDetails `json:"requestDetails,omitempty"`
	Action         ReviewRequestOpAction         `json:"action,omitempty"`
	Reason         Longstring                    `json:"reason,omitempty"`
	Ext            ReviewRequestOpExt            `json:"ext,omitempty"`
}

// ReviewRequestResultCode is an XDR Enum defines as:
//
//   enum ReviewRequestResultCode
//    {
//        // Codes considered as "success" for the operation
//        SUCCESS = 0,
//
//        // Codes considered as "failure" for the operation
//        INVALID_REASON = -1,        // reason must be empty if approving and not empty if rejecting
//    	INVALID_ACTION = -2,
//    	HASH_MISMATCHED = -3,
//    	NOT_FOUND = -4,
//    	TYPE_MISMATCHED = -5,
//    	REJECT_NOT_ALLOWED = -6, // reject not allowed, use permanent reject
//    	INVALID_EXTERNAL_DETAILS = -7,
//    	REQUESTOR_IS_BLOCKED = -8,
//    	PERMANENT_REJECT_NOT_ALLOWED = -9, // permanent reject not allowed, use reject
//
//    	// Asset requests
//    	ASSET_ALREADY_EXISTS = -20,
//    	ASSET_DOES_NOT_EXISTS = -21,
//
//    	// Issuance requests
//    	MAX_ISSUANCE_AMOUNT_EXCEEDED = -40,
//    	INSUFFICIENT_AVAILABLE_FOR_ISSUANCE_AMOUNT = -41,
//    	FULL_LINE = -42, // can't fund balance - total funds exceed UINT64_MAX
//    	SYSTEM_TASKS_NOT_ALLOWED = -43,
//
//    	// Sale creation requests
//    	BASE_ASSET_DOES_NOT_EXISTS = -50,
//    	HARD_CAP_WILL_EXCEED_MAX_ISSUANCE = -51,
//    	INSUFFICIENT_PREISSUED_FOR_HARD_CAP = -52,
//
//    	// Update KYC requests
//    	NON_ZERO_TASKS_TO_REMOVE_NOT_ALLOWED = -60,
//
//    	// Update sale details, end time and promotion requests
//    	SALE_NOT_FOUND = -70,
//
//    	// Promotion update requests
//    	INVALID_SALE_STATE = -80, // sale state must be "PROMOTION"
//
//    	// Update sale end time requests
//        INVALID_SALE_NEW_END_TIME = -90, // new end time is before start time or current ledger close time
//
//        // Invoice requests
//        AMOUNT_MISMATCHED = -101, // amount does not match
//        DESTINATION_BALANCE_MISMATCHED = -102, // invoice balance and payment balance do not match
//        NOT_ALLOWED_ACCOUNT_DESTINATION = -103,
//        REQUIRED_SOURCE_PAY_FOR_DESTINATION = -104, // not allowed shift fee responsibility to destination
//        SOURCE_BALANCE_MISMATCHED = -105, // source balance must match invoice sender account
//        CONTRACT_NOT_FOUND = -106,
//        INVOICE_RECEIVER_BALANCE_LOCK_AMOUNT_OVERFLOW = -107,
//        INVOICE_ALREADY_APPROVED = -108,
//
//        // codes considered as "failure" for the payment operation
//        PAYMENT_V2_MALFORMED = -110, // bad input, requestID must be > 0
//        UNDERFUNDED = -111, // not enough funds in source account
//        LINE_FULL = -112, // destination would go above their limit
//        DESTINATION_BALANCE_NOT_FOUND = -113,
//        BALANCE_ASSETS_MISMATCHED = -114,
//        SRC_BALANCE_NOT_FOUND = -115, // source balance not found
//        REFERENCE_DUPLICATION = -116,
//        STATS_OVERFLOW = -117,
//        LIMITS_EXCEEDED = -118,
//        NOT_ALLOWED_BY_ASSET_POLICY = -119,
//        INVALID_DESTINATION_FEE = -120,
//        INVALID_DESTINATION_FEE_ASSET = -121, // destination fee asset must be the same as source balance asset
//        FEE_ASSET_MISMATCHED = -122,
//        INSUFFICIENT_FEE_AMOUNT = -123,
//        BALANCE_TO_CHARGE_FEE_FROM_NOT_FOUND = -124,
//        PAYMENT_AMOUNT_IS_LESS_THAN_DEST_FEE = -125,
//        DESTINATION_ACCOUNT_NOT_FOUND = -126,
//
//        // Limits update requests
//        CANNOT_CREATE_FOR_ACC_ID_AND_ACC_TYPE = 130, // limits cannot be created for account ID and account type simultaneously
//        INVALID_LIMITS = 131,
//
//        // Contract requests
//        CONTRACT_DETAILS_TOO_LONG = -140 // customer details reached length limit
//    };
//
type ReviewRequestResultCode int32

const (
	ReviewRequestResultCodeSuccess                                  ReviewRequestResultCode = 0
	ReviewRequestResultCodeInvalidReason                            ReviewRequestResultCode = -1
	ReviewRequestResultCodeInvalidAction                            ReviewRequestResultCode = -2
	ReviewRequestResultCodeHashMismatched                           ReviewRequestResultCode = -3
	ReviewRequestResultCodeNotFound                                 ReviewRequestResultCode = -4
	ReviewRequestResultCodeTypeMismatched                           ReviewRequestResultCode = -5
	ReviewRequestResultCodeRejectNotAllowed                         ReviewRequestResultCode = -6
	ReviewRequestResultCodeInvalidExternalDetails                   ReviewRequestResultCode = -7
	ReviewRequestResultCodeRequestorIsBlocked                       ReviewRequestResultCode = -8
	ReviewRequestResultCodePermanentRejectNotAllowed                ReviewRequestResultCode = -9
	ReviewRequestResultCodeAssetAlreadyExists                       ReviewRequestResultCode = -20
	ReviewRequestResultCodeAssetDoesNotExists                       ReviewRequestResultCode = -21
	ReviewRequestResultCodeMaxIssuanceAmountExceeded                ReviewRequestResultCode = -40
	ReviewRequestResultCodeInsufficientAvailableForIssuanceAmount   ReviewRequestResultCode = -41
	ReviewRequestResultCodeFullLine                                 ReviewRequestResultCode = -42
	ReviewRequestResultCodeSystemTasksNotAllowed                    ReviewRequestResultCode = -43
	ReviewRequestResultCodeBaseAssetDoesNotExists                   ReviewRequestResultCode = -50
	ReviewRequestResultCodeHardCapWillExceedMaxIssuance             ReviewRequestResultCode = -51
	ReviewRequestResultCodeInsufficientPreissuedForHardCap          ReviewRequestResultCode = -52
	ReviewRequestResultCodeNonZeroTasksToRemoveNotAllowed           ReviewRequestResultCode = -60
	ReviewRequestResultCodeSaleNotFound                             ReviewRequestResultCode = -70
	ReviewRequestResultCodeInvalidSaleState                         ReviewRequestResultCode = -80
	ReviewRequestResultCodeInvalidSaleNewEndTime                    ReviewRequestResultCode = -90
	ReviewRequestResultCodeAmountMismatched                         ReviewRequestResultCode = -101
	ReviewRequestResultCodeDestinationBalanceMismatched             ReviewRequestResultCode = -102
	ReviewRequestResultCodeNotAllowedAccountDestination             ReviewRequestResultCode = -103
	ReviewRequestResultCodeRequiredSourcePayForDestination          ReviewRequestResultCode = -104
	ReviewRequestResultCodeSourceBalanceMismatched                  ReviewRequestResultCode = -105
	ReviewRequestResultCodeContractNotFound                         ReviewRequestResultCode = -106
	ReviewRequestResultCodeInvoiceReceiverBalanceLockAmountOverflow ReviewRequestResultCode = -107
	ReviewRequestResultCodeInvoiceAlreadyApproved                   ReviewRequestResultCode = -108
	ReviewRequestResultCodePaymentV2Malformed                       ReviewRequestResultCode = -110
	ReviewRequestResultCodeUnderfunded                              ReviewRequestResultCode = -111
	ReviewRequestResultCodeLineFull                                 ReviewRequestResultCode = -112
	ReviewRequestResultCodeDestinationBalanceNotFound               ReviewRequestResultCode = -113
	ReviewRequestResultCodeBalanceAssetsMismatched                  ReviewRequestResultCode = -114
	ReviewRequestResultCodeSrcBalanceNotFound                       ReviewRequestResultCode = -115
	ReviewRequestResultCodeReferenceDuplication                     ReviewRequestResultCode = -116
	ReviewRequestResultCodeStatsOverflow                            ReviewRequestResultCode = -117
	ReviewRequestResultCodeLimitsExceeded                           ReviewRequestResultCode = -118
	ReviewRequestResultCodeNotAllowedByAssetPolicy                  ReviewRequestResultCode = -119
	ReviewRequestResultCodeInvalidDestinationFee                    ReviewRequestResultCode = -120
	ReviewRequestResultCodeInvalidDestinationFeeAsset               ReviewRequestResultCode = -121
	ReviewRequestResultCodeFeeAssetMismatched                       ReviewRequestResultCode = -122
	ReviewRequestResultCodeInsufficientFeeAmount                    ReviewRequestResultCode = -123
	ReviewRequestResultCodeBalanceToChargeFeeFromNotFound           ReviewRequestResultCode = -124
	ReviewRequestResultCodePaymentAmountIsLessThanDestFee           ReviewRequestResultCode = -125
	ReviewRequestResultCodeDestinationAccountNotFound               ReviewRequestResultCode = -126
	ReviewRequestResultCodeCannotCreateForAccIdAndAccType           ReviewRequestResultCode = 130
	ReviewRequestResultCodeInvalidLimits                            ReviewRequestResultCode = 131
	ReviewRequestResultCodeContractDetailsTooLong                   ReviewRequestResultCode = -140
)

var ReviewRequestResultCodeAll = []ReviewRequestResultCode{
	ReviewRequestResultCodeSuccess,
	ReviewRequestResultCodeInvalidReason,
	ReviewRequestResultCodeInvalidAction,
	ReviewRequestResultCodeHashMismatched,
	ReviewRequestResultCodeNotFound,
	ReviewRequestResultCodeTypeMismatched,
	ReviewRequestResultCodeRejectNotAllowed,
	ReviewRequestResultCodeInvalidExternalDetails,
	ReviewRequestResultCodeRequestorIsBlocked,
	ReviewRequestResultCodePermanentRejectNotAllowed,
	ReviewRequestResultCodeAssetAlreadyExists,
	ReviewRequestResultCodeAssetDoesNotExists,
	ReviewRequestResultCodeMaxIssuanceAmountExceeded,
	ReviewRequestResultCodeInsufficientAvailableForIssuanceAmount,
	ReviewRequestResultCodeFullLine,
	ReviewRequestResultCodeSystemTasksNotAllowed,
	ReviewRequestResultCodeBaseAssetDoesNotExists,
	ReviewRequestResultCodeHardCapWillExceedMaxIssuance,
	ReviewRequestResultCodeInsufficientPreissuedForHardCap,
	ReviewRequestResultCodeNonZeroTasksToRemoveNotAllowed,
	ReviewRequestResultCodeSaleNotFound,
	ReviewRequestResultCodeInvalidSaleState,
	ReviewRequestResultCodeInvalidSaleNewEndTime,
	ReviewRequestResultCodeAmountMismatched,
	ReviewRequestResultCodeDestinationBalanceMismatched,
	ReviewRequestResultCodeNotAllowedAccountDestination,
	ReviewRequestResultCodeRequiredSourcePayForDestination,
	ReviewRequestResultCodeSourceBalanceMismatched,
	ReviewRequestResultCodeContractNotFound,
	ReviewRequestResultCodeInvoiceReceiverBalanceLockAmountOverflow,
	ReviewRequestResultCodeInvoiceAlreadyApproved,
	ReviewRequestResultCodePaymentV2Malformed,
	ReviewRequestResultCodeUnderfunded,
	ReviewRequestResultCodeLineFull,
	ReviewRequestResultCodeDestinationBalanceNotFound,
	ReviewRequestResultCodeBalanceAssetsMismatched,
	ReviewRequestResultCodeSrcBalanceNotFound,
	ReviewRequestResultCodeReferenceDuplication,
	ReviewRequestResultCodeStatsOverflow,
	ReviewRequestResultCodeLimitsExceeded,
	ReviewRequestResultCodeNotAllowedByAssetPolicy,
	ReviewRequestResultCodeInvalidDestinationFee,
	ReviewRequestResultCodeInvalidDestinationFeeAsset,
	ReviewRequestResultCodeFeeAssetMismatched,
	ReviewRequestResultCodeInsufficientFeeAmount,
	ReviewRequestResultCodeBalanceToChargeFeeFromNotFound,
	ReviewRequestResultCodePaymentAmountIsLessThanDestFee,
	ReviewRequestResultCodeDestinationAccountNotFound,
	ReviewRequestResultCodeCannotCreateForAccIdAndAccType,
	ReviewRequestResultCodeInvalidLimits,
	ReviewRequestResultCodeContractDetailsTooLong,
}

var reviewRequestResultCodeMap = map[int32]string{
	0:    "ReviewRequestResultCodeSuccess",
	-1:   "ReviewRequestResultCodeInvalidReason",
	-2:   "ReviewRequestResultCodeInvalidAction",
	-3:   "ReviewRequestResultCodeHashMismatched",
	-4:   "ReviewRequestResultCodeNotFound",
	-5:   "ReviewRequestResultCodeTypeMismatched",
	-6:   "ReviewRequestResultCodeRejectNotAllowed",
	-7:   "ReviewRequestResultCodeInvalidExternalDetails",
	-8:   "ReviewRequestResultCodeRequestorIsBlocked",
	-9:   "ReviewRequestResultCodePermanentRejectNotAllowed",
	-20:  "ReviewRequestResultCodeAssetAlreadyExists",
	-21:  "ReviewRequestResultCodeAssetDoesNotExists",
	-40:  "ReviewRequestResultCodeMaxIssuanceAmountExceeded",
	-41:  "ReviewRequestResultCodeInsufficientAvailableForIssuanceAmount",
	-42:  "ReviewRequestResultCodeFullLine",
	-43:  "ReviewRequestResultCodeSystemTasksNotAllowed",
	-50:  "ReviewRequestResultCodeBaseAssetDoesNotExists",
	-51:  "ReviewRequestResultCodeHardCapWillExceedMaxIssuance",
	-52:  "ReviewRequestResultCodeInsufficientPreissuedForHardCap",
	-60:  "ReviewRequestResultCodeNonZeroTasksToRemoveNotAllowed",
	-70:  "ReviewRequestResultCodeSaleNotFound",
	-80:  "ReviewRequestResultCodeInvalidSaleState",
	-90:  "ReviewRequestResultCodeInvalidSaleNewEndTime",
	-101: "ReviewRequestResultCodeAmountMismatched",
	-102: "ReviewRequestResultCodeDestinationBalanceMismatched",
	-103: "ReviewRequestResultCodeNotAllowedAccountDestination",
	-104: "ReviewRequestResultCodeRequiredSourcePayForDestination",
	-105: "ReviewRequestResultCodeSourceBalanceMismatched",
	-106: "ReviewRequestResultCodeContractNotFound",
	-107: "ReviewRequestResultCodeInvoiceReceiverBalanceLockAmountOverflow",
	-108: "ReviewRequestResultCodeInvoiceAlreadyApproved",
	-110: "ReviewRequestResultCodePaymentV2Malformed",
	-111: "ReviewRequestResultCodeUnderfunded",
	-112: "ReviewRequestResultCodeLineFull",
	-113: "ReviewRequestResultCodeDestinationBalanceNotFound",
	-114: "ReviewRequestResultCodeBalanceAssetsMismatched",
	-115: "ReviewRequestResultCodeSrcBalanceNotFound",
	-116: "ReviewRequestResultCodeReferenceDuplication",
	-117: "ReviewRequestResultCodeStatsOverflow",
	-118: "ReviewRequestResultCodeLimitsExceeded",
	-119: "ReviewRequestResultCodeNotAllowedByAssetPolicy",
	-120: "ReviewRequestResultCodeInvalidDestinationFee",
	-121: "ReviewRequestResultCodeInvalidDestinationFeeAsset",
	-122: "ReviewRequestResultCodeFeeAssetMismatched",
	-123: "ReviewRequestResultCodeInsufficientFeeAmount",
	-124: "ReviewRequestResultCodeBalanceToChargeFeeFromNotFound",
	-125: "ReviewRequestResultCodePaymentAmountIsLessThanDestFee",
	-126: "ReviewRequestResultCodeDestinationAccountNotFound",
	130:  "ReviewRequestResultCodeCannotCreateForAccIdAndAccType",
	131:  "ReviewRequestResultCodeInvalidLimits",
	-140: "ReviewRequestResultCodeContractDetailsTooLong",
}

var reviewRequestResultCodeShortMap = map[int32]string{
	0:    "success",
	-1:   "invalid_reason",
	-2:   "invalid_action",
	-3:   "hash_mismatched",
	-4:   "not_found",
	-5:   "type_mismatched",
	-6:   "reject_not_allowed",
	-7:   "invalid_external_details",
	-8:   "requestor_is_blocked",
	-9:   "permanent_reject_not_allowed",
	-20:  "asset_already_exists",
	-21:  "asset_does_not_exists",
	-40:  "max_issuance_amount_exceeded",
	-41:  "insufficient_available_for_issuance_amount",
	-42:  "full_line",
	-43:  "system_tasks_not_allowed",
	-50:  "base_asset_does_not_exists",
	-51:  "hard_cap_will_exceed_max_issuance",
	-52:  "insufficient_preissued_for_hard_cap",
	-60:  "non_zero_tasks_to_remove_not_allowed",
	-70:  "sale_not_found",
	-80:  "invalid_sale_state",
	-90:  "invalid_sale_new_end_time",
	-101: "amount_mismatched",
	-102: "destination_balance_mismatched",
	-103: "not_allowed_account_destination",
	-104: "required_source_pay_for_destination",
	-105: "source_balance_mismatched",
	-106: "contract_not_found",
	-107: "invoice_receiver_balance_lock_amount_overflow",
	-108: "invoice_already_approved",
	-110: "payment_v2_malformed",
	-111: "underfunded",
	-112: "line_full",
	-113: "destination_balance_not_found",
	-114: "balance_assets_mismatched",
	-115: "src_balance_not_found",
	-116: "reference_duplication",
	-117: "stats_overflow",
	-118: "limits_exceeded",
	-119: "not_allowed_by_asset_policy",
	-120: "invalid_destination_fee",
	-121: "invalid_destination_fee_asset",
	-122: "fee_asset_mismatched",
	-123: "insufficient_fee_amount",
	-124: "balance_to_charge_fee_from_not_found",
	-125: "payment_amount_is_less_than_dest_fee",
	-126: "destination_account_not_found",
	130:  "cannot_create_for_acc_id_and_acc_type",
	131:  "invalid_limits",
	-140: "contract_details_too_long",
}

var reviewRequestResultCodeRevMap = map[string]int32{
	"ReviewRequestResultCodeSuccess":                                  0,
	"ReviewRequestResultCodeInvalidReason":                            -1,
	"ReviewRequestResultCodeInvalidAction":                            -2,
	"ReviewRequestResultCodeHashMismatched":                           -3,
	"ReviewRequestResultCodeNotFound":                                 -4,
	"ReviewRequestResultCodeTypeMismatched":                           -5,
	"ReviewRequestResultCodeRejectNotAllowed":                         -6,
	"ReviewRequestResultCodeInvalidExternalDetails":                   -7,
	"ReviewRequestResultCodeRequestorIsBlocked":                       -8,
	"ReviewRequestResultCodePermanentRejectNotAllowed":                -9,
	"ReviewRequestResultCodeAssetAlreadyExists":                       -20,
	"ReviewRequestResultCodeAssetDoesNotExists":                       -21,
	"ReviewRequestResultCodeMaxIssuanceAmountExceeded":                -40,
	"ReviewRequestResultCodeInsufficientAvailableForIssuanceAmount":   -41,
	"ReviewRequestResultCodeFullLine":                                 -42,
	"ReviewRequestResultCodeSystemTasksNotAllowed":                    -43,
	"ReviewRequestResultCodeBaseAssetDoesNotExists":                   -50,
	"ReviewRequestResultCodeHardCapWillExceedMaxIssuance":             -51,
	"ReviewRequestResultCodeInsufficientPreissuedForHardCap":          -52,
	"ReviewRequestResultCodeNonZeroTasksToRemoveNotAllowed":           -60,
	"ReviewRequestResultCodeSaleNotFound":                             -70,
	"ReviewRequestResultCodeInvalidSaleState":                         -80,
	"ReviewRequestResultCodeInvalidSaleNewEndTime":                    -90,
	"ReviewRequestResultCodeAmountMismatched":                         -101,
	"ReviewRequestResultCodeDestinationBalanceMismatched":             -102,
	"ReviewRequestResultCodeNotAllowedAccountDestination":             -103,
	"ReviewRequestResultCodeRequiredSourcePayForDestination":          -104,
	"ReviewRequestResultCodeSourceBalanceMismatched":                  -105,
	"ReviewRequestResultCodeContractNotFound":                         -106,
	"ReviewRequestResultCodeInvoiceReceiverBalanceLockAmountOverflow": -107,
	"ReviewRequestResultCodeInvoiceAlreadyApproved":                   -108,
	"ReviewRequestResultCodePaymentV2Malformed":                       -110,
	"ReviewRequestResultCodeUnderfunded":                              -111,
	"ReviewRequestResultCodeLineFull":                                 -112,
	"ReviewRequestResultCodeDestinationBalanceNotFound":               -113,
	"ReviewRequestResultCodeBalanceAssetsMismatched":                  -114,
	"ReviewRequestResultCodeSrcBalanceNotFound":                       -115,
	"ReviewRequestResultCodeReferenceDuplication":                     -116,
	"ReviewRequestResultCodeStatsOverflow":                            -117,
	"ReviewRequestResultCodeLimitsExceeded":                           -118,
	"ReviewRequestResultCodeNotAllowedByAssetPolicy":                  -119,
	"ReviewRequestResultCodeInvalidDestinationFee":                    -120,
	"ReviewRequestResultCodeInvalidDestinationFeeAsset":               -121,
	"ReviewRequestResultCodeFeeAssetMismatched":                       -122,
	"ReviewRequestResultCodeInsufficientFeeAmount":                    -123,
	"ReviewRequestResultCodeBalanceToChargeFeeFromNotFound":           -124,
	"ReviewRequestResultCodePaymentAmountIsLessThanDestFee":           -125,
	"ReviewRequestResultCodeDestinationAccountNotFound":               -126,
	"ReviewRequestResultCodeCannotCreateForAccIdAndAccType":           130,
	"ReviewRequestResultCodeInvalidLimits":                            131,
	"ReviewRequestResultCodeContractDetailsTooLong":                   -140,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ReviewRequestResultCode
func (e ReviewRequestResultCode) ValidEnum(v int32) bool {
	_, ok := reviewRequestResultCodeMap[v]
	return ok
}
func (e ReviewRequestResultCode) isFlag() bool {
	for i := len(ReviewRequestResultCodeAll) - 1; i >= 0; i-- {
		expected := ReviewRequestResultCode(2) << uint64(len(ReviewRequestResultCodeAll)-1) >> uint64(len(ReviewRequestResultCodeAll)-i)
		if expected != ReviewRequestResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ReviewRequestResultCode) String() string {
	name, _ := reviewRequestResultCodeMap[int32(e)]
	return name
}

func (e ReviewRequestResultCode) ShortString() string {
	name, _ := reviewRequestResultCodeShortMap[int32(e)]
	return name
}

func (e ReviewRequestResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range ReviewRequestResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ReviewRequestResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ReviewRequestResultCode(t.Value)
	return nil
}

// ReviewRequestResultSuccessExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//    		{
//    		case ADD_SALE_ID_REVIEW_REQUEST_RESULT:
//    		    uint64 saleID;
//    		case ADD_REVIEW_INVOICE_REQUEST_PAYMENT_RESPONSE:
//    		    PaymentV2Response paymentV2Response;
//    		case ADD_CONTRACT_ID_REVIEW_REQUEST_RESULT:
//    		    uint64 contractID;
//    		case EMPTY_VERSION:
//    			void;
//            case ADD_TASKS_TO_REVIEWABLE_REQUEST:
//                ExtendedResult extendedResult;
//    		}
//
type ReviewRequestResultSuccessExt struct {
	V                 LedgerVersion      `json:"v,omitempty"`
	SaleId            *Uint64            `json:"saleID,omitempty"`
	PaymentV2Response *PaymentV2Response `json:"paymentV2Response,omitempty"`
	ContractId        *Uint64            `json:"contractID,omitempty"`
	ExtendedResult    *ExtendedResult    `json:"extendedResult,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ReviewRequestResultSuccessExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ReviewRequestResultSuccessExt
func (u ReviewRequestResultSuccessExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionAddSaleIdReviewRequestResult:
		return "SaleId", true
	case LedgerVersionAddReviewInvoiceRequestPaymentResponse:
		return "PaymentV2Response", true
	case LedgerVersionAddContractIdReviewRequestResult:
		return "ContractId", true
	case LedgerVersionEmptyVersion:
		return "", true
	case LedgerVersionAddTasksToReviewableRequest:
		return "ExtendedResult", true
	}
	return "-", false
}

// NewReviewRequestResultSuccessExt creates a new  ReviewRequestResultSuccessExt.
func NewReviewRequestResultSuccessExt(v LedgerVersion, value interface{}) (result ReviewRequestResultSuccessExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionAddSaleIdReviewRequestResult:
		tv, ok := value.(Uint64)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint64")
			return
		}
		result.SaleId = &tv
	case LedgerVersionAddReviewInvoiceRequestPaymentResponse:
		tv, ok := value.(PaymentV2Response)
		if !ok {
			err = fmt.Errorf("invalid value, must be PaymentV2Response")
			return
		}
		result.PaymentV2Response = &tv
	case LedgerVersionAddContractIdReviewRequestResult:
		tv, ok := value.(Uint64)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint64")
			return
		}
		result.ContractId = &tv
	case LedgerVersionEmptyVersion:
		// void
	case LedgerVersionAddTasksToReviewableRequest:
		tv, ok := value.(ExtendedResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be ExtendedResult")
			return
		}
		result.ExtendedResult = &tv
	}
	return
}

// MustSaleId retrieves the SaleId value from the union,
// panicing if the value is not set.
func (u ReviewRequestResultSuccessExt) MustSaleId() Uint64 {
	val, ok := u.GetSaleId()

	if !ok {
		panic("arm SaleId is not set")
	}

	return val
}

// GetSaleId retrieves the SaleId value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ReviewRequestResultSuccessExt) GetSaleId() (result Uint64, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.V))

	if armName == "SaleId" {
		result = *u.SaleId
		ok = true
	}

	return
}

// MustPaymentV2Response retrieves the PaymentV2Response value from the union,
// panicing if the value is not set.
func (u ReviewRequestResultSuccessExt) MustPaymentV2Response() PaymentV2Response {
	val, ok := u.GetPaymentV2Response()

	if !ok {
		panic("arm PaymentV2Response is not set")
	}

	return val
}

// GetPaymentV2Response retrieves the PaymentV2Response value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ReviewRequestResultSuccessExt) GetPaymentV2Response() (result PaymentV2Response, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.V))

	if armName == "PaymentV2Response" {
		result = *u.PaymentV2Response
		ok = true
	}

	return
}

// MustContractId retrieves the ContractId value from the union,
// panicing if the value is not set.
func (u ReviewRequestResultSuccessExt) MustContractId() Uint64 {
	val, ok := u.GetContractId()

	if !ok {
		panic("arm ContractId is not set")
	}

	return val
}

// GetContractId retrieves the ContractId value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ReviewRequestResultSuccessExt) GetContractId() (result Uint64, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.V))

	if armName == "ContractId" {
		result = *u.ContractId
		ok = true
	}

	return
}

// MustExtendedResult retrieves the ExtendedResult value from the union,
// panicing if the value is not set.
func (u ReviewRequestResultSuccessExt) MustExtendedResult() ExtendedResult {
	val, ok := u.GetExtendedResult()

	if !ok {
		panic("arm ExtendedResult is not set")
	}

	return val
}

// GetExtendedResult retrieves the ExtendedResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ReviewRequestResultSuccessExt) GetExtendedResult() (result ExtendedResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.V))

	if armName == "ExtendedResult" {
		result = *u.ExtendedResult
		ok = true
	}

	return
}

// ReviewRequestResultSuccess is an XDR NestedStruct defines as:
//
//   struct {
//    		// reserved for future use
//    		union switch (LedgerVersion v)
//    		{
//    		case ADD_SALE_ID_REVIEW_REQUEST_RESULT:
//    		    uint64 saleID;
//    		case ADD_REVIEW_INVOICE_REQUEST_PAYMENT_RESPONSE:
//    		    PaymentV2Response paymentV2Response;
//    		case ADD_CONTRACT_ID_REVIEW_REQUEST_RESULT:
//    		    uint64 contractID;
//    		case EMPTY_VERSION:
//    			void;
//            case ADD_TASKS_TO_REVIEWABLE_REQUEST:
//                ExtendedResult extendedResult;
//    		}
//    		ext;
//    	}
//
type ReviewRequestResultSuccess struct {
	Ext ReviewRequestResultSuccessExt `json:"ext,omitempty"`
}

// ReviewRequestResult is an XDR Union defines as:
//
//   union ReviewRequestResult switch (ReviewRequestResultCode code)
//    {
//    case SUCCESS:
//    	struct {
//    		// reserved for future use
//    		union switch (LedgerVersion v)
//    		{
//    		case ADD_SALE_ID_REVIEW_REQUEST_RESULT:
//    		    uint64 saleID;
//    		case ADD_REVIEW_INVOICE_REQUEST_PAYMENT_RESPONSE:
//    		    PaymentV2Response paymentV2Response;
//    		case ADD_CONTRACT_ID_REVIEW_REQUEST_RESULT:
//    		    uint64 contractID;
//    		case EMPTY_VERSION:
//    			void;
//            case ADD_TASKS_TO_REVIEWABLE_REQUEST:
//                ExtendedResult extendedResult;
//    		}
//    		ext;
//    	} success;
//    default:
//        void;
//    };
//
type ReviewRequestResult struct {
	Code    ReviewRequestResultCode     `json:"code,omitempty"`
	Success *ReviewRequestResultSuccess `json:"success,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ReviewRequestResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ReviewRequestResult
func (u ReviewRequestResult) ArmForSwitch(sw int32) (string, bool) {
	switch ReviewRequestResultCode(sw) {
	case ReviewRequestResultCodeSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewReviewRequestResult creates a new  ReviewRequestResult.
func NewReviewRequestResult(code ReviewRequestResultCode, value interface{}) (result ReviewRequestResult, err error) {
	result.Code = code
	switch ReviewRequestResultCode(code) {
	case ReviewRequestResultCodeSuccess:
		tv, ok := value.(ReviewRequestResultSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be ReviewRequestResultSuccess")
			return
		}
		result.Success = &tv
	default:
		// void
	}
	return
}

// MustSuccess retrieves the Success value from the union,
// panicing if the value is not set.
func (u ReviewRequestResult) MustSuccess() ReviewRequestResultSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ReviewRequestResult) GetSuccess() (result ReviewRequestResultSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// SetFeesOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//
type SetFeesOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u SetFeesOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of SetFeesOpExt
func (u SetFeesOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewSetFeesOpExt creates a new  SetFeesOpExt.
func NewSetFeesOpExt(v LedgerVersion, value interface{}) (result SetFeesOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// SetFeesOp is an XDR Struct defines as:
//
//   struct SetFeesOp
//        {
//            FeeEntry* fee;
//    		bool isDelete;
//    		// reserved for future use
//    		union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//    		ext;
//        };
//
type SetFeesOp struct {
	Fee      *FeeEntry    `json:"fee,omitempty"`
	IsDelete bool         `json:"isDelete,omitempty"`
	Ext      SetFeesOpExt `json:"ext,omitempty"`
}

// SetFeesResultCode is an XDR Enum defines as:
//
//   enum SetFeesResultCode
//        {
//            // codes considered as "success" for the operation
//            SUCCESS = 0,
//
//            // codes considered as "failure" for the operation
//            INVALID_AMOUNT = -1,      // amount is negative
//    		INVALID_FEE_TYPE = -2,     // operation type is invalid
//            ASSET_NOT_FOUND = -3,
//            INVALID_ASSET = -4,
//            MALFORMED = -5,
//    		MALFORMED_RANGE = -6,
//    		RANGE_OVERLAP = -7,
//    		NOT_FOUND = -8,
//    		SUB_TYPE_NOT_EXIST = -9,
//    		INVALID_FEE_VERSION = -10, // version of fee entry is greater than ledger version
//    		INVALID_FEE_ASSET = -11,
//    		FEE_ASSET_NOT_ALLOWED = -12, // feeAsset can be set only if feeType is PAYMENT
//    		CROSS_ASSET_FEE_NOT_ALLOWED = -13, // feeAsset on payment fee type can differ from asset only if payment fee subtype is OUTGOING
//    		FEE_ASSET_NOT_FOUND = -14,
//    		ASSET_PAIR_NOT_FOUND = -15, // cannot create cross asset fee entry without existing asset pair
//    		INVALID_ASSET_PAIR_PRICE = -16
//        };
//
type SetFeesResultCode int32

const (
	SetFeesResultCodeSuccess                 SetFeesResultCode = 0
	SetFeesResultCodeInvalidAmount           SetFeesResultCode = -1
	SetFeesResultCodeInvalidFeeType          SetFeesResultCode = -2
	SetFeesResultCodeAssetNotFound           SetFeesResultCode = -3
	SetFeesResultCodeInvalidAsset            SetFeesResultCode = -4
	SetFeesResultCodeMalformed               SetFeesResultCode = -5
	SetFeesResultCodeMalformedRange          SetFeesResultCode = -6
	SetFeesResultCodeRangeOverlap            SetFeesResultCode = -7
	SetFeesResultCodeNotFound                SetFeesResultCode = -8
	SetFeesResultCodeSubTypeNotExist         SetFeesResultCode = -9
	SetFeesResultCodeInvalidFeeVersion       SetFeesResultCode = -10
	SetFeesResultCodeInvalidFeeAsset         SetFeesResultCode = -11
	SetFeesResultCodeFeeAssetNotAllowed      SetFeesResultCode = -12
	SetFeesResultCodeCrossAssetFeeNotAllowed SetFeesResultCode = -13
	SetFeesResultCodeFeeAssetNotFound        SetFeesResultCode = -14
	SetFeesResultCodeAssetPairNotFound       SetFeesResultCode = -15
	SetFeesResultCodeInvalidAssetPairPrice   SetFeesResultCode = -16
)

var SetFeesResultCodeAll = []SetFeesResultCode{
	SetFeesResultCodeSuccess,
	SetFeesResultCodeInvalidAmount,
	SetFeesResultCodeInvalidFeeType,
	SetFeesResultCodeAssetNotFound,
	SetFeesResultCodeInvalidAsset,
	SetFeesResultCodeMalformed,
	SetFeesResultCodeMalformedRange,
	SetFeesResultCodeRangeOverlap,
	SetFeesResultCodeNotFound,
	SetFeesResultCodeSubTypeNotExist,
	SetFeesResultCodeInvalidFeeVersion,
	SetFeesResultCodeInvalidFeeAsset,
	SetFeesResultCodeFeeAssetNotAllowed,
	SetFeesResultCodeCrossAssetFeeNotAllowed,
	SetFeesResultCodeFeeAssetNotFound,
	SetFeesResultCodeAssetPairNotFound,
	SetFeesResultCodeInvalidAssetPairPrice,
}

var setFeesResultCodeMap = map[int32]string{
	0:   "SetFeesResultCodeSuccess",
	-1:  "SetFeesResultCodeInvalidAmount",
	-2:  "SetFeesResultCodeInvalidFeeType",
	-3:  "SetFeesResultCodeAssetNotFound",
	-4:  "SetFeesResultCodeInvalidAsset",
	-5:  "SetFeesResultCodeMalformed",
	-6:  "SetFeesResultCodeMalformedRange",
	-7:  "SetFeesResultCodeRangeOverlap",
	-8:  "SetFeesResultCodeNotFound",
	-9:  "SetFeesResultCodeSubTypeNotExist",
	-10: "SetFeesResultCodeInvalidFeeVersion",
	-11: "SetFeesResultCodeInvalidFeeAsset",
	-12: "SetFeesResultCodeFeeAssetNotAllowed",
	-13: "SetFeesResultCodeCrossAssetFeeNotAllowed",
	-14: "SetFeesResultCodeFeeAssetNotFound",
	-15: "SetFeesResultCodeAssetPairNotFound",
	-16: "SetFeesResultCodeInvalidAssetPairPrice",
}

var setFeesResultCodeShortMap = map[int32]string{
	0:   "success",
	-1:  "invalid_amount",
	-2:  "invalid_fee_type",
	-3:  "asset_not_found",
	-4:  "invalid_asset",
	-5:  "malformed",
	-6:  "malformed_range",
	-7:  "range_overlap",
	-8:  "not_found",
	-9:  "sub_type_not_exist",
	-10: "invalid_fee_version",
	-11: "invalid_fee_asset",
	-12: "fee_asset_not_allowed",
	-13: "cross_asset_fee_not_allowed",
	-14: "fee_asset_not_found",
	-15: "asset_pair_not_found",
	-16: "invalid_asset_pair_price",
}

var setFeesResultCodeRevMap = map[string]int32{
	"SetFeesResultCodeSuccess":                 0,
	"SetFeesResultCodeInvalidAmount":           -1,
	"SetFeesResultCodeInvalidFeeType":          -2,
	"SetFeesResultCodeAssetNotFound":           -3,
	"SetFeesResultCodeInvalidAsset":            -4,
	"SetFeesResultCodeMalformed":               -5,
	"SetFeesResultCodeMalformedRange":          -6,
	"SetFeesResultCodeRangeOverlap":            -7,
	"SetFeesResultCodeNotFound":                -8,
	"SetFeesResultCodeSubTypeNotExist":         -9,
	"SetFeesResultCodeInvalidFeeVersion":       -10,
	"SetFeesResultCodeInvalidFeeAsset":         -11,
	"SetFeesResultCodeFeeAssetNotAllowed":      -12,
	"SetFeesResultCodeCrossAssetFeeNotAllowed": -13,
	"SetFeesResultCodeFeeAssetNotFound":        -14,
	"SetFeesResultCodeAssetPairNotFound":       -15,
	"SetFeesResultCodeInvalidAssetPairPrice":   -16,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for SetFeesResultCode
func (e SetFeesResultCode) ValidEnum(v int32) bool {
	_, ok := setFeesResultCodeMap[v]
	return ok
}
func (e SetFeesResultCode) isFlag() bool {
	for i := len(SetFeesResultCodeAll) - 1; i >= 0; i-- {
		expected := SetFeesResultCode(2) << uint64(len(SetFeesResultCodeAll)-1) >> uint64(len(SetFeesResultCodeAll)-i)
		if expected != SetFeesResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e SetFeesResultCode) String() string {
	name, _ := setFeesResultCodeMap[int32(e)]
	return name
}

func (e SetFeesResultCode) ShortString() string {
	name, _ := setFeesResultCodeShortMap[int32(e)]
	return name
}

func (e SetFeesResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range SetFeesResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *SetFeesResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = SetFeesResultCode(t.Value)
	return nil
}

// SetFeesResultSuccessExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//    				{
//    				case EMPTY_VERSION:
//    					void;
//    				}
//
type SetFeesResultSuccessExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u SetFeesResultSuccessExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of SetFeesResultSuccessExt
func (u SetFeesResultSuccessExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewSetFeesResultSuccessExt creates a new  SetFeesResultSuccessExt.
func NewSetFeesResultSuccessExt(v LedgerVersion, value interface{}) (result SetFeesResultSuccessExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// SetFeesResultSuccess is an XDR NestedStruct defines as:
//
//   struct {
//    				// reserved for future use
//    				union switch (LedgerVersion v)
//    				{
//    				case EMPTY_VERSION:
//    					void;
//    				}
//    				ext;
//    			}
//
type SetFeesResultSuccess struct {
	Ext SetFeesResultSuccessExt `json:"ext,omitempty"`
}

// SetFeesResult is an XDR Union defines as:
//
//   union SetFeesResult switch (SetFeesResultCode code)
//        {
//            case SUCCESS:
//                struct {
//    				// reserved for future use
//    				union switch (LedgerVersion v)
//    				{
//    				case EMPTY_VERSION:
//    					void;
//    				}
//    				ext;
//    			} success;
//            default:
//                void;
//        };
//
type SetFeesResult struct {
	Code    SetFeesResultCode     `json:"code,omitempty"`
	Success *SetFeesResultSuccess `json:"success,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u SetFeesResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of SetFeesResult
func (u SetFeesResult) ArmForSwitch(sw int32) (string, bool) {
	switch SetFeesResultCode(sw) {
	case SetFeesResultCodeSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewSetFeesResult creates a new  SetFeesResult.
func NewSetFeesResult(code SetFeesResultCode, value interface{}) (result SetFeesResult, err error) {
	result.Code = code
	switch SetFeesResultCode(code) {
	case SetFeesResultCodeSuccess:
		tv, ok := value.(SetFeesResultSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be SetFeesResultSuccess")
			return
		}
		result.Success = &tv
	default:
		// void
	}
	return
}

// MustSuccess retrieves the Success value from the union,
// panicing if the value is not set.
func (u SetFeesResult) MustSuccess() SetFeesResultSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u SetFeesResult) GetSuccess() (result SetFeesResultSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// ManageTrustAction is an XDR Enum defines as:
//
//   enum ManageTrustAction
//    {
//        TRUST_ADD = 0,
//        TRUST_REMOVE = 1
//    };
//
type ManageTrustAction int32

const (
	ManageTrustActionTrustAdd    ManageTrustAction = 0
	ManageTrustActionTrustRemove ManageTrustAction = 1
)

var ManageTrustActionAll = []ManageTrustAction{
	ManageTrustActionTrustAdd,
	ManageTrustActionTrustRemove,
}

var manageTrustActionMap = map[int32]string{
	0: "ManageTrustActionTrustAdd",
	1: "ManageTrustActionTrustRemove",
}

var manageTrustActionShortMap = map[int32]string{
	0: "trust_add",
	1: "trust_remove",
}

var manageTrustActionRevMap = map[string]int32{
	"ManageTrustActionTrustAdd":    0,
	"ManageTrustActionTrustRemove": 1,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ManageTrustAction
func (e ManageTrustAction) ValidEnum(v int32) bool {
	_, ok := manageTrustActionMap[v]
	return ok
}
func (e ManageTrustAction) isFlag() bool {
	for i := len(ManageTrustActionAll) - 1; i >= 0; i-- {
		expected := ManageTrustAction(2) << uint64(len(ManageTrustActionAll)-1) >> uint64(len(ManageTrustActionAll)-i)
		if expected != ManageTrustActionAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ManageTrustAction) String() string {
	name, _ := manageTrustActionMap[int32(e)]
	return name
}

func (e ManageTrustAction) ShortString() string {
	name, _ := manageTrustActionShortMap[int32(e)]
	return name
}

func (e ManageTrustAction) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range ManageTrustActionAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ManageTrustAction) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ManageTrustAction(t.Value)
	return nil
}

// TrustDataExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//    	{
//    	case EMPTY_VERSION:
//    		void;
//    	}
//
type TrustDataExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u TrustDataExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of TrustDataExt
func (u TrustDataExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewTrustDataExt creates a new  TrustDataExt.
func NewTrustDataExt(v LedgerVersion, value interface{}) (result TrustDataExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// TrustData is an XDR Struct defines as:
//
//   struct TrustData {
//        TrustEntry trust;
//        ManageTrustAction action;
//    	// reserved for future use
//    	union switch (LedgerVersion v)
//    	{
//    	case EMPTY_VERSION:
//    		void;
//    	}
//    	ext;
//    };
//
type TrustData struct {
	Trust  TrustEntry        `json:"trust,omitempty"`
	Action ManageTrustAction `json:"action,omitempty"`
	Ext    TrustDataExt      `json:"ext,omitempty"`
}

// LimitsUpdateRequestDataExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type LimitsUpdateRequestDataExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LimitsUpdateRequestDataExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LimitsUpdateRequestDataExt
func (u LimitsUpdateRequestDataExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewLimitsUpdateRequestDataExt creates a new  LimitsUpdateRequestDataExt.
func NewLimitsUpdateRequestDataExt(v LedgerVersion, value interface{}) (result LimitsUpdateRequestDataExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// LimitsUpdateRequestData is an XDR Struct defines as:
//
//   struct LimitsUpdateRequestData {
//        Hash documentHash;
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type LimitsUpdateRequestData struct {
	DocumentHash Hash                       `json:"documentHash,omitempty"`
	Ext          LimitsUpdateRequestDataExt `json:"ext,omitempty"`
}

// SetOptionsOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//    	{
//    	case EMPTY_VERSION:
//    		void;
//    	}
//
type SetOptionsOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u SetOptionsOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of SetOptionsOpExt
func (u SetOptionsOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewSetOptionsOpExt creates a new  SetOptionsOpExt.
func NewSetOptionsOpExt(v LedgerVersion, value interface{}) (result SetOptionsOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// SetOptionsOp is an XDR Struct defines as:
//
//   struct SetOptionsOp
//    {
//        // account threshold manipulation
//        uint32* masterWeight; // weight of the master account
//        uint32* lowThreshold;
//        uint32* medThreshold;
//        uint32* highThreshold;
//
//        // Add, update or remove a signer for the account
//        // signer is deleted if the weight is 0
//        Signer* signer;
//
//        TrustData* trustData;
//
//        // Create LimitsUpdateRequest for account
//        LimitsUpdateRequestData* limitsUpdateRequestData;
//
//    	// reserved for future use
//    	union switch (LedgerVersion v)
//    	{
//    	case EMPTY_VERSION:
//    		void;
//    	}
//    	ext;
//
//    };
//
type SetOptionsOp struct {
	MasterWeight            *Uint32                  `json:"masterWeight,omitempty"`
	LowThreshold            *Uint32                  `json:"lowThreshold,omitempty"`
	MedThreshold            *Uint32                  `json:"medThreshold,omitempty"`
	HighThreshold           *Uint32                  `json:"highThreshold,omitempty"`
	Signer                  *Signer                  `json:"signer,omitempty"`
	TrustData               *TrustData               `json:"trustData,omitempty"`
	LimitsUpdateRequestData *LimitsUpdateRequestData `json:"limitsUpdateRequestData,omitempty"`
	Ext                     SetOptionsOpExt          `json:"ext,omitempty"`
}

// SetOptionsResultCode is an XDR Enum defines as:
//
//   enum SetOptionsResultCode
//    {
//        // codes considered as "success" for the operation
//        SUCCESS = 0,
//        // codes considered as "failure" for the operation
//        TOO_MANY_SIGNERS = -1, // max number of signers already reached
//        THRESHOLD_OUT_OF_RANGE = -2, // bad value for weight/threshold
//        BAD_SIGNER = -3,             // signer cannot be masterkey
//        BALANCE_NOT_FOUND = -4,
//        TRUST_MALFORMED = -5,
//    	TRUST_TOO_MANY = -6,
//    	INVALID_SIGNER_VERSION = -7, // if signer version is higher than ledger version
//    	LIMITS_UPDATE_REQUEST_REFERENCE_DUPLICATION = -8
//    };
//
type SetOptionsResultCode int32

const (
	SetOptionsResultCodeSuccess                                 SetOptionsResultCode = 0
	SetOptionsResultCodeTooManySigners                          SetOptionsResultCode = -1
	SetOptionsResultCodeThresholdOutOfRange                     SetOptionsResultCode = -2
	SetOptionsResultCodeBadSigner                               SetOptionsResultCode = -3
	SetOptionsResultCodeBalanceNotFound                         SetOptionsResultCode = -4
	SetOptionsResultCodeTrustMalformed                          SetOptionsResultCode = -5
	SetOptionsResultCodeTrustTooMany                            SetOptionsResultCode = -6
	SetOptionsResultCodeInvalidSignerVersion                    SetOptionsResultCode = -7
	SetOptionsResultCodeLimitsUpdateRequestReferenceDuplication SetOptionsResultCode = -8
)

var SetOptionsResultCodeAll = []SetOptionsResultCode{
	SetOptionsResultCodeSuccess,
	SetOptionsResultCodeTooManySigners,
	SetOptionsResultCodeThresholdOutOfRange,
	SetOptionsResultCodeBadSigner,
	SetOptionsResultCodeBalanceNotFound,
	SetOptionsResultCodeTrustMalformed,
	SetOptionsResultCodeTrustTooMany,
	SetOptionsResultCodeInvalidSignerVersion,
	SetOptionsResultCodeLimitsUpdateRequestReferenceDuplication,
}

var setOptionsResultCodeMap = map[int32]string{
	0:  "SetOptionsResultCodeSuccess",
	-1: "SetOptionsResultCodeTooManySigners",
	-2: "SetOptionsResultCodeThresholdOutOfRange",
	-3: "SetOptionsResultCodeBadSigner",
	-4: "SetOptionsResultCodeBalanceNotFound",
	-5: "SetOptionsResultCodeTrustMalformed",
	-6: "SetOptionsResultCodeTrustTooMany",
	-7: "SetOptionsResultCodeInvalidSignerVersion",
	-8: "SetOptionsResultCodeLimitsUpdateRequestReferenceDuplication",
}

var setOptionsResultCodeShortMap = map[int32]string{
	0:  "success",
	-1: "too_many_signers",
	-2: "threshold_out_of_range",
	-3: "bad_signer",
	-4: "balance_not_found",
	-5: "trust_malformed",
	-6: "trust_too_many",
	-7: "invalid_signer_version",
	-8: "limits_update_request_reference_duplication",
}

var setOptionsResultCodeRevMap = map[string]int32{
	"SetOptionsResultCodeSuccess":                                 0,
	"SetOptionsResultCodeTooManySigners":                          -1,
	"SetOptionsResultCodeThresholdOutOfRange":                     -2,
	"SetOptionsResultCodeBadSigner":                               -3,
	"SetOptionsResultCodeBalanceNotFound":                         -4,
	"SetOptionsResultCodeTrustMalformed":                          -5,
	"SetOptionsResultCodeTrustTooMany":                            -6,
	"SetOptionsResultCodeInvalidSignerVersion":                    -7,
	"SetOptionsResultCodeLimitsUpdateRequestReferenceDuplication": -8,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for SetOptionsResultCode
func (e SetOptionsResultCode) ValidEnum(v int32) bool {
	_, ok := setOptionsResultCodeMap[v]
	return ok
}
func (e SetOptionsResultCode) isFlag() bool {
	for i := len(SetOptionsResultCodeAll) - 1; i >= 0; i-- {
		expected := SetOptionsResultCode(2) << uint64(len(SetOptionsResultCodeAll)-1) >> uint64(len(SetOptionsResultCodeAll)-i)
		if expected != SetOptionsResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e SetOptionsResultCode) String() string {
	name, _ := setOptionsResultCodeMap[int32(e)]
	return name
}

func (e SetOptionsResultCode) ShortString() string {
	name, _ := setOptionsResultCodeShortMap[int32(e)]
	return name
}

func (e SetOptionsResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range SetOptionsResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *SetOptionsResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = SetOptionsResultCode(t.Value)
	return nil
}

// SetOptionsResultSuccessExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//
type SetOptionsResultSuccessExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u SetOptionsResultSuccessExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of SetOptionsResultSuccessExt
func (u SetOptionsResultSuccessExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewSetOptionsResultSuccessExt creates a new  SetOptionsResultSuccessExt.
func NewSetOptionsResultSuccessExt(v LedgerVersion, value interface{}) (result SetOptionsResultSuccessExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// SetOptionsResultSuccess is an XDR NestedStruct defines as:
//
//   struct {
//            uint64 limitsUpdateRequestID;
//    		// reserved for future use
//    		union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//    		ext;
//    	}
//
type SetOptionsResultSuccess struct {
	LimitsUpdateRequestId Uint64                     `json:"limitsUpdateRequestID,omitempty"`
	Ext                   SetOptionsResultSuccessExt `json:"ext,omitempty"`
}

// SetOptionsResult is an XDR Union defines as:
//
//   union SetOptionsResult switch (SetOptionsResultCode code)
//    {
//    case SUCCESS:
//        struct {
//            uint64 limitsUpdateRequestID;
//    		// reserved for future use
//    		union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//    		ext;
//    	} success;
//    default:
//        void;
//    };
//
type SetOptionsResult struct {
	Code    SetOptionsResultCode     `json:"code,omitempty"`
	Success *SetOptionsResultSuccess `json:"success,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u SetOptionsResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of SetOptionsResult
func (u SetOptionsResult) ArmForSwitch(sw int32) (string, bool) {
	switch SetOptionsResultCode(sw) {
	case SetOptionsResultCodeSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewSetOptionsResult creates a new  SetOptionsResult.
func NewSetOptionsResult(code SetOptionsResultCode, value interface{}) (result SetOptionsResult, err error) {
	result.Code = code
	switch SetOptionsResultCode(code) {
	case SetOptionsResultCodeSuccess:
		tv, ok := value.(SetOptionsResultSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be SetOptionsResultSuccess")
			return
		}
		result.Success = &tv
	default:
		// void
	}
	return
}

// MustSuccess retrieves the Success value from the union,
// panicing if the value is not set.
func (u SetOptionsResult) MustSuccess() SetOptionsResultSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u SetOptionsResult) GetSuccess() (result SetOptionsResultSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// ErrorCode is an XDR Enum defines as:
//
//   enum ErrorCode
//    {
//        MISC = 0, // Unspecific error
//        DATA = 1, // Malformed data
//        CONF = 2, // Misconfiguration error
//        AUTH = 3, // Authentication failure
//        LOAD = 4  // System overloaded
//    };
//
type ErrorCode int32

const (
	ErrorCodeMisc ErrorCode = 0
	ErrorCodeData ErrorCode = 1
	ErrorCodeConf ErrorCode = 2
	ErrorCodeAuth ErrorCode = 3
	ErrorCodeLoad ErrorCode = 4
)

var ErrorCodeAll = []ErrorCode{
	ErrorCodeMisc,
	ErrorCodeData,
	ErrorCodeConf,
	ErrorCodeAuth,
	ErrorCodeLoad,
}

var errorCodeMap = map[int32]string{
	0: "ErrorCodeMisc",
	1: "ErrorCodeData",
	2: "ErrorCodeConf",
	3: "ErrorCodeAuth",
	4: "ErrorCodeLoad",
}

var errorCodeShortMap = map[int32]string{
	0: "misc",
	1: "data",
	2: "conf",
	3: "auth",
	4: "load",
}

var errorCodeRevMap = map[string]int32{
	"ErrorCodeMisc": 0,
	"ErrorCodeData": 1,
	"ErrorCodeConf": 2,
	"ErrorCodeAuth": 3,
	"ErrorCodeLoad": 4,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ErrorCode
func (e ErrorCode) ValidEnum(v int32) bool {
	_, ok := errorCodeMap[v]
	return ok
}
func (e ErrorCode) isFlag() bool {
	for i := len(ErrorCodeAll) - 1; i >= 0; i-- {
		expected := ErrorCode(2) << uint64(len(ErrorCodeAll)-1) >> uint64(len(ErrorCodeAll)-i)
		if expected != ErrorCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ErrorCode) String() string {
	name, _ := errorCodeMap[int32(e)]
	return name
}

func (e ErrorCode) ShortString() string {
	name, _ := errorCodeShortMap[int32(e)]
	return name
}

func (e ErrorCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range ErrorCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *ErrorCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ErrorCode(t.Value)
	return nil
}

// Error is an XDR Struct defines as:
//
//   struct Error
//    {
//        ErrorCode code;
//        string msg<100>;
//    };
//
type Error struct {
	Code ErrorCode `json:"code,omitempty"`
	Msg  string    `json:"msg,omitempty" xdrmaxsize:"100"`
}

// AuthCert is an XDR Struct defines as:
//
//   struct AuthCert
//    {
//        Curve25519Public pubkey;
//        uint64 expiration;
//        Signature sig;
//    };
//
type AuthCert struct {
	Pubkey     Curve25519Public `json:"pubkey,omitempty"`
	Expiration Uint64           `json:"expiration,omitempty"`
	Sig        Signature        `json:"sig,omitempty"`
}

// Hello is an XDR Struct defines as:
//
//   struct Hello
//    {
//        uint32 ledgerVersion;
//        uint32 overlayVersion;
//        uint32 overlayMinVersion;
//        Hash networkID;
//        string versionStr<100>;
//        int listeningPort;
//        NodeID peerID;
//        AuthCert cert;
//        uint256 nonce;
//    };
//
type Hello struct {
	LedgerVersion     Uint32   `json:"ledgerVersion,omitempty"`
	OverlayVersion    Uint32   `json:"overlayVersion,omitempty"`
	OverlayMinVersion Uint32   `json:"overlayMinVersion,omitempty"`
	NetworkId         Hash     `json:"networkID,omitempty"`
	VersionStr        string   `json:"versionStr,omitempty" xdrmaxsize:"100"`
	ListeningPort     int32    `json:"listeningPort,omitempty"`
	PeerId            NodeId   `json:"peerID,omitempty"`
	Cert              AuthCert `json:"cert,omitempty"`
	Nonce             Uint256  `json:"nonce,omitempty"`
}

// Auth is an XDR Struct defines as:
//
//   struct Auth
//    {
//        // Empty message, just to confirm
//        // establishment of MAC keys.
//        int unused;
//    };
//
type Auth struct {
	Unused int32 `json:"unused,omitempty"`
}

// IpAddrType is an XDR Enum defines as:
//
//   enum IPAddrType
//    {
//        IPv4 = 0,
//        IPv6 = 1
//    };
//
type IpAddrType int32

const (
	IpAddrTypeIPv4 IpAddrType = 0
	IpAddrTypeIPv6 IpAddrType = 1
)

var IpAddrTypeAll = []IpAddrType{
	IpAddrTypeIPv4,
	IpAddrTypeIPv6,
}

var ipAddrTypeMap = map[int32]string{
	0: "IpAddrTypeIPv4",
	1: "IpAddrTypeIPv6",
}

var ipAddrTypeShortMap = map[int32]string{
	0: "i_pv4",
	1: "i_pv6",
}

var ipAddrTypeRevMap = map[string]int32{
	"IpAddrTypeIPv4": 0,
	"IpAddrTypeIPv6": 1,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for IpAddrType
func (e IpAddrType) ValidEnum(v int32) bool {
	_, ok := ipAddrTypeMap[v]
	return ok
}
func (e IpAddrType) isFlag() bool {
	for i := len(IpAddrTypeAll) - 1; i >= 0; i-- {
		expected := IpAddrType(2) << uint64(len(IpAddrTypeAll)-1) >> uint64(len(IpAddrTypeAll)-i)
		if expected != IpAddrTypeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e IpAddrType) String() string {
	name, _ := ipAddrTypeMap[int32(e)]
	return name
}

func (e IpAddrType) ShortString() string {
	name, _ := ipAddrTypeShortMap[int32(e)]
	return name
}

func (e IpAddrType) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range IpAddrTypeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *IpAddrType) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = IpAddrType(t.Value)
	return nil
}

// PeerAddressIp is an XDR NestedUnion defines as:
//
//   union switch (IPAddrType type)
//        {
//        case IPv4:
//            opaque ipv4[4];
//        case IPv6:
//            opaque ipv6[16];
//        }
//
type PeerAddressIp struct {
	Type IpAddrType `json:"type,omitempty"`
	Ipv4 *[4]byte   `json:"ipv4,omitempty"`
	Ipv6 *[16]byte  `json:"ipv6,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u PeerAddressIp) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of PeerAddressIp
func (u PeerAddressIp) ArmForSwitch(sw int32) (string, bool) {
	switch IpAddrType(sw) {
	case IpAddrTypeIPv4:
		return "Ipv4", true
	case IpAddrTypeIPv6:
		return "Ipv6", true
	}
	return "-", false
}

// NewPeerAddressIp creates a new  PeerAddressIp.
func NewPeerAddressIp(aType IpAddrType, value interface{}) (result PeerAddressIp, err error) {
	result.Type = aType
	switch IpAddrType(aType) {
	case IpAddrTypeIPv4:
		tv, ok := value.([4]byte)
		if !ok {
			err = fmt.Errorf("invalid value, must be [4]byte")
			return
		}
		result.Ipv4 = &tv
	case IpAddrTypeIPv6:
		tv, ok := value.([16]byte)
		if !ok {
			err = fmt.Errorf("invalid value, must be [16]byte")
			return
		}
		result.Ipv6 = &tv
	}
	return
}

// MustIpv4 retrieves the Ipv4 value from the union,
// panicing if the value is not set.
func (u PeerAddressIp) MustIpv4() [4]byte {
	val, ok := u.GetIpv4()

	if !ok {
		panic("arm Ipv4 is not set")
	}

	return val
}

// GetIpv4 retrieves the Ipv4 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u PeerAddressIp) GetIpv4() (result [4]byte, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Ipv4" {
		result = *u.Ipv4
		ok = true
	}

	return
}

// MustIpv6 retrieves the Ipv6 value from the union,
// panicing if the value is not set.
func (u PeerAddressIp) MustIpv6() [16]byte {
	val, ok := u.GetIpv6()

	if !ok {
		panic("arm Ipv6 is not set")
	}

	return val
}

// GetIpv6 retrieves the Ipv6 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u PeerAddressIp) GetIpv6() (result [16]byte, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Ipv6" {
		result = *u.Ipv6
		ok = true
	}

	return
}

// PeerAddress is an XDR Struct defines as:
//
//   struct PeerAddress
//    {
//        union switch (IPAddrType type)
//        {
//        case IPv4:
//            opaque ipv4[4];
//        case IPv6:
//            opaque ipv6[16];
//        }
//        ip;
//        uint32 port;
//        uint32 numFailures;
//    };
//
type PeerAddress struct {
	Ip          PeerAddressIp `json:"ip,omitempty"`
	Port        Uint32        `json:"port,omitempty"`
	NumFailures Uint32        `json:"numFailures,omitempty"`
}

// MessageType is an XDR Enum defines as:
//
//   enum MessageType
//    {
//        ERROR_MSG = 0,
//        AUTH = 2,
//        DONT_HAVE = 3,
//
//        GET_PEERS = 4, // gets a list of peers this guy knows about
//        PEERS = 5,
//
//        GET_TX_SET = 6, // gets a particular txset by hash
//        TX_SET = 7,
//
//        TRANSACTION = 8, // pass on a tx you have heard about
//
//        // SCP
//        GET_SCP_QUORUMSET = 9,
//        SCP_QUORUMSET = 10,
//        SCP_MESSAGE = 11,
//        GET_SCP_STATE = 12,
//
//        // new messages
//        HELLO = 13
//    };
//
type MessageType int32

const (
	MessageTypeErrorMsg        MessageType = 0
	MessageTypeAuth            MessageType = 2
	MessageTypeDontHave        MessageType = 3
	MessageTypeGetPeers        MessageType = 4
	MessageTypePeers           MessageType = 5
	MessageTypeGetTxSet        MessageType = 6
	MessageTypeTxSet           MessageType = 7
	MessageTypeTransaction     MessageType = 8
	MessageTypeGetScpQuorumset MessageType = 9
	MessageTypeScpQuorumset    MessageType = 10
	MessageTypeScpMessage      MessageType = 11
	MessageTypeGetScpState     MessageType = 12
	MessageTypeHello           MessageType = 13
)

var MessageTypeAll = []MessageType{
	MessageTypeErrorMsg,
	MessageTypeAuth,
	MessageTypeDontHave,
	MessageTypeGetPeers,
	MessageTypePeers,
	MessageTypeGetTxSet,
	MessageTypeTxSet,
	MessageTypeTransaction,
	MessageTypeGetScpQuorumset,
	MessageTypeScpQuorumset,
	MessageTypeScpMessage,
	MessageTypeGetScpState,
	MessageTypeHello,
}

var messageTypeMap = map[int32]string{
	0:  "MessageTypeErrorMsg",
	2:  "MessageTypeAuth",
	3:  "MessageTypeDontHave",
	4:  "MessageTypeGetPeers",
	5:  "MessageTypePeers",
	6:  "MessageTypeGetTxSet",
	7:  "MessageTypeTxSet",
	8:  "MessageTypeTransaction",
	9:  "MessageTypeGetScpQuorumset",
	10: "MessageTypeScpQuorumset",
	11: "MessageTypeScpMessage",
	12: "MessageTypeGetScpState",
	13: "MessageTypeHello",
}

var messageTypeShortMap = map[int32]string{
	0:  "error_msg",
	2:  "auth",
	3:  "dont_have",
	4:  "get_peers",
	5:  "peers",
	6:  "get_tx_set",
	7:  "tx_set",
	8:  "transaction",
	9:  "get_scp_quorumset",
	10: "scp_quorumset",
	11: "scp_message",
	12: "get_scp_state",
	13: "hello",
}

var messageTypeRevMap = map[string]int32{
	"MessageTypeErrorMsg":        0,
	"MessageTypeAuth":            2,
	"MessageTypeDontHave":        3,
	"MessageTypeGetPeers":        4,
	"MessageTypePeers":           5,
	"MessageTypeGetTxSet":        6,
	"MessageTypeTxSet":           7,
	"MessageTypeTransaction":     8,
	"MessageTypeGetScpQuorumset": 9,
	"MessageTypeScpQuorumset":    10,
	"MessageTypeScpMessage":      11,
	"MessageTypeGetScpState":     12,
	"MessageTypeHello":           13,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for MessageType
func (e MessageType) ValidEnum(v int32) bool {
	_, ok := messageTypeMap[v]
	return ok
}
func (e MessageType) isFlag() bool {
	for i := len(MessageTypeAll) - 1; i >= 0; i-- {
		expected := MessageType(2) << uint64(len(MessageTypeAll)-1) >> uint64(len(MessageTypeAll)-i)
		if expected != MessageTypeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e MessageType) String() string {
	name, _ := messageTypeMap[int32(e)]
	return name
}

func (e MessageType) ShortString() string {
	name, _ := messageTypeShortMap[int32(e)]
	return name
}

func (e MessageType) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range MessageTypeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *MessageType) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = MessageType(t.Value)
	return nil
}

// DontHave is an XDR Struct defines as:
//
//   struct DontHave
//    {
//        MessageType type;
//        uint256 reqHash;
//    };
//
type DontHave struct {
	Type    MessageType `json:"type,omitempty"`
	ReqHash Uint256     `json:"reqHash,omitempty"`
}

// StellarMessage is an XDR Union defines as:
//
//   union StellarMessage switch (MessageType type)
//    {
//    case ERROR_MSG:
//        Error error;
//    case HELLO:
//        Hello hello;
//    case AUTH:
//        Auth auth;
//    case DONT_HAVE:
//        DontHave dontHave;
//    case GET_PEERS:
//        void;
//    case PEERS:
//        PeerAddress peers<>;
//
//    case GET_TX_SET:
//        uint256 txSetHash;
//    case TX_SET:
//        TransactionSet txSet;
//
//    case TRANSACTION:
//        TransactionEnvelope transaction;
//
//    // SCP
//    case GET_SCP_QUORUMSET:
//        uint256 qSetHash;
//    case SCP_QUORUMSET:
//        SCPQuorumSet qSet;
//    case SCP_MESSAGE:
//        SCPEnvelope envelope;
//    case GET_SCP_STATE:
//        uint32 getSCPLedgerSeq; // ledger seq requested ; if 0, requests the latest
//    };
//
type StellarMessage struct {
	Type            MessageType          `json:"type,omitempty"`
	Error           *Error               `json:"error,omitempty"`
	Hello           *Hello               `json:"hello,omitempty"`
	Auth            *Auth                `json:"auth,omitempty"`
	DontHave        *DontHave            `json:"dontHave,omitempty"`
	Peers           *[]PeerAddress       `json:"peers,omitempty"`
	TxSetHash       *Uint256             `json:"txSetHash,omitempty"`
	TxSet           *TransactionSet      `json:"txSet,omitempty"`
	Transaction     *TransactionEnvelope `json:"transaction,omitempty"`
	QSetHash        *Uint256             `json:"qSetHash,omitempty"`
	QSet            *ScpQuorumSet        `json:"qSet,omitempty"`
	Envelope        *ScpEnvelope         `json:"envelope,omitempty"`
	GetScpLedgerSeq *Uint32              `json:"getSCPLedgerSeq,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u StellarMessage) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of StellarMessage
func (u StellarMessage) ArmForSwitch(sw int32) (string, bool) {
	switch MessageType(sw) {
	case MessageTypeErrorMsg:
		return "Error", true
	case MessageTypeHello:
		return "Hello", true
	case MessageTypeAuth:
		return "Auth", true
	case MessageTypeDontHave:
		return "DontHave", true
	case MessageTypeGetPeers:
		return "", true
	case MessageTypePeers:
		return "Peers", true
	case MessageTypeGetTxSet:
		return "TxSetHash", true
	case MessageTypeTxSet:
		return "TxSet", true
	case MessageTypeTransaction:
		return "Transaction", true
	case MessageTypeGetScpQuorumset:
		return "QSetHash", true
	case MessageTypeScpQuorumset:
		return "QSet", true
	case MessageTypeScpMessage:
		return "Envelope", true
	case MessageTypeGetScpState:
		return "GetScpLedgerSeq", true
	}
	return "-", false
}

// NewStellarMessage creates a new  StellarMessage.
func NewStellarMessage(aType MessageType, value interface{}) (result StellarMessage, err error) {
	result.Type = aType
	switch MessageType(aType) {
	case MessageTypeErrorMsg:
		tv, ok := value.(Error)
		if !ok {
			err = fmt.Errorf("invalid value, must be Error")
			return
		}
		result.Error = &tv
	case MessageTypeHello:
		tv, ok := value.(Hello)
		if !ok {
			err = fmt.Errorf("invalid value, must be Hello")
			return
		}
		result.Hello = &tv
	case MessageTypeAuth:
		tv, ok := value.(Auth)
		if !ok {
			err = fmt.Errorf("invalid value, must be Auth")
			return
		}
		result.Auth = &tv
	case MessageTypeDontHave:
		tv, ok := value.(DontHave)
		if !ok {
			err = fmt.Errorf("invalid value, must be DontHave")
			return
		}
		result.DontHave = &tv
	case MessageTypeGetPeers:
		// void
	case MessageTypePeers:
		tv, ok := value.([]PeerAddress)
		if !ok {
			err = fmt.Errorf("invalid value, must be []PeerAddress")
			return
		}
		result.Peers = &tv
	case MessageTypeGetTxSet:
		tv, ok := value.(Uint256)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint256")
			return
		}
		result.TxSetHash = &tv
	case MessageTypeTxSet:
		tv, ok := value.(TransactionSet)
		if !ok {
			err = fmt.Errorf("invalid value, must be TransactionSet")
			return
		}
		result.TxSet = &tv
	case MessageTypeTransaction:
		tv, ok := value.(TransactionEnvelope)
		if !ok {
			err = fmt.Errorf("invalid value, must be TransactionEnvelope")
			return
		}
		result.Transaction = &tv
	case MessageTypeGetScpQuorumset:
		tv, ok := value.(Uint256)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint256")
			return
		}
		result.QSetHash = &tv
	case MessageTypeScpQuorumset:
		tv, ok := value.(ScpQuorumSet)
		if !ok {
			err = fmt.Errorf("invalid value, must be ScpQuorumSet")
			return
		}
		result.QSet = &tv
	case MessageTypeScpMessage:
		tv, ok := value.(ScpEnvelope)
		if !ok {
			err = fmt.Errorf("invalid value, must be ScpEnvelope")
			return
		}
		result.Envelope = &tv
	case MessageTypeGetScpState:
		tv, ok := value.(Uint32)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint32")
			return
		}
		result.GetScpLedgerSeq = &tv
	}
	return
}

// MustError retrieves the Error value from the union,
// panicing if the value is not set.
func (u StellarMessage) MustError() Error {
	val, ok := u.GetError()

	if !ok {
		panic("arm Error is not set")
	}

	return val
}

// GetError retrieves the Error value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u StellarMessage) GetError() (result Error, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Error" {
		result = *u.Error
		ok = true
	}

	return
}

// MustHello retrieves the Hello value from the union,
// panicing if the value is not set.
func (u StellarMessage) MustHello() Hello {
	val, ok := u.GetHello()

	if !ok {
		panic("arm Hello is not set")
	}

	return val
}

// GetHello retrieves the Hello value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u StellarMessage) GetHello() (result Hello, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Hello" {
		result = *u.Hello
		ok = true
	}

	return
}

// MustAuth retrieves the Auth value from the union,
// panicing if the value is not set.
func (u StellarMessage) MustAuth() Auth {
	val, ok := u.GetAuth()

	if !ok {
		panic("arm Auth is not set")
	}

	return val
}

// GetAuth retrieves the Auth value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u StellarMessage) GetAuth() (result Auth, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Auth" {
		result = *u.Auth
		ok = true
	}

	return
}

// MustDontHave retrieves the DontHave value from the union,
// panicing if the value is not set.
func (u StellarMessage) MustDontHave() DontHave {
	val, ok := u.GetDontHave()

	if !ok {
		panic("arm DontHave is not set")
	}

	return val
}

// GetDontHave retrieves the DontHave value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u StellarMessage) GetDontHave() (result DontHave, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "DontHave" {
		result = *u.DontHave
		ok = true
	}

	return
}

// MustPeers retrieves the Peers value from the union,
// panicing if the value is not set.
func (u StellarMessage) MustPeers() []PeerAddress {
	val, ok := u.GetPeers()

	if !ok {
		panic("arm Peers is not set")
	}

	return val
}

// GetPeers retrieves the Peers value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u StellarMessage) GetPeers() (result []PeerAddress, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Peers" {
		result = *u.Peers
		ok = true
	}

	return
}

// MustTxSetHash retrieves the TxSetHash value from the union,
// panicing if the value is not set.
func (u StellarMessage) MustTxSetHash() Uint256 {
	val, ok := u.GetTxSetHash()

	if !ok {
		panic("arm TxSetHash is not set")
	}

	return val
}

// GetTxSetHash retrieves the TxSetHash value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u StellarMessage) GetTxSetHash() (result Uint256, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "TxSetHash" {
		result = *u.TxSetHash
		ok = true
	}

	return
}

// MustTxSet retrieves the TxSet value from the union,
// panicing if the value is not set.
func (u StellarMessage) MustTxSet() TransactionSet {
	val, ok := u.GetTxSet()

	if !ok {
		panic("arm TxSet is not set")
	}

	return val
}

// GetTxSet retrieves the TxSet value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u StellarMessage) GetTxSet() (result TransactionSet, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "TxSet" {
		result = *u.TxSet
		ok = true
	}

	return
}

// MustTransaction retrieves the Transaction value from the union,
// panicing if the value is not set.
func (u StellarMessage) MustTransaction() TransactionEnvelope {
	val, ok := u.GetTransaction()

	if !ok {
		panic("arm Transaction is not set")
	}

	return val
}

// GetTransaction retrieves the Transaction value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u StellarMessage) GetTransaction() (result TransactionEnvelope, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Transaction" {
		result = *u.Transaction
		ok = true
	}

	return
}

// MustQSetHash retrieves the QSetHash value from the union,
// panicing if the value is not set.
func (u StellarMessage) MustQSetHash() Uint256 {
	val, ok := u.GetQSetHash()

	if !ok {
		panic("arm QSetHash is not set")
	}

	return val
}

// GetQSetHash retrieves the QSetHash value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u StellarMessage) GetQSetHash() (result Uint256, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "QSetHash" {
		result = *u.QSetHash
		ok = true
	}

	return
}

// MustQSet retrieves the QSet value from the union,
// panicing if the value is not set.
func (u StellarMessage) MustQSet() ScpQuorumSet {
	val, ok := u.GetQSet()

	if !ok {
		panic("arm QSet is not set")
	}

	return val
}

// GetQSet retrieves the QSet value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u StellarMessage) GetQSet() (result ScpQuorumSet, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "QSet" {
		result = *u.QSet
		ok = true
	}

	return
}

// MustEnvelope retrieves the Envelope value from the union,
// panicing if the value is not set.
func (u StellarMessage) MustEnvelope() ScpEnvelope {
	val, ok := u.GetEnvelope()

	if !ok {
		panic("arm Envelope is not set")
	}

	return val
}

// GetEnvelope retrieves the Envelope value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u StellarMessage) GetEnvelope() (result ScpEnvelope, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Envelope" {
		result = *u.Envelope
		ok = true
	}

	return
}

// MustGetScpLedgerSeq retrieves the GetScpLedgerSeq value from the union,
// panicing if the value is not set.
func (u StellarMessage) MustGetScpLedgerSeq() Uint32 {
	val, ok := u.GetGetScpLedgerSeq()

	if !ok {
		panic("arm GetScpLedgerSeq is not set")
	}

	return val
}

// GetGetScpLedgerSeq retrieves the GetScpLedgerSeq value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u StellarMessage) GetGetScpLedgerSeq() (result Uint32, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "GetScpLedgerSeq" {
		result = *u.GetScpLedgerSeq
		ok = true
	}

	return
}

// AuthenticatedMessageV0 is an XDR NestedStruct defines as:
//
//   struct
//    {
//       uint64 sequence;
//       StellarMessage message;
//       HmacSha256Mac mac;
//        }
//
type AuthenticatedMessageV0 struct {
	Sequence Uint64         `json:"sequence,omitempty"`
	Message  StellarMessage `json:"message,omitempty"`
	Mac      HmacSha256Mac  `json:"mac,omitempty"`
}

// AuthenticatedMessage is an XDR Union defines as:
//
//   union AuthenticatedMessage switch (LedgerVersion v)
//    {
//    case EMPTY_VERSION:
//        struct
//    {
//       uint64 sequence;
//       StellarMessage message;
//       HmacSha256Mac mac;
//        } v0;
//    };
//
type AuthenticatedMessage struct {
	V  LedgerVersion           `json:"v,omitempty"`
	V0 *AuthenticatedMessageV0 `json:"v0,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u AuthenticatedMessage) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of AuthenticatedMessage
func (u AuthenticatedMessage) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "V0", true
	}
	return "-", false
}

// NewAuthenticatedMessage creates a new  AuthenticatedMessage.
func NewAuthenticatedMessage(v LedgerVersion, value interface{}) (result AuthenticatedMessage, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		tv, ok := value.(AuthenticatedMessageV0)
		if !ok {
			err = fmt.Errorf("invalid value, must be AuthenticatedMessageV0")
			return
		}
		result.V0 = &tv
	}
	return
}

// MustV0 retrieves the V0 value from the union,
// panicing if the value is not set.
func (u AuthenticatedMessage) MustV0() AuthenticatedMessageV0 {
	val, ok := u.GetV0()

	if !ok {
		panic("arm V0 is not set")
	}

	return val
}

// GetV0 retrieves the V0 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u AuthenticatedMessage) GetV0() (result AuthenticatedMessageV0, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.V))

	if armName == "V0" {
		result = *u.V0
		ok = true
	}

	return
}

// AmlAlertRequestExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type AmlAlertRequestExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u AmlAlertRequestExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of AmlAlertRequestExt
func (u AmlAlertRequestExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewAmlAlertRequestExt creates a new  AmlAlertRequestExt.
func NewAmlAlertRequestExt(v LedgerVersion, value interface{}) (result AmlAlertRequestExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// AmlAlertRequest is an XDR Struct defines as:
//
//   struct AMLAlertRequest {
//        BalanceID balanceID;
//        uint64 amount;
//        longstring reason;
//    	union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type AmlAlertRequest struct {
	BalanceId BalanceId          `json:"balanceID,omitempty"`
	Amount    Uint64             `json:"amount,omitempty"`
	Reason    Longstring         `json:"reason,omitempty"`
	Ext       AmlAlertRequestExt `json:"ext,omitempty"`
}

// AssetCreationRequestExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type AssetCreationRequestExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u AssetCreationRequestExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of AssetCreationRequestExt
func (u AssetCreationRequestExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewAssetCreationRequestExt creates a new  AssetCreationRequestExt.
func NewAssetCreationRequestExt(v LedgerVersion, value interface{}) (result AssetCreationRequestExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// AssetCreationRequest is an XDR Struct defines as:
//
//   struct AssetCreationRequest {
//
//    	AssetCode code;
//    	AccountID preissuedAssetSigner;
//    	uint64 maxIssuanceAmount;
//    	uint64 initialPreissuedAmount;
//        uint32 policies;
//        longstring details;
//
//    	// reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type AssetCreationRequest struct {
	Code                   AssetCode               `json:"code,omitempty"`
	PreissuedAssetSigner   AccountId               `json:"preissuedAssetSigner,omitempty"`
	MaxIssuanceAmount      Uint64                  `json:"maxIssuanceAmount,omitempty"`
	InitialPreissuedAmount Uint64                  `json:"initialPreissuedAmount,omitempty"`
	Policies               Uint32                  `json:"policies,omitempty"`
	Details                Longstring              `json:"details,omitempty"`
	Ext                    AssetCreationRequestExt `json:"ext,omitempty"`
}

// AssetUpdateRequestExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type AssetUpdateRequestExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u AssetUpdateRequestExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of AssetUpdateRequestExt
func (u AssetUpdateRequestExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewAssetUpdateRequestExt creates a new  AssetUpdateRequestExt.
func NewAssetUpdateRequestExt(v LedgerVersion, value interface{}) (result AssetUpdateRequestExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// AssetUpdateRequest is an XDR Struct defines as:
//
//   struct AssetUpdateRequest {
//    	AssetCode code;
//    	longstring details;
//    	uint32 policies;
//
//    	// reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type AssetUpdateRequest struct {
	Code     AssetCode             `json:"code,omitempty"`
	Details  Longstring            `json:"details,omitempty"`
	Policies Uint32                `json:"policies,omitempty"`
	Ext      AssetUpdateRequestExt `json:"ext,omitempty"`
}

// AssetChangePreissuedSignerExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type AssetChangePreissuedSignerExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u AssetChangePreissuedSignerExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of AssetChangePreissuedSignerExt
func (u AssetChangePreissuedSignerExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewAssetChangePreissuedSignerExt creates a new  AssetChangePreissuedSignerExt.
func NewAssetChangePreissuedSignerExt(v LedgerVersion, value interface{}) (result AssetChangePreissuedSignerExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// AssetChangePreissuedSigner is an XDR Struct defines as:
//
//   struct AssetChangePreissuedSigner {
//    	AssetCode code;
//    	AccountID accountID;
//    	// reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type AssetChangePreissuedSigner struct {
	Code      AssetCode                     `json:"code,omitempty"`
	AccountId AccountId                     `json:"accountID,omitempty"`
	Ext       AssetChangePreissuedSignerExt `json:"ext,omitempty"`
}

// ContractRequestExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type ContractRequestExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ContractRequestExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ContractRequestExt
func (u ContractRequestExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewContractRequestExt creates a new  ContractRequestExt.
func NewContractRequestExt(v LedgerVersion, value interface{}) (result ContractRequestExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ContractRequest is an XDR Struct defines as:
//
//   struct ContractRequest
//    {
//        AccountID customer;
//        AccountID escrow;
//        longstring details;
//
//        uint64 startTime;
//        uint64 endTime;
//
//        // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type ContractRequest struct {
	Customer  AccountId          `json:"customer,omitempty"`
	Escrow    AccountId          `json:"escrow,omitempty"`
	Details   Longstring         `json:"details,omitempty"`
	StartTime Uint64             `json:"startTime,omitempty"`
	EndTime   Uint64             `json:"endTime,omitempty"`
	Ext       ContractRequestExt `json:"ext,omitempty"`
}

// InvoiceRequestExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type InvoiceRequestExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u InvoiceRequestExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of InvoiceRequestExt
func (u InvoiceRequestExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewInvoiceRequestExt creates a new  InvoiceRequestExt.
func NewInvoiceRequestExt(v LedgerVersion, value interface{}) (result InvoiceRequestExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// InvoiceRequest is an XDR Struct defines as:
//
//   struct InvoiceRequest
//    {
//        AssetCode asset;
//        uint64 amount; // not allowed to set 0
//        BalanceID senderBalance;
//        BalanceID receiverBalance;
//
//        uint64 *contractID;
//        bool isApproved;
//        longstring details;
//
//        // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type InvoiceRequest struct {
	Asset           AssetCode         `json:"asset,omitempty"`
	Amount          Uint64            `json:"amount,omitempty"`
	SenderBalance   BalanceId         `json:"senderBalance,omitempty"`
	ReceiverBalance BalanceId         `json:"receiverBalance,omitempty"`
	ContractId      *Uint64           `json:"contractID,omitempty"`
	IsApproved      bool              `json:"isApproved,omitempty"`
	Details         Longstring        `json:"details,omitempty"`
	Ext             InvoiceRequestExt `json:"ext,omitempty"`
}

// PreIssuanceRequestExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type PreIssuanceRequestExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u PreIssuanceRequestExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of PreIssuanceRequestExt
func (u PreIssuanceRequestExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewPreIssuanceRequestExt creates a new  PreIssuanceRequestExt.
func NewPreIssuanceRequestExt(v LedgerVersion, value interface{}) (result PreIssuanceRequestExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// PreIssuanceRequest is an XDR Struct defines as:
//
//   struct PreIssuanceRequest {
//    	AssetCode asset;
//    	uint64 amount;
//    	DecoratedSignature signature;
//    	string64 reference;
//
//    	// reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type PreIssuanceRequest struct {
	Asset     AssetCode             `json:"asset,omitempty"`
	Amount    Uint64                `json:"amount,omitempty"`
	Signature DecoratedSignature    `json:"signature,omitempty"`
	Reference String64              `json:"reference,omitempty"`
	Ext       PreIssuanceRequestExt `json:"ext,omitempty"`
}

// IssuanceRequestExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type IssuanceRequestExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u IssuanceRequestExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of IssuanceRequestExt
func (u IssuanceRequestExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewIssuanceRequestExt creates a new  IssuanceRequestExt.
func NewIssuanceRequestExt(v LedgerVersion, value interface{}) (result IssuanceRequestExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// IssuanceRequest is an XDR Struct defines as:
//
//   struct IssuanceRequest {
//    	AssetCode asset;
//    	uint64 amount;
//    	BalanceID receiver;
//    	longstring externalDetails; // details of the issuance (External system id, etc.)
//    	Fee fee; //totalFee to be payed (calculated automatically)
//    	// reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type IssuanceRequest struct {
	Asset           AssetCode          `json:"asset,omitempty"`
	Amount          Uint64             `json:"amount,omitempty"`
	Receiver        BalanceId          `json:"receiver,omitempty"`
	ExternalDetails Longstring         `json:"externalDetails,omitempty"`
	Fee             Fee                `json:"fee,omitempty"`
	Ext             IssuanceRequestExt `json:"ext,omitempty"`
}

// LimitsUpdateRequestExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        case LIMITS_UPDATE_REQUEST_DEPRECATED_DOCUMENT_HASH:
//            longstring details;
//        }
//
type LimitsUpdateRequestExt struct {
	V       LedgerVersion `json:"v,omitempty"`
	Details *Longstring   `json:"details,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LimitsUpdateRequestExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LimitsUpdateRequestExt
func (u LimitsUpdateRequestExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	case LedgerVersionLimitsUpdateRequestDeprecatedDocumentHash:
		return "Details", true
	}
	return "-", false
}

// NewLimitsUpdateRequestExt creates a new  LimitsUpdateRequestExt.
func NewLimitsUpdateRequestExt(v LedgerVersion, value interface{}) (result LimitsUpdateRequestExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	case LedgerVersionLimitsUpdateRequestDeprecatedDocumentHash:
		tv, ok := value.(Longstring)
		if !ok {
			err = fmt.Errorf("invalid value, must be Longstring")
			return
		}
		result.Details = &tv
	}
	return
}

// MustDetails retrieves the Details value from the union,
// panicing if the value is not set.
func (u LimitsUpdateRequestExt) MustDetails() Longstring {
	val, ok := u.GetDetails()

	if !ok {
		panic("arm Details is not set")
	}

	return val
}

// GetDetails retrieves the Details value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LimitsUpdateRequestExt) GetDetails() (result Longstring, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.V))

	if armName == "Details" {
		result = *u.Details
		ok = true
	}

	return
}

// LimitsUpdateRequest is an XDR Struct defines as:
//
//   struct LimitsUpdateRequest {
//        Hash deprecatedDocumentHash;
//
//        // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        case LIMITS_UPDATE_REQUEST_DEPRECATED_DOCUMENT_HASH:
//            longstring details;
//        }
//        ext;
//    };
//
type LimitsUpdateRequest struct {
	DeprecatedDocumentHash Hash                   `json:"deprecatedDocumentHash,omitempty"`
	Ext                    LimitsUpdateRequestExt `json:"ext,omitempty"`
}

// SaleCreationRequestQuoteAssetExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type SaleCreationRequestQuoteAssetExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u SaleCreationRequestQuoteAssetExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of SaleCreationRequestQuoteAssetExt
func (u SaleCreationRequestQuoteAssetExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewSaleCreationRequestQuoteAssetExt creates a new  SaleCreationRequestQuoteAssetExt.
func NewSaleCreationRequestQuoteAssetExt(v LedgerVersion, value interface{}) (result SaleCreationRequestQuoteAssetExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// SaleCreationRequestQuoteAsset is an XDR Struct defines as:
//
//   struct SaleCreationRequestQuoteAsset {
//    	AssetCode quoteAsset; // asset in which participation will be accepted
//    	uint64 price; // price for 1 baseAsset in terms of quote asset
//    	union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type SaleCreationRequestQuoteAsset struct {
	QuoteAsset AssetCode                        `json:"quoteAsset,omitempty"`
	Price      Uint64                           `json:"price,omitempty"`
	Ext        SaleCreationRequestQuoteAssetExt `json:"ext,omitempty"`
}

// SaleCreationRequestExtV2 is an XDR NestedStruct defines as:
//
//   struct {
//                SaleTypeExt saleTypeExt;
//                uint64 requiredBaseAssetForHardCap;
//            }
//
type SaleCreationRequestExtV2 struct {
	SaleTypeExt                 SaleTypeExt `json:"saleTypeExt,omitempty"`
	RequiredBaseAssetForHardCap Uint64      `json:"requiredBaseAssetForHardCap,omitempty"`
}

// SaleCreationRequestExtV3 is an XDR NestedStruct defines as:
//
//   struct {
//    			SaleTypeExt saleTypeExt;
//                uint64 requiredBaseAssetForHardCap;
//    			SaleState state;
//    		}
//
type SaleCreationRequestExtV3 struct {
	SaleTypeExt                 SaleTypeExt `json:"saleTypeExt,omitempty"`
	RequiredBaseAssetForHardCap Uint64      `json:"requiredBaseAssetForHardCap,omitempty"`
	State                       SaleState   `json:"state,omitempty"`
}

// SaleCreationRequestExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//    	case TYPED_SALE:
//    		SaleTypeExt saleTypeExt;
//        case ALLOW_TO_SPECIFY_REQUIRED_BASE_ASSET_AMOUNT_FOR_HARD_CAP:
//            struct {
//                SaleTypeExt saleTypeExt;
//                uint64 requiredBaseAssetForHardCap;
//            } extV2;
//    	case STATABLE_SALES:
//    		struct {
//    			SaleTypeExt saleTypeExt;
//                uint64 requiredBaseAssetForHardCap;
//    			SaleState state;
//    		} extV3;
//        }
//
type SaleCreationRequestExt struct {
	V           LedgerVersion             `json:"v,omitempty"`
	SaleTypeExt *SaleTypeExt              `json:"saleTypeExt,omitempty"`
	ExtV2       *SaleCreationRequestExtV2 `json:"extV2,omitempty"`
	ExtV3       *SaleCreationRequestExtV3 `json:"extV3,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u SaleCreationRequestExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of SaleCreationRequestExt
func (u SaleCreationRequestExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	case LedgerVersionTypedSale:
		return "SaleTypeExt", true
	case LedgerVersionAllowToSpecifyRequiredBaseAssetAmountForHardCap:
		return "ExtV2", true
	case LedgerVersionStatableSales:
		return "ExtV3", true
	}
	return "-", false
}

// NewSaleCreationRequestExt creates a new  SaleCreationRequestExt.
func NewSaleCreationRequestExt(v LedgerVersion, value interface{}) (result SaleCreationRequestExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	case LedgerVersionTypedSale:
		tv, ok := value.(SaleTypeExt)
		if !ok {
			err = fmt.Errorf("invalid value, must be SaleTypeExt")
			return
		}
		result.SaleTypeExt = &tv
	case LedgerVersionAllowToSpecifyRequiredBaseAssetAmountForHardCap:
		tv, ok := value.(SaleCreationRequestExtV2)
		if !ok {
			err = fmt.Errorf("invalid value, must be SaleCreationRequestExtV2")
			return
		}
		result.ExtV2 = &tv
	case LedgerVersionStatableSales:
		tv, ok := value.(SaleCreationRequestExtV3)
		if !ok {
			err = fmt.Errorf("invalid value, must be SaleCreationRequestExtV3")
			return
		}
		result.ExtV3 = &tv
	}
	return
}

// MustSaleTypeExt retrieves the SaleTypeExt value from the union,
// panicing if the value is not set.
func (u SaleCreationRequestExt) MustSaleTypeExt() SaleTypeExt {
	val, ok := u.GetSaleTypeExt()

	if !ok {
		panic("arm SaleTypeExt is not set")
	}

	return val
}

// GetSaleTypeExt retrieves the SaleTypeExt value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u SaleCreationRequestExt) GetSaleTypeExt() (result SaleTypeExt, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.V))

	if armName == "SaleTypeExt" {
		result = *u.SaleTypeExt
		ok = true
	}

	return
}

// MustExtV2 retrieves the ExtV2 value from the union,
// panicing if the value is not set.
func (u SaleCreationRequestExt) MustExtV2() SaleCreationRequestExtV2 {
	val, ok := u.GetExtV2()

	if !ok {
		panic("arm ExtV2 is not set")
	}

	return val
}

// GetExtV2 retrieves the ExtV2 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u SaleCreationRequestExt) GetExtV2() (result SaleCreationRequestExtV2, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.V))

	if armName == "ExtV2" {
		result = *u.ExtV2
		ok = true
	}

	return
}

// MustExtV3 retrieves the ExtV3 value from the union,
// panicing if the value is not set.
func (u SaleCreationRequestExt) MustExtV3() SaleCreationRequestExtV3 {
	val, ok := u.GetExtV3()

	if !ok {
		panic("arm ExtV3 is not set")
	}

	return val
}

// GetExtV3 retrieves the ExtV3 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u SaleCreationRequestExt) GetExtV3() (result SaleCreationRequestExtV3, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.V))

	if armName == "ExtV3" {
		result = *u.ExtV3
		ok = true
	}

	return
}

// SaleCreationRequest is an XDR Struct defines as:
//
//   struct SaleCreationRequest {
//    	AssetCode baseAsset; // asset for which sale will be performed
//    	AssetCode defaultQuoteAsset; // asset for soft and hard cap
//    	uint64 startTime; // start time of the sale
//    	uint64 endTime; // close time of the sale
//    	uint64 softCap; // minimum amount of quote asset to be received at which sale will be considered a successful
//    	uint64 hardCap; // max amount of quote asset to be received
//    	longstring details; // sale specific details
//
//    	SaleCreationRequestQuoteAsset quoteAssets<100>;
//
//    	union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//    	case TYPED_SALE:
//    		SaleTypeExt saleTypeExt;
//        case ALLOW_TO_SPECIFY_REQUIRED_BASE_ASSET_AMOUNT_FOR_HARD_CAP:
//            struct {
//                SaleTypeExt saleTypeExt;
//                uint64 requiredBaseAssetForHardCap;
//            } extV2;
//    	case STATABLE_SALES:
//    		struct {
//    			SaleTypeExt saleTypeExt;
//                uint64 requiredBaseAssetForHardCap;
//    			SaleState state;
//    		} extV3;
//        }
//        ext;
//    };
//
type SaleCreationRequest struct {
	BaseAsset         AssetCode                       `json:"baseAsset,omitempty"`
	DefaultQuoteAsset AssetCode                       `json:"defaultQuoteAsset,omitempty"`
	StartTime         Uint64                          `json:"startTime,omitempty"`
	EndTime           Uint64                          `json:"endTime,omitempty"`
	SoftCap           Uint64                          `json:"softCap,omitempty"`
	HardCap           Uint64                          `json:"hardCap,omitempty"`
	Details           Longstring                      `json:"details,omitempty"`
	QuoteAssets       []SaleCreationRequestQuoteAsset `json:"quoteAssets,omitempty" xdrmaxsize:"100"`
	Ext               SaleCreationRequestExt          `json:"ext,omitempty"`
}

// UpdateKycRequestExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type UpdateKycRequestExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u UpdateKycRequestExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of UpdateKycRequestExt
func (u UpdateKycRequestExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewUpdateKycRequestExt creates a new  UpdateKycRequestExt.
func NewUpdateKycRequestExt(v LedgerVersion, value interface{}) (result UpdateKycRequestExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// UpdateKycRequest is an XDR Struct defines as:
//
//   struct UpdateKYCRequest {
//    	AccountID accountToUpdateKYC;
//    	AccountType accountTypeToSet;
//    	uint32 kycLevel;
//    	longstring kycData;
//
//    	// Tasks are represented by a bit mask. Each flag(task) in mask refers to specific KYC data validity checker
//    	uint32 allTasks;
//    	uint32 pendingTasks;
//
//    	// Sequence number increases when request is rejected
//    	uint32 sequenceNumber;
//
//    	// External details vector consists of comments written by KYC data validity checkers
//    	longstring externalDetails<>;
//
//    	// Reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type UpdateKycRequest struct {
	AccountToUpdateKyc AccountId           `json:"accountToUpdateKYC,omitempty"`
	AccountTypeToSet   AccountType         `json:"accountTypeToSet,omitempty"`
	KycLevel           Uint32              `json:"kycLevel,omitempty"`
	KycData            Longstring          `json:"kycData,omitempty"`
	AllTasks           Uint32              `json:"allTasks,omitempty"`
	PendingTasks       Uint32              `json:"pendingTasks,omitempty"`
	SequenceNumber     Uint32              `json:"sequenceNumber,omitempty"`
	ExternalDetails    []Longstring        `json:"externalDetails,omitempty"`
	Ext                UpdateKycRequestExt `json:"ext,omitempty"`
}

// PromotionUpdateRequestExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type PromotionUpdateRequestExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u PromotionUpdateRequestExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of PromotionUpdateRequestExt
func (u PromotionUpdateRequestExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewPromotionUpdateRequestExt creates a new  PromotionUpdateRequestExt.
func NewPromotionUpdateRequestExt(v LedgerVersion, value interface{}) (result PromotionUpdateRequestExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// PromotionUpdateRequest is an XDR Struct defines as:
//
//   struct PromotionUpdateRequest {
//        uint64 promotionID;
//        SaleCreationRequest newPromotionData;
//
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type PromotionUpdateRequest struct {
	PromotionId      Uint64                    `json:"promotionID,omitempty"`
	NewPromotionData SaleCreationRequest       `json:"newPromotionData,omitempty"`
	Ext              PromotionUpdateRequestExt `json:"ext,omitempty"`
}

// UpdateSaleDetailsRequestExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type UpdateSaleDetailsRequestExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u UpdateSaleDetailsRequestExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of UpdateSaleDetailsRequestExt
func (u UpdateSaleDetailsRequestExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewUpdateSaleDetailsRequestExt creates a new  UpdateSaleDetailsRequestExt.
func NewUpdateSaleDetailsRequestExt(v LedgerVersion, value interface{}) (result UpdateSaleDetailsRequestExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// UpdateSaleDetailsRequest is an XDR Struct defines as:
//
//   struct UpdateSaleDetailsRequest {
//        uint64 saleID; // ID of sale to update details
//        longstring newDetails;
//
//        // Reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type UpdateSaleDetailsRequest struct {
	SaleId     Uint64                      `json:"saleID,omitempty"`
	NewDetails Longstring                  `json:"newDetails,omitempty"`
	Ext        UpdateSaleDetailsRequestExt `json:"ext,omitempty"`
}

// UpdateSaleEndTimeRequestExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type UpdateSaleEndTimeRequestExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u UpdateSaleEndTimeRequestExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of UpdateSaleEndTimeRequestExt
func (u UpdateSaleEndTimeRequestExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewUpdateSaleEndTimeRequestExt creates a new  UpdateSaleEndTimeRequestExt.
func NewUpdateSaleEndTimeRequestExt(v LedgerVersion, value interface{}) (result UpdateSaleEndTimeRequestExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// UpdateSaleEndTimeRequest is an XDR Struct defines as:
//
//   struct UpdateSaleEndTimeRequest {
//        uint64 saleID; // ID of the sale to update end time
//        uint64 newEndTime;
//
//        // Reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type UpdateSaleEndTimeRequest struct {
	SaleId     Uint64                      `json:"saleID,omitempty"`
	NewEndTime Uint64                      `json:"newEndTime,omitempty"`
	Ext        UpdateSaleEndTimeRequestExt `json:"ext,omitempty"`
}

// WithdrawalType is an XDR Enum defines as:
//
//   enum WithdrawalType {
//    	AUTO_CONVERSION = 0
//    };
//
type WithdrawalType int32

const (
	WithdrawalTypeAutoConversion WithdrawalType = 0
)

var WithdrawalTypeAll = []WithdrawalType{
	WithdrawalTypeAutoConversion,
}

var withdrawalTypeMap = map[int32]string{
	0: "WithdrawalTypeAutoConversion",
}

var withdrawalTypeShortMap = map[int32]string{
	0: "auto_conversion",
}

var withdrawalTypeRevMap = map[string]int32{
	"WithdrawalTypeAutoConversion": 0,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for WithdrawalType
func (e WithdrawalType) ValidEnum(v int32) bool {
	_, ok := withdrawalTypeMap[v]
	return ok
}
func (e WithdrawalType) isFlag() bool {
	for i := len(WithdrawalTypeAll) - 1; i >= 0; i-- {
		expected := WithdrawalType(2) << uint64(len(WithdrawalTypeAll)-1) >> uint64(len(WithdrawalTypeAll)-i)
		if expected != WithdrawalTypeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e WithdrawalType) String() string {
	name, _ := withdrawalTypeMap[int32(e)]
	return name
}

func (e WithdrawalType) ShortString() string {
	name, _ := withdrawalTypeShortMap[int32(e)]
	return name
}

func (e WithdrawalType) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range WithdrawalTypeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *WithdrawalType) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = WithdrawalType(t.Value)
	return nil
}

// AutoConversionWithdrawalDetailsExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type AutoConversionWithdrawalDetailsExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u AutoConversionWithdrawalDetailsExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of AutoConversionWithdrawalDetailsExt
func (u AutoConversionWithdrawalDetailsExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewAutoConversionWithdrawalDetailsExt creates a new  AutoConversionWithdrawalDetailsExt.
func NewAutoConversionWithdrawalDetailsExt(v LedgerVersion, value interface{}) (result AutoConversionWithdrawalDetailsExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// AutoConversionWithdrawalDetails is an XDR Struct defines as:
//
//   struct AutoConversionWithdrawalDetails {
//    	AssetCode destAsset; // asset in which withdrawal will be converted
//    	uint64 expectedAmount; // expected amount to be received in specified asset
//
//    	union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type AutoConversionWithdrawalDetails struct {
	DestAsset      AssetCode                          `json:"destAsset,omitempty"`
	ExpectedAmount Uint64                             `json:"expectedAmount,omitempty"`
	Ext            AutoConversionWithdrawalDetailsExt `json:"ext,omitempty"`
}

// WithdrawalRequestDetails is an XDR NestedUnion defines as:
//
//   union switch (WithdrawalType withdrawalType) {
//    	case AUTO_CONVERSION:
//    		AutoConversionWithdrawalDetails autoConversion;
//    	}
//
type WithdrawalRequestDetails struct {
	WithdrawalType WithdrawalType                   `json:"withdrawalType,omitempty"`
	AutoConversion *AutoConversionWithdrawalDetails `json:"autoConversion,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u WithdrawalRequestDetails) SwitchFieldName() string {
	return "WithdrawalType"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of WithdrawalRequestDetails
func (u WithdrawalRequestDetails) ArmForSwitch(sw int32) (string, bool) {
	switch WithdrawalType(sw) {
	case WithdrawalTypeAutoConversion:
		return "AutoConversion", true
	}
	return "-", false
}

// NewWithdrawalRequestDetails creates a new  WithdrawalRequestDetails.
func NewWithdrawalRequestDetails(withdrawalType WithdrawalType, value interface{}) (result WithdrawalRequestDetails, err error) {
	result.WithdrawalType = withdrawalType
	switch WithdrawalType(withdrawalType) {
	case WithdrawalTypeAutoConversion:
		tv, ok := value.(AutoConversionWithdrawalDetails)
		if !ok {
			err = fmt.Errorf("invalid value, must be AutoConversionWithdrawalDetails")
			return
		}
		result.AutoConversion = &tv
	}
	return
}

// MustAutoConversion retrieves the AutoConversion value from the union,
// panicing if the value is not set.
func (u WithdrawalRequestDetails) MustAutoConversion() AutoConversionWithdrawalDetails {
	val, ok := u.GetAutoConversion()

	if !ok {
		panic("arm AutoConversion is not set")
	}

	return val
}

// GetAutoConversion retrieves the AutoConversion value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u WithdrawalRequestDetails) GetAutoConversion() (result AutoConversionWithdrawalDetails, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.WithdrawalType))

	if armName == "AutoConversion" {
		result = *u.AutoConversion
		ok = true
	}

	return
}

// WithdrawalRequestExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type WithdrawalRequestExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u WithdrawalRequestExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of WithdrawalRequestExt
func (u WithdrawalRequestExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewWithdrawalRequestExt creates a new  WithdrawalRequestExt.
func NewWithdrawalRequestExt(v LedgerVersion, value interface{}) (result WithdrawalRequestExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// WithdrawalRequest is an XDR Struct defines as:
//
//   struct WithdrawalRequest {
//    	BalanceID balance; // balance id from which withdrawal will be performed
//        uint64 amount; // amount to be withdrawn
//        uint64 universalAmount; // amount in stats asset
//    	Fee fee; // expected fee to be paid
//        longstring externalDetails; // details of the withdrawal (External system id, etc.)
//    	longstring preConfirmationDetails; // details provided by PSIM if two step withdrwal is required
//    	union switch (WithdrawalType withdrawalType) {
//    	case AUTO_CONVERSION:
//    		AutoConversionWithdrawalDetails autoConversion;
//    	} details;
//
//    	union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type WithdrawalRequest struct {
	Balance                BalanceId                `json:"balance,omitempty"`
	Amount                 Uint64                   `json:"amount,omitempty"`
	UniversalAmount        Uint64                   `json:"universalAmount,omitempty"`
	Fee                    Fee                      `json:"fee,omitempty"`
	ExternalDetails        Longstring               `json:"externalDetails,omitempty"`
	PreConfirmationDetails Longstring               `json:"preConfirmationDetails,omitempty"`
	Details                WithdrawalRequestDetails `json:"details,omitempty"`
	Ext                    WithdrawalRequestExt     `json:"ext,omitempty"`
}

// OperationBody is an XDR NestedUnion defines as:
//
//   union switch (OperationType type)
//        {
//        case CREATE_ACCOUNT:
//            CreateAccountOp createAccountOp;
//        case PAYMENT:
//            PaymentOp paymentOp;
//        case SET_OPTIONS:
//            SetOptionsOp setOptionsOp;
//    	case CREATE_ISSUANCE_REQUEST:
//    		CreateIssuanceRequestOp createIssuanceRequestOp;
//        case SET_FEES:
//            SetFeesOp setFeesOp;
//    	case MANAGE_ACCOUNT:
//    		ManageAccountOp manageAccountOp;
//    	case CREATE_WITHDRAWAL_REQUEST:
//    		CreateWithdrawalRequestOp createWithdrawalRequestOp;
//    	case MANAGE_BALANCE:
//    		ManageBalanceOp manageBalanceOp;
//        case MANAGE_ASSET:
//            ManageAssetOp manageAssetOp;
//        case CREATE_PREISSUANCE_REQUEST:
//            CreatePreIssuanceRequestOp createPreIssuanceRequest;
//        case MANAGE_LIMITS:
//            ManageLimitsOp manageLimitsOp;
//        case DIRECT_DEBIT:
//            DirectDebitOp directDebitOp;
//    	case MANAGE_ASSET_PAIR:
//    		ManageAssetPairOp manageAssetPairOp;
//    	case MANAGE_OFFER:
//    		ManageOfferOp manageOfferOp;
//        case MANAGE_INVOICE_REQUEST:
//            ManageInvoiceRequestOp manageInvoiceRequestOp;
//    	case REVIEW_REQUEST:
//    		ReviewRequestOp reviewRequestOp;
//    	case CREATE_SALE_REQUEST:
//    		CreateSaleCreationRequestOp createSaleCreationRequestOp;
//    	case CHECK_SALE_STATE:
//    		CheckSaleStateOp checkSaleStateOp;
//    	case CREATE_AML_ALERT:
//    	    CreateAMLAlertRequestOp createAMLAlertRequestOp;
//    	case MANAGE_KEY_VALUE:
//    	    ManageKeyValueOp manageKeyValueOp;
//    	case CREATE_KYC_REQUEST:
//    		CreateUpdateKYCRequestOp createUpdateKYCRequestOp;
//        case MANAGE_EXTERNAL_SYSTEM_ACCOUNT_ID_POOL_ENTRY:
//            ManageExternalSystemAccountIdPoolEntryOp manageExternalSystemAccountIdPoolEntryOp;
//        case BIND_EXTERNAL_SYSTEM_ACCOUNT_ID:
//            BindExternalSystemAccountIdOp bindExternalSystemAccountIdOp;
//        case PAYMENT_V2:
//            PaymentOpV2 paymentOpV2;
//        case MANAGE_SALE:
//            ManageSaleOp manageSaleOp;
//        case CREATE_MANAGE_LIMITS_REQUEST:
//            CreateManageLimitsRequestOp createManageLimitsRequestOp;
//        case MANAGE_CONTRACT_REQUEST:
//            ManageContractRequestOp manageContractRequestOp;
//        case MANAGE_CONTRACT:
//            ManageContractOp manageContractOp;
//        }
//
type OperationBody struct {
	Type                                     OperationType                             `json:"type,omitempty"`
	CreateAccountOp                          *CreateAccountOp                          `json:"createAccountOp,omitempty"`
	PaymentOp                                *PaymentOp                                `json:"paymentOp,omitempty"`
	SetOptionsOp                             *SetOptionsOp                             `json:"setOptionsOp,omitempty"`
	CreateIssuanceRequestOp                  *CreateIssuanceRequestOp                  `json:"createIssuanceRequestOp,omitempty"`
	SetFeesOp                                *SetFeesOp                                `json:"setFeesOp,omitempty"`
	ManageAccountOp                          *ManageAccountOp                          `json:"manageAccountOp,omitempty"`
	CreateWithdrawalRequestOp                *CreateWithdrawalRequestOp                `json:"createWithdrawalRequestOp,omitempty"`
	ManageBalanceOp                          *ManageBalanceOp                          `json:"manageBalanceOp,omitempty"`
	ManageAssetOp                            *ManageAssetOp                            `json:"manageAssetOp,omitempty"`
	CreatePreIssuanceRequest                 *CreatePreIssuanceRequestOp               `json:"createPreIssuanceRequest,omitempty"`
	ManageLimitsOp                           *ManageLimitsOp                           `json:"manageLimitsOp,omitempty"`
	DirectDebitOp                            *DirectDebitOp                            `json:"directDebitOp,omitempty"`
	ManageAssetPairOp                        *ManageAssetPairOp                        `json:"manageAssetPairOp,omitempty"`
	ManageOfferOp                            *ManageOfferOp                            `json:"manageOfferOp,omitempty"`
	ManageInvoiceRequestOp                   *ManageInvoiceRequestOp                   `json:"manageInvoiceRequestOp,omitempty"`
	ReviewRequestOp                          *ReviewRequestOp                          `json:"reviewRequestOp,omitempty"`
	CreateSaleCreationRequestOp              *CreateSaleCreationRequestOp              `json:"createSaleCreationRequestOp,omitempty"`
	CheckSaleStateOp                         *CheckSaleStateOp                         `json:"checkSaleStateOp,omitempty"`
	CreateAmlAlertRequestOp                  *CreateAmlAlertRequestOp                  `json:"createAMLAlertRequestOp,omitempty"`
	ManageKeyValueOp                         *ManageKeyValueOp                         `json:"manageKeyValueOp,omitempty"`
	CreateUpdateKycRequestOp                 *CreateUpdateKycRequestOp                 `json:"createUpdateKYCRequestOp,omitempty"`
	ManageExternalSystemAccountIdPoolEntryOp *ManageExternalSystemAccountIdPoolEntryOp `json:"manageExternalSystemAccountIdPoolEntryOp,omitempty"`
	BindExternalSystemAccountIdOp            *BindExternalSystemAccountIdOp            `json:"bindExternalSystemAccountIdOp,omitempty"`
	PaymentOpV2                              *PaymentOpV2                              `json:"paymentOpV2,omitempty"`
	ManageSaleOp                             *ManageSaleOp                             `json:"manageSaleOp,omitempty"`
	CreateManageLimitsRequestOp              *CreateManageLimitsRequestOp              `json:"createManageLimitsRequestOp,omitempty"`
	ManageContractRequestOp                  *ManageContractRequestOp                  `json:"manageContractRequestOp,omitempty"`
	ManageContractOp                         *ManageContractOp                         `json:"manageContractOp,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u OperationBody) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of OperationBody
func (u OperationBody) ArmForSwitch(sw int32) (string, bool) {
	switch OperationType(sw) {
	case OperationTypeCreateAccount:
		return "CreateAccountOp", true
	case OperationTypePayment:
		return "PaymentOp", true
	case OperationTypeSetOptions:
		return "SetOptionsOp", true
	case OperationTypeCreateIssuanceRequest:
		return "CreateIssuanceRequestOp", true
	case OperationTypeSetFees:
		return "SetFeesOp", true
	case OperationTypeManageAccount:
		return "ManageAccountOp", true
	case OperationTypeCreateWithdrawalRequest:
		return "CreateWithdrawalRequestOp", true
	case OperationTypeManageBalance:
		return "ManageBalanceOp", true
	case OperationTypeManageAsset:
		return "ManageAssetOp", true
	case OperationTypeCreatePreissuanceRequest:
		return "CreatePreIssuanceRequest", true
	case OperationTypeManageLimits:
		return "ManageLimitsOp", true
	case OperationTypeDirectDebit:
		return "DirectDebitOp", true
	case OperationTypeManageAssetPair:
		return "ManageAssetPairOp", true
	case OperationTypeManageOffer:
		return "ManageOfferOp", true
	case OperationTypeManageInvoiceRequest:
		return "ManageInvoiceRequestOp", true
	case OperationTypeReviewRequest:
		return "ReviewRequestOp", true
	case OperationTypeCreateSaleRequest:
		return "CreateSaleCreationRequestOp", true
	case OperationTypeCheckSaleState:
		return "CheckSaleStateOp", true
	case OperationTypeCreateAmlAlert:
		return "CreateAmlAlertRequestOp", true
	case OperationTypeManageKeyValue:
		return "ManageKeyValueOp", true
	case OperationTypeCreateKycRequest:
		return "CreateUpdateKycRequestOp", true
	case OperationTypeManageExternalSystemAccountIdPoolEntry:
		return "ManageExternalSystemAccountIdPoolEntryOp", true
	case OperationTypeBindExternalSystemAccountId:
		return "BindExternalSystemAccountIdOp", true
	case OperationTypePaymentV2:
		return "PaymentOpV2", true
	case OperationTypeManageSale:
		return "ManageSaleOp", true
	case OperationTypeCreateManageLimitsRequest:
		return "CreateManageLimitsRequestOp", true
	case OperationTypeManageContractRequest:
		return "ManageContractRequestOp", true
	case OperationTypeManageContract:
		return "ManageContractOp", true
	}
	return "-", false
}

// NewOperationBody creates a new  OperationBody.
func NewOperationBody(aType OperationType, value interface{}) (result OperationBody, err error) {
	result.Type = aType
	switch OperationType(aType) {
	case OperationTypeCreateAccount:
		tv, ok := value.(CreateAccountOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateAccountOp")
			return
		}
		result.CreateAccountOp = &tv
	case OperationTypePayment:
		tv, ok := value.(PaymentOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be PaymentOp")
			return
		}
		result.PaymentOp = &tv
	case OperationTypeSetOptions:
		tv, ok := value.(SetOptionsOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be SetOptionsOp")
			return
		}
		result.SetOptionsOp = &tv
	case OperationTypeCreateIssuanceRequest:
		tv, ok := value.(CreateIssuanceRequestOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateIssuanceRequestOp")
			return
		}
		result.CreateIssuanceRequestOp = &tv
	case OperationTypeSetFees:
		tv, ok := value.(SetFeesOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be SetFeesOp")
			return
		}
		result.SetFeesOp = &tv
	case OperationTypeManageAccount:
		tv, ok := value.(ManageAccountOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageAccountOp")
			return
		}
		result.ManageAccountOp = &tv
	case OperationTypeCreateWithdrawalRequest:
		tv, ok := value.(CreateWithdrawalRequestOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateWithdrawalRequestOp")
			return
		}
		result.CreateWithdrawalRequestOp = &tv
	case OperationTypeManageBalance:
		tv, ok := value.(ManageBalanceOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageBalanceOp")
			return
		}
		result.ManageBalanceOp = &tv
	case OperationTypeManageAsset:
		tv, ok := value.(ManageAssetOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageAssetOp")
			return
		}
		result.ManageAssetOp = &tv
	case OperationTypeCreatePreissuanceRequest:
		tv, ok := value.(CreatePreIssuanceRequestOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreatePreIssuanceRequestOp")
			return
		}
		result.CreatePreIssuanceRequest = &tv
	case OperationTypeManageLimits:
		tv, ok := value.(ManageLimitsOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageLimitsOp")
			return
		}
		result.ManageLimitsOp = &tv
	case OperationTypeDirectDebit:
		tv, ok := value.(DirectDebitOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be DirectDebitOp")
			return
		}
		result.DirectDebitOp = &tv
	case OperationTypeManageAssetPair:
		tv, ok := value.(ManageAssetPairOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageAssetPairOp")
			return
		}
		result.ManageAssetPairOp = &tv
	case OperationTypeManageOffer:
		tv, ok := value.(ManageOfferOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageOfferOp")
			return
		}
		result.ManageOfferOp = &tv
	case OperationTypeManageInvoiceRequest:
		tv, ok := value.(ManageInvoiceRequestOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageInvoiceRequestOp")
			return
		}
		result.ManageInvoiceRequestOp = &tv
	case OperationTypeReviewRequest:
		tv, ok := value.(ReviewRequestOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be ReviewRequestOp")
			return
		}
		result.ReviewRequestOp = &tv
	case OperationTypeCreateSaleRequest:
		tv, ok := value.(CreateSaleCreationRequestOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateSaleCreationRequestOp")
			return
		}
		result.CreateSaleCreationRequestOp = &tv
	case OperationTypeCheckSaleState:
		tv, ok := value.(CheckSaleStateOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be CheckSaleStateOp")
			return
		}
		result.CheckSaleStateOp = &tv
	case OperationTypeCreateAmlAlert:
		tv, ok := value.(CreateAmlAlertRequestOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateAmlAlertRequestOp")
			return
		}
		result.CreateAmlAlertRequestOp = &tv
	case OperationTypeManageKeyValue:
		tv, ok := value.(ManageKeyValueOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageKeyValueOp")
			return
		}
		result.ManageKeyValueOp = &tv
	case OperationTypeCreateKycRequest:
		tv, ok := value.(CreateUpdateKycRequestOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateUpdateKycRequestOp")
			return
		}
		result.CreateUpdateKycRequestOp = &tv
	case OperationTypeManageExternalSystemAccountIdPoolEntry:
		tv, ok := value.(ManageExternalSystemAccountIdPoolEntryOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageExternalSystemAccountIdPoolEntryOp")
			return
		}
		result.ManageExternalSystemAccountIdPoolEntryOp = &tv
	case OperationTypeBindExternalSystemAccountId:
		tv, ok := value.(BindExternalSystemAccountIdOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be BindExternalSystemAccountIdOp")
			return
		}
		result.BindExternalSystemAccountIdOp = &tv
	case OperationTypePaymentV2:
		tv, ok := value.(PaymentOpV2)
		if !ok {
			err = fmt.Errorf("invalid value, must be PaymentOpV2")
			return
		}
		result.PaymentOpV2 = &tv
	case OperationTypeManageSale:
		tv, ok := value.(ManageSaleOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageSaleOp")
			return
		}
		result.ManageSaleOp = &tv
	case OperationTypeCreateManageLimitsRequest:
		tv, ok := value.(CreateManageLimitsRequestOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateManageLimitsRequestOp")
			return
		}
		result.CreateManageLimitsRequestOp = &tv
	case OperationTypeManageContractRequest:
		tv, ok := value.(ManageContractRequestOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageContractRequestOp")
			return
		}
		result.ManageContractRequestOp = &tv
	case OperationTypeManageContract:
		tv, ok := value.(ManageContractOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageContractOp")
			return
		}
		result.ManageContractOp = &tv
	}
	return
}

// MustCreateAccountOp retrieves the CreateAccountOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustCreateAccountOp() CreateAccountOp {
	val, ok := u.GetCreateAccountOp()

	if !ok {
		panic("arm CreateAccountOp is not set")
	}

	return val
}

// GetCreateAccountOp retrieves the CreateAccountOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetCreateAccountOp() (result CreateAccountOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CreateAccountOp" {
		result = *u.CreateAccountOp
		ok = true
	}

	return
}

// MustPaymentOp retrieves the PaymentOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustPaymentOp() PaymentOp {
	val, ok := u.GetPaymentOp()

	if !ok {
		panic("arm PaymentOp is not set")
	}

	return val
}

// GetPaymentOp retrieves the PaymentOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetPaymentOp() (result PaymentOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "PaymentOp" {
		result = *u.PaymentOp
		ok = true
	}

	return
}

// MustSetOptionsOp retrieves the SetOptionsOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustSetOptionsOp() SetOptionsOp {
	val, ok := u.GetSetOptionsOp()

	if !ok {
		panic("arm SetOptionsOp is not set")
	}

	return val
}

// GetSetOptionsOp retrieves the SetOptionsOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetSetOptionsOp() (result SetOptionsOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "SetOptionsOp" {
		result = *u.SetOptionsOp
		ok = true
	}

	return
}

// MustCreateIssuanceRequestOp retrieves the CreateIssuanceRequestOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustCreateIssuanceRequestOp() CreateIssuanceRequestOp {
	val, ok := u.GetCreateIssuanceRequestOp()

	if !ok {
		panic("arm CreateIssuanceRequestOp is not set")
	}

	return val
}

// GetCreateIssuanceRequestOp retrieves the CreateIssuanceRequestOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetCreateIssuanceRequestOp() (result CreateIssuanceRequestOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CreateIssuanceRequestOp" {
		result = *u.CreateIssuanceRequestOp
		ok = true
	}

	return
}

// MustSetFeesOp retrieves the SetFeesOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustSetFeesOp() SetFeesOp {
	val, ok := u.GetSetFeesOp()

	if !ok {
		panic("arm SetFeesOp is not set")
	}

	return val
}

// GetSetFeesOp retrieves the SetFeesOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetSetFeesOp() (result SetFeesOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "SetFeesOp" {
		result = *u.SetFeesOp
		ok = true
	}

	return
}

// MustManageAccountOp retrieves the ManageAccountOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustManageAccountOp() ManageAccountOp {
	val, ok := u.GetManageAccountOp()

	if !ok {
		panic("arm ManageAccountOp is not set")
	}

	return val
}

// GetManageAccountOp retrieves the ManageAccountOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetManageAccountOp() (result ManageAccountOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ManageAccountOp" {
		result = *u.ManageAccountOp
		ok = true
	}

	return
}

// MustCreateWithdrawalRequestOp retrieves the CreateWithdrawalRequestOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustCreateWithdrawalRequestOp() CreateWithdrawalRequestOp {
	val, ok := u.GetCreateWithdrawalRequestOp()

	if !ok {
		panic("arm CreateWithdrawalRequestOp is not set")
	}

	return val
}

// GetCreateWithdrawalRequestOp retrieves the CreateWithdrawalRequestOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetCreateWithdrawalRequestOp() (result CreateWithdrawalRequestOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CreateWithdrawalRequestOp" {
		result = *u.CreateWithdrawalRequestOp
		ok = true
	}

	return
}

// MustManageBalanceOp retrieves the ManageBalanceOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustManageBalanceOp() ManageBalanceOp {
	val, ok := u.GetManageBalanceOp()

	if !ok {
		panic("arm ManageBalanceOp is not set")
	}

	return val
}

// GetManageBalanceOp retrieves the ManageBalanceOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetManageBalanceOp() (result ManageBalanceOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ManageBalanceOp" {
		result = *u.ManageBalanceOp
		ok = true
	}

	return
}

// MustManageAssetOp retrieves the ManageAssetOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustManageAssetOp() ManageAssetOp {
	val, ok := u.GetManageAssetOp()

	if !ok {
		panic("arm ManageAssetOp is not set")
	}

	return val
}

// GetManageAssetOp retrieves the ManageAssetOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetManageAssetOp() (result ManageAssetOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ManageAssetOp" {
		result = *u.ManageAssetOp
		ok = true
	}

	return
}

// MustCreatePreIssuanceRequest retrieves the CreatePreIssuanceRequest value from the union,
// panicing if the value is not set.
func (u OperationBody) MustCreatePreIssuanceRequest() CreatePreIssuanceRequestOp {
	val, ok := u.GetCreatePreIssuanceRequest()

	if !ok {
		panic("arm CreatePreIssuanceRequest is not set")
	}

	return val
}

// GetCreatePreIssuanceRequest retrieves the CreatePreIssuanceRequest value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetCreatePreIssuanceRequest() (result CreatePreIssuanceRequestOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CreatePreIssuanceRequest" {
		result = *u.CreatePreIssuanceRequest
		ok = true
	}

	return
}

// MustManageLimitsOp retrieves the ManageLimitsOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustManageLimitsOp() ManageLimitsOp {
	val, ok := u.GetManageLimitsOp()

	if !ok {
		panic("arm ManageLimitsOp is not set")
	}

	return val
}

// GetManageLimitsOp retrieves the ManageLimitsOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetManageLimitsOp() (result ManageLimitsOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ManageLimitsOp" {
		result = *u.ManageLimitsOp
		ok = true
	}

	return
}

// MustDirectDebitOp retrieves the DirectDebitOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustDirectDebitOp() DirectDebitOp {
	val, ok := u.GetDirectDebitOp()

	if !ok {
		panic("arm DirectDebitOp is not set")
	}

	return val
}

// GetDirectDebitOp retrieves the DirectDebitOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetDirectDebitOp() (result DirectDebitOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "DirectDebitOp" {
		result = *u.DirectDebitOp
		ok = true
	}

	return
}

// MustManageAssetPairOp retrieves the ManageAssetPairOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustManageAssetPairOp() ManageAssetPairOp {
	val, ok := u.GetManageAssetPairOp()

	if !ok {
		panic("arm ManageAssetPairOp is not set")
	}

	return val
}

// GetManageAssetPairOp retrieves the ManageAssetPairOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetManageAssetPairOp() (result ManageAssetPairOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ManageAssetPairOp" {
		result = *u.ManageAssetPairOp
		ok = true
	}

	return
}

// MustManageOfferOp retrieves the ManageOfferOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustManageOfferOp() ManageOfferOp {
	val, ok := u.GetManageOfferOp()

	if !ok {
		panic("arm ManageOfferOp is not set")
	}

	return val
}

// GetManageOfferOp retrieves the ManageOfferOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetManageOfferOp() (result ManageOfferOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ManageOfferOp" {
		result = *u.ManageOfferOp
		ok = true
	}

	return
}

// MustManageInvoiceRequestOp retrieves the ManageInvoiceRequestOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustManageInvoiceRequestOp() ManageInvoiceRequestOp {
	val, ok := u.GetManageInvoiceRequestOp()

	if !ok {
		panic("arm ManageInvoiceRequestOp is not set")
	}

	return val
}

// GetManageInvoiceRequestOp retrieves the ManageInvoiceRequestOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetManageInvoiceRequestOp() (result ManageInvoiceRequestOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ManageInvoiceRequestOp" {
		result = *u.ManageInvoiceRequestOp
		ok = true
	}

	return
}

// MustReviewRequestOp retrieves the ReviewRequestOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustReviewRequestOp() ReviewRequestOp {
	val, ok := u.GetReviewRequestOp()

	if !ok {
		panic("arm ReviewRequestOp is not set")
	}

	return val
}

// GetReviewRequestOp retrieves the ReviewRequestOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetReviewRequestOp() (result ReviewRequestOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ReviewRequestOp" {
		result = *u.ReviewRequestOp
		ok = true
	}

	return
}

// MustCreateSaleCreationRequestOp retrieves the CreateSaleCreationRequestOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustCreateSaleCreationRequestOp() CreateSaleCreationRequestOp {
	val, ok := u.GetCreateSaleCreationRequestOp()

	if !ok {
		panic("arm CreateSaleCreationRequestOp is not set")
	}

	return val
}

// GetCreateSaleCreationRequestOp retrieves the CreateSaleCreationRequestOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetCreateSaleCreationRequestOp() (result CreateSaleCreationRequestOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CreateSaleCreationRequestOp" {
		result = *u.CreateSaleCreationRequestOp
		ok = true
	}

	return
}

// MustCheckSaleStateOp retrieves the CheckSaleStateOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustCheckSaleStateOp() CheckSaleStateOp {
	val, ok := u.GetCheckSaleStateOp()

	if !ok {
		panic("arm CheckSaleStateOp is not set")
	}

	return val
}

// GetCheckSaleStateOp retrieves the CheckSaleStateOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetCheckSaleStateOp() (result CheckSaleStateOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CheckSaleStateOp" {
		result = *u.CheckSaleStateOp
		ok = true
	}

	return
}

// MustCreateAmlAlertRequestOp retrieves the CreateAmlAlertRequestOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustCreateAmlAlertRequestOp() CreateAmlAlertRequestOp {
	val, ok := u.GetCreateAmlAlertRequestOp()

	if !ok {
		panic("arm CreateAmlAlertRequestOp is not set")
	}

	return val
}

// GetCreateAmlAlertRequestOp retrieves the CreateAmlAlertRequestOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetCreateAmlAlertRequestOp() (result CreateAmlAlertRequestOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CreateAmlAlertRequestOp" {
		result = *u.CreateAmlAlertRequestOp
		ok = true
	}

	return
}

// MustManageKeyValueOp retrieves the ManageKeyValueOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustManageKeyValueOp() ManageKeyValueOp {
	val, ok := u.GetManageKeyValueOp()

	if !ok {
		panic("arm ManageKeyValueOp is not set")
	}

	return val
}

// GetManageKeyValueOp retrieves the ManageKeyValueOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetManageKeyValueOp() (result ManageKeyValueOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ManageKeyValueOp" {
		result = *u.ManageKeyValueOp
		ok = true
	}

	return
}

// MustCreateUpdateKycRequestOp retrieves the CreateUpdateKycRequestOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustCreateUpdateKycRequestOp() CreateUpdateKycRequestOp {
	val, ok := u.GetCreateUpdateKycRequestOp()

	if !ok {
		panic("arm CreateUpdateKycRequestOp is not set")
	}

	return val
}

// GetCreateUpdateKycRequestOp retrieves the CreateUpdateKycRequestOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetCreateUpdateKycRequestOp() (result CreateUpdateKycRequestOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CreateUpdateKycRequestOp" {
		result = *u.CreateUpdateKycRequestOp
		ok = true
	}

	return
}

// MustManageExternalSystemAccountIdPoolEntryOp retrieves the ManageExternalSystemAccountIdPoolEntryOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustManageExternalSystemAccountIdPoolEntryOp() ManageExternalSystemAccountIdPoolEntryOp {
	val, ok := u.GetManageExternalSystemAccountIdPoolEntryOp()

	if !ok {
		panic("arm ManageExternalSystemAccountIdPoolEntryOp is not set")
	}

	return val
}

// GetManageExternalSystemAccountIdPoolEntryOp retrieves the ManageExternalSystemAccountIdPoolEntryOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetManageExternalSystemAccountIdPoolEntryOp() (result ManageExternalSystemAccountIdPoolEntryOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ManageExternalSystemAccountIdPoolEntryOp" {
		result = *u.ManageExternalSystemAccountIdPoolEntryOp
		ok = true
	}

	return
}

// MustBindExternalSystemAccountIdOp retrieves the BindExternalSystemAccountIdOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustBindExternalSystemAccountIdOp() BindExternalSystemAccountIdOp {
	val, ok := u.GetBindExternalSystemAccountIdOp()

	if !ok {
		panic("arm BindExternalSystemAccountIdOp is not set")
	}

	return val
}

// GetBindExternalSystemAccountIdOp retrieves the BindExternalSystemAccountIdOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetBindExternalSystemAccountIdOp() (result BindExternalSystemAccountIdOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "BindExternalSystemAccountIdOp" {
		result = *u.BindExternalSystemAccountIdOp
		ok = true
	}

	return
}

// MustPaymentOpV2 retrieves the PaymentOpV2 value from the union,
// panicing if the value is not set.
func (u OperationBody) MustPaymentOpV2() PaymentOpV2 {
	val, ok := u.GetPaymentOpV2()

	if !ok {
		panic("arm PaymentOpV2 is not set")
	}

	return val
}

// GetPaymentOpV2 retrieves the PaymentOpV2 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetPaymentOpV2() (result PaymentOpV2, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "PaymentOpV2" {
		result = *u.PaymentOpV2
		ok = true
	}

	return
}

// MustManageSaleOp retrieves the ManageSaleOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustManageSaleOp() ManageSaleOp {
	val, ok := u.GetManageSaleOp()

	if !ok {
		panic("arm ManageSaleOp is not set")
	}

	return val
}

// GetManageSaleOp retrieves the ManageSaleOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetManageSaleOp() (result ManageSaleOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ManageSaleOp" {
		result = *u.ManageSaleOp
		ok = true
	}

	return
}

// MustCreateManageLimitsRequestOp retrieves the CreateManageLimitsRequestOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustCreateManageLimitsRequestOp() CreateManageLimitsRequestOp {
	val, ok := u.GetCreateManageLimitsRequestOp()

	if !ok {
		panic("arm CreateManageLimitsRequestOp is not set")
	}

	return val
}

// GetCreateManageLimitsRequestOp retrieves the CreateManageLimitsRequestOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetCreateManageLimitsRequestOp() (result CreateManageLimitsRequestOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CreateManageLimitsRequestOp" {
		result = *u.CreateManageLimitsRequestOp
		ok = true
	}

	return
}

// MustManageContractRequestOp retrieves the ManageContractRequestOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustManageContractRequestOp() ManageContractRequestOp {
	val, ok := u.GetManageContractRequestOp()

	if !ok {
		panic("arm ManageContractRequestOp is not set")
	}

	return val
}

// GetManageContractRequestOp retrieves the ManageContractRequestOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetManageContractRequestOp() (result ManageContractRequestOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ManageContractRequestOp" {
		result = *u.ManageContractRequestOp
		ok = true
	}

	return
}

// MustManageContractOp retrieves the ManageContractOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustManageContractOp() ManageContractOp {
	val, ok := u.GetManageContractOp()

	if !ok {
		panic("arm ManageContractOp is not set")
	}

	return val
}

// GetManageContractOp retrieves the ManageContractOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetManageContractOp() (result ManageContractOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ManageContractOp" {
		result = *u.ManageContractOp
		ok = true
	}

	return
}

// Operation is an XDR Struct defines as:
//
//   struct Operation
//    {
//        // sourceAccount is the account used to run the operation
//        // if not set, the runtime defaults to "sourceAccount" specified at
//        // the transaction level
//        AccountID* sourceAccount;
//
//        union switch (OperationType type)
//        {
//        case CREATE_ACCOUNT:
//            CreateAccountOp createAccountOp;
//        case PAYMENT:
//            PaymentOp paymentOp;
//        case SET_OPTIONS:
//            SetOptionsOp setOptionsOp;
//    	case CREATE_ISSUANCE_REQUEST:
//    		CreateIssuanceRequestOp createIssuanceRequestOp;
//        case SET_FEES:
//            SetFeesOp setFeesOp;
//    	case MANAGE_ACCOUNT:
//    		ManageAccountOp manageAccountOp;
//    	case CREATE_WITHDRAWAL_REQUEST:
//    		CreateWithdrawalRequestOp createWithdrawalRequestOp;
//    	case MANAGE_BALANCE:
//    		ManageBalanceOp manageBalanceOp;
//        case MANAGE_ASSET:
//            ManageAssetOp manageAssetOp;
//        case CREATE_PREISSUANCE_REQUEST:
//            CreatePreIssuanceRequestOp createPreIssuanceRequest;
//        case MANAGE_LIMITS:
//            ManageLimitsOp manageLimitsOp;
//        case DIRECT_DEBIT:
//            DirectDebitOp directDebitOp;
//    	case MANAGE_ASSET_PAIR:
//    		ManageAssetPairOp manageAssetPairOp;
//    	case MANAGE_OFFER:
//    		ManageOfferOp manageOfferOp;
//        case MANAGE_INVOICE_REQUEST:
//            ManageInvoiceRequestOp manageInvoiceRequestOp;
//    	case REVIEW_REQUEST:
//    		ReviewRequestOp reviewRequestOp;
//    	case CREATE_SALE_REQUEST:
//    		CreateSaleCreationRequestOp createSaleCreationRequestOp;
//    	case CHECK_SALE_STATE:
//    		CheckSaleStateOp checkSaleStateOp;
//    	case CREATE_AML_ALERT:
//    	    CreateAMLAlertRequestOp createAMLAlertRequestOp;
//    	case MANAGE_KEY_VALUE:
//    	    ManageKeyValueOp manageKeyValueOp;
//    	case CREATE_KYC_REQUEST:
//    		CreateUpdateKYCRequestOp createUpdateKYCRequestOp;
//        case MANAGE_EXTERNAL_SYSTEM_ACCOUNT_ID_POOL_ENTRY:
//            ManageExternalSystemAccountIdPoolEntryOp manageExternalSystemAccountIdPoolEntryOp;
//        case BIND_EXTERNAL_SYSTEM_ACCOUNT_ID:
//            BindExternalSystemAccountIdOp bindExternalSystemAccountIdOp;
//        case PAYMENT_V2:
//            PaymentOpV2 paymentOpV2;
//        case MANAGE_SALE:
//            ManageSaleOp manageSaleOp;
//        case CREATE_MANAGE_LIMITS_REQUEST:
//            CreateManageLimitsRequestOp createManageLimitsRequestOp;
//        case MANAGE_CONTRACT_REQUEST:
//            ManageContractRequestOp manageContractRequestOp;
//        case MANAGE_CONTRACT:
//            ManageContractOp manageContractOp;
//        }
//        body;
//    };
//
type Operation struct {
	SourceAccount *AccountId    `json:"sourceAccount,omitempty"`
	Body          OperationBody `json:"body,omitempty"`
}

// MemoType is an XDR Enum defines as:
//
//   enum MemoType
//    {
//        MEMO_NONE = 0,
//        MEMO_TEXT = 1,
//        MEMO_ID = 2,
//        MEMO_HASH = 3,
//        MEMO_RETURN = 4
//    };
//
type MemoType int32

const (
	MemoTypeMemoNone   MemoType = 0
	MemoTypeMemoText   MemoType = 1
	MemoTypeMemoId     MemoType = 2
	MemoTypeMemoHash   MemoType = 3
	MemoTypeMemoReturn MemoType = 4
)

var MemoTypeAll = []MemoType{
	MemoTypeMemoNone,
	MemoTypeMemoText,
	MemoTypeMemoId,
	MemoTypeMemoHash,
	MemoTypeMemoReturn,
}

var memoTypeMap = map[int32]string{
	0: "MemoTypeMemoNone",
	1: "MemoTypeMemoText",
	2: "MemoTypeMemoId",
	3: "MemoTypeMemoHash",
	4: "MemoTypeMemoReturn",
}

var memoTypeShortMap = map[int32]string{
	0: "memo_none",
	1: "memo_text",
	2: "memo_id",
	3: "memo_hash",
	4: "memo_return",
}

var memoTypeRevMap = map[string]int32{
	"MemoTypeMemoNone":   0,
	"MemoTypeMemoText":   1,
	"MemoTypeMemoId":     2,
	"MemoTypeMemoHash":   3,
	"MemoTypeMemoReturn": 4,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for MemoType
func (e MemoType) ValidEnum(v int32) bool {
	_, ok := memoTypeMap[v]
	return ok
}
func (e MemoType) isFlag() bool {
	for i := len(MemoTypeAll) - 1; i >= 0; i-- {
		expected := MemoType(2) << uint64(len(MemoTypeAll)-1) >> uint64(len(MemoTypeAll)-i)
		if expected != MemoTypeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e MemoType) String() string {
	name, _ := memoTypeMap[int32(e)]
	return name
}

func (e MemoType) ShortString() string {
	name, _ := memoTypeShortMap[int32(e)]
	return name
}

func (e MemoType) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range MemoTypeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *MemoType) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = MemoType(t.Value)
	return nil
}

// Memo is an XDR Union defines as:
//
//   union Memo switch (MemoType type)
//    {
//    case MEMO_NONE:
//        void;
//    case MEMO_TEXT:
//        string text<28>;
//    case MEMO_ID:
//        uint64 id;
//    case MEMO_HASH:
//        Hash hash; // the hash of what to pull from the content server
//    case MEMO_RETURN:
//        Hash retHash; // the hash of the tx you are rejecting
//    };
//
type Memo struct {
	Type    MemoType `json:"type,omitempty"`
	Text    *string  `json:"text,omitempty" xdrmaxsize:"28"`
	Id      *Uint64  `json:"id,omitempty"`
	Hash    *Hash    `json:"hash,omitempty"`
	RetHash *Hash    `json:"retHash,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u Memo) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of Memo
func (u Memo) ArmForSwitch(sw int32) (string, bool) {
	switch MemoType(sw) {
	case MemoTypeMemoNone:
		return "", true
	case MemoTypeMemoText:
		return "Text", true
	case MemoTypeMemoId:
		return "Id", true
	case MemoTypeMemoHash:
		return "Hash", true
	case MemoTypeMemoReturn:
		return "RetHash", true
	}
	return "-", false
}

// NewMemo creates a new  Memo.
func NewMemo(aType MemoType, value interface{}) (result Memo, err error) {
	result.Type = aType
	switch MemoType(aType) {
	case MemoTypeMemoNone:
		// void
	case MemoTypeMemoText:
		tv, ok := value.(string)
		if !ok {
			err = fmt.Errorf("invalid value, must be string")
			return
		}
		result.Text = &tv
	case MemoTypeMemoId:
		tv, ok := value.(Uint64)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint64")
			return
		}
		result.Id = &tv
	case MemoTypeMemoHash:
		tv, ok := value.(Hash)
		if !ok {
			err = fmt.Errorf("invalid value, must be Hash")
			return
		}
		result.Hash = &tv
	case MemoTypeMemoReturn:
		tv, ok := value.(Hash)
		if !ok {
			err = fmt.Errorf("invalid value, must be Hash")
			return
		}
		result.RetHash = &tv
	}
	return
}

// MustText retrieves the Text value from the union,
// panicing if the value is not set.
func (u Memo) MustText() string {
	val, ok := u.GetText()

	if !ok {
		panic("arm Text is not set")
	}

	return val
}

// GetText retrieves the Text value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u Memo) GetText() (result string, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Text" {
		result = *u.Text
		ok = true
	}

	return
}

// MustId retrieves the Id value from the union,
// panicing if the value is not set.
func (u Memo) MustId() Uint64 {
	val, ok := u.GetId()

	if !ok {
		panic("arm Id is not set")
	}

	return val
}

// GetId retrieves the Id value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u Memo) GetId() (result Uint64, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Id" {
		result = *u.Id
		ok = true
	}

	return
}

// MustHash retrieves the Hash value from the union,
// panicing if the value is not set.
func (u Memo) MustHash() Hash {
	val, ok := u.GetHash()

	if !ok {
		panic("arm Hash is not set")
	}

	return val
}

// GetHash retrieves the Hash value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u Memo) GetHash() (result Hash, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Hash" {
		result = *u.Hash
		ok = true
	}

	return
}

// MustRetHash retrieves the RetHash value from the union,
// panicing if the value is not set.
func (u Memo) MustRetHash() Hash {
	val, ok := u.GetRetHash()

	if !ok {
		panic("arm RetHash is not set")
	}

	return val
}

// GetRetHash retrieves the RetHash value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u Memo) GetRetHash() (result Hash, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "RetHash" {
		result = *u.RetHash
		ok = true
	}

	return
}

// TimeBounds is an XDR Struct defines as:
//
//   struct TimeBounds
//    {
//        uint64 minTime;
//        uint64 maxTime; // 0 here means no maxTime
//    };
//
type TimeBounds struct {
	MinTime Uint64 `json:"minTime,omitempty"`
	MaxTime Uint64 `json:"maxTime,omitempty"`
}

// TransactionExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type TransactionExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u TransactionExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of TransactionExt
func (u TransactionExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewTransactionExt creates a new  TransactionExt.
func NewTransactionExt(v LedgerVersion, value interface{}) (result TransactionExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// Transaction is an XDR Struct defines as:
//
//   struct Transaction
//    {
//        // account used to run the transaction
//        AccountID sourceAccount;
//
//        Salt salt;
//
//        // validity range (inclusive) for the last ledger close time
//        TimeBounds timeBounds;
//
//        Memo memo;
//
//        Operation operations<100>;
//
//        // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type Transaction struct {
	SourceAccount AccountId      `json:"sourceAccount,omitempty"`
	Salt          Salt           `json:"salt,omitempty"`
	TimeBounds    TimeBounds     `json:"timeBounds,omitempty"`
	Memo          Memo           `json:"memo,omitempty"`
	Operations    []Operation    `json:"operations,omitempty" xdrmaxsize:"100"`
	Ext           TransactionExt `json:"ext,omitempty"`
}

// TransactionEnvelope is an XDR Struct defines as:
//
//   struct TransactionEnvelope
//    {
//        Transaction tx;
//        DecoratedSignature signatures<20>;
//    };
//
type TransactionEnvelope struct {
	Tx         Transaction          `json:"tx,omitempty"`
	Signatures []DecoratedSignature `json:"signatures,omitempty" xdrmaxsize:"20"`
}

// OperationResultCode is an XDR Enum defines as:
//
//   enum OperationResultCode
//    {
//        opINNER = 0, // inner object result is valid
//
//        opBAD_AUTH = -1,      // too few valid signatures / wrong network
//        opNO_ACCOUNT = -2,    // source account was not found
//    	opNOT_ALLOWED = -3,   // operation is not allowed for this type of source account
//    	opACCOUNT_BLOCKED = -4, // account is blocked
//        opNO_COUNTERPARTY = -5,
//        opCOUNTERPARTY_BLOCKED = -6,
//        opCOUNTERPARTY_WRONG_TYPE = -7,
//    	opBAD_AUTH_EXTRA = -8
//    };
//
type OperationResultCode int32

const (
	OperationResultCodeOpInner                 OperationResultCode = 0
	OperationResultCodeOpBadAuth               OperationResultCode = -1
	OperationResultCodeOpNoAccount             OperationResultCode = -2
	OperationResultCodeOpNotAllowed            OperationResultCode = -3
	OperationResultCodeOpAccountBlocked        OperationResultCode = -4
	OperationResultCodeOpNoCounterparty        OperationResultCode = -5
	OperationResultCodeOpCounterpartyBlocked   OperationResultCode = -6
	OperationResultCodeOpCounterpartyWrongType OperationResultCode = -7
	OperationResultCodeOpBadAuthExtra          OperationResultCode = -8
)

var OperationResultCodeAll = []OperationResultCode{
	OperationResultCodeOpInner,
	OperationResultCodeOpBadAuth,
	OperationResultCodeOpNoAccount,
	OperationResultCodeOpNotAllowed,
	OperationResultCodeOpAccountBlocked,
	OperationResultCodeOpNoCounterparty,
	OperationResultCodeOpCounterpartyBlocked,
	OperationResultCodeOpCounterpartyWrongType,
	OperationResultCodeOpBadAuthExtra,
}

var operationResultCodeMap = map[int32]string{
	0:  "OperationResultCodeOpInner",
	-1: "OperationResultCodeOpBadAuth",
	-2: "OperationResultCodeOpNoAccount",
	-3: "OperationResultCodeOpNotAllowed",
	-4: "OperationResultCodeOpAccountBlocked",
	-5: "OperationResultCodeOpNoCounterparty",
	-6: "OperationResultCodeOpCounterpartyBlocked",
	-7: "OperationResultCodeOpCounterpartyWrongType",
	-8: "OperationResultCodeOpBadAuthExtra",
}

var operationResultCodeShortMap = map[int32]string{
	0:  "op_inner",
	-1: "op_bad_auth",
	-2: "op_no_account",
	-3: "op_not_allowed",
	-4: "op_account_blocked",
	-5: "op_no_counterparty",
	-6: "op_counterparty_blocked",
	-7: "op_counterparty_wrong_type",
	-8: "op_bad_auth_extra",
}

var operationResultCodeRevMap = map[string]int32{
	"OperationResultCodeOpInner":                 0,
	"OperationResultCodeOpBadAuth":               -1,
	"OperationResultCodeOpNoAccount":             -2,
	"OperationResultCodeOpNotAllowed":            -3,
	"OperationResultCodeOpAccountBlocked":        -4,
	"OperationResultCodeOpNoCounterparty":        -5,
	"OperationResultCodeOpCounterpartyBlocked":   -6,
	"OperationResultCodeOpCounterpartyWrongType": -7,
	"OperationResultCodeOpBadAuthExtra":          -8,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for OperationResultCode
func (e OperationResultCode) ValidEnum(v int32) bool {
	_, ok := operationResultCodeMap[v]
	return ok
}
func (e OperationResultCode) isFlag() bool {
	for i := len(OperationResultCodeAll) - 1; i >= 0; i-- {
		expected := OperationResultCode(2) << uint64(len(OperationResultCodeAll)-1) >> uint64(len(OperationResultCodeAll)-i)
		if expected != OperationResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e OperationResultCode) String() string {
	name, _ := operationResultCodeMap[int32(e)]
	return name
}

func (e OperationResultCode) ShortString() string {
	name, _ := operationResultCodeShortMap[int32(e)]
	return name
}

func (e OperationResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range OperationResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *OperationResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = OperationResultCode(t.Value)
	return nil
}

// OperationResultTr is an XDR NestedUnion defines as:
//
//   union switch (OperationType type)
//        {
//        case CREATE_ACCOUNT:
//            CreateAccountResult createAccountResult;
//        case PAYMENT:
//            PaymentResult paymentResult;
//        case SET_OPTIONS:
//            SetOptionsResult setOptionsResult;
//    	case CREATE_ISSUANCE_REQUEST:
//    		CreateIssuanceRequestResult createIssuanceRequestResult;
//        case SET_FEES:
//            SetFeesResult setFeesResult;
//    	case MANAGE_ACCOUNT:
//    		ManageAccountResult manageAccountResult;
//        case CREATE_WITHDRAWAL_REQUEST:
//    		CreateWithdrawalRequestResult createWithdrawalRequestResult;
//        case MANAGE_BALANCE:
//            ManageBalanceResult manageBalanceResult;
//        case MANAGE_ASSET:
//            ManageAssetResult manageAssetResult;
//        case CREATE_PREISSUANCE_REQUEST:
//            CreatePreIssuanceRequestResult createPreIssuanceRequestResult;
//        case MANAGE_LIMITS:
//            ManageLimitsResult manageLimitsResult;
//        case DIRECT_DEBIT:
//            DirectDebitResult directDebitResult;
//    	case MANAGE_ASSET_PAIR:
//    		ManageAssetPairResult manageAssetPairResult;
//    	case MANAGE_OFFER:
//    		ManageOfferResult manageOfferResult;
//    	case MANAGE_INVOICE_REQUEST:
//    		ManageInvoiceRequestResult manageInvoiceRequestResult;
//    	case REVIEW_REQUEST:
//    		ReviewRequestResult reviewRequestResult;
//    	case CREATE_SALE_REQUEST:
//    		CreateSaleCreationRequestResult createSaleCreationRequestResult;
//    	case CHECK_SALE_STATE:
//    		CheckSaleStateResult checkSaleStateResult;
//    	case CREATE_AML_ALERT:
//    	    CreateAMLAlertRequestResult createAMLAlertRequestResult;
//    	case MANAGE_KEY_VALUE:
//    	    ManageKeyValueResult manageKeyValueResult;
//    	case CREATE_KYC_REQUEST:
//    	    CreateUpdateKYCRequestResult createUpdateKYCRequestResult;
//        case MANAGE_EXTERNAL_SYSTEM_ACCOUNT_ID_POOL_ENTRY:
//            ManageExternalSystemAccountIdPoolEntryResult manageExternalSystemAccountIdPoolEntryResult;
//        case BIND_EXTERNAL_SYSTEM_ACCOUNT_ID:
//            BindExternalSystemAccountIdResult bindExternalSystemAccountIdResult;
//        case PAYMENT_V2:
//            PaymentV2Result paymentV2Result;
//        case MANAGE_SALE:
//            ManageSaleResult manageSaleResult;
//        case CREATE_MANAGE_LIMITS_REQUEST:
//            CreateManageLimitsRequestResult createManageLimitsRequestResult;
//        case MANAGE_CONTRACT_REQUEST:
//            ManageContractRequestResult manageContractRequestResult;
//        case MANAGE_CONTRACT:
//            ManageContractResult manageContractResult;
//        }
//
type OperationResultTr struct {
	Type                                         OperationType                                 `json:"type,omitempty"`
	CreateAccountResult                          *CreateAccountResult                          `json:"createAccountResult,omitempty"`
	PaymentResult                                *PaymentResult                                `json:"paymentResult,omitempty"`
	SetOptionsResult                             *SetOptionsResult                             `json:"setOptionsResult,omitempty"`
	CreateIssuanceRequestResult                  *CreateIssuanceRequestResult                  `json:"createIssuanceRequestResult,omitempty"`
	SetFeesResult                                *SetFeesResult                                `json:"setFeesResult,omitempty"`
	ManageAccountResult                          *ManageAccountResult                          `json:"manageAccountResult,omitempty"`
	CreateWithdrawalRequestResult                *CreateWithdrawalRequestResult                `json:"createWithdrawalRequestResult,omitempty"`
	ManageBalanceResult                          *ManageBalanceResult                          `json:"manageBalanceResult,omitempty"`
	ManageAssetResult                            *ManageAssetResult                            `json:"manageAssetResult,omitempty"`
	CreatePreIssuanceRequestResult               *CreatePreIssuanceRequestResult               `json:"createPreIssuanceRequestResult,omitempty"`
	ManageLimitsResult                           *ManageLimitsResult                           `json:"manageLimitsResult,omitempty"`
	DirectDebitResult                            *DirectDebitResult                            `json:"directDebitResult,omitempty"`
	ManageAssetPairResult                        *ManageAssetPairResult                        `json:"manageAssetPairResult,omitempty"`
	ManageOfferResult                            *ManageOfferResult                            `json:"manageOfferResult,omitempty"`
	ManageInvoiceRequestResult                   *ManageInvoiceRequestResult                   `json:"manageInvoiceRequestResult,omitempty"`
	ReviewRequestResult                          *ReviewRequestResult                          `json:"reviewRequestResult,omitempty"`
	CreateSaleCreationRequestResult              *CreateSaleCreationRequestResult              `json:"createSaleCreationRequestResult,omitempty"`
	CheckSaleStateResult                         *CheckSaleStateResult                         `json:"checkSaleStateResult,omitempty"`
	CreateAmlAlertRequestResult                  *CreateAmlAlertRequestResult                  `json:"createAMLAlertRequestResult,omitempty"`
	ManageKeyValueResult                         *ManageKeyValueResult                         `json:"manageKeyValueResult,omitempty"`
	CreateUpdateKycRequestResult                 *CreateUpdateKycRequestResult                 `json:"createUpdateKYCRequestResult,omitempty"`
	ManageExternalSystemAccountIdPoolEntryResult *ManageExternalSystemAccountIdPoolEntryResult `json:"manageExternalSystemAccountIdPoolEntryResult,omitempty"`
	BindExternalSystemAccountIdResult            *BindExternalSystemAccountIdResult            `json:"bindExternalSystemAccountIdResult,omitempty"`
	PaymentV2Result                              *PaymentV2Result                              `json:"paymentV2Result,omitempty"`
	ManageSaleResult                             *ManageSaleResult                             `json:"manageSaleResult,omitempty"`
	CreateManageLimitsRequestResult              *CreateManageLimitsRequestResult              `json:"createManageLimitsRequestResult,omitempty"`
	ManageContractRequestResult                  *ManageContractRequestResult                  `json:"manageContractRequestResult,omitempty"`
	ManageContractResult                         *ManageContractResult                         `json:"manageContractResult,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u OperationResultTr) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of OperationResultTr
func (u OperationResultTr) ArmForSwitch(sw int32) (string, bool) {
	switch OperationType(sw) {
	case OperationTypeCreateAccount:
		return "CreateAccountResult", true
	case OperationTypePayment:
		return "PaymentResult", true
	case OperationTypeSetOptions:
		return "SetOptionsResult", true
	case OperationTypeCreateIssuanceRequest:
		return "CreateIssuanceRequestResult", true
	case OperationTypeSetFees:
		return "SetFeesResult", true
	case OperationTypeManageAccount:
		return "ManageAccountResult", true
	case OperationTypeCreateWithdrawalRequest:
		return "CreateWithdrawalRequestResult", true
	case OperationTypeManageBalance:
		return "ManageBalanceResult", true
	case OperationTypeManageAsset:
		return "ManageAssetResult", true
	case OperationTypeCreatePreissuanceRequest:
		return "CreatePreIssuanceRequestResult", true
	case OperationTypeManageLimits:
		return "ManageLimitsResult", true
	case OperationTypeDirectDebit:
		return "DirectDebitResult", true
	case OperationTypeManageAssetPair:
		return "ManageAssetPairResult", true
	case OperationTypeManageOffer:
		return "ManageOfferResult", true
	case OperationTypeManageInvoiceRequest:
		return "ManageInvoiceRequestResult", true
	case OperationTypeReviewRequest:
		return "ReviewRequestResult", true
	case OperationTypeCreateSaleRequest:
		return "CreateSaleCreationRequestResult", true
	case OperationTypeCheckSaleState:
		return "CheckSaleStateResult", true
	case OperationTypeCreateAmlAlert:
		return "CreateAmlAlertRequestResult", true
	case OperationTypeManageKeyValue:
		return "ManageKeyValueResult", true
	case OperationTypeCreateKycRequest:
		return "CreateUpdateKycRequestResult", true
	case OperationTypeManageExternalSystemAccountIdPoolEntry:
		return "ManageExternalSystemAccountIdPoolEntryResult", true
	case OperationTypeBindExternalSystemAccountId:
		return "BindExternalSystemAccountIdResult", true
	case OperationTypePaymentV2:
		return "PaymentV2Result", true
	case OperationTypeManageSale:
		return "ManageSaleResult", true
	case OperationTypeCreateManageLimitsRequest:
		return "CreateManageLimitsRequestResult", true
	case OperationTypeManageContractRequest:
		return "ManageContractRequestResult", true
	case OperationTypeManageContract:
		return "ManageContractResult", true
	}
	return "-", false
}

// NewOperationResultTr creates a new  OperationResultTr.
func NewOperationResultTr(aType OperationType, value interface{}) (result OperationResultTr, err error) {
	result.Type = aType
	switch OperationType(aType) {
	case OperationTypeCreateAccount:
		tv, ok := value.(CreateAccountResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateAccountResult")
			return
		}
		result.CreateAccountResult = &tv
	case OperationTypePayment:
		tv, ok := value.(PaymentResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be PaymentResult")
			return
		}
		result.PaymentResult = &tv
	case OperationTypeSetOptions:
		tv, ok := value.(SetOptionsResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be SetOptionsResult")
			return
		}
		result.SetOptionsResult = &tv
	case OperationTypeCreateIssuanceRequest:
		tv, ok := value.(CreateIssuanceRequestResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateIssuanceRequestResult")
			return
		}
		result.CreateIssuanceRequestResult = &tv
	case OperationTypeSetFees:
		tv, ok := value.(SetFeesResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be SetFeesResult")
			return
		}
		result.SetFeesResult = &tv
	case OperationTypeManageAccount:
		tv, ok := value.(ManageAccountResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageAccountResult")
			return
		}
		result.ManageAccountResult = &tv
	case OperationTypeCreateWithdrawalRequest:
		tv, ok := value.(CreateWithdrawalRequestResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateWithdrawalRequestResult")
			return
		}
		result.CreateWithdrawalRequestResult = &tv
	case OperationTypeManageBalance:
		tv, ok := value.(ManageBalanceResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageBalanceResult")
			return
		}
		result.ManageBalanceResult = &tv
	case OperationTypeManageAsset:
		tv, ok := value.(ManageAssetResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageAssetResult")
			return
		}
		result.ManageAssetResult = &tv
	case OperationTypeCreatePreissuanceRequest:
		tv, ok := value.(CreatePreIssuanceRequestResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreatePreIssuanceRequestResult")
			return
		}
		result.CreatePreIssuanceRequestResult = &tv
	case OperationTypeManageLimits:
		tv, ok := value.(ManageLimitsResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageLimitsResult")
			return
		}
		result.ManageLimitsResult = &tv
	case OperationTypeDirectDebit:
		tv, ok := value.(DirectDebitResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be DirectDebitResult")
			return
		}
		result.DirectDebitResult = &tv
	case OperationTypeManageAssetPair:
		tv, ok := value.(ManageAssetPairResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageAssetPairResult")
			return
		}
		result.ManageAssetPairResult = &tv
	case OperationTypeManageOffer:
		tv, ok := value.(ManageOfferResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageOfferResult")
			return
		}
		result.ManageOfferResult = &tv
	case OperationTypeManageInvoiceRequest:
		tv, ok := value.(ManageInvoiceRequestResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageInvoiceRequestResult")
			return
		}
		result.ManageInvoiceRequestResult = &tv
	case OperationTypeReviewRequest:
		tv, ok := value.(ReviewRequestResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be ReviewRequestResult")
			return
		}
		result.ReviewRequestResult = &tv
	case OperationTypeCreateSaleRequest:
		tv, ok := value.(CreateSaleCreationRequestResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateSaleCreationRequestResult")
			return
		}
		result.CreateSaleCreationRequestResult = &tv
	case OperationTypeCheckSaleState:
		tv, ok := value.(CheckSaleStateResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be CheckSaleStateResult")
			return
		}
		result.CheckSaleStateResult = &tv
	case OperationTypeCreateAmlAlert:
		tv, ok := value.(CreateAmlAlertRequestResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateAmlAlertRequestResult")
			return
		}
		result.CreateAmlAlertRequestResult = &tv
	case OperationTypeManageKeyValue:
		tv, ok := value.(ManageKeyValueResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageKeyValueResult")
			return
		}
		result.ManageKeyValueResult = &tv
	case OperationTypeCreateKycRequest:
		tv, ok := value.(CreateUpdateKycRequestResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateUpdateKycRequestResult")
			return
		}
		result.CreateUpdateKycRequestResult = &tv
	case OperationTypeManageExternalSystemAccountIdPoolEntry:
		tv, ok := value.(ManageExternalSystemAccountIdPoolEntryResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageExternalSystemAccountIdPoolEntryResult")
			return
		}
		result.ManageExternalSystemAccountIdPoolEntryResult = &tv
	case OperationTypeBindExternalSystemAccountId:
		tv, ok := value.(BindExternalSystemAccountIdResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be BindExternalSystemAccountIdResult")
			return
		}
		result.BindExternalSystemAccountIdResult = &tv
	case OperationTypePaymentV2:
		tv, ok := value.(PaymentV2Result)
		if !ok {
			err = fmt.Errorf("invalid value, must be PaymentV2Result")
			return
		}
		result.PaymentV2Result = &tv
	case OperationTypeManageSale:
		tv, ok := value.(ManageSaleResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageSaleResult")
			return
		}
		result.ManageSaleResult = &tv
	case OperationTypeCreateManageLimitsRequest:
		tv, ok := value.(CreateManageLimitsRequestResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be CreateManageLimitsRequestResult")
			return
		}
		result.CreateManageLimitsRequestResult = &tv
	case OperationTypeManageContractRequest:
		tv, ok := value.(ManageContractRequestResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageContractRequestResult")
			return
		}
		result.ManageContractRequestResult = &tv
	case OperationTypeManageContract:
		tv, ok := value.(ManageContractResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageContractResult")
			return
		}
		result.ManageContractResult = &tv
	}
	return
}

// MustCreateAccountResult retrieves the CreateAccountResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustCreateAccountResult() CreateAccountResult {
	val, ok := u.GetCreateAccountResult()

	if !ok {
		panic("arm CreateAccountResult is not set")
	}

	return val
}

// GetCreateAccountResult retrieves the CreateAccountResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetCreateAccountResult() (result CreateAccountResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CreateAccountResult" {
		result = *u.CreateAccountResult
		ok = true
	}

	return
}

// MustPaymentResult retrieves the PaymentResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustPaymentResult() PaymentResult {
	val, ok := u.GetPaymentResult()

	if !ok {
		panic("arm PaymentResult is not set")
	}

	return val
}

// GetPaymentResult retrieves the PaymentResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetPaymentResult() (result PaymentResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "PaymentResult" {
		result = *u.PaymentResult
		ok = true
	}

	return
}

// MustSetOptionsResult retrieves the SetOptionsResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustSetOptionsResult() SetOptionsResult {
	val, ok := u.GetSetOptionsResult()

	if !ok {
		panic("arm SetOptionsResult is not set")
	}

	return val
}

// GetSetOptionsResult retrieves the SetOptionsResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetSetOptionsResult() (result SetOptionsResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "SetOptionsResult" {
		result = *u.SetOptionsResult
		ok = true
	}

	return
}

// MustCreateIssuanceRequestResult retrieves the CreateIssuanceRequestResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustCreateIssuanceRequestResult() CreateIssuanceRequestResult {
	val, ok := u.GetCreateIssuanceRequestResult()

	if !ok {
		panic("arm CreateIssuanceRequestResult is not set")
	}

	return val
}

// GetCreateIssuanceRequestResult retrieves the CreateIssuanceRequestResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetCreateIssuanceRequestResult() (result CreateIssuanceRequestResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CreateIssuanceRequestResult" {
		result = *u.CreateIssuanceRequestResult
		ok = true
	}

	return
}

// MustSetFeesResult retrieves the SetFeesResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustSetFeesResult() SetFeesResult {
	val, ok := u.GetSetFeesResult()

	if !ok {
		panic("arm SetFeesResult is not set")
	}

	return val
}

// GetSetFeesResult retrieves the SetFeesResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetSetFeesResult() (result SetFeesResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "SetFeesResult" {
		result = *u.SetFeesResult
		ok = true
	}

	return
}

// MustManageAccountResult retrieves the ManageAccountResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustManageAccountResult() ManageAccountResult {
	val, ok := u.GetManageAccountResult()

	if !ok {
		panic("arm ManageAccountResult is not set")
	}

	return val
}

// GetManageAccountResult retrieves the ManageAccountResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetManageAccountResult() (result ManageAccountResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ManageAccountResult" {
		result = *u.ManageAccountResult
		ok = true
	}

	return
}

// MustCreateWithdrawalRequestResult retrieves the CreateWithdrawalRequestResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustCreateWithdrawalRequestResult() CreateWithdrawalRequestResult {
	val, ok := u.GetCreateWithdrawalRequestResult()

	if !ok {
		panic("arm CreateWithdrawalRequestResult is not set")
	}

	return val
}

// GetCreateWithdrawalRequestResult retrieves the CreateWithdrawalRequestResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetCreateWithdrawalRequestResult() (result CreateWithdrawalRequestResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CreateWithdrawalRequestResult" {
		result = *u.CreateWithdrawalRequestResult
		ok = true
	}

	return
}

// MustManageBalanceResult retrieves the ManageBalanceResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustManageBalanceResult() ManageBalanceResult {
	val, ok := u.GetManageBalanceResult()

	if !ok {
		panic("arm ManageBalanceResult is not set")
	}

	return val
}

// GetManageBalanceResult retrieves the ManageBalanceResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetManageBalanceResult() (result ManageBalanceResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ManageBalanceResult" {
		result = *u.ManageBalanceResult
		ok = true
	}

	return
}

// MustManageAssetResult retrieves the ManageAssetResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustManageAssetResult() ManageAssetResult {
	val, ok := u.GetManageAssetResult()

	if !ok {
		panic("arm ManageAssetResult is not set")
	}

	return val
}

// GetManageAssetResult retrieves the ManageAssetResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetManageAssetResult() (result ManageAssetResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ManageAssetResult" {
		result = *u.ManageAssetResult
		ok = true
	}

	return
}

// MustCreatePreIssuanceRequestResult retrieves the CreatePreIssuanceRequestResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustCreatePreIssuanceRequestResult() CreatePreIssuanceRequestResult {
	val, ok := u.GetCreatePreIssuanceRequestResult()

	if !ok {
		panic("arm CreatePreIssuanceRequestResult is not set")
	}

	return val
}

// GetCreatePreIssuanceRequestResult retrieves the CreatePreIssuanceRequestResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetCreatePreIssuanceRequestResult() (result CreatePreIssuanceRequestResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CreatePreIssuanceRequestResult" {
		result = *u.CreatePreIssuanceRequestResult
		ok = true
	}

	return
}

// MustManageLimitsResult retrieves the ManageLimitsResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustManageLimitsResult() ManageLimitsResult {
	val, ok := u.GetManageLimitsResult()

	if !ok {
		panic("arm ManageLimitsResult is not set")
	}

	return val
}

// GetManageLimitsResult retrieves the ManageLimitsResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetManageLimitsResult() (result ManageLimitsResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ManageLimitsResult" {
		result = *u.ManageLimitsResult
		ok = true
	}

	return
}

// MustDirectDebitResult retrieves the DirectDebitResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustDirectDebitResult() DirectDebitResult {
	val, ok := u.GetDirectDebitResult()

	if !ok {
		panic("arm DirectDebitResult is not set")
	}

	return val
}

// GetDirectDebitResult retrieves the DirectDebitResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetDirectDebitResult() (result DirectDebitResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "DirectDebitResult" {
		result = *u.DirectDebitResult
		ok = true
	}

	return
}

// MustManageAssetPairResult retrieves the ManageAssetPairResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustManageAssetPairResult() ManageAssetPairResult {
	val, ok := u.GetManageAssetPairResult()

	if !ok {
		panic("arm ManageAssetPairResult is not set")
	}

	return val
}

// GetManageAssetPairResult retrieves the ManageAssetPairResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetManageAssetPairResult() (result ManageAssetPairResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ManageAssetPairResult" {
		result = *u.ManageAssetPairResult
		ok = true
	}

	return
}

// MustManageOfferResult retrieves the ManageOfferResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustManageOfferResult() ManageOfferResult {
	val, ok := u.GetManageOfferResult()

	if !ok {
		panic("arm ManageOfferResult is not set")
	}

	return val
}

// GetManageOfferResult retrieves the ManageOfferResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetManageOfferResult() (result ManageOfferResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ManageOfferResult" {
		result = *u.ManageOfferResult
		ok = true
	}

	return
}

// MustManageInvoiceRequestResult retrieves the ManageInvoiceRequestResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustManageInvoiceRequestResult() ManageInvoiceRequestResult {
	val, ok := u.GetManageInvoiceRequestResult()

	if !ok {
		panic("arm ManageInvoiceRequestResult is not set")
	}

	return val
}

// GetManageInvoiceRequestResult retrieves the ManageInvoiceRequestResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetManageInvoiceRequestResult() (result ManageInvoiceRequestResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ManageInvoiceRequestResult" {
		result = *u.ManageInvoiceRequestResult
		ok = true
	}

	return
}

// MustReviewRequestResult retrieves the ReviewRequestResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustReviewRequestResult() ReviewRequestResult {
	val, ok := u.GetReviewRequestResult()

	if !ok {
		panic("arm ReviewRequestResult is not set")
	}

	return val
}

// GetReviewRequestResult retrieves the ReviewRequestResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetReviewRequestResult() (result ReviewRequestResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ReviewRequestResult" {
		result = *u.ReviewRequestResult
		ok = true
	}

	return
}

// MustCreateSaleCreationRequestResult retrieves the CreateSaleCreationRequestResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustCreateSaleCreationRequestResult() CreateSaleCreationRequestResult {
	val, ok := u.GetCreateSaleCreationRequestResult()

	if !ok {
		panic("arm CreateSaleCreationRequestResult is not set")
	}

	return val
}

// GetCreateSaleCreationRequestResult retrieves the CreateSaleCreationRequestResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetCreateSaleCreationRequestResult() (result CreateSaleCreationRequestResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CreateSaleCreationRequestResult" {
		result = *u.CreateSaleCreationRequestResult
		ok = true
	}

	return
}

// MustCheckSaleStateResult retrieves the CheckSaleStateResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustCheckSaleStateResult() CheckSaleStateResult {
	val, ok := u.GetCheckSaleStateResult()

	if !ok {
		panic("arm CheckSaleStateResult is not set")
	}

	return val
}

// GetCheckSaleStateResult retrieves the CheckSaleStateResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetCheckSaleStateResult() (result CheckSaleStateResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CheckSaleStateResult" {
		result = *u.CheckSaleStateResult
		ok = true
	}

	return
}

// MustCreateAmlAlertRequestResult retrieves the CreateAmlAlertRequestResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustCreateAmlAlertRequestResult() CreateAmlAlertRequestResult {
	val, ok := u.GetCreateAmlAlertRequestResult()

	if !ok {
		panic("arm CreateAmlAlertRequestResult is not set")
	}

	return val
}

// GetCreateAmlAlertRequestResult retrieves the CreateAmlAlertRequestResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetCreateAmlAlertRequestResult() (result CreateAmlAlertRequestResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CreateAmlAlertRequestResult" {
		result = *u.CreateAmlAlertRequestResult
		ok = true
	}

	return
}

// MustManageKeyValueResult retrieves the ManageKeyValueResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustManageKeyValueResult() ManageKeyValueResult {
	val, ok := u.GetManageKeyValueResult()

	if !ok {
		panic("arm ManageKeyValueResult is not set")
	}

	return val
}

// GetManageKeyValueResult retrieves the ManageKeyValueResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetManageKeyValueResult() (result ManageKeyValueResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ManageKeyValueResult" {
		result = *u.ManageKeyValueResult
		ok = true
	}

	return
}

// MustCreateUpdateKycRequestResult retrieves the CreateUpdateKycRequestResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustCreateUpdateKycRequestResult() CreateUpdateKycRequestResult {
	val, ok := u.GetCreateUpdateKycRequestResult()

	if !ok {
		panic("arm CreateUpdateKycRequestResult is not set")
	}

	return val
}

// GetCreateUpdateKycRequestResult retrieves the CreateUpdateKycRequestResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetCreateUpdateKycRequestResult() (result CreateUpdateKycRequestResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CreateUpdateKycRequestResult" {
		result = *u.CreateUpdateKycRequestResult
		ok = true
	}

	return
}

// MustManageExternalSystemAccountIdPoolEntryResult retrieves the ManageExternalSystemAccountIdPoolEntryResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustManageExternalSystemAccountIdPoolEntryResult() ManageExternalSystemAccountIdPoolEntryResult {
	val, ok := u.GetManageExternalSystemAccountIdPoolEntryResult()

	if !ok {
		panic("arm ManageExternalSystemAccountIdPoolEntryResult is not set")
	}

	return val
}

// GetManageExternalSystemAccountIdPoolEntryResult retrieves the ManageExternalSystemAccountIdPoolEntryResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetManageExternalSystemAccountIdPoolEntryResult() (result ManageExternalSystemAccountIdPoolEntryResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ManageExternalSystemAccountIdPoolEntryResult" {
		result = *u.ManageExternalSystemAccountIdPoolEntryResult
		ok = true
	}

	return
}

// MustBindExternalSystemAccountIdResult retrieves the BindExternalSystemAccountIdResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustBindExternalSystemAccountIdResult() BindExternalSystemAccountIdResult {
	val, ok := u.GetBindExternalSystemAccountIdResult()

	if !ok {
		panic("arm BindExternalSystemAccountIdResult is not set")
	}

	return val
}

// GetBindExternalSystemAccountIdResult retrieves the BindExternalSystemAccountIdResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetBindExternalSystemAccountIdResult() (result BindExternalSystemAccountIdResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "BindExternalSystemAccountIdResult" {
		result = *u.BindExternalSystemAccountIdResult
		ok = true
	}

	return
}

// MustPaymentV2Result retrieves the PaymentV2Result value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustPaymentV2Result() PaymentV2Result {
	val, ok := u.GetPaymentV2Result()

	if !ok {
		panic("arm PaymentV2Result is not set")
	}

	return val
}

// GetPaymentV2Result retrieves the PaymentV2Result value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetPaymentV2Result() (result PaymentV2Result, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "PaymentV2Result" {
		result = *u.PaymentV2Result
		ok = true
	}

	return
}

// MustManageSaleResult retrieves the ManageSaleResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustManageSaleResult() ManageSaleResult {
	val, ok := u.GetManageSaleResult()

	if !ok {
		panic("arm ManageSaleResult is not set")
	}

	return val
}

// GetManageSaleResult retrieves the ManageSaleResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetManageSaleResult() (result ManageSaleResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ManageSaleResult" {
		result = *u.ManageSaleResult
		ok = true
	}

	return
}

// MustCreateManageLimitsRequestResult retrieves the CreateManageLimitsRequestResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustCreateManageLimitsRequestResult() CreateManageLimitsRequestResult {
	val, ok := u.GetCreateManageLimitsRequestResult()

	if !ok {
		panic("arm CreateManageLimitsRequestResult is not set")
	}

	return val
}

// GetCreateManageLimitsRequestResult retrieves the CreateManageLimitsRequestResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetCreateManageLimitsRequestResult() (result CreateManageLimitsRequestResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "CreateManageLimitsRequestResult" {
		result = *u.CreateManageLimitsRequestResult
		ok = true
	}

	return
}

// MustManageContractRequestResult retrieves the ManageContractRequestResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustManageContractRequestResult() ManageContractRequestResult {
	val, ok := u.GetManageContractRequestResult()

	if !ok {
		panic("arm ManageContractRequestResult is not set")
	}

	return val
}

// GetManageContractRequestResult retrieves the ManageContractRequestResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetManageContractRequestResult() (result ManageContractRequestResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ManageContractRequestResult" {
		result = *u.ManageContractRequestResult
		ok = true
	}

	return
}

// MustManageContractResult retrieves the ManageContractResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustManageContractResult() ManageContractResult {
	val, ok := u.GetManageContractResult()

	if !ok {
		panic("arm ManageContractResult is not set")
	}

	return val
}

// GetManageContractResult retrieves the ManageContractResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetManageContractResult() (result ManageContractResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ManageContractResult" {
		result = *u.ManageContractResult
		ok = true
	}

	return
}

// OperationResult is an XDR Union defines as:
//
//   union OperationResult switch (OperationResultCode code)
//    {
//    case opINNER:
//        union switch (OperationType type)
//        {
//        case CREATE_ACCOUNT:
//            CreateAccountResult createAccountResult;
//        case PAYMENT:
//            PaymentResult paymentResult;
//        case SET_OPTIONS:
//            SetOptionsResult setOptionsResult;
//    	case CREATE_ISSUANCE_REQUEST:
//    		CreateIssuanceRequestResult createIssuanceRequestResult;
//        case SET_FEES:
//            SetFeesResult setFeesResult;
//    	case MANAGE_ACCOUNT:
//    		ManageAccountResult manageAccountResult;
//        case CREATE_WITHDRAWAL_REQUEST:
//    		CreateWithdrawalRequestResult createWithdrawalRequestResult;
//        case MANAGE_BALANCE:
//            ManageBalanceResult manageBalanceResult;
//        case MANAGE_ASSET:
//            ManageAssetResult manageAssetResult;
//        case CREATE_PREISSUANCE_REQUEST:
//            CreatePreIssuanceRequestResult createPreIssuanceRequestResult;
//        case MANAGE_LIMITS:
//            ManageLimitsResult manageLimitsResult;
//        case DIRECT_DEBIT:
//            DirectDebitResult directDebitResult;
//    	case MANAGE_ASSET_PAIR:
//    		ManageAssetPairResult manageAssetPairResult;
//    	case MANAGE_OFFER:
//    		ManageOfferResult manageOfferResult;
//    	case MANAGE_INVOICE_REQUEST:
//    		ManageInvoiceRequestResult manageInvoiceRequestResult;
//    	case REVIEW_REQUEST:
//    		ReviewRequestResult reviewRequestResult;
//    	case CREATE_SALE_REQUEST:
//    		CreateSaleCreationRequestResult createSaleCreationRequestResult;
//    	case CHECK_SALE_STATE:
//    		CheckSaleStateResult checkSaleStateResult;
//    	case CREATE_AML_ALERT:
//    	    CreateAMLAlertRequestResult createAMLAlertRequestResult;
//    	case MANAGE_KEY_VALUE:
//    	    ManageKeyValueResult manageKeyValueResult;
//    	case CREATE_KYC_REQUEST:
//    	    CreateUpdateKYCRequestResult createUpdateKYCRequestResult;
//        case MANAGE_EXTERNAL_SYSTEM_ACCOUNT_ID_POOL_ENTRY:
//            ManageExternalSystemAccountIdPoolEntryResult manageExternalSystemAccountIdPoolEntryResult;
//        case BIND_EXTERNAL_SYSTEM_ACCOUNT_ID:
//            BindExternalSystemAccountIdResult bindExternalSystemAccountIdResult;
//        case PAYMENT_V2:
//            PaymentV2Result paymentV2Result;
//        case MANAGE_SALE:
//            ManageSaleResult manageSaleResult;
//        case CREATE_MANAGE_LIMITS_REQUEST:
//            CreateManageLimitsRequestResult createManageLimitsRequestResult;
//        case MANAGE_CONTRACT_REQUEST:
//            ManageContractRequestResult manageContractRequestResult;
//        case MANAGE_CONTRACT:
//            ManageContractResult manageContractResult;
//        }
//        tr;
//    default:
//        void;
//    };
//
type OperationResult struct {
	Code OperationResultCode `json:"code,omitempty"`
	Tr   *OperationResultTr  `json:"tr,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u OperationResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of OperationResult
func (u OperationResult) ArmForSwitch(sw int32) (string, bool) {
	switch OperationResultCode(sw) {
	case OperationResultCodeOpInner:
		return "Tr", true
	default:
		return "", true
	}
}

// NewOperationResult creates a new  OperationResult.
func NewOperationResult(code OperationResultCode, value interface{}) (result OperationResult, err error) {
	result.Code = code
	switch OperationResultCode(code) {
	case OperationResultCodeOpInner:
		tv, ok := value.(OperationResultTr)
		if !ok {
			err = fmt.Errorf("invalid value, must be OperationResultTr")
			return
		}
		result.Tr = &tv
	default:
		// void
	}
	return
}

// MustTr retrieves the Tr value from the union,
// panicing if the value is not set.
func (u OperationResult) MustTr() OperationResultTr {
	val, ok := u.GetTr()

	if !ok {
		panic("arm Tr is not set")
	}

	return val
}

// GetTr retrieves the Tr value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResult) GetTr() (result OperationResultTr, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Tr" {
		result = *u.Tr
		ok = true
	}

	return
}

// TransactionResultCode is an XDR Enum defines as:
//
//   enum TransactionResultCode
//    {
//        txSUCCESS = 0, // all operations succeeded
//
//        txFAILED = -1, // one of the operations failed (none were applied)
//
//        txTOO_EARLY = -2,         // ledger closeTime before minTime
//        txTOO_LATE = -3,          // ledger closeTime after maxTime
//        txMISSING_OPERATION = -4, // no operation was specified
//
//        txBAD_AUTH = -5,             // too few valid signatures / wrong network
//        txNO_ACCOUNT = -6,           // source account not found
//        txBAD_AUTH_EXTRA = -7,      // unused signatures attached to transaction
//        txINTERNAL_ERROR = -8,      // an unknown error occured
//    	txACCOUNT_BLOCKED = -9,     // account is blocked and cannot be source of tx
//        txDUPLICATION = -10         // if timing is stored
//    };
//
type TransactionResultCode int32

const (
	TransactionResultCodeTxSuccess          TransactionResultCode = 0
	TransactionResultCodeTxFailed           TransactionResultCode = -1
	TransactionResultCodeTxTooEarly         TransactionResultCode = -2
	TransactionResultCodeTxTooLate          TransactionResultCode = -3
	TransactionResultCodeTxMissingOperation TransactionResultCode = -4
	TransactionResultCodeTxBadAuth          TransactionResultCode = -5
	TransactionResultCodeTxNoAccount        TransactionResultCode = -6
	TransactionResultCodeTxBadAuthExtra     TransactionResultCode = -7
	TransactionResultCodeTxInternalError    TransactionResultCode = -8
	TransactionResultCodeTxAccountBlocked   TransactionResultCode = -9
	TransactionResultCodeTxDuplication      TransactionResultCode = -10
)

var TransactionResultCodeAll = []TransactionResultCode{
	TransactionResultCodeTxSuccess,
	TransactionResultCodeTxFailed,
	TransactionResultCodeTxTooEarly,
	TransactionResultCodeTxTooLate,
	TransactionResultCodeTxMissingOperation,
	TransactionResultCodeTxBadAuth,
	TransactionResultCodeTxNoAccount,
	TransactionResultCodeTxBadAuthExtra,
	TransactionResultCodeTxInternalError,
	TransactionResultCodeTxAccountBlocked,
	TransactionResultCodeTxDuplication,
}

var transactionResultCodeMap = map[int32]string{
	0:   "TransactionResultCodeTxSuccess",
	-1:  "TransactionResultCodeTxFailed",
	-2:  "TransactionResultCodeTxTooEarly",
	-3:  "TransactionResultCodeTxTooLate",
	-4:  "TransactionResultCodeTxMissingOperation",
	-5:  "TransactionResultCodeTxBadAuth",
	-6:  "TransactionResultCodeTxNoAccount",
	-7:  "TransactionResultCodeTxBadAuthExtra",
	-8:  "TransactionResultCodeTxInternalError",
	-9:  "TransactionResultCodeTxAccountBlocked",
	-10: "TransactionResultCodeTxDuplication",
}

var transactionResultCodeShortMap = map[int32]string{
	0:   "tx_success",
	-1:  "tx_failed",
	-2:  "tx_too_early",
	-3:  "tx_too_late",
	-4:  "tx_missing_operation",
	-5:  "tx_bad_auth",
	-6:  "tx_no_account",
	-7:  "tx_bad_auth_extra",
	-8:  "tx_internal_error",
	-9:  "tx_account_blocked",
	-10: "tx_duplication",
}

var transactionResultCodeRevMap = map[string]int32{
	"TransactionResultCodeTxSuccess":          0,
	"TransactionResultCodeTxFailed":           -1,
	"TransactionResultCodeTxTooEarly":         -2,
	"TransactionResultCodeTxTooLate":          -3,
	"TransactionResultCodeTxMissingOperation": -4,
	"TransactionResultCodeTxBadAuth":          -5,
	"TransactionResultCodeTxNoAccount":        -6,
	"TransactionResultCodeTxBadAuthExtra":     -7,
	"TransactionResultCodeTxInternalError":    -8,
	"TransactionResultCodeTxAccountBlocked":   -9,
	"TransactionResultCodeTxDuplication":      -10,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for TransactionResultCode
func (e TransactionResultCode) ValidEnum(v int32) bool {
	_, ok := transactionResultCodeMap[v]
	return ok
}
func (e TransactionResultCode) isFlag() bool {
	for i := len(TransactionResultCodeAll) - 1; i >= 0; i-- {
		expected := TransactionResultCode(2) << uint64(len(TransactionResultCodeAll)-1) >> uint64(len(TransactionResultCodeAll)-i)
		if expected != TransactionResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e TransactionResultCode) String() string {
	name, _ := transactionResultCodeMap[int32(e)]
	return name
}

func (e TransactionResultCode) ShortString() string {
	name, _ := transactionResultCodeShortMap[int32(e)]
	return name
}

func (e TransactionResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range TransactionResultCodeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *TransactionResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = TransactionResultCode(t.Value)
	return nil
}

// TransactionResultResult is an XDR NestedUnion defines as:
//
//   union switch (TransactionResultCode code)
//        {
//        case txSUCCESS:
//        case txFAILED:
//            OperationResult results<>;
//        default:
//            void;
//        }
//
type TransactionResultResult struct {
	Code    TransactionResultCode `json:"code,omitempty"`
	Results *[]OperationResult    `json:"results,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u TransactionResultResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of TransactionResultResult
func (u TransactionResultResult) ArmForSwitch(sw int32) (string, bool) {
	switch TransactionResultCode(sw) {
	case TransactionResultCodeTxSuccess:
		return "Results", true
	case TransactionResultCodeTxFailed:
		return "Results", true
	default:
		return "", true
	}
}

// NewTransactionResultResult creates a new  TransactionResultResult.
func NewTransactionResultResult(code TransactionResultCode, value interface{}) (result TransactionResultResult, err error) {
	result.Code = code
	switch TransactionResultCode(code) {
	case TransactionResultCodeTxSuccess:
		tv, ok := value.([]OperationResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be []OperationResult")
			return
		}
		result.Results = &tv
	case TransactionResultCodeTxFailed:
		tv, ok := value.([]OperationResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be []OperationResult")
			return
		}
		result.Results = &tv
	default:
		// void
	}
	return
}

// MustResults retrieves the Results value from the union,
// panicing if the value is not set.
func (u TransactionResultResult) MustResults() []OperationResult {
	val, ok := u.GetResults()

	if !ok {
		panic("arm Results is not set")
	}

	return val
}

// GetResults retrieves the Results value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u TransactionResultResult) GetResults() (result []OperationResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Results" {
		result = *u.Results
		ok = true
	}

	return
}

// TransactionResultExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type TransactionResultExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u TransactionResultExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of TransactionResultExt
func (u TransactionResultExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewTransactionResultExt creates a new  TransactionResultExt.
func NewTransactionResultExt(v LedgerVersion, value interface{}) (result TransactionResultExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// TransactionResult is an XDR Struct defines as:
//
//   struct TransactionResult
//    {
//        int64 feeCharged; // actual fee charged for the transaction
//
//        union switch (TransactionResultCode code)
//        {
//        case txSUCCESS:
//        case txFAILED:
//            OperationResult results<>;
//        default:
//            void;
//        }
//        result;
//
//        // reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type TransactionResult struct {
	FeeCharged Int64                   `json:"feeCharged,omitempty"`
	Result     TransactionResultResult `json:"result,omitempty"`
	Ext        TransactionResultExt    `json:"ext,omitempty"`
}

// Hash is an XDR Typedef defines as:
//
//   typedef opaque Hash[32];
//
type Hash [32]byte

// Uint256 is an XDR Typedef defines as:
//
//   typedef opaque uint256[32];
//
type Uint256 [32]byte

// Uint32 is an XDR Typedef defines as:
//
//   typedef unsigned int uint32;
//
type Uint32 uint32

// Int32 is an XDR Typedef defines as:
//
//   typedef int int32;
//
type Int32 int32

// Uint64 is an XDR Typedef defines as:
//
//   typedef unsigned hyper uint64;
//
type Uint64 uint64

// Int64 is an XDR Typedef defines as:
//
//   typedef hyper int64;
//
type Int64 int64

// CryptoKeyType is an XDR Enum defines as:
//
//   enum CryptoKeyType
//    {
//        KEY_TYPE_ED25519 = 0
//    };
//
type CryptoKeyType int32

const (
	CryptoKeyTypeKeyTypeEd25519 CryptoKeyType = 0
)

var CryptoKeyTypeAll = []CryptoKeyType{
	CryptoKeyTypeKeyTypeEd25519,
}

var cryptoKeyTypeMap = map[int32]string{
	0: "CryptoKeyTypeKeyTypeEd25519",
}

var cryptoKeyTypeShortMap = map[int32]string{
	0: "key_type_ed25519",
}

var cryptoKeyTypeRevMap = map[string]int32{
	"CryptoKeyTypeKeyTypeEd25519": 0,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for CryptoKeyType
func (e CryptoKeyType) ValidEnum(v int32) bool {
	_, ok := cryptoKeyTypeMap[v]
	return ok
}
func (e CryptoKeyType) isFlag() bool {
	for i := len(CryptoKeyTypeAll) - 1; i >= 0; i-- {
		expected := CryptoKeyType(2) << uint64(len(CryptoKeyTypeAll)-1) >> uint64(len(CryptoKeyTypeAll)-i)
		if expected != CryptoKeyTypeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e CryptoKeyType) String() string {
	name, _ := cryptoKeyTypeMap[int32(e)]
	return name
}

func (e CryptoKeyType) ShortString() string {
	name, _ := cryptoKeyTypeShortMap[int32(e)]
	return name
}

func (e CryptoKeyType) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range CryptoKeyTypeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *CryptoKeyType) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = CryptoKeyType(t.Value)
	return nil
}

// PublicKeyType is an XDR Enum defines as:
//
//   enum PublicKeyType
//    {
//    	PUBLIC_KEY_TYPE_ED25519 = 0
//    };
//
type PublicKeyType int32

const (
	PublicKeyTypePublicKeyTypeEd25519 PublicKeyType = 0
)

var PublicKeyTypeAll = []PublicKeyType{
	PublicKeyTypePublicKeyTypeEd25519,
}

var publicKeyTypeMap = map[int32]string{
	0: "PublicKeyTypePublicKeyTypeEd25519",
}

var publicKeyTypeShortMap = map[int32]string{
	0: "public_key_type_ed25519",
}

var publicKeyTypeRevMap = map[string]int32{
	"PublicKeyTypePublicKeyTypeEd25519": 0,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for PublicKeyType
func (e PublicKeyType) ValidEnum(v int32) bool {
	_, ok := publicKeyTypeMap[v]
	return ok
}
func (e PublicKeyType) isFlag() bool {
	for i := len(PublicKeyTypeAll) - 1; i >= 0; i-- {
		expected := PublicKeyType(2) << uint64(len(PublicKeyTypeAll)-1) >> uint64(len(PublicKeyTypeAll)-i)
		if expected != PublicKeyTypeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e PublicKeyType) String() string {
	name, _ := publicKeyTypeMap[int32(e)]
	return name
}

func (e PublicKeyType) ShortString() string {
	name, _ := publicKeyTypeShortMap[int32(e)]
	return name
}

func (e PublicKeyType) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range PublicKeyTypeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *PublicKeyType) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = PublicKeyType(t.Value)
	return nil
}

// PublicKey is an XDR Union defines as:
//
//   union PublicKey switch (CryptoKeyType type)
//    {
//    case KEY_TYPE_ED25519:
//        uint256 ed25519;
//    };
//
type PublicKey struct {
	Type    CryptoKeyType `json:"type,omitempty"`
	Ed25519 *Uint256      `json:"ed25519,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u PublicKey) SwitchFieldName() string {
	return "Type"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of PublicKey
func (u PublicKey) ArmForSwitch(sw int32) (string, bool) {
	switch CryptoKeyType(sw) {
	case CryptoKeyTypeKeyTypeEd25519:
		return "Ed25519", true
	}
	return "-", false
}

// NewPublicKey creates a new  PublicKey.
func NewPublicKey(aType CryptoKeyType, value interface{}) (result PublicKey, err error) {
	result.Type = aType
	switch CryptoKeyType(aType) {
	case CryptoKeyTypeKeyTypeEd25519:
		tv, ok := value.(Uint256)
		if !ok {
			err = fmt.Errorf("invalid value, must be Uint256")
			return
		}
		result.Ed25519 = &tv
	}
	return
}

// MustEd25519 retrieves the Ed25519 value from the union,
// panicing if the value is not set.
func (u PublicKey) MustEd25519() Uint256 {
	val, ok := u.GetEd25519()

	if !ok {
		panic("arm Ed25519 is not set")
	}

	return val
}

// GetEd25519 retrieves the Ed25519 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u PublicKey) GetEd25519() (result Uint256, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Ed25519" {
		result = *u.Ed25519
		ok = true
	}

	return
}

// LedgerVersion is an XDR Enum defines as:
//
//   enum LedgerVersion {
//    	EMPTY_VERSION = 0,
//    	PASS_EXTERNAL_SYS_ACC_ID_IN_CREATE_ACC = 1,
//    	DETAILED_LEDGER_CHANGES = 2, // write more all ledger changes to transaction meta
//    	NEW_SIGNER_TYPES = 3, // use more comprehensive list of signer types
//    	TYPED_SALE = 4, // sales can have type
//    	UNIQUE_BALANCE_CREATION = 5, // allows to specify in manage balance that balance should not be created if one for such asset and account exists
//    	ASSET_PREISSUER_MIGRATION = 6,
//    	ASSET_PREISSUER_MIGRATED = 7,
//    	USE_KYC_LEVEL = 8,
//    	ERROR_ON_NON_ZERO_TASKS_TO_REMOVE_IN_REJECT_KYC = 9,
//    	ALLOW_ACCOUNT_MANAGER_TO_CHANGE_KYC = 10,
//    	CHANGE_ASSET_ISSUER_BAD_AUTH_EXTRA_FIXED = 11,
//    	AUTO_CREATE_COMMISSION_BALANCE_ON_TRANSFER = 12,
//        ALLOW_REJECT_REQUEST_OF_BLOCKED_REQUESTOR = 13,
//    	ASSET_UPDATE_CHECK_REFERENCE_EXISTS = 14,
//    	CROSS_ASSET_FEE = 15,
//    	USE_PAYMENT_V2 = 16,
//    	ALLOW_SYNDICATE_TO_UPDATE_KYC = 17,
//    	DO_NOT_BUILD_ACCOUNT_IF_VERSION_EQUALS_OR_GREATER = 18,
//    	ALLOW_TO_SPECIFY_REQUIRED_BASE_ASSET_AMOUNT_FOR_HARD_CAP = 19,
//    	KYC_RULES = 20,
//    	ALLOW_TO_CREATE_SEVERAL_SALES = 21,
//    	KEY_VALUE_POOL_ENTRY_EXPIRES_AT = 22,
//    	KEY_VALUE_UPDATE = 23,
//    	ALLOW_TO_CANCEL_SALE_PARTICIP_WITHOUT_SPECIFING_BALANCE = 24,
//    	DETAILS_MAX_LENGTH_EXTENDED = 25,
//    	ALLOW_MASTER_TO_MANAGE_SALE = 26,
//    	USE_SALE_ANTE = 27,
//    	FIX_ASSET_PAIRS_CREATION_IN_SALE_CREATION = 28,
//    	STATABLE_SALES = 29,
//    	CREATE_ONLY_STATISTICS_V2 = 30,
//    	LIMITS_UPDATE_REQUEST_DEPRECATED_DOCUMENT_HASH = 31,
//    	FIX_PAYMENT_V2_FEE = 32,
//    	ADD_SALE_ID_REVIEW_REQUEST_RESULT = 33,
//    	FIX_SET_SALE_STATE_AND_CHECK_SALE_STATE_OPS = 34, // only master allowed to set sale state, max issuance after sale closure = pending + issued
//    	FIX_UPDATE_MAX_ISSUANCE = 35,
//    	ALLOW_CLOSE_SALE_WITH_NON_ZERO_BALANCE = 36,
//    	ALLOW_TO_UPDATE_VOTING_SALES_AS_PROMOTION = 37,
//    	ALLOW_TO_ISSUE_AFTER_SALE = 38,
//    	FIX_PAYMENT_V2_SEND_TO_SELF = 39,
//    	FIX_PAYMENT_V2_DEST_ACCOUNT_NOT_FOUND = 40,
//    	FIX_CREATE_KYC_REQUEST_AUTO_APPROVE = 41,
//    	ADD_TASKS_TO_REVIEWABLE_REQUEST = 42,
//    	USE_ONLY_PAYMENT_V2 = 43,
//        ADD_REVIEW_INVOICE_REQUEST_PAYMENT_RESPONSE = 44,
//        ADD_CONTRACT_ID_REVIEW_REQUEST_RESULT = 45,
//        ALLOW_TO_UPDATE_AND_REJECT_LIMITS_UPDATE_REQUESTS = 46,
//        ADD_CUSTOMER_DETAILS_TO_CONTRACT = 47
//    };
//
type LedgerVersion int32

const (
	LedgerVersionEmptyVersion                                     LedgerVersion = 0
	LedgerVersionPassExternalSysAccIdInCreateAcc                  LedgerVersion = 1
	LedgerVersionDetailedLedgerChanges                            LedgerVersion = 2
	LedgerVersionNewSignerTypes                                   LedgerVersion = 3
	LedgerVersionTypedSale                                        LedgerVersion = 4
	LedgerVersionUniqueBalanceCreation                            LedgerVersion = 5
	LedgerVersionAssetPreissuerMigration                          LedgerVersion = 6
	LedgerVersionAssetPreissuerMigrated                           LedgerVersion = 7
	LedgerVersionUseKycLevel                                      LedgerVersion = 8
	LedgerVersionErrorOnNonZeroTasksToRemoveInRejectKyc           LedgerVersion = 9
	LedgerVersionAllowAccountManagerToChangeKyc                   LedgerVersion = 10
	LedgerVersionChangeAssetIssuerBadAuthExtraFixed               LedgerVersion = 11
	LedgerVersionAutoCreateCommissionBalanceOnTransfer            LedgerVersion = 12
	LedgerVersionAllowRejectRequestOfBlockedRequestor             LedgerVersion = 13
	LedgerVersionAssetUpdateCheckReferenceExists                  LedgerVersion = 14
	LedgerVersionCrossAssetFee                                    LedgerVersion = 15
	LedgerVersionUsePaymentV2                                     LedgerVersion = 16
	LedgerVersionAllowSyndicateToUpdateKyc                        LedgerVersion = 17
	LedgerVersionDoNotBuildAccountIfVersionEqualsOrGreater        LedgerVersion = 18
	LedgerVersionAllowToSpecifyRequiredBaseAssetAmountForHardCap  LedgerVersion = 19
	LedgerVersionKycRules                                         LedgerVersion = 20
	LedgerVersionAllowToCreateSeveralSales                        LedgerVersion = 21
	LedgerVersionKeyValuePoolEntryExpiresAt                       LedgerVersion = 22
	LedgerVersionKeyValueUpdate                                   LedgerVersion = 23
	LedgerVersionAllowToCancelSaleParticipWithoutSpecifingBalance LedgerVersion = 24
	LedgerVersionDetailsMaxLengthExtended                         LedgerVersion = 25
	LedgerVersionAllowMasterToManageSale                          LedgerVersion = 26
	LedgerVersionUseSaleAnte                                      LedgerVersion = 27
	LedgerVersionFixAssetPairsCreationInSaleCreation              LedgerVersion = 28
	LedgerVersionStatableSales                                    LedgerVersion = 29
	LedgerVersionCreateOnlyStatisticsV2                           LedgerVersion = 30
	LedgerVersionLimitsUpdateRequestDeprecatedDocumentHash        LedgerVersion = 31
	LedgerVersionFixPaymentV2Fee                                  LedgerVersion = 32
	LedgerVersionAddSaleIdReviewRequestResult                     LedgerVersion = 33
	LedgerVersionFixSetSaleStateAndCheckSaleStateOps              LedgerVersion = 34
	LedgerVersionFixUpdateMaxIssuance                             LedgerVersion = 35
	LedgerVersionAllowCloseSaleWithNonZeroBalance                 LedgerVersion = 36
	LedgerVersionAllowToUpdateVotingSalesAsPromotion              LedgerVersion = 37
	LedgerVersionAllowToIssueAfterSale                            LedgerVersion = 38
	LedgerVersionFixPaymentV2SendToSelf                           LedgerVersion = 39
	LedgerVersionFixPaymentV2DestAccountNotFound                  LedgerVersion = 40
	LedgerVersionFixCreateKycRequestAutoApprove                   LedgerVersion = 41
	LedgerVersionAddTasksToReviewableRequest                      LedgerVersion = 42
	LedgerVersionUseOnlyPaymentV2                                 LedgerVersion = 43
	LedgerVersionAddReviewInvoiceRequestPaymentResponse           LedgerVersion = 44
	LedgerVersionAddContractIdReviewRequestResult                 LedgerVersion = 45
	LedgerVersionAllowToUpdateAndRejectLimitsUpdateRequests       LedgerVersion = 46
	LedgerVersionAddCustomerDetailsToContract                     LedgerVersion = 47
)

var LedgerVersionAll = []LedgerVersion{
	LedgerVersionEmptyVersion,
	LedgerVersionPassExternalSysAccIdInCreateAcc,
	LedgerVersionDetailedLedgerChanges,
	LedgerVersionNewSignerTypes,
	LedgerVersionTypedSale,
	LedgerVersionUniqueBalanceCreation,
	LedgerVersionAssetPreissuerMigration,
	LedgerVersionAssetPreissuerMigrated,
	LedgerVersionUseKycLevel,
	LedgerVersionErrorOnNonZeroTasksToRemoveInRejectKyc,
	LedgerVersionAllowAccountManagerToChangeKyc,
	LedgerVersionChangeAssetIssuerBadAuthExtraFixed,
	LedgerVersionAutoCreateCommissionBalanceOnTransfer,
	LedgerVersionAllowRejectRequestOfBlockedRequestor,
	LedgerVersionAssetUpdateCheckReferenceExists,
	LedgerVersionCrossAssetFee,
	LedgerVersionUsePaymentV2,
	LedgerVersionAllowSyndicateToUpdateKyc,
	LedgerVersionDoNotBuildAccountIfVersionEqualsOrGreater,
	LedgerVersionAllowToSpecifyRequiredBaseAssetAmountForHardCap,
	LedgerVersionKycRules,
	LedgerVersionAllowToCreateSeveralSales,
	LedgerVersionKeyValuePoolEntryExpiresAt,
	LedgerVersionKeyValueUpdate,
	LedgerVersionAllowToCancelSaleParticipWithoutSpecifingBalance,
	LedgerVersionDetailsMaxLengthExtended,
	LedgerVersionAllowMasterToManageSale,
	LedgerVersionUseSaleAnte,
	LedgerVersionFixAssetPairsCreationInSaleCreation,
	LedgerVersionStatableSales,
	LedgerVersionCreateOnlyStatisticsV2,
	LedgerVersionLimitsUpdateRequestDeprecatedDocumentHash,
	LedgerVersionFixPaymentV2Fee,
	LedgerVersionAddSaleIdReviewRequestResult,
	LedgerVersionFixSetSaleStateAndCheckSaleStateOps,
	LedgerVersionFixUpdateMaxIssuance,
	LedgerVersionAllowCloseSaleWithNonZeroBalance,
	LedgerVersionAllowToUpdateVotingSalesAsPromotion,
	LedgerVersionAllowToIssueAfterSale,
	LedgerVersionFixPaymentV2SendToSelf,
	LedgerVersionFixPaymentV2DestAccountNotFound,
	LedgerVersionFixCreateKycRequestAutoApprove,
	LedgerVersionAddTasksToReviewableRequest,
	LedgerVersionUseOnlyPaymentV2,
	LedgerVersionAddReviewInvoiceRequestPaymentResponse,
	LedgerVersionAddContractIdReviewRequestResult,
	LedgerVersionAllowToUpdateAndRejectLimitsUpdateRequests,
	LedgerVersionAddCustomerDetailsToContract,
}

var ledgerVersionMap = map[int32]string{
	0:  "LedgerVersionEmptyVersion",
	1:  "LedgerVersionPassExternalSysAccIdInCreateAcc",
	2:  "LedgerVersionDetailedLedgerChanges",
	3:  "LedgerVersionNewSignerTypes",
	4:  "LedgerVersionTypedSale",
	5:  "LedgerVersionUniqueBalanceCreation",
	6:  "LedgerVersionAssetPreissuerMigration",
	7:  "LedgerVersionAssetPreissuerMigrated",
	8:  "LedgerVersionUseKycLevel",
	9:  "LedgerVersionErrorOnNonZeroTasksToRemoveInRejectKyc",
	10: "LedgerVersionAllowAccountManagerToChangeKyc",
	11: "LedgerVersionChangeAssetIssuerBadAuthExtraFixed",
	12: "LedgerVersionAutoCreateCommissionBalanceOnTransfer",
	13: "LedgerVersionAllowRejectRequestOfBlockedRequestor",
	14: "LedgerVersionAssetUpdateCheckReferenceExists",
	15: "LedgerVersionCrossAssetFee",
	16: "LedgerVersionUsePaymentV2",
	17: "LedgerVersionAllowSyndicateToUpdateKyc",
	18: "LedgerVersionDoNotBuildAccountIfVersionEqualsOrGreater",
	19: "LedgerVersionAllowToSpecifyRequiredBaseAssetAmountForHardCap",
	20: "LedgerVersionKycRules",
	21: "LedgerVersionAllowToCreateSeveralSales",
	22: "LedgerVersionKeyValuePoolEntryExpiresAt",
	23: "LedgerVersionKeyValueUpdate",
	24: "LedgerVersionAllowToCancelSaleParticipWithoutSpecifingBalance",
	25: "LedgerVersionDetailsMaxLengthExtended",
	26: "LedgerVersionAllowMasterToManageSale",
	27: "LedgerVersionUseSaleAnte",
	28: "LedgerVersionFixAssetPairsCreationInSaleCreation",
	29: "LedgerVersionStatableSales",
	30: "LedgerVersionCreateOnlyStatisticsV2",
	31: "LedgerVersionLimitsUpdateRequestDeprecatedDocumentHash",
	32: "LedgerVersionFixPaymentV2Fee",
	33: "LedgerVersionAddSaleIdReviewRequestResult",
	34: "LedgerVersionFixSetSaleStateAndCheckSaleStateOps",
	35: "LedgerVersionFixUpdateMaxIssuance",
	36: "LedgerVersionAllowCloseSaleWithNonZeroBalance",
	37: "LedgerVersionAllowToUpdateVotingSalesAsPromotion",
	38: "LedgerVersionAllowToIssueAfterSale",
	39: "LedgerVersionFixPaymentV2SendToSelf",
	40: "LedgerVersionFixPaymentV2DestAccountNotFound",
	41: "LedgerVersionFixCreateKycRequestAutoApprove",
	42: "LedgerVersionAddTasksToReviewableRequest",
	43: "LedgerVersionUseOnlyPaymentV2",
	44: "LedgerVersionAddReviewInvoiceRequestPaymentResponse",
	45: "LedgerVersionAddContractIdReviewRequestResult",
	46: "LedgerVersionAllowToUpdateAndRejectLimitsUpdateRequests",
	47: "LedgerVersionAddCustomerDetailsToContract",
}

var ledgerVersionShortMap = map[int32]string{
	0:  "empty_version",
	1:  "pass_external_sys_acc_id_in_create_acc",
	2:  "detailed_ledger_changes",
	3:  "new_signer_types",
	4:  "typed_sale",
	5:  "unique_balance_creation",
	6:  "asset_preissuer_migration",
	7:  "asset_preissuer_migrated",
	8:  "use_kyc_level",
	9:  "error_on_non_zero_tasks_to_remove_in_reject_kyc",
	10: "allow_account_manager_to_change_kyc",
	11: "change_asset_issuer_bad_auth_extra_fixed",
	12: "auto_create_commission_balance_on_transfer",
	13: "allow_reject_request_of_blocked_requestor",
	14: "asset_update_check_reference_exists",
	15: "cross_asset_fee",
	16: "use_payment_v2",
	17: "allow_syndicate_to_update_kyc",
	18: "do_not_build_account_if_version_equals_or_greater",
	19: "allow_to_specify_required_base_asset_amount_for_hard_cap",
	20: "kyc_rules",
	21: "allow_to_create_several_sales",
	22: "key_value_pool_entry_expires_at",
	23: "key_value_update",
	24: "allow_to_cancel_sale_particip_without_specifing_balance",
	25: "details_max_length_extended",
	26: "allow_master_to_manage_sale",
	27: "use_sale_ante",
	28: "fix_asset_pairs_creation_in_sale_creation",
	29: "statable_sales",
	30: "create_only_statistics_v2",
	31: "limits_update_request_deprecated_document_hash",
	32: "fix_payment_v2_fee",
	33: "add_sale_id_review_request_result",
	34: "fix_set_sale_state_and_check_sale_state_ops",
	35: "fix_update_max_issuance",
	36: "allow_close_sale_with_non_zero_balance",
	37: "allow_to_update_voting_sales_as_promotion",
	38: "allow_to_issue_after_sale",
	39: "fix_payment_v2_send_to_self",
	40: "fix_payment_v2_dest_account_not_found",
	41: "fix_create_kyc_request_auto_approve",
	42: "add_tasks_to_reviewable_request",
	43: "use_only_payment_v2",
	44: "add_review_invoice_request_payment_response",
	45: "add_contract_id_review_request_result",
	46: "allow_to_update_and_reject_limits_update_requests",
	47: "add_customer_details_to_contract",
}

var ledgerVersionRevMap = map[string]int32{
	"LedgerVersionEmptyVersion":                                     0,
	"LedgerVersionPassExternalSysAccIdInCreateAcc":                  1,
	"LedgerVersionDetailedLedgerChanges":                            2,
	"LedgerVersionNewSignerTypes":                                   3,
	"LedgerVersionTypedSale":                                        4,
	"LedgerVersionUniqueBalanceCreation":                            5,
	"LedgerVersionAssetPreissuerMigration":                          6,
	"LedgerVersionAssetPreissuerMigrated":                           7,
	"LedgerVersionUseKycLevel":                                      8,
	"LedgerVersionErrorOnNonZeroTasksToRemoveInRejectKyc":           9,
	"LedgerVersionAllowAccountManagerToChangeKyc":                   10,
	"LedgerVersionChangeAssetIssuerBadAuthExtraFixed":               11,
	"LedgerVersionAutoCreateCommissionBalanceOnTransfer":            12,
	"LedgerVersionAllowRejectRequestOfBlockedRequestor":             13,
	"LedgerVersionAssetUpdateCheckReferenceExists":                  14,
	"LedgerVersionCrossAssetFee":                                    15,
	"LedgerVersionUsePaymentV2":                                     16,
	"LedgerVersionAllowSyndicateToUpdateKyc":                        17,
	"LedgerVersionDoNotBuildAccountIfVersionEqualsOrGreater":        18,
	"LedgerVersionAllowToSpecifyRequiredBaseAssetAmountForHardCap":  19,
	"LedgerVersionKycRules":                                         20,
	"LedgerVersionAllowToCreateSeveralSales":                        21,
	"LedgerVersionKeyValuePoolEntryExpiresAt":                       22,
	"LedgerVersionKeyValueUpdate":                                   23,
	"LedgerVersionAllowToCancelSaleParticipWithoutSpecifingBalance": 24,
	"LedgerVersionDetailsMaxLengthExtended":                         25,
	"LedgerVersionAllowMasterToManageSale":                          26,
	"LedgerVersionUseSaleAnte":                                      27,
	"LedgerVersionFixAssetPairsCreationInSaleCreation":              28,
	"LedgerVersionStatableSales":                                    29,
	"LedgerVersionCreateOnlyStatisticsV2":                           30,
	"LedgerVersionLimitsUpdateRequestDeprecatedDocumentHash":        31,
	"LedgerVersionFixPaymentV2Fee":                                  32,
	"LedgerVersionAddSaleIdReviewRequestResult":                     33,
	"LedgerVersionFixSetSaleStateAndCheckSaleStateOps":              34,
	"LedgerVersionFixUpdateMaxIssuance":                             35,
	"LedgerVersionAllowCloseSaleWithNonZeroBalance":                 36,
	"LedgerVersionAllowToUpdateVotingSalesAsPromotion":              37,
	"LedgerVersionAllowToIssueAfterSale":                            38,
	"LedgerVersionFixPaymentV2SendToSelf":                           39,
	"LedgerVersionFixPaymentV2DestAccountNotFound":                  40,
	"LedgerVersionFixCreateKycRequestAutoApprove":                   41,
	"LedgerVersionAddTasksToReviewableRequest":                      42,
	"LedgerVersionUseOnlyPaymentV2":                                 43,
	"LedgerVersionAddReviewInvoiceRequestPaymentResponse":           44,
	"LedgerVersionAddContractIdReviewRequestResult":                 45,
	"LedgerVersionAllowToUpdateAndRejectLimitsUpdateRequests":       46,
	"LedgerVersionAddCustomerDetailsToContract":                     47,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for LedgerVersion
func (e LedgerVersion) ValidEnum(v int32) bool {
	_, ok := ledgerVersionMap[v]
	return ok
}
func (e LedgerVersion) isFlag() bool {
	for i := len(LedgerVersionAll) - 1; i >= 0; i-- {
		expected := LedgerVersion(2) << uint64(len(LedgerVersionAll)-1) >> uint64(len(LedgerVersionAll)-i)
		if expected != LedgerVersionAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e LedgerVersion) String() string {
	name, _ := ledgerVersionMap[int32(e)]
	return name
}

func (e LedgerVersion) ShortString() string {
	name, _ := ledgerVersionShortMap[int32(e)]
	return name
}

func (e LedgerVersion) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range LedgerVersionAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *LedgerVersion) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = LedgerVersion(t.Value)
	return nil
}

// Signature is an XDR Typedef defines as:
//
//   typedef opaque Signature<64>;
//
type Signature []byte

// SignatureHint is an XDR Typedef defines as:
//
//   typedef opaque SignatureHint[4];
//
type SignatureHint [4]byte

// NodeId is an XDR Typedef defines as:
//
//   typedef PublicKey NodeID;
//
type NodeId PublicKey

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u NodeId) SwitchFieldName() string {
	return PublicKey(u).SwitchFieldName()
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of PublicKey
func (u NodeId) ArmForSwitch(sw int32) (string, bool) {
	return PublicKey(u).ArmForSwitch(sw)
}

// NewNodeId creates a new  NodeId.
func NewNodeId(aType CryptoKeyType, value interface{}) (result NodeId, err error) {
	u, err := NewPublicKey(aType, value)
	result = NodeId(u)
	return
}

// MustEd25519 retrieves the Ed25519 value from the union,
// panicing if the value is not set.
func (u NodeId) MustEd25519() Uint256 {
	return PublicKey(u).MustEd25519()
}

// GetEd25519 retrieves the Ed25519 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u NodeId) GetEd25519() (result Uint256, ok bool) {
	return PublicKey(u).GetEd25519()
}

// Curve25519Secret is an XDR Struct defines as:
//
//   struct Curve25519Secret
//    {
//            opaque key[32];
//    };
//
type Curve25519Secret struct {
	Key [32]byte `json:"key,omitempty"`
}

// Curve25519Public is an XDR Struct defines as:
//
//   struct Curve25519Public
//    {
//            opaque key[32];
//    };
//
type Curve25519Public struct {
	Key [32]byte `json:"key,omitempty"`
}

// HmacSha256Key is an XDR Struct defines as:
//
//   struct HmacSha256Key
//    {
//            opaque key[32];
//    };
//
type HmacSha256Key struct {
	Key [32]byte `json:"key,omitempty"`
}

// HmacSha256Mac is an XDR Struct defines as:
//
//   struct HmacSha256Mac
//    {
//            opaque mac[32];
//    };
//
type HmacSha256Mac struct {
	Mac [32]byte `json:"mac,omitempty"`
}

// AccountId is an XDR Typedef defines as:
//
//   typedef PublicKey AccountID;
//
type AccountId PublicKey

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u AccountId) SwitchFieldName() string {
	return PublicKey(u).SwitchFieldName()
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of PublicKey
func (u AccountId) ArmForSwitch(sw int32) (string, bool) {
	return PublicKey(u).ArmForSwitch(sw)
}

// NewAccountId creates a new  AccountId.
func NewAccountId(aType CryptoKeyType, value interface{}) (result AccountId, err error) {
	u, err := NewPublicKey(aType, value)
	result = AccountId(u)
	return
}

// MustEd25519 retrieves the Ed25519 value from the union,
// panicing if the value is not set.
func (u AccountId) MustEd25519() Uint256 {
	return PublicKey(u).MustEd25519()
}

// GetEd25519 retrieves the Ed25519 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u AccountId) GetEd25519() (result Uint256, ok bool) {
	return PublicKey(u).GetEd25519()
}

// BalanceId is an XDR Typedef defines as:
//
//   typedef PublicKey BalanceID;
//
type BalanceId PublicKey

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u BalanceId) SwitchFieldName() string {
	return PublicKey(u).SwitchFieldName()
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of PublicKey
func (u BalanceId) ArmForSwitch(sw int32) (string, bool) {
	return PublicKey(u).ArmForSwitch(sw)
}

// NewBalanceId creates a new  BalanceId.
func NewBalanceId(aType CryptoKeyType, value interface{}) (result BalanceId, err error) {
	u, err := NewPublicKey(aType, value)
	result = BalanceId(u)
	return
}

// MustEd25519 retrieves the Ed25519 value from the union,
// panicing if the value is not set.
func (u BalanceId) MustEd25519() Uint256 {
	return PublicKey(u).MustEd25519()
}

// GetEd25519 retrieves the Ed25519 value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u BalanceId) GetEd25519() (result Uint256, ok bool) {
	return PublicKey(u).GetEd25519()
}

// Thresholds is an XDR Typedef defines as:
//
//   typedef opaque Thresholds[4];
//
type Thresholds [4]byte

// String32 is an XDR Typedef defines as:
//
//   typedef string string32<32>;
//
type String32 string

// XDRMaxSize implements the Sized interface for String32
func (e String32) XDRMaxSize() int {
	return 32
}

// String64 is an XDR Typedef defines as:
//
//   typedef string string64<64>;
//
type String64 string

// XDRMaxSize implements the Sized interface for String64
func (e String64) XDRMaxSize() int {
	return 64
}

// String256 is an XDR Typedef defines as:
//
//   typedef string string256<256>;
//
type String256 string

// XDRMaxSize implements the Sized interface for String256
func (e String256) XDRMaxSize() int {
	return 256
}

// Longstring is an XDR Typedef defines as:
//
//   typedef string longstring<>;
//
type Longstring string

// AssetCode is an XDR Typedef defines as:
//
//   typedef string AssetCode<16>;
//
type AssetCode string

// XDRMaxSize implements the Sized interface for AssetCode
func (e AssetCode) XDRMaxSize() int {
	return 16
}

// Salt is an XDR Typedef defines as:
//
//   typedef uint64 Salt;
//
type Salt Uint64

// DataValue is an XDR Typedef defines as:
//
//   typedef opaque DataValue<64>;
//
type DataValue []byte

// FeeExt is an XDR NestedUnion defines as:
//
//   union switch(LedgerVersion v)
//        {
//            case EMPTY_VERSION:
//                void;
//        }
//
type FeeExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u FeeExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of FeeExt
func (u FeeExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewFeeExt creates a new  FeeExt.
func NewFeeExt(v LedgerVersion, value interface{}) (result FeeExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// Fee is an XDR Struct defines as:
//
//   struct Fee {
//    	uint64 fixed;
//    	uint64 percent;
//
//        // reserved for future use
//        union switch(LedgerVersion v)
//        {
//            case EMPTY_VERSION:
//                void;
//        }
//        ext;
//    };
//
type Fee struct {
	Fixed   Uint64 `json:"fixed,omitempty"`
	Percent Uint64 `json:"percent,omitempty"`
	Ext     FeeExt `json:"ext,omitempty"`
}

// OperationType is an XDR Enum defines as:
//
//   enum OperationType
//    {
//        CREATE_ACCOUNT = 0,
//        PAYMENT = 1,
//        SET_OPTIONS = 2,
//        CREATE_ISSUANCE_REQUEST = 3,
//        SET_FEES = 5,
//    	MANAGE_ACCOUNT = 6,
//        CREATE_WITHDRAWAL_REQUEST = 7,
//        MANAGE_BALANCE = 9,
//        MANAGE_ASSET = 11,
//        CREATE_PREISSUANCE_REQUEST = 12,
//        MANAGE_LIMITS = 13,
//        DIRECT_DEBIT = 14,
//    	MANAGE_ASSET_PAIR = 15,
//    	MANAGE_OFFER = 16,
//        MANAGE_INVOICE_REQUEST = 17,
//    	REVIEW_REQUEST = 18,
//    	CREATE_SALE_REQUEST = 19,
//    	CHECK_SALE_STATE = 20,
//        CREATE_AML_ALERT = 21,
//        CREATE_KYC_REQUEST = 22,
//        PAYMENT_V2 = 23,
//        MANAGE_EXTERNAL_SYSTEM_ACCOUNT_ID_POOL_ENTRY = 24,
//        BIND_EXTERNAL_SYSTEM_ACCOUNT_ID = 25,
//        MANAGE_SALE = 26,
//        MANAGE_KEY_VALUE = 27,
//        CREATE_MANAGE_LIMITS_REQUEST = 28,
//        MANAGE_CONTRACT_REQUEST = 29,
//        MANAGE_CONTRACT = 30
//    };
//
type OperationType int32

const (
	OperationTypeCreateAccount                          OperationType = 0
	OperationTypePayment                                OperationType = 1
	OperationTypeSetOptions                             OperationType = 2
	OperationTypeCreateIssuanceRequest                  OperationType = 3
	OperationTypeSetFees                                OperationType = 5
	OperationTypeManageAccount                          OperationType = 6
	OperationTypeCreateWithdrawalRequest                OperationType = 7
	OperationTypeManageBalance                          OperationType = 9
	OperationTypeManageAsset                            OperationType = 11
	OperationTypeCreatePreissuanceRequest               OperationType = 12
	OperationTypeManageLimits                           OperationType = 13
	OperationTypeDirectDebit                            OperationType = 14
	OperationTypeManageAssetPair                        OperationType = 15
	OperationTypeManageOffer                            OperationType = 16
	OperationTypeManageInvoiceRequest                   OperationType = 17
	OperationTypeReviewRequest                          OperationType = 18
	OperationTypeCreateSaleRequest                      OperationType = 19
	OperationTypeCheckSaleState                         OperationType = 20
	OperationTypeCreateAmlAlert                         OperationType = 21
	OperationTypeCreateKycRequest                       OperationType = 22
	OperationTypePaymentV2                              OperationType = 23
	OperationTypeManageExternalSystemAccountIdPoolEntry OperationType = 24
	OperationTypeBindExternalSystemAccountId            OperationType = 25
	OperationTypeManageSale                             OperationType = 26
	OperationTypeManageKeyValue                         OperationType = 27
	OperationTypeCreateManageLimitsRequest              OperationType = 28
	OperationTypeManageContractRequest                  OperationType = 29
	OperationTypeManageContract                         OperationType = 30
)

var OperationTypeAll = []OperationType{
	OperationTypeCreateAccount,
	OperationTypePayment,
	OperationTypeSetOptions,
	OperationTypeCreateIssuanceRequest,
	OperationTypeSetFees,
	OperationTypeManageAccount,
	OperationTypeCreateWithdrawalRequest,
	OperationTypeManageBalance,
	OperationTypeManageAsset,
	OperationTypeCreatePreissuanceRequest,
	OperationTypeManageLimits,
	OperationTypeDirectDebit,
	OperationTypeManageAssetPair,
	OperationTypeManageOffer,
	OperationTypeManageInvoiceRequest,
	OperationTypeReviewRequest,
	OperationTypeCreateSaleRequest,
	OperationTypeCheckSaleState,
	OperationTypeCreateAmlAlert,
	OperationTypeCreateKycRequest,
	OperationTypePaymentV2,
	OperationTypeManageExternalSystemAccountIdPoolEntry,
	OperationTypeBindExternalSystemAccountId,
	OperationTypeManageSale,
	OperationTypeManageKeyValue,
	OperationTypeCreateManageLimitsRequest,
	OperationTypeManageContractRequest,
	OperationTypeManageContract,
}

var operationTypeMap = map[int32]string{
	0:  "OperationTypeCreateAccount",
	1:  "OperationTypePayment",
	2:  "OperationTypeSetOptions",
	3:  "OperationTypeCreateIssuanceRequest",
	5:  "OperationTypeSetFees",
	6:  "OperationTypeManageAccount",
	7:  "OperationTypeCreateWithdrawalRequest",
	9:  "OperationTypeManageBalance",
	11: "OperationTypeManageAsset",
	12: "OperationTypeCreatePreissuanceRequest",
	13: "OperationTypeManageLimits",
	14: "OperationTypeDirectDebit",
	15: "OperationTypeManageAssetPair",
	16: "OperationTypeManageOffer",
	17: "OperationTypeManageInvoiceRequest",
	18: "OperationTypeReviewRequest",
	19: "OperationTypeCreateSaleRequest",
	20: "OperationTypeCheckSaleState",
	21: "OperationTypeCreateAmlAlert",
	22: "OperationTypeCreateKycRequest",
	23: "OperationTypePaymentV2",
	24: "OperationTypeManageExternalSystemAccountIdPoolEntry",
	25: "OperationTypeBindExternalSystemAccountId",
	26: "OperationTypeManageSale",
	27: "OperationTypeManageKeyValue",
	28: "OperationTypeCreateManageLimitsRequest",
	29: "OperationTypeManageContractRequest",
	30: "OperationTypeManageContract",
}

var operationTypeShortMap = map[int32]string{
	0:  "create_account",
	1:  "payment",
	2:  "set_options",
	3:  "create_issuance_request",
	5:  "set_fees",
	6:  "manage_account",
	7:  "create_withdrawal_request",
	9:  "manage_balance",
	11: "manage_asset",
	12: "create_preissuance_request",
	13: "manage_limits",
	14: "direct_debit",
	15: "manage_asset_pair",
	16: "manage_offer",
	17: "manage_invoice_request",
	18: "review_request",
	19: "create_sale_request",
	20: "check_sale_state",
	21: "create_aml_alert",
	22: "create_kyc_request",
	23: "payment_v2",
	24: "manage_external_system_account_id_pool_entry",
	25: "bind_external_system_account_id",
	26: "manage_sale",
	27: "manage_key_value",
	28: "create_manage_limits_request",
	29: "manage_contract_request",
	30: "manage_contract",
}

var operationTypeRevMap = map[string]int32{
	"OperationTypeCreateAccount":                          0,
	"OperationTypePayment":                                1,
	"OperationTypeSetOptions":                             2,
	"OperationTypeCreateIssuanceRequest":                  3,
	"OperationTypeSetFees":                                5,
	"OperationTypeManageAccount":                          6,
	"OperationTypeCreateWithdrawalRequest":                7,
	"OperationTypeManageBalance":                          9,
	"OperationTypeManageAsset":                            11,
	"OperationTypeCreatePreissuanceRequest":               12,
	"OperationTypeManageLimits":                           13,
	"OperationTypeDirectDebit":                            14,
	"OperationTypeManageAssetPair":                        15,
	"OperationTypeManageOffer":                            16,
	"OperationTypeManageInvoiceRequest":                   17,
	"OperationTypeReviewRequest":                          18,
	"OperationTypeCreateSaleRequest":                      19,
	"OperationTypeCheckSaleState":                         20,
	"OperationTypeCreateAmlAlert":                         21,
	"OperationTypeCreateKycRequest":                       22,
	"OperationTypePaymentV2":                              23,
	"OperationTypeManageExternalSystemAccountIdPoolEntry": 24,
	"OperationTypeBindExternalSystemAccountId":            25,
	"OperationTypeManageSale":                             26,
	"OperationTypeManageKeyValue":                         27,
	"OperationTypeCreateManageLimitsRequest":              28,
	"OperationTypeManageContractRequest":                  29,
	"OperationTypeManageContract":                         30,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for OperationType
func (e OperationType) ValidEnum(v int32) bool {
	_, ok := operationTypeMap[v]
	return ok
}
func (e OperationType) isFlag() bool {
	for i := len(OperationTypeAll) - 1; i >= 0; i-- {
		expected := OperationType(2) << uint64(len(OperationTypeAll)-1) >> uint64(len(OperationTypeAll)-i)
		if expected != OperationTypeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e OperationType) String() string {
	name, _ := operationTypeMap[int32(e)]
	return name
}

func (e OperationType) ShortString() string {
	name, _ := operationTypeShortMap[int32(e)]
	return name
}

func (e OperationType) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range OperationTypeAll {
			if (value & e) == value {
				result.Flags = append(result.Flags, flagValue{
					Value: int32(value),
					Name:  value.ShortString(),
				})
			}
		}
		return json.Marshal(&result)
	} else {
		// marshal as enum
		result := enum{
			Value:  int32(e),
			String: e.ShortString(),
		}
		return json.Marshal(&result)
	}
}

func (e *OperationType) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = OperationType(t.Value)
	return nil
}

// DecoratedSignature is an XDR Struct defines as:
//
//   struct DecoratedSignature
//    {
//        SignatureHint hint;  // last 4 bytes of the public key, used as a hint
//        Signature signature; // actual signature
//    };
//
type DecoratedSignature struct {
	Hint      SignatureHint `json:"hint,omitempty"`
	Signature Signature     `json:"signature,omitempty"`
}

var fmtTest = fmt.Sprint("this is a dummy usage of fmt")
