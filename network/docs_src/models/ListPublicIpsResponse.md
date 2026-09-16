# ListPublicIpsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicIps** | [**[]PublicIp**](PublicIp.md) | 퍼블릭 IP 목록 | 
**Pagination** | [**Pagination**](Pagination.md) | 페이지네이션 정보 | 

## Methods

### NewListPublicIpsResponse

`func NewListPublicIpsResponse(publicIps []PublicIp, pagination Pagination, ) *ListPublicIpsResponse`

NewListPublicIpsResponse instantiates a new ListPublicIpsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPublicIpsResponseWithDefaults

`func NewListPublicIpsResponseWithDefaults() *ListPublicIpsResponse`

NewListPublicIpsResponseWithDefaults instantiates a new ListPublicIpsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicIps

`func (o *ListPublicIpsResponse) GetPublicIps() []PublicIp`

GetPublicIps returns the PublicIps field if non-nil, zero value otherwise.

### GetPublicIpsOk

`func (o *ListPublicIpsResponse) GetPublicIpsOk() (*[]PublicIp, bool)`

GetPublicIpsOk returns a tuple with the PublicIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIps

`func (o *ListPublicIpsResponse) SetPublicIps(v []PublicIp)`

SetPublicIps sets PublicIps field to given value.


### GetPagination

`func (o *ListPublicIpsResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListPublicIpsResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListPublicIpsResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


