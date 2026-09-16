# BulkScheduleKeyVersionDestruction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Versions** | **[]int32** | 처리할 키 버전 목록 - [List key versions](/openapi/security/kms/list-key-versions)에서 확인 | 
**DestructionScheduledAt** | **string** | 키 버전 폐기 예약 일시 - ISO 8601 형식 - UTC 기준 | 

## Methods

### NewBulkScheduleKeyVersionDestruction

`func NewBulkScheduleKeyVersionDestruction(versions []int32, destructionScheduledAt string, ) *BulkScheduleKeyVersionDestruction`

NewBulkScheduleKeyVersionDestruction instantiates a new BulkScheduleKeyVersionDestruction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkScheduleKeyVersionDestructionWithDefaults

`func NewBulkScheduleKeyVersionDestructionWithDefaults() *BulkScheduleKeyVersionDestruction`

NewBulkScheduleKeyVersionDestructionWithDefaults instantiates a new BulkScheduleKeyVersionDestruction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersions

`func (o *BulkScheduleKeyVersionDestruction) GetVersions() []int32`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *BulkScheduleKeyVersionDestruction) GetVersionsOk() (*[]int32, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *BulkScheduleKeyVersionDestruction) SetVersions(v []int32)`

SetVersions sets Versions field to given value.


### GetDestructionScheduledAt

`func (o *BulkScheduleKeyVersionDestruction) GetDestructionScheduledAt() string`

GetDestructionScheduledAt returns the DestructionScheduledAt field if non-nil, zero value otherwise.

### GetDestructionScheduledAtOk

`func (o *BulkScheduleKeyVersionDestruction) GetDestructionScheduledAtOk() (*string, bool)`

GetDestructionScheduledAtOk returns a tuple with the DestructionScheduledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestructionScheduledAt

`func (o *BulkScheduleKeyVersionDestruction) SetDestructionScheduledAt(v string)`

SetDestructionScheduledAt sets DestructionScheduledAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


