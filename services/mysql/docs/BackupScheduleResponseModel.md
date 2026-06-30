# BackupScheduleResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 백업 스케줄 ID | 
**StartTime** | Pointer to **NullableString** |  | [optional] 
**ExpiryDuration** | Pointer to **NullableInt32** |  | [optional] 
**Enabled** | **bool** | 백업 활성화 여부 | 
**Type** | Pointer to [**NullableBackupScheduleType**](BackupScheduleType.md) |  | [optional] 

## Methods

### NewBackupScheduleResponseModel

`func NewBackupScheduleResponseModel(id string, enabled bool, ) *BackupScheduleResponseModel`

NewBackupScheduleResponseModel instantiates a new BackupScheduleResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupScheduleResponseModelWithDefaults

`func NewBackupScheduleResponseModelWithDefaults() *BackupScheduleResponseModel`

NewBackupScheduleResponseModelWithDefaults instantiates a new BackupScheduleResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BackupScheduleResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BackupScheduleResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BackupScheduleResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetStartTime

`func (o *BackupScheduleResponseModel) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *BackupScheduleResponseModel) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *BackupScheduleResponseModel) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *BackupScheduleResponseModel) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *BackupScheduleResponseModel) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *BackupScheduleResponseModel) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetExpiryDuration

`func (o *BackupScheduleResponseModel) GetExpiryDuration() int32`

GetExpiryDuration returns the ExpiryDuration field if non-nil, zero value otherwise.

### GetExpiryDurationOk

`func (o *BackupScheduleResponseModel) GetExpiryDurationOk() (*int32, bool)`

GetExpiryDurationOk returns a tuple with the ExpiryDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDuration

`func (o *BackupScheduleResponseModel) SetExpiryDuration(v int32)`

SetExpiryDuration sets ExpiryDuration field to given value.

### HasExpiryDuration

`func (o *BackupScheduleResponseModel) HasExpiryDuration() bool`

HasExpiryDuration returns a boolean if a field has been set.

### SetExpiryDurationNil

`func (o *BackupScheduleResponseModel) SetExpiryDurationNil(b bool)`

 SetExpiryDurationNil sets the value for ExpiryDuration to be an explicit nil

### UnsetExpiryDuration
`func (o *BackupScheduleResponseModel) UnsetExpiryDuration()`

UnsetExpiryDuration ensures that no value is present for ExpiryDuration, not even an explicit nil
### GetEnabled

`func (o *BackupScheduleResponseModel) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BackupScheduleResponseModel) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BackupScheduleResponseModel) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetType

`func (o *BackupScheduleResponseModel) GetType() BackupScheduleType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BackupScheduleResponseModel) GetTypeOk() (*BackupScheduleType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BackupScheduleResponseModel) SetType(v BackupScheduleType)`

SetType sets Type field to given value.

### HasType

`func (o *BackupScheduleResponseModel) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *BackupScheduleResponseModel) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *BackupScheduleResponseModel) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


