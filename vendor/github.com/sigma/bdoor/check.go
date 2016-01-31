package bdoor

// #include <limits.h>
// #include <stdint.h>
// #include <signal.h>
// #include <setjmp.h>
//
// #define VMWARE_BDOOR_MAGIC 0x564D5868
// #define VMWARE_BDOOR_PORT 0x5658
// #define VMWARE_BDOOR_CMD_GETVERSION 10
// #define VMWARE_BDOOR_RUN(cmd, eax, ebx, ecx, edx)        \
//     __asm__("inl %%dx, %%eax" :                          \
//             "=a"(eax), "=c"(ecx), "=d"(edx), "=b"(ebx) : \
//             "0"(VMWARE_BDOOR_MAGIC),                     \
//             "1"(VMWARE_BDOOR_CMD_##cmd),                 \
//             "2"(VMWARE_BDOOR_PORT),                      \
//             "3"(UINT_MAX) :                              \
//             "memory");
//
// static sigjmp_buf env;
//
// void handler(int signal) {
//     siglongjmp(env, 1);
// }
//
// int hypervisor_bdoor_check(void) {
//     uint32_t eax, ebx, ecx, edx;
//     struct sigaction sa;
//     int res = 0;
//
//     sa.sa_handler = handler;
//     sigemptyset(&sa.sa_mask);
//     sa.sa_flags = SA_RESTART;
//     sigaction(SIGSEGV, &sa, 0);
//     sigaction(SIGILL, &sa, 0);
//
//     if(sigsetjmp(env, 1)==0) {
//         VMWARE_BDOOR_RUN(GETVERSION, eax, ebx, ecx, edx);
//         if (ebx == VMWARE_BDOOR_MAGIC)
//             res = 1;
//     }
//
//     sa.sa_handler = SIG_DFL;
//     sigaction(SIGSEGV, &sa, 0);
//     sigaction(SIGILL, &sa, 0);
//     return res;
// }
//
import "C"

func HypervisorPortCheck() bool {
	return C.hypervisor_bdoor_check() != 0
}
