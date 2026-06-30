# ClientEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ServiceEndpoints** | **map[string]string** |  | 
**Priority** | **int32** |  | 
**CreatedBy** | **string** |  | 
**UpdatedBy** | **string** |  | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewClientEndpoint

`func NewClientEndpoint(id string, serviceEndpoints map[string]string, priority int32, createdBy string, updatedBy string, ) *ClientEndpoint`

NewClientEndpoint instantiates a new ClientEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientEndpointWithDefaults

`func NewClientEndpointWithDefaults() *ClientEndpoint`

NewClientEndpointWithDefaults instantiates a new ClientEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClientEndpoint) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClientEndpoint) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClientEndpoint) SetId(v string)`

SetId sets Id field to given value.


### GetServiceEndpoints

`func (o *ClientEndpoint) GetServiceEndpoints() map[string]string`

GetServiceEndpoints returns the ServiceEndpoints field if non-nil, zero value otherwise.

### GetServiceEndpointsOk

`func (o *ClientEndpoint) GetServiceEndpointsOk() (*map[string]string, bool)`

GetServiceEndpointsOk returns a tuple with the ServiceEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEndpoints

`func (o *ClientEndpoint) SetServiceEndpoints(v map[string]string)`

SetServiceEndpoints sets ServiceEndpoints field to given value.


### GetPriority

`func (o *ClientEndpoint) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ClientEndpoint) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ClientEndpoint) SetPriority(v int32)`

SetPriority sets Priority field to given value.


### GetCreatedBy

`func (o *ClientEndpoint) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ClientEndpoint) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ClientEndpoint) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetUpdatedBy

`func (o *ClientEndpoint) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *ClientEndpoint) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *ClientEndpoint) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.


### GetCreatedAt

`func (o *ClientEndpoint) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ClientEndpoint) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ClientEndpoint) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ClientEndpoint) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ClientEndpoint) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ClientEndpoint) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ClientEndpoint) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ClientEndpoint) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


