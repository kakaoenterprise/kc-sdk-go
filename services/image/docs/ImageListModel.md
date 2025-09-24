# ImageListModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Images** | Pointer to [**[]BcsImageV1ApiListImagesModelImageModel**](BcsImageV1ApiListImagesModelImageModel.md) |  | [optional] 
**Pagination** | [**PaginationModel**](PaginationModel.md) |  | 

## Methods

### NewImageListModel

`func NewImageListModel(pagination PaginationModel, ) *ImageListModel`

NewImageListModel instantiates a new ImageListModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageListModelWithDefaults

`func NewImageListModelWithDefaults() *ImageListModel`

NewImageListModelWithDefaults instantiates a new ImageListModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImages

`func (o *ImageListModel) GetImages() []BcsImageV1ApiListImagesModelImageModel`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *ImageListModel) GetImagesOk() (*[]BcsImageV1ApiListImagesModelImageModel, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *ImageListModel) SetImages(v []BcsImageV1ApiListImagesModelImageModel)`

SetImages sets Images field to given value.

### HasImages

`func (o *ImageListModel) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetPagination

`func (o *ImageListModel) GetPagination() PaginationModel`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ImageListModel) GetPaginationOk() (*PaginationModel, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ImageListModel) SetPagination(v PaginationModel)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


