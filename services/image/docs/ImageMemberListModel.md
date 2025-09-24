# ImageMemberListModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Members** | Pointer to [**[]BcsImageV1ApiListImageSharedProjectsModelImageMemberModel**](BcsImageV1ApiListImageSharedProjectsModelImageMemberModel.md) |  | [optional] 
**Pagination** | [**PaginationModel**](PaginationModel.md) |  | 

## Methods

### NewImageMemberListModel

`func NewImageMemberListModel(pagination PaginationModel, ) *ImageMemberListModel`

NewImageMemberListModel instantiates a new ImageMemberListModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageMemberListModelWithDefaults

`func NewImageMemberListModelWithDefaults() *ImageMemberListModel`

NewImageMemberListModelWithDefaults instantiates a new ImageMemberListModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMembers

`func (o *ImageMemberListModel) GetMembers() []BcsImageV1ApiListImageSharedProjectsModelImageMemberModel`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *ImageMemberListModel) GetMembersOk() (*[]BcsImageV1ApiListImageSharedProjectsModelImageMemberModel, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *ImageMemberListModel) SetMembers(v []BcsImageV1ApiListImageSharedProjectsModelImageMemberModel)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *ImageMemberListModel) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetPagination

`func (o *ImageMemberListModel) GetPagination() PaginationModel`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ImageMemberListModel) GetPaginationOk() (*PaginationModel, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ImageMemberListModel) SetPagination(v PaginationModel)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


