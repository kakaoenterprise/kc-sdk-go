# L7PolicyListModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**L7Policies** | Pointer to [**[]BnsLoadBalancerV1ApiListL7PoliciesModelL7PolicyModel**](BnsLoadBalancerV1ApiListL7PoliciesModelL7PolicyModel.md) |  | [optional] 
**Pagination** | [**PaginationModel**](PaginationModel.md) |  | 

## Methods

### NewL7PolicyListModel

`func NewL7PolicyListModel(pagination PaginationModel, ) *L7PolicyListModel`

NewL7PolicyListModel instantiates a new L7PolicyListModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewL7PolicyListModelWithDefaults

`func NewL7PolicyListModelWithDefaults() *L7PolicyListModel`

NewL7PolicyListModelWithDefaults instantiates a new L7PolicyListModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetL7Policies

`func (o *L7PolicyListModel) GetL7Policies() []BnsLoadBalancerV1ApiListL7PoliciesModelL7PolicyModel`

GetL7Policies returns the L7Policies field if non-nil, zero value otherwise.

### GetL7PoliciesOk

`func (o *L7PolicyListModel) GetL7PoliciesOk() (*[]BnsLoadBalancerV1ApiListL7PoliciesModelL7PolicyModel, bool)`

GetL7PoliciesOk returns a tuple with the L7Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL7Policies

`func (o *L7PolicyListModel) SetL7Policies(v []BnsLoadBalancerV1ApiListL7PoliciesModelL7PolicyModel)`

SetL7Policies sets L7Policies field to given value.

### HasL7Policies

`func (o *L7PolicyListModel) HasL7Policies() bool`

HasL7Policies returns a boolean if a field has been set.

### GetPagination

`func (o *L7PolicyListModel) GetPagination() PaginationModel`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *L7PolicyListModel) GetPaginationOk() (*PaginationModel, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *L7PolicyListModel) SetPagination(v PaginationModel)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


