# NodePoolStatusInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailableNodes** | **int32** | 사용 가능한 노드 수 | 
**Phase** | [**NodePoolStatus**](NodePoolStatus.md) | 노드 풀의 현재 상태 | 
**UnavailableNodes** | **int32** | 사용 불가 노드 수 | 

## Methods

### NewNodePoolStatusInfo

`func NewNodePoolStatusInfo(availableNodes int32, phase NodePoolStatus, unavailableNodes int32, ) *NodePoolStatusInfo`

NewNodePoolStatusInfo instantiates a new NodePoolStatusInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodePoolStatusInfoWithDefaults

`func NewNodePoolStatusInfoWithDefaults() *NodePoolStatusInfo`

NewNodePoolStatusInfoWithDefaults instantiates a new NodePoolStatusInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailableNodes

`func (o *NodePoolStatusInfo) GetAvailableNodes() int32`

GetAvailableNodes returns the AvailableNodes field if non-nil, zero value otherwise.

### GetAvailableNodesOk

`func (o *NodePoolStatusInfo) GetAvailableNodesOk() (*int32, bool)`

GetAvailableNodesOk returns a tuple with the AvailableNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableNodes

`func (o *NodePoolStatusInfo) SetAvailableNodes(v int32)`

SetAvailableNodes sets AvailableNodes field to given value.


### GetPhase

`func (o *NodePoolStatusInfo) GetPhase() NodePoolStatus`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *NodePoolStatusInfo) GetPhaseOk() (*NodePoolStatus, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *NodePoolStatusInfo) SetPhase(v NodePoolStatus)`

SetPhase sets Phase field to given value.


### GetUnavailableNodes

`func (o *NodePoolStatusInfo) GetUnavailableNodes() int32`

GetUnavailableNodes returns the UnavailableNodes field if non-nil, zero value otherwise.

### GetUnavailableNodesOk

`func (o *NodePoolStatusInfo) GetUnavailableNodesOk() (*int32, bool)`

GetUnavailableNodesOk returns a tuple with the UnavailableNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnavailableNodes

`func (o *NodePoolStatusInfo) SetUnavailableNodes(v int32)`

SetUnavailableNodes sets UnavailableNodes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


