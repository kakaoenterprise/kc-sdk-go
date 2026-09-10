# CreateNodePoolScheduledScalingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScheduledScaling** | [**CreateNodePoolScheduledScaling**](CreateNodePoolScheduledScaling.md) | 예약 기반 오토스케일링 정보 | 

## Methods

### NewCreateNodePoolScheduledScalingRequest

`func NewCreateNodePoolScheduledScalingRequest(scheduledScaling CreateNodePoolScheduledScaling, ) *CreateNodePoolScheduledScalingRequest`

NewCreateNodePoolScheduledScalingRequest instantiates a new CreateNodePoolScheduledScalingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNodePoolScheduledScalingRequestWithDefaults

`func NewCreateNodePoolScheduledScalingRequestWithDefaults() *CreateNodePoolScheduledScalingRequest`

NewCreateNodePoolScheduledScalingRequestWithDefaults instantiates a new CreateNodePoolScheduledScalingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScheduledScaling

`func (o *CreateNodePoolScheduledScalingRequest) GetScheduledScaling() CreateNodePoolScheduledScaling`

GetScheduledScaling returns the ScheduledScaling field if non-nil, zero value otherwise.

### GetScheduledScalingOk

`func (o *CreateNodePoolScheduledScalingRequest) GetScheduledScalingOk() (*CreateNodePoolScheduledScaling, bool)`

GetScheduledScalingOk returns a tuple with the ScheduledScaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledScaling

`func (o *CreateNodePoolScheduledScalingRequest) SetScheduledScaling(v CreateNodePoolScheduledScaling)`

SetScheduledScaling sets ScheduledScaling field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


