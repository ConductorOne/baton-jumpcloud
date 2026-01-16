# CommandresultslistResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the command result. | [optional] 
**Command** | Pointer to **string** | The command that was executed on the system. | [optional] 
**ExitCode** | Pointer to **int32** | The stderr output from the command that ran. | [optional] 
**Name** | Pointer to **string** | The name of the command. | [optional] 
**RequestTime** | Pointer to **time.Time** | The time (UTC) that the command was sent. | [optional] 
**ResponseTime** | Pointer to **time.Time** | The time (UTC) that the command was completed. | [optional] 
**Sudo** | Pointer to **bool** | If the user had sudo rights. | [optional] 
**System** | Pointer to **string** | The display name of the system the command was executed on. | [optional] 
**SystemId** | Pointer to **string** | The id of the system the command was executed on. | [optional] 
**User** | Pointer to **string** | The user the command ran as. | [optional] 
**WorkflowId** | Pointer to **string** | The id for the command that ran on the system. | [optional] 

## Methods

### NewCommandresultslistResultsInner

`func NewCommandresultslistResultsInner() *CommandresultslistResultsInner`

NewCommandresultslistResultsInner instantiates a new CommandresultslistResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommandresultslistResultsInnerWithDefaults

`func NewCommandresultslistResultsInnerWithDefaults() *CommandresultslistResultsInner`

NewCommandresultslistResultsInnerWithDefaults instantiates a new CommandresultslistResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CommandresultslistResultsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommandresultslistResultsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommandresultslistResultsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CommandresultslistResultsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCommand

`func (o *CommandresultslistResultsInner) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *CommandresultslistResultsInner) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *CommandresultslistResultsInner) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *CommandresultslistResultsInner) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetExitCode

`func (o *CommandresultslistResultsInner) GetExitCode() int32`

GetExitCode returns the ExitCode field if non-nil, zero value otherwise.

### GetExitCodeOk

`func (o *CommandresultslistResultsInner) GetExitCodeOk() (*int32, bool)`

GetExitCodeOk returns a tuple with the ExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCode

`func (o *CommandresultslistResultsInner) SetExitCode(v int32)`

SetExitCode sets ExitCode field to given value.

### HasExitCode

`func (o *CommandresultslistResultsInner) HasExitCode() bool`

HasExitCode returns a boolean if a field has been set.

### GetName

`func (o *CommandresultslistResultsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CommandresultslistResultsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CommandresultslistResultsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CommandresultslistResultsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRequestTime

`func (o *CommandresultslistResultsInner) GetRequestTime() time.Time`

GetRequestTime returns the RequestTime field if non-nil, zero value otherwise.

### GetRequestTimeOk

`func (o *CommandresultslistResultsInner) GetRequestTimeOk() (*time.Time, bool)`

GetRequestTimeOk returns a tuple with the RequestTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTime

`func (o *CommandresultslistResultsInner) SetRequestTime(v time.Time)`

SetRequestTime sets RequestTime field to given value.

### HasRequestTime

`func (o *CommandresultslistResultsInner) HasRequestTime() bool`

HasRequestTime returns a boolean if a field has been set.

### GetResponseTime

`func (o *CommandresultslistResultsInner) GetResponseTime() time.Time`

GetResponseTime returns the ResponseTime field if non-nil, zero value otherwise.

### GetResponseTimeOk

`func (o *CommandresultslistResultsInner) GetResponseTimeOk() (*time.Time, bool)`

GetResponseTimeOk returns a tuple with the ResponseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTime

`func (o *CommandresultslistResultsInner) SetResponseTime(v time.Time)`

SetResponseTime sets ResponseTime field to given value.

### HasResponseTime

`func (o *CommandresultslistResultsInner) HasResponseTime() bool`

HasResponseTime returns a boolean if a field has been set.

### GetSudo

`func (o *CommandresultslistResultsInner) GetSudo() bool`

GetSudo returns the Sudo field if non-nil, zero value otherwise.

### GetSudoOk

`func (o *CommandresultslistResultsInner) GetSudoOk() (*bool, bool)`

GetSudoOk returns a tuple with the Sudo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSudo

`func (o *CommandresultslistResultsInner) SetSudo(v bool)`

SetSudo sets Sudo field to given value.

### HasSudo

`func (o *CommandresultslistResultsInner) HasSudo() bool`

HasSudo returns a boolean if a field has been set.

### GetSystem

`func (o *CommandresultslistResultsInner) GetSystem() string`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *CommandresultslistResultsInner) GetSystemOk() (*string, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *CommandresultslistResultsInner) SetSystem(v string)`

SetSystem sets System field to given value.

### HasSystem

`func (o *CommandresultslistResultsInner) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetSystemId

`func (o *CommandresultslistResultsInner) GetSystemId() string`

GetSystemId returns the SystemId field if non-nil, zero value otherwise.

### GetSystemIdOk

`func (o *CommandresultslistResultsInner) GetSystemIdOk() (*string, bool)`

GetSystemIdOk returns a tuple with the SystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemId

`func (o *CommandresultslistResultsInner) SetSystemId(v string)`

SetSystemId sets SystemId field to given value.

### HasSystemId

`func (o *CommandresultslistResultsInner) HasSystemId() bool`

HasSystemId returns a boolean if a field has been set.

### GetUser

`func (o *CommandresultslistResultsInner) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *CommandresultslistResultsInner) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *CommandresultslistResultsInner) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *CommandresultslistResultsInner) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetWorkflowId

`func (o *CommandresultslistResultsInner) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *CommandresultslistResultsInner) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *CommandresultslistResultsInner) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *CommandresultslistResultsInner) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


