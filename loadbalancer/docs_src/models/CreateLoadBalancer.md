# CreateLoadBalancer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | 로드 밸런서 이름 (프로젝트 내 고유) | 
**Description** | Pointer to **NullableString** | 로드 밸런서에 대한 설명 | [optional] 
**SubnetId** | **string** | 로드 밸런서를 배치할 서브넷 ID - [List subnets](/openapi/networking/vpc/list-subnets)에서 확인 | 
**AvailabilityZone** | [**AvailabilityZone**](AvailabilityZone.md) | 로드 밸런서를 생성할 가용 영역 | 
**FlavorId** | **string** | 로드 밸런서 유형 ID - &#x60;flavors.id&#x60; 값은 [List load balancer types](/openapi/networking/lb/list-load-balancer-types)에서 확인 | 
**Listeners** | Pointer to [**[]CreateLoadBalancerListenerRequest**](CreateLoadBalancerListenerRequest.md) | 로드 밸런서에 설정할 초기 리스너 목록 | [optional] 
**TargetGroups** | Pointer to [**[]CreateLoadBalancerPoolRequest**](CreateLoadBalancerPoolRequest.md) | 초기 대상 그룹 설정 목록 | [optional] 

## Methods

### NewCreateLoadBalancer

`func NewCreateLoadBalancer(name string, subnetId string, availabilityZone AvailabilityZone, flavorId string, ) *CreateLoadBalancer`

NewCreateLoadBalancer instantiates a new CreateLoadBalancer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLoadBalancerWithDefaults

`func NewCreateLoadBalancerWithDefaults() *CreateLoadBalancer`

NewCreateLoadBalancerWithDefaults instantiates a new CreateLoadBalancer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateLoadBalancer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateLoadBalancer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateLoadBalancer) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateLoadBalancer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateLoadBalancer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateLoadBalancer) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateLoadBalancer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateLoadBalancer) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateLoadBalancer) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSubnetId

`func (o *CreateLoadBalancer) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *CreateLoadBalancer) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *CreateLoadBalancer) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.


### GetAvailabilityZone

`func (o *CreateLoadBalancer) GetAvailabilityZone() AvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *CreateLoadBalancer) GetAvailabilityZoneOk() (*AvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *CreateLoadBalancer) SetAvailabilityZone(v AvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.


### GetFlavorId

`func (o *CreateLoadBalancer) GetFlavorId() string`

GetFlavorId returns the FlavorId field if non-nil, zero value otherwise.

### GetFlavorIdOk

`func (o *CreateLoadBalancer) GetFlavorIdOk() (*string, bool)`

GetFlavorIdOk returns a tuple with the FlavorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavorId

`func (o *CreateLoadBalancer) SetFlavorId(v string)`

SetFlavorId sets FlavorId field to given value.


### GetListeners

`func (o *CreateLoadBalancer) GetListeners() []CreateLoadBalancerListenerRequest`

GetListeners returns the Listeners field if non-nil, zero value otherwise.

### GetListenersOk

`func (o *CreateLoadBalancer) GetListenersOk() (*[]CreateLoadBalancerListenerRequest, bool)`

GetListenersOk returns a tuple with the Listeners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListeners

`func (o *CreateLoadBalancer) SetListeners(v []CreateLoadBalancerListenerRequest)`

SetListeners sets Listeners field to given value.

### HasListeners

`func (o *CreateLoadBalancer) HasListeners() bool`

HasListeners returns a boolean if a field has been set.

### SetListenersNil

`func (o *CreateLoadBalancer) SetListenersNil(b bool)`

 SetListenersNil sets the value for Listeners to be an explicit nil

### UnsetListeners
`func (o *CreateLoadBalancer) UnsetListeners()`

UnsetListeners ensures that no value is present for Listeners, not even an explicit nil
### GetTargetGroups

`func (o *CreateLoadBalancer) GetTargetGroups() []CreateLoadBalancerPoolRequest`

GetTargetGroups returns the TargetGroups field if non-nil, zero value otherwise.

### GetTargetGroupsOk

`func (o *CreateLoadBalancer) GetTargetGroupsOk() (*[]CreateLoadBalancerPoolRequest, bool)`

GetTargetGroupsOk returns a tuple with the TargetGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroups

`func (o *CreateLoadBalancer) SetTargetGroups(v []CreateLoadBalancerPoolRequest)`

SetTargetGroups sets TargetGroups field to given value.

### HasTargetGroups

`func (o *CreateLoadBalancer) HasTargetGroups() bool`

HasTargetGroups returns a boolean if a field has been set.

### SetTargetGroupsNil

`func (o *CreateLoadBalancer) SetTargetGroupsNil(b bool)`

 SetTargetGroupsNil sets the value for TargetGroups to be an explicit nil

### UnsetTargetGroups
`func (o *CreateLoadBalancer) UnsetTargetGroups()`

UnsetTargetGroups ensures that no value is present for TargetGroups, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


