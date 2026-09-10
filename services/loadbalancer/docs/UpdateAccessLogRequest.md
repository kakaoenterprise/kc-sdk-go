# UpdateAccessLogRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessLogs** | Pointer to [**NullableUpdateAccessLog**](UpdateAccessLog.md) |  | [optional] 

## Methods

### NewUpdateAccessLogRequest

`func NewUpdateAccessLogRequest() *UpdateAccessLogRequest`

NewUpdateAccessLogRequest instantiates a new UpdateAccessLogRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAccessLogRequestWithDefaults

`func NewUpdateAccessLogRequestWithDefaults() *UpdateAccessLogRequest`

NewUpdateAccessLogRequestWithDefaults instantiates a new UpdateAccessLogRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessLogs

`func (o *UpdateAccessLogRequest) GetAccessLogs() UpdateAccessLog`

GetAccessLogs returns the AccessLogs field if non-nil, zero value otherwise.

### GetAccessLogsOk

`func (o *UpdateAccessLogRequest) GetAccessLogsOk() (*UpdateAccessLog, bool)`

GetAccessLogsOk returns a tuple with the AccessLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLogs

`func (o *UpdateAccessLogRequest) SetAccessLogs(v UpdateAccessLog)`

SetAccessLogs sets AccessLogs field to given value.

### HasAccessLogs

`func (o *UpdateAccessLogRequest) HasAccessLogs() bool`

HasAccessLogs returns a boolean if a field has been set.

### SetAccessLogsNil

`func (o *UpdateAccessLogRequest) SetAccessLogsNil(b bool)`

 SetAccessLogsNil sets the value for AccessLogs to be an explicit nil

### UnsetAccessLogs
`func (o *UpdateAccessLogRequest) UnsetAccessLogs()`

UnsetAccessLogs ensures that no value is present for AccessLogs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


