# UpdateSecretKmsKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Secret** | [**UpdateSecretKmsKey**](UpdateSecretKmsKey.md) | 시크릿 KMS 키 변경 정보 | 

## Methods

### NewUpdateSecretKmsKeyRequest

`func NewUpdateSecretKmsKeyRequest(secret UpdateSecretKmsKey, ) *UpdateSecretKmsKeyRequest`

NewUpdateSecretKmsKeyRequest instantiates a new UpdateSecretKmsKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSecretKmsKeyRequestWithDefaults

`func NewUpdateSecretKmsKeyRequestWithDefaults() *UpdateSecretKmsKeyRequest`

NewUpdateSecretKmsKeyRequestWithDefaults instantiates a new UpdateSecretKmsKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecret

`func (o *UpdateSecretKmsKeyRequest) GetSecret() UpdateSecretKmsKey`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *UpdateSecretKmsKeyRequest) GetSecretOk() (*UpdateSecretKmsKey, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *UpdateSecretKmsKeyRequest) SetSecret(v UpdateSecretKmsKey)`

SetSecret sets Secret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


