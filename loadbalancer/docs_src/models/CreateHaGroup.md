# CreateHaGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | 생성할 고가용성 그룹의 이름 | 
**Description** | Pointer to **NullableString** | 고가용성 그룹에 대한 설명 | [optional] 
**TypeId** | **string** | 사용할 로드 밸런서 유형 ID - [List load balancer types](/openapi/networking/lb/list-load-balancer-types)에서 확인 | 
**Scheme** | [**BeyondLoadBalancerScheme**](BeyondLoadBalancerScheme.md) | 로드 밸런서 통신 방식 | 
**VpcId** | **string** | 고가용성 그룹이 소속될 VPC의 ID - [List VPCs](/openapi/networking/vpc/list-vpcs)에서 확인 | 
**Subnets** | [**[]VpcSubnetRequest**](VpcSubnetRequest.md) | 연결할 서브넷 목록 | 

## Methods

### NewCreateHaGroup

`func NewCreateHaGroup(name string, typeId string, scheme BeyondLoadBalancerScheme, vpcId string, subnets []VpcSubnetRequest, ) *CreateHaGroup`

NewCreateHaGroup instantiates a new CreateHaGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateHaGroupWithDefaults

`func NewCreateHaGroupWithDefaults() *CreateHaGroup`

NewCreateHaGroupWithDefaults instantiates a new CreateHaGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateHaGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateHaGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateHaGroup) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateHaGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateHaGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateHaGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateHaGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateHaGroup) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateHaGroup) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTypeId

`func (o *CreateHaGroup) GetTypeId() string`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *CreateHaGroup) GetTypeIdOk() (*string, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *CreateHaGroup) SetTypeId(v string)`

SetTypeId sets TypeId field to given value.


### GetScheme

`func (o *CreateHaGroup) GetScheme() BeyondLoadBalancerScheme`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *CreateHaGroup) GetSchemeOk() (*BeyondLoadBalancerScheme, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *CreateHaGroup) SetScheme(v BeyondLoadBalancerScheme)`

SetScheme sets Scheme field to given value.


### GetVpcId

`func (o *CreateHaGroup) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *CreateHaGroup) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *CreateHaGroup) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.


### GetSubnets

`func (o *CreateHaGroup) GetSubnets() []VpcSubnetRequest`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *CreateHaGroup) GetSubnetsOk() (*[]VpcSubnetRequest, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *CreateHaGroup) SetSubnets(v []VpcSubnetRequest)`

SetSubnets sets Subnets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


