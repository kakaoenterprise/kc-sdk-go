# SwaggerAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identity** | Pointer to [**SwaggerIdPwd**](SwaggerIdPwd.md) |  | [optional] 
**Scope** | Pointer to [**SwaggerScope**](SwaggerScope.md) |  | [optional] 

## Methods

### NewSwaggerAuth

`func NewSwaggerAuth() *SwaggerAuth`

NewSwaggerAuth instantiates a new SwaggerAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwaggerAuthWithDefaults

`func NewSwaggerAuthWithDefaults() *SwaggerAuth`

NewSwaggerAuthWithDefaults instantiates a new SwaggerAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentity

`func (o *SwaggerAuth) GetIdentity() SwaggerIdPwd`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *SwaggerAuth) GetIdentityOk() (*SwaggerIdPwd, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *SwaggerAuth) SetIdentity(v SwaggerIdPwd)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *SwaggerAuth) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetScope

`func (o *SwaggerAuth) GetScope() SwaggerScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *SwaggerAuth) GetScopeOk() (*SwaggerScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *SwaggerAuth) SetScope(v SwaggerScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *SwaggerAuth) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


