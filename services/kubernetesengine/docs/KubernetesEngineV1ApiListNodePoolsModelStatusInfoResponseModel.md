# KubernetesEngineV1ApiListNodePoolsModelStatusInfoResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailableNodes** | **int32** | 사용 가능한 노드 수 | 
**Phase** | [**NodePoolStatus**](NodePoolStatus.md) | 노드 풀의 현재 상태 &lt;br/&gt;- &#x60;Pending&#x60;: 대기 중 &lt;br/&gt;- &#x60;ScalingUp&#x60;: 확장 중 &lt;br/&gt;- &#x60;ScalingDown&#x60;: 축소 중 &lt;br/&gt;- &#x60;Running&#x60;: 실행 중 &lt;br/&gt;- &#x60;Failed&#x60;: 오류 발생 &lt;br/&gt;- &#x60;Deleting&#x60;: 삭제 중 &lt;br/&gt;- &#x60;Updating&#x60;: 업데이트 중 &lt;br/&gt;- &#x60;Running (Scheduling Disable)&#x60;: 실행 중이나 스케줄링 불가 | 
**UnavailableNodes** | **int32** | 사용 불가 노드 수 | 

## Methods

### NewKubernetesEngineV1ApiListNodePoolsModelStatusInfoResponseModel

`func NewKubernetesEngineV1ApiListNodePoolsModelStatusInfoResponseModel(availableNodes int32, phase NodePoolStatus, unavailableNodes int32, ) *KubernetesEngineV1ApiListNodePoolsModelStatusInfoResponseModel`

NewKubernetesEngineV1ApiListNodePoolsModelStatusInfoResponseModel instantiates a new KubernetesEngineV1ApiListNodePoolsModelStatusInfoResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesEngineV1ApiListNodePoolsModelStatusInfoResponseModelWithDefaults

`func NewKubernetesEngineV1ApiListNodePoolsModelStatusInfoResponseModelWithDefaults() *KubernetesEngineV1ApiListNodePoolsModelStatusInfoResponseModel`

NewKubernetesEngineV1ApiListNodePoolsModelStatusInfoResponseModelWithDefaults instantiates a new KubernetesEngineV1ApiListNodePoolsModelStatusInfoResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailableNodes

`func (o *KubernetesEngineV1ApiListNodePoolsModelStatusInfoResponseModel) GetAvailableNodes() int32`

GetAvailableNodes returns the AvailableNodes field if non-nil, zero value otherwise.

### GetAvailableNodesOk

`func (o *KubernetesEngineV1ApiListNodePoolsModelStatusInfoResponseModel) GetAvailableNodesOk() (*int32, bool)`

GetAvailableNodesOk returns a tuple with the AvailableNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableNodes

`func (o *KubernetesEngineV1ApiListNodePoolsModelStatusInfoResponseModel) SetAvailableNodes(v int32)`

SetAvailableNodes sets AvailableNodes field to given value.


### GetPhase

`func (o *KubernetesEngineV1ApiListNodePoolsModelStatusInfoResponseModel) GetPhase() NodePoolStatus`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *KubernetesEngineV1ApiListNodePoolsModelStatusInfoResponseModel) GetPhaseOk() (*NodePoolStatus, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *KubernetesEngineV1ApiListNodePoolsModelStatusInfoResponseModel) SetPhase(v NodePoolStatus)`

SetPhase sets Phase field to given value.


### GetUnavailableNodes

`func (o *KubernetesEngineV1ApiListNodePoolsModelStatusInfoResponseModel) GetUnavailableNodes() int32`

GetUnavailableNodes returns the UnavailableNodes field if non-nil, zero value otherwise.

### GetUnavailableNodesOk

`func (o *KubernetesEngineV1ApiListNodePoolsModelStatusInfoResponseModel) GetUnavailableNodesOk() (*int32, bool)`

GetUnavailableNodesOk returns a tuple with the UnavailableNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnavailableNodes

`func (o *KubernetesEngineV1ApiListNodePoolsModelStatusInfoResponseModel) SetUnavailableNodes(v int32)`

SetUnavailableNodes sets UnavailableNodes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


