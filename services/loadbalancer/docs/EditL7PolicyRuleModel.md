# EditL7PolicyRuleModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompareType** | [**L7RuleCompareType**](L7RuleCompareType.md) | 규칙 값 비교 방식 | 
**IsInverted** | Pointer to **NullableBool** |  | [optional] 
**Key** | Pointer to **NullableString** |  | [optional] 
**Type** | [**L7RuleType**](L7RuleType.md) | 요청의 어느 속성을 기준으로 규칙을 평가할지 정의 &lt;br/&gt; - &#x60;COOKIE&#x60;: 쿠키 값 &lt;br/&gt; - &#x60;FILE_TYPE&#x60;: 파일 확장자 &lt;br/&gt; - &#x60;HEADER&#x60;: HTTP 헤더 &lt;br/&gt; - &#x60;HOST_NAME&#x60;: 호스트 이름 &lt;br/&gt; - &#x60;PATH&#x60;: 요청 경로 | 
**Value** | **string** | 비교할 대상 값 | 

## Methods

### NewEditL7PolicyRuleModel

`func NewEditL7PolicyRuleModel(compareType L7RuleCompareType, type_ L7RuleType, value string, ) *EditL7PolicyRuleModel`

NewEditL7PolicyRuleModel instantiates a new EditL7PolicyRuleModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditL7PolicyRuleModelWithDefaults

`func NewEditL7PolicyRuleModelWithDefaults() *EditL7PolicyRuleModel`

NewEditL7PolicyRuleModelWithDefaults instantiates a new EditL7PolicyRuleModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompareType

`func (o *EditL7PolicyRuleModel) GetCompareType() L7RuleCompareType`

GetCompareType returns the CompareType field if non-nil, zero value otherwise.

### GetCompareTypeOk

`func (o *EditL7PolicyRuleModel) GetCompareTypeOk() (*L7RuleCompareType, bool)`

GetCompareTypeOk returns a tuple with the CompareType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompareType

`func (o *EditL7PolicyRuleModel) SetCompareType(v L7RuleCompareType)`

SetCompareType sets CompareType field to given value.


### GetIsInverted

`func (o *EditL7PolicyRuleModel) GetIsInverted() bool`

GetIsInverted returns the IsInverted field if non-nil, zero value otherwise.

### GetIsInvertedOk

`func (o *EditL7PolicyRuleModel) GetIsInvertedOk() (*bool, bool)`

GetIsInvertedOk returns a tuple with the IsInverted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInverted

`func (o *EditL7PolicyRuleModel) SetIsInverted(v bool)`

SetIsInverted sets IsInverted field to given value.

### HasIsInverted

`func (o *EditL7PolicyRuleModel) HasIsInverted() bool`

HasIsInverted returns a boolean if a field has been set.

### SetIsInvertedNil

`func (o *EditL7PolicyRuleModel) SetIsInvertedNil(b bool)`

 SetIsInvertedNil sets the value for IsInverted to be an explicit nil

### UnsetIsInverted
`func (o *EditL7PolicyRuleModel) UnsetIsInverted()`

UnsetIsInverted ensures that no value is present for IsInverted, not even an explicit nil
### GetKey

`func (o *EditL7PolicyRuleModel) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *EditL7PolicyRuleModel) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *EditL7PolicyRuleModel) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *EditL7PolicyRuleModel) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *EditL7PolicyRuleModel) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *EditL7PolicyRuleModel) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetType

`func (o *EditL7PolicyRuleModel) GetType() L7RuleType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EditL7PolicyRuleModel) GetTypeOk() (*L7RuleType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EditL7PolicyRuleModel) SetType(v L7RuleType)`

SetType sets Type field to given value.


### GetValue

`func (o *EditL7PolicyRuleModel) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EditL7PolicyRuleModel) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EditL7PolicyRuleModel) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


