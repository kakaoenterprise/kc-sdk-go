# BcsInstanceV1ApiAttachNetworkInterfaceModelInstanceNetworkInterfaceModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 네트워크 인터페이스의 ID | 
**Status** | **string** | 네트워크 인터페이스의 상태 | 
**MacAddress** | **string** | 네트워크 인터페이스의 MAC 주소 | 
**PrivateIps** | Pointer to [**[]FixedIpModel**](FixedIpModel.md) | 네트워크 인터페이스에 연결된 프라이빗 IP 목록 | [optional] 

## Methods

### NewBcsInstanceV1ApiAttachNetworkInterfaceModelInstanceNetworkInterfaceModel

`func NewBcsInstanceV1ApiAttachNetworkInterfaceModelInstanceNetworkInterfaceModel(id string, status string, macAddress string, ) *BcsInstanceV1ApiAttachNetworkInterfaceModelInstanceNetworkInterfaceModel`

NewBcsInstanceV1ApiAttachNetworkInterfaceModelInstanceNetworkInterfaceModel instantiates a new BcsInstanceV1ApiAttachNetworkInterfaceModelInstanceNetworkInterfaceModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBcsInstanceV1ApiAttachNetworkInterfaceModelInstanceNetworkInterfaceModelWithDefaults

`func NewBcsInstanceV1ApiAttachNetworkInterfaceModelInstanceNetworkInterfaceModelWithDefaults() *BcsInstanceV1ApiAttachNetworkInterfaceModelInstanceNetworkInterfaceModel`

NewBcsInstanceV1ApiAttachNetworkInterfaceModelInstanceNetworkInterfaceModelWithDefaults instantiates a new BcsInstanceV1ApiAttachNetworkInterfaceModelInstanceNetworkInterfaceModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BcsInstanceV1ApiAttachNetworkInterfaceModelInstanceNetworkInterfaceModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BcsInstanceV1ApiAttachNetworkInterfaceModelInstanceNetworkInterfaceModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BcsInstanceV1ApiAttachNetworkInterfaceModelInstanceNetworkInterfaceModel) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *BcsInstanceV1ApiAttachNetworkInterfaceModelInstanceNetworkInterfaceModel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BcsInstanceV1ApiAttachNetworkInterfaceModelInstanceNetworkInterfaceModel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BcsInstanceV1ApiAttachNetworkInterfaceModelInstanceNetworkInterfaceModel) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetMacAddress

`func (o *BcsInstanceV1ApiAttachNetworkInterfaceModelInstanceNetworkInterfaceModel) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *BcsInstanceV1ApiAttachNetworkInterfaceModelInstanceNetworkInterfaceModel) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *BcsInstanceV1ApiAttachNetworkInterfaceModelInstanceNetworkInterfaceModel) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.


### GetPrivateIps

`func (o *BcsInstanceV1ApiAttachNetworkInterfaceModelInstanceNetworkInterfaceModel) GetPrivateIps() []FixedIpModel`

GetPrivateIps returns the PrivateIps field if non-nil, zero value otherwise.

### GetPrivateIpsOk

`func (o *BcsInstanceV1ApiAttachNetworkInterfaceModelInstanceNetworkInterfaceModel) GetPrivateIpsOk() (*[]FixedIpModel, bool)`

GetPrivateIpsOk returns a tuple with the PrivateIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIps

`func (o *BcsInstanceV1ApiAttachNetworkInterfaceModelInstanceNetworkInterfaceModel) SetPrivateIps(v []FixedIpModel)`

SetPrivateIps sets PrivateIps field to given value.

### HasPrivateIps

`func (o *BcsInstanceV1ApiAttachNetworkInterfaceModelInstanceNetworkInterfaceModel) HasPrivateIps() bool`

HasPrivateIps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


