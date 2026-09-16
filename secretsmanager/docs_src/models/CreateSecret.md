# CreateSecret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | 시크릿 이름 | 
**Description** | Pointer to **NullableString** | 시크릿 설명 | [optional] 
**SecretData** | **interface{}** |  | 
**KmsKeyId** | Pointer to **NullableString** | 시크릿 보호에 사용하는 KMS 키의 고유 ID - [List user keys](/openapi/security/kms/list-user-keys)에서 확인 | [optional] 
**IsAccessControlEnabled** | Pointer to **NullableBool** | 시크릿 접근 제어 활성화 여부 | [optional] 
**AccessTargetIds** | Pointer to **[]string** | 접근을 허용할 IAM 사용자 또는 서비스 계정의 고유 ID 목록 | [optional] 

## Methods

### NewCreateSecret

`func NewCreateSecret(name string, secretData interface{}, ) *CreateSecret`

NewCreateSecret instantiates a new CreateSecret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSecretWithDefaults

`func NewCreateSecretWithDefaults() *CreateSecret`

NewCreateSecretWithDefaults instantiates a new CreateSecret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateSecret) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateSecret) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateSecret) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateSecret) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateSecret) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateSecret) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateSecret) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateSecret) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateSecret) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSecretData

`func (o *CreateSecret) GetSecretData() interface{}`

GetSecretData returns the SecretData field if non-nil, zero value otherwise.

### GetSecretDataOk

`func (o *CreateSecret) GetSecretDataOk() (*interface{}, bool)`

GetSecretDataOk returns a tuple with the SecretData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretData

`func (o *CreateSecret) SetSecretData(v interface{})`

SetSecretData sets SecretData field to given value.


### SetSecretDataNil

`func (o *CreateSecret) SetSecretDataNil(b bool)`

 SetSecretDataNil sets the value for SecretData to be an explicit nil

### UnsetSecretData
`func (o *CreateSecret) UnsetSecretData()`

UnsetSecretData ensures that no value is present for SecretData, not even an explicit nil
### GetKmsKeyId

`func (o *CreateSecret) GetKmsKeyId() string`

GetKmsKeyId returns the KmsKeyId field if non-nil, zero value otherwise.

### GetKmsKeyIdOk

`func (o *CreateSecret) GetKmsKeyIdOk() (*string, bool)`

GetKmsKeyIdOk returns a tuple with the KmsKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmsKeyId

`func (o *CreateSecret) SetKmsKeyId(v string)`

SetKmsKeyId sets KmsKeyId field to given value.

### HasKmsKeyId

`func (o *CreateSecret) HasKmsKeyId() bool`

HasKmsKeyId returns a boolean if a field has been set.

### SetKmsKeyIdNil

`func (o *CreateSecret) SetKmsKeyIdNil(b bool)`

 SetKmsKeyIdNil sets the value for KmsKeyId to be an explicit nil

### UnsetKmsKeyId
`func (o *CreateSecret) UnsetKmsKeyId()`

UnsetKmsKeyId ensures that no value is present for KmsKeyId, not even an explicit nil
### GetIsAccessControlEnabled

`func (o *CreateSecret) GetIsAccessControlEnabled() bool`

GetIsAccessControlEnabled returns the IsAccessControlEnabled field if non-nil, zero value otherwise.

### GetIsAccessControlEnabledOk

`func (o *CreateSecret) GetIsAccessControlEnabledOk() (*bool, bool)`

GetIsAccessControlEnabledOk returns a tuple with the IsAccessControlEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAccessControlEnabled

`func (o *CreateSecret) SetIsAccessControlEnabled(v bool)`

SetIsAccessControlEnabled sets IsAccessControlEnabled field to given value.

### HasIsAccessControlEnabled

`func (o *CreateSecret) HasIsAccessControlEnabled() bool`

HasIsAccessControlEnabled returns a boolean if a field has been set.

### SetIsAccessControlEnabledNil

`func (o *CreateSecret) SetIsAccessControlEnabledNil(b bool)`

 SetIsAccessControlEnabledNil sets the value for IsAccessControlEnabled to be an explicit nil

### UnsetIsAccessControlEnabled
`func (o *CreateSecret) UnsetIsAccessControlEnabled()`

UnsetIsAccessControlEnabled ensures that no value is present for IsAccessControlEnabled, not even an explicit nil
### GetAccessTargetIds

`func (o *CreateSecret) GetAccessTargetIds() []string`

GetAccessTargetIds returns the AccessTargetIds field if non-nil, zero value otherwise.

### GetAccessTargetIdsOk

`func (o *CreateSecret) GetAccessTargetIdsOk() (*[]string, bool)`

GetAccessTargetIdsOk returns a tuple with the AccessTargetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTargetIds

`func (o *CreateSecret) SetAccessTargetIds(v []string)`

SetAccessTargetIds sets AccessTargetIds field to given value.

### HasAccessTargetIds

`func (o *CreateSecret) HasAccessTargetIds() bool`

HasAccessTargetIds returns a boolean if a field has been set.

### SetAccessTargetIdsNil

`func (o *CreateSecret) SetAccessTargetIdsNil(b bool)`

 SetAccessTargetIdsNil sets the value for AccessTargetIds to be an explicit nil

### UnsetAccessTargetIds
`func (o *CreateSecret) UnsetAccessTargetIds()`

UnsetAccessTargetIds ensures that no value is present for AccessTargetIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


