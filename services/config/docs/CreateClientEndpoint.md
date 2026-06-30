# CreateClientEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ServiceEndpoints** | **map[string]string** |  | 

## Methods

### NewCreateClientEndpoint

`func NewCreateClientEndpoint(id string, serviceEndpoints map[string]string, ) *CreateClientEndpoint`

NewCreateClientEndpoint instantiates a new CreateClientEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateClientEndpointWithDefaults

`func NewCreateClientEndpointWithDefaults() *CreateClientEndpoint`

NewCreateClientEndpointWithDefaults instantiates a new CreateClientEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateClientEndpoint) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateClientEndpoint) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateClientEndpoint) SetId(v string)`

SetId sets Id field to given value.


### GetServiceEndpoints

`func (o *CreateClientEndpoint) GetServiceEndpoints() map[string]string`

GetServiceEndpoints returns the ServiceEndpoints field if non-nil, zero value otherwise.

### GetServiceEndpointsOk

`func (o *CreateClientEndpoint) GetServiceEndpointsOk() (*map[string]string, bool)`

GetServiceEndpointsOk returns a tuple with the ServiceEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEndpoints

`func (o *CreateClientEndpoint) SetServiceEndpoints(v map[string]string)`

SetServiceEndpoints sets ServiceEndpoints field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


