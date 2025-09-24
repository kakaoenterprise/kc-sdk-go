# BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 생성된 정책의 ID | 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ProvisioningStatus** | Pointer to [**NullableProvisioningStatus**](ProvisioningStatus.md) |  | [optional] 
**OperatingStatus** | Pointer to [**NullableLoadBalancerOperatingStatus**](LoadBalancerOperatingStatus.md) |  | [optional] 
**ProjectId** | Pointer to **NullableString** |  | [optional] 
**Action** | [**L7PolicyAction**](L7PolicyAction.md) | 정책 동작 유형 | 
**ListenerId** | **string** | 정책을 적용할 대상 리스너 ID | 
**RedirectTargetGroupId** | Pointer to **NullableString** |  | [optional] 
**RedirectUrl** | Pointer to **NullableString** |  | [optional] 
**RedirectPrefix** | Pointer to **NullableString** |  | [optional] 
**Position** | Pointer to **NullableInt32** |  | [optional] 
**Rules** | Pointer to [**[]BnsLoadBalancerV1ApiCreateL7PolicyModelRuleModel**](BnsLoadBalancerV1ApiCreateL7PolicyModelRuleModel.md) |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewBnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel

`func NewBnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel(id string, action L7PolicyAction, listenerId string, ) *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel`

NewBnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel instantiates a new BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModelWithDefaults

`func NewBnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModelWithDefaults() *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel`

NewBnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModelWithDefaults instantiates a new BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetProvisioningStatus

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) GetProvisioningStatus() ProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) GetProvisioningStatusOk() (*ProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) SetProvisioningStatus(v ProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.

### HasProvisioningStatus

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) HasProvisioningStatus() bool`

HasProvisioningStatus returns a boolean if a field has been set.

### SetProvisioningStatusNil

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) SetProvisioningStatusNil(b bool)`

 SetProvisioningStatusNil sets the value for ProvisioningStatus to be an explicit nil

### UnsetProvisioningStatus
`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) UnsetProvisioningStatus()`

UnsetProvisioningStatus ensures that no value is present for ProvisioningStatus, not even an explicit nil
### GetOperatingStatus

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) GetOperatingStatus() LoadBalancerOperatingStatus`

GetOperatingStatus returns the OperatingStatus field if non-nil, zero value otherwise.

### GetOperatingStatusOk

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) GetOperatingStatusOk() (*LoadBalancerOperatingStatus, bool)`

GetOperatingStatusOk returns a tuple with the OperatingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingStatus

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) SetOperatingStatus(v LoadBalancerOperatingStatus)`

SetOperatingStatus sets OperatingStatus field to given value.

### HasOperatingStatus

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) HasOperatingStatus() bool`

HasOperatingStatus returns a boolean if a field has been set.

### SetOperatingStatusNil

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) SetOperatingStatusNil(b bool)`

 SetOperatingStatusNil sets the value for OperatingStatus to be an explicit nil

### UnsetOperatingStatus
`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) UnsetOperatingStatus()`

UnsetOperatingStatus ensures that no value is present for OperatingStatus, not even an explicit nil
### GetProjectId

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetAction

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) GetAction() L7PolicyAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) GetActionOk() (*L7PolicyAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) SetAction(v L7PolicyAction)`

SetAction sets Action field to given value.


### GetListenerId

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) GetListenerId() string`

GetListenerId returns the ListenerId field if non-nil, zero value otherwise.

### GetListenerIdOk

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) GetListenerIdOk() (*string, bool)`

GetListenerIdOk returns a tuple with the ListenerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenerId

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) SetListenerId(v string)`

SetListenerId sets ListenerId field to given value.


### GetRedirectTargetGroupId

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) GetRedirectTargetGroupId() string`

GetRedirectTargetGroupId returns the RedirectTargetGroupId field if non-nil, zero value otherwise.

### GetRedirectTargetGroupIdOk

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) GetRedirectTargetGroupIdOk() (*string, bool)`

GetRedirectTargetGroupIdOk returns a tuple with the RedirectTargetGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectTargetGroupId

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) SetRedirectTargetGroupId(v string)`

SetRedirectTargetGroupId sets RedirectTargetGroupId field to given value.

### HasRedirectTargetGroupId

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) HasRedirectTargetGroupId() bool`

HasRedirectTargetGroupId returns a boolean if a field has been set.

### SetRedirectTargetGroupIdNil

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) SetRedirectTargetGroupIdNil(b bool)`

 SetRedirectTargetGroupIdNil sets the value for RedirectTargetGroupId to be an explicit nil

### UnsetRedirectTargetGroupId
`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) UnsetRedirectTargetGroupId()`

UnsetRedirectTargetGroupId ensures that no value is present for RedirectTargetGroupId, not even an explicit nil
### GetRedirectUrl

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### SetRedirectUrlNil

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) SetRedirectUrlNil(b bool)`

 SetRedirectUrlNil sets the value for RedirectUrl to be an explicit nil

### UnsetRedirectUrl
`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) UnsetRedirectUrl()`

UnsetRedirectUrl ensures that no value is present for RedirectUrl, not even an explicit nil
### GetRedirectPrefix

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) GetRedirectPrefix() string`

GetRedirectPrefix returns the RedirectPrefix field if non-nil, zero value otherwise.

### GetRedirectPrefixOk

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) GetRedirectPrefixOk() (*string, bool)`

GetRedirectPrefixOk returns a tuple with the RedirectPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectPrefix

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) SetRedirectPrefix(v string)`

SetRedirectPrefix sets RedirectPrefix field to given value.

### HasRedirectPrefix

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) HasRedirectPrefix() bool`

HasRedirectPrefix returns a boolean if a field has been set.

### SetRedirectPrefixNil

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) SetRedirectPrefixNil(b bool)`

 SetRedirectPrefixNil sets the value for RedirectPrefix to be an explicit nil

### UnsetRedirectPrefix
`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) UnsetRedirectPrefix()`

UnsetRedirectPrefix ensures that no value is present for RedirectPrefix, not even an explicit nil
### GetPosition

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### SetPositionNil

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) SetPositionNil(b bool)`

 SetPositionNil sets the value for Position to be an explicit nil

### UnsetPosition
`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) UnsetPosition()`

UnsetPosition ensures that no value is present for Position, not even an explicit nil
### GetRules

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) GetRules() []BnsLoadBalancerV1ApiCreateL7PolicyModelRuleModel`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) GetRulesOk() (*[]BnsLoadBalancerV1ApiCreateL7PolicyModelRuleModel, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) SetRules(v []BnsLoadBalancerV1ApiCreateL7PolicyModelRuleModel)`

SetRules sets Rules field to given value.

### HasRules

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) HasRules() bool`

HasRules returns a boolean if a field has been set.

### SetRulesNil

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) SetRulesNil(b bool)`

 SetRulesNil sets the value for Rules to be an explicit nil

### UnsetRules
`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) UnsetRules()`

UnsetRules ensures that no value is present for Rules, not even an explicit nil
### GetCreatedAt

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *BnsLoadBalancerV1ApiCreateL7PolicyModelL7PolicyModel) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


