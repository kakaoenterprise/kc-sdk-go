# EncryptDataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Encrypt** | [**EncryptData**](EncryptData.md) | 데이터 암호화 요청 또는 결과 정보 | 

## Methods

### NewEncryptDataRequest

`func NewEncryptDataRequest(encrypt EncryptData, ) *EncryptDataRequest`

NewEncryptDataRequest instantiates a new EncryptDataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEncryptDataRequestWithDefaults

`func NewEncryptDataRequestWithDefaults() *EncryptDataRequest`

NewEncryptDataRequestWithDefaults instantiates a new EncryptDataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncrypt

`func (o *EncryptDataRequest) GetEncrypt() EncryptData`

GetEncrypt returns the Encrypt field if non-nil, zero value otherwise.

### GetEncryptOk

`func (o *EncryptDataRequest) GetEncryptOk() (*EncryptData, bool)`

GetEncryptOk returns a tuple with the Encrypt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypt

`func (o *EncryptDataRequest) SetEncrypt(v EncryptData)`

SetEncrypt sets Encrypt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


