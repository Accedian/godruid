package godruid

import (
	"encoding/json"
)

type Aggregation struct {
	Type        string       `json:"type"`
	Name        string       `json:"name,omitempty"`
	FieldName   string       `json:"fieldName,omitempty"`
	FieldNames  []string     `json:"fieldNames,omitempty"`
	FnAggregate string       `json:"fnAggregate,omitempty"`
	FnCombine   string       `json:"fnCombine,omitempty"`
	FnReset     string       `json:"fnReset,omitempty"`
	ByRow       bool         `json:"byRow,omitempty"`
	Filter      *Filter      `json:"filter,omitempty"`
	Resolution  int32        `json:"resolution,omitempty"`
	NumBuckets  int32        `json:"numBuckets,omitempty"`
	LowerLimit  string       `json:"lowerLimit,omitempty"`
	UpperLimit  string       `json:"upperLimit,omitempty"`
	Aggregator  *Aggregation `json:"aggregator,omitempty"`
	Round       bool         `json:"round,omitempty"`
	K           int32        `json:"k,omitempty"` // druid-datasketches extension
}

func AggRawJson(rawJson string) *Aggregation {
	agg := &Aggregation{}
	json.Unmarshal([]byte(rawJson), agg)
	return agg
}

func AggCount(name string) *Aggregation {
	return &Aggregation{
		Type: "count",
		Name: name,
	}
}

func AggLongSum(name, fieldName string) *Aggregation {
	return &Aggregation{
		Type:      "longSum",
		Name:      name,
		FieldName: fieldName,
	}
}

func AggDoubleSum(name, fieldName string) *Aggregation {
	return &Aggregation{
		Type:      "doubleSum",
		Name:      name,
		FieldName: fieldName,
	}
}

func AggMin(name, fieldName string) *Aggregation {
	return &Aggregation{
		Type:      "min",
		Name:      name,
		FieldName: fieldName,
	}
}

func AggMax(name, fieldName string) *Aggregation {
	return &Aggregation{
		Type:      "max",
		Name:      name,
		FieldName: fieldName,
	}
}

func AggDoubleMax(name, fieldName string) *Aggregation {
	return &Aggregation{
		Type:      "doubleMax",
		Name:      name,
		FieldName: fieldName,
	}
}

func AggDoubleMin(name, fieldName string) *Aggregation {
	return &Aggregation{
		Type:      "doubleMin",
		Name:      name,
		FieldName: fieldName,
	}
}

func AggLongMin(name, fieldName string) *Aggregation {
	return &Aggregation{
		Type:      "longMin",
		Name:      name,
		FieldName: fieldName,
	}
}

func AggLongMax(name, fieldName string) *Aggregation {
	return &Aggregation{
		Type:      "longMax",
		Name:      name,
		FieldName: fieldName,
	}
}

func AggFiltered(filter *Filter, aggregator *Aggregation) Aggregation {
	return Aggregation{
		Type:       "filtered",
		Filter:     filter,
		Aggregator: aggregator,
	}
}

func AggHistoFold(name string, fieldName string, resolution int32, numBuckets int32, lowerLimit string, upperLimit string) Aggregation {
	return Aggregation{
		Type:       "approxHistogramFold",
		Name:       name,
		Resolution: resolution,
		NumBuckets: numBuckets,
		FieldName:  fieldName,
		LowerLimit: lowerLimit,
		UpperLimit: upperLimit,
	}
}

func AggJavaScript(name, fnAggregate, fnCombine, fnReset string, fieldNames []string) Aggregation {
	return Aggregation{
		Type:        "javascript",
		Name:        name,
		FieldNames:  fieldNames,
		FnAggregate: fnAggregate,
		FnCombine:   fnCombine,
		FnReset:     fnReset,
	}
}

func AggCardinality(name string, fieldNames []string, byRow ...bool) Aggregation {
	isByRow := false
	if len(byRow) != 0 {
		isByRow = byRow[0]
	}
	return Aggregation{
		Type:       "cardinality",
		Name:       name,
		FieldNames: fieldNames,
		ByRow:      isByRow,
	}
}

// druid-stats extension
func ExtAggVariance(name, fieldName string) *Aggregation {
	return &Aggregation{
		Type:      "variance",
		Name:      name,
		FieldName: fieldName,
	}
}

// druid-datasketches extension
func ExtAggQuantile(name string, fieldName string, k int32) *Aggregation {

	retAgg := Aggregation{
		Type:      "quantilesDoublesSketch",
		Name:      name,
		FieldName: fieldName,
	}

	if k > 0 {
		retAgg.K = k
	}

	return &retAgg
}

// druid-datasketches extension
func ExtAggThetaSketch(name string, fieldName string) *Aggregation {
	return &Aggregation{
		Type:      "thetaSketch",
		Name:      name,
		FieldName: fieldName,
	}
}
