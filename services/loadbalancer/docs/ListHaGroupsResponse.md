# ListHaGroupsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BeyondLoadBalancers** | [**[]BeyondLoadBalancer**](BeyondLoadBalancer.md) | 고가용성 그룹 목록 | 
**Pagination** | [**Pagination**](Pagination.md) | 페이지네이션 정보 | 

## Methods

### NewListHaGroupsResponse

`func NewListHaGroupsResponse(beyondLoadBalancers []BeyondLoadBalancer, pagination Pagination, ) *ListHaGroupsResponse`

NewListHaGroupsResponse instantiates a new ListHaGroupsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListHaGroupsResponseWithDefaults

`func NewListHaGroupsResponseWithDefaults() *ListHaGroupsResponse`

NewListHaGroupsResponseWithDefaults instantiates a new ListHaGroupsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBeyondLoadBalancers

`func (o *ListHaGroupsResponse) GetBeyondLoadBalancers() []BeyondLoadBalancer`

GetBeyondLoadBalancers returns the BeyondLoadBalancers field if non-nil, zero value otherwise.

### GetBeyondLoadBalancersOk

`func (o *ListHaGroupsResponse) GetBeyondLoadBalancersOk() (*[]BeyondLoadBalancer, bool)`

GetBeyondLoadBalancersOk returns a tuple with the BeyondLoadBalancers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeyondLoadBalancers

`func (o *ListHaGroupsResponse) SetBeyondLoadBalancers(v []BeyondLoadBalancer)`

SetBeyondLoadBalancers sets BeyondLoadBalancers field to given value.


### GetPagination

`func (o *ListHaGroupsResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListHaGroupsResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListHaGroupsResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


