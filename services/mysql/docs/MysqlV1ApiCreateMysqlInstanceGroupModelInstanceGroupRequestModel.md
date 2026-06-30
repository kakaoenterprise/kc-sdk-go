# MysqlV1ApiCreateMysqlInstanceGroupModelInstanceGroupRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | MySQL 인스턴스 그룹 이름 | 
**Description** | Pointer to **NullableString** |  | [optional] 
**NetworkInfo** | [**NetworkInfoRequestModel**](NetworkInfoRequestModel.md) | 인스턴스 그룹의 네트워크 설정 정보 | 
**SpecContent** | [**SpecContentRequestModel**](SpecContentRequestModel.md) | MySQL 인스턴스 사양 및 엔진 설정 | 
**Source** | Pointer to [**NullableRestoreSourceRequestModel**](RestoreSourceRequestModel.md) |  | [optional] 
**BackupSchedule** | [**MysqlV1ApiCreateMysqlInstanceGroupModelBackupScheduleRequestModel**](MysqlV1ApiCreateMysqlInstanceGroupModelBackupScheduleRequestModel.md) | 백업 스케줄 설정 정보 | 
**ParameterGroup** | [**MysqlV1ApiCreateMysqlInstanceGroupModelParameterGroupRequestModel**](MysqlV1ApiCreateMysqlInstanceGroupModelParameterGroupRequestModel.md) | MySQL 인스턴스 그룹에 적용할 파라미터 그룹 정보 | 
**ExtraInfo** | Pointer to [**NullableExtraInfoRequestModel**](ExtraInfoRequestModel.md) |  | [optional] 

## Methods

### NewMysqlV1ApiCreateMysqlInstanceGroupModelInstanceGroupRequestModel

`func NewMysqlV1ApiCreateMysqlInstanceGroupModelInstanceGroupRequestModel(name string, networkInfo NetworkInfoRequestModel, specContent SpecContentRequestModel, backupSchedule MysqlV1ApiCreateMysqlInstanceGroupModelBackupScheduleRequestModel, parameterGroup MysqlV1ApiCreateMysqlInstanceGroupModelParameterGroupRequestModel, ) *MysqlV1ApiCreateMysqlInstanceGroupModelInstanceGroupRequestModel`

NewMysqlV1ApiCreateMysqlInstanceGroupModelInstanceGroupRequestModel instantiates a new MysqlV1ApiCreateMysqlInstanceGroupModelInstanceGroupRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMysqlV1ApiCreateMysqlInstanceGroupModelInstanceGroupRequestModelWithDefaults

`func NewMysqlV1ApiCreateMysqlInstanceGroupModelInstanceGroupRequestModelWithDefaults() *MysqlV1ApiCreateMysqlInstanceGroupModelInstanceGroupRequestModel`

NewMysqlV1ApiCreateMysqlInstanceGroupModelInstanceGroupRequestModelWithDefaults instantiates a new MysqlV1ApiCreateMysqlInstanceGroupModelInstanceGroupRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelInstanceGroupRequestModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelInstanceGroupRequestModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelInstanceGroupRequestModel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelInstanceGroupRequestModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelInstanceGroupRequestModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelInstanceGroupRequestModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelInstanceGroupRequestModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelInstanceGroupRequestModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelInstanceGroupRequestModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetNetworkInfo

`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelInstanceGroupRequestModel) GetNetworkInfo() NetworkInfoRequestModel`

GetNetworkInfo returns the NetworkInfo field if non-nil, zero value otherwise.

### GetNetworkInfoOk

`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelInstanceGroupRequestModel) GetNetworkInfoOk() (*NetworkInfoRequestModel, bool)`

GetNetworkInfoOk returns a tuple with the NetworkInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInfo

`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelInstanceGroupRequestModel) SetNetworkInfo(v NetworkInfoRequestModel)`

SetNetworkInfo sets NetworkInfo field to given value.


### GetSpecContent

`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelInstanceGroupRequestModel) GetSpecContent() SpecContentRequestModel`

GetSpecContent returns the SpecContent field if non-nil, zero value otherwise.

### GetSpecContentOk

`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelInstanceGroupRequestModel) GetSpecContentOk() (*SpecContentRequestModel, bool)`

GetSpecContentOk returns a tuple with the SpecContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecContent

`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelInstanceGroupRequestModel) SetSpecContent(v SpecContentRequestModel)`

SetSpecContent sets SpecContent field to given value.


### GetSource

`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelInstanceGroupRequestModel) GetSource() RestoreSourceRequestModel`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelInstanceGroupRequestModel) GetSourceOk() (*RestoreSourceRequestModel, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelInstanceGroupRequestModel) SetSource(v RestoreSourceRequestModel)`

SetSource sets Source field to given value.

### HasSource

`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelInstanceGroupRequestModel) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSourceNil

`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelInstanceGroupRequestModel) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelInstanceGroupRequestModel) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetBackupSchedule

`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelInstanceGroupRequestModel) GetBackupSchedule() MysqlV1ApiCreateMysqlInstanceGroupModelBackupScheduleRequestModel`

GetBackupSchedule returns the BackupSchedule field if non-nil, zero value otherwise.

### GetBackupScheduleOk

`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelInstanceGroupRequestModel) GetBackupScheduleOk() (*MysqlV1ApiCreateMysqlInstanceGroupModelBackupScheduleRequestModel, bool)`

GetBackupScheduleOk returns a tuple with the BackupSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupSchedule

`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelInstanceGroupRequestModel) SetBackupSchedule(v MysqlV1ApiCreateMysqlInstanceGroupModelBackupScheduleRequestModel)`

SetBackupSchedule sets BackupSchedule field to given value.


### GetParameterGroup

`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelInstanceGroupRequestModel) GetParameterGroup() MysqlV1ApiCreateMysqlInstanceGroupModelParameterGroupRequestModel`

GetParameterGroup returns the ParameterGroup field if non-nil, zero value otherwise.

### GetParameterGroupOk

`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelInstanceGroupRequestModel) GetParameterGroupOk() (*MysqlV1ApiCreateMysqlInstanceGroupModelParameterGroupRequestModel, bool)`

GetParameterGroupOk returns a tuple with the ParameterGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterGroup

`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelInstanceGroupRequestModel) SetParameterGroup(v MysqlV1ApiCreateMysqlInstanceGroupModelParameterGroupRequestModel)`

SetParameterGroup sets ParameterGroup field to given value.


### GetExtraInfo

`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelInstanceGroupRequestModel) GetExtraInfo() ExtraInfoRequestModel`

GetExtraInfo returns the ExtraInfo field if non-nil, zero value otherwise.

### GetExtraInfoOk

`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelInstanceGroupRequestModel) GetExtraInfoOk() (*ExtraInfoRequestModel, bool)`

GetExtraInfoOk returns a tuple with the ExtraInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraInfo

`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelInstanceGroupRequestModel) SetExtraInfo(v ExtraInfoRequestModel)`

SetExtraInfo sets ExtraInfo field to given value.

### HasExtraInfo

`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelInstanceGroupRequestModel) HasExtraInfo() bool`

HasExtraInfo returns a boolean if a field has been set.

### SetExtraInfoNil

`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelInstanceGroupRequestModel) SetExtraInfoNil(b bool)`

 SetExtraInfoNil sets the value for ExtraInfo to be an explicit nil

### UnsetExtraInfo
`func (o *MysqlV1ApiCreateMysqlInstanceGroupModelInstanceGroupRequestModel) UnsetExtraInfo()`

UnsetExtraInfo ensures that no value is present for ExtraInfo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


