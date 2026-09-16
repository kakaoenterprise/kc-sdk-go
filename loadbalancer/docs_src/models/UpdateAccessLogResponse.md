# UpdateAccessLogResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoadBalancer** | [**UpdatedLoadBalancer**](UpdatedLoadBalancer.md) | 수정된 로드 밸런서 정보 | 

## Methods

### NewUpdateAccessLogResponse

`func NewUpdateAccessLogResponse(loadBalancer UpdatedLoadBalancer, ) *UpdateAccessLogResponse`

NewUpdateAccessLogResponse instantiates a new UpdateAccessLogResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAccessLogResponseWithDefaults

`func NewUpdateAccessLogResponseWithDefaults() *UpdateAccessLogResponse`

NewUpdateAccessLogResponseWithDefaults instantiates a new UpdateAccessLogResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoadBalancer

`func (o *UpdateAccessLogResponse) GetLoadBalancer() UpdatedLoadBalancer`

GetLoadBalancer returns the LoadBalancer field if non-nil, zero value otherwise.

### GetLoadBalancerOk

`func (o *UpdateAccessLogResponse) GetLoadBalancerOk() (*UpdatedLoadBalancer, bool)`

GetLoadBalancerOk returns a tuple with the LoadBalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancer

`func (o *UpdateAccessLogResponse) SetLoadBalancer(v UpdatedLoadBalancer)`

SetLoadBalancer sets LoadBalancer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


