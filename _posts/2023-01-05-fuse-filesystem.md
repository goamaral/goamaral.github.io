---
layout: post
title: FUSE Filesystem
categories: linux filesystems golang
---
Filesystem in USErspace (FUSE) is a software interface for Unix and Unix-like computer operating systems that lets non-privileged users create their own file systems without editing kernel code. This is achieved by running file system code in user space while the FUSE module provides only a bridge to the actual kernel interfaces.

FUSE is available for Linux, FreeBSD, OpenBSD, NetBSD, OpenSolaris, Minix 3, macOS, and Windows.

### How does it work?

To implement a new file system, a handler program should use the libfuse library. This handler program should implement the required methods.

When the filesystem is mounted, the handler is registered with the kernel. Now, when a user calls an operation on this filesystem, the kernel will proxy these requests to the handler.

FUSE is particularly useful for writing virtual filesystems. Unlike traditional filesystems that essentially work with data on mass storage, virtual filesystems don't actually store data themselves. They act as a view or translation of an existing filesystem or storage device.

In principle, any resource available to a FUSE implementation can be exported as a file system.

### Where is this used?

Check these pages for a great examples of where FUSE is used.

[https://en.wikipedia.org/wiki/Filesystem_in_Userspace#Applications](https://en.wikipedia.org/wiki/Filesystem_in_Userspace#Applications)

[https://wiki.archlinux.org/title/FUSE](https://wiki.archlinux.org/title/FUSE)

### Basic implementation

[https://github.com/Goamaral/fuse-filesystem](https://github.com/Goamaral/fuse-filesystem)

For this first implementation I used Go. After a reviewing some solutions I decided to use [https://github.com/bazil/fuse](https://github.com/bazil/fuse). It seemed to be the easiest way to prototype.

This library implements the communication with the kernel from scratch in Go (without using libfuse) and enables an incremental implementation of a custom filesystem. It takes advantage of interfaces and if the implementation does not implement an interface (does not have a method), it has a fallback.

My goal for this implementation was to be able to list directory contents, create file, create directory.

I encourage you to check the code, it always seems harder before implementing.

**Implemented interfaces**

```go
type Node interface {
	Attr(ctx context.Context, attr *fuse.Attr) error
}
```

Get the file/directory attributes (permissions, ownership, size, …)

```go
type NodeStringLookuper interface {
	Lookup(ctx context.Context, name string) (Node, error)
}
```

Lookup file/directory by name inside a file/directory (of course, looking up anything inside a file should return an error)

```go
type NodeCreater interface {
	Create(ctx context.Context, req *fuse.CreateRequest, resp *fuse.CreateResponse) (Node, Handle, error)
}
```

Creates a file (not sure if it can create directories)

```go
type HandleReadDirAller interface {
	ReadDirAll(ctx context.Context) ([]fuse.Dirent, error)
}
```

List files and directories inside a directory

```go
type NodeMkdirer interface {
	Mkdir(ctx context.Context, req *fuse.MkdirRequest) (Node, error)
}
```

Create directory

### How do I unmount the FUSE filesystem?

```
$ fusermount3 -u MOUNTPOINT
```

### What’s next?

I will definitely continue implementing move interfaces like writing/reading to a file, get file size and explore the POSIX syscalls to find new features.

After that I will probably implement the same but in C (with libfuse probably) and register the handler in the kernel.

### References

[https://en.wikipedia.org/wiki/Filesystem_in_Userspace](https://en.wikipedia.org/wiki/Filesystem_in_Userspace)

[https://wiki.archlinux.org/title/FUSE](https://wiki.archlinux.org/title/FUSE)

[https://man7.org/linux/man-pages/man3/errno.3.html](https://man7.org/linux/man-pages/man3/errno.3.html)

[https://github.com/libfuse/libfuse/wiki/FAQ](https://github.com/libfuse/libfuse/wiki/FAQ)

[https://github.com/bazil/fuse](https://github.com/bazil/fuse)