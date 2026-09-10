# AccessTargets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **NullableInt32** |  | [optional] 
**Requested** | Pointer to **NullableInt32** |  | [optional] 
**Succeeded** | Pointer to **NullableInt32** |  | [optional] 
**Failed** | Pointer to **NullableInt32** |  | [optional] 
**Results** | Pointer to [**[]UpdatedResult**](UpdatedResult.md) |  | [optional] 

## Methods

### NewAccessTargets

`func NewAccessTargets() *AccessTargets`

NewAccessTargets instantiates a new AccessTargets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessTargetsWithDefaults

`func NewAccessTargetsWithDefaults() *AccessTargets`

NewAccessTargetsWithDefaults instantiates a new AccessTargets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *AccessTargets) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AccessTargets) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AccessTargets) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *AccessTargets) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotalNil

`func (o *AccessTargets) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *AccessTargets) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil
### GetRequested

`func (o *AccessTargets) GetRequested() int32`

GetRequested returns the Requested field if non-nil, zero value otherwise.

### GetRequestedOk

`func (o *AccessTargets) GetRequestedOk() (*int32, bool)`

GetRequestedOk returns a tuple with the Requested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequested

`func (o *AccessTargets) SetRequested(v int32)`

SetRequested sets Requested field to given value.

### HasRequested

`func (o *AccessTargets) HasRequested() bool`

HasRequested returns a boolean if a field has been set.

### SetRequestedNil

`func (o *AccessTargets) SetRequestedNil(b bool)`

 SetRequestedNil sets the value for Requested to be an explicit nil

### UnsetRequested
`func (o *AccessTargets) UnsetRequested()`

UnsetRequested ensures that no value is present for Requested, not even an explicit nil
### GetSucceeded

`func (o *AccessTargets) GetSucceeded() int32`

GetSucceeded returns the Succeeded field if non-nil, zero value otherwise.

### GetSucceededOk

`func (o *AccessTargets) GetSucceededOk() (*int32, bool)`

GetSucceededOk returns a tuple with the Succeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSucceeded

`func (o *AccessTargets) SetSucceeded(v int32)`

SetSucceeded sets Succeeded field to given value.

### HasSucceeded

`func (o *AccessTargets) HasSucceeded() bool`

HasSucceeded returns a boolean if a field has been set.

### SetSucceededNil

`func (o *AccessTargets) SetSucceededNil(b bool)`

 SetSucceededNil sets the value for Succeeded to be an explicit nil

### UnsetSucceeded
`func (o *AccessTargets) UnsetSucceeded()`

UnsetSucceeded ensures that no value is present for Succeeded, not even an explicit nil
### GetFailed

`func (o *AccessTargets) GetFailed() int32`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *AccessTargets) GetFailedOk() (*int32, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *AccessTargets) SetFailed(v int32)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *AccessTargets) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### SetFailedNil

`func (o *AccessTargets) SetFailedNil(b bool)`

 SetFailedNil sets the value for Failed to be an explicit nil

### UnsetFailed
`func (o *AccessTargets) UnsetFailed()`

UnsetFailed ensures that no value is present for Failed, not even an explicit nil
### GetResults

`func (o *AccessTargets) GetResults() []UpdatedResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *AccessTargets) GetResultsOk() (*[]UpdatedResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *AccessTargets) SetResults(v []UpdatedResult)`

SetResults sets Results field to given value.

### HasResults

`func (o *AccessTargets) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *AccessTargets) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *AccessTargets) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


