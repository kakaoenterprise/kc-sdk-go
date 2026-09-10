# CreateMysqlCustomParameterGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | MySQL 커스텀 파라미터 그룹 이름 | 
**SourceParameterGroupId** | **string** | 기준이 되는 MySQL 파라미터 그룹 ID &lt;br/&gt;- &#x60;source_parameter_group_type&#x3D;DEFAULT&#x60;: [List MySQL default parameter groups](https://docs.kakaocloud.com/openapi/data-store/mysql/list-mysql-default-parameter-groups)에서 확인 &lt;br/&gt;- &#x60;source_parameter_group_type&#x3D;CUSTOM&#x60;: [List MySQL custom parameter groups](https://docs.kakaocloud.com/openapi/data-store/mysql/list-mysql-custom-parameter-groups)에서 확인 | 
**SourceParameterGroupType** | [**ParameterGroupType**](ParameterGroupType.md) | 기준이 되는 MySQL 파라미터 그룹 유형 | 
**Description** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreateMysqlCustomParameterGroup

`func NewCreateMysqlCustomParameterGroup(name string, sourceParameterGroupId string, sourceParameterGroupType ParameterGroupType, ) *CreateMysqlCustomParameterGroup`

NewCreateMysqlCustomParameterGroup instantiates a new CreateMysqlCustomParameterGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMysqlCustomParameterGroupWithDefaults

`func NewCreateMysqlCustomParameterGroupWithDefaults() *CreateMysqlCustomParameterGroup`

NewCreateMysqlCustomParameterGroupWithDefaults instantiates a new CreateMysqlCustomParameterGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateMysqlCustomParameterGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateMysqlCustomParameterGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateMysqlCustomParameterGroup) SetName(v string)`

SetName sets Name field to given value.


### GetSourceParameterGroupId

`func (o *CreateMysqlCustomParameterGroup) GetSourceParameterGroupId() string`

GetSourceParameterGroupId returns the SourceParameterGroupId field if non-nil, zero value otherwise.

### GetSourceParameterGroupIdOk

`func (o *CreateMysqlCustomParameterGroup) GetSourceParameterGroupIdOk() (*string, bool)`

GetSourceParameterGroupIdOk returns a tuple with the SourceParameterGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceParameterGroupId

`func (o *CreateMysqlCustomParameterGroup) SetSourceParameterGroupId(v string)`

SetSourceParameterGroupId sets SourceParameterGroupId field to given value.


### GetSourceParameterGroupType

`func (o *CreateMysqlCustomParameterGroup) GetSourceParameterGroupType() ParameterGroupType`

GetSourceParameterGroupType returns the SourceParameterGroupType field if non-nil, zero value otherwise.

### GetSourceParameterGroupTypeOk

`func (o *CreateMysqlCustomParameterGroup) GetSourceParameterGroupTypeOk() (*ParameterGroupType, bool)`

GetSourceParameterGroupTypeOk returns a tuple with the SourceParameterGroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceParameterGroupType

`func (o *CreateMysqlCustomParameterGroup) SetSourceParameterGroupType(v ParameterGroupType)`

SetSourceParameterGroupType sets SourceParameterGroupType field to given value.


### GetDescription

`func (o *CreateMysqlCustomParameterGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateMysqlCustomParameterGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateMysqlCustomParameterGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateMysqlCustomParameterGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateMysqlCustomParameterGroup) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateMysqlCustomParameterGroup) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


