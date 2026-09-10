# SecurityGroupResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 보안 그룹의 ID | 
**Name** | **string** | 보안 그룹의 이름 | 
**IsStateful** | **bool** | 보안 그룹의 상태 기반(Stateful) 여부 | 
**Description** | **string** | 보안 그룹에 대한 설명 | 
**CreatedAt** | **time.Time** | 리소스가 생성된 시간 &lt;br/&gt; - ISO 8601 형식 &lt;br/&gt; - UTC 기준 | 
**UpdatedAt** | **time.Time** | 리소스가 마지막으로 수정된 시간 &lt;br/&gt; - ISO 8601 형식 &lt;br/&gt; - UTC 기준 | 
**Tags** | **[]string** | 태그 목록 | 
**Rules** | [**[]RuleResult**](RuleResult.md) | 보안 그룹에 속하는 규칙 | 

## Methods

### NewSecurityGroupResult

`func NewSecurityGroupResult(id string, name string, isStateful bool, description string, createdAt time.Time, updatedAt time.Time, tags []string, rules []RuleResult, ) *SecurityGroupResult`

NewSecurityGroupResult instantiates a new SecurityGroupResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityGroupResultWithDefaults

`func NewSecurityGroupResultWithDefaults() *SecurityGroupResult`

NewSecurityGroupResultWithDefaults instantiates a new SecurityGroupResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SecurityGroupResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecurityGroupResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecurityGroupResult) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *SecurityGroupResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityGroupResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityGroupResult) SetName(v string)`

SetName sets Name field to given value.


### GetIsStateful

`func (o *SecurityGroupResult) GetIsStateful() bool`

GetIsStateful returns the IsStateful field if non-nil, zero value otherwise.

### GetIsStatefulOk

`func (o *SecurityGroupResult) GetIsStatefulOk() (*bool, bool)`

GetIsStatefulOk returns a tuple with the IsStateful field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStateful

`func (o *SecurityGroupResult) SetIsStateful(v bool)`

SetIsStateful sets IsStateful field to given value.


### GetDescription

`func (o *SecurityGroupResult) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SecurityGroupResult) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SecurityGroupResult) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCreatedAt

`func (o *SecurityGroupResult) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SecurityGroupResult) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SecurityGroupResult) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *SecurityGroupResult) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SecurityGroupResult) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SecurityGroupResult) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetTags

`func (o *SecurityGroupResult) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SecurityGroupResult) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SecurityGroupResult) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetRules

`func (o *SecurityGroupResult) GetRules() []RuleResult`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *SecurityGroupResult) GetRulesOk() (*[]RuleResult, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *SecurityGroupResult) SetRules(v []RuleResult)`

SetRules sets Rules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


