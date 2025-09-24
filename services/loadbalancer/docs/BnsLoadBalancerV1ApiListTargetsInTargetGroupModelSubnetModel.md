# BnsLoadBalancerV1ApiListTargetsInTargetGroupModelSubnetModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 서브넷의 고유 ID | 
**Name** | Pointer to **NullableString** |  | [optional] 
**CidrBlock** | Pointer to **NullableString** |  | [optional] 
**AvailabilityZone** | Pointer to [**NullableAvailabilityZone**](AvailabilityZone.md) |  | [optional] 
**HealthCheckIps** | Pointer to **[]string** |  | [optional] 

## Methods

### NewBnsLoadBalancerV1ApiListTargetsInTargetGroupModelSubnetModel

`func NewBnsLoadBalancerV1ApiListTargetsInTargetGroupModelSubnetModel(id string, ) *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelSubnetModel`

NewBnsLoadBalancerV1ApiListTargetsInTargetGroupModelSubnetModel instantiates a new BnsLoadBalancerV1ApiListTargetsInTargetGroupModelSubnetModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsLoadBalancerV1ApiListTargetsInTargetGroupModelSubnetModelWithDefaults

`func NewBnsLoadBalancerV1ApiListTargetsInTargetGroupModelSubnetModelWithDefaults() *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelSubnetModel`

NewBnsLoadBalancerV1ApiListTargetsInTargetGroupModelSubnetModelWithDefaults instantiates a new BnsLoadBalancerV1ApiListTargetsInTargetGroupModelSubnetModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelSubnetModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelSubnetModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelSubnetModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelSubnetModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelSubnetModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelSubnetModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelSubnetModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelSubnetModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelSubnetModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCidrBlock

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelSubnetModel) GetCidrBlock() string`

GetCidrBlock returns the CidrBlock field if non-nil, zero value otherwise.

### GetCidrBlockOk

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelSubnetModel) GetCidrBlockOk() (*string, bool)`

GetCidrBlockOk returns a tuple with the CidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrBlock

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelSubnetModel) SetCidrBlock(v string)`

SetCidrBlock sets CidrBlock field to given value.

### HasCidrBlock

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelSubnetModel) HasCidrBlock() bool`

HasCidrBlock returns a boolean if a field has been set.

### SetCidrBlockNil

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelSubnetModel) SetCidrBlockNil(b bool)`

 SetCidrBlockNil sets the value for CidrBlock to be an explicit nil

### UnsetCidrBlock
`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelSubnetModel) UnsetCidrBlock()`

UnsetCidrBlock ensures that no value is present for CidrBlock, not even an explicit nil
### GetAvailabilityZone

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelSubnetModel) GetAvailabilityZone() AvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelSubnetModel) GetAvailabilityZoneOk() (*AvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelSubnetModel) SetAvailabilityZone(v AvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelSubnetModel) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### SetAvailabilityZoneNil

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelSubnetModel) SetAvailabilityZoneNil(b bool)`

 SetAvailabilityZoneNil sets the value for AvailabilityZone to be an explicit nil

### UnsetAvailabilityZone
`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelSubnetModel) UnsetAvailabilityZone()`

UnsetAvailabilityZone ensures that no value is present for AvailabilityZone, not even an explicit nil
### GetHealthCheckIps

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelSubnetModel) GetHealthCheckIps() []string`

GetHealthCheckIps returns the HealthCheckIps field if non-nil, zero value otherwise.

### GetHealthCheckIpsOk

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelSubnetModel) GetHealthCheckIpsOk() (*[]string, bool)`

GetHealthCheckIpsOk returns a tuple with the HealthCheckIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckIps

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelSubnetModel) SetHealthCheckIps(v []string)`

SetHealthCheckIps sets HealthCheckIps field to given value.

### HasHealthCheckIps

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelSubnetModel) HasHealthCheckIps() bool`

HasHealthCheckIps returns a boolean if a field has been set.

### SetHealthCheckIpsNil

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelSubnetModel) SetHealthCheckIpsNil(b bool)`

 SetHealthCheckIpsNil sets the value for HealthCheckIps to be an explicit nil

### UnsetHealthCheckIps
`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelSubnetModel) UnsetHealthCheckIps()`

UnsetHealthCheckIps ensures that no value is present for HealthCheckIps, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


