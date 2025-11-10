# GetTgwsResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tgws** | [**[]BnsTgwV1ApiListTransitGatewaysModelTgwResponseModel**](BnsTgwV1ApiListTransitGatewaysModelTgwResponseModel.md) | Transit Gateway 목록 | 
**Pagination** | [**PaginationModel**](PaginationModel.md) | 페이지네이션 메타테이터 | 

## Methods

### NewGetTgwsResponseModel

`func NewGetTgwsResponseModel(tgws []BnsTgwV1ApiListTransitGatewaysModelTgwResponseModel, pagination PaginationModel, ) *GetTgwsResponseModel`

NewGetTgwsResponseModel instantiates a new GetTgwsResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTgwsResponseModelWithDefaults

`func NewGetTgwsResponseModelWithDefaults() *GetTgwsResponseModel`

NewGetTgwsResponseModelWithDefaults instantiates a new GetTgwsResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTgws

`func (o *GetTgwsResponseModel) GetTgws() []BnsTgwV1ApiListTransitGatewaysModelTgwResponseModel`

GetTgws returns the Tgws field if non-nil, zero value otherwise.

### GetTgwsOk

`func (o *GetTgwsResponseModel) GetTgwsOk() (*[]BnsTgwV1ApiListTransitGatewaysModelTgwResponseModel, bool)`

GetTgwsOk returns a tuple with the Tgws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgws

`func (o *GetTgwsResponseModel) SetTgws(v []BnsTgwV1ApiListTransitGatewaysModelTgwResponseModel)`

SetTgws sets Tgws field to given value.


### GetPagination

`func (o *GetTgwsResponseModel) GetPagination() PaginationModel`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetTgwsResponseModel) GetPaginationOk() (*PaginationModel, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetTgwsResponseModel) SetPagination(v PaginationModel)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


