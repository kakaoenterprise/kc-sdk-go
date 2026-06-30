# CreateRateLimitPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyName** | **string** |  | 
**Priority** | **int32** |  | 
**KeyType** | **string** |  | 
**KeyTargets** | **[]string** |  | 
**SubType** | Pointer to **string** |  | [optional] 
**Rules** | [**[]RateLimitRule**](RateLimitRule.md) |  | 
**Operator** | **string** |  | 
**Description** | **string** |  | 

## Methods

### NewCreateRateLimitPolicyRequest

`func NewCreateRateLimitPolicyRequest(policyName string, priority int32, keyType string, keyTargets []string, rules []RateLimitRule, operator string, description string, ) *CreateRateLimitPolicyRequest`

NewCreateRateLimitPolicyRequest instantiates a new CreateRateLimitPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRateLimitPolicyRequestWithDefaults

`func NewCreateRateLimitPolicyRequestWithDefaults() *CreateRateLimitPolicyRequest`

NewCreateRateLimitPolicyRequestWithDefaults instantiates a new CreateRateLimitPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyName

`func (o *CreateRateLimitPolicyRequest) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *CreateRateLimitPolicyRequest) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *CreateRateLimitPolicyRequest) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.


### GetPriority

`func (o *CreateRateLimitPolicyRequest) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *CreateRateLimitPolicyRequest) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *CreateRateLimitPolicyRequest) SetPriority(v int32)`

SetPriority sets Priority field to given value.


### GetKeyType

`func (o *CreateRateLimitPolicyRequest) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *CreateRateLimitPolicyRequest) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *CreateRateLimitPolicyRequest) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.


### GetKeyTargets

`func (o *CreateRateLimitPolicyRequest) GetKeyTargets() []string`

GetKeyTargets returns the KeyTargets field if non-nil, zero value otherwise.

### GetKeyTargetsOk

`func (o *CreateRateLimitPolicyRequest) GetKeyTargetsOk() (*[]string, bool)`

GetKeyTargetsOk returns a tuple with the KeyTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyTargets

`func (o *CreateRateLimitPolicyRequest) SetKeyTargets(v []string)`

SetKeyTargets sets KeyTargets field to given value.


### GetSubType

`func (o *CreateRateLimitPolicyRequest) GetSubType() string`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *CreateRateLimitPolicyRequest) GetSubTypeOk() (*string, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *CreateRateLimitPolicyRequest) SetSubType(v string)`

SetSubType sets SubType field to given value.

### HasSubType

`func (o *CreateRateLimitPolicyRequest) HasSubType() bool`

HasSubType returns a boolean if a field has been set.

### GetRules

`func (o *CreateRateLimitPolicyRequest) GetRules() []RateLimitRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *CreateRateLimitPolicyRequest) GetRulesOk() (*[]RateLimitRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *CreateRateLimitPolicyRequest) SetRules(v []RateLimitRule)`

SetRules sets Rules field to given value.


### GetOperator

`func (o *CreateRateLimitPolicyRequest) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *CreateRateLimitPolicyRequest) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *CreateRateLimitPolicyRequest) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetDescription

`func (o *CreateRateLimitPolicyRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateRateLimitPolicyRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateRateLimitPolicyRequest) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


