package transfer

type ProductAttribute struct {
	Attribute  string `json:"attribute"`
	LocaleCode string `json:"localeCode"`
	Value      string `json:"value"`
}