from os import environ, execl, path
from sys import executable
from time import sleep

from psutil import virtual_memory

ram_to_consume = int(virtual_memory().total * float(f"0.{int(TAKE) if (TAKE := environ.get("TAKE")) else 15}"))
allocated_mem = bytearray(ram_to_consume)

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
