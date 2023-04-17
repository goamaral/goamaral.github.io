---
layout: post
title: Static vs dynamic linking
categories: linux compilers
---
### What is static linking?

Static linking links libraries at compile time, copying them to the final binary.

### What is dynamic linking?

Dynamic linking loads and links libraries at runtime, loading them to memory.

Only the name of the shared libraries is saved at compile time.

These names are saved in a PLT (Procedure Linkage Table)

### Static vs dynamic linking

**Static**

- Bigger binaries

**Dynamic**

- Depend on external libraries to be installed and be compatible
- Shared libraries are shared across processes
- Shared library code can be updated/patched without new compilation
- Updates to shared library code can add breaking changes and prevent the program from running

### How to create a statically linked binary?

```
$ ld [options] objfile
```

`ld` combines several object and archive files, relocates their data and ties up symbol references. Usually, the last step in compiling a program is to run `ld`.

```
$ gcc hello.c -static -o hello
```

### How to create a dynamically linked binary?

```
$ gcc hello.c -o hello
```

### How to know if a binary is statically or dynamically linked?

**Check the type of linking**

```
$ file /usr/bin/gcc

/usr/bin/gcc: ELF 64-bit LSB executable, x86-64, version 1 (SYSV), dynamically linked, interpreter /lib64/ld-linux-x86-64.so.2, BuildID[sha1]=017fc52acbca077c9bc6a4e8f04dd90eb5385243, for GNU/Linux 4.4.0, stripped
```

**Check dynamically linked libraries**

```
$ ldd /bin/gcc

linux-vdso.so.1 (0x00007fff6377e000)
libc.so.6 => /usr/lib/libc.so.6 (0x00007fcd238f2000)
/lib64/ld-linux-x86-64.so.2 => /usr/lib64/ld-linux-x86-64.so.2 (0x00007fcd23b02000)
```