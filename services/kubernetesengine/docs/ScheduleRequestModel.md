# ScheduleRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | 예약 기반 오토스케일링 규칙 이름 | 
**ScheduleType** | [**SchedulingType**](SchedulingType.md) | 예약 기반 오토스케일링 반복 설정&lt;br/&gt;- &#x60;cron&#x60;: 주기적 실행(CRON 표현식 기반)&lt;br/&gt;- &#x60;once&#x60;: 지정된 시작 시간에 한 번만 실행 | 
**Schedule** | Pointer to **NullableString** |  | [optional] 
**DesiredNodes** | **int32** | 예약 기반 오토스케일링이 수행될 때 원하는 노드 수 설정 | 
**StartTime** | **string** | 예약 기반 오토스케일링이 실제로 실행될 수 있는 기준 시각 (ISO 8601, UTC) &lt;br/&gt; - 지정한 &#x60;start_time&#x60; 이후부터 설정된 주기에 따라 규칙 실행 &lt;br/&gt; - 분 단위까지 설정 가능하며, 초 단위는 설정할 수 없음 | 

## Methods

### NewScheduleRequestModel

`func NewScheduleRequestModel(name string, scheduleType SchedulingType, desiredNodes int32, startTime string, ) *ScheduleRequestModel`

NewScheduleRequestModel instantiates a new ScheduleRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleRequestModelWithDefaults

`func NewScheduleRequestModelWithDefaults() *ScheduleRequestModel`

NewScheduleRequestModelWithDefaults instantiates a new ScheduleRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ScheduleRequestModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScheduleRequestModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScheduleRequestModel) SetName(v string)`

SetName sets Name field to given value.


### GetScheduleType

`func (o *ScheduleRequestModel) GetScheduleType() SchedulingType`

GetScheduleType returns the ScheduleType field if non-nil, zero value otherwise.

### GetScheduleTypeOk

`func (o *ScheduleRequestModel) GetScheduleTypeOk() (*SchedulingType, bool)`

GetScheduleTypeOk returns a tuple with the ScheduleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleType

`func (o *ScheduleRequestModel) SetScheduleType(v SchedulingType)`

SetScheduleType sets ScheduleType field to given value.


### GetSchedule

`func (o *ScheduleRequestModel) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ScheduleRequestModel) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ScheduleRequestModel) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *ScheduleRequestModel) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### SetScheduleNil

`func (o *ScheduleRequestModel) SetScheduleNil(b bool)`

 SetScheduleNil sets the value for Schedule to be an explicit nil

### UnsetSchedule
`func (o *ScheduleRequestModel) UnsetSchedule()`

UnsetSchedule ensures that no value is present for Schedule, not even an explicit nil
### GetDesiredNodes

`func (o *ScheduleRequestModel) GetDesiredNodes() int32`

GetDesiredNodes returns the DesiredNodes field if non-nil, zero value otherwise.

### GetDesiredNodesOk

`func (o *ScheduleRequestModel) GetDesiredNodesOk() (*int32, bool)`

GetDesiredNodesOk returns a tuple with the DesiredNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredNodes

`func (o *ScheduleRequestModel) SetDesiredNodes(v int32)`

SetDesiredNodes sets DesiredNodes field to given value.


### GetStartTime

`func (o *ScheduleRequestModel) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ScheduleRequestModel) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ScheduleRequestModel) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


