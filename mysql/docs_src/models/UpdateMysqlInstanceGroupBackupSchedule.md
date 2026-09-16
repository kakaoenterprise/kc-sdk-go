# UpdateMysqlInstanceGroupBackupSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**NullableBackupScheduleType**](BackupScheduleType.md) | 백업 주기 유형 | [optional] 
**StartTime** | Pointer to **NullableString** | 백업 시작 시간 - &#x60;HH:mm&#x60; 형식 - UTC 기준 | [optional] 
**ExpiryDuration** | Pointer to **NullableInt32** | 백업 보관 기간 (단위: 일) | [optional] 
**Enabled** | **bool** | 백업 스케줄 활성화 여부 | 

## Methods

### NewUpdateMysqlInstanceGroupBackupSchedule

`func NewUpdateMysqlInstanceGroupBackupSchedule(enabled bool, ) *UpdateMysqlInstanceGroupBackupSchedule`

NewUpdateMysqlInstanceGroupBackupSchedule instantiates a new UpdateMysqlInstanceGroupBackupSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMysqlInstanceGroupBackupScheduleWithDefaults

`func NewUpdateMysqlInstanceGroupBackupScheduleWithDefaults() *UpdateMysqlInstanceGroupBackupSchedule`

NewUpdateMysqlInstanceGroupBackupScheduleWithDefaults instantiates a new UpdateMysqlInstanceGroupBackupSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UpdateMysqlInstanceGroupBackupSchedule) GetType() BackupScheduleType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateMysqlInstanceGroupBackupSchedule) GetTypeOk() (*BackupScheduleType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateMysqlInstanceGroupBackupSchedule) SetType(v BackupScheduleType)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateMysqlInstanceGroupBackupSchedule) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *UpdateMysqlInstanceGroupBackupSchedule) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *UpdateMysqlInstanceGroupBackupSchedule) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetStartTime

`func (o *UpdateMysqlInstanceGroupBackupSchedule) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *UpdateMysqlInstanceGroupBackupSchedule) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *UpdateMysqlInstanceGroupBackupSchedule) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *UpdateMysqlInstanceGroupBackupSchedule) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *UpdateMysqlInstanceGroupBackupSchedule) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *UpdateMysqlInstanceGroupBackupSchedule) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetExpiryDuration

`func (o *UpdateMysqlInstanceGroupBackupSchedule) GetExpiryDuration() int32`

GetExpiryDuration returns the ExpiryDuration field if non-nil, zero value otherwise.

### GetExpiryDurationOk

`func (o *UpdateMysqlInstanceGroupBackupSchedule) GetExpiryDurationOk() (*int32, bool)`

GetExpiryDurationOk returns a tuple with the ExpiryDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDuration

`func (o *UpdateMysqlInstanceGroupBackupSchedule) SetExpiryDuration(v int32)`

SetExpiryDuration sets ExpiryDuration field to given value.

### HasExpiryDuration

`func (o *UpdateMysqlInstanceGroupBackupSchedule) HasExpiryDuration() bool`

HasExpiryDuration returns a boolean if a field has been set.

### SetExpiryDurationNil

`func (o *UpdateMysqlInstanceGroupBackupSchedule) SetExpiryDurationNil(b bool)`

 SetExpiryDurationNil sets the value for ExpiryDuration to be an explicit nil

### UnsetExpiryDuration
`func (o *UpdateMysqlInstanceGroupBackupSchedule) UnsetExpiryDuration()`

UnsetExpiryDuration ensures that no value is present for ExpiryDuration, not even an explicit nil
### GetEnabled

`func (o *UpdateMysqlInstanceGroupBackupSchedule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateMysqlInstanceGroupBackupSchedule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateMysqlInstanceGroupBackupSchedule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


