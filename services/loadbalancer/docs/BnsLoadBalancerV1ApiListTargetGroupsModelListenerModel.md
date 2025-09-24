# BnsLoadBalancerV1ApiListTargetGroupsModelListenerModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 리스너 ID | 
**Protocol** | Pointer to [**NullableProtocol**](Protocol.md) |  | [optional] 
**ProtocolPort** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewBnsLoadBalancerV1ApiListTargetGroupsModelListenerModel

`func NewBnsLoadBalancerV1ApiListTargetGroupsModelListenerModel(id string, ) *BnsLoadBalancerV1ApiListTargetGroupsModelListenerModel`

NewBnsLoadBalancerV1ApiListTargetGroupsModelListenerModel instantiates a new BnsLoadBalancerV1ApiListTargetGroupsModelListenerModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsLoadBalancerV1ApiListTargetGroupsModelListenerModelWithDefaults

`func NewBnsLoadBalancerV1ApiListTargetGroupsModelListenerModelWithDefaults() *BnsLoadBalancerV1ApiListTargetGroupsModelListenerModel`

NewBnsLoadBalancerV1ApiListTargetGroupsModelListenerModelWithDefaults instantiates a new BnsLoadBalancerV1ApiListTargetGroupsModelListenerModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelListenerModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelListenerModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelListenerModel) SetId(v string)`

SetId sets Id field to given value.


### GetProtocol

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelListenerModel) GetProtocol() Protocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelListenerModel) GetProtocolOk() (*Protocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelListenerModel) SetProtocol(v Protocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelListenerModel) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### SetProtocolNil

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelListenerModel) SetProtocolNil(b bool)`

 SetProtocolNil sets the value for Protocol to be an explicit nil

### UnsetProtocol
`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelListenerModel) UnsetProtocol()`

UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil
### GetProtocolPort

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelListenerModel) GetProtocolPort() int32`

GetProtocolPort returns the ProtocolPort field if non-nil, zero value otherwise.

### GetProtocolPortOk

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelListenerModel) GetProtocolPortOk() (*int32, bool)`

GetProtocolPortOk returns a tuple with the ProtocolPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolPort

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelListenerModel) SetProtocolPort(v int32)`

SetProtocolPort sets ProtocolPort field to given value.

### HasProtocolPort

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelListenerModel) HasProtocolPort() bool`

HasProtocolPort returns a boolean if a field has been set.

### SetProtocolPortNil

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelListenerModel) SetProtocolPortNil(b bool)`

 SetProtocolPortNil sets the value for ProtocolPort to be an explicit nil

### UnsetProtocolPort
`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelListenerModel) UnsetProtocolPort()`

UnsetProtocolPort ensures that no value is present for ProtocolPort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


