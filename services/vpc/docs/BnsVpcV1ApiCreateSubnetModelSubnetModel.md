# BnsVpcV1ApiCreateSubnetModelSubnetModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 서브넷의 고유 식별 ID | 
**Name** | **string** | 서브넷 이름 | 
**VpcId** | **string** | 서브넷이 속한 VPC의 ID | 
**ProvisioningStatus** | [**ProvisioningStatus**](ProvisioningStatus.md) | 프로비저닝 상태 | 
**AvailabilityZone** | [**AvailabilityZone**](AvailabilityZone.md) | 서브넷이 위치한 가용 영역 | 
**CidrBlock** | **string** | 서브넷의 IPv4 CIDR 블록 | 
**ProjectId** | **string** | 서브넷이 속한 프로젝트의 ID | 
**CreatedAt** | **time.Time** | 리소스가 생성된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
**UpdatedAt** | **time.Time** | 리소스가 마지막으로 수정된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 

## Methods

### NewBnsVpcV1ApiCreateSubnetModelSubnetModel

`func NewBnsVpcV1ApiCreateSubnetModelSubnetModel(id string, name string, vpcId string, provisioningStatus ProvisioningStatus, availabilityZone AvailabilityZone, cidrBlock string, projectId string, createdAt time.Time, updatedAt time.Time, ) *BnsVpcV1ApiCreateSubnetModelSubnetModel`

NewBnsVpcV1ApiCreateSubnetModelSubnetModel instantiates a new BnsVpcV1ApiCreateSubnetModelSubnetModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsVpcV1ApiCreateSubnetModelSubnetModelWithDefaults

`func NewBnsVpcV1ApiCreateSubnetModelSubnetModelWithDefaults() *BnsVpcV1ApiCreateSubnetModelSubnetModel`

NewBnsVpcV1ApiCreateSubnetModelSubnetModelWithDefaults instantiates a new BnsVpcV1ApiCreateSubnetModelSubnetModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsVpcV1ApiCreateSubnetModelSubnetModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsVpcV1ApiCreateSubnetModelSubnetModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsVpcV1ApiCreateSubnetModelSubnetModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BnsVpcV1ApiCreateSubnetModelSubnetModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BnsVpcV1ApiCreateSubnetModelSubnetModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BnsVpcV1ApiCreateSubnetModelSubnetModel) SetName(v string)`

SetName sets Name field to given value.


### GetVpcId

`func (o *BnsVpcV1ApiCreateSubnetModelSubnetModel) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *BnsVpcV1ApiCreateSubnetModelSubnetModel) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *BnsVpcV1ApiCreateSubnetModelSubnetModel) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.


### GetProvisioningStatus

`func (o *BnsVpcV1ApiCreateSubnetModelSubnetModel) GetProvisioningStatus() ProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *BnsVpcV1ApiCreateSubnetModelSubnetModel) GetProvisioningStatusOk() (*ProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *BnsVpcV1ApiCreateSubnetModelSubnetModel) SetProvisioningStatus(v ProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.


### GetAvailabilityZone

`func (o *BnsVpcV1ApiCreateSubnetModelSubnetModel) GetAvailabilityZone() AvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *BnsVpcV1ApiCreateSubnetModelSubnetModel) GetAvailabilityZoneOk() (*AvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *BnsVpcV1ApiCreateSubnetModelSubnetModel) SetAvailabilityZone(v AvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.


### GetCidrBlock

`func (o *BnsVpcV1ApiCreateSubnetModelSubnetModel) GetCidrBlock() string`

GetCidrBlock returns the CidrBlock field if non-nil, zero value otherwise.

### GetCidrBlockOk

`func (o *BnsVpcV1ApiCreateSubnetModelSubnetModel) GetCidrBlockOk() (*string, bool)`

GetCidrBlockOk returns a tuple with the CidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrBlock

`func (o *BnsVpcV1ApiCreateSubnetModelSubnetModel) SetCidrBlock(v string)`

SetCidrBlock sets CidrBlock field to given value.


### GetProjectId

`func (o *BnsVpcV1ApiCreateSubnetModelSubnetModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BnsVpcV1ApiCreateSubnetModelSubnetModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BnsVpcV1ApiCreateSubnetModelSubnetModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetCreatedAt

`func (o *BnsVpcV1ApiCreateSubnetModelSubnetModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BnsVpcV1ApiCreateSubnetModelSubnetModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BnsVpcV1ApiCreateSubnetModelSubnetModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *BnsVpcV1ApiCreateSubnetModelSubnetModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BnsVpcV1ApiCreateSubnetModelSubnetModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BnsVpcV1ApiCreateSubnetModelSubnetModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


