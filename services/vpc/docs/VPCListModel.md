# VPCListModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vpcs** | Pointer to [**[]BnsVpcV1ApiListVpcsModelVpcModel**](BnsVpcV1ApiListVpcsModelVpcModel.md) |  | [optional] 
**Pagination** | [**PaginationModel**](PaginationModel.md) |  | 

## Methods

### NewVPCListModel

`func NewVPCListModel(pagination PaginationModel, ) *VPCListModel`

NewVPCListModel instantiates a new VPCListModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVPCListModelWithDefaults

`func NewVPCListModelWithDefaults() *VPCListModel`

NewVPCListModelWithDefaults instantiates a new VPCListModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVpcs

`func (o *VPCListModel) GetVpcs() []BnsVpcV1ApiListVpcsModelVpcModel`

GetVpcs returns the Vpcs field if non-nil, zero value otherwise.

### GetVpcsOk

`func (o *VPCListModel) GetVpcsOk() (*[]BnsVpcV1ApiListVpcsModelVpcModel, bool)`

GetVpcsOk returns a tuple with the Vpcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcs

`func (o *VPCListModel) SetVpcs(v []BnsVpcV1ApiListVpcsModelVpcModel)`

SetVpcs sets Vpcs field to given value.

### HasVpcs

`func (o *VPCListModel) HasVpcs() bool`

HasVpcs returns a boolean if a field has been set.

### GetPagination

`func (o *VPCListModel) GetPagination() PaginationModel`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *VPCListModel) GetPaginationOk() (*PaginationModel, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *VPCListModel) SetPagination(v PaginationModel)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


