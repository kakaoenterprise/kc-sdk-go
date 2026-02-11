# TgwRouteTableAssociationResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Association ID | 
**ResourceId** | **string** | 연결된 리소스의 ID | 
**ResourceType** | [**ResourceType**](ResourceType.md) | 연결된 리소스 유형 | 
**TgwRouteTableId** | **string** | 연결된 TGW 라우팅 테이블 ID | 
**TgwAttachmentId** | **string** | Transit Gateway와 연결된 Attachment의 연결 ID | 
**ProvisioningStatus** | [**TGWProvisioningStatus**](TGWProvisioningStatus.md) | Association 프로비저닝 상태 | 
**ProjectId** | **string** | 프로젝트 ID | 

## Methods

### NewTgwRouteTableAssociationResponseModel

`func NewTgwRouteTableAssociationResponseModel(id string, resourceId string, resourceType ResourceType, tgwRouteTableId string, tgwAttachmentId string, provisioningStatus TGWProvisioningStatus, projectId string, ) *TgwRouteTableAssociationResponseModel`

NewTgwRouteTableAssociationResponseModel instantiates a new TgwRouteTableAssociationResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTgwRouteTableAssociationResponseModelWithDefaults

`func NewTgwRouteTableAssociationResponseModelWithDefaults() *TgwRouteTableAssociationResponseModel`

NewTgwRouteTableAssociationResponseModelWithDefaults instantiates a new TgwRouteTableAssociationResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TgwRouteTableAssociationResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TgwRouteTableAssociationResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TgwRouteTableAssociationResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetResourceId

`func (o *TgwRouteTableAssociationResponseModel) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *TgwRouteTableAssociationResponseModel) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *TgwRouteTableAssociationResponseModel) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetResourceType

`func (o *TgwRouteTableAssociationResponseModel) GetResourceType() ResourceType`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *TgwRouteTableAssociationResponseModel) GetResourceTypeOk() (*ResourceType, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *TgwRouteTableAssociationResponseModel) SetResourceType(v ResourceType)`

SetResourceType sets ResourceType field to given value.


### GetTgwRouteTableId

`func (o *TgwRouteTableAssociationResponseModel) GetTgwRouteTableId() string`

GetTgwRouteTableId returns the TgwRouteTableId field if non-nil, zero value otherwise.

### GetTgwRouteTableIdOk

`func (o *TgwRouteTableAssociationResponseModel) GetTgwRouteTableIdOk() (*string, bool)`

GetTgwRouteTableIdOk returns a tuple with the TgwRouteTableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgwRouteTableId

`func (o *TgwRouteTableAssociationResponseModel) SetTgwRouteTableId(v string)`

SetTgwRouteTableId sets TgwRouteTableId field to given value.


### GetTgwAttachmentId

`func (o *TgwRouteTableAssociationResponseModel) GetTgwAttachmentId() string`

GetTgwAttachmentId returns the TgwAttachmentId field if non-nil, zero value otherwise.

### GetTgwAttachmentIdOk

`func (o *TgwRouteTableAssociationResponseModel) GetTgwAttachmentIdOk() (*string, bool)`

GetTgwAttachmentIdOk returns a tuple with the TgwAttachmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgwAttachmentId

`func (o *TgwRouteTableAssociationResponseModel) SetTgwAttachmentId(v string)`

SetTgwAttachmentId sets TgwAttachmentId field to given value.


### GetProvisioningStatus

`func (o *TgwRouteTableAssociationResponseModel) GetProvisioningStatus() TGWProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *TgwRouteTableAssociationResponseModel) GetProvisioningStatusOk() (*TGWProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *TgwRouteTableAssociationResponseModel) SetProvisioningStatus(v TGWProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.


### GetProjectId

`func (o *TgwRouteTableAssociationResponseModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *TgwRouteTableAssociationResponseModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *TgwRouteTableAssociationResponseModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


