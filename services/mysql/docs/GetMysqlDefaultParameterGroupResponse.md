# GetMysqlDefaultParameterGroupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultParameterGroup** | [**GetMysqlDefaultParameterGroupDefaultParameterGroup**](GetMysqlDefaultParameterGroupDefaultParameterGroup.md) | 조회된 MySQL 기본 파라미터 그룹 정보 | 

## Methods

### NewGetMysqlDefaultParameterGroupResponse

`func NewGetMysqlDefaultParameterGroupResponse(defaultParameterGroup GetMysqlDefaultParameterGroupDefaultParameterGroup, ) *GetMysqlDefaultParameterGroupResponse`

NewGetMysqlDefaultParameterGroupResponse instantiates a new GetMysqlDefaultParameterGroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMysqlDefaultParameterGroupResponseWithDefaults

`func NewGetMysqlDefaultParameterGroupResponseWithDefaults() *GetMysqlDefaultParameterGroupResponse`

NewGetMysqlDefaultParameterGroupResponseWithDefaults instantiates a new GetMysqlDefaultParameterGroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultParameterGroup

`func (o *GetMysqlDefaultParameterGroupResponse) GetDefaultParameterGroup() GetMysqlDefaultParameterGroupDefaultParameterGroup`

GetDefaultParameterGroup returns the DefaultParameterGroup field if non-nil, zero value otherwise.

### GetDefaultParameterGroupOk

`func (o *GetMysqlDefaultParameterGroupResponse) GetDefaultParameterGroupOk() (*GetMysqlDefaultParameterGroupDefaultParameterGroup, bool)`

GetDefaultParameterGroupOk returns a tuple with the DefaultParameterGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultParameterGroup

`func (o *GetMysqlDefaultParameterGroupResponse) SetDefaultParameterGroup(v GetMysqlDefaultParameterGroupDefaultParameterGroup)`

SetDefaultParameterGroup sets DefaultParameterGroup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


