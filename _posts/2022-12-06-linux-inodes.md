---
layout: post
title: What are linux inodes?
categories: linux filesystems
---
An inode is an index node for every file and directory in the filesystem. Inodes do not store actual data. Instead, they store the metadata where you can find the storage blocks of each file’s data.

### Metadata

- File type
- Permissions
- Owner ID
- Group ID
- Size of file
- Time last accessed
- Time last modified
- Soft/Hard Links
- Access Control List (ACLs)

### How to check inode information?

```
$ stat /bin/gcc
  File: /bin/gcc
  Size: 956032    	Blocks: 1872       IO Block: 4096   regular file
Device: 8,1	Inode: 4993952     Links: 3
Access: (0755/-rwxr-xr-x)  Uid: (    0/    root)   Gid: (    0/    root)
Access: 2022-09-13 00:03:33.000000000 +0100
Modify: 2022-08-20 02:12:31.000000000 +0100
Change: 2022-09-13 00:03:33.715344081 +0100
 Birth: 2022-09-13 00:03:33.705344003 +0100

$ ls -i /bin/gcc
4993952 /bin/gcc
```

### How to check the inode usage on filesystems?

```
$ df -i
Filesystem      Inodes   IUsed   IFree IUse% Mounted on
dev             878882     564  878318    1% /dev
run             880930     941  879989    1% /run
/dev/sda1      6553600 1260316 5293284   20% /
tmpfs           880930     271  880659    1% /dev/shm
tmpfs           880930      60  880870    1% /tmp
tmpfs           176186     129  176057    1% /run/user/1000
```

### What happens to the inode assigned when moving or copying a file?

When you copy a file, linux assigns a different inode to the new file.

```
$ touch file1
$ ls -i file1
2674409 file1
$ cp file1 file2
$ ls -i
2674409 file1  2674178 file2
```

When moving a file, the inode remains the same, as long as the file does not change filesystems.

```
$ touch file1
$ ls -i file1
2674409 file1
$ mkdir dir
$ mv file1 dir/
$ ls -i dir/file1
2674409 dir/file1
```

If we change filesystems, the inode changes.

```
$ touch file1
$ ls -i file1
2674409 file1
$ mv file1 /run/media/architect/123253A832538F99
$ ls -i /run/media/architect/123253A832538F99/file1
37316 /run/media/architect/123253A832538F99/file1
```

Hard links connect directly to the same inode. Soft links creates a new inode.

```
$ touch file1
$ ls -i file1
2674409 file1
$ ln file1 file1_hl
$ ls -i file1_hl
2674409 file1_hl
$ ln -s file1 file1_sl
$ ls -i file1_sl
2674103 file_sl
```

### What is the maximum inode value?

In the kernel source code it is coded as a 32-bit unsigned long integer, so the theoretical value would be 2³² (4,294,967,295).

That’s the theoretical maximum. In practice, the number of inodes in an ext4 file system is determined when the file system is created at a default ratio of one inode per 16 KB of file system capacity. Directory structures are created on the fly when the file system is in use, as files and directories are created within the file system.

### References

[https://docs.rackspace.com/support/how-to/what-are-inodes-in-linux/](https://docs.rackspace.com/support/how-to/what-are-inodes-in-linux/)

[https://www.howtogeek.com/465350/everything-you-ever-wanted-to-know-about-inodes-on-linux/](https://www.howtogeek.com/465350/everything-you-ever-wanted-to-know-about-inodes-on-linux/)

[https://www.site24x7.com/learn/linux/inode.html](https://www.site24x7.com/learn/linux/inode.html)

[https://en.wikipedia.org/wiki/Inode](https://en.wikipedia.org/wiki/Inode)