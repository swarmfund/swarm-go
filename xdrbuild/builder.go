package xdrbuild

import "gitlab.com/tokend/keypair"

type Builder struct {
	passphrase         string
	txExpirationPeriod int64
}

func NewBuilder(passphrase string, txExpire int64) *Builder {
	return &Builder{
		passphrase:         passphrase,
		txExpirationPeriod: txExpire,
	}
}

func (b *Builder) Transaction(source keypair.Address) *Transaction {
	return &Transaction{
		builder: b,
		source:  source,
	}
}
