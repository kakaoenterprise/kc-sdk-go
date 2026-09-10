# ExtendMysqlInstanceGroupVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogDiskSize** | **int32** | 확장할 로그 디스크 크기 (GB, 100 ~ 16384) | 
**DataDiskSize** | **int32** | 확장할 데이터 디스크 크기 (GB, 100 ~ 16384) | 

## Methods

### NewExtendMysqlInstanceGroupVolume

`func NewExtendMysqlInstanceGroupVolume(logDiskSize int32, dataDiskSize int32, ) *ExtendMysqlInstanceGroupVolume`

NewExtendMysqlInstanceGroupVolume instantiates a new ExtendMysqlInstanceGroupVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtendMysqlInstanceGroupVolumeWithDefaults

`func NewExtendMysqlInstanceGroupVolumeWithDefaults() *ExtendMysqlInstanceGroupVolume`

NewExtendMysqlInstanceGroupVolumeWithDefaults instantiates a new ExtendMysqlInstanceGroupVolume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogDiskSize

`func (o *ExtendMysqlInstanceGroupVolume) GetLogDiskSize() int32`

GetLogDiskSize returns the LogDiskSize field if non-nil, zero value otherwise.

### GetLogDiskSizeOk

`func (o *ExtendMysqlInstanceGroupVolume) GetLogDiskSizeOk() (*int32, bool)`

GetLogDiskSizeOk returns a tuple with the LogDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogDiskSize

`func (o *ExtendMysqlInstanceGroupVolume) SetLogDiskSize(v int32)`

SetLogDiskSize sets LogDiskSize field to given value.


### GetDataDiskSize

`func (o *ExtendMysqlInstanceGroupVolume) GetDataDiskSize() int32`

GetDataDiskSize returns the DataDiskSize field if non-nil, zero value otherwise.

### GetDataDiskSizeOk

`func (o *ExtendMysqlInstanceGroupVolume) GetDataDiskSizeOk() (*int32, bool)`

GetDataDiskSizeOk returns a tuple with the DataDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDiskSize

`func (o *ExtendMysqlInstanceGroupVolume) SetDataDiskSize(v int32)`

SetDataDiskSize sets DataDiskSize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


