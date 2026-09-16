# CreatedInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 생성된 인스턴스의 ID | 
**DiskConfig** | **string** | 디스크 설정 방식 | 
**SecurityGroups** | [**[]SecurityGroup**](SecurityGroup.md) | 인스턴스에 적용된 보안 그룹 목록 | 
**AdminPass** | Pointer to **NullableString** | 생성된 인스턴스의 초기 관리자 비밀번호 - 인스턴스 접속 시 기본 계정의 비밀번호로 사용 | [optional] 

## Methods

### NewCreatedInstance

`func NewCreatedInstance(id string, diskConfig string, securityGroups []SecurityGroup, ) *CreatedInstance`

NewCreatedInstance instantiates a new CreatedInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatedInstanceWithDefaults

`func NewCreatedInstanceWithDefaults() *CreatedInstance`

NewCreatedInstanceWithDefaults instantiates a new CreatedInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreatedInstance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreatedInstance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreatedInstance) SetId(v string)`

SetId sets Id field to given value.


### GetDiskConfig

`func (o *CreatedInstance) GetDiskConfig() string`

GetDiskConfig returns the DiskConfig field if non-nil, zero value otherwise.

### GetDiskConfigOk

`func (o *CreatedInstance) GetDiskConfigOk() (*string, bool)`

GetDiskConfigOk returns a tuple with the DiskConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskConfig

`func (o *CreatedInstance) SetDiskConfig(v string)`

SetDiskConfig sets DiskConfig field to given value.


### GetSecurityGroups

`func (o *CreatedInstance) GetSecurityGroups() []SecurityGroup`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *CreatedInstance) GetSecurityGroupsOk() (*[]SecurityGroup, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *CreatedInstance) SetSecurityGroups(v []SecurityGroup)`

SetSecurityGroups sets SecurityGroups field to given value.


### GetAdminPass

`func (o *CreatedInstance) GetAdminPass() string`

GetAdminPass returns the AdminPass field if non-nil, zero value otherwise.

### GetAdminPassOk

`func (o *CreatedInstance) GetAdminPassOk() (*string, bool)`

GetAdminPassOk returns a tuple with the AdminPass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPass

`func (o *CreatedInstance) SetAdminPass(v string)`

SetAdminPass sets AdminPass field to given value.

### HasAdminPass

`func (o *CreatedInstance) HasAdminPass() bool`

HasAdminPass returns a boolean if a field has been set.

### SetAdminPassNil

`func (o *CreatedInstance) SetAdminPassNil(b bool)`

 SetAdminPassNil sets the value for AdminPass to be an explicit nil

### UnsetAdminPass
`func (o *CreatedInstance) UnsetAdminPass()`

UnsetAdminPass ensures that no value is present for AdminPass, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


