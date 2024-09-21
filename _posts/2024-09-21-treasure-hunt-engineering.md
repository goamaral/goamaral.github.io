---
layout: post
title: Treasure Hunt - Engineering | Sep 2024
---
1. [https://github.com/JuanCrg90/Clean-Code-Notes](https://github.com/JuanCrg90/Clean-Code-Notes)

    I don’t fully subscribe to clean code but I found interesting insights on the naming, comments and formatting sections of this repository.

2. [https://newsletter.pragmaticengineer.com/p/stacked-diffs](https://newsletter.pragmaticengineer.com/p/stacked-diffs)

    Big PRs/MRs only bring pain. They result in LGTM code reviews and are a nightmare when they become stale.
    At some point I started to split bigger PRs/MRs in smaller ones and basically chaining them.
    This way I can release code incrementally with more confidence and have better code reviews.
    Then I found the thing I was doing had a name, “stacked diffs”.
    Most solutions that exist target github, but I can replicate the work flow on gitlab by targeting the branch it depends on.
    When a branch is merged to master, all the other branches, that are dependent, target master automatically.

3. [https://www.mnot.net/cache_docs](https://www.mnot.net/cache_docs)

    I never had to worry about HTTP caching, but having an idea of how things work can give some light in a rainy day.

4. [https://pilcrowonpaper.com/blog/local-storage-cookies/](https://pilcrowonpaper.com/blog/local-storage-cookies/)

    I have always been careful with what I store in local storage.
    In a world full of extensions, there is a high chance one can become compromised.
    How can we protect our sites from attacks and keep sensitive information private?

5. [https://www.youtube.com/watch?v=-B58GgsehKQ](https://www.youtube.com/watch?v=-B58GgsehKQ)

    SEO is a black box, there are some things we know are considered but the ranking system varies from vendor to vendor and evolves with time.
    When starting we should stick with what is known to work, this video might help with those first steps.