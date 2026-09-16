# CreateSubnet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailabilityZone** | [**AvailabilityZone**](AvailabilityZone.md) | 서브넷을 생성할 가용 영역 | 
**CidrBlock** | **string** | 서브넷의 IP 주소 범위 (CIDR 형식) - 예: &#x60;10.0.1.0/24&#x60; | 
**Name** | **string** | 생성할 서브넷 이름 | 
**Description** | Pointer to **NullableString** | 서브넷의 용도 또는 구성을 설명하기 위한 추가 정보 | [optional] 
**VpcId** | **string** | 서브넷이 속할 VPC의 ID - [List VPCs](/openapi/networking/vpc/list-vpcs)에서 확인 | 

## Methods

### NewCreateSubnet

`func NewCreateSubnet(availabilityZone AvailabilityZone, cidrBlock string, name string, vpcId string, ) *CreateSubnet`

NewCreateSubnet instantiates a new CreateSubnet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSubnetWithDefaults

`func NewCreateSubnetWithDefaults() *CreateSubnet`

NewCreateSubnetWithDefaults instantiates a new CreateSubnet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailabilityZone

`func (o *CreateSubnet) GetAvailabilityZone() AvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *CreateSubnet) GetAvailabilityZoneOk() (*AvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *CreateSubnet) SetAvailabilityZone(v AvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.


### GetCidrBlock

`func (o *CreateSubnet) GetCidrBlock() string`

GetCidrBlock returns the CidrBlock field if non-nil, zero value otherwise.

### GetCidrBlockOk

`func (o *CreateSubnet) GetCidrBlockOk() (*string, bool)`

GetCidrBlockOk returns a tuple with the CidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrBlock

`func (o *CreateSubnet) SetCidrBlock(v string)`

SetCidrBlock sets CidrBlock field to given value.


### GetName

`func (o *CreateSubnet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateSubnet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateSubnet) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateSubnet) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateSubnet) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateSubnet) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateSubnet) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateSubnet) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateSubnet) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetVpcId

`func (o *CreateSubnet) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *CreateSubnet) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *CreateSubnet) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


