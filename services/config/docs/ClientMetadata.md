# ClientMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Upstreams** | **map[string]string** |  | 
**References** | **map[string]string** |  | 
**Cors** | [**Cors**](Cors.md) |  | 
**CreatedBy** | **string** |  | 
**UpdatedBy** | **string** |  | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewClientMetadata

`func NewClientMetadata(id string, upstreams map[string]string, references map[string]string, cors Cors, createdBy string, updatedBy string, ) *ClientMetadata`

NewClientMetadata instantiates a new ClientMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientMetadataWithDefaults

`func NewClientMetadataWithDefaults() *ClientMetadata`

NewClientMetadataWithDefaults instantiates a new ClientMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClientMetadata) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClientMetadata) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClientMetadata) SetId(v string)`

SetId sets Id field to given value.


### GetUpstreams

`func (o *ClientMetadata) GetUpstreams() map[string]string`

GetUpstreams returns the Upstreams field if non-nil, zero value otherwise.

### GetUpstreamsOk

`func (o *ClientMetadata) GetUpstreamsOk() (*map[string]string, bool)`

GetUpstreamsOk returns a tuple with the Upstreams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreams

`func (o *ClientMetadata) SetUpstreams(v map[string]string)`

SetUpstreams sets Upstreams field to given value.


### GetReferences

`func (o *ClientMetadata) GetReferences() map[string]string`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *ClientMetadata) GetReferencesOk() (*map[string]string, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *ClientMetadata) SetReferences(v map[string]string)`

SetReferences sets References field to given value.


### GetCors

`func (o *ClientMetadata) GetCors() Cors`

GetCors returns the Cors field if non-nil, zero value otherwise.

### GetCorsOk

`func (o *ClientMetadata) GetCorsOk() (*Cors, bool)`

GetCorsOk returns a tuple with the Cors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCors

`func (o *ClientMetadata) SetCors(v Cors)`

SetCors sets Cors field to given value.


### GetCreatedBy

`func (o *ClientMetadata) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ClientMetadata) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ClientMetadata) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetUpdatedBy

`func (o *ClientMetadata) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *ClientMetadata) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *ClientMetadata) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.


### GetCreatedAt

`func (o *ClientMetadata) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ClientMetadata) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ClientMetadata) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ClientMetadata) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ClientMetadata) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ClientMetadata) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ClientMetadata) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ClientMetadata) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


