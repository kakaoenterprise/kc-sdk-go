# RouteTableListModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VpcRouteTables** | Pointer to [**[]BnsVpcV1ApiListRouteTablesModelRouteTableModel**](BnsVpcV1ApiListRouteTablesModelRouteTableModel.md) |  | [optional] 
**Pagination** | [**PaginationModel**](PaginationModel.md) |  | 

## Methods

### NewRouteTableListModel

`func NewRouteTableListModel(pagination PaginationModel, ) *RouteTableListModel`

NewRouteTableListModel instantiates a new RouteTableListModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteTableListModelWithDefaults

`func NewRouteTableListModelWithDefaults() *RouteTableListModel`

NewRouteTableListModelWithDefaults instantiates a new RouteTableListModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVpcRouteTables

`func (o *RouteTableListModel) GetVpcRouteTables() []BnsVpcV1ApiListRouteTablesModelRouteTableModel`

GetVpcRouteTables returns the VpcRouteTables field if non-nil, zero value otherwise.

### GetVpcRouteTablesOk

`func (o *RouteTableListModel) GetVpcRouteTablesOk() (*[]BnsVpcV1ApiListRouteTablesModelRouteTableModel, bool)`

GetVpcRouteTablesOk returns a tuple with the VpcRouteTables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcRouteTables

`func (o *RouteTableListModel) SetVpcRouteTables(v []BnsVpcV1ApiListRouteTablesModelRouteTableModel)`

SetVpcRouteTables sets VpcRouteTables field to given value.

### HasVpcRouteTables

`func (o *RouteTableListModel) HasVpcRouteTables() bool`

HasVpcRouteTables returns a boolean if a field has been set.

### GetPagination

`func (o *RouteTableListModel) GetPagination() PaginationModel`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *RouteTableListModel) GetPaginationOk() (*PaginationModel, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *RouteTableListModel) SetPagination(v PaginationModel)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


