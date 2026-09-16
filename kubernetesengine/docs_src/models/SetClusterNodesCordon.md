# SetClusterNodesCordon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsCordon** | **bool** | 노드를 cordon 상태로 설정할지 여부 - &#x60;true&#x60;: 해당 노드를 cordon 상태로 설정하여 신규 파드 스케줄링을 막음 - &#x60;false&#x60;: 노드를 cordon 상태로 설정하지 않음 (기본값) | 
**NodeNames** | **[]string** | 대상 노드 이름 목록 | 

## Methods

### NewSetClusterNodesCordon

`func NewSetClusterNodesCordon(isCordon bool, nodeNames []string, ) *SetClusterNodesCordon`

NewSetClusterNodesCordon instantiates a new SetClusterNodesCordon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetClusterNodesCordonWithDefaults

`func NewSetClusterNodesCordonWithDefaults() *SetClusterNodesCordon`

NewSetClusterNodesCordonWithDefaults instantiates a new SetClusterNodesCordon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsCordon

`func (o *SetClusterNodesCordon) GetIsCordon() bool`

GetIsCordon returns the IsCordon field if non-nil, zero value otherwise.

### GetIsCordonOk

`func (o *SetClusterNodesCordon) GetIsCordonOk() (*bool, bool)`

GetIsCordonOk returns a tuple with the IsCordon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCordon

`func (o *SetClusterNodesCordon) SetIsCordon(v bool)`

SetIsCordon sets IsCordon field to given value.


### GetNodeNames

`func (o *SetClusterNodesCordon) GetNodeNames() []string`

GetNodeNames returns the NodeNames field if non-nil, zero value otherwise.

### GetNodeNamesOk

`func (o *SetClusterNodesCordon) GetNodeNamesOk() (*[]string, bool)`

GetNodeNamesOk returns a tuple with the NodeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeNames

`func (o *SetClusterNodesCordon) SetNodeNames(v []string)`

SetNodeNames sets NodeNames field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


