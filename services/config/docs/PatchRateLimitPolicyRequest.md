# PatchRateLimitPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyName** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**KeyType** | Pointer to **string** |  | [optional] 
**KeyTargets** | Pointer to **[]string** |  | [optional] 
**SubType** | Pointer to **string** |  | [optional] 
**Rules** | Pointer to [**[]RateLimitRule**](RateLimitRule.md) |  | [optional] 
**Operator** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewPatchRateLimitPolicyRequest

`func NewPatchRateLimitPolicyRequest() *PatchRateLimitPolicyRequest`

NewPatchRateLimitPolicyRequest instantiates a new PatchRateLimitPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchRateLimitPolicyRequestWithDefaults

`func NewPatchRateLimitPolicyRequestWithDefaults() *PatchRateLimitPolicyRequest`

NewPatchRateLimitPolicyRequestWithDefaults instantiates a new PatchRateLimitPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyName

`func (o *PatchRateLimitPolicyRequest) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *PatchRateLimitPolicyRequest) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *PatchRateLimitPolicyRequest) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *PatchRateLimitPolicyRequest) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### GetPriority

`func (o *PatchRateLimitPolicyRequest) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *PatchRateLimitPolicyRequest) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *PatchRateLimitPolicyRequest) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *PatchRateLimitPolicyRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetKeyType

`func (o *PatchRateLimitPolicyRequest) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *PatchRateLimitPolicyRequest) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *PatchRateLimitPolicyRequest) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *PatchRateLimitPolicyRequest) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### GetKeyTargets

`func (o *PatchRateLimitPolicyRequest) GetKeyTargets() []string`

GetKeyTargets returns the KeyTargets field if non-nil, zero value otherwise.

### GetKeyTargetsOk

`func (o *PatchRateLimitPolicyRequest) GetKeyTargetsOk() (*[]string, bool)`

GetKeyTargetsOk returns a tuple with the KeyTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyTargets

`func (o *PatchRateLimitPolicyRequest) SetKeyTargets(v []string)`

SetKeyTargets sets KeyTargets field to given value.

### HasKeyTargets

`func (o *PatchRateLimitPolicyRequest) HasKeyTargets() bool`

HasKeyTargets returns a boolean if a field has been set.

### GetSubType

`func (o *PatchRateLimitPolicyRequest) GetSubType() string`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *PatchRateLimitPolicyRequest) GetSubTypeOk() (*string, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *PatchRateLimitPolicyRequest) SetSubType(v string)`

SetSubType sets SubType field to given value.

### HasSubType

`func (o *PatchRateLimitPolicyRequest) HasSubType() bool`

HasSubType returns a boolean if a field has been set.

### GetRules

`func (o *PatchRateLimitPolicyRequest) GetRules() []RateLimitRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *PatchRateLimitPolicyRequest) GetRulesOk() (*[]RateLimitRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *PatchRateLimitPolicyRequest) SetRules(v []RateLimitRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *PatchRateLimitPolicyRequest) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetOperator

`func (o *PatchRateLimitPolicyRequest) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *PatchRateLimitPolicyRequest) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *PatchRateLimitPolicyRequest) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *PatchRateLimitPolicyRequest) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetDescription

`func (o *PatchRateLimitPolicyRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchRateLimitPolicyRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchRateLimitPolicyRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchRateLimitPolicyRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


