# UpdateSnapshotResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Snapshot** | [**SnapshotResult**](SnapshotResult.md) | 수정된 스냅샷 정보 | 

## Methods

### NewUpdateSnapshotResponse

`func NewUpdateSnapshotResponse(snapshot SnapshotResult, ) *UpdateSnapshotResponse`

NewUpdateSnapshotResponse instantiates a new UpdateSnapshotResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSnapshotResponseWithDefaults

`func NewUpdateSnapshotResponseWithDefaults() *UpdateSnapshotResponse`

NewUpdateSnapshotResponseWithDefaults instantiates a new UpdateSnapshotResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnapshot

`func (o *UpdateSnapshotResponse) GetSnapshot() SnapshotResult`

GetSnapshot returns the Snapshot field if non-nil, zero value otherwise.

### GetSnapshotOk

`func (o *UpdateSnapshotResponse) GetSnapshotOk() (*SnapshotResult, bool)`

GetSnapshotOk returns a tuple with the Snapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshot

`func (o *UpdateSnapshotResponse) SetSnapshot(v SnapshotResult)`

SetSnapshot sets Snapshot field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


