# UpdateHaGroupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BeyondLoadBalancer** | [**BeyondLoadBalancerResult**](BeyondLoadBalancerResult.md) | 수정된 고가용성 그룹 정보 | 

## Methods

### NewUpdateHaGroupResponse

`func NewUpdateHaGroupResponse(beyondLoadBalancer BeyondLoadBalancerResult, ) *UpdateHaGroupResponse`

NewUpdateHaGroupResponse instantiates a new UpdateHaGroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateHaGroupResponseWithDefaults

`func NewUpdateHaGroupResponseWithDefaults() *UpdateHaGroupResponse`

NewUpdateHaGroupResponseWithDefaults instantiates a new UpdateHaGroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBeyondLoadBalancer

`func (o *UpdateHaGroupResponse) GetBeyondLoadBalancer() BeyondLoadBalancerResult`

GetBeyondLoadBalancer returns the BeyondLoadBalancer field if non-nil, zero value otherwise.

### GetBeyondLoadBalancerOk

`func (o *UpdateHaGroupResponse) GetBeyondLoadBalancerOk() (*BeyondLoadBalancerResult, bool)`

GetBeyondLoadBalancerOk returns a tuple with the BeyondLoadBalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeyondLoadBalancer

`func (o *UpdateHaGroupResponse) SetBeyondLoadBalancer(v BeyondLoadBalancerResult)`

SetBeyondLoadBalancer sets BeyondLoadBalancer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


