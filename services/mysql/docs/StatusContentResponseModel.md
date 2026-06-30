# StatusContentResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NeedsRestart** | **bool** | 재시작 필요 여부 | 
**NeedsRestartReason** | **[]string** | 재시작이 필요한 사유 목록 | 

## Methods

### NewStatusContentResponseModel

`func NewStatusContentResponseModel(needsRestart bool, needsRestartReason []string, ) *StatusContentResponseModel`

NewStatusContentResponseModel instantiates a new StatusContentResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusContentResponseModelWithDefaults

`func NewStatusContentResponseModelWithDefaults() *StatusContentResponseModel`

NewStatusContentResponseModelWithDefaults instantiates a new StatusContentResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNeedsRestart

`func (o *StatusContentResponseModel) GetNeedsRestart() bool`

GetNeedsRestart returns the NeedsRestart field if non-nil, zero value otherwise.

### GetNeedsRestartOk

`func (o *StatusContentResponseModel) GetNeedsRestartOk() (*bool, bool)`

GetNeedsRestartOk returns a tuple with the NeedsRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedsRestart

`func (o *StatusContentResponseModel) SetNeedsRestart(v bool)`

SetNeedsRestart sets NeedsRestart field to given value.


### GetNeedsRestartReason

`func (o *StatusContentResponseModel) GetNeedsRestartReason() []string`

GetNeedsRestartReason returns the NeedsRestartReason field if non-nil, zero value otherwise.

### GetNeedsRestartReasonOk

`func (o *StatusContentResponseModel) GetNeedsRestartReasonOk() (*[]string, bool)`

GetNeedsRestartReasonOk returns a tuple with the NeedsRestartReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedsRestartReason

`func (o *StatusContentResponseModel) SetNeedsRestartReason(v []string)`

SetNeedsRestartReason sets NeedsRestartReason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


