# CreateVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | 생성할 볼륨의 이름 | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Size** | **int32** | 볼륨 크기 (GB 단위)&lt;br/&gt;- Linux 계열: 1 ~ 16,384 GB &lt;br/&gt;- Windows 계열: 1 ~ 2,048 GB | 
**AvailabilityZone** | [**AvailabilityZone**](AvailabilityZone.md) | 볼륨을 생성할 가용 영역 | 
**VolumeTypeId** | Pointer to **NullableString** |  | [optional] 
**SourceVolumeId** | Pointer to **NullableString** |  | [optional] 
**EncryptionSecretId** | Pointer to **NullableString** |  | [optional] 
**ImageId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreateVolume

`func NewCreateVolume(name string, size int32, availabilityZone AvailabilityZone, ) *CreateVolume`

NewCreateVolume instantiates a new CreateVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVolumeWithDefaults

`func NewCreateVolumeWithDefaults() *CreateVolume`

NewCreateVolumeWithDefaults instantiates a new CreateVolume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateVolume) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateVolume) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateVolume) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateVolume) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateVolume) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateVolume) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateVolume) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateVolume) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateVolume) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSize

`func (o *CreateVolume) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CreateVolume) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CreateVolume) SetSize(v int32)`

SetSize sets Size field to given value.


### GetAvailabilityZone

`func (o *CreateVolume) GetAvailabilityZone() AvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *CreateVolume) GetAvailabilityZoneOk() (*AvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *CreateVolume) SetAvailabilityZone(v AvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.


### GetVolumeTypeId

`func (o *CreateVolume) GetVolumeTypeId() string`

GetVolumeTypeId returns the VolumeTypeId field if non-nil, zero value otherwise.

### GetVolumeTypeIdOk

`func (o *CreateVolume) GetVolumeTypeIdOk() (*string, bool)`

GetVolumeTypeIdOk returns a tuple with the VolumeTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeTypeId

`func (o *CreateVolume) SetVolumeTypeId(v string)`

SetVolumeTypeId sets VolumeTypeId field to given value.

### HasVolumeTypeId

`func (o *CreateVolume) HasVolumeTypeId() bool`

HasVolumeTypeId returns a boolean if a field has been set.

### SetVolumeTypeIdNil

`func (o *CreateVolume) SetVolumeTypeIdNil(b bool)`

 SetVolumeTypeIdNil sets the value for VolumeTypeId to be an explicit nil

### UnsetVolumeTypeId
`func (o *CreateVolume) UnsetVolumeTypeId()`

UnsetVolumeTypeId ensures that no value is present for VolumeTypeId, not even an explicit nil
### GetSourceVolumeId

`func (o *CreateVolume) GetSourceVolumeId() string`

GetSourceVolumeId returns the SourceVolumeId field if non-nil, zero value otherwise.

### GetSourceVolumeIdOk

`func (o *CreateVolume) GetSourceVolumeIdOk() (*string, bool)`

GetSourceVolumeIdOk returns a tuple with the SourceVolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVolumeId

`func (o *CreateVolume) SetSourceVolumeId(v string)`

SetSourceVolumeId sets SourceVolumeId field to given value.

### HasSourceVolumeId

`func (o *CreateVolume) HasSourceVolumeId() bool`

HasSourceVolumeId returns a boolean if a field has been set.

### SetSourceVolumeIdNil

`func (o *CreateVolume) SetSourceVolumeIdNil(b bool)`

 SetSourceVolumeIdNil sets the value for SourceVolumeId to be an explicit nil

### UnsetSourceVolumeId
`func (o *CreateVolume) UnsetSourceVolumeId()`

UnsetSourceVolumeId ensures that no value is present for SourceVolumeId, not even an explicit nil
### GetEncryptionSecretId

`func (o *CreateVolume) GetEncryptionSecretId() string`

GetEncryptionSecretId returns the EncryptionSecretId field if non-nil, zero value otherwise.

### GetEncryptionSecretIdOk

`func (o *CreateVolume) GetEncryptionSecretIdOk() (*string, bool)`

GetEncryptionSecretIdOk returns a tuple with the EncryptionSecretId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionSecretId

`func (o *CreateVolume) SetEncryptionSecretId(v string)`

SetEncryptionSecretId sets EncryptionSecretId field to given value.

### HasEncryptionSecretId

`func (o *CreateVolume) HasEncryptionSecretId() bool`

HasEncryptionSecretId returns a boolean if a field has been set.

### SetEncryptionSecretIdNil

`func (o *CreateVolume) SetEncryptionSecretIdNil(b bool)`

 SetEncryptionSecretIdNil sets the value for EncryptionSecretId to be an explicit nil

### UnsetEncryptionSecretId
`func (o *CreateVolume) UnsetEncryptionSecretId()`

UnsetEncryptionSecretId ensures that no value is present for EncryptionSecretId, not even an explicit nil
### GetImageId

`func (o *CreateVolume) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *CreateVolume) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *CreateVolume) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *CreateVolume) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### SetImageIdNil

`func (o *CreateVolume) SetImageIdNil(b bool)`

 SetImageIdNil sets the value for ImageId to be an explicit nil

### UnsetImageId
`func (o *CreateVolume) UnsetImageId()`

UnsetImageId ensures that no value is present for ImageId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


