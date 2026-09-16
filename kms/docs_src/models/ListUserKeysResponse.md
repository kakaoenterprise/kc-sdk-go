# ListUserKeysResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keys** | [**[]Key**](Key.md) | KMS 키 목록 | 
**Pagination** | [**Pagination**](Pagination.md) | 목록 조회의 페이지네이션 정보 | 

## Methods

### NewListUserKeysResponse

`func NewListUserKeysResponse(keys []Key, pagination Pagination, ) *ListUserKeysResponse`

NewListUserKeysResponse instantiates a new ListUserKeysResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUserKeysResponseWithDefaults

`func NewListUserKeysResponseWithDefaults() *ListUserKeysResponse`

NewListUserKeysResponseWithDefaults instantiates a new ListUserKeysResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeys

`func (o *ListUserKeysResponse) GetKeys() []Key`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *ListUserKeysResponse) GetKeysOk() (*[]Key, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *ListUserKeysResponse) SetKeys(v []Key)`

SetKeys sets Keys field to given value.


### GetPagination

`func (o *ListUserKeysResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListUserKeysResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListUserKeysResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


