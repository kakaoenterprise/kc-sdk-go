# VolumeListModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Volumes** | Pointer to [**[]BcsVolumeV1ApiListVolumesModelVolumeModel**](BcsVolumeV1ApiListVolumesModelVolumeModel.md) |  | [optional] 
**Pagination** | [**PaginationModel**](PaginationModel.md) |  | 

## Methods

### NewVolumeListModel

`func NewVolumeListModel(pagination PaginationModel, ) *VolumeListModel`

NewVolumeListModel instantiates a new VolumeListModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeListModelWithDefaults

`func NewVolumeListModelWithDefaults() *VolumeListModel`

NewVolumeListModelWithDefaults instantiates a new VolumeListModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVolumes

`func (o *VolumeListModel) GetVolumes() []BcsVolumeV1ApiListVolumesModelVolumeModel`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *VolumeListModel) GetVolumesOk() (*[]BcsVolumeV1ApiListVolumesModelVolumeModel, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *VolumeListModel) SetVolumes(v []BcsVolumeV1ApiListVolumesModelVolumeModel)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *VolumeListModel) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetPagination

`func (o *VolumeListModel) GetPagination() PaginationModel`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *VolumeListModel) GetPaginationOk() (*PaginationModel, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *VolumeListModel) SetPagination(v PaginationModel)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


