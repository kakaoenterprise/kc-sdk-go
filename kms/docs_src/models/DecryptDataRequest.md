# DecryptDataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Decrypt** | [**DecryptData**](DecryptData.md) | 데이터 복호화 요청 또는 결과 정보 | 

## Methods

### NewDecryptDataRequest

`func NewDecryptDataRequest(decrypt DecryptData, ) *DecryptDataRequest`

NewDecryptDataRequest instantiates a new DecryptDataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecryptDataRequestWithDefaults

`func NewDecryptDataRequestWithDefaults() *DecryptDataRequest`

NewDecryptDataRequestWithDefaults instantiates a new DecryptDataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDecrypt

`func (o *DecryptDataRequest) GetDecrypt() DecryptData`

GetDecrypt returns the Decrypt field if non-nil, zero value otherwise.

### GetDecryptOk

`func (o *DecryptDataRequest) GetDecryptOk() (*DecryptData, bool)`

GetDecryptOk returns a tuple with the Decrypt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecrypt

`func (o *DecryptDataRequest) SetDecrypt(v DecryptData)`

SetDecrypt sets Decrypt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


