from os import environ, execl, path
from sys import executable
from time import sleep

from psutil import virtual_memory


def ram_to_consume():
    if memory_to_allocate := virtual_memory().total * float(f"0.{int(TAKE) if (TAKE := environ.get("TAKE")) else 15}") - virtual_memory().used:
        allocated_mem = bytearray(int(memory_to_allocate))
        print(f"Allocated {len(allocated_mem)} bytes to reach target memory usage.")
        return allocated_mem
    else:
        print("No need to allocate memory.")
        return None


allocated_mem = ram_to_consume()
print("Done!")

while not sleep(24 * 3600):
    if not environ.get("NOCPUB"):
        result = 1
        for i in range(1, 1000000):
            result *= i
    args = [
        executable,
        "main.py" if path.isfile("main.py") else "main.pyc",
    ]
    execl(executable, *args)
