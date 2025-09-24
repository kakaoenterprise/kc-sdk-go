# CreateVolumeModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | 생성할 볼륨의 이름 | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Size** | **int32** | 볼륨 크기 (GB 단위)&lt;br/&gt;- Linux 계열: 1 ~ 16,384 GB &lt;br/&gt;- Windows 계열: 1 ~ 2,048 GB | 
**AvailabilityZone** | [**AvailabilityZone**](AvailabilityZone.md) | 볼륨을 생성할 가용 영역&lt;br/&gt; - &#x60;kr-central-2-a&#x60;: kr-central-2-a 가용 영역 &lt;br/&gt; - &#x60;kr-central-2-b&#x60;: kr-central-2-b 가용 영역 &lt;br/&gt; - &#x60;kr-central-2-c&#x60;: kr-central-2-c 가용 영역 | 
**VolumeTypeId** | Pointer to **NullableString** |  | [optional] 
**SourceVolumeId** | Pointer to **NullableString** |  | [optional] 
**EncryptionSecretId** | Pointer to **NullableString** |  | [optional] 
**ImageId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreateVolumeModel

`func NewCreateVolumeModel(name string, size int32, availabilityZone AvailabilityZone, ) *CreateVolumeModel`

NewCreateVolumeModel instantiates a new CreateVolumeModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVolumeModelWithDefaults

`func NewCreateVolumeModelWithDefaults() *CreateVolumeModel`

NewCreateVolumeModelWithDefaults instantiates a new CreateVolumeModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateVolumeModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateVolumeModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateVolumeModel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateVolumeModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateVolumeModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateVolumeModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateVolumeModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateVolumeModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateVolumeModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSize

`func (o *CreateVolumeModel) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CreateVolumeModel) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CreateVolumeModel) SetSize(v int32)`

SetSize sets Size field to given value.


### GetAvailabilityZone

`func (o *CreateVolumeModel) GetAvailabilityZone() AvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *CreateVolumeModel) GetAvailabilityZoneOk() (*AvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *CreateVolumeModel) SetAvailabilityZone(v AvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.


### GetVolumeTypeId

`func (o *CreateVolumeModel) GetVolumeTypeId() string`

GetVolumeTypeId returns the VolumeTypeId field if non-nil, zero value otherwise.

### GetVolumeTypeIdOk

`func (o *CreateVolumeModel) GetVolumeTypeIdOk() (*string, bool)`

GetVolumeTypeIdOk returns a tuple with the VolumeTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeTypeId

`func (o *CreateVolumeModel) SetVolumeTypeId(v string)`

SetVolumeTypeId sets VolumeTypeId field to given value.

### HasVolumeTypeId

`func (o *CreateVolumeModel) HasVolumeTypeId() bool`

HasVolumeTypeId returns a boolean if a field has been set.

### SetVolumeTypeIdNil

`func (o *CreateVolumeModel) SetVolumeTypeIdNil(b bool)`

 SetVolumeTypeIdNil sets the value for VolumeTypeId to be an explicit nil

### UnsetVolumeTypeId
`func (o *CreateVolumeModel) UnsetVolumeTypeId()`

UnsetVolumeTypeId ensures that no value is present for VolumeTypeId, not even an explicit nil
### GetSourceVolumeId

`func (o *CreateVolumeModel) GetSourceVolumeId() string`

GetSourceVolumeId returns the SourceVolumeId field if non-nil, zero value otherwise.

### GetSourceVolumeIdOk

`func (o *CreateVolumeModel) GetSourceVolumeIdOk() (*string, bool)`

GetSourceVolumeIdOk returns a tuple with the SourceVolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVolumeId

`func (o *CreateVolumeModel) SetSourceVolumeId(v string)`

SetSourceVolumeId sets SourceVolumeId field to given value.

### HasSourceVolumeId

`func (o *CreateVolumeModel) HasSourceVolumeId() bool`

HasSourceVolumeId returns a boolean if a field has been set.

### SetSourceVolumeIdNil

`func (o *CreateVolumeModel) SetSourceVolumeIdNil(b bool)`

 SetSourceVolumeIdNil sets the value for SourceVolumeId to be an explicit nil

### UnsetSourceVolumeId
`func (o *CreateVolumeModel) UnsetSourceVolumeId()`

UnsetSourceVolumeId ensures that no value is present for SourceVolumeId, not even an explicit nil
### GetEncryptionSecretId

`func (o *CreateVolumeModel) GetEncryptionSecretId() string`

GetEncryptionSecretId returns the EncryptionSecretId field if non-nil, zero value otherwise.

### GetEncryptionSecretIdOk

`func (o *CreateVolumeModel) GetEncryptionSecretIdOk() (*string, bool)`

GetEncryptionSecretIdOk returns a tuple with the EncryptionSecretId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionSecretId

`func (o *CreateVolumeModel) SetEncryptionSecretId(v string)`

SetEncryptionSecretId sets EncryptionSecretId field to given value.

### HasEncryptionSecretId

`func (o *CreateVolumeModel) HasEncryptionSecretId() bool`

HasEncryptionSecretId returns a boolean if a field has been set.

### SetEncryptionSecretIdNil

`func (o *CreateVolumeModel) SetEncryptionSecretIdNil(b bool)`

 SetEncryptionSecretIdNil sets the value for EncryptionSecretId to be an explicit nil

### UnsetEncryptionSecretId
`func (o *CreateVolumeModel) UnsetEncryptionSecretId()`

UnsetEncryptionSecretId ensures that no value is present for EncryptionSecretId, not even an explicit nil
### GetImageId

`func (o *CreateVolumeModel) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *CreateVolumeModel) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *CreateVolumeModel) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *CreateVolumeModel) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### SetImageIdNil

`func (o *CreateVolumeModel) SetImageIdNil(b bool)`

 SetImageIdNil sets the value for ImageId to be an explicit nil

### UnsetImageId
`func (o *CreateVolumeModel) UnsetImageId()`

UnsetImageId ensures that no value is present for ImageId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


