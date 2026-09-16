# ListMysqlInstancesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instances** | [**[]Instance**](Instance.md) | MySQL 인스턴스 그룹에 포함된 MySQL 인스턴스 목록 | 

## Methods

### NewListMysqlInstancesResponse

`func NewListMysqlInstancesResponse(instances []Instance, ) *ListMysqlInstancesResponse`

NewListMysqlInstancesResponse instantiates a new ListMysqlInstancesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMysqlInstancesResponseWithDefaults

`func NewListMysqlInstancesResponseWithDefaults() *ListMysqlInstancesResponse`

NewListMysqlInstancesResponseWithDefaults instantiates a new ListMysqlInstancesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstances

`func (o *ListMysqlInstancesResponse) GetInstances() []Instance`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *ListMysqlInstancesResponse) GetInstancesOk() (*[]Instance, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *ListMysqlInstancesResponse) SetInstances(v []Instance)`

SetInstances sets Instances field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


