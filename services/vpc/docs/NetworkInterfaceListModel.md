# NetworkInterfaceListModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkInterfaces** | Pointer to [**[]BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel**](BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel.md) |  | [optional] 
**Pagination** | [**PaginationModel**](PaginationModel.md) |  | 

## Methods

### NewNetworkInterfaceListModel

`func NewNetworkInterfaceListModel(pagination PaginationModel, ) *NetworkInterfaceListModel`

NewNetworkInterfaceListModel instantiates a new NetworkInterfaceListModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkInterfaceListModelWithDefaults

`func NewNetworkInterfaceListModelWithDefaults() *NetworkInterfaceListModel`

NewNetworkInterfaceListModelWithDefaults instantiates a new NetworkInterfaceListModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkInterfaces

`func (o *NetworkInterfaceListModel) GetNetworkInterfaces() []BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *NetworkInterfaceListModel) GetNetworkInterfacesOk() (*[]BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *NetworkInterfaceListModel) SetNetworkInterfaces(v []BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel)`

SetNetworkInterfaces sets NetworkInterfaces field to given value.

### HasNetworkInterfaces

`func (o *NetworkInterfaceListModel) HasNetworkInterfaces() bool`

HasNetworkInterfaces returns a boolean if a field has been set.

### GetPagination

`func (o *NetworkInterfaceListModel) GetPagination() PaginationModel`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *NetworkInterfaceListModel) GetPaginationOk() (*PaginationModel, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *NetworkInterfaceListModel) SetPagination(v PaginationModel)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


