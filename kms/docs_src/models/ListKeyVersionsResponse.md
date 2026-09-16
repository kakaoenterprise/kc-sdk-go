# ListKeyVersionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Versions** | [**[]KeyVersion**](KeyVersion.md) | 키 버전 목록 | 
**Pagination** | [**Pagination**](Pagination.md) | 목록 조회의 페이지네이션 정보 | 

## Methods

### NewListKeyVersionsResponse

`func NewListKeyVersionsResponse(versions []KeyVersion, pagination Pagination, ) *ListKeyVersionsResponse`

NewListKeyVersionsResponse instantiates a new ListKeyVersionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListKeyVersionsResponseWithDefaults

`func NewListKeyVersionsResponseWithDefaults() *ListKeyVersionsResponse`

NewListKeyVersionsResponseWithDefaults instantiates a new ListKeyVersionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersions

`func (o *ListKeyVersionsResponse) GetVersions() []KeyVersion`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *ListKeyVersionsResponse) GetVersionsOk() (*[]KeyVersion, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *ListKeyVersionsResponse) SetVersions(v []KeyVersion)`

SetVersions sets Versions field to given value.


### GetPagination

`func (o *ListKeyVersionsResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListKeyVersionsResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListKeyVersionsResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


