package process

// Length of the process.command_args array
//
// This field can be useful for querying or performing bucket analysis on how many arguments were provided to start a process. More arguments may be an indication of suspicious activity
type AttrArgsCount string // process.args_count

func (AttrArgsCount) Development()    {}
func (AttrArgsCount) Recommended()    {}
func (AttrArgsCount) Key() string     { return "process_args_count" }
func (a AttrArgsCount) Value() string { return string(a) }

// The command used to launch the process (i.e. the command name). On Linux based systems, can be set to the zeroth string in `proc/[pid]/cmdline`. On Windows, can be set to the first parameter extracted from `GetCommandLineW`
type AttrCommand string // process.command

func (AttrCommand) Development()    {}
func (AttrCommand) Recommended()    {}
func (AttrCommand) Key() string     { return "process_command" }
func (a AttrCommand) Value() string { return string(a) }

// All the command arguments (including the command/executable itself) as received by the process. On Linux-based systems (and some other Unixoid systems supporting procfs), can be set according to the list of null-delimited strings extracted from `proc/[pid]/cmdline`. For libc-based executables, this would be the full argv vector passed to `main`. SHOULD NOT be collected by default unless there is sanitization that excludes sensitive data
type AttrCommandArgs string // process.command_args

func (AttrCommandArgs) Development()    {}
func (AttrCommandArgs) Recommended()    {}
func (AttrCommandArgs) Key() string     { return "process_command_args" }
func (a AttrCommandArgs) Value() string { return string(a) }

// The full command used to launch the process as a single string representing the full command. On Windows, can be set to the result of `GetCommandLineW`. Do not set this if you have to assemble it just for monitoring; use `process.command_args` instead. SHOULD NOT be collected by default unless there is sanitization that excludes sensitive data
type AttrCommandLine string // process.command_line

func (AttrCommandLine) Development()    {}
func (AttrCommandLine) Recommended()    {}
func (AttrCommandLine) Key() string     { return "process_command_line" }
func (a AttrCommandLine) Value() string { return string(a) }

// Specifies whether the context switches for this data point were voluntary or involuntary
type AttrContextSwitchType string // process.context_switch_type

func (AttrContextSwitchType) Development()    {}
func (AttrContextSwitchType) Recommended()    {}
func (AttrContextSwitchType) Key() string     { return "process_context_switch_type" }
func (a AttrContextSwitchType) Value() string { return string(a) }

const ContextSwitchTypeVoluntary AttrContextSwitchType = "voluntary"
const ContextSwitchTypeInvoluntary AttrContextSwitchType = "involuntary"

// Deprecated, use `cpu.mode` instead
type AttrCpuState string // process.cpu.state

func (AttrCpuState) Development()    {}
func (AttrCpuState) Recommended()    {}
func (AttrCpuState) Key() string     { return "process_cpu_state" }
func (a AttrCpuState) Value() string { return string(a) }

const CpuStateSystem AttrCpuState = "system"
const CpuStateUser AttrCpuState = "user"
const CpuStateWait AttrCpuState = "wait"

// The date and time the process was created, in ISO 8601 format
type AttrCreationTime string // process.creation.time

func (AttrCreationTime) Development()    {}
func (AttrCreationTime) Recommended()    {}
func (AttrCreationTime) Key() string     { return "process_creation_time" }
func (a AttrCreationTime) Value() string { return string(a) }

// Process environment variables, `<key>` being the environment variable name, the value being the environment variable value.
//
// Examples:
//
//   - an environment variable `USER` with value `"ubuntu"` SHOULD be recorded
//     as the `process.environment_variable.USER` attribute with value `"ubuntu"`.
//   - an environment variable `PATH` with value `"/usr/local/bin:/usr/bin"`
//     SHOULD be recorded as the `process.environment_variable.PATH` attribute
//     with value `"/usr/local/bin:/usr/bin"`
type AttrEnvironmentVariable string // process.environment_variable

func (AttrEnvironmentVariable) Development()    {}
func (AttrEnvironmentVariable) Recommended()    {}
func (AttrEnvironmentVariable) Key() string     { return "process_environment_variable" }
func (a AttrEnvironmentVariable) Value() string { return string(a) }

// The GNU build ID as found in the `.note.gnu.build-id` ELF section (hex string)
type AttrExecutableBuildIdGnu string // process.executable.build_id.gnu

func (AttrExecutableBuildIdGnu) Development()    {}
func (AttrExecutableBuildIdGnu) Recommended()    {}
func (AttrExecutableBuildIdGnu) Key() string     { return "process_executable_build_id_gnu" }
func (a AttrExecutableBuildIdGnu) Value() string { return string(a) }

// The Go build ID as retrieved by `go tool buildid <go executable>`
type AttrExecutableBuildIdGo string // process.executable.build_id.go

func (AttrExecutableBuildIdGo) Development()    {}
func (AttrExecutableBuildIdGo) Recommended()    {}
func (AttrExecutableBuildIdGo) Key() string     { return "process_executable_build_id_go" }
func (a AttrExecutableBuildIdGo) Value() string { return string(a) }

// Profiling specific build ID for executables. See the OTel specification for Profiles for more information
type AttrExecutableBuildIdHtlhash string // process.executable.build_id.htlhash

func (AttrExecutableBuildIdHtlhash) Development()    {}
func (AttrExecutableBuildIdHtlhash) Recommended()    {}
func (AttrExecutableBuildIdHtlhash) Key() string     { return "process_executable_build_id_htlhash" }
func (a AttrExecutableBuildIdHtlhash) Value() string { return string(a) }

// "Deprecated, use `process.executable.build_id.htlhash` instead."
type AttrExecutableBuildIdProfiling string // process.executable.build_id.profiling

func (AttrExecutableBuildIdProfiling) Development()    {}
func (AttrExecutableBuildIdProfiling) Recommended()    {}
func (AttrExecutableBuildIdProfiling) Key() string     { return "process_executable_build_id_profiling" }
func (a AttrExecutableBuildIdProfiling) Value() string { return string(a) }

// The name of the process executable. On Linux based systems, this SHOULD be set to the base name of the target of `/proc/[pid]/exe`. On Windows, this SHOULD be set to the base name of `GetProcessImageFileNameW`
type AttrExecutableName string // process.executable.name

func (AttrExecutableName) Development()    {}
func (AttrExecutableName) Recommended()    {}
func (AttrExecutableName) Key() string     { return "process_executable_name" }
func (a AttrExecutableName) Value() string { return string(a) }

// The full path to the process executable. On Linux based systems, can be set to the target of `proc/[pid]/exe`. On Windows, can be set to the result of `GetProcessImageFileNameW`
type AttrExecutablePath string // process.executable.path

func (AttrExecutablePath) Development()    {}
func (AttrExecutablePath) Recommended()    {}
func (AttrExecutablePath) Key() string     { return "process_executable_path" }
func (a AttrExecutablePath) Value() string { return string(a) }

// The exit code of the process
type AttrExitCode string // process.exit.code

func (AttrExitCode) Development()    {}
func (AttrExitCode) Recommended()    {}
func (AttrExitCode) Key() string     { return "process_exit_code" }
func (a AttrExitCode) Value() string { return string(a) }

// The date and time the process exited, in ISO 8601 format
type AttrExitTime string // process.exit.time

func (AttrExitTime) Development()    {}
func (AttrExitTime) Recommended()    {}
func (AttrExitTime) Key() string     { return "process_exit_time" }
func (a AttrExitTime) Value() string { return string(a) }

// The PID of the process's group leader. This is also the process group ID (PGID) of the process
type AttrGroupLeaderPid string // process.group_leader.pid

func (AttrGroupLeaderPid) Development()    {}
func (AttrGroupLeaderPid) Recommended()    {}
func (AttrGroupLeaderPid) Key() string     { return "process_group_leader_pid" }
func (a AttrGroupLeaderPid) Value() string { return string(a) }

// Whether the process is connected to an interactive shell
type AttrInteractive string // process.interactive

func (AttrInteractive) Development()    {}
func (AttrInteractive) Recommended()    {}
func (AttrInteractive) Key() string     { return "process_interactive" }
func (a AttrInteractive) Value() string { return string(a) }

// The control group associated with the process.
// Control groups (cgroups) are a kernel feature used to organize and manage process resources. This attribute provides the path(s) to the cgroup(s) associated with the process, which should match the contents of the [/proc/[PID]/cgroup] file
//
// [/proc/[PID]/cgroup]: https://man7.org/linux/man-pages/man7/cgroups.7.html
type AttrLinuxCgroup string // process.linux.cgroup

func (AttrLinuxCgroup) Development()    {}
func (AttrLinuxCgroup) Recommended()    {}
func (AttrLinuxCgroup) Key() string     { return "process_linux_cgroup" }
func (a AttrLinuxCgroup) Value() string { return string(a) }

// The username of the user that owns the process
type AttrOwner string // process.owner

func (AttrOwner) Development()    {}
func (AttrOwner) Recommended()    {}
func (AttrOwner) Key() string     { return "process_owner" }
func (a AttrOwner) Value() string { return string(a) }

// The type of page fault for this data point. Type `major` is for major/hard page faults, and `minor` is for minor/soft page faults
type AttrPagingFaultType string // process.paging.fault_type

func (AttrPagingFaultType) Development()    {}
func (AttrPagingFaultType) Recommended()    {}
func (AttrPagingFaultType) Key() string     { return "process_paging_fault_type" }
func (a AttrPagingFaultType) Value() string { return string(a) }

const PagingFaultTypeMajor AttrPagingFaultType = "major"
const PagingFaultTypeMinor AttrPagingFaultType = "minor"

// Parent Process identifier (PPID)
type AttrParentPid string // process.parent_pid

func (AttrParentPid) Development()    {}
func (AttrParentPid) Recommended()    {}
func (AttrParentPid) Key() string     { return "process_parent_pid" }
func (a AttrParentPid) Value() string { return string(a) }

// Process identifier (PID)
type AttrPid string // process.pid

func (AttrPid) Development()    {}
func (AttrPid) Recommended()    {}
func (AttrPid) Key() string     { return "process_pid" }
func (a AttrPid) Value() string { return string(a) }

// The real user ID (RUID) of the process
type AttrRealUserId string // process.real_user.id

func (AttrRealUserId) Development()    {}
func (AttrRealUserId) Recommended()    {}
func (AttrRealUserId) Key() string     { return "process_real_user_id" }
func (a AttrRealUserId) Value() string { return string(a) }

// The username of the real user of the process
type AttrRealUserName string // process.real_user.name

func (AttrRealUserName) Development()    {}
func (AttrRealUserName) Recommended()    {}
func (AttrRealUserName) Key() string     { return "process_real_user_name" }
func (a AttrRealUserName) Value() string { return string(a) }

// An additional description about the runtime of the process, for example a specific vendor customization of the runtime environment
type AttrRuntimeDescription string // process.runtime.description

func (AttrRuntimeDescription) Development()    {}
func (AttrRuntimeDescription) Recommended()    {}
func (AttrRuntimeDescription) Key() string     { return "process_runtime_description" }
func (a AttrRuntimeDescription) Value() string { return string(a) }

// The name of the runtime of this process
type AttrRuntimeName string // process.runtime.name

func (AttrRuntimeName) Development()    {}
func (AttrRuntimeName) Recommended()    {}
func (AttrRuntimeName) Key() string     { return "process_runtime_name" }
func (a AttrRuntimeName) Value() string { return string(a) }

// The version of the runtime of this process, as returned by the runtime without modification
type AttrRuntimeVersion string // process.runtime.version

func (AttrRuntimeVersion) Development()    {}
func (AttrRuntimeVersion) Recommended()    {}
func (AttrRuntimeVersion) Key() string     { return "process_runtime_version" }
func (a AttrRuntimeVersion) Value() string { return string(a) }

// The saved user ID (SUID) of the process
type AttrSavedUserId string // process.saved_user.id

func (AttrSavedUserId) Development()    {}
func (AttrSavedUserId) Recommended()    {}
func (AttrSavedUserId) Key() string     { return "process_saved_user_id" }
func (a AttrSavedUserId) Value() string { return string(a) }

// The username of the saved user
type AttrSavedUserName string // process.saved_user.name

func (AttrSavedUserName) Development()    {}
func (AttrSavedUserName) Recommended()    {}
func (AttrSavedUserName) Key() string     { return "process_saved_user_name" }
func (a AttrSavedUserName) Value() string { return string(a) }

// The PID of the process's session leader. This is also the session ID (SID) of the process
type AttrSessionLeaderPid string // process.session_leader.pid

func (AttrSessionLeaderPid) Development()    {}
func (AttrSessionLeaderPid) Recommended()    {}
func (AttrSessionLeaderPid) Key() string     { return "process_session_leader_pid" }
func (a AttrSessionLeaderPid) Value() string { return string(a) }

// Process title (proctitle)
//
// In many Unix-like systems, process title (proctitle), is the string that represents the name or command line of a running process, displayed by system monitoring tools like ps, top, and htop
type AttrTitle string // process.title

func (AttrTitle) Development()    {}
func (AttrTitle) Recommended()    {}
func (AttrTitle) Key() string     { return "process_title" }
func (a AttrTitle) Value() string { return string(a) }

// The effective user ID (EUID) of the process
type AttrUserId string // process.user.id

func (AttrUserId) Development()    {}
func (AttrUserId) Recommended()    {}
func (AttrUserId) Key() string     { return "process_user_id" }
func (a AttrUserId) Value() string { return string(a) }

// The username of the effective user of the process
type AttrUserName string // process.user.name

func (AttrUserName) Development()    {}
func (AttrUserName) Recommended()    {}
func (AttrUserName) Key() string     { return "process_user_name" }
func (a AttrUserName) Value() string { return string(a) }

// Virtual process identifier.
//
// The process ID within a PID namespace. This is not necessarily unique across all processes on the host but it is unique within the process namespace that the process exists within
type AttrVpid string // process.vpid

func (AttrVpid) Development()    {}
func (AttrVpid) Recommended()    {}
func (AttrVpid) Key() string     { return "process_vpid" }
func (a AttrVpid) Value() string { return string(a) }

// The working directory of the process
type AttrWorkingDirectory string // process.working_directory

func (AttrWorkingDirectory) Development()    {}
func (AttrWorkingDirectory) Recommended()    {}
func (AttrWorkingDirectory) Key() string     { return "process_working_directory" }
func (a AttrWorkingDirectory) Value() string { return string(a) }

/* State {
    name: "attr.go.j2",
    current_block: None,
    auto_escape: None,
    ctx: {
        "ctx": {
            "attributes": [
                {
                    "brief": "Length of the process.command_args array\n",
                    "examples": [
                        4,
                    ],
                    "name": "process.args_count",
                    "note": "This field can be useful for querying or performing bucket analysis on how many arguments were provided to start a process. More arguments may be an indication of suspicious activity.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "process",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "The command used to launch the process (i.e. the command name). On Linux based systems, can be set to the zeroth string in `proc/[pid]/cmdline`. On Windows, can be set to the first parameter extracted from `GetCommandLineW`.\n",
                    "examples": [
                        "cmd/otelcol",
                    ],
                    "name": "process.command",
                    "requirement_level": "recommended",
                    "root_namespace": "process",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "All the command arguments (including the command/executable itself) as received by the process. On Linux-based systems (and some other Unixoid systems supporting procfs), can be set according to the list of null-delimited strings extracted from `proc/[pid]/cmdline`. For libc-based executables, this would be the full argv vector passed to `main`. SHOULD NOT be collected by default unless there is sanitization that excludes sensitive data.\n",
                    "examples": [
                        [
                            "cmd/otecol",
                            "--config=config.yaml",
                        ],
                    ],
                    "name": "process.command_args",
                    "requirement_level": "recommended",
                    "root_namespace": "process",
                    "stability": "development",
                    "type": "string[]",
                },
                {
                    "brief": "The full command used to launch the process as a single string representing the full command. On Windows, can be set to the result of `GetCommandLineW`. Do not set this if you have to assemble it just for monitoring; use `process.command_args` instead. SHOULD NOT be collected by default unless there is sanitization that excludes sensitive data.\n",
                    "examples": [
                        "C:\\cmd\\otecol --config=\"my directory\\config.yaml\"",
                    ],
                    "name": "process.command_line",
                    "requirement_level": "recommended",
                    "root_namespace": "process",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Specifies whether the context switches for this data point were voluntary or involuntary.",
                    "name": "process.context_switch_type",
                    "requirement_level": "recommended",
                    "root_namespace": "process",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "voluntary",
                                "note": none,
                                "stability": "development",
                                "value": "voluntary",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "involuntary",
                                "note": none,
                                "stability": "development",
                                "value": "involuntary",
                            },
                        ],
                    },
                },
                {
                    "brief": "Deprecated, use `cpu.mode` instead.",
                    "deprecated": {
                        "note": "Replaced by `cpu.mode`.",
                        "reason": "renamed",
                        "renamed_to": "cpu.mode",
                    },
                    "name": "process.cpu.state",
                    "requirement_level": "recommended",
                    "root_namespace": "process",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "system",
                                "note": none,
                                "stability": "development",
                                "value": "system",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "user",
                                "note": none,
                                "stability": "development",
                                "value": "user",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "wait",
                                "note": none,
                                "stability": "development",
                                "value": "wait",
                            },
                        ],
                    },
                },
                {
                    "brief": "The date and time the process was created, in ISO 8601 format.\n",
                    "examples": [
                        "2023-11-21T09:25:34.853Z",
                    ],
                    "name": "process.creation.time",
                    "requirement_level": "recommended",
                    "root_namespace": "process",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Process environment variables, `<key>` being the environment variable name, the value being the environment variable value.\n",
                    "examples": [
                        "ubuntu",
                        "/usr/local/bin:/usr/bin",
                    ],
                    "name": "process.environment_variable",
                    "note": "Examples:\n\n- an environment variable `USER` with value `\"ubuntu\"` SHOULD be recorded\nas the `process.environment_variable.USER` attribute with value `\"ubuntu\"`.\n\n- an environment variable `PATH` with value `\"/usr/local/bin:/usr/bin\"`\nSHOULD be recorded as the `process.environment_variable.PATH` attribute\nwith value `\"/usr/local/bin:/usr/bin\"`.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "process",
                    "stability": "development",
                    "type": "template[string]",
                },
                {
                    "brief": "The GNU build ID as found in the `.note.gnu.build-id` ELF section (hex string).\n",
                    "examples": [
                        "c89b11207f6479603b0d49bf291c092c2b719293",
                    ],
                    "name": "process.executable.build_id.gnu",
                    "requirement_level": "recommended",
                    "root_namespace": "process",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The Go build ID as retrieved by `go tool buildid <go executable>`.\n",
                    "examples": [
                        "foh3mEXu7BLZjsN9pOwG/kATcXlYVCDEFouRMQed_/WwRFB1hPo9LBkekthSPG/x8hMC8emW2cCjXD0_1aY",
                    ],
                    "name": "process.executable.build_id.go",
                    "requirement_level": "recommended",
                    "root_namespace": "process",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Profiling specific build ID for executables. See the OTel specification for Profiles for more information.\n",
                    "examples": [
                        "600DCAFE4A110000F2BF38C493F5FB92",
                    ],
                    "name": "process.executable.build_id.htlhash",
                    "requirement_level": "recommended",
                    "root_namespace": "process",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "\"Deprecated, use `process.executable.build_id.htlhash` instead.\"\n",
                    "deprecated": {
                        "note": "Replaced by `process.executable.build_id.htlhash`.",
                        "reason": "renamed",
                        "renamed_to": "process.executable.build_id.htlhash",
                    },
                    "examples": [
                        "600DCAFE4A110000F2BF38C493F5FB92",
                    ],
                    "name": "process.executable.build_id.profiling",
                    "requirement_level": "recommended",
                    "root_namespace": "process",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The name of the process executable. On Linux based systems, this SHOULD be set to the base name of the target of `/proc/[pid]/exe`. On Windows, this SHOULD be set to the base name of `GetProcessImageFileNameW`.\n",
                    "examples": [
                        "otelcol",
                    ],
                    "name": "process.executable.name",
                    "requirement_level": "recommended",
                    "root_namespace": "process",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The full path to the process executable. On Linux based systems, can be set to the target of `proc/[pid]/exe`. On Windows, can be set to the result of `GetProcessImageFileNameW`.\n",
                    "examples": [
                        "/usr/bin/cmd/otelcol",
                    ],
                    "name": "process.executable.path",
                    "requirement_level": "recommended",
                    "root_namespace": "process",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The exit code of the process.\n",
                    "examples": [
                        127,
                    ],
                    "name": "process.exit.code",
                    "requirement_level": "recommended",
                    "root_namespace": "process",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "The date and time the process exited, in ISO 8601 format.\n",
                    "examples": [
                        "2023-11-21T09:26:12.315Z",
                    ],
                    "name": "process.exit.time",
                    "requirement_level": "recommended",
                    "root_namespace": "process",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The PID of the process's group leader. This is also the process group ID (PGID) of the process.\n",
                    "examples": [
                        23,
                    ],
                    "name": "process.group_leader.pid",
                    "requirement_level": "recommended",
                    "root_namespace": "process",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "Whether the process is connected to an interactive shell.\n",
                    "name": "process.interactive",
                    "requirement_level": "recommended",
                    "root_namespace": "process",
                    "stability": "development",
                    "type": "boolean",
                },
                {
                    "brief": "The control group associated with the process.",
                    "examples": [
                        "1:name=systemd:/user.slice/user-1000.slice/session-3.scope",
                        "0::/user.slice/user-1000.slice/user@1000.service/tmux-spawn-0267755b-4639-4a27-90ed-f19f88e53748.scope",
                    ],
                    "name": "process.linux.cgroup",
                    "note": "Control groups (cgroups) are a kernel feature used to organize and manage process resources. This attribute provides the path(s) to the cgroup(s) associated with the process, which should match the contents of the [/proc/\\[PID\\]/cgroup](https://man7.org/linux/man-pages/man7/cgroups.7.html) file.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "process",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The username of the user that owns the process.\n",
                    "examples": [
                        "root",
                    ],
                    "name": "process.owner",
                    "requirement_level": "recommended",
                    "root_namespace": "process",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The type of page fault for this data point. Type `major` is for major/hard page faults, and `minor` is for minor/soft page faults.\n",
                    "name": "process.paging.fault_type",
                    "requirement_level": "recommended",
                    "root_namespace": "process",
                    "stability": "development",
                    "type": {
                        "members": [
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "major",
                                "note": none,
                                "stability": "development",
                                "value": "major",
                            },
                            {
                                "brief": none,
                                "deprecated": none,
                                "id": "minor",
                                "note": none,
                                "stability": "development",
                                "value": "minor",
                            },
                        ],
                    },
                },
                {
                    "brief": "Parent Process identifier (PPID).\n",
                    "examples": [
                        111,
                    ],
                    "name": "process.parent_pid",
                    "requirement_level": "recommended",
                    "root_namespace": "process",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "Process identifier (PID).\n",
                    "examples": [
                        1234,
                    ],
                    "name": "process.pid",
                    "requirement_level": "recommended",
                    "root_namespace": "process",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "The real user ID (RUID) of the process.\n",
                    "examples": [
                        1000,
                    ],
                    "name": "process.real_user.id",
                    "requirement_level": "recommended",
                    "root_namespace": "process",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "The username of the real user of the process.\n",
                    "examples": [
                        "operator",
                    ],
                    "name": "process.real_user.name",
                    "requirement_level": "recommended",
                    "root_namespace": "process",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "An additional description about the runtime of the process, for example a specific vendor customization of the runtime environment.\n",
                    "examples": "Eclipse OpenJ9 Eclipse OpenJ9 VM openj9-0.21.0",
                    "name": "process.runtime.description",
                    "requirement_level": "recommended",
                    "root_namespace": "process",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The name of the runtime of this process.\n",
                    "examples": [
                        "OpenJDK Runtime Environment",
                    ],
                    "name": "process.runtime.name",
                    "requirement_level": "recommended",
                    "root_namespace": "process",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The version of the runtime of this process, as returned by the runtime without modification.\n",
                    "examples": "14.0.2",
                    "name": "process.runtime.version",
                    "requirement_level": "recommended",
                    "root_namespace": "process",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The saved user ID (SUID) of the process.\n",
                    "examples": [
                        1002,
                    ],
                    "name": "process.saved_user.id",
                    "requirement_level": "recommended",
                    "root_namespace": "process",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "The username of the saved user.\n",
                    "examples": [
                        "operator",
                    ],
                    "name": "process.saved_user.name",
                    "requirement_level": "recommended",
                    "root_namespace": "process",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The PID of the process's session leader. This is also the session ID (SID) of the process.\n",
                    "examples": [
                        14,
                    ],
                    "name": "process.session_leader.pid",
                    "requirement_level": "recommended",
                    "root_namespace": "process",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "Process title (proctitle)\n",
                    "examples": [
                        "cat /etc/hostname",
                        "xfce4-session",
                        "bash",
                    ],
                    "name": "process.title",
                    "note": "In many Unix-like systems, process title (proctitle), is the string that represents the name or command line of a running process, displayed by system monitoring tools like ps, top, and htop.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "process",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "The effective user ID (EUID) of the process.\n",
                    "examples": [
                        1001,
                    ],
                    "name": "process.user.id",
                    "requirement_level": "recommended",
                    "root_namespace": "process",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "The username of the effective user of the process.\n",
                    "examples": [
                        "root",
                    ],
                    "name": "process.user.name",
                    "requirement_level": "recommended",
                    "root_namespace": "process",
                    "stability": "development",
                    "type": "string",
                },
                {
                    "brief": "Virtual process identifier.\n",
                    "examples": [
                        12,
                    ],
                    "name": "process.vpid",
                    "note": "The process ID within a PID namespace. This is not necessarily unique across all processes on the host but it is unique within the process namespace that the process exists within.\n",
                    "requirement_level": "recommended",
                    "root_namespace": "process",
                    "stability": "development",
                    "type": "int",
                },
                {
                    "brief": "The working directory of the process.\n",
                    "examples": [
                        "/root",
                    ],
                    "name": "process.working_directory",
                    "requirement_level": "recommended",
                    "root_namespace": "process",
                    "stability": "development",
                    "type": "string",
                },
            ],
            "root_namespace": "process",
        },
    },
    env: Environment {
        globals: {
            "concat_if": weaver_forge::extensions::util::concat_if,
            "cycler": minijinja_contrib::globals::cycler,
            "debug": minijinja::functions::builtins::debug,
            "dict": minijinja::functions::builtins::dict,
            "joiner": minijinja_contrib::globals::joiner,
            "namespace": minijinja::functions::builtins::namespace,
            "params": {
                "params": {},
            },
            "range": minijinja::functions::builtins::range,
            "template": {},
        },
        tests: [
            "!=",
            "<",
            "<=",
            "==",
            ">",
            ">=",
            "array",
            "boolean",
            "defined",
            "deprecated",
            "divisibleby",
            "endingwith",
            "enum",
            "enum_type",
            "eq",
            "equalto",
            "escaped",
            "even",
            "experimental",
            "false",
            "filter",
            "float",
            "ge",
            "greaterthan",
            "gt",
            "in",
            "int",
            "integer",
            "iterable",
            "le",
            "lessthan",
            "lower",
            "lt",
            "mapping",
            "ne",
            "none",
            "number",
            "odd",
            "safe",
            "sameas",
            "sequence",
            "simple_type",
            "stable",
            "startingwith",
            "string",
            "template_type",
            "test",
            "true",
            "undefined",
            "upper",
        ],
        filters: [
            "abs",
            "acronym",
            "ansi_bg_black",
            "ansi_bg_blue",
            "ansi_bg_bright_black",
            "ansi_bg_bright_blue",
            "ansi_bg_bright_cyan",
            "ansi_bg_bright_green",
            "ansi_bg_bright_magenta",
            "ansi_bg_bright_red",
            "ansi_bg_bright_white",
            "ansi_bg_bright_yellow",
            "ansi_bg_cyan",
            "ansi_bg_green",
            "ansi_bg_magenta",
            "ansi_bg_red",
            "ansi_bg_white",
            "ansi_bg_yellow",
            "ansi_black",
            "ansi_blue",
            "ansi_bold",
            "ansi_bright_black",
            "ansi_bright_blue",
            "ansi_bright_cyan",
            "ansi_bright_green",
            "ansi_bright_magenta",
            "ansi_bright_red",
            "ansi_bright_white",
            "ansi_bright_yellow",
            "ansi_cyan",
            "ansi_green",
            "ansi_italic",
            "ansi_magenta",
            "ansi_red",
            "ansi_strikethrough",
            "ansi_underline",
            "ansi_white",
            "ansi_yellow",
            "attr",
            "attribute_id",
            "attribute_namespace",
            "attribute_registry_file",
            "attribute_registry_namespace",
            "attribute_registry_title",
            "attribute_sort",
            "batch",
            "body_fields",
            "bool",
            "camel_case",
            "camel_case_const",
            "capitalize",
            "capitalize_first",
            "comment",
            "comment_with_prefix",
            "count",
            "d",
            "default",
            "dictsort",
            "e",
            "enum_type",
            "escape",
            "filesizeformat",
            "first",
            "flatten",
            "float",
            "groupby",
            "indent",
            "instantiated_type",
            "int",
            "items",
            "join",
            "kebab_case",
            "kebab_case_const",
            "last",
            "length",
            "lines",
            "list",
            "lower",
            "lower_case",
            "map",
            "map_text",
            "markdown_to_html",
            "max",
            "metric_namespace",
            "min",
            "not_required",
            "pascal_case",
            "pascal_case_const",
            "pluralize",
            "pprint",
            "print_member_value",
            "regex_replace",
            "reject",
            "rejectattr",
            "replace",
            "required",
            "reverse",
            "round",
            "safe",
            "screaming_kebab_case",
            "screaming_snake_case",
            "screaming_snake_case_const",
            "select",
            "selectattr",
            "slice",
            "snake_case",
            "snake_case_const",
            "sort",
            "split",
            "split_id",
            "string",
            "striptags",
            "sum",
            "title",
            "title_case",
            "tojson",
            "toyaml",
            "trim",
            "truncate",
            "type_mapping",
            "unique",
            "upper",
            "upper_case",
            "urlencode",
        ],
        templates: [
            "attr.go.j2",
        ],
    },
} */
