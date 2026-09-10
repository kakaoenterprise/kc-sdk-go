# RestoreSnapshotResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Restore** | [**Restore**](Restore.md) | 스냅샷 복원 결과 정보 | 

## Methods

### NewRestoreSnapshotResponse

`func NewRestoreSnapshotResponse(restore Restore, ) *RestoreSnapshotResponse`

NewRestoreSnapshotResponse instantiates a new RestoreSnapshotResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreSnapshotResponseWithDefaults

`func NewRestoreSnapshotResponseWithDefaults() *RestoreSnapshotResponse`

NewRestoreSnapshotResponseWithDefaults instantiates a new RestoreSnapshotResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRestore

`func (o *RestoreSnapshotResponse) GetRestore() Restore`

GetRestore returns the Restore field if non-nil, zero value otherwise.

### GetRestoreOk

`func (o *RestoreSnapshotResponse) GetRestoreOk() (*Restore, bool)`

GetRestoreOk returns a tuple with the Restore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestore

`func (o *RestoreSnapshotResponse) SetRestore(v Restore)`

SetRestore sets Restore field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


