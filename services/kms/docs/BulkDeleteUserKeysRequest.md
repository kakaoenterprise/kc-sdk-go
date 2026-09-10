# BulkDeleteUserKeysRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keys** | [**BulkDeleteUserKeys**](BulkDeleteUserKeys.md) | 일괄 삭제할 KMS 키 정보 | 

## Methods

### NewBulkDeleteUserKeysRequest

`func NewBulkDeleteUserKeysRequest(keys BulkDeleteUserKeys, ) *BulkDeleteUserKeysRequest`

NewBulkDeleteUserKeysRequest instantiates a new BulkDeleteUserKeysRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkDeleteUserKeysRequestWithDefaults

`func NewBulkDeleteUserKeysRequestWithDefaults() *BulkDeleteUserKeysRequest`

NewBulkDeleteUserKeysRequestWithDefaults instantiates a new BulkDeleteUserKeysRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeys

`func (o *BulkDeleteUserKeysRequest) GetKeys() BulkDeleteUserKeys`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *BulkDeleteUserKeysRequest) GetKeysOk() (*BulkDeleteUserKeys, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *BulkDeleteUserKeysRequest) SetKeys(v BulkDeleteUserKeys)`

SetKeys sets Keys field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


