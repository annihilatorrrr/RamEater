from time import sleep

from psutil import virtual_memory

ram_to_consume = int(virtual_memory().total * 0.2)
allocated_mem = bytearray(ram_to_consume)

while 1:
    sleep(12 * 3600)
    pass
