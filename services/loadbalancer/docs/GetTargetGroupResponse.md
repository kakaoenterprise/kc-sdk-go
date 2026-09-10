# GetTargetGroupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetGroup** | [**LoadBalancerPool**](LoadBalancerPool.md) | 대상 그룹 | 

## Methods

### NewGetTargetGroupResponse

`func NewGetTargetGroupResponse(targetGroup LoadBalancerPool, ) *GetTargetGroupResponse`

NewGetTargetGroupResponse instantiates a new GetTargetGroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTargetGroupResponseWithDefaults

`func NewGetTargetGroupResponseWithDefaults() *GetTargetGroupResponse`

NewGetTargetGroupResponseWithDefaults instantiates a new GetTargetGroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetGroup

`func (o *GetTargetGroupResponse) GetTargetGroup() LoadBalancerPool`

GetTargetGroup returns the TargetGroup field if non-nil, zero value otherwise.

### GetTargetGroupOk

`func (o *GetTargetGroupResponse) GetTargetGroupOk() (*LoadBalancerPool, bool)`

GetTargetGroupOk returns a tuple with the TargetGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroup

`func (o *GetTargetGroupResponse) SetTargetGroup(v LoadBalancerPool)`

SetTargetGroup sets TargetGroup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


