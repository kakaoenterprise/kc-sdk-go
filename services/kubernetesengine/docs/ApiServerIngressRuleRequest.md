# ApiServerIngressRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cidr** | **string** | API 서버 접근을 허용할 네트워크의 CIDR | 
**Description** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewApiServerIngressRuleRequest

`func NewApiServerIngressRuleRequest(cidr string, ) *ApiServerIngressRuleRequest`

NewApiServerIngressRuleRequest instantiates a new ApiServerIngressRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiServerIngressRuleRequestWithDefaults

`func NewApiServerIngressRuleRequestWithDefaults() *ApiServerIngressRuleRequest`

NewApiServerIngressRuleRequestWithDefaults instantiates a new ApiServerIngressRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCidr

`func (o *ApiServerIngressRuleRequest) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *ApiServerIngressRuleRequest) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *ApiServerIngressRuleRequest) SetCidr(v string)`

SetCidr sets Cidr field to given value.


### GetDescription

`func (o *ApiServerIngressRuleRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiServerIngressRuleRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiServerIngressRuleRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiServerIngressRuleRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ApiServerIngressRuleRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ApiServerIngressRuleRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


