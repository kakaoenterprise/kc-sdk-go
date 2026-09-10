# DecryptData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CipherText** | **string** | 복호화할 Base64 인코딩 암호문 &lt;br/&gt; - [Encrypt data](https://docs.kakaocloud.com/openapi/security/kms/encrypt-data) 응답에서 확인 | 
**Aad** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDecryptData

`func NewDecryptData(cipherText string, ) *DecryptData`

NewDecryptData instantiates a new DecryptData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecryptDataWithDefaults

`func NewDecryptDataWithDefaults() *DecryptData`

NewDecryptDataWithDefaults instantiates a new DecryptData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCipherText

`func (o *DecryptData) GetCipherText() string`

GetCipherText returns the CipherText field if non-nil, zero value otherwise.

### GetCipherTextOk

`func (o *DecryptData) GetCipherTextOk() (*string, bool)`

GetCipherTextOk returns a tuple with the CipherText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCipherText

`func (o *DecryptData) SetCipherText(v string)`

SetCipherText sets CipherText field to given value.


### GetAad

`func (o *DecryptData) GetAad() string`

GetAad returns the Aad field if non-nil, zero value otherwise.

### GetAadOk

`func (o *DecryptData) GetAadOk() (*string, bool)`

GetAadOk returns a tuple with the Aad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAad

`func (o *DecryptData) SetAad(v string)`

SetAad sets Aad field to given value.

### HasAad

`func (o *DecryptData) HasAad() bool`

HasAad returns a boolean if a field has been set.

### SetAadNil

`func (o *DecryptData) SetAadNil(b bool)`

 SetAadNil sets the value for Aad to be an explicit nil

### UnsetAad
`func (o *DecryptData) UnsetAad()`

UnsetAad ensures that no value is present for Aad, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


