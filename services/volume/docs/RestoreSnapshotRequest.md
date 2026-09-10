# RestoreSnapshotRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Restore** | [**RestoreSnapshot**](RestoreSnapshot.md) | 스냅샷 복원 요청 정보 | 

## Methods

### NewRestoreSnapshotRequest

`func NewRestoreSnapshotRequest(restore RestoreSnapshot, ) *RestoreSnapshotRequest`

NewRestoreSnapshotRequest instantiates a new RestoreSnapshotRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreSnapshotRequestWithDefaults

`func NewRestoreSnapshotRequestWithDefaults() *RestoreSnapshotRequest`

NewRestoreSnapshotRequestWithDefaults instantiates a new RestoreSnapshotRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRestore

`func (o *RestoreSnapshotRequest) GetRestore() RestoreSnapshot`

GetRestore returns the Restore field if non-nil, zero value otherwise.

### GetRestoreOk

`func (o *RestoreSnapshotRequest) GetRestoreOk() (*RestoreSnapshot, bool)`

GetRestoreOk returns a tuple with the Restore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestore

`func (o *RestoreSnapshotRequest) SetRestore(v RestoreSnapshot)`

SetRestore sets Restore field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


