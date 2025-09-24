# NodePoolScalingResourceRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoscalerDesiredNodeCount** | Pointer to **NullableInt32** |  | [optional] 
**IsAutoscalerEnable** | **bool** | 리소스 기반 오토스케일링 활성화 여부 | 
**AutoscalerMaxNodeCount** | Pointer to **NullableInt32** |  | [optional] 
**AutoscalerMinNodeCount** | Pointer to **NullableInt32** |  | [optional] 
**AutoscalerScaleDownThreshold** | Pointer to **NullableFloat32** |  | [optional] 
**AutoscalerScaleDownUnneededTime** | Pointer to **NullableInt32** |  | [optional] 
**AutoscalerScaleDownUnreadyTime** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewNodePoolScalingResourceRequestModel

`func NewNodePoolScalingResourceRequestModel(isAutoscalerEnable bool, ) *NodePoolScalingResourceRequestModel`

NewNodePoolScalingResourceRequestModel instantiates a new NodePoolScalingResourceRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodePoolScalingResourceRequestModelWithDefaults

`func NewNodePoolScalingResourceRequestModelWithDefaults() *NodePoolScalingResourceRequestModel`

NewNodePoolScalingResourceRequestModelWithDefaults instantiates a new NodePoolScalingResourceRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoscalerDesiredNodeCount

`func (o *NodePoolScalingResourceRequestModel) GetAutoscalerDesiredNodeCount() int32`

GetAutoscalerDesiredNodeCount returns the AutoscalerDesiredNodeCount field if non-nil, zero value otherwise.

### GetAutoscalerDesiredNodeCountOk

`func (o *NodePoolScalingResourceRequestModel) GetAutoscalerDesiredNodeCountOk() (*int32, bool)`

GetAutoscalerDesiredNodeCountOk returns a tuple with the AutoscalerDesiredNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscalerDesiredNodeCount

`func (o *NodePoolScalingResourceRequestModel) SetAutoscalerDesiredNodeCount(v int32)`

SetAutoscalerDesiredNodeCount sets AutoscalerDesiredNodeCount field to given value.

### HasAutoscalerDesiredNodeCount

`func (o *NodePoolScalingResourceRequestModel) HasAutoscalerDesiredNodeCount() bool`

HasAutoscalerDesiredNodeCount returns a boolean if a field has been set.

### SetAutoscalerDesiredNodeCountNil

`func (o *NodePoolScalingResourceRequestModel) SetAutoscalerDesiredNodeCountNil(b bool)`

 SetAutoscalerDesiredNodeCountNil sets the value for AutoscalerDesiredNodeCount to be an explicit nil

### UnsetAutoscalerDesiredNodeCount
`func (o *NodePoolScalingResourceRequestModel) UnsetAutoscalerDesiredNodeCount()`

UnsetAutoscalerDesiredNodeCount ensures that no value is present for AutoscalerDesiredNodeCount, not even an explicit nil
### GetIsAutoscalerEnable

`func (o *NodePoolScalingResourceRequestModel) GetIsAutoscalerEnable() bool`

GetIsAutoscalerEnable returns the IsAutoscalerEnable field if non-nil, zero value otherwise.

### GetIsAutoscalerEnableOk

`func (o *NodePoolScalingResourceRequestModel) GetIsAutoscalerEnableOk() (*bool, bool)`

GetIsAutoscalerEnableOk returns a tuple with the IsAutoscalerEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoscalerEnable

`func (o *NodePoolScalingResourceRequestModel) SetIsAutoscalerEnable(v bool)`

SetIsAutoscalerEnable sets IsAutoscalerEnable field to given value.


### GetAutoscalerMaxNodeCount

`func (o *NodePoolScalingResourceRequestModel) GetAutoscalerMaxNodeCount() int32`

GetAutoscalerMaxNodeCount returns the AutoscalerMaxNodeCount field if non-nil, zero value otherwise.

### GetAutoscalerMaxNodeCountOk

`func (o *NodePoolScalingResourceRequestModel) GetAutoscalerMaxNodeCountOk() (*int32, bool)`

GetAutoscalerMaxNodeCountOk returns a tuple with the AutoscalerMaxNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscalerMaxNodeCount

`func (o *NodePoolScalingResourceRequestModel) SetAutoscalerMaxNodeCount(v int32)`

SetAutoscalerMaxNodeCount sets AutoscalerMaxNodeCount field to given value.

### HasAutoscalerMaxNodeCount

`func (o *NodePoolScalingResourceRequestModel) HasAutoscalerMaxNodeCount() bool`

HasAutoscalerMaxNodeCount returns a boolean if a field has been set.

### SetAutoscalerMaxNodeCountNil

`func (o *NodePoolScalingResourceRequestModel) SetAutoscalerMaxNodeCountNil(b bool)`

 SetAutoscalerMaxNodeCountNil sets the value for AutoscalerMaxNodeCount to be an explicit nil

### UnsetAutoscalerMaxNodeCount
`func (o *NodePoolScalingResourceRequestModel) UnsetAutoscalerMaxNodeCount()`

UnsetAutoscalerMaxNodeCount ensures that no value is present for AutoscalerMaxNodeCount, not even an explicit nil
### GetAutoscalerMinNodeCount

`func (o *NodePoolScalingResourceRequestModel) GetAutoscalerMinNodeCount() int32`

GetAutoscalerMinNodeCount returns the AutoscalerMinNodeCount field if non-nil, zero value otherwise.

### GetAutoscalerMinNodeCountOk

`func (o *NodePoolScalingResourceRequestModel) GetAutoscalerMinNodeCountOk() (*int32, bool)`

GetAutoscalerMinNodeCountOk returns a tuple with the AutoscalerMinNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscalerMinNodeCount

`func (o *NodePoolScalingResourceRequestModel) SetAutoscalerMinNodeCount(v int32)`

SetAutoscalerMinNodeCount sets AutoscalerMinNodeCount field to given value.

### HasAutoscalerMinNodeCount

`func (o *NodePoolScalingResourceRequestModel) HasAutoscalerMinNodeCount() bool`

HasAutoscalerMinNodeCount returns a boolean if a field has been set.

### SetAutoscalerMinNodeCountNil

`func (o *NodePoolScalingResourceRequestModel) SetAutoscalerMinNodeCountNil(b bool)`

 SetAutoscalerMinNodeCountNil sets the value for AutoscalerMinNodeCount to be an explicit nil

### UnsetAutoscalerMinNodeCount
`func (o *NodePoolScalingResourceRequestModel) UnsetAutoscalerMinNodeCount()`

UnsetAutoscalerMinNodeCount ensures that no value is present for AutoscalerMinNodeCount, not even an explicit nil
### GetAutoscalerScaleDownThreshold

`func (o *NodePoolScalingResourceRequestModel) GetAutoscalerScaleDownThreshold() float32`

GetAutoscalerScaleDownThreshold returns the AutoscalerScaleDownThreshold field if non-nil, zero value otherwise.

### GetAutoscalerScaleDownThresholdOk

`func (o *NodePoolScalingResourceRequestModel) GetAutoscalerScaleDownThresholdOk() (*float32, bool)`

GetAutoscalerScaleDownThresholdOk returns a tuple with the AutoscalerScaleDownThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscalerScaleDownThreshold

`func (o *NodePoolScalingResourceRequestModel) SetAutoscalerScaleDownThreshold(v float32)`

SetAutoscalerScaleDownThreshold sets AutoscalerScaleDownThreshold field to given value.

### HasAutoscalerScaleDownThreshold

`func (o *NodePoolScalingResourceRequestModel) HasAutoscalerScaleDownThreshold() bool`

HasAutoscalerScaleDownThreshold returns a boolean if a field has been set.

### SetAutoscalerScaleDownThresholdNil

`func (o *NodePoolScalingResourceRequestModel) SetAutoscalerScaleDownThresholdNil(b bool)`

 SetAutoscalerScaleDownThresholdNil sets the value for AutoscalerScaleDownThreshold to be an explicit nil

### UnsetAutoscalerScaleDownThreshold
`func (o *NodePoolScalingResourceRequestModel) UnsetAutoscalerScaleDownThreshold()`

UnsetAutoscalerScaleDownThreshold ensures that no value is present for AutoscalerScaleDownThreshold, not even an explicit nil
### GetAutoscalerScaleDownUnneededTime

`func (o *NodePoolScalingResourceRequestModel) GetAutoscalerScaleDownUnneededTime() int32`

GetAutoscalerScaleDownUnneededTime returns the AutoscalerScaleDownUnneededTime field if non-nil, zero value otherwise.

### GetAutoscalerScaleDownUnneededTimeOk

`func (o *NodePoolScalingResourceRequestModel) GetAutoscalerScaleDownUnneededTimeOk() (*int32, bool)`

GetAutoscalerScaleDownUnneededTimeOk returns a tuple with the AutoscalerScaleDownUnneededTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscalerScaleDownUnneededTime

`func (o *NodePoolScalingResourceRequestModel) SetAutoscalerScaleDownUnneededTime(v int32)`

SetAutoscalerScaleDownUnneededTime sets AutoscalerScaleDownUnneededTime field to given value.

### HasAutoscalerScaleDownUnneededTime

`func (o *NodePoolScalingResourceRequestModel) HasAutoscalerScaleDownUnneededTime() bool`

HasAutoscalerScaleDownUnneededTime returns a boolean if a field has been set.

### SetAutoscalerScaleDownUnneededTimeNil

`func (o *NodePoolScalingResourceRequestModel) SetAutoscalerScaleDownUnneededTimeNil(b bool)`

 SetAutoscalerScaleDownUnneededTimeNil sets the value for AutoscalerScaleDownUnneededTime to be an explicit nil

### UnsetAutoscalerScaleDownUnneededTime
`func (o *NodePoolScalingResourceRequestModel) UnsetAutoscalerScaleDownUnneededTime()`

UnsetAutoscalerScaleDownUnneededTime ensures that no value is present for AutoscalerScaleDownUnneededTime, not even an explicit nil
### GetAutoscalerScaleDownUnreadyTime

`func (o *NodePoolScalingResourceRequestModel) GetAutoscalerScaleDownUnreadyTime() int32`

GetAutoscalerScaleDownUnreadyTime returns the AutoscalerScaleDownUnreadyTime field if non-nil, zero value otherwise.

### GetAutoscalerScaleDownUnreadyTimeOk

`func (o *NodePoolScalingResourceRequestModel) GetAutoscalerScaleDownUnreadyTimeOk() (*int32, bool)`

GetAutoscalerScaleDownUnreadyTimeOk returns a tuple with the AutoscalerScaleDownUnreadyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscalerScaleDownUnreadyTime

`func (o *NodePoolScalingResourceRequestModel) SetAutoscalerScaleDownUnreadyTime(v int32)`

SetAutoscalerScaleDownUnreadyTime sets AutoscalerScaleDownUnreadyTime field to given value.

### HasAutoscalerScaleDownUnreadyTime

`func (o *NodePoolScalingResourceRequestModel) HasAutoscalerScaleDownUnreadyTime() bool`

HasAutoscalerScaleDownUnreadyTime returns a boolean if a field has been set.

### SetAutoscalerScaleDownUnreadyTimeNil

`func (o *NodePoolScalingResourceRequestModel) SetAutoscalerScaleDownUnreadyTimeNil(b bool)`

 SetAutoscalerScaleDownUnreadyTimeNil sets the value for AutoscalerScaleDownUnreadyTime to be an explicit nil

### UnsetAutoscalerScaleDownUnreadyTime
`func (o *NodePoolScalingResourceRequestModel) UnsetAutoscalerScaleDownUnreadyTime()`

UnsetAutoscalerScaleDownUnreadyTime ensures that no value is present for AutoscalerScaleDownUnreadyTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


