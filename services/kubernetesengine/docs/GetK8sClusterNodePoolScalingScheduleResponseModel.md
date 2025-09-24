# GetK8sClusterNodePoolScalingScheduleResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScheduledScaling** | [**[]ScheduledScaleResponseModel**](ScheduledScaleResponseModel.md) | 예약 기반 오토스케일링 정보 | 

## Methods

### NewGetK8sClusterNodePoolScalingScheduleResponseModel

`func NewGetK8sClusterNodePoolScalingScheduleResponseModel(scheduledScaling []ScheduledScaleResponseModel, ) *GetK8sClusterNodePoolScalingScheduleResponseModel`

NewGetK8sClusterNodePoolScalingScheduleResponseModel instantiates a new GetK8sClusterNodePoolScalingScheduleResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetK8sClusterNodePoolScalingScheduleResponseModelWithDefaults

`func NewGetK8sClusterNodePoolScalingScheduleResponseModelWithDefaults() *GetK8sClusterNodePoolScalingScheduleResponseModel`

NewGetK8sClusterNodePoolScalingScheduleResponseModelWithDefaults instantiates a new GetK8sClusterNodePoolScalingScheduleResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScheduledScaling

`func (o *GetK8sClusterNodePoolScalingScheduleResponseModel) GetScheduledScaling() []ScheduledScaleResponseModel`

GetScheduledScaling returns the ScheduledScaling field if non-nil, zero value otherwise.

### GetScheduledScalingOk

`func (o *GetK8sClusterNodePoolScalingScheduleResponseModel) GetScheduledScalingOk() (*[]ScheduledScaleResponseModel, bool)`

GetScheduledScalingOk returns a tuple with the ScheduledScaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledScaling

`func (o *GetK8sClusterNodePoolScalingScheduleResponseModel) SetScheduledScaling(v []ScheduledScaleResponseModel)`

SetScheduledScaling sets ScheduledScaling field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


