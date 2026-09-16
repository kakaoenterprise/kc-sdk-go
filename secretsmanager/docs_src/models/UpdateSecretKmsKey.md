# UpdateSecretKmsKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KmsKeyId** | Pointer to **NullableString** | 시크릿 보호에 사용하는 KMS 키의 고유 ID - [List user keys](/openapi/security/kms/list-user-keys)에서 확인 | [optional] 

## Methods

### NewUpdateSecretKmsKey

`func NewUpdateSecretKmsKey() *UpdateSecretKmsKey`

NewUpdateSecretKmsKey instantiates a new UpdateSecretKmsKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSecretKmsKeyWithDefaults

`func NewUpdateSecretKmsKeyWithDefaults() *UpdateSecretKmsKey`

NewUpdateSecretKmsKeyWithDefaults instantiates a new UpdateSecretKmsKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKmsKeyId

`func (o *UpdateSecretKmsKey) GetKmsKeyId() string`

GetKmsKeyId returns the KmsKeyId field if non-nil, zero value otherwise.

### GetKmsKeyIdOk

`func (o *UpdateSecretKmsKey) GetKmsKeyIdOk() (*string, bool)`

GetKmsKeyIdOk returns a tuple with the KmsKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmsKeyId

`func (o *UpdateSecretKmsKey) SetKmsKeyId(v string)`

SetKmsKeyId sets KmsKeyId field to given value.

### HasKmsKeyId

`func (o *UpdateSecretKmsKey) HasKmsKeyId() bool`

HasKmsKeyId returns a boolean if a field has been set.

### SetKmsKeyIdNil

`func (o *UpdateSecretKmsKey) SetKmsKeyIdNil(b bool)`

 SetKmsKeyIdNil sets the value for KmsKeyId to be an explicit nil

### UnsetKmsKeyId
`func (o *UpdateSecretKmsKey) UnsetKmsKeyId()`

UnsetKmsKeyId ensures that no value is present for KmsKeyId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


