# CreateK8sClusterNodePoolScalingScheduleRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScheduledScaling** | [**ScheduleRequestModel**](ScheduleRequestModel.md) | 예약 기반 오토스케일링 정보 | 

## Methods

### NewCreateK8sClusterNodePoolScalingScheduleRequestModel

`func NewCreateK8sClusterNodePoolScalingScheduleRequestModel(scheduledScaling ScheduleRequestModel, ) *CreateK8sClusterNodePoolScalingScheduleRequestModel`

NewCreateK8sClusterNodePoolScalingScheduleRequestModel instantiates a new CreateK8sClusterNodePoolScalingScheduleRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateK8sClusterNodePoolScalingScheduleRequestModelWithDefaults

`func NewCreateK8sClusterNodePoolScalingScheduleRequestModelWithDefaults() *CreateK8sClusterNodePoolScalingScheduleRequestModel`

NewCreateK8sClusterNodePoolScalingScheduleRequestModelWithDefaults instantiates a new CreateK8sClusterNodePoolScalingScheduleRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScheduledScaling

`func (o *CreateK8sClusterNodePoolScalingScheduleRequestModel) GetScheduledScaling() ScheduleRequestModel`

GetScheduledScaling returns the ScheduledScaling field if non-nil, zero value otherwise.

### GetScheduledScalingOk

`func (o *CreateK8sClusterNodePoolScalingScheduleRequestModel) GetScheduledScalingOk() (*ScheduleRequestModel, bool)`

GetScheduledScalingOk returns a tuple with the ScheduledScaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledScaling

`func (o *CreateK8sClusterNodePoolScalingScheduleRequestModel) SetScheduledScaling(v ScheduleRequestModel)`

SetScheduledScaling sets ScheduledScaling field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


