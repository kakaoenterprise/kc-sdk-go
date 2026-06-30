# MysqlV1ApiCreateMysqlCustomParameterGroupModelCustomParameterGroupRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | MySQL 커스텀 파라미터 그룹 이름 | 
**SourceParameterGroupId** | **string** | # source_parameter_group_id - 타입: String - 설명: (필드 설명을 여기에 작성하세요) | 
**SourceParameterGroupType** | [**ParameterGroupType**](ParameterGroupType.md) | # source_parameter_group_type - 타입: ParameterGroupType - 설명: (필드 설명을 여기에 작성하세요) | 
**Description** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewMysqlV1ApiCreateMysqlCustomParameterGroupModelCustomParameterGroupRequestModel

`func NewMysqlV1ApiCreateMysqlCustomParameterGroupModelCustomParameterGroupRequestModel(name string, sourceParameterGroupId string, sourceParameterGroupType ParameterGroupType, ) *MysqlV1ApiCreateMysqlCustomParameterGroupModelCustomParameterGroupRequestModel`

NewMysqlV1ApiCreateMysqlCustomParameterGroupModelCustomParameterGroupRequestModel instantiates a new MysqlV1ApiCreateMysqlCustomParameterGroupModelCustomParameterGroupRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMysqlV1ApiCreateMysqlCustomParameterGroupModelCustomParameterGroupRequestModelWithDefaults

`func NewMysqlV1ApiCreateMysqlCustomParameterGroupModelCustomParameterGroupRequestModelWithDefaults() *MysqlV1ApiCreateMysqlCustomParameterGroupModelCustomParameterGroupRequestModel`

NewMysqlV1ApiCreateMysqlCustomParameterGroupModelCustomParameterGroupRequestModelWithDefaults instantiates a new MysqlV1ApiCreateMysqlCustomParameterGroupModelCustomParameterGroupRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MysqlV1ApiCreateMysqlCustomParameterGroupModelCustomParameterGroupRequestModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MysqlV1ApiCreateMysqlCustomParameterGroupModelCustomParameterGroupRequestModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MysqlV1ApiCreateMysqlCustomParameterGroupModelCustomParameterGroupRequestModel) SetName(v string)`

SetName sets Name field to given value.


### GetSourceParameterGroupId

`func (o *MysqlV1ApiCreateMysqlCustomParameterGroupModelCustomParameterGroupRequestModel) GetSourceParameterGroupId() string`

GetSourceParameterGroupId returns the SourceParameterGroupId field if non-nil, zero value otherwise.

### GetSourceParameterGroupIdOk

`func (o *MysqlV1ApiCreateMysqlCustomParameterGroupModelCustomParameterGroupRequestModel) GetSourceParameterGroupIdOk() (*string, bool)`

GetSourceParameterGroupIdOk returns a tuple with the SourceParameterGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceParameterGroupId

`func (o *MysqlV1ApiCreateMysqlCustomParameterGroupModelCustomParameterGroupRequestModel) SetSourceParameterGroupId(v string)`

SetSourceParameterGroupId sets SourceParameterGroupId field to given value.


### GetSourceParameterGroupType

`func (o *MysqlV1ApiCreateMysqlCustomParameterGroupModelCustomParameterGroupRequestModel) GetSourceParameterGroupType() ParameterGroupType`

GetSourceParameterGroupType returns the SourceParameterGroupType field if non-nil, zero value otherwise.

### GetSourceParameterGroupTypeOk

`func (o *MysqlV1ApiCreateMysqlCustomParameterGroupModelCustomParameterGroupRequestModel) GetSourceParameterGroupTypeOk() (*ParameterGroupType, bool)`

GetSourceParameterGroupTypeOk returns a tuple with the SourceParameterGroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceParameterGroupType

`func (o *MysqlV1ApiCreateMysqlCustomParameterGroupModelCustomParameterGroupRequestModel) SetSourceParameterGroupType(v ParameterGroupType)`

SetSourceParameterGroupType sets SourceParameterGroupType field to given value.


### GetDescription

`func (o *MysqlV1ApiCreateMysqlCustomParameterGroupModelCustomParameterGroupRequestModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MysqlV1ApiCreateMysqlCustomParameterGroupModelCustomParameterGroupRequestModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MysqlV1ApiCreateMysqlCustomParameterGroupModelCustomParameterGroupRequestModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MysqlV1ApiCreateMysqlCustomParameterGroupModelCustomParameterGroupRequestModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MysqlV1ApiCreateMysqlCustomParameterGroupModelCustomParameterGroupRequestModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MysqlV1ApiCreateMysqlCustomParameterGroupModelCustomParameterGroupRequestModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


