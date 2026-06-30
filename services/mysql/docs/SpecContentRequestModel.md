# SpecContentRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatabaseUserName** | **string** | 초기 MySQL 관리자 계정 이름 | 
**DatabaseUserPassword** | **string** | 초기 MySQL 관리자 계정 비밀번호 | 
**PrimaryPort** | **int32** | Primary 인스턴스용 MySQL 포트(1024 ~ 65535) | 
**StandbyPort** | Pointer to **NullableInt32** |  | [optional] 
**EngineVersion** | **string** | [MySQL 엔진 버전](/openapi/data-store/mysql/list-available-mysql-engine-versions) | 
**FlavorId** | **string** | [인스턴스에 적용할 Flavor ID](/openapi/data-store/mysql/list-mysql-instance-types-flavors) | 
**LogDiskSize** | **int32** | 로그 디스크 크기 (GB, 100 ~ 16384) | 
**DataDiskSize** | **int32** | 데이터 디스크 크기 (GB, 100 ~ 16384) | 

## Methods

### NewSpecContentRequestModel

`func NewSpecContentRequestModel(databaseUserName string, databaseUserPassword string, primaryPort int32, engineVersion string, flavorId string, logDiskSize int32, dataDiskSize int32, ) *SpecContentRequestModel`

NewSpecContentRequestModel instantiates a new SpecContentRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpecContentRequestModelWithDefaults

`func NewSpecContentRequestModelWithDefaults() *SpecContentRequestModel`

NewSpecContentRequestModelWithDefaults instantiates a new SpecContentRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabaseUserName

`func (o *SpecContentRequestModel) GetDatabaseUserName() string`

GetDatabaseUserName returns the DatabaseUserName field if non-nil, zero value otherwise.

### GetDatabaseUserNameOk

`func (o *SpecContentRequestModel) GetDatabaseUserNameOk() (*string, bool)`

GetDatabaseUserNameOk returns a tuple with the DatabaseUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseUserName

`func (o *SpecContentRequestModel) SetDatabaseUserName(v string)`

SetDatabaseUserName sets DatabaseUserName field to given value.


### GetDatabaseUserPassword

`func (o *SpecContentRequestModel) GetDatabaseUserPassword() string`

GetDatabaseUserPassword returns the DatabaseUserPassword field if non-nil, zero value otherwise.

### GetDatabaseUserPasswordOk

`func (o *SpecContentRequestModel) GetDatabaseUserPasswordOk() (*string, bool)`

GetDatabaseUserPasswordOk returns a tuple with the DatabaseUserPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseUserPassword

`func (o *SpecContentRequestModel) SetDatabaseUserPassword(v string)`

SetDatabaseUserPassword sets DatabaseUserPassword field to given value.


### GetPrimaryPort

`func (o *SpecContentRequestModel) GetPrimaryPort() int32`

GetPrimaryPort returns the PrimaryPort field if non-nil, zero value otherwise.

### GetPrimaryPortOk

`func (o *SpecContentRequestModel) GetPrimaryPortOk() (*int32, bool)`

GetPrimaryPortOk returns a tuple with the PrimaryPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryPort

`func (o *SpecContentRequestModel) SetPrimaryPort(v int32)`

SetPrimaryPort sets PrimaryPort field to given value.


### GetStandbyPort

`func (o *SpecContentRequestModel) GetStandbyPort() int32`

GetStandbyPort returns the StandbyPort field if non-nil, zero value otherwise.

### GetStandbyPortOk

`func (o *SpecContentRequestModel) GetStandbyPortOk() (*int32, bool)`

GetStandbyPortOk returns a tuple with the StandbyPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandbyPort

`func (o *SpecContentRequestModel) SetStandbyPort(v int32)`

SetStandbyPort sets StandbyPort field to given value.

### HasStandbyPort

`func (o *SpecContentRequestModel) HasStandbyPort() bool`

HasStandbyPort returns a boolean if a field has been set.

### SetStandbyPortNil

`func (o *SpecContentRequestModel) SetStandbyPortNil(b bool)`

 SetStandbyPortNil sets the value for StandbyPort to be an explicit nil

### UnsetStandbyPort
`func (o *SpecContentRequestModel) UnsetStandbyPort()`

UnsetStandbyPort ensures that no value is present for StandbyPort, not even an explicit nil
### GetEngineVersion

`func (o *SpecContentRequestModel) GetEngineVersion() string`

GetEngineVersion returns the EngineVersion field if non-nil, zero value otherwise.

### GetEngineVersionOk

`func (o *SpecContentRequestModel) GetEngineVersionOk() (*string, bool)`

GetEngineVersionOk returns a tuple with the EngineVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineVersion

`func (o *SpecContentRequestModel) SetEngineVersion(v string)`

SetEngineVersion sets EngineVersion field to given value.


### GetFlavorId

`func (o *SpecContentRequestModel) GetFlavorId() string`

GetFlavorId returns the FlavorId field if non-nil, zero value otherwise.

### GetFlavorIdOk

`func (o *SpecContentRequestModel) GetFlavorIdOk() (*string, bool)`

GetFlavorIdOk returns a tuple with the FlavorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavorId

`func (o *SpecContentRequestModel) SetFlavorId(v string)`

SetFlavorId sets FlavorId field to given value.


### GetLogDiskSize

`func (o *SpecContentRequestModel) GetLogDiskSize() int32`

GetLogDiskSize returns the LogDiskSize field if non-nil, zero value otherwise.

### GetLogDiskSizeOk

`func (o *SpecContentRequestModel) GetLogDiskSizeOk() (*int32, bool)`

GetLogDiskSizeOk returns a tuple with the LogDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogDiskSize

`func (o *SpecContentRequestModel) SetLogDiskSize(v int32)`

SetLogDiskSize sets LogDiskSize field to given value.


### GetDataDiskSize

`func (o *SpecContentRequestModel) GetDataDiskSize() int32`

GetDataDiskSize returns the DataDiskSize field if non-nil, zero value otherwise.

### GetDataDiskSizeOk

`func (o *SpecContentRequestModel) GetDataDiskSizeOk() (*int32, bool)`

GetDataDiskSizeOk returns a tuple with the DataDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDiskSize

`func (o *SpecContentRequestModel) SetDataDiskSize(v int32)`

SetDataDiskSize sets DataDiskSize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


