# ListTargetGroupsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetGroups** | [**[]LoadBalancerPool**](LoadBalancerPool.md) | 대상 그룹 목록 | 
**Pagination** | [**Pagination**](Pagination.md) | 페이지네이션 정보 | 

## Methods

### NewListTargetGroupsResponse

`func NewListTargetGroupsResponse(targetGroups []LoadBalancerPool, pagination Pagination, ) *ListTargetGroupsResponse`

NewListTargetGroupsResponse instantiates a new ListTargetGroupsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTargetGroupsResponseWithDefaults

`func NewListTargetGroupsResponseWithDefaults() *ListTargetGroupsResponse`

NewListTargetGroupsResponseWithDefaults instantiates a new ListTargetGroupsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetGroups

`func (o *ListTargetGroupsResponse) GetTargetGroups() []LoadBalancerPool`

GetTargetGroups returns the TargetGroups field if non-nil, zero value otherwise.

### GetTargetGroupsOk

`func (o *ListTargetGroupsResponse) GetTargetGroupsOk() (*[]LoadBalancerPool, bool)`

GetTargetGroupsOk returns a tuple with the TargetGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroups

`func (o *ListTargetGroupsResponse) SetTargetGroups(v []LoadBalancerPool)`

SetTargetGroups sets TargetGroups field to given value.


### GetPagination

`func (o *ListTargetGroupsResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListTargetGroupsResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListTargetGroupsResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


