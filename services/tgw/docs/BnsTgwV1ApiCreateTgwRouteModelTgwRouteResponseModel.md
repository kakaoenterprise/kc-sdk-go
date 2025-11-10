# BnsTgwV1ApiCreateTgwRouteModelTgwRouteResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Route ID | 
**DestinationCidrBlock** | **string** | 대상 네트워크 CIDR 블록 | 
**TgwRouteTableId** | **string** | Route가 속한 Transit Gateway 라우팅 테이블 ID | 
**TgwRouteAttachments** | [**[]BnsTgwV1ApiCreateTgwRouteModelTgwRouteAttachmentResponseModel**](BnsTgwV1ApiCreateTgwRouteModelTgwRouteAttachmentResponseModel.md) | Route에 연결된 Transit Gateway Attachment 목록 | 
**ProvisioningStatus** | [**TGWProvisioningStatus**](TGWProvisioningStatus.md) | Route의 프로비저닝 상태 | 
**ProjectId** | **string** | 프로젝트 ID | 

## Methods

### NewBnsTgwV1ApiCreateTgwRouteModelTgwRouteResponseModel

`func NewBnsTgwV1ApiCreateTgwRouteModelTgwRouteResponseModel(id string, destinationCidrBlock string, tgwRouteTableId string, tgwRouteAttachments []BnsTgwV1ApiCreateTgwRouteModelTgwRouteAttachmentResponseModel, provisioningStatus TGWProvisioningStatus, projectId string, ) *BnsTgwV1ApiCreateTgwRouteModelTgwRouteResponseModel`

NewBnsTgwV1ApiCreateTgwRouteModelTgwRouteResponseModel instantiates a new BnsTgwV1ApiCreateTgwRouteModelTgwRouteResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsTgwV1ApiCreateTgwRouteModelTgwRouteResponseModelWithDefaults

`func NewBnsTgwV1ApiCreateTgwRouteModelTgwRouteResponseModelWithDefaults() *BnsTgwV1ApiCreateTgwRouteModelTgwRouteResponseModel`

NewBnsTgwV1ApiCreateTgwRouteModelTgwRouteResponseModelWithDefaults instantiates a new BnsTgwV1ApiCreateTgwRouteModelTgwRouteResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsTgwV1ApiCreateTgwRouteModelTgwRouteResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsTgwV1ApiCreateTgwRouteModelTgwRouteResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsTgwV1ApiCreateTgwRouteModelTgwRouteResponseModel) SetId(v string)`

SetId sets Id field to given value.


### GetDestinationCidrBlock

`func (o *BnsTgwV1ApiCreateTgwRouteModelTgwRouteResponseModel) GetDestinationCidrBlock() string`

GetDestinationCidrBlock returns the DestinationCidrBlock field if non-nil, zero value otherwise.

### GetDestinationCidrBlockOk

`func (o *BnsTgwV1ApiCreateTgwRouteModelTgwRouteResponseModel) GetDestinationCidrBlockOk() (*string, bool)`

GetDestinationCidrBlockOk returns a tuple with the DestinationCidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationCidrBlock

`func (o *BnsTgwV1ApiCreateTgwRouteModelTgwRouteResponseModel) SetDestinationCidrBlock(v string)`

SetDestinationCidrBlock sets DestinationCidrBlock field to given value.


### GetTgwRouteTableId

`func (o *BnsTgwV1ApiCreateTgwRouteModelTgwRouteResponseModel) GetTgwRouteTableId() string`

GetTgwRouteTableId returns the TgwRouteTableId field if non-nil, zero value otherwise.

### GetTgwRouteTableIdOk

`func (o *BnsTgwV1ApiCreateTgwRouteModelTgwRouteResponseModel) GetTgwRouteTableIdOk() (*string, bool)`

GetTgwRouteTableIdOk returns a tuple with the TgwRouteTableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgwRouteTableId

`func (o *BnsTgwV1ApiCreateTgwRouteModelTgwRouteResponseModel) SetTgwRouteTableId(v string)`

SetTgwRouteTableId sets TgwRouteTableId field to given value.


### GetTgwRouteAttachments

`func (o *BnsTgwV1ApiCreateTgwRouteModelTgwRouteResponseModel) GetTgwRouteAttachments() []BnsTgwV1ApiCreateTgwRouteModelTgwRouteAttachmentResponseModel`

GetTgwRouteAttachments returns the TgwRouteAttachments field if non-nil, zero value otherwise.

### GetTgwRouteAttachmentsOk

`func (o *BnsTgwV1ApiCreateTgwRouteModelTgwRouteResponseModel) GetTgwRouteAttachmentsOk() (*[]BnsTgwV1ApiCreateTgwRouteModelTgwRouteAttachmentResponseModel, bool)`

GetTgwRouteAttachmentsOk returns a tuple with the TgwRouteAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgwRouteAttachments

`func (o *BnsTgwV1ApiCreateTgwRouteModelTgwRouteResponseModel) SetTgwRouteAttachments(v []BnsTgwV1ApiCreateTgwRouteModelTgwRouteAttachmentResponseModel)`

SetTgwRouteAttachments sets TgwRouteAttachments field to given value.


### GetProvisioningStatus

`func (o *BnsTgwV1ApiCreateTgwRouteModelTgwRouteResponseModel) GetProvisioningStatus() TGWProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *BnsTgwV1ApiCreateTgwRouteModelTgwRouteResponseModel) GetProvisioningStatusOk() (*TGWProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *BnsTgwV1ApiCreateTgwRouteModelTgwRouteResponseModel) SetProvisioningStatus(v TGWProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.


### GetProjectId

`func (o *BnsTgwV1ApiCreateTgwRouteModelTgwRouteResponseModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BnsTgwV1ApiCreateTgwRouteModelTgwRouteResponseModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BnsTgwV1ApiCreateTgwRouteModelTgwRouteResponseModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


