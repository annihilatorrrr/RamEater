#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <sys/sysinfo.h>

#define ONE_DAY_SECONDS (24 * 3600)

void allocate_memory() {
    struct sysinfo info;
    sysinfo(&info);

    const char *take_str = getenv("TAKE");
    int take = take_str ? atoi(take_str) : 15;

    unsigned long total_memory = info.totalram;
    unsigned long used_memory = info.totalram - info.freeram;
    unsigned long memory_to_allocate = total_memory * (take / 100.0) - used_memory;

    if (memory_to_allocate > 0) {
        void *allocated_mem = malloc(memory_to_allocate);
        if (allocated_mem) {
            memset(allocated_mem, 0, memory_to_allocate);
            printf("Allocated %lu bytes to reach target memory usage.\n", memory_to_allocate);
        } else {
            perror("malloc");
        }
    } else {
        printf("No need to allocate memory.\n");
    }
}

int main(int argc, char *argv[]) {
    allocate_memory();
    printf("Done!\n");

    while (1) {
        sleep(ONE_DAY_SECONDS);

        if (!getenv("NOCPUB")) {
            long result = 1;
            for (int i = 1; i < 1000000; i++) {
                result *= i;
            }
        }

        // Restart the program
        execvp(argv[0], argv);
        perror("execvp");  // If execvp returns, it must have failed
        exit(EXIT_FAILURE);
    }

    return 0;
}
