# InterfaceAttachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 네트워크 인터페이스의 ID | 
**MacAddress** | **string** | 네트워크 인터페이스의 MAC 주소 | 
**Status** | **string** | 네트워크 인터페이스의 상태 | 
**PrivateIps** | [**[]FixedIp**](FixedIp.md) | 네트워크 인터페이스에 연결된 프라이빗 IP 목록 | 

## Methods

### NewInterfaceAttachment

`func NewInterfaceAttachment(id string, macAddress string, status string, privateIps []FixedIp, ) *InterfaceAttachment`

NewInterfaceAttachment instantiates a new InterfaceAttachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterfaceAttachmentWithDefaults

`func NewInterfaceAttachmentWithDefaults() *InterfaceAttachment`

NewInterfaceAttachmentWithDefaults instantiates a new InterfaceAttachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InterfaceAttachment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InterfaceAttachment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InterfaceAttachment) SetId(v string)`

SetId sets Id field to given value.


### GetMacAddress

`func (o *InterfaceAttachment) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *InterfaceAttachment) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *InterfaceAttachment) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.


### GetStatus

`func (o *InterfaceAttachment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InterfaceAttachment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InterfaceAttachment) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetPrivateIps

`func (o *InterfaceAttachment) GetPrivateIps() []FixedIp`

GetPrivateIps returns the PrivateIps field if non-nil, zero value otherwise.

### GetPrivateIpsOk

`func (o *InterfaceAttachment) GetPrivateIpsOk() (*[]FixedIp, bool)`

GetPrivateIpsOk returns a tuple with the PrivateIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIps

`func (o *InterfaceAttachment) SetPrivateIps(v []FixedIp)`

SetPrivateIps sets PrivateIps field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


