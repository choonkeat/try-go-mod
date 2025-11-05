```
.
├── go.mod
├── go.sum
├── main.go
├── README.md
└── vendor
    ├── github.com
    │   ├── choonkeat
    │   │   └── try-go-util
    │   │       └── pkg
    │   │           └── logutil
    │   │               ├── basic.go
    │   │               └── logrusutil
    │   │                   └── lib.go
    │   └── sirupsen
    │       └── logrus
    │           ├── alt_exit.go
    │           ├── appveyor.yml
    │           ├── buffer_pool.go
    │           ├── CHANGELOG.md
    │           ├── doc.go
    │           ├── entry.go
    │           ├── exported.go
    │           ├── formatter.go
    │           ├── hooks.go
    │           ├── json_formatter.go
    │           ├── LICENSE
    │           ├── logger.go
    │           ├── logrus.go
    │           ├── README.md
    │           ├── terminal_check_appengine.go
    │           ├── terminal_check_bsd.go
    │           ├── terminal_check_js.go
    │           ├── terminal_check_no_terminal.go
    │           ├── terminal_check_notappengine.go
    │           ├── terminal_check_solaris.go
    │           ├── terminal_check_unix.go
    │           ├── terminal_check_windows.go
    │           ├── text_formatter.go
    │           └── writer.go
    ├── golang.org
    │   └── x
    │       └── sys
    │           ├── AUTHORS
    │           ├── CONTRIBUTORS
    │           ├── internal
    │           │   └── unsafeheader
    │           │       └── unsafeheader.go
    │           ├── LICENSE
    │           ├── PATENTS
    │           ├── unix
    │           │   ├── affinity_linux.go
    │           │   ├── aliases.go
    │           │   ├── asm_aix_ppc64.s
    │           │   ├── asm_bsd_386.s
    │           │   ├── asm_bsd_amd64.s
    │           │   ├── asm_bsd_arm.s
    │           │   ├── asm_bsd_arm64.s
    │           │   ├── asm_bsd_riscv64.s
    │           │   ├── asm_linux_386.s
    │           │   ├── asm_linux_amd64.s
    │           │   ├── asm_linux_arm.s
    │           │   ├── asm_linux_arm64.s
    │           │   ├── asm_linux_loong64.s
    │           │   ├── asm_linux_mips64x.s
    │           │   ├── asm_linux_mipsx.s
    │           │   ├── asm_linux_ppc64x.s
    │           │   ├── asm_linux_riscv64.s
    │           │   ├── asm_linux_s390x.s
    │           │   ├── asm_openbsd_mips64.s
    │           │   ├── asm_solaris_amd64.s
    │           │   ├── asm_zos_s390x.s
    │           │   ├── bluetooth_linux.go
    │           │   ├── cap_freebsd.go
    │           │   ├── constants.go
    │           │   ├── dev_aix_ppc.go
    │           │   ├── dev_aix_ppc64.go
    │           │   ├── dev_darwin.go
    │           │   ├── dev_dragonfly.go
    │           │   ├── dev_freebsd.go
    │           │   ├── dev_linux.go
    │           │   ├── dev_netbsd.go
    │           │   ├── dev_openbsd.go
    │           │   ├── dev_zos.go
    │           │   ├── dirent.go
    │           │   ├── endian_big.go
    │           │   ├── endian_little.go
    │           │   ├── env_unix.go
    │           │   ├── epoll_zos.go
    │           │   ├── fcntl_darwin.go
    │           │   ├── fcntl_linux_32bit.go
    │           │   ├── fcntl.go
    │           │   ├── fdset.go
    │           │   ├── fstatfs_zos.go
    │           │   ├── gccgo_c.c
    │           │   ├── gccgo_linux_amd64.go
    │           │   ├── gccgo.go
    │           │   ├── ifreq_linux.go
    │           │   ├── ioctl_linux.go
    │           │   ├── ioctl_zos.go
    │           │   ├── ioctl.go
    │           │   ├── mkall.sh
    │           │   ├── mkerrors.sh
    │           │   ├── pagesize_unix.go
    │           │   ├── pledge_openbsd.go
    │           │   ├── ptrace_darwin.go
    │           │   ├── ptrace_ios.go
    │           │   ├── race.go
    │           │   ├── race0.go
    │           │   ├── readdirent_getdents.go
    │           │   ├── readdirent_getdirentries.go
    │           │   ├── README.md
    │           │   ├── sockcmsg_dragonfly.go
    │           │   ├── sockcmsg_linux.go
    │           │   ├── sockcmsg_unix_other.go
    │           │   ├── sockcmsg_unix.go
    │           │   ├── str.go
    │           │   ├── syscall_aix_ppc.go
    │           │   ├── syscall_aix_ppc64.go
    │           │   ├── syscall_aix.go
    │           │   ├── syscall_bsd.go
    │           │   ├── syscall_darwin_amd64.go
    │           │   ├── syscall_darwin_arm64.go
    │           │   ├── syscall_darwin_libSystem.go
    │           │   ├── syscall_darwin.1_12.go
    │           │   ├── syscall_darwin.1_13.go
    │           │   ├── syscall_darwin.go
    │           │   ├── syscall_dragonfly_amd64.go
    │           │   ├── syscall_dragonfly.go
    │           │   ├── syscall_freebsd_386.go
    │           │   ├── syscall_freebsd_amd64.go
    │           │   ├── syscall_freebsd_arm.go
    │           │   ├── syscall_freebsd_arm64.go
    │           │   ├── syscall_freebsd_riscv64.go
    │           │   ├── syscall_freebsd.go
    │           │   ├── syscall_illumos.go
    │           │   ├── syscall_linux_386.go
    │           │   ├── syscall_linux_alarm.go
    │           │   ├── syscall_linux_amd64_gc.go
    │           │   ├── syscall_linux_amd64.go
    │           │   ├── syscall_linux_arm.go
    │           │   ├── syscall_linux_arm64.go
    │           │   ├── syscall_linux_gc_386.go
    │           │   ├── syscall_linux_gc_arm.go
    │           │   ├── syscall_linux_gc.go
    │           │   ├── syscall_linux_gccgo_386.go
    │           │   ├── syscall_linux_gccgo_arm.go
    │           │   ├── syscall_linux_loong64.go
    │           │   ├── syscall_linux_mips64x.go
    │           │   ├── syscall_linux_mipsx.go
    │           │   ├── syscall_linux_ppc.go
    │           │   ├── syscall_linux_ppc64x.go
    │           │   ├── syscall_linux_riscv64.go
    │           │   ├── syscall_linux_s390x.go
    │           │   ├── syscall_linux_sparc64.go
    │           │   ├── syscall_linux.go
    │           │   ├── syscall_netbsd_386.go
    │           │   ├── syscall_netbsd_amd64.go
    │           │   ├── syscall_netbsd_arm.go
    │           │   ├── syscall_netbsd_arm64.go
    │           │   ├── syscall_netbsd.go
    │           │   ├── syscall_openbsd_386.go
    │           │   ├── syscall_openbsd_amd64.go
    │           │   ├── syscall_openbsd_arm.go
    │           │   ├── syscall_openbsd_arm64.go
    │           │   ├── syscall_openbsd_mips64.go
    │           │   ├── syscall_openbsd.go
    │           │   ├── syscall_solaris_amd64.go
    │           │   ├── syscall_solaris.go
    │           │   ├── syscall_unix_gc_ppc64x.go
    │           │   ├── syscall_unix_gc.go
    │           │   ├── syscall_unix.go
    │           │   ├── syscall_zos_s390x.go
    │           │   ├── syscall.go
    │           │   ├── sysvshm_linux.go
    │           │   ├── sysvshm_unix_other.go
    │           │   ├── sysvshm_unix.go
    │           │   ├── timestruct.go
    │           │   ├── unveil_openbsd.go
    │           │   ├── xattr_bsd.go
    │           │   ├── zerrors_aix_ppc.go
    │           │   ├── zerrors_aix_ppc64.go
    │           │   ├── zerrors_darwin_amd64.go
    │           │   ├── zerrors_darwin_arm64.go
    │           │   ├── zerrors_dragonfly_amd64.go
    │           │   ├── zerrors_freebsd_386.go
    │           │   ├── zerrors_freebsd_amd64.go
    │           │   ├── zerrors_freebsd_arm.go
    │           │   ├── zerrors_freebsd_arm64.go
    │           │   ├── zerrors_freebsd_riscv64.go
    │           │   ├── zerrors_linux_386.go
    │           │   ├── zerrors_linux_amd64.go
    │           │   ├── zerrors_linux_arm.go
    │           │   ├── zerrors_linux_arm64.go
    │           │   ├── zerrors_linux_loong64.go
    │           │   ├── zerrors_linux_mips.go
    │           │   ├── zerrors_linux_mips64.go
    │           │   ├── zerrors_linux_mips64le.go
    │           │   ├── zerrors_linux_mipsle.go
    │           │   ├── zerrors_linux_ppc.go
    │           │   ├── zerrors_linux_ppc64.go
    │           │   ├── zerrors_linux_ppc64le.go
    │           │   ├── zerrors_linux_riscv64.go
    │           │   ├── zerrors_linux_s390x.go
    │           │   ├── zerrors_linux_sparc64.go
    │           │   ├── zerrors_linux.go
    │           │   ├── zerrors_netbsd_386.go
    │           │   ├── zerrors_netbsd_amd64.go
    │           │   ├── zerrors_netbsd_arm.go
    │           │   ├── zerrors_netbsd_arm64.go
    │           │   ├── zerrors_openbsd_386.go
    │           │   ├── zerrors_openbsd_amd64.go
    │           │   ├── zerrors_openbsd_arm.go
    │           │   ├── zerrors_openbsd_arm64.go
    │           │   ├── zerrors_openbsd_mips64.go
    │           │   ├── zerrors_solaris_amd64.go
    │           │   ├── zerrors_zos_s390x.go
    │           │   ├── zptrace_armnn_linux.go
    │           │   ├── zptrace_linux_arm64.go
    │           │   ├── zptrace_mipsnn_linux.go
    │           │   ├── zptrace_mipsnnle_linux.go
    │           │   ├── zptrace_x86_linux.go
    │           │   ├── zsyscall_aix_ppc.go
    │           │   ├── zsyscall_aix_ppc64_gc.go
    │           │   ├── zsyscall_aix_ppc64_gccgo.go
    │           │   ├── zsyscall_aix_ppc64.go
    │           │   ├── zsyscall_darwin_amd64.1_13.go
    │           │   ├── zsyscall_darwin_amd64.1_13.s
    │           │   ├── zsyscall_darwin_amd64.go
    │           │   ├── zsyscall_darwin_amd64.s
    │           │   ├── zsyscall_darwin_arm64.1_13.go
    │           │   ├── zsyscall_darwin_arm64.1_13.s
    │           │   ├── zsyscall_darwin_arm64.go
    │           │   ├── zsyscall_darwin_arm64.s
    │           │   ├── zsyscall_dragonfly_amd64.go
    │           │   ├── zsyscall_freebsd_386.go
    │           │   ├── zsyscall_freebsd_amd64.go
    │           │   ├── zsyscall_freebsd_arm.go
    │           │   ├── zsyscall_freebsd_arm64.go
    │           │   ├── zsyscall_freebsd_riscv64.go
    │           │   ├── zsyscall_illumos_amd64.go
    │           │   ├── zsyscall_linux_386.go
    │           │   ├── zsyscall_linux_amd64.go
    │           │   ├── zsyscall_linux_arm.go
    │           │   ├── zsyscall_linux_arm64.go
    │           │   ├── zsyscall_linux_loong64.go
    │           │   ├── zsyscall_linux_mips.go
    │           │   ├── zsyscall_linux_mips64.go
    │           │   ├── zsyscall_linux_mips64le.go
    │           │   ├── zsyscall_linux_mipsle.go
    │           │   ├── zsyscall_linux_ppc.go
    │           │   ├── zsyscall_linux_ppc64.go
    │           │   ├── zsyscall_linux_ppc64le.go
    │           │   ├── zsyscall_linux_riscv64.go
    │           │   ├── zsyscall_linux_s390x.go
    │           │   ├── zsyscall_linux_sparc64.go
    │           │   ├── zsyscall_linux.go
    │           │   ├── zsyscall_netbsd_386.go
    │           │   ├── zsyscall_netbsd_amd64.go
    │           │   ├── zsyscall_netbsd_arm.go
    │           │   ├── zsyscall_netbsd_arm64.go
    │           │   ├── zsyscall_openbsd_386.go
    │           │   ├── zsyscall_openbsd_amd64.go
    │           │   ├── zsyscall_openbsd_arm.go
    │           │   ├── zsyscall_openbsd_arm64.go
    │           │   ├── zsyscall_openbsd_mips64.go
    │           │   ├── zsyscall_solaris_amd64.go
    │           │   ├── zsyscall_zos_s390x.go
    │           │   ├── zsysctl_openbsd_386.go
    │           │   ├── zsysctl_openbsd_amd64.go
    │           │   ├── zsysctl_openbsd_arm.go
    │           │   ├── zsysctl_openbsd_arm64.go
    │           │   ├── zsysctl_openbsd_mips64.go
    │           │   ├── zsysnum_darwin_amd64.go
    │           │   ├── zsysnum_darwin_arm64.go
    │           │   ├── zsysnum_dragonfly_amd64.go
    │           │   ├── zsysnum_freebsd_386.go
    │           │   ├── zsysnum_freebsd_amd64.go
    │           │   ├── zsysnum_freebsd_arm.go
    │           │   ├── zsysnum_freebsd_arm64.go
    │           │   ├── zsysnum_freebsd_riscv64.go
    │           │   ├── zsysnum_linux_386.go
    │           │   ├── zsysnum_linux_amd64.go
    │           │   ├── zsysnum_linux_arm.go
    │           │   ├── zsysnum_linux_arm64.go
    │           │   ├── zsysnum_linux_loong64.go
    │           │   ├── zsysnum_linux_mips.go
    │           │   ├── zsysnum_linux_mips64.go
    │           │   ├── zsysnum_linux_mips64le.go
    │           │   ├── zsysnum_linux_mipsle.go
    │           │   ├── zsysnum_linux_ppc.go
    │           │   ├── zsysnum_linux_ppc64.go
    │           │   ├── zsysnum_linux_ppc64le.go
    │           │   ├── zsysnum_linux_riscv64.go
    │           │   ├── zsysnum_linux_s390x.go
    │           │   ├── zsysnum_linux_sparc64.go
    │           │   ├── zsysnum_netbsd_386.go
    │           │   ├── zsysnum_netbsd_amd64.go
    │           │   ├── zsysnum_netbsd_arm.go
    │           │   ├── zsysnum_netbsd_arm64.go
    │           │   ├── zsysnum_openbsd_386.go
    │           │   ├── zsysnum_openbsd_amd64.go
    │           │   ├── zsysnum_openbsd_arm.go
    │           │   ├── zsysnum_openbsd_arm64.go
    │           │   ├── zsysnum_openbsd_mips64.go
    │           │   ├── zsysnum_zos_s390x.go
    │           │   ├── ztypes_aix_ppc.go
    │           │   ├── ztypes_aix_ppc64.go
    │           │   ├── ztypes_darwin_amd64.go
    │           │   ├── ztypes_darwin_arm64.go
    │           │   ├── ztypes_dragonfly_amd64.go
    │           │   ├── ztypes_freebsd_386.go
    │           │   ├── ztypes_freebsd_amd64.go
    │           │   ├── ztypes_freebsd_arm.go
    │           │   ├── ztypes_freebsd_arm64.go
    │           │   ├── ztypes_freebsd_riscv64.go
    │           │   ├── ztypes_illumos_amd64.go
    │           │   ├── ztypes_linux_386.go
    │           │   ├── ztypes_linux_amd64.go
    │           │   ├── ztypes_linux_arm.go
    │           │   ├── ztypes_linux_arm64.go
    │           │   ├── ztypes_linux_loong64.go
    │           │   ├── ztypes_linux_mips.go
    │           │   ├── ztypes_linux_mips64.go
    │           │   ├── ztypes_linux_mips64le.go
    │           │   ├── ztypes_linux_mipsle.go
    │           │   ├── ztypes_linux_ppc.go
    │           │   ├── ztypes_linux_ppc64.go
    │           │   ├── ztypes_linux_ppc64le.go
    │           │   ├── ztypes_linux_riscv64.go
    │           │   ├── ztypes_linux_s390x.go
    │           │   ├── ztypes_linux_sparc64.go
    │           │   ├── ztypes_linux.go
    │           │   ├── ztypes_netbsd_386.go
    │           │   ├── ztypes_netbsd_amd64.go
    │           │   ├── ztypes_netbsd_arm.go
    │           │   ├── ztypes_netbsd_arm64.go
    │           │   ├── ztypes_openbsd_386.go
    │           │   ├── ztypes_openbsd_amd64.go
    │           │   ├── ztypes_openbsd_arm.go
    │           │   ├── ztypes_openbsd_arm64.go
    │           │   ├── ztypes_openbsd_mips64.go
    │           │   ├── ztypes_solaris_amd64.go
    │           │   └── ztypes_zos_s390x.go
    │           └── windows
    │               ├── aliases.go
    │               ├── dll_windows.go
    │               ├── empty.s
    │               ├── env_windows.go
    │               ├── eventlog.go
    │               ├── exec_windows.go
    │               ├── memory_windows.go
    │               ├── mkerrors.bash
    │               ├── mkknownfolderids.bash
    │               ├── mksyscall.go
    │               ├── race.go
    │               ├── race0.go
    │               ├── security_windows.go
    │               ├── service.go
    │               ├── setupapi_windows.go
    │               ├── str.go
    │               ├── syscall_windows.go
    │               ├── syscall.go
    │               ├── types_windows_386.go
    │               ├── types_windows_amd64.go
    │               ├── types_windows_arm.go
    │               ├── types_windows_arm64.go
    │               ├── types_windows.go
    │               ├── zerrors_windows.go
    │               ├── zknownfolderids_windows.go
    │               └── zsyscall_windows.go
    └── modules.txt

17 directories, 355 files
```
