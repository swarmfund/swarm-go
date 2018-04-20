package xdr

// package contains helper struct for xdr enums marshaling

type value struct {
	Value int32 `json:"int"`
}

type enum struct {
	Value  int32  `json:"int"`
	String string `json:"string"`
}

type flag struct {
	Value int32       `json:"int"`
	Flags []flagValue `json:"flags"`
}

type flagValue struct {
	Name  string `json:"name"`
	Value int32  `json:"value"`
}
