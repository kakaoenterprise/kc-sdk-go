# BulkCancelScheduledSecretVersionDestruction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Versions** | **[]int32** | 처리할 시크릿 버전 목록 &lt;br/&gt;- [List secret versions](https://docs.kakaocloud.com/openapi/security/secrets-manager/list-secret-versions)에서 확인 | 

## Methods

### NewBulkCancelScheduledSecretVersionDestruction

`func NewBulkCancelScheduledSecretVersionDestruction(versions []int32, ) *BulkCancelScheduledSecretVersionDestruction`

NewBulkCancelScheduledSecretVersionDestruction instantiates a new BulkCancelScheduledSecretVersionDestruction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkCancelScheduledSecretVersionDestructionWithDefaults

`func NewBulkCancelScheduledSecretVersionDestructionWithDefaults() *BulkCancelScheduledSecretVersionDestruction`

NewBulkCancelScheduledSecretVersionDestructionWithDefaults instantiates a new BulkCancelScheduledSecretVersionDestruction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersions

`func (o *BulkCancelScheduledSecretVersionDestruction) GetVersions() []int32`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *BulkCancelScheduledSecretVersionDestruction) GetVersionsOk() (*[]int32, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *BulkCancelScheduledSecretVersionDestruction) SetVersions(v []int32)`

SetVersions sets Versions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


