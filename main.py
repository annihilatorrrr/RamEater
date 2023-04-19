from psutil import virtual_memory

# By @sanjitsinha
ram_to_consume = int(virtual_memory().total * 0.3)

allocated_mem = bytearray(ram_to_consume)

while 1:
    pass
