---
layout: post
title: General Dump - Jul 2023
---
This will be the first post of a series of posts I will call *Dumps*. In each post, I will dump about 5 items that caught my attention (articles, libraries, any kind of link/reference really). This post will be a *General Dump*, where I list items that are not related to a specific topic. In the future, there will be specific *Dump* series (ruby, linux and containers, databases, go). Hope you enjoy this new format.

1. [https://www.16elt.com/2023/01/06/logging-practices-I-follow](https://www.16elt.com/2023/01/06/logging-practices-I-follow)
    
    A few months ago I had a production bug that required reading logs to track down the root cause. Unfortunately, the logs were useless. Not only because of the quantity but also the quality. In a sea of logs, we need a way to track what logs belong to the same flow and get something useful from a flow that we can test against (an id, a date, SQL, etc.).
    
    This incident made me think about a better way to do it and experiment while developing the next features.
    
    The article is a great shortcut. It sums up what I had to find out on my own.
    

2. [https://www.quantamagazine.org/how-to-prove-you-know-a-secret-without-giving-it-away-20221011](https://www.quantamagazine.org/how-to-prove-you-know-a-secret-without-giving-it-away-20221011)
    
    I first ran against zero-knowledge proofs while diving into the world of blockchains. But they are not only applicable in that space. If you think about a system, many proofs exist. Things like authenticating a user, verifying if a user owns an asset and anything that a user needs to prove.
    
    This article explains this topic while keeping things beginner friendly.
    

3. [https://theconversation.com/how-to-test-if-were-living-in-a-computer-simulation-194929](https://theconversation.com/how-to-test-if-were-living-in-a-computer-simulation-194929)
    
    If you have ever thought about our existence, this article might interest you. It gives a really interesting take on the simulation hypothesis (aka we live in a computer simulation).
    

4. [https://github.com/alex/what-happens-when](https://github.com/alex/what-happens-when)
    
    Explains what happens when we search on the web browser. From when *the"g" key is pressed* until the end of the first browser paint.
    

5. [https://endoflife.date](https://endoflife.date)
    
    Sometimes it can be hard to know when a version of a package/service has reached the end of life or when the day will come. Recently, I found this site that gives us these important dates. This way, we can plan our upgrades instead of stressing out when a warning appears.