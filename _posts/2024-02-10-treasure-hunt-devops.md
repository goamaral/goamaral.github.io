---
layout: post
title: Treasure Hunt - Devops | Feb 2024
---
1. [https://unixism.net/loti/async_intro.html](https://unixism.net/loti/async_intro.html)
    
    This article is related to the [io_ring article](/posts/2023/04/07/io-uring), basically it takes a short tour around linux async io and different solutions on how to handle io. We use this technology without knowing, for example, most web servers implement this strategy. Specially interpreted languages with global locks (python, ruby) or using an event loop model (nodejs).
    
2. [https://earthly.dev/blog/chroot](https://earthly.dev/blog/chroot)
    
    What if docker is just a magic trick, something that seems so complex but if you lift the curtain it becomes quite easy to understand. This article lifts that curtain revealing docker and containers in general are just a facade on top of chroot.
    
3. [https://github.com/bibendi/dip](https://github.com/bibendi/dip) + [https://github.com/evilmartians/ruby-on-whales](https://github.com/evilmartians/ruby-on-whales)
    
    I have tried many times, to some success, to create a dev environment that would have a fast and automated setup and that could run in any OS. I leaned heavily on docker and recently with vscode dev containers on top of that. But I always felt it should be simpler.
    
    That’s where dip comes in, with a Dockerfile and a dip config file (very similar to docker-compose.yml), a docker dev environment is setup and with the help of the dip cli, a near native/local experience can be achieved. As an example, this is what I have to type to start a rails console that runs inside a container `dip run rails c` .
    
    And to further simplify things, we have [https://github.com/evilmartians/ruby-on-whales](https://github.com/evilmartians/ruby-on-whales), a rails app template that comes with default a Dockerfile and dip config.
    

4. [https://biriukov.dev/docs/fd-pipe-session-terminal/0-sre-should-know-about-gnu-linux-shell-related-internals-file-descriptors-pipes-terminals-user-sessions-process-groups-and-daemons/](https://biriukov.dev/docs/fd-pipe-session-terminal/0-sre-should-know-about-gnu-linux-shell-related-internals-file-descriptors-pipes-terminals-user-sessions-process-groups-and-daemons/)
    
    Deep dive into file descriptors, pipes, processes, sessions, jobs, terminals and pseudoterminals. I haven’t finished this article yet but I can only say good things. It has a clear explanation of the topics with pictures and code to go along.
    

5. [https://12factor.net/](https://12factor.net/)
    
    Best practices on how to build/manage an application. Most of them are intuitive, specially in a cloud world. But is always important to keep things in perspective, things that are obvious now weren’t in the past.