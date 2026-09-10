# TgwRouteTableResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 라우팅 테이블 ID | 
**Name** | **string** | 라우팅 테이블 이름 | 
**TgwId** | **string** | 연결된 Transit Gateway ID | 
**IsDefaultAssociationRouteTable** | **bool** | 기본 Association 라우팅 테이블 여부 | 
**IsDefaultPropagationRouteTable** | **bool** | 기본 Propagation 라우팅 테이블 여부 | 
**Region** | **string** | 라우팅 테이블이 위치한 리전 (예: &#x60;kr-central-2&#x60;) | 
**ProjectId** | **string** | 라우팅 테이블이 속한 프로젝트 ID | 
**ProvisioningStatus** | [**TGWRouteTableProvisioningStatus**](TGWRouteTableProvisioningStatus.md) | 프로비저닝 상태 | 
**CreatedAt** | **time.Time** | 리소스가 생성된 시간 &lt;br/&gt; - ISO 8601 형식 &lt;br/&gt; - UTC 기준 | 
**UpdatedAt** | **time.Time** | 리소스가 마지막으로 수정된 시간 &lt;br/&gt; - ISO 8601 형식 &lt;br/&gt; - UTC 기준 | 

## Methods

### NewTgwRouteTableResult

`func NewTgwRouteTableResult(id string, name string, tgwId string, isDefaultAssociationRouteTable bool, isDefaultPropagationRouteTable bool, region string, projectId string, provisioningStatus TGWRouteTableProvisioningStatus, createdAt time.Time, updatedAt time.Time, ) *TgwRouteTableResult`

NewTgwRouteTableResult instantiates a new TgwRouteTableResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTgwRouteTableResultWithDefaults

`func NewTgwRouteTableResultWithDefaults() *TgwRouteTableResult`

NewTgwRouteTableResultWithDefaults instantiates a new TgwRouteTableResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TgwRouteTableResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TgwRouteTableResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TgwRouteTableResult) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *TgwRouteTableResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TgwRouteTableResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TgwRouteTableResult) SetName(v string)`

SetName sets Name field to given value.


### GetTgwId

`func (o *TgwRouteTableResult) GetTgwId() string`

GetTgwId returns the TgwId field if non-nil, zero value otherwise.

### GetTgwIdOk

`func (o *TgwRouteTableResult) GetTgwIdOk() (*string, bool)`

GetTgwIdOk returns a tuple with the TgwId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgwId

`func (o *TgwRouteTableResult) SetTgwId(v string)`

SetTgwId sets TgwId field to given value.


### GetIsDefaultAssociationRouteTable

`func (o *TgwRouteTableResult) GetIsDefaultAssociationRouteTable() bool`

GetIsDefaultAssociationRouteTable returns the IsDefaultAssociationRouteTable field if non-nil, zero value otherwise.

### GetIsDefaultAssociationRouteTableOk

`func (o *TgwRouteTableResult) GetIsDefaultAssociationRouteTableOk() (*bool, bool)`

GetIsDefaultAssociationRouteTableOk returns a tuple with the IsDefaultAssociationRouteTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultAssociationRouteTable

`func (o *TgwRouteTableResult) SetIsDefaultAssociationRouteTable(v bool)`

SetIsDefaultAssociationRouteTable sets IsDefaultAssociationRouteTable field to given value.


### GetIsDefaultPropagationRouteTable

`func (o *TgwRouteTableResult) GetIsDefaultPropagationRouteTable() bool`

GetIsDefaultPropagationRouteTable returns the IsDefaultPropagationRouteTable field if non-nil, zero value otherwise.

### GetIsDefaultPropagationRouteTableOk

`func (o *TgwRouteTableResult) GetIsDefaultPropagationRouteTableOk() (*bool, bool)`

GetIsDefaultPropagationRouteTableOk returns a tuple with the IsDefaultPropagationRouteTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultPropagationRouteTable

`func (o *TgwRouteTableResult) SetIsDefaultPropagationRouteTable(v bool)`

SetIsDefaultPropagationRouteTable sets IsDefaultPropagationRouteTable field to given value.


### GetRegion

`func (o *TgwRouteTableResult) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *TgwRouteTableResult) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *TgwRouteTableResult) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetProjectId

`func (o *TgwRouteTableResult) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *TgwRouteTableResult) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *TgwRouteTableResult) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetProvisioningStatus

`func (o *TgwRouteTableResult) GetProvisioningStatus() TGWRouteTableProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *TgwRouteTableResult) GetProvisioningStatusOk() (*TGWRouteTableProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *TgwRouteTableResult) SetProvisioningStatus(v TGWRouteTableProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.


### GetCreatedAt

`func (o *TgwRouteTableResult) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TgwRouteTableResult) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TgwRouteTableResult) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *TgwRouteTableResult) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TgwRouteTableResult) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TgwRouteTableResult) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


