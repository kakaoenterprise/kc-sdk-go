# CreateVolumeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Volume** | [**VolumeResult**](VolumeResult.md) | 생성된 볼륨 정보 | 

## Methods

### NewCreateVolumeResponse

`func NewCreateVolumeResponse(volume VolumeResult, ) *CreateVolumeResponse`

NewCreateVolumeResponse instantiates a new CreateVolumeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVolumeResponseWithDefaults

`func NewCreateVolumeResponseWithDefaults() *CreateVolumeResponse`

NewCreateVolumeResponseWithDefaults instantiates a new CreateVolumeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVolume

`func (o *CreateVolumeResponse) GetVolume() VolumeResult`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *CreateVolumeResponse) GetVolumeOk() (*VolumeResult, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *CreateVolumeResponse) SetVolume(v VolumeResult)`

SetVolume sets Volume field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


