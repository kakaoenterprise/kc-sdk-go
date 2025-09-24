# Token

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OSTRUSTTrust** | Pointer to [**TokenOSTrust**](TokenOSTrust.md) |  | [optional] 
**ApplicationCredential** | Pointer to [**TokenApplicationCredential**](TokenApplicationCredential.md) |  | [optional] 
**AuditIds** | Pointer to **[]string** |  | [optional] 
**Catalog** | Pointer to [**[]TokenCatalog**](TokenCatalog.md) |  | [optional] 
**Domain** | Pointer to [**TokenDomain**](TokenDomain.md) |  | [optional] 
**ExpiresAt** | Pointer to **string** |  | [optional] 
**IsDomain** | Pointer to **bool** |  | [optional] 
**IssuedAt** | Pointer to **string** |  | [optional] 
**Methods** | Pointer to **[]string** |  | [optional] 
**Project** | Pointer to [**TokenProject**](TokenProject.md) |  | [optional] 
**Roles** | Pointer to [**[]TokenRole**](TokenRole.md) |  | [optional] 
**System** | Pointer to [**TokenSystem**](TokenSystem.md) |  | [optional] 
**User** | Pointer to [**TokenUser**](TokenUser.md) |  | [optional] 

## Methods

### NewToken

`func NewToken() *Token`

NewToken instantiates a new Token object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenWithDefaults

`func NewTokenWithDefaults() *Token`

NewTokenWithDefaults instantiates a new Token object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOSTRUSTTrust

`func (o *Token) GetOSTRUSTTrust() TokenOSTrust`

GetOSTRUSTTrust returns the OSTRUSTTrust field if non-nil, zero value otherwise.

### GetOSTRUSTTrustOk

`func (o *Token) GetOSTRUSTTrustOk() (*TokenOSTrust, bool)`

GetOSTRUSTTrustOk returns a tuple with the OSTRUSTTrust field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOSTRUSTTrust

`func (o *Token) SetOSTRUSTTrust(v TokenOSTrust)`

SetOSTRUSTTrust sets OSTRUSTTrust field to given value.

### HasOSTRUSTTrust

`func (o *Token) HasOSTRUSTTrust() bool`

HasOSTRUSTTrust returns a boolean if a field has been set.

### GetApplicationCredential

`func (o *Token) GetApplicationCredential() TokenApplicationCredential`

GetApplicationCredential returns the ApplicationCredential field if non-nil, zero value otherwise.

### GetApplicationCredentialOk

`func (o *Token) GetApplicationCredentialOk() (*TokenApplicationCredential, bool)`

GetApplicationCredentialOk returns a tuple with the ApplicationCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationCredential

`func (o *Token) SetApplicationCredential(v TokenApplicationCredential)`

SetApplicationCredential sets ApplicationCredential field to given value.

### HasApplicationCredential

`func (o *Token) HasApplicationCredential() bool`

HasApplicationCredential returns a boolean if a field has been set.

### GetAuditIds

`func (o *Token) GetAuditIds() []string`

GetAuditIds returns the AuditIds field if non-nil, zero value otherwise.

### GetAuditIdsOk

`func (o *Token) GetAuditIdsOk() (*[]string, bool)`

GetAuditIdsOk returns a tuple with the AuditIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditIds

`func (o *Token) SetAuditIds(v []string)`

SetAuditIds sets AuditIds field to given value.

### HasAuditIds

`func (o *Token) HasAuditIds() bool`

HasAuditIds returns a boolean if a field has been set.

### GetCatalog

`func (o *Token) GetCatalog() []TokenCatalog`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *Token) GetCatalogOk() (*[]TokenCatalog, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *Token) SetCatalog(v []TokenCatalog)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *Token) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.

### GetDomain

`func (o *Token) GetDomain() TokenDomain`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *Token) GetDomainOk() (*TokenDomain, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *Token) SetDomain(v TokenDomain)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *Token) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetExpiresAt

`func (o *Token) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *Token) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *Token) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *Token) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetIsDomain

`func (o *Token) GetIsDomain() bool`

GetIsDomain returns the IsDomain field if non-nil, zero value otherwise.

### GetIsDomainOk

`func (o *Token) GetIsDomainOk() (*bool, bool)`

GetIsDomainOk returns a tuple with the IsDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDomain

`func (o *Token) SetIsDomain(v bool)`

SetIsDomain sets IsDomain field to given value.

### HasIsDomain

`func (o *Token) HasIsDomain() bool`

HasIsDomain returns a boolean if a field has been set.

### GetIssuedAt

`func (o *Token) GetIssuedAt() string`

GetIssuedAt returns the IssuedAt field if non-nil, zero value otherwise.

### GetIssuedAtOk

`func (o *Token) GetIssuedAtOk() (*string, bool)`

GetIssuedAtOk returns a tuple with the IssuedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedAt

`func (o *Token) SetIssuedAt(v string)`

SetIssuedAt sets IssuedAt field to given value.

### HasIssuedAt

`func (o *Token) HasIssuedAt() bool`

HasIssuedAt returns a boolean if a field has been set.

### GetMethods

`func (o *Token) GetMethods() []string`

GetMethods returns the Methods field if non-nil, zero value otherwise.

### GetMethodsOk

`func (o *Token) GetMethodsOk() (*[]string, bool)`

GetMethodsOk returns a tuple with the Methods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethods

`func (o *Token) SetMethods(v []string)`

SetMethods sets Methods field to given value.

### HasMethods

`func (o *Token) HasMethods() bool`

HasMethods returns a boolean if a field has been set.

### GetProject

`func (o *Token) GetProject() TokenProject`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *Token) GetProjectOk() (*TokenProject, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *Token) SetProject(v TokenProject)`

SetProject sets Project field to given value.

### HasProject

`func (o *Token) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetRoles

`func (o *Token) GetRoles() []TokenRole`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *Token) GetRolesOk() (*[]TokenRole, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *Token) SetRoles(v []TokenRole)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *Token) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetSystem

`func (o *Token) GetSystem() TokenSystem`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *Token) GetSystemOk() (*TokenSystem, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *Token) SetSystem(v TokenSystem)`

SetSystem sets System field to given value.

### HasSystem

`func (o *Token) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetUser

`func (o *Token) GetUser() TokenUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Token) GetUserOk() (*TokenUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Token) SetUser(v TokenUser)`

SetUser sets User field to given value.

### HasUser

`func (o *Token) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


