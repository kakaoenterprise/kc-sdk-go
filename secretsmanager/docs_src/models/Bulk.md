# Bulk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Failed** | Pointer to **NullableInt32** | 처리에 실패한 항목 수 | [optional] 
**Requested** | Pointer to **NullableInt32** | 처리를 요청한 항목 수 | [optional] 
**Results** | Pointer to **[]interface{}** | 항목별 처리 결과 목록 | [optional] 
**Succeeded** | Pointer to **NullableInt32** | 처리에 성공한 항목 수 | [optional] 
**Total** | Pointer to **NullableInt32** | 전체 항목 수 | [optional] 

## Methods

### NewBulk

`func NewBulk() *Bulk`

NewBulk instantiates a new Bulk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWithDefaults

`func NewBulkWithDefaults() *Bulk`

NewBulkWithDefaults instantiates a new Bulk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailed

`func (o *Bulk) GetFailed() int32`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *Bulk) GetFailedOk() (*int32, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *Bulk) SetFailed(v int32)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *Bulk) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### SetFailedNil

`func (o *Bulk) SetFailedNil(b bool)`

 SetFailedNil sets the value for Failed to be an explicit nil

### UnsetFailed
`func (o *Bulk) UnsetFailed()`

UnsetFailed ensures that no value is present for Failed, not even an explicit nil
### GetRequested

`func (o *Bulk) GetRequested() int32`

GetRequested returns the Requested field if non-nil, zero value otherwise.

### GetRequestedOk

`func (o *Bulk) GetRequestedOk() (*int32, bool)`

GetRequestedOk returns a tuple with the Requested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequested

`func (o *Bulk) SetRequested(v int32)`

SetRequested sets Requested field to given value.

### HasRequested

`func (o *Bulk) HasRequested() bool`

HasRequested returns a boolean if a field has been set.

### SetRequestedNil

`func (o *Bulk) SetRequestedNil(b bool)`

 SetRequestedNil sets the value for Requested to be an explicit nil

### UnsetRequested
`func (o *Bulk) UnsetRequested()`

UnsetRequested ensures that no value is present for Requested, not even an explicit nil
### GetResults

`func (o *Bulk) GetResults() []interface{}`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *Bulk) GetResultsOk() (*[]interface{}, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *Bulk) SetResults(v []interface{})`

SetResults sets Results field to given value.

### HasResults

`func (o *Bulk) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *Bulk) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *Bulk) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil
### GetSucceeded

`func (o *Bulk) GetSucceeded() int32`

GetSucceeded returns the Succeeded field if non-nil, zero value otherwise.

### GetSucceededOk

`func (o *Bulk) GetSucceededOk() (*int32, bool)`

GetSucceededOk returns a tuple with the Succeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSucceeded

`func (o *Bulk) SetSucceeded(v int32)`

SetSucceeded sets Succeeded field to given value.

### HasSucceeded

`func (o *Bulk) HasSucceeded() bool`

HasSucceeded returns a boolean if a field has been set.

### SetSucceededNil

`func (o *Bulk) SetSucceededNil(b bool)`

 SetSucceededNil sets the value for Succeeded to be an explicit nil

### UnsetSucceeded
`func (o *Bulk) UnsetSucceeded()`

UnsetSucceeded ensures that no value is present for Succeeded, not even an explicit nil
### GetTotal

`func (o *Bulk) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *Bulk) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *Bulk) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *Bulk) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotalNil

`func (o *Bulk) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *Bulk) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


