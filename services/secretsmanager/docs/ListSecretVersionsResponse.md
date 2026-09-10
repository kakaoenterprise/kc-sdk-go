# ListSecretVersionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Versions** | [**[]Version**](Version.md) | 시크릿 버전 목록 | 
**Pagination** | [**Pagination**](Pagination.md) | 목록 조회의 페이지네이션 정보 | 

## Methods

### NewListSecretVersionsResponse

`func NewListSecretVersionsResponse(versions []Version, pagination Pagination, ) *ListSecretVersionsResponse`

NewListSecretVersionsResponse instantiates a new ListSecretVersionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSecretVersionsResponseWithDefaults

`func NewListSecretVersionsResponseWithDefaults() *ListSecretVersionsResponse`

NewListSecretVersionsResponseWithDefaults instantiates a new ListSecretVersionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersions

`func (o *ListSecretVersionsResponse) GetVersions() []Version`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *ListSecretVersionsResponse) GetVersionsOk() (*[]Version, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *ListSecretVersionsResponse) SetVersions(v []Version)`

SetVersions sets Versions field to given value.


### GetPagination

`func (o *ListSecretVersionsResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListSecretVersionsResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListSecretVersionsResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


