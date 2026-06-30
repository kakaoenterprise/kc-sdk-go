# MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | MySQL 인스턴스 그룹 ID | 
**CreatedAt** | **string** | 리소스가 생성된 시간 &lt;br/&gt; - ISO 8601 형식 &lt;br/&gt; - UTC 기준 | 
**UpdatedAt** | **string** | 리소스가 마지막으로 수정된 시간 &lt;br/&gt; - ISO 8601 형식 &lt;br/&gt; - UTC 기준 | 
**License** | **string** | MySQL 라이선스 유형 | 
**Name** | **string** | MySQL 인스턴스 그룹 이름 | 
**ProjectId** | **string** | 프로젝트 ID | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Creator** | Pointer to **NullableString** |  | [optional] 
**SourceBackupId** | Pointer to **NullableString** |  | [optional] 
**IsMultiAz** | **bool** | 다중 AZ 구성 여부 | 
**NetworkInfo** | Pointer to [**NullableNetworkInfoResponseModel**](NetworkInfoResponseModel.md) |  | [optional] 
**SpecContent** | [**MysqlV1ApiListMysqlInstanceGroupsModelSpecContentResponseModel**](MysqlV1ApiListMysqlInstanceGroupsModelSpecContentResponseModel.md) | MySQL 인스턴스 그룹 구성 정보 | 
**Instances** | Pointer to [**NullableMysqlV1ApiListMysqlInstanceGroupsModelTopologyResponseModel**](MysqlV1ApiListMysqlInstanceGroupsModelTopologyResponseModel.md) |  | [optional] 
**Endpoint** | Pointer to **[]string** |  | [optional] 
**ExtraInfo** | Pointer to [**NullableMysqlV1ApiListMysqlInstanceGroupsModelExtraInfoResponseModel**](MysqlV1ApiListMysqlInstanceGroupsModelExtraInfoResponseModel.md) |  | [optional] 
**Status** | [**InstanceGroupStatus**](InstanceGroupStatus.md) | 리소스의 현재 상태 | 
**BackupSchedule** | Pointer to [**NullableBackupScheduleResponseModel**](BackupScheduleResponseModel.md) |  | [optional] 
**ParameterGroup** | [**ParameterGroupResponseModel**](ParameterGroupResponseModel.md) | 파라미터 그룹 | 

## Methods

### NewMysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel

`func NewMysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel(id string, createdAt string, updatedAt string, license string, name string, projectId string, isMultiAz bool, specContent MysqlV1ApiListMysqlInstanceGroupsModelSpecContentResponseModel, status InstanceGroupStatus, parameterGroup ParameterGroupResponseModel, ) *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel`

NewMysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel instantiates a new MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModelWithDefaults

`func NewMysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModelWithDefaults() *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel`

NewMysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModelWithDefaults instantiates a new MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetLicense

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) SetLicense(v string)`

SetLicense sets License field to given value.


### GetName

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) SetName(v string)`

SetName sets Name field to given value.


### GetProjectId

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetDescription

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCreator

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) GetCreator() string`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) GetCreatorOk() (*string, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) SetCreator(v string)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### SetCreatorNil

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) SetCreatorNil(b bool)`

 SetCreatorNil sets the value for Creator to be an explicit nil

### UnsetCreator
`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) UnsetCreator()`

UnsetCreator ensures that no value is present for Creator, not even an explicit nil
### GetSourceBackupId

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) GetSourceBackupId() string`

GetSourceBackupId returns the SourceBackupId field if non-nil, zero value otherwise.

### GetSourceBackupIdOk

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) GetSourceBackupIdOk() (*string, bool)`

GetSourceBackupIdOk returns a tuple with the SourceBackupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceBackupId

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) SetSourceBackupId(v string)`

SetSourceBackupId sets SourceBackupId field to given value.

### HasSourceBackupId

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) HasSourceBackupId() bool`

HasSourceBackupId returns a boolean if a field has been set.

### SetSourceBackupIdNil

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) SetSourceBackupIdNil(b bool)`

 SetSourceBackupIdNil sets the value for SourceBackupId to be an explicit nil

### UnsetSourceBackupId
`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) UnsetSourceBackupId()`

UnsetSourceBackupId ensures that no value is present for SourceBackupId, not even an explicit nil
### GetIsMultiAz

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) GetIsMultiAz() bool`

GetIsMultiAz returns the IsMultiAz field if non-nil, zero value otherwise.

### GetIsMultiAzOk

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) GetIsMultiAzOk() (*bool, bool)`

GetIsMultiAzOk returns a tuple with the IsMultiAz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMultiAz

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) SetIsMultiAz(v bool)`

SetIsMultiAz sets IsMultiAz field to given value.


### GetNetworkInfo

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) GetNetworkInfo() NetworkInfoResponseModel`

GetNetworkInfo returns the NetworkInfo field if non-nil, zero value otherwise.

### GetNetworkInfoOk

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) GetNetworkInfoOk() (*NetworkInfoResponseModel, bool)`

GetNetworkInfoOk returns a tuple with the NetworkInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInfo

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) SetNetworkInfo(v NetworkInfoResponseModel)`

SetNetworkInfo sets NetworkInfo field to given value.

### HasNetworkInfo

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) HasNetworkInfo() bool`

HasNetworkInfo returns a boolean if a field has been set.

### SetNetworkInfoNil

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) SetNetworkInfoNil(b bool)`

 SetNetworkInfoNil sets the value for NetworkInfo to be an explicit nil

### UnsetNetworkInfo
`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) UnsetNetworkInfo()`

UnsetNetworkInfo ensures that no value is present for NetworkInfo, not even an explicit nil
### GetSpecContent

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) GetSpecContent() MysqlV1ApiListMysqlInstanceGroupsModelSpecContentResponseModel`

GetSpecContent returns the SpecContent field if non-nil, zero value otherwise.

### GetSpecContentOk

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) GetSpecContentOk() (*MysqlV1ApiListMysqlInstanceGroupsModelSpecContentResponseModel, bool)`

GetSpecContentOk returns a tuple with the SpecContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecContent

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) SetSpecContent(v MysqlV1ApiListMysqlInstanceGroupsModelSpecContentResponseModel)`

SetSpecContent sets SpecContent field to given value.


### GetInstances

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) GetInstances() MysqlV1ApiListMysqlInstanceGroupsModelTopologyResponseModel`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) GetInstancesOk() (*MysqlV1ApiListMysqlInstanceGroupsModelTopologyResponseModel, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) SetInstances(v MysqlV1ApiListMysqlInstanceGroupsModelTopologyResponseModel)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### SetInstancesNil

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) SetInstancesNil(b bool)`

 SetInstancesNil sets the value for Instances to be an explicit nil

### UnsetInstances
`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) UnsetInstances()`

UnsetInstances ensures that no value is present for Instances, not even an explicit nil
### GetEndpoint

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) GetEndpoint() []string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) GetEndpointOk() (*[]string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) SetEndpoint(v []string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### SetEndpointNil

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) SetEndpointNil(b bool)`

 SetEndpointNil sets the value for Endpoint to be an explicit nil

### UnsetEndpoint
`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) UnsetEndpoint()`

UnsetEndpoint ensures that no value is present for Endpoint, not even an explicit nil
### GetExtraInfo

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) GetExtraInfo() MysqlV1ApiListMysqlInstanceGroupsModelExtraInfoResponseModel`

GetExtraInfo returns the ExtraInfo field if non-nil, zero value otherwise.

### GetExtraInfoOk

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) GetExtraInfoOk() (*MysqlV1ApiListMysqlInstanceGroupsModelExtraInfoResponseModel, bool)`

GetExtraInfoOk returns a tuple with the ExtraInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraInfo

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) SetExtraInfo(v MysqlV1ApiListMysqlInstanceGroupsModelExtraInfoResponseModel)`

SetExtraInfo sets ExtraInfo field to given value.

### HasExtraInfo

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) HasExtraInfo() bool`

HasExtraInfo returns a boolean if a field has been set.

### SetExtraInfoNil

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) SetExtraInfoNil(b bool)`

 SetExtraInfoNil sets the value for ExtraInfo to be an explicit nil

### UnsetExtraInfo
`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) UnsetExtraInfo()`

UnsetExtraInfo ensures that no value is present for ExtraInfo, not even an explicit nil
### GetStatus

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) GetStatus() InstanceGroupStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) GetStatusOk() (*InstanceGroupStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) SetStatus(v InstanceGroupStatus)`

SetStatus sets Status field to given value.


### GetBackupSchedule

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) GetBackupSchedule() BackupScheduleResponseModel`

GetBackupSchedule returns the BackupSchedule field if non-nil, zero value otherwise.

### GetBackupScheduleOk

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) GetBackupScheduleOk() (*BackupScheduleResponseModel, bool)`

GetBackupScheduleOk returns a tuple with the BackupSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupSchedule

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) SetBackupSchedule(v BackupScheduleResponseModel)`

SetBackupSchedule sets BackupSchedule field to given value.

### HasBackupSchedule

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) HasBackupSchedule() bool`

HasBackupSchedule returns a boolean if a field has been set.

### SetBackupScheduleNil

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) SetBackupScheduleNil(b bool)`

 SetBackupScheduleNil sets the value for BackupSchedule to be an explicit nil

### UnsetBackupSchedule
`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) UnsetBackupSchedule()`

UnsetBackupSchedule ensures that no value is present for BackupSchedule, not even an explicit nil
### GetParameterGroup

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) GetParameterGroup() ParameterGroupResponseModel`

GetParameterGroup returns the ParameterGroup field if non-nil, zero value otherwise.

### GetParameterGroupOk

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) GetParameterGroupOk() (*ParameterGroupResponseModel, bool)`

GetParameterGroupOk returns a tuple with the ParameterGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterGroup

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelInstanceGroupResponseModel) SetParameterGroup(v ParameterGroupResponseModel)`

SetParameterGroup sets ParameterGroup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


