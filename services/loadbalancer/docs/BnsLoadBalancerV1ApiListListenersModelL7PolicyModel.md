# BnsLoadBalancerV1ApiListListenersModelL7PolicyModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | L7 정책의 ID | 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ProvisioningStatus** | Pointer to [**NullableProvisioningStatus**](ProvisioningStatus.md) |  | [optional] 
**OperatingStatus** | Pointer to [**NullableLoadBalancerOperatingStatus**](LoadBalancerOperatingStatus.md) |  | [optional] 
**ProjectId** | Pointer to **NullableString** |  | [optional] 
**Action** | Pointer to [**NullableL7PolicyAction**](L7PolicyAction.md) |  | [optional] 
**Position** | Pointer to **NullableInt32** |  | [optional] 
**Rules** | Pointer to [**[]BnsLoadBalancerV1ApiListListenersModelRuleModel**](BnsLoadBalancerV1ApiListListenersModelRuleModel.md) |  | [optional] 
**RedirectTargetGroupId** | Pointer to **NullableString** |  | [optional] 
**RedirectUrl** | Pointer to **NullableString** |  | [optional] 
**RedirectPrefix** | Pointer to **NullableString** |  | [optional] 
**RedirectHttpCode** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewBnsLoadBalancerV1ApiListListenersModelL7PolicyModel

`func NewBnsLoadBalancerV1ApiListListenersModelL7PolicyModel(id string, ) *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel`

NewBnsLoadBalancerV1ApiListListenersModelL7PolicyModel instantiates a new BnsLoadBalancerV1ApiListListenersModelL7PolicyModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsLoadBalancerV1ApiListListenersModelL7PolicyModelWithDefaults

`func NewBnsLoadBalancerV1ApiListListenersModelL7PolicyModelWithDefaults() *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel`

NewBnsLoadBalancerV1ApiListListenersModelL7PolicyModelWithDefaults instantiates a new BnsLoadBalancerV1ApiListListenersModelL7PolicyModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetProvisioningStatus

`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) GetProvisioningStatus() ProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) GetProvisioningStatusOk() (*ProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) SetProvisioningStatus(v ProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.

### HasProvisioningStatus

`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) HasProvisioningStatus() bool`

HasProvisioningStatus returns a boolean if a field has been set.

### SetProvisioningStatusNil

`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) SetProvisioningStatusNil(b bool)`

 SetProvisioningStatusNil sets the value for ProvisioningStatus to be an explicit nil

### UnsetProvisioningStatus
`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) UnsetProvisioningStatus()`

UnsetProvisioningStatus ensures that no value is present for ProvisioningStatus, not even an explicit nil
### GetOperatingStatus

`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) GetOperatingStatus() LoadBalancerOperatingStatus`

GetOperatingStatus returns the OperatingStatus field if non-nil, zero value otherwise.

### GetOperatingStatusOk

`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) GetOperatingStatusOk() (*LoadBalancerOperatingStatus, bool)`

GetOperatingStatusOk returns a tuple with the OperatingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingStatus

`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) SetOperatingStatus(v LoadBalancerOperatingStatus)`

SetOperatingStatus sets OperatingStatus field to given value.

### HasOperatingStatus

`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) HasOperatingStatus() bool`

HasOperatingStatus returns a boolean if a field has been set.

### SetOperatingStatusNil

`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) SetOperatingStatusNil(b bool)`

 SetOperatingStatusNil sets the value for OperatingStatus to be an explicit nil

### UnsetOperatingStatus
`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) UnsetOperatingStatus()`

UnsetOperatingStatus ensures that no value is present for OperatingStatus, not even an explicit nil
### GetProjectId

`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetAction

`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) GetAction() L7PolicyAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) GetActionOk() (*L7PolicyAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) SetAction(v L7PolicyAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetActionNil

`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetPosition

`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### SetPositionNil

`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) SetPositionNil(b bool)`

 SetPositionNil sets the value for Position to be an explicit nil

### UnsetPosition
`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) UnsetPosition()`

UnsetPosition ensures that no value is present for Position, not even an explicit nil
### GetRules

`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) GetRules() []BnsLoadBalancerV1ApiListListenersModelRuleModel`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) GetRulesOk() (*[]BnsLoadBalancerV1ApiListListenersModelRuleModel, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) SetRules(v []BnsLoadBalancerV1ApiListListenersModelRuleModel)`

SetRules sets Rules field to given value.

### HasRules

`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) HasRules() bool`

HasRules returns a boolean if a field has been set.

### SetRulesNil

`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) SetRulesNil(b bool)`

 SetRulesNil sets the value for Rules to be an explicit nil

### UnsetRules
`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) UnsetRules()`

UnsetRules ensures that no value is present for Rules, not even an explicit nil
### GetRedirectTargetGroupId

`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) GetRedirectTargetGroupId() string`

GetRedirectTargetGroupId returns the RedirectTargetGroupId field if non-nil, zero value otherwise.

### GetRedirectTargetGroupIdOk

`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) GetRedirectTargetGroupIdOk() (*string, bool)`

GetRedirectTargetGroupIdOk returns a tuple with the RedirectTargetGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectTargetGroupId

`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) SetRedirectTargetGroupId(v string)`

SetRedirectTargetGroupId sets RedirectTargetGroupId field to given value.

### HasRedirectTargetGroupId

`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) HasRedirectTargetGroupId() bool`

HasRedirectTargetGroupId returns a boolean if a field has been set.

### SetRedirectTargetGroupIdNil

`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) SetRedirectTargetGroupIdNil(b bool)`

 SetRedirectTargetGroupIdNil sets the value for RedirectTargetGroupId to be an explicit nil

### UnsetRedirectTargetGroupId
`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) UnsetRedirectTargetGroupId()`

UnsetRedirectTargetGroupId ensures that no value is present for RedirectTargetGroupId, not even an explicit nil
### GetRedirectUrl

`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### SetRedirectUrlNil

`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) SetRedirectUrlNil(b bool)`

 SetRedirectUrlNil sets the value for RedirectUrl to be an explicit nil

### UnsetRedirectUrl
`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) UnsetRedirectUrl()`

UnsetRedirectUrl ensures that no value is present for RedirectUrl, not even an explicit nil
### GetRedirectPrefix

`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) GetRedirectPrefix() string`

GetRedirectPrefix returns the RedirectPrefix field if non-nil, zero value otherwise.

### GetRedirectPrefixOk

`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) GetRedirectPrefixOk() (*string, bool)`

GetRedirectPrefixOk returns a tuple with the RedirectPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectPrefix

`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) SetRedirectPrefix(v string)`

SetRedirectPrefix sets RedirectPrefix field to given value.

### HasRedirectPrefix

`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) HasRedirectPrefix() bool`

HasRedirectPrefix returns a boolean if a field has been set.

### SetRedirectPrefixNil

`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) SetRedirectPrefixNil(b bool)`

 SetRedirectPrefixNil sets the value for RedirectPrefix to be an explicit nil

### UnsetRedirectPrefix
`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) UnsetRedirectPrefix()`

UnsetRedirectPrefix ensures that no value is present for RedirectPrefix, not even an explicit nil
### GetRedirectHttpCode

`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) GetRedirectHttpCode() int32`

GetRedirectHttpCode returns the RedirectHttpCode field if non-nil, zero value otherwise.

### GetRedirectHttpCodeOk

`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) GetRedirectHttpCodeOk() (*int32, bool)`

GetRedirectHttpCodeOk returns a tuple with the RedirectHttpCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectHttpCode

`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) SetRedirectHttpCode(v int32)`

SetRedirectHttpCode sets RedirectHttpCode field to given value.

### HasRedirectHttpCode

`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) HasRedirectHttpCode() bool`

HasRedirectHttpCode returns a boolean if a field has been set.

### SetRedirectHttpCodeNil

`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) SetRedirectHttpCodeNil(b bool)`

 SetRedirectHttpCodeNil sets the value for RedirectHttpCode to be an explicit nil

### UnsetRedirectHttpCode
`func (o *BnsLoadBalancerV1ApiListListenersModelL7PolicyModel) UnsetRedirectHttpCode()`

UnsetRedirectHttpCode ensures that no value is present for RedirectHttpCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


