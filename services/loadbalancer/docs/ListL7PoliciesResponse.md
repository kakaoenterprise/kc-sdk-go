# ListL7PoliciesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**L7Policies** | [**[]L7Policy**](L7Policy.md) | L7 정책 목록 | 
**Pagination** | [**Pagination**](Pagination.md) | 페이지네이션 정보 | 

## Methods

### NewListL7PoliciesResponse

`func NewListL7PoliciesResponse(l7Policies []L7Policy, pagination Pagination, ) *ListL7PoliciesResponse`

NewListL7PoliciesResponse instantiates a new ListL7PoliciesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListL7PoliciesResponseWithDefaults

`func NewListL7PoliciesResponseWithDefaults() *ListL7PoliciesResponse`

NewListL7PoliciesResponseWithDefaults instantiates a new ListL7PoliciesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetL7Policies

`func (o *ListL7PoliciesResponse) GetL7Policies() []L7Policy`

GetL7Policies returns the L7Policies field if non-nil, zero value otherwise.

### GetL7PoliciesOk

`func (o *ListL7PoliciesResponse) GetL7PoliciesOk() (*[]L7Policy, bool)`

GetL7PoliciesOk returns a tuple with the L7Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL7Policies

`func (o *ListL7PoliciesResponse) SetL7Policies(v []L7Policy)`

SetL7Policies sets L7Policies field to given value.


### GetPagination

`func (o *ListL7PoliciesResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListL7PoliciesResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListL7PoliciesResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


