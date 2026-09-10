# Algorithm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to [**NullableKeyAlgorithm**](KeyAlgorithm.md) |  | [optional] 
**Purpose** | Pointer to [**NullableKeyPurpose**](KeyPurpose.md) |  | [optional] 
**Type** | Pointer to [**NullableKeyType**](KeyType.md) |  | [optional] 

## Methods

### NewAlgorithm

`func NewAlgorithm() *Algorithm`

NewAlgorithm instantiates a new Algorithm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlgorithmWithDefaults

`func NewAlgorithmWithDefaults() *Algorithm`

NewAlgorithmWithDefaults instantiates a new Algorithm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Algorithm) GetName() KeyAlgorithm`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Algorithm) GetNameOk() (*KeyAlgorithm, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Algorithm) SetName(v KeyAlgorithm)`

SetName sets Name field to given value.

### HasName

`func (o *Algorithm) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Algorithm) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Algorithm) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPurpose

`func (o *Algorithm) GetPurpose() KeyPurpose`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *Algorithm) GetPurposeOk() (*KeyPurpose, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *Algorithm) SetPurpose(v KeyPurpose)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *Algorithm) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### SetPurposeNil

`func (o *Algorithm) SetPurposeNil(b bool)`

 SetPurposeNil sets the value for Purpose to be an explicit nil

### UnsetPurpose
`func (o *Algorithm) UnsetPurpose()`

UnsetPurpose ensures that no value is present for Purpose, not even an explicit nil
### GetType

`func (o *Algorithm) GetType() KeyType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Algorithm) GetTypeOk() (*KeyType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Algorithm) SetType(v KeyType)`

SetType sets Type field to given value.

### HasType

`func (o *Algorithm) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *Algorithm) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *Algorithm) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


