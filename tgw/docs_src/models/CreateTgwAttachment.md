# CreateTgwAttachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | 생성할 Transit Gateway Attachment 이름 | [optional] 
**TgwId** | **string** | 연결할 Transit Gateway ID | 
**VpcId** | **string** | 연결할 VPC ID - [List VPCs](/openapi/networking/vpc/list-vpcs)에서 확인 | 
**SubnetIds** | **[]string** | 연결할 서브넷 ID 목록 - [List subnets](/openapi/networking/vpc/list-subnets)에서 확인 | 

## Methods

### NewCreateTgwAttachment

`func NewCreateTgwAttachment(tgwId string, vpcId string, subnetIds []string, ) *CreateTgwAttachment`

NewCreateTgwAttachment instantiates a new CreateTgwAttachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTgwAttachmentWithDefaults

`func NewCreateTgwAttachmentWithDefaults() *CreateTgwAttachment`

NewCreateTgwAttachmentWithDefaults instantiates a new CreateTgwAttachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateTgwAttachment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTgwAttachment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTgwAttachment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateTgwAttachment) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreateTgwAttachment) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateTgwAttachment) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTgwId

`func (o *CreateTgwAttachment) GetTgwId() string`

GetTgwId returns the TgwId field if non-nil, zero value otherwise.

### GetTgwIdOk

`func (o *CreateTgwAttachment) GetTgwIdOk() (*string, bool)`

GetTgwIdOk returns a tuple with the TgwId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgwId

`func (o *CreateTgwAttachment) SetTgwId(v string)`

SetTgwId sets TgwId field to given value.


### GetVpcId

`func (o *CreateTgwAttachment) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *CreateTgwAttachment) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *CreateTgwAttachment) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.


### GetSubnetIds

`func (o *CreateTgwAttachment) GetSubnetIds() []string`

GetSubnetIds returns the SubnetIds field if non-nil, zero value otherwise.

### GetSubnetIdsOk

`func (o *CreateTgwAttachment) GetSubnetIdsOk() (*[]string, bool)`

GetSubnetIdsOk returns a tuple with the SubnetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetIds

`func (o *CreateTgwAttachment) SetSubnetIds(v []string)`

SetSubnetIds sets SubnetIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


