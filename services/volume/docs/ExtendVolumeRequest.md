# ExtendVolumeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Volume** | [**ExtendVolume**](ExtendVolume.md) | 확장할 볼륨 크기 정보 | 

## Methods

### NewExtendVolumeRequest

`func NewExtendVolumeRequest(volume ExtendVolume, ) *ExtendVolumeRequest`

NewExtendVolumeRequest instantiates a new ExtendVolumeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtendVolumeRequestWithDefaults

`func NewExtendVolumeRequestWithDefaults() *ExtendVolumeRequest`

NewExtendVolumeRequestWithDefaults instantiates a new ExtendVolumeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVolume

`func (o *ExtendVolumeRequest) GetVolume() ExtendVolume`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *ExtendVolumeRequest) GetVolumeOk() (*ExtendVolume, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *ExtendVolumeRequest) SetVolume(v ExtendVolume)`

SetVolume sets Volume field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


