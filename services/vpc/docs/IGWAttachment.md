# IGWAttachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | IGW Attachment의 고유 ID | 
**VpcId** | **string** | 인터넷 게이트웨이에 연결된 VPC의 고유 ID | 
**ProvisioningStatus** | **string** | IGW Attachment의 프로비저닝 상태 | 
**ProjectId** | **string** | IGW Attachment가 속한 프로젝트 ID | 

## Methods

### NewIGWAttachment

`func NewIGWAttachment(id string, vpcId string, provisioningStatus string, projectId string, ) *IGWAttachment`

NewIGWAttachment instantiates a new IGWAttachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIGWAttachmentWithDefaults

`func NewIGWAttachmentWithDefaults() *IGWAttachment`

NewIGWAttachmentWithDefaults instantiates a new IGWAttachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IGWAttachment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IGWAttachment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IGWAttachment) SetId(v string)`

SetId sets Id field to given value.


### GetVpcId

`func (o *IGWAttachment) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *IGWAttachment) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *IGWAttachment) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.


### GetProvisioningStatus

`func (o *IGWAttachment) GetProvisioningStatus() string`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *IGWAttachment) GetProvisioningStatusOk() (*string, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *IGWAttachment) SetProvisioningStatus(v string)`

SetProvisioningStatus sets ProvisioningStatus field to given value.


### GetProjectId

`func (o *IGWAttachment) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *IGWAttachment) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *IGWAttachment) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


