# ListVpcsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vpcs** | [**[]Vpc**](Vpc.md) | VPC 목록 | 
**Pagination** | [**Pagination**](Pagination.md) | 페이지네이션 정보 | 

## Methods

### NewListVpcsResponse

`func NewListVpcsResponse(vpcs []Vpc, pagination Pagination, ) *ListVpcsResponse`

NewListVpcsResponse instantiates a new ListVpcsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListVpcsResponseWithDefaults

`func NewListVpcsResponseWithDefaults() *ListVpcsResponse`

NewListVpcsResponseWithDefaults instantiates a new ListVpcsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVpcs

`func (o *ListVpcsResponse) GetVpcs() []Vpc`

GetVpcs returns the Vpcs field if non-nil, zero value otherwise.

### GetVpcsOk

`func (o *ListVpcsResponse) GetVpcsOk() (*[]Vpc, bool)`

GetVpcsOk returns a tuple with the Vpcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcs

`func (o *ListVpcsResponse) SetVpcs(v []Vpc)`

SetVpcs sets Vpcs field to given value.


### GetPagination

`func (o *ListVpcsResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListVpcsResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListVpcsResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


