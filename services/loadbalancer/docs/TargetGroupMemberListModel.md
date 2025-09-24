# TargetGroupMemberListModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Members** | Pointer to [**[]BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel**](BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel.md) |  | [optional] 
**Pagination** | [**PaginationModel**](PaginationModel.md) |  | 

## Methods

### NewTargetGroupMemberListModel

`func NewTargetGroupMemberListModel(pagination PaginationModel, ) *TargetGroupMemberListModel`

NewTargetGroupMemberListModel instantiates a new TargetGroupMemberListModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetGroupMemberListModelWithDefaults

`func NewTargetGroupMemberListModelWithDefaults() *TargetGroupMemberListModel`

NewTargetGroupMemberListModelWithDefaults instantiates a new TargetGroupMemberListModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMembers

`func (o *TargetGroupMemberListModel) GetMembers() []BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *TargetGroupMemberListModel) GetMembersOk() (*[]BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *TargetGroupMemberListModel) SetMembers(v []BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *TargetGroupMemberListModel) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetPagination

`func (o *TargetGroupMemberListModel) GetPagination() PaginationModel`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *TargetGroupMemberListModel) GetPaginationOk() (*PaginationModel, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *TargetGroupMemberListModel) SetPagination(v PaginationModel)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


