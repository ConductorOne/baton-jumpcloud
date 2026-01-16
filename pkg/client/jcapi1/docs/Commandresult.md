# Commandresult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the command. | [optional] 
**Command** | Pointer to **string** | The command that was executed on the system. | [optional] 
**Files** | Pointer to **[]string** | An array of file ids that were included in the command | [optional] 
**Name** | Pointer to **string** | The name of the command. | [optional] 
**Organization** | Pointer to **string** | The ID of the organization. | [optional] 
**RequestTime** | Pointer to **time.Time** | The time that the command was sent. | [optional] 
**Response** | Pointer to [**CommandresultResponse**](CommandresultResponse.md) |  | [optional] 
**ResponseTime** | Pointer to **time.Time** | The time that the command was completed. | [optional] 
**Sudo** | Pointer to **bool** | If the user had sudo rights | [optional] 
**System** | Pointer to **string** | The name of the system the command was executed on. | [optional] 
**SystemId** | Pointer to **string** | The id of the system the command was executed on. | [optional] 
**User** | Pointer to **string** | The user the command ran as. | [optional] 
**WorkflowId** | Pointer to **string** |  | [optional] 
**WorkflowInstanceId** | Pointer to **string** |  | [optional] 

## Methods

### NewCommandresult

`func NewCommandresult() *Commandresult`

NewCommandresult instantiates a new Commandresult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommandresultWithDefaults

`func NewCommandresultWithDefaults() *Commandresult`

NewCommandresultWithDefaults instantiates a new Commandresult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Commandresult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Commandresult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Commandresult) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Commandresult) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCommand

`func (o *Commandresult) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *Commandresult) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *Commandresult) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *Commandresult) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetFiles

`func (o *Commandresult) GetFiles() []string`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *Commandresult) GetFilesOk() (*[]string, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *Commandresult) SetFiles(v []string)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *Commandresult) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetName

`func (o *Commandresult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Commandresult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Commandresult) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Commandresult) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganization

`func (o *Commandresult) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *Commandresult) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *Commandresult) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *Commandresult) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetRequestTime

`func (o *Commandresult) GetRequestTime() time.Time`

GetRequestTime returns the RequestTime field if non-nil, zero value otherwise.

### GetRequestTimeOk

`func (o *Commandresult) GetRequestTimeOk() (*time.Time, bool)`

GetRequestTimeOk returns a tuple with the RequestTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTime

`func (o *Commandresult) SetRequestTime(v time.Time)`

SetRequestTime sets RequestTime field to given value.

### HasRequestTime

`func (o *Commandresult) HasRequestTime() bool`

HasRequestTime returns a boolean if a field has been set.

### GetResponse

`func (o *Commandresult) GetResponse() CommandresultResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *Commandresult) GetResponseOk() (*CommandresultResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *Commandresult) SetResponse(v CommandresultResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *Commandresult) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetResponseTime

`func (o *Commandresult) GetResponseTime() time.Time`

GetResponseTime returns the ResponseTime field if non-nil, zero value otherwise.

### GetResponseTimeOk

`func (o *Commandresult) GetResponseTimeOk() (*time.Time, bool)`

GetResponseTimeOk returns a tuple with the ResponseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTime

`func (o *Commandresult) SetResponseTime(v time.Time)`

SetResponseTime sets ResponseTime field to given value.

### HasResponseTime

`func (o *Commandresult) HasResponseTime() bool`

HasResponseTime returns a boolean if a field has been set.

### GetSudo

`func (o *Commandresult) GetSudo() bool`

GetSudo returns the Sudo field if non-nil, zero value otherwise.

### GetSudoOk

`func (o *Commandresult) GetSudoOk() (*bool, bool)`

GetSudoOk returns a tuple with the Sudo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSudo

`func (o *Commandresult) SetSudo(v bool)`

SetSudo sets Sudo field to given value.

### HasSudo

`func (o *Commandresult) HasSudo() bool`

HasSudo returns a boolean if a field has been set.

### GetSystem

`func (o *Commandresult) GetSystem() string`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *Commandresult) GetSystemOk() (*string, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *Commandresult) SetSystem(v string)`

SetSystem sets System field to given value.

### HasSystem

`func (o *Commandresult) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetSystemId

`func (o *Commandresult) GetSystemId() string`

GetSystemId returns the SystemId field if non-nil, zero value otherwise.

### GetSystemIdOk

`func (o *Commandresult) GetSystemIdOk() (*string, bool)`

GetSystemIdOk returns a tuple with the SystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemId

`func (o *Commandresult) SetSystemId(v string)`

SetSystemId sets SystemId field to given value.

### HasSystemId

`func (o *Commandresult) HasSystemId() bool`

HasSystemId returns a boolean if a field has been set.

### GetUser

`func (o *Commandresult) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Commandresult) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Commandresult) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *Commandresult) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetWorkflowId

`func (o *Commandresult) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *Commandresult) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *Commandresult) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *Commandresult) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.

### GetWorkflowInstanceId

`func (o *Commandresult) GetWorkflowInstanceId() string`

GetWorkflowInstanceId returns the WorkflowInstanceId field if non-nil, zero value otherwise.

### GetWorkflowInstanceIdOk

`func (o *Commandresult) GetWorkflowInstanceIdOk() (*string, bool)`

GetWorkflowInstanceIdOk returns a tuple with the WorkflowInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowInstanceId

`func (o *Commandresult) SetWorkflowInstanceId(v string)`

SetWorkflowInstanceId sets WorkflowInstanceId field to given value.

### HasWorkflowInstanceId

`func (o *Commandresult) HasWorkflowInstanceId() bool`

HasWorkflowInstanceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


