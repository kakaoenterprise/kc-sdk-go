# CreateMysqlInstanceGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | MySQL 인스턴스 그룹 이름 | 
**Description** | Pointer to **NullableString** | MySQL 인스턴스 그룹에 대한 설명 | [optional] 
**NetworkInfo** | [**NetworkInfoRequest**](NetworkInfoRequest.md) | 인스턴스 그룹의 네트워크 설정 정보 | 
**SpecContent** | [**SpecContentRequest**](SpecContentRequest.md) | MySQL 인스턴스 사양 및 엔진 설정 | 
**Source** | Pointer to [**NullableRestoreSourceRequest**](RestoreSourceRequest.md) | 인스턴스 그룹 생성 시 사용할 소스 정보 | [optional] 
**BackupSchedule** | [**BackupScheduleRequest**](BackupScheduleRequest.md) | 백업 스케줄 설정 정보 | 
**ParameterGroup** | [**ParameterGroupRequest**](ParameterGroupRequest.md) | MySQL 인스턴스 그룹에 적용할 파라미터 그룹 정보 | 
**ExtraInfo** | Pointer to [**NullableExtraInfoRequest**](ExtraInfoRequest.md) | 추가 설정 정보 | [optional] 

## Methods

### NewCreateMysqlInstanceGroup

`func NewCreateMysqlInstanceGroup(name string, networkInfo NetworkInfoRequest, specContent SpecContentRequest, backupSchedule BackupScheduleRequest, parameterGroup ParameterGroupRequest, ) *CreateMysqlInstanceGroup`

NewCreateMysqlInstanceGroup instantiates a new CreateMysqlInstanceGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMysqlInstanceGroupWithDefaults

`func NewCreateMysqlInstanceGroupWithDefaults() *CreateMysqlInstanceGroup`

NewCreateMysqlInstanceGroupWithDefaults instantiates a new CreateMysqlInstanceGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateMysqlInstanceGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateMysqlInstanceGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateMysqlInstanceGroup) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateMysqlInstanceGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateMysqlInstanceGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateMysqlInstanceGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateMysqlInstanceGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateMysqlInstanceGroup) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateMysqlInstanceGroup) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetNetworkInfo

`func (o *CreateMysqlInstanceGroup) GetNetworkInfo() NetworkInfoRequest`

GetNetworkInfo returns the NetworkInfo field if non-nil, zero value otherwise.

### GetNetworkInfoOk

`func (o *CreateMysqlInstanceGroup) GetNetworkInfoOk() (*NetworkInfoRequest, bool)`

GetNetworkInfoOk returns a tuple with the NetworkInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInfo

`func (o *CreateMysqlInstanceGroup) SetNetworkInfo(v NetworkInfoRequest)`

SetNetworkInfo sets NetworkInfo field to given value.


### GetSpecContent

`func (o *CreateMysqlInstanceGroup) GetSpecContent() SpecContentRequest`

GetSpecContent returns the SpecContent field if non-nil, zero value otherwise.

### GetSpecContentOk

`func (o *CreateMysqlInstanceGroup) GetSpecContentOk() (*SpecContentRequest, bool)`

GetSpecContentOk returns a tuple with the SpecContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecContent

`func (o *CreateMysqlInstanceGroup) SetSpecContent(v SpecContentRequest)`

SetSpecContent sets SpecContent field to given value.


### GetSource

`func (o *CreateMysqlInstanceGroup) GetSource() RestoreSourceRequest`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CreateMysqlInstanceGroup) GetSourceOk() (*RestoreSourceRequest, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CreateMysqlInstanceGroup) SetSource(v RestoreSourceRequest)`

SetSource sets Source field to given value.

### HasSource

`func (o *CreateMysqlInstanceGroup) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *CreateMysqlInstanceGroup) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *CreateMysqlInstanceGroup) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetBackupSchedule

`func (o *CreateMysqlInstanceGroup) GetBackupSchedule() BackupScheduleRequest`

GetBackupSchedule returns the BackupSchedule field if non-nil, zero value otherwise.

### GetBackupScheduleOk

`func (o *CreateMysqlInstanceGroup) GetBackupScheduleOk() (*BackupScheduleRequest, bool)`

GetBackupScheduleOk returns a tuple with the BackupSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupSchedule

`func (o *CreateMysqlInstanceGroup) SetBackupSchedule(v BackupScheduleRequest)`

SetBackupSchedule sets BackupSchedule field to given value.


### GetParameterGroup

`func (o *CreateMysqlInstanceGroup) GetParameterGroup() ParameterGroupRequest`

GetParameterGroup returns the ParameterGroup field if non-nil, zero value otherwise.

### GetParameterGroupOk

`func (o *CreateMysqlInstanceGroup) GetParameterGroupOk() (*ParameterGroupRequest, bool)`

GetParameterGroupOk returns a tuple with the ParameterGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterGroup

`func (o *CreateMysqlInstanceGroup) SetParameterGroup(v ParameterGroupRequest)`

SetParameterGroup sets ParameterGroup field to given value.


### GetExtraInfo

`func (o *CreateMysqlInstanceGroup) GetExtraInfo() ExtraInfoRequest`

GetExtraInfo returns the ExtraInfo field if non-nil, zero value otherwise.

### GetExtraInfoOk

`func (o *CreateMysqlInstanceGroup) GetExtraInfoOk() (*ExtraInfoRequest, bool)`

GetExtraInfoOk returns a tuple with the ExtraInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraInfo

`func (o *CreateMysqlInstanceGroup) SetExtraInfo(v ExtraInfoRequest)`

SetExtraInfo sets ExtraInfo field to given value.

### HasExtraInfo

`func (o *CreateMysqlInstanceGroup) HasExtraInfo() bool`

HasExtraInfo returns a boolean if a field has been set.

### SetExtraInfoNil

`func (o *CreateMysqlInstanceGroup) SetExtraInfoNil(b bool)`

 SetExtraInfoNil sets the value for ExtraInfo to be an explicit nil

### UnsetExtraInfo
`func (o *CreateMysqlInstanceGroup) UnsetExtraInfo()`

UnsetExtraInfo ensures that no value is present for ExtraInfo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


