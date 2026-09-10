# ListServiceKeysResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceKeys** | [**[]ServiceKey**](ServiceKey.md) | 현재 프로젝트의 서비스 키 목록 | 
**Pagination** | [**Pagination**](Pagination.md) | 목록 조회의 페이지네이션 정보 | 

## Methods

### NewListServiceKeysResponse

`func NewListServiceKeysResponse(serviceKeys []ServiceKey, pagination Pagination, ) *ListServiceKeysResponse`

NewListServiceKeysResponse instantiates a new ListServiceKeysResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListServiceKeysResponseWithDefaults

`func NewListServiceKeysResponseWithDefaults() *ListServiceKeysResponse`

NewListServiceKeysResponseWithDefaults instantiates a new ListServiceKeysResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceKeys

`func (o *ListServiceKeysResponse) GetServiceKeys() []ServiceKey`

GetServiceKeys returns the ServiceKeys field if non-nil, zero value otherwise.

### GetServiceKeysOk

`func (o *ListServiceKeysResponse) GetServiceKeysOk() (*[]ServiceKey, bool)`

GetServiceKeysOk returns a tuple with the ServiceKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceKeys

`func (o *ListServiceKeysResponse) SetServiceKeys(v []ServiceKey)`

SetServiceKeys sets ServiceKeys field to given value.


### GetPagination

`func (o *ListServiceKeysResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListServiceKeysResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListServiceKeysResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


