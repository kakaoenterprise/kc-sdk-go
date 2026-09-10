# CreateTargetGroupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetGroup** | [**Pool**](Pool.md) | 생성된 대상 그룹 정보 | 

## Methods

### NewCreateTargetGroupResponse

`func NewCreateTargetGroupResponse(targetGroup Pool, ) *CreateTargetGroupResponse`

NewCreateTargetGroupResponse instantiates a new CreateTargetGroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTargetGroupResponseWithDefaults

`func NewCreateTargetGroupResponseWithDefaults() *CreateTargetGroupResponse`

NewCreateTargetGroupResponseWithDefaults instantiates a new CreateTargetGroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetGroup

`func (o *CreateTargetGroupResponse) GetTargetGroup() Pool`

GetTargetGroup returns the TargetGroup field if non-nil, zero value otherwise.

### GetTargetGroupOk

`func (o *CreateTargetGroupResponse) GetTargetGroupOk() (*Pool, bool)`

GetTargetGroupOk returns a tuple with the TargetGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroup

`func (o *CreateTargetGroupResponse) SetTargetGroup(v Pool)`

SetTargetGroup sets TargetGroup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


