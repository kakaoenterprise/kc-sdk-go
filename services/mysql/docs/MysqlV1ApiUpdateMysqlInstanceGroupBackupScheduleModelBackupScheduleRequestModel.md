# MysqlV1ApiUpdateMysqlInstanceGroupBackupScheduleModelBackupScheduleRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**NullableBackupScheduleType**](BackupScheduleType.md) |  | [optional] 
**StartTime** | Pointer to **NullableString** |  | [optional] 
**ExpiryDuration** | Pointer to **NullableInt32** |  | [optional] 
**Enabled** | **bool** | # enabled - 타입: Boolean - 설명: (필드 설명을 여기에 작성하세요) | 

## Methods

### NewMysqlV1ApiUpdateMysqlInstanceGroupBackupScheduleModelBackupScheduleRequestModel

`func NewMysqlV1ApiUpdateMysqlInstanceGroupBackupScheduleModelBackupScheduleRequestModel(enabled bool, ) *MysqlV1ApiUpdateMysqlInstanceGroupBackupScheduleModelBackupScheduleRequestModel`

NewMysqlV1ApiUpdateMysqlInstanceGroupBackupScheduleModelBackupScheduleRequestModel instantiates a new MysqlV1ApiUpdateMysqlInstanceGroupBackupScheduleModelBackupScheduleRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMysqlV1ApiUpdateMysqlInstanceGroupBackupScheduleModelBackupScheduleRequestModelWithDefaults

`func NewMysqlV1ApiUpdateMysqlInstanceGroupBackupScheduleModelBackupScheduleRequestModelWithDefaults() *MysqlV1ApiUpdateMysqlInstanceGroupBackupScheduleModelBackupScheduleRequestModel`

NewMysqlV1ApiUpdateMysqlInstanceGroupBackupScheduleModelBackupScheduleRequestModelWithDefaults instantiates a new MysqlV1ApiUpdateMysqlInstanceGroupBackupScheduleModelBackupScheduleRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MysqlV1ApiUpdateMysqlInstanceGroupBackupScheduleModelBackupScheduleRequestModel) GetType() BackupScheduleType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MysqlV1ApiUpdateMysqlInstanceGroupBackupScheduleModelBackupScheduleRequestModel) GetTypeOk() (*BackupScheduleType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MysqlV1ApiUpdateMysqlInstanceGroupBackupScheduleModelBackupScheduleRequestModel) SetType(v BackupScheduleType)`

SetType sets Type field to given value.

### HasType

`func (o *MysqlV1ApiUpdateMysqlInstanceGroupBackupScheduleModelBackupScheduleRequestModel) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *MysqlV1ApiUpdateMysqlInstanceGroupBackupScheduleModelBackupScheduleRequestModel) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *MysqlV1ApiUpdateMysqlInstanceGroupBackupScheduleModelBackupScheduleRequestModel) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetStartTime

`func (o *MysqlV1ApiUpdateMysqlInstanceGroupBackupScheduleModelBackupScheduleRequestModel) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *MysqlV1ApiUpdateMysqlInstanceGroupBackupScheduleModelBackupScheduleRequestModel) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *MysqlV1ApiUpdateMysqlInstanceGroupBackupScheduleModelBackupScheduleRequestModel) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *MysqlV1ApiUpdateMysqlInstanceGroupBackupScheduleModelBackupScheduleRequestModel) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *MysqlV1ApiUpdateMysqlInstanceGroupBackupScheduleModelBackupScheduleRequestModel) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *MysqlV1ApiUpdateMysqlInstanceGroupBackupScheduleModelBackupScheduleRequestModel) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetExpiryDuration

`func (o *MysqlV1ApiUpdateMysqlInstanceGroupBackupScheduleModelBackupScheduleRequestModel) GetExpiryDuration() int32`

GetExpiryDuration returns the ExpiryDuration field if non-nil, zero value otherwise.

### GetExpiryDurationOk

`func (o *MysqlV1ApiUpdateMysqlInstanceGroupBackupScheduleModelBackupScheduleRequestModel) GetExpiryDurationOk() (*int32, bool)`

GetExpiryDurationOk returns a tuple with the ExpiryDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDuration

`func (o *MysqlV1ApiUpdateMysqlInstanceGroupBackupScheduleModelBackupScheduleRequestModel) SetExpiryDuration(v int32)`

SetExpiryDuration sets ExpiryDuration field to given value.

### HasExpiryDuration

`func (o *MysqlV1ApiUpdateMysqlInstanceGroupBackupScheduleModelBackupScheduleRequestModel) HasExpiryDuration() bool`

HasExpiryDuration returns a boolean if a field has been set.

### SetExpiryDurationNil

`func (o *MysqlV1ApiUpdateMysqlInstanceGroupBackupScheduleModelBackupScheduleRequestModel) SetExpiryDurationNil(b bool)`

 SetExpiryDurationNil sets the value for ExpiryDuration to be an explicit nil

### UnsetExpiryDuration
`func (o *MysqlV1ApiUpdateMysqlInstanceGroupBackupScheduleModelBackupScheduleRequestModel) UnsetExpiryDuration()`

UnsetExpiryDuration ensures that no value is present for ExpiryDuration, not even an explicit nil
### GetEnabled

`func (o *MysqlV1ApiUpdateMysqlInstanceGroupBackupScheduleModelBackupScheduleRequestModel) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MysqlV1ApiUpdateMysqlInstanceGroupBackupScheduleModelBackupScheduleRequestModel) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MysqlV1ApiUpdateMysqlInstanceGroupBackupScheduleModelBackupScheduleRequestModel) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


