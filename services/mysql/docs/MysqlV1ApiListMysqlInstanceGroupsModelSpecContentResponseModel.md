# MysqlV1ApiListMysqlInstanceGroupsModelSpecContentResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatabaseUserName** | **string** | 기본 데이터베이스 사용자 이름 | 
**PrimaryPort** | **int32** | Primary 포트 번호 | 
**StandbyPort** | **int32** | Standby 포트 번호 | 
**EngineVersion** | **string** | MySQL 엔진 버전 | 
**FlavorId** | **string** | 인스턴스 Flavor ID | 
**Vcpu** | **int32** | vCPU 개수 | 
**Memory** | **int32** | 메모리 크기 (MB) | 
**LogDiskSize** | **int32** | 로그 디스크 크기 (GB) | 
**DataDiskSize** | **int32** | 데이터 디스크 크기 (GB) | 
**InstanceGroupType** | [**ClusterType**](ClusterType.md) | MySQL 인스턴스 그룹 유형 | 
**NodeSize** | **int32** | MySQL 인스턴스 그룹의 노드 크기 | 

## Methods

### NewMysqlV1ApiListMysqlInstanceGroupsModelSpecContentResponseModel

`func NewMysqlV1ApiListMysqlInstanceGroupsModelSpecContentResponseModel(databaseUserName string, primaryPort int32, standbyPort int32, engineVersion string, flavorId string, vcpu int32, memory int32, logDiskSize int32, dataDiskSize int32, instanceGroupType ClusterType, nodeSize int32, ) *MysqlV1ApiListMysqlInstanceGroupsModelSpecContentResponseModel`

NewMysqlV1ApiListMysqlInstanceGroupsModelSpecContentResponseModel instantiates a new MysqlV1ApiListMysqlInstanceGroupsModelSpecContentResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMysqlV1ApiListMysqlInstanceGroupsModelSpecContentResponseModelWithDefaults

`func NewMysqlV1ApiListMysqlInstanceGroupsModelSpecContentResponseModelWithDefaults() *MysqlV1ApiListMysqlInstanceGroupsModelSpecContentResponseModel`

NewMysqlV1ApiListMysqlInstanceGroupsModelSpecContentResponseModelWithDefaults instantiates a new MysqlV1ApiListMysqlInstanceGroupsModelSpecContentResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabaseUserName

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelSpecContentResponseModel) GetDatabaseUserName() string`

GetDatabaseUserName returns the DatabaseUserName field if non-nil, zero value otherwise.

### GetDatabaseUserNameOk

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelSpecContentResponseModel) GetDatabaseUserNameOk() (*string, bool)`

GetDatabaseUserNameOk returns a tuple with the DatabaseUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseUserName

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelSpecContentResponseModel) SetDatabaseUserName(v string)`

SetDatabaseUserName sets DatabaseUserName field to given value.


### GetPrimaryPort

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelSpecContentResponseModel) GetPrimaryPort() int32`

GetPrimaryPort returns the PrimaryPort field if non-nil, zero value otherwise.

### GetPrimaryPortOk

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelSpecContentResponseModel) GetPrimaryPortOk() (*int32, bool)`

GetPrimaryPortOk returns a tuple with the PrimaryPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryPort

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelSpecContentResponseModel) SetPrimaryPort(v int32)`

SetPrimaryPort sets PrimaryPort field to given value.


### GetStandbyPort

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelSpecContentResponseModel) GetStandbyPort() int32`

GetStandbyPort returns the StandbyPort field if non-nil, zero value otherwise.

### GetStandbyPortOk

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelSpecContentResponseModel) GetStandbyPortOk() (*int32, bool)`

GetStandbyPortOk returns a tuple with the StandbyPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandbyPort

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelSpecContentResponseModel) SetStandbyPort(v int32)`

SetStandbyPort sets StandbyPort field to given value.


### GetEngineVersion

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelSpecContentResponseModel) GetEngineVersion() string`

GetEngineVersion returns the EngineVersion field if non-nil, zero value otherwise.

### GetEngineVersionOk

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelSpecContentResponseModel) GetEngineVersionOk() (*string, bool)`

GetEngineVersionOk returns a tuple with the EngineVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineVersion

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelSpecContentResponseModel) SetEngineVersion(v string)`

SetEngineVersion sets EngineVersion field to given value.


### GetFlavorId

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelSpecContentResponseModel) GetFlavorId() string`

GetFlavorId returns the FlavorId field if non-nil, zero value otherwise.

### GetFlavorIdOk

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelSpecContentResponseModel) GetFlavorIdOk() (*string, bool)`

GetFlavorIdOk returns a tuple with the FlavorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavorId

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelSpecContentResponseModel) SetFlavorId(v string)`

SetFlavorId sets FlavorId field to given value.


### GetVcpu

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelSpecContentResponseModel) GetVcpu() int32`

GetVcpu returns the Vcpu field if non-nil, zero value otherwise.

### GetVcpuOk

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelSpecContentResponseModel) GetVcpuOk() (*int32, bool)`

GetVcpuOk returns a tuple with the Vcpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcpu

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelSpecContentResponseModel) SetVcpu(v int32)`

SetVcpu sets Vcpu field to given value.


### GetMemory

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelSpecContentResponseModel) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelSpecContentResponseModel) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelSpecContentResponseModel) SetMemory(v int32)`

SetMemory sets Memory field to given value.


### GetLogDiskSize

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelSpecContentResponseModel) GetLogDiskSize() int32`

GetLogDiskSize returns the LogDiskSize field if non-nil, zero value otherwise.

### GetLogDiskSizeOk

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelSpecContentResponseModel) GetLogDiskSizeOk() (*int32, bool)`

GetLogDiskSizeOk returns a tuple with the LogDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogDiskSize

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelSpecContentResponseModel) SetLogDiskSize(v int32)`

SetLogDiskSize sets LogDiskSize field to given value.


### GetDataDiskSize

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelSpecContentResponseModel) GetDataDiskSize() int32`

GetDataDiskSize returns the DataDiskSize field if non-nil, zero value otherwise.

### GetDataDiskSizeOk

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelSpecContentResponseModel) GetDataDiskSizeOk() (*int32, bool)`

GetDataDiskSizeOk returns a tuple with the DataDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDiskSize

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelSpecContentResponseModel) SetDataDiskSize(v int32)`

SetDataDiskSize sets DataDiskSize field to given value.


### GetInstanceGroupType

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelSpecContentResponseModel) GetInstanceGroupType() ClusterType`

GetInstanceGroupType returns the InstanceGroupType field if non-nil, zero value otherwise.

### GetInstanceGroupTypeOk

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelSpecContentResponseModel) GetInstanceGroupTypeOk() (*ClusterType, bool)`

GetInstanceGroupTypeOk returns a tuple with the InstanceGroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceGroupType

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelSpecContentResponseModel) SetInstanceGroupType(v ClusterType)`

SetInstanceGroupType sets InstanceGroupType field to given value.


### GetNodeSize

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelSpecContentResponseModel) GetNodeSize() int32`

GetNodeSize returns the NodeSize field if non-nil, zero value otherwise.

### GetNodeSizeOk

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelSpecContentResponseModel) GetNodeSizeOk() (*int32, bool)`

GetNodeSizeOk returns a tuple with the NodeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeSize

`func (o *MysqlV1ApiListMysqlInstanceGroupsModelSpecContentResponseModel) SetNodeSize(v int32)`

SetNodeSize sets NodeSize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


