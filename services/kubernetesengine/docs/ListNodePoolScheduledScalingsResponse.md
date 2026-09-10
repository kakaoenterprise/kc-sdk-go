# ListNodePoolScheduledScalingsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScheduledScaling** | [**[]ScheduledScale**](ScheduledScale.md) | 예약 기반 오토스케일링 정보 | 

## Methods

### NewListNodePoolScheduledScalingsResponse

`func NewListNodePoolScheduledScalingsResponse(scheduledScaling []ScheduledScale, ) *ListNodePoolScheduledScalingsResponse`

NewListNodePoolScheduledScalingsResponse instantiates a new ListNodePoolScheduledScalingsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListNodePoolScheduledScalingsResponseWithDefaults

`func NewListNodePoolScheduledScalingsResponseWithDefaults() *ListNodePoolScheduledScalingsResponse`

NewListNodePoolScheduledScalingsResponseWithDefaults instantiates a new ListNodePoolScheduledScalingsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScheduledScaling

`func (o *ListNodePoolScheduledScalingsResponse) GetScheduledScaling() []ScheduledScale`

GetScheduledScaling returns the ScheduledScaling field if non-nil, zero value otherwise.

### GetScheduledScalingOk

`func (o *ListNodePoolScheduledScalingsResponse) GetScheduledScalingOk() (*[]ScheduledScale, bool)`

GetScheduledScalingOk returns a tuple with the ScheduledScaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledScaling

`func (o *ListNodePoolScheduledScalingsResponse) SetScheduledScaling(v []ScheduledScale)`

SetScheduledScaling sets ScheduledScaling field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


