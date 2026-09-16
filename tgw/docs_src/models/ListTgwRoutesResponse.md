# ListTgwRoutesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Routes** | [**[]Route**](Route.md) | 조회된 Route 목록 | 
**Pagination** | [**Pagination**](Pagination.md) | 페이지네이션 메타데이터 | 

## Methods

### NewListTgwRoutesResponse

`func NewListTgwRoutesResponse(routes []Route, pagination Pagination, ) *ListTgwRoutesResponse`

NewListTgwRoutesResponse instantiates a new ListTgwRoutesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTgwRoutesResponseWithDefaults

`func NewListTgwRoutesResponseWithDefaults() *ListTgwRoutesResponse`

NewListTgwRoutesResponseWithDefaults instantiates a new ListTgwRoutesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoutes

`func (o *ListTgwRoutesResponse) GetRoutes() []Route`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *ListTgwRoutesResponse) GetRoutesOk() (*[]Route, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *ListTgwRoutesResponse) SetRoutes(v []Route)`

SetRoutes sets Routes field to given value.


### GetPagination

`func (o *ListTgwRoutesResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListTgwRoutesResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListTgwRoutesResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


