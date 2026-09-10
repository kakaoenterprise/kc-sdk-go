# GetMysqlCustomParameterGroupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomParameterGroup** | [**GetMysqlCustomParameterGroupCustomParameterGroup**](GetMysqlCustomParameterGroupCustomParameterGroup.md) | 조회된 MySQL 커스텀 파라미터 그룹 정보 | 

## Methods

### NewGetMysqlCustomParameterGroupResponse

`func NewGetMysqlCustomParameterGroupResponse(customParameterGroup GetMysqlCustomParameterGroupCustomParameterGroup, ) *GetMysqlCustomParameterGroupResponse`

NewGetMysqlCustomParameterGroupResponse instantiates a new GetMysqlCustomParameterGroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMysqlCustomParameterGroupResponseWithDefaults

`func NewGetMysqlCustomParameterGroupResponseWithDefaults() *GetMysqlCustomParameterGroupResponse`

NewGetMysqlCustomParameterGroupResponseWithDefaults instantiates a new GetMysqlCustomParameterGroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomParameterGroup

`func (o *GetMysqlCustomParameterGroupResponse) GetCustomParameterGroup() GetMysqlCustomParameterGroupCustomParameterGroup`

GetCustomParameterGroup returns the CustomParameterGroup field if non-nil, zero value otherwise.

### GetCustomParameterGroupOk

`func (o *GetMysqlCustomParameterGroupResponse) GetCustomParameterGroupOk() (*GetMysqlCustomParameterGroupCustomParameterGroup, bool)`

GetCustomParameterGroupOk returns a tuple with the CustomParameterGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomParameterGroup

`func (o *GetMysqlCustomParameterGroupResponse) SetCustomParameterGroup(v GetMysqlCustomParameterGroupCustomParameterGroup)`

SetCustomParameterGroup sets CustomParameterGroup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


