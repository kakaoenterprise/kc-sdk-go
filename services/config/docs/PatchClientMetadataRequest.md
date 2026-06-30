# PatchClientMetadataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Upstreams** | Pointer to **map[string]string** |  | [optional] 
**References** | Pointer to **map[string]string** |  | [optional] 
**Cors** | Pointer to [**Cors**](Cors.md) |  | [optional] 

## Methods

### NewPatchClientMetadataRequest

`func NewPatchClientMetadataRequest() *PatchClientMetadataRequest`

NewPatchClientMetadataRequest instantiates a new PatchClientMetadataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchClientMetadataRequestWithDefaults

`func NewPatchClientMetadataRequestWithDefaults() *PatchClientMetadataRequest`

NewPatchClientMetadataRequestWithDefaults instantiates a new PatchClientMetadataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpstreams

`func (o *PatchClientMetadataRequest) GetUpstreams() map[string]string`

GetUpstreams returns the Upstreams field if non-nil, zero value otherwise.

### GetUpstreamsOk

`func (o *PatchClientMetadataRequest) GetUpstreamsOk() (*map[string]string, bool)`

GetUpstreamsOk returns a tuple with the Upstreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreams

`func (o *PatchClientMetadataRequest) SetUpstreams(v map[string]string)`

SetUpstreams sets Upstreams field to given value.

### HasUpstreams

`func (o *PatchClientMetadataRequest) HasUpstreams() bool`

HasUpstreams returns a boolean if a field has been set.

### GetReferences

`func (o *PatchClientMetadataRequest) GetReferences() map[string]string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *PatchClientMetadataRequest) GetReferencesOk() (*map[string]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *PatchClientMetadataRequest) SetReferences(v map[string]string)`

SetReferences sets References field to given value.

### HasReferences

`func (o *PatchClientMetadataRequest) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetCors

`func (o *PatchClientMetadataRequest) GetCors() Cors`

GetCors returns the Cors field if non-nil, zero value otherwise.

### GetCorsOk

`func (o *PatchClientMetadataRequest) GetCorsOk() (*Cors, bool)`

GetCorsOk returns a tuple with the Cors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCors

`func (o *PatchClientMetadataRequest) SetCors(v Cors)`

SetCors sets Cors field to given value.

### HasCors

`func (o *PatchClientMetadataRequest) HasCors() bool`

HasCors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


