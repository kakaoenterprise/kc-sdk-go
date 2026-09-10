# Encrypt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CipherText** | **string** | Base64로 인코딩된 암호문 | 
**Version** | **int32** | 키 버전 | 

## Methods

### NewEncrypt

`func NewEncrypt(cipherText string, version int32, ) *Encrypt`

NewEncrypt instantiates a new Encrypt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEncryptWithDefaults

`func NewEncryptWithDefaults() *Encrypt`

NewEncryptWithDefaults instantiates a new Encrypt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCipherText

`func (o *Encrypt) GetCipherText() string`

GetCipherText returns the CipherText field if non-nil, zero value otherwise.

### GetCipherTextOk

`func (o *Encrypt) GetCipherTextOk() (*string, bool)`

GetCipherTextOk returns a tuple with the CipherText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCipherText

`func (o *Encrypt) SetCipherText(v string)`

SetCipherText sets CipherText field to given value.


### GetVersion

`func (o *Encrypt) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Encrypt) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Encrypt) SetVersion(v int32)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


