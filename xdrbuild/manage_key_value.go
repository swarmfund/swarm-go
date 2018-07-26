package xdrbuild

import (
	. "github.com/go-ozzo/ozzo-validation"
	"github.com/pkg/errors"
	"gitlab.com/tokend/go/xdr"
)

type (
	ManageKeyValueOp struct {
		Key    string
		Uint32 *uint32
		String *string
	}
)

func (mkv ManageKeyValueOp) Validate() error {
	var fieldRules []*FieldRules

	fieldRules = append(fieldRules, Field(&mkv.Key, Required))
	if mkv.Uint32 == nil && mkv.String != nil {
		fieldRules = append(fieldRules, Field(&mkv.String, Required))
		return ValidateStruct(&mkv, fieldRules...)
	} else {
		fieldRules = append(fieldRules, Field(&mkv.Uint32, Required))
		return ValidateStruct(&mkv, fieldRules...)
	}
	return errors.New("Only Uint32 or String required, not both")
}

func (mkv ManageKeyValueOp) XDR() (*xdr.Operation, error) {
	op := &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeManageKeyValue,
			ManageKeyValueOp: &xdr.ManageKeyValueOp{
				Key: xdr.String256(mkv.Key),
				Action: xdr.ManageKeyValueOpAction{
					Action: xdr.ManageKvActionPut,
					Value: &xdr.KeyValueEntry{
						Key: xdr.Longstring(mkv.Key),
					},
				},
			},
		},
	}

	if mkv.Uint32 != nil {
		op.Body.ManageKeyValueOp.Action.Value.Value.Type = xdr.KeyValueEntryTypeUint32
		val := xdr.Uint32(*mkv.Uint32)
		op.Body.ManageKeyValueOp.Action.Value.Value.Ui32Value = &val
	} else {
		op.Body.ManageKeyValueOp.Action.Value.Value.Type = xdr.KeyValueEntryTypeString
		op.Body.ManageKeyValueOp.Action.Value.Value.StringValue = mkv.String
	}
	return op, nil
}
