# ListInstanceActionLogsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | [**[]Action**](Action.md) | 인스턴스에 수행한 작업 | 

## Methods

### NewListInstanceActionLogsResponse

`func NewListInstanceActionLogsResponse(actions []Action, ) *ListInstanceActionLogsResponse`

NewListInstanceActionLogsResponse instantiates a new ListInstanceActionLogsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInstanceActionLogsResponseWithDefaults

`func NewListInstanceActionLogsResponseWithDefaults() *ListInstanceActionLogsResponse`

NewListInstanceActionLogsResponseWithDefaults instantiates a new ListInstanceActionLogsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *ListInstanceActionLogsResponse) GetActions() []Action`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ListInstanceActionLogsResponse) GetActionsOk() (*[]Action, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ListInstanceActionLogsResponse) SetActions(v []Action)`

SetActions sets Actions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


