# ListTgwRouteTablesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TgwRouteTables** | [**[]TgwRouteTable**](TgwRouteTable.md) | 조회된 Transit Gateway 라우팅 테이블 목록 | 
**Pagination** | [**Pagination**](Pagination.md) | 목록 조회 결과의 페이지네이션 메타데이터 | 

## Methods

### NewListTgwRouteTablesResponse

`func NewListTgwRouteTablesResponse(tgwRouteTables []TgwRouteTable, pagination Pagination, ) *ListTgwRouteTablesResponse`

NewListTgwRouteTablesResponse instantiates a new ListTgwRouteTablesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTgwRouteTablesResponseWithDefaults

`func NewListTgwRouteTablesResponseWithDefaults() *ListTgwRouteTablesResponse`

NewListTgwRouteTablesResponseWithDefaults instantiates a new ListTgwRouteTablesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTgwRouteTables

`func (o *ListTgwRouteTablesResponse) GetTgwRouteTables() []TgwRouteTable`

GetTgwRouteTables returns the TgwRouteTables field if non-nil, zero value otherwise.

### GetTgwRouteTablesOk

`func (o *ListTgwRouteTablesResponse) GetTgwRouteTablesOk() (*[]TgwRouteTable, bool)`

GetTgwRouteTablesOk returns a tuple with the TgwRouteTables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgwRouteTables

`func (o *ListTgwRouteTablesResponse) SetTgwRouteTables(v []TgwRouteTable)`

SetTgwRouteTables sets TgwRouteTables field to given value.


### GetPagination

`func (o *ListTgwRouteTablesResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListTgwRouteTablesResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListTgwRouteTablesResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


