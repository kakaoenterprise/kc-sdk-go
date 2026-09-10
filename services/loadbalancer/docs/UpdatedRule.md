# UpdatedRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompareType** | Pointer to [**NullableL7RuleCompareType**](L7RuleCompareType.md) |  | [optional] 
**Id** | Pointer to **NullableString** |  | [optional] 
**IsInverted** | **bool** | 규칙 반전 여부&lt;br/&gt;- &#x60;true&#x60;인 경우 비교 규칙의 반대 결과가 일치로 간주됨 | 
**Key** | Pointer to **NullableString** |  | [optional] 
**Value** | Pointer to **NullableString** |  | [optional] 
**ProvisioningStatus** | Pointer to [**NullableProvisioningStatus**](ProvisioningStatus.md) |  | [optional] 
**OperatingStatus** | Pointer to [**NullableLoadBalancerOperatingStatus**](LoadBalancerOperatingStatus.md) |  | [optional] 
**ProjectId** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to [**NullableL7RuleType**](L7RuleType.md) |  | [optional] 

## Methods

### NewUpdatedRule

`func NewUpdatedRule(isInverted bool, ) *UpdatedRule`

NewUpdatedRule instantiates a new UpdatedRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatedRuleWithDefaults

`func NewUpdatedRuleWithDefaults() *UpdatedRule`

NewUpdatedRuleWithDefaults instantiates a new UpdatedRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompareType

`func (o *UpdatedRule) GetCompareType() L7RuleCompareType`

GetCompareType returns the CompareType field if non-nil, zero value otherwise.

### GetCompareTypeOk

`func (o *UpdatedRule) GetCompareTypeOk() (*L7RuleCompareType, bool)`

GetCompareTypeOk returns a tuple with the CompareType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompareType

`func (o *UpdatedRule) SetCompareType(v L7RuleCompareType)`

SetCompareType sets CompareType field to given value.

### HasCompareType

`func (o *UpdatedRule) HasCompareType() bool`

HasCompareType returns a boolean if a field has been set.

### SetCompareTypeNil

`func (o *UpdatedRule) SetCompareTypeNil(b bool)`

 SetCompareTypeNil sets the value for CompareType to be an explicit nil

### UnsetCompareType
`func (o *UpdatedRule) UnsetCompareType()`

UnsetCompareType ensures that no value is present for CompareType, not even an explicit nil
### GetId

`func (o *UpdatedRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdatedRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdatedRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UpdatedRule) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *UpdatedRule) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *UpdatedRule) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIsInverted

`func (o *UpdatedRule) GetIsInverted() bool`

GetIsInverted returns the IsInverted field if non-nil, zero value otherwise.

### GetIsInvertedOk

`func (o *UpdatedRule) GetIsInvertedOk() (*bool, bool)`

GetIsInvertedOk returns a tuple with the IsInverted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInverted

`func (o *UpdatedRule) SetIsInverted(v bool)`

SetIsInverted sets IsInverted field to given value.


### GetKey

`func (o *UpdatedRule) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UpdatedRule) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UpdatedRule) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *UpdatedRule) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *UpdatedRule) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *UpdatedRule) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetValue

`func (o *UpdatedRule) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UpdatedRule) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UpdatedRule) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *UpdatedRule) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *UpdatedRule) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *UpdatedRule) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetProvisioningStatus

`func (o *UpdatedRule) GetProvisioningStatus() ProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *UpdatedRule) GetProvisioningStatusOk() (*ProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *UpdatedRule) SetProvisioningStatus(v ProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.

### HasProvisioningStatus

`func (o *UpdatedRule) HasProvisioningStatus() bool`

HasProvisioningStatus returns a boolean if a field has been set.

### SetProvisioningStatusNil

`func (o *UpdatedRule) SetProvisioningStatusNil(b bool)`

 SetProvisioningStatusNil sets the value for ProvisioningStatus to be an explicit nil

### UnsetProvisioningStatus
`func (o *UpdatedRule) UnsetProvisioningStatus()`

UnsetProvisioningStatus ensures that no value is present for ProvisioningStatus, not even an explicit nil
### GetOperatingStatus

`func (o *UpdatedRule) GetOperatingStatus() LoadBalancerOperatingStatus`

GetOperatingStatus returns the OperatingStatus field if non-nil, zero value otherwise.

### GetOperatingStatusOk

`func (o *UpdatedRule) GetOperatingStatusOk() (*LoadBalancerOperatingStatus, bool)`

GetOperatingStatusOk returns a tuple with the OperatingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingStatus

`func (o *UpdatedRule) SetOperatingStatus(v LoadBalancerOperatingStatus)`

SetOperatingStatus sets OperatingStatus field to given value.

### HasOperatingStatus

`func (o *UpdatedRule) HasOperatingStatus() bool`

HasOperatingStatus returns a boolean if a field has been set.

### SetOperatingStatusNil

`func (o *UpdatedRule) SetOperatingStatusNil(b bool)`

 SetOperatingStatusNil sets the value for OperatingStatus to be an explicit nil

### UnsetOperatingStatus
`func (o *UpdatedRule) UnsetOperatingStatus()`

UnsetOperatingStatus ensures that no value is present for OperatingStatus, not even an explicit nil
### GetProjectId

`func (o *UpdatedRule) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *UpdatedRule) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *UpdatedRule) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *UpdatedRule) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *UpdatedRule) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *UpdatedRule) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetType

`func (o *UpdatedRule) GetType() L7RuleType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdatedRule) GetTypeOk() (*L7RuleType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdatedRule) SetType(v L7RuleType)`

SetType sets Type field to given value.

### HasType

`func (o *UpdatedRule) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *UpdatedRule) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *UpdatedRule) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


