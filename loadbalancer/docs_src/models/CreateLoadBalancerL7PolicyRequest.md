# CreateLoadBalancerL7PolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | L7 정책 이름 | [optional] 
**Action** | [**L7PolicyAction**](L7PolicyAction.md) | 트래픽 처리 방식 | 
**Description** | Pointer to **NullableString** | L7 정책 설명 | [optional] 
**Position** | Pointer to **NullableInt32** | 정책 적용 순서 | [optional] 
**RedirectTargetGroup** | Pointer to [**NullablePoolRefRequest**](PoolRefRequest.md) | 리다이렉트 대상 그룹 | [optional] 
**RedirectPrefix** | Pointer to **NullableString** | 리다이렉트 경로 prefix | [optional] 
**RedirectUrl** | Pointer to **NullableString** | 리다이렉트 URL | [optional] 

## Methods

### NewCreateLoadBalancerL7PolicyRequest

`func NewCreateLoadBalancerL7PolicyRequest(action L7PolicyAction, ) *CreateLoadBalancerL7PolicyRequest`

NewCreateLoadBalancerL7PolicyRequest instantiates a new CreateLoadBalancerL7PolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLoadBalancerL7PolicyRequestWithDefaults

`func NewCreateLoadBalancerL7PolicyRequestWithDefaults() *CreateLoadBalancerL7PolicyRequest`

NewCreateLoadBalancerL7PolicyRequestWithDefaults instantiates a new CreateLoadBalancerL7PolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateLoadBalancerL7PolicyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateLoadBalancerL7PolicyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateLoadBalancerL7PolicyRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateLoadBalancerL7PolicyRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreateLoadBalancerL7PolicyRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateLoadBalancerL7PolicyRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAction

`func (o *CreateLoadBalancerL7PolicyRequest) GetAction() L7PolicyAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *CreateLoadBalancerL7PolicyRequest) GetActionOk() (*L7PolicyAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *CreateLoadBalancerL7PolicyRequest) SetAction(v L7PolicyAction)`

SetAction sets Action field to given value.


### GetDescription

`func (o *CreateLoadBalancerL7PolicyRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateLoadBalancerL7PolicyRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateLoadBalancerL7PolicyRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateLoadBalancerL7PolicyRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateLoadBalancerL7PolicyRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateLoadBalancerL7PolicyRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPosition

`func (o *CreateLoadBalancerL7PolicyRequest) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *CreateLoadBalancerL7PolicyRequest) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *CreateLoadBalancerL7PolicyRequest) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *CreateLoadBalancerL7PolicyRequest) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### SetPositionNil

`func (o *CreateLoadBalancerL7PolicyRequest) SetPositionNil(b bool)`

 SetPositionNil sets the value for Position to be an explicit nil

### UnsetPosition
`func (o *CreateLoadBalancerL7PolicyRequest) UnsetPosition()`

UnsetPosition ensures that no value is present for Position, not even an explicit nil
### GetRedirectTargetGroup

`func (o *CreateLoadBalancerL7PolicyRequest) GetRedirectTargetGroup() PoolRefRequest`

GetRedirectTargetGroup returns the RedirectTargetGroup field if non-nil, zero value otherwise.

### GetRedirectTargetGroupOk

`func (o *CreateLoadBalancerL7PolicyRequest) GetRedirectTargetGroupOk() (*PoolRefRequest, bool)`

GetRedirectTargetGroupOk returns a tuple with the RedirectTargetGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectTargetGroup

`func (o *CreateLoadBalancerL7PolicyRequest) SetRedirectTargetGroup(v PoolRefRequest)`

SetRedirectTargetGroup sets RedirectTargetGroup field to given value.

### HasRedirectTargetGroup

`func (o *CreateLoadBalancerL7PolicyRequest) HasRedirectTargetGroup() bool`

HasRedirectTargetGroup returns a boolean if a field has been set.

### SetRedirectTargetGroupNil

`func (o *CreateLoadBalancerL7PolicyRequest) SetRedirectTargetGroupNil(b bool)`

 SetRedirectTargetGroupNil sets the value for RedirectTargetGroup to be an explicit nil

### UnsetRedirectTargetGroup
`func (o *CreateLoadBalancerL7PolicyRequest) UnsetRedirectTargetGroup()`

UnsetRedirectTargetGroup ensures that no value is present for RedirectTargetGroup, not even an explicit nil
### GetRedirectPrefix

`func (o *CreateLoadBalancerL7PolicyRequest) GetRedirectPrefix() string`

GetRedirectPrefix returns the RedirectPrefix field if non-nil, zero value otherwise.

### GetRedirectPrefixOk

`func (o *CreateLoadBalancerL7PolicyRequest) GetRedirectPrefixOk() (*string, bool)`

GetRedirectPrefixOk returns a tuple with the RedirectPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectPrefix

`func (o *CreateLoadBalancerL7PolicyRequest) SetRedirectPrefix(v string)`

SetRedirectPrefix sets RedirectPrefix field to given value.

### HasRedirectPrefix

`func (o *CreateLoadBalancerL7PolicyRequest) HasRedirectPrefix() bool`

HasRedirectPrefix returns a boolean if a field has been set.

### SetRedirectPrefixNil

`func (o *CreateLoadBalancerL7PolicyRequest) SetRedirectPrefixNil(b bool)`

 SetRedirectPrefixNil sets the value for RedirectPrefix to be an explicit nil

### UnsetRedirectPrefix
`func (o *CreateLoadBalancerL7PolicyRequest) UnsetRedirectPrefix()`

UnsetRedirectPrefix ensures that no value is present for RedirectPrefix, not even an explicit nil
### GetRedirectUrl

`func (o *CreateLoadBalancerL7PolicyRequest) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *CreateLoadBalancerL7PolicyRequest) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *CreateLoadBalancerL7PolicyRequest) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *CreateLoadBalancerL7PolicyRequest) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### SetRedirectUrlNil

`func (o *CreateLoadBalancerL7PolicyRequest) SetRedirectUrlNil(b bool)`

 SetRedirectUrlNil sets the value for RedirectUrl to be an explicit nil

### UnsetRedirectUrl
`func (o *CreateLoadBalancerL7PolicyRequest) UnsetRedirectUrl()`

UnsetRedirectUrl ensures that no value is present for RedirectUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


