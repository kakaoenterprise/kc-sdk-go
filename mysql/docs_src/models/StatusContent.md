# StatusContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NeedsRestart** | **bool** | 재시작 필요 여부 | 
**NeedsRestartReason** | **[]string** | 재시작이 필요한 사유 목록 | 

## Methods

### NewStatusContent

`func NewStatusContent(needsRestart bool, needsRestartReason []string, ) *StatusContent`

NewStatusContent instantiates a new StatusContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusContentWithDefaults

`func NewStatusContentWithDefaults() *StatusContent`

NewStatusContentWithDefaults instantiates a new StatusContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNeedsRestart

`func (o *StatusContent) GetNeedsRestart() bool`

GetNeedsRestart returns the NeedsRestart field if non-nil, zero value otherwise.

### GetNeedsRestartOk

`func (o *StatusContent) GetNeedsRestartOk() (*bool, bool)`

GetNeedsRestartOk returns a tuple with the NeedsRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedsRestart

`func (o *StatusContent) SetNeedsRestart(v bool)`

SetNeedsRestart sets NeedsRestart field to given value.


### GetNeedsRestartReason

`func (o *StatusContent) GetNeedsRestartReason() []string`

GetNeedsRestartReason returns the NeedsRestartReason field if non-nil, zero value otherwise.

### GetNeedsRestartReasonOk

`func (o *StatusContent) GetNeedsRestartReasonOk() (*[]string, bool)`

GetNeedsRestartReasonOk returns a tuple with the NeedsRestartReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedsRestartReason

`func (o *StatusContent) SetNeedsRestartReason(v []string)`

SetNeedsRestartReason sets NeedsRestartReason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


