# ListenerListModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Listeners** | Pointer to [**[]BnsLoadBalancerV1ApiListListenersModelListenerModel**](BnsLoadBalancerV1ApiListListenersModelListenerModel.md) |  | [optional] 
**Pagination** | [**PaginationModel**](PaginationModel.md) |  | 

## Methods

### NewListenerListModel

`func NewListenerListModel(pagination PaginationModel, ) *ListenerListModel`

NewListenerListModel instantiates a new ListenerListModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListenerListModelWithDefaults

`func NewListenerListModelWithDefaults() *ListenerListModel`

NewListenerListModelWithDefaults instantiates a new ListenerListModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetListeners

`func (o *ListenerListModel) GetListeners() []BnsLoadBalancerV1ApiListListenersModelListenerModel`

GetListeners returns the Listeners field if non-nil, zero value otherwise.

### GetListenersOk

`func (o *ListenerListModel) GetListenersOk() (*[]BnsLoadBalancerV1ApiListListenersModelListenerModel, bool)`

GetListenersOk returns a tuple with the Listeners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListeners

`func (o *ListenerListModel) SetListeners(v []BnsLoadBalancerV1ApiListListenersModelListenerModel)`

SetListeners sets Listeners field to given value.

### HasListeners

`func (o *ListenerListModel) HasListeners() bool`

HasListeners returns a boolean if a field has been set.

### GetPagination

`func (o *ListenerListModel) GetPagination() PaginationModel`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListenerListModel) GetPaginationOk() (*PaginationModel, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListenerListModel) SetPagination(v PaginationModel)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


