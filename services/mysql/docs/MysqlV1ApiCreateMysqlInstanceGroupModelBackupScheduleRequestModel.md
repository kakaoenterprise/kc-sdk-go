# MysqlV1ApiCreateMysqlInstanceGroupModelBackupScheduleRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**NullableBackupScheduleType**](BackupScheduleType.md) |  | [optional] 
**StartTime** | Pointer to **NullableString** |  | [optional] 
**ExpiryDuration** | Pointer to **NullableInt32** |  | [optional] 
**Enabled** | **bool** | 백업 스케줄 활성화 여부 | 

## Methods

### NewMysqlV1ApiCreateMysqlInstanceGroupModelBackupScheduleRequestModel

`func NewMysqlV1ApiCreateMysqlInstanceGroupModelBackupScheduleRequestModel(enabled bool, ) *MysqlV1ApiCreateMysqlInstanceGroupModelBackupScheduleRequestModel`

NewMysqlV1ApiCreateMysqlInstanceGroupModelBackupScheduleRequestModel instantiates a new MysqlV1ApiCreateMysqlInstanceGroupModelBackupScheduleRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMysqlV1ApiCreateMysqlInstanceGroupModelBackupScheduleRequestModelWithDefaults

`func NewMysqlV1ApiCreateMysqlInstanceGroupModelBackupScheduleRequestModelWithDefaults() *MysqlV1ApiCreateMysqlInstanceGroupModelBackupScheduleRequestModel`

NewMysqlV1ApiCreateMysqlInstanceGroupModelBackupScheduleRequestModelWithDefaults instantiates a new MysqlV1ApiCreateMysqlInstanceGroupModelBackupScheduleRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelBackupScheduleRequestModel) GetType() BackupScheduleType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelBackupScheduleRequestModel) GetTypeOk() (*BackupScheduleType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelBackupScheduleRequestModel) SetType(v BackupScheduleType)`

SetType sets Type field to given value.

### HasType

`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelBackupScheduleRequestModel) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelBackupScheduleRequestModel) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelBackupScheduleRequestModel) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetStartTime

`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelBackupScheduleRequestModel) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelBackupScheduleRequestModel) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelBackupScheduleRequestModel) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelBackupScheduleRequestModel) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelBackupScheduleRequestModel) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelBackupScheduleRequestModel) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetExpiryDuration

`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelBackupScheduleRequestModel) GetExpiryDuration() int32`

GetExpiryDuration returns the ExpiryDuration field if non-nil, zero value otherwise.

### GetExpiryDurationOk

`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelBackupScheduleRequestModel) GetExpiryDurationOk() (*int32, bool)`

GetExpiryDurationOk returns a tuple with the ExpiryDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDuration

`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelBackupScheduleRequestModel) SetExpiryDuration(v int32)`

SetExpiryDuration sets ExpiryDuration field to given value.

### HasExpiryDuration

`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelBackupScheduleRequestModel) HasExpiryDuration() bool`

HasExpiryDuration returns a boolean if a field has been set.

### SetExpiryDurationNil

`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelBackupScheduleRequestModel) SetExpiryDurationNil(b bool)`

 SetExpiryDurationNil sets the value for ExpiryDuration to be an explicit nil

### UnsetExpiryDuration
`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelBackupScheduleRequestModel) UnsetExpiryDuration()`

UnsetExpiryDuration ensures that no value is present for ExpiryDuration, not even an explicit nil
### GetEnabled

`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelBackupScheduleRequestModel) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelBackupScheduleRequestModel) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelBackupScheduleRequestModel) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


