# CreateSecurityGroupRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecurityGroupRule** | [**CreateSecurityGroupRule**](CreateSecurityGroupRule.md) | 생성할 보안 그룹 규칙 정보 | 

## Methods

### NewCreateSecurityGroupRuleRequest

`func NewCreateSecurityGroupRuleRequest(securityGroupRule CreateSecurityGroupRule, ) *CreateSecurityGroupRuleRequest`

NewCreateSecurityGroupRuleRequest instantiates a new CreateSecurityGroupRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSecurityGroupRuleRequestWithDefaults

`func NewCreateSecurityGroupRuleRequestWithDefaults() *CreateSecurityGroupRuleRequest`

NewCreateSecurityGroupRuleRequestWithDefaults instantiates a new CreateSecurityGroupRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecurityGroupRule

`func (o *CreateSecurityGroupRuleRequest) GetSecurityGroupRule() CreateSecurityGroupRule`

GetSecurityGroupRule returns the SecurityGroupRule field if non-nil, zero value otherwise.

### GetSecurityGroupRuleOk

`func (o *CreateSecurityGroupRuleRequest) GetSecurityGroupRuleOk() (*CreateSecurityGroupRule, bool)`

GetSecurityGroupRuleOk returns a tuple with the SecurityGroupRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupRule

`func (o *CreateSecurityGroupRuleRequest) SetSecurityGroupRule(v CreateSecurityGroupRule)`

SetSecurityGroupRule sets SecurityGroupRule field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


