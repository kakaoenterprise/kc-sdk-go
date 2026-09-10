# ListRouteTablesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VpcRouteTables** | [**[]VpcRouteTable**](VpcRouteTable.md) | 라우팅 테이블 목록 | 
**Pagination** | [**Pagination**](Pagination.md) | 페이지네이션 정보 | 

## Methods

### NewListRouteTablesResponse

`func NewListRouteTablesResponse(vpcRouteTables []VpcRouteTable, pagination Pagination, ) *ListRouteTablesResponse`

NewListRouteTablesResponse instantiates a new ListRouteTablesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListRouteTablesResponseWithDefaults

`func NewListRouteTablesResponseWithDefaults() *ListRouteTablesResponse`

NewListRouteTablesResponseWithDefaults instantiates a new ListRouteTablesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVpcRouteTables

`func (o *ListRouteTablesResponse) GetVpcRouteTables() []VpcRouteTable`

GetVpcRouteTables returns the VpcRouteTables field if non-nil, zero value otherwise.

### GetVpcRouteTablesOk

`func (o *ListRouteTablesResponse) GetVpcRouteTablesOk() (*[]VpcRouteTable, bool)`

GetVpcRouteTablesOk returns a tuple with the VpcRouteTables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcRouteTables

`func (o *ListRouteTablesResponse) SetVpcRouteTables(v []VpcRouteTable)`

SetVpcRouteTables sets VpcRouteTables field to given value.


### GetPagination

`func (o *ListRouteTablesResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListRouteTablesResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListRouteTablesResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


