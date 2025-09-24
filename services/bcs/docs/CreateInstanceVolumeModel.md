# CreateInstanceVolumeModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsDeleteOnTermination** | Pointer to **bool** | 인스턴스 삭제 시 해당 볼륨을 자동으로 삭제할지 여부 &lt;br/&gt; - 첫 번째 볼륨: 기본값 &#x60;true&#x60; &lt;br/&gt; - 두 번째 이후 볼륨: 기본값 &#x60;false&#x60; | [optional] [default to false]
**Size** | **int32** | 볼륨의 크기(GB) &lt;br/&gt; - 허용 범위: 1 ~ 16384 &lt;br/&gt; - &#x60;source_type &#x3D; image&#x60;일 경우, 해당 이미지의 &#x60;min_disk&#x60; 이상이어야 하며, Windows 이미지의 경우 최대 2048GB까지 지원 | 
**SourceType** | Pointer to [**SourceType**](SourceType.md) | 볼륨의 소스 유형 &lt;br/&gt; - &#x60;image&#x60;: 이미지로부터 생성 &lt;br/&gt; - &#x60;blank&#x60;: 빈 볼륨으로 생성 (첫 번째 이후 볼륨의 기본값) &lt;br/&gt; - &#x60;volume&#x60;: 기존 볼륨으로부터 생성 | [optional] 
**Uuid** | Pointer to **NullableString** |  | [optional] 
**TypeId** | Pointer to **NullableString** |  | [optional] 
**EncryptionSecretId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreateInstanceVolumeModel

`func NewCreateInstanceVolumeModel(size int32, ) *CreateInstanceVolumeModel`

NewCreateInstanceVolumeModel instantiates a new CreateInstanceVolumeModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInstanceVolumeModelWithDefaults

`func NewCreateInstanceVolumeModelWithDefaults() *CreateInstanceVolumeModel`

NewCreateInstanceVolumeModelWithDefaults instantiates a new CreateInstanceVolumeModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsDeleteOnTermination

`func (o *CreateInstanceVolumeModel) GetIsDeleteOnTermination() bool`

GetIsDeleteOnTermination returns the IsDeleteOnTermination field if non-nil, zero value otherwise.

### GetIsDeleteOnTerminationOk

`func (o *CreateInstanceVolumeModel) GetIsDeleteOnTerminationOk() (*bool, bool)`

GetIsDeleteOnTerminationOk returns a tuple with the IsDeleteOnTermination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleteOnTermination

`func (o *CreateInstanceVolumeModel) SetIsDeleteOnTermination(v bool)`

SetIsDeleteOnTermination sets IsDeleteOnTermination field to given value.

### HasIsDeleteOnTermination

`func (o *CreateInstanceVolumeModel) HasIsDeleteOnTermination() bool`

HasIsDeleteOnTermination returns a boolean if a field has been set.

### GetSize

`func (o *CreateInstanceVolumeModel) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CreateInstanceVolumeModel) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CreateInstanceVolumeModel) SetSize(v int32)`

SetSize sets Size field to given value.


### GetSourceType

`func (o *CreateInstanceVolumeModel) GetSourceType() SourceType`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *CreateInstanceVolumeModel) GetSourceTypeOk() (*SourceType, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *CreateInstanceVolumeModel) SetSourceType(v SourceType)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *CreateInstanceVolumeModel) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetUuid

`func (o *CreateInstanceVolumeModel) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *CreateInstanceVolumeModel) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *CreateInstanceVolumeModel) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *CreateInstanceVolumeModel) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *CreateInstanceVolumeModel) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *CreateInstanceVolumeModel) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetTypeId

`func (o *CreateInstanceVolumeModel) GetTypeId() string`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *CreateInstanceVolumeModel) GetTypeIdOk() (*string, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *CreateInstanceVolumeModel) SetTypeId(v string)`

SetTypeId sets TypeId field to given value.

### HasTypeId

`func (o *CreateInstanceVolumeModel) HasTypeId() bool`

HasTypeId returns a boolean if a field has been set.

### SetTypeIdNil

`func (o *CreateInstanceVolumeModel) SetTypeIdNil(b bool)`

 SetTypeIdNil sets the value for TypeId to be an explicit nil

### UnsetTypeId
`func (o *CreateInstanceVolumeModel) UnsetTypeId()`

UnsetTypeId ensures that no value is present for TypeId, not even an explicit nil
### GetEncryptionSecretId

`func (o *CreateInstanceVolumeModel) GetEncryptionSecretId() string`

GetEncryptionSecretId returns the EncryptionSecretId field if non-nil, zero value otherwise.

### GetEncryptionSecretIdOk

`func (o *CreateInstanceVolumeModel) GetEncryptionSecretIdOk() (*string, bool)`

GetEncryptionSecretIdOk returns a tuple with the EncryptionSecretId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionSecretId

`func (o *CreateInstanceVolumeModel) SetEncryptionSecretId(v string)`

SetEncryptionSecretId sets EncryptionSecretId field to given value.

### HasEncryptionSecretId

`func (o *CreateInstanceVolumeModel) HasEncryptionSecretId() bool`

HasEncryptionSecretId returns a boolean if a field has been set.

### SetEncryptionSecretIdNil

`func (o *CreateInstanceVolumeModel) SetEncryptionSecretIdNil(b bool)`

 SetEncryptionSecretIdNil sets the value for EncryptionSecretId to be an explicit nil

### UnsetEncryptionSecretId
`func (o *CreateInstanceVolumeModel) UnsetEncryptionSecretId()`

UnsetEncryptionSecretId ensures that no value is present for EncryptionSecretId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


