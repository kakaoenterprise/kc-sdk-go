# BcsInstanceV1ApiCreateInstanceModelInstanceModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 생성된 인스턴스의 ID | 
**SecurityGroups** | [**[]CreateInstanceSecurityGroupModel**](CreateInstanceSecurityGroupModel.md) | 인스턴스에 적용된 보안 그룹 목록 | 
**AdminPass** | Pointer to **NullableString** |  | [optional] 
**DiskConfig** | **string** | 디스크 설정 방식 | 

## Methods

### NewBcsInstanceV1ApiCreateInstanceModelInstanceModel

`func NewBcsInstanceV1ApiCreateInstanceModelInstanceModel(id string, securityGroups []CreateInstanceSecurityGroupModel, diskConfig string, ) *BcsInstanceV1ApiCreateInstanceModelInstanceModel`

NewBcsInstanceV1ApiCreateInstanceModelInstanceModel instantiates a new BcsInstanceV1ApiCreateInstanceModelInstanceModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBcsInstanceV1ApiCreateInstanceModelInstanceModelWithDefaults

`func NewBcsInstanceV1ApiCreateInstanceModelInstanceModelWithDefaults() *BcsInstanceV1ApiCreateInstanceModelInstanceModel`

NewBcsInstanceV1ApiCreateInstanceModelInstanceModelWithDefaults instantiates a new BcsInstanceV1ApiCreateInstanceModelInstanceModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BcsInstanceV1ApiCreateInstanceModelInstanceModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BcsInstanceV1ApiCreateInstanceModelInstanceModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BcsInstanceV1ApiCreateInstanceModelInstanceModel) SetId(v string)`

SetId sets Id field to given value.


### GetSecurityGroups

`func (o *BcsInstanceV1ApiCreateInstanceModelInstanceModel) GetSecurityGroups() []CreateInstanceSecurityGroupModel`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *BcsInstanceV1ApiCreateInstanceModelInstanceModel) GetSecurityGroupsOk() (*[]CreateInstanceSecurityGroupModel, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *BcsInstanceV1ApiCreateInstanceModelInstanceModel) SetSecurityGroups(v []CreateInstanceSecurityGroupModel)`

SetSecurityGroups sets SecurityGroups field to given value.


### GetAdminPass

`func (o *BcsInstanceV1ApiCreateInstanceModelInstanceModel) GetAdminPass() string`

GetAdminPass returns the AdminPass field if non-nil, zero value otherwise.

### GetAdminPassOk

`func (o *BcsInstanceV1ApiCreateInstanceModelInstanceModel) GetAdminPassOk() (*string, bool)`

GetAdminPassOk returns a tuple with the AdminPass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPass

`func (o *BcsInstanceV1ApiCreateInstanceModelInstanceModel) SetAdminPass(v string)`

SetAdminPass sets AdminPass field to given value.

### HasAdminPass

`func (o *BcsInstanceV1ApiCreateInstanceModelInstanceModel) HasAdminPass() bool`

HasAdminPass returns a boolean if a field has been set.

### SetAdminPassNil

`func (o *BcsInstanceV1ApiCreateInstanceModelInstanceModel) SetAdminPassNil(b bool)`

 SetAdminPassNil sets the value for AdminPass to be an explicit nil

### UnsetAdminPass
`func (o *BcsInstanceV1ApiCreateInstanceModelInstanceModel) UnsetAdminPass()`

UnsetAdminPass ensures that no value is present for AdminPass, not even an explicit nil
### GetDiskConfig

`func (o *BcsInstanceV1ApiCreateInstanceModelInstanceModel) GetDiskConfig() string`

GetDiskConfig returns the DiskConfig field if non-nil, zero value otherwise.

### GetDiskConfigOk

`func (o *BcsInstanceV1ApiCreateInstanceModelInstanceModel) GetDiskConfigOk() (*string, bool)`

GetDiskConfigOk returns a tuple with the DiskConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskConfig

`func (o *BcsInstanceV1ApiCreateInstanceModelInstanceModel) SetDiskConfig(v string)`

SetDiskConfig sets DiskConfig field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


