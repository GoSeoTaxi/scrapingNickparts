package structures

type JsonImport struct {
	Result []struct {
		Num     string `json:"num"`
		Brand   string `json:"brand"`
		AddInfo int    `json:"addInfo"`
	} `json:"result"`
	Query []interface{} `json:"query"`
}
