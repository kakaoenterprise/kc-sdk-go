# Key

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | KMS 키의 고유 ID | [optional] 
**Name** | Pointer to **NullableString** | KMS 키 이름 | [optional] 
**ProjectId** | Pointer to **NullableString** | KMS 키가 속한 프로젝트의 고유 ID | [optional] 
**OwnerName** | Pointer to **NullableString** | 키 소유자 이름 | [optional] 
**Description** | Pointer to **NullableString** | KMS 키 설명 | [optional] 
**Purpose** | Pointer to [**NullableKeyPurpose**](KeyPurpose.md) | 키 용도 | [optional] 
**Status** | Pointer to [**NullableKeyStatus**](KeyStatus.md) | KMS 키 상태 | [optional] 
**Type** | Pointer to [**NullableKeyType**](KeyType.md) | 키 유형 | [optional] 
**Algorithm** | Pointer to [**NullableKeyAlgorithm**](KeyAlgorithm.md) | 키 알고리즘 | [optional] 
**DefaultVersion** | Pointer to **NullableInt32** | KMS 키의 기본 버전 | [optional] 
**IsRotationEnabled** | Pointer to **NullableBool** | 키 자동 순환 활성화 여부 | [optional] 
**RotationPeriod** | Pointer to **NullableInt32** | 키 자동 순환 주기(일) | [optional] 
**IsImmediateDeleteEnabled** | Pointer to **NullableBool** | 키를 즉시 삭제할 수 있는지 여부 | [optional] 
**LastRotationAt** | Pointer to **NullableString** | 키가 마지막으로 순환된 일시 | [optional] 
**NextRotationAt** | Pointer to **NullableString** | 키가 다음에 자동 순환될 예정인 일시 | [optional] 
**CreatedBy** | Pointer to **NullableString** | 키를 생성한 사용자 또는 주체 | [optional] 
**CreatedAt** | Pointer to **NullableString** | 키가 생성된 일시 | [optional] 
**VersionQuota** | Pointer to **NullableInt32** | 생성할 수 있는 최대 버전 수 | [optional] 
**VersionCount** | Pointer to **NullableInt32** | 현재 생성된 버전 수 | [optional] 
**IsVersionCreationAllowed** | Pointer to **NullableBool** | 새 버전을 생성할 수 있는지 여부 | [optional] 
**IsAccessControlEnabled** | Pointer to **NullableBool** | 키 접근 제어 활성화 여부 | [optional] 
**IsAccessible** | Pointer to **NullableBool** | 요청자가 KMS 키에 접근할 수 있는지 여부 | [optional] 

## Methods

### NewKey

`func NewKey() *Key`

NewKey instantiates a new Key object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyWithDefaults

`func NewKeyWithDefaults() *Key`

NewKeyWithDefaults instantiates a new Key object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Key) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Key) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Key) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Key) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Key) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Key) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *Key) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Key) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Key) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Key) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Key) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Key) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetProjectId

`func (o *Key) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *Key) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *Key) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *Key) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *Key) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *Key) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetOwnerName

`func (o *Key) GetOwnerName() string`

GetOwnerName returns the OwnerName field if non-nil, zero value otherwise.

### GetOwnerNameOk

`func (o *Key) GetOwnerNameOk() (*string, bool)`

GetOwnerNameOk returns a tuple with the OwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerName

`func (o *Key) SetOwnerName(v string)`

SetOwnerName sets OwnerName field to given value.

### HasOwnerName

`func (o *Key) HasOwnerName() bool`

HasOwnerName returns a boolean if a field has been set.

### SetOwnerNameNil

`func (o *Key) SetOwnerNameNil(b bool)`

 SetOwnerNameNil sets the value for OwnerName to be an explicit nil

### UnsetOwnerName
`func (o *Key) UnsetOwnerName()`

UnsetOwnerName ensures that no value is present for OwnerName, not even an explicit nil
### GetDescription

`func (o *Key) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Key) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Key) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Key) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Key) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Key) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPurpose

`func (o *Key) GetPurpose() KeyPurpose`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *Key) GetPurposeOk() (*KeyPurpose, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *Key) SetPurpose(v KeyPurpose)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *Key) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### SetPurposeNil

`func (o *Key) SetPurposeNil(b bool)`

 SetPurposeNil sets the value for Purpose to be an explicit nil

### UnsetPurpose
`func (o *Key) UnsetPurpose()`

UnsetPurpose ensures that no value is present for Purpose, not even an explicit nil
### GetStatus

`func (o *Key) GetStatus() KeyStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Key) GetStatusOk() (*KeyStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Key) SetStatus(v KeyStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Key) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *Key) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *Key) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetType

`func (o *Key) GetType() KeyType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Key) GetTypeOk() (*KeyType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Key) SetType(v KeyType)`

SetType sets Type field to given value.

### HasType

`func (o *Key) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *Key) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *Key) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetAlgorithm

`func (o *Key) GetAlgorithm() KeyAlgorithm`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *Key) GetAlgorithmOk() (*KeyAlgorithm, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *Key) SetAlgorithm(v KeyAlgorithm)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *Key) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### SetAlgorithmNil

`func (o *Key) SetAlgorithmNil(b bool)`

 SetAlgorithmNil sets the value for Algorithm to be an explicit nil

### UnsetAlgorithm
`func (o *Key) UnsetAlgorithm()`

UnsetAlgorithm ensures that no value is present for Algorithm, not even an explicit nil
### GetDefaultVersion

`func (o *Key) GetDefaultVersion() int32`

GetDefaultVersion returns the DefaultVersion field if non-nil, zero value otherwise.

### GetDefaultVersionOk

`func (o *Key) GetDefaultVersionOk() (*int32, bool)`

GetDefaultVersionOk returns a tuple with the DefaultVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVersion

`func (o *Key) SetDefaultVersion(v int32)`

SetDefaultVersion sets DefaultVersion field to given value.

### HasDefaultVersion

`func (o *Key) HasDefaultVersion() bool`

HasDefaultVersion returns a boolean if a field has been set.

### SetDefaultVersionNil

`func (o *Key) SetDefaultVersionNil(b bool)`

 SetDefaultVersionNil sets the value for DefaultVersion to be an explicit nil

### UnsetDefaultVersion
`func (o *Key) UnsetDefaultVersion()`

UnsetDefaultVersion ensures that no value is present for DefaultVersion, not even an explicit nil
### GetIsRotationEnabled

`func (o *Key) GetIsRotationEnabled() bool`

GetIsRotationEnabled returns the IsRotationEnabled field if non-nil, zero value otherwise.

### GetIsRotationEnabledOk

`func (o *Key) GetIsRotationEnabledOk() (*bool, bool)`

GetIsRotationEnabledOk returns a tuple with the IsRotationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRotationEnabled

`func (o *Key) SetIsRotationEnabled(v bool)`

SetIsRotationEnabled sets IsRotationEnabled field to given value.

### HasIsRotationEnabled

`func (o *Key) HasIsRotationEnabled() bool`

HasIsRotationEnabled returns a boolean if a field has been set.

### SetIsRotationEnabledNil

`func (o *Key) SetIsRotationEnabledNil(b bool)`

 SetIsRotationEnabledNil sets the value for IsRotationEnabled to be an explicit nil

### UnsetIsRotationEnabled
`func (o *Key) UnsetIsRotationEnabled()`

UnsetIsRotationEnabled ensures that no value is present for IsRotationEnabled, not even an explicit nil
### GetRotationPeriod

`func (o *Key) GetRotationPeriod() int32`

GetRotationPeriod returns the RotationPeriod field if non-nil, zero value otherwise.

### GetRotationPeriodOk

`func (o *Key) GetRotationPeriodOk() (*int32, bool)`

GetRotationPeriodOk returns a tuple with the RotationPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationPeriod

`func (o *Key) SetRotationPeriod(v int32)`

SetRotationPeriod sets RotationPeriod field to given value.

### HasRotationPeriod

`func (o *Key) HasRotationPeriod() bool`

HasRotationPeriod returns a boolean if a field has been set.

### SetRotationPeriodNil

`func (o *Key) SetRotationPeriodNil(b bool)`

 SetRotationPeriodNil sets the value for RotationPeriod to be an explicit nil

### UnsetRotationPeriod
`func (o *Key) UnsetRotationPeriod()`

UnsetRotationPeriod ensures that no value is present for RotationPeriod, not even an explicit nil
### GetIsImmediateDeleteEnabled

`func (o *Key) GetIsImmediateDeleteEnabled() bool`

GetIsImmediateDeleteEnabled returns the IsImmediateDeleteEnabled field if non-nil, zero value otherwise.

### GetIsImmediateDeleteEnabledOk

`func (o *Key) GetIsImmediateDeleteEnabledOk() (*bool, bool)`

GetIsImmediateDeleteEnabledOk returns a tuple with the IsImmediateDeleteEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsImmediateDeleteEnabled

`func (o *Key) SetIsImmediateDeleteEnabled(v bool)`

SetIsImmediateDeleteEnabled sets IsImmediateDeleteEnabled field to given value.

### HasIsImmediateDeleteEnabled

`func (o *Key) HasIsImmediateDeleteEnabled() bool`

HasIsImmediateDeleteEnabled returns a boolean if a field has been set.

### SetIsImmediateDeleteEnabledNil

`func (o *Key) SetIsImmediateDeleteEnabledNil(b bool)`

 SetIsImmediateDeleteEnabledNil sets the value for IsImmediateDeleteEnabled to be an explicit nil

### UnsetIsImmediateDeleteEnabled
`func (o *Key) UnsetIsImmediateDeleteEnabled()`

UnsetIsImmediateDeleteEnabled ensures that no value is present for IsImmediateDeleteEnabled, not even an explicit nil
### GetLastRotationAt

`func (o *Key) GetLastRotationAt() string`

GetLastRotationAt returns the LastRotationAt field if non-nil, zero value otherwise.

### GetLastRotationAtOk

`func (o *Key) GetLastRotationAtOk() (*string, bool)`

GetLastRotationAtOk returns a tuple with the LastRotationAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRotationAt

`func (o *Key) SetLastRotationAt(v string)`

SetLastRotationAt sets LastRotationAt field to given value.

### HasLastRotationAt

`func (o *Key) HasLastRotationAt() bool`

HasLastRotationAt returns a boolean if a field has been set.

### SetLastRotationAtNil

`func (o *Key) SetLastRotationAtNil(b bool)`

 SetLastRotationAtNil sets the value for LastRotationAt to be an explicit nil

### UnsetLastRotationAt
`func (o *Key) UnsetLastRotationAt()`

UnsetLastRotationAt ensures that no value is present for LastRotationAt, not even an explicit nil
### GetNextRotationAt

`func (o *Key) GetNextRotationAt() string`

GetNextRotationAt returns the NextRotationAt field if non-nil, zero value otherwise.

### GetNextRotationAtOk

`func (o *Key) GetNextRotationAtOk() (*string, bool)`

GetNextRotationAtOk returns a tuple with the NextRotationAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRotationAt

`func (o *Key) SetNextRotationAt(v string)`

SetNextRotationAt sets NextRotationAt field to given value.

### HasNextRotationAt

`func (o *Key) HasNextRotationAt() bool`

HasNextRotationAt returns a boolean if a field has been set.

### SetNextRotationAtNil

`func (o *Key) SetNextRotationAtNil(b bool)`

 SetNextRotationAtNil sets the value for NextRotationAt to be an explicit nil

### UnsetNextRotationAt
`func (o *Key) UnsetNextRotationAt()`

UnsetNextRotationAt ensures that no value is present for NextRotationAt, not even an explicit nil
### GetCreatedBy

`func (o *Key) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Key) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Key) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Key) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *Key) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *Key) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetCreatedAt

`func (o *Key) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Key) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Key) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Key) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *Key) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *Key) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetVersionQuota

`func (o *Key) GetVersionQuota() int32`

GetVersionQuota returns the VersionQuota field if non-nil, zero value otherwise.

### GetVersionQuotaOk

`func (o *Key) GetVersionQuotaOk() (*int32, bool)`

GetVersionQuotaOk returns a tuple with the VersionQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionQuota

`func (o *Key) SetVersionQuota(v int32)`

SetVersionQuota sets VersionQuota field to given value.

### HasVersionQuota

`func (o *Key) HasVersionQuota() bool`

HasVersionQuota returns a boolean if a field has been set.

### SetVersionQuotaNil

`func (o *Key) SetVersionQuotaNil(b bool)`

 SetVersionQuotaNil sets the value for VersionQuota to be an explicit nil

### UnsetVersionQuota
`func (o *Key) UnsetVersionQuota()`

UnsetVersionQuota ensures that no value is present for VersionQuota, not even an explicit nil
### GetVersionCount

`func (o *Key) GetVersionCount() int32`

GetVersionCount returns the VersionCount field if non-nil, zero value otherwise.

### GetVersionCountOk

`func (o *Key) GetVersionCountOk() (*int32, bool)`

GetVersionCountOk returns a tuple with the VersionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionCount

`func (o *Key) SetVersionCount(v int32)`

SetVersionCount sets VersionCount field to given value.

### HasVersionCount

`func (o *Key) HasVersionCount() bool`

HasVersionCount returns a boolean if a field has been set.

### SetVersionCountNil

`func (o *Key) SetVersionCountNil(b bool)`

 SetVersionCountNil sets the value for VersionCount to be an explicit nil

### UnsetVersionCount
`func (o *Key) UnsetVersionCount()`

UnsetVersionCount ensures that no value is present for VersionCount, not even an explicit nil
### GetIsVersionCreationAllowed

`func (o *Key) GetIsVersionCreationAllowed() bool`

GetIsVersionCreationAllowed returns the IsVersionCreationAllowed field if non-nil, zero value otherwise.

### GetIsVersionCreationAllowedOk

`func (o *Key) GetIsVersionCreationAllowedOk() (*bool, bool)`

GetIsVersionCreationAllowedOk returns a tuple with the IsVersionCreationAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVersionCreationAllowed

`func (o *Key) SetIsVersionCreationAllowed(v bool)`

SetIsVersionCreationAllowed sets IsVersionCreationAllowed field to given value.

### HasIsVersionCreationAllowed

`func (o *Key) HasIsVersionCreationAllowed() bool`

HasIsVersionCreationAllowed returns a boolean if a field has been set.

### SetIsVersionCreationAllowedNil

`func (o *Key) SetIsVersionCreationAllowedNil(b bool)`

 SetIsVersionCreationAllowedNil sets the value for IsVersionCreationAllowed to be an explicit nil

### UnsetIsVersionCreationAllowed
`func (o *Key) UnsetIsVersionCreationAllowed()`

UnsetIsVersionCreationAllowed ensures that no value is present for IsVersionCreationAllowed, not even an explicit nil
### GetIsAccessControlEnabled

`func (o *Key) GetIsAccessControlEnabled() bool`

GetIsAccessControlEnabled returns the IsAccessControlEnabled field if non-nil, zero value otherwise.

### GetIsAccessControlEnabledOk

`func (o *Key) GetIsAccessControlEnabledOk() (*bool, bool)`

GetIsAccessControlEnabledOk returns a tuple with the IsAccessControlEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAccessControlEnabled

`func (o *Key) SetIsAccessControlEnabled(v bool)`

SetIsAccessControlEnabled sets IsAccessControlEnabled field to given value.

### HasIsAccessControlEnabled

`func (o *Key) HasIsAccessControlEnabled() bool`

HasIsAccessControlEnabled returns a boolean if a field has been set.

### SetIsAccessControlEnabledNil

`func (o *Key) SetIsAccessControlEnabledNil(b bool)`

 SetIsAccessControlEnabledNil sets the value for IsAccessControlEnabled to be an explicit nil

### UnsetIsAccessControlEnabled
`func (o *Key) UnsetIsAccessControlEnabled()`

UnsetIsAccessControlEnabled ensures that no value is present for IsAccessControlEnabled, not even an explicit nil
### GetIsAccessible

`func (o *Key) GetIsAccessible() bool`

GetIsAccessible returns the IsAccessible field if non-nil, zero value otherwise.

### GetIsAccessibleOk

`func (o *Key) GetIsAccessibleOk() (*bool, bool)`

GetIsAccessibleOk returns a tuple with the IsAccessible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAccessible

`func (o *Key) SetIsAccessible(v bool)`

SetIsAccessible sets IsAccessible field to given value.

### HasIsAccessible

`func (o *Key) HasIsAccessible() bool`

HasIsAccessible returns a boolean if a field has been set.

### SetIsAccessibleNil

`func (o *Key) SetIsAccessibleNil(b bool)`

 SetIsAccessibleNil sets the value for IsAccessible to be an explicit nil

### UnsetIsAccessible
`func (o *Key) UnsetIsAccessible()`

UnsetIsAccessible ensures that no value is present for IsAccessible, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


