# BnsVpcV1ApiUpdateRouteTableAssociationModelAssociationModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 연결(association)의 ID | 
**RouteTableId** | **string** | 연결된 라우팅 테이블의 ID | 
**VpcId** | **string** | 연결이 속한 VPC의 ID | 
**ProvisioningStatus** | [**ProvisioningStatus**](ProvisioningStatus.md) | 프로비저닝 상태 | 
**SubnetId** | Pointer to **NullableString** |  | [optional] 
**ProjectId** | **string** | 연결이 속한 프로젝트의 ID | 
**Description** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | **time.Time** | 리소스가 생성된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
**UpdatedAt** | **time.Time** | 리소스가 마지막으로 수정된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 

## Methods

### NewBnsVpcV1ApiUpdateRouteTableAssociationModelAssociationModel

`func NewBnsVpcV1ApiUpdateRouteTableAssociationModelAssociationModel(id string, routeTableId string, vpcId string, provisioningStatus ProvisioningStatus, projectId string, createdAt time.Time, updatedAt time.Time, ) *BnsVpcV1ApiUpdateRouteTableAssociationModelAssociationModel`

NewBnsVpcV1ApiUpdateRouteTableAssociationModelAssociationModel instantiates a new BnsVpcV1ApiUpdateRouteTableAssociationModelAssociationModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsVpcV1ApiUpdateRouteTableAssociationModelAssociationModelWithDefaults

`func NewBnsVpcV1ApiUpdateRouteTableAssociationModelAssociationModelWithDefaults() *BnsVpcV1ApiUpdateRouteTableAssociationModelAssociationModel`

NewBnsVpcV1ApiUpdateRouteTableAssociationModelAssociationModelWithDefaults instantiates a new BnsVpcV1ApiUpdateRouteTableAssociationModelAssociationModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsVpcV1ApiUpdateRouteTableAssociationModelAssociationModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsVpcV1ApiUpdateRouteTableAssociationModelAssociationModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsVpcV1ApiUpdateRouteTableAssociationModelAssociationModel) SetId(v string)`

SetId sets Id field to given value.


### GetRouteTableId

`func (o *BnsVpcV1ApiUpdateRouteTableAssociationModelAssociationModel) GetRouteTableId() string`

GetRouteTableId returns the RouteTableId field if non-nil, zero value otherwise.

### GetRouteTableIdOk

`func (o *BnsVpcV1ApiUpdateRouteTableAssociationModelAssociationModel) GetRouteTableIdOk() (*string, bool)`

GetRouteTableIdOk returns a tuple with the RouteTableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteTableId

`func (o *BnsVpcV1ApiUpdateRouteTableAssociationModelAssociationModel) SetRouteTableId(v string)`

SetRouteTableId sets RouteTableId field to given value.


### GetVpcId

`func (o *BnsVpcV1ApiUpdateRouteTableAssociationModelAssociationModel) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *BnsVpcV1ApiUpdateRouteTableAssociationModelAssociationModel) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *BnsVpcV1ApiUpdateRouteTableAssociationModelAssociationModel) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.


### GetProvisioningStatus

`func (o *BnsVpcV1ApiUpdateRouteTableAssociationModelAssociationModel) GetProvisioningStatus() ProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *BnsVpcV1ApiUpdateRouteTableAssociationModelAssociationModel) GetProvisioningStatusOk() (*ProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *BnsVpcV1ApiUpdateRouteTableAssociationModelAssociationModel) SetProvisioningStatus(v ProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.


### GetSubnetId

`func (o *BnsVpcV1ApiUpdateRouteTableAssociationModelAssociationModel) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *BnsVpcV1ApiUpdateRouteTableAssociationModelAssociationModel) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *BnsVpcV1ApiUpdateRouteTableAssociationModelAssociationModel) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.

### HasSubnetId

`func (o *BnsVpcV1ApiUpdateRouteTableAssociationModelAssociationModel) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.

### SetSubnetIdNil

`func (o *BnsVpcV1ApiUpdateRouteTableAssociationModelAssociationModel) SetSubnetIdNil(b bool)`

 SetSubnetIdNil sets the value for SubnetId to be an explicit nil

### UnsetSubnetId
`func (o *BnsVpcV1ApiUpdateRouteTableAssociationModelAssociationModel) UnsetSubnetId()`

UnsetSubnetId ensures that no value is present for SubnetId, not even an explicit nil
### GetProjectId

`func (o *BnsVpcV1ApiUpdateRouteTableAssociationModelAssociationModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BnsVpcV1ApiUpdateRouteTableAssociationModelAssociationModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BnsVpcV1ApiUpdateRouteTableAssociationModelAssociationModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetDescription

`func (o *BnsVpcV1ApiUpdateRouteTableAssociationModelAssociationModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BnsVpcV1ApiUpdateRouteTableAssociationModelAssociationModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BnsVpcV1ApiUpdateRouteTableAssociationModelAssociationModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BnsVpcV1ApiUpdateRouteTableAssociationModelAssociationModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BnsVpcV1ApiUpdateRouteTableAssociationModelAssociationModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BnsVpcV1ApiUpdateRouteTableAssociationModelAssociationModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCreatedAt

`func (o *BnsVpcV1ApiUpdateRouteTableAssociationModelAssociationModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BnsVpcV1ApiUpdateRouteTableAssociationModelAssociationModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BnsVpcV1ApiUpdateRouteTableAssociationModelAssociationModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *BnsVpcV1ApiUpdateRouteTableAssociationModelAssociationModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BnsVpcV1ApiUpdateRouteTableAssociationModelAssociationModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BnsVpcV1ApiUpdateRouteTableAssociationModelAssociationModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


