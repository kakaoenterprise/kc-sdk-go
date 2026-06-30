# MysqlV1ApiListMysqlInstancesModelSpecContentResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailabilityZone** | **string** | MySQL 인스턴스가 속한 가용 영역 | 
**FlavorId** | **string** | 인스턴스 유형(Flavor) ID | 
**DataDiskSize** | **int32** | 데이터 디스크 크기 (GB) | 
**LogDiskSize** | **int32** | 로그 디스크 크기 (GB) | 
**EngineVersion** | **string** | MySQL 인스턴스에 적용된 MySQL 엔진 버전 | 
**NetworkPorts** | [**[]NetworkPortResponseModel**](NetworkPortResponseModel.md) | 네트워크 포트 정보 목록 | 

## Methods

### NewMysqlV1ApiListMysqlInstancesModelSpecContentResponseModel

`func NewMysqlV1ApiListMysqlInstancesModelSpecContentResponseModel(availabilityZone string, flavorId string, dataDiskSize int32, logDiskSize int32, engineVersion string, networkPorts []NetworkPortResponseModel, ) *MysqlV1ApiListMysqlInstancesModelSpecContentResponseModel`

NewMysqlV1ApiListMysqlInstancesModelSpecContentResponseModel instantiates a new MysqlV1ApiListMysqlInstancesModelSpecContentResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMysqlV1ApiListMysqlInstancesModelSpecContentResponseModelWithDefaults

`func NewMysqlV1ApiListMysqlInstancesModelSpecContentResponseModelWithDefaults() *MysqlV1ApiListMysqlInstancesModelSpecContentResponseModel`

NewMysqlV1ApiListMysqlInstancesModelSpecContentResponseModelWithDefaults instantiates a new MysqlV1ApiListMysqlInstancesModelSpecContentResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailabilityZone

`func (o *MysqlV1ApiListMysqlInstancesModelSpecContentResponseModel) GetAvailabilityZone() string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *MysqlV1ApiListMysqlInstancesModelSpecContentResponseModel) GetAvailabilityZoneOk() (*string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *MysqlV1ApiListMysqlInstancesModelSpecContentResponseModel) SetAvailabilityZone(v string)`

SetAvailabilityZone sets AvailabilityZone field to given value.


### GetFlavorId

`func (o *MysqlV1ApiListMysqlInstancesModelSpecContentResponseModel) GetFlavorId() string`

GetFlavorId returns the FlavorId field if non-nil, zero value otherwise.

### GetFlavorIdOk

`func (o *MysqlV1ApiListMysqlInstancesModelSpecContentResponseModel) GetFlavorIdOk() (*string, bool)`

GetFlavorIdOk returns a tuple with the FlavorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavorId

`func (o *MysqlV1ApiListMysqlInstancesModelSpecContentResponseModel) SetFlavorId(v string)`

SetFlavorId sets FlavorId field to given value.


### GetDataDiskSize

`func (o *MysqlV1ApiListMysqlInstancesModelSpecContentResponseModel) GetDataDiskSize() int32`

GetDataDiskSize returns the DataDiskSize field if non-nil, zero value otherwise.

### GetDataDiskSizeOk

`func (o *MysqlV1ApiListMysqlInstancesModelSpecContentResponseModel) GetDataDiskSizeOk() (*int32, bool)`

GetDataDiskSizeOk returns a tuple with the DataDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDiskSize

`func (o *MysqlV1ApiListMysqlInstancesModelSpecContentResponseModel) SetDataDiskSize(v int32)`

SetDataDiskSize sets DataDiskSize field to given value.


### GetLogDiskSize

`func (o *MysqlV1ApiListMysqlInstancesModelSpecContentResponseModel) GetLogDiskSize() int32`

GetLogDiskSize returns the LogDiskSize field if non-nil, zero value otherwise.

### GetLogDiskSizeOk

`func (o *MysqlV1ApiListMysqlInstancesModelSpecContentResponseModel) GetLogDiskSizeOk() (*int32, bool)`

GetLogDiskSizeOk returns a tuple with the LogDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogDiskSize

`func (o *MysqlV1ApiListMysqlInstancesModelSpecContentResponseModel) SetLogDiskSize(v int32)`

SetLogDiskSize sets LogDiskSize field to given value.


### GetEngineVersion

`func (o *MysqlV1ApiListMysqlInstancesModelSpecContentResponseModel) GetEngineVersion() string`

GetEngineVersion returns the EngineVersion field if non-nil, zero value otherwise.

### GetEngineVersionOk

`func (o *MysqlV1ApiListMysqlInstancesModelSpecContentResponseModel) GetEngineVersionOk() (*string, bool)`

GetEngineVersionOk returns a tuple with the EngineVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineVersion

`func (o *MysqlV1ApiListMysqlInstancesModelSpecContentResponseModel) SetEngineVersion(v string)`

SetEngineVersion sets EngineVersion field to given value.


### GetNetworkPorts

`func (o *MysqlV1ApiListMysqlInstancesModelSpecContentResponseModel) GetNetworkPorts() []NetworkPortResponseModel`

GetNetworkPorts returns the NetworkPorts field if non-nil, zero value otherwise.

### GetNetworkPortsOk

`func (o *MysqlV1ApiListMysqlInstancesModelSpecContentResponseModel) GetNetworkPortsOk() (*[]NetworkPortResponseModel, bool)`

GetNetworkPortsOk returns a tuple with the NetworkPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPorts

`func (o *MysqlV1ApiListMysqlInstancesModelSpecContentResponseModel) SetNetworkPorts(v []NetworkPortResponseModel)`

SetNetworkPorts sets NetworkPorts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


