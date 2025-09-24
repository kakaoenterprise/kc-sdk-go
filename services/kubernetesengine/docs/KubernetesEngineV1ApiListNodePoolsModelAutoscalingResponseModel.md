# KubernetesEngineV1ApiListNodePoolsModelAutoscalingResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsAutoscalerEnable** | **bool** | 리소스 기반 오토스케일링 사용 여부 | 
**AutoscalerDesiredNodeCount** | Pointer to **NullableInt32** |  | [optional] 
**AutoscalerMaxNodeCount** | Pointer to **NullableInt32** |  | [optional] 
**AutoscalerMinNodeCount** | Pointer to **NullableInt32** |  | [optional] 
**AutoscalerScaleDownThreshold** | Pointer to **NullableFloat32** |  | [optional] 
**AutoscalerScaleDownUnneededTime** | Pointer to **NullableInt32** |  | [optional] 
**AutoscalerScaleDownUnreadyTime** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewKubernetesEngineV1ApiListNodePoolsModelAutoscalingResponseModel

`func NewKubernetesEngineV1ApiListNodePoolsModelAutoscalingResponseModel(isAutoscalerEnable bool, ) *KubernetesEngineV1ApiListNodePoolsModelAutoscalingResponseModel`

NewKubernetesEngineV1ApiListNodePoolsModelAutoscalingResponseModel instantiates a new KubernetesEngineV1ApiListNodePoolsModelAutoscalingResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesEngineV1ApiListNodePoolsModelAutoscalingResponseModelWithDefaults

`func NewKubernetesEngineV1ApiListNodePoolsModelAutoscalingResponseModelWithDefaults() *KubernetesEngineV1ApiListNodePoolsModelAutoscalingResponseModel`

NewKubernetesEngineV1ApiListNodePoolsModelAutoscalingResponseModelWithDefaults instantiates a new KubernetesEngineV1ApiListNodePoolsModelAutoscalingResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsAutoscalerEnable

`func (o *KubernetesEngineV1ApiListNodePoolsModelAutoscalingResponseModel) GetIsAutoscalerEnable() bool`

GetIsAutoscalerEnable returns the IsAutoscalerEnable field if non-nil, zero value otherwise.

### GetIsAutoscalerEnableOk

`func (o *KubernetesEngineV1ApiListNodePoolsModelAutoscalingResponseModel) GetIsAutoscalerEnableOk() (*bool, bool)`

GetIsAutoscalerEnableOk returns a tuple with the IsAutoscalerEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoscalerEnable

`func (o *KubernetesEngineV1ApiListNodePoolsModelAutoscalingResponseModel) SetIsAutoscalerEnable(v bool)`

SetIsAutoscalerEnable sets IsAutoscalerEnable field to given value.


### GetAutoscalerDesiredNodeCount

`func (o *KubernetesEngineV1ApiListNodePoolsModelAutoscalingResponseModel) GetAutoscalerDesiredNodeCount() int32`

GetAutoscalerDesiredNodeCount returns the AutoscalerDesiredNodeCount field if non-nil, zero value otherwise.

### GetAutoscalerDesiredNodeCountOk

`func (o *KubernetesEngineV1ApiListNodePoolsModelAutoscalingResponseModel) GetAutoscalerDesiredNodeCountOk() (*int32, bool)`

GetAutoscalerDesiredNodeCountOk returns a tuple with the AutoscalerDesiredNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscalerDesiredNodeCount

`func (o *KubernetesEngineV1ApiListNodePoolsModelAutoscalingResponseModel) SetAutoscalerDesiredNodeCount(v int32)`

SetAutoscalerDesiredNodeCount sets AutoscalerDesiredNodeCount field to given value.

### HasAutoscalerDesiredNodeCount

`func (o *KubernetesEngineV1ApiListNodePoolsModelAutoscalingResponseModel) HasAutoscalerDesiredNodeCount() bool`

HasAutoscalerDesiredNodeCount returns a boolean if a field has been set.

### SetAutoscalerDesiredNodeCountNil

`func (o *KubernetesEngineV1ApiListNodePoolsModelAutoscalingResponseModel) SetAutoscalerDesiredNodeCountNil(b bool)`

 SetAutoscalerDesiredNodeCountNil sets the value for AutoscalerDesiredNodeCount to be an explicit nil

### UnsetAutoscalerDesiredNodeCount
`func (o *KubernetesEngineV1ApiListNodePoolsModelAutoscalingResponseModel) UnsetAutoscalerDesiredNodeCount()`

UnsetAutoscalerDesiredNodeCount ensures that no value is present for AutoscalerDesiredNodeCount, not even an explicit nil
### GetAutoscalerMaxNodeCount

`func (o *KubernetesEngineV1ApiListNodePoolsModelAutoscalingResponseModel) GetAutoscalerMaxNodeCount() int32`

GetAutoscalerMaxNodeCount returns the AutoscalerMaxNodeCount field if non-nil, zero value otherwise.

### GetAutoscalerMaxNodeCountOk

`func (o *KubernetesEngineV1ApiListNodePoolsModelAutoscalingResponseModel) GetAutoscalerMaxNodeCountOk() (*int32, bool)`

GetAutoscalerMaxNodeCountOk returns a tuple with the AutoscalerMaxNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscalerMaxNodeCount

`func (o *KubernetesEngineV1ApiListNodePoolsModelAutoscalingResponseModel) SetAutoscalerMaxNodeCount(v int32)`

SetAutoscalerMaxNodeCount sets AutoscalerMaxNodeCount field to given value.

### HasAutoscalerMaxNodeCount

`func (o *KubernetesEngineV1ApiListNodePoolsModelAutoscalingResponseModel) HasAutoscalerMaxNodeCount() bool`

HasAutoscalerMaxNodeCount returns a boolean if a field has been set.

### SetAutoscalerMaxNodeCountNil

`func (o *KubernetesEngineV1ApiListNodePoolsModelAutoscalingResponseModel) SetAutoscalerMaxNodeCountNil(b bool)`

 SetAutoscalerMaxNodeCountNil sets the value for AutoscalerMaxNodeCount to be an explicit nil

### UnsetAutoscalerMaxNodeCount
`func (o *KubernetesEngineV1ApiListNodePoolsModelAutoscalingResponseModel) UnsetAutoscalerMaxNodeCount()`

UnsetAutoscalerMaxNodeCount ensures that no value is present for AutoscalerMaxNodeCount, not even an explicit nil
### GetAutoscalerMinNodeCount

`func (o *KubernetesEngineV1ApiListNodePoolsModelAutoscalingResponseModel) GetAutoscalerMinNodeCount() int32`

GetAutoscalerMinNodeCount returns the AutoscalerMinNodeCount field if non-nil, zero value otherwise.

### GetAutoscalerMinNodeCountOk

`func (o *KubernetesEngineV1ApiListNodePoolsModelAutoscalingResponseModel) GetAutoscalerMinNodeCountOk() (*int32, bool)`

GetAutoscalerMinNodeCountOk returns a tuple with the AutoscalerMinNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscalerMinNodeCount

`func (o *KubernetesEngineV1ApiListNodePoolsModelAutoscalingResponseModel) SetAutoscalerMinNodeCount(v int32)`

SetAutoscalerMinNodeCount sets AutoscalerMinNodeCount field to given value.

### HasAutoscalerMinNodeCount

`func (o *KubernetesEngineV1ApiListNodePoolsModelAutoscalingResponseModel) HasAutoscalerMinNodeCount() bool`

HasAutoscalerMinNodeCount returns a boolean if a field has been set.

### SetAutoscalerMinNodeCountNil

`func (o *KubernetesEngineV1ApiListNodePoolsModelAutoscalingResponseModel) SetAutoscalerMinNodeCountNil(b bool)`

 SetAutoscalerMinNodeCountNil sets the value for AutoscalerMinNodeCount to be an explicit nil

### UnsetAutoscalerMinNodeCount
`func (o *KubernetesEngineV1ApiListNodePoolsModelAutoscalingResponseModel) UnsetAutoscalerMinNodeCount()`

UnsetAutoscalerMinNodeCount ensures that no value is present for AutoscalerMinNodeCount, not even an explicit nil
### GetAutoscalerScaleDownThreshold

`func (o *KubernetesEngineV1ApiListNodePoolsModelAutoscalingResponseModel) GetAutoscalerScaleDownThreshold() float32`

GetAutoscalerScaleDownThreshold returns the AutoscalerScaleDownThreshold field if non-nil, zero value otherwise.

### GetAutoscalerScaleDownThresholdOk

`func (o *KubernetesEngineV1ApiListNodePoolsModelAutoscalingResponseModel) GetAutoscalerScaleDownThresholdOk() (*float32, bool)`

GetAutoscalerScaleDownThresholdOk returns a tuple with the AutoscalerScaleDownThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscalerScaleDownThreshold

`func (o *KubernetesEngineV1ApiListNodePoolsModelAutoscalingResponseModel) SetAutoscalerScaleDownThreshold(v float32)`

SetAutoscalerScaleDownThreshold sets AutoscalerScaleDownThreshold field to given value.

### HasAutoscalerScaleDownThreshold

`func (o *KubernetesEngineV1ApiListNodePoolsModelAutoscalingResponseModel) HasAutoscalerScaleDownThreshold() bool`

HasAutoscalerScaleDownThreshold returns a boolean if a field has been set.

### SetAutoscalerScaleDownThresholdNil

`func (o *KubernetesEngineV1ApiListNodePoolsModelAutoscalingResponseModel) SetAutoscalerScaleDownThresholdNil(b bool)`

 SetAutoscalerScaleDownThresholdNil sets the value for AutoscalerScaleDownThreshold to be an explicit nil

### UnsetAutoscalerScaleDownThreshold
`func (o *KubernetesEngineV1ApiListNodePoolsModelAutoscalingResponseModel) UnsetAutoscalerScaleDownThreshold()`

UnsetAutoscalerScaleDownThreshold ensures that no value is present for AutoscalerScaleDownThreshold, not even an explicit nil
### GetAutoscalerScaleDownUnneededTime

`func (o *KubernetesEngineV1ApiListNodePoolsModelAutoscalingResponseModel) GetAutoscalerScaleDownUnneededTime() int32`

GetAutoscalerScaleDownUnneededTime returns the AutoscalerScaleDownUnneededTime field if non-nil, zero value otherwise.

### GetAutoscalerScaleDownUnneededTimeOk

`func (o *KubernetesEngineV1ApiListNodePoolsModelAutoscalingResponseModel) GetAutoscalerScaleDownUnneededTimeOk() (*int32, bool)`

GetAutoscalerScaleDownUnneededTimeOk returns a tuple with the AutoscalerScaleDownUnneededTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscalerScaleDownUnneededTime

`func (o *KubernetesEngineV1ApiListNodePoolsModelAutoscalingResponseModel) SetAutoscalerScaleDownUnneededTime(v int32)`

SetAutoscalerScaleDownUnneededTime sets AutoscalerScaleDownUnneededTime field to given value.

### HasAutoscalerScaleDownUnneededTime

`func (o *KubernetesEngineV1ApiListNodePoolsModelAutoscalingResponseModel) HasAutoscalerScaleDownUnneededTime() bool`

HasAutoscalerScaleDownUnneededTime returns a boolean if a field has been set.

### SetAutoscalerScaleDownUnneededTimeNil

`func (o *KubernetesEngineV1ApiListNodePoolsModelAutoscalingResponseModel) SetAutoscalerScaleDownUnneededTimeNil(b bool)`

 SetAutoscalerScaleDownUnneededTimeNil sets the value for AutoscalerScaleDownUnneededTime to be an explicit nil

### UnsetAutoscalerScaleDownUnneededTime
`func (o *KubernetesEngineV1ApiListNodePoolsModelAutoscalingResponseModel) UnsetAutoscalerScaleDownUnneededTime()`

UnsetAutoscalerScaleDownUnneededTime ensures that no value is present for AutoscalerScaleDownUnneededTime, not even an explicit nil
### GetAutoscalerScaleDownUnreadyTime

`func (o *KubernetesEngineV1ApiListNodePoolsModelAutoscalingResponseModel) GetAutoscalerScaleDownUnreadyTime() int32`

GetAutoscalerScaleDownUnreadyTime returns the AutoscalerScaleDownUnreadyTime field if non-nil, zero value otherwise.

### GetAutoscalerScaleDownUnreadyTimeOk

`func (o *KubernetesEngineV1ApiListNodePoolsModelAutoscalingResponseModel) GetAutoscalerScaleDownUnreadyTimeOk() (*int32, bool)`

GetAutoscalerScaleDownUnreadyTimeOk returns a tuple with the AutoscalerScaleDownUnreadyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscalerScaleDownUnreadyTime

`func (o *KubernetesEngineV1ApiListNodePoolsModelAutoscalingResponseModel) SetAutoscalerScaleDownUnreadyTime(v int32)`

SetAutoscalerScaleDownUnreadyTime sets AutoscalerScaleDownUnreadyTime field to given value.

### HasAutoscalerScaleDownUnreadyTime

`func (o *KubernetesEngineV1ApiListNodePoolsModelAutoscalingResponseModel) HasAutoscalerScaleDownUnreadyTime() bool`

HasAutoscalerScaleDownUnreadyTime returns a boolean if a field has been set.

### SetAutoscalerScaleDownUnreadyTimeNil

`func (o *KubernetesEngineV1ApiListNodePoolsModelAutoscalingResponseModel) SetAutoscalerScaleDownUnreadyTimeNil(b bool)`

 SetAutoscalerScaleDownUnreadyTimeNil sets the value for AutoscalerScaleDownUnreadyTime to be an explicit nil

### UnsetAutoscalerScaleDownUnreadyTime
`func (o *KubernetesEngineV1ApiListNodePoolsModelAutoscalingResponseModel) UnsetAutoscalerScaleDownUnreadyTime()`

UnsetAutoscalerScaleDownUnreadyTime ensures that no value is present for AutoscalerScaleDownUnreadyTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


