# SystemInsightsLaunchd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionTime** | Pointer to **string** |  | [optional] 
**Disabled** | Pointer to **string** |  | [optional] 
**Groupname** | Pointer to **string** |  | [optional] 
**InetdCompatibility** | Pointer to **string** |  | [optional] 
**KeepAlive** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OnDemand** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**ProcessType** | Pointer to **string** |  | [optional] 
**Program** | Pointer to **string** |  | [optional] 
**ProgramArguments** | Pointer to **string** |  | [optional] 
**QueueDirectories** | Pointer to **string** |  | [optional] 
**RootDirectory** | Pointer to **string** |  | [optional] 
**RunAtLoad** | Pointer to **string** |  | [optional] 
**StartInterval** | Pointer to **string** |  | [optional] 
**StartOnMount** | Pointer to **string** |  | [optional] 
**StderrPath** | Pointer to **string** |  | [optional] 
**StdoutPath** | Pointer to **string** |  | [optional] 
**SystemId** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**WatchPaths** | Pointer to **string** |  | [optional] 
**WorkingDirectory** | Pointer to **string** |  | [optional] 

## Methods

### NewSystemInsightsLaunchd

`func NewSystemInsightsLaunchd() *SystemInsightsLaunchd`

NewSystemInsightsLaunchd instantiates a new SystemInsightsLaunchd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemInsightsLaunchdWithDefaults

`func NewSystemInsightsLaunchdWithDefaults() *SystemInsightsLaunchd`

NewSystemInsightsLaunchdWithDefaults instantiates a new SystemInsightsLaunchd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionTime

`func (o *SystemInsightsLaunchd) GetCollectionTime() string`

GetCollectionTime returns the CollectionTime field if non-nil, zero value otherwise.

### GetCollectionTimeOk

`func (o *SystemInsightsLaunchd) GetCollectionTimeOk() (*string, bool)`

GetCollectionTimeOk returns a tuple with the CollectionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionTime

`func (o *SystemInsightsLaunchd) SetCollectionTime(v string)`

SetCollectionTime sets CollectionTime field to given value.

### HasCollectionTime

`func (o *SystemInsightsLaunchd) HasCollectionTime() bool`

HasCollectionTime returns a boolean if a field has been set.

### GetDisabled

`func (o *SystemInsightsLaunchd) GetDisabled() string`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *SystemInsightsLaunchd) GetDisabledOk() (*string, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *SystemInsightsLaunchd) SetDisabled(v string)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *SystemInsightsLaunchd) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetGroupname

`func (o *SystemInsightsLaunchd) GetGroupname() string`

GetGroupname returns the Groupname field if non-nil, zero value otherwise.

### GetGroupnameOk

`func (o *SystemInsightsLaunchd) GetGroupnameOk() (*string, bool)`

GetGroupnameOk returns a tuple with the Groupname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupname

`func (o *SystemInsightsLaunchd) SetGroupname(v string)`

SetGroupname sets Groupname field to given value.

### HasGroupname

`func (o *SystemInsightsLaunchd) HasGroupname() bool`

HasGroupname returns a boolean if a field has been set.

### GetInetdCompatibility

`func (o *SystemInsightsLaunchd) GetInetdCompatibility() string`

GetInetdCompatibility returns the InetdCompatibility field if non-nil, zero value otherwise.

### GetInetdCompatibilityOk

`func (o *SystemInsightsLaunchd) GetInetdCompatibilityOk() (*string, bool)`

GetInetdCompatibilityOk returns a tuple with the InetdCompatibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInetdCompatibility

`func (o *SystemInsightsLaunchd) SetInetdCompatibility(v string)`

SetInetdCompatibility sets InetdCompatibility field to given value.

### HasInetdCompatibility

`func (o *SystemInsightsLaunchd) HasInetdCompatibility() bool`

HasInetdCompatibility returns a boolean if a field has been set.

### GetKeepAlive

`func (o *SystemInsightsLaunchd) GetKeepAlive() string`

GetKeepAlive returns the KeepAlive field if non-nil, zero value otherwise.

### GetKeepAliveOk

`func (o *SystemInsightsLaunchd) GetKeepAliveOk() (*string, bool)`

GetKeepAliveOk returns a tuple with the KeepAlive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepAlive

`func (o *SystemInsightsLaunchd) SetKeepAlive(v string)`

SetKeepAlive sets KeepAlive field to given value.

### HasKeepAlive

`func (o *SystemInsightsLaunchd) HasKeepAlive() bool`

HasKeepAlive returns a boolean if a field has been set.

### GetLabel

`func (o *SystemInsightsLaunchd) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *SystemInsightsLaunchd) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *SystemInsightsLaunchd) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *SystemInsightsLaunchd) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *SystemInsightsLaunchd) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SystemInsightsLaunchd) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SystemInsightsLaunchd) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SystemInsightsLaunchd) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOnDemand

`func (o *SystemInsightsLaunchd) GetOnDemand() string`

GetOnDemand returns the OnDemand field if non-nil, zero value otherwise.

### GetOnDemandOk

`func (o *SystemInsightsLaunchd) GetOnDemandOk() (*string, bool)`

GetOnDemandOk returns a tuple with the OnDemand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnDemand

`func (o *SystemInsightsLaunchd) SetOnDemand(v string)`

SetOnDemand sets OnDemand field to given value.

### HasOnDemand

`func (o *SystemInsightsLaunchd) HasOnDemand() bool`

HasOnDemand returns a boolean if a field has been set.

### GetPath

`func (o *SystemInsightsLaunchd) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *SystemInsightsLaunchd) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *SystemInsightsLaunchd) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *SystemInsightsLaunchd) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetProcessType

`func (o *SystemInsightsLaunchd) GetProcessType() string`

GetProcessType returns the ProcessType field if non-nil, zero value otherwise.

### GetProcessTypeOk

`func (o *SystemInsightsLaunchd) GetProcessTypeOk() (*string, bool)`

GetProcessTypeOk returns a tuple with the ProcessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessType

`func (o *SystemInsightsLaunchd) SetProcessType(v string)`

SetProcessType sets ProcessType field to given value.

### HasProcessType

`func (o *SystemInsightsLaunchd) HasProcessType() bool`

HasProcessType returns a boolean if a field has been set.

### GetProgram

`func (o *SystemInsightsLaunchd) GetProgram() string`

GetProgram returns the Program field if non-nil, zero value otherwise.

### GetProgramOk

`func (o *SystemInsightsLaunchd) GetProgramOk() (*string, bool)`

GetProgramOk returns a tuple with the Program field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgram

`func (o *SystemInsightsLaunchd) SetProgram(v string)`

SetProgram sets Program field to given value.

### HasProgram

`func (o *SystemInsightsLaunchd) HasProgram() bool`

HasProgram returns a boolean if a field has been set.

### GetProgramArguments

`func (o *SystemInsightsLaunchd) GetProgramArguments() string`

GetProgramArguments returns the ProgramArguments field if non-nil, zero value otherwise.

### GetProgramArgumentsOk

`func (o *SystemInsightsLaunchd) GetProgramArgumentsOk() (*string, bool)`

GetProgramArgumentsOk returns a tuple with the ProgramArguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramArguments

`func (o *SystemInsightsLaunchd) SetProgramArguments(v string)`

SetProgramArguments sets ProgramArguments field to given value.

### HasProgramArguments

`func (o *SystemInsightsLaunchd) HasProgramArguments() bool`

HasProgramArguments returns a boolean if a field has been set.

### GetQueueDirectories

`func (o *SystemInsightsLaunchd) GetQueueDirectories() string`

GetQueueDirectories returns the QueueDirectories field if non-nil, zero value otherwise.

### GetQueueDirectoriesOk

`func (o *SystemInsightsLaunchd) GetQueueDirectoriesOk() (*string, bool)`

GetQueueDirectoriesOk returns a tuple with the QueueDirectories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueDirectories

`func (o *SystemInsightsLaunchd) SetQueueDirectories(v string)`

SetQueueDirectories sets QueueDirectories field to given value.

### HasQueueDirectories

`func (o *SystemInsightsLaunchd) HasQueueDirectories() bool`

HasQueueDirectories returns a boolean if a field has been set.

### GetRootDirectory

`func (o *SystemInsightsLaunchd) GetRootDirectory() string`

GetRootDirectory returns the RootDirectory field if non-nil, zero value otherwise.

### GetRootDirectoryOk

`func (o *SystemInsightsLaunchd) GetRootDirectoryOk() (*string, bool)`

GetRootDirectoryOk returns a tuple with the RootDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDirectory

`func (o *SystemInsightsLaunchd) SetRootDirectory(v string)`

SetRootDirectory sets RootDirectory field to given value.

### HasRootDirectory

`func (o *SystemInsightsLaunchd) HasRootDirectory() bool`

HasRootDirectory returns a boolean if a field has been set.

### GetRunAtLoad

`func (o *SystemInsightsLaunchd) GetRunAtLoad() string`

GetRunAtLoad returns the RunAtLoad field if non-nil, zero value otherwise.

### GetRunAtLoadOk

`func (o *SystemInsightsLaunchd) GetRunAtLoadOk() (*string, bool)`

GetRunAtLoadOk returns a tuple with the RunAtLoad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunAtLoad

`func (o *SystemInsightsLaunchd) SetRunAtLoad(v string)`

SetRunAtLoad sets RunAtLoad field to given value.

### HasRunAtLoad

`func (o *SystemInsightsLaunchd) HasRunAtLoad() bool`

HasRunAtLoad returns a boolean if a field has been set.

### GetStartInterval

`func (o *SystemInsightsLaunchd) GetStartInterval() string`

GetStartInterval returns the StartInterval field if non-nil, zero value otherwise.

### GetStartIntervalOk

`func (o *SystemInsightsLaunchd) GetStartIntervalOk() (*string, bool)`

GetStartIntervalOk returns a tuple with the StartInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartInterval

`func (o *SystemInsightsLaunchd) SetStartInterval(v string)`

SetStartInterval sets StartInterval field to given value.

### HasStartInterval

`func (o *SystemInsightsLaunchd) HasStartInterval() bool`

HasStartInterval returns a boolean if a field has been set.

### GetStartOnMount

`func (o *SystemInsightsLaunchd) GetStartOnMount() string`

GetStartOnMount returns the StartOnMount field if non-nil, zero value otherwise.

### GetStartOnMountOk

`func (o *SystemInsightsLaunchd) GetStartOnMountOk() (*string, bool)`

GetStartOnMountOk returns a tuple with the StartOnMount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartOnMount

`func (o *SystemInsightsLaunchd) SetStartOnMount(v string)`

SetStartOnMount sets StartOnMount field to given value.

### HasStartOnMount

`func (o *SystemInsightsLaunchd) HasStartOnMount() bool`

HasStartOnMount returns a boolean if a field has been set.

### GetStderrPath

`func (o *SystemInsightsLaunchd) GetStderrPath() string`

GetStderrPath returns the StderrPath field if non-nil, zero value otherwise.

### GetStderrPathOk

`func (o *SystemInsightsLaunchd) GetStderrPathOk() (*string, bool)`

GetStderrPathOk returns a tuple with the StderrPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStderrPath

`func (o *SystemInsightsLaunchd) SetStderrPath(v string)`

SetStderrPath sets StderrPath field to given value.

### HasStderrPath

`func (o *SystemInsightsLaunchd) HasStderrPath() bool`

HasStderrPath returns a boolean if a field has been set.

### GetStdoutPath

`func (o *SystemInsightsLaunchd) GetStdoutPath() string`

GetStdoutPath returns the StdoutPath field if non-nil, zero value otherwise.

### GetStdoutPathOk

`func (o *SystemInsightsLaunchd) GetStdoutPathOk() (*string, bool)`

GetStdoutPathOk returns a tuple with the StdoutPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdoutPath

`func (o *SystemInsightsLaunchd) SetStdoutPath(v string)`

SetStdoutPath sets StdoutPath field to given value.

### HasStdoutPath

`func (o *SystemInsightsLaunchd) HasStdoutPath() bool`

HasStdoutPath returns a boolean if a field has been set.

### GetSystemId

`func (o *SystemInsightsLaunchd) GetSystemId() string`

GetSystemId returns the SystemId field if non-nil, zero value otherwise.

### GetSystemIdOk

`func (o *SystemInsightsLaunchd) GetSystemIdOk() (*string, bool)`

GetSystemIdOk returns a tuple with the SystemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemId

`func (o *SystemInsightsLaunchd) SetSystemId(v string)`

SetSystemId sets SystemId field to given value.

### HasSystemId

`func (o *SystemInsightsLaunchd) HasSystemId() bool`

HasSystemId returns a boolean if a field has been set.

### GetUsername

`func (o *SystemInsightsLaunchd) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SystemInsightsLaunchd) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SystemInsightsLaunchd) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *SystemInsightsLaunchd) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetWatchPaths

`func (o *SystemInsightsLaunchd) GetWatchPaths() string`

GetWatchPaths returns the WatchPaths field if non-nil, zero value otherwise.

### GetWatchPathsOk

`func (o *SystemInsightsLaunchd) GetWatchPathsOk() (*string, bool)`

GetWatchPathsOk returns a tuple with the WatchPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchPaths

`func (o *SystemInsightsLaunchd) SetWatchPaths(v string)`

SetWatchPaths sets WatchPaths field to given value.

### HasWatchPaths

`func (o *SystemInsightsLaunchd) HasWatchPaths() bool`

HasWatchPaths returns a boolean if a field has been set.

### GetWorkingDirectory

`func (o *SystemInsightsLaunchd) GetWorkingDirectory() string`

GetWorkingDirectory returns the WorkingDirectory field if non-nil, zero value otherwise.

### GetWorkingDirectoryOk

`func (o *SystemInsightsLaunchd) GetWorkingDirectoryOk() (*string, bool)`

GetWorkingDirectoryOk returns a tuple with the WorkingDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingDirectory

`func (o *SystemInsightsLaunchd) SetWorkingDirectory(v string)`

SetWorkingDirectory sets WorkingDirectory field to given value.

### HasWorkingDirectory

`func (o *SystemInsightsLaunchd) HasWorkingDirectory() bool`

HasWorkingDirectory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


