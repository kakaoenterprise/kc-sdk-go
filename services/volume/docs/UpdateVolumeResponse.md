# UpdateVolumeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Volume** | [**VolumeResult**](VolumeResult.md) | 수정된 볼륨 정보 | 

## Methods

### NewUpdateVolumeResponse

`func NewUpdateVolumeResponse(volume VolumeResult, ) *UpdateVolumeResponse`

NewUpdateVolumeResponse instantiates a new UpdateVolumeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateVolumeResponseWithDefaults

`func NewUpdateVolumeResponseWithDefaults() *UpdateVolumeResponse`

NewUpdateVolumeResponseWithDefaults instantiates a new UpdateVolumeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVolume

`func (o *UpdateVolumeResponse) GetVolume() VolumeResult`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *UpdateVolumeResponse) GetVolumeOk() (*VolumeResult, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *UpdateVolumeResponse) SetVolume(v VolumeResult)`

SetVolume sets Volume field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


