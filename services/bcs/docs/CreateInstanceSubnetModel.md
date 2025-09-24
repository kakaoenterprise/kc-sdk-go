# CreateInstanceSubnetModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 인스턴스를 연결할 서브넷의 ID | 
**NetworkInterfaceId** | Pointer to **NullableString** |  | [optional] 
**PrivateIp** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreateInstanceSubnetModel

`func NewCreateInstanceSubnetModel(id string, ) *CreateInstanceSubnetModel`

NewCreateInstanceSubnetModel instantiates a new CreateInstanceSubnetModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInstanceSubnetModelWithDefaults

`func NewCreateInstanceSubnetModelWithDefaults() *CreateInstanceSubnetModel`

NewCreateInstanceSubnetModelWithDefaults instantiates a new CreateInstanceSubnetModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateInstanceSubnetModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateInstanceSubnetModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateInstanceSubnetModel) SetId(v string)`

SetId sets Id field to given value.


### GetNetworkInterfaceId

`func (o *CreateInstanceSubnetModel) GetNetworkInterfaceId() string`

GetNetworkInterfaceId returns the NetworkInterfaceId field if non-nil, zero value otherwise.

### GetNetworkInterfaceIdOk

`func (o *CreateInstanceSubnetModel) GetNetworkInterfaceIdOk() (*string, bool)`

GetNetworkInterfaceIdOk returns a tuple with the NetworkInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaceId

`func (o *CreateInstanceSubnetModel) SetNetworkInterfaceId(v string)`

SetNetworkInterfaceId sets NetworkInterfaceId field to given value.

### HasNetworkInterfaceId

`func (o *CreateInstanceSubnetModel) HasNetworkInterfaceId() bool`

HasNetworkInterfaceId returns a boolean if a field has been set.

### SetNetworkInterfaceIdNil

`func (o *CreateInstanceSubnetModel) SetNetworkInterfaceIdNil(b bool)`

 SetNetworkInterfaceIdNil sets the value for NetworkInterfaceId to be an explicit nil

### UnsetNetworkInterfaceId
`func (o *CreateInstanceSubnetModel) UnsetNetworkInterfaceId()`

UnsetNetworkInterfaceId ensures that no value is present for NetworkInterfaceId, not even an explicit nil
### GetPrivateIp

`func (o *CreateInstanceSubnetModel) GetPrivateIp() string`

GetPrivateIp returns the PrivateIp field if non-nil, zero value otherwise.

### GetPrivateIpOk

`func (o *CreateInstanceSubnetModel) GetPrivateIpOk() (*string, bool)`

GetPrivateIpOk returns a tuple with the PrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIp

`func (o *CreateInstanceSubnetModel) SetPrivateIp(v string)`

SetPrivateIp sets PrivateIp field to given value.

### HasPrivateIp

`func (o *CreateInstanceSubnetModel) HasPrivateIp() bool`

HasPrivateIp returns a boolean if a field has been set.

### SetPrivateIpNil

`func (o *CreateInstanceSubnetModel) SetPrivateIpNil(b bool)`

 SetPrivateIpNil sets the value for PrivateIp to be an explicit nil

### UnsetPrivateIp
`func (o *CreateInstanceSubnetModel) UnsetPrivateIp()`

UnsetPrivateIp ensures that no value is present for PrivateIp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


