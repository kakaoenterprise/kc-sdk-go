# FixedIpModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpAddress** | **string** | 할당된 프라이빗 IP 주소 | 
**SubnetId** | **string** | 해당 IP가 소속된 서브넷의 ID | 

## Methods

### NewFixedIpModel

`func NewFixedIpModel(ipAddress string, subnetId string, ) *FixedIpModel`

NewFixedIpModel instantiates a new FixedIpModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFixedIpModelWithDefaults

`func NewFixedIpModelWithDefaults() *FixedIpModel`

NewFixedIpModelWithDefaults instantiates a new FixedIpModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpAddress

`func (o *FixedIpModel) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *FixedIpModel) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *FixedIpModel) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.


### GetSubnetId

`func (o *FixedIpModel) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *FixedIpModel) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *FixedIpModel) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


