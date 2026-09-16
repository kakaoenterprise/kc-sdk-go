# ListImageSharedProjectsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Members** | [**[]ListImageSharedProjectsMember**](ListImageSharedProjectsMember.md) | 이미지 공유 멤버 목록 | 
**Pagination** | [**Pagination**](Pagination.md) | 페이지네이션 정보 | 

## Methods

### NewListImageSharedProjectsResponse

`func NewListImageSharedProjectsResponse(members []ListImageSharedProjectsMember, pagination Pagination, ) *ListImageSharedProjectsResponse`

NewListImageSharedProjectsResponse instantiates a new ListImageSharedProjectsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListImageSharedProjectsResponseWithDefaults

`func NewListImageSharedProjectsResponseWithDefaults() *ListImageSharedProjectsResponse`

NewListImageSharedProjectsResponseWithDefaults instantiates a new ListImageSharedProjectsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMembers

`func (o *ListImageSharedProjectsResponse) GetMembers() []ListImageSharedProjectsMember`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *ListImageSharedProjectsResponse) GetMembersOk() (*[]ListImageSharedProjectsMember, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *ListImageSharedProjectsResponse) SetMembers(v []ListImageSharedProjectsMember)`

SetMembers sets Members field to given value.


### GetPagination

`func (o *ListImageSharedProjectsResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListImageSharedProjectsResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListImageSharedProjectsResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


