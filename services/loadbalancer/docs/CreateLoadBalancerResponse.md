# CreateLoadBalancerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoadBalancer** | [**CreatedLoadBalancer**](CreatedLoadBalancer.md) | 생성된 로드 밸런서 정보 | 

## Methods

### NewCreateLoadBalancerResponse

`func NewCreateLoadBalancerResponse(loadBalancer CreatedLoadBalancer, ) *CreateLoadBalancerResponse`

NewCreateLoadBalancerResponse instantiates a new CreateLoadBalancerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLoadBalancerResponseWithDefaults

`func NewCreateLoadBalancerResponseWithDefaults() *CreateLoadBalancerResponse`

NewCreateLoadBalancerResponseWithDefaults instantiates a new CreateLoadBalancerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoadBalancer

`func (o *CreateLoadBalancerResponse) GetLoadBalancer() CreatedLoadBalancer`

GetLoadBalancer returns the LoadBalancer field if non-nil, zero value otherwise.

### GetLoadBalancerOk

`func (o *CreateLoadBalancerResponse) GetLoadBalancerOk() (*CreatedLoadBalancer, bool)`

GetLoadBalancerOk returns a tuple with the LoadBalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancer

`func (o *CreateLoadBalancerResponse) SetLoadBalancer(v CreatedLoadBalancer)`

SetLoadBalancer sets LoadBalancer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


