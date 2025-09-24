# BnsVpcV1ApiCreateVpcModelVpcModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | VPC의 고유 ID | 
**ProjectId** | **string** | VPC가 속한 프로젝트의 ID | 
**Name** | **string** | VPC 이름 | 
**IsDefault** | **bool** | 기본 VPC 여부 | 
**Region** | [**Region**](Region.md) | VPC가 속한 리전 | 
**CidrBlock** | **string** | VPC의 IP 주소 범위를 나타내는 CIDR | 
**ProvisioningStatus** | [**ProvisioningStatus**](ProvisioningStatus.md) | 프로비저닝 상태 | 
**CreatedAt** | **time.Time** | 리소스가 생성된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
**UpdatedAt** | **time.Time** | 리소스가 마지막으로 수정된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
**Description** | **string** | VPC에 대한 설명 | 

## Methods

### NewBnsVpcV1ApiCreateVpcModelVpcModel

`func NewBnsVpcV1ApiCreateVpcModelVpcModel(id string, projectId string, name string, isDefault bool, region Region, cidrBlock string, provisioningStatus ProvisioningStatus, createdAt time.Time, updatedAt time.Time, description string, ) *BnsVpcV1ApiCreateVpcModelVpcModel`

NewBnsVpcV1ApiCreateVpcModelVpcModel instantiates a new BnsVpcV1ApiCreateVpcModelVpcModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsVpcV1ApiCreateVpcModelVpcModelWithDefaults

`func NewBnsVpcV1ApiCreateVpcModelVpcModelWithDefaults() *BnsVpcV1ApiCreateVpcModelVpcModel`

NewBnsVpcV1ApiCreateVpcModelVpcModelWithDefaults instantiates a new BnsVpcV1ApiCreateVpcModelVpcModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsVpcV1ApiCreateVpcModelVpcModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsVpcV1ApiCreateVpcModelVpcModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsVpcV1ApiCreateVpcModelVpcModel) SetId(v string)`

SetId sets Id field to given value.


### GetProjectId

`func (o *BnsVpcV1ApiCreateVpcModelVpcModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BnsVpcV1ApiCreateVpcModelVpcModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BnsVpcV1ApiCreateVpcModelVpcModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetName

`func (o *BnsVpcV1ApiCreateVpcModelVpcModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BnsVpcV1ApiCreateVpcModelVpcModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BnsVpcV1ApiCreateVpcModelVpcModel) SetName(v string)`

SetName sets Name field to given value.


### GetIsDefault

`func (o *BnsVpcV1ApiCreateVpcModelVpcModel) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *BnsVpcV1ApiCreateVpcModelVpcModel) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *BnsVpcV1ApiCreateVpcModelVpcModel) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetRegion

`func (o *BnsVpcV1ApiCreateVpcModelVpcModel) GetRegion() Region`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *BnsVpcV1ApiCreateVpcModelVpcModel) GetRegionOk() (*Region, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *BnsVpcV1ApiCreateVpcModelVpcModel) SetRegion(v Region)`

SetRegion sets Region field to given value.


### GetCidrBlock

`func (o *BnsVpcV1ApiCreateVpcModelVpcModel) GetCidrBlock() string`

GetCidrBlock returns the CidrBlock field if non-nil, zero value otherwise.

### GetCidrBlockOk

`func (o *BnsVpcV1ApiCreateVpcModelVpcModel) GetCidrBlockOk() (*string, bool)`

GetCidrBlockOk returns a tuple with the CidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrBlock

`func (o *BnsVpcV1ApiCreateVpcModelVpcModel) SetCidrBlock(v string)`

SetCidrBlock sets CidrBlock field to given value.


### GetProvisioningStatus

`func (o *BnsVpcV1ApiCreateVpcModelVpcModel) GetProvisioningStatus() ProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *BnsVpcV1ApiCreateVpcModelVpcModel) GetProvisioningStatusOk() (*ProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *BnsVpcV1ApiCreateVpcModelVpcModel) SetProvisioningStatus(v ProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.


### GetCreatedAt

`func (o *BnsVpcV1ApiCreateVpcModelVpcModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BnsVpcV1ApiCreateVpcModelVpcModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BnsVpcV1ApiCreateVpcModelVpcModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *BnsVpcV1ApiCreateVpcModelVpcModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BnsVpcV1ApiCreateVpcModelVpcModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BnsVpcV1ApiCreateVpcModelVpcModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDescription

`func (o *BnsVpcV1ApiCreateVpcModelVpcModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BnsVpcV1ApiCreateVpcModelVpcModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BnsVpcV1ApiCreateVpcModelVpcModel) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


