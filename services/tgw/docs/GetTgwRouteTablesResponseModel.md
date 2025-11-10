# GetTgwRouteTablesResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TgwRouteTables** | [**[]BnsTgwV1ApiListTgwRouteTablesModelTgwRouteTableResponseModel**](BnsTgwV1ApiListTgwRouteTablesModelTgwRouteTableResponseModel.md) | 조회된 Transit Gateway 라우팅 테이블 목록 | 
**Pagination** | [**PaginationModel**](PaginationModel.md) | 목록 조회 결과의 페이지네이션 메타데이터 | 

## Methods

### NewGetTgwRouteTablesResponseModel

`func NewGetTgwRouteTablesResponseModel(tgwRouteTables []BnsTgwV1ApiListTgwRouteTablesModelTgwRouteTableResponseModel, pagination PaginationModel, ) *GetTgwRouteTablesResponseModel`

NewGetTgwRouteTablesResponseModel instantiates a new GetTgwRouteTablesResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTgwRouteTablesResponseModelWithDefaults

`func NewGetTgwRouteTablesResponseModelWithDefaults() *GetTgwRouteTablesResponseModel`

NewGetTgwRouteTablesResponseModelWithDefaults instantiates a new GetTgwRouteTablesResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTgwRouteTables

`func (o *GetTgwRouteTablesResponseModel) GetTgwRouteTables() []BnsTgwV1ApiListTgwRouteTablesModelTgwRouteTableResponseModel`

GetTgwRouteTables returns the TgwRouteTables field if non-nil, zero value otherwise.

### GetTgwRouteTablesOk

`func (o *GetTgwRouteTablesResponseModel) GetTgwRouteTablesOk() (*[]BnsTgwV1ApiListTgwRouteTablesModelTgwRouteTableResponseModel, bool)`

GetTgwRouteTablesOk returns a tuple with the TgwRouteTables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgwRouteTables

`func (o *GetTgwRouteTablesResponseModel) SetTgwRouteTables(v []BnsTgwV1ApiListTgwRouteTablesModelTgwRouteTableResponseModel)`

SetTgwRouteTables sets TgwRouteTables field to given value.


### GetPagination

`func (o *GetTgwRouteTablesResponseModel) GetPagination() PaginationModel`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetTgwRouteTablesResponseModel) GetPaginationOk() (*PaginationModel, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetTgwRouteTablesResponseModel) SetPagination(v PaginationModel)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


