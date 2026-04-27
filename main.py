from os import environ, execl, path
from sys import executable
from time import sleep

from psutil import virtual_memory


def ram_to_consume():
    mem = virtual_memory()
    take = int(environ.get("TAKE") or 15)
    memory_to_allocate = int(mem.total * take / 100 - mem.used)
    if memory_to_allocate > 0:
        allocated_memo = bytearray(memory_to_allocate)
        print(f"Allocated {len(allocated_memo)} bytes to reach target memory usage.")
        return allocated_memo
    print("No need to allocate memory.")
    return None


allocated_mem = ram_to_consume()
print("Done!")

while True:
    sleep(24 * 3600)
    if not environ.get("NOCPUB"):
        result = 1
        for i in range(1, 1_000_000):
            result *= i
    script = "main.py" if path.isfile("main.py") else "main.pyc"
    execl(executable, executable, script)
