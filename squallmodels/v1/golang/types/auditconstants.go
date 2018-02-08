package types

import (
	"github.com/aporeto-inc/elemental"
)

// AuditFilterType are the valid IDs of the audit filters.
type AuditFilterType string

// Values of AuditFilterType.
const (
	AuditFilterTypeA0         AuditFilterType = "a0"
	AuditFilterTypeA1         AuditFilterType = "a1"
	AuditFilterTypeA2         AuditFilterType = "a2"
	AuditFilterTypeA3         AuditFilterType = "a3"
	AuditFilterTypeArch       AuditFilterType = "arch"
	AuditFilterTypeDevMajor   AuditFilterType = "devmajor"
	AuditFilterTypeDevMinor   AuditFilterType = "devminor"
	AuditFilterTypeDir        AuditFilterType = "dir"
	AuditFilterTypeEgid       AuditFilterType = "egid"
	AuditFilterTypeEuid       AuditFilterType = "euid"
	AuditFilterTypeExit       AuditFilterType = "exit"
	AuditFilterTypeFsgid      AuditFilterType = "fsgid"
	AuditFilterTypeFsuid      AuditFilterType = "fsuid"
	AuditFilterTypeFiletype   AuditFilterType = "filetye"
	AuditFilterTypeGid        AuditFilterType = "gid"
	AuditFilterTypeInode      AuditFilterType = "inode"
	AuditFilterTypeMsgtype    AuditFilterType = "msg_type"
	AuditFilterTypeObjuid     AuditFilterType = "obj_uid"
	AuditFilterTypeObjgid     AuditFilterType = "obj_gid"
	AuditFilterTypeObjuser    AuditFilterType = "obj_user"
	AuditFilterTypeObjrole    AuditFilterType = "obj_role"
	AuditFilterTypeObjtype    AuditFilterType = "obj_type"
	AuditFilterTypeObjlevlow  AuditFilterType = "obj_lev_low"
	AuditFilterTypeObjlevhigh AuditFilterType = "obj_lev_highj"
	AuditFilterTypePath       AuditFilterType = "path"
	AuditFilterTypePerm       AuditFilterType = "perm"
	AuditFilterTypePers       AuditFilterType = "pers"
	AuditFilterTypePid        AuditFilterType = "pid"
	AuditFilterTypePpid       AuditFilterType = "ppid"
	AuditFilterTypeSubuser    AuditFilterType = "subj_user"
	AuditFilterTypeSubjrole   AuditFilterType = "subj_role"
	AuditFilterTypeSubjtype   AuditFilterType = "subj_type"
	AuditFilterTypeSubsen     AuditFilterType = "subj_sen"
	AuditFilterTypeSubclr     AuditFilterType = "subj_clr"
	AuditFilterTypeSgid       AuditFilterType = "sgid"
	AuditFilterTypeSuccess    AuditFilterType = "success"
	AuditFilterTypeSuid       AuditFilterType = "suid"
	AuditFilterTypeUserid     AuditFilterType = "uid"
)

var auditFilterTypeReverse = map[string]interface{}{
	"a0":            nil,
	"a1":            nil,
	"a2":            nil,
	"a3":            nil,
	"arch":          nil,
	"devmajor":      nil,
	"devminor":      nil,
	"dir":           nil,
	"egid":          nil,
	"euid":          nil,
	"exit":          nil,
	"fsgid":         nil,
	"fsuid":         nil,
	"filetye":       nil,
	"gid":           nil,
	"inode":         nil,
	"msg_type":      nil,
	"obj_uid":       nil,
	"obj_gid":       nil,
	"obj_user":      nil,
	"obj_role":      nil,
	"obj_type":      nil,
	"obj_lev_low":   nil,
	"obj_lev_highj": nil,
	"path":          nil,
	"perm":          nil,
	"pers":          nil,
	"pid":           nil,
	"ppid":          nil,
	"subj_user":     nil,
	"subj_role":     nil,
	"subj_type":     nil,
	"subj_sen":      nil,
	"subj_clr":      nil,
	"sgid":          nil,
	"success":       nil,
	"suid":          nil,
	"uid":           nil,
}

// Validate validates the AuditFilterType
func (a AuditFilterType) Validate(attribute string) error {
	return elemental.ValidateStringInMap(attribute, string(a), auditFilterTypeReverse, false)
}

// AuditFilterListType are the types allowed in the list argument of audit filters
type AuditFilterListType string

// Values of AuditFilterListType
const (
	AuditFilterListTypeTask    AuditFilterListType = "task"
	AuditFilterListTypeExit    AuditFilterListType = "exit"
	AuditFilterListTypeUser    AuditFilterListType = "user"
	AuditFilterListTypeExclude AuditFilterListType = "exclude"
)

var auditFilterListTypeReverse = map[string]interface{}{
	"task":    nil,
	"exit":    nil,
	"user":    nil,
	"exclude": nil,
}

// Validate validates the AuditFilterListType
func (a AuditFilterListType) Validate(attribute string) error {
	return elemental.ValidateStringInMap(attribute, string(a), auditFilterListTypeReverse, false)
}

// AuditFilterActionType are the types allowed in the audit filter action
type AuditFilterActionType string

// Values of AuditFilterActionType
const (
	AuditFilterActionTypeNever  AuditFilterActionType = "never"
	AuditFilterActionTypeAlways AuditFilterActionType = "always"
)

// Validate validates the AuditFilterActionType
func (a AuditFilterActionType) Validate(attribute string) error {
	return elemental.ValidateStringInList(attribute, string(a), []string{string(AuditFilterActionTypeNever), string(AuditFilterActionTypeAlways)}, false)
}

// AuditFilterOperator is the operator for filters.
type AuditFilterOperator string

// Values of AuditFilterOperator
const (
	AuditFilterOperatorEqual           AuditFilterOperator = "="
	AuditFilterOperatorNotEqual        AuditFilterOperator = "!="
	AuditFilterOperatorGreater         AuditFilterOperator = ">"
	AuditFilterOperatorLessThan        AuditFilterOperator = "<"
	AuditFilterOperatorGreaterOrEqual  AuditFilterOperator = ">="
	AuditFilterOperatorLessThanOrEqual AuditFilterOperator = "<="
	AuditFilterOperatorBitMask         AuditFilterOperator = "&"
	AuditFilterOperatorBitTest         AuditFilterOperator = "&="
)

var auditFilterOperatorReverse = map[string]interface{}{
	"=":  nil,
	"!=": nil,
	">":  nil,
	"<":  nil,
	">=": nil,
	"<=": nil,
	"&":  nil,
	"&=": nil,
}

// Validate validates the audit filter operator
func (a AuditFilterOperator) Validate(attribute string) error {
	return elemental.ValidateStringInMap(attribute, string(a), auditFilterOperatorReverse, false)
}

// AuditFilePermissions is the type of file permissions
type AuditFilePermissions string

// Values of AuditFilePermissions
const (
	AuditFilePermissionsWrite     AuditFilePermissions = "w"
	AuditFilePermissionsRead      AuditFilePermissions = "r"
	AuditFilePermissionsExecute   AuditFilePermissions = "x"
	AuditFilePermissionsAttribute AuditFilePermissions = "a"
)

// Validate validates the audit file permissions
func (a AuditFilePermissions) Validate(attribute string) error {
	enums := []string{
		string(AuditFilePermissionsWrite),
		string(AuditFilePermissionsRead),
		string(AuditFilePermissionsExecute),
		string(AuditFilePermissionsAttribute),
	}

	return elemental.ValidateStringInList(attribute, string(a), enums, false)
}

// AuditSystemCallType is the type for the system calls.
type AuditSystemCallType string

// Validate validates the AuditSystemCallType
func (a AuditSystemCallType) Validate(attribute string) error {
	return elemental.ValidateStringInMap(attribute, string(a), auditSystemCallTypeReverse, false)
}

// Values of AuditSystemCallType.
const (
	AuditSystemCallTypeREAD                   AuditSystemCallType = "read"
	AuditSystemCallTypeWRITE                  AuditSystemCallType = "write"
	AuditSystemCallTypeOPEN                   AuditSystemCallType = "open"
	AuditSystemCallTypeCLOSE                  AuditSystemCallType = "close"
	AuditSystemCallTypeSTAT                   AuditSystemCallType = "stat"
	AuditSystemCallTypeFSTAT                  AuditSystemCallType = "fstat"
	AuditSystemCallTypeLSTAT                  AuditSystemCallType = "lstat"
	AuditSystemCallTypePOLL                   AuditSystemCallType = "poll"
	AuditSystemCallTypeLSEEK                  AuditSystemCallType = "lseek"
	AuditSystemCallTypeMMAP                   AuditSystemCallType = "mmap"
	AuditSystemCallTypeMPROTECT               AuditSystemCallType = "mprotect"
	AuditSystemCallTypeMUNMAP                 AuditSystemCallType = "munmap"
	AuditSystemCallTypeBRK                    AuditSystemCallType = "brk"
	AuditSystemCallTypeRT_SIGACTION           AuditSystemCallType = "rt_sigaction"
	AuditSystemCallTypeRT_SIGPROCMASK         AuditSystemCallType = "rt_sigprocmask"
	AuditSystemCallTypeRT_SIGRETURN           AuditSystemCallType = "rt_sigreturn"
	AuditSystemCallTypeIOCTL                  AuditSystemCallType = "ioctl"
	AuditSystemCallTypePREAD64                AuditSystemCallType = "pread64"
	AuditSystemCallTypePWRITE64               AuditSystemCallType = "pwrite64"
	AuditSystemCallTypeREADV                  AuditSystemCallType = "readv"
	AuditSystemCallTypeWRITEV                 AuditSystemCallType = "writev"
	AuditSystemCallTypeACCESS                 AuditSystemCallType = "access"
	AuditSystemCallTypePIPE                   AuditSystemCallType = "pipe"
	AuditSystemCallTypeSELECT                 AuditSystemCallType = "select"
	AuditSystemCallTypeSCHED_YIELD            AuditSystemCallType = "sched_yield"
	AuditSystemCallTypeMREMAP                 AuditSystemCallType = "mremap"
	AuditSystemCallTypeMSYNC                  AuditSystemCallType = "msync"
	AuditSystemCallTypeMINCORE                AuditSystemCallType = "mincore"
	AuditSystemCallTypeMADVISE                AuditSystemCallType = "madvise"
	AuditSystemCallTypeSHMGET                 AuditSystemCallType = "shmget"
	AuditSystemCallTypeSHMAT                  AuditSystemCallType = "shmat"
	AuditSystemCallTypeSHMCTL                 AuditSystemCallType = "shmctl"
	AuditSystemCallTypeDUP                    AuditSystemCallType = "dup"
	AuditSystemCallTypeDUP2                   AuditSystemCallType = "dup2"
	AuditSystemCallTypePAUSE                  AuditSystemCallType = "pause"
	AuditSystemCallTypeNANOSLEEP              AuditSystemCallType = "nanosleep"
	AuditSystemCallTypeGETITIMER              AuditSystemCallType = "getitimer"
	AuditSystemCallTypeALARM                  AuditSystemCallType = "alarm"
	AuditSystemCallTypeSETITIMER              AuditSystemCallType = "setitimer"
	AuditSystemCallTypeGETPID                 AuditSystemCallType = "getpid"
	AuditSystemCallTypeSENDFILE               AuditSystemCallType = "sendfile"
	AuditSystemCallTypeSOCKET                 AuditSystemCallType = "socket"
	AuditSystemCallTypeCONNECT                AuditSystemCallType = "connect"
	AuditSystemCallTypeACCEPT                 AuditSystemCallType = "accept"
	AuditSystemCallTypeSENDTO                 AuditSystemCallType = "sendto"
	AuditSystemCallTypeRECVFROM               AuditSystemCallType = "recvfrom"
	AuditSystemCallTypeSENDMSG                AuditSystemCallType = "sendmsg"
	AuditSystemCallTypeRECVMSG                AuditSystemCallType = "recvmsg"
	AuditSystemCallTypeSHUTDOWN               AuditSystemCallType = "shutdown"
	AuditSystemCallTypeBIND                   AuditSystemCallType = "bind"
	AuditSystemCallTypeLISTEN                 AuditSystemCallType = "listen"
	AuditSystemCallTypeGETSOCKNAME            AuditSystemCallType = "getsockname"
	AuditSystemCallTypeGETPEERNAME            AuditSystemCallType = "getpeername"
	AuditSystemCallTypeSOCKETPAIR             AuditSystemCallType = "socketpair"
	AuditSystemCallTypeSETSOCKOPT             AuditSystemCallType = "setsockopt"
	AuditSystemCallTypeGETSOCKOPT             AuditSystemCallType = "getsockopt"
	AuditSystemCallTypeCLONE                  AuditSystemCallType = "clone"
	AuditSystemCallTypeFORK                   AuditSystemCallType = "fork"
	AuditSystemCallTypeVFORK                  AuditSystemCallType = "vfork"
	AuditSystemCallTypeEXECVE                 AuditSystemCallType = "execve"
	AuditSystemCallTypeEXIT                   AuditSystemCallType = "exit"
	AuditSystemCallTypeWAIT4                  AuditSystemCallType = "wait4"
	AuditSystemCallTypeKILL                   AuditSystemCallType = "kill"
	AuditSystemCallTypeUNAME                  AuditSystemCallType = "uname"
	AuditSystemCallTypeSEMGET                 AuditSystemCallType = "semget"
	AuditSystemCallTypeSEMOP                  AuditSystemCallType = "semop"
	AuditSystemCallTypeSEMCTL                 AuditSystemCallType = "semctl"
	AuditSystemCallTypeSHMDT                  AuditSystemCallType = "shmdt"
	AuditSystemCallTypeMSGGET                 AuditSystemCallType = "msgget"
	AuditSystemCallTypeMSGSND                 AuditSystemCallType = "msgsnd"
	AuditSystemCallTypeMSGRCV                 AuditSystemCallType = "msgrcv"
	AuditSystemCallTypeMSGCTL                 AuditSystemCallType = "msgctl"
	AuditSystemCallTypeFCNTL                  AuditSystemCallType = "fcntl"
	AuditSystemCallTypeFLOCK                  AuditSystemCallType = "flock"
	AuditSystemCallTypeFSYNC                  AuditSystemCallType = "fsync"
	AuditSystemCallTypeFDATASYNC              AuditSystemCallType = "fdatasync"
	AuditSystemCallTypeTRUNCATE               AuditSystemCallType = "truncate"
	AuditSystemCallTypeFTRUNCATE              AuditSystemCallType = "ftruncate"
	AuditSystemCallTypeGETDENTS               AuditSystemCallType = "getdents"
	AuditSystemCallTypeGETCWD                 AuditSystemCallType = "getcwd"
	AuditSystemCallTypeCHDIR                  AuditSystemCallType = "chdir"
	AuditSystemCallTypeFCHDIR                 AuditSystemCallType = "fchdir"
	AuditSystemCallTypeRENAME                 AuditSystemCallType = "rename"
	AuditSystemCallTypeMKDIR                  AuditSystemCallType = "mkdir"
	AuditSystemCallTypeRMDIR                  AuditSystemCallType = "rmdir"
	AuditSystemCallTypeCREAT                  AuditSystemCallType = "creat"
	AuditSystemCallTypeLINK                   AuditSystemCallType = "link"
	AuditSystemCallTypeUNLINK                 AuditSystemCallType = "unlink"
	AuditSystemCallTypeSYMLINK                AuditSystemCallType = "symlink"
	AuditSystemCallTypeREADLINK               AuditSystemCallType = "readlink"
	AuditSystemCallTypeCHMOD                  AuditSystemCallType = "chmod"
	AuditSystemCallTypeFCHMOD                 AuditSystemCallType = "fchmod"
	AuditSystemCallTypeCHOWN                  AuditSystemCallType = "chown"
	AuditSystemCallTypeFCHOWN                 AuditSystemCallType = "fchown"
	AuditSystemCallTypeLCHOWN                 AuditSystemCallType = "lchown"
	AuditSystemCallTypeUMASK                  AuditSystemCallType = "umask"
	AuditSystemCallTypeGETTIMEOFDAY           AuditSystemCallType = "gettimeofday"
	AuditSystemCallTypeGETRLIMIT              AuditSystemCallType = "getrlimit"
	AuditSystemCallTypeGETRUSAGE              AuditSystemCallType = "getrusage"
	AuditSystemCallTypeSYSINFO                AuditSystemCallType = "sysinfo"
	AuditSystemCallTypeTIMES                  AuditSystemCallType = "times"
	AuditSystemCallTypePTRACE                 AuditSystemCallType = "ptrace"
	AuditSystemCallTypeGETUID                 AuditSystemCallType = "getuid"
	AuditSystemCallTypeSYSLOG                 AuditSystemCallType = "syslog"
	AuditSystemCallTypeGETGID                 AuditSystemCallType = "getgid"
	AuditSystemCallTypeSETUID                 AuditSystemCallType = "setuid"
	AuditSystemCallTypeSETGID                 AuditSystemCallType = "setgid"
	AuditSystemCallTypeGETEUID                AuditSystemCallType = "geteuid"
	AuditSystemCallTypeGETEGID                AuditSystemCallType = "getegid"
	AuditSystemCallTypeSETPGID                AuditSystemCallType = "setpgid"
	AuditSystemCallTypeGETPPID                AuditSystemCallType = "getppid"
	AuditSystemCallTypeGETPGRP                AuditSystemCallType = "getpgrp"
	AuditSystemCallTypeSETSID                 AuditSystemCallType = "setsid"
	AuditSystemCallTypeSETREUID               AuditSystemCallType = "setreuid"
	AuditSystemCallTypeSETREGID               AuditSystemCallType = "setregid"
	AuditSystemCallTypeGETGROUPS              AuditSystemCallType = "getgroups"
	AuditSystemCallTypeSETGROUPS              AuditSystemCallType = "setgroups"
	AuditSystemCallTypeSETRESUID              AuditSystemCallType = "setresuid"
	AuditSystemCallTypeGETRESUID              AuditSystemCallType = "getresuid"
	AuditSystemCallTypeSETRESGID              AuditSystemCallType = "setresgid"
	AuditSystemCallTypeGETRESGID              AuditSystemCallType = "getresgid"
	AuditSystemCallTypeGETPGID                AuditSystemCallType = "getpgid"
	AuditSystemCallTypeSETFSUID               AuditSystemCallType = "setfsuid"
	AuditSystemCallTypeSETFSGID               AuditSystemCallType = "setfsgid"
	AuditSystemCallTypeGETSID                 AuditSystemCallType = "getsid"
	AuditSystemCallTypeCAPGET                 AuditSystemCallType = "capget"
	AuditSystemCallTypeCAPSET                 AuditSystemCallType = "capset"
	AuditSystemCallTypeRT_SIGPENDING          AuditSystemCallType = "rt_sigpending"
	AuditSystemCallTypeRT_SIGTIMEDWAIT        AuditSystemCallType = "rt_sigtimedwait"
	AuditSystemCallTypeRT_SIGQUEUEINFO        AuditSystemCallType = "rt_sigqueueinfo"
	AuditSystemCallTypeRT_SIGSUSPEND          AuditSystemCallType = "rt_sigsuspend"
	AuditSystemCallTypeSIGALTSTACK            AuditSystemCallType = "sigaltstack"
	AuditSystemCallTypeUTIME                  AuditSystemCallType = "utime"
	AuditSystemCallTypeMKNOD                  AuditSystemCallType = "mknod"
	AuditSystemCallTypeUSELIB                 AuditSystemCallType = "uselib"
	AuditSystemCallTypePERSONALITY            AuditSystemCallType = "personality"
	AuditSystemCallTypeUSTAT                  AuditSystemCallType = "ustat"
	AuditSystemCallTypeSTATFS                 AuditSystemCallType = "statfs"
	AuditSystemCallTypeFSTATFS                AuditSystemCallType = "fstatfs"
	AuditSystemCallTypeSYSFS                  AuditSystemCallType = "sysfs"
	AuditSystemCallTypeGETPRIORITY            AuditSystemCallType = "getpriority"
	AuditSystemCallTypeSETPRIORITY            AuditSystemCallType = "setpriority"
	AuditSystemCallTypeSCHED_SETPARAM         AuditSystemCallType = "sched_setparam"
	AuditSystemCallTypeSCHED_GETPARAM         AuditSystemCallType = "sched_getparam"
	AuditSystemCallTypeSCHED_SETSCHEDULER     AuditSystemCallType = "sched_setscheduler"
	AuditSystemCallTypeSCHED_GETSCHEDULER     AuditSystemCallType = "sched_getscheduler"
	AuditSystemCallTypeSCHED_GET_PRIORITY_MAX AuditSystemCallType = "sched_get_priority_max"
	AuditSystemCallTypeSCHED_GET_PRIORITY_MIN AuditSystemCallType = "sched_get_priority_min"
	AuditSystemCallTypeSCHED_RR_GET_INTERVAL  AuditSystemCallType = "sched_rr_get_interval"
	AuditSystemCallTypeMLOCK                  AuditSystemCallType = "mlock"
	AuditSystemCallTypeMUNLOCK                AuditSystemCallType = "munlock"
	AuditSystemCallTypeMLOCKALL               AuditSystemCallType = "mlockall"
	AuditSystemCallTypeMUNLOCKALL             AuditSystemCallType = "munlockall"
	AuditSystemCallTypeVHANGUP                AuditSystemCallType = "vhangup"
	AuditSystemCallTypeMODIFY_LDT             AuditSystemCallType = "modify_ldt"
	AuditSystemCallTypePIVOT_ROOT             AuditSystemCallType = "pivot_root"
	AuditSystemCallType_SYSCTL                AuditSystemCallType = "_sysctl"
	AuditSystemCallTypePRCTL                  AuditSystemCallType = "prctl"
	AuditSystemCallTypeARCH_PRCTL             AuditSystemCallType = "arch_prctl"
	AuditSystemCallTypeADJTIMEX               AuditSystemCallType = "adjtimex"
	AuditSystemCallTypeSETRLIMIT              AuditSystemCallType = "setrlimit"
	AuditSystemCallTypeCHROOT                 AuditSystemCallType = "chroot"
	AuditSystemCallTypeSYNC                   AuditSystemCallType = "sync"
	AuditSystemCallTypeACCT                   AuditSystemCallType = "acct"
	AuditSystemCallTypeSETTIMEOFDAY           AuditSystemCallType = "settimeofday"
	AuditSystemCallTypeMOUNT                  AuditSystemCallType = "mount"
	AuditSystemCallTypeUMOUNT2                AuditSystemCallType = "umount2"
	AuditSystemCallTypeSWAPON                 AuditSystemCallType = "swapon"
	AuditSystemCallTypeSWAPOFF                AuditSystemCallType = "swapoff"
	AuditSystemCallTypeREBOOT                 AuditSystemCallType = "reboot"
	AuditSystemCallTypeSETHOSTNAME            AuditSystemCallType = "sethostname"
	AuditSystemCallTypeSETDOMAINNAME          AuditSystemCallType = "setdomainname"
	AuditSystemCallTypeIOPL                   AuditSystemCallType = "iopl"
	AuditSystemCallTypeIOPERM                 AuditSystemCallType = "ioperm"
	AuditSystemCallTypeCREATE_MODULE          AuditSystemCallType = "create_module"
	AuditSystemCallTypeINIT_MODULE            AuditSystemCallType = "init_module"
	AuditSystemCallTypeDELETE_MODULE          AuditSystemCallType = "delete_module"
	AuditSystemCallTypeGET_KERNEL_SYMS        AuditSystemCallType = "get_kernel_syms"
	AuditSystemCallTypeQUERY_MODULE           AuditSystemCallType = "query_module"
	AuditSystemCallTypeQUOTACTL               AuditSystemCallType = "quotactl"
	AuditSystemCallTypeNFSSERVCTL             AuditSystemCallType = "nfsservctl"
	AuditSystemCallTypeGETPMSG                AuditSystemCallType = "getpmsg"
	AuditSystemCallTypePUTPMSG                AuditSystemCallType = "putpmsg"
	AuditSystemCallTypeAFS_SYSCALL            AuditSystemCallType = "afs_syscall"
	AuditSystemCallTypeTUXCALL                AuditSystemCallType = "tuxcall"
	AuditSystemCallTypeSECURITY               AuditSystemCallType = "security"
	AuditSystemCallTypeGETTID                 AuditSystemCallType = "gettid"
	AuditSystemCallTypeREADAHEAD              AuditSystemCallType = "readahead"
	AuditSystemCallTypeSETXATTR               AuditSystemCallType = "setxattr"
	AuditSystemCallTypeLSETXATTR              AuditSystemCallType = "lsetxattr"
	AuditSystemCallTypeFSETXATTR              AuditSystemCallType = "fsetxattr"
	AuditSystemCallTypeGETXATTR               AuditSystemCallType = "getxattr"
	AuditSystemCallTypeLGETXATTR              AuditSystemCallType = "lgetxattr"
	AuditSystemCallTypeFGETXATTR              AuditSystemCallType = "fgetxattr"
	AuditSystemCallTypeLISTXATTR              AuditSystemCallType = "listxattr"
	AuditSystemCallTypeLLISTXATTR             AuditSystemCallType = "llistxattr"
	AuditSystemCallTypeFLISTXATTR             AuditSystemCallType = "flistxattr"
	AuditSystemCallTypeREMOVEXATTR            AuditSystemCallType = "removexattr"
	AuditSystemCallTypeLREMOVEXATTR           AuditSystemCallType = "lremovexattr"
	AuditSystemCallTypeFREMOVEXATTR           AuditSystemCallType = "fremovexattr"
	AuditSystemCallTypeTKILL                  AuditSystemCallType = "tkill"
	AuditSystemCallTypeTIME                   AuditSystemCallType = "time"
	AuditSystemCallTypeFUTEX                  AuditSystemCallType = "futex"
	AuditSystemCallTypeSCHED_SETAFFINITY      AuditSystemCallType = "sched_setaffinity"
	AuditSystemCallTypeSCHED_GETAFFINITY      AuditSystemCallType = "sched_getaffinity"
	AuditSystemCallTypeSET_THREAD_AREA        AuditSystemCallType = "set_thread_area"
	AuditSystemCallTypeIO_SETUP               AuditSystemCallType = "io_setup"
	AuditSystemCallTypeIO_DESTROY             AuditSystemCallType = "io_destroy"
	AuditSystemCallTypeIO_GETEVENTS           AuditSystemCallType = "io_getevents"
	AuditSystemCallTypeIO_SUBMIT              AuditSystemCallType = "io_submit"
	AuditSystemCallTypeIO_CANCEL              AuditSystemCallType = "io_cancel"
	AuditSystemCallTypeGET_THREAD_AREA        AuditSystemCallType = "get_thread_area"
	AuditSystemCallTypeLOOKUP_DCOOKIE         AuditSystemCallType = "lookup_dcookie"
	AuditSystemCallTypeEPOLL_CREATE           AuditSystemCallType = "epoll_create"
	AuditSystemCallTypeEPOLL_CTL_OLD          AuditSystemCallType = "epoll_ctl_old"
	AuditSystemCallTypeEPOLL_WAIT_OLD         AuditSystemCallType = "epoll_wait_old"
	AuditSystemCallTypeREMAP_FILE_PAGES       AuditSystemCallType = "remap_file_pages"
	AuditSystemCallTypeGETDENTS64             AuditSystemCallType = "getdents64"
	AuditSystemCallTypeSET_TID_ADDRESS        AuditSystemCallType = "set_tid_address"
	AuditSystemCallTypeRESTART_SYSCALL        AuditSystemCallType = "restart_syscall"
	AuditSystemCallTypeSEMTIMEDOP             AuditSystemCallType = "semtimedop"
	AuditSystemCallTypeFADVISE64              AuditSystemCallType = "fadvise64"
	AuditSystemCallTypeTIMER_CREATE           AuditSystemCallType = "timer_create"
	AuditSystemCallTypeTIMER_SETTIME          AuditSystemCallType = "timer_settime"
	AuditSystemCallTypeTIMER_GETTIME          AuditSystemCallType = "timer_gettime"
	AuditSystemCallTypeTIMER_GETOVERRUN       AuditSystemCallType = "timer_getoverrun"
	AuditSystemCallTypeTIMER_DELETE           AuditSystemCallType = "timer_delete"
	AuditSystemCallTypeCLOCK_SETTIME          AuditSystemCallType = "clock_settime"
	AuditSystemCallTypeCLOCK_GETTIME          AuditSystemCallType = "clock_gettime"
	AuditSystemCallTypeCLOCK_GETRES           AuditSystemCallType = "clock_getres"
	AuditSystemCallTypeCLOCK_NANOSLEEP        AuditSystemCallType = "clock_nanosleep"
	AuditSystemCallTypeEXIT_GROUP             AuditSystemCallType = "exit_group"
	AuditSystemCallTypeEPOLL_WAIT             AuditSystemCallType = "epoll_wait"
	AuditSystemCallTypeEPOLL_CTL              AuditSystemCallType = "epoll_ctl"
	AuditSystemCallTypeTGKILL                 AuditSystemCallType = "tgkill"
	AuditSystemCallTypeUTIMES                 AuditSystemCallType = "utimes"
	AuditSystemCallTypeVSERVER                AuditSystemCallType = "vserver"
	AuditSystemCallTypeMBIND                  AuditSystemCallType = "mbind"
	AuditSystemCallTypeSET_MEMPOLICY          AuditSystemCallType = "set_mempolicy"
	AuditSystemCallTypeGET_MEMPOLICY          AuditSystemCallType = "get_mempolicy"
	AuditSystemCallTypeMQ_OPEN                AuditSystemCallType = "mq_open"
	AuditSystemCallTypeMQ_UNLINK              AuditSystemCallType = "mq_unlink"
	AuditSystemCallTypeMQ_TIMEDSEND           AuditSystemCallType = "mq_timedsend"
	AuditSystemCallTypeMQ_TIMEDRECEIVE        AuditSystemCallType = "mq_timedreceive"
	AuditSystemCallTypeMQ_NOTIFY              AuditSystemCallType = "mq_notify"
	AuditSystemCallTypeMQ_GETSETATTR          AuditSystemCallType = "mq_getsetattr"
	AuditSystemCallTypeKEXEC_LOAD             AuditSystemCallType = "kexec_load"
	AuditSystemCallTypeWAITID                 AuditSystemCallType = "waitid"
	AuditSystemCallTypeADD_KEY                AuditSystemCallType = "add_key"
	AuditSystemCallTypeREQUEST_KEY            AuditSystemCallType = "request_key"
	AuditSystemCallTypeKEYCTL                 AuditSystemCallType = "keyctl"
	AuditSystemCallTypeIOPRIO_SET             AuditSystemCallType = "ioprio_set"
	AuditSystemCallTypeIOPRIO_GET             AuditSystemCallType = "ioprio_get"
	AuditSystemCallTypeINOTIFY_INIT           AuditSystemCallType = "inotify_init"
	AuditSystemCallTypeINOTIFY_ADD_WATCH      AuditSystemCallType = "inotify_add_watch"
	AuditSystemCallTypeINOTIFY_RM_WATCH       AuditSystemCallType = "inotify_rm_watch"
	AuditSystemCallTypeMIGRATE_PAGES          AuditSystemCallType = "migrate_pages"
	AuditSystemCallTypeOPENAT                 AuditSystemCallType = "openat"
	AuditSystemCallTypeMKDIRAT                AuditSystemCallType = "mkdirat"
	AuditSystemCallTypeMKNODAT                AuditSystemCallType = "mknodat"
	AuditSystemCallTypeFCHOWNAT               AuditSystemCallType = "fchownat"
	AuditSystemCallTypeFUTIMESAT              AuditSystemCallType = "futimesat"
	AuditSystemCallTypeNEWFSTATAT             AuditSystemCallType = "newfstatat"
	AuditSystemCallTypeUNLINKAT               AuditSystemCallType = "unlinkat"
	AuditSystemCallTypeRENAMEAT               AuditSystemCallType = "renameat"
	AuditSystemCallTypeLINKAT                 AuditSystemCallType = "linkat"
	AuditSystemCallTypeSYMLINKAT              AuditSystemCallType = "symlinkat"
	AuditSystemCallTypeREADLINKAT             AuditSystemCallType = "readlinkat"
	AuditSystemCallTypeFCHMODAT               AuditSystemCallType = "fchmodat"
	AuditSystemCallTypeFACCESSAT              AuditSystemCallType = "faccessat"
	AuditSystemCallTypePSELECT6               AuditSystemCallType = "pselect6"
	AuditSystemCallTypePPOLL                  AuditSystemCallType = "ppoll"
	AuditSystemCallTypeUNSHARE                AuditSystemCallType = "unshare"
	AuditSystemCallTypeSET_ROBUST_LIST        AuditSystemCallType = "set_robust_list"
	AuditSystemCallTypeGET_ROBUST_LIST        AuditSystemCallType = "get_robust_list"
	AuditSystemCallTypeSPLICE                 AuditSystemCallType = "splice"
	AuditSystemCallTypeTEE                    AuditSystemCallType = "tee"
	AuditSystemCallTypeSYNC_FILE_RANGE        AuditSystemCallType = "sync_file_range"
	AuditSystemCallTypeVMSPLICE               AuditSystemCallType = "vmsplice"
	AuditSystemCallTypeMOVE_PAGES             AuditSystemCallType = "move_pages"
	AuditSystemCallTypeUTIMENSAT              AuditSystemCallType = "utimensat"
	AuditSystemCallTypeEPOLL_PWAIT            AuditSystemCallType = "epoll_pwait"
	AuditSystemCallTypeSIGNALFD               AuditSystemCallType = "signalfd"
	AuditSystemCallTypeTIMERFD_CREATE         AuditSystemCallType = "timerfd_create"
	AuditSystemCallTypeEVENTFD                AuditSystemCallType = "eventfd"
	AuditSystemCallTypeFALLOCATE              AuditSystemCallType = "fallocate"
	AuditSystemCallTypeTIMERFD_SETTIME        AuditSystemCallType = "timerfd_settime"
	AuditSystemCallTypeTIMERFD_GETTIME        AuditSystemCallType = "timerfd_gettime"
	AuditSystemCallTypeACCEPT4                AuditSystemCallType = "accept4"
	AuditSystemCallTypeSIGNALFD4              AuditSystemCallType = "signalfd4"
	AuditSystemCallTypeEVENTFD2               AuditSystemCallType = "eventfd2"
	AuditSystemCallTypeEPOLL_CREATE1          AuditSystemCallType = "epoll_create1"
	AuditSystemCallTypeDUP3                   AuditSystemCallType = "dup3"
	AuditSystemCallTypePIPE2                  AuditSystemCallType = "pipe2"
	AuditSystemCallTypeINOTIFY_INIT1          AuditSystemCallType = "inotify_init1"
	AuditSystemCallTypePREADV                 AuditSystemCallType = "preadv"
	AuditSystemCallTypePWRITEV                AuditSystemCallType = "pwritev"
	AuditSystemCallTypeRT_TGSIGQUEUEINFO      AuditSystemCallType = "rt_tgsigqueueinfo"
	AuditSystemCallTypePERF_EVENT_OPEN        AuditSystemCallType = "perf_event_open"
	AuditSystemCallTypeRECVMMSG               AuditSystemCallType = "recvmmsg"
	AuditSystemCallTypeFANOTIFY_INIT          AuditSystemCallType = "fanotify_init"
	AuditSystemCallTypeFANOTIFY_MARK          AuditSystemCallType = "fanotify_mark"
	AuditSystemCallTypePRLIMIT64              AuditSystemCallType = "prlimit64"
	AuditSystemCallTypeNAME_TO_HANDLE_AT      AuditSystemCallType = "name_to_handle_at"
	AuditSystemCallTypeOPEN_BY_HANDLE_AT      AuditSystemCallType = "open_by_handle_at"
	AuditSystemCallTypeCLOCK_ADJTIME          AuditSystemCallType = "clock_adjtime"
	AuditSystemCallTypeSYNCFS                 AuditSystemCallType = "syncfs"
	AuditSystemCallTypeSENDMMSG               AuditSystemCallType = "sendmmsg"
	AuditSystemCallTypeSETNS                  AuditSystemCallType = "setns"
	AuditSystemCallTypeGETCPU                 AuditSystemCallType = "getcpu"
	AuditSystemCallTypePROCESS_VM_READV       AuditSystemCallType = "process_vm_readv"
	AuditSystemCallTypePROCESS_VM_WRITEV      AuditSystemCallType = "process_vm_writev"
	AuditSystemCallTypeKCMP                   AuditSystemCallType = "kcmp"
	AuditSystemCallTypeFINIT_MODULE           AuditSystemCallType = "finit_module"
)

var auditSystemCallTypeReverse = map[string]interface{}{
	"read":                   nil,
	"write":                  nil,
	"open":                   nil,
	"close":                  nil,
	"stat":                   nil,
	"fstat":                  nil,
	"lstat":                  nil,
	"poll":                   nil,
	"lseek":                  nil,
	"mmap":                   nil,
	"mprotect":               nil,
	"munmap":                 nil,
	"brk":                    nil,
	"rt_sigaction":           nil,
	"rt_sigprocmask":         nil,
	"rt_sigreturn":           nil,
	"ioctl":                  nil,
	"pread64":                nil,
	"pwrite64":               nil,
	"readv":                  nil,
	"writev":                 nil,
	"access":                 nil,
	"pipe":                   nil,
	"select":                 nil,
	"sched_yield":            nil,
	"mremap":                 nil,
	"msync":                  nil,
	"mincore":                nil,
	"madvise":                nil,
	"shmget":                 nil,
	"shmat":                  nil,
	"shmctl":                 nil,
	"dup":                    nil,
	"dup2":                   nil,
	"pause":                  nil,
	"nanosleep":              nil,
	"getitimer":              nil,
	"alarm":                  nil,
	"setitimer":              nil,
	"getpid":                 nil,
	"sendfile":               nil,
	"socket":                 nil,
	"connect":                nil,
	"accept":                 nil,
	"sendto":                 nil,
	"recvfrom":               nil,
	"sendmsg":                nil,
	"recvmsg":                nil,
	"shutdown":               nil,
	"bind":                   nil,
	"listen":                 nil,
	"getsockname":            nil,
	"getpeername":            nil,
	"socketpair":             nil,
	"setsockopt":             nil,
	"getsockopt":             nil,
	"clone":                  nil,
	"fork":                   nil,
	"vfork":                  nil,
	"execve":                 nil,
	"exit":                   nil,
	"wait4":                  nil,
	"kill":                   nil,
	"uname":                  nil,
	"semget":                 nil,
	"semop":                  nil,
	"semctl":                 nil,
	"shmdt":                  nil,
	"msgget":                 nil,
	"msgsnd":                 nil,
	"msgrcv":                 nil,
	"msgctl":                 nil,
	"fcntl":                  nil,
	"flock":                  nil,
	"fsync":                  nil,
	"fdatasync":              nil,
	"truncate":               nil,
	"ftruncate":              nil,
	"getdents":               nil,
	"getcwd":                 nil,
	"chdir":                  nil,
	"fchdir":                 nil,
	"rename":                 nil,
	"mkdir":                  nil,
	"rmdir":                  nil,
	"creat":                  nil,
	"link":                   nil,
	"unlink":                 nil,
	"symlink":                nil,
	"readlink":               nil,
	"chmod":                  nil,
	"fchmod":                 nil,
	"chown":                  nil,
	"fchown":                 nil,
	"lchown":                 nil,
	"umask":                  nil,
	"gettimeofday":           nil,
	"getrlimit":              nil,
	"getrusage":              nil,
	"sysinfo":                nil,
	"times":                  nil,
	"ptrace":                 nil,
	"getuid":                 nil,
	"syslog":                 nil,
	"getgid":                 nil,
	"setuid":                 nil,
	"setgid":                 nil,
	"geteuid":                nil,
	"getegid":                nil,
	"setpgid":                nil,
	"getppid":                nil,
	"getpgrp":                nil,
	"setsid":                 nil,
	"setreuid":               nil,
	"setregid":               nil,
	"getgroups":              nil,
	"setgroups":              nil,
	"setresuid":              nil,
	"getresuid":              nil,
	"setresgid":              nil,
	"getresgid":              nil,
	"getpgid":                nil,
	"setfsuid":               nil,
	"setfsgid":               nil,
	"getsid":                 nil,
	"capget":                 nil,
	"capset":                 nil,
	"rt_sigpending":          nil,
	"rt_sigtimedwait":        nil,
	"rt_sigqueueinfo":        nil,
	"rt_sigsuspend":          nil,
	"sigaltstack":            nil,
	"utime":                  nil,
	"mknod":                  nil,
	"uselib":                 nil,
	"personality":            nil,
	"ustat":                  nil,
	"statfs":                 nil,
	"fstatfs":                nil,
	"sysfs":                  nil,
	"getpriority":            nil,
	"setpriority":            nil,
	"sched_setparam":         nil,
	"sched_getparam":         nil,
	"sched_setscheduler":     nil,
	"sched_getscheduler":     nil,
	"sched_get_priority_max": nil,
	"sched_get_priority_min": nil,
	"sched_rr_get_interval":  nil,
	"mlock":                  nil,
	"munlock":                nil,
	"mlockall":               nil,
	"munlockall":             nil,
	"vhangup":                nil,
	"modify_ldt":             nil,
	"pivot_root":             nil,
	"_sysctl":                nil,
	"prctl":                  nil,
	"arch_prctl":             nil,
	"adjtimex":               nil,
	"setrlimit":              nil,
	"chroot":                 nil,
	"sync":                   nil,
	"acct":                   nil,
	"settimeofday":           nil,
	"mount":                  nil,
	"umount2":                nil,
	"swapon":                 nil,
	"swapoff":                nil,
	"reboot":                 nil,
	"sethostname":            nil,
	"setdomainname":          nil,
	"iopl":                   nil,
	"ioperm":                 nil,
	"create_module":          nil,
	"init_module":            nil,
	"delete_module":          nil,
	"get_kernel_syms":        nil,
	"query_module":           nil,
	"quotactl":               nil,
	"nfsservctl":             nil,
	"getpmsg":                nil,
	"putpmsg":                nil,
	"afs_syscall":            nil,
	"tuxcall":                nil,
	"security":               nil,
	"gettid":                 nil,
	"readahead":              nil,
	"setxattr":               nil,
	"lsetxattr":              nil,
	"fsetxattr":              nil,
	"getxattr":               nil,
	"lgetxattr":              nil,
	"fgetxattr":              nil,
	"listxattr":              nil,
	"llistxattr":             nil,
	"flistxattr":             nil,
	"removexattr":            nil,
	"lremovexattr":           nil,
	"fremovexattr":           nil,
	"tkill":                  nil,
	"time":                   nil,
	"futex":                  nil,
	"sched_setaffinity":      nil,
	"sched_getaffinity":      nil,
	"set_thread_area":        nil,
	"io_setup":               nil,
	"io_destroy":             nil,
	"io_getevents":           nil,
	"io_submit":              nil,
	"io_cancel":              nil,
	"get_thread_area":        nil,
	"lookup_dcookie":         nil,
	"epoll_create":           nil,
	"epoll_ctl_old":          nil,
	"epoll_wait_old":         nil,
	"remap_file_pages":       nil,
	"getdents64":             nil,
	"set_tid_address":        nil,
	"restart_syscall":        nil,
	"semtimedop":             nil,
	"fadvise64":              nil,
	"timer_create":           nil,
	"timer_settime":          nil,
	"timer_gettime":          nil,
	"timer_getoverrun":       nil,
	"timer_delete":           nil,
	"clock_settime":          nil,
	"clock_gettime":          nil,
	"clock_getres":           nil,
	"clock_nanosleep":        nil,
	"exit_group":             nil,
	"epoll_wait":             nil,
	"epoll_ctl":              nil,
	"tgkill":                 nil,
	"utimes":                 nil,
	"vserver":                nil,
	"mbind":                  nil,
	"set_mempolicy":          nil,
	"get_mempolicy":          nil,
	"mq_open":                nil,
	"mq_unlink":              nil,
	"mq_timedsend":           nil,
	"mq_timedreceive":        nil,
	"mq_notify":              nil,
	"mq_getsetattr":          nil,
	"kexec_load":             nil,
	"waitid":                 nil,
	"add_key":                nil,
	"request_key":            nil,
	"keyctl":                 nil,
	"ioprio_set":             nil,
	"ioprio_get":             nil,
	"inotify_init":           nil,
	"inotify_add_watch":      nil,
	"inotify_rm_watch":       nil,
	"migrate_pages":          nil,
	"openat":                 nil,
	"mkdirat":                nil,
	"mknodat":                nil,
	"fchownat":               nil,
	"futimesat":              nil,
	"newfstatat":             nil,
	"unlinkat":               nil,
	"renameat":               nil,
	"linkat":                 nil,
	"symlinkat":              nil,
	"readlinkat":             nil,
	"fchmodat":               nil,
	"faccessat":              nil,
	"pselect6":               nil,
	"ppoll":                  nil,
	"unshare":                nil,
	"set_robust_list":        nil,
	"get_robust_list":        nil,
	"splice":                 nil,
	"tee":                    nil,
	"sync_file_range":        nil,
	"vmsplice":               nil,
	"move_pages":             nil,
	"utimensat":              nil,
	"epoll_pwait":            nil,
	"signalfd":               nil,
	"timerfd_create":         nil,
	"eventfd":                nil,
	"fallocate":              nil,
	"timerfd_settime":        nil,
	"timerfd_gettime":        nil,
	"accept4":                nil,
	"signalfd4":              nil,
	"eventfd2":               nil,
	"epoll_create1":          nil,
	"dup3":                   nil,
	"pipe2":                  nil,
	"inotify_init1":          nil,
	"preadv":                 nil,
	"pwritev":                nil,
	"rt_tgsigqueueinfo":      nil,
	"perf_event_open":        nil,
	"recvmmsg":               nil,
	"fanotify_init":          nil,
	"fanotify_mark":          nil,
	"prlimit64":              nil,
	"name_to_handle_at":      nil,
	"open_by_handle_at":      nil,
	"clock_adjtime":          nil,
	"syncfs":                 nil,
	"sendmmsg":               nil,
	"setns":                  nil,
	"getcpu":                 nil,
	"process_vm_readv":       nil,
	"process_vm_writev":      nil,
	"kcmp":                   nil,
	"finit_module":           nil,
}