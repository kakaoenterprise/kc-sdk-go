# ApiServerIngressRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cidr** | **string** | API 서버 접근을 허용하는 네트워크의 CIDR | 
**Description** | Pointer to **NullableString** | API 서버 접근 규칙에 대한 설명 | [optional] 

## Methods

### NewApiServerIngressRule

`func NewApiServerIngressRule(cidr string, ) *ApiServerIngressRule`

NewApiServerIngressRule instantiates a new ApiServerIngressRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiServerIngressRuleWithDefaults

`func NewApiServerIngressRuleWithDefaults() *ApiServerIngressRule`

NewApiServerIngressRuleWithDefaults instantiates a new ApiServerIngressRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCidr

`func (o *ApiServerIngressRule) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *ApiServerIngressRule) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *ApiServerIngressRule) SetCidr(v string)`

SetCidr sets Cidr field to given value.


### GetDescription

`func (o *ApiServerIngressRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiServerIngressRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiServerIngressRule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiServerIngressRule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ApiServerIngressRule) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ApiServerIngressRule) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


