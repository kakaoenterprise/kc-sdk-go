# BnsVpcV1ApiCreateRouteTableModelRouteTableModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 라우팅 테이블 ID | 
**Name** | **string** | 라우팅 테이블 이름 | 
**VpcId** | **string** | 라우팅 테이블이 속한 VPC의 ID | 
**ProvisioningStatus** | [**ProvisioningStatus**](ProvisioningStatus.md) | 프로비저닝 상태 | 
**ProjectId** | **string** | 라우팅 테이블이 속한 프로젝트 ID | 
**CreatedAt** | **time.Time** | 리소스가 생성된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
**UpdatedAt** | **time.Time** | 리소스가 마지막으로 수정된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 

## Methods

### NewBnsVpcV1ApiCreateRouteTableModelRouteTableModel

`func NewBnsVpcV1ApiCreateRouteTableModelRouteTableModel(id string, name string, vpcId string, provisioningStatus ProvisioningStatus, projectId string, createdAt time.Time, updatedAt time.Time, ) *BnsVpcV1ApiCreateRouteTableModelRouteTableModel`

NewBnsVpcV1ApiCreateRouteTableModelRouteTableModel instantiates a new BnsVpcV1ApiCreateRouteTableModelRouteTableModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsVpcV1ApiCreateRouteTableModelRouteTableModelWithDefaults

`func NewBnsVpcV1ApiCreateRouteTableModelRouteTableModelWithDefaults() *BnsVpcV1ApiCreateRouteTableModelRouteTableModel`

NewBnsVpcV1ApiCreateRouteTableModelRouteTableModelWithDefaults instantiates a new BnsVpcV1ApiCreateRouteTableModelRouteTableModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsVpcV1ApiCreateRouteTableModelRouteTableModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsVpcV1ApiCreateRouteTableModelRouteTableModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsVpcV1ApiCreateRouteTableModelRouteTableModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BnsVpcV1ApiCreateRouteTableModelRouteTableModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BnsVpcV1ApiCreateRouteTableModelRouteTableModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BnsVpcV1ApiCreateRouteTableModelRouteTableModel) SetName(v string)`

SetName sets Name field to given value.


### GetVpcId

`func (o *BnsVpcV1ApiCreateRouteTableModelRouteTableModel) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *BnsVpcV1ApiCreateRouteTableModelRouteTableModel) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *BnsVpcV1ApiCreateRouteTableModelRouteTableModel) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.


### GetProvisioningStatus

`func (o *BnsVpcV1ApiCreateRouteTableModelRouteTableModel) GetProvisioningStatus() ProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *BnsVpcV1ApiCreateRouteTableModelRouteTableModel) GetProvisioningStatusOk() (*ProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *BnsVpcV1ApiCreateRouteTableModelRouteTableModel) SetProvisioningStatus(v ProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.


### GetProjectId

`func (o *BnsVpcV1ApiCreateRouteTableModelRouteTableModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BnsVpcV1ApiCreateRouteTableModelRouteTableModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BnsVpcV1ApiCreateRouteTableModelRouteTableModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetCreatedAt

`func (o *BnsVpcV1ApiCreateRouteTableModelRouteTableModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BnsVpcV1ApiCreateRouteTableModelRouteTableModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BnsVpcV1ApiCreateRouteTableModelRouteTableModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *BnsVpcV1ApiCreateRouteTableModelRouteTableModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BnsVpcV1ApiCreateRouteTableModelRouteTableModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BnsVpcV1ApiCreateRouteTableModelRouteTableModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


