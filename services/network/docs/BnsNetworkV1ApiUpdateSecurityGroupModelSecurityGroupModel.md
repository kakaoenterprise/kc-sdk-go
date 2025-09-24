# BnsNetworkV1ApiUpdateSecurityGroupModelSecurityGroupModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 보안 그룹의 ID | 
**Name** | **string** | 보안 그룹의 이름 | 
**IsStateful** | **bool** | 보안 그룹의 상태 기반(Stateful) 여부 | 
**Description** | **string** | 보안 그룹에 대한 설명 | 
**Rules** | [**[]BnsNetworkV1ApiUpdateSecurityGroupModelSecurityGroupRuleModel**](BnsNetworkV1ApiUpdateSecurityGroupModelSecurityGroupRuleModel.md) | 보안 그룹에 속하는 규칙 | 
**CreatedAt** | **time.Time** | 리소스가 생성된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
**UpdatedAt** | **time.Time** | 리소스가 마지막으로 수정된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 

## Methods

### NewBnsNetworkV1ApiUpdateSecurityGroupModelSecurityGroupModel

`func NewBnsNetworkV1ApiUpdateSecurityGroupModelSecurityGroupModel(id string, name string, isStateful bool, description string, rules []BnsNetworkV1ApiUpdateSecurityGroupModelSecurityGroupRuleModel, createdAt time.Time, updatedAt time.Time, ) *BnsNetworkV1ApiUpdateSecurityGroupModelSecurityGroupModel`

NewBnsNetworkV1ApiUpdateSecurityGroupModelSecurityGroupModel instantiates a new BnsNetworkV1ApiUpdateSecurityGroupModelSecurityGroupModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsNetworkV1ApiUpdateSecurityGroupModelSecurityGroupModelWithDefaults

`func NewBnsNetworkV1ApiUpdateSecurityGroupModelSecurityGroupModelWithDefaults() *BnsNetworkV1ApiUpdateSecurityGroupModelSecurityGroupModel`

NewBnsNetworkV1ApiUpdateSecurityGroupModelSecurityGroupModelWithDefaults instantiates a new BnsNetworkV1ApiUpdateSecurityGroupModelSecurityGroupModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsNetworkV1ApiUpdateSecurityGroupModelSecurityGroupModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsNetworkV1ApiUpdateSecurityGroupModelSecurityGroupModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsNetworkV1ApiUpdateSecurityGroupModelSecurityGroupModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BnsNetworkV1ApiUpdateSecurityGroupModelSecurityGroupModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BnsNetworkV1ApiUpdateSecurityGroupModelSecurityGroupModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BnsNetworkV1ApiUpdateSecurityGroupModelSecurityGroupModel) SetName(v string)`

SetName sets Name field to given value.


### GetIsStateful

`func (o *BnsNetworkV1ApiUpdateSecurityGroupModelSecurityGroupModel) GetIsStateful() bool`

GetIsStateful returns the IsStateful field if non-nil, zero value otherwise.

### GetIsStatefulOk

`func (o *BnsNetworkV1ApiUpdateSecurityGroupModelSecurityGroupModel) GetIsStatefulOk() (*bool, bool)`

GetIsStatefulOk returns a tuple with the IsStateful field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStateful

`func (o *BnsNetworkV1ApiUpdateSecurityGroupModelSecurityGroupModel) SetIsStateful(v bool)`

SetIsStateful sets IsStateful field to given value.


### GetDescription

`func (o *BnsNetworkV1ApiUpdateSecurityGroupModelSecurityGroupModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BnsNetworkV1ApiUpdateSecurityGroupModelSecurityGroupModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BnsNetworkV1ApiUpdateSecurityGroupModelSecurityGroupModel) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetRules

`func (o *BnsNetworkV1ApiUpdateSecurityGroupModelSecurityGroupModel) GetRules() []BnsNetworkV1ApiUpdateSecurityGroupModelSecurityGroupRuleModel`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *BnsNetworkV1ApiUpdateSecurityGroupModelSecurityGroupModel) GetRulesOk() (*[]BnsNetworkV1ApiUpdateSecurityGroupModelSecurityGroupRuleModel, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *BnsNetworkV1ApiUpdateSecurityGroupModelSecurityGroupModel) SetRules(v []BnsNetworkV1ApiUpdateSecurityGroupModelSecurityGroupRuleModel)`

SetRules sets Rules field to given value.


### GetCreatedAt

`func (o *BnsNetworkV1ApiUpdateSecurityGroupModelSecurityGroupModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BnsNetworkV1ApiUpdateSecurityGroupModelSecurityGroupModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BnsNetworkV1ApiUpdateSecurityGroupModelSecurityGroupModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *BnsNetworkV1ApiUpdateSecurityGroupModelSecurityGroupModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BnsNetworkV1ApiUpdateSecurityGroupModelSecurityGroupModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BnsNetworkV1ApiUpdateSecurityGroupModelSecurityGroupModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


