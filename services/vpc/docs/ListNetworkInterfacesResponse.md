# ListNetworkInterfacesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkInterfaces** | [**[]NetworkInterface**](NetworkInterface.md) | 네트워크 인터페이스 목록 | 
**Pagination** | [**Pagination**](Pagination.md) | 페이지네이션 정보 | 

## Methods

### NewListNetworkInterfacesResponse

`func NewListNetworkInterfacesResponse(networkInterfaces []NetworkInterface, pagination Pagination, ) *ListNetworkInterfacesResponse`

NewListNetworkInterfacesResponse instantiates a new ListNetworkInterfacesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListNetworkInterfacesResponseWithDefaults

`func NewListNetworkInterfacesResponseWithDefaults() *ListNetworkInterfacesResponse`

NewListNetworkInterfacesResponseWithDefaults instantiates a new ListNetworkInterfacesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkInterfaces

`func (o *ListNetworkInterfacesResponse) GetNetworkInterfaces() []NetworkInterface`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *ListNetworkInterfacesResponse) GetNetworkInterfacesOk() (*[]NetworkInterface, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *ListNetworkInterfacesResponse) SetNetworkInterfaces(v []NetworkInterface)`

SetNetworkInterfaces sets NetworkInterfaces field to given value.


### GetPagination

`func (o *ListNetworkInterfacesResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListNetworkInterfacesResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListNetworkInterfacesResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


