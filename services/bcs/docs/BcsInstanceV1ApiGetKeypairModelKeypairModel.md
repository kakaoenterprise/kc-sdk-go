# BcsInstanceV1ApiGetKeypairModelKeypairModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 키 페어의 고유 ID | 
**Name** | Pointer to **NullableString** |  | [optional] 
**UserId** | Pointer to **NullableString** |  | [optional] 
**Fingerprint** | Pointer to **NullableString** |  | [optional] 
**PublicKey** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewBcsInstanceV1ApiGetKeypairModelKeypairModel

`func NewBcsInstanceV1ApiGetKeypairModelKeypairModel(id string, ) *BcsInstanceV1ApiGetKeypairModelKeypairModel`

NewBcsInstanceV1ApiGetKeypairModelKeypairModel instantiates a new BcsInstanceV1ApiGetKeypairModelKeypairModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBcsInstanceV1ApiGetKeypairModelKeypairModelWithDefaults

`func NewBcsInstanceV1ApiGetKeypairModelKeypairModelWithDefaults() *BcsInstanceV1ApiGetKeypairModelKeypairModel`

NewBcsInstanceV1ApiGetKeypairModelKeypairModelWithDefaults instantiates a new BcsInstanceV1ApiGetKeypairModelKeypairModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BcsInstanceV1ApiGetKeypairModelKeypairModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BcsInstanceV1ApiGetKeypairModelKeypairModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BcsInstanceV1ApiGetKeypairModelKeypairModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BcsInstanceV1ApiGetKeypairModelKeypairModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BcsInstanceV1ApiGetKeypairModelKeypairModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BcsInstanceV1ApiGetKeypairModelKeypairModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BcsInstanceV1ApiGetKeypairModelKeypairModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BcsInstanceV1ApiGetKeypairModelKeypairModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BcsInstanceV1ApiGetKeypairModelKeypairModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetUserId

`func (o *BcsInstanceV1ApiGetKeypairModelKeypairModel) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *BcsInstanceV1ApiGetKeypairModelKeypairModel) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *BcsInstanceV1ApiGetKeypairModelKeypairModel) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *BcsInstanceV1ApiGetKeypairModelKeypairModel) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *BcsInstanceV1ApiGetKeypairModelKeypairModel) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *BcsInstanceV1ApiGetKeypairModelKeypairModel) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetFingerprint

`func (o *BcsInstanceV1ApiGetKeypairModelKeypairModel) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *BcsInstanceV1ApiGetKeypairModelKeypairModel) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *BcsInstanceV1ApiGetKeypairModelKeypairModel) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *BcsInstanceV1ApiGetKeypairModelKeypairModel) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### SetFingerprintNil

`func (o *BcsInstanceV1ApiGetKeypairModelKeypairModel) SetFingerprintNil(b bool)`

 SetFingerprintNil sets the value for Fingerprint to be an explicit nil

### UnsetFingerprint
`func (o *BcsInstanceV1ApiGetKeypairModelKeypairModel) UnsetFingerprint()`

UnsetFingerprint ensures that no value is present for Fingerprint, not even an explicit nil
### GetPublicKey

`func (o *BcsInstanceV1ApiGetKeypairModelKeypairModel) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *BcsInstanceV1ApiGetKeypairModelKeypairModel) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *BcsInstanceV1ApiGetKeypairModelKeypairModel) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *BcsInstanceV1ApiGetKeypairModelKeypairModel) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### SetPublicKeyNil

`func (o *BcsInstanceV1ApiGetKeypairModelKeypairModel) SetPublicKeyNil(b bool)`

 SetPublicKeyNil sets the value for PublicKey to be an explicit nil

### UnsetPublicKey
`func (o *BcsInstanceV1ApiGetKeypairModelKeypairModel) UnsetPublicKey()`

UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil
### GetType

`func (o *BcsInstanceV1ApiGetKeypairModelKeypairModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BcsInstanceV1ApiGetKeypairModelKeypairModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BcsInstanceV1ApiGetKeypairModelKeypairModel) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BcsInstanceV1ApiGetKeypairModelKeypairModel) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *BcsInstanceV1ApiGetKeypairModelKeypairModel) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *BcsInstanceV1ApiGetKeypairModelKeypairModel) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetCreatedAt

`func (o *BcsInstanceV1ApiGetKeypairModelKeypairModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BcsInstanceV1ApiGetKeypairModelKeypairModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BcsInstanceV1ApiGetKeypairModelKeypairModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BcsInstanceV1ApiGetKeypairModelKeypairModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *BcsInstanceV1ApiGetKeypairModelKeypairModel) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *BcsInstanceV1ApiGetKeypairModelKeypairModel) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


