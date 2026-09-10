# Autoscaling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsAutoscalerEnable** | Pointer to **NullableBool** |  | [optional] 
**AutoscalerDesiredNodeCount** | Pointer to **NullableInt32** |  | [optional] 
**AutoscalerMaxNodeCount** | Pointer to **NullableInt32** |  | [optional] 
**AutoscalerMinNodeCount** | Pointer to **NullableInt32** |  | [optional] 
**AutoscalerScaleDownThreshold** | Pointer to **NullableFloat32** |  | [optional] 
**AutoscalerScaleDownUnneededTime** | Pointer to **NullableInt32** |  | [optional] 
**AutoscalerScaleDownUnreadyTime** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewAutoscaling

`func NewAutoscaling() *Autoscaling`

NewAutoscaling instantiates a new Autoscaling object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoscalingWithDefaults

`func NewAutoscalingWithDefaults() *Autoscaling`

NewAutoscalingWithDefaults instantiates a new Autoscaling object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsAutoscalerEnable

`func (o *Autoscaling) GetIsAutoscalerEnable() bool`

GetIsAutoscalerEnable returns the IsAutoscalerEnable field if non-nil, zero value otherwise.

### GetIsAutoscalerEnableOk

`func (o *Autoscaling) GetIsAutoscalerEnableOk() (*bool, bool)`

GetIsAutoscalerEnableOk returns a tuple with the IsAutoscalerEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoscalerEnable

`func (o *Autoscaling) SetIsAutoscalerEnable(v bool)`

SetIsAutoscalerEnable sets IsAutoscalerEnable field to given value.

### HasIsAutoscalerEnable

`func (o *Autoscaling) HasIsAutoscalerEnable() bool`

HasIsAutoscalerEnable returns a boolean if a field has been set.

### SetIsAutoscalerEnableNil

`func (o *Autoscaling) SetIsAutoscalerEnableNil(b bool)`

 SetIsAutoscalerEnableNil sets the value for IsAutoscalerEnable to be an explicit nil

### UnsetIsAutoscalerEnable
`func (o *Autoscaling) UnsetIsAutoscalerEnable()`

UnsetIsAutoscalerEnable ensures that no value is present for IsAutoscalerEnable, not even an explicit nil
### GetAutoscalerDesiredNodeCount

`func (o *Autoscaling) GetAutoscalerDesiredNodeCount() int32`

GetAutoscalerDesiredNodeCount returns the AutoscalerDesiredNodeCount field if non-nil, zero value otherwise.

### GetAutoscalerDesiredNodeCountOk

`func (o *Autoscaling) GetAutoscalerDesiredNodeCountOk() (*int32, bool)`

GetAutoscalerDesiredNodeCountOk returns a tuple with the AutoscalerDesiredNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscalerDesiredNodeCount

`func (o *Autoscaling) SetAutoscalerDesiredNodeCount(v int32)`

SetAutoscalerDesiredNodeCount sets AutoscalerDesiredNodeCount field to given value.

### HasAutoscalerDesiredNodeCount

`func (o *Autoscaling) HasAutoscalerDesiredNodeCount() bool`

HasAutoscalerDesiredNodeCount returns a boolean if a field has been set.

### SetAutoscalerDesiredNodeCountNil

`func (o *Autoscaling) SetAutoscalerDesiredNodeCountNil(b bool)`

 SetAutoscalerDesiredNodeCountNil sets the value for AutoscalerDesiredNodeCount to be an explicit nil

### UnsetAutoscalerDesiredNodeCount
`func (o *Autoscaling) UnsetAutoscalerDesiredNodeCount()`

UnsetAutoscalerDesiredNodeCount ensures that no value is present for AutoscalerDesiredNodeCount, not even an explicit nil
### GetAutoscalerMaxNodeCount

`func (o *Autoscaling) GetAutoscalerMaxNodeCount() int32`

GetAutoscalerMaxNodeCount returns the AutoscalerMaxNodeCount field if non-nil, zero value otherwise.

### GetAutoscalerMaxNodeCountOk

`func (o *Autoscaling) GetAutoscalerMaxNodeCountOk() (*int32, bool)`

GetAutoscalerMaxNodeCountOk returns a tuple with the AutoscalerMaxNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscalerMaxNodeCount

`func (o *Autoscaling) SetAutoscalerMaxNodeCount(v int32)`

SetAutoscalerMaxNodeCount sets AutoscalerMaxNodeCount field to given value.

### HasAutoscalerMaxNodeCount

`func (o *Autoscaling) HasAutoscalerMaxNodeCount() bool`

HasAutoscalerMaxNodeCount returns a boolean if a field has been set.

### SetAutoscalerMaxNodeCountNil

`func (o *Autoscaling) SetAutoscalerMaxNodeCountNil(b bool)`

 SetAutoscalerMaxNodeCountNil sets the value for AutoscalerMaxNodeCount to be an explicit nil

### UnsetAutoscalerMaxNodeCount
`func (o *Autoscaling) UnsetAutoscalerMaxNodeCount()`

UnsetAutoscalerMaxNodeCount ensures that no value is present for AutoscalerMaxNodeCount, not even an explicit nil
### GetAutoscalerMinNodeCount

`func (o *Autoscaling) GetAutoscalerMinNodeCount() int32`

GetAutoscalerMinNodeCount returns the AutoscalerMinNodeCount field if non-nil, zero value otherwise.

### GetAutoscalerMinNodeCountOk

`func (o *Autoscaling) GetAutoscalerMinNodeCountOk() (*int32, bool)`

GetAutoscalerMinNodeCountOk returns a tuple with the AutoscalerMinNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscalerMinNodeCount

`func (o *Autoscaling) SetAutoscalerMinNodeCount(v int32)`

SetAutoscalerMinNodeCount sets AutoscalerMinNodeCount field to given value.

### HasAutoscalerMinNodeCount

`func (o *Autoscaling) HasAutoscalerMinNodeCount() bool`

HasAutoscalerMinNodeCount returns a boolean if a field has been set.

### SetAutoscalerMinNodeCountNil

`func (o *Autoscaling) SetAutoscalerMinNodeCountNil(b bool)`

 SetAutoscalerMinNodeCountNil sets the value for AutoscalerMinNodeCount to be an explicit nil

### UnsetAutoscalerMinNodeCount
`func (o *Autoscaling) UnsetAutoscalerMinNodeCount()`

UnsetAutoscalerMinNodeCount ensures that no value is present for AutoscalerMinNodeCount, not even an explicit nil
### GetAutoscalerScaleDownThreshold

`func (o *Autoscaling) GetAutoscalerScaleDownThreshold() float32`

GetAutoscalerScaleDownThreshold returns the AutoscalerScaleDownThreshold field if non-nil, zero value otherwise.

### GetAutoscalerScaleDownThresholdOk

`func (o *Autoscaling) GetAutoscalerScaleDownThresholdOk() (*float32, bool)`

GetAutoscalerScaleDownThresholdOk returns a tuple with the AutoscalerScaleDownThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscalerScaleDownThreshold

`func (o *Autoscaling) SetAutoscalerScaleDownThreshold(v float32)`

SetAutoscalerScaleDownThreshold sets AutoscalerScaleDownThreshold field to given value.

### HasAutoscalerScaleDownThreshold

`func (o *Autoscaling) HasAutoscalerScaleDownThreshold() bool`

HasAutoscalerScaleDownThreshold returns a boolean if a field has been set.

### SetAutoscalerScaleDownThresholdNil

`func (o *Autoscaling) SetAutoscalerScaleDownThresholdNil(b bool)`

 SetAutoscalerScaleDownThresholdNil sets the value for AutoscalerScaleDownThreshold to be an explicit nil

### UnsetAutoscalerScaleDownThreshold
`func (o *Autoscaling) UnsetAutoscalerScaleDownThreshold()`

UnsetAutoscalerScaleDownThreshold ensures that no value is present for AutoscalerScaleDownThreshold, not even an explicit nil
### GetAutoscalerScaleDownUnneededTime

`func (o *Autoscaling) GetAutoscalerScaleDownUnneededTime() int32`

GetAutoscalerScaleDownUnneededTime returns the AutoscalerScaleDownUnneededTime field if non-nil, zero value otherwise.

### GetAutoscalerScaleDownUnneededTimeOk

`func (o *Autoscaling) GetAutoscalerScaleDownUnneededTimeOk() (*int32, bool)`

GetAutoscalerScaleDownUnneededTimeOk returns a tuple with the AutoscalerScaleDownUnneededTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscalerScaleDownUnneededTime

`func (o *Autoscaling) SetAutoscalerScaleDownUnneededTime(v int32)`

SetAutoscalerScaleDownUnneededTime sets AutoscalerScaleDownUnneededTime field to given value.

### HasAutoscalerScaleDownUnneededTime

`func (o *Autoscaling) HasAutoscalerScaleDownUnneededTime() bool`

HasAutoscalerScaleDownUnneededTime returns a boolean if a field has been set.

### SetAutoscalerScaleDownUnneededTimeNil

`func (o *Autoscaling) SetAutoscalerScaleDownUnneededTimeNil(b bool)`

 SetAutoscalerScaleDownUnneededTimeNil sets the value for AutoscalerScaleDownUnneededTime to be an explicit nil

### UnsetAutoscalerScaleDownUnneededTime
`func (o *Autoscaling) UnsetAutoscalerScaleDownUnneededTime()`

UnsetAutoscalerScaleDownUnneededTime ensures that no value is present for AutoscalerScaleDownUnneededTime, not even an explicit nil
### GetAutoscalerScaleDownUnreadyTime

`func (o *Autoscaling) GetAutoscalerScaleDownUnreadyTime() int32`

GetAutoscalerScaleDownUnreadyTime returns the AutoscalerScaleDownUnreadyTime field if non-nil, zero value otherwise.

### GetAutoscalerScaleDownUnreadyTimeOk

`func (o *Autoscaling) GetAutoscalerScaleDownUnreadyTimeOk() (*int32, bool)`

GetAutoscalerScaleDownUnreadyTimeOk returns a tuple with the AutoscalerScaleDownUnreadyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscalerScaleDownUnreadyTime

`func (o *Autoscaling) SetAutoscalerScaleDownUnreadyTime(v int32)`

SetAutoscalerScaleDownUnreadyTime sets AutoscalerScaleDownUnreadyTime field to given value.

### HasAutoscalerScaleDownUnreadyTime

`func (o *Autoscaling) HasAutoscalerScaleDownUnreadyTime() bool`

HasAutoscalerScaleDownUnreadyTime returns a boolean if a field has been set.

### SetAutoscalerScaleDownUnreadyTimeNil

`func (o *Autoscaling) SetAutoscalerScaleDownUnreadyTimeNil(b bool)`

 SetAutoscalerScaleDownUnreadyTimeNil sets the value for AutoscalerScaleDownUnreadyTime to be an explicit nil

### UnsetAutoscalerScaleDownUnreadyTime
`func (o *Autoscaling) UnsetAutoscalerScaleDownUnreadyTime()`

UnsetAutoscalerScaleDownUnreadyTime ensures that no value is present for AutoscalerScaleDownUnreadyTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


