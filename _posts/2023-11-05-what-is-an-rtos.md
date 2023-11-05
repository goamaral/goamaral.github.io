---
layout: post
title: What is an RTOS?
---
RTOS (Real-Time Operating System) is an OS for critical systems that need to process data and events e a defined time.

In this system, we trade speed for predictability. All processing must occur within the defined constraints. We have the guarantee that task x will run in n time. A task in this context is a set of program instructions loaded into memory.

These systems can be hard (operate within tens of milliseconds or less) or soft (operate within a few hundred milliseconds, at the scale of a human reaction), depending on how predictable they need to be. In a hard RTOS, a late answer is a wrong answer.

Size is a lot smaller being in the megabyte range instead of the gigabyte.

### Task switching
Tasks are only switched if a higher priority needs to be run, instead of switching in a regular clocked interrupt, in a round-robin fashion. Even with minimal thread switching, interrupt and thread switching latency is kept at a minimum.

### Interrupts
Only interrupt handlers have a higher priority than tasks and block the highest priority task from running. Because of this, they are typically kept as short as possible. Because they interrupt the running task, the internal OS object database can be in an inconsistent state. To deal with this, RTOS either disables interrupts, while the internal database is being updated (this can cause interrupts to be ignored) or creates a top-priority task to process the interrupt handler (increases latency).

### Memory allocation
Memory allocation is especially important because the device should work indefinitely, without ever needing a reboot. For this reason, dynamic memory allocation is frowned upon because it can easily lead to memory leaks. Dynamic allocation and releasing of small chunks of memory will cause memory fragmentation, reducing the efficient use of free memory (if a big chunk of memory is requested, it might not be possible to be allocated although enough global free memory is available) and allocation latency (allocated memory is typically represented as a linked list, with more fragments, the number of iterations required to find a free memory fragment increase). This is unacceptable in an RTOS since memory allocation has to occur within a certain amount of time.

Memory swap is not used because mechanical disks have much longer and unpredictable response times.

## Where are they used?
- Flight display controller
- Extraterrestrial rovers
- Emergency braking systems
- Engine warning systems
- Magnetic resonance imaging
- Surgery equipment
- Factory robotics systems

### References
[https://en.wikipedia.org/wiki/Real-time_operating_system](https://en.wikipedia.org/wiki/Real-time_operating_system)

[https://www.windriver.com/solutions/learning/rtos](https://www.windriver.com/solutions/learning/rtos)

[https://www.digikey.com/en/maker/projects/what-is-a-realtime-operating-system-rtos/28d8087f53844decafa5000d89608016](https://www.digikey.com/en/maker/projects/what-is-a-realtime-operating-system-rtos/28d8087f53844decafa5000d89608016)