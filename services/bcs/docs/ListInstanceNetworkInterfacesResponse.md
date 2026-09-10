# ListInstanceNetworkInterfacesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkInterfaces** | [**[]ListInstanceNetworkInterface**](ListInstanceNetworkInterface.md) | 인스턴스에 연결된 네트워크 인터페이스 목록 | 
**Pagination** | [**Pagination**](Pagination.md) | 페이지네이션 정보 | 

## Methods

### NewListInstanceNetworkInterfacesResponse

`func NewListInstanceNetworkInterfacesResponse(networkInterfaces []ListInstanceNetworkInterface, pagination Pagination, ) *ListInstanceNetworkInterfacesResponse`

NewListInstanceNetworkInterfacesResponse instantiates a new ListInstanceNetworkInterfacesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInstanceNetworkInterfacesResponseWithDefaults

`func NewListInstanceNetworkInterfacesResponseWithDefaults() *ListInstanceNetworkInterfacesResponse`

NewListInstanceNetworkInterfacesResponseWithDefaults instantiates a new ListInstanceNetworkInterfacesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkInterfaces

`func (o *ListInstanceNetworkInterfacesResponse) GetNetworkInterfaces() []ListInstanceNetworkInterface`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *ListInstanceNetworkInterfacesResponse) GetNetworkInterfacesOk() (*[]ListInstanceNetworkInterface, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *ListInstanceNetworkInterfacesResponse) SetNetworkInterfaces(v []ListInstanceNetworkInterface)`

SetNetworkInterfaces sets NetworkInterfaces field to given value.


### GetPagination

`func (o *ListInstanceNetworkInterfacesResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListInstanceNetworkInterfacesResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListInstanceNetworkInterfacesResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


