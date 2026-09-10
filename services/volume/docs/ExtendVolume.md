# ExtendVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewSize** | **int32** | 변경할 볼륨의 새로운 크기 (GB 단위) &lt;br/&gt; - 기존 크기보다 커야 함 | 

## Methods

### NewExtendVolume

`func NewExtendVolume(newSize int32, ) *ExtendVolume`

NewExtendVolume instantiates a new ExtendVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtendVolumeWithDefaults

`func NewExtendVolumeWithDefaults() *ExtendVolume`

NewExtendVolumeWithDefaults instantiates a new ExtendVolume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewSize

`func (o *ExtendVolume) GetNewSize() int32`

GetNewSize returns the NewSize field if non-nil, zero value otherwise.

### GetNewSizeOk

`func (o *ExtendVolume) GetNewSizeOk() (*int32, bool)`

GetNewSizeOk returns a tuple with the NewSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSize

`func (o *ExtendVolume) SetNewSize(v int32)`

SetNewSize sets NewSize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


