# ListSnapshotsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Snapshots** | [**[]Snapshot**](Snapshot.md) | 스냅샷 목록 | 
**Pagination** | [**Pagination**](Pagination.md) | 페이지네이션 정보 | 

## Methods

### NewListSnapshotsResponse

`func NewListSnapshotsResponse(snapshots []Snapshot, pagination Pagination, ) *ListSnapshotsResponse`

NewListSnapshotsResponse instantiates a new ListSnapshotsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSnapshotsResponseWithDefaults

`func NewListSnapshotsResponseWithDefaults() *ListSnapshotsResponse`

NewListSnapshotsResponseWithDefaults instantiates a new ListSnapshotsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnapshots

`func (o *ListSnapshotsResponse) GetSnapshots() []Snapshot`

GetSnapshots returns the Snapshots field if non-nil, zero value otherwise.

### GetSnapshotsOk

`func (o *ListSnapshotsResponse) GetSnapshotsOk() (*[]Snapshot, bool)`

GetSnapshotsOk returns a tuple with the Snapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshots

`func (o *ListSnapshotsResponse) SetSnapshots(v []Snapshot)`

SetSnapshots sets Snapshots field to given value.


### GetPagination

`func (o *ListSnapshotsResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListSnapshotsResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListSnapshotsResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


