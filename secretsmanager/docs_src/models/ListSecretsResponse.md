# ListSecretsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Secrets** | [**[]Secret**](Secret.md) | 시크릿 목록 | 
**Pagination** | [**Pagination**](Pagination.md) | 목록 조회의 페이지네이션 정보 | 

## Methods

### NewListSecretsResponse

`func NewListSecretsResponse(secrets []Secret, pagination Pagination, ) *ListSecretsResponse`

NewListSecretsResponse instantiates a new ListSecretsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSecretsResponseWithDefaults

`func NewListSecretsResponseWithDefaults() *ListSecretsResponse`

NewListSecretsResponseWithDefaults instantiates a new ListSecretsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecrets

`func (o *ListSecretsResponse) GetSecrets() []Secret`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *ListSecretsResponse) GetSecretsOk() (*[]Secret, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *ListSecretsResponse) SetSecrets(v []Secret)`

SetSecrets sets Secrets field to given value.


### GetPagination

`func (o *ListSecretsResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListSecretsResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListSecretsResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


