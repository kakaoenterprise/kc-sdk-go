# BnsTgwV1ApiUpdateTgwRouteModelTgwRouteResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Route ID | 
**DestinationCidrBlock** | **string** | 목적지 CIDR 블록 (예: 10.0.0.0/16) | 
**TgwRouteTableId** | **string** | Route가 속한 Transit Gateway 라우팅 테이블 ID | 
**TgwRouteAttachments** | [**[]BnsTgwV1ApiUpdateTgwRouteModelTgwRouteAttachmentResponseModel**](BnsTgwV1ApiUpdateTgwRouteModelTgwRouteAttachmentResponseModel.md) | Route에 연결된 Transit Gateway Attachment 목록 | 
**ProvisioningStatus** | [**TGWProvisioningStatus**](TGWProvisioningStatus.md) | 프로비저닝 상태 | 
**ProjectId** | **string** | Route가 속한 프로젝트 ID | 

## Methods

### NewBnsTgwV1ApiUpdateTgwRouteModelTgwRouteResponseModel

`func NewBnsTgwV1ApiUpdateTgwRouteModelTgwRouteResponseModel(id string, destinationCidrBlock string, tgwRouteTableId string, tgwRouteAttachments []BnsTgwV1ApiUpdateTgwRouteModelTgwRouteAttachmentResponseModel, provisioningStatus TGWProvisioningStatus, projectId string, ) *BnsTgwV1ApiUpdateTgwRouteModelTgwRouteResponseModel`

NewBnsTgwV1ApiUpdateTgwRouteModelTgwRouteResponseModel instantiates a new BnsTgwV1ApiUpdateTgwRouteModelTgwRouteResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsTgwV1ApiUpdateTgwRouteModelTgwRouteResponseModelWithDefaults

`func NewBnsTgwV1ApiUpdateTgwRouteModelTgwRouteResponseModelWithDefaults() *BnsTgwV1ApiUpdateTgwRouteModelTgwRouteResponseModel`

NewBnsTgwV1ApiUpdateTgwRouteModelTgwRouteResponseModelWithDefaults instantiates a new BnsTgwV1ApiUpdateTgwRouteModelTgwRouteResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsTgwV1ApiUpdateTgwRouteModelTgwRouteResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsTgwV1ApiUpdateTgwRouteModelTgwRouteResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsTgwV1ApiUpdateTgwRouteModelTgwRouteResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetDestinationCidrBlock

`func (o *BnsTgwV1ApiUpdateTgwRouteModelTgwRouteResponseModel) GetDestinationCidrBlock() string`

GetDestinationCidrBlock returns the DestinationCidrBlock field if non-nil, zero value otherwise.

### GetDestinationCidrBlockOk

`func (o *BnsTgwV1ApiUpdateTgwRouteModelTgwRouteResponseModel) GetDestinationCidrBlockOk() (*string, bool)`

GetDestinationCidrBlockOk returns a tuple with the DestinationCidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationCidrBlock

`func (o *BnsTgwV1ApiUpdateTgwRouteModelTgwRouteResponseModel) SetDestinationCidrBlock(v string)`

SetDestinationCidrBlock sets DestinationCidrBlock field to given value.


### GetTgwRouteTableId

`func (o *BnsTgwV1ApiUpdateTgwRouteModelTgwRouteResponseModel) GetTgwRouteTableId() string`

GetTgwRouteTableId returns the TgwRouteTableId field if non-nil, zero value otherwise.

### GetTgwRouteTableIdOk

`func (o *BnsTgwV1ApiUpdateTgwRouteModelTgwRouteResponseModel) GetTgwRouteTableIdOk() (*string, bool)`

GetTgwRouteTableIdOk returns a tuple with the TgwRouteTableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgwRouteTableId

`func (o *BnsTgwV1ApiUpdateTgwRouteModelTgwRouteResponseModel) SetTgwRouteTableId(v string)`

SetTgwRouteTableId sets TgwRouteTableId field to given value.


### GetTgwRouteAttachments

`func (o *BnsTgwV1ApiUpdateTgwRouteModelTgwRouteResponseModel) GetTgwRouteAttachments() []BnsTgwV1ApiUpdateTgwRouteModelTgwRouteAttachmentResponseModel`

GetTgwRouteAttachments returns the TgwRouteAttachments field if non-nil, zero value otherwise.

### GetTgwRouteAttachmentsOk

`func (o *BnsTgwV1ApiUpdateTgwRouteModelTgwRouteResponseModel) GetTgwRouteAttachmentsOk() (*[]BnsTgwV1ApiUpdateTgwRouteModelTgwRouteAttachmentResponseModel, bool)`

GetTgwRouteAttachmentsOk returns a tuple with the TgwRouteAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgwRouteAttachments

`func (o *BnsTgwV1ApiUpdateTgwRouteModelTgwRouteResponseModel) SetTgwRouteAttachments(v []BnsTgwV1ApiUpdateTgwRouteModelTgwRouteAttachmentResponseModel)`

SetTgwRouteAttachments sets TgwRouteAttachments field to given value.


### GetProvisioningStatus

`func (o *BnsTgwV1ApiUpdateTgwRouteModelTgwRouteResponseModel) GetProvisioningStatus() TGWProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *BnsTgwV1ApiUpdateTgwRouteModelTgwRouteResponseModel) GetProvisioningStatusOk() (*TGWProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *BnsTgwV1ApiUpdateTgwRouteModelTgwRouteResponseModel) SetProvisioningStatus(v TGWProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.


### GetProjectId

`func (o *BnsTgwV1ApiUpdateTgwRouteModelTgwRouteResponseModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BnsTgwV1ApiUpdateTgwRouteModelTgwRouteResponseModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BnsTgwV1ApiUpdateTgwRouteModelTgwRouteResponseModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


