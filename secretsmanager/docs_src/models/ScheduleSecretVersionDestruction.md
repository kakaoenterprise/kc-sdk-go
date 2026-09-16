# ScheduleSecretVersionDestruction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestructionScheduledAt** | **time.Time** | 시크릿 버전 폐기 예약 일시 - ISO 8601 형식 - UTC 기준 | 

## Methods

### NewScheduleSecretVersionDestruction

`func NewScheduleSecretVersionDestruction(destructionScheduledAt time.Time, ) *ScheduleSecretVersionDestruction`

NewScheduleSecretVersionDestruction instantiates a new ScheduleSecretVersionDestruction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleSecretVersionDestructionWithDefaults

`func NewScheduleSecretVersionDestructionWithDefaults() *ScheduleSecretVersionDestruction`

NewScheduleSecretVersionDestructionWithDefaults instantiates a new ScheduleSecretVersionDestruction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestructionScheduledAt

`func (o *ScheduleSecretVersionDestruction) GetDestructionScheduledAt() time.Time`

GetDestructionScheduledAt returns the DestructionScheduledAt field if non-nil, zero value otherwise.

### GetDestructionScheduledAtOk

`func (o *ScheduleSecretVersionDestruction) GetDestructionScheduledAtOk() (*time.Time, bool)`

GetDestructionScheduledAtOk returns a tuple with the DestructionScheduledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestructionScheduledAt

`func (o *ScheduleSecretVersionDestruction) SetDestructionScheduledAt(v time.Time)`

SetDestructionScheduledAt sets DestructionScheduledAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


