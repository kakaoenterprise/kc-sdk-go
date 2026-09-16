# NodePoolScript

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserData** | Pointer to **NullableString** | 설정된 사용자 스크립트 (base64 인코딩된 값) | [optional] 
**Name** | **string** | 노드 풀 이름 | 
**Id** | **string** | 노드 풀의 고유 ID | 

## Methods

### NewNodePoolScript

`func NewNodePoolScript(name string, id string, ) *NodePoolScript`

NewNodePoolScript instantiates a new NodePoolScript object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodePoolScriptWithDefaults

`func NewNodePoolScriptWithDefaults() *NodePoolScript`

NewNodePoolScriptWithDefaults instantiates a new NodePoolScript object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserData

`func (o *NodePoolScript) GetUserData() string`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *NodePoolScript) GetUserDataOk() (*string, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *NodePoolScript) SetUserData(v string)`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *NodePoolScript) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### SetUserDataNil

`func (o *NodePoolScript) SetUserDataNil(b bool)`

 SetUserDataNil sets the value for UserData to be an explicit nil

### UnsetUserData
`func (o *NodePoolScript) UnsetUserData()`

UnsetUserData ensures that no value is present for UserData, not even an explicit nil
### GetName

`func (o *NodePoolScript) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NodePoolScript) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NodePoolScript) SetName(v string)`

SetName sets Name field to given value.


### GetId

`func (o *NodePoolScript) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NodePoolScript) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NodePoolScript) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


