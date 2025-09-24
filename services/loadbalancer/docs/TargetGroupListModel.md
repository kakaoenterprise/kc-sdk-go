# TargetGroupListModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetGroups** | Pointer to [**[]BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel**](BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel.md) |  | [optional] 
**Pagination** | [**PaginationModel**](PaginationModel.md) |  | 

## Methods

### NewTargetGroupListModel

`func NewTargetGroupListModel(pagination PaginationModel, ) *TargetGroupListModel`

NewTargetGroupListModel instantiates a new TargetGroupListModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetGroupListModelWithDefaults

`func NewTargetGroupListModelWithDefaults() *TargetGroupListModel`

NewTargetGroupListModelWithDefaults instantiates a new TargetGroupListModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetGroups

`func (o *TargetGroupListModel) GetTargetGroups() []BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel`

GetTargetGroups returns the TargetGroups field if non-nil, zero value otherwise.

### GetTargetGroupsOk

`func (o *TargetGroupListModel) GetTargetGroupsOk() (*[]BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel, bool)`

GetTargetGroupsOk returns a tuple with the TargetGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroups

`func (o *TargetGroupListModel) SetTargetGroups(v []BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel)`

SetTargetGroups sets TargetGroups field to given value.

### HasTargetGroups

`func (o *TargetGroupListModel) HasTargetGroups() bool`

HasTargetGroups returns a boolean if a field has been set.

### GetPagination

`func (o *TargetGroupListModel) GetPagination() PaginationModel`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *TargetGroupListModel) GetPaginationOk() (*PaginationModel, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *TargetGroupListModel) SetPagination(v PaginationModel)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


