# BnsLoadBalancerV1ApiCreateHaGroupModelSubnetModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoadBalancerId** | **string** | 해당 서브넷에 연결할 로드 밸런서 ID | 
**AvailabilityZone** | [**AvailabilityZone**](AvailabilityZone.md) | 서브넷이 속한 가용 영역&lt;br/&gt; - &#x60;kr-central-2-a&#x60;: kr-central-2-a 가용 영역 &lt;br/&gt; - &#x60;kr-central-2-b&#x60;: kr-central-2-b 가용 영역 &lt;br/&gt; - &#x60;kr-central-2-c&#x60;: kr-central-2-c 가용 영역 | 

## Methods

### NewBnsLoadBalancerV1ApiCreateHaGroupModelSubnetModel

`func NewBnsLoadBalancerV1ApiCreateHaGroupModelSubnetModel(loadBalancerId string, availabilityZone AvailabilityZone, ) *BnsLoadBalancerV1ApiCreateHaGroupModelSubnetModel`

NewBnsLoadBalancerV1ApiCreateHaGroupModelSubnetModel instantiates a new BnsLoadBalancerV1ApiCreateHaGroupModelSubnetModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsLoadBalancerV1ApiCreateHaGroupModelSubnetModelWithDefaults

`func NewBnsLoadBalancerV1ApiCreateHaGroupModelSubnetModelWithDefaults() *BnsLoadBalancerV1ApiCreateHaGroupModelSubnetModel`

NewBnsLoadBalancerV1ApiCreateHaGroupModelSubnetModelWithDefaults instantiates a new BnsLoadBalancerV1ApiCreateHaGroupModelSubnetModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoadBalancerId

`func (o *BnsLoadBalancerV1ApiCreateHaGroupModelSubnetModel) GetLoadBalancerId() string`

GetLoadBalancerId returns the LoadBalancerId field if non-nil, zero value otherwise.

### GetLoadBalancerIdOk

`func (o *BnsLoadBalancerV1ApiCreateHaGroupModelSubnetModel) GetLoadBalancerIdOk() (*string, bool)`

GetLoadBalancerIdOk returns a tuple with the LoadBalancerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerId

`func (o *BnsLoadBalancerV1ApiCreateHaGroupModelSubnetModel) SetLoadBalancerId(v string)`

SetLoadBalancerId sets LoadBalancerId field to given value.


### GetAvailabilityZone

`func (o *BnsLoadBalancerV1ApiCreateHaGroupModelSubnetModel) GetAvailabilityZone() AvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *BnsLoadBalancerV1ApiCreateHaGroupModelSubnetModel) GetAvailabilityZoneOk() (*AvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *BnsLoadBalancerV1ApiCreateHaGroupModelSubnetModel) SetAvailabilityZone(v AvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


