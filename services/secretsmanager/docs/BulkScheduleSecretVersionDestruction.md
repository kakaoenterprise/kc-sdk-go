# BulkScheduleSecretVersionDestruction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Versions** | **[]int32** | 처리할 시크릿 버전 목록 &lt;br/&gt;- [List secret versions](https://docs.kakaocloud.com/openapi/security/secrets-manager/list-secret-versions)에서 확인 | 
**DestructionScheduledAt** | **time.Time** | 시크릿 버전 폐기 예약 일시 &lt;br/&gt; - ISO 8601 형식 &lt;br/&gt; - UTC 기준 | 

## Methods

### NewBulkScheduleSecretVersionDestruction

`func NewBulkScheduleSecretVersionDestruction(versions []int32, destructionScheduledAt time.Time, ) *BulkScheduleSecretVersionDestruction`

NewBulkScheduleSecretVersionDestruction instantiates a new BulkScheduleSecretVersionDestruction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkScheduleSecretVersionDestructionWithDefaults

`func NewBulkScheduleSecretVersionDestructionWithDefaults() *BulkScheduleSecretVersionDestruction`

NewBulkScheduleSecretVersionDestructionWithDefaults instantiates a new BulkScheduleSecretVersionDestruction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersions

`func (o *BulkScheduleSecretVersionDestruction) GetVersions() []int32`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *BulkScheduleSecretVersionDestruction) GetVersionsOk() (*[]int32, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *BulkScheduleSecretVersionDestruction) SetVersions(v []int32)`

SetVersions sets Versions field to given value.


### GetDestructionScheduledAt

`func (o *BulkScheduleSecretVersionDestruction) GetDestructionScheduledAt() time.Time`

GetDestructionScheduledAt returns the DestructionScheduledAt field if non-nil, zero value otherwise.

### GetDestructionScheduledAtOk

`func (o *BulkScheduleSecretVersionDestruction) GetDestructionScheduledAtOk() (*time.Time, bool)`

GetDestructionScheduledAtOk returns a tuple with the DestructionScheduledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestructionScheduledAt

`func (o *BulkScheduleSecretVersionDestruction) SetDestructionScheduledAt(v time.Time)`

SetDestructionScheduledAt sets DestructionScheduledAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


