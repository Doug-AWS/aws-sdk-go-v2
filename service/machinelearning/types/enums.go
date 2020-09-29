// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type Algorithm string

// Enum values for Algorithm
const (
	AlgorithmSgd Algorithm = "sgd"
)

type BatchPredictionFilterVariable string

// Enum values for BatchPredictionFilterVariable
const (
	BatchPredictionFilterVariableCreated_at      BatchPredictionFilterVariable = "CreatedAt"
	BatchPredictionFilterVariableLast_updated_at BatchPredictionFilterVariable = "LastUpdatedAt"
	BatchPredictionFilterVariableStatus          BatchPredictionFilterVariable = "Status"
	BatchPredictionFilterVariableName            BatchPredictionFilterVariable = "Name"
	BatchPredictionFilterVariableIam_user        BatchPredictionFilterVariable = "IAMUser"
	BatchPredictionFilterVariableMl_model_id     BatchPredictionFilterVariable = "MLModelId"
	BatchPredictionFilterVariableDatasource_id   BatchPredictionFilterVariable = "DataSourceId"
	BatchPredictionFilterVariableData_uri        BatchPredictionFilterVariable = "DataURI"
)

type DataSourceFilterVariable string

// Enum values for DataSourceFilterVariable
const (
	DataSourceFilterVariableCreated_at      DataSourceFilterVariable = "CreatedAt"
	DataSourceFilterVariableLast_updated_at DataSourceFilterVariable = "LastUpdatedAt"
	DataSourceFilterVariableStatus          DataSourceFilterVariable = "Status"
	DataSourceFilterVariableName            DataSourceFilterVariable = "Name"
	DataSourceFilterVariableData_uri        DataSourceFilterVariable = "DataLocationS3"
	DataSourceFilterVariableIam_user        DataSourceFilterVariable = "IAMUser"
)

type DetailsAttributes string

// Enum values for DetailsAttributes
const (
	DetailsAttributesPredictive_model_type DetailsAttributes = "PredictiveModelType"
	DetailsAttributesAlgorithm             DetailsAttributes = "Algorithm"
)

type EntityStatus string

// Enum values for EntityStatus
const (
	EntityStatusPending    EntityStatus = "PENDING"
	EntityStatusInprogress EntityStatus = "INPROGRESS"
	EntityStatusFailed     EntityStatus = "FAILED"
	EntityStatusCompleted  EntityStatus = "COMPLETED"
	EntityStatusDeleted    EntityStatus = "DELETED"
)

type EvaluationFilterVariable string

// Enum values for EvaluationFilterVariable
const (
	EvaluationFilterVariableCreated_at      EvaluationFilterVariable = "CreatedAt"
	EvaluationFilterVariableLast_updated_at EvaluationFilterVariable = "LastUpdatedAt"
	EvaluationFilterVariableStatus          EvaluationFilterVariable = "Status"
	EvaluationFilterVariableName            EvaluationFilterVariable = "Name"
	EvaluationFilterVariableIam_user        EvaluationFilterVariable = "IAMUser"
	EvaluationFilterVariableMl_model_id     EvaluationFilterVariable = "MLModelId"
	EvaluationFilterVariableDatasource_id   EvaluationFilterVariable = "DataSourceId"
	EvaluationFilterVariableData_uri        EvaluationFilterVariable = "DataURI"
)

type MLModelFilterVariable string

// Enum values for MLModelFilterVariable
const (
	MLModelFilterVariableCreated_at                MLModelFilterVariable = "CreatedAt"
	MLModelFilterVariableLast_updated_at           MLModelFilterVariable = "LastUpdatedAt"
	MLModelFilterVariableStatus                    MLModelFilterVariable = "Status"
	MLModelFilterVariableName                      MLModelFilterVariable = "Name"
	MLModelFilterVariableIam_user                  MLModelFilterVariable = "IAMUser"
	MLModelFilterVariableTraining_datasource_id    MLModelFilterVariable = "TrainingDataSourceId"
	MLModelFilterVariableReal_time_endpoint_status MLModelFilterVariable = "RealtimeEndpointStatus"
	MLModelFilterVariableMl_model_type             MLModelFilterVariable = "MLModelType"
	MLModelFilterVariableAlgorithm                 MLModelFilterVariable = "Algorithm"
	MLModelFilterVariableTraining_data_uri         MLModelFilterVariable = "TrainingDataURI"
)

type MLModelType string

// Enum values for MLModelType
const (
	MLModelTypeRegression MLModelType = "REGRESSION"
	MLModelTypeBinary     MLModelType = "BINARY"
	MLModelTypeMulticlass MLModelType = "MULTICLASS"
)

type RealtimeEndpointStatus string

// Enum values for RealtimeEndpointStatus
const (
	RealtimeEndpointStatusNone     RealtimeEndpointStatus = "NONE"
	RealtimeEndpointStatusReady    RealtimeEndpointStatus = "READY"
	RealtimeEndpointStatusUpdating RealtimeEndpointStatus = "UPDATING"
	RealtimeEndpointStatusFailed   RealtimeEndpointStatus = "FAILED"
)

type SortOrder string

// Enum values for SortOrder
const (
	SortOrderAsc SortOrder = "asc"
	SortOrderDsc SortOrder = "dsc"
)

type TaggableResourceType string

// Enum values for TaggableResourceType
const (
	TaggableResourceTypeBatch_prediction TaggableResourceType = "BatchPrediction"
	TaggableResourceTypeDatasource       TaggableResourceType = "DataSource"
	TaggableResourceTypeEvaluation       TaggableResourceType = "Evaluation"
	TaggableResourceTypeMl_model         TaggableResourceType = "MLModel"
)
