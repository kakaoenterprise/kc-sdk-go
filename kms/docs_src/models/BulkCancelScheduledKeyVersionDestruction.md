# BulkCancelScheduledKeyVersionDestruction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Versions** | **[]int32** | 처리할 키 버전 목록 - [List key versions](/openapi/security/kms/list-key-versions)에서 확인 | 

## Methods

### NewBulkCancelScheduledKeyVersionDestruction

`func NewBulkCancelScheduledKeyVersionDestruction(versions []int32, ) *BulkCancelScheduledKeyVersionDestruction`

NewBulkCancelScheduledKeyVersionDestruction instantiates a new BulkCancelScheduledKeyVersionDestruction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkCancelScheduledKeyVersionDestructionWithDefaults

`func NewBulkCancelScheduledKeyVersionDestructionWithDefaults() *BulkCancelScheduledKeyVersionDestruction`

NewBulkCancelScheduledKeyVersionDestructionWithDefaults instantiates a new BulkCancelScheduledKeyVersionDestruction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersions

`func (o *BulkCancelScheduledKeyVersionDestruction) GetVersions() []int32`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *BulkCancelScheduledKeyVersionDestruction) GetVersionsOk() (*[]int32, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *BulkCancelScheduledKeyVersionDestruction) SetVersions(v []int32)`

SetVersions sets Versions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


