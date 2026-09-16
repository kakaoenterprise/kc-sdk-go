# TgwRouteAttachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceId** | **string** | 연결된 리소스 ID | 
**ResourceType** | [**ResourceType**](ResourceType.md) | 연결된 리소스 유형 | 
**TgwAttachmentId** | **string** | Transit Gateway와 연결된 Attachment의 연결 ID | 

## Methods

### NewTgwRouteAttachment

`func NewTgwRouteAttachment(resourceId string, resourceType ResourceType, tgwAttachmentId string, ) *TgwRouteAttachment`

NewTgwRouteAttachment instantiates a new TgwRouteAttachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTgwRouteAttachmentWithDefaults

`func NewTgwRouteAttachmentWithDefaults() *TgwRouteAttachment`

NewTgwRouteAttachmentWithDefaults instantiates a new TgwRouteAttachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceId

`func (o *TgwRouteAttachment) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *TgwRouteAttachment) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *TgwRouteAttachment) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetResourceType

`func (o *TgwRouteAttachment) GetResourceType() ResourceType`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *TgwRouteAttachment) GetResourceTypeOk() (*ResourceType, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *TgwRouteAttachment) SetResourceType(v ResourceType)`

SetResourceType sets ResourceType field to given value.


### GetTgwAttachmentId

`func (o *TgwRouteAttachment) GetTgwAttachmentId() string`

GetTgwAttachmentId returns the TgwAttachmentId field if non-nil, zero value otherwise.

### GetTgwAttachmentIdOk

`func (o *TgwRouteAttachment) GetTgwAttachmentIdOk() (*string, bool)`

GetTgwAttachmentIdOk returns a tuple with the TgwAttachmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTgwAttachmentId

`func (o *TgwRouteAttachment) SetTgwAttachmentId(v string)`

SetTgwAttachmentId sets TgwAttachmentId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


