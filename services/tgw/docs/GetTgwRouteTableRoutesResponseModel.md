# GetTgwRouteTableRoutesResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Routes** | [**[]BnsTgwV1ApiListTgwRoutesModelRouteResponseModel**](BnsTgwV1ApiListTgwRoutesModelRouteResponseModel.md) | 조회된 Route 목록 | 
**Pagination** | [**PaginationModel**](PaginationModel.md) | 페이지네이션 메타데이터 | 

## Methods

### NewGetTgwRouteTableRoutesResponseModel

`func NewGetTgwRouteTableRoutesResponseModel(routes []BnsTgwV1ApiListTgwRoutesModelRouteResponseModel, pagination PaginationModel, ) *GetTgwRouteTableRoutesResponseModel`

NewGetTgwRouteTableRoutesResponseModel instantiates a new GetTgwRouteTableRoutesResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTgwRouteTableRoutesResponseModelWithDefaults

`func NewGetTgwRouteTableRoutesResponseModelWithDefaults() *GetTgwRouteTableRoutesResponseModel`

NewGetTgwRouteTableRoutesResponseModelWithDefaults instantiates a new GetTgwRouteTableRoutesResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoutes

`func (o *GetTgwRouteTableRoutesResponseModel) GetRoutes() []BnsTgwV1ApiListTgwRoutesModelRouteResponseModel`

GetRoutes returns the Routes field if non-nil, zero value otherwise.

### GetRoutesOk

`func (o *GetTgwRouteTableRoutesResponseModel) GetRoutesOk() (*[]BnsTgwV1ApiListTgwRoutesModelRouteResponseModel, bool)`

GetRoutesOk returns a tuple with the Routes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutes

`func (o *GetTgwRouteTableRoutesResponseModel) SetRoutes(v []BnsTgwV1ApiListTgwRoutesModelRouteResponseModel)`

SetRoutes sets Routes field to given value.


### GetPagination

`func (o *GetTgwRouteTableRoutesResponseModel) GetPagination() PaginationModel`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetTgwRouteTableRoutesResponseModel) GetPaginationOk() (*PaginationModel, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetTgwRouteTableRoutesResponseModel) SetPagination(v PaginationModel)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


