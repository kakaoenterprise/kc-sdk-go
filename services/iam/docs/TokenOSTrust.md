# TokenOSTrust

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Impersonation** | Pointer to **bool** |  | [optional] 
**Links** | Pointer to [**TokenLinks**](TokenLinks.md) |  | [optional] 
**TrusteeUser** | Pointer to [**TokenTrustUser**](TokenTrustUser.md) |  | [optional] 
**TrustorUser** | Pointer to [**TokenTrustUser**](TokenTrustUser.md) |  | [optional] 

## Methods

### NewTokenOSTrust

`func NewTokenOSTrust() *TokenOSTrust`

NewTokenOSTrust instantiates a new TokenOSTrust object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenOSTrustWithDefaults

`func NewTokenOSTrustWithDefaults() *TokenOSTrust`

NewTokenOSTrustWithDefaults instantiates a new TokenOSTrust object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TokenOSTrust) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TokenOSTrust) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TokenOSTrust) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TokenOSTrust) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImpersonation

`func (o *TokenOSTrust) GetImpersonation() bool`

GetImpersonation returns the Impersonation field if non-nil, zero value otherwise.

### GetImpersonationOk

`func (o *TokenOSTrust) GetImpersonationOk() (*bool, bool)`

GetImpersonationOk returns a tuple with the Impersonation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpersonation

`func (o *TokenOSTrust) SetImpersonation(v bool)`

SetImpersonation sets Impersonation field to given value.

### HasImpersonation

`func (o *TokenOSTrust) HasImpersonation() bool`

HasImpersonation returns a boolean if a field has been set.

### GetLinks

`func (o *TokenOSTrust) GetLinks() TokenLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TokenOSTrust) GetLinksOk() (*TokenLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TokenOSTrust) SetLinks(v TokenLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TokenOSTrust) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetTrusteeUser

`func (o *TokenOSTrust) GetTrusteeUser() TokenTrustUser`

GetTrusteeUser returns the TrusteeUser field if non-nil, zero value otherwise.

### GetTrusteeUserOk

`func (o *TokenOSTrust) GetTrusteeUserOk() (*TokenTrustUser, bool)`

GetTrusteeUserOk returns a tuple with the TrusteeUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrusteeUser

`func (o *TokenOSTrust) SetTrusteeUser(v TokenTrustUser)`

SetTrusteeUser sets TrusteeUser field to given value.

### HasTrusteeUser

`func (o *TokenOSTrust) HasTrusteeUser() bool`

HasTrusteeUser returns a boolean if a field has been set.

### GetTrustorUser

`func (o *TokenOSTrust) GetTrustorUser() TokenTrustUser`

GetTrustorUser returns the TrustorUser field if non-nil, zero value otherwise.

### GetTrustorUserOk

`func (o *TokenOSTrust) GetTrustorUserOk() (*TokenTrustUser, bool)`

GetTrustorUserOk returns a tuple with the TrustorUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustorUser

`func (o *TokenOSTrust) SetTrustorUser(v TokenTrustUser)`

SetTrustorUser sets TrustorUser field to given value.

### HasTrustorUser

`func (o *TokenOSTrust) HasTrustorUser() bool`

HasTrustorUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


