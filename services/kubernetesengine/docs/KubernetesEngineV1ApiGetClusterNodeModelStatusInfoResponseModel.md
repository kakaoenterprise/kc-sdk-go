# KubernetesEngineV1ApiGetClusterNodeModelStatusInfoResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phase** | [**NodeStatus**](NodeStatus.md) | 노드의 현재 상태 &lt;br/&gt;- &#x60;Running&#x60;: 정상 실행 중 &lt;br/&gt;- &#x60;Running (Scheduling Disable)&#x60;: 실행 중이나 스케줄링 불가 &lt;br/&gt;- &#x60;Provisioned&#x60;: 생성 완료 &lt;br/&gt;- &#x60;Deleted&#x60;: 삭제됨 &lt;br/&gt;- &#x60;Pending&#x60;: 대기 중 &lt;br/&gt;- &#x60;Provisioning&#x60;: 생성 중 &lt;br/&gt;- &#x60;Deleting&#x60;: 삭제 중 &lt;br/&gt;- &#x60;Failed&#x60;: 오류 발생 | 

## Methods

### NewKubernetesEngineV1ApiGetClusterNodeModelStatusInfoResponseModel

`func NewKubernetesEngineV1ApiGetClusterNodeModelStatusInfoResponseModel(phase NodeStatus, ) *KubernetesEngineV1ApiGetClusterNodeModelStatusInfoResponseModel`

NewKubernetesEngineV1ApiGetClusterNodeModelStatusInfoResponseModel instantiates a new KubernetesEngineV1ApiGetClusterNodeModelStatusInfoResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesEngineV1ApiGetClusterNodeModelStatusInfoResponseModelWithDefaults

`func NewKubernetesEngineV1ApiGetClusterNodeModelStatusInfoResponseModelWithDefaults() *KubernetesEngineV1ApiGetClusterNodeModelStatusInfoResponseModel`

NewKubernetesEngineV1ApiGetClusterNodeModelStatusInfoResponseModelWithDefaults instantiates a new KubernetesEngineV1ApiGetClusterNodeModelStatusInfoResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhase

`func (o *KubernetesEngineV1ApiGetClusterNodeModelStatusInfoResponseModel) GetPhase() NodeStatus`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *KubernetesEngineV1ApiGetClusterNodeModelStatusInfoResponseModel) GetPhaseOk() (*NodeStatus, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *KubernetesEngineV1ApiGetClusterNodeModelStatusInfoResponseModel) SetPhase(v NodeStatus)`

SetPhase sets Phase field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


