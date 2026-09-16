# BulkDeactivateSecretVersionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Versions** | [**Bulk**](Bulk.md) | 시크릿 버전별 처리 결과 | 

## Methods

### NewBulkDeactivateSecretVersionsResponse

`func NewBulkDeactivateSecretVersionsResponse(versions Bulk, ) *BulkDeactivateSecretVersionsResponse`

NewBulkDeactivateSecretVersionsResponse instantiates a new BulkDeactivateSecretVersionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkDeactivateSecretVersionsResponseWithDefaults

`func NewBulkDeactivateSecretVersionsResponseWithDefaults() *BulkDeactivateSecretVersionsResponse`

NewBulkDeactivateSecretVersionsResponseWithDefaults instantiates a new BulkDeactivateSecretVersionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersions

`func (o *BulkDeactivateSecretVersionsResponse) GetVersions() Bulk`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *BulkDeactivateSecretVersionsResponse) GetVersionsOk() (*Bulk, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *BulkDeactivateSecretVersionsResponse) SetVersions(v Bulk)`

SetVersions sets Versions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


