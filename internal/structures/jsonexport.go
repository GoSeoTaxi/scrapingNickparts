package structures

type JsonExport struct {
	RequestItemData   RequestItem
	OriginalAnalogs   []OriginalAnalog
	NoOriginalAnalogs []NoOriginalAnalog
}

type RequestItemparameter struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type RequestItem struct {
	Num        string `json:"num"`
	Brand      string `json:"brand"`
	Text       string `json:"text"`
	Pic        string `json:"pic"`
	Price      string `json:"price"`
	Parameters []RequestItemparameter
}

type OriginalAnalog struct {
	Num   string `json:"num"`
	Brand string `json:"brand"`
	Text  string `json:"text"`
	Pic   string `json:"pic"`
	Price string `json:"price"`
}

type NoOriginalAnalog struct {
	Num   string `json:"num"`
	Brand string `json:"brand"`
	Text  string `json:"text"`
	Pic   string `json:"pic"`
	Price string `json:"price"`
}
