package xdrbuild

import (
	. "github.com/go-ozzo/ozzo-validation"
	"github.com/pkg/errors"
	"gitlab.com/tokend/go/xdr"
)

type ManageKeyValueOp struct {
	Key    string
	Uint32 *uint32
	String *string
}

func (mkv ManageKeyValueOp) Validate() error {
	if (mkv.String != nil) && (mkv.Uint32 != nil) {
		return errors.New("unexpected state, Both not nil Uint32 and String not allowed")
	}

	return ValidateStruct(&mkv,
		Field(&mkv.Key, Required),
		Field(&mkv.Uint32, NilOrNotEmpty),
		Field(&mkv.String, NilOrNotEmpty),
	)
}

func (mkv ManageKeyValueOp) XDR() (*xdr.Operation, error) {
	manageKvAction := xdr.ManageKvActionRemove
	keyValueEntry := (*xdr.KeyValueEntry)(nil)
	switch {
	case mkv.Uint32 != nil:
		manageKvAction = xdr.ManageKvActionPut
		val := xdr.Uint32(*mkv.Uint32)
		keyValueEntry = &xdr.KeyValueEntry{
			Key: xdr.Longstring(mkv.Key),
			Value: xdr.KeyValueEntryValue{
				Type:      xdr.KeyValueEntryTypeUint32,
				Ui32Value: &val,
			},
		}
	case mkv.String != nil:
		manageKvAction = xdr.ManageKvActionPut
		keyValueEntry = &xdr.KeyValueEntry{
			Key: xdr.Longstring(mkv.Key),
			Value: xdr.KeyValueEntryValue{
				Type:        xdr.KeyValueEntryTypeString,
				StringValue: mkv.String,
			},
		}
	}

	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeManageKeyValue,
			ManageKeyValueOp: &xdr.ManageKeyValueOp{
				Key: xdr.String256(mkv.Key),
				Action: xdr.ManageKeyValueOpAction{
					Action: manageKvAction,
					Value:  keyValueEntry,
				},
			},
		},
	}, nil
}
