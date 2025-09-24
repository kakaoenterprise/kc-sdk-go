# BnsLoadBalancerV1ApiCreateHaGroupModelCreateBeyondLoadBalancerModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | 생성할 고가용성 그룹의 이름 | 
**Description** | Pointer to **NullableString** |  | [optional] 
**TypeId** | **string** | 사용할 로드 밸런서 유형 ID | 
**Scheme** | [**Scheme**](Scheme.md) | 로드 밸런서 통신 방식 &lt;br/&gt; - &#x60;internet-facing&#x60;: 인터넷과 연결됨 &lt;br/&gt; - &#x60;internal&#x60;: 내부 전용 | 
**VpcId** | **string** | 고가용성 그룹이 소속될 VPC의 ID | 
**Subnets** | [**[]BnsLoadBalancerV1ApiCreateHaGroupModelSubnetModel**](BnsLoadBalancerV1ApiCreateHaGroupModelSubnetModel.md) | 연결할 서브넷 목록 | 

## Methods

### NewBnsLoadBalancerV1ApiCreateHaGroupModelCreateBeyondLoadBalancerModel

`func NewBnsLoadBalancerV1ApiCreateHaGroupModelCreateBeyondLoadBalancerModel(name string, typeId string, scheme Scheme, vpcId string, subnets []BnsLoadBalancerV1ApiCreateHaGroupModelSubnetModel, ) *BnsLoadBalancerV1ApiCreateHaGroupModelCreateBeyondLoadBalancerModel`

NewBnsLoadBalancerV1ApiCreateHaGroupModelCreateBeyondLoadBalancerModel instantiates a new BnsLoadBalancerV1ApiCreateHaGroupModelCreateBeyondLoadBalancerModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsLoadBalancerV1ApiCreateHaGroupModelCreateBeyondLoadBalancerModelWithDefaults

`func NewBnsLoadBalancerV1ApiCreateHaGroupModelCreateBeyondLoadBalancerModelWithDefaults() *BnsLoadBalancerV1ApiCreateHaGroupModelCreateBeyondLoadBalancerModel`

NewBnsLoadBalancerV1ApiCreateHaGroupModelCreateBeyondLoadBalancerModelWithDefaults instantiates a new BnsLoadBalancerV1ApiCreateHaGroupModelCreateBeyondLoadBalancerModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BnsLoadBalancerV1ApiCreateHaGroupModelCreateBeyondLoadBalancerModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BnsLoadBalancerV1ApiCreateHaGroupModelCreateBeyondLoadBalancerModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BnsLoadBalancerV1ApiCreateHaGroupModelCreateBeyondLoadBalancerModel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *BnsLoadBalancerV1ApiCreateHaGroupModelCreateBeyondLoadBalancerModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BnsLoadBalancerV1ApiCreateHaGroupModelCreateBeyondLoadBalancerModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BnsLoadBalancerV1ApiCreateHaGroupModelCreateBeyondLoadBalancerModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BnsLoadBalancerV1ApiCreateHaGroupModelCreateBeyondLoadBalancerModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BnsLoadBalancerV1ApiCreateHaGroupModelCreateBeyondLoadBalancerModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BnsLoadBalancerV1ApiCreateHaGroupModelCreateBeyondLoadBalancerModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTypeId

`func (o *BnsLoadBalancerV1ApiCreateHaGroupModelCreateBeyondLoadBalancerModel) GetTypeId() string`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *BnsLoadBalancerV1ApiCreateHaGroupModelCreateBeyondLoadBalancerModel) GetTypeIdOk() (*string, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *BnsLoadBalancerV1ApiCreateHaGroupModelCreateBeyondLoadBalancerModel) SetTypeId(v string)`

SetTypeId sets TypeId field to given value.


### GetScheme

`func (o *BnsLoadBalancerV1ApiCreateHaGroupModelCreateBeyondLoadBalancerModel) GetScheme() Scheme`

GetScheme returns the Scheme field if non-nil, zero value otherwise.

### GetSchemeOk

`func (o *BnsLoadBalancerV1ApiCreateHaGroupModelCreateBeyondLoadBalancerModel) GetSchemeOk() (*Scheme, bool)`

GetSchemeOk returns a tuple with the Scheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheme

`func (o *BnsLoadBalancerV1ApiCreateHaGroupModelCreateBeyondLoadBalancerModel) SetScheme(v Scheme)`

SetScheme sets Scheme field to given value.


### GetVpcId

`func (o *BnsLoadBalancerV1ApiCreateHaGroupModelCreateBeyondLoadBalancerModel) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *BnsLoadBalancerV1ApiCreateHaGroupModelCreateBeyondLoadBalancerModel) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *BnsLoadBalancerV1ApiCreateHaGroupModelCreateBeyondLoadBalancerModel) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.


### GetSubnets

`func (o *BnsLoadBalancerV1ApiCreateHaGroupModelCreateBeyondLoadBalancerModel) GetSubnets() []BnsLoadBalancerV1ApiCreateHaGroupModelSubnetModel`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *BnsLoadBalancerV1ApiCreateHaGroupModelCreateBeyondLoadBalancerModel) GetSubnetsOk() (*[]BnsLoadBalancerV1ApiCreateHaGroupModelSubnetModel, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *BnsLoadBalancerV1ApiCreateHaGroupModelCreateBeyondLoadBalancerModel) SetSubnets(v []BnsLoadBalancerV1ApiCreateHaGroupModelSubnetModel)`

SetSubnets sets Subnets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


