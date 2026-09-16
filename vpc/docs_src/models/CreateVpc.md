# CreateVpc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | 생성할 VPC의 이름 | 
**Description** | Pointer to **NullableString** | VPC에 대한 설명 | [optional] 
**CidrBlock** | **string** | VPC의 IPv4 CIDR 블록 (예: &#x60;10.0.0.0/16&#x60;) | 

## Methods

### NewCreateVpc

`func NewCreateVpc(name string, cidrBlock string, ) *CreateVpc`

NewCreateVpc instantiates a new CreateVpc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateVpcWithDefaults

`func NewCreateVpcWithDefaults() *CreateVpc`

NewCreateVpcWithDefaults instantiates a new CreateVpc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateVpc) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateVpc) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateVpc) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateVpc) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateVpc) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateVpc) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateVpc) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateVpc) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateVpc) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCidrBlock

`func (o *CreateVpc) GetCidrBlock() string`

GetCidrBlock returns the CidrBlock field if non-nil, zero value otherwise.

### GetCidrBlockOk

`func (o *CreateVpc) GetCidrBlockOk() (*string, bool)`

GetCidrBlockOk returns a tuple with the CidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrBlock

`func (o *CreateVpc) SetCidrBlock(v string)`

SetCidrBlock sets CidrBlock field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


