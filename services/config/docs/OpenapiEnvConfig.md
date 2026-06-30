# OpenapiEnvConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Profile** | **string** |  | 
**Endpoints** | **[]map[string]string** |  | 
**References** | Pointer to [**References**](References.md) |  | [optional] 
**AzPolicies** | Pointer to [**AZPolicies**](AZPolicies.md) |  | [optional] 

## Methods

### NewOpenapiEnvConfig

`func NewOpenapiEnvConfig(profile string, endpoints []map[string]string, ) *OpenapiEnvConfig`

NewOpenapiEnvConfig instantiates a new OpenapiEnvConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenapiEnvConfigWithDefaults

`func NewOpenapiEnvConfigWithDefaults() *OpenapiEnvConfig`

NewOpenapiEnvConfigWithDefaults instantiates a new OpenapiEnvConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OpenapiEnvConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OpenapiEnvConfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OpenapiEnvConfig) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OpenapiEnvConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProfile

`func (o *OpenapiEnvConfig) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *OpenapiEnvConfig) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *OpenapiEnvConfig) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetEndpoints

`func (o *OpenapiEnvConfig) GetEndpoints() []map[string]string`

GetEndpoints returns the Endpoints field if non-nil, zero value otherwise.

### GetEndpointsOk

`func (o *OpenapiEnvConfig) GetEndpointsOk() (*[]map[string]string, bool)`

GetEndpointsOk returns a tuple with the Endpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoints

`func (o *OpenapiEnvConfig) SetEndpoints(v []map[string]string)`

SetEndpoints sets Endpoints field to given value.


### GetReferences

`func (o *OpenapiEnvConfig) GetReferences() References`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *OpenapiEnvConfig) GetReferencesOk() (*References, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *OpenapiEnvConfig) SetReferences(v References)`

SetReferences sets References field to given value.

### HasReferences

`func (o *OpenapiEnvConfig) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetAzPolicies

`func (o *OpenapiEnvConfig) GetAzPolicies() AZPolicies`

GetAzPolicies returns the AzPolicies field if non-nil, zero value otherwise.

### GetAzPoliciesOk

`func (o *OpenapiEnvConfig) GetAzPoliciesOk() (*AZPolicies, bool)`

GetAzPoliciesOk returns a tuple with the AzPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzPolicies

`func (o *OpenapiEnvConfig) SetAzPolicies(v AZPolicies)`

SetAzPolicies sets AzPolicies field to given value.

### HasAzPolicies

`func (o *OpenapiEnvConfig) HasAzPolicies() bool`

HasAzPolicies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


