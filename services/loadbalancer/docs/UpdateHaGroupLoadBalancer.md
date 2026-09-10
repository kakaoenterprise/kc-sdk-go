# UpdateHaGroupLoadBalancer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subnets** | [**[]VpcSubnetRequest**](VpcSubnetRequest.md) | 고가용성 로드 밸런서에 연결할 서브넷 목록 | 

## Methods

### NewUpdateHaGroupLoadBalancer

`func NewUpdateHaGroupLoadBalancer(subnets []VpcSubnetRequest, ) *UpdateHaGroupLoadBalancer`

NewUpdateHaGroupLoadBalancer instantiates a new UpdateHaGroupLoadBalancer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateHaGroupLoadBalancerWithDefaults

`func NewUpdateHaGroupLoadBalancerWithDefaults() *UpdateHaGroupLoadBalancer`

NewUpdateHaGroupLoadBalancerWithDefaults instantiates a new UpdateHaGroupLoadBalancer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubnets

`func (o *UpdateHaGroupLoadBalancer) GetSubnets() []VpcSubnetRequest`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *UpdateHaGroupLoadBalancer) GetSubnetsOk() (*[]VpcSubnetRequest, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *UpdateHaGroupLoadBalancer) SetSubnets(v []VpcSubnetRequest)`

SetSubnets sets Subnets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


