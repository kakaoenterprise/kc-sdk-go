# BnsLoadBalancerV1ApiGetL7PolicyModelRuleModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | L7 규칙 ID | 
**CompareType** | Pointer to [**NullableL7RuleCompareType**](L7RuleCompareType.md) |  | [optional] 
**IsInverted** | **bool** | 규칙 반전 여부&lt;br/&gt;- &#x60;true&#x60;인 경우 비교 규칙의 반대 결과가 일치로 간주됨 | 
**Key** | Pointer to **NullableString** |  | [optional] 
**Value** | Pointer to **NullableString** |  | [optional] 
**ProvisioningStatus** | Pointer to [**NullableProvisioningStatus**](ProvisioningStatus.md) |  | [optional] 
**OperatingStatus** | Pointer to [**NullableLoadBalancerOperatingStatus**](LoadBalancerOperatingStatus.md) |  | [optional] 
**ProjectId** | **string** | 프로젝트 ID | 
**Type** | Pointer to [**NullableL7RuleType**](L7RuleType.md) |  | [optional] 

## Methods

### NewBnsLoadBalancerV1ApiGetL7PolicyModelRuleModel

`func NewBnsLoadBalancerV1ApiGetL7PolicyModelRuleModel(id string, isInverted bool, projectId string, ) *BnsLoadBalancerV1ApiGetL7PolicyModelRuleModel`

NewBnsLoadBalancerV1ApiGetL7PolicyModelRuleModel instantiates a new BnsLoadBalancerV1ApiGetL7PolicyModelRuleModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsLoadBalancerV1ApiGetL7PolicyModelRuleModelWithDefaults

`func NewBnsLoadBalancerV1ApiGetL7PolicyModelRuleModelWithDefaults() *BnsLoadBalancerV1ApiGetL7PolicyModelRuleModel`

NewBnsLoadBalancerV1ApiGetL7PolicyModelRuleModelWithDefaults instantiates a new BnsLoadBalancerV1ApiGetL7PolicyModelRuleModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsLoadBalancerV1ApiGetL7PolicyModelRuleModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsLoadBalancerV1ApiGetL7PolicyModelRuleModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsLoadBalancerV1ApiGetL7PolicyModelRuleModel) SetId(v string)`

SetId sets Id field to given value.


### GetCompareType

`func (o *BnsLoadBalancerV1ApiGetL7PolicyModelRuleModel) GetCompareType() L7RuleCompareType`

GetCompareType returns the CompareType field if non-nil, zero value otherwise.

### GetCompareTypeOk

`func (o *BnsLoadBalancerV1ApiGetL7PolicyModelRuleModel) GetCompareTypeOk() (*L7RuleCompareType, bool)`

GetCompareTypeOk returns a tuple with the CompareType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompareType

`func (o *BnsLoadBalancerV1ApiGetL7PolicyModelRuleModel) SetCompareType(v L7RuleCompareType)`

SetCompareType sets CompareType field to given value.

### HasCompareType

`func (o *BnsLoadBalancerV1ApiGetL7PolicyModelRuleModel) HasCompareType() bool`

HasCompareType returns a boolean if a field has been set.

### SetCompareTypeNil

`func (o *BnsLoadBalancerV1ApiGetL7PolicyModelRuleModel) SetCompareTypeNil(b bool)`

 SetCompareTypeNil sets the value for CompareType to be an explicit nil

### UnsetCompareType
`func (o *BnsLoadBalancerV1ApiGetL7PolicyModelRuleModel) UnsetCompareType()`

UnsetCompareType ensures that no value is present for CompareType, not even an explicit nil
### GetIsInverted

`func (o *BnsLoadBalancerV1ApiGetL7PolicyModelRuleModel) GetIsInverted() bool`

GetIsInverted returns the IsInverted field if non-nil, zero value otherwise.

### GetIsInvertedOk

`func (o *BnsLoadBalancerV1ApiGetL7PolicyModelRuleModel) GetIsInvertedOk() (*bool, bool)`

GetIsInvertedOk returns a tuple with the IsInverted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInverted

`func (o *BnsLoadBalancerV1ApiGetL7PolicyModelRuleModel) SetIsInverted(v bool)`

SetIsInverted sets IsInverted field to given value.


### GetKey

`func (o *BnsLoadBalancerV1ApiGetL7PolicyModelRuleModel) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *BnsLoadBalancerV1ApiGetL7PolicyModelRuleModel) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *BnsLoadBalancerV1ApiGetL7PolicyModelRuleModel) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *BnsLoadBalancerV1ApiGetL7PolicyModelRuleModel) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *BnsLoadBalancerV1ApiGetL7PolicyModelRuleModel) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *BnsLoadBalancerV1ApiGetL7PolicyModelRuleModel) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetValue

`func (o *BnsLoadBalancerV1ApiGetL7PolicyModelRuleModel) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BnsLoadBalancerV1ApiGetL7PolicyModelRuleModel) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BnsLoadBalancerV1ApiGetL7PolicyModelRuleModel) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *BnsLoadBalancerV1ApiGetL7PolicyModelRuleModel) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *BnsLoadBalancerV1ApiGetL7PolicyModelRuleModel) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *BnsLoadBalancerV1ApiGetL7PolicyModelRuleModel) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetProvisioningStatus

`func (o *BnsLoadBalancerV1ApiGetL7PolicyModelRuleModel) GetProvisioningStatus() ProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *BnsLoadBalancerV1ApiGetL7PolicyModelRuleModel) GetProvisioningStatusOk() (*ProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *BnsLoadBalancerV1ApiGetL7PolicyModelRuleModel) SetProvisioningStatus(v ProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.

### HasProvisioningStatus

`func (o *BnsLoadBalancerV1ApiGetL7PolicyModelRuleModel) HasProvisioningStatus() bool`

HasProvisioningStatus returns a boolean if a field has been set.

### SetProvisioningStatusNil

`func (o *BnsLoadBalancerV1ApiGetL7PolicyModelRuleModel) SetProvisioningStatusNil(b bool)`

 SetProvisioningStatusNil sets the value for ProvisioningStatus to be an explicit nil

### UnsetProvisioningStatus
`func (o *BnsLoadBalancerV1ApiGetL7PolicyModelRuleModel) UnsetProvisioningStatus()`

UnsetProvisioningStatus ensures that no value is present for ProvisioningStatus, not even an explicit nil
### GetOperatingStatus

`func (o *BnsLoadBalancerV1ApiGetL7PolicyModelRuleModel) GetOperatingStatus() LoadBalancerOperatingStatus`

GetOperatingStatus returns the OperatingStatus field if non-nil, zero value otherwise.

### GetOperatingStatusOk

`func (o *BnsLoadBalancerV1ApiGetL7PolicyModelRuleModel) GetOperatingStatusOk() (*LoadBalancerOperatingStatus, bool)`

GetOperatingStatusOk returns a tuple with the OperatingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingStatus

`func (o *BnsLoadBalancerV1ApiGetL7PolicyModelRuleModel) SetOperatingStatus(v LoadBalancerOperatingStatus)`

SetOperatingStatus sets OperatingStatus field to given value.

### HasOperatingStatus

`func (o *BnsLoadBalancerV1ApiGetL7PolicyModelRuleModel) HasOperatingStatus() bool`

HasOperatingStatus returns a boolean if a field has been set.

### SetOperatingStatusNil

`func (o *BnsLoadBalancerV1ApiGetL7PolicyModelRuleModel) SetOperatingStatusNil(b bool)`

 SetOperatingStatusNil sets the value for OperatingStatus to be an explicit nil

### UnsetOperatingStatus
`func (o *BnsLoadBalancerV1ApiGetL7PolicyModelRuleModel) UnsetOperatingStatus()`

UnsetOperatingStatus ensures that no value is present for OperatingStatus, not even an explicit nil
### GetProjectId

`func (o *BnsLoadBalancerV1ApiGetL7PolicyModelRuleModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BnsLoadBalancerV1ApiGetL7PolicyModelRuleModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BnsLoadBalancerV1ApiGetL7PolicyModelRuleModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetType

`func (o *BnsLoadBalancerV1ApiGetL7PolicyModelRuleModel) GetType() L7RuleType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BnsLoadBalancerV1ApiGetL7PolicyModelRuleModel) GetTypeOk() (*L7RuleType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BnsLoadBalancerV1ApiGetL7PolicyModelRuleModel) SetType(v L7RuleType)`

SetType sets Type field to given value.

### HasType

`func (o *BnsLoadBalancerV1ApiGetL7PolicyModelRuleModel) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *BnsLoadBalancerV1ApiGetL7PolicyModelRuleModel) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *BnsLoadBalancerV1ApiGetL7PolicyModelRuleModel) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


