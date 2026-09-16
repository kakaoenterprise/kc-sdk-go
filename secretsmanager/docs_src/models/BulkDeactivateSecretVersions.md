# BulkDeactivateSecretVersions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Versions** | **[]int32** | 처리할 시크릿 버전 목록 - [List secret versions](/openapi/security/secrets-manager/list-secret-versions)에서 확인 | 

## Methods

### NewBulkDeactivateSecretVersions

`func NewBulkDeactivateSecretVersions(versions []int32, ) *BulkDeactivateSecretVersions`

NewBulkDeactivateSecretVersions instantiates a new BulkDeactivateSecretVersions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkDeactivateSecretVersionsWithDefaults

`func NewBulkDeactivateSecretVersionsWithDefaults() *BulkDeactivateSecretVersions`

NewBulkDeactivateSecretVersionsWithDefaults instantiates a new BulkDeactivateSecretVersions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersions

`func (o *BulkDeactivateSecretVersions) GetVersions() []int32`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *BulkDeactivateSecretVersions) GetVersionsOk() (*[]int32, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *BulkDeactivateSecretVersions) SetVersions(v []int32)`

SetVersions sets Versions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


