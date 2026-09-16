# OsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **NullableString** | 운영체제 유형 | [optional] 
**Distro** | Pointer to **NullableString** | 배포판 이름 | [optional] 
**Architecture** | Pointer to **NullableString** | 운영체제 아키텍처 | [optional] 
**AdminUser** | Pointer to **NullableString** | 이미지의 기본 관리자 계정명 | [optional] 
**IsHidden** | Pointer to **NullableBool** | 이미지 숨김 여부 | [optional] 

## Methods

### NewOsInfo

`func NewOsInfo() *OsInfo`

NewOsInfo instantiates a new OsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsInfoWithDefaults

`func NewOsInfoWithDefaults() *OsInfo`

NewOsInfoWithDefaults instantiates a new OsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OsInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OsInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OsInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OsInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *OsInfo) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *OsInfo) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetDistro

`func (o *OsInfo) GetDistro() string`

GetDistro returns the Distro field if non-nil, zero value otherwise.

### GetDistroOk

`func (o *OsInfo) GetDistroOk() (*string, bool)`

GetDistroOk returns a tuple with the Distro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistro

`func (o *OsInfo) SetDistro(v string)`

SetDistro sets Distro field to given value.

### HasDistro

`func (o *OsInfo) HasDistro() bool`

HasDistro returns a boolean if a field has been set.

### SetDistroNil

`func (o *OsInfo) SetDistroNil(b bool)`

 SetDistroNil sets the value for Distro to be an explicit nil

### UnsetDistro
`func (o *OsInfo) UnsetDistro()`

UnsetDistro ensures that no value is present for Distro, not even an explicit nil
### GetArchitecture

`func (o *OsInfo) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *OsInfo) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *OsInfo) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.

### HasArchitecture

`func (o *OsInfo) HasArchitecture() bool`

HasArchitecture returns a boolean if a field has been set.

### SetArchitectureNil

`func (o *OsInfo) SetArchitectureNil(b bool)`

 SetArchitectureNil sets the value for Architecture to be an explicit nil

### UnsetArchitecture
`func (o *OsInfo) UnsetArchitecture()`

UnsetArchitecture ensures that no value is present for Architecture, not even an explicit nil
### GetAdminUser

`func (o *OsInfo) GetAdminUser() string`

GetAdminUser returns the AdminUser field if non-nil, zero value otherwise.

### GetAdminUserOk

`func (o *OsInfo) GetAdminUserOk() (*string, bool)`

GetAdminUserOk returns a tuple with the AdminUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminUser

`func (o *OsInfo) SetAdminUser(v string)`

SetAdminUser sets AdminUser field to given value.

### HasAdminUser

`func (o *OsInfo) HasAdminUser() bool`

HasAdminUser returns a boolean if a field has been set.

### SetAdminUserNil

`func (o *OsInfo) SetAdminUserNil(b bool)`

 SetAdminUserNil sets the value for AdminUser to be an explicit nil

### UnsetAdminUser
`func (o *OsInfo) UnsetAdminUser()`

UnsetAdminUser ensures that no value is present for AdminUser, not even an explicit nil
### GetIsHidden

`func (o *OsInfo) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *OsInfo) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *OsInfo) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.

### HasIsHidden

`func (o *OsInfo) HasIsHidden() bool`

HasIsHidden returns a boolean if a field has been set.

### SetIsHiddenNil

`func (o *OsInfo) SetIsHiddenNil(b bool)`

 SetIsHiddenNil sets the value for IsHidden to be an explicit nil

### UnsetIsHidden
`func (o *OsInfo) UnsetIsHidden()`

UnsetIsHidden ensures that no value is present for IsHidden, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


