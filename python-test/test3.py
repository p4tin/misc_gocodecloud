from cffi import FFI
ffi = FFI()
ffi.cdef("""
    typedef struct { int x, y; } IntStruct;
    int add(int a, int b);
    int structadd(IntStruct s);
""")
lib = ffi.dlopen("./libadd.so")


print(lib.add(2, 2))

struct = ffi.new("IntStruct *", [1, 2])
print(lib.structadd(struct))