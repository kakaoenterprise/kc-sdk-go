# AttachVolumeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Volume** | [**VolumeAttachment**](VolumeAttachment.md) | 볼륨 정보 | 

## Methods

### NewAttachVolumeResponse

`func NewAttachVolumeResponse(volume VolumeAttachment, ) *AttachVolumeResponse`

NewAttachVolumeResponse instantiates a new AttachVolumeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachVolumeResponseWithDefaults

`func NewAttachVolumeResponseWithDefaults() *AttachVolumeResponse`

NewAttachVolumeResponseWithDefaults instantiates a new AttachVolumeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVolume

`func (o *AttachVolumeResponse) GetVolume() VolumeAttachment`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *AttachVolumeResponse) GetVolumeOk() (*VolumeAttachment, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *AttachVolumeResponse) SetVolume(v VolumeAttachment)`

SetVolume sets Volume field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


