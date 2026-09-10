# EncryptDataResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Encrypt** | [**Encrypt**](Encrypt.md) | 데이터 암호화 요청 또는 결과 정보 | 

## Methods

### NewEncryptDataResponse

`func NewEncryptDataResponse(encrypt Encrypt, ) *EncryptDataResponse`

NewEncryptDataResponse instantiates a new EncryptDataResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEncryptDataResponseWithDefaults

`func NewEncryptDataResponseWithDefaults() *EncryptDataResponse`

NewEncryptDataResponseWithDefaults instantiates a new EncryptDataResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncrypt

`func (o *EncryptDataResponse) GetEncrypt() Encrypt`

GetEncrypt returns the Encrypt field if non-nil, zero value otherwise.

### GetEncryptOk

`func (o *EncryptDataResponse) GetEncryptOk() (*Encrypt, bool)`

GetEncryptOk returns a tuple with the Encrypt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypt

`func (o *EncryptDataResponse) SetEncrypt(v Encrypt)`

SetEncrypt sets Encrypt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


