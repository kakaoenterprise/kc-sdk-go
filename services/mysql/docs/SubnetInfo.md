# SubnetInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailabilityZone** | **string** | 서브넷이 속한 가용 영역 | 
**Replicas** | **int32** | 해당 서브넷에 생성된 인스턴스 수 | 
**SubnetId** | **string** | 서브넷 ID | 

## Methods

### NewSubnetInfo

`func NewSubnetInfo(availabilityZone string, replicas int32, subnetId string, ) *SubnetInfo`

NewSubnetInfo instantiates a new SubnetInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubnetInfoWithDefaults

`func NewSubnetInfoWithDefaults() *SubnetInfo`

NewSubnetInfoWithDefaults instantiates a new SubnetInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailabilityZone

`func (o *SubnetInfo) GetAvailabilityZone() string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *SubnetInfo) GetAvailabilityZoneOk() (*string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *SubnetInfo) SetAvailabilityZone(v string)`

SetAvailabilityZone sets AvailabilityZone field to given value.


### GetReplicas

`func (o *SubnetInfo) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *SubnetInfo) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *SubnetInfo) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.


### GetSubnetId

`func (o *SubnetInfo) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *SubnetInfo) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *SubnetInfo) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


