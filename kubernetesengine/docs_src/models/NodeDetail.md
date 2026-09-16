# NodeDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | API 버전 | 
**Events** | Pointer to [**[]Event**](Event.md) | 노드 관련 이벤트 목록 (예: 스케줄링, 상태 변경 등) | [optional] 
**Kind** | **string** | 리소스 종류 | 
**Metadata** | Pointer to [**NullableMetadata**](Metadata.md) | 노드의 메타데이터 | [optional] 
**Pods** | Pointer to [**[]Pod**](Pod.md) | 노드에 배치된 파드 목록 | [optional] 
**Spec** | Pointer to [**NullableSpec**](Spec.md) | 노드 spec | [optional] 
**Status** | Pointer to [**NullableGetClusterNodeDetailsStatus**](GetClusterNodeDetailsStatus.md) | 노드 status   &lt;!---표현 의도적으로 영문 사용---&gt; | [optional] 

## Methods

### NewNodeDetail

`func NewNodeDetail(apiVersion string, kind string, ) *NodeDetail`

NewNodeDetail instantiates a new NodeDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeDetailWithDefaults

`func NewNodeDetailWithDefaults() *NodeDetail`

NewNodeDetailWithDefaults instantiates a new NodeDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *NodeDetail) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NodeDetail) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NodeDetail) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetEvents

`func (o *NodeDetail) GetEvents() []Event`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *NodeDetail) GetEventsOk() (*[]Event, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *NodeDetail) SetEvents(v []Event)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *NodeDetail) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### SetEventsNil

`func (o *NodeDetail) SetEventsNil(b bool)`

 SetEventsNil sets the value for Events to be an explicit nil

### UnsetEvents
`func (o *NodeDetail) UnsetEvents()`

UnsetEvents ensures that no value is present for Events, not even an explicit nil
### GetKind

`func (o *NodeDetail) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *NodeDetail) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *NodeDetail) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMetadata

`func (o *NodeDetail) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NodeDetail) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NodeDetail) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *NodeDetail) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *NodeDetail) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *NodeDetail) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetPods

`func (o *NodeDetail) GetPods() []Pod`

GetPods returns the Pods field if non-nil, zero value otherwise.

### GetPodsOk

`func (o *NodeDetail) GetPodsOk() (*[]Pod, bool)`

GetPodsOk returns a tuple with the Pods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPods

`func (o *NodeDetail) SetPods(v []Pod)`

SetPods sets Pods field to given value.

### HasPods

`func (o *NodeDetail) HasPods() bool`

HasPods returns a boolean if a field has been set.

### SetPodsNil

`func (o *NodeDetail) SetPodsNil(b bool)`

 SetPodsNil sets the value for Pods to be an explicit nil

### UnsetPods
`func (o *NodeDetail) UnsetPods()`

UnsetPods ensures that no value is present for Pods, not even an explicit nil
### GetSpec

`func (o *NodeDetail) GetSpec() Spec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *NodeDetail) GetSpecOk() (*Spec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *NodeDetail) SetSpec(v Spec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *NodeDetail) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### SetSpecNil

`func (o *NodeDetail) SetSpecNil(b bool)`

 SetSpecNil sets the value for Spec to be an explicit nil

### UnsetSpec
`func (o *NodeDetail) UnsetSpec()`

UnsetSpec ensures that no value is present for Spec, not even an explicit nil
### GetStatus

`func (o *NodeDetail) GetStatus() GetClusterNodeDetailsStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NodeDetail) GetStatusOk() (*GetClusterNodeDetailsStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NodeDetail) SetStatus(v GetClusterNodeDetailsStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NodeDetail) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *NodeDetail) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *NodeDetail) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


