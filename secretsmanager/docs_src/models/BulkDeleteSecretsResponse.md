# BulkDeleteSecretsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Secrets** | [**Bulk**](Bulk.md) | 시크릿별 일괄 처리 결과 | 

## Methods

### NewBulkDeleteSecretsResponse

`func NewBulkDeleteSecretsResponse(secrets Bulk, ) *BulkDeleteSecretsResponse`

NewBulkDeleteSecretsResponse instantiates a new BulkDeleteSecretsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkDeleteSecretsResponseWithDefaults

`func NewBulkDeleteSecretsResponseWithDefaults() *BulkDeleteSecretsResponse`

NewBulkDeleteSecretsResponseWithDefaults instantiates a new BulkDeleteSecretsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecrets

`func (o *BulkDeleteSecretsResponse) GetSecrets() Bulk`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *BulkDeleteSecretsResponse) GetSecretsOk() (*Bulk, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *BulkDeleteSecretsResponse) SetSecrets(v Bulk)`

SetSecrets sets Secrets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


