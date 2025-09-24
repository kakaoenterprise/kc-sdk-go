# InstanceNetworkInterfaceListModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkInterfaces** | Pointer to [**[]BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel**](BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel.md) |  | [optional] 
**Pagination** | [**PaginationModel**](PaginationModel.md) |  | 

## Methods

### NewInstanceNetworkInterfaceListModel

`func NewInstanceNetworkInterfaceListModel(pagination PaginationModel, ) *InstanceNetworkInterfaceListModel`

NewInstanceNetworkInterfaceListModel instantiates a new InstanceNetworkInterfaceListModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceNetworkInterfaceListModelWithDefaults

`func NewInstanceNetworkInterfaceListModelWithDefaults() *InstanceNetworkInterfaceListModel`

NewInstanceNetworkInterfaceListModelWithDefaults instantiates a new InstanceNetworkInterfaceListModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkInterfaces

`func (o *InstanceNetworkInterfaceListModel) GetNetworkInterfaces() []BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *InstanceNetworkInterfaceListModel) GetNetworkInterfacesOk() (*[]BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *InstanceNetworkInterfaceListModel) SetNetworkInterfaces(v []BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel)`

SetNetworkInterfaces sets NetworkInterfaces field to given value.

### HasNetworkInterfaces

`func (o *InstanceNetworkInterfaceListModel) HasNetworkInterfaces() bool`

HasNetworkInterfaces returns a boolean if a field has been set.

### GetPagination

`func (o *InstanceNetworkInterfaceListModel) GetPagination() PaginationModel`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *InstanceNetworkInterfaceListModel) GetPaginationOk() (*PaginationModel, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *InstanceNetworkInterfaceListModel) SetPagination(v PaginationModel)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


