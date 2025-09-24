# TokenCatalog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endpoints** | Pointer to [**[]TokenEndpoint**](TokenEndpoint.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewTokenCatalog

`func NewTokenCatalog() *TokenCatalog`

NewTokenCatalog instantiates a new TokenCatalog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenCatalogWithDefaults

`func NewTokenCatalogWithDefaults() *TokenCatalog`

NewTokenCatalogWithDefaults instantiates a new TokenCatalog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpoints

`func (o *TokenCatalog) GetEndpoints() []TokenEndpoint`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *TokenCatalog) GetEndpointsOk() (*[]TokenEndpoint, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *TokenCatalog) SetEndpoints(v []TokenEndpoint)`

SetEndpoints sets Endpoints field to given value.

### HasEndpoints

`func (o *TokenCatalog) HasEndpoints() bool`

HasEndpoints returns a boolean if a field has been set.

### GetId

`func (o *TokenCatalog) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TokenCatalog) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TokenCatalog) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TokenCatalog) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TokenCatalog) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TokenCatalog) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TokenCatalog) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TokenCatalog) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *TokenCatalog) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TokenCatalog) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TokenCatalog) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TokenCatalog) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


