# BnsTgwV1ApiCreateTgwRouteTableModelTgwRouteTableResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 라우팅 테이블 ID | 
**Name** | **string** | 라우팅 테이블 이름 | 
**TgwId** | **string** | 라우팅 테이블이 속한 Transit Gateway ID | 
**IsDefaultAssociationRouteTable** | **bool** | 기본 Association 라우팅 테이블 여부 | 
**IsDefaultPropagationRouteTable** | **bool** | 기본 Propagation 라우팅 테이블 여부 | 
**Region** | [**Region**](Region.md) | 라우팅 테이블이 생성된 리전 | 
**ProjectId** | **string** | 프로젝트 ID | 
**ProvisioningStatus** | [**TGWProvisioningStatus**](TGWProvisioningStatus.md) | 라우팅 테이블의 프로비저닝 상태 | 
**CreatedAt** | **time.Time** | 리소스가 생성된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
**UpdatedAt** | **time.Time** | 리소스가 마지막으로 수정된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 

## Methods

### NewBnsTgwV1ApiCreateTgwRouteTableModelTgwRouteTableResponseModel

`func NewBnsTgwV1ApiCreateTgwRouteTableModelTgwRouteTableResponseModel(id string, name string, tgwId string, isDefaultAssociationRouteTable bool, isDefaultPropagationRouteTable bool, region Region, projectId string, provisioningStatus TGWProvisioningStatus, createdAt time.Time, updatedAt time.Time, ) *BnsTgwV1ApiCreateTgwRouteTableModelTgwRouteTableResponseModel`

NewBnsTgwV1ApiCreateTgwRouteTableModelTgwRouteTableResponseModel instantiates a new BnsTgwV1ApiCreateTgwRouteTableModelTgwRouteTableResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsTgwV1ApiCreateTgwRouteTableModelTgwRouteTableResponseModelWithDefaults

`func NewBnsTgwV1ApiCreateTgwRouteTableModelTgwRouteTableResponseModelWithDefaults() *BnsTgwV1ApiCreateTgwRouteTableModelTgwRouteTableResponseModel`

NewBnsTgwV1ApiCreateTgwRouteTableModelTgwRouteTableResponseModelWithDefaults instantiates a new BnsTgwV1ApiCreateTgwRouteTableModelTgwRouteTableResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsTgwV1ApiCreateTgwRouteTableModelTgwRouteTableResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsTgwV1ApiCreateTgwRouteTableModelTgwRouteTableResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsTgwV1ApiCreateTgwRouteTableModelTgwRouteTableResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BnsTgwV1ApiCreateTgwRouteTableModelTgwRouteTableResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BnsTgwV1ApiCreateTgwRouteTableModelTgwRouteTableResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BnsTgwV1ApiCreateTgwRouteTableModelTgwRouteTableResponseModel) SetName(v string)`

SetName sets Name field to given value.


### GetTgwId

`func (o *BnsTgwV1ApiCreateTgwRouteTableModelTgwRouteTableResponseModel) GetTgwId() string`

GetTgwId returns the TgwId field if non-nil, zero value otherwise.

### GetTgwIdOk

`func (o *BnsTgwV1ApiCreateTgwRouteTableModelTgwRouteTableResponseModel) GetTgwIdOk() (*string, bool)`

GetTgwIdOk returns a tuple with the TgwId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgwId

`func (o *BnsTgwV1ApiCreateTgwRouteTableModelTgwRouteTableResponseModel) SetTgwId(v string)`

SetTgwId sets TgwId field to given value.


### GetIsDefaultAssociationRouteTable

`func (o *BnsTgwV1ApiCreateTgwRouteTableModelTgwRouteTableResponseModel) GetIsDefaultAssociationRouteTable() bool`

GetIsDefaultAssociationRouteTable returns the IsDefaultAssociationRouteTable field if non-nil, zero value otherwise.

### GetIsDefaultAssociationRouteTableOk

`func (o *BnsTgwV1ApiCreateTgwRouteTableModelTgwRouteTableResponseModel) GetIsDefaultAssociationRouteTableOk() (*bool, bool)`

GetIsDefaultAssociationRouteTableOk returns a tuple with the IsDefaultAssociationRouteTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultAssociationRouteTable

`func (o *BnsTgwV1ApiCreateTgwRouteTableModelTgwRouteTableResponseModel) SetIsDefaultAssociationRouteTable(v bool)`

SetIsDefaultAssociationRouteTable sets IsDefaultAssociationRouteTable field to given value.


### GetIsDefaultPropagationRouteTable

`func (o *BnsTgwV1ApiCreateTgwRouteTableModelTgwRouteTableResponseModel) GetIsDefaultPropagationRouteTable() bool`

GetIsDefaultPropagationRouteTable returns the IsDefaultPropagationRouteTable field if non-nil, zero value otherwise.

### GetIsDefaultPropagationRouteTableOk

`func (o *BnsTgwV1ApiCreateTgwRouteTableModelTgwRouteTableResponseModel) GetIsDefaultPropagationRouteTableOk() (*bool, bool)`

GetIsDefaultPropagationRouteTableOk returns a tuple with the IsDefaultPropagationRouteTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultPropagationRouteTable

`func (o *BnsTgwV1ApiCreateTgwRouteTableModelTgwRouteTableResponseModel) SetIsDefaultPropagationRouteTable(v bool)`

SetIsDefaultPropagationRouteTable sets IsDefaultPropagationRouteTable field to given value.


### GetRegion

`func (o *BnsTgwV1ApiCreateTgwRouteTableModelTgwRouteTableResponseModel) GetRegion() Region`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *BnsTgwV1ApiCreateTgwRouteTableModelTgwRouteTableResponseModel) GetRegionOk() (*Region, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *BnsTgwV1ApiCreateTgwRouteTableModelTgwRouteTableResponseModel) SetRegion(v Region)`

SetRegion sets Region field to given value.


### GetProjectId

`func (o *BnsTgwV1ApiCreateTgwRouteTableModelTgwRouteTableResponseModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BnsTgwV1ApiCreateTgwRouteTableModelTgwRouteTableResponseModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BnsTgwV1ApiCreateTgwRouteTableModelTgwRouteTableResponseModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetProvisioningStatus

`func (o *BnsTgwV1ApiCreateTgwRouteTableModelTgwRouteTableResponseModel) GetProvisioningStatus() TGWProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *BnsTgwV1ApiCreateTgwRouteTableModelTgwRouteTableResponseModel) GetProvisioningStatusOk() (*TGWProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *BnsTgwV1ApiCreateTgwRouteTableModelTgwRouteTableResponseModel) SetProvisioningStatus(v TGWProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.


### GetCreatedAt

`func (o *BnsTgwV1ApiCreateTgwRouteTableModelTgwRouteTableResponseModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BnsTgwV1ApiCreateTgwRouteTableModelTgwRouteTableResponseModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BnsTgwV1ApiCreateTgwRouteTableModelTgwRouteTableResponseModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *BnsTgwV1ApiCreateTgwRouteTableModelTgwRouteTableResponseModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BnsTgwV1ApiCreateTgwRouteTableModelTgwRouteTableResponseModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BnsTgwV1ApiCreateTgwRouteTableModelTgwRouteTableResponseModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


