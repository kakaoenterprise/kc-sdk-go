# UpdateUserKeyAccessControl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | 접근 제어 활성화 여부 | 
**TargetIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewUpdateUserKeyAccessControl

`func NewUpdateUserKeyAccessControl(isEnabled bool, ) *UpdateUserKeyAccessControl`

NewUpdateUserKeyAccessControl instantiates a new UpdateUserKeyAccessControl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserKeyAccessControlWithDefaults

`func NewUpdateUserKeyAccessControlWithDefaults() *UpdateUserKeyAccessControl`

NewUpdateUserKeyAccessControlWithDefaults instantiates a new UpdateUserKeyAccessControl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsEnabled

`func (o *UpdateUserKeyAccessControl) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *UpdateUserKeyAccessControl) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *UpdateUserKeyAccessControl) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.


### GetTargetIds

`func (o *UpdateUserKeyAccessControl) GetTargetIds() []string`

GetTargetIds returns the TargetIds field if non-nil, zero value otherwise.

### GetTargetIdsOk

`func (o *UpdateUserKeyAccessControl) GetTargetIdsOk() (*[]string, bool)`

GetTargetIdsOk returns a tuple with the TargetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetIds

`func (o *UpdateUserKeyAccessControl) SetTargetIds(v []string)`

SetTargetIds sets TargetIds field to given value.

### HasTargetIds

`func (o *UpdateUserKeyAccessControl) HasTargetIds() bool`

HasTargetIds returns a boolean if a field has been set.

### SetTargetIdsNil

`func (o *UpdateUserKeyAccessControl) SetTargetIdsNil(b bool)`

 SetTargetIdsNil sets the value for TargetIds to be an explicit nil

### UnsetTargetIds
`func (o *UpdateUserKeyAccessControl) UnsetTargetIds()`

UnsetTargetIds ensures that no value is present for TargetIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


