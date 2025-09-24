# TokenApplicationCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessRules** | Pointer to [**[]AccessRule**](AccessRule.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ExpiresAt** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Links** | Pointer to [**TokenLinks**](TokenLinks.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**Restricted** | Pointer to **bool** |  | [optional] 
**Roles** | Pointer to [**[]TokenRole**](TokenRole.md) |  | [optional] 
**Secret** | Pointer to **string** |  | [optional] 
**Unrestricted** | Pointer to **bool** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 

## Methods

### NewTokenApplicationCredential

`func NewTokenApplicationCredential() *TokenApplicationCredential`

NewTokenApplicationCredential instantiates a new TokenApplicationCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenApplicationCredentialWithDefaults

`func NewTokenApplicationCredentialWithDefaults() *TokenApplicationCredential`

NewTokenApplicationCredentialWithDefaults instantiates a new TokenApplicationCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessRules

`func (o *TokenApplicationCredential) GetAccessRules() []AccessRule`

GetAccessRules returns the AccessRules field if non-nil, zero value otherwise.

### GetAccessRulesOk

`func (o *TokenApplicationCredential) GetAccessRulesOk() (*[]AccessRule, bool)`

GetAccessRulesOk returns a tuple with the AccessRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessRules

`func (o *TokenApplicationCredential) SetAccessRules(v []AccessRule)`

SetAccessRules sets AccessRules field to given value.

### HasAccessRules

`func (o *TokenApplicationCredential) HasAccessRules() bool`

HasAccessRules returns a boolean if a field has been set.

### GetDescription

`func (o *TokenApplicationCredential) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TokenApplicationCredential) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TokenApplicationCredential) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TokenApplicationCredential) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpiresAt

`func (o *TokenApplicationCredential) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *TokenApplicationCredential) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *TokenApplicationCredential) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *TokenApplicationCredential) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetId

`func (o *TokenApplicationCredential) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TokenApplicationCredential) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TokenApplicationCredential) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TokenApplicationCredential) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLinks

`func (o *TokenApplicationCredential) GetLinks() TokenLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TokenApplicationCredential) GetLinksOk() (*TokenLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TokenApplicationCredential) SetLinks(v TokenLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TokenApplicationCredential) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetName

`func (o *TokenApplicationCredential) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TokenApplicationCredential) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TokenApplicationCredential) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TokenApplicationCredential) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProjectId

`func (o *TokenApplicationCredential) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *TokenApplicationCredential) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *TokenApplicationCredential) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *TokenApplicationCredential) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetRestricted

`func (o *TokenApplicationCredential) GetRestricted() bool`

GetRestricted returns the Restricted field if non-nil, zero value otherwise.

### GetRestrictedOk

`func (o *TokenApplicationCredential) GetRestrictedOk() (*bool, bool)`

GetRestrictedOk returns a tuple with the Restricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestricted

`func (o *TokenApplicationCredential) SetRestricted(v bool)`

SetRestricted sets Restricted field to given value.

### HasRestricted

`func (o *TokenApplicationCredential) HasRestricted() bool`

HasRestricted returns a boolean if a field has been set.

### GetRoles

`func (o *TokenApplicationCredential) GetRoles() []TokenRole`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *TokenApplicationCredential) GetRolesOk() (*[]TokenRole, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *TokenApplicationCredential) SetRoles(v []TokenRole)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *TokenApplicationCredential) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetSecret

`func (o *TokenApplicationCredential) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *TokenApplicationCredential) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *TokenApplicationCredential) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *TokenApplicationCredential) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetUnrestricted

`func (o *TokenApplicationCredential) GetUnrestricted() bool`

GetUnrestricted returns the Unrestricted field if non-nil, zero value otherwise.

### GetUnrestrictedOk

`func (o *TokenApplicationCredential) GetUnrestrictedOk() (*bool, bool)`

GetUnrestrictedOk returns a tuple with the Unrestricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnrestricted

`func (o *TokenApplicationCredential) SetUnrestricted(v bool)`

SetUnrestricted sets Unrestricted field to given value.

### HasUnrestricted

`func (o *TokenApplicationCredential) HasUnrestricted() bool`

HasUnrestricted returns a boolean if a field has been set.

### GetUserId

`func (o *TokenApplicationCredential) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *TokenApplicationCredential) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *TokenApplicationCredential) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *TokenApplicationCredential) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


