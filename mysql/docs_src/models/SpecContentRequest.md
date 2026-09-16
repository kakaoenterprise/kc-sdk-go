# SpecContentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatabaseUserName** | **string** | 초기 MySQL 관리자 계정 이름 | 
**DatabaseUserPassword** | **string** | 초기 MySQL 관리자 계정 비밀번호 | 
**PrimaryPort** | **int32** | Primary 인스턴스용 MySQL 포트(1024 ~ 65535) | 
**StandbyPort** | Pointer to **NullableInt32** | Standby 인스턴스용 MySQL 포트(1024 ~ 65535) | [optional] 
**EngineVersion** | **string** | [MySQL 엔진 버전](/openapi/data-store/mysql/list-available-mysql-engine-versions) | 
**FlavorId** | **string** | [인스턴스에 적용할 Flavor ID](/openapi/data-store/mysql/list-mysql-instance-types-flavors) | 
**LogDiskSize** | **int32** | 로그 디스크 크기 (GB, 100 ~ 16384) | 
**DataDiskSize** | **int32** | 데이터 디스크 크기 (GB, 100 ~ 16384) | 

## Methods

### NewSpecContentRequest

`func NewSpecContentRequest(databaseUserName string, databaseUserPassword string, primaryPort int32, engineVersion string, flavorId string, logDiskSize int32, dataDiskSize int32, ) *SpecContentRequest`

NewSpecContentRequest instantiates a new SpecContentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpecContentRequestWithDefaults

`func NewSpecContentRequestWithDefaults() *SpecContentRequest`

NewSpecContentRequestWithDefaults instantiates a new SpecContentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabaseUserName

`func (o *SpecContentRequest) GetDatabaseUserName() string`

GetDatabaseUserName returns the DatabaseUserName field if non-nil, zero value otherwise.

### GetDatabaseUserNameOk

`func (o *SpecContentRequest) GetDatabaseUserNameOk() (*string, bool)`

GetDatabaseUserNameOk returns a tuple with the DatabaseUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseUserName

`func (o *SpecContentRequest) SetDatabaseUserName(v string)`

SetDatabaseUserName sets DatabaseUserName field to given value.


### GetDatabaseUserPassword

`func (o *SpecContentRequest) GetDatabaseUserPassword() string`

GetDatabaseUserPassword returns the DatabaseUserPassword field if non-nil, zero value otherwise.

### GetDatabaseUserPasswordOk

`func (o *SpecContentRequest) GetDatabaseUserPasswordOk() (*string, bool)`

GetDatabaseUserPasswordOk returns a tuple with the DatabaseUserPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseUserPassword

`func (o *SpecContentRequest) SetDatabaseUserPassword(v string)`

SetDatabaseUserPassword sets DatabaseUserPassword field to given value.


### GetPrimaryPort

`func (o *SpecContentRequest) GetPrimaryPort() int32`

GetPrimaryPort returns the PrimaryPort field if non-nil, zero value otherwise.

### GetPrimaryPortOk

`func (o *SpecContentRequest) GetPrimaryPortOk() (*int32, bool)`

GetPrimaryPortOk returns a tuple with the PrimaryPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryPort

`func (o *SpecContentRequest) SetPrimaryPort(v int32)`

SetPrimaryPort sets PrimaryPort field to given value.


### GetStandbyPort

`func (o *SpecContentRequest) GetStandbyPort() int32`

GetStandbyPort returns the StandbyPort field if non-nil, zero value otherwise.

### GetStandbyPortOk

`func (o *SpecContentRequest) GetStandbyPortOk() (*int32, bool)`

GetStandbyPortOk returns a tuple with the StandbyPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandbyPort

`func (o *SpecContentRequest) SetStandbyPort(v int32)`

SetStandbyPort sets StandbyPort field to given value.

### HasStandbyPort

`func (o *SpecContentRequest) HasStandbyPort() bool`

HasStandbyPort returns a boolean if a field has been set.

### SetStandbyPortNil

`func (o *SpecContentRequest) SetStandbyPortNil(b bool)`

 SetStandbyPortNil sets the value for StandbyPort to be an explicit nil

### UnsetStandbyPort
`func (o *SpecContentRequest) UnsetStandbyPort()`

UnsetStandbyPort ensures that no value is present for StandbyPort, not even an explicit nil
### GetEngineVersion

`func (o *SpecContentRequest) GetEngineVersion() string`

GetEngineVersion returns the EngineVersion field if non-nil, zero value otherwise.

### GetEngineVersionOk

`func (o *SpecContentRequest) GetEngineVersionOk() (*string, bool)`

GetEngineVersionOk returns a tuple with the EngineVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineVersion

`func (o *SpecContentRequest) SetEngineVersion(v string)`

SetEngineVersion sets EngineVersion field to given value.


### GetFlavorId

`func (o *SpecContentRequest) GetFlavorId() string`

GetFlavorId returns the FlavorId field if non-nil, zero value otherwise.

### GetFlavorIdOk

`func (o *SpecContentRequest) GetFlavorIdOk() (*string, bool)`

GetFlavorIdOk returns a tuple with the FlavorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavorId

`func (o *SpecContentRequest) SetFlavorId(v string)`

SetFlavorId sets FlavorId field to given value.


### GetLogDiskSize

`func (o *SpecContentRequest) GetLogDiskSize() int32`

GetLogDiskSize returns the LogDiskSize field if non-nil, zero value otherwise.

### GetLogDiskSizeOk

`func (o *SpecContentRequest) GetLogDiskSizeOk() (*int32, bool)`

GetLogDiskSizeOk returns a tuple with the LogDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogDiskSize

`func (o *SpecContentRequest) SetLogDiskSize(v int32)`

SetLogDiskSize sets LogDiskSize field to given value.


### GetDataDiskSize

`func (o *SpecContentRequest) GetDataDiskSize() int32`

GetDataDiskSize returns the DataDiskSize field if non-nil, zero value otherwise.

### GetDataDiskSizeOk

`func (o *SpecContentRequest) GetDataDiskSizeOk() (*int32, bool)`

GetDataDiskSizeOk returns a tuple with the DataDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDiskSize

`func (o *SpecContentRequest) SetDataDiskSize(v int32)`

SetDataDiskSize sets DataDiskSize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


