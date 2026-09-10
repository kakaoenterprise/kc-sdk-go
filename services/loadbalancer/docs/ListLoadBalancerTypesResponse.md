# ListLoadBalancerTypesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Flavors** | Pointer to [**[]Flavor**](Flavor.md) |  | [optional] 

## Methods

### NewListLoadBalancerTypesResponse

`func NewListLoadBalancerTypesResponse() *ListLoadBalancerTypesResponse`

NewListLoadBalancerTypesResponse instantiates a new ListLoadBalancerTypesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLoadBalancerTypesResponseWithDefaults

`func NewListLoadBalancerTypesResponseWithDefaults() *ListLoadBalancerTypesResponse`

NewListLoadBalancerTypesResponseWithDefaults instantiates a new ListLoadBalancerTypesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlavors

`func (o *ListLoadBalancerTypesResponse) GetFlavors() []Flavor`

GetFlavors returns the Flavors field if non-nil, zero value otherwise.

### GetFlavorsOk

`func (o *ListLoadBalancerTypesResponse) GetFlavorsOk() (*[]Flavor, bool)`

GetFlavorsOk returns a tuple with the Flavors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavors

`func (o *ListLoadBalancerTypesResponse) SetFlavors(v []Flavor)`

SetFlavors sets Flavors field to given value.

### HasFlavors

`func (o *ListLoadBalancerTypesResponse) HasFlavors() bool`

HasFlavors returns a boolean if a field has been set.

### SetFlavorsNil

`func (o *ListLoadBalancerTypesResponse) SetFlavorsNil(b bool)`

 SetFlavorsNil sets the value for Flavors to be an explicit nil

### UnsetFlavors
`func (o *ListLoadBalancerTypesResponse) UnsetFlavors()`

UnsetFlavors ensures that no value is present for Flavors, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


