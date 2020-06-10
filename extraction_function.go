package godruid

type ExtractionFn interface{}

type LookupExtractionFn struct {
	Type                    string      `json:"type"`
	LookUp                  interface{} `json:"lookup,omitempty"`
	RetainMissingValue      bool        `json:"retainMissingValue,omitempty"`
	ReplaceMissingValueWith string      `json:"replaceMissingValueWith,omitempty"`
	Injective               bool        `json:"injective,omitempty"`
	Optimize                bool        `json:"optimize,omitempty"`
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
	return &LookupExtractionFn{
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

func DimExFnRegisteredLookup(lookup string, retainMissingValue bool, replaceMissingValueWith string, optimize bool) ExtractionFn {
	return &LookupExtractionFn{
		Type:                    "registeredLookup",
		RetainMissingValue:      retainMissingValue,
		ReplaceMissingValueWith: replaceMissingValueWith,
		Optimize:                optimize,
		LookUp:                  lookup,
	}
}