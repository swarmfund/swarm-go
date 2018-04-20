package resources

type Signer struct {
	AccountID  string `json:"public_key"`
	Weight     int    `json:"weight"`
	SignerType int    `json:"signer_type_i"`
	Identity   int    `json:"signer_identity"`
	Name       string `json:"signer_name"`
}
