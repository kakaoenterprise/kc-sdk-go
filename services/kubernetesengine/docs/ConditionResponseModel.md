# ConditionResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastHeartbeatTime** | **string** | 마지막 하트비트 수신 시간 | 
**LastTransitionTime** | **string** | 상태가 마지막으로 변경된 시간 | 
**Message** | **string** | 상태 관련 메시지 | 
**Reason** | **string** | 상태 변경 사유 | 
**Status** | **string** | 상태 정보 | 
**Type** | **string** | 상태 유형 | 

## Methods

### NewConditionResponseModel

`func NewConditionResponseModel(lastHeartbeatTime string, lastTransitionTime string, message string, reason string, status string, type_ string, ) *ConditionResponseModel`

NewConditionResponseModel instantiates a new ConditionResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionResponseModelWithDefaults

`func NewConditionResponseModelWithDefaults() *ConditionResponseModel`

NewConditionResponseModelWithDefaults instantiates a new ConditionResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastHeartbeatTime

`func (o *ConditionResponseModel) GetLastHeartbeatTime() string`

GetLastHeartbeatTime returns the LastHeartbeatTime field if non-nil, zero value otherwise.

### GetLastHeartbeatTimeOk

`func (o *ConditionResponseModel) GetLastHeartbeatTimeOk() (*string, bool)`

GetLastHeartbeatTimeOk returns a tuple with the LastHeartbeatTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastHeartbeatTime

`func (o *ConditionResponseModel) SetLastHeartbeatTime(v string)`

SetLastHeartbeatTime sets LastHeartbeatTime field to given value.


### GetLastTransitionTime

`func (o *ConditionResponseModel) GetLastTransitionTime() string`

GetLastTransitionTime returns the LastTransitionTime field if non-nil, zero value otherwise.

### GetLastTransitionTimeOk

`func (o *ConditionResponseModel) GetLastTransitionTimeOk() (*string, bool)`

GetLastTransitionTimeOk returns a tuple with the LastTransitionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTransitionTime

`func (o *ConditionResponseModel) SetLastTransitionTime(v string)`

SetLastTransitionTime sets LastTransitionTime field to given value.


### GetMessage

`func (o *ConditionResponseModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ConditionResponseModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ConditionResponseModel) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetReason

`func (o *ConditionResponseModel) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ConditionResponseModel) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ConditionResponseModel) SetReason(v string)`

SetReason sets Reason field to given value.


### GetStatus

`func (o *ConditionResponseModel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConditionResponseModel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConditionResponseModel) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetType

`func (o *ConditionResponseModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConditionResponseModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConditionResponseModel) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


