# GetMysqlInstanceGroupInstanceGroup

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
**Instances** | Pointer to [**NullableGetMysqlInstanceGroupTopology**](GetMysqlInstanceGroupTopology.md) | MySQL 인스턴스 목록 | [optional] 
**Endpoint** | Pointer to **[]string** | MySQL 접속 엔드포인트 목록 | [optional] 
**ExtraInfo** | Pointer to [**NullableExtraInfo**](ExtraInfo.md) | 추가 정보 | [optional] 
**Status** | [**InstanceGroupStatus**](InstanceGroupStatus.md) | 리소스의 현재 상태 | 
**BackupSchedule** | Pointer to [**NullableBackupSchedule**](BackupSchedule.md) | 백업 스케줄 | [optional] 
**ParameterGroup** | [**ParameterGroup**](ParameterGroup.md) | 파라미터 그룹 | 

## Methods

### NewGetMysqlInstanceGroupInstanceGroup

`func NewGetMysqlInstanceGroupInstanceGroup(id string, createdAt string, updatedAt string, license string, name string, projectId string, isMultiAz bool, specContent SpecContent, status InstanceGroupStatus, parameterGroup ParameterGroup, ) *GetMysqlInstanceGroupInstanceGroup`

NewGetMysqlInstanceGroupInstanceGroup instantiates a new GetMysqlInstanceGroupInstanceGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMysqlInstanceGroupInstanceGroupWithDefaults

`func NewGetMysqlInstanceGroupInstanceGroupWithDefaults() *GetMysqlInstanceGroupInstanceGroup`

NewGetMysqlInstanceGroupInstanceGroupWithDefaults instantiates a new GetMysqlInstanceGroupInstanceGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetMysqlInstanceGroupInstanceGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetMysqlInstanceGroupInstanceGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetMysqlInstanceGroupInstanceGroup) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *GetMysqlInstanceGroupInstanceGroup) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetMysqlInstanceGroupInstanceGroup) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetMysqlInstanceGroupInstanceGroup) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *GetMysqlInstanceGroupInstanceGroup) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetMysqlInstanceGroupInstanceGroup) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetMysqlInstanceGroupInstanceGroup) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetLicense

`func (o *GetMysqlInstanceGroupInstanceGroup) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *GetMysqlInstanceGroupInstanceGroup) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *GetMysqlInstanceGroupInstanceGroup) SetLicense(v string)`

SetLicense sets License field to given value.


### GetName

`func (o *GetMysqlInstanceGroupInstanceGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetMysqlInstanceGroupInstanceGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetMysqlInstanceGroupInstanceGroup) SetName(v string)`

SetName sets Name field to given value.


### GetProjectId

`func (o *GetMysqlInstanceGroupInstanceGroup) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GetMysqlInstanceGroupInstanceGroup) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *GetMysqlInstanceGroupInstanceGroup) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetDescription

`func (o *GetMysqlInstanceGroupInstanceGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetMysqlInstanceGroupInstanceGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetMysqlInstanceGroupInstanceGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetMysqlInstanceGroupInstanceGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetMysqlInstanceGroupInstanceGroup) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetMysqlInstanceGroupInstanceGroup) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCreator

`func (o *GetMysqlInstanceGroupInstanceGroup) GetCreator() string`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *GetMysqlInstanceGroupInstanceGroup) GetCreatorOk() (*string, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *GetMysqlInstanceGroupInstanceGroup) SetCreator(v string)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *GetMysqlInstanceGroupInstanceGroup) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### SetCreatorNil

`func (o *GetMysqlInstanceGroupInstanceGroup) SetCreatorNil(b bool)`

 SetCreatorNil sets the value for Creator to be an explicit nil

### UnsetCreator
`func (o *GetMysqlInstanceGroupInstanceGroup) UnsetCreator()`

UnsetCreator ensures that no value is present for Creator, not even an explicit nil
### GetSourceBackupId

`func (o *GetMysqlInstanceGroupInstanceGroup) GetSourceBackupId() string`

GetSourceBackupId returns the SourceBackupId field if non-nil, zero value otherwise.

### GetSourceBackupIdOk

`func (o *GetMysqlInstanceGroupInstanceGroup) GetSourceBackupIdOk() (*string, bool)`

GetSourceBackupIdOk returns a tuple with the SourceBackupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceBackupId

`func (o *GetMysqlInstanceGroupInstanceGroup) SetSourceBackupId(v string)`

SetSourceBackupId sets SourceBackupId field to given value.

### HasSourceBackupId

`func (o *GetMysqlInstanceGroupInstanceGroup) HasSourceBackupId() bool`

HasSourceBackupId returns a boolean if a field has been set.

### SetSourceBackupIdNil

`func (o *GetMysqlInstanceGroupInstanceGroup) SetSourceBackupIdNil(b bool)`

 SetSourceBackupIdNil sets the value for SourceBackupId to be an explicit nil

### UnsetSourceBackupId
`func (o *GetMysqlInstanceGroupInstanceGroup) UnsetSourceBackupId()`

UnsetSourceBackupId ensures that no value is present for SourceBackupId, not even an explicit nil
### GetIsMultiAz

`func (o *GetMysqlInstanceGroupInstanceGroup) GetIsMultiAz() bool`

GetIsMultiAz returns the IsMultiAz field if non-nil, zero value otherwise.

### GetIsMultiAzOk

`func (o *GetMysqlInstanceGroupInstanceGroup) GetIsMultiAzOk() (*bool, bool)`

GetIsMultiAzOk returns a tuple with the IsMultiAz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMultiAz

`func (o *GetMysqlInstanceGroupInstanceGroup) SetIsMultiAz(v bool)`

SetIsMultiAz sets IsMultiAz field to given value.


### GetNetworkInfo

`func (o *GetMysqlInstanceGroupInstanceGroup) GetNetworkInfo() NetworkInfo`

GetNetworkInfo returns the NetworkInfo field if non-nil, zero value otherwise.

### GetNetworkInfoOk

`func (o *GetMysqlInstanceGroupInstanceGroup) GetNetworkInfoOk() (*NetworkInfo, bool)`

GetNetworkInfoOk returns a tuple with the NetworkInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInfo

`func (o *GetMysqlInstanceGroupInstanceGroup) SetNetworkInfo(v NetworkInfo)`

SetNetworkInfo sets NetworkInfo field to given value.

### HasNetworkInfo

`func (o *GetMysqlInstanceGroupInstanceGroup) HasNetworkInfo() bool`

HasNetworkInfo returns a boolean if a field has been set.

### SetNetworkInfoNil

`func (o *GetMysqlInstanceGroupInstanceGroup) SetNetworkInfoNil(b bool)`

 SetNetworkInfoNil sets the value for NetworkInfo to be an explicit nil

### UnsetNetworkInfo
`func (o *GetMysqlInstanceGroupInstanceGroup) UnsetNetworkInfo()`

UnsetNetworkInfo ensures that no value is present for NetworkInfo, not even an explicit nil
### GetSpecContent

`func (o *GetMysqlInstanceGroupInstanceGroup) GetSpecContent() SpecContent`

GetSpecContent returns the SpecContent field if non-nil, zero value otherwise.

### GetSpecContentOk

`func (o *GetMysqlInstanceGroupInstanceGroup) GetSpecContentOk() (*SpecContent, bool)`

GetSpecContentOk returns a tuple with the SpecContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecContent

`func (o *GetMysqlInstanceGroupInstanceGroup) SetSpecContent(v SpecContent)`

SetSpecContent sets SpecContent field to given value.


### GetInstances

`func (o *GetMysqlInstanceGroupInstanceGroup) GetInstances() GetMysqlInstanceGroupTopology`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *GetMysqlInstanceGroupInstanceGroup) GetInstancesOk() (*GetMysqlInstanceGroupTopology, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *GetMysqlInstanceGroupInstanceGroup) SetInstances(v GetMysqlInstanceGroupTopology)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *GetMysqlInstanceGroupInstanceGroup) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### SetInstancesNil

`func (o *GetMysqlInstanceGroupInstanceGroup) SetInstancesNil(b bool)`

 SetInstancesNil sets the value for Instances to be an explicit nil

### UnsetInstances
`func (o *GetMysqlInstanceGroupInstanceGroup) UnsetInstances()`

UnsetInstances ensures that no value is present for Instances, not even an explicit nil
### GetEndpoint

`func (o *GetMysqlInstanceGroupInstanceGroup) GetEndpoint() []string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *GetMysqlInstanceGroupInstanceGroup) GetEndpointOk() (*[]string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *GetMysqlInstanceGroupInstanceGroup) SetEndpoint(v []string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *GetMysqlInstanceGroupInstanceGroup) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### SetEndpointNil

`func (o *GetMysqlInstanceGroupInstanceGroup) SetEndpointNil(b bool)`

 SetEndpointNil sets the value for Endpoint to be an explicit nil

### UnsetEndpoint
`func (o *GetMysqlInstanceGroupInstanceGroup) UnsetEndpoint()`

UnsetEndpoint ensures that no value is present for Endpoint, not even an explicit nil
### GetExtraInfo

`func (o *GetMysqlInstanceGroupInstanceGroup) GetExtraInfo() ExtraInfo`

GetExtraInfo returns the ExtraInfo field if non-nil, zero value otherwise.

### GetExtraInfoOk

`func (o *GetMysqlInstanceGroupInstanceGroup) GetExtraInfoOk() (*ExtraInfo, bool)`

GetExtraInfoOk returns a tuple with the ExtraInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraInfo

`func (o *GetMysqlInstanceGroupInstanceGroup) SetExtraInfo(v ExtraInfo)`

SetExtraInfo sets ExtraInfo field to given value.

### HasExtraInfo

`func (o *GetMysqlInstanceGroupInstanceGroup) HasExtraInfo() bool`

HasExtraInfo returns a boolean if a field has been set.

### SetExtraInfoNil

`func (o *GetMysqlInstanceGroupInstanceGroup) SetExtraInfoNil(b bool)`

 SetExtraInfoNil sets the value for ExtraInfo to be an explicit nil

### UnsetExtraInfo
`func (o *GetMysqlInstanceGroupInstanceGroup) UnsetExtraInfo()`

UnsetExtraInfo ensures that no value is present for ExtraInfo, not even an explicit nil
### GetStatus

`func (o *GetMysqlInstanceGroupInstanceGroup) GetStatus() InstanceGroupStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetMysqlInstanceGroupInstanceGroup) GetStatusOk() (*InstanceGroupStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetMysqlInstanceGroupInstanceGroup) SetStatus(v InstanceGroupStatus)`

SetStatus sets Status field to given value.


### GetBackupSchedule

`func (o *GetMysqlInstanceGroupInstanceGroup) GetBackupSchedule() BackupSchedule`

GetBackupSchedule returns the BackupSchedule field if non-nil, zero value otherwise.

### GetBackupScheduleOk

`func (o *GetMysqlInstanceGroupInstanceGroup) GetBackupScheduleOk() (*BackupSchedule, bool)`

GetBackupScheduleOk returns a tuple with the BackupSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupSchedule

`func (o *GetMysqlInstanceGroupInstanceGroup) SetBackupSchedule(v BackupSchedule)`

SetBackupSchedule sets BackupSchedule field to given value.

### HasBackupSchedule

`func (o *GetMysqlInstanceGroupInstanceGroup) HasBackupSchedule() bool`

HasBackupSchedule returns a boolean if a field has been set.

### SetBackupScheduleNil

`func (o *GetMysqlInstanceGroupInstanceGroup) SetBackupScheduleNil(b bool)`

 SetBackupScheduleNil sets the value for BackupSchedule to be an explicit nil

### UnsetBackupSchedule
`func (o *GetMysqlInstanceGroupInstanceGroup) UnsetBackupSchedule()`

UnsetBackupSchedule ensures that no value is present for BackupSchedule, not even an explicit nil
### GetParameterGroup

`func (o *GetMysqlInstanceGroupInstanceGroup) GetParameterGroup() ParameterGroup`

GetParameterGroup returns the ParameterGroup field if non-nil, zero value otherwise.

### GetParameterGroupOk

`func (o *GetMysqlInstanceGroupInstanceGroup) GetParameterGroupOk() (*ParameterGroup, bool)`

GetParameterGroupOk returns a tuple with the ParameterGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterGroup

`func (o *GetMysqlInstanceGroupInstanceGroup) SetParameterGroup(v ParameterGroup)`

SetParameterGroup sets ParameterGroup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


