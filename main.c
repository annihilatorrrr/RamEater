#include <stdlib.h>
#include <stdio.h>

int main() {
  // Get total memory in bytes
  long total = sysconf(_SC_PHYS_PAGES) * sysconf(_SC_PAGE_SIZE);
  // Calculate the amount of memory to use (10% of total)
  long use = total / 10;
  // Allocate memory
  char *memory = malloc(use);
  // Check if allocation was successful
  if (memory == NULL) {
    printf_s("Failed to allocate memory\n");
  }
  return 0;
}
