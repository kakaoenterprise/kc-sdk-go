# CreateSecurityGroupRuleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecurityGroupRule** | [**SecurityGroupRule**](SecurityGroupRule.md) | 생성된 보안 그룹 규칙 정보 | 

## Methods

### NewCreateSecurityGroupRuleResponse

`func NewCreateSecurityGroupRuleResponse(securityGroupRule SecurityGroupRule, ) *CreateSecurityGroupRuleResponse`

NewCreateSecurityGroupRuleResponse instantiates a new CreateSecurityGroupRuleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSecurityGroupRuleResponseWithDefaults

`func NewCreateSecurityGroupRuleResponseWithDefaults() *CreateSecurityGroupRuleResponse`

NewCreateSecurityGroupRuleResponseWithDefaults instantiates a new CreateSecurityGroupRuleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecurityGroupRule

`func (o *CreateSecurityGroupRuleResponse) GetSecurityGroupRule() SecurityGroupRule`

GetSecurityGroupRule returns the SecurityGroupRule field if non-nil, zero value otherwise.

### GetSecurityGroupRuleOk

`func (o *CreateSecurityGroupRuleResponse) GetSecurityGroupRuleOk() (*SecurityGroupRule, bool)`

GetSecurityGroupRuleOk returns a tuple with the SecurityGroupRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupRule

`func (o *CreateSecurityGroupRuleResponse) SetSecurityGroupRule(v SecurityGroupRule)`

SetSecurityGroupRule sets SecurityGroupRule field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


