# NodeDetailResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | API 버전 | 
**Events** | Pointer to [**[]EventResponseModel**](EventResponseModel.md) |  | [optional] 
**Kind** | **string** | 리소스 종류 | 
**Metadata** | Pointer to [**NullableMetadataResponseModel**](MetadataResponseModel.md) |  | [optional] 
**Pods** | Pointer to [**[]PodResponseModel**](PodResponseModel.md) |  | [optional] 
**Spec** | Pointer to [**NullableSpecResponseModel**](SpecResponseModel.md) |  | [optional] 
**Status** | Pointer to [**NullableKubernetesEngineV1ApiGetClusterNodeDetailsModelStatusResponseModel**](KubernetesEngineV1ApiGetClusterNodeDetailsModelStatusResponseModel.md) |  | [optional] 

## Methods

### NewNodeDetailResponseModel

`func NewNodeDetailResponseModel(apiVersion string, kind string, ) *NodeDetailResponseModel`

NewNodeDetailResponseModel instantiates a new NodeDetailResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeDetailResponseModelWithDefaults

`func NewNodeDetailResponseModelWithDefaults() *NodeDetailResponseModel`

NewNodeDetailResponseModelWithDefaults instantiates a new NodeDetailResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *NodeDetailResponseModel) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NodeDetailResponseModel) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NodeDetailResponseModel) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetEvents

`func (o *NodeDetailResponseModel) GetEvents() []EventResponseModel`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *NodeDetailResponseModel) GetEventsOk() (*[]EventResponseModel, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *NodeDetailResponseModel) SetEvents(v []EventResponseModel)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *NodeDetailResponseModel) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### SetEventsNil

`func (o *NodeDetailResponseModel) SetEventsNil(b bool)`

 SetEventsNil sets the value for Events to be an explicit nil

### UnsetEvents
`func (o *NodeDetailResponseModel) UnsetEvents()`

UnsetEvents ensures that no value is present for Events, not even an explicit nil
### GetKind

`func (o *NodeDetailResponseModel) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NodeDetailResponseModel) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NodeDetailResponseModel) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *NodeDetailResponseModel) GetMetadata() MetadataResponseModel`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NodeDetailResponseModel) GetMetadataOk() (*MetadataResponseModel, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NodeDetailResponseModel) SetMetadata(v MetadataResponseModel)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *NodeDetailResponseModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *NodeDetailResponseModel) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *NodeDetailResponseModel) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetPods

`func (o *NodeDetailResponseModel) GetPods() []PodResponseModel`

GetPods returns the Pods field if non-nil, zero value otherwise.

### GetPodsOk

`func (o *NodeDetailResponseModel) GetPodsOk() (*[]PodResponseModel, bool)`

GetPodsOk returns a tuple with the Pods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPods

`func (o *NodeDetailResponseModel) SetPods(v []PodResponseModel)`

SetPods sets Pods field to given value.

### HasPods

`func (o *NodeDetailResponseModel) HasPods() bool`

HasPods returns a boolean if a field has been set.

### SetPodsNil

`func (o *NodeDetailResponseModel) SetPodsNil(b bool)`

 SetPodsNil sets the value for Pods to be an explicit nil

### UnsetPods
`func (o *NodeDetailResponseModel) UnsetPods()`

UnsetPods ensures that no value is present for Pods, not even an explicit nil
### GetSpec

`func (o *NodeDetailResponseModel) GetSpec() SpecResponseModel`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *NodeDetailResponseModel) GetSpecOk() (*SpecResponseModel, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *NodeDetailResponseModel) SetSpec(v SpecResponseModel)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *NodeDetailResponseModel) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### SetSpecNil

`func (o *NodeDetailResponseModel) SetSpecNil(b bool)`

 SetSpecNil sets the value for Spec to be an explicit nil

### UnsetSpec
`func (o *NodeDetailResponseModel) UnsetSpec()`

UnsetSpec ensures that no value is present for Spec, not even an explicit nil
### GetStatus

`func (o *NodeDetailResponseModel) GetStatus() KubernetesEngineV1ApiGetClusterNodeDetailsModelStatusResponseModel`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NodeDetailResponseModel) GetStatusOk() (*KubernetesEngineV1ApiGetClusterNodeDetailsModelStatusResponseModel, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NodeDetailResponseModel) SetStatus(v KubernetesEngineV1ApiGetClusterNodeDetailsModelStatusResponseModel)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NodeDetailResponseModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *NodeDetailResponseModel) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *NodeDetailResponseModel) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


