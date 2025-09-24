# BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelSubnetModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoadBalancerId** | **string** | 각 서브넷에 해당하는 로드 밸런서 인스턴스 ID | 
**AvailabilityZone** | [**AvailabilityZone**](AvailabilityZone.md) | 해당 서브넷이 속한 가용 영역&lt;br/&gt; - &#x60;kr-central-2-a&#x60;: kr-central-2-a 가용 영역 &lt;br/&gt; - &#x60;kr-central-2-b&#x60;: kr-central-2-b 가용 영역 &lt;br/&gt; - &#x60;kr-central-2-c&#x60;: kr-central-2-c 가용 영역 | 

## Methods

### NewBnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelSubnetModel

`func NewBnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelSubnetModel(loadBalancerId string, availabilityZone AvailabilityZone, ) *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelSubnetModel`

NewBnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelSubnetModel instantiates a new BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelSubnetModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelSubnetModelWithDefaults

`func NewBnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelSubnetModelWithDefaults() *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelSubnetModel`

NewBnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelSubnetModelWithDefaults instantiates a new BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelSubnetModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoadBalancerId

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelSubnetModel) GetLoadBalancerId() string`

GetLoadBalancerId returns the LoadBalancerId field if non-nil, zero value otherwise.

### GetLoadBalancerIdOk

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelSubnetModel) GetLoadBalancerIdOk() (*string, bool)`

GetLoadBalancerIdOk returns a tuple with the LoadBalancerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerId

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelSubnetModel) SetLoadBalancerId(v string)`

SetLoadBalancerId sets LoadBalancerId field to given value.


### GetAvailabilityZone

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelSubnetModel) GetAvailabilityZone() AvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelSubnetModel) GetAvailabilityZoneOk() (*AvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelSubnetModel) SetAvailabilityZone(v AvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


