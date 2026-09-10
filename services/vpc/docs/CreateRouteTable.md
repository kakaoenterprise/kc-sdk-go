# CreateRouteTable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | 생성할 라우팅 테이블의 이름 | 
**Description** | Pointer to **NullableString** |  | [optional] 
**VpcId** | **string** | 라우팅 테이블을 생성할 대상 VPC의 ID &lt;br/&gt;- [List VPCs](https://docs.kakaocloud.com/openapi/networking/vpc/list-vpcs)에서 확인 | 

## Methods

### NewCreateRouteTable

`func NewCreateRouteTable(name string, vpcId string, ) *CreateRouteTable`

NewCreateRouteTable instantiates a new CreateRouteTable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRouteTableWithDefaults

`func NewCreateRouteTableWithDefaults() *CreateRouteTable`

NewCreateRouteTableWithDefaults instantiates a new CreateRouteTable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateRouteTable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateRouteTable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateRouteTable) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateRouteTable) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateRouteTable) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateRouteTable) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateRouteTable) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateRouteTable) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateRouteTable) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetVpcId

`func (o *CreateRouteTable) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *CreateRouteTable) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *CreateRouteTable) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


