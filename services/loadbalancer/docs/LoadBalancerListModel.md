# LoadBalancerListModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoadBalancers** | Pointer to [**[]BnsLoadBalancerV1ApiListLoadBalancersModelLoadBalancerModel**](BnsLoadBalancerV1ApiListLoadBalancersModelLoadBalancerModel.md) |  | [optional] 
**Pagination** | [**PaginationModel**](PaginationModel.md) |  | 

## Methods

### NewLoadBalancerListModel

`func NewLoadBalancerListModel(pagination PaginationModel, ) *LoadBalancerListModel`

NewLoadBalancerListModel instantiates a new LoadBalancerListModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadBalancerListModelWithDefaults

`func NewLoadBalancerListModelWithDefaults() *LoadBalancerListModel`

NewLoadBalancerListModelWithDefaults instantiates a new LoadBalancerListModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoadBalancers

`func (o *LoadBalancerListModel) GetLoadBalancers() []BnsLoadBalancerV1ApiListLoadBalancersModelLoadBalancerModel`

GetLoadBalancers returns the LoadBalancers field if non-nil, zero value otherwise.

### GetLoadBalancersOk

`func (o *LoadBalancerListModel) GetLoadBalancersOk() (*[]BnsLoadBalancerV1ApiListLoadBalancersModelLoadBalancerModel, bool)`

GetLoadBalancersOk returns a tuple with the LoadBalancers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancers

`func (o *LoadBalancerListModel) SetLoadBalancers(v []BnsLoadBalancerV1ApiListLoadBalancersModelLoadBalancerModel)`

SetLoadBalancers sets LoadBalancers field to given value.

### HasLoadBalancers

`func (o *LoadBalancerListModel) HasLoadBalancers() bool`

HasLoadBalancers returns a boolean if a field has been set.

### GetPagination

`func (o *LoadBalancerListModel) GetPagination() PaginationModel`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *LoadBalancerListModel) GetPaginationOk() (*PaginationModel, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *LoadBalancerListModel) SetPagination(v PaginationModel)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


