# BeyondLoadBalancerListModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BeyondLoadBalancers** | Pointer to [**[]BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel**](BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel.md) |  | [optional] 
**Pagination** | [**PaginationModel**](PaginationModel.md) |  | 

## Methods

### NewBeyondLoadBalancerListModel

`func NewBeyondLoadBalancerListModel(pagination PaginationModel, ) *BeyondLoadBalancerListModel`

NewBeyondLoadBalancerListModel instantiates a new BeyondLoadBalancerListModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBeyondLoadBalancerListModelWithDefaults

`func NewBeyondLoadBalancerListModelWithDefaults() *BeyondLoadBalancerListModel`

NewBeyondLoadBalancerListModelWithDefaults instantiates a new BeyondLoadBalancerListModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBeyondLoadBalancers

`func (o *BeyondLoadBalancerListModel) GetBeyondLoadBalancers() []BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel`

GetBeyondLoadBalancers returns the BeyondLoadBalancers field if non-nil, zero value otherwise.

### GetBeyondLoadBalancersOk

`func (o *BeyondLoadBalancerListModel) GetBeyondLoadBalancersOk() (*[]BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel, bool)`

GetBeyondLoadBalancersOk returns a tuple with the BeyondLoadBalancers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeyondLoadBalancers

`func (o *BeyondLoadBalancerListModel) SetBeyondLoadBalancers(v []BnsLoadBalancerV1ApiListHaGroupsModelBeyondLoadBalancerModel)`

SetBeyondLoadBalancers sets BeyondLoadBalancers field to given value.

### HasBeyondLoadBalancers

`func (o *BeyondLoadBalancerListModel) HasBeyondLoadBalancers() bool`

HasBeyondLoadBalancers returns a boolean if a field has been set.

### GetPagination

`func (o *BeyondLoadBalancerListModel) GetPagination() PaginationModel`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *BeyondLoadBalancerListModel) GetPaginationOk() (*PaginationModel, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *BeyondLoadBalancerListModel) SetPagination(v PaginationModel)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


