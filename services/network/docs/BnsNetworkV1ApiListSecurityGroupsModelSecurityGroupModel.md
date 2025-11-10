# BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 보안 그룹의 ID | 
**Name** | Pointer to **NullableString** |  | [optional] 
**IsStateful** | Pointer to **NullableBool** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Rules** | Pointer to [**[]BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel**](BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel.md) |  | [optional] 
**ProjectId** | Pointer to **NullableString** |  | [optional] 
**ProjectName** | Pointer to **NullableString** |  | [optional] 
**Tags** | **[]string** | 태그 목록 | 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewBnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel

`func NewBnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel(id string, tags []string, ) *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel`

NewBnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel instantiates a new BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModelWithDefaults

`func NewBnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModelWithDefaults() *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel`

NewBnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModelWithDefaults instantiates a new BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIsStateful

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel) GetIsStateful() bool`

GetIsStateful returns the IsStateful field if non-nil, zero value otherwise.

### GetIsStatefulOk

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel) GetIsStatefulOk() (*bool, bool)`

GetIsStatefulOk returns a tuple with the IsStateful field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStateful

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel) SetIsStateful(v bool)`

SetIsStateful sets IsStateful field to given value.

### HasIsStateful

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel) HasIsStateful() bool`

HasIsStateful returns a boolean if a field has been set.

### SetIsStatefulNil

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel) SetIsStatefulNil(b bool)`

 SetIsStatefulNil sets the value for IsStateful to be an explicit nil

### UnsetIsStateful
`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel) UnsetIsStateful()`

UnsetIsStateful ensures that no value is present for IsStateful, not even an explicit nil
### GetDescription

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetRules

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel) GetRules() []BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel) GetRulesOk() (*[]BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel) SetRules(v []BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupRuleModel)`

SetRules sets Rules field to given value.

### HasRules

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel) HasRules() bool`

HasRules returns a boolean if a field has been set.

### SetRulesNil

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel) SetRulesNil(b bool)`

 SetRulesNil sets the value for Rules to be an explicit nil

### UnsetRules
`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel) UnsetRules()`

UnsetRules ensures that no value is present for Rules, not even an explicit nil
### GetProjectId

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetProjectName

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### SetProjectNameNil

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel) SetProjectNameNil(b bool)`

 SetProjectNameNil sets the value for ProjectName to be an explicit nil

### UnsetProjectName
`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel) UnsetProjectName()`

UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil
### GetTags

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetCreatedAt

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *BnsNetworkV1ApiListSecurityGroupsModelSecurityGroupModel) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


