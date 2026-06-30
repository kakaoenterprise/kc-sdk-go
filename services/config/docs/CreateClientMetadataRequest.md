# CreateClientMetadataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Upstreams** | **map[string]string** |  | 
**References** | **map[string]string** |  | 
**Cors** | [**Cors**](Cors.md) |  | 

## Methods

### NewCreateClientMetadataRequest

`func NewCreateClientMetadataRequest(id string, upstreams map[string]string, references map[string]string, cors Cors, ) *CreateClientMetadataRequest`

NewCreateClientMetadataRequest instantiates a new CreateClientMetadataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateClientMetadataRequestWithDefaults

`func NewCreateClientMetadataRequestWithDefaults() *CreateClientMetadataRequest`

NewCreateClientMetadataRequestWithDefaults instantiates a new CreateClientMetadataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateClientMetadataRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateClientMetadataRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateClientMetadataRequest) SetId(v string)`

SetId sets Id field to given value.


### GetUpstreams

`func (o *CreateClientMetadataRequest) GetUpstreams() map[string]string`

GetUpstreams returns the Upstreams field if non-nil, zero value otherwise.

### GetUpstreamsOk

`func (o *CreateClientMetadataRequest) GetUpstreamsOk() (*map[string]string, bool)`

GetUpstreamsOk returns a tuple with the Upstreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreams

`func (o *CreateClientMetadataRequest) SetUpstreams(v map[string]string)`

SetUpstreams sets Upstreams field to given value.


### GetReferences

`func (o *CreateClientMetadataRequest) GetReferences() map[string]string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *CreateClientMetadataRequest) GetReferencesOk() (*map[string]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *CreateClientMetadataRequest) SetReferences(v map[string]string)`

SetReferences sets References field to given value.


### GetCors

`func (o *CreateClientMetadataRequest) GetCors() Cors`

GetCors returns the Cors field if non-nil, zero value otherwise.

### GetCorsOk

`func (o *CreateClientMetadataRequest) GetCorsOk() (*Cors, bool)`

GetCorsOk returns a tuple with the Cors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCors

`func (o *CreateClientMetadataRequest) SetCors(v Cors)`

SetCors sets Cors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


