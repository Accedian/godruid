package godruid

type Filter struct {
	Type         string             `json:"type"`
	Dimension    string             `json:"dimension,omitempty"`
	Dimensions   DimSpec            `json:"dimensions,omitempty"`
	Value        interface{}        `json:"value,omitempty"`
	Values       interface{}        `json:"values,omitempty"`
	Pattern      string             `json:"pattern,omitempty"`
	Function     string             `json:"function,omitempty"`
	Field        *Filter            `json:"field,omitempty"`
	Fields       []*Filter          `json:"fields,omitempty"`
	Upper        string             `json:"upper,omitempty"`
	Lower        string             `json:"lower,omitempty"`
	Ordering     Ordering           `json:"ordering,omitempty"`
	UpperStrict  bool               `json:"upperStrict,omitempty"`
	LowerStrict  bool               `json:"lowerStrict,omitempty"`
	ExtractionFn ExtractionFn       `json:"extractionFn,omitempty"`
	Bound        *Bound             `json:"bound,omitempty"`
	Query        *SearchFilterQuery `json:"query,omitempty"`
}

type Bound struct {
	Type      string    `json:"type"`
	MinCoords []float64 `json:"minCoords,omitempty"`
	MaxCoords []float64 `json:"maxCoords,omitempty"`
	Coords    []float64 `json:"coords,omitempty"`
	Radius    float64   `json:"radius,omitempty"`
}

type Ordering string

const (
	LEXICOGRAPHIC Ordering = "lexicographic"
	ALPHANUMERIC  Ordering = "alphanumeric"
	NUMERIC       Ordering = "numeric"
	STRLEN        Ordering = "strlen"
)

// Filter constants
const (
	LOWERSTRICT = "lowerStrict"
	UPPERSTRICT = "upperStrict"
	LOWERLIMIT  = "lowerLimit"
	UPPERLIMIT  = "upperLimit"
)

type SpatialCoordinates struct {
	Latitude  float64
	Longitude float64
}

type SearchFilterQuery struct {
	Type          string   `json:"type"`
	Value         string   `json:"value,omitempty"`
	Values        []string `json:"values,omitempty"`
	CaseSensitive bool     `json:"caseSensitive,omitempty"`
}

func FilterSpatialRectangle(dimension string, minCoords SpatialCoordinates, maxCoords SpatialCoordinates) *Filter {
	return &Filter{
		Type:      "spatial",
		Dimension: dimension,
		Bound: &Bound{
			Type:      "rectangular",
			MinCoords: []float64{minCoords.Latitude, minCoords.Longitude},
			MaxCoords: []float64{maxCoords.Latitude, maxCoords.Longitude},
		},
	}
}

func FilterSpatialRadius(dimension string, coords SpatialCoordinates, radius float64) *Filter {
	return &Filter{
		Type:      "spatial",
		Dimension: dimension,
		Bound: &Bound{
			Type:   "radius",
			Coords: []float64{coords.Latitude, coords.Longitude},
			Radius: radius,
		},
	}
}

func FilterSelector(dimension string, value interface{}) *Filter {
	return &Filter{
		Type:      "selector",
		Dimension: dimension,
		Value:     value,
	}
}

func FilterUpperBound(dimension string, ordering Ordering, bound string, strict bool) *Filter {
	return &Filter{
		Type:        "bound",
		Dimension:   dimension,
		Ordering:    ordering,
		Upper:       bound,
		UpperStrict: strict,
	}
}

func FilterLowerBound(dimension string, ordering Ordering, bound string, strict bool) *Filter {
	return &Filter{
		Type:        "bound",
		Dimension:   dimension,
		Ordering:    ordering,
		Lower:       bound,
		LowerStrict: strict,
	}
}

func FilterLowerUpperBound(dimension string, ordering Ordering, lowerBound string, lowerStrict bool, upperBound string, upperStrict bool) *Filter {
	return &Filter{
		Type:        "bound",
		Dimension:   dimension,
		Ordering:    ordering,
		Lower:       lowerBound,
		LowerStrict: lowerStrict,
		Upper:       upperBound,
		UpperStrict: upperStrict,
	}
}

func FilterColumnComparison(dimensions []DimSpec) *Filter {
	return &Filter{
		Type:        "columnComparison",
		Dimensions:   dimensions,
	}
}

func FilterRegex(dimension, pattern string) *Filter {
	return &Filter{
		Type:      "regex",
		Dimension: dimension,
		Pattern:   pattern,
	}
}

func FilterJavaScript(dimension, function string) *Filter {
	return &Filter{
		Type:      "javascript",
		Dimension: dimension,
		Function:  function,
	}
}

func FilterSearch(dimension string, caseSensitive bool, values ...string) *Filter {
	query := &SearchFilterQuery{
		CaseSensitive: caseSensitive,
	}

	if len(values) == 1 {
		query.Type = "contains"
		query.Value = values[0]
	} else {
		query.Type = "fragment"
		query.Values = values
	}

	return &Filter{
		Type:      "search",
		Dimension: dimension,
		Query:     query,
	}
}

func FilterLike(dimension, pattern string) *Filter {
	return &Filter{
		Type:      "like",
		Dimension: dimension,
		Pattern:   pattern,
	}
}

func FilterAnd(filters ...*Filter) *Filter {
	return joinFilters(filters, "and")
}

func FilterOr(filters ...*Filter) *Filter {
	return joinFilters(filters, "or")
}

func FilterNot(filter *Filter) *Filter {
	return &Filter{
		Type:  "not",
		Field: filter,
	}
}

func joinFilters(filters []*Filter, connector string) *Filter {
	// Remove null filters.
	p := 0
	for _, f := range filters {
		if f != nil {
			filters[p] = f
			p++
		}
	}
	filters = filters[0:p]

	fLen := len(filters)
	if fLen == 0 {
		return nil
	}
	if fLen == 1 {
		return filters[0]
	}

	return &Filter{
		Type:   connector,
		Fields: filters,
	}
}
