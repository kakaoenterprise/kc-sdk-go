# BulkScheduleKeyVersionDestructionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Versions** | [**BulkScheduleKeyVersionDestruction**](BulkScheduleKeyVersionDestruction.md) | 처리할 키 버전 정보와 폐기 예약 일시 | 

## Methods

### NewBulkScheduleKeyVersionDestructionRequest

`func NewBulkScheduleKeyVersionDestructionRequest(versions BulkScheduleKeyVersionDestruction, ) *BulkScheduleKeyVersionDestructionRequest`

NewBulkScheduleKeyVersionDestructionRequest instantiates a new BulkScheduleKeyVersionDestructionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkScheduleKeyVersionDestructionRequestWithDefaults

`func NewBulkScheduleKeyVersionDestructionRequestWithDefaults() *BulkScheduleKeyVersionDestructionRequest`

NewBulkScheduleKeyVersionDestructionRequestWithDefaults instantiates a new BulkScheduleKeyVersionDestructionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersions

`func (o *BulkScheduleKeyVersionDestructionRequest) GetVersions() BulkScheduleKeyVersionDestruction`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *BulkScheduleKeyVersionDestructionRequest) GetVersionsOk() (*BulkScheduleKeyVersionDestruction, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *BulkScheduleKeyVersionDestructionRequest) SetVersions(v BulkScheduleKeyVersionDestruction)`

SetVersions sets Versions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


