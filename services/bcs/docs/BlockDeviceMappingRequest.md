# BlockDeviceMappingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsDeleteOnTermination** | Pointer to **NullableBool** |  | [optional] 
**Size** | Pointer to **NullableInt32** |  | [optional] 
**SourceType** | Pointer to [**NullableSourceType**](SourceType.md) |  | [optional] 
**Uuid** | Pointer to **NullableString** |  | [optional] 
**TypeId** | Pointer to **NullableString** |  | [optional] 
**EncryptionSecretId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBlockDeviceMappingRequest

`func NewBlockDeviceMappingRequest() *BlockDeviceMappingRequest`

NewBlockDeviceMappingRequest instantiates a new BlockDeviceMappingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockDeviceMappingRequestWithDefaults

`func NewBlockDeviceMappingRequestWithDefaults() *BlockDeviceMappingRequest`

NewBlockDeviceMappingRequestWithDefaults instantiates a new BlockDeviceMappingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsDeleteOnTermination

`func (o *BlockDeviceMappingRequest) GetIsDeleteOnTermination() bool`

GetIsDeleteOnTermination returns the IsDeleteOnTermination field if non-nil, zero value otherwise.

### GetIsDeleteOnTerminationOk

`func (o *BlockDeviceMappingRequest) GetIsDeleteOnTerminationOk() (*bool, bool)`

GetIsDeleteOnTerminationOk returns a tuple with the IsDeleteOnTermination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleteOnTermination

`func (o *BlockDeviceMappingRequest) SetIsDeleteOnTermination(v bool)`

SetIsDeleteOnTermination sets IsDeleteOnTermination field to given value.

### HasIsDeleteOnTermination

`func (o *BlockDeviceMappingRequest) HasIsDeleteOnTermination() bool`

HasIsDeleteOnTermination returns a boolean if a field has been set.

### SetIsDeleteOnTerminationNil

`func (o *BlockDeviceMappingRequest) SetIsDeleteOnTerminationNil(b bool)`

 SetIsDeleteOnTerminationNil sets the value for IsDeleteOnTermination to be an explicit nil

### UnsetIsDeleteOnTermination
`func (o *BlockDeviceMappingRequest) UnsetIsDeleteOnTermination()`

UnsetIsDeleteOnTermination ensures that no value is present for IsDeleteOnTermination, not even an explicit nil
### GetSize

`func (o *BlockDeviceMappingRequest) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *BlockDeviceMappingRequest) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *BlockDeviceMappingRequest) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *BlockDeviceMappingRequest) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *BlockDeviceMappingRequest) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *BlockDeviceMappingRequest) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetSourceType

`func (o *BlockDeviceMappingRequest) GetSourceType() SourceType`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *BlockDeviceMappingRequest) GetSourceTypeOk() (*SourceType, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *BlockDeviceMappingRequest) SetSourceType(v SourceType)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *BlockDeviceMappingRequest) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### SetSourceTypeNil

`func (o *BlockDeviceMappingRequest) SetSourceTypeNil(b bool)`

 SetSourceTypeNil sets the value for SourceType to be an explicit nil

### UnsetSourceType
`func (o *BlockDeviceMappingRequest) UnsetSourceType()`

UnsetSourceType ensures that no value is present for SourceType, not even an explicit nil
### GetUuid

`func (o *BlockDeviceMappingRequest) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *BlockDeviceMappingRequest) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *BlockDeviceMappingRequest) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *BlockDeviceMappingRequest) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuidNil

`func (o *BlockDeviceMappingRequest) SetUuidNil(b bool)`

 SetUuidNil sets the value for Uuid to be an explicit nil

### UnsetUuid
`func (o *BlockDeviceMappingRequest) UnsetUuid()`

UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
### GetTypeId

`func (o *BlockDeviceMappingRequest) GetTypeId() string`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *BlockDeviceMappingRequest) GetTypeIdOk() (*string, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *BlockDeviceMappingRequest) SetTypeId(v string)`

SetTypeId sets TypeId field to given value.

### HasTypeId

`func (o *BlockDeviceMappingRequest) HasTypeId() bool`

HasTypeId returns a boolean if a field has been set.

### SetTypeIdNil

`func (o *BlockDeviceMappingRequest) SetTypeIdNil(b bool)`

 SetTypeIdNil sets the value for TypeId to be an explicit nil

### UnsetTypeId
`func (o *BlockDeviceMappingRequest) UnsetTypeId()`

UnsetTypeId ensures that no value is present for TypeId, not even an explicit nil
### GetEncryptionSecretId

`func (o *BlockDeviceMappingRequest) GetEncryptionSecretId() string`

GetEncryptionSecretId returns the EncryptionSecretId field if non-nil, zero value otherwise.

### GetEncryptionSecretIdOk

`func (o *BlockDeviceMappingRequest) GetEncryptionSecretIdOk() (*string, bool)`

GetEncryptionSecretIdOk returns a tuple with the EncryptionSecretId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionSecretId

`func (o *BlockDeviceMappingRequest) SetEncryptionSecretId(v string)`

SetEncryptionSecretId sets EncryptionSecretId field to given value.

### HasEncryptionSecretId

`func (o *BlockDeviceMappingRequest) HasEncryptionSecretId() bool`

HasEncryptionSecretId returns a boolean if a field has been set.

### SetEncryptionSecretIdNil

`func (o *BlockDeviceMappingRequest) SetEncryptionSecretIdNil(b bool)`

 SetEncryptionSecretIdNil sets the value for EncryptionSecretId to be an explicit nil

### UnsetEncryptionSecretId
`func (o *BlockDeviceMappingRequest) UnsetEncryptionSecretId()`

UnsetEncryptionSecretId ensures that no value is present for EncryptionSecretId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


