# BnsTgwV1ApiCreateTransitGatewayModelTgwResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Transit Gateway ID | 
**Name** | **string** | Transit Gateway 이름 | 
**Region** | [**Region**](Region.md) | Transit Gateway가 생성된 리전 | 
**Options** | [**TgwOptionsResponseModel**](TgwOptionsResponseModel.md) | Transit Gateway 옵션 설정 | 
**ProvisioningStatus** | [**TGWProvisioningStatus**](TGWProvisioningStatus.md) | Transit Gateway의 프로비저닝 상태 | 
**ProjectId** | **string** | 프로젝트 ID | 
**CreatedAt** | **time.Time** | 리소스가 생성된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
**UpdatedAt** | **time.Time** | 리소스가 마지막으로 수정된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 

## Methods

### NewBnsTgwV1ApiCreateTransitGatewayModelTgwResponseModel

`func NewBnsTgwV1ApiCreateTransitGatewayModelTgwResponseModel(id string, name string, region Region, options TgwOptionsResponseModel, provisioningStatus TGWProvisioningStatus, projectId string, createdAt time.Time, updatedAt time.Time, ) *BnsTgwV1ApiCreateTransitGatewayModelTgwResponseModel`

NewBnsTgwV1ApiCreateTransitGatewayModelTgwResponseModel instantiates a new BnsTgwV1ApiCreateTransitGatewayModelTgwResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsTgwV1ApiCreateTransitGatewayModelTgwResponseModelWithDefaults

`func NewBnsTgwV1ApiCreateTransitGatewayModelTgwResponseModelWithDefaults() *BnsTgwV1ApiCreateTransitGatewayModelTgwResponseModel`

NewBnsTgwV1ApiCreateTransitGatewayModelTgwResponseModelWithDefaults instantiates a new BnsTgwV1ApiCreateTransitGatewayModelTgwResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsTgwV1ApiCreateTransitGatewayModelTgwResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsTgwV1ApiCreateTransitGatewayModelTgwResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsTgwV1ApiCreateTransitGatewayModelTgwResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BnsTgwV1ApiCreateTransitGatewayModelTgwResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BnsTgwV1ApiCreateTransitGatewayModelTgwResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BnsTgwV1ApiCreateTransitGatewayModelTgwResponseModel) SetName(v string)`

SetName sets Name field to given value.


### GetRegion

`func (o *BnsTgwV1ApiCreateTransitGatewayModelTgwResponseModel) GetRegion() Region`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *BnsTgwV1ApiCreateTransitGatewayModelTgwResponseModel) GetRegionOk() (*Region, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *BnsTgwV1ApiCreateTransitGatewayModelTgwResponseModel) SetRegion(v Region)`

SetRegion sets Region field to given value.


### GetOptions

`func (o *BnsTgwV1ApiCreateTransitGatewayModelTgwResponseModel) GetOptions() TgwOptionsResponseModel`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *BnsTgwV1ApiCreateTransitGatewayModelTgwResponseModel) GetOptionsOk() (*TgwOptionsResponseModel, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *BnsTgwV1ApiCreateTransitGatewayModelTgwResponseModel) SetOptions(v TgwOptionsResponseModel)`

SetOptions sets Options field to given value.


### GetProvisioningStatus

`func (o *BnsTgwV1ApiCreateTransitGatewayModelTgwResponseModel) GetProvisioningStatus() TGWProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *BnsTgwV1ApiCreateTransitGatewayModelTgwResponseModel) GetProvisioningStatusOk() (*TGWProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *BnsTgwV1ApiCreateTransitGatewayModelTgwResponseModel) SetProvisioningStatus(v TGWProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.


### GetProjectId

`func (o *BnsTgwV1ApiCreateTransitGatewayModelTgwResponseModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BnsTgwV1ApiCreateTransitGatewayModelTgwResponseModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BnsTgwV1ApiCreateTransitGatewayModelTgwResponseModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetCreatedAt

`func (o *BnsTgwV1ApiCreateTransitGatewayModelTgwResponseModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BnsTgwV1ApiCreateTransitGatewayModelTgwResponseModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BnsTgwV1ApiCreateTransitGatewayModelTgwResponseModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *BnsTgwV1ApiCreateTransitGatewayModelTgwResponseModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BnsTgwV1ApiCreateTransitGatewayModelTgwResponseModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BnsTgwV1ApiCreateTransitGatewayModelTgwResponseModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


