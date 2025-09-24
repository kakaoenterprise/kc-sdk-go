# UpdateKubernetesEngineClusterNodePoolScalingResourceRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Autoscaling** | [**NodePoolScalingResourceRequestModel**](NodePoolScalingResourceRequestModel.md) | 노드 풀의 리소스 기반 오토스케일링 설정 정보 | 

## Methods

### NewUpdateKubernetesEngineClusterNodePoolScalingResourceRequestModel

`func NewUpdateKubernetesEngineClusterNodePoolScalingResourceRequestModel(autoscaling NodePoolScalingResourceRequestModel, ) *UpdateKubernetesEngineClusterNodePoolScalingResourceRequestModel`

NewUpdateKubernetesEngineClusterNodePoolScalingResourceRequestModel instantiates a new UpdateKubernetesEngineClusterNodePoolScalingResourceRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateKubernetesEngineClusterNodePoolScalingResourceRequestModelWithDefaults

`func NewUpdateKubernetesEngineClusterNodePoolScalingResourceRequestModelWithDefaults() *UpdateKubernetesEngineClusterNodePoolScalingResourceRequestModel`

NewUpdateKubernetesEngineClusterNodePoolScalingResourceRequestModelWithDefaults instantiates a new UpdateKubernetesEngineClusterNodePoolScalingResourceRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoscaling

`func (o *UpdateKubernetesEngineClusterNodePoolScalingResourceRequestModel) GetAutoscaling() NodePoolScalingResourceRequestModel`

GetAutoscaling returns the Autoscaling field if non-nil, zero value otherwise.

### GetAutoscalingOk

`func (o *UpdateKubernetesEngineClusterNodePoolScalingResourceRequestModel) GetAutoscalingOk() (*NodePoolScalingResourceRequestModel, bool)`

GetAutoscalingOk returns a tuple with the Autoscaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscaling

`func (o *UpdateKubernetesEngineClusterNodePoolScalingResourceRequestModel) SetAutoscaling(v NodePoolScalingResourceRequestModel)`

SetAutoscaling sets Autoscaling field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


