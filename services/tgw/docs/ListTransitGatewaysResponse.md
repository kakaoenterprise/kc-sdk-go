# ListTransitGatewaysResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tgws** | [**[]Tgw**](Tgw.md) | Transit Gateway 목록 | 
**Pagination** | [**Pagination**](Pagination.md) | 페이지네이션 메타테이터 | 

## Methods

### NewListTransitGatewaysResponse

`func NewListTransitGatewaysResponse(tgws []Tgw, pagination Pagination, ) *ListTransitGatewaysResponse`

NewListTransitGatewaysResponse instantiates a new ListTransitGatewaysResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTransitGatewaysResponseWithDefaults

`func NewListTransitGatewaysResponseWithDefaults() *ListTransitGatewaysResponse`

NewListTransitGatewaysResponseWithDefaults instantiates a new ListTransitGatewaysResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTgws

`func (o *ListTransitGatewaysResponse) GetTgws() []Tgw`

GetTgws returns the Tgws field if non-nil, zero value otherwise.

### GetTgwsOk

`func (o *ListTransitGatewaysResponse) GetTgwsOk() (*[]Tgw, bool)`

GetTgwsOk returns a tuple with the Tgws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgws

`func (o *ListTransitGatewaysResponse) SetTgws(v []Tgw)`

SetTgws sets Tgws field to given value.


### GetPagination

`func (o *ListTransitGatewaysResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListTransitGatewaysResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListTransitGatewaysResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


