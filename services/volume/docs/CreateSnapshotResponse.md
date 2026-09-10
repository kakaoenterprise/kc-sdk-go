# CreateSnapshotResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Snapshot** | [**SnapshotResult**](SnapshotResult.md) | 생성된 스냅샷 정보 | 

## Methods

### NewCreateSnapshotResponse

`func NewCreateSnapshotResponse(snapshot SnapshotResult, ) *CreateSnapshotResponse`

NewCreateSnapshotResponse instantiates a new CreateSnapshotResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSnapshotResponseWithDefaults

`func NewCreateSnapshotResponseWithDefaults() *CreateSnapshotResponse`

NewCreateSnapshotResponseWithDefaults instantiates a new CreateSnapshotResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnapshot

`func (o *CreateSnapshotResponse) GetSnapshot() SnapshotResult`

GetSnapshot returns the Snapshot field if non-nil, zero value otherwise.

### GetSnapshotOk

`func (o *CreateSnapshotResponse) GetSnapshotOk() (*SnapshotResult, bool)`

GetSnapshotOk returns a tuple with the Snapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshot

`func (o *CreateSnapshotResponse) SetSnapshot(v SnapshotResult)`

SetSnapshot sets Snapshot field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


