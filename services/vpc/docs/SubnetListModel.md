# SubnetListModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subnets** | Pointer to [**[]BnsVpcV1ApiListSubnetsModelSubnetModel**](BnsVpcV1ApiListSubnetsModelSubnetModel.md) |  | [optional] 
**Pagination** | [**PaginationModel**](PaginationModel.md) |  | 

## Methods

### NewSubnetListModel

`func NewSubnetListModel(pagination PaginationModel, ) *SubnetListModel`

NewSubnetListModel instantiates a new SubnetListModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubnetListModelWithDefaults

`func NewSubnetListModelWithDefaults() *SubnetListModel`

NewSubnetListModelWithDefaults instantiates a new SubnetListModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubnets

`func (o *SubnetListModel) GetSubnets() []BnsVpcV1ApiListSubnetsModelSubnetModel`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *SubnetListModel) GetSubnetsOk() (*[]BnsVpcV1ApiListSubnetsModelSubnetModel, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *SubnetListModel) SetSubnets(v []BnsVpcV1ApiListSubnetsModelSubnetModel)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *SubnetListModel) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.

### GetPagination

`func (o *SubnetListModel) GetPagination() PaginationModel`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *SubnetListModel) GetPaginationOk() (*PaginationModel, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *SubnetListModel) SetPagination(v PaginationModel)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


