# Secret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | 시크릿 이름 | 
**Id** | **string** | 시크릿의 고유 ID | 
**Description** | Pointer to **NullableString** | 시크릿 설명 | [optional] 
**Status** | [**Status**](Status.md) | 시크릿 상태 | 
**KmsKeyName** | Pointer to **NullableString** | 시크릿 보호에 사용하는 KMS 키 이름 | [optional] 
**Algorithm** | Pointer to **NullableString** | 시크릿 보호에 사용하는 암호화 알고리즘 | [optional] 
**DefaultVersion** | **int32** | 시크릿의 기본 버전 | 
**IsDeletionAllowed** | **bool** | 시크릿을 삭제할 수 있는지 여부 | 
**CreatedAt** | **string** | 시크릿이 생성된 일시 | 
**CreatedBy** | **string** | 시크릿을 생성한 사용자 또는 주체 | 
**VersionQuota** | **int32** | 생성할 수 있는 최대 버전 수 | 
**VersionCount** | **int32** | 현재 생성된 버전 수 | 
**IsVersionCreationAllowed** | **bool** | 새 버전을 생성할 수 있는지 여부 | 
**IsAccessControlEnabled** | Pointer to **NullableBool** | 시크릿 접근 제어 활성화 여부 | [optional] 
**IsAccessible** | Pointer to **NullableBool** | 요청자가 시크릿에 접근할 수 있는지 여부 | [optional] 

## Methods

### NewSecret

`func NewSecret(name string, id string, status Status, defaultVersion int32, isDeletionAllowed bool, createdAt string, createdBy string, versionQuota int32, versionCount int32, isVersionCreationAllowed bool, ) *Secret`

NewSecret instantiates a new Secret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretWithDefaults

`func NewSecretWithDefaults() *Secret`

NewSecretWithDefaults instantiates a new Secret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Secret) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Secret) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Secret) SetName(v string)`

SetName sets Name field to given value.


### GetId

`func (o *Secret) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Secret) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Secret) SetId(v string)`

SetId sets Id field to given value.


### GetDescription

`func (o *Secret) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Secret) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Secret) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Secret) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Secret) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Secret) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetStatus

`func (o *Secret) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Secret) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Secret) SetStatus(v Status)`

SetStatus sets Status field to given value.


### GetKmsKeyName

`func (o *Secret) GetKmsKeyName() string`

GetKmsKeyName returns the KmsKeyName field if non-nil, zero value otherwise.

### GetKmsKeyNameOk

`func (o *Secret) GetKmsKeyNameOk() (*string, bool)`

GetKmsKeyNameOk returns a tuple with the KmsKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmsKeyName

`func (o *Secret) SetKmsKeyName(v string)`

SetKmsKeyName sets KmsKeyName field to given value.

### HasKmsKeyName

`func (o *Secret) HasKmsKeyName() bool`

HasKmsKeyName returns a boolean if a field has been set.

### SetKmsKeyNameNil

`func (o *Secret) SetKmsKeyNameNil(b bool)`

 SetKmsKeyNameNil sets the value for KmsKeyName to be an explicit nil

### UnsetKmsKeyName
`func (o *Secret) UnsetKmsKeyName()`

UnsetKmsKeyName ensures that no value is present for KmsKeyName, not even an explicit nil
### GetAlgorithm

`func (o *Secret) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *Secret) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *Secret) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *Secret) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### SetAlgorithmNil

`func (o *Secret) SetAlgorithmNil(b bool)`

 SetAlgorithmNil sets the value for Algorithm to be an explicit nil

### UnsetAlgorithm
`func (o *Secret) UnsetAlgorithm()`

UnsetAlgorithm ensures that no value is present for Algorithm, not even an explicit nil
### GetDefaultVersion

`func (o *Secret) GetDefaultVersion() int32`

GetDefaultVersion returns the DefaultVersion field if non-nil, zero value otherwise.

### GetDefaultVersionOk

`func (o *Secret) GetDefaultVersionOk() (*int32, bool)`

GetDefaultVersionOk returns a tuple with the DefaultVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVersion

`func (o *Secret) SetDefaultVersion(v int32)`

SetDefaultVersion sets DefaultVersion field to given value.


### GetIsDeletionAllowed

`func (o *Secret) GetIsDeletionAllowed() bool`

GetIsDeletionAllowed returns the IsDeletionAllowed field if non-nil, zero value otherwise.

### GetIsDeletionAllowedOk

`func (o *Secret) GetIsDeletionAllowedOk() (*bool, bool)`

GetIsDeletionAllowedOk returns a tuple with the IsDeletionAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeletionAllowed

`func (o *Secret) SetIsDeletionAllowed(v bool)`

SetIsDeletionAllowed sets IsDeletionAllowed field to given value.


### GetCreatedAt

`func (o *Secret) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Secret) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Secret) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *Secret) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Secret) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Secret) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetVersionQuota

`func (o *Secret) GetVersionQuota() int32`

GetVersionQuota returns the VersionQuota field if non-nil, zero value otherwise.

### GetVersionQuotaOk

`func (o *Secret) GetVersionQuotaOk() (*int32, bool)`

GetVersionQuotaOk returns a tuple with the VersionQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionQuota

`func (o *Secret) SetVersionQuota(v int32)`

SetVersionQuota sets VersionQuota field to given value.


### GetVersionCount

`func (o *Secret) GetVersionCount() int32`

GetVersionCount returns the VersionCount field if non-nil, zero value otherwise.

### GetVersionCountOk

`func (o *Secret) GetVersionCountOk() (*int32, bool)`

GetVersionCountOk returns a tuple with the VersionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionCount

`func (o *Secret) SetVersionCount(v int32)`

SetVersionCount sets VersionCount field to given value.


### GetIsVersionCreationAllowed

`func (o *Secret) GetIsVersionCreationAllowed() bool`

GetIsVersionCreationAllowed returns the IsVersionCreationAllowed field if non-nil, zero value otherwise.

### GetIsVersionCreationAllowedOk

`func (o *Secret) GetIsVersionCreationAllowedOk() (*bool, bool)`

GetIsVersionCreationAllowedOk returns a tuple with the IsVersionCreationAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVersionCreationAllowed

`func (o *Secret) SetIsVersionCreationAllowed(v bool)`

SetIsVersionCreationAllowed sets IsVersionCreationAllowed field to given value.


### GetIsAccessControlEnabled

`func (o *Secret) GetIsAccessControlEnabled() bool`

GetIsAccessControlEnabled returns the IsAccessControlEnabled field if non-nil, zero value otherwise.

### GetIsAccessControlEnabledOk

`func (o *Secret) GetIsAccessControlEnabledOk() (*bool, bool)`

GetIsAccessControlEnabledOk returns a tuple with the IsAccessControlEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAccessControlEnabled

`func (o *Secret) SetIsAccessControlEnabled(v bool)`

SetIsAccessControlEnabled sets IsAccessControlEnabled field to given value.

### HasIsAccessControlEnabled

`func (o *Secret) HasIsAccessControlEnabled() bool`

HasIsAccessControlEnabled returns a boolean if a field has been set.

### SetIsAccessControlEnabledNil

`func (o *Secret) SetIsAccessControlEnabledNil(b bool)`

 SetIsAccessControlEnabledNil sets the value for IsAccessControlEnabled to be an explicit nil

### UnsetIsAccessControlEnabled
`func (o *Secret) UnsetIsAccessControlEnabled()`

UnsetIsAccessControlEnabled ensures that no value is present for IsAccessControlEnabled, not even an explicit nil
### GetIsAccessible

`func (o *Secret) GetIsAccessible() bool`

GetIsAccessible returns the IsAccessible field if non-nil, zero value otherwise.

### GetIsAccessibleOk

`func (o *Secret) GetIsAccessibleOk() (*bool, bool)`

GetIsAccessibleOk returns a tuple with the IsAccessible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAccessible

`func (o *Secret) SetIsAccessible(v bool)`

SetIsAccessible sets IsAccessible field to given value.

### HasIsAccessible

`func (o *Secret) HasIsAccessible() bool`

HasIsAccessible returns a boolean if a field has been set.

### SetIsAccessibleNil

`func (o *Secret) SetIsAccessibleNil(b bool)`

 SetIsAccessibleNil sets the value for IsAccessible to be an explicit nil

### UnsetIsAccessible
`func (o *Secret) UnsetIsAccessible()`

UnsetIsAccessible ensures that no value is present for IsAccessible, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


