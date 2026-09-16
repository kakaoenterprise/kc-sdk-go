# UpdateCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | 클러스터에 대한 설명 (기존 설명을 새 값으로 덮어씀) | 
**ApiServerIngressRules** | Pointer to [**[]ApiServerIngressRuleRequest**](ApiServerIngressRuleRequest.md) | API 서버에 접근을 허용할 CIDR 목록 | [optional] 

## Methods

### NewUpdateCluster

`func NewUpdateCluster(description string, ) *UpdateCluster`

NewUpdateCluster instantiates a new UpdateCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateClusterWithDefaults

`func NewUpdateClusterWithDefaults() *UpdateCluster`

NewUpdateClusterWithDefaults instantiates a new UpdateCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateCluster) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateCluster) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateCluster) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetApiServerIngressRules

`func (o *UpdateCluster) GetApiServerIngressRules() []ApiServerIngressRuleRequest`

GetApiServerIngressRules returns the ApiServerIngressRules field if non-nil, zero value otherwise.

### GetApiServerIngressRulesOk

`func (o *UpdateCluster) GetApiServerIngressRulesOk() (*[]ApiServerIngressRuleRequest, bool)`

GetApiServerIngressRulesOk returns a tuple with the ApiServerIngressRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiServerIngressRules

`func (o *UpdateCluster) SetApiServerIngressRules(v []ApiServerIngressRuleRequest)`

SetApiServerIngressRules sets ApiServerIngressRules field to given value.

### HasApiServerIngressRules

`func (o *UpdateCluster) HasApiServerIngressRules() bool`

HasApiServerIngressRules returns a boolean if a field has been set.

### SetApiServerIngressRulesNil

`func (o *UpdateCluster) SetApiServerIngressRulesNil(b bool)`

 SetApiServerIngressRulesNil sets the value for ApiServerIngressRules to be an explicit nil

### UnsetApiServerIngressRules
`func (o *UpdateCluster) UnsetApiServerIngressRules()`

UnsetApiServerIngressRules ensures that no value is present for ApiServerIngressRules, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


