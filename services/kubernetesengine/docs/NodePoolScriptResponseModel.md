# NodePoolScriptResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserData** | **string** | 설정된 사용자 스크립트 (base64 인코딩된 값) | 
**Name** | **string** | 노드 풀 이름 | 
**Id** | **string** | 노드 풀의 고유 ID | 

## Methods

### NewNodePoolScriptResponseModel

`func NewNodePoolScriptResponseModel(userData string, name string, id string, ) *NodePoolScriptResponseModel`

NewNodePoolScriptResponseModel instantiates a new NodePoolScriptResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodePoolScriptResponseModelWithDefaults

`func NewNodePoolScriptResponseModelWithDefaults() *NodePoolScriptResponseModel`

NewNodePoolScriptResponseModelWithDefaults instantiates a new NodePoolScriptResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserData

`func (o *NodePoolScriptResponseModel) GetUserData() string`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *NodePoolScriptResponseModel) GetUserDataOk() (*string, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *NodePoolScriptResponseModel) SetUserData(v string)`

SetUserData sets UserData field to given value.


### GetName

`func (o *NodePoolScriptResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NodePoolScriptResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NodePoolScriptResponseModel) SetName(v string)`

SetName sets Name field to given value.


### GetId

`func (o *NodePoolScriptResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NodePoolScriptResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NodePoolScriptResponseModel) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


