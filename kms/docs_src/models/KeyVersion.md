# KeyVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | 키 버전의 고유 ID | [optional] 
**Version** | Pointer to **NullableInt32** | 키 버전 | [optional] 
**Status** | Pointer to [**NullableKeyVersionStatus**](KeyVersionStatus.md) | 키 버전 상태 | [optional] 
**Algorithm** | Pointer to [**NullableKeyAlgorithm**](KeyAlgorithm.md) | 키 알고리즘 | [optional] 
**ResolvedDestructionAt** | Pointer to **NullableString** | 키 버전이 실제로 폐기된 일시 | [optional] 
**CreatedAt** | Pointer to **NullableString** | 버전이 생성된 일시 | [optional] 

## Methods

### NewKeyVersion

`func NewKeyVersion() *KeyVersion`

NewKeyVersion instantiates a new KeyVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyVersionWithDefaults

`func NewKeyVersionWithDefaults() *KeyVersion`

NewKeyVersionWithDefaults instantiates a new KeyVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KeyVersion) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyVersion) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyVersion) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KeyVersion) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *KeyVersion) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *KeyVersion) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetVersion

`func (o *KeyVersion) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *KeyVersion) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *KeyVersion) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *KeyVersion) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *KeyVersion) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *KeyVersion) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetStatus

`func (o *KeyVersion) GetStatus() KeyVersionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KeyVersion) GetStatusOk() (*KeyVersionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KeyVersion) SetStatus(v KeyVersionStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KeyVersion) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *KeyVersion) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *KeyVersion) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetAlgorithm

`func (o *KeyVersion) GetAlgorithm() KeyAlgorithm`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *KeyVersion) GetAlgorithmOk() (*KeyAlgorithm, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *KeyVersion) SetAlgorithm(v KeyAlgorithm)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *KeyVersion) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### SetAlgorithmNil

`func (o *KeyVersion) SetAlgorithmNil(b bool)`

 SetAlgorithmNil sets the value for Algorithm to be an explicit nil

### UnsetAlgorithm
`func (o *KeyVersion) UnsetAlgorithm()`

UnsetAlgorithm ensures that no value is present for Algorithm, not even an explicit nil
### GetResolvedDestructionAt

`func (o *KeyVersion) GetResolvedDestructionAt() string`

GetResolvedDestructionAt returns the ResolvedDestructionAt field if non-nil, zero value otherwise.

### GetResolvedDestructionAtOk

`func (o *KeyVersion) GetResolvedDestructionAtOk() (*string, bool)`

GetResolvedDestructionAtOk returns a tuple with the ResolvedDestructionAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedDestructionAt

`func (o *KeyVersion) SetResolvedDestructionAt(v string)`

SetResolvedDestructionAt sets ResolvedDestructionAt field to given value.

### HasResolvedDestructionAt

`func (o *KeyVersion) HasResolvedDestructionAt() bool`

HasResolvedDestructionAt returns a boolean if a field has been set.

### SetResolvedDestructionAtNil

`func (o *KeyVersion) SetResolvedDestructionAtNil(b bool)`

 SetResolvedDestructionAtNil sets the value for ResolvedDestructionAt to be an explicit nil

### UnsetResolvedDestructionAt
`func (o *KeyVersion) UnsetResolvedDestructionAt()`

UnsetResolvedDestructionAt ensures that no value is present for ResolvedDestructionAt, not even an explicit nil
### GetCreatedAt

`func (o *KeyVersion) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KeyVersion) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KeyVersion) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KeyVersion) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *KeyVersion) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *KeyVersion) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


