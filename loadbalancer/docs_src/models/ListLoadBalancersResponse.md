# ListLoadBalancersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoadBalancers** | [**[]LoadBalancer**](LoadBalancer.md) | 로드 밸런서 목록 | 
**Pagination** | [**Pagination**](Pagination.md) | 페이지네이션 정보 | 

## Methods

### NewListLoadBalancersResponse

`func NewListLoadBalancersResponse(loadBalancers []LoadBalancer, pagination Pagination, ) *ListLoadBalancersResponse`

NewListLoadBalancersResponse instantiates a new ListLoadBalancersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLoadBalancersResponseWithDefaults

`func NewListLoadBalancersResponseWithDefaults() *ListLoadBalancersResponse`

NewListLoadBalancersResponseWithDefaults instantiates a new ListLoadBalancersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoadBalancers

`func (o *ListLoadBalancersResponse) GetLoadBalancers() []LoadBalancer`

GetLoadBalancers returns the LoadBalancers field if non-nil, zero value otherwise.

### GetLoadBalancersOk

`func (o *ListLoadBalancersResponse) GetLoadBalancersOk() (*[]LoadBalancer, bool)`

GetLoadBalancersOk returns a tuple with the LoadBalancers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancers

`func (o *ListLoadBalancersResponse) SetLoadBalancers(v []LoadBalancer)`

SetLoadBalancers sets LoadBalancers field to given value.


### GetPagination

`func (o *ListLoadBalancersResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListLoadBalancersResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListLoadBalancersResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


