# ListMysqlInstanceGroupsInstanceGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | MySQL 인스턴스 그룹 ID | 
**CreatedAt** | **string** | 리소스가 생성된 시간 - ISO 8601 형식 - UTC 기준 | 
**UpdatedAt** | **string** | 리소스가 마지막으로 수정된 시간 - ISO 8601 형식 - UTC 기준 | 
**License** | **string** | MySQL 라이선스 유형 | 
**Name** | **string** | MySQL 인스턴스 그룹 이름 | 
**ProjectId** | **string** | 프로젝트 ID | 
**Description** | Pointer to **NullableString** | 인스턴스 그룹 설명 | [optional] 
**Creator** | Pointer to **NullableString** | 인스턴스 그룹 생성자 | [optional] 
**SourceBackupId** | Pointer to **NullableString** | 인스턴스 그룹 생성에 사용된 백업 ID | [optional] 
**IsMultiAz** | **bool** | 다중 AZ 구성 여부 | 
**NetworkInfo** | Pointer to [**NullableNetworkInfo**](NetworkInfo.md) | 네트워크 구성 정보 | [optional] 
**SpecContent** | [**SpecContent**](SpecContent.md) | MySQL 인스턴스 그룹 구성 정보 | 
**Instances** | Pointer to [**NullableListMysqlInstanceGroupsTopology**](ListMysqlInstanceGroupsTopology.md) | MySQL 인스턴스 목록 | [optional] 
**Endpoint** | Pointer to **[]string** | MySQL 접속 엔드포인트 목록 | [optional] 
**ExtraInfo** | Pointer to [**NullableExtraInfo**](ExtraInfo.md) | 추가 정보 | [optional] 
**Status** | [**InstanceGroupStatus**](InstanceGroupStatus.md) | 리소스의 현재 상태 | 
**BackupSchedule** | Pointer to [**NullableBackupSchedule**](BackupSchedule.md) | 백업 스케줄 | [optional] 
**ParameterGroup** | [**ParameterGroup**](ParameterGroup.md) | 파라미터 그룹 | 

## Methods

### NewListMysqlInstanceGroupsInstanceGroup

`func NewListMysqlInstanceGroupsInstanceGroup(id string, createdAt string, updatedAt string, license string, name string, projectId string, isMultiAz bool, specContent SpecContent, status InstanceGroupStatus, parameterGroup ParameterGroup, ) *ListMysqlInstanceGroupsInstanceGroup`

NewListMysqlInstanceGroupsInstanceGroup instantiates a new ListMysqlInstanceGroupsInstanceGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMysqlInstanceGroupsInstanceGroupWithDefaults

`func NewListMysqlInstanceGroupsInstanceGroupWithDefaults() *ListMysqlInstanceGroupsInstanceGroup`

NewListMysqlInstanceGroupsInstanceGroupWithDefaults instantiates a new ListMysqlInstanceGroupsInstanceGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListMysqlInstanceGroupsInstanceGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListMysqlInstanceGroupsInstanceGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListMysqlInstanceGroupsInstanceGroup) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *ListMysqlInstanceGroupsInstanceGroup) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListMysqlInstanceGroupsInstanceGroup) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListMysqlInstanceGroupsInstanceGroup) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ListMysqlInstanceGroupsInstanceGroup) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ListMysqlInstanceGroupsInstanceGroup) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ListMysqlInstanceGroupsInstanceGroup) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetLicense

`func (o *ListMysqlInstanceGroupsInstanceGroup) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *ListMysqlInstanceGroupsInstanceGroup) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *ListMysqlInstanceGroupsInstanceGroup) SetLicense(v string)`

SetLicense sets License field to given value.


### GetName

`func (o *ListMysqlInstanceGroupsInstanceGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListMysqlInstanceGroupsInstanceGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListMysqlInstanceGroupsInstanceGroup) SetName(v string)`

SetName sets Name field to given value.


### GetProjectId

`func (o *ListMysqlInstanceGroupsInstanceGroup) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ListMysqlInstanceGroupsInstanceGroup) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ListMysqlInstanceGroupsInstanceGroup) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetDescription

`func (o *ListMysqlInstanceGroupsInstanceGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListMysqlInstanceGroupsInstanceGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListMysqlInstanceGroupsInstanceGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListMysqlInstanceGroupsInstanceGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ListMysqlInstanceGroupsInstanceGroup) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ListMysqlInstanceGroupsInstanceGroup) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCreator

`func (o *ListMysqlInstanceGroupsInstanceGroup) GetCreator() string`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *ListMysqlInstanceGroupsInstanceGroup) GetCreatorOk() (*string, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *ListMysqlInstanceGroupsInstanceGroup) SetCreator(v string)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *ListMysqlInstanceGroupsInstanceGroup) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### SetCreatorNil

`func (o *ListMysqlInstanceGroupsInstanceGroup) SetCreatorNil(b bool)`

 SetCreatorNil sets the value for Creator to be an explicit nil

### UnsetCreator
`func (o *ListMysqlInstanceGroupsInstanceGroup) UnsetCreator()`

UnsetCreator ensures that no value is present for Creator, not even an explicit nil
### GetSourceBackupId

`func (o *ListMysqlInstanceGroupsInstanceGroup) GetSourceBackupId() string`

GetSourceBackupId returns the SourceBackupId field if non-nil, zero value otherwise.

### GetSourceBackupIdOk

`func (o *ListMysqlInstanceGroupsInstanceGroup) GetSourceBackupIdOk() (*string, bool)`

GetSourceBackupIdOk returns a tuple with the SourceBackupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceBackupId

`func (o *ListMysqlInstanceGroupsInstanceGroup) SetSourceBackupId(v string)`

SetSourceBackupId sets SourceBackupId field to given value.

### HasSourceBackupId

`func (o *ListMysqlInstanceGroupsInstanceGroup) HasSourceBackupId() bool`

HasSourceBackupId returns a boolean if a field has been set.

### SetSourceBackupIdNil

`func (o *ListMysqlInstanceGroupsInstanceGroup) SetSourceBackupIdNil(b bool)`

 SetSourceBackupIdNil sets the value for SourceBackupId to be an explicit nil

### UnsetSourceBackupId
`func (o *ListMysqlInstanceGroupsInstanceGroup) UnsetSourceBackupId()`

UnsetSourceBackupId ensures that no value is present for SourceBackupId, not even an explicit nil
### GetIsMultiAz

`func (o *ListMysqlInstanceGroupsInstanceGroup) GetIsMultiAz() bool`

GetIsMultiAz returns the IsMultiAz field if non-nil, zero value otherwise.

### GetIsMultiAzOk

`func (o *ListMysqlInstanceGroupsInstanceGroup) GetIsMultiAzOk() (*bool, bool)`

GetIsMultiAzOk returns a tuple with the IsMultiAz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMultiAz

`func (o *ListMysqlInstanceGroupsInstanceGroup) SetIsMultiAz(v bool)`

SetIsMultiAz sets IsMultiAz field to given value.


### GetNetworkInfo

`func (o *ListMysqlInstanceGroupsInstanceGroup) GetNetworkInfo() NetworkInfo`

GetNetworkInfo returns the NetworkInfo field if non-nil, zero value otherwise.

### GetNetworkInfoOk

`func (o *ListMysqlInstanceGroupsInstanceGroup) GetNetworkInfoOk() (*NetworkInfo, bool)`

GetNetworkInfoOk returns a tuple with the NetworkInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInfo

`func (o *ListMysqlInstanceGroupsInstanceGroup) SetNetworkInfo(v NetworkInfo)`

SetNetworkInfo sets NetworkInfo field to given value.

### HasNetworkInfo

`func (o *ListMysqlInstanceGroupsInstanceGroup) HasNetworkInfo() bool`

HasNetworkInfo returns a boolean if a field has been set.

### SetNetworkInfoNil

`func (o *ListMysqlInstanceGroupsInstanceGroup) SetNetworkInfoNil(b bool)`

 SetNetworkInfoNil sets the value for NetworkInfo to be an explicit nil

### UnsetNetworkInfo
`func (o *ListMysqlInstanceGroupsInstanceGroup) UnsetNetworkInfo()`

UnsetNetworkInfo ensures that no value is present for NetworkInfo, not even an explicit nil
### GetSpecContent

`func (o *ListMysqlInstanceGroupsInstanceGroup) GetSpecContent() SpecContent`

GetSpecContent returns the SpecContent field if non-nil, zero value otherwise.

### GetSpecContentOk

`func (o *ListMysqlInstanceGroupsInstanceGroup) GetSpecContentOk() (*SpecContent, bool)`

GetSpecContentOk returns a tuple with the SpecContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecContent

`func (o *ListMysqlInstanceGroupsInstanceGroup) SetSpecContent(v SpecContent)`

SetSpecContent sets SpecContent field to given value.


### GetInstances

`func (o *ListMysqlInstanceGroupsInstanceGroup) GetInstances() ListMysqlInstanceGroupsTopology`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *ListMysqlInstanceGroupsInstanceGroup) GetInstancesOk() (*ListMysqlInstanceGroupsTopology, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *ListMysqlInstanceGroupsInstanceGroup) SetInstances(v ListMysqlInstanceGroupsTopology)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *ListMysqlInstanceGroupsInstanceGroup) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### SetInstancesNil

`func (o *ListMysqlInstanceGroupsInstanceGroup) SetInstancesNil(b bool)`

 SetInstancesNil sets the value for Instances to be an explicit nil

### UnsetInstances
`func (o *ListMysqlInstanceGroupsInstanceGroup) UnsetInstances()`

UnsetInstances ensures that no value is present for Instances, not even an explicit nil
### GetEndpoint

`func (o *ListMysqlInstanceGroupsInstanceGroup) GetEndpoint() []string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *ListMysqlInstanceGroupsInstanceGroup) GetEndpointOk() (*[]string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *ListMysqlInstanceGroupsInstanceGroup) SetEndpoint(v []string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *ListMysqlInstanceGroupsInstanceGroup) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### SetEndpointNil

`func (o *ListMysqlInstanceGroupsInstanceGroup) SetEndpointNil(b bool)`

 SetEndpointNil sets the value for Endpoint to be an explicit nil

### UnsetEndpoint
`func (o *ListMysqlInstanceGroupsInstanceGroup) UnsetEndpoint()`

UnsetEndpoint ensures that no value is present for Endpoint, not even an explicit nil
### GetExtraInfo

`func (o *ListMysqlInstanceGroupsInstanceGroup) GetExtraInfo() ExtraInfo`

GetExtraInfo returns the ExtraInfo field if non-nil, zero value otherwise.

### GetExtraInfoOk

`func (o *ListMysqlInstanceGroupsInstanceGroup) GetExtraInfoOk() (*ExtraInfo, bool)`

GetExtraInfoOk returns a tuple with the ExtraInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraInfo

`func (o *ListMysqlInstanceGroupsInstanceGroup) SetExtraInfo(v ExtraInfo)`

SetExtraInfo sets ExtraInfo field to given value.

### HasExtraInfo

`func (o *ListMysqlInstanceGroupsInstanceGroup) HasExtraInfo() bool`

HasExtraInfo returns a boolean if a field has been set.

### SetExtraInfoNil

`func (o *ListMysqlInstanceGroupsInstanceGroup) SetExtraInfoNil(b bool)`

 SetExtraInfoNil sets the value for ExtraInfo to be an explicit nil

### UnsetExtraInfo
`func (o *ListMysqlInstanceGroupsInstanceGroup) UnsetExtraInfo()`

UnsetExtraInfo ensures that no value is present for ExtraInfo, not even an explicit nil
### GetStatus

`func (o *ListMysqlInstanceGroupsInstanceGroup) GetStatus() InstanceGroupStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListMysqlInstanceGroupsInstanceGroup) GetStatusOk() (*InstanceGroupStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListMysqlInstanceGroupsInstanceGroup) SetStatus(v InstanceGroupStatus)`

SetStatus sets Status field to given value.


### GetBackupSchedule

`func (o *ListMysqlInstanceGroupsInstanceGroup) GetBackupSchedule() BackupSchedule`

GetBackupSchedule returns the BackupSchedule field if non-nil, zero value otherwise.

### GetBackupScheduleOk

`func (o *ListMysqlInstanceGroupsInstanceGroup) GetBackupScheduleOk() (*BackupSchedule, bool)`

GetBackupScheduleOk returns a tuple with the BackupSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupSchedule

`func (o *ListMysqlInstanceGroupsInstanceGroup) SetBackupSchedule(v BackupSchedule)`

SetBackupSchedule sets BackupSchedule field to given value.

### HasBackupSchedule

`func (o *ListMysqlInstanceGroupsInstanceGroup) HasBackupSchedule() bool`

HasBackupSchedule returns a boolean if a field has been set.

### SetBackupScheduleNil

`func (o *ListMysqlInstanceGroupsInstanceGroup) SetBackupScheduleNil(b bool)`

 SetBackupScheduleNil sets the value for BackupSchedule to be an explicit nil

### UnsetBackupSchedule
`func (o *ListMysqlInstanceGroupsInstanceGroup) UnsetBackupSchedule()`

UnsetBackupSchedule ensures that no value is present for BackupSchedule, not even an explicit nil
### GetParameterGroup

`func (o *ListMysqlInstanceGroupsInstanceGroup) GetParameterGroup() ParameterGroup`

GetParameterGroup returns the ParameterGroup field if non-nil, zero value otherwise.

### GetParameterGroupOk

`func (o *ListMysqlInstanceGroupsInstanceGroup) GetParameterGroupOk() (*ParameterGroup, bool)`

GetParameterGroupOk returns a tuple with the ParameterGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterGroup

`func (o *ListMysqlInstanceGroupsInstanceGroup) SetParameterGroup(v ParameterGroup)`

SetParameterGroup sets ParameterGroup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


