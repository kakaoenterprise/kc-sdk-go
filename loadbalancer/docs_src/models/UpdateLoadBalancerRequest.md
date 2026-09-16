# UpdateLoadBalancerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoadBalancer** | [**UpdateLoadBalancer**](UpdateLoadBalancer.md) | 수정할 로드 밸런서 정보 | 

## Methods

### NewUpdateLoadBalancerRequest

`func NewUpdateLoadBalancerRequest(loadBalancer UpdateLoadBalancer, ) *UpdateLoadBalancerRequest`

NewUpdateLoadBalancerRequest instantiates a new UpdateLoadBalancerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateLoadBalancerRequestWithDefaults

`func NewUpdateLoadBalancerRequestWithDefaults() *UpdateLoadBalancerRequest`

NewUpdateLoadBalancerRequestWithDefaults instantiates a new UpdateLoadBalancerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoadBalancer

`func (o *UpdateLoadBalancerRequest) GetLoadBalancer() UpdateLoadBalancer`

GetLoadBalancer returns the LoadBalancer field if non-nil, zero value otherwise.

### GetLoadBalancerOk

`func (o *UpdateLoadBalancerRequest) GetLoadBalancerOk() (*UpdateLoadBalancer, bool)`

GetLoadBalancerOk returns a tuple with the LoadBalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancer

`func (o *UpdateLoadBalancerRequest) SetLoadBalancer(v UpdateLoadBalancer)`

SetLoadBalancer sets LoadBalancer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


