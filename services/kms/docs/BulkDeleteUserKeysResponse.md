# BulkDeleteUserKeysResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keys** | [**Result**](Result.md) | KMS 키별 일괄 처리 결과 | 

## Methods

### NewBulkDeleteUserKeysResponse

`func NewBulkDeleteUserKeysResponse(keys Result, ) *BulkDeleteUserKeysResponse`

NewBulkDeleteUserKeysResponse instantiates a new BulkDeleteUserKeysResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkDeleteUserKeysResponseWithDefaults

`func NewBulkDeleteUserKeysResponseWithDefaults() *BulkDeleteUserKeysResponse`

NewBulkDeleteUserKeysResponseWithDefaults instantiates a new BulkDeleteUserKeysResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeys

`func (o *BulkDeleteUserKeysResponse) GetKeys() Result`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *BulkDeleteUserKeysResponse) GetKeysOk() (*Result, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *BulkDeleteUserKeysResponse) SetKeys(v Result)`

SetKeys sets Keys field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


