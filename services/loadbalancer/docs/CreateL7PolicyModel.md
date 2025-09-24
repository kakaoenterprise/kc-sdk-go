# CreateL7PolicyModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Action** | [**L7PolicyAction**](L7PolicyAction.md) | 정책 동작 유형 &lt;br/&gt; - &#x60;REDIRECT_PREFIX&#x60;: Prefix로 리디렉션 &lt;br/&gt; - &#x60;REDIRECT_TO_POOL&#x60;: 대상 풀로 리디렉션 &lt;br/&gt; - &#x60;REDIRECT_TO_URL&#x60;: URL로 리디렉션 | 
**Description** | Pointer to **NullableString** |  | [optional] 
**ListenerId** | **string** | 정책을 적용할 대상 리스너 ID | 
**Position** | Pointer to **NullableInt32** |  | [optional] 
**RedirectTargetGroupId** | Pointer to **NullableString** |  | [optional] 
**RedirectPrefix** | Pointer to **NullableString** |  | [optional] 
**RedirectUrl** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreateL7PolicyModel

`func NewCreateL7PolicyModel(action L7PolicyAction, listenerId string, ) *CreateL7PolicyModel`

NewCreateL7PolicyModel instantiates a new CreateL7PolicyModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateL7PolicyModelWithDefaults

`func NewCreateL7PolicyModelWithDefaults() *CreateL7PolicyModel`

NewCreateL7PolicyModelWithDefaults instantiates a new CreateL7PolicyModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateL7PolicyModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateL7PolicyModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateL7PolicyModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateL7PolicyModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreateL7PolicyModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateL7PolicyModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAction

`func (o *CreateL7PolicyModel) GetAction() L7PolicyAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *CreateL7PolicyModel) GetActionOk() (*L7PolicyAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *CreateL7PolicyModel) SetAction(v L7PolicyAction)`

SetAction sets Action field to given value.


### GetDescription

`func (o *CreateL7PolicyModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateL7PolicyModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateL7PolicyModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateL7PolicyModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateL7PolicyModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateL7PolicyModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetListenerId

`func (o *CreateL7PolicyModel) GetListenerId() string`

GetListenerId returns the ListenerId field if non-nil, zero value otherwise.

### GetListenerIdOk

`func (o *CreateL7PolicyModel) GetListenerIdOk() (*string, bool)`

GetListenerIdOk returns a tuple with the ListenerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenerId

`func (o *CreateL7PolicyModel) SetListenerId(v string)`

SetListenerId sets ListenerId field to given value.


### GetPosition

`func (o *CreateL7PolicyModel) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *CreateL7PolicyModel) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *CreateL7PolicyModel) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *CreateL7PolicyModel) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### SetPositionNil

`func (o *CreateL7PolicyModel) SetPositionNil(b bool)`

 SetPositionNil sets the value for Position to be an explicit nil

### UnsetPosition
`func (o *CreateL7PolicyModel) UnsetPosition()`

UnsetPosition ensures that no value is present for Position, not even an explicit nil
### GetRedirectTargetGroupId

`func (o *CreateL7PolicyModel) GetRedirectTargetGroupId() string`

GetRedirectTargetGroupId returns the RedirectTargetGroupId field if non-nil, zero value otherwise.

### GetRedirectTargetGroupIdOk

`func (o *CreateL7PolicyModel) GetRedirectTargetGroupIdOk() (*string, bool)`

GetRedirectTargetGroupIdOk returns a tuple with the RedirectTargetGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectTargetGroupId

`func (o *CreateL7PolicyModel) SetRedirectTargetGroupId(v string)`

SetRedirectTargetGroupId sets RedirectTargetGroupId field to given value.

### HasRedirectTargetGroupId

`func (o *CreateL7PolicyModel) HasRedirectTargetGroupId() bool`

HasRedirectTargetGroupId returns a boolean if a field has been set.

### SetRedirectTargetGroupIdNil

`func (o *CreateL7PolicyModel) SetRedirectTargetGroupIdNil(b bool)`

 SetRedirectTargetGroupIdNil sets the value for RedirectTargetGroupId to be an explicit nil

### UnsetRedirectTargetGroupId
`func (o *CreateL7PolicyModel) UnsetRedirectTargetGroupId()`

UnsetRedirectTargetGroupId ensures that no value is present for RedirectTargetGroupId, not even an explicit nil
### GetRedirectPrefix

`func (o *CreateL7PolicyModel) GetRedirectPrefix() string`

GetRedirectPrefix returns the RedirectPrefix field if non-nil, zero value otherwise.

### GetRedirectPrefixOk

`func (o *CreateL7PolicyModel) GetRedirectPrefixOk() (*string, bool)`

GetRedirectPrefixOk returns a tuple with the RedirectPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectPrefix

`func (o *CreateL7PolicyModel) SetRedirectPrefix(v string)`

SetRedirectPrefix sets RedirectPrefix field to given value.

### HasRedirectPrefix

`func (o *CreateL7PolicyModel) HasRedirectPrefix() bool`

HasRedirectPrefix returns a boolean if a field has been set.

### SetRedirectPrefixNil

`func (o *CreateL7PolicyModel) SetRedirectPrefixNil(b bool)`

 SetRedirectPrefixNil sets the value for RedirectPrefix to be an explicit nil

### UnsetRedirectPrefix
`func (o *CreateL7PolicyModel) UnsetRedirectPrefix()`

UnsetRedirectPrefix ensures that no value is present for RedirectPrefix, not even an explicit nil
### GetRedirectUrl

`func (o *CreateL7PolicyModel) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *CreateL7PolicyModel) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *CreateL7PolicyModel) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *CreateL7PolicyModel) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### SetRedirectUrlNil

`func (o *CreateL7PolicyModel) SetRedirectUrlNil(b bool)`

 SetRedirectUrlNil sets the value for RedirectUrl to be an explicit nil

### UnsetRedirectUrl
`func (o *CreateL7PolicyModel) UnsetRedirectUrl()`

UnsetRedirectUrl ensures that no value is present for RedirectUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


