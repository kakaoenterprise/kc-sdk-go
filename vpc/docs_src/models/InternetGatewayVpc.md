# InternetGatewayVpc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | VPC의 고유 ID | 
**Name** | Pointer to **NullableString** | VPC 이름 | [optional] 
**CidrBlock** | Pointer to **NullableString** | VPC의 IPv4 CIDR 블록 | [optional] 

## Methods

### NewInternetGatewayVpc

`func NewInternetGatewayVpc(id string, ) *InternetGatewayVpc`

NewInternetGatewayVpc instantiates a new InternetGatewayVpc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternetGatewayVpcWithDefaults

`func NewInternetGatewayVpcWithDefaults() *InternetGatewayVpc`

NewInternetGatewayVpcWithDefaults instantiates a new InternetGatewayVpc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InternetGatewayVpc) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InternetGatewayVpc) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InternetGatewayVpc) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *InternetGatewayVpc) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InternetGatewayVpc) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InternetGatewayVpc) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InternetGatewayVpc) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *InternetGatewayVpc) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *InternetGatewayVpc) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCidrBlock

`func (o *InternetGatewayVpc) GetCidrBlock() string`

GetCidrBlock returns the CidrBlock field if non-nil, zero value otherwise.

### GetCidrBlockOk

`func (o *InternetGatewayVpc) GetCidrBlockOk() (*string, bool)`

GetCidrBlockOk returns a tuple with the CidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrBlock

`func (o *InternetGatewayVpc) SetCidrBlock(v string)`

SetCidrBlock sets CidrBlock field to given value.

### HasCidrBlock

`func (o *InternetGatewayVpc) HasCidrBlock() bool`

HasCidrBlock returns a boolean if a field has been set.

### SetCidrBlockNil

`func (o *InternetGatewayVpc) SetCidrBlockNil(b bool)`

 SetCidrBlockNil sets the value for CidrBlock to be an explicit nil

### UnsetCidrBlock
`func (o *InternetGatewayVpc) UnsetCidrBlock()`

UnsetCidrBlock ensures that no value is present for CidrBlock, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


