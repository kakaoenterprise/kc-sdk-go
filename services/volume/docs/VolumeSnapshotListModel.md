# VolumeSnapshotListModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Snapshots** | Pointer to [**[]BcsVolumeV1ApiListSnapshotsModelVolumeSnapshotModel**](BcsVolumeV1ApiListSnapshotsModelVolumeSnapshotModel.md) |  | [optional] 
**Pagination** | [**PaginationModel**](PaginationModel.md) |  | 

## Methods

### NewVolumeSnapshotListModel

`func NewVolumeSnapshotListModel(pagination PaginationModel, ) *VolumeSnapshotListModel`

NewVolumeSnapshotListModel instantiates a new VolumeSnapshotListModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeSnapshotListModelWithDefaults

`func NewVolumeSnapshotListModelWithDefaults() *VolumeSnapshotListModel`

NewVolumeSnapshotListModelWithDefaults instantiates a new VolumeSnapshotListModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnapshots

`func (o *VolumeSnapshotListModel) GetSnapshots() []BcsVolumeV1ApiListSnapshotsModelVolumeSnapshotModel`

GetSnapshots returns the Snapshots field if non-nil, zero value otherwise.

### GetSnapshotsOk

`func (o *VolumeSnapshotListModel) GetSnapshotsOk() (*[]BcsVolumeV1ApiListSnapshotsModelVolumeSnapshotModel, bool)`

GetSnapshotsOk returns a tuple with the Snapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshots

`func (o *VolumeSnapshotListModel) SetSnapshots(v []BcsVolumeV1ApiListSnapshotsModelVolumeSnapshotModel)`

SetSnapshots sets Snapshots field to given value.

### HasSnapshots

`func (o *VolumeSnapshotListModel) HasSnapshots() bool`

HasSnapshots returns a boolean if a field has been set.

### GetPagination

`func (o *VolumeSnapshotListModel) GetPagination() PaginationModel`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *VolumeSnapshotListModel) GetPaginationOk() (*PaginationModel, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *VolumeSnapshotListModel) SetPagination(v PaginationModel)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


