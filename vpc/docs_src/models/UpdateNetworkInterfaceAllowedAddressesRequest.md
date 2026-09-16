# UpdateNetworkInterfaceAllowedAddressesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedAddressPairs** | [**[]UpdateNetworkInterfaceAllowedAddresses**](UpdateNetworkInterfaceAllowedAddresses.md) | 네트워크 인터페이스에 허용할 IP 주소 목록 | 

## Methods

### NewUpdateNetworkInterfaceAllowedAddressesRequest

`func NewUpdateNetworkInterfaceAllowedAddressesRequest(allowedAddressPairs []UpdateNetworkInterfaceAllowedAddresses, ) *UpdateNetworkInterfaceAllowedAddressesRequest`

NewUpdateNetworkInterfaceAllowedAddressesRequest instantiates a new UpdateNetworkInterfaceAllowedAddressesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkInterfaceAllowedAddressesRequestWithDefaults

`func NewUpdateNetworkInterfaceAllowedAddressesRequestWithDefaults() *UpdateNetworkInterfaceAllowedAddressesRequest`

NewUpdateNetworkInterfaceAllowedAddressesRequestWithDefaults instantiates a new UpdateNetworkInterfaceAllowedAddressesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedAddressPairs

`func (o *UpdateNetworkInterfaceAllowedAddressesRequest) GetAllowedAddressPairs() []UpdateNetworkInterfaceAllowedAddresses`

GetAllowedAddressPairs returns the AllowedAddressPairs field if non-nil, zero value otherwise.

### GetAllowedAddressPairsOk

`func (o *UpdateNetworkInterfaceAllowedAddressesRequest) GetAllowedAddressPairsOk() (*[]UpdateNetworkInterfaceAllowedAddresses, bool)`

GetAllowedAddressPairsOk returns a tuple with the AllowedAddressPairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedAddressPairs

`func (o *UpdateNetworkInterfaceAllowedAddressesRequest) SetAllowedAddressPairs(v []UpdateNetworkInterfaceAllowedAddresses)`

SetAllowedAddressPairs sets AllowedAddressPairs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


