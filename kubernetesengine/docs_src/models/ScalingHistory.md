# ScalingHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | 이력에 대한 설명 | 
**OccurredTime** | **string** | 이벤트 발생 시간 | 
**State** | **string** | 상태값 | 

## Methods

### NewScalingHistory

`func NewScalingHistory(description string, occurredTime string, state string, ) *ScalingHistory`

NewScalingHistory instantiates a new ScalingHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScalingHistoryWithDefaults

`func NewScalingHistoryWithDefaults() *ScalingHistory`

NewScalingHistoryWithDefaults instantiates a new ScalingHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ScalingHistory) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ScalingHistory) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ScalingHistory) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetOccurredTime

`func (o *ScalingHistory) GetOccurredTime() string`

GetOccurredTime returns the OccurredTime field if non-nil, zero value otherwise.

### GetOccurredTimeOk

`func (o *ScalingHistory) GetOccurredTimeOk() (*string, bool)`

GetOccurredTimeOk returns a tuple with the OccurredTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredTime

`func (o *ScalingHistory) SetOccurredTime(v string)`

SetOccurredTime sets OccurredTime field to given value.


### GetState

`func (o *ScalingHistory) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ScalingHistory) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ScalingHistory) SetState(v string)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


