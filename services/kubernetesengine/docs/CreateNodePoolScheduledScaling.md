# CreateNodePoolScheduledScaling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | 예약 기반 오토스케일링 규칙 이름 | 
**ScheduleType** | [**SchedulingType**](SchedulingType.md) | 예약 기반 오토스케일링 반복 설정 | 
**Schedule** | Pointer to **NullableString** |  | [optional] 
**DesiredNodes** | **int32** | 예약 기반 오토스케일링이 수행될 때 원하는 노드 수 설정 | 
**StartTime** | **string** | 예약 기반 오토스케일링이 실제로 실행될 수 있는 기준 시각 &lt;br/&gt; - ISO 8601 형식 &lt;br/&gt; - UTC 기준 &lt;br/&gt; - 지정한 &#x60;start_time&#x60; 이후부터 설정된 주기에 따라 규칙 실행 &lt;br/&gt; - 분 단위까지 설정 가능하며, 초 단위는 설정할 수 없음 | 

## Methods

### NewCreateNodePoolScheduledScaling

`func NewCreateNodePoolScheduledScaling(name string, scheduleType SchedulingType, desiredNodes int32, startTime string, ) *CreateNodePoolScheduledScaling`

NewCreateNodePoolScheduledScaling instantiates a new CreateNodePoolScheduledScaling object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNodePoolScheduledScalingWithDefaults

`func NewCreateNodePoolScheduledScalingWithDefaults() *CreateNodePoolScheduledScaling`

NewCreateNodePoolScheduledScalingWithDefaults instantiates a new CreateNodePoolScheduledScaling object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateNodePoolScheduledScaling) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNodePoolScheduledScaling) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNodePoolScheduledScaling) SetName(v string)`

SetName sets Name field to given value.


### GetScheduleType

`func (o *CreateNodePoolScheduledScaling) GetScheduleType() SchedulingType`

GetScheduleType returns the ScheduleType field if non-nil, zero value otherwise.

### GetScheduleTypeOk

`func (o *CreateNodePoolScheduledScaling) GetScheduleTypeOk() (*SchedulingType, bool)`

GetScheduleTypeOk returns a tuple with the ScheduleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleType

`func (o *CreateNodePoolScheduledScaling) SetScheduleType(v SchedulingType)`

SetScheduleType sets ScheduleType field to given value.


### GetSchedule

`func (o *CreateNodePoolScheduledScaling) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *CreateNodePoolScheduledScaling) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *CreateNodePoolScheduledScaling) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *CreateNodePoolScheduledScaling) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### SetScheduleNil

`func (o *CreateNodePoolScheduledScaling) SetScheduleNil(b bool)`

 SetScheduleNil sets the value for Schedule to be an explicit nil

### UnsetSchedule
`func (o *CreateNodePoolScheduledScaling) UnsetSchedule()`

UnsetSchedule ensures that no value is present for Schedule, not even an explicit nil
### GetDesiredNodes

`func (o *CreateNodePoolScheduledScaling) GetDesiredNodes() int32`

GetDesiredNodes returns the DesiredNodes field if non-nil, zero value otherwise.

### GetDesiredNodesOk

`func (o *CreateNodePoolScheduledScaling) GetDesiredNodesOk() (*int32, bool)`

GetDesiredNodesOk returns a tuple with the DesiredNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredNodes

`func (o *CreateNodePoolScheduledScaling) SetDesiredNodes(v int32)`

SetDesiredNodes sets DesiredNodes field to given value.


### GetStartTime

`func (o *CreateNodePoolScheduledScaling) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *CreateNodePoolScheduledScaling) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *CreateNodePoolScheduledScaling) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


