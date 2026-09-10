# BulkDeleteSecrets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | **[]string** | 시크릿의 고유 ID 목록 &lt;br/&gt;- [List secrets](https://docs.kakaocloud.com/openapi/security/secrets-manager/list-secrets)에서 확인 | 

## Methods

### NewBulkDeleteSecrets

`func NewBulkDeleteSecrets(ids []string, ) *BulkDeleteSecrets`

NewBulkDeleteSecrets instantiates a new BulkDeleteSecrets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkDeleteSecretsWithDefaults

`func NewBulkDeleteSecretsWithDefaults() *BulkDeleteSecrets`

NewBulkDeleteSecretsWithDefaults instantiates a new BulkDeleteSecrets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *BulkDeleteSecrets) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *BulkDeleteSecrets) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *BulkDeleteSecrets) SetIds(v []string)`

SetIds sets Ids field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


