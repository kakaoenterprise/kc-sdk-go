# UpdateL7PolicyRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**L7Rule** | [**UpdateL7PolicyRule**](UpdateL7PolicyRule.md) | L7 정책에 연결된 규칙의 규칙 및 상태를 수정하는 요청 본문 | 

## Methods

### NewUpdateL7PolicyRuleRequest

`func NewUpdateL7PolicyRuleRequest(l7Rule UpdateL7PolicyRule, ) *UpdateL7PolicyRuleRequest`

NewUpdateL7PolicyRuleRequest instantiates a new UpdateL7PolicyRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateL7PolicyRuleRequestWithDefaults

`func NewUpdateL7PolicyRuleRequestWithDefaults() *UpdateL7PolicyRuleRequest`

NewUpdateL7PolicyRuleRequestWithDefaults instantiates a new UpdateL7PolicyRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetL7Rule

`func (o *UpdateL7PolicyRuleRequest) GetL7Rule() UpdateL7PolicyRule`

GetL7Rule returns the L7Rule field if non-nil, zero value otherwise.

### GetL7RuleOk

`func (o *UpdateL7PolicyRuleRequest) GetL7RuleOk() (*UpdateL7PolicyRule, bool)`

GetL7RuleOk returns a tuple with the L7Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL7Rule

`func (o *UpdateL7PolicyRuleRequest) SetL7Rule(v UpdateL7PolicyRule)`

SetL7Rule sets L7Rule field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


