---
layout: post
title: What is an ELF file?
categories: linux compilers
---
*(The information in this article might be incomplete, I only include information I understood or considered most relevant. Please visit the references for more information)*

Is a file format used for executable files, object code, shared libraries, and core dumps.

By design, the ELF format is flexible, extensible, and cross-platform. For instance, it supports different endiannesses and address sizes, so it does not exclude any particular CPU or instruction set architecture.

## File layout

![ELF file](https://dev-to-uploads.s3.amazonaws.com/uploads/articles/p6jp47u343jhvfjdyexi.png)

Each ELF file is made up of one ELF header, followed by file data. The data can include:

- Program header table (PHT), describing zero or more memory segments
- Section header table (SHT), describing zero or more sections
- Data referred to by entries in the program header table or section header table

The segments contain information that is needed for run time execution of the file, while sections contain important data for linking and relocation.

### ELF header

- 32/64 bit format
- endianness
- target ABI
- file type (relocatable, executable, shared, core, others…)
- instruction set
- entry point address
- program header address
- section header address
- size of this header
- program header table entry size
- program header table entry count
- section header table entry size
- section header table entry count
- index of section entry that contains the section names

### Program header

- type
- offset of the segment in the file image
- virtual address of the segment in memory
- size in bytes of the segment in the file image
- size in bytes of the segment in memory

### Section header

- name
- type
- virtual address of the section in memory
- offset of the section in the file image
- size in bytes of the section in the file image

## How to check ELF file content?

### ELF header

```
$ readelf -h /bin/gcc

ELF Header:
  Magic:   7f 45 4c 46 02 01 01 00 00 00 00 00 00 00 00 00
  Class:                             ELF64
  Data:                              2's complement, little endian
  Version:                           1 (current)
  OS/ABI:                            UNIX - System V
  ABI Version:                       0
  Type:                              EXEC (Executable file)
  Machine:                           Advanced Micro Devices X86-64
  Version:                           0x1
  Entry point address:               0x4077a0
  Start of program headers:          64 (bytes into file)
  Start of section headers:          954048 (bytes into file)
  Flags:                             0x0
  Size of this header:               64 (bytes)
  Size of program headers:           56 (bytes)
  Number of program headers:         14
  Size of section headers:           64 (bytes)
  Number of section headers:         31
  Section header string table index: 30
```

### Program headers

```
$ readelf -l /usr/bin/gcc

Program Headers:
  Type           Offset             VirtAddr           PhysAddr
                 FileSiz            MemSiz              Flags  Align
  PHDR           0x0000000000000040 0x0000000000400040 0x0000000000400040
                 0x0000000000000310 0x0000000000000310  R      0x8
  INTERP         0x0000000000000350 0x0000000000400350 0x0000000000400350
                 0x000000000000001c 0x000000000000001c  R      0x1
      [Requesting program interpreter: /lib64/ld-linux-x86-64.so.2]
  LOAD           0x0000000000000000 0x0000000000400000 0x0000000000400000
                 0x0000000000002538 0x0000000000002538  R      0x1000
  LOAD           0x0000000000003000 0x0000000000403000 0x0000000000403000
                 0x0000000000061301 0x0000000000061301  R E    0x1000
  LOAD           0x0000000000065000 0x0000000000465000 0x0000000000465000
                 0x0000000000080bf8 0x0000000000080bf8  R      0x1000
  LOAD           0x00000000000e6a40 0x00000000004e6a40 0x00000000004e6a40
                 0x00000000000021e8 0x0000000000005a60  RW     0x1000
  DYNAMIC        0x00000000000e7a38 0x00000000004e7a38 0x00000000004e7a38
                 0x00000000000001c0 0x00000000000001c0  RW     0x8
  NOTE           0x0000000000000370 0x0000000000400370 0x0000000000400370
                 0x0000000000000050 0x0000000000000050  R      0x8
  NOTE           0x00000000000003c0 0x00000000004003c0 0x00000000004003c0
                 0x0000000000000044 0x0000000000000044  R      0x4
  TLS            0x00000000000e6a40 0x00000000004e6a40 0x00000000004e6a40
                 0x0000000000000000 0x0000000000000010  R      0x8
  GNU_PROPERTY   0x0000000000000370 0x0000000000400370 0x0000000000400370
                 0x0000000000000050 0x0000000000000050  R      0x8
  GNU_EH_FRAME   0x00000000000da014 0x00000000004da014 0x00000000004da014
                 0x0000000000001acc 0x0000000000001acc  R      0x4
  GNU_STACK      0x0000000000000000 0x0000000000000000 0x0000000000000000
                 0x0000000000000000 0x0000000000000000  RW     0x10
  GNU_RELRO      0x00000000000e6a40 0x00000000004e6a40 0x00000000004e6a40
                 0x00000000000015c0 0x00000000000015c0  R      0x1
```

### Section headers

```
$ readelf -S /usr/bin/gcc

Section Headers:
  [Nr] Name              Type             Address           Offset
       Size              EntSize          Flags  Link  Info  Align
  [ 0]                   NULL             0000000000000000  00000000
       0000000000000000  0000000000000000           0     0     0
  [ 1] .interp           PROGBITS         0000000000400350  00000350
       000000000000001c  0000000000000000   A       0     0     1
  [ 2] .note.gnu.pr[...] NOTE             0000000000400370  00000370
       0000000000000050  0000000000000000   A       0     0     8
  [ 3] .note.gnu.bu[...] NOTE             00000000004003c0  000003c0
       0000000000000024  0000000000000000   A       0     0     4
  [ 4] .note.ABI-tag     NOTE             00000000004003e4  000003e4
       0000000000000020  0000000000000000   A       0     0     4
  [ 5] .gnu.hash         GNU_HASH         0000000000400408  00000408
       00000000000000a4  0000000000000000   A       6     0     8
  [ 6] .dynsym           DYNSYM           00000000004004b0  000004b0
       0000000000000cd8  0000000000000018   A       7     1     8
  [ 7] .dynstr           STRTAB           0000000000401188  00001188
       0000000000000591  0000000000000000   A       0     0     1
  [ 8] .gnu.version      VERSYM           000000000040171a  0000171a
       0000000000000112  0000000000000002   A       6     0     2
  [ 9] .gnu.version_r    VERNEED          0000000000401830  00001830
       00000000000000f0  0000000000000000   A       7     2     8
  [10] .rela.dyn         RELA             0000000000401920  00001920
       0000000000000c18  0000000000000018   A       6     0     8
  [11] .init             PROGBITS         0000000000403000  00003000
       000000000000001b  0000000000000000  AX       0     0     4
  [12] .text             PROGBITS         0000000000403020  00003020
       00000000000612d3  0000000000000000  AX       0     0     16
  [13] .fini             PROGBITS         00000000004642f4  000642f4
       000000000000000d  0000000000000000  AX       0     0     4
  [14] .rodata           PROGBITS         0000000000465000  00065000
       0000000000075010  0000000000000000   A       0     0     32
  [15] .stapsdt.base     PROGBITS         00000000004da010  000da010
       0000000000000001  0000000000000000   A       0     0     1
  [16] .eh_frame_hdr     PROGBITS         00000000004da014  000da014
       0000000000001acc  0000000000000000   A       0     0     4
  [17] .eh_frame         PROGBITS         00000000004dbae0  000dbae0
       000000000000a048  0000000000000000   A       0     0     8
  [18] .gcc_except_table PROGBITS         00000000004e5b28  000e5b28
       00000000000000d0  0000000000000000   A       0     0     4
  [19] .tbss             NOBITS           00000000004e6a40  000e6a40
       0000000000000010  0000000000000000 WAT       0     0     8
  [20] .init_array       INIT_ARRAY       00000000004e6a40  000e6a40
       0000000000000018  0000000000000008  WA       0     0     8
  [21] .fini_array       FINI_ARRAY       00000000004e6a58  000e6a58
       0000000000000008  0000000000000008  WA       0     0     8
  [22] .data.rel.ro      PROGBITS         00000000004e6a60  000e6a60
       0000000000000fd8  0000000000000000  WA       0     0     32
  [23] .dynamic          DYNAMIC          00000000004e7a38  000e7a38
       00000000000001c0  0000000000000010  WA       7     0     8
  [24] .got              PROGBITS         00000000004e7bf8  000e7bf8
       0000000000000400  0000000000000008  WA       0     0     8
  [25] .data             PROGBITS         00000000004e8000  000e8000
       0000000000000c28  0000000000000000  WA       0     0     32
  [26] .bss              NOBITS           00000000004e8c40  000e8c28
       0000000000003860  0000000000000000  WA       0     0     32
  [27] .comment          PROGBITS         0000000000000000  000e8c28
       0000000000000012  0000000000000001  MS       0     0     1
  [28] .note.stapsdt     NOTE             0000000000000000  000e8c3c
       0000000000000130  0000000000000000           0     0     4
  [29] .gnu_debuglink    PROGBITS         0000000000000000  000e8d6c
       0000000000000010  0000000000000000           0     0     4
  [30] .shstrtab         STRTAB           0000000000000000  000e8d7c
       0000000000000143  0000000000000000           0     0     1
```

### Everything

```
$ readelf -a /usr/bin/gcc
$ objdump -x /usr/bin/gcc
$ file /usr/bin/gcc
```

## References

[https://en.wikipedia.org/wiki/Executable_and_Linkable_Format](https://en.wikipedia.org/wiki/Executable_and_Linkable_Format)

[https://linuxhint.com/understanding_elf_file_format](https://linuxhint.com/understanding_elf_file_format/)

[https://man7.org/linux/man-pages/man5/elf.5.html](https://man7.org/linux/man-pages/man5/elf.5.html)