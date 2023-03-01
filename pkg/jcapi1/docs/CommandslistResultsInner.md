# CommandslistResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the command. | [optional] 
**Command** | Pointer to **string** | The Command to execute. | [optional] 
**CommandType** | Pointer to **string** | The Command OS. | [optional] 
**LaunchType** | Pointer to **string** | How the Command is executed. | [optional] 
**ListensTo** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | The name of the Command. | [optional] 
**Organization** | Pointer to **string** | The ID of the Organization. | [optional] 
**Schedule** | Pointer to **string** | A crontab that consists of: [ (seconds) (minutes) (hours) (days of month) (months) (weekdays) ] or [ immediate ]. If you send this as an empty string, it will run immediately. | [optional] 
**ScheduleRepeatType** | Pointer to **string** | When the command will repeat. | [optional] 
**Trigger** | Pointer to **string** | Trigger to execute command. | [optional] 

## Methods

### NewCommandslistResultsInner

`func NewCommandslistResultsInner() *CommandslistResultsInner`

NewCommandslistResultsInner instantiates a new CommandslistResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommandslistResultsInnerWithDefaults

`func NewCommandslistResultsInnerWithDefaults() *CommandslistResultsInner`

NewCommandslistResultsInnerWithDefaults instantiates a new CommandslistResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CommandslistResultsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommandslistResultsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommandslistResultsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CommandslistResultsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCommand

`func (o *CommandslistResultsInner) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *CommandslistResultsInner) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *CommandslistResultsInner) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *CommandslistResultsInner) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetCommandType

`func (o *CommandslistResultsInner) GetCommandType() string`

GetCommandType returns the CommandType field if non-nil, zero value otherwise.

### GetCommandTypeOk

`func (o *CommandslistResultsInner) GetCommandTypeOk() (*string, bool)`

GetCommandTypeOk returns a tuple with the CommandType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandType

`func (o *CommandslistResultsInner) SetCommandType(v string)`

SetCommandType sets CommandType field to given value.

### HasCommandType

`func (o *CommandslistResultsInner) HasCommandType() bool`

HasCommandType returns a boolean if a field has been set.

### GetLaunchType

`func (o *CommandslistResultsInner) GetLaunchType() string`

GetLaunchType returns the LaunchType field if non-nil, zero value otherwise.

### GetLaunchTypeOk

`func (o *CommandslistResultsInner) GetLaunchTypeOk() (*string, bool)`

GetLaunchTypeOk returns a tuple with the LaunchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchType

`func (o *CommandslistResultsInner) SetLaunchType(v string)`

SetLaunchType sets LaunchType field to given value.

### HasLaunchType

`func (o *CommandslistResultsInner) HasLaunchType() bool`

HasLaunchType returns a boolean if a field has been set.

### GetListensTo

`func (o *CommandslistResultsInner) GetListensTo() string`

GetListensTo returns the ListensTo field if non-nil, zero value otherwise.

### GetListensToOk

`func (o *CommandslistResultsInner) GetListensToOk() (*string, bool)`

GetListensToOk returns a tuple with the ListensTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListensTo

`func (o *CommandslistResultsInner) SetListensTo(v string)`

SetListensTo sets ListensTo field to given value.

### HasListensTo

`func (o *CommandslistResultsInner) HasListensTo() bool`

HasListensTo returns a boolean if a field has been set.

### GetName

`func (o *CommandslistResultsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CommandslistResultsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CommandslistResultsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CommandslistResultsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganization

`func (o *CommandslistResultsInner) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *CommandslistResultsInner) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *CommandslistResultsInner) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *CommandslistResultsInner) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetSchedule

`func (o *CommandslistResultsInner) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *CommandslistResultsInner) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *CommandslistResultsInner) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *CommandslistResultsInner) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetScheduleRepeatType

`func (o *CommandslistResultsInner) GetScheduleRepeatType() string`

GetScheduleRepeatType returns the ScheduleRepeatType field if non-nil, zero value otherwise.

### GetScheduleRepeatTypeOk

`func (o *CommandslistResultsInner) GetScheduleRepeatTypeOk() (*string, bool)`

GetScheduleRepeatTypeOk returns a tuple with the ScheduleRepeatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleRepeatType

`func (o *CommandslistResultsInner) SetScheduleRepeatType(v string)`

SetScheduleRepeatType sets ScheduleRepeatType field to given value.

### HasScheduleRepeatType

`func (o *CommandslistResultsInner) HasScheduleRepeatType() bool`

HasScheduleRepeatType returns a boolean if a field has been set.

### GetTrigger

`func (o *CommandslistResultsInner) GetTrigger() string`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *CommandslistResultsInner) GetTriggerOk() (*string, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *CommandslistResultsInner) SetTrigger(v string)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *CommandslistResultsInner) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


