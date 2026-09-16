# GetServiceAgentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceAgent** | [**ServiceAgent**](ServiceAgent.md) | 서비스 에이전트 정보 | 

## Methods

### NewGetServiceAgentResponse

`func NewGetServiceAgentResponse(serviceAgent ServiceAgent, ) *GetServiceAgentResponse`

NewGetServiceAgentResponse instantiates a new GetServiceAgentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetServiceAgentResponseWithDefaults

`func NewGetServiceAgentResponseWithDefaults() *GetServiceAgentResponse`

NewGetServiceAgentResponseWithDefaults instantiates a new GetServiceAgentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceAgent

`func (o *GetServiceAgentResponse) GetServiceAgent() ServiceAgent`

GetServiceAgent returns the ServiceAgent field if non-nil, zero value otherwise.

### GetServiceAgentOk

`func (o *GetServiceAgentResponse) GetServiceAgentOk() (*ServiceAgent, bool)`

GetServiceAgentOk returns a tuple with the ServiceAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAgent

`func (o *GetServiceAgentResponse) SetServiceAgent(v ServiceAgent)`

SetServiceAgent sets ServiceAgent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


