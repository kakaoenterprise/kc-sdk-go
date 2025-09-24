# FlavorListModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Flavors** | Pointer to [**[]BcsInstanceV1ApiListInstanceTypesModelFlavorModel**](BcsInstanceV1ApiListInstanceTypesModelFlavorModel.md) |  | [optional] 
**Pagination** | [**PaginationModel**](PaginationModel.md) |  | 

## Methods

### NewFlavorListModel

`func NewFlavorListModel(pagination PaginationModel, ) *FlavorListModel`

NewFlavorListModel instantiates a new FlavorListModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlavorListModelWithDefaults

`func NewFlavorListModelWithDefaults() *FlavorListModel`

NewFlavorListModelWithDefaults instantiates a new FlavorListModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlavors

`func (o *FlavorListModel) GetFlavors() []BcsInstanceV1ApiListInstanceTypesModelFlavorModel`

GetFlavors returns the Flavors field if non-nil, zero value otherwise.

### GetFlavorsOk

`func (o *FlavorListModel) GetFlavorsOk() (*[]BcsInstanceV1ApiListInstanceTypesModelFlavorModel, bool)`

GetFlavorsOk returns a tuple with the Flavors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavors

`func (o *FlavorListModel) SetFlavors(v []BcsInstanceV1ApiListInstanceTypesModelFlavorModel)`

SetFlavors sets Flavors field to given value.

### HasFlavors

`func (o *FlavorListModel) HasFlavors() bool`

HasFlavors returns a boolean if a field has been set.

### GetPagination

`func (o *FlavorListModel) GetPagination() PaginationModel`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *FlavorListModel) GetPaginationOk() (*PaginationModel, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *FlavorListModel) SetPagination(v PaginationModel)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


