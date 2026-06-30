# AzPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Services** | **map[string][]string** |  | 
**Priority** | **int32** |  | 
**CreatedBy** | **string** |  | 
**UpdatedBy** | **string** |  | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAzPolicy

`func NewAzPolicy(id string, services map[string][]string, priority int32, createdBy string, updatedBy string, ) *AzPolicy`

NewAzPolicy instantiates a new AzPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzPolicyWithDefaults

`func NewAzPolicyWithDefaults() *AzPolicy`

NewAzPolicyWithDefaults instantiates a new AzPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AzPolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AzPolicy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AzPolicy) SetId(v string)`

SetId sets Id field to given value.


### GetServices

`func (o *AzPolicy) GetServices() map[string][]string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *AzPolicy) GetServicesOk() (*map[string][]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *AzPolicy) SetServices(v map[string][]string)`

SetServices sets Services field to given value.


### GetPriority

`func (o *AzPolicy) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *AzPolicy) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *AzPolicy) SetPriority(v int32)`

SetPriority sets Priority field to given value.


### GetCreatedBy

`func (o *AzPolicy) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AzPolicy) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AzPolicy) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetUpdatedBy

`func (o *AzPolicy) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *AzPolicy) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *AzPolicy) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.


### GetCreatedAt

`func (o *AzPolicy) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AzPolicy) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AzPolicy) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AzPolicy) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AzPolicy) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AzPolicy) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AzPolicy) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AzPolicy) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


