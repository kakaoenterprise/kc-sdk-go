# ListTargetsInTargetGroupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Members** | [**[]PoolMember**](PoolMember.md) | 대상 그룹 멤버 목록 | 
**Pagination** | [**Pagination**](Pagination.md) | 페이지네이션 정보 | 

## Methods

### NewListTargetsInTargetGroupResponse

`func NewListTargetsInTargetGroupResponse(members []PoolMember, pagination Pagination, ) *ListTargetsInTargetGroupResponse`

NewListTargetsInTargetGroupResponse instantiates a new ListTargetsInTargetGroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTargetsInTargetGroupResponseWithDefaults

`func NewListTargetsInTargetGroupResponseWithDefaults() *ListTargetsInTargetGroupResponse`

NewListTargetsInTargetGroupResponseWithDefaults instantiates a new ListTargetsInTargetGroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMembers

`func (o *ListTargetsInTargetGroupResponse) GetMembers() []PoolMember`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *ListTargetsInTargetGroupResponse) GetMembersOk() (*[]PoolMember, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *ListTargetsInTargetGroupResponse) SetMembers(v []PoolMember)`

SetMembers sets Members field to given value.


### GetPagination

`func (o *ListTargetsInTargetGroupResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListTargetsInTargetGroupResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListTargetsInTargetGroupResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


