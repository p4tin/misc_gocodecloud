from ctypes import *
from ctypes import cdll
import time

class go_string(Structure):
    _fields_ = [
        ("p", c_char_p),
        ("n", c_int)]

lib = cdll.LoadLibrary('./lib_disco.so')


def start_discovery():
    lib.startDiscovery()


def get_service_url(str):
    b = go_string(c_char_p(str), len(str))
    lib.getServiceURL.restype = c_char_p
    a = lib.getServiceURL(b, c_char_p(str))
    print a

start_discovery()
get_service_url("checkout")
get_service_url("pricing")

while True:
    time.sleep(5)
    print(lib.listAllNodes())
