# SymmetricKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | KMS 키의 고유 ID | [optional] 
**Name** | Pointer to **NullableString** | KMS 키 이름 | [optional] 
**Type** | Pointer to **NullableString** | 키 유형 | [optional] 

## Methods

### NewSymmetricKey

`func NewSymmetricKey() *SymmetricKey`

NewSymmetricKey instantiates a new SymmetricKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSymmetricKeyWithDefaults

`func NewSymmetricKeyWithDefaults() *SymmetricKey`

NewSymmetricKeyWithDefaults instantiates a new SymmetricKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SymmetricKey) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SymmetricKey) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SymmetricKey) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SymmetricKey) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SymmetricKey) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SymmetricKey) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *SymmetricKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SymmetricKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SymmetricKey) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SymmetricKey) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SymmetricKey) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SymmetricKey) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetType

`func (o *SymmetricKey) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SymmetricKey) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SymmetricKey) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SymmetricKey) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *SymmetricKey) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *SymmetricKey) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


