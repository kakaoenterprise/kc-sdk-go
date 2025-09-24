# BnsNetworkV1ApiCreateSecurityGroupModelSecurityGroupModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 보안 그룹의 ID | 
**Name** | **string** | 보안 그룹 이름 | 
**IsStateful** | **bool** | 보안 그룹의 상태 기반(Stateful) 여부 | 
**Description** | **string** | 보안 그룹에 대한 설명 | 
**Rules** | [**[]BnsNetworkV1ApiCreateSecurityGroupModelSecurityGroupRuleModel**](BnsNetworkV1ApiCreateSecurityGroupModelSecurityGroupRuleModel.md) | 보안 그룹에 포함된 규칙 목록 | 
**CreatedAt** | **time.Time** | 리소스가 생성된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
**UpdatedAt** | **time.Time** | 리소스가 마지막으로 수정된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 

## Methods

### NewBnsNetworkV1ApiCreateSecurityGroupModelSecurityGroupModel

`func NewBnsNetworkV1ApiCreateSecurityGroupModelSecurityGroupModel(id string, name string, isStateful bool, description string, rules []BnsNetworkV1ApiCreateSecurityGroupModelSecurityGroupRuleModel, createdAt time.Time, updatedAt time.Time, ) *BnsNetworkV1ApiCreateSecurityGroupModelSecurityGroupModel`

NewBnsNetworkV1ApiCreateSecurityGroupModelSecurityGroupModel instantiates a new BnsNetworkV1ApiCreateSecurityGroupModelSecurityGroupModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsNetworkV1ApiCreateSecurityGroupModelSecurityGroupModelWithDefaults

`func NewBnsNetworkV1ApiCreateSecurityGroupModelSecurityGroupModelWithDefaults() *BnsNetworkV1ApiCreateSecurityGroupModelSecurityGroupModel`

NewBnsNetworkV1ApiCreateSecurityGroupModelSecurityGroupModelWithDefaults instantiates a new BnsNetworkV1ApiCreateSecurityGroupModelSecurityGroupModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsNetworkV1ApiCreateSecurityGroupModelSecurityGroupModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsNetworkV1ApiCreateSecurityGroupModelSecurityGroupModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsNetworkV1ApiCreateSecurityGroupModelSecurityGroupModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BnsNetworkV1ApiCreateSecurityGroupModelSecurityGroupModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BnsNetworkV1ApiCreateSecurityGroupModelSecurityGroupModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BnsNetworkV1ApiCreateSecurityGroupModelSecurityGroupModel) SetName(v string)`

SetName sets Name field to given value.


### GetIsStateful

`func (o *BnsNetworkV1ApiCreateSecurityGroupModelSecurityGroupModel) GetIsStateful() bool`

GetIsStateful returns the IsStateful field if non-nil, zero value otherwise.

### GetIsStatefulOk

`func (o *BnsNetworkV1ApiCreateSecurityGroupModelSecurityGroupModel) GetIsStatefulOk() (*bool, bool)`

GetIsStatefulOk returns a tuple with the IsStateful field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStateful

`func (o *BnsNetworkV1ApiCreateSecurityGroupModelSecurityGroupModel) SetIsStateful(v bool)`

SetIsStateful sets IsStateful field to given value.


### GetDescription

`func (o *BnsNetworkV1ApiCreateSecurityGroupModelSecurityGroupModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BnsNetworkV1ApiCreateSecurityGroupModelSecurityGroupModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BnsNetworkV1ApiCreateSecurityGroupModelSecurityGroupModel) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetRules

`func (o *BnsNetworkV1ApiCreateSecurityGroupModelSecurityGroupModel) GetRules() []BnsNetworkV1ApiCreateSecurityGroupModelSecurityGroupRuleModel`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *BnsNetworkV1ApiCreateSecurityGroupModelSecurityGroupModel) GetRulesOk() (*[]BnsNetworkV1ApiCreateSecurityGroupModelSecurityGroupRuleModel, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *BnsNetworkV1ApiCreateSecurityGroupModelSecurityGroupModel) SetRules(v []BnsNetworkV1ApiCreateSecurityGroupModelSecurityGroupRuleModel)`

SetRules sets Rules field to given value.


### GetCreatedAt

`func (o *BnsNetworkV1ApiCreateSecurityGroupModelSecurityGroupModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BnsNetworkV1ApiCreateSecurityGroupModelSecurityGroupModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BnsNetworkV1ApiCreateSecurityGroupModelSecurityGroupModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *BnsNetworkV1ApiCreateSecurityGroupModelSecurityGroupModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BnsNetworkV1ApiCreateSecurityGroupModelSecurityGroupModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BnsNetworkV1ApiCreateSecurityGroupModelSecurityGroupModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


