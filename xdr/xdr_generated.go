// Package xdr is generated from:
//
//  xdr/raw/Stellar-SCP.x
//  xdr/raw/Stellar-ledger-entries-account-limits.x
//  xdr/raw/Stellar-ledger-entries-account-type-limits.x
//  xdr/raw/Stellar-ledger-entries-account.x
//  xdr/raw/Stellar-ledger-entries-asset-pair.x
//  xdr/raw/Stellar-ledger-entries-asset.x
//  xdr/raw/Stellar-ledger-entries-balance.x
//  xdr/raw/Stellar-ledger-entries-external-system-id.x
//  xdr/raw/Stellar-ledger-entries-fee.x
//  xdr/raw/Stellar-ledger-entries-invoice.x
//  xdr/raw/Stellar-ledger-entries-offer.x
//  xdr/raw/Stellar-ledger-entries-payment-request.x
//  xdr/raw/Stellar-ledger-entries-reference.x
//  xdr/raw/Stellar-ledger-entries-reviewable-request.x
//  xdr/raw/Stellar-ledger-entries-sale.x
//  xdr/raw/Stellar-ledger-entries-statistics.x
//  xdr/raw/Stellar-ledger-entries.x
//  xdr/raw/Stellar-ledger.x
//  xdr/raw/Stellar-operation-create-account.x
//  xdr/raw/Stellar-operation-create-issuance-request.x
//  xdr/raw/Stellar-operation-create-preissuance-request.x
//  xdr/raw/Stellar-operation-create-sale-creation-request.x
//  xdr/raw/Stellar-operation-create-withdrawal-request.x
//  xdr/raw/Stellar-operation-direct-debit.x
//  xdr/raw/Stellar-operation-manage-account.x
//  xdr/raw/Stellar-operation-manage-asset-pair.x
//  xdr/raw/Stellar-operation-manage-asset.x
//  xdr/raw/Stellar-operation-manage-balance.x
//  xdr/raw/Stellar-operation-manage-invoice.x
//  xdr/raw/Stellar-operation-manage-offer.x
//  xdr/raw/Stellar-operation-payment.x
//  xdr/raw/Stellar-operation-recover.x
//  xdr/raw/Stellar-operation-review-payment-request.x
//  xdr/raw/Stellar-operation-review-request.x
//  xdr/raw/Stellar-operation-set-fees.x
//  xdr/raw/Stellar-operation-set-limits.x
//  xdr/raw/Stellar-operation-set-options.x
//  xdr/raw/Stellar-overlay.x
//  xdr/raw/Stellar-reviewable-request-asset.x
//  xdr/raw/Stellar-reviewable-request-issuance.x
//  xdr/raw/Stellar-reviewable-request-sale.x
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
//    	ISSUANCE_MANAGER = 128, // allowed to make preissuance request, review issuance
//    	INVOICE_MANAGER = 256, // allowed to create payment requests to other accounts
//    	PAYMENT_OPERATOR = 512, // allowed to review payment requests
//    	LIMITS_MANAGER = 1024, // allowed to change limits
//    	ACCOUNT_MANAGER = 2048, // allowed to add/delete signers and trust
//    	COMMISSION_BALANCE_MANAGER  = 4096,// allowed to spend from commission balances
//    	OPERATIONAL_BALANCE_MANAGER = 8192 // allowed to spend from operational balances
//    };
//
type SignerType int32

const (
	SignerTypeReader                    SignerType = 1
	SignerTypeNotVerifiedAccManager     SignerType = 2
	SignerTypeGeneralAccManager         SignerType = 4
	SignerTypeDirectDebitOperator       SignerType = 8
	SignerTypeAssetManager              SignerType = 16
	SignerTypeAssetRateManager          SignerType = 32
	SignerTypeBalanceManager            SignerType = 64
	SignerTypeIssuanceManager           SignerType = 128
	SignerTypeInvoiceManager            SignerType = 256
	SignerTypePaymentOperator           SignerType = 512
	SignerTypeLimitsManager             SignerType = 1024
	SignerTypeAccountManager            SignerType = 2048
	SignerTypeCommissionBalanceManager  SignerType = 4096
	SignerTypeOperationalBalanceManager SignerType = 8192
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
}

var signerTypeMap = map[int32]string{
	1:    "SignerTypeReader",
	2:    "SignerTypeNotVerifiedAccManager",
	4:    "SignerTypeGeneralAccManager",
	8:    "SignerTypeDirectDebitOperator",
	16:   "SignerTypeAssetManager",
	32:   "SignerTypeAssetRateManager",
	64:   "SignerTypeBalanceManager",
	128:  "SignerTypeIssuanceManager",
	256:  "SignerTypeInvoiceManager",
	512:  "SignerTypePaymentOperator",
	1024: "SignerTypeLimitsManager",
	2048: "SignerTypeAccountManager",
	4096: "SignerTypeCommissionBalanceManager",
	8192: "SignerTypeOperationalBalanceManager",
}

var signerTypeShortMap = map[int32]string{
	1:    "reader",
	2:    "not_verified_acc_manager",
	4:    "general_acc_manager",
	8:    "direct_debit_operator",
	16:   "asset_manager",
	32:   "asset_rate_manager",
	64:   "balance_manager",
	128:  "issuance_manager",
	256:  "invoice_manager",
	512:  "payment_operator",
	1024: "limits_manager",
	2048: "account_manager",
	4096: "commission_balance_manager",
	8192: "operational_balance_manager",
}

var signerTypeRevMap = map[string]int32{
	"SignerTypeReader":                    1,
	"SignerTypeNotVerifiedAccManager":     2,
	"SignerTypeGeneralAccManager":         4,
	"SignerTypeDirectDebitOperator":       8,
	"SignerTypeAssetManager":              16,
	"SignerTypeAssetRateManager":          32,
	"SignerTypeBalanceManager":            64,
	"SignerTypeIssuanceManager":           128,
	"SignerTypeInvoiceManager":            256,
	"SignerTypePaymentOperator":           512,
	"SignerTypeLimitsManager":             1024,
	"SignerTypeAccountManager":            2048,
	"SignerTypeCommissionBalanceManager":  4096,
	"SignerTypeOperationalBalanceManager": 8192,
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
//    	SYNDICATE = 6 // can create asset
//    };
//
type AccountType int32

const (
	AccountTypeOperational AccountType = 1
	AccountTypeGeneral     AccountType = 2
	AccountTypeCommission  AccountType = 3
	AccountTypeMaster      AccountType = 4
	AccountTypeNotVerified AccountType = 5
	AccountTypeSyndicate   AccountType = 6
)

var AccountTypeAll = []AccountType{
	AccountTypeOperational,
	AccountTypeGeneral,
	AccountTypeCommission,
	AccountTypeMaster,
	AccountTypeNotVerified,
	AccountTypeSyndicate,
}

var accountTypeMap = map[int32]string{
	1: "AccountTypeOperational",
	2: "AccountTypeGeneral",
	3: "AccountTypeCommission",
	4: "AccountTypeMaster",
	5: "AccountTypeNotVerified",
	6: "AccountTypeSyndicate",
}

var accountTypeShortMap = map[int32]string{
	1: "operational",
	2: "general",
	3: "commission",
	4: "master",
	5: "not_verified",
	6: "syndicate",
}

var accountTypeRevMap = map[string]int32{
	"AccountTypeOperational": 1,
	"AccountTypeGeneral":     2,
	"AccountTypeCommission":  3,
	"AccountTypeMaster":      4,
	"AccountTypeNotVerified": 5,
	"AccountTypeSyndicate":   6,
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
//    	SUSPICIOUS_BEHAVIOR = 4
//    };
//
type BlockReasons int32

const (
	BlockReasonsRecoveryRequest    BlockReasons = 1
	BlockReasonsKycUpdate          BlockReasons = 2
	BlockReasonsSuspiciousBehavior BlockReasons = 4
)

var BlockReasonsAll = []BlockReasons{
	BlockReasonsRecoveryRequest,
	BlockReasonsKycUpdate,
	BlockReasonsSuspiciousBehavior,
}

var blockReasonsMap = map[int32]string{
	1: "BlockReasonsRecoveryRequest",
	2: "BlockReasonsKycUpdate",
	4: "BlockReasonsSuspiciousBehavior",
}

var blockReasonsShortMap = map[int32]string{
	1: "recovery_request",
	2: "kyc_update",
	4: "suspicious_behavior",
}

var blockReasonsRevMap = map[string]int32{
	"BlockReasonsRecoveryRequest":    1,
	"BlockReasonsKycUpdate":          2,
	"BlockReasonsSuspiciousBehavior": 4,
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
//        }
//
type AccountEntryExt struct {
	V LedgerVersion `json:"v,omitempty"`
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
	}
	return "-", false
}

// NewAccountEntryExt creates a new  AccountEntryExt.
func NewAccountEntryExt(v LedgerVersion, value interface{}) (result AccountEntryExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// AccountEntry is an XDR Struct defines as:
//
//   struct AccountEntry
//    {
//        AccountID accountID;      // master public key for this account
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
//        }
//        ext;
//    };
//
type AccountEntry struct {
	AccountId    AccountId       `json:"accountID,omitempty"`
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
//    	TRADEABLE = 1, // if not set pair can not be traided
//    	PHYSICAL_PRICE_RESTRICTION = 2, // if set, then prices for new offers must be greater then physical price with correction
//    	CURRENT_PRICE_RESTRICTION = 4 // if set, then price for new offers must be in interval of (1 +- maxPriceStep)*currentPrice
//    };
//
type AssetPairPolicy int32

const (
	AssetPairPolicyTradeable                AssetPairPolicy = 1
	AssetPairPolicyPhysicalPriceRestriction AssetPairPolicy = 2
	AssetPairPolicyCurrentPriceRestriction  AssetPairPolicy = 4
)

var AssetPairPolicyAll = []AssetPairPolicy{
	AssetPairPolicyTradeable,
	AssetPairPolicyPhysicalPriceRestriction,
	AssetPairPolicyCurrentPriceRestriction,
}

var assetPairPolicyMap = map[int32]string{
	1: "AssetPairPolicyTradeable",
	2: "AssetPairPolicyPhysicalPriceRestriction",
	4: "AssetPairPolicyCurrentPriceRestriction",
}

var assetPairPolicyShortMap = map[int32]string{
	1: "tradeable",
	2: "physical_price_restriction",
	4: "current_price_restriction",
}

var assetPairPolicyRevMap = map[string]int32{
	"AssetPairPolicyTradeable":                1,
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
//    	WITHDRAWABLE = 8
//    };
//
type AssetPolicy int32

const (
	AssetPolicyTransferable    AssetPolicy = 1
	AssetPolicyBaseAsset       AssetPolicy = 2
	AssetPolicyStatsQuoteAsset AssetPolicy = 4
	AssetPolicyWithdrawable    AssetPolicy = 8
)

var AssetPolicyAll = []AssetPolicy{
	AssetPolicyTransferable,
	AssetPolicyBaseAsset,
	AssetPolicyStatsQuoteAsset,
	AssetPolicyWithdrawable,
}

var assetPolicyMap = map[int32]string{
	1: "AssetPolicyTransferable",
	2: "AssetPolicyBaseAsset",
	4: "AssetPolicyStatsQuoteAsset",
	8: "AssetPolicyWithdrawable",
}

var assetPolicyShortMap = map[int32]string{
	1: "transferable",
	2: "base_asset",
	4: "stats_quote_asset",
	8: "withdrawable",
}

var assetPolicyRevMap = map[string]int32{
	"AssetPolicyTransferable":    1,
	"AssetPolicyBaseAsset":       2,
	"AssetPolicyStatsQuoteAsset": 4,
	"AssetPolicyWithdrawable":    8,
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
//    	uint64 lockedIssuance; // number of tokens locked for entries like token sale. lockedIssuance + issued can not be > maxIssuanceAmount
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
	LockedIssuance        Uint64        `json:"lockedIssuance,omitempty"`
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

// ExternalSystemType is an XDR Enum defines as:
//
//   enum ExternalSystemType
//    {
//    	BITCOIN = 1,
//    	ETHEREUM = 2
//    };
//
type ExternalSystemType int32

const (
	ExternalSystemTypeBitcoin  ExternalSystemType = 1
	ExternalSystemTypeEthereum ExternalSystemType = 2
)

var ExternalSystemTypeAll = []ExternalSystemType{
	ExternalSystemTypeBitcoin,
	ExternalSystemTypeEthereum,
}

var externalSystemTypeMap = map[int32]string{
	1: "ExternalSystemTypeBitcoin",
	2: "ExternalSystemTypeEthereum",
}

var externalSystemTypeShortMap = map[int32]string{
	1: "bitcoin",
	2: "ethereum",
}

var externalSystemTypeRevMap = map[string]int32{
	"ExternalSystemTypeBitcoin":  1,
	"ExternalSystemTypeEthereum": 2,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ExternalSystemType
func (e ExternalSystemType) ValidEnum(v int32) bool {
	_, ok := externalSystemTypeMap[v]
	return ok
}
func (e ExternalSystemType) isFlag() bool {
	for i := len(ExternalSystemTypeAll) - 1; i >= 0; i-- {
		expected := ExternalSystemType(2) << uint64(len(ExternalSystemTypeAll)-1) >> uint64(len(ExternalSystemTypeAll)-i)
		if expected != ExternalSystemTypeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ExternalSystemType) String() string {
	name, _ := externalSystemTypeMap[int32(e)]
	return name
}

func (e ExternalSystemType) ShortString() string {
	name, _ := externalSystemTypeShortMap[int32(e)]
	return name
}

func (e ExternalSystemType) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range ExternalSystemTypeAll {
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

func (e *ExternalSystemType) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ExternalSystemType(t.Value)
	return nil
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
//        ExternalSystemType externalSystemType;
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
	ExternalSystemType ExternalSystemType         `json:"externalSystemType,omitempty"`
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
//        ISSUANCE_FEE = 3
//    };
//
type FeeType int32

const (
	FeeTypePaymentFee    FeeType = 0
	FeeTypeOfferFee      FeeType = 1
	FeeTypeWithdrawalFee FeeType = 2
	FeeTypeIssuanceFee   FeeType = 3
)

var FeeTypeAll = []FeeType{
	FeeTypePaymentFee,
	FeeTypeOfferFee,
	FeeTypeWithdrawalFee,
	FeeTypeIssuanceFee,
}

var feeTypeMap = map[int32]string{
	0: "FeeTypePaymentFee",
	1: "FeeTypeOfferFee",
	2: "FeeTypeWithdrawalFee",
	3: "FeeTypeIssuanceFee",
}

var feeTypeShortMap = map[int32]string{
	0: "payment_fee",
	1: "offer_fee",
	2: "withdrawal_fee",
	3: "issuance_fee",
}

var feeTypeRevMap = map[string]int32{
	"FeeTypePaymentFee":    0,
	"FeeTypeOfferFee":      1,
	"FeeTypeWithdrawalFee": 2,
	"FeeTypeIssuanceFee":   3,
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

// FeeEntryExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type FeeEntryExt struct {
	V LedgerVersion `json:"v,omitempty"`
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
	}
	return "-", false
}

// NewFeeEntryExt creates a new  FeeEntryExt.
func NewFeeEntryExt(v LedgerVersion, value interface{}) (result FeeEntryExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
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
//        }
//        ext;
//
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

// InvoiceState is an XDR Enum defines as:
//
//   enum InvoiceState
//    {
//        INVOICE_NEEDS_PAYMENT = 0,
//        INVOICE_NEEDS_PAYMENT_REVIEW = 1
//    };
//
type InvoiceState int32

const (
	InvoiceStateInvoiceNeedsPayment       InvoiceState = 0
	InvoiceStateInvoiceNeedsPaymentReview InvoiceState = 1
)

var InvoiceStateAll = []InvoiceState{
	InvoiceStateInvoiceNeedsPayment,
	InvoiceStateInvoiceNeedsPaymentReview,
}

var invoiceStateMap = map[int32]string{
	0: "InvoiceStateInvoiceNeedsPayment",
	1: "InvoiceStateInvoiceNeedsPaymentReview",
}

var invoiceStateShortMap = map[int32]string{
	0: "invoice_needs_payment",
	1: "invoice_needs_payment_review",
}

var invoiceStateRevMap = map[string]int32{
	"InvoiceStateInvoiceNeedsPayment":       0,
	"InvoiceStateInvoiceNeedsPaymentReview": 1,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for InvoiceState
func (e InvoiceState) ValidEnum(v int32) bool {
	_, ok := invoiceStateMap[v]
	return ok
}
func (e InvoiceState) isFlag() bool {
	for i := len(InvoiceStateAll) - 1; i >= 0; i-- {
		expected := InvoiceState(2) << uint64(len(InvoiceStateAll)-1) >> uint64(len(InvoiceStateAll)-i)
		if expected != InvoiceStateAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e InvoiceState) String() string {
	name, _ := invoiceStateMap[int32(e)]
	return name
}

func (e InvoiceState) ShortString() string {
	name, _ := invoiceStateShortMap[int32(e)]
	return name
}

func (e InvoiceState) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range InvoiceStateAll {
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

func (e *InvoiceState) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = InvoiceState(t.Value)
	return nil
}

// InvoiceEntryExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type InvoiceEntryExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u InvoiceEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of InvoiceEntryExt
func (u InvoiceEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewInvoiceEntryExt creates a new  InvoiceEntryExt.
func NewInvoiceEntryExt(v LedgerVersion, value interface{}) (result InvoiceEntryExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// InvoiceEntry is an XDR Struct defines as:
//
//   struct InvoiceEntry
//    {
//        uint64 invoiceID;
//        AccountID receiverAccount;
//        BalanceID receiverBalance;
//    	AccountID sender;
//        int64 amount;
//
//        InvoiceState state;
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
type InvoiceEntry struct {
	InvoiceId       Uint64          `json:"invoiceID,omitempty"`
	ReceiverAccount AccountId       `json:"receiverAccount,omitempty"`
	ReceiverBalance BalanceId       `json:"receiverBalance,omitempty"`
	Sender          AccountId       `json:"sender,omitempty"`
	Amount          Int64           `json:"amount,omitempty"`
	State           InvoiceState    `json:"state,omitempty"`
	Ext             InvoiceEntryExt `json:"ext,omitempty"`
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

// RequestType is an XDR Enum defines as:
//
//   enum RequestType
//    {
//        REQUEST_TYPE_SALE = 0,
//        REQUEST_TYPE_WITHDRAWAL = 1,
//        REQUEST_TYPE_REDEEM = 2,
//        REQUEST_TYPE_PAYMENT = 3
//    };
//
type RequestType int32

const (
	RequestTypeRequestTypeSale       RequestType = 0
	RequestTypeRequestTypeWithdrawal RequestType = 1
	RequestTypeRequestTypeRedeem     RequestType = 2
	RequestTypeRequestTypePayment    RequestType = 3
)

var RequestTypeAll = []RequestType{
	RequestTypeRequestTypeSale,
	RequestTypeRequestTypeWithdrawal,
	RequestTypeRequestTypeRedeem,
	RequestTypeRequestTypePayment,
}

var requestTypeMap = map[int32]string{
	0: "RequestTypeRequestTypeSale",
	1: "RequestTypeRequestTypeWithdrawal",
	2: "RequestTypeRequestTypeRedeem",
	3: "RequestTypeRequestTypePayment",
}

var requestTypeShortMap = map[int32]string{
	0: "request_type_sale",
	1: "request_type_withdrawal",
	2: "request_type_redeem",
	3: "request_type_payment",
}

var requestTypeRevMap = map[string]int32{
	"RequestTypeRequestTypeSale":       0,
	"RequestTypeRequestTypeWithdrawal": 1,
	"RequestTypeRequestTypeRedeem":     2,
	"RequestTypeRequestTypePayment":    3,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for RequestType
func (e RequestType) ValidEnum(v int32) bool {
	_, ok := requestTypeMap[v]
	return ok
}
func (e RequestType) isFlag() bool {
	for i := len(RequestTypeAll) - 1; i >= 0; i-- {
		expected := RequestType(2) << uint64(len(RequestTypeAll)-1) >> uint64(len(RequestTypeAll)-i)
		if expected != RequestTypeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e RequestType) String() string {
	name, _ := requestTypeMap[int32(e)]
	return name
}

func (e RequestType) ShortString() string {
	name, _ := requestTypeShortMap[int32(e)]
	return name
}

func (e RequestType) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range RequestTypeAll {
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

func (e *RequestType) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = RequestType(t.Value)
	return nil
}

// PaymentRequestEntryExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type PaymentRequestEntryExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u PaymentRequestEntryExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of PaymentRequestEntryExt
func (u PaymentRequestEntryExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewPaymentRequestEntryExt creates a new  PaymentRequestEntryExt.
func NewPaymentRequestEntryExt(v LedgerVersion, value interface{}) (result PaymentRequestEntryExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// PaymentRequestEntry is an XDR Struct defines as:
//
//   struct PaymentRequestEntry
//    {
//        uint64 paymentID;
//        BalanceID sourceBalance;
//        BalanceID* destinationBalance;
//        int64 sourceSend;
//        int64 sourceSendUniversal;
//        int64 destinationReceive;
//
//        uint64 createdAt;
//
//        uint64* invoiceID;
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
type PaymentRequestEntry struct {
	PaymentId           Uint64                 `json:"paymentID,omitempty"`
	SourceBalance       BalanceId              `json:"sourceBalance,omitempty"`
	DestinationBalance  *BalanceId             `json:"destinationBalance,omitempty"`
	SourceSend          Int64                  `json:"sourceSend,omitempty"`
	SourceSendUniversal Int64                  `json:"sourceSendUniversal,omitempty"`
	DestinationReceive  Int64                  `json:"destinationReceive,omitempty"`
	CreatedAt           Uint64                 `json:"createdAt,omitempty"`
	InvoiceId           *Uint64                `json:"invoiceID,omitempty"`
	Ext                 PaymentRequestEntryExt `json:"ext,omitempty"`
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
//    	SALE = 5
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
)

var ReviewableRequestTypeAll = []ReviewableRequestType{
	ReviewableRequestTypeAssetCreate,
	ReviewableRequestTypeAssetUpdate,
	ReviewableRequestTypePreIssuanceCreate,
	ReviewableRequestTypeIssuanceCreate,
	ReviewableRequestTypeWithdraw,
	ReviewableRequestTypeSale,
}

var reviewableRequestTypeMap = map[int32]string{
	0: "ReviewableRequestTypeAssetCreate",
	1: "ReviewableRequestTypeAssetUpdate",
	2: "ReviewableRequestTypePreIssuanceCreate",
	3: "ReviewableRequestTypeIssuanceCreate",
	4: "ReviewableRequestTypeWithdraw",
	5: "ReviewableRequestTypeSale",
}

var reviewableRequestTypeShortMap = map[int32]string{
	0: "asset_create",
	1: "asset_update",
	2: "pre_issuance_create",
	3: "issuance_create",
	4: "withdraw",
	5: "sale",
}

var reviewableRequestTypeRevMap = map[string]int32{
	"ReviewableRequestTypeAssetCreate":       0,
	"ReviewableRequestTypeAssetUpdate":       1,
	"ReviewableRequestTypePreIssuanceCreate": 2,
	"ReviewableRequestTypeIssuanceCreate":    3,
	"ReviewableRequestTypeWithdraw":          4,
	"ReviewableRequestTypeSale":              5,
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
//    	}
//
type ReviewableRequestEntryBody struct {
	Type                 ReviewableRequestType `json:"type,omitempty"`
	AssetCreationRequest *AssetCreationRequest `json:"assetCreationRequest,omitempty"`
	AssetUpdateRequest   *AssetUpdateRequest   `json:"assetUpdateRequest,omitempty"`
	PreIssuanceRequest   *PreIssuanceRequest   `json:"preIssuanceRequest,omitempty"`
	IssuanceRequest      *IssuanceRequest      `json:"issuanceRequest,omitempty"`
	WithdrawalRequest    *WithdrawalRequest    `json:"withdrawalRequest,omitempty"`
	SaleCreationRequest  *SaleCreationRequest  `json:"saleCreationRequest,omitempty"`
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

// ReviewableRequestEntryExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type ReviewableRequestEntryExt struct {
	V LedgerVersion `json:"v,omitempty"`
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
	}
	return "-", false
}

// NewReviewableRequestEntryExt creates a new  ReviewableRequestEntryExt.
func NewReviewableRequestEntryExt(v LedgerVersion, value interface{}) (result ReviewableRequestEntryExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ReviewableRequestEntry is an XDR Struct defines as:
//
//   struct ReviewableRequestEntry {
//    	uint64 requestID;
//    	Hash hash; // hash of the request body
//    	AccountID requestor;
//    	string256 rejectReason;
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
//    	} body;
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
type ReviewableRequestEntry struct {
	RequestId    Uint64                     `json:"requestID,omitempty"`
	Hash         Hash                       `json:"hash,omitempty"`
	Requestor    AccountId                  `json:"requestor,omitempty"`
	RejectReason String256                  `json:"rejectReason,omitempty"`
	Reviewer     AccountId                  `json:"reviewer,omitempty"`
	Reference    *String64                  `json:"reference,omitempty"`
	CreatedAt    Int64                      `json:"createdAt,omitempty"`
	Body         ReviewableRequestEntryBody `json:"body,omitempty"`
	Ext          ReviewableRequestEntryExt  `json:"ext,omitempty"`
}

// SaleEntryExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type SaleEntryExt struct {
	V LedgerVersion `json:"v,omitempty"`
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
	}
	return "-", false
}

// NewSaleEntryExt creates a new  SaleEntryExt.
func NewSaleEntryExt(v LedgerVersion, value interface{}) (result SaleEntryExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
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
//    	AssetCode quoteAsset; // asset in which participation will be accepted
//    	uint64 startTime; // start time of the sale
//    	uint64 endTime; // close time of the sale
//    	uint64 price; // price for 1 baseAsset in terms of quote asset
//    	uint64 softCap; // minimum amount of quote asset to be received at which sale will be considered a successful
//    	uint64 hardCap; // max amount of quote asset to be received
//    	longstring details; // sale specific details
//
//    	uint64 currentCap; // current capitalization
//
//    	union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type SaleEntry struct {
	SaleId     Uint64       `json:"saleID,omitempty"`
	OwnerId    AccountId    `json:"ownerID,omitempty"`
	BaseAsset  AssetCode    `json:"baseAsset,omitempty"`
	QuoteAsset AssetCode    `json:"quoteAsset,omitempty"`
	StartTime  Uint64       `json:"startTime,omitempty"`
	EndTime    Uint64       `json:"endTime,omitempty"`
	Price      Uint64       `json:"price,omitempty"`
	SoftCap    Uint64       `json:"softCap,omitempty"`
	HardCap    Uint64       `json:"hardCap,omitempty"`
	Details    Longstring   `json:"details,omitempty"`
	CurrentCap Uint64       `json:"currentCap,omitempty"`
	Ext        SaleEntryExt `json:"ext,omitempty"`
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
//    	int64 dailyOutcome;
//    	int64 weeklyOutcome;
//    	int64 monthlyOutcome;
//    	int64 annualOutcome;
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
	DailyOutcome   Int64              `json:"dailyOutcome,omitempty"`
	WeeklyOutcome  Int64              `json:"weeklyOutcome,omitempty"`
	MonthlyOutcome Int64              `json:"monthlyOutcome,omitempty"`
	AnnualOutcome  Int64              `json:"annualOutcome,omitempty"`
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
//        INVOICE = 14,
//    	REVIEWABLE_REQUEST = 15,
//    	EXTERNAL_SYSTEM_ACCOUNT_ID = 16,
//    	SALE = 17
//    };
//
type LedgerEntryType int32

const (
	LedgerEntryTypeAccount                 LedgerEntryType = 0
	LedgerEntryTypeFee                     LedgerEntryType = 2
	LedgerEntryTypeBalance                 LedgerEntryType = 4
	LedgerEntryTypePaymentRequest          LedgerEntryType = 5
	LedgerEntryTypeAsset                   LedgerEntryType = 6
	LedgerEntryTypeReferenceEntry          LedgerEntryType = 7
	LedgerEntryTypeAccountTypeLimits       LedgerEntryType = 8
	LedgerEntryTypeStatistics              LedgerEntryType = 9
	LedgerEntryTypeTrust                   LedgerEntryType = 10
	LedgerEntryTypeAccountLimits           LedgerEntryType = 11
	LedgerEntryTypeAssetPair               LedgerEntryType = 12
	LedgerEntryTypeOfferEntry              LedgerEntryType = 13
	LedgerEntryTypeInvoice                 LedgerEntryType = 14
	LedgerEntryTypeReviewableRequest       LedgerEntryType = 15
	LedgerEntryTypeExternalSystemAccountId LedgerEntryType = 16
	LedgerEntryTypeSale                    LedgerEntryType = 17
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
	LedgerEntryTypeInvoice,
	LedgerEntryTypeReviewableRequest,
	LedgerEntryTypeExternalSystemAccountId,
	LedgerEntryTypeSale,
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
	14: "LedgerEntryTypeInvoice",
	15: "LedgerEntryTypeReviewableRequest",
	16: "LedgerEntryTypeExternalSystemAccountId",
	17: "LedgerEntryTypeSale",
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
	14: "invoice",
	15: "reviewable_request",
	16: "external_system_account_id",
	17: "sale",
}

var ledgerEntryTypeRevMap = map[string]int32{
	"LedgerEntryTypeAccount":                 0,
	"LedgerEntryTypeFee":                     2,
	"LedgerEntryTypeBalance":                 4,
	"LedgerEntryTypePaymentRequest":          5,
	"LedgerEntryTypeAsset":                   6,
	"LedgerEntryTypeReferenceEntry":          7,
	"LedgerEntryTypeAccountTypeLimits":       8,
	"LedgerEntryTypeStatistics":              9,
	"LedgerEntryTypeTrust":                   10,
	"LedgerEntryTypeAccountLimits":           11,
	"LedgerEntryTypeAssetPair":               12,
	"LedgerEntryTypeOfferEntry":              13,
	"LedgerEntryTypeInvoice":                 14,
	"LedgerEntryTypeReviewableRequest":       15,
	"LedgerEntryTypeExternalSystemAccountId": 16,
	"LedgerEntryTypeSale":                    17,
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
//        case PAYMENT_REQUEST:
//            PaymentRequestEntry paymentRequest;
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
//        case INVOICE:
//            InvoiceEntry invoice;
//    	case REVIEWABLE_REQUEST:
//    		ReviewableRequestEntry reviewableRequest;
//    	case EXTERNAL_SYSTEM_ACCOUNT_ID:
//    		ExternalSystemAccountID externalSystemAccountID;
//    	case SALE:
//    		SaleEntry sale;
//        }
//
type LedgerEntryData struct {
	Type                    LedgerEntryType          `json:"type,omitempty"`
	Account                 *AccountEntry            `json:"account,omitempty"`
	FeeState                *FeeEntry                `json:"feeState,omitempty"`
	Balance                 *BalanceEntry            `json:"balance,omitempty"`
	PaymentRequest          *PaymentRequestEntry     `json:"paymentRequest,omitempty"`
	Asset                   *AssetEntry              `json:"asset,omitempty"`
	Reference               *ReferenceEntry          `json:"reference,omitempty"`
	AccountTypeLimits       *AccountTypeLimitsEntry  `json:"accountTypeLimits,omitempty"`
	Stats                   *StatisticsEntry         `json:"stats,omitempty"`
	Trust                   *TrustEntry              `json:"trust,omitempty"`
	AccountLimits           *AccountLimitsEntry      `json:"accountLimits,omitempty"`
	AssetPair               *AssetPairEntry          `json:"assetPair,omitempty"`
	Offer                   *OfferEntry              `json:"offer,omitempty"`
	Invoice                 *InvoiceEntry            `json:"invoice,omitempty"`
	ReviewableRequest       *ReviewableRequestEntry  `json:"reviewableRequest,omitempty"`
	ExternalSystemAccountId *ExternalSystemAccountId `json:"externalSystemAccountID,omitempty"`
	Sale                    *SaleEntry               `json:"sale,omitempty"`
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
	case LedgerEntryTypePaymentRequest:
		return "PaymentRequest", true
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
	case LedgerEntryTypeInvoice:
		return "Invoice", true
	case LedgerEntryTypeReviewableRequest:
		return "ReviewableRequest", true
	case LedgerEntryTypeExternalSystemAccountId:
		return "ExternalSystemAccountId", true
	case LedgerEntryTypeSale:
		return "Sale", true
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
	case LedgerEntryTypePaymentRequest:
		tv, ok := value.(PaymentRequestEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be PaymentRequestEntry")
			return
		}
		result.PaymentRequest = &tv
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
	case LedgerEntryTypeInvoice:
		tv, ok := value.(InvoiceEntry)
		if !ok {
			err = fmt.Errorf("invalid value, must be InvoiceEntry")
			return
		}
		result.Invoice = &tv
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

// MustPaymentRequest retrieves the PaymentRequest value from the union,
// panicing if the value is not set.
func (u LedgerEntryData) MustPaymentRequest() PaymentRequestEntry {
	val, ok := u.GetPaymentRequest()

	if !ok {
		panic("arm PaymentRequest is not set")
	}

	return val
}

// GetPaymentRequest retrieves the PaymentRequest value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryData) GetPaymentRequest() (result PaymentRequestEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "PaymentRequest" {
		result = *u.PaymentRequest
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

// MustInvoice retrieves the Invoice value from the union,
// panicing if the value is not set.
func (u LedgerEntryData) MustInvoice() InvoiceEntry {
	val, ok := u.GetInvoice()

	if !ok {
		panic("arm Invoice is not set")
	}

	return val
}

// GetInvoice retrieves the Invoice value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerEntryData) GetInvoice() (result InvoiceEntry, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Invoice" {
		result = *u.Invoice
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
//        case PAYMENT_REQUEST:
//            PaymentRequestEntry paymentRequest;
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
//        case INVOICE:
//            InvoiceEntry invoice;
//    	case REVIEWABLE_REQUEST:
//    		ReviewableRequestEntry reviewableRequest;
//    	case EXTERNAL_SYSTEM_ACCOUNT_ID:
//    		ExternalSystemAccountID externalSystemAccountID;
//    	case SALE:
//    		SaleEntry sale;
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

// LedgerKeyPaymentRequestExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//
type LedgerKeyPaymentRequestExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerKeyPaymentRequestExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerKeyPaymentRequestExt
func (u LedgerKeyPaymentRequestExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewLedgerKeyPaymentRequestExt creates a new  LedgerKeyPaymentRequestExt.
func NewLedgerKeyPaymentRequestExt(v LedgerVersion, value interface{}) (result LedgerKeyPaymentRequestExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// LedgerKeyPaymentRequest is an XDR NestedStruct defines as:
//
//   struct
//        {
//    		uint64 paymentID;
//    		union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//    		ext;
//        }
//
type LedgerKeyPaymentRequest struct {
	PaymentId Uint64                     `json:"paymentID,omitempty"`
	Ext       LedgerKeyPaymentRequestExt `json:"ext,omitempty"`
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

// LedgerKeyInvoiceExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//
type LedgerKeyInvoiceExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u LedgerKeyInvoiceExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of LedgerKeyInvoiceExt
func (u LedgerKeyInvoiceExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewLedgerKeyInvoiceExt creates a new  LedgerKeyInvoiceExt.
func NewLedgerKeyInvoiceExt(v LedgerVersion, value interface{}) (result LedgerKeyInvoiceExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// LedgerKeyInvoice is an XDR NestedStruct defines as:
//
//   struct {
//            uint64 invoiceID;
//    		union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//    		ext;
//        }
//
type LedgerKeyInvoice struct {
	InvoiceId Uint64              `json:"invoiceID,omitempty"`
	Ext       LedgerKeyInvoiceExt `json:"ext,omitempty"`
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
//    		ExternalSystemType externalSystemType;
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
	ExternalSystemType ExternalSystemType                  `json:"externalSystemType,omitempty"`
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
//    case PAYMENT_REQUEST:
//        struct
//        {
//    		uint64 paymentID;
//    		union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//    		ext;
//        } paymentRequest;
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
//    case INVOICE:
//        struct {
//            uint64 invoiceID;
//    		union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//    		ext;
//        } invoice;
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
//    		ExternalSystemType externalSystemType;
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
//    };
//
type LedgerKey struct {
	Type                    LedgerEntryType                   `json:"type,omitempty"`
	Account                 *LedgerKeyAccount                 `json:"account,omitempty"`
	FeeState                *LedgerKeyFeeState                `json:"feeState,omitempty"`
	Balance                 *LedgerKeyBalance                 `json:"balance,omitempty"`
	PaymentRequest          *LedgerKeyPaymentRequest          `json:"paymentRequest,omitempty"`
	Asset                   *LedgerKeyAsset                   `json:"asset,omitempty"`
	Reference               *LedgerKeyReference               `json:"reference,omitempty"`
	AccountTypeLimits       *LedgerKeyAccountTypeLimits       `json:"accountTypeLimits,omitempty"`
	Stats                   *LedgerKeyStats                   `json:"stats,omitempty"`
	Trust                   *LedgerKeyTrust                   `json:"trust,omitempty"`
	AccountLimits           *LedgerKeyAccountLimits           `json:"accountLimits,omitempty"`
	AssetPair               *LedgerKeyAssetPair               `json:"assetPair,omitempty"`
	Offer                   *LedgerKeyOffer                   `json:"offer,omitempty"`
	Invoice                 *LedgerKeyInvoice                 `json:"invoice,omitempty"`
	ReviewableRequest       *LedgerKeyReviewableRequest       `json:"reviewableRequest,omitempty"`
	ExternalSystemAccountId *LedgerKeyExternalSystemAccountId `json:"externalSystemAccountID,omitempty"`
	Sale                    *LedgerKeySale                    `json:"sale,omitempty"`
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
	case LedgerEntryTypePaymentRequest:
		return "PaymentRequest", true
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
	case LedgerEntryTypeInvoice:
		return "Invoice", true
	case LedgerEntryTypeReviewableRequest:
		return "ReviewableRequest", true
	case LedgerEntryTypeExternalSystemAccountId:
		return "ExternalSystemAccountId", true
	case LedgerEntryTypeSale:
		return "Sale", true
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
	case LedgerEntryTypePaymentRequest:
		tv, ok := value.(LedgerKeyPaymentRequest)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKeyPaymentRequest")
			return
		}
		result.PaymentRequest = &tv
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
	case LedgerEntryTypeInvoice:
		tv, ok := value.(LedgerKeyInvoice)
		if !ok {
			err = fmt.Errorf("invalid value, must be LedgerKeyInvoice")
			return
		}
		result.Invoice = &tv
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

// MustPaymentRequest retrieves the PaymentRequest value from the union,
// panicing if the value is not set.
func (u LedgerKey) MustPaymentRequest() LedgerKeyPaymentRequest {
	val, ok := u.GetPaymentRequest()

	if !ok {
		panic("arm PaymentRequest is not set")
	}

	return val
}

// GetPaymentRequest retrieves the PaymentRequest value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerKey) GetPaymentRequest() (result LedgerKeyPaymentRequest, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "PaymentRequest" {
		result = *u.PaymentRequest
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

// MustInvoice retrieves the Invoice value from the union,
// panicing if the value is not set.
func (u LedgerKey) MustInvoice() LedgerKeyInvoice {
	val, ok := u.GetInvoice()

	if !ok {
		panic("arm Invoice is not set")
	}

	return val
}

// GetInvoice retrieves the Invoice value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u LedgerKey) GetInvoice() (result LedgerKeyInvoice, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "Invoice" {
		result = *u.Invoice
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

// CreateAccountOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type CreateAccountOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
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
	}
	return "-", false
}

// NewCreateAccountOpExt creates a new  CreateAccountOpExt.
func NewCreateAccountOpExt(v LedgerVersion, value interface{}) (result CreateAccountOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// CreateAccountOp is an XDR Struct defines as:
//
//   struct CreateAccountOp
//    {
//        AccountID destination; // account to create
//        AccountID* referrer;     // parent account
//    	AccountType accountType;
//    	uint32 policies;
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
type CreateAccountOp struct {
	Destination AccountId          `json:"destination,omitempty"`
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
//    	INVALID_ACCOUNT_VERSION = -6 // if account version is higher than ledger version
//    };
//
type CreateAccountResultCode int32

const (
	CreateAccountResultCodeSuccess               CreateAccountResultCode = 0
	CreateAccountResultCodeMalformed             CreateAccountResultCode = -1
	CreateAccountResultCodeAccountTypeMismatched CreateAccountResultCode = -2
	CreateAccountResultCodeTypeNotAllowed        CreateAccountResultCode = -3
	CreateAccountResultCodeNameDuplication       CreateAccountResultCode = -4
	CreateAccountResultCodeReferrerNotFound      CreateAccountResultCode = -5
	CreateAccountResultCodeInvalidAccountVersion CreateAccountResultCode = -6
)

var CreateAccountResultCodeAll = []CreateAccountResultCode{
	CreateAccountResultCodeSuccess,
	CreateAccountResultCodeMalformed,
	CreateAccountResultCodeAccountTypeMismatched,
	CreateAccountResultCodeTypeNotAllowed,
	CreateAccountResultCodeNameDuplication,
	CreateAccountResultCodeReferrerNotFound,
	CreateAccountResultCodeInvalidAccountVersion,
}

var createAccountResultCodeMap = map[int32]string{
	0:  "CreateAccountResultCodeSuccess",
	-1: "CreateAccountResultCodeMalformed",
	-2: "CreateAccountResultCodeAccountTypeMismatched",
	-3: "CreateAccountResultCodeTypeNotAllowed",
	-4: "CreateAccountResultCodeNameDuplication",
	-5: "CreateAccountResultCodeReferrerNotFound",
	-6: "CreateAccountResultCodeInvalidAccountVersion",
}

var createAccountResultCodeShortMap = map[int32]string{
	0:  "success",
	-1: "malformed",
	-2: "account_type_mismatched",
	-3: "type_not_allowed",
	-4: "name_duplication",
	-5: "referrer_not_found",
	-6: "invalid_account_version",
}

var createAccountResultCodeRevMap = map[string]int32{
	"CreateAccountResultCodeSuccess":               0,
	"CreateAccountResultCodeMalformed":             -1,
	"CreateAccountResultCodeAccountTypeMismatched": -2,
	"CreateAccountResultCodeTypeNotAllowed":        -3,
	"CreateAccountResultCodeNameDuplication":       -4,
	"CreateAccountResultCodeReferrerNotFound":      -5,
	"CreateAccountResultCodeInvalidAccountVersion": -6,
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
//        }
//
type CreateIssuanceRequestOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
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
	}
	return "-", false
}

// NewCreateIssuanceRequestOpExt creates a new  CreateIssuanceRequestOpExt.
func NewCreateIssuanceRequestOpExt(v LedgerVersion, value interface{}) (result CreateIssuanceRequestOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
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
//    	FEE_EXCEEDS_AMOUNT = -8 // fee more than amount to issue
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
	CreateIssuanceRequestResultCodeFeeExceedsAmount         CreateIssuanceRequestResultCode = -8
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
	CreateIssuanceRequestResultCodeFeeExceedsAmount,
}

var createIssuanceRequestResultCodeMap = map[int32]string{
	0:  "CreateIssuanceRequestResultCodeSuccess",
	-1: "CreateIssuanceRequestResultCodeAssetNotFound",
	-2: "CreateIssuanceRequestResultCodeInvalidAmount",
	-3: "CreateIssuanceRequestResultCodeReferenceDuplication",
	-4: "CreateIssuanceRequestResultCodeNoCounterparty",
	-5: "CreateIssuanceRequestResultCodeNotAuthorized",
	-6: "CreateIssuanceRequestResultCodeExceedsMaxIssuanceAmount",
	-7: "CreateIssuanceRequestResultCodeReceiverFullLine",
	-8: "CreateIssuanceRequestResultCodeFeeExceedsAmount",
}

var createIssuanceRequestResultCodeShortMap = map[int32]string{
	0:  "success",
	-1: "asset_not_found",
	-2: "invalid_amount",
	-3: "reference_duplication",
	-4: "no_counterparty",
	-5: "not_authorized",
	-6: "exceeds_max_issuance_amount",
	-7: "receiver_full_line",
	-8: "fee_exceeds_amount",
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
	"CreateIssuanceRequestResultCodeFeeExceedsAmount":         -8,
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
//    	REQUEST_OR_SALE_ALREADY_EXISTS = -10
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
//        LIMITS_EXCEEDED = -13 // withdraw exceeds limits for source account
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
//    	CANCEL_ASSET_REQUEST = 2
//    };
//
type ManageAssetAction int32

const (
	ManageAssetActionCreateAssetCreationRequest ManageAssetAction = 0
	ManageAssetActionCreateAssetUpdateRequest   ManageAssetAction = 1
	ManageAssetActionCancelAssetRequest         ManageAssetAction = 2
)

var ManageAssetActionAll = []ManageAssetAction{
	ManageAssetActionCreateAssetCreationRequest,
	ManageAssetActionCreateAssetUpdateRequest,
	ManageAssetActionCancelAssetRequest,
}

var manageAssetActionMap = map[int32]string{
	0: "ManageAssetActionCreateAssetCreationRequest",
	1: "ManageAssetActionCreateAssetUpdateRequest",
	2: "ManageAssetActionCancelAssetRequest",
}

var manageAssetActionShortMap = map[int32]string{
	0: "create_asset_creation_request",
	1: "create_asset_update_request",
	2: "cancel_asset_request",
}

var manageAssetActionRevMap = map[string]int32{
	"ManageAssetActionCreateAssetCreationRequest": 0,
	"ManageAssetActionCreateAssetUpdateRequest":   1,
	"ManageAssetActionCancelAssetRequest":         2,
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
//    	}
//
type ManageAssetOpRequest struct {
	Action        ManageAssetAction     `json:"action,omitempty"`
	CreateAsset   *AssetCreationRequest `json:"createAsset,omitempty"`
	UpdateAsset   *AssetUpdateRequest   `json:"updateAsset,omitempty"`
	CancelRequest *CancelAssetRequest   `json:"cancelRequest,omitempty"`
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
//    	INITIAL_PREISSUED_EXCEEDS_MAX_ISSUANCE = -11 // initial pre issued amount exceeds max issuance amount
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
//        DELETE_BALANCE = 1
//    };
//
type ManageBalanceAction int32

const (
	ManageBalanceActionCreate        ManageBalanceAction = 0
	ManageBalanceActionDeleteBalance ManageBalanceAction = 1
)

var ManageBalanceActionAll = []ManageBalanceAction{
	ManageBalanceActionCreate,
	ManageBalanceActionDeleteBalance,
}

var manageBalanceActionMap = map[int32]string{
	0: "ManageBalanceActionCreate",
	1: "ManageBalanceActionDeleteBalance",
}

var manageBalanceActionShortMap = map[int32]string{
	0: "create",
	1: "delete_balance",
}

var manageBalanceActionRevMap = map[string]int32{
	"ManageBalanceActionCreate":        0,
	"ManageBalanceActionDeleteBalance": 1,
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
//        INVALID_ASSET = -5
//    };
//
type ManageBalanceResultCode int32

const (
	ManageBalanceResultCodeSuccess             ManageBalanceResultCode = 0
	ManageBalanceResultCodeMalformed           ManageBalanceResultCode = -1
	ManageBalanceResultCodeNotFound            ManageBalanceResultCode = -2
	ManageBalanceResultCodeDestinationNotFound ManageBalanceResultCode = -3
	ManageBalanceResultCodeAssetNotFound       ManageBalanceResultCode = -4
	ManageBalanceResultCodeInvalidAsset        ManageBalanceResultCode = -5
)

var ManageBalanceResultCodeAll = []ManageBalanceResultCode{
	ManageBalanceResultCodeSuccess,
	ManageBalanceResultCodeMalformed,
	ManageBalanceResultCodeNotFound,
	ManageBalanceResultCodeDestinationNotFound,
	ManageBalanceResultCodeAssetNotFound,
	ManageBalanceResultCodeInvalidAsset,
}

var manageBalanceResultCodeMap = map[int32]string{
	0:  "ManageBalanceResultCodeSuccess",
	-1: "ManageBalanceResultCodeMalformed",
	-2: "ManageBalanceResultCodeNotFound",
	-3: "ManageBalanceResultCodeDestinationNotFound",
	-4: "ManageBalanceResultCodeAssetNotFound",
	-5: "ManageBalanceResultCodeInvalidAsset",
}

var manageBalanceResultCodeShortMap = map[int32]string{
	0:  "success",
	-1: "malformed",
	-2: "not_found",
	-3: "destination_not_found",
	-4: "asset_not_found",
	-5: "invalid_asset",
}

var manageBalanceResultCodeRevMap = map[string]int32{
	"ManageBalanceResultCodeSuccess":             0,
	"ManageBalanceResultCodeMalformed":           -1,
	"ManageBalanceResultCodeNotFound":            -2,
	"ManageBalanceResultCodeDestinationNotFound": -3,
	"ManageBalanceResultCodeAssetNotFound":       -4,
	"ManageBalanceResultCodeInvalidAsset":        -5,
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

// ManageInvoiceOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type ManageInvoiceOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageInvoiceOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageInvoiceOpExt
func (u ManageInvoiceOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewManageInvoiceOpExt creates a new  ManageInvoiceOpExt.
func NewManageInvoiceOpExt(v LedgerVersion, value interface{}) (result ManageInvoiceOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ManageInvoiceOp is an XDR Struct defines as:
//
//   struct ManageInvoiceOp
//    {
//        BalanceID receiverBalance;
//    	AccountID sender;
//        int64 amount; // if set to 0, delete the invoice
//
//        // 0=create a new invoice, otherwise edit an existing invoice
//        uint64 invoiceID;
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
type ManageInvoiceOp struct {
	ReceiverBalance BalanceId          `json:"receiverBalance,omitempty"`
	Sender          AccountId          `json:"sender,omitempty"`
	Amount          Int64              `json:"amount,omitempty"`
	InvoiceId       Uint64             `json:"invoiceID,omitempty"`
	Ext             ManageInvoiceOpExt `json:"ext,omitempty"`
}

// ManageInvoiceResultCode is an XDR Enum defines as:
//
//   enum ManageInvoiceResultCode
//    {
//        // codes considered as "success" for the operation
//        SUCCESS = 0,
//
//        // codes considered as "failure" for the operation
//        MALFORMED = -1,
//        BALANCE_NOT_FOUND = -2,
//    	INVOICE_OVERFLOW = -3,
//
//        NOT_FOUND = -4,
//        TOO_MANY_INVOICES = -5,
//        CAN_NOT_DELETE_IN_PROGRESS = -6
//    };
//
type ManageInvoiceResultCode int32

const (
	ManageInvoiceResultCodeSuccess                ManageInvoiceResultCode = 0
	ManageInvoiceResultCodeMalformed              ManageInvoiceResultCode = -1
	ManageInvoiceResultCodeBalanceNotFound        ManageInvoiceResultCode = -2
	ManageInvoiceResultCodeInvoiceOverflow        ManageInvoiceResultCode = -3
	ManageInvoiceResultCodeNotFound               ManageInvoiceResultCode = -4
	ManageInvoiceResultCodeTooManyInvoices        ManageInvoiceResultCode = -5
	ManageInvoiceResultCodeCanNotDeleteInProgress ManageInvoiceResultCode = -6
)

var ManageInvoiceResultCodeAll = []ManageInvoiceResultCode{
	ManageInvoiceResultCodeSuccess,
	ManageInvoiceResultCodeMalformed,
	ManageInvoiceResultCodeBalanceNotFound,
	ManageInvoiceResultCodeInvoiceOverflow,
	ManageInvoiceResultCodeNotFound,
	ManageInvoiceResultCodeTooManyInvoices,
	ManageInvoiceResultCodeCanNotDeleteInProgress,
}

var manageInvoiceResultCodeMap = map[int32]string{
	0:  "ManageInvoiceResultCodeSuccess",
	-1: "ManageInvoiceResultCodeMalformed",
	-2: "ManageInvoiceResultCodeBalanceNotFound",
	-3: "ManageInvoiceResultCodeInvoiceOverflow",
	-4: "ManageInvoiceResultCodeNotFound",
	-5: "ManageInvoiceResultCodeTooManyInvoices",
	-6: "ManageInvoiceResultCodeCanNotDeleteInProgress",
}

var manageInvoiceResultCodeShortMap = map[int32]string{
	0:  "success",
	-1: "malformed",
	-2: "balance_not_found",
	-3: "invoice_overflow",
	-4: "not_found",
	-5: "too_many_invoices",
	-6: "can_not_delete_in_progress",
}

var manageInvoiceResultCodeRevMap = map[string]int32{
	"ManageInvoiceResultCodeSuccess":                0,
	"ManageInvoiceResultCodeMalformed":              -1,
	"ManageInvoiceResultCodeBalanceNotFound":        -2,
	"ManageInvoiceResultCodeInvoiceOverflow":        -3,
	"ManageInvoiceResultCodeNotFound":               -4,
	"ManageInvoiceResultCodeTooManyInvoices":        -5,
	"ManageInvoiceResultCodeCanNotDeleteInProgress": -6,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ManageInvoiceResultCode
func (e ManageInvoiceResultCode) ValidEnum(v int32) bool {
	_, ok := manageInvoiceResultCodeMap[v]
	return ok
}
func (e ManageInvoiceResultCode) isFlag() bool {
	for i := len(ManageInvoiceResultCodeAll) - 1; i >= 0; i-- {
		expected := ManageInvoiceResultCode(2) << uint64(len(ManageInvoiceResultCodeAll)-1) >> uint64(len(ManageInvoiceResultCodeAll)-i)
		if expected != ManageInvoiceResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ManageInvoiceResultCode) String() string {
	name, _ := manageInvoiceResultCodeMap[int32(e)]
	return name
}

func (e ManageInvoiceResultCode) ShortString() string {
	name, _ := manageInvoiceResultCodeShortMap[int32(e)]
	return name
}

func (e ManageInvoiceResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range ManageInvoiceResultCodeAll {
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

func (e *ManageInvoiceResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ManageInvoiceResultCode(t.Value)
	return nil
}

// ManageInvoiceSuccessResultExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type ManageInvoiceSuccessResultExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageInvoiceSuccessResultExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageInvoiceSuccessResultExt
func (u ManageInvoiceSuccessResultExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewManageInvoiceSuccessResultExt creates a new  ManageInvoiceSuccessResultExt.
func NewManageInvoiceSuccessResultExt(v LedgerVersion, value interface{}) (result ManageInvoiceSuccessResultExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ManageInvoiceSuccessResult is an XDR Struct defines as:
//
//   struct ManageInvoiceSuccessResult
//    {
//    	uint64 invoiceID;
//    	AssetCode asset;
//    	BalanceID senderBalance;
//
//    	union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type ManageInvoiceSuccessResult struct {
	InvoiceId     Uint64                        `json:"invoiceID,omitempty"`
	Asset         AssetCode                     `json:"asset,omitempty"`
	SenderBalance BalanceId                     `json:"senderBalance,omitempty"`
	Ext           ManageInvoiceSuccessResultExt `json:"ext,omitempty"`
}

// ManageInvoiceResult is an XDR Union defines as:
//
//   union ManageInvoiceResult switch (ManageInvoiceResultCode code)
//    {
//    case SUCCESS:
//        ManageInvoiceSuccessResult success;
//    default:
//        void;
//    };
//
type ManageInvoiceResult struct {
	Code    ManageInvoiceResultCode     `json:"code,omitempty"`
	Success *ManageInvoiceSuccessResult `json:"success,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ManageInvoiceResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ManageInvoiceResult
func (u ManageInvoiceResult) ArmForSwitch(sw int32) (string, bool) {
	switch ManageInvoiceResultCode(sw) {
	case ManageInvoiceResultCodeSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewManageInvoiceResult creates a new  ManageInvoiceResult.
func NewManageInvoiceResult(code ManageInvoiceResultCode, value interface{}) (result ManageInvoiceResult, err error) {
	result.Code = code
	switch ManageInvoiceResultCode(code) {
	case ManageInvoiceResultCodeSuccess:
		tv, ok := value.(ManageInvoiceSuccessResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageInvoiceSuccessResult")
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
func (u ManageInvoiceResult) MustSuccess() ManageInvoiceSuccessResult {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ManageInvoiceResult) GetSuccess() (result ManageInvoiceSuccessResult, ok bool) {
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
//    	INSUFFICIENT_PRICE = -12
//    };
//
type ManageOfferResultCode int32

const (
	ManageOfferResultCodeSuccess                  ManageOfferResultCode = 0
	ManageOfferResultCodeMalformed                ManageOfferResultCode = -1
	ManageOfferResultCodePairNotTraded            ManageOfferResultCode = -2
	ManageOfferResultCodeBalanceNotFound          ManageOfferResultCode = -3
	ManageOfferResultCodeUnderfunded              ManageOfferResultCode = -4
	ManageOfferResultCodeCrossSelf                ManageOfferResultCode = -5
	ManageOfferResultCodeOfferOverflow            ManageOfferResultCode = -6
	ManageOfferResultCodeAssetPairNotTradable     ManageOfferResultCode = -7
	ManageOfferResultCodePhysicalPriceRestriction ManageOfferResultCode = -8
	ManageOfferResultCodeCurrentPriceRestriction  ManageOfferResultCode = -9
	ManageOfferResultCodeNotFound                 ManageOfferResultCode = -10
	ManageOfferResultCodeInvalidPercentFee        ManageOfferResultCode = -11
	ManageOfferResultCodeInsufficientPrice        ManageOfferResultCode = -12
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
}

var manageOfferResultCodeRevMap = map[string]int32{
	"ManageOfferResultCodeSuccess":                  0,
	"ManageOfferResultCodeMalformed":                -1,
	"ManageOfferResultCodePairNotTraded":            -2,
	"ManageOfferResultCodeBalanceNotFound":          -3,
	"ManageOfferResultCodeUnderfunded":              -4,
	"ManageOfferResultCodeCrossSelf":                -5,
	"ManageOfferResultCodeOfferOverflow":            -6,
	"ManageOfferResultCodeAssetPairNotTradable":     -7,
	"ManageOfferResultCodePhysicalPriceRestriction": -8,
	"ManageOfferResultCodeCurrentPriceRestriction":  -9,
	"ManageOfferResultCodeNotFound":                 -10,
	"ManageOfferResultCodeInvalidPercentFee":        -11,
	"ManageOfferResultCodeInsufficientPrice":        -12,
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
//        INVOICE_ALREADY_PAID = -17
//    };
//
type PaymentResultCode int32

const (
	PaymentResultCodeSuccess                  PaymentResultCode = 0
	PaymentResultCodeMalformed                PaymentResultCode = -1
	PaymentResultCodeUnderfunded              PaymentResultCode = -2
	PaymentResultCodeLineFull                 PaymentResultCode = -3
	PaymentResultCodeFeeMismatched            PaymentResultCode = -4
	PaymentResultCodeBalanceNotFound          PaymentResultCode = -5
	PaymentResultCodeBalanceAccountMismatched PaymentResultCode = -6
	PaymentResultCodeBalanceAssetsMismatched  PaymentResultCode = -7
	PaymentResultCodeSrcBalanceNotFound       PaymentResultCode = -8
	PaymentResultCodeReferenceDuplication     PaymentResultCode = -9
	PaymentResultCodeStatsOverflow            PaymentResultCode = -10
	PaymentResultCodeLimitsExceeded           PaymentResultCode = -11
	PaymentResultCodeNotAllowedByAssetPolicy  PaymentResultCode = -12
	PaymentResultCodeInvoiceNotFound          PaymentResultCode = -13
	PaymentResultCodeInvoiceWrongAmount       PaymentResultCode = -14
	PaymentResultCodeInvoiceBalanceMismatch   PaymentResultCode = -15
	PaymentResultCodeInvoiceAccountMismatch   PaymentResultCode = -16
	PaymentResultCodeInvoiceAlreadyPaid       PaymentResultCode = -17
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
}

var paymentResultCodeRevMap = map[string]int32{
	"PaymentResultCodeSuccess":                  0,
	"PaymentResultCodeMalformed":                -1,
	"PaymentResultCodeUnderfunded":              -2,
	"PaymentResultCodeLineFull":                 -3,
	"PaymentResultCodeFeeMismatched":            -4,
	"PaymentResultCodeBalanceNotFound":          -5,
	"PaymentResultCodeBalanceAccountMismatched": -6,
	"PaymentResultCodeBalanceAssetsMismatched":  -7,
	"PaymentResultCodeSrcBalanceNotFound":       -8,
	"PaymentResultCodeReferenceDuplication":     -9,
	"PaymentResultCodeStatsOverflow":            -10,
	"PaymentResultCodeLimitsExceeded":           -11,
	"PaymentResultCodeNotAllowedByAssetPolicy":  -12,
	"PaymentResultCodeInvoiceNotFound":          -13,
	"PaymentResultCodeInvoiceWrongAmount":       -14,
	"PaymentResultCodeInvoiceBalanceMismatch":   -15,
	"PaymentResultCodeInvoiceAccountMismatch":   -16,
	"PaymentResultCodeInvoiceAlreadyPaid":       -17,
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

// RecoverOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type RecoverOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u RecoverOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of RecoverOpExt
func (u RecoverOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewRecoverOpExt creates a new  RecoverOpExt.
func NewRecoverOpExt(v LedgerVersion, value interface{}) (result RecoverOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// RecoverOp is an XDR Struct defines as:
//
//   struct RecoverOp
//    {
//        AccountID account;
//        PublicKey oldSigner;
//        PublicKey newSigner;
//    	// reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type RecoverOp struct {
	Account   AccountId    `json:"account,omitempty"`
	OldSigner PublicKey    `json:"oldSigner,omitempty"`
	NewSigner PublicKey    `json:"newSigner,omitempty"`
	Ext       RecoverOpExt `json:"ext,omitempty"`
}

// RecoverResultCode is an XDR Enum defines as:
//
//   enum RecoverResultCode
//    {
//        // codes considered as "success" for the operation
//        SUCCESS = 0,
//
//        // codes considered as "failure" for the operation
//
//        MALFORMED = -1,
//        OLD_SIGNER_NOT_FOUND = -2,
//        SIGNER_ALREADY_EXISTS = -3
//    };
//
type RecoverResultCode int32

const (
	RecoverResultCodeSuccess             RecoverResultCode = 0
	RecoverResultCodeMalformed           RecoverResultCode = -1
	RecoverResultCodeOldSignerNotFound   RecoverResultCode = -2
	RecoverResultCodeSignerAlreadyExists RecoverResultCode = -3
)

var RecoverResultCodeAll = []RecoverResultCode{
	RecoverResultCodeSuccess,
	RecoverResultCodeMalformed,
	RecoverResultCodeOldSignerNotFound,
	RecoverResultCodeSignerAlreadyExists,
}

var recoverResultCodeMap = map[int32]string{
	0:  "RecoverResultCodeSuccess",
	-1: "RecoverResultCodeMalformed",
	-2: "RecoverResultCodeOldSignerNotFound",
	-3: "RecoverResultCodeSignerAlreadyExists",
}

var recoverResultCodeShortMap = map[int32]string{
	0:  "success",
	-1: "malformed",
	-2: "old_signer_not_found",
	-3: "signer_already_exists",
}

var recoverResultCodeRevMap = map[string]int32{
	"RecoverResultCodeSuccess":             0,
	"RecoverResultCodeMalformed":           -1,
	"RecoverResultCodeOldSignerNotFound":   -2,
	"RecoverResultCodeSignerAlreadyExists": -3,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for RecoverResultCode
func (e RecoverResultCode) ValidEnum(v int32) bool {
	_, ok := recoverResultCodeMap[v]
	return ok
}
func (e RecoverResultCode) isFlag() bool {
	for i := len(RecoverResultCodeAll) - 1; i >= 0; i-- {
		expected := RecoverResultCode(2) << uint64(len(RecoverResultCodeAll)-1) >> uint64(len(RecoverResultCodeAll)-i)
		if expected != RecoverResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e RecoverResultCode) String() string {
	name, _ := recoverResultCodeMap[int32(e)]
	return name
}

func (e RecoverResultCode) ShortString() string {
	name, _ := recoverResultCodeShortMap[int32(e)]
	return name
}

func (e RecoverResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range RecoverResultCodeAll {
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

func (e *RecoverResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = RecoverResultCode(t.Value)
	return nil
}

// RecoverResultSuccessExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//
type RecoverResultSuccessExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u RecoverResultSuccessExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of RecoverResultSuccessExt
func (u RecoverResultSuccessExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewRecoverResultSuccessExt creates a new  RecoverResultSuccessExt.
func NewRecoverResultSuccessExt(v LedgerVersion, value interface{}) (result RecoverResultSuccessExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// RecoverResultSuccess is an XDR NestedStruct defines as:
//
//   struct {
//    		// reserved for future use
//    		union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//    		ext;
//    	}
//
type RecoverResultSuccess struct {
	Ext RecoverResultSuccessExt `json:"ext,omitempty"`
}

// RecoverResult is an XDR Union defines as:
//
//   union RecoverResult switch (RecoverResultCode code)
//    {
//    case SUCCESS:
//        struct {
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
type RecoverResult struct {
	Code    RecoverResultCode     `json:"code,omitempty"`
	Success *RecoverResultSuccess `json:"success,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u RecoverResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of RecoverResult
func (u RecoverResult) ArmForSwitch(sw int32) (string, bool) {
	switch RecoverResultCode(sw) {
	case RecoverResultCodeSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewRecoverResult creates a new  RecoverResult.
func NewRecoverResult(code RecoverResultCode, value interface{}) (result RecoverResult, err error) {
	result.Code = code
	switch RecoverResultCode(code) {
	case RecoverResultCodeSuccess:
		tv, ok := value.(RecoverResultSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be RecoverResultSuccess")
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
func (u RecoverResult) MustSuccess() RecoverResultSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u RecoverResult) GetSuccess() (result RecoverResultSuccess, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "Success" {
		result = *u.Success
		ok = true
	}

	return
}

// ReviewPaymentRequestOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//    	{
//    	case EMPTY_VERSION:
//    		void;
//    	}
//
type ReviewPaymentRequestOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ReviewPaymentRequestOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ReviewPaymentRequestOpExt
func (u ReviewPaymentRequestOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewReviewPaymentRequestOpExt creates a new  ReviewPaymentRequestOpExt.
func NewReviewPaymentRequestOpExt(v LedgerVersion, value interface{}) (result ReviewPaymentRequestOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ReviewPaymentRequestOp is an XDR Struct defines as:
//
//   struct ReviewPaymentRequestOp
//    {
//        uint64 paymentID;
//
//    	bool accept;
//        string256* rejectReason;
//    	// reserved for future use
//    	union switch (LedgerVersion v)
//    	{
//    	case EMPTY_VERSION:
//    		void;
//    	}
//    	ext;
//    };
//
type ReviewPaymentRequestOp struct {
	PaymentId    Uint64                    `json:"paymentID,omitempty"`
	Accept       bool                      `json:"accept,omitempty"`
	RejectReason *String256                `json:"rejectReason,omitempty"`
	Ext          ReviewPaymentRequestOpExt `json:"ext,omitempty"`
}

// ReviewPaymentRequestResultCode is an XDR Enum defines as:
//
//   enum ReviewPaymentRequestResultCode
//    {
//        // codes considered as "success" for the operation
//        SUCCESS = 0,
//
//        // codes considered as "failure" for the operation
//    	NOT_FOUND = -1,           // failed to find Recovery request with such ID
//        LINE_FULL = -2
//    };
//
type ReviewPaymentRequestResultCode int32

const (
	ReviewPaymentRequestResultCodeSuccess  ReviewPaymentRequestResultCode = 0
	ReviewPaymentRequestResultCodeNotFound ReviewPaymentRequestResultCode = -1
	ReviewPaymentRequestResultCodeLineFull ReviewPaymentRequestResultCode = -2
)

var ReviewPaymentRequestResultCodeAll = []ReviewPaymentRequestResultCode{
	ReviewPaymentRequestResultCodeSuccess,
	ReviewPaymentRequestResultCodeNotFound,
	ReviewPaymentRequestResultCodeLineFull,
}

var reviewPaymentRequestResultCodeMap = map[int32]string{
	0:  "ReviewPaymentRequestResultCodeSuccess",
	-1: "ReviewPaymentRequestResultCodeNotFound",
	-2: "ReviewPaymentRequestResultCodeLineFull",
}

var reviewPaymentRequestResultCodeShortMap = map[int32]string{
	0:  "success",
	-1: "not_found",
	-2: "line_full",
}

var reviewPaymentRequestResultCodeRevMap = map[string]int32{
	"ReviewPaymentRequestResultCodeSuccess":  0,
	"ReviewPaymentRequestResultCodeNotFound": -1,
	"ReviewPaymentRequestResultCodeLineFull": -2,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for ReviewPaymentRequestResultCode
func (e ReviewPaymentRequestResultCode) ValidEnum(v int32) bool {
	_, ok := reviewPaymentRequestResultCodeMap[v]
	return ok
}
func (e ReviewPaymentRequestResultCode) isFlag() bool {
	for i := len(ReviewPaymentRequestResultCodeAll) - 1; i >= 0; i-- {
		expected := ReviewPaymentRequestResultCode(2) << uint64(len(ReviewPaymentRequestResultCodeAll)-1) >> uint64(len(ReviewPaymentRequestResultCodeAll)-i)
		if expected != ReviewPaymentRequestResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e ReviewPaymentRequestResultCode) String() string {
	name, _ := reviewPaymentRequestResultCodeMap[int32(e)]
	return name
}

func (e ReviewPaymentRequestResultCode) ShortString() string {
	name, _ := reviewPaymentRequestResultCodeShortMap[int32(e)]
	return name
}

func (e ReviewPaymentRequestResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range ReviewPaymentRequestResultCodeAll {
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

func (e *ReviewPaymentRequestResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = ReviewPaymentRequestResultCode(t.Value)
	return nil
}

// PaymentState is an XDR Enum defines as:
//
//   enum PaymentState
//    {
//        PENDING = 0,
//        PROCESSED = 1,
//        REJECTED = 2
//    };
//
type PaymentState int32

const (
	PaymentStatePending   PaymentState = 0
	PaymentStateProcessed PaymentState = 1
	PaymentStateRejected  PaymentState = 2
)

var PaymentStateAll = []PaymentState{
	PaymentStatePending,
	PaymentStateProcessed,
	PaymentStateRejected,
}

var paymentStateMap = map[int32]string{
	0: "PaymentStatePending",
	1: "PaymentStateProcessed",
	2: "PaymentStateRejected",
}

var paymentStateShortMap = map[int32]string{
	0: "pending",
	1: "processed",
	2: "rejected",
}

var paymentStateRevMap = map[string]int32{
	"PaymentStatePending":   0,
	"PaymentStateProcessed": 1,
	"PaymentStateRejected":  2,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for PaymentState
func (e PaymentState) ValidEnum(v int32) bool {
	_, ok := paymentStateMap[v]
	return ok
}
func (e PaymentState) isFlag() bool {
	for i := len(PaymentStateAll) - 1; i >= 0; i-- {
		expected := PaymentState(2) << uint64(len(PaymentStateAll)-1) >> uint64(len(PaymentStateAll)-i)
		if expected != PaymentStateAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e PaymentState) String() string {
	name, _ := paymentStateMap[int32(e)]
	return name
}

func (e PaymentState) ShortString() string {
	name, _ := paymentStateShortMap[int32(e)]
	return name
}

func (e PaymentState) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range PaymentStateAll {
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

func (e *PaymentState) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = PaymentState(t.Value)
	return nil
}

// ReviewPaymentResponseExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//    	{
//    	case EMPTY_VERSION:
//    		void;
//    	}
//
type ReviewPaymentResponseExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ReviewPaymentResponseExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ReviewPaymentResponseExt
func (u ReviewPaymentResponseExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewReviewPaymentResponseExt creates a new  ReviewPaymentResponseExt.
func NewReviewPaymentResponseExt(v LedgerVersion, value interface{}) (result ReviewPaymentResponseExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ReviewPaymentResponse is an XDR Struct defines as:
//
//   struct ReviewPaymentResponse {
//        PaymentState state;
//
//        uint64* relatedInvoiceID;
//    	// reserved for future use
//    	union switch (LedgerVersion v)
//    	{
//    	case EMPTY_VERSION:
//    		void;
//    	}
//    	ext;
//    };
//
type ReviewPaymentResponse struct {
	State            PaymentState             `json:"state,omitempty"`
	RelatedInvoiceId *Uint64                  `json:"relatedInvoiceID,omitempty"`
	Ext              ReviewPaymentResponseExt `json:"ext,omitempty"`
}

// ReviewPaymentRequestResult is an XDR Union defines as:
//
//   union ReviewPaymentRequestResult switch (ReviewPaymentRequestResultCode code)
//    {
//    case SUCCESS:
//        ReviewPaymentResponse reviewPaymentResponse;
//    default:
//        void;
//    };
//
type ReviewPaymentRequestResult struct {
	Code                  ReviewPaymentRequestResultCode `json:"code,omitempty"`
	ReviewPaymentResponse *ReviewPaymentResponse         `json:"reviewPaymentResponse,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u ReviewPaymentRequestResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of ReviewPaymentRequestResult
func (u ReviewPaymentRequestResult) ArmForSwitch(sw int32) (string, bool) {
	switch ReviewPaymentRequestResultCode(sw) {
	case ReviewPaymentRequestResultCodeSuccess:
		return "ReviewPaymentResponse", true
	default:
		return "", true
	}
}

// NewReviewPaymentRequestResult creates a new  ReviewPaymentRequestResult.
func NewReviewPaymentRequestResult(code ReviewPaymentRequestResultCode, value interface{}) (result ReviewPaymentRequestResult, err error) {
	result.Code = code
	switch ReviewPaymentRequestResultCode(code) {
	case ReviewPaymentRequestResultCodeSuccess:
		tv, ok := value.(ReviewPaymentResponse)
		if !ok {
			err = fmt.Errorf("invalid value, must be ReviewPaymentResponse")
			return
		}
		result.ReviewPaymentResponse = &tv
	default:
		// void
	}
	return
}

// MustReviewPaymentResponse retrieves the ReviewPaymentResponse value from the union,
// panicing if the value is not set.
func (u ReviewPaymentRequestResult) MustReviewPaymentResponse() ReviewPaymentResponse {
	val, ok := u.GetReviewPaymentResponse()

	if !ok {
		panic("arm ReviewPaymentResponse is not set")
	}

	return val
}

// GetReviewPaymentResponse retrieves the ReviewPaymentResponse value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u ReviewPaymentRequestResult) GetReviewPaymentResponse() (result ReviewPaymentResponse, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Code))

	if armName == "ReviewPaymentResponse" {
		result = *u.ReviewPaymentResponse
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

// ReviewRequestOpRequestDetails is an XDR NestedUnion defines as:
//
//   union switch(ReviewableRequestType requestType) {
//    	case WITHDRAW:
//    		WithdrawalDetails withdrawal;
//    	default:
//    		void;
//    	}
//
type ReviewRequestOpRequestDetails struct {
	RequestType ReviewableRequestType `json:"requestType,omitempty"`
	Withdrawal  *WithdrawalDetails    `json:"withdrawal,omitempty"`
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

// ReviewRequestOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type ReviewRequestOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
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
	}
	return "-", false
}

// NewReviewRequestOpExt creates a new  ReviewRequestOpExt.
func NewReviewRequestOpExt(v LedgerVersion, value interface{}) (result ReviewRequestOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
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
//    	default:
//    		void;
//    	} requestDetails;
//    	ReviewRequestOpAction action;
//    	string256 reason;
//    	// reserved for future use
//        union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type ReviewRequestOp struct {
	RequestId      Uint64                        `json:"requestID,omitempty"`
	RequestHash    Hash                          `json:"requestHash,omitempty"`
	RequestDetails ReviewRequestOpRequestDetails `json:"requestDetails,omitempty"`
	Action         ReviewRequestOpAction         `json:"action,omitempty"`
	Reason         String256                     `json:"reason,omitempty"`
	Ext            ReviewRequestOpExt            `json:"ext,omitempty"`
}

// ReviewRequestResultCode is an XDR Enum defines as:
//
//   enum ReviewRequestResultCode
//    {
//        // codes considered as "success" for the operation
//        SUCCESS = 0,
//
//        // codes considered as "failure" for the operation
//        INVALID_REASON = -1,        // reason must be empty if approving and not empty if rejecting
//    	INVALID_ACTION = -2,
//    	HASH_MISMATCHED = -3,
//    	NOT_FOUND = -4,
//    	TYPE_MISMATCHED = -5,
//    	REJECT_NOT_ALLOWED = -6, // reject not allowed, use permanent reject
//
//    	// Asset requests
//    	ASSET_ALREADY_EXISTS = -20,
//    	ASSET_DOES_NOT_EXISTS = -21,
//
//    	// Issuance requests
//    	MAX_ISSUANCE_AMOUNT_EXCEEDED = -40,
//    	INSUFFICIENT_AVAILABLE_FOR_ISSUANCE_AMOUNT = -41,
//    	FULL_LINE = -42, // can't fund balance - total funds exceed UINT64_MAX
//
//    	// sale creation reuqests
//    	QUOTE_ASSET_DOES_NOT_EXISTS = -50,
//    	BASE_ASSET_DOES_NOT_EXISTS = -51,
//    	SOFT_CAP_WILL_EXCEED_MAX_ISSUANCE = -52
//
//    };
//
type ReviewRequestResultCode int32

const (
	ReviewRequestResultCodeSuccess                                ReviewRequestResultCode = 0
	ReviewRequestResultCodeInvalidReason                          ReviewRequestResultCode = -1
	ReviewRequestResultCodeInvalidAction                          ReviewRequestResultCode = -2
	ReviewRequestResultCodeHashMismatched                         ReviewRequestResultCode = -3
	ReviewRequestResultCodeNotFound                               ReviewRequestResultCode = -4
	ReviewRequestResultCodeTypeMismatched                         ReviewRequestResultCode = -5
	ReviewRequestResultCodeRejectNotAllowed                       ReviewRequestResultCode = -6
	ReviewRequestResultCodeAssetAlreadyExists                     ReviewRequestResultCode = -20
	ReviewRequestResultCodeAssetDoesNotExists                     ReviewRequestResultCode = -21
	ReviewRequestResultCodeMaxIssuanceAmountExceeded              ReviewRequestResultCode = -40
	ReviewRequestResultCodeInsufficientAvailableForIssuanceAmount ReviewRequestResultCode = -41
	ReviewRequestResultCodeFullLine                               ReviewRequestResultCode = -42
	ReviewRequestResultCodeQuoteAssetDoesNotExists                ReviewRequestResultCode = -50
	ReviewRequestResultCodeBaseAssetDoesNotExists                 ReviewRequestResultCode = -51
	ReviewRequestResultCodeSoftCapWillExceedMaxIssuance           ReviewRequestResultCode = -52
)

var ReviewRequestResultCodeAll = []ReviewRequestResultCode{
	ReviewRequestResultCodeSuccess,
	ReviewRequestResultCodeInvalidReason,
	ReviewRequestResultCodeInvalidAction,
	ReviewRequestResultCodeHashMismatched,
	ReviewRequestResultCodeNotFound,
	ReviewRequestResultCodeTypeMismatched,
	ReviewRequestResultCodeRejectNotAllowed,
	ReviewRequestResultCodeAssetAlreadyExists,
	ReviewRequestResultCodeAssetDoesNotExists,
	ReviewRequestResultCodeMaxIssuanceAmountExceeded,
	ReviewRequestResultCodeInsufficientAvailableForIssuanceAmount,
	ReviewRequestResultCodeFullLine,
	ReviewRequestResultCodeQuoteAssetDoesNotExists,
	ReviewRequestResultCodeBaseAssetDoesNotExists,
	ReviewRequestResultCodeSoftCapWillExceedMaxIssuance,
}

var reviewRequestResultCodeMap = map[int32]string{
	0:   "ReviewRequestResultCodeSuccess",
	-1:  "ReviewRequestResultCodeInvalidReason",
	-2:  "ReviewRequestResultCodeInvalidAction",
	-3:  "ReviewRequestResultCodeHashMismatched",
	-4:  "ReviewRequestResultCodeNotFound",
	-5:  "ReviewRequestResultCodeTypeMismatched",
	-6:  "ReviewRequestResultCodeRejectNotAllowed",
	-20: "ReviewRequestResultCodeAssetAlreadyExists",
	-21: "ReviewRequestResultCodeAssetDoesNotExists",
	-40: "ReviewRequestResultCodeMaxIssuanceAmountExceeded",
	-41: "ReviewRequestResultCodeInsufficientAvailableForIssuanceAmount",
	-42: "ReviewRequestResultCodeFullLine",
	-50: "ReviewRequestResultCodeQuoteAssetDoesNotExists",
	-51: "ReviewRequestResultCodeBaseAssetDoesNotExists",
	-52: "ReviewRequestResultCodeSoftCapWillExceedMaxIssuance",
}

var reviewRequestResultCodeShortMap = map[int32]string{
	0:   "success",
	-1:  "invalid_reason",
	-2:  "invalid_action",
	-3:  "hash_mismatched",
	-4:  "not_found",
	-5:  "type_mismatched",
	-6:  "reject_not_allowed",
	-20: "asset_already_exists",
	-21: "asset_does_not_exists",
	-40: "max_issuance_amount_exceeded",
	-41: "insufficient_available_for_issuance_amount",
	-42: "full_line",
	-50: "quote_asset_does_not_exists",
	-51: "base_asset_does_not_exists",
	-52: "soft_cap_will_exceed_max_issuance",
}

var reviewRequestResultCodeRevMap = map[string]int32{
	"ReviewRequestResultCodeSuccess":                                0,
	"ReviewRequestResultCodeInvalidReason":                          -1,
	"ReviewRequestResultCodeInvalidAction":                          -2,
	"ReviewRequestResultCodeHashMismatched":                         -3,
	"ReviewRequestResultCodeNotFound":                               -4,
	"ReviewRequestResultCodeTypeMismatched":                         -5,
	"ReviewRequestResultCodeRejectNotAllowed":                       -6,
	"ReviewRequestResultCodeAssetAlreadyExists":                     -20,
	"ReviewRequestResultCodeAssetDoesNotExists":                     -21,
	"ReviewRequestResultCodeMaxIssuanceAmountExceeded":              -40,
	"ReviewRequestResultCodeInsufficientAvailableForIssuanceAmount": -41,
	"ReviewRequestResultCodeFullLine":                               -42,
	"ReviewRequestResultCodeQuoteAssetDoesNotExists":                -50,
	"ReviewRequestResultCodeBaseAssetDoesNotExists":                 -51,
	"ReviewRequestResultCodeSoftCapWillExceedMaxIssuance":           -52,
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
//    		case EMPTY_VERSION:
//    			void;
//    		}
//
type ReviewRequestResultSuccessExt struct {
	V LedgerVersion `json:"v,omitempty"`
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
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewReviewRequestResultSuccessExt creates a new  ReviewRequestResultSuccessExt.
func NewReviewRequestResultSuccessExt(v LedgerVersion, value interface{}) (result ReviewRequestResultSuccessExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// ReviewRequestResultSuccess is an XDR NestedStruct defines as:
//
//   struct {
//    		// reserved for future use
//    		union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
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
//    		case EMPTY_VERSION:
//    			void;
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
//    		SUB_TYPE_NOT_EXIST = -9
//        };
//
type SetFeesResultCode int32

const (
	SetFeesResultCodeSuccess         SetFeesResultCode = 0
	SetFeesResultCodeInvalidAmount   SetFeesResultCode = -1
	SetFeesResultCodeInvalidFeeType  SetFeesResultCode = -2
	SetFeesResultCodeAssetNotFound   SetFeesResultCode = -3
	SetFeesResultCodeInvalidAsset    SetFeesResultCode = -4
	SetFeesResultCodeMalformed       SetFeesResultCode = -5
	SetFeesResultCodeMalformedRange  SetFeesResultCode = -6
	SetFeesResultCodeRangeOverlap    SetFeesResultCode = -7
	SetFeesResultCodeNotFound        SetFeesResultCode = -8
	SetFeesResultCodeSubTypeNotExist SetFeesResultCode = -9
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
}

var setFeesResultCodeMap = map[int32]string{
	0:  "SetFeesResultCodeSuccess",
	-1: "SetFeesResultCodeInvalidAmount",
	-2: "SetFeesResultCodeInvalidFeeType",
	-3: "SetFeesResultCodeAssetNotFound",
	-4: "SetFeesResultCodeInvalidAsset",
	-5: "SetFeesResultCodeMalformed",
	-6: "SetFeesResultCodeMalformedRange",
	-7: "SetFeesResultCodeRangeOverlap",
	-8: "SetFeesResultCodeNotFound",
	-9: "SetFeesResultCodeSubTypeNotExist",
}

var setFeesResultCodeShortMap = map[int32]string{
	0:  "success",
	-1: "invalid_amount",
	-2: "invalid_fee_type",
	-3: "asset_not_found",
	-4: "invalid_asset",
	-5: "malformed",
	-6: "malformed_range",
	-7: "range_overlap",
	-8: "not_found",
	-9: "sub_type_not_exist",
}

var setFeesResultCodeRevMap = map[string]int32{
	"SetFeesResultCodeSuccess":         0,
	"SetFeesResultCodeInvalidAmount":   -1,
	"SetFeesResultCodeInvalidFeeType":  -2,
	"SetFeesResultCodeAssetNotFound":   -3,
	"SetFeesResultCodeInvalidAsset":    -4,
	"SetFeesResultCodeMalformed":       -5,
	"SetFeesResultCodeMalformedRange":  -6,
	"SetFeesResultCodeRangeOverlap":    -7,
	"SetFeesResultCodeNotFound":        -8,
	"SetFeesResultCodeSubTypeNotExist": -9,
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

// SetLimitsOpExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//    	{
//    	case EMPTY_VERSION:
//    		void;
//    	}
//
type SetLimitsOpExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u SetLimitsOpExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of SetLimitsOpExt
func (u SetLimitsOpExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewSetLimitsOpExt creates a new  SetLimitsOpExt.
func NewSetLimitsOpExt(v LedgerVersion, value interface{}) (result SetLimitsOpExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// SetLimitsOp is an XDR Struct defines as:
//
//   struct SetLimitsOp
//    {
//        AccountID* account;
//        AccountType* accountType;
//
//        Limits limits;
//    	// reserved for future use
//    	union switch (LedgerVersion v)
//    	{
//    	case EMPTY_VERSION:
//    		void;
//    	}
//    	ext;
//    };
//
type SetLimitsOp struct {
	Account     *AccountId     `json:"account,omitempty"`
	AccountType *AccountType   `json:"accountType,omitempty"`
	Limits      Limits         `json:"limits,omitempty"`
	Ext         SetLimitsOpExt `json:"ext,omitempty"`
}

// SetLimitsResultCode is an XDR Enum defines as:
//
//   enum SetLimitsResultCode
//    {
//        // codes considered as "success" for the operation
//        SUCCESS = 0,
//        // codes considered as "failure" for the operation
//        MALFORMED = -1
//    };
//
type SetLimitsResultCode int32

const (
	SetLimitsResultCodeSuccess   SetLimitsResultCode = 0
	SetLimitsResultCodeMalformed SetLimitsResultCode = -1
)

var SetLimitsResultCodeAll = []SetLimitsResultCode{
	SetLimitsResultCodeSuccess,
	SetLimitsResultCodeMalformed,
}

var setLimitsResultCodeMap = map[int32]string{
	0:  "SetLimitsResultCodeSuccess",
	-1: "SetLimitsResultCodeMalformed",
}

var setLimitsResultCodeShortMap = map[int32]string{
	0:  "success",
	-1: "malformed",
}

var setLimitsResultCodeRevMap = map[string]int32{
	"SetLimitsResultCodeSuccess":   0,
	"SetLimitsResultCodeMalformed": -1,
}

// ValidEnum validates a proposed value for this enum.  Implements
// the Enum interface for SetLimitsResultCode
func (e SetLimitsResultCode) ValidEnum(v int32) bool {
	_, ok := setLimitsResultCodeMap[v]
	return ok
}
func (e SetLimitsResultCode) isFlag() bool {
	for i := len(SetLimitsResultCodeAll) - 1; i >= 0; i-- {
		expected := SetLimitsResultCode(2) << uint64(len(SetLimitsResultCodeAll)-1) >> uint64(len(SetLimitsResultCodeAll)-i)
		if expected != SetLimitsResultCodeAll[i] {
			return false
		}
	}
	return true
}

// String returns the name of `e`
func (e SetLimitsResultCode) String() string {
	name, _ := setLimitsResultCodeMap[int32(e)]
	return name
}

func (e SetLimitsResultCode) ShortString() string {
	name, _ := setLimitsResultCodeShortMap[int32(e)]
	return name
}

func (e SetLimitsResultCode) MarshalJSON() ([]byte, error) {
	if e.isFlag() {
		// marshal as mask
		result := flag{
			Value: int32(e),
		}
		for _, value := range SetLimitsResultCodeAll {
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

func (e *SetLimitsResultCode) UnmarshalJSON(data []byte) error {
	var t value
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*e = SetLimitsResultCode(t.Value)
	return nil
}

// SetLimitsResultSuccessExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//
type SetLimitsResultSuccessExt struct {
	V LedgerVersion `json:"v,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u SetLimitsResultSuccessExt) SwitchFieldName() string {
	return "V"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of SetLimitsResultSuccessExt
func (u SetLimitsResultSuccessExt) ArmForSwitch(sw int32) (string, bool) {
	switch LedgerVersion(sw) {
	case LedgerVersionEmptyVersion:
		return "", true
	}
	return "-", false
}

// NewSetLimitsResultSuccessExt creates a new  SetLimitsResultSuccessExt.
func NewSetLimitsResultSuccessExt(v LedgerVersion, value interface{}) (result SetLimitsResultSuccessExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// SetLimitsResultSuccess is an XDR NestedStruct defines as:
//
//   struct {
//    		// reserved for future use
//    		union switch (LedgerVersion v)
//    		{
//    		case EMPTY_VERSION:
//    			void;
//    		}
//    		ext;
//    	}
//
type SetLimitsResultSuccess struct {
	Ext SetLimitsResultSuccessExt `json:"ext,omitempty"`
}

// SetLimitsResult is an XDR Union defines as:
//
//   union SetLimitsResult switch (SetLimitsResultCode code)
//    {
//    case SUCCESS:
//        struct {
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
type SetLimitsResult struct {
	Code    SetLimitsResultCode     `json:"code,omitempty"`
	Success *SetLimitsResultSuccess `json:"success,omitempty"`
}

// SwitchFieldName returns the field name in which this union's
// discriminant is stored
func (u SetLimitsResult) SwitchFieldName() string {
	return "Code"
}

// ArmForSwitch returns which field name should be used for storing
// the value for an instance of SetLimitsResult
func (u SetLimitsResult) ArmForSwitch(sw int32) (string, bool) {
	switch SetLimitsResultCode(sw) {
	case SetLimitsResultCodeSuccess:
		return "Success", true
	default:
		return "", true
	}
}

// NewSetLimitsResult creates a new  SetLimitsResult.
func NewSetLimitsResult(code SetLimitsResultCode, value interface{}) (result SetLimitsResult, err error) {
	result.Code = code
	switch SetLimitsResultCode(code) {
	case SetLimitsResultCodeSuccess:
		tv, ok := value.(SetLimitsResultSuccess)
		if !ok {
			err = fmt.Errorf("invalid value, must be SetLimitsResultSuccess")
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
func (u SetLimitsResult) MustSuccess() SetLimitsResultSuccess {
	val, ok := u.GetSuccess()

	if !ok {
		panic("arm Success is not set")
	}

	return val
}

// GetSuccess retrieves the Success value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u SetLimitsResult) GetSuccess() (result SetLimitsResultSuccess, ok bool) {
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
	MasterWeight  *Uint32         `json:"masterWeight,omitempty"`
	LowThreshold  *Uint32         `json:"lowThreshold,omitempty"`
	MedThreshold  *Uint32         `json:"medThreshold,omitempty"`
	HighThreshold *Uint32         `json:"highThreshold,omitempty"`
	Signer        *Signer         `json:"signer,omitempty"`
	TrustData     *TrustData      `json:"trustData,omitempty"`
	Ext           SetOptionsOpExt `json:"ext,omitempty"`
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
//    	INVALID_SIGNER_VERSION = -7 // if signer version is higher than ledger version
//    };
//
type SetOptionsResultCode int32

const (
	SetOptionsResultCodeSuccess              SetOptionsResultCode = 0
	SetOptionsResultCodeTooManySigners       SetOptionsResultCode = -1
	SetOptionsResultCodeThresholdOutOfRange  SetOptionsResultCode = -2
	SetOptionsResultCodeBadSigner            SetOptionsResultCode = -3
	SetOptionsResultCodeBalanceNotFound      SetOptionsResultCode = -4
	SetOptionsResultCodeTrustMalformed       SetOptionsResultCode = -5
	SetOptionsResultCodeTrustTooMany         SetOptionsResultCode = -6
	SetOptionsResultCodeInvalidSignerVersion SetOptionsResultCode = -7
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
}

var setOptionsResultCodeRevMap = map[string]int32{
	"SetOptionsResultCodeSuccess":              0,
	"SetOptionsResultCodeTooManySigners":       -1,
	"SetOptionsResultCodeThresholdOutOfRange":  -2,
	"SetOptionsResultCodeBadSigner":            -3,
	"SetOptionsResultCodeBalanceNotFound":      -4,
	"SetOptionsResultCodeTrustMalformed":       -5,
	"SetOptionsResultCodeTrustTooMany":         -6,
	"SetOptionsResultCodeInvalidSignerVersion": -7,
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
	Ext SetOptionsResultSuccessExt `json:"ext,omitempty"`
}

// SetOptionsResult is an XDR Union defines as:
//
//   union SetOptionsResult switch (SetOptionsResultCode code)
//    {
//    case SUCCESS:
//        struct {
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
//
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
	Asset    AssetCode          `json:"asset,omitempty"`
	Amount   Uint64             `json:"amount,omitempty"`
	Receiver BalanceId          `json:"receiver,omitempty"`
	Fee      Fee                `json:"fee,omitempty"`
	Ext      IssuanceRequestExt `json:"ext,omitempty"`
}

// SaleCreationRequestExt is an XDR NestedUnion defines as:
//
//   union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//
type SaleCreationRequestExt struct {
	V LedgerVersion `json:"v,omitempty"`
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
	}
	return "-", false
}

// NewSaleCreationRequestExt creates a new  SaleCreationRequestExt.
func NewSaleCreationRequestExt(v LedgerVersion, value interface{}) (result SaleCreationRequestExt, err error) {
	result.V = v
	switch LedgerVersion(v) {
	case LedgerVersionEmptyVersion:
		// void
	}
	return
}

// SaleCreationRequest is an XDR Struct defines as:
//
//   struct SaleCreationRequest {
//    	AssetCode baseAsset; // asset for which sale will be performed
//    	AssetCode quoteAsset; // asset in which participation will be accepted
//    	uint64 startTime; // start time of the sale
//    	uint64 endTime; // close time of the sale
//    	uint64 price; // price for 1 baseAsset in terms of quote asset
//    	uint64 softCap; // minimum amount of quote asset to be received at which sale will be considered a successful
//    	uint64 hardCap; // max amount of quote asset to be received
//    	longstring details; // sale specific details
//
//    	union switch (LedgerVersion v)
//        {
//        case EMPTY_VERSION:
//            void;
//        }
//        ext;
//    };
//
type SaleCreationRequest struct {
	BaseAsset  AssetCode              `json:"baseAsset,omitempty"`
	QuoteAsset AssetCode              `json:"quoteAsset,omitempty"`
	StartTime  Uint64                 `json:"startTime,omitempty"`
	EndTime    Uint64                 `json:"endTime,omitempty"`
	Price      Uint64                 `json:"price,omitempty"`
	SoftCap    Uint64                 `json:"softCap,omitempty"`
	HardCap    Uint64                 `json:"hardCap,omitempty"`
	Details    Longstring             `json:"details,omitempty"`
	Ext        SaleCreationRequestExt `json:"ext,omitempty"`
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
	Balance         BalanceId                `json:"balance,omitempty"`
	Amount          Uint64                   `json:"amount,omitempty"`
	UniversalAmount Uint64                   `json:"universalAmount,omitempty"`
	Fee             Fee                      `json:"fee,omitempty"`
	ExternalDetails Longstring               `json:"externalDetails,omitempty"`
	Details         WithdrawalRequestDetails `json:"details,omitempty"`
	Ext             WithdrawalRequestExt     `json:"ext,omitempty"`
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
//    	case RECOVER:
//    		RecoverOp recoverOp;
//    	case MANAGE_BALANCE:
//    		ManageBalanceOp manageBalanceOp;
//    	case REVIEW_PAYMENT_REQUEST:
//    		ReviewPaymentRequestOp reviewPaymentRequestOp;
//        case MANAGE_ASSET:
//            ManageAssetOp manageAssetOp;
//        case CREATE_PREISSUANCE_REQUEST:
//            CreatePreIssuanceRequestOp createPreIssuanceRequest;
//        case SET_LIMITS:
//            SetLimitsOp setLimitsOp;
//        case DIRECT_DEBIT:
//            DirectDebitOp directDebitOp;
//    	case MANAGE_ASSET_PAIR:
//    		ManageAssetPairOp manageAssetPairOp;
//    	case MANAGE_OFFER:
//    		ManageOfferOp manageOfferOp;
//        case MANAGE_INVOICE:
//            ManageInvoiceOp manageInvoiceOp;
//    	case REVIEW_REQUEST:
//    		ReviewRequestOp reviewRequestOp;
//    	case CREATE_SALE_REQUEST:
//    		CreateSaleCreationRequestOp createSaleCreationRequestOp;
//        }
//
type OperationBody struct {
	Type                        OperationType                `json:"type,omitempty"`
	CreateAccountOp             *CreateAccountOp             `json:"createAccountOp,omitempty"`
	PaymentOp                   *PaymentOp                   `json:"paymentOp,omitempty"`
	SetOptionsOp                *SetOptionsOp                `json:"setOptionsOp,omitempty"`
	CreateIssuanceRequestOp     *CreateIssuanceRequestOp     `json:"createIssuanceRequestOp,omitempty"`
	SetFeesOp                   *SetFeesOp                   `json:"setFeesOp,omitempty"`
	ManageAccountOp             *ManageAccountOp             `json:"manageAccountOp,omitempty"`
	CreateWithdrawalRequestOp   *CreateWithdrawalRequestOp   `json:"createWithdrawalRequestOp,omitempty"`
	RecoverOp                   *RecoverOp                   `json:"recoverOp,omitempty"`
	ManageBalanceOp             *ManageBalanceOp             `json:"manageBalanceOp,omitempty"`
	ReviewPaymentRequestOp      *ReviewPaymentRequestOp      `json:"reviewPaymentRequestOp,omitempty"`
	ManageAssetOp               *ManageAssetOp               `json:"manageAssetOp,omitempty"`
	CreatePreIssuanceRequest    *CreatePreIssuanceRequestOp  `json:"createPreIssuanceRequest,omitempty"`
	SetLimitsOp                 *SetLimitsOp                 `json:"setLimitsOp,omitempty"`
	DirectDebitOp               *DirectDebitOp               `json:"directDebitOp,omitempty"`
	ManageAssetPairOp           *ManageAssetPairOp           `json:"manageAssetPairOp,omitempty"`
	ManageOfferOp               *ManageOfferOp               `json:"manageOfferOp,omitempty"`
	ManageInvoiceOp             *ManageInvoiceOp             `json:"manageInvoiceOp,omitempty"`
	ReviewRequestOp             *ReviewRequestOp             `json:"reviewRequestOp,omitempty"`
	CreateSaleCreationRequestOp *CreateSaleCreationRequestOp `json:"createSaleCreationRequestOp,omitempty"`
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
	case OperationTypeRecover:
		return "RecoverOp", true
	case OperationTypeManageBalance:
		return "ManageBalanceOp", true
	case OperationTypeReviewPaymentRequest:
		return "ReviewPaymentRequestOp", true
	case OperationTypeManageAsset:
		return "ManageAssetOp", true
	case OperationTypeCreatePreissuanceRequest:
		return "CreatePreIssuanceRequest", true
	case OperationTypeSetLimits:
		return "SetLimitsOp", true
	case OperationTypeDirectDebit:
		return "DirectDebitOp", true
	case OperationTypeManageAssetPair:
		return "ManageAssetPairOp", true
	case OperationTypeManageOffer:
		return "ManageOfferOp", true
	case OperationTypeManageInvoice:
		return "ManageInvoiceOp", true
	case OperationTypeReviewRequest:
		return "ReviewRequestOp", true
	case OperationTypeCreateSaleRequest:
		return "CreateSaleCreationRequestOp", true
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
	case OperationTypeRecover:
		tv, ok := value.(RecoverOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be RecoverOp")
			return
		}
		result.RecoverOp = &tv
	case OperationTypeManageBalance:
		tv, ok := value.(ManageBalanceOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageBalanceOp")
			return
		}
		result.ManageBalanceOp = &tv
	case OperationTypeReviewPaymentRequest:
		tv, ok := value.(ReviewPaymentRequestOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be ReviewPaymentRequestOp")
			return
		}
		result.ReviewPaymentRequestOp = &tv
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
	case OperationTypeSetLimits:
		tv, ok := value.(SetLimitsOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be SetLimitsOp")
			return
		}
		result.SetLimitsOp = &tv
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
	case OperationTypeManageInvoice:
		tv, ok := value.(ManageInvoiceOp)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageInvoiceOp")
			return
		}
		result.ManageInvoiceOp = &tv
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

// MustRecoverOp retrieves the RecoverOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustRecoverOp() RecoverOp {
	val, ok := u.GetRecoverOp()

	if !ok {
		panic("arm RecoverOp is not set")
	}

	return val
}

// GetRecoverOp retrieves the RecoverOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetRecoverOp() (result RecoverOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "RecoverOp" {
		result = *u.RecoverOp
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

// MustReviewPaymentRequestOp retrieves the ReviewPaymentRequestOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustReviewPaymentRequestOp() ReviewPaymentRequestOp {
	val, ok := u.GetReviewPaymentRequestOp()

	if !ok {
		panic("arm ReviewPaymentRequestOp is not set")
	}

	return val
}

// GetReviewPaymentRequestOp retrieves the ReviewPaymentRequestOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetReviewPaymentRequestOp() (result ReviewPaymentRequestOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ReviewPaymentRequestOp" {
		result = *u.ReviewPaymentRequestOp
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

// MustSetLimitsOp retrieves the SetLimitsOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustSetLimitsOp() SetLimitsOp {
	val, ok := u.GetSetLimitsOp()

	if !ok {
		panic("arm SetLimitsOp is not set")
	}

	return val
}

// GetSetLimitsOp retrieves the SetLimitsOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetSetLimitsOp() (result SetLimitsOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "SetLimitsOp" {
		result = *u.SetLimitsOp
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

// MustManageInvoiceOp retrieves the ManageInvoiceOp value from the union,
// panicing if the value is not set.
func (u OperationBody) MustManageInvoiceOp() ManageInvoiceOp {
	val, ok := u.GetManageInvoiceOp()

	if !ok {
		panic("arm ManageInvoiceOp is not set")
	}

	return val
}

// GetManageInvoiceOp retrieves the ManageInvoiceOp value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationBody) GetManageInvoiceOp() (result ManageInvoiceOp, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ManageInvoiceOp" {
		result = *u.ManageInvoiceOp
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
//    	case RECOVER:
//    		RecoverOp recoverOp;
//    	case MANAGE_BALANCE:
//    		ManageBalanceOp manageBalanceOp;
//    	case REVIEW_PAYMENT_REQUEST:
//    		ReviewPaymentRequestOp reviewPaymentRequestOp;
//        case MANAGE_ASSET:
//            ManageAssetOp manageAssetOp;
//        case CREATE_PREISSUANCE_REQUEST:
//            CreatePreIssuanceRequestOp createPreIssuanceRequest;
//        case SET_LIMITS:
//            SetLimitsOp setLimitsOp;
//        case DIRECT_DEBIT:
//            DirectDebitOp directDebitOp;
//    	case MANAGE_ASSET_PAIR:
//    		ManageAssetPairOp manageAssetPairOp;
//    	case MANAGE_OFFER:
//    		ManageOfferOp manageOfferOp;
//        case MANAGE_INVOICE:
//            ManageInvoiceOp manageInvoiceOp;
//    	case REVIEW_REQUEST:
//    		ReviewRequestOp reviewRequestOp;
//    	case CREATE_SALE_REQUEST:
//    		CreateSaleCreationRequestOp createSaleCreationRequestOp;
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
//        case RECOVER:
//    		RecoverResult recoverResult;
//        case MANAGE_BALANCE:
//            ManageBalanceResult manageBalanceResult;
//        case REVIEW_PAYMENT_REQUEST:
//            ReviewPaymentRequestResult reviewPaymentRequestResult;
//        case MANAGE_ASSET:
//            ManageAssetResult manageAssetResult;
//        case CREATE_PREISSUANCE_REQUEST:
//            CreatePreIssuanceRequestResult createPreIssuanceRequestResult;
//        case SET_LIMITS:
//            SetLimitsResult setLimitsResult;
//        case DIRECT_DEBIT:
//            DirectDebitResult directDebitResult;
//    	case MANAGE_ASSET_PAIR:
//    		ManageAssetPairResult manageAssetPairResult;
//    	case MANAGE_OFFER:
//    		ManageOfferResult manageOfferResult;
//    	case MANAGE_INVOICE:
//    		ManageInvoiceResult manageInvoiceResult;
//    	case REVIEW_REQUEST:
//    		ReviewRequestResult reviewRequestResult;
//    	case CREATE_SALE_REQUEST:
//    		CreateSaleCreationRequestResult createSaleCreationRequestResult;
//        }
//
type OperationResultTr struct {
	Type                            OperationType                    `json:"type,omitempty"`
	CreateAccountResult             *CreateAccountResult             `json:"createAccountResult,omitempty"`
	PaymentResult                   *PaymentResult                   `json:"paymentResult,omitempty"`
	SetOptionsResult                *SetOptionsResult                `json:"setOptionsResult,omitempty"`
	CreateIssuanceRequestResult     *CreateIssuanceRequestResult     `json:"createIssuanceRequestResult,omitempty"`
	SetFeesResult                   *SetFeesResult                   `json:"setFeesResult,omitempty"`
	ManageAccountResult             *ManageAccountResult             `json:"manageAccountResult,omitempty"`
	CreateWithdrawalRequestResult   *CreateWithdrawalRequestResult   `json:"createWithdrawalRequestResult,omitempty"`
	RecoverResult                   *RecoverResult                   `json:"recoverResult,omitempty"`
	ManageBalanceResult             *ManageBalanceResult             `json:"manageBalanceResult,omitempty"`
	ReviewPaymentRequestResult      *ReviewPaymentRequestResult      `json:"reviewPaymentRequestResult,omitempty"`
	ManageAssetResult               *ManageAssetResult               `json:"manageAssetResult,omitempty"`
	CreatePreIssuanceRequestResult  *CreatePreIssuanceRequestResult  `json:"createPreIssuanceRequestResult,omitempty"`
	SetLimitsResult                 *SetLimitsResult                 `json:"setLimitsResult,omitempty"`
	DirectDebitResult               *DirectDebitResult               `json:"directDebitResult,omitempty"`
	ManageAssetPairResult           *ManageAssetPairResult           `json:"manageAssetPairResult,omitempty"`
	ManageOfferResult               *ManageOfferResult               `json:"manageOfferResult,omitempty"`
	ManageInvoiceResult             *ManageInvoiceResult             `json:"manageInvoiceResult,omitempty"`
	ReviewRequestResult             *ReviewRequestResult             `json:"reviewRequestResult,omitempty"`
	CreateSaleCreationRequestResult *CreateSaleCreationRequestResult `json:"createSaleCreationRequestResult,omitempty"`
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
	case OperationTypeRecover:
		return "RecoverResult", true
	case OperationTypeManageBalance:
		return "ManageBalanceResult", true
	case OperationTypeReviewPaymentRequest:
		return "ReviewPaymentRequestResult", true
	case OperationTypeManageAsset:
		return "ManageAssetResult", true
	case OperationTypeCreatePreissuanceRequest:
		return "CreatePreIssuanceRequestResult", true
	case OperationTypeSetLimits:
		return "SetLimitsResult", true
	case OperationTypeDirectDebit:
		return "DirectDebitResult", true
	case OperationTypeManageAssetPair:
		return "ManageAssetPairResult", true
	case OperationTypeManageOffer:
		return "ManageOfferResult", true
	case OperationTypeManageInvoice:
		return "ManageInvoiceResult", true
	case OperationTypeReviewRequest:
		return "ReviewRequestResult", true
	case OperationTypeCreateSaleRequest:
		return "CreateSaleCreationRequestResult", true
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
	case OperationTypeRecover:
		tv, ok := value.(RecoverResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be RecoverResult")
			return
		}
		result.RecoverResult = &tv
	case OperationTypeManageBalance:
		tv, ok := value.(ManageBalanceResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageBalanceResult")
			return
		}
		result.ManageBalanceResult = &tv
	case OperationTypeReviewPaymentRequest:
		tv, ok := value.(ReviewPaymentRequestResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be ReviewPaymentRequestResult")
			return
		}
		result.ReviewPaymentRequestResult = &tv
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
	case OperationTypeSetLimits:
		tv, ok := value.(SetLimitsResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be SetLimitsResult")
			return
		}
		result.SetLimitsResult = &tv
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
	case OperationTypeManageInvoice:
		tv, ok := value.(ManageInvoiceResult)
		if !ok {
			err = fmt.Errorf("invalid value, must be ManageInvoiceResult")
			return
		}
		result.ManageInvoiceResult = &tv
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

// MustRecoverResult retrieves the RecoverResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustRecoverResult() RecoverResult {
	val, ok := u.GetRecoverResult()

	if !ok {
		panic("arm RecoverResult is not set")
	}

	return val
}

// GetRecoverResult retrieves the RecoverResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetRecoverResult() (result RecoverResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "RecoverResult" {
		result = *u.RecoverResult
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

// MustReviewPaymentRequestResult retrieves the ReviewPaymentRequestResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustReviewPaymentRequestResult() ReviewPaymentRequestResult {
	val, ok := u.GetReviewPaymentRequestResult()

	if !ok {
		panic("arm ReviewPaymentRequestResult is not set")
	}

	return val
}

// GetReviewPaymentRequestResult retrieves the ReviewPaymentRequestResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetReviewPaymentRequestResult() (result ReviewPaymentRequestResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ReviewPaymentRequestResult" {
		result = *u.ReviewPaymentRequestResult
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

// MustSetLimitsResult retrieves the SetLimitsResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustSetLimitsResult() SetLimitsResult {
	val, ok := u.GetSetLimitsResult()

	if !ok {
		panic("arm SetLimitsResult is not set")
	}

	return val
}

// GetSetLimitsResult retrieves the SetLimitsResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetSetLimitsResult() (result SetLimitsResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "SetLimitsResult" {
		result = *u.SetLimitsResult
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

// MustManageInvoiceResult retrieves the ManageInvoiceResult value from the union,
// panicing if the value is not set.
func (u OperationResultTr) MustManageInvoiceResult() ManageInvoiceResult {
	val, ok := u.GetManageInvoiceResult()

	if !ok {
		panic("arm ManageInvoiceResult is not set")
	}

	return val
}

// GetManageInvoiceResult retrieves the ManageInvoiceResult value from the union,
// returning ok if the union's switch indicated the value is valid.
func (u OperationResultTr) GetManageInvoiceResult() (result ManageInvoiceResult, ok bool) {
	armName, _ := u.ArmForSwitch(int32(u.Type))

	if armName == "ManageInvoiceResult" {
		result = *u.ManageInvoiceResult
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
//        case RECOVER:
//    		RecoverResult recoverResult;
//        case MANAGE_BALANCE:
//            ManageBalanceResult manageBalanceResult;
//        case REVIEW_PAYMENT_REQUEST:
//            ReviewPaymentRequestResult reviewPaymentRequestResult;
//        case MANAGE_ASSET:
//            ManageAssetResult manageAssetResult;
//        case CREATE_PREISSUANCE_REQUEST:
//            CreatePreIssuanceRequestResult createPreIssuanceRequestResult;
//        case SET_LIMITS:
//            SetLimitsResult setLimitsResult;
//        case DIRECT_DEBIT:
//            DirectDebitResult directDebitResult;
//    	case MANAGE_ASSET_PAIR:
//    		ManageAssetPairResult manageAssetPairResult;
//    	case MANAGE_OFFER:
//    		ManageOfferResult manageOfferResult;
//    	case MANAGE_INVOICE:
//    		ManageInvoiceResult manageInvoiceResult;
//    	case REVIEW_REQUEST:
//    		ReviewRequestResult reviewRequestResult;
//    	case CREATE_SALE_REQUEST:
//    		CreateSaleCreationRequestResult createSaleCreationRequestResult;
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
//    	IMPROVED_STATS_CALCULATION = 4,
//    	EMISSION_REQUEST_BALANCE_ID = 5,
//    	IMPROVED_TRANSFER_FEES_CALC = 8,
//    	USE_IMPROVED_SIGNATURE_CHECK = 9
//    };
//
type LedgerVersion int32

const (
	LedgerVersionEmptyVersion              LedgerVersion = 0
	LedgerVersionImprovedStatsCalculation  LedgerVersion = 4
	LedgerVersionEmissionRequestBalanceId  LedgerVersion = 5
	LedgerVersionImprovedTransferFeesCalc  LedgerVersion = 8
	LedgerVersionUseImprovedSignatureCheck LedgerVersion = 9
)

var LedgerVersionAll = []LedgerVersion{
	LedgerVersionEmptyVersion,
	LedgerVersionImprovedStatsCalculation,
	LedgerVersionEmissionRequestBalanceId,
	LedgerVersionImprovedTransferFeesCalc,
	LedgerVersionUseImprovedSignatureCheck,
}

var ledgerVersionMap = map[int32]string{
	0: "LedgerVersionEmptyVersion",
	4: "LedgerVersionImprovedStatsCalculation",
	5: "LedgerVersionEmissionRequestBalanceId",
	8: "LedgerVersionImprovedTransferFeesCalc",
	9: "LedgerVersionUseImprovedSignatureCheck",
}

var ledgerVersionShortMap = map[int32]string{
	0: "empty_version",
	4: "improved_stats_calculation",
	5: "emission_request_balance_id",
	8: "improved_transfer_fees_calc",
	9: "use_improved_signature_check",
}

var ledgerVersionRevMap = map[string]int32{
	"LedgerVersionEmptyVersion":              0,
	"LedgerVersionImprovedStatsCalculation":  4,
	"LedgerVersionEmissionRequestBalanceId":  5,
	"LedgerVersionImprovedTransferFeesCalc":  8,
	"LedgerVersionUseImprovedSignatureCheck": 9,
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
//        RECOVER = 8,
//        MANAGE_BALANCE = 9,
//        REVIEW_PAYMENT_REQUEST = 10,
//        MANAGE_ASSET = 11,
//        CREATE_PREISSUANCE_REQUEST = 12,
//        SET_LIMITS = 13,
//        DIRECT_DEBIT = 14,
//    	MANAGE_ASSET_PAIR = 15,
//    	MANAGE_OFFER = 16,
//        MANAGE_INVOICE = 17,
//    	REVIEW_REQUEST = 18,
//    	CREATE_SALE_REQUEST = 19
//    };
//
type OperationType int32

const (
	OperationTypeCreateAccount            OperationType = 0
	OperationTypePayment                  OperationType = 1
	OperationTypeSetOptions               OperationType = 2
	OperationTypeCreateIssuanceRequest    OperationType = 3
	OperationTypeSetFees                  OperationType = 5
	OperationTypeManageAccount            OperationType = 6
	OperationTypeCreateWithdrawalRequest  OperationType = 7
	OperationTypeRecover                  OperationType = 8
	OperationTypeManageBalance            OperationType = 9
	OperationTypeReviewPaymentRequest     OperationType = 10
	OperationTypeManageAsset              OperationType = 11
	OperationTypeCreatePreissuanceRequest OperationType = 12
	OperationTypeSetLimits                OperationType = 13
	OperationTypeDirectDebit              OperationType = 14
	OperationTypeManageAssetPair          OperationType = 15
	OperationTypeManageOffer              OperationType = 16
	OperationTypeManageInvoice            OperationType = 17
	OperationTypeReviewRequest            OperationType = 18
	OperationTypeCreateSaleRequest        OperationType = 19
)

var OperationTypeAll = []OperationType{
	OperationTypeCreateAccount,
	OperationTypePayment,
	OperationTypeSetOptions,
	OperationTypeCreateIssuanceRequest,
	OperationTypeSetFees,
	OperationTypeManageAccount,
	OperationTypeCreateWithdrawalRequest,
	OperationTypeRecover,
	OperationTypeManageBalance,
	OperationTypeReviewPaymentRequest,
	OperationTypeManageAsset,
	OperationTypeCreatePreissuanceRequest,
	OperationTypeSetLimits,
	OperationTypeDirectDebit,
	OperationTypeManageAssetPair,
	OperationTypeManageOffer,
	OperationTypeManageInvoice,
	OperationTypeReviewRequest,
	OperationTypeCreateSaleRequest,
}

var operationTypeMap = map[int32]string{
	0:  "OperationTypeCreateAccount",
	1:  "OperationTypePayment",
	2:  "OperationTypeSetOptions",
	3:  "OperationTypeCreateIssuanceRequest",
	5:  "OperationTypeSetFees",
	6:  "OperationTypeManageAccount",
	7:  "OperationTypeCreateWithdrawalRequest",
	8:  "OperationTypeRecover",
	9:  "OperationTypeManageBalance",
	10: "OperationTypeReviewPaymentRequest",
	11: "OperationTypeManageAsset",
	12: "OperationTypeCreatePreissuanceRequest",
	13: "OperationTypeSetLimits",
	14: "OperationTypeDirectDebit",
	15: "OperationTypeManageAssetPair",
	16: "OperationTypeManageOffer",
	17: "OperationTypeManageInvoice",
	18: "OperationTypeReviewRequest",
	19: "OperationTypeCreateSaleRequest",
}

var operationTypeShortMap = map[int32]string{
	0:  "create_account",
	1:  "payment",
	2:  "set_options",
	3:  "create_issuance_request",
	5:  "set_fees",
	6:  "manage_account",
	7:  "create_withdrawal_request",
	8:  "recover",
	9:  "manage_balance",
	10: "review_payment_request",
	11: "manage_asset",
	12: "create_preissuance_request",
	13: "set_limits",
	14: "direct_debit",
	15: "manage_asset_pair",
	16: "manage_offer",
	17: "manage_invoice",
	18: "review_request",
	19: "create_sale_request",
}

var operationTypeRevMap = map[string]int32{
	"OperationTypeCreateAccount":            0,
	"OperationTypePayment":                  1,
	"OperationTypeSetOptions":               2,
	"OperationTypeCreateIssuanceRequest":    3,
	"OperationTypeSetFees":                  5,
	"OperationTypeManageAccount":            6,
	"OperationTypeCreateWithdrawalRequest":  7,
	"OperationTypeRecover":                  8,
	"OperationTypeManageBalance":            9,
	"OperationTypeReviewPaymentRequest":     10,
	"OperationTypeManageAsset":              11,
	"OperationTypeCreatePreissuanceRequest": 12,
	"OperationTypeSetLimits":                13,
	"OperationTypeDirectDebit":              14,
	"OperationTypeManageAssetPair":          15,
	"OperationTypeManageOffer":              16,
	"OperationTypeManageInvoice":            17,
	"OperationTypeReviewRequest":            18,
	"OperationTypeCreateSaleRequest":        19,
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
