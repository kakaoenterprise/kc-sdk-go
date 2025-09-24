# InstanceListModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instances** | [**[]BcsInstanceV1ApiListInstancesModelInstanceModel**](BcsInstanceV1ApiListInstancesModelInstanceModel.md) |  | 
**Pagination** | [**PaginationModel**](PaginationModel.md) |  | 

## Methods

### NewInstanceListModel

`func NewInstanceListModel(instances []BcsInstanceV1ApiListInstancesModelInstanceModel, pagination PaginationModel, ) *InstanceListModel`

NewInstanceListModel instantiates a new InstanceListModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceListModelWithDefaults

`func NewInstanceListModelWithDefaults() *InstanceListModel`

NewInstanceListModelWithDefaults instantiates a new InstanceListModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstances

`func (o *InstanceListModel) GetInstances() []BcsInstanceV1ApiListInstancesModelInstanceModel`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *InstanceListModel) GetInstancesOk() (*[]BcsInstanceV1ApiListInstancesModelInstanceModel, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *InstanceListModel) SetInstances(v []BcsInstanceV1ApiListInstancesModelInstanceModel)`

SetInstances sets Instances field to given value.


### GetPagination

`func (o *InstanceListModel) GetPagination() PaginationModel`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *InstanceListModel) GetPaginationOk() (*PaginationModel, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *InstanceListModel) SetPagination(v PaginationModel)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


