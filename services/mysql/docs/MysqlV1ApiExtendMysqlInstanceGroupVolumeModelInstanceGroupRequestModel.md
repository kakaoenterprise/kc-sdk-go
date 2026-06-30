# MysqlV1ApiExtendMysqlInstanceGroupVolumeModelInstanceGroupRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogDiskSize** | **int32** | 확장할 로그 디스크 크기 (GB, 100 ~ 16384) | 
**DataDiskSize** | **int32** | 확장할 데이터 디스크 크기 (GB, 100 ~ 16384) | 

## Methods

### NewMysqlV1ApiExtendMysqlInstanceGroupVolumeModelInstanceGroupRequestModel

`func NewMysqlV1ApiExtendMysqlInstanceGroupVolumeModelInstanceGroupRequestModel(logDiskSize int32, dataDiskSize int32, ) *MysqlV1ApiExtendMysqlInstanceGroupVolumeModelInstanceGroupRequestModel`

NewMysqlV1ApiExtendMysqlInstanceGroupVolumeModelInstanceGroupRequestModel instantiates a new MysqlV1ApiExtendMysqlInstanceGroupVolumeModelInstanceGroupRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMysqlV1ApiExtendMysqlInstanceGroupVolumeModelInstanceGroupRequestModelWithDefaults

`func NewMysqlV1ApiExtendMysqlInstanceGroupVolumeModelInstanceGroupRequestModelWithDefaults() *MysqlV1ApiExtendMysqlInstanceGroupVolumeModelInstanceGroupRequestModel`

NewMysqlV1ApiExtendMysqlInstanceGroupVolumeModelInstanceGroupRequestModelWithDefaults instantiates a new MysqlV1ApiExtendMysqlInstanceGroupVolumeModelInstanceGroupRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogDiskSize

`func (o *MysqlV1ApiExtendMysqlInstanceGroupVolumeModelInstanceGroupRequestModel) GetLogDiskSize() int32`

GetLogDiskSize returns the LogDiskSize field if non-nil, zero value otherwise.

### GetLogDiskSizeOk

`func (o *MysqlV1ApiExtendMysqlInstanceGroupVolumeModelInstanceGroupRequestModel) GetLogDiskSizeOk() (*int32, bool)`

GetLogDiskSizeOk returns a tuple with the LogDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogDiskSize

`func (o *MysqlV1ApiExtendMysqlInstanceGroupVolumeModelInstanceGroupRequestModel) SetLogDiskSize(v int32)`

SetLogDiskSize sets LogDiskSize field to given value.


### GetDataDiskSize

`func (o *MysqlV1ApiExtendMysqlInstanceGroupVolumeModelInstanceGroupRequestModel) GetDataDiskSize() int32`

GetDataDiskSize returns the DataDiskSize field if non-nil, zero value otherwise.

### GetDataDiskSizeOk

`func (o *MysqlV1ApiExtendMysqlInstanceGroupVolumeModelInstanceGroupRequestModel) GetDataDiskSizeOk() (*int32, bool)`

GetDataDiskSizeOk returns a tuple with the DataDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDiskSize

`func (o *MysqlV1ApiExtendMysqlInstanceGroupVolumeModelInstanceGroupRequestModel) SetDataDiskSize(v int32)`

SetDataDiskSize sets DataDiskSize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


