#include <stdio.h>
#include <stdlib.h>

#ifdef _WIN32
#include <windows.h>
#else
#include <unistd.h>
#endif

int main() {
    long long total, use;
#ifdef _WIN32
    MEMORYSTATUSEX status;
    status.dwLength = sizeof(status);
    GlobalMemoryStatusEx(&status);
    total = status.ullTotalPhys;
#else
    total = sysconf(_SC_PHYS_PAGES) * sysconf(_SC_PAGE_SIZE);
#endif
    use = total / 20;
    char *memory = (char *) malloc(use);
    printf_s("\nDone.\n");
    while (1) {
#ifdef _WIN32
        Sleep(10000);
#else
        sleep(10);
#endif
    }
    return 0;
}
