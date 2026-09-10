# SetNodePoolUserScript

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserData** | **string** | 노드 풀 내 노드 생성 시 실행할 사용자 스크립트(base64 인코딩 필요) | 

## Methods

### NewSetNodePoolUserScript

`func NewSetNodePoolUserScript(userData string, ) *SetNodePoolUserScript`

NewSetNodePoolUserScript instantiates a new SetNodePoolUserScript object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetNodePoolUserScriptWithDefaults

`func NewSetNodePoolUserScriptWithDefaults() *SetNodePoolUserScript`

NewSetNodePoolUserScriptWithDefaults instantiates a new SetNodePoolUserScript object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserData

`func (o *SetNodePoolUserScript) GetUserData() string`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *SetNodePoolUserScript) GetUserDataOk() (*string, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *SetNodePoolUserScript) SetUserData(v string)`

SetUserData sets UserData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


