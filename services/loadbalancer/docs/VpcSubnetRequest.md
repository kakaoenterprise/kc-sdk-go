# VpcSubnetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoadBalancerId** | **string** | 각 서브넷에 해당하는 로드 밸런서 인스턴스 ID &lt;br/&gt;- [List load balancers](https://docs.kakaocloud.com/openapi/networking/lb/list-load-balancers)에서 확인 | 
**AvailabilityZone** | [**AvailabilityZone**](AvailabilityZone.md) | 해당 서브넷이 속한 가용 영역 | 

## Methods

### NewVpcSubnetRequest

`func NewVpcSubnetRequest(loadBalancerId string, availabilityZone AvailabilityZone, ) *VpcSubnetRequest`

NewVpcSubnetRequest instantiates a new VpcSubnetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpcSubnetRequestWithDefaults

`func NewVpcSubnetRequestWithDefaults() *VpcSubnetRequest`

NewVpcSubnetRequestWithDefaults instantiates a new VpcSubnetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoadBalancerId

`func (o *VpcSubnetRequest) GetLoadBalancerId() string`

GetLoadBalancerId returns the LoadBalancerId field if non-nil, zero value otherwise.

### GetLoadBalancerIdOk

`func (o *VpcSubnetRequest) GetLoadBalancerIdOk() (*string, bool)`

GetLoadBalancerIdOk returns a tuple with the LoadBalancerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerId

`func (o *VpcSubnetRequest) SetLoadBalancerId(v string)`

SetLoadBalancerId sets LoadBalancerId field to given value.


### GetAvailabilityZone

`func (o *VpcSubnetRequest) GetAvailabilityZone() AvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *VpcSubnetRequest) GetAvailabilityZoneOk() (*AvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *VpcSubnetRequest) SetAvailabilityZone(v AvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


