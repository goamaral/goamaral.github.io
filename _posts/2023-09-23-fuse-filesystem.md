---
layout: post
title: FUSE Filesystem - Sep 2023
---
This is a follow-up to a previous article exploring/implementing a FUSE filesystem. There is still a lot of work so this will become a series.

**Series**
- [Jan 2023](https://goamaral.github.io/posts/2023/01/05/fuse-filesystem)
- [Sep 2023](https://goamaral.github.io/posts/2023/09/23/fuse-filesystem) (this one)

**What was done**
- Started writing tests
- Improved filesystem mounting/unmounting flow
- Added logging functions
- Open - Mark node as open
- Write - Write data to a file
- Setattr - Change node mode
- Getxattr - Get extended attribute
- Remove - Remove file or directory

I started by writing some tests, to explore what interfaces I should implement next.

First tried to mount a unique filesystem in each test but started having trouble because I could not unmount the filesystem properly. This would be a huge pain as my tests grew. For now, I start the filesystem manually and run the tests against it.

The fuse library includes a `fstestutil` package that provides some functions to do exactly what I am trying to do but for some reason, the filesystem server hangs. In the future, I might give starting and mounting a filesystem in each test another try. I am running Linux in a VM, it should not cause problems but you never know. I also found a small bug in this package. Once I get this working I will contribute to the fuse repository.

Started by testing the `Write` method. Firstly I started with a basic success test, writing to a regular file. Note that the filesystem is mounted at `/tmp/fusefs`

```go
t.Run("Success", func(t *testing.T) {
		generatedFile := GenerateTestFile(t, "/tmp/fusefs")
		str := "hello"
		n, err := generatedFile.WriteString(str)
		assert.Equal(t, len(str), n)
		require.NoError(t, err)
	})
```

Then a failure test, trying to write to a read-only file

```go
t.Run("FileIsReadOnly", func(t *testing.T) {
		generatedFile := GenerateTestFile(t, "/tmp/fusefs")
		file, err := os.OpenFile(generatedFile.Name(), os.O_RDONLY, 0)
		require.NoError(t, err)
		t.Cleanup(func() { require.NoError(t, file.Close()) })

		n, err := file.WriteString("hello")
		assert.Zero(t, n)
		require.Error(t, err)
		pathErr, ok := err.(*fs.PathError)
		require.True(t, ok, "err is not *fs.PathError")
		errno, ok := pathErr.Err.(syscall.Errno)
		require.True(t, ok, "err is not syscall.Errno")
		assert.Equal(t, syscall.EBADF, errno)
	})
```

I tried using `chmod` on the file but realised I needed to implement the `fs.NodeSetattrer` interface to change the node permissions. I will probably explore node permissions after this series ends.

The `fuse.SetattrRequest` gives us a lot of fields but we will only use `Mode` for now. In the fuse source code, there was a comment (“*The type of the node is not guaranteed to be sent by the kernel, in which case os.ModeIrregular will be set.”),* I am not sure in what cases this could happen so I added an error log. I normally use man pages as a reference to how the function should behave and what error codes it should return. I suppose `chmod` triggers the method `Setattr` but could not find any info about this case.

```go
func (n *fuseFSNode) Setattr(ctx context.Context, req *fuse.SetattrRequest, resp *fuse.SetattrResponse) error {
	// NOTE: res.Atrr is filled by Attr method

	if req.Mode&os.ModeIrregular != 0 {
		Errorf("call to Setattr with mode irregular")
		return nil
	}

	n.Mode = req.Mode
	return nil
}
```

After this, I followed the filesystem logs and it made a call to the `fs.NodeGetxattrer` interface. Not sure what was calling this but implemented it anyway. After reading the man pages I think implementing it was not necessary cause not all filesystems need to implement it (there is a `ENOTSUP` error code which indicates that `xattrs` are not supported). I have some vague idea of `xattrs`, so I will explore them in the future (probably along node permissions).

```go
func (n fuseFSNode) Getxattr(ctx context.Context, req *fuse.GetxattrRequest, res *fuse.GetxattrResponse) error {
	// NOTE: req.Size is the size of res.Xattr. Size check is performed by fuse library

	if n.Xattrs == nil {
		return syscall.ENODATA
	}

	value, found := n.Xattrs[req.Name]
	if !found {
		return syscall.ENODATA
	}

	res.Xattr = []byte(value)
	return nil
}
```

Finally for the test to end successfully a call to delete the file is needed, so the `fs.NodeRemover` was implemented.

```go
func (n *fuseFSNode) Remove(ctx context.Context, req *fuse.RemoveRequest) error {
	for i, node := range n.Nodes {
		if node.Name == req.Name {
			// TODO: Test if rmdir fills req.Dir
			if req.Dir {
				if !node.Mode.IsDir() {
					return syscall.ENOTDIR
				}
				if len(req.Name) != 0 && req.Name[len(req.Name)-1] == '.' {
					return syscall.EINVAL
				}
			} else {
				if node.Mode.IsDir() {
					return syscall.EISDIR
				}
			}

			n.Nodes = append(n.Nodes[:i], n.Nodes[i+1:]...)
			return nil
		}
	}
	return syscall.ENOENT
}
```

At this point, the test was not returning the expected error. After some digging around I found that the `Write` method was checking the file `OpenFlags` . As the name suggests these flags check how the file was opened. Just had to open the file in read-only mode to make the test pass. This also made me realise I needed to check the file permissions. I will implement this in the future because I still need to figure out how to obtain the file/group owner and de request owner.

Our first method test is implemented, in the next article I will test the `Remove` method by using the syscalls `rm`, `rmdir`, `unlink`.

### References
https://www.gnu.org/software/libc/manual/html_node/Error-Codes.html
https://man7.org/linux/man-pages/index.html