# BcsInstanceV1ApiListKeypairsModelKeypairModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 키 ID | 
**Name** | Pointer to **NullableString** |  | [optional] 
**UserId** | Pointer to **NullableString** |  | [optional] 
**Fingerprint** | Pointer to **NullableString** |  | [optional] 
**PublicKey** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewBcsInstanceV1ApiListKeypairsModelKeypairModel

`func NewBcsInstanceV1ApiListKeypairsModelKeypairModel(id string, ) *BcsInstanceV1ApiListKeypairsModelKeypairModel`

NewBcsInstanceV1ApiListKeypairsModelKeypairModel instantiates a new BcsInstanceV1ApiListKeypairsModelKeypairModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBcsInstanceV1ApiListKeypairsModelKeypairModelWithDefaults

`func NewBcsInstanceV1ApiListKeypairsModelKeypairModelWithDefaults() *BcsInstanceV1ApiListKeypairsModelKeypairModel`

NewBcsInstanceV1ApiListKeypairsModelKeypairModelWithDefaults instantiates a new BcsInstanceV1ApiListKeypairsModelKeypairModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BcsInstanceV1ApiListKeypairsModelKeypairModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BcsInstanceV1ApiListKeypairsModelKeypairModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BcsInstanceV1ApiListKeypairsModelKeypairModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BcsInstanceV1ApiListKeypairsModelKeypairModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BcsInstanceV1ApiListKeypairsModelKeypairModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BcsInstanceV1ApiListKeypairsModelKeypairModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BcsInstanceV1ApiListKeypairsModelKeypairModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BcsInstanceV1ApiListKeypairsModelKeypairModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BcsInstanceV1ApiListKeypairsModelKeypairModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetUserId

`func (o *BcsInstanceV1ApiListKeypairsModelKeypairModel) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *BcsInstanceV1ApiListKeypairsModelKeypairModel) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *BcsInstanceV1ApiListKeypairsModelKeypairModel) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *BcsInstanceV1ApiListKeypairsModelKeypairModel) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *BcsInstanceV1ApiListKeypairsModelKeypairModel) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *BcsInstanceV1ApiListKeypairsModelKeypairModel) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetFingerprint

`func (o *BcsInstanceV1ApiListKeypairsModelKeypairModel) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *BcsInstanceV1ApiListKeypairsModelKeypairModel) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *BcsInstanceV1ApiListKeypairsModelKeypairModel) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *BcsInstanceV1ApiListKeypairsModelKeypairModel) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### SetFingerprintNil

`func (o *BcsInstanceV1ApiListKeypairsModelKeypairModel) SetFingerprintNil(b bool)`

 SetFingerprintNil sets the value for Fingerprint to be an explicit nil

### UnsetFingerprint
`func (o *BcsInstanceV1ApiListKeypairsModelKeypairModel) UnsetFingerprint()`

UnsetFingerprint ensures that no value is present for Fingerprint, not even an explicit nil
### GetPublicKey

`func (o *BcsInstanceV1ApiListKeypairsModelKeypairModel) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *BcsInstanceV1ApiListKeypairsModelKeypairModel) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *BcsInstanceV1ApiListKeypairsModelKeypairModel) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *BcsInstanceV1ApiListKeypairsModelKeypairModel) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### SetPublicKeyNil

`func (o *BcsInstanceV1ApiListKeypairsModelKeypairModel) SetPublicKeyNil(b bool)`

 SetPublicKeyNil sets the value for PublicKey to be an explicit nil

### UnsetPublicKey
`func (o *BcsInstanceV1ApiListKeypairsModelKeypairModel) UnsetPublicKey()`

UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil
### GetType

`func (o *BcsInstanceV1ApiListKeypairsModelKeypairModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BcsInstanceV1ApiListKeypairsModelKeypairModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BcsInstanceV1ApiListKeypairsModelKeypairModel) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BcsInstanceV1ApiListKeypairsModelKeypairModel) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *BcsInstanceV1ApiListKeypairsModelKeypairModel) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *BcsInstanceV1ApiListKeypairsModelKeypairModel) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetCreatedAt

`func (o *BcsInstanceV1ApiListKeypairsModelKeypairModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BcsInstanceV1ApiListKeypairsModelKeypairModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BcsInstanceV1ApiListKeypairsModelKeypairModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BcsInstanceV1ApiListKeypairsModelKeypairModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *BcsInstanceV1ApiListKeypairsModelKeypairModel) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *BcsInstanceV1ApiListKeypairsModelKeypairModel) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


