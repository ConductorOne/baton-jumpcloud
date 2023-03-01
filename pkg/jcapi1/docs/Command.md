# Command

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | **string** | The command to execute on the server. | 
**CommandRunners** | Pointer to **[]string** | An array of IDs of the Command Runner Users that can execute this command. | [optional] 
**CommandType** | **string** | The Command OS | [default to "linux"]
**Files** | Pointer to **[]string** | An array of file IDs to include with the command. | [optional] 
**LaunchType** | Pointer to **string** | How the command will execute. | [optional] 
**ListensTo** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Organization** | Pointer to **string** | The ID of the organization. | [optional] 
**Schedule** | Pointer to **string** | A crontab that consists of: [ (seconds) (minutes) (hours) (days of month) (months) (weekdays) ] or [ immediate ]. If you send this as an empty string, it will run immediately.  | [optional] 
**ScheduleRepeatType** | Pointer to **string** | When the command will repeat. | [optional] 
**ScheduleYear** | Pointer to **int32** | The year that a scheduled command will launch in. | [optional] 
**Shell** | Pointer to **string** | The shell used to run the command. | [optional] 
**Sudo** | Pointer to **bool** |  | [optional] 
**Systems** | Pointer to **[]string** | Not used. Use /api/v2/commands/{id}/associations to bind commands to systems. | [optional] 
**Template** | Pointer to **string** | The template this command was created from | [optional] 
**TimeToLiveSeconds** | Pointer to **int32** | Time in seconds a command can wait in the queue to be run before timing out | [optional] 
**Timeout** | Pointer to **string** | The time in seconds to allow the command to run for. | [optional] 
**Trigger** | Pointer to **string** | The name of the command trigger. | [optional] 
**User** | Pointer to **string** | The ID of the system user to run the command as. This field is required when creating a command with a commandType of \&quot;mac\&quot; or \&quot;linux\&quot;. | [optional] 

## Methods

### NewCommand

`func NewCommand(command string, commandType string, name string, ) *Command`

NewCommand instantiates a new Command object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommandWithDefaults

`func NewCommandWithDefaults() *Command`

NewCommandWithDefaults instantiates a new Command object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *Command) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *Command) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *Command) SetCommand(v string)`

SetCommand sets Command field to given value.


### GetCommandRunners

`func (o *Command) GetCommandRunners() []string`

GetCommandRunners returns the CommandRunners field if non-nil, zero value otherwise.

### GetCommandRunnersOk

`func (o *Command) GetCommandRunnersOk() (*[]string, bool)`

GetCommandRunnersOk returns a tuple with the CommandRunners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandRunners

`func (o *Command) SetCommandRunners(v []string)`

SetCommandRunners sets CommandRunners field to given value.

### HasCommandRunners

`func (o *Command) HasCommandRunners() bool`

HasCommandRunners returns a boolean if a field has been set.

### GetCommandType

`func (o *Command) GetCommandType() string`

GetCommandType returns the CommandType field if non-nil, zero value otherwise.

### GetCommandTypeOk

`func (o *Command) GetCommandTypeOk() (*string, bool)`

GetCommandTypeOk returns a tuple with the CommandType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandType

`func (o *Command) SetCommandType(v string)`

SetCommandType sets CommandType field to given value.


### GetFiles

`func (o *Command) GetFiles() []string`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *Command) GetFilesOk() (*[]string, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *Command) SetFiles(v []string)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *Command) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetLaunchType

`func (o *Command) GetLaunchType() string`

GetLaunchType returns the LaunchType field if non-nil, zero value otherwise.

### GetLaunchTypeOk

`func (o *Command) GetLaunchTypeOk() (*string, bool)`

GetLaunchTypeOk returns a tuple with the LaunchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchType

`func (o *Command) SetLaunchType(v string)`

SetLaunchType sets LaunchType field to given value.

### HasLaunchType

`func (o *Command) HasLaunchType() bool`

HasLaunchType returns a boolean if a field has been set.

### GetListensTo

`func (o *Command) GetListensTo() string`

GetListensTo returns the ListensTo field if non-nil, zero value otherwise.

### GetListensToOk

`func (o *Command) GetListensToOk() (*string, bool)`

GetListensToOk returns a tuple with the ListensTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListensTo

`func (o *Command) SetListensTo(v string)`

SetListensTo sets ListensTo field to given value.

### HasListensTo

`func (o *Command) HasListensTo() bool`

HasListensTo returns a boolean if a field has been set.

### GetName

`func (o *Command) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Command) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Command) SetName(v string)`

SetName sets Name field to given value.


### GetOrganization

`func (o *Command) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *Command) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *Command) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *Command) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetSchedule

`func (o *Command) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *Command) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *Command) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *Command) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetScheduleRepeatType

`func (o *Command) GetScheduleRepeatType() string`

GetScheduleRepeatType returns the ScheduleRepeatType field if non-nil, zero value otherwise.

### GetScheduleRepeatTypeOk

`func (o *Command) GetScheduleRepeatTypeOk() (*string, bool)`

GetScheduleRepeatTypeOk returns a tuple with the ScheduleRepeatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleRepeatType

`func (o *Command) SetScheduleRepeatType(v string)`

SetScheduleRepeatType sets ScheduleRepeatType field to given value.

### HasScheduleRepeatType

`func (o *Command) HasScheduleRepeatType() bool`

HasScheduleRepeatType returns a boolean if a field has been set.

### GetScheduleYear

`func (o *Command) GetScheduleYear() int32`

GetScheduleYear returns the ScheduleYear field if non-nil, zero value otherwise.

### GetScheduleYearOk

`func (o *Command) GetScheduleYearOk() (*int32, bool)`

GetScheduleYearOk returns a tuple with the ScheduleYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleYear

`func (o *Command) SetScheduleYear(v int32)`

SetScheduleYear sets ScheduleYear field to given value.

### HasScheduleYear

`func (o *Command) HasScheduleYear() bool`

HasScheduleYear returns a boolean if a field has been set.

### GetShell

`func (o *Command) GetShell() string`

GetShell returns the Shell field if non-nil, zero value otherwise.

### GetShellOk

`func (o *Command) GetShellOk() (*string, bool)`

GetShellOk returns a tuple with the Shell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShell

`func (o *Command) SetShell(v string)`

SetShell sets Shell field to given value.

### HasShell

`func (o *Command) HasShell() bool`

HasShell returns a boolean if a field has been set.

### GetSudo

`func (o *Command) GetSudo() bool`

GetSudo returns the Sudo field if non-nil, zero value otherwise.

### GetSudoOk

`func (o *Command) GetSudoOk() (*bool, bool)`

GetSudoOk returns a tuple with the Sudo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSudo

`func (o *Command) SetSudo(v bool)`

SetSudo sets Sudo field to given value.

### HasSudo

`func (o *Command) HasSudo() bool`

HasSudo returns a boolean if a field has been set.

### GetSystems

`func (o *Command) GetSystems() []string`

GetSystems returns the Systems field if non-nil, zero value otherwise.

### GetSystemsOk

`func (o *Command) GetSystemsOk() (*[]string, bool)`

GetSystemsOk returns a tuple with the Systems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystems

`func (o *Command) SetSystems(v []string)`

SetSystems sets Systems field to given value.

### HasSystems

`func (o *Command) HasSystems() bool`

HasSystems returns a boolean if a field has been set.

### GetTemplate

`func (o *Command) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *Command) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *Command) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *Command) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetTimeToLiveSeconds

`func (o *Command) GetTimeToLiveSeconds() int32`

GetTimeToLiveSeconds returns the TimeToLiveSeconds field if non-nil, zero value otherwise.

### GetTimeToLiveSecondsOk

`func (o *Command) GetTimeToLiveSecondsOk() (*int32, bool)`

GetTimeToLiveSecondsOk returns a tuple with the TimeToLiveSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeToLiveSeconds

`func (o *Command) SetTimeToLiveSeconds(v int32)`

SetTimeToLiveSeconds sets TimeToLiveSeconds field to given value.

### HasTimeToLiveSeconds

`func (o *Command) HasTimeToLiveSeconds() bool`

HasTimeToLiveSeconds returns a boolean if a field has been set.

### GetTimeout

`func (o *Command) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *Command) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *Command) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *Command) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetTrigger

`func (o *Command) GetTrigger() string`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *Command) GetTriggerOk() (*string, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *Command) SetTrigger(v string)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *Command) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.

### GetUser

`func (o *Command) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Command) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Command) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *Command) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


