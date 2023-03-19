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
  // Get total memory in bytes
  total = status.ullTotalPhys;
  #else
    // Get total memory in bytes
  total = sysconf(_SC_PHYS_PAGES) * sysconf(_SC_PAGE_SIZE);
  #endif
  // Calculate the amount of memory to use (10% of total)
  use = total / 10;
  // Allocate memory
  char *memory = (char *) malloc(use);
  // Sleep for 10 seconds
  #ifdef _WIN32
    Sleep(10000);
  #else
  sleep(10);
  #endif
  // Free memory
  free(memory);
  return 0;
}
