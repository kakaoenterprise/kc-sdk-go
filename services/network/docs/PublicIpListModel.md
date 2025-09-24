# PublicIpListModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicIps** | Pointer to [**[]BnsNetworkV1ApiListPublicIpsModelFloatingIpModel**](BnsNetworkV1ApiListPublicIpsModelFloatingIpModel.md) |  | [optional] 
**Pagination** | [**PaginationModel**](PaginationModel.md) |  | 

## Methods

### NewPublicIpListModel

`func NewPublicIpListModel(pagination PaginationModel, ) *PublicIpListModel`

NewPublicIpListModel instantiates a new PublicIpListModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicIpListModelWithDefaults

`func NewPublicIpListModelWithDefaults() *PublicIpListModel`

NewPublicIpListModelWithDefaults instantiates a new PublicIpListModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicIps

`func (o *PublicIpListModel) GetPublicIps() []BnsNetworkV1ApiListPublicIpsModelFloatingIpModel`

GetPublicIps returns the PublicIps field if non-nil, zero value otherwise.

### GetPublicIpsOk

`func (o *PublicIpListModel) GetPublicIpsOk() (*[]BnsNetworkV1ApiListPublicIpsModelFloatingIpModel, bool)`

GetPublicIpsOk returns a tuple with the PublicIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIps

`func (o *PublicIpListModel) SetPublicIps(v []BnsNetworkV1ApiListPublicIpsModelFloatingIpModel)`

SetPublicIps sets PublicIps field to given value.

### HasPublicIps

`func (o *PublicIpListModel) HasPublicIps() bool`

HasPublicIps returns a boolean if a field has been set.

### GetPagination

`func (o *PublicIpListModel) GetPagination() PaginationModel`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *PublicIpListModel) GetPaginationOk() (*PaginationModel, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *PublicIpListModel) SetPagination(v PaginationModel)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


