# QueuedCommandListResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | Pointer to **string** | The ID of the command, from savedAgentCommands. | [optional] 
**Id** | Pointer to **string** | The workflowInstanceId. | [optional] 
**PendingCount** | Pointer to **int32** | The number of devices that still haven&#39;t received the directive. | [optional] 
**System** | Pointer to **string** | The ID of the device the command is bound to. | [optional] 

## Methods

### NewQueuedCommandListResultsInner

`func NewQueuedCommandListResultsInner() *QueuedCommandListResultsInner`

NewQueuedCommandListResultsInner instantiates a new QueuedCommandListResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueuedCommandListResultsInnerWithDefaults

`func NewQueuedCommandListResultsInnerWithDefaults() *QueuedCommandListResultsInner`

NewQueuedCommandListResultsInnerWithDefaults instantiates a new QueuedCommandListResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *QueuedCommandListResultsInner) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *QueuedCommandListResultsInner) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *QueuedCommandListResultsInner) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *QueuedCommandListResultsInner) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetId

`func (o *QueuedCommandListResultsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QueuedCommandListResultsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QueuedCommandListResultsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *QueuedCommandListResultsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPendingCount

`func (o *QueuedCommandListResultsInner) GetPendingCount() int32`

GetPendingCount returns the PendingCount field if non-nil, zero value otherwise.

### GetPendingCountOk

`func (o *QueuedCommandListResultsInner) GetPendingCountOk() (*int32, bool)`

GetPendingCountOk returns a tuple with the PendingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingCount

`func (o *QueuedCommandListResultsInner) SetPendingCount(v int32)`

SetPendingCount sets PendingCount field to given value.

### HasPendingCount

`func (o *QueuedCommandListResultsInner) HasPendingCount() bool`

HasPendingCount returns a boolean if a field has been set.

### GetSystem

`func (o *QueuedCommandListResultsInner) GetSystem() string`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *QueuedCommandListResultsInner) GetSystemOk() (*string, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *QueuedCommandListResultsInner) SetSystem(v string)`

SetSystem sets System field to given value.

### HasSystem

`func (o *QueuedCommandListResultsInner) HasSystem() bool`

HasSystem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


