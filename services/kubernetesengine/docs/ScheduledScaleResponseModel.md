# ScheduledScaleResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | 리소스가 생성된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
**DesiredNodes** | **int32** | 해당 스케줄 실행 시 노드 풀에서 유지할 목표 노드 수 | 
**Name** | **string** | 예약 기반 오토스케일링 규칙 이름 | 
**Schedule** | **string** | 예약 기반 오토스케일링 실행 주기 | 
**ScheduleType** | **string** | 예약 기반 오토스케일링 유형 | 
**StartTime** | **string** | 예약 기반 오토스케일링이 실제로 실행될 수 있는 기준 시각 (ISO 8601, UTC) | 
**Status** | Pointer to [**NullableScalingStatusResponseModel**](ScalingStatusResponseModel.md) |  | [optional] 

## Methods

### NewScheduledScaleResponseModel

`func NewScheduledScaleResponseModel(createdAt time.Time, desiredNodes int32, name string, schedule string, scheduleType string, startTime string, ) *ScheduledScaleResponseModel`

NewScheduledScaleResponseModel instantiates a new ScheduledScaleResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduledScaleResponseModelWithDefaults

`func NewScheduledScaleResponseModelWithDefaults() *ScheduledScaleResponseModel`

NewScheduledScaleResponseModelWithDefaults instantiates a new ScheduledScaleResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ScheduledScaleResponseModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ScheduledScaleResponseModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ScheduledScaleResponseModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDesiredNodes

`func (o *ScheduledScaleResponseModel) GetDesiredNodes() int32`

GetDesiredNodes returns the DesiredNodes field if non-nil, zero value otherwise.

### GetDesiredNodesOk

`func (o *ScheduledScaleResponseModel) GetDesiredNodesOk() (*int32, bool)`

GetDesiredNodesOk returns a tuple with the DesiredNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredNodes

`func (o *ScheduledScaleResponseModel) SetDesiredNodes(v int32)`

SetDesiredNodes sets DesiredNodes field to given value.


### GetName

`func (o *ScheduledScaleResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScheduledScaleResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScheduledScaleResponseModel) SetName(v string)`

SetName sets Name field to given value.


### GetSchedule

`func (o *ScheduledScaleResponseModel) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ScheduledScaleResponseModel) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ScheduledScaleResponseModel) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.


### GetScheduleType

`func (o *ScheduledScaleResponseModel) GetScheduleType() string`

GetScheduleType returns the ScheduleType field if non-nil, zero value otherwise.

### GetScheduleTypeOk

`func (o *ScheduledScaleResponseModel) GetScheduleTypeOk() (*string, bool)`

GetScheduleTypeOk returns a tuple with the ScheduleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleType

`func (o *ScheduledScaleResponseModel) SetScheduleType(v string)`

SetScheduleType sets ScheduleType field to given value.


### GetStartTime

`func (o *ScheduledScaleResponseModel) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ScheduledScaleResponseModel) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ScheduledScaleResponseModel) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.


### GetStatus

`func (o *ScheduledScaleResponseModel) GetStatus() ScalingStatusResponseModel`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ScheduledScaleResponseModel) GetStatusOk() (*ScalingStatusResponseModel, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ScheduledScaleResponseModel) SetStatus(v ScalingStatusResponseModel)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ScheduledScaleResponseModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ScheduledScaleResponseModel) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ScheduledScaleResponseModel) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


