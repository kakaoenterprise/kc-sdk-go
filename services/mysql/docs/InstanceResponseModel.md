# InstanceResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | MySQL 인스턴스 ID | 
**ProjectId** | **string** | 프로젝트 ID | 
**InstanceGroupId** | **string** | 대상 MySQL 인스턴스 그룹 ID | 
**InstanceGroupName** | **string** | MySQL 인스턴스 그룹 이름 | 
**Name** | **string** | MySQL 인스턴스 이름 | 
**Status** | [**InstanceStatus**](InstanceStatus.md) | 리소스의 현재 상태 | 
**AvailabilityStatus** | Pointer to **NullableString** |  | [optional] 
**StatusContent** | [**StatusContentResponseModel**](StatusContentResponseModel.md) | MySQL 인스턴스의 상태 정보 | 
**Role** | [**InstanceRole**](InstanceRole.md) | MySQL 인스턴스 역할 | 
**DataDiskUsage** | **int32** | 데이터 디스크 사용량 | 
**LogDiskUsage** | **int32** | 로그 디스크 사용량 | 
**SpecContent** | [**MysqlV1ApiListMysqlInstancesModelSpecContentResponseModel**](MysqlV1ApiListMysqlInstancesModelSpecContentResponseModel.md) | MySQL 인스턴스의 사양 정보 | 
**CreatedAt** | **string** | 리소스가 생성된 시간 &lt;br/&gt; - ISO 8601 형식 &lt;br/&gt; - UTC 기준 | 
**UpdatedAt** | **string** | 리소스가 마지막으로 수정된 시간 &lt;br/&gt; - ISO 8601 형식 &lt;br/&gt; - UTC 기준 | 
**StartTime** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewInstanceResponseModel

`func NewInstanceResponseModel(id string, projectId string, instanceGroupId string, instanceGroupName string, name string, status InstanceStatus, statusContent StatusContentResponseModel, role InstanceRole, dataDiskUsage int32, logDiskUsage int32, specContent MysqlV1ApiListMysqlInstancesModelSpecContentResponseModel, createdAt string, updatedAt string, ) *InstanceResponseModel`

NewInstanceResponseModel instantiates a new InstanceResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceResponseModelWithDefaults

`func NewInstanceResponseModelWithDefaults() *InstanceResponseModel`

NewInstanceResponseModelWithDefaults instantiates a new InstanceResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstanceResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetProjectId

`func (o *InstanceResponseModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *InstanceResponseModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *InstanceResponseModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetInstanceGroupId

`func (o *InstanceResponseModel) GetInstanceGroupId() string`

GetInstanceGroupId returns the InstanceGroupId field if non-nil, zero value otherwise.

### GetInstanceGroupIdOk

`func (o *InstanceResponseModel) GetInstanceGroupIdOk() (*string, bool)`

GetInstanceGroupIdOk returns a tuple with the InstanceGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceGroupId

`func (o *InstanceResponseModel) SetInstanceGroupId(v string)`

SetInstanceGroupId sets InstanceGroupId field to given value.


### GetInstanceGroupName

`func (o *InstanceResponseModel) GetInstanceGroupName() string`

GetInstanceGroupName returns the InstanceGroupName field if non-nil, zero value otherwise.

### GetInstanceGroupNameOk

`func (o *InstanceResponseModel) GetInstanceGroupNameOk() (*string, bool)`

GetInstanceGroupNameOk returns a tuple with the InstanceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceGroupName

`func (o *InstanceResponseModel) SetInstanceGroupName(v string)`

SetInstanceGroupName sets InstanceGroupName field to given value.


### GetName

`func (o *InstanceResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceResponseModel) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *InstanceResponseModel) GetStatus() InstanceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InstanceResponseModel) GetStatusOk() (*InstanceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InstanceResponseModel) SetStatus(v InstanceStatus)`

SetStatus sets Status field to given value.


### GetAvailabilityStatus

`func (o *InstanceResponseModel) GetAvailabilityStatus() string`

GetAvailabilityStatus returns the AvailabilityStatus field if non-nil, zero value otherwise.

### GetAvailabilityStatusOk

`func (o *InstanceResponseModel) GetAvailabilityStatusOk() (*string, bool)`

GetAvailabilityStatusOk returns a tuple with the AvailabilityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityStatus

`func (o *InstanceResponseModel) SetAvailabilityStatus(v string)`

SetAvailabilityStatus sets AvailabilityStatus field to given value.

### HasAvailabilityStatus

`func (o *InstanceResponseModel) HasAvailabilityStatus() bool`

HasAvailabilityStatus returns a boolean if a field has been set.

### SetAvailabilityStatusNil

`func (o *InstanceResponseModel) SetAvailabilityStatusNil(b bool)`

 SetAvailabilityStatusNil sets the value for AvailabilityStatus to be an explicit nil

### UnsetAvailabilityStatus
`func (o *InstanceResponseModel) UnsetAvailabilityStatus()`

UnsetAvailabilityStatus ensures that no value is present for AvailabilityStatus, not even an explicit nil
### GetStatusContent

`func (o *InstanceResponseModel) GetStatusContent() StatusContentResponseModel`

GetStatusContent returns the StatusContent field if non-nil, zero value otherwise.

### GetStatusContentOk

`func (o *InstanceResponseModel) GetStatusContentOk() (*StatusContentResponseModel, bool)`

GetStatusContentOk returns a tuple with the StatusContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusContent

`func (o *InstanceResponseModel) SetStatusContent(v StatusContentResponseModel)`

SetStatusContent sets StatusContent field to given value.


### GetRole

`func (o *InstanceResponseModel) GetRole() InstanceRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *InstanceResponseModel) GetRoleOk() (*InstanceRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *InstanceResponseModel) SetRole(v InstanceRole)`

SetRole sets Role field to given value.


### GetDataDiskUsage

`func (o *InstanceResponseModel) GetDataDiskUsage() int32`

GetDataDiskUsage returns the DataDiskUsage field if non-nil, zero value otherwise.

### GetDataDiskUsageOk

`func (o *InstanceResponseModel) GetDataDiskUsageOk() (*int32, bool)`

GetDataDiskUsageOk returns a tuple with the DataDiskUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDiskUsage

`func (o *InstanceResponseModel) SetDataDiskUsage(v int32)`

SetDataDiskUsage sets DataDiskUsage field to given value.


### GetLogDiskUsage

`func (o *InstanceResponseModel) GetLogDiskUsage() int32`

GetLogDiskUsage returns the LogDiskUsage field if non-nil, zero value otherwise.

### GetLogDiskUsageOk

`func (o *InstanceResponseModel) GetLogDiskUsageOk() (*int32, bool)`

GetLogDiskUsageOk returns a tuple with the LogDiskUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogDiskUsage

`func (o *InstanceResponseModel) SetLogDiskUsage(v int32)`

SetLogDiskUsage sets LogDiskUsage field to given value.


### GetSpecContent

`func (o *InstanceResponseModel) GetSpecContent() MysqlV1ApiListMysqlInstancesModelSpecContentResponseModel`

GetSpecContent returns the SpecContent field if non-nil, zero value otherwise.

### GetSpecContentOk

`func (o *InstanceResponseModel) GetSpecContentOk() (*MysqlV1ApiListMysqlInstancesModelSpecContentResponseModel, bool)`

GetSpecContentOk returns a tuple with the SpecContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecContent

`func (o *InstanceResponseModel) SetSpecContent(v MysqlV1ApiListMysqlInstancesModelSpecContentResponseModel)`

SetSpecContent sets SpecContent field to given value.


### GetCreatedAt

`func (o *InstanceResponseModel) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InstanceResponseModel) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InstanceResponseModel) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *InstanceResponseModel) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InstanceResponseModel) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InstanceResponseModel) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetStartTime

`func (o *InstanceResponseModel) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *InstanceResponseModel) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *InstanceResponseModel) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *InstanceResponseModel) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *InstanceResponseModel) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *InstanceResponseModel) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


