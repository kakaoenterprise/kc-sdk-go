# Instance

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
**StatusContent** | [**StatusContent**](StatusContent.md) | MySQL 인스턴스의 상태 정보 | 
**Role** | [**InstanceRole**](InstanceRole.md) | MySQL 인스턴스 역할 | 
**DataDiskUsage** | **int32** | 데이터 디스크 사용량 | 
**LogDiskUsage** | **int32** | 로그 디스크 사용량 | 
**SpecContent** | [**ListMysqlInstancesSpecContent**](ListMysqlInstancesSpecContent.md) | MySQL 인스턴스의 사양 정보 | 
**CreatedAt** | **string** | 리소스가 생성된 시간 &lt;br/&gt; - ISO 8601 형식 &lt;br/&gt; - UTC 기준 | 
**UpdatedAt** | **string** | 리소스가 마지막으로 수정된 시간 &lt;br/&gt; - ISO 8601 형식 &lt;br/&gt; - UTC 기준 | 
**StartTime** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewInstance

`func NewInstance(id string, projectId string, instanceGroupId string, instanceGroupName string, name string, status InstanceStatus, statusContent StatusContent, role InstanceRole, dataDiskUsage int32, logDiskUsage int32, specContent ListMysqlInstancesSpecContent, createdAt string, updatedAt string, ) *Instance`

NewInstance instantiates a new Instance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceWithDefaults

`func NewInstanceWithDefaults() *Instance`

NewInstanceWithDefaults instantiates a new Instance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Instance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Instance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Instance) SetId(v string)`

SetId sets Id field to given value.


### GetProjectId

`func (o *Instance) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *Instance) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *Instance) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetInstanceGroupId

`func (o *Instance) GetInstanceGroupId() string`

GetInstanceGroupId returns the InstanceGroupId field if non-nil, zero value otherwise.

### GetInstanceGroupIdOk

`func (o *Instance) GetInstanceGroupIdOk() (*string, bool)`

GetInstanceGroupIdOk returns a tuple with the InstanceGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceGroupId

`func (o *Instance) SetInstanceGroupId(v string)`

SetInstanceGroupId sets InstanceGroupId field to given value.


### GetInstanceGroupName

`func (o *Instance) GetInstanceGroupName() string`

GetInstanceGroupName returns the InstanceGroupName field if non-nil, zero value otherwise.

### GetInstanceGroupNameOk

`func (o *Instance) GetInstanceGroupNameOk() (*string, bool)`

GetInstanceGroupNameOk returns a tuple with the InstanceGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceGroupName

`func (o *Instance) SetInstanceGroupName(v string)`

SetInstanceGroupName sets InstanceGroupName field to given value.


### GetName

`func (o *Instance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Instance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Instance) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *Instance) GetStatus() InstanceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Instance) GetStatusOk() (*InstanceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Instance) SetStatus(v InstanceStatus)`

SetStatus sets Status field to given value.


### GetAvailabilityStatus

`func (o *Instance) GetAvailabilityStatus() string`

GetAvailabilityStatus returns the AvailabilityStatus field if non-nil, zero value otherwise.

### GetAvailabilityStatusOk

`func (o *Instance) GetAvailabilityStatusOk() (*string, bool)`

GetAvailabilityStatusOk returns a tuple with the AvailabilityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityStatus

`func (o *Instance) SetAvailabilityStatus(v string)`

SetAvailabilityStatus sets AvailabilityStatus field to given value.

### HasAvailabilityStatus

`func (o *Instance) HasAvailabilityStatus() bool`

HasAvailabilityStatus returns a boolean if a field has been set.

### SetAvailabilityStatusNil

`func (o *Instance) SetAvailabilityStatusNil(b bool)`

 SetAvailabilityStatusNil sets the value for AvailabilityStatus to be an explicit nil

### UnsetAvailabilityStatus
`func (o *Instance) UnsetAvailabilityStatus()`

UnsetAvailabilityStatus ensures that no value is present for AvailabilityStatus, not even an explicit nil
### GetStatusContent

`func (o *Instance) GetStatusContent() StatusContent`

GetStatusContent returns the StatusContent field if non-nil, zero value otherwise.

### GetStatusContentOk

`func (o *Instance) GetStatusContentOk() (*StatusContent, bool)`

GetStatusContentOk returns a tuple with the StatusContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusContent

`func (o *Instance) SetStatusContent(v StatusContent)`

SetStatusContent sets StatusContent field to given value.


### GetRole

`func (o *Instance) GetRole() InstanceRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *Instance) GetRoleOk() (*InstanceRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *Instance) SetRole(v InstanceRole)`

SetRole sets Role field to given value.


### GetDataDiskUsage

`func (o *Instance) GetDataDiskUsage() int32`

GetDataDiskUsage returns the DataDiskUsage field if non-nil, zero value otherwise.

### GetDataDiskUsageOk

`func (o *Instance) GetDataDiskUsageOk() (*int32, bool)`

GetDataDiskUsageOk returns a tuple with the DataDiskUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDiskUsage

`func (o *Instance) SetDataDiskUsage(v int32)`

SetDataDiskUsage sets DataDiskUsage field to given value.


### GetLogDiskUsage

`func (o *Instance) GetLogDiskUsage() int32`

GetLogDiskUsage returns the LogDiskUsage field if non-nil, zero value otherwise.

### GetLogDiskUsageOk

`func (o *Instance) GetLogDiskUsageOk() (*int32, bool)`

GetLogDiskUsageOk returns a tuple with the LogDiskUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogDiskUsage

`func (o *Instance) SetLogDiskUsage(v int32)`

SetLogDiskUsage sets LogDiskUsage field to given value.


### GetSpecContent

`func (o *Instance) GetSpecContent() ListMysqlInstancesSpecContent`

GetSpecContent returns the SpecContent field if non-nil, zero value otherwise.

### GetSpecContentOk

`func (o *Instance) GetSpecContentOk() (*ListMysqlInstancesSpecContent, bool)`

GetSpecContentOk returns a tuple with the SpecContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecContent

`func (o *Instance) SetSpecContent(v ListMysqlInstancesSpecContent)`

SetSpecContent sets SpecContent field to given value.


### GetCreatedAt

`func (o *Instance) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Instance) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Instance) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Instance) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Instance) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Instance) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetStartTime

`func (o *Instance) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *Instance) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *Instance) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *Instance) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *Instance) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *Instance) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


