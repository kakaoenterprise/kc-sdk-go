# UpdateHaGroupLoadBalancerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BeyondLoadBalancer** | [**UpdateHaGroupLoadBalancer**](UpdateHaGroupLoadBalancer.md) | 수정할 고가용성 그룹 서브넷 연결 정보 | 

## Methods

### NewUpdateHaGroupLoadBalancerRequest

`func NewUpdateHaGroupLoadBalancerRequest(beyondLoadBalancer UpdateHaGroupLoadBalancer, ) *UpdateHaGroupLoadBalancerRequest`

NewUpdateHaGroupLoadBalancerRequest instantiates a new UpdateHaGroupLoadBalancerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateHaGroupLoadBalancerRequestWithDefaults

`func NewUpdateHaGroupLoadBalancerRequestWithDefaults() *UpdateHaGroupLoadBalancerRequest`

NewUpdateHaGroupLoadBalancerRequestWithDefaults instantiates a new UpdateHaGroupLoadBalancerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBeyondLoadBalancer

`func (o *UpdateHaGroupLoadBalancerRequest) GetBeyondLoadBalancer() UpdateHaGroupLoadBalancer`

GetBeyondLoadBalancer returns the BeyondLoadBalancer field if non-nil, zero value otherwise.

### GetBeyondLoadBalancerOk

`func (o *UpdateHaGroupLoadBalancerRequest) GetBeyondLoadBalancerOk() (*UpdateHaGroupLoadBalancer, bool)`

GetBeyondLoadBalancerOk returns a tuple with the BeyondLoadBalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeyondLoadBalancer

`func (o *UpdateHaGroupLoadBalancerRequest) SetBeyondLoadBalancer(v UpdateHaGroupLoadBalancer)`

SetBeyondLoadBalancer sets BeyondLoadBalancer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


