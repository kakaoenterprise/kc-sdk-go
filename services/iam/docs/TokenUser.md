# TokenUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to [**TokenDomain**](TokenDomain.md) |  | [optional] 
**DomainId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**PasswordExpiresAt** | Pointer to **string** |  | [optional] 

## Methods

### NewTokenUser

`func NewTokenUser() *TokenUser`

NewTokenUser instantiates a new TokenUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenUserWithDefaults

`func NewTokenUserWithDefaults() *TokenUser`

NewTokenUserWithDefaults instantiates a new TokenUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *TokenUser) GetDomain() TokenDomain`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *TokenUser) GetDomainOk() (*TokenDomain, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *TokenUser) SetDomain(v TokenDomain)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *TokenUser) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetDomainId

`func (o *TokenUser) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *TokenUser) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *TokenUser) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.

### HasDomainId

`func (o *TokenUser) HasDomainId() bool`

HasDomainId returns a boolean if a field has been set.

### GetId

`func (o *TokenUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TokenUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TokenUser) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TokenUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TokenUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TokenUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TokenUser) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TokenUser) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPassword

`func (o *TokenUser) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *TokenUser) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *TokenUser) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *TokenUser) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPasswordExpiresAt

`func (o *TokenUser) GetPasswordExpiresAt() string`

GetPasswordExpiresAt returns the PasswordExpiresAt field if non-nil, zero value otherwise.

### GetPasswordExpiresAtOk

`func (o *TokenUser) GetPasswordExpiresAtOk() (*string, bool)`

GetPasswordExpiresAtOk returns a tuple with the PasswordExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordExpiresAt

`func (o *TokenUser) SetPasswordExpiresAt(v string)`

SetPasswordExpiresAt sets PasswordExpiresAt field to given value.

### HasPasswordExpiresAt

`func (o *TokenUser) HasPasswordExpiresAt() bool`

HasPasswordExpiresAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


