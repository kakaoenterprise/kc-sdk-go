# SecurityGroupListModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecurityGroups** | Pointer to [**[]BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel**](BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel.md) |  | [optional] 
**Pagination** | [**PaginationModel**](PaginationModel.md) |  | 

## Methods

### NewSecurityGroupListModel

`func NewSecurityGroupListModel(pagination PaginationModel, ) *SecurityGroupListModel`

NewSecurityGroupListModel instantiates a new SecurityGroupListModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityGroupListModelWithDefaults

`func NewSecurityGroupListModelWithDefaults() *SecurityGroupListModel`

NewSecurityGroupListModelWithDefaults instantiates a new SecurityGroupListModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecurityGroups

`func (o *SecurityGroupListModel) GetSecurityGroups() []BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *SecurityGroupListModel) GetSecurityGroupsOk() (*[]BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *SecurityGroupListModel) SetSecurityGroups(v []BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *SecurityGroupListModel) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### GetPagination

`func (o *SecurityGroupListModel) GetPagination() PaginationModel`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *SecurityGroupListModel) GetPaginationOk() (*PaginationModel, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *SecurityGroupListModel) SetPagination(v PaginationModel)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


