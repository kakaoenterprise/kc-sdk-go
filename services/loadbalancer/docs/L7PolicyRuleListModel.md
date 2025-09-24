# L7PolicyRuleListModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**L7Rules** | Pointer to [**[]BnsLoadBalancerV1ApiListL7PolicyRulesModelL7PolicyRuleModel**](BnsLoadBalancerV1ApiListL7PolicyRulesModelL7PolicyRuleModel.md) |  | [optional] 
**Pagination** | [**PaginationModel**](PaginationModel.md) |  | 

## Methods

### NewL7PolicyRuleListModel

`func NewL7PolicyRuleListModel(pagination PaginationModel, ) *L7PolicyRuleListModel`

NewL7PolicyRuleListModel instantiates a new L7PolicyRuleListModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewL7PolicyRuleListModelWithDefaults

`func NewL7PolicyRuleListModelWithDefaults() *L7PolicyRuleListModel`

NewL7PolicyRuleListModelWithDefaults instantiates a new L7PolicyRuleListModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetL7Rules

`func (o *L7PolicyRuleListModel) GetL7Rules() []BnsLoadBalancerV1ApiListL7PolicyRulesModelL7PolicyRuleModel`

GetL7Rules returns the L7Rules field if non-nil, zero value otherwise.

### GetL7RulesOk

`func (o *L7PolicyRuleListModel) GetL7RulesOk() (*[]BnsLoadBalancerV1ApiListL7PolicyRulesModelL7PolicyRuleModel, bool)`

GetL7RulesOk returns a tuple with the L7Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL7Rules

`func (o *L7PolicyRuleListModel) SetL7Rules(v []BnsLoadBalancerV1ApiListL7PolicyRulesModelL7PolicyRuleModel)`

SetL7Rules sets L7Rules field to given value.

### HasL7Rules

`func (o *L7PolicyRuleListModel) HasL7Rules() bool`

HasL7Rules returns a boolean if a field has been set.

### GetPagination

`func (o *L7PolicyRuleListModel) GetPagination() PaginationModel`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *L7PolicyRuleListModel) GetPaginationOk() (*PaginationModel, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *L7PolicyRuleListModel) SetPagination(v PaginationModel)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


