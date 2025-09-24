# EventResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Component** | **string** | 이벤트를 발생시킨 컴포넌트 | 
**CreationTimestamp** | **string** | 이벤트 발생 시각 | 
**Message** | **string** | 이벤트 메시지 | 
**Reason** | **string** | 이벤트 발생 원인 | 
**Type** | **string** | 이벤트 종류 | 

## Methods

### NewEventResponseModel

`func NewEventResponseModel(component string, creationTimestamp string, message string, reason string, type_ string, ) *EventResponseModel`

NewEventResponseModel instantiates a new EventResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventResponseModelWithDefaults

`func NewEventResponseModelWithDefaults() *EventResponseModel`

NewEventResponseModelWithDefaults instantiates a new EventResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponent

`func (o *EventResponseModel) GetComponent() string`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *EventResponseModel) GetComponentOk() (*string, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *EventResponseModel) SetComponent(v string)`

SetComponent sets Component field to given value.


### GetCreationTimestamp

`func (o *EventResponseModel) GetCreationTimestamp() string`

GetCreationTimestamp returns the CreationTimestamp field if non-nil, zero value otherwise.

### GetCreationTimestampOk

`func (o *EventResponseModel) GetCreationTimestampOk() (*string, bool)`

GetCreationTimestampOk returns a tuple with the CreationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTimestamp

`func (o *EventResponseModel) SetCreationTimestamp(v string)`

SetCreationTimestamp sets CreationTimestamp field to given value.


### GetMessage

`func (o *EventResponseModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *EventResponseModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *EventResponseModel) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetReason

`func (o *EventResponseModel) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *EventResponseModel) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *EventResponseModel) SetReason(v string)`

SetReason sets Reason field to given value.


### GetType

`func (o *EventResponseModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventResponseModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventResponseModel) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


