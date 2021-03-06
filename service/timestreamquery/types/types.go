// Code generated by smithy-go-codegen DO NOT EDIT.

package types

// Contains the meta data for query results such as the column names, data types,
// and other attributes.
type ColumnInfo struct {

	// The data type of the result set column. The data type can be a scalar or
	// complex. Scalar data types are integers, strings, doubles, booleans, and others.
	// Complex data types are types such as arrays, rows, and others.
	//
	// This member is required.
	Type *Type

	// The name of the result set column. The name of the result set is available for
	// columns of all data types except for arrays.
	Name *string
}

// Datum represents a single data point in a query result.
type Datum struct {

	// Indicates if the data point is an array.
	ArrayValue []*Datum

	// Indicates if the data point is null.
	NullValue *bool

	// Indicates if the data point is a row.
	RowValue *Row

	// Indicates if the data point is a scalar value such as integer, string, double,
	// or boolean.
	ScalarValue *string

	// Indicates if the data point is of timeseries data type.
	TimeSeriesValue []*TimeSeriesDataPoint
}

// Represents an available endpoint against which to make API calls agaisnt, as
// well as the TTL for that endpoint.
type Endpoint struct {

	// An endpoint address.
	//
	// This member is required.
	Address *string

	// The TTL for the endpoint, in minutes.
	//
	// This member is required.
	CachePeriodInMinutes *int64
}

// Represents a single row in the query results.
type Row struct {

	// List of data points in a single row of the result set.
	//
	// This member is required.
	Data []*Datum
}

// The timeseries datatype represents the values of a measure over time. A time
// series is an array of rows of timestamps and measure values, with rows sorted in
// ascending order of time. A TimeSeriesDataPoint is a single data point in the
// timeseries. It represents a tuple of (time, measure value) in a timeseries.
type TimeSeriesDataPoint struct {

	// The timestamp when the measure value was collected.
	//
	// This member is required.
	Time *string

	// The measure value for the data point.
	//
	// This member is required.
	Value *Datum
}

// Contains the data type of a column in a query result set. The data type can be
// scalar or complex. The supported scalar data types are integers, boolean,
// string, double, timestamp, date, time, and intervals. The supported complex data
// types are arrays, rows, and timeseries.
type Type struct {

	// Indicates if the column is an array.
	ArrayColumnInfo *ColumnInfo

	// Indicates if the column is a row.
	RowColumnInfo []*ColumnInfo

	// Indicates if the column is of type string, integer, boolean, double, timestamp,
	// date, time.
	ScalarType ScalarType

	// Indicates if the column is a timeseries data type.
	TimeSeriesMeasureValueColumnInfo *ColumnInfo
}
