# NodeUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodepoolName** | **string** | 노드 풀 이름 | 
**NodeCount** | **int32** | 노드 풀에서 사용 중인 노드 수 | 
**IsBm** | **bool** | 베어메탈 노드인지 여부 | 

## Methods

### NewNodeUsage

`func NewNodeUsage(nodepoolName string, nodeCount int32, isBm bool, ) *NodeUsage`

NewNodeUsage instantiates a new NodeUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeUsageWithDefaults

`func NewNodeUsageWithDefaults() *NodeUsage`

NewNodeUsageWithDefaults instantiates a new NodeUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodepoolName

`func (o *NodeUsage) GetNodepoolName() string`

GetNodepoolName returns the NodepoolName field if non-nil, zero value otherwise.

### GetNodepoolNameOk

`func (o *NodeUsage) GetNodepoolNameOk() (*string, bool)`

GetNodepoolNameOk returns a tuple with the NodepoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodepoolName

`func (o *NodeUsage) SetNodepoolName(v string)`

SetNodepoolName sets NodepoolName field to given value.


### GetNodeCount

`func (o *NodeUsage) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *NodeUsage) GetNodeCountOk() (*int32, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *NodeUsage) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.


### GetIsBm

`func (o *NodeUsage) GetIsBm() bool`

GetIsBm returns the IsBm field if non-nil, zero value otherwise.

### GetIsBmOk

`func (o *NodeUsage) GetIsBmOk() (*bool, bool)`

GetIsBmOk returns a tuple with the IsBm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBm

`func (o *NodeUsage) SetIsBm(v bool)`

SetIsBm sets IsBm field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


