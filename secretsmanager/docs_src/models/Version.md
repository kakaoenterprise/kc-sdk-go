# Version

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **int32** | 시크릿 버전 | 
**Status** | [**Status**](Status.md) | 시크릿 버전 상태 | 
**KmsKeyName** | Pointer to **NullableString** | 시크릿 보호에 사용하는 KMS 키 이름 | [optional] 
**KmsKeyVersion** | Pointer to **NullableInt32** | 시크릿 암호화에 사용된 KMS 키 버전 | [optional] 
**CreatedAt** | **string** | 버전이 생성된 일시 | 
**ResolvedDestructionAt** | Pointer to **NullableString** | 시크릿 버전이 실제로 폐기된 일시 | [optional] 

## Methods

### NewVersion

`func NewVersion(version int32, status Status, createdAt string, ) *Version`

NewVersion instantiates a new Version object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionWithDefaults

`func NewVersionWithDefaults() *Version`

NewVersionWithDefaults instantiates a new Version object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *Version) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Version) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Version) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetStatus

`func (o *Version) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Version) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Version) SetStatus(v Status)`

SetStatus sets Status field to given value.


### GetKmsKeyName

`func (o *Version) GetKmsKeyName() string`

GetKmsKeyName returns the KmsKeyName field if non-nil, zero value otherwise.

### GetKmsKeyNameOk

`func (o *Version) GetKmsKeyNameOk() (*string, bool)`

GetKmsKeyNameOk returns a tuple with the KmsKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmsKeyName

`func (o *Version) SetKmsKeyName(v string)`

SetKmsKeyName sets KmsKeyName field to given value.

### HasKmsKeyName

`func (o *Version) HasKmsKeyName() bool`

HasKmsKeyName returns a boolean if a field has been set.

### SetKmsKeyNameNil

`func (o *Version) SetKmsKeyNameNil(b bool)`

 SetKmsKeyNameNil sets the value for KmsKeyName to be an explicit nil

### UnsetKmsKeyName
`func (o *Version) UnsetKmsKeyName()`

UnsetKmsKeyName ensures that no value is present for KmsKeyName, not even an explicit nil
### GetKmsKeyVersion

`func (o *Version) GetKmsKeyVersion() int32`

GetKmsKeyVersion returns the KmsKeyVersion field if non-nil, zero value otherwise.

### GetKmsKeyVersionOk

`func (o *Version) GetKmsKeyVersionOk() (*int32, bool)`

GetKmsKeyVersionOk returns a tuple with the KmsKeyVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmsKeyVersion

`func (o *Version) SetKmsKeyVersion(v int32)`

SetKmsKeyVersion sets KmsKeyVersion field to given value.

### HasKmsKeyVersion

`func (o *Version) HasKmsKeyVersion() bool`

HasKmsKeyVersion returns a boolean if a field has been set.

### SetKmsKeyVersionNil

`func (o *Version) SetKmsKeyVersionNil(b bool)`

 SetKmsKeyVersionNil sets the value for KmsKeyVersion to be an explicit nil

### UnsetKmsKeyVersion
`func (o *Version) UnsetKmsKeyVersion()`

UnsetKmsKeyVersion ensures that no value is present for KmsKeyVersion, not even an explicit nil
### GetCreatedAt

`func (o *Version) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Version) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Version) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetResolvedDestructionAt

`func (o *Version) GetResolvedDestructionAt() string`

GetResolvedDestructionAt returns the ResolvedDestructionAt field if non-nil, zero value otherwise.

### GetResolvedDestructionAtOk

`func (o *Version) GetResolvedDestructionAtOk() (*string, bool)`

GetResolvedDestructionAtOk returns a tuple with the ResolvedDestructionAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedDestructionAt

`func (o *Version) SetResolvedDestructionAt(v string)`

SetResolvedDestructionAt sets ResolvedDestructionAt field to given value.

### HasResolvedDestructionAt

`func (o *Version) HasResolvedDestructionAt() bool`

HasResolvedDestructionAt returns a boolean if a field has been set.

### SetResolvedDestructionAtNil

`func (o *Version) SetResolvedDestructionAtNil(b bool)`

 SetResolvedDestructionAtNil sets the value for ResolvedDestructionAt to be an explicit nil

### UnsetResolvedDestructionAt
`func (o *Version) UnsetResolvedDestructionAt()`

UnsetResolvedDestructionAt ensures that no value is present for ResolvedDestructionAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


