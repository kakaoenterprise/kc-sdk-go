# CreatedL7Policy2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | L7 정책 ID | 
**Name** | **string** | L7 정책 이름 | 
**Description** | Pointer to **NullableString** | L7 정책 설명 | [optional] 
**ProvisioningStatus** | Pointer to [**NullableProvisioningStatus**](ProvisioningStatus.md) | 정책 구성 상태 | [optional] 
**OperatingStatus** | Pointer to [**NullableLoadBalancerOperatingStatus**](LoadBalancerOperatingStatus.md) | 정책 운영 상태 | [optional] 
**ProjectId** | Pointer to **NullableString** | 프로젝트 ID | [optional] 
**Action** | [**L7PolicyAction**](L7PolicyAction.md) | 트래픽 처리 방식 | 
**ListenerId** | Pointer to **NullableString** | 연결된 리스너 ID | [optional] 
**Position** | **int32** | 정책 적용 순서 | 
**RedirectTargetGroupId** | Pointer to **NullableString** | 리다이렉트 대상 그룹 ID | [optional] 
**RedirectUrl** | Pointer to **NullableString** | 리다이렉트 URL | [optional] 
**RedirectPrefix** | Pointer to **NullableString** | 리다이렉트 경로 | [optional] 
**RedirectHttpCode** | Pointer to **NullableInt32** | 리다이렉트 HTTP 상태 코드 | [optional] 

## Methods

### NewCreatedL7Policy2

`func NewCreatedL7Policy2(id string, name string, action L7PolicyAction, position int32, ) *CreatedL7Policy2`

NewCreatedL7Policy2 instantiates a new CreatedL7Policy2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatedL7Policy2WithDefaults

`func NewCreatedL7Policy2WithDefaults() *CreatedL7Policy2`

NewCreatedL7Policy2WithDefaults instantiates a new CreatedL7Policy2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreatedL7Policy2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreatedL7Policy2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreatedL7Policy2) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *CreatedL7Policy2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatedL7Policy2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatedL7Policy2) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreatedL7Policy2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreatedL7Policy2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreatedL7Policy2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreatedL7Policy2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreatedL7Policy2) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreatedL7Policy2) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetProvisioningStatus

`func (o *CreatedL7Policy2) GetProvisioningStatus() ProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *CreatedL7Policy2) GetProvisioningStatusOk() (*ProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *CreatedL7Policy2) SetProvisioningStatus(v ProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.

### HasProvisioningStatus

`func (o *CreatedL7Policy2) HasProvisioningStatus() bool`

HasProvisioningStatus returns a boolean if a field has been set.

### SetProvisioningStatusNil

`func (o *CreatedL7Policy2) SetProvisioningStatusNil(b bool)`

 SetProvisioningStatusNil sets the value for ProvisioningStatus to be an explicit nil

### UnsetProvisioningStatus
`func (o *CreatedL7Policy2) UnsetProvisioningStatus()`

UnsetProvisioningStatus ensures that no value is present for ProvisioningStatus, not even an explicit nil
### GetOperatingStatus

`func (o *CreatedL7Policy2) GetOperatingStatus() LoadBalancerOperatingStatus`

GetOperatingStatus returns the OperatingStatus field if non-nil, zero value otherwise.

### GetOperatingStatusOk

`func (o *CreatedL7Policy2) GetOperatingStatusOk() (*LoadBalancerOperatingStatus, bool)`

GetOperatingStatusOk returns a tuple with the OperatingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingStatus

`func (o *CreatedL7Policy2) SetOperatingStatus(v LoadBalancerOperatingStatus)`

SetOperatingStatus sets OperatingStatus field to given value.

### HasOperatingStatus

`func (o *CreatedL7Policy2) HasOperatingStatus() bool`

HasOperatingStatus returns a boolean if a field has been set.

### SetOperatingStatusNil

`func (o *CreatedL7Policy2) SetOperatingStatusNil(b bool)`

 SetOperatingStatusNil sets the value for OperatingStatus to be an explicit nil

### UnsetOperatingStatus
`func (o *CreatedL7Policy2) UnsetOperatingStatus()`

UnsetOperatingStatus ensures that no value is present for OperatingStatus, not even an explicit nil
### GetProjectId

`func (o *CreatedL7Policy2) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CreatedL7Policy2) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CreatedL7Policy2) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *CreatedL7Policy2) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *CreatedL7Policy2) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *CreatedL7Policy2) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetAction

`func (o *CreatedL7Policy2) GetAction() L7PolicyAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *CreatedL7Policy2) GetActionOk() (*L7PolicyAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *CreatedL7Policy2) SetAction(v L7PolicyAction)`

SetAction sets Action field to given value.


### GetListenerId

`func (o *CreatedL7Policy2) GetListenerId() string`

GetListenerId returns the ListenerId field if non-nil, zero value otherwise.

### GetListenerIdOk

`func (o *CreatedL7Policy2) GetListenerIdOk() (*string, bool)`

GetListenerIdOk returns a tuple with the ListenerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenerId

`func (o *CreatedL7Policy2) SetListenerId(v string)`

SetListenerId sets ListenerId field to given value.

### HasListenerId

`func (o *CreatedL7Policy2) HasListenerId() bool`

HasListenerId returns a boolean if a field has been set.

### SetListenerIdNil

`func (o *CreatedL7Policy2) SetListenerIdNil(b bool)`

 SetListenerIdNil sets the value for ListenerId to be an explicit nil

### UnsetListenerId
`func (o *CreatedL7Policy2) UnsetListenerId()`

UnsetListenerId ensures that no value is present for ListenerId, not even an explicit nil
### GetPosition

`func (o *CreatedL7Policy2) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *CreatedL7Policy2) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *CreatedL7Policy2) SetPosition(v int32)`

SetPosition sets Position field to given value.


### GetRedirectTargetGroupId

`func (o *CreatedL7Policy2) GetRedirectTargetGroupId() string`

GetRedirectTargetGroupId returns the RedirectTargetGroupId field if non-nil, zero value otherwise.

### GetRedirectTargetGroupIdOk

`func (o *CreatedL7Policy2) GetRedirectTargetGroupIdOk() (*string, bool)`

GetRedirectTargetGroupIdOk returns a tuple with the RedirectTargetGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectTargetGroupId

`func (o *CreatedL7Policy2) SetRedirectTargetGroupId(v string)`

SetRedirectTargetGroupId sets RedirectTargetGroupId field to given value.

### HasRedirectTargetGroupId

`func (o *CreatedL7Policy2) HasRedirectTargetGroupId() bool`

HasRedirectTargetGroupId returns a boolean if a field has been set.

### SetRedirectTargetGroupIdNil

`func (o *CreatedL7Policy2) SetRedirectTargetGroupIdNil(b bool)`

 SetRedirectTargetGroupIdNil sets the value for RedirectTargetGroupId to be an explicit nil

### UnsetRedirectTargetGroupId
`func (o *CreatedL7Policy2) UnsetRedirectTargetGroupId()`

UnsetRedirectTargetGroupId ensures that no value is present for RedirectTargetGroupId, not even an explicit nil
### GetRedirectUrl

`func (o *CreatedL7Policy2) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *CreatedL7Policy2) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *CreatedL7Policy2) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *CreatedL7Policy2) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### SetRedirectUrlNil

`func (o *CreatedL7Policy2) SetRedirectUrlNil(b bool)`

 SetRedirectUrlNil sets the value for RedirectUrl to be an explicit nil

### UnsetRedirectUrl
`func (o *CreatedL7Policy2) UnsetRedirectUrl()`

UnsetRedirectUrl ensures that no value is present for RedirectUrl, not even an explicit nil
### GetRedirectPrefix

`func (o *CreatedL7Policy2) GetRedirectPrefix() string`

GetRedirectPrefix returns the RedirectPrefix field if non-nil, zero value otherwise.

### GetRedirectPrefixOk

`func (o *CreatedL7Policy2) GetRedirectPrefixOk() (*string, bool)`

GetRedirectPrefixOk returns a tuple with the RedirectPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectPrefix

`func (o *CreatedL7Policy2) SetRedirectPrefix(v string)`

SetRedirectPrefix sets RedirectPrefix field to given value.

### HasRedirectPrefix

`func (o *CreatedL7Policy2) HasRedirectPrefix() bool`

HasRedirectPrefix returns a boolean if a field has been set.

### SetRedirectPrefixNil

`func (o *CreatedL7Policy2) SetRedirectPrefixNil(b bool)`

 SetRedirectPrefixNil sets the value for RedirectPrefix to be an explicit nil

### UnsetRedirectPrefix
`func (o *CreatedL7Policy2) UnsetRedirectPrefix()`

UnsetRedirectPrefix ensures that no value is present for RedirectPrefix, not even an explicit nil
### GetRedirectHttpCode

`func (o *CreatedL7Policy2) GetRedirectHttpCode() int32`

GetRedirectHttpCode returns the RedirectHttpCode field if non-nil, zero value otherwise.

### GetRedirectHttpCodeOk

`func (o *CreatedL7Policy2) GetRedirectHttpCodeOk() (*int32, bool)`

GetRedirectHttpCodeOk returns a tuple with the RedirectHttpCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectHttpCode

`func (o *CreatedL7Policy2) SetRedirectHttpCode(v int32)`

SetRedirectHttpCode sets RedirectHttpCode field to given value.

### HasRedirectHttpCode

`func (o *CreatedL7Policy2) HasRedirectHttpCode() bool`

HasRedirectHttpCode returns a boolean if a field has been set.

### SetRedirectHttpCodeNil

`func (o *CreatedL7Policy2) SetRedirectHttpCodeNil(b bool)`

 SetRedirectHttpCodeNil sets the value for RedirectHttpCode to be an explicit nil

### UnsetRedirectHttpCode
`func (o *CreatedL7Policy2) UnsetRedirectHttpCode()`

UnsetRedirectHttpCode ensures that no value is present for RedirectHttpCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


