# BackupScheduleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**NullableBackupScheduleType**](BackupScheduleType.md) | 백업 주기 유형 (DAY 등) | [optional] 
**StartTime** | Pointer to **NullableString** | 백업 시작 시간 - &#x60;HH:mm&#x60; 형식 - UTC 기준 | [optional] 
**ExpiryDuration** | Pointer to **NullableInt32** | 백업 데이터를 보존할 기간 (단위: 일) | [optional] 
**Enabled** | **bool** | 백업 스케줄 활성화 여부 | 

## Methods

### NewBackupScheduleRequest

`func NewBackupScheduleRequest(enabled bool, ) *BackupScheduleRequest`

NewBackupScheduleRequest instantiates a new BackupScheduleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupScheduleRequestWithDefaults

`func NewBackupScheduleRequestWithDefaults() *BackupScheduleRequest`

NewBackupScheduleRequestWithDefaults instantiates a new BackupScheduleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BackupScheduleRequest) GetType() BackupScheduleType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BackupScheduleRequest) GetTypeOk() (*BackupScheduleType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BackupScheduleRequest) SetType(v BackupScheduleType)`

SetType sets Type field to given value.

### HasType

`func (o *BackupScheduleRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *BackupScheduleRequest) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *BackupScheduleRequest) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetStartTime

`func (o *BackupScheduleRequest) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *BackupScheduleRequest) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *BackupScheduleRequest) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *BackupScheduleRequest) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *BackupScheduleRequest) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *BackupScheduleRequest) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetExpiryDuration

`func (o *BackupScheduleRequest) GetExpiryDuration() int32`

GetExpiryDuration returns the ExpiryDuration field if non-nil, zero value otherwise.

### GetExpiryDurationOk

`func (o *BackupScheduleRequest) GetExpiryDurationOk() (*int32, bool)`

GetExpiryDurationOk returns a tuple with the ExpiryDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDuration

`func (o *BackupScheduleRequest) SetExpiryDuration(v int32)`

SetExpiryDuration sets ExpiryDuration field to given value.

### HasExpiryDuration

`func (o *BackupScheduleRequest) HasExpiryDuration() bool`

HasExpiryDuration returns a boolean if a field has been set.

### SetExpiryDurationNil

`func (o *BackupScheduleRequest) SetExpiryDurationNil(b bool)`

 SetExpiryDurationNil sets the value for ExpiryDuration to be an explicit nil

### UnsetExpiryDuration
`func (o *BackupScheduleRequest) UnsetExpiryDuration()`

UnsetExpiryDuration ensures that no value is present for ExpiryDuration, not even an explicit nil
### GetEnabled

`func (o *BackupScheduleRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BackupScheduleRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BackupScheduleRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


