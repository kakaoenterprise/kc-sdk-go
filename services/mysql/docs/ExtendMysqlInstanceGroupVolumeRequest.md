# ExtendMysqlInstanceGroupVolumeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceGroup** | [**ExtendMysqlInstanceGroupVolume**](ExtendMysqlInstanceGroupVolume.md) | MySQL 인스턴스 그룹 볼륨 확장 정보 | 

## Methods

### NewExtendMysqlInstanceGroupVolumeRequest

`func NewExtendMysqlInstanceGroupVolumeRequest(instanceGroup ExtendMysqlInstanceGroupVolume, ) *ExtendMysqlInstanceGroupVolumeRequest`

NewExtendMysqlInstanceGroupVolumeRequest instantiates a new ExtendMysqlInstanceGroupVolumeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtendMysqlInstanceGroupVolumeRequestWithDefaults

`func NewExtendMysqlInstanceGroupVolumeRequestWithDefaults() *ExtendMysqlInstanceGroupVolumeRequest`

NewExtendMysqlInstanceGroupVolumeRequestWithDefaults instantiates a new ExtendMysqlInstanceGroupVolumeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceGroup

`func (o *ExtendMysqlInstanceGroupVolumeRequest) GetInstanceGroup() ExtendMysqlInstanceGroupVolume`

GetInstanceGroup returns the InstanceGroup field if non-nil, zero value otherwise.

### GetInstanceGroupOk

`func (o *ExtendMysqlInstanceGroupVolumeRequest) GetInstanceGroupOk() (*ExtendMysqlInstanceGroupVolume, bool)`

GetInstanceGroupOk returns a tuple with the InstanceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceGroup

`func (o *ExtendMysqlInstanceGroupVolumeRequest) SetInstanceGroup(v ExtendMysqlInstanceGroupVolume)`

SetInstanceGroup sets InstanceGroup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


