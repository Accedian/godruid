package godruid

type DimSpec interface{}

type Dimension struct {
	Type         string       `json:"type"`
	Dimension    string       `json:"dimension"`
	OutputName   string       `json:"outputName"`
	Name         string       `json:"name"`
	ExtractionFn ExtractionFn `json:"extractionFn,omitempty"`
}

type DimExtractionFn struct {
	Type        string       `json:"type"`
	Expr        string       `json:"expr,omitempty"`
	Query       *SearchQuery `json:"query,omitempty"`
	Format      string       `json:"format,omitempty"`
	Function    string       `json:"function,omitempty"`
	TimeZone    string       `json:"timeZone,omitempty"`
	Locale      string       `json:"locale,omitempty"`
	Granularity string       `json:"granularity,omitempty"`
	AsMillis    bool         `json:"asMillis,omitempty"`
}

type TimeExtractionDimensionSpec struct {
	Type               string       `json:"type"`
	Dimension          string       `json:"dimension"`
	OutputName         string       `json:"outputName"`
	ExtractionFunction ExtractionFn `json:"extractionFn"`
}

func DimDefault(dimension, outputName string) DimSpec {
	return &Dimension{
		Type:       "default",
		Dimension:  dimension,
		OutputName: outputName,
	}
}

func DimExtraction(dimension, outputName string, fn *DimExtractionFn) DimSpec {
	return &Dimension{
		Type:         "extraction",
		Dimension:    dimension,
		OutputName:   outputName,
		ExtractionFn: fn,
	}
}

func DimLookup(dimension, outputName, name string) DimSpec {
	return &Dimension{
		Type:       "lookup",
		Dimension:  dimension,
		OutputName: outputName,
		Name:       name,
	}
}

func DimExFnRegex(expr string) *DimExtractionFn {
	return &DimExtractionFn{
		Type: "regex",
		Expr: expr,
	}
}

func DimExFnPartial(expr string) *DimExtractionFn {
	return &DimExtractionFn{
		Type: "partial",
		Expr: expr,
	}
}

func DimExFnSearchQuerySpec(query *SearchQuery) *DimExtractionFn {
	return &DimExtractionFn{
		Type:  "searchQuery",
		Query: query,
	}
}

func DimExFnTime(timeFormat, timeZone string, locale string, granularity string, asMillis bool) *DimExtractionFn {
	return &DimExtractionFn{
		Type:        "timeFormat",
		Format:      timeFormat,
		TimeZone:    timeZone,
		Locale:      locale,
		Granularity: granularity,
		AsMillis:    asMillis,
	}
}

func DimExFnJavascript(function string) *DimExtractionFn {
	return &DimExtractionFn{
		Type:     "javascript",
		Function: function,
	}
}
