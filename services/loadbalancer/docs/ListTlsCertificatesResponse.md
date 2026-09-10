# ListTlsCertificatesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Secrets** | [**[]ListTlsCertificatesSecret**](ListTlsCertificatesSecret.md) | TLS 인증서 목록 | 
**Pagination** | [**Pagination**](Pagination.md) | 페이지네이션 정보 | 

## Methods

### NewListTlsCertificatesResponse

`func NewListTlsCertificatesResponse(secrets []ListTlsCertificatesSecret, pagination Pagination, ) *ListTlsCertificatesResponse`

NewListTlsCertificatesResponse instantiates a new ListTlsCertificatesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTlsCertificatesResponseWithDefaults

`func NewListTlsCertificatesResponseWithDefaults() *ListTlsCertificatesResponse`

NewListTlsCertificatesResponseWithDefaults instantiates a new ListTlsCertificatesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecrets

`func (o *ListTlsCertificatesResponse) GetSecrets() []ListTlsCertificatesSecret`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *ListTlsCertificatesResponse) GetSecretsOk() (*[]ListTlsCertificatesSecret, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *ListTlsCertificatesResponse) SetSecrets(v []ListTlsCertificatesSecret)`

SetSecrets sets Secrets field to given value.


### GetPagination

`func (o *ListTlsCertificatesResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListTlsCertificatesResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListTlsCertificatesResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


