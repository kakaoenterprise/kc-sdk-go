# BulkCancelScheduledKeyVersionDestructionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Versions** | [**Result**](Result.md) | 키 버전별 처리 결과 | 

## Methods

### NewBulkCancelScheduledKeyVersionDestructionResponse

`func NewBulkCancelScheduledKeyVersionDestructionResponse(versions Result, ) *BulkCancelScheduledKeyVersionDestructionResponse`

NewBulkCancelScheduledKeyVersionDestructionResponse instantiates a new BulkCancelScheduledKeyVersionDestructionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkCancelScheduledKeyVersionDestructionResponseWithDefaults

`func NewBulkCancelScheduledKeyVersionDestructionResponseWithDefaults() *BulkCancelScheduledKeyVersionDestructionResponse`

NewBulkCancelScheduledKeyVersionDestructionResponseWithDefaults instantiates a new BulkCancelScheduledKeyVersionDestructionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersions

`func (o *BulkCancelScheduledKeyVersionDestructionResponse) GetVersions() Result`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *BulkCancelScheduledKeyVersionDestructionResponse) GetVersionsOk() (*Result, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *BulkCancelScheduledKeyVersionDestructionResponse) SetVersions(v Result)`

SetVersions sets Versions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


