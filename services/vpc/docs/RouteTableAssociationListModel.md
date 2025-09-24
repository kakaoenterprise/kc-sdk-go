# RouteTableAssociationListModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Associations** | Pointer to [**[]BnsVpcV1ApiListRouteTableAssociationsModelAssociationModel**](BnsVpcV1ApiListRouteTableAssociationsModelAssociationModel.md) |  | [optional] 
**Pagination** | [**PaginationModel**](PaginationModel.md) |  | 

## Methods

### NewRouteTableAssociationListModel

`func NewRouteTableAssociationListModel(pagination PaginationModel, ) *RouteTableAssociationListModel`

NewRouteTableAssociationListModel instantiates a new RouteTableAssociationListModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteTableAssociationListModelWithDefaults

`func NewRouteTableAssociationListModelWithDefaults() *RouteTableAssociationListModel`

NewRouteTableAssociationListModelWithDefaults instantiates a new RouteTableAssociationListModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociations

`func (o *RouteTableAssociationListModel) GetAssociations() []BnsVpcV1ApiListRouteTableAssociationsModelAssociationModel`

GetAssociations returns the Associations field if non-nil, zero value otherwise.

### GetAssociationsOk

`func (o *RouteTableAssociationListModel) GetAssociationsOk() (*[]BnsVpcV1ApiListRouteTableAssociationsModelAssociationModel, bool)`

GetAssociationsOk returns a tuple with the Associations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociations

`func (o *RouteTableAssociationListModel) SetAssociations(v []BnsVpcV1ApiListRouteTableAssociationsModelAssociationModel)`

SetAssociations sets Associations field to given value.

### HasAssociations

`func (o *RouteTableAssociationListModel) HasAssociations() bool`

HasAssociations returns a boolean if a field has been set.

### GetPagination

`func (o *RouteTableAssociationListModel) GetPagination() PaginationModel`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *RouteTableAssociationListModel) GetPaginationOk() (*PaginationModel, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *RouteTableAssociationListModel) SetPagination(v PaginationModel)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


