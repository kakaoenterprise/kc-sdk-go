# BnsLoadBalancerV1ApiGetTargetGroupModelListenerModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 리스너 ID | 
**Protocol** | Pointer to [**NullableProtocol**](Protocol.md) |  | [optional] 
**ProtocolPort** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewBnsLoadBalancerV1ApiGetTargetGroupModelListenerModel

`func NewBnsLoadBalancerV1ApiGetTargetGroupModelListenerModel(id string, ) *BnsLoadBalancerV1ApiGetTargetGroupModelListenerModel`

NewBnsLoadBalancerV1ApiGetTargetGroupModelListenerModel instantiates a new BnsLoadBalancerV1ApiGetTargetGroupModelListenerModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsLoadBalancerV1ApiGetTargetGroupModelListenerModelWithDefaults

`func NewBnsLoadBalancerV1ApiGetTargetGroupModelListenerModelWithDefaults() *BnsLoadBalancerV1ApiGetTargetGroupModelListenerModel`

NewBnsLoadBalancerV1ApiGetTargetGroupModelListenerModelWithDefaults instantiates a new BnsLoadBalancerV1ApiGetTargetGroupModelListenerModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelListenerModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelListenerModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelListenerModel) SetId(v string)`

SetId sets Id field to given value.


### GetProtocol

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelListenerModel) GetProtocol() Protocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelListenerModel) GetProtocolOk() (*Protocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelListenerModel) SetProtocol(v Protocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelListenerModel) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### SetProtocolNil

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelListenerModel) SetProtocolNil(b bool)`

 SetProtocolNil sets the value for Protocol to be an explicit nil

### UnsetProtocol
`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelListenerModel) UnsetProtocol()`

UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil
### GetProtocolPort

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelListenerModel) GetProtocolPort() int32`

GetProtocolPort returns the ProtocolPort field if non-nil, zero value otherwise.

### GetProtocolPortOk

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelListenerModel) GetProtocolPortOk() (*int32, bool)`

GetProtocolPortOk returns a tuple with the ProtocolPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolPort

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelListenerModel) SetProtocolPort(v int32)`

SetProtocolPort sets ProtocolPort field to given value.

### HasProtocolPort

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelListenerModel) HasProtocolPort() bool`

HasProtocolPort returns a boolean if a field has been set.

### SetProtocolPortNil

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelListenerModel) SetProtocolPortNil(b bool)`

 SetProtocolPortNil sets the value for ProtocolPort to be an explicit nil

### UnsetProtocolPort
`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelListenerModel) UnsetProtocolPort()`

UnsetProtocolPort ensures that no value is present for ProtocolPort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


