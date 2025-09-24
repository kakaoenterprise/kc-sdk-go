# CreateSubnetModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | 생성할 서브넷 이름 | 
**VpcId** | **string** | 서브넷이 속할 VPC의 ID | 
**CidrBlock** | **string** | 서브넷의 IP 주소 범위 (CIDR 형식)&lt;br/&gt; - 예: &#x60;10.0.1.0/24&#x60; | 
**AvailabilityZone** | [**AvailabilityZone**](AvailabilityZone.md) | 서브넷을 생성할 가용 영역&lt;br/&gt; - &#x60;kr-central-2-a&#x60;: kr-central-2-a 가용 영역 &lt;br/&gt; - &#x60;kr-central-2-b&#x60;: kr-central-2-b 가용 영역 &lt;br/&gt; - &#x60;kr-central-2-c&#x60;: kr-central-2-c 가용 영역 | 

## Methods

### NewCreateSubnetModel

`func NewCreateSubnetModel(name string, vpcId string, cidrBlock string, availabilityZone AvailabilityZone, ) *CreateSubnetModel`

NewCreateSubnetModel instantiates a new CreateSubnetModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSubnetModelWithDefaults

`func NewCreateSubnetModelWithDefaults() *CreateSubnetModel`

NewCreateSubnetModelWithDefaults instantiates a new CreateSubnetModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateSubnetModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateSubnetModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateSubnetModel) SetName(v string)`

SetName sets Name field to given value.


### GetVpcId

`func (o *CreateSubnetModel) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *CreateSubnetModel) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *CreateSubnetModel) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.


### GetCidrBlock

`func (o *CreateSubnetModel) GetCidrBlock() string`

GetCidrBlock returns the CidrBlock field if non-nil, zero value otherwise.

### GetCidrBlockOk

`func (o *CreateSubnetModel) GetCidrBlockOk() (*string, bool)`

GetCidrBlockOk returns a tuple with the CidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrBlock

`func (o *CreateSubnetModel) SetCidrBlock(v string)`

SetCidrBlock sets CidrBlock field to given value.


### GetAvailabilityZone

`func (o *CreateSubnetModel) GetAvailabilityZone() AvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *CreateSubnetModel) GetAvailabilityZoneOk() (*AvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *CreateSubnetModel) SetAvailabilityZone(v AvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


