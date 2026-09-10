# UpdateVolumeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Volume** | [**UpdateVolume**](UpdateVolume.md) | 수정할 볼륨 정보 | 

## Methods

### NewUpdateVolumeRequest

`func NewUpdateVolumeRequest(volume UpdateVolume, ) *UpdateVolumeRequest`

NewUpdateVolumeRequest instantiates a new UpdateVolumeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateVolumeRequestWithDefaults

`func NewUpdateVolumeRequestWithDefaults() *UpdateVolumeRequest`

NewUpdateVolumeRequestWithDefaults instantiates a new UpdateVolumeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVolume

`func (o *UpdateVolumeRequest) GetVolume() UpdateVolume`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *UpdateVolumeRequest) GetVolumeOk() (*UpdateVolume, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *UpdateVolumeRequest) SetVolume(v UpdateVolume)`

SetVolume sets Volume field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


