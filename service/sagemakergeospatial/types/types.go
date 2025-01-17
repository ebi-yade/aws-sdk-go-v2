// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// The following types satisfy this interface:
//
//	AreaOfInterestMemberAreaOfInterestGeometry
type AreaOfInterest interface {
	isAreaOfInterest()
}

type AreaOfInterestMemberAreaOfInterestGeometry struct {
	Value AreaOfInterestGeometry

	noSmithyDocumentSerde
}

func (*AreaOfInterestMemberAreaOfInterestGeometry) isAreaOfInterest() {}

// The following types satisfy this interface:
//
//	AreaOfInterestGeometryMemberMultiPolygonGeometry
//	AreaOfInterestGeometryMemberPolygonGeometry
type AreaOfInterestGeometry interface {
	isAreaOfInterestGeometry()
}

type AreaOfInterestGeometryMemberMultiPolygonGeometry struct {
	Value MultiPolygonGeometryInput

	noSmithyDocumentSerde
}

func (*AreaOfInterestGeometryMemberMultiPolygonGeometry) isAreaOfInterestGeometry() {}

type AreaOfInterestGeometryMemberPolygonGeometry struct {
	Value PolygonGeometryInput

	noSmithyDocumentSerde
}

func (*AreaOfInterestGeometryMemberPolygonGeometry) isAreaOfInterestGeometry() {}

type AssetValue struct {

	//
	Href *string

	noSmithyDocumentSerde
}

type BandMathConfigInput struct {

	//
	CustomIndices *CustomIndicesInput

	//
	PredefinedIndices []string

	noSmithyDocumentSerde
}

type CloudMaskingConfigInput struct {
	noSmithyDocumentSerde
}

type CloudRemovalConfigInput struct {

	// The name of the algorithm used for cloud removal.
	AlgorithmName AlgorithmNameCloudRemoval

	// The interpolation value you provide for cloud removal.
	InterpolationValue *string

	//
	TargetBands []string

	noSmithyDocumentSerde
}

type CustomIndicesInput struct {

	//
	Operations []Operation

	noSmithyDocumentSerde
}

// The structure representing the errors in an EarthObservationJob.
type EarthObservationJobErrorDetails struct {

	//
	Message *string

	//
	Type EarthObservationJobErrorType

	noSmithyDocumentSerde
}

type EoCloudCoverInput struct {

	//
	//
	// This member is required.
	LowerBound *float32

	//
	//
	// This member is required.
	UpperBound *float32

	noSmithyDocumentSerde
}

// The following types satisfy this interface:
//
//	EojDataSourceConfigInputMemberS3Data
type EojDataSourceConfigInput interface {
	isEojDataSourceConfigInput()
}

type EojDataSourceConfigInputMemberS3Data struct {
	Value S3DataInput

	noSmithyDocumentSerde
}

func (*EojDataSourceConfigInputMemberS3Data) isEojDataSourceConfigInput() {}

// The structure for returning the export error details in a
// GetEarthObservationJob.
type ExportErrorDetails struct {

	//
	ExportResults *ExportErrorDetailsOutput

	//
	ExportSourceImages *ExportErrorDetailsOutput

	noSmithyDocumentSerde
}

type ExportErrorDetailsOutput struct {

	//
	Message *string

	//
	Type ExportErrorType

	noSmithyDocumentSerde
}

type ExportS3DataInput struct {

	// The URL to the Amazon S3 data input.
	//
	// This member is required.
	S3Uri *string

	// The Amazon Key Management Service (KMS) key ID for server-side encryption.
	KmsKeyId *string

	noSmithyDocumentSerde
}

// An object containing information about the output file.
type ExportVectorEnrichmentJobOutputConfig struct {

	//
	//
	// This member is required.
	S3Data *VectorEnrichmentJobS3Data

	noSmithyDocumentSerde
}

// The structure representing the filters supported by a RasterDataCollection.
type Filter struct {

	// The name of the filter.
	//
	// This member is required.
	Name *string

	// The type of the filter being used.
	//
	// This member is required.
	Type *string

	// The maximum value of the filter.
	Maximum *float32

	// The minimum value of the filter.
	Minimum *float32

	noSmithyDocumentSerde
}

type Geometry struct {

	//
	//
	// This member is required.
	Coordinates [][][]float64

	//
	//
	// This member is required.
	Type *string

	noSmithyDocumentSerde
}

// Input configuration information for the geomosaic.
type GeoMosaicConfigInput struct {

	// The name of the algorithm being used for geomosaic.
	AlgorithmName AlgorithmNameGeoMosaic

	// The target bands for geomosaic.
	TargetBands []string

	noSmithyDocumentSerde
}

// Input configuration information.
type InputConfigInput struct {

	// The location of the input data.>
	DataSourceConfig EojDataSourceConfigInput

	// The Amazon Resource Name (ARN) of the previous Earth Observation job.
	PreviousEarthObservationJobArn *string

	//
	RasterDataCollectionQuery *RasterDataCollectionQueryInput

	noSmithyDocumentSerde
}

// The InputConfig for an EarthObservationJob response.
type InputConfigOutput struct {

	// The location of the input data.
	DataSourceConfig EojDataSourceConfigInput

	// The Amazon Resource Name (ARN) of the previous Earth Observation job.
	PreviousEarthObservationJobArn *string

	//
	RasterDataCollectionQuery *RasterDataCollectionQueryOutput

	noSmithyDocumentSerde
}

// Structure representing the items in the response for SearchRasterDataCollection.
type ItemSource struct {

	//
	//
	// This member is required.
	DateTime *time.Time

	//
	//
	// This member is required.
	Geometry *Geometry

	// A unique Id for the source item.
	//
	// This member is required.
	Id *string

	//
	Assets map[string]AssetValue

	//
	Properties *Properties

	noSmithyDocumentSerde
}

// The input structure for the JobConfig in an EarthObservationJob.
//
// The following types satisfy this interface:
//
//	JobConfigInputMemberBandMathConfig
//	JobConfigInputMemberCloudMaskingConfig
//	JobConfigInputMemberCloudRemovalConfig
//	JobConfigInputMemberGeoMosaicConfig
//	JobConfigInputMemberLandCoverSegmentationConfig
//	JobConfigInputMemberResamplingConfig
//	JobConfigInputMemberStackConfig
//	JobConfigInputMemberTemporalStatisticsConfig
//	JobConfigInputMemberZonalStatisticsConfig
type JobConfigInput interface {
	isJobConfigInput()
}

type JobConfigInputMemberBandMathConfig struct {
	Value BandMathConfigInput

	noSmithyDocumentSerde
}

func (*JobConfigInputMemberBandMathConfig) isJobConfigInput() {}

// An object containing information about the job configuration for cloud masking.
type JobConfigInputMemberCloudMaskingConfig struct {
	Value CloudMaskingConfigInput

	noSmithyDocumentSerde
}

func (*JobConfigInputMemberCloudMaskingConfig) isJobConfigInput() {}

// An object containing information about the job configuration for cloud removal.
type JobConfigInputMemberCloudRemovalConfig struct {
	Value CloudRemovalConfigInput

	noSmithyDocumentSerde
}

func (*JobConfigInputMemberCloudRemovalConfig) isJobConfigInput() {}

// An object containing information about the job configuration for geomosaic.
type JobConfigInputMemberGeoMosaicConfig struct {
	Value GeoMosaicConfigInput

	noSmithyDocumentSerde
}

func (*JobConfigInputMemberGeoMosaicConfig) isJobConfigInput() {}

// An object containing information about the job configuration for land cover
// segmentation.
type JobConfigInputMemberLandCoverSegmentationConfig struct {
	Value LandCoverSegmentationConfigInput

	noSmithyDocumentSerde
}

func (*JobConfigInputMemberLandCoverSegmentationConfig) isJobConfigInput() {}

// An object containing information about the job configuration for resampling.
type JobConfigInputMemberResamplingConfig struct {
	Value ResamplingConfigInput

	noSmithyDocumentSerde
}

func (*JobConfigInputMemberResamplingConfig) isJobConfigInput() {}

type JobConfigInputMemberStackConfig struct {
	Value StackConfigInput

	noSmithyDocumentSerde
}

func (*JobConfigInputMemberStackConfig) isJobConfigInput() {}

// An object containing information about the job configuration for temporal
// statistics.
type JobConfigInputMemberTemporalStatisticsConfig struct {
	Value TemporalStatisticsConfigInput

	noSmithyDocumentSerde
}

func (*JobConfigInputMemberTemporalStatisticsConfig) isJobConfigInput() {}

// An object containing information about the job configuration for zonal
// statistics.
type JobConfigInputMemberZonalStatisticsConfig struct {
	Value ZonalStatisticsConfigInput

	noSmithyDocumentSerde
}

func (*JobConfigInputMemberZonalStatisticsConfig) isJobConfigInput() {}

type LandCoverSegmentationConfigInput struct {
	noSmithyDocumentSerde
}

type LandsatCloudCoverLandInput struct {

	//
	//
	// This member is required.
	LowerBound *float32

	//
	//
	// This member is required.
	UpperBound *float32

	noSmithyDocumentSerde
}

// An object containing information about the output file.
type ListEarthObservationJobOutputConfig struct {

	// The Amazon Resource Name (ARN) of the list of the Earth Observation jobs.
	//
	// This member is required.
	Arn *string

	// The creation time.
	//
	// This member is required.
	CreationTime *time.Time

	// The duration of the session, in seconds.
	//
	// This member is required.
	DurationInSeconds *int32

	// The names of the Earth Observation jobs in the list.
	//
	// This member is required.
	Name *string

	//
	//
	// This member is required.
	OperationType *string

	// The status of the list of the Earth Observation jobs.
	//
	// This member is required.
	Status EarthObservationJobStatus

	// Each tag consists of a key and a value.
	Tags map[string]string

	noSmithyDocumentSerde
}

// An object containing information about the output file.
type ListVectorEnrichmentJobOutputConfig struct {

	// The Amazon Resource Name (ARN) of the list of the Vector Enrichment jobs.
	//
	// This member is required.
	Arn *string

	// The creation time.
	//
	// This member is required.
	CreationTime *time.Time

	// The duration of the session, in seconds.
	//
	// This member is required.
	DurationInSeconds *int32

	// The names of the Vector Enrichment jobs in the list.
	//
	// This member is required.
	Name *string

	// The status of the Vector Enrichment jobs list.
	//
	// This member is required.
	Status VectorEnrichmentJobStatus

	// The type of the list of Vector Enrichment jobs.
	//
	// This member is required.
	Type VectorEnrichmentJobType

	// Each tag consists of a key and a value.
	Tags map[string]string

	noSmithyDocumentSerde
}

type MapMatchingConfig struct {

	//
	//
	// This member is required.
	IdAttributeName *string

	// The name of the timestamp attribute.
	//
	// This member is required.
	TimestampAttributeName *string

	// The name of the X-attribute
	//
	// This member is required.
	XAttributeName *string

	// The name of the Y-attribute
	//
	// This member is required.
	YAttributeName *string

	noSmithyDocumentSerde
}

type MultiPolygonGeometryInput struct {

	// The coordinates of the multipolygon geometry.
	//
	// This member is required.
	Coordinates [][][][]float64

	noSmithyDocumentSerde
}

type Operation struct {

	//
	//
	// This member is required.
	Equation *string

	// The name of the operation.
	//
	// This member is required.
	Name *string

	// The type of the operation.
	OutputType OutputType

	noSmithyDocumentSerde
}

// A single EarthObservationJob output band.
type OutputBand struct {

	// The name of the band.
	//
	// This member is required.
	BandName *string

	// The datatype of the output band.
	//
	// This member is required.
	OutputDataType OutputType

	noSmithyDocumentSerde
}

// The response structure for an OutputConfig returned by an
// ExportEarthObservationJob.
type OutputConfigInput struct {

	// Path to Amazon S3 storage location for the output configuration file.
	//
	// This member is required.
	S3Data *ExportS3DataInput

	noSmithyDocumentSerde
}

type OutputResolutionResamplingInput struct {

	//
	//
	// This member is required.
	UserDefined *UserDefined

	noSmithyDocumentSerde
}

type OutputResolutionStackInput struct {

	//
	Predefined PredefinedResolution

	//
	UserDefined *UserDefined

	noSmithyDocumentSerde
}

type PlatformInput struct {

	// The value of the platform.
	//
	// This member is required.
	Value *string

	//
	ComparisonOperator ComparisonOperator

	noSmithyDocumentSerde
}

type PolygonGeometryInput struct {

	//
	//
	// This member is required.
	Coordinates [][][]float64

	noSmithyDocumentSerde
}

type Properties struct {

	//
	EoCloudCover *float32

	//
	LandsatCloudCoverLand *float32

	//
	Platform *string

	//
	ViewOffNadir *float32

	//
	ViewSunAzimuth *float32

	//
	ViewSunElevation *float32

	noSmithyDocumentSerde
}

// The following types satisfy this interface:
//
//	PropertyMemberEoCloudCover
//	PropertyMemberLandsatCloudCoverLand
//	PropertyMemberPlatform
//	PropertyMemberViewOffNadir
//	PropertyMemberViewSunAzimuth
//	PropertyMemberViewSunElevation
type Property interface {
	isProperty()
}

type PropertyMemberEoCloudCover struct {
	Value EoCloudCoverInput

	noSmithyDocumentSerde
}

func (*PropertyMemberEoCloudCover) isProperty() {}

type PropertyMemberLandsatCloudCoverLand struct {
	Value LandsatCloudCoverLandInput

	noSmithyDocumentSerde
}

func (*PropertyMemberLandsatCloudCoverLand) isProperty() {}

type PropertyMemberPlatform struct {
	Value PlatformInput

	noSmithyDocumentSerde
}

func (*PropertyMemberPlatform) isProperty() {}

type PropertyMemberViewOffNadir struct {
	Value ViewOffNadirInput

	noSmithyDocumentSerde
}

func (*PropertyMemberViewOffNadir) isProperty() {}

type PropertyMemberViewSunAzimuth struct {
	Value ViewSunAzimuthInput

	noSmithyDocumentSerde
}

func (*PropertyMemberViewSunAzimuth) isProperty() {}

type PropertyMemberViewSunElevation struct {
	Value ViewSunElevationInput

	noSmithyDocumentSerde
}

func (*PropertyMemberViewSunElevation) isProperty() {}

type PropertyFilter struct {

	//
	//
	// This member is required.
	Property Property

	noSmithyDocumentSerde
}

type PropertyFilters struct {

	//
	LogicalOperator LogicalOperator

	//
	Properties []PropertyFilter

	noSmithyDocumentSerde
}

// Response object containing details for a specific RasterDataCollection.
type RasterDataCollectionMetadata struct {

	// The Amazon Resource Name (ARN) of the raster data collection.
	//
	// This member is required.
	Arn *string

	// A description of the raster data collection.
	//
	// This member is required.
	Description *string

	// The name of the raster data collection.
	//
	// This member is required.
	Name *string

	//
	//
	// This member is required.
	SupportedFilters []Filter

	// The type of raster data collection.
	//
	// This member is required.
	Type DataCollectionType

	// The description URL of the raster data collection.
	DescriptionPageUrl *string

	// Each tag consists of a key and a value.
	Tags map[string]string

	noSmithyDocumentSerde
}

type RasterDataCollectionQueryInput struct {

	// The Amazon Resource Name (ARN) of the raster data collection.
	//
	// This member is required.
	RasterDataCollectionArn *string

	//
	//
	// This member is required.
	TimeRangeFilter *TimeRangeFilterInput

	// The area of interest being queried for the raster data collection.
	AreaOfInterest AreaOfInterest

	//
	PropertyFilters *PropertyFilters

	noSmithyDocumentSerde
}

type RasterDataCollectionQueryOutput struct {

	//
	//
	// This member is required.
	RasterDataCollectionArn *string

	// The name of the raster data collection.
	//
	// This member is required.
	RasterDataCollectionName *string

	//
	//
	// This member is required.
	TimeRangeFilter *TimeRangeFilterInput

	//
	AreaOfInterest AreaOfInterest

	//
	PropertyFilters *PropertyFilters

	noSmithyDocumentSerde
}

// This is a RasterDataCollectionQueryInput containing AreaOfInterest, Time Range
// filter and Property filters.
type RasterDataCollectionQueryWithBandFilterInput struct {

	//
	//
	// This member is required.
	TimeRangeFilter *TimeRangeFilterInput

	//
	AreaOfInterest AreaOfInterest

	//
	BandFilter []string

	//
	PropertyFilters *PropertyFilters

	noSmithyDocumentSerde
}

type ResamplingConfigInput struct {

	//
	//
	// This member is required.
	OutputResolution *OutputResolutionResamplingInput

	// The name of the algorithm used for resampling.
	AlgorithmName AlgorithmNameResampling

	//
	TargetBands []string

	noSmithyDocumentSerde
}

type ReverseGeocodingConfig struct {

	//
	//
	// This member is required.
	XAttributeName *string

	//
	//
	// This member is required.
	YAttributeName *string

	noSmithyDocumentSerde
}

// Path to Amazon S3 storage location for input data.
type S3DataInput struct {

	//
	//
	// This member is required.
	MetadataProvider MetadataProvider

	// The URL to the Amazon S3 input.
	//
	// This member is required.
	S3Uri *string

	// The Amazon Key Management Service (KMS) key ID for server-side encryption.
	KmsKeyId *string

	noSmithyDocumentSerde
}

type StackConfigInput struct {

	//
	OutputResolution *OutputResolutionStackInput

	//
	TargetBands []string

	noSmithyDocumentSerde
}

type TemporalStatisticsConfigInput struct {

	//
	//
	// This member is required.
	Statistics []TemporalStatistics

	//
	GroupBy GroupBy

	//
	TargetBands []string

	noSmithyDocumentSerde
}

// The input for the time-range filter.
type TimeRangeFilterInput struct {

	// The end time for the time-range filter.
	//
	// This member is required.
	EndTime *time.Time

	// The start time for the time-range filter.
	//
	// This member is required.
	StartTime *time.Time

	noSmithyDocumentSerde
}

type UserDefined struct {

	//
	//
	// This member is required.
	Unit Unit

	//
	//
	// This member is required.
	Value *float32

	noSmithyDocumentSerde
}

// It contains configs such as ReverseGeocodingConfig and MapMatchingConfig.
//
// The following types satisfy this interface:
//
//	VectorEnrichmentJobConfigMemberMapMatchingConfig
//	VectorEnrichmentJobConfigMemberReverseGeocodingConfig
type VectorEnrichmentJobConfig interface {
	isVectorEnrichmentJobConfig()
}

type VectorEnrichmentJobConfigMemberMapMatchingConfig struct {
	Value MapMatchingConfig

	noSmithyDocumentSerde
}

func (*VectorEnrichmentJobConfigMemberMapMatchingConfig) isVectorEnrichmentJobConfig() {}

type VectorEnrichmentJobConfigMemberReverseGeocodingConfig struct {
	Value ReverseGeocodingConfig

	noSmithyDocumentSerde
}

func (*VectorEnrichmentJobConfigMemberReverseGeocodingConfig) isVectorEnrichmentJobConfig() {}

// The following types satisfy this interface:
//
//	VectorEnrichmentJobDataSourceConfigInputMemberS3Data
type VectorEnrichmentJobDataSourceConfigInput interface {
	isVectorEnrichmentJobDataSourceConfigInput()
}

type VectorEnrichmentJobDataSourceConfigInputMemberS3Data struct {
	Value VectorEnrichmentJobS3Data

	noSmithyDocumentSerde
}

func (*VectorEnrichmentJobDataSourceConfigInputMemberS3Data) isVectorEnrichmentJobDataSourceConfigInput() {
}

// VectorEnrichmentJob error details in response from GetVectorEnrichmentJob.
type VectorEnrichmentJobErrorDetails struct {

	// A message that you define and then is processed and rendered by the Vector
	// Enrichment job when the error occurs.
	ErrorMessage *string

	// The type of error generated during the Vector Enrichment job.
	ErrorType VectorEnrichmentJobErrorType

	noSmithyDocumentSerde
}

// VectorEnrichmentJob export error details in response from
// GetVectorEnrichmentJob.
type VectorEnrichmentJobExportErrorDetails struct {

	// The message providing details about the errors generated during the Vector
	// Enrichment job.
	Message *string

	//
	Type VectorEnrichmentJobExportErrorType

	noSmithyDocumentSerde
}

// The input structure for the InputConfig in a VectorEnrichmentJob.
type VectorEnrichmentJobInputConfig struct {

	//
	//
	// This member is required.
	DataSourceConfig VectorEnrichmentJobDataSourceConfigInput

	//
	//
	// This member is required.
	DocumentType VectorEnrichmentJobDocumentType

	noSmithyDocumentSerde
}

// The Amazon S3 data for the Vector Enrichment job.
type VectorEnrichmentJobS3Data struct {

	// The URL to the Amazon S3 data for the Vector Enrichment job.
	//
	// This member is required.
	S3Uri *string

	// The Amazon Key Management Service (KMS) key ID for server-side encryption.
	KmsKeyId *string

	noSmithyDocumentSerde
}

type ViewOffNadirInput struct {

	//
	//
	// This member is required.
	LowerBound *float32

	//
	//
	// This member is required.
	UpperBound *float32

	noSmithyDocumentSerde
}

type ViewSunAzimuthInput struct {

	//
	//
	// This member is required.
	LowerBound *float32

	//
	//
	// This member is required.
	UpperBound *float32

	noSmithyDocumentSerde
}

type ViewSunElevationInput struct {

	// The lower bound to view the sun elevation.
	//
	// This member is required.
	LowerBound *float32

	// The upper bound to view the sun elevation.
	//
	// This member is required.
	UpperBound *float32

	noSmithyDocumentSerde
}

type ZonalStatisticsConfigInput struct {

	//
	//
	// This member is required.
	Statistics []ZonalStatistics

	//
	//
	// This member is required.
	ZoneS3Path *string

	//
	TargetBands []string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde

// UnknownUnionMember is returned when a union member is returned over the wire,
// but has an unknown tag.
type UnknownUnionMember struct {
	Tag   string
	Value []byte

	noSmithyDocumentSerde
}

func (*UnknownUnionMember) isAreaOfInterest()                           {}
func (*UnknownUnionMember) isAreaOfInterestGeometry()                   {}
func (*UnknownUnionMember) isEojDataSourceConfigInput()                 {}
func (*UnknownUnionMember) isJobConfigInput()                           {}
func (*UnknownUnionMember) isProperty()                                 {}
func (*UnknownUnionMember) isVectorEnrichmentJobConfig()                {}
func (*UnknownUnionMember) isVectorEnrichmentJobDataSourceConfigInput() {}
