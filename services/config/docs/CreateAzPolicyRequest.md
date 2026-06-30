# CreateAzPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Services** | **map[string][]string** |  | 

## Methods

### NewCreateAzPolicyRequest

`func NewCreateAzPolicyRequest(id string, services map[string][]string, ) *CreateAzPolicyRequest`

NewCreateAzPolicyRequest instantiates a new CreateAzPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAzPolicyRequestWithDefaults

`func NewCreateAzPolicyRequestWithDefaults() *CreateAzPolicyRequest`

NewCreateAzPolicyRequestWithDefaults instantiates a new CreateAzPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateAzPolicyRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateAzPolicyRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateAzPolicyRequest) SetId(v string)`

SetId sets Id field to given value.


### GetServices

`func (o *CreateAzPolicyRequest) GetServices() map[string][]string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *CreateAzPolicyRequest) GetServicesOk() (*map[string][]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *CreateAzPolicyRequest) SetServices(v map[string][]string)`

SetServices sets Services field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


