# BnsTgwV1ApiCreateTgwRouteModelTgwRouteAttachmentResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceId** | **string** | Route가 연결된 리소스 ID | 
**ResourceType** | **string** | 리소스 유형 | 
**TgwAttachmentId** | **string** | Transit Gateway와 연결된 Attachment의 연결 ID&lt;br/&gt;- [List TGW attachments](/openapi/bns/tgw/list-tgw-attachments) API에서 조회한 &#x60;attachments.tgw.attachment_id&#x60; 확인 | 

## Methods

### NewBnsTgwV1ApiCreateTgwRouteModelTgwRouteAttachmentResponseModel

`func NewBnsTgwV1ApiCreateTgwRouteModelTgwRouteAttachmentResponseModel(resourceId string, resourceType string, tgwAttachmentId string, ) *BnsTgwV1ApiCreateTgwRouteModelTgwRouteAttachmentResponseModel`

NewBnsTgwV1ApiCreateTgwRouteModelTgwRouteAttachmentResponseModel instantiates a new BnsTgwV1ApiCreateTgwRouteModelTgwRouteAttachmentResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsTgwV1ApiCreateTgwRouteModelTgwRouteAttachmentResponseModelWithDefaults

`func NewBnsTgwV1ApiCreateTgwRouteModelTgwRouteAttachmentResponseModelWithDefaults() *BnsTgwV1ApiCreateTgwRouteModelTgwRouteAttachmentResponseModel`

NewBnsTgwV1ApiCreateTgwRouteModelTgwRouteAttachmentResponseModelWithDefaults instantiates a new BnsTgwV1ApiCreateTgwRouteModelTgwRouteAttachmentResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceId

`func (o *BnsTgwV1ApiCreateTgwRouteModelTgwRouteAttachmentResponseModel) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *BnsTgwV1ApiCreateTgwRouteModelTgwRouteAttachmentResponseModel) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *BnsTgwV1ApiCreateTgwRouteModelTgwRouteAttachmentResponseModel) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetResourceType

`func (o *BnsTgwV1ApiCreateTgwRouteModelTgwRouteAttachmentResponseModel) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *BnsTgwV1ApiCreateTgwRouteModelTgwRouteAttachmentResponseModel) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *BnsTgwV1ApiCreateTgwRouteModelTgwRouteAttachmentResponseModel) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetTgwAttachmentId

`func (o *BnsTgwV1ApiCreateTgwRouteModelTgwRouteAttachmentResponseModel) GetTgwAttachmentId() string`

GetTgwAttachmentId returns the TgwAttachmentId field if non-nil, zero value otherwise.

### GetTgwAttachmentIdOk

`func (o *BnsTgwV1ApiCreateTgwRouteModelTgwRouteAttachmentResponseModel) GetTgwAttachmentIdOk() (*string, bool)`

GetTgwAttachmentIdOk returns a tuple with the TgwAttachmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgwAttachmentId

`func (o *BnsTgwV1ApiCreateTgwRouteModelTgwRouteAttachmentResponseModel) SetTgwAttachmentId(v string)`

SetTgwAttachmentId sets TgwAttachmentId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


