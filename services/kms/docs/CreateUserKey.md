# CreateUserKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | KMS 키 이름 | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Type** | [**KeyType**](KeyType.md) | 키 유형 | 
**Purpose** | [**KeyPurpose**](KeyPurpose.md) | 키 용도 | 
**Algorithm** | [**KeyAlgorithm**](KeyAlgorithm.md) | 키 알고리즘 | 
**RotationPeriod** | Pointer to **NullableInt32** |  | [optional] 
**IsPendingActivation** | Pointer to **NullableBool** |  | [optional] 
**IsAccessControlEnabled** | **bool** | 키 접근 제어 활성화 여부 | 
**AccessTargetIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCreateUserKey

`func NewCreateUserKey(name string, type_ KeyType, purpose KeyPurpose, algorithm KeyAlgorithm, isAccessControlEnabled bool, ) *CreateUserKey`

NewCreateUserKey instantiates a new CreateUserKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserKeyWithDefaults

`func NewCreateUserKeyWithDefaults() *CreateUserKey`

NewCreateUserKeyWithDefaults instantiates a new CreateUserKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateUserKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateUserKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateUserKey) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateUserKey) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateUserKey) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateUserKey) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateUserKey) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateUserKey) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateUserKey) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetType

`func (o *CreateUserKey) GetType() KeyType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateUserKey) GetTypeOk() (*KeyType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateUserKey) SetType(v KeyType)`

SetType sets Type field to given value.


### GetPurpose

`func (o *CreateUserKey) GetPurpose() KeyPurpose`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *CreateUserKey) GetPurposeOk() (*KeyPurpose, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *CreateUserKey) SetPurpose(v KeyPurpose)`

SetPurpose sets Purpose field to given value.


### GetAlgorithm

`func (o *CreateUserKey) GetAlgorithm() KeyAlgorithm`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *CreateUserKey) GetAlgorithmOk() (*KeyAlgorithm, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *CreateUserKey) SetAlgorithm(v KeyAlgorithm)`

SetAlgorithm sets Algorithm field to given value.


### GetRotationPeriod

`func (o *CreateUserKey) GetRotationPeriod() int32`

GetRotationPeriod returns the RotationPeriod field if non-nil, zero value otherwise.

### GetRotationPeriodOk

`func (o *CreateUserKey) GetRotationPeriodOk() (*int32, bool)`

GetRotationPeriodOk returns a tuple with the RotationPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationPeriod

`func (o *CreateUserKey) SetRotationPeriod(v int32)`

SetRotationPeriod sets RotationPeriod field to given value.

### HasRotationPeriod

`func (o *CreateUserKey) HasRotationPeriod() bool`

HasRotationPeriod returns a boolean if a field has been set.

### SetRotationPeriodNil

`func (o *CreateUserKey) SetRotationPeriodNil(b bool)`

 SetRotationPeriodNil sets the value for RotationPeriod to be an explicit nil

### UnsetRotationPeriod
`func (o *CreateUserKey) UnsetRotationPeriod()`

UnsetRotationPeriod ensures that no value is present for RotationPeriod, not even an explicit nil
### GetIsPendingActivation

`func (o *CreateUserKey) GetIsPendingActivation() bool`

GetIsPendingActivation returns the IsPendingActivation field if non-nil, zero value otherwise.

### GetIsPendingActivationOk

`func (o *CreateUserKey) GetIsPendingActivationOk() (*bool, bool)`

GetIsPendingActivationOk returns a tuple with the IsPendingActivation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPendingActivation

`func (o *CreateUserKey) SetIsPendingActivation(v bool)`

SetIsPendingActivation sets IsPendingActivation field to given value.

### HasIsPendingActivation

`func (o *CreateUserKey) HasIsPendingActivation() bool`

HasIsPendingActivation returns a boolean if a field has been set.

### SetIsPendingActivationNil

`func (o *CreateUserKey) SetIsPendingActivationNil(b bool)`

 SetIsPendingActivationNil sets the value for IsPendingActivation to be an explicit nil

### UnsetIsPendingActivation
`func (o *CreateUserKey) UnsetIsPendingActivation()`

UnsetIsPendingActivation ensures that no value is present for IsPendingActivation, not even an explicit nil
### GetIsAccessControlEnabled

`func (o *CreateUserKey) GetIsAccessControlEnabled() bool`

GetIsAccessControlEnabled returns the IsAccessControlEnabled field if non-nil, zero value otherwise.

### GetIsAccessControlEnabledOk

`func (o *CreateUserKey) GetIsAccessControlEnabledOk() (*bool, bool)`

GetIsAccessControlEnabledOk returns a tuple with the IsAccessControlEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAccessControlEnabled

`func (o *CreateUserKey) SetIsAccessControlEnabled(v bool)`

SetIsAccessControlEnabled sets IsAccessControlEnabled field to given value.


### GetAccessTargetIds

`func (o *CreateUserKey) GetAccessTargetIds() []string`

GetAccessTargetIds returns the AccessTargetIds field if non-nil, zero value otherwise.

### GetAccessTargetIdsOk

`func (o *CreateUserKey) GetAccessTargetIdsOk() (*[]string, bool)`

GetAccessTargetIdsOk returns a tuple with the AccessTargetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTargetIds

`func (o *CreateUserKey) SetAccessTargetIds(v []string)`

SetAccessTargetIds sets AccessTargetIds field to given value.

### HasAccessTargetIds

`func (o *CreateUserKey) HasAccessTargetIds() bool`

HasAccessTargetIds returns a boolean if a field has been set.

### SetAccessTargetIdsNil

`func (o *CreateUserKey) SetAccessTargetIdsNil(b bool)`

 SetAccessTargetIdsNil sets the value for AccessTargetIds to be an explicit nil

### UnsetAccessTargetIds
`func (o *CreateUserKey) UnsetAccessTargetIds()`

UnsetAccessTargetIds ensures that no value is present for AccessTargetIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


