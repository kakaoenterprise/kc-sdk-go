# ListInternetGatewaysResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Igws** | [**[]Igw**](Igw.md) | 인터넷 게이트웨이 목록 | 
**Pagination** | [**Pagination**](Pagination.md) | 페이지네이션 정보 | 

## Methods

### NewListInternetGatewaysResponse

`func NewListInternetGatewaysResponse(igws []Igw, pagination Pagination, ) *ListInternetGatewaysResponse`

NewListInternetGatewaysResponse instantiates a new ListInternetGatewaysResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInternetGatewaysResponseWithDefaults

`func NewListInternetGatewaysResponseWithDefaults() *ListInternetGatewaysResponse`

NewListInternetGatewaysResponseWithDefaults instantiates a new ListInternetGatewaysResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIgws

`func (o *ListInternetGatewaysResponse) GetIgws() []Igw`

GetIgws returns the Igws field if non-nil, zero value otherwise.

### GetIgwsOk

`func (o *ListInternetGatewaysResponse) GetIgwsOk() (*[]Igw, bool)`

GetIgwsOk returns a tuple with the Igws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgws

`func (o *ListInternetGatewaysResponse) SetIgws(v []Igw)`

SetIgws sets Igws field to given value.


### GetPagination

`func (o *ListInternetGatewaysResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListInternetGatewaysResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListInternetGatewaysResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


