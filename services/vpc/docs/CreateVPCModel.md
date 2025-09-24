# CreateVPCModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | 생성할 VPC의 이름 | 
**CidrBlock** | **string** | VPC의 IPv4 CIDR 블록 (예: &#x60;10.0.0.0/16&#x60;) | 
**Subnet** | Pointer to [**MainSubnet**](MainSubnet.md) | 서브넷 정보 | [optional] 

## Methods

### NewCreateVPCModel

`func NewCreateVPCModel(name string, cidrBlock string, ) *CreateVPCModel`

NewCreateVPCModel instantiates a new CreateVPCModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVPCModelWithDefaults

`func NewCreateVPCModelWithDefaults() *CreateVPCModel`

NewCreateVPCModelWithDefaults instantiates a new CreateVPCModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateVPCModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateVPCModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateVPCModel) SetName(v string)`

SetName sets Name field to given value.


### GetCidrBlock

`func (o *CreateVPCModel) GetCidrBlock() string`

GetCidrBlock returns the CidrBlock field if non-nil, zero value otherwise.

### GetCidrBlockOk

`func (o *CreateVPCModel) GetCidrBlockOk() (*string, bool)`

GetCidrBlockOk returns a tuple with the CidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrBlock

`func (o *CreateVPCModel) SetCidrBlock(v string)`

SetCidrBlock sets CidrBlock field to given value.


### GetSubnet

`func (o *CreateVPCModel) GetSubnet() MainSubnet`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *CreateVPCModel) GetSubnetOk() (*MainSubnet, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *CreateVPCModel) SetSubnet(v MainSubnet)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *CreateVPCModel) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


