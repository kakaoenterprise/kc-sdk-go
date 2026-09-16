# GetLoadBalancerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoadBalancer** | [**LoadBalancer**](LoadBalancer.md) | 로드 밸런서 상세 정보 | 

## Methods

### NewGetLoadBalancerResponse

`func NewGetLoadBalancerResponse(loadBalancer LoadBalancer, ) *GetLoadBalancerResponse`

NewGetLoadBalancerResponse instantiates a new GetLoadBalancerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLoadBalancerResponseWithDefaults

`func NewGetLoadBalancerResponseWithDefaults() *GetLoadBalancerResponse`

NewGetLoadBalancerResponseWithDefaults instantiates a new GetLoadBalancerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoadBalancer

`func (o *GetLoadBalancerResponse) GetLoadBalancer() LoadBalancer`

GetLoadBalancer returns the LoadBalancer field if non-nil, zero value otherwise.

### GetLoadBalancerOk

`func (o *GetLoadBalancerResponse) GetLoadBalancerOk() (*LoadBalancer, bool)`

GetLoadBalancerOk returns a tuple with the LoadBalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancer

`func (o *GetLoadBalancerResponse) SetLoadBalancer(v LoadBalancer)`

SetLoadBalancer sets LoadBalancer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


