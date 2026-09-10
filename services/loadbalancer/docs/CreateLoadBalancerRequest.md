# CreateLoadBalancerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoadBalancer** | [**CreateLoadBalancer**](CreateLoadBalancer.md) | 생성할 로드 밸런서 정보 | 

## Methods

### NewCreateLoadBalancerRequest

`func NewCreateLoadBalancerRequest(loadBalancer CreateLoadBalancer, ) *CreateLoadBalancerRequest`

NewCreateLoadBalancerRequest instantiates a new CreateLoadBalancerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLoadBalancerRequestWithDefaults

`func NewCreateLoadBalancerRequestWithDefaults() *CreateLoadBalancerRequest`

NewCreateLoadBalancerRequestWithDefaults instantiates a new CreateLoadBalancerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoadBalancer

`func (o *CreateLoadBalancerRequest) GetLoadBalancer() CreateLoadBalancer`

GetLoadBalancer returns the LoadBalancer field if non-nil, zero value otherwise.

### GetLoadBalancerOk

`func (o *CreateLoadBalancerRequest) GetLoadBalancerOk() (*CreateLoadBalancer, bool)`

GetLoadBalancerOk returns a tuple with the LoadBalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancer

`func (o *CreateLoadBalancerRequest) SetLoadBalancer(v CreateLoadBalancer)`

SetLoadBalancer sets LoadBalancer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


