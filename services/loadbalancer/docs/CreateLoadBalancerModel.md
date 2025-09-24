# CreateLoadBalancerModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | 로드 밸런서 이름 (프로젝트 내 고유) | 
**Description** | Pointer to **NullableString** |  | [optional] 
**SubnetId** | **string** | 로드 밸런서를 배치할 서브넷 ID | 
**AvailabilityZone** | [**AvailabilityZone**](AvailabilityZone.md) | 로드 밸런서를 생성할 가용 영역&lt;br/&gt; - &#x60;kr-central-2-a&#x60;: kr-central-2-a 가용 영역 &lt;br/&gt; - &#x60;kr-central-2-b&#x60;: kr-central-2-b 가용 영역 &lt;br/&gt; - &#x60;kr-central-2-c&#x60;: kr-central-2-c 가용 영역 | 
**FlavorId** | **string** | 로드 밸런서 유형 &lt;br/&gt; - [List load balancer types](https://docs.kakaocloud.com/openapi/bns/lb/list-load-balancer-types) API에서 조회한 &#x60;flavors.id&#x60; 사용 | 
**Listeners** | Pointer to [**[]ListenerModelInput**](ListenerModelInput.md) |  | [optional] 

## Methods

### NewCreateLoadBalancerModel

`func NewCreateLoadBalancerModel(name string, subnetId string, availabilityZone AvailabilityZone, flavorId string, ) *CreateLoadBalancerModel`

NewCreateLoadBalancerModel instantiates a new CreateLoadBalancerModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLoadBalancerModelWithDefaults

`func NewCreateLoadBalancerModelWithDefaults() *CreateLoadBalancerModel`

NewCreateLoadBalancerModelWithDefaults instantiates a new CreateLoadBalancerModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateLoadBalancerModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateLoadBalancerModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateLoadBalancerModel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateLoadBalancerModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateLoadBalancerModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateLoadBalancerModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateLoadBalancerModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateLoadBalancerModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateLoadBalancerModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSubnetId

`func (o *CreateLoadBalancerModel) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *CreateLoadBalancerModel) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *CreateLoadBalancerModel) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.


### GetAvailabilityZone

`func (o *CreateLoadBalancerModel) GetAvailabilityZone() AvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *CreateLoadBalancerModel) GetAvailabilityZoneOk() (*AvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *CreateLoadBalancerModel) SetAvailabilityZone(v AvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.


### GetFlavorId

`func (o *CreateLoadBalancerModel) GetFlavorId() string`

GetFlavorId returns the FlavorId field if non-nil, zero value otherwise.

### GetFlavorIdOk

`func (o *CreateLoadBalancerModel) GetFlavorIdOk() (*string, bool)`

GetFlavorIdOk returns a tuple with the FlavorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavorId

`func (o *CreateLoadBalancerModel) SetFlavorId(v string)`

SetFlavorId sets FlavorId field to given value.


### GetListeners

`func (o *CreateLoadBalancerModel) GetListeners() []ListenerModelInput`

GetListeners returns the Listeners field if non-nil, zero value otherwise.

### GetListenersOk

`func (o *CreateLoadBalancerModel) GetListenersOk() (*[]ListenerModelInput, bool)`

GetListenersOk returns a tuple with the Listeners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListeners

`func (o *CreateLoadBalancerModel) SetListeners(v []ListenerModelInput)`

SetListeners sets Listeners field to given value.

### HasListeners

`func (o *CreateLoadBalancerModel) HasListeners() bool`

HasListeners returns a boolean if a field has been set.

### SetListenersNil

`func (o *CreateLoadBalancerModel) SetListenersNil(b bool)`

 SetListenersNil sets the value for Listeners to be an explicit nil

### UnsetListeners
`func (o *CreateLoadBalancerModel) UnsetListeners()`

UnsetListeners ensures that no value is present for Listeners, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


