# ListMysqlInstancesSpecContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailabilityZone** | **string** | MySQL 인스턴스가 속한 가용 영역 | 
**FlavorId** | **string** | 인스턴스 유형(Flavor) ID | 
**DataDiskSize** | **int32** | 데이터 디스크 크기 (GB) | 
**LogDiskSize** | **int32** | 로그 디스크 크기 (GB) | 
**EngineVersion** | **string** | MySQL 인스턴스에 적용된 MySQL 엔진 버전 | 
**NetworkPorts** | [**[]NetworkPort**](NetworkPort.md) | 네트워크 포트 정보 목록 | 

## Methods

### NewListMysqlInstancesSpecContent

`func NewListMysqlInstancesSpecContent(availabilityZone string, flavorId string, dataDiskSize int32, logDiskSize int32, engineVersion string, networkPorts []NetworkPort, ) *ListMysqlInstancesSpecContent`

NewListMysqlInstancesSpecContent instantiates a new ListMysqlInstancesSpecContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMysqlInstancesSpecContentWithDefaults

`func NewListMysqlInstancesSpecContentWithDefaults() *ListMysqlInstancesSpecContent`

NewListMysqlInstancesSpecContentWithDefaults instantiates a new ListMysqlInstancesSpecContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailabilityZone

`func (o *ListMysqlInstancesSpecContent) GetAvailabilityZone() string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *ListMysqlInstancesSpecContent) GetAvailabilityZoneOk() (*string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *ListMysqlInstancesSpecContent) SetAvailabilityZone(v string)`

SetAvailabilityZone sets AvailabilityZone field to given value.


### GetFlavorId

`func (o *ListMysqlInstancesSpecContent) GetFlavorId() string`

GetFlavorId returns the FlavorId field if non-nil, zero value otherwise.

### GetFlavorIdOk

`func (o *ListMysqlInstancesSpecContent) GetFlavorIdOk() (*string, bool)`

GetFlavorIdOk returns a tuple with the FlavorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavorId

`func (o *ListMysqlInstancesSpecContent) SetFlavorId(v string)`

SetFlavorId sets FlavorId field to given value.


### GetDataDiskSize

`func (o *ListMysqlInstancesSpecContent) GetDataDiskSize() int32`

GetDataDiskSize returns the DataDiskSize field if non-nil, zero value otherwise.

### GetDataDiskSizeOk

`func (o *ListMysqlInstancesSpecContent) GetDataDiskSizeOk() (*int32, bool)`

GetDataDiskSizeOk returns a tuple with the DataDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDiskSize

`func (o *ListMysqlInstancesSpecContent) SetDataDiskSize(v int32)`

SetDataDiskSize sets DataDiskSize field to given value.


### GetLogDiskSize

`func (o *ListMysqlInstancesSpecContent) GetLogDiskSize() int32`

GetLogDiskSize returns the LogDiskSize field if non-nil, zero value otherwise.

### GetLogDiskSizeOk

`func (o *ListMysqlInstancesSpecContent) GetLogDiskSizeOk() (*int32, bool)`

GetLogDiskSizeOk returns a tuple with the LogDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogDiskSize

`func (o *ListMysqlInstancesSpecContent) SetLogDiskSize(v int32)`

SetLogDiskSize sets LogDiskSize field to given value.


### GetEngineVersion

`func (o *ListMysqlInstancesSpecContent) GetEngineVersion() string`

GetEngineVersion returns the EngineVersion field if non-nil, zero value otherwise.

### GetEngineVersionOk

`func (o *ListMysqlInstancesSpecContent) GetEngineVersionOk() (*string, bool)`

GetEngineVersionOk returns a tuple with the EngineVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineVersion

`func (o *ListMysqlInstancesSpecContent) SetEngineVersion(v string)`

SetEngineVersion sets EngineVersion field to given value.


### GetNetworkPorts

`func (o *ListMysqlInstancesSpecContent) GetNetworkPorts() []NetworkPort`

GetNetworkPorts returns the NetworkPorts field if non-nil, zero value otherwise.

### GetNetworkPortsOk

`func (o *ListMysqlInstancesSpecContent) GetNetworkPortsOk() (*[]NetworkPort, bool)`

GetNetworkPortsOk returns a tuple with the NetworkPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPorts

`func (o *ListMysqlInstancesSpecContent) SetNetworkPorts(v []NetworkPort)`

SetNetworkPorts sets NetworkPorts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


