# SetNodePoolResourceBasedAutoScaling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoscalerDesiredNodeCount** | Pointer to **NullableInt32** | 오토스케일러가 목표로 하는 적정 노드 수 | [optional] 
**IsAutoscalerEnable** | **bool** | 리소스 기반 오토스케일링 활성화 여부 | 
**AutoscalerMaxNodeCount** | Pointer to **NullableInt32** | 리소스 기반 오토스케일링이 허용하는 최대 노드 수 | [optional] 
**AutoscalerMinNodeCount** | Pointer to **NullableInt32** | 리소스 기반 오토스케일링이 허용하는 최소 노드 수 | [optional] 
**AutoscalerScaleDownThreshold** | Pointer to **NullableFloat32** | 리소스 기반 오토스케일링 자동 축소 임계치 조건 - 비율(0~1) 값으로, 소수점 둘째 자리까지 설정 가능 | [optional] 
**AutoscalerScaleDownUnneededTime** | Pointer to **NullableInt32** | 리소스 기반 오토스케일링 자동 축소 임계치 지속 시간 (초) - 입력 가능 범위: 1 ~ 86400 | [optional] 
**AutoscalerScaleDownUnreadyTime** | Pointer to **NullableInt32** | 리소스 기반 오토스케일링 자동 축소 모니터링 제외 시간 (초) - 입력 가능 범위: 1 ~ 86400 | [optional] 

## Methods

### NewSetNodePoolResourceBasedAutoScaling

`func NewSetNodePoolResourceBasedAutoScaling(isAutoscalerEnable bool, ) *SetNodePoolResourceBasedAutoScaling`

NewSetNodePoolResourceBasedAutoScaling instantiates a new SetNodePoolResourceBasedAutoScaling object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetNodePoolResourceBasedAutoScalingWithDefaults

`func NewSetNodePoolResourceBasedAutoScalingWithDefaults() *SetNodePoolResourceBasedAutoScaling`

NewSetNodePoolResourceBasedAutoScalingWithDefaults instantiates a new SetNodePoolResourceBasedAutoScaling object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoscalerDesiredNodeCount

`func (o *SetNodePoolResourceBasedAutoScaling) GetAutoscalerDesiredNodeCount() int32`

GetAutoscalerDesiredNodeCount returns the AutoscalerDesiredNodeCount field if non-nil, zero value otherwise.

### GetAutoscalerDesiredNodeCountOk

`func (o *SetNodePoolResourceBasedAutoScaling) GetAutoscalerDesiredNodeCountOk() (*int32, bool)`

GetAutoscalerDesiredNodeCountOk returns a tuple with the AutoscalerDesiredNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscalerDesiredNodeCount

`func (o *SetNodePoolResourceBasedAutoScaling) SetAutoscalerDesiredNodeCount(v int32)`

SetAutoscalerDesiredNodeCount sets AutoscalerDesiredNodeCount field to given value.

### HasAutoscalerDesiredNodeCount

`func (o *SetNodePoolResourceBasedAutoScaling) HasAutoscalerDesiredNodeCount() bool`

HasAutoscalerDesiredNodeCount returns a boolean if a field has been set.

### SetAutoscalerDesiredNodeCountNil

`func (o *SetNodePoolResourceBasedAutoScaling) SetAutoscalerDesiredNodeCountNil(b bool)`

 SetAutoscalerDesiredNodeCountNil sets the value for AutoscalerDesiredNodeCount to be an explicit nil

### UnsetAutoscalerDesiredNodeCount
`func (o *SetNodePoolResourceBasedAutoScaling) UnsetAutoscalerDesiredNodeCount()`

UnsetAutoscalerDesiredNodeCount ensures that no value is present for AutoscalerDesiredNodeCount, not even an explicit nil
### GetIsAutoscalerEnable

`func (o *SetNodePoolResourceBasedAutoScaling) GetIsAutoscalerEnable() bool`

GetIsAutoscalerEnable returns the IsAutoscalerEnable field if non-nil, zero value otherwise.

### GetIsAutoscalerEnableOk

`func (o *SetNodePoolResourceBasedAutoScaling) GetIsAutoscalerEnableOk() (*bool, bool)`

GetIsAutoscalerEnableOk returns a tuple with the IsAutoscalerEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoscalerEnable

`func (o *SetNodePoolResourceBasedAutoScaling) SetIsAutoscalerEnable(v bool)`

SetIsAutoscalerEnable sets IsAutoscalerEnable field to given value.


### GetAutoscalerMaxNodeCount

`func (o *SetNodePoolResourceBasedAutoScaling) GetAutoscalerMaxNodeCount() int32`

GetAutoscalerMaxNodeCount returns the AutoscalerMaxNodeCount field if non-nil, zero value otherwise.

### GetAutoscalerMaxNodeCountOk

`func (o *SetNodePoolResourceBasedAutoScaling) GetAutoscalerMaxNodeCountOk() (*int32, bool)`

GetAutoscalerMaxNodeCountOk returns a tuple with the AutoscalerMaxNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscalerMaxNodeCount

`func (o *SetNodePoolResourceBasedAutoScaling) SetAutoscalerMaxNodeCount(v int32)`

SetAutoscalerMaxNodeCount sets AutoscalerMaxNodeCount field to given value.

### HasAutoscalerMaxNodeCount

`func (o *SetNodePoolResourceBasedAutoScaling) HasAutoscalerMaxNodeCount() bool`

HasAutoscalerMaxNodeCount returns a boolean if a field has been set.

### SetAutoscalerMaxNodeCountNil

`func (o *SetNodePoolResourceBasedAutoScaling) SetAutoscalerMaxNodeCountNil(b bool)`

 SetAutoscalerMaxNodeCountNil sets the value for AutoscalerMaxNodeCount to be an explicit nil

### UnsetAutoscalerMaxNodeCount
`func (o *SetNodePoolResourceBasedAutoScaling) UnsetAutoscalerMaxNodeCount()`

UnsetAutoscalerMaxNodeCount ensures that no value is present for AutoscalerMaxNodeCount, not even an explicit nil
### GetAutoscalerMinNodeCount

`func (o *SetNodePoolResourceBasedAutoScaling) GetAutoscalerMinNodeCount() int32`

GetAutoscalerMinNodeCount returns the AutoscalerMinNodeCount field if non-nil, zero value otherwise.

### GetAutoscalerMinNodeCountOk

`func (o *SetNodePoolResourceBasedAutoScaling) GetAutoscalerMinNodeCountOk() (*int32, bool)`

GetAutoscalerMinNodeCountOk returns a tuple with the AutoscalerMinNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscalerMinNodeCount

`func (o *SetNodePoolResourceBasedAutoScaling) SetAutoscalerMinNodeCount(v int32)`

SetAutoscalerMinNodeCount sets AutoscalerMinNodeCount field to given value.

### HasAutoscalerMinNodeCount

`func (o *SetNodePoolResourceBasedAutoScaling) HasAutoscalerMinNodeCount() bool`

HasAutoscalerMinNodeCount returns a boolean if a field has been set.

### SetAutoscalerMinNodeCountNil

`func (o *SetNodePoolResourceBasedAutoScaling) SetAutoscalerMinNodeCountNil(b bool)`

 SetAutoscalerMinNodeCountNil sets the value for AutoscalerMinNodeCount to be an explicit nil

### UnsetAutoscalerMinNodeCount
`func (o *SetNodePoolResourceBasedAutoScaling) UnsetAutoscalerMinNodeCount()`

UnsetAutoscalerMinNodeCount ensures that no value is present for AutoscalerMinNodeCount, not even an explicit nil
### GetAutoscalerScaleDownThreshold

`func (o *SetNodePoolResourceBasedAutoScaling) GetAutoscalerScaleDownThreshold() float32`

GetAutoscalerScaleDownThreshold returns the AutoscalerScaleDownThreshold field if non-nil, zero value otherwise.

### GetAutoscalerScaleDownThresholdOk

`func (o *SetNodePoolResourceBasedAutoScaling) GetAutoscalerScaleDownThresholdOk() (*float32, bool)`

GetAutoscalerScaleDownThresholdOk returns a tuple with the AutoscalerScaleDownThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscalerScaleDownThreshold

`func (o *SetNodePoolResourceBasedAutoScaling) SetAutoscalerScaleDownThreshold(v float32)`

SetAutoscalerScaleDownThreshold sets AutoscalerScaleDownThreshold field to given value.

### HasAutoscalerScaleDownThreshold

`func (o *SetNodePoolResourceBasedAutoScaling) HasAutoscalerScaleDownThreshold() bool`

HasAutoscalerScaleDownThreshold returns a boolean if a field has been set.

### SetAutoscalerScaleDownThresholdNil

`func (o *SetNodePoolResourceBasedAutoScaling) SetAutoscalerScaleDownThresholdNil(b bool)`

 SetAutoscalerScaleDownThresholdNil sets the value for AutoscalerScaleDownThreshold to be an explicit nil

### UnsetAutoscalerScaleDownThreshold
`func (o *SetNodePoolResourceBasedAutoScaling) UnsetAutoscalerScaleDownThreshold()`

UnsetAutoscalerScaleDownThreshold ensures that no value is present for AutoscalerScaleDownThreshold, not even an explicit nil
### GetAutoscalerScaleDownUnneededTime

`func (o *SetNodePoolResourceBasedAutoScaling) GetAutoscalerScaleDownUnneededTime() int32`

GetAutoscalerScaleDownUnneededTime returns the AutoscalerScaleDownUnneededTime field if non-nil, zero value otherwise.

### GetAutoscalerScaleDownUnneededTimeOk

`func (o *SetNodePoolResourceBasedAutoScaling) GetAutoscalerScaleDownUnneededTimeOk() (*int32, bool)`

GetAutoscalerScaleDownUnneededTimeOk returns a tuple with the AutoscalerScaleDownUnneededTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscalerScaleDownUnneededTime

`func (o *SetNodePoolResourceBasedAutoScaling) SetAutoscalerScaleDownUnneededTime(v int32)`

SetAutoscalerScaleDownUnneededTime sets AutoscalerScaleDownUnneededTime field to given value.

### HasAutoscalerScaleDownUnneededTime

`func (o *SetNodePoolResourceBasedAutoScaling) HasAutoscalerScaleDownUnneededTime() bool`

HasAutoscalerScaleDownUnneededTime returns a boolean if a field has been set.

### SetAutoscalerScaleDownUnneededTimeNil

`func (o *SetNodePoolResourceBasedAutoScaling) SetAutoscalerScaleDownUnneededTimeNil(b bool)`

 SetAutoscalerScaleDownUnneededTimeNil sets the value for AutoscalerScaleDownUnneededTime to be an explicit nil

### UnsetAutoscalerScaleDownUnneededTime
`func (o *SetNodePoolResourceBasedAutoScaling) UnsetAutoscalerScaleDownUnneededTime()`

UnsetAutoscalerScaleDownUnneededTime ensures that no value is present for AutoscalerScaleDownUnneededTime, not even an explicit nil
### GetAutoscalerScaleDownUnreadyTime

`func (o *SetNodePoolResourceBasedAutoScaling) GetAutoscalerScaleDownUnreadyTime() int32`

GetAutoscalerScaleDownUnreadyTime returns the AutoscalerScaleDownUnreadyTime field if non-nil, zero value otherwise.

### GetAutoscalerScaleDownUnreadyTimeOk

`func (o *SetNodePoolResourceBasedAutoScaling) GetAutoscalerScaleDownUnreadyTimeOk() (*int32, bool)`

GetAutoscalerScaleDownUnreadyTimeOk returns a tuple with the AutoscalerScaleDownUnreadyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscalerScaleDownUnreadyTime

`func (o *SetNodePoolResourceBasedAutoScaling) SetAutoscalerScaleDownUnreadyTime(v int32)`

SetAutoscalerScaleDownUnreadyTime sets AutoscalerScaleDownUnreadyTime field to given value.

### HasAutoscalerScaleDownUnreadyTime

`func (o *SetNodePoolResourceBasedAutoScaling) HasAutoscalerScaleDownUnreadyTime() bool`

HasAutoscalerScaleDownUnreadyTime returns a boolean if a field has been set.

### SetAutoscalerScaleDownUnreadyTimeNil

`func (o *SetNodePoolResourceBasedAutoScaling) SetAutoscalerScaleDownUnreadyTimeNil(b bool)`

 SetAutoscalerScaleDownUnreadyTimeNil sets the value for AutoscalerScaleDownUnreadyTime to be an explicit nil

### UnsetAutoscalerScaleDownUnreadyTime
`func (o *SetNodePoolResourceBasedAutoScaling) UnsetAutoscalerScaleDownUnreadyTime()`

UnsetAutoscalerScaleDownUnreadyTime ensures that no value is present for AutoscalerScaleDownUnreadyTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


