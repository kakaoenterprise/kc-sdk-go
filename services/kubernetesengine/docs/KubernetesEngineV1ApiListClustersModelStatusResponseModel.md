# KubernetesEngineV1ApiListClustersModelStatusResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phase** | [**ClusterStatus**](ClusterStatus.md) | 클러스터의 현재 상태 &lt;br/&gt;- &#x60;Pending&#x60;: 대기 중 &lt;br/&gt;- &#x60;Provisioning&#x60;: 생성 중 &lt;br/&gt;- &#x60;Provisioned&#x60;: 생성 완료 &lt;br/&gt;- &#x60;Updating&#x60;: 업데이트 중 &lt;br/&gt;- &#x60;Deleting&#x60;: 삭제 중 &lt;br/&gt;- &#x60;Failed&#x60;: 오류 발생 | 

## Methods

### NewKubernetesEngineV1ApiListClustersModelStatusResponseModel

`func NewKubernetesEngineV1ApiListClustersModelStatusResponseModel(phase ClusterStatus, ) *KubernetesEngineV1ApiListClustersModelStatusResponseModel`

NewKubernetesEngineV1ApiListClustersModelStatusResponseModel instantiates a new KubernetesEngineV1ApiListClustersModelStatusResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesEngineV1ApiListClustersModelStatusResponseModelWithDefaults

`func NewKubernetesEngineV1ApiListClustersModelStatusResponseModelWithDefaults() *KubernetesEngineV1ApiListClustersModelStatusResponseModel`

NewKubernetesEngineV1ApiListClustersModelStatusResponseModelWithDefaults instantiates a new KubernetesEngineV1ApiListClustersModelStatusResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhase

`func (o *KubernetesEngineV1ApiListClustersModelStatusResponseModel) GetPhase() ClusterStatus`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *KubernetesEngineV1ApiListClustersModelStatusResponseModel) GetPhaseOk() (*ClusterStatus, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *KubernetesEngineV1ApiListClustersModelStatusResponseModel) SetPhase(v ClusterStatus)`

SetPhase sets Phase field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


