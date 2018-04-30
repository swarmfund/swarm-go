package xdrbuild

import (
	"math"

	. "github.com/go-ozzo/ozzo-validation"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/go/xdr"
)

var ErrThresholdOutOfRange = errors.New("threshold out of range")

type Signer struct {
	PublicKey  string
	Weight     uint32
	SignerType uint32
	Identity   uint32
	Name       string
}

type SetOptions struct {
	Signer        *Signer
	MasterWeight  *uint32
	LowThreshold  *uint32
	MedThreshold  *uint32
	HighThreshold *uint32
}

func (op SetOptions) Validate() error {
	errs := Errors{}

	if op.Signer != nil {
		errs["/op/signer/public_key"] = Validate(op.Signer.PublicKey, Required)
		errs["/op/signer/signer_type"] = Validate(op.Signer.SignerType, Required)
		errs["/op/signer/identity"] = Validate(op.Signer.Identity, Required)
		errs["/op/signer/name"] = Validate(op.Signer.Name, Required)
	}

	if op.LowThreshold != nil {
		errs["/op/low_threshold"] = validateThreshold(*op.LowThreshold)
	}

	if op.MedThreshold != nil {
		errs["/op/med_threshold"] = validateThreshold(*op.MedThreshold)
	}

	if op.HighThreshold != nil {
		errs["/op/high_threshold"] = validateThreshold(*op.HighThreshold)
	}

	return errs.Filter()
}

func (op SetOptions) XDR() (*xdr.Operation, error) {
	signer := &xdr.Signer{
		Weight:     xdr.Uint32(op.Signer.Weight),
		SignerType: xdr.Uint32(op.Signer.SignerType),
		Identity:   xdr.Uint32(op.Signer.Identity),
		Name:       xdr.String256(op.Signer.Name),
	}

	var signerPubKey xdr.AccountId
	if err := signerPubKey.SetAddress(op.Signer.PublicKey); err != nil {
		return nil, errors.Wrap(err, "failed to set signer public key")
	}

	signer.PubKey = signerPubKey

	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeSetOptions,
			SetOptionsOp: &xdr.SetOptionsOp{
				MasterWeight:  (*xdr.Uint32)(op.MasterWeight),
				LowThreshold:  (*xdr.Uint32)(op.LowThreshold),
				MedThreshold:  (*xdr.Uint32)(op.MedThreshold),
				HighThreshold: (*xdr.Uint32)(op.HighThreshold),
				Signer:        signer,
			},
		},
	}, nil
}

//DeleteSigner create new SetOptions without signer weight
//by default if signer weight is zero, it mean delete signer
func DeleteSigner(publicKey, name string, signerType, identity uint32) *SetOptions {
	return &SetOptions{
		Signer: &Signer{
			PublicKey:  publicKey,
			SignerType: signerType,
			Identity:   identity,
			Name:       name,
		},
	}
}

func AddSigner(publicKey, name string, weight, signerType, identity uint32) *SetOptions {
	return &SetOptions{
		Signer: &Signer{
			PublicKey:  publicKey,
			Weight:     weight,
			SignerType: signerType,
			Identity:   identity,
			Name:       name,
		},
	}
}

func SetThreshold(masterWeight, lowThreshold, medThreshold, highThreshold uint32) *SetOptions {
	return &SetOptions{
		MasterWeight:  &masterWeight,
		LowThreshold:  &lowThreshold,
		MedThreshold:  &medThreshold,
		HighThreshold: &highThreshold,
	}
}

func validateThreshold(value uint32) error {
	if value > math.MaxUint8 {
		return ErrThresholdOutOfRange
	}
	return nil
}
