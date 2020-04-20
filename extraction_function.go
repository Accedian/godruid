package godruid

type ExtractionFn interface{}

type RegisteredLookupExtractionFn struct {
	Type   string `json:"type"`
	Lookup string `json:"lookup"`
}

type InlineLookupExtractionFn struct {
	Type                    string               `json:"type"`
	LookUp                  *LookUpExtractionMap `json:"lookup,omitempty"`
	RetainMissingValue      bool                 `json:"retainMissingValue,omitempty"`
	ReplaceMissingValueWith string               `json:"replaceMissingValueWith,omitempty"`
	Injective               bool                 `json:"injective,omitempty"`
}

type LookUpExtractionMap struct {
	Type string                 `json:"type"`
	Map  map[string]interface{} `json:"map"`
}

type TimeExtractionFn struct {
	Type     string `json:"type"`
	Format   string `json:"format"`
	TimeZone string `json:"timeZone"`
	Locale   string `json:"locale"`
}

func DimExFnInlineLookUp(lookups map[string]interface{}, retainMissingValue bool, replaceMissingValueWith string, injective bool) ExtractionFn {
	return &InlineLookupExtractionFn{
		Type:                    "lookup",
		RetainMissingValue:      retainMissingValue,
		ReplaceMissingValueWith: replaceMissingValueWith,
		Injective:               injective,
		LookUp: &LookUpExtractionMap{
			Type: "map",
			Map:  lookups,
		},
	}
}