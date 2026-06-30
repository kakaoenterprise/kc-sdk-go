# RateLimitPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyId** | **string** |  | 
**PolicyName** | **string** |  | 
**Status** | **string** |  | 
**Priority** | **int32** |  | 
**KeyType** | **string** |  | 
**KeyTargets** | **[]string** |  | 
**SubType** | Pointer to **string** |  | [optional] 
**Rules** | [**[]RateLimitRule**](RateLimitRule.md) |  | 
**Operator** | **string** |  | 
**Description** | **string** |  | 
**CreatedBy** | **string** |  | 
**UpdatedBy** | **string** |  | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Version** | **[]int32** |  | 

## Methods

### NewRateLimitPolicy

`func NewRateLimitPolicy(policyId string, policyName string, status string, priority int32, keyType string, keyTargets []string, rules []RateLimitRule, operator string, description string, createdBy string, updatedBy string, version []int32, ) *RateLimitPolicy`

NewRateLimitPolicy instantiates a new RateLimitPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateLimitPolicyWithDefaults

`func NewRateLimitPolicyWithDefaults() *RateLimitPolicy`

NewRateLimitPolicyWithDefaults instantiates a new RateLimitPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyId

`func (o *RateLimitPolicy) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *RateLimitPolicy) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *RateLimitPolicy) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.


### GetPolicyName

`func (o *RateLimitPolicy) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *RateLimitPolicy) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *RateLimitPolicy) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.


### GetStatus

`func (o *RateLimitPolicy) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RateLimitPolicy) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RateLimitPolicy) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetPriority

`func (o *RateLimitPolicy) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *RateLimitPolicy) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *RateLimitPolicy) SetPriority(v int32)`

SetPriority sets Priority field to given value.


### GetKeyType

`func (o *RateLimitPolicy) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *RateLimitPolicy) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *RateLimitPolicy) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.


### GetKeyTargets

`func (o *RateLimitPolicy) GetKeyTargets() []string`

GetKeyTargets returns the KeyTargets field if non-nil, zero value otherwise.

### GetKeyTargetsOk

`func (o *RateLimitPolicy) GetKeyTargetsOk() (*[]string, bool)`

GetKeyTargetsOk returns a tuple with the KeyTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyTargets

`func (o *RateLimitPolicy) SetKeyTargets(v []string)`

SetKeyTargets sets KeyTargets field to given value.


### GetSubType

`func (o *RateLimitPolicy) GetSubType() string`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *RateLimitPolicy) GetSubTypeOk() (*string, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *RateLimitPolicy) SetSubType(v string)`

SetSubType sets SubType field to given value.

### HasSubType

`func (o *RateLimitPolicy) HasSubType() bool`

HasSubType returns a boolean if a field has been set.

### GetRules

`func (o *RateLimitPolicy) GetRules() []RateLimitRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *RateLimitPolicy) GetRulesOk() (*[]RateLimitRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *RateLimitPolicy) SetRules(v []RateLimitRule)`

SetRules sets Rules field to given value.


### GetOperator

`func (o *RateLimitPolicy) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *RateLimitPolicy) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *RateLimitPolicy) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetDescription

`func (o *RateLimitPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RateLimitPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RateLimitPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCreatedBy

`func (o *RateLimitPolicy) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *RateLimitPolicy) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *RateLimitPolicy) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetUpdatedBy

`func (o *RateLimitPolicy) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *RateLimitPolicy) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *RateLimitPolicy) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.


### GetCreatedAt

`func (o *RateLimitPolicy) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RateLimitPolicy) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RateLimitPolicy) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RateLimitPolicy) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RateLimitPolicy) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RateLimitPolicy) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RateLimitPolicy) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RateLimitPolicy) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVersion

`func (o *RateLimitPolicy) GetVersion() []int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RateLimitPolicy) GetVersionOk() (*[]int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RateLimitPolicy) SetVersion(v []int32)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


