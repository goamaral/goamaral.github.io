---
layout: post
title: Job Adventures - PDF generation | Jun 2024
---
Well, here we are with a new series. This one is called *Job Adventures* where I will talk about some challenges I encountered on my day to day job.

In this article we will explore PDF generation. This is one of those classic tasks you rarely need to do but when the task eventually arrives, I get PTSD.

My first contact with building PDFs was with rails using https://github.com/mileszs/wicked_pdf. The task always seems easy, you just build HTML and render that to pdf. And in fact, the part of rendering the info to the pdf is easy. The nightmare comes when implementing what is on the mockups. How will CSS behave in printing mode? What if we have a component that can’t split on a page break, it should jump in its entirety to the next page? What if our cover page does not count to the page total? What if the cover page does not have an header/footer? Why is the pdf so big?

Some of those problems I had in the past, but at the time I was just rendering tables for a financial report. The main problem I remember having was the CSS part and the long generation time. Because I was not implementing the styling at the time, the CSS part was not really my problem, and I am sure wicked_pdf provides some default styles to help in this part. The long processing times were a problem because we were generating pdfs with over 100 pages, this process would take about 5 min and would get worse if more pdfs were being requested in parallel. I can’t remember what the solution was at the time but I think we ended up generating some pdfs in the background and sending them by email when ready. The wicked_pdf gem uses an instance of https://github.com/wkhtmltopdf/wkhtmltopdf under the hood. This causes problems because it can only generate pdfs one by one. The solution would probably be having a dedicated service that would orchestrate multiple wkhtmltopdf instances.

Jumping to today, I am using Go and my first instinct was to find a binding to wkhtmltopdf and go from there. I remember trying to find better solutions to wicked_pdf at the time and none was better, so I started with what I knew worked. What a big surprise it was when I opened wkhtmltopdf github page and found it archived. Basically, it was based on QtWebKit that stopped being maintained long ago. You can find a longer explanation [here](https://wkhtmltopdf.org/status.html).

After some searching, I found https://github.com/gotenberg/gotenberg. It ticked a lot of boxes.
- It is an independent service that communicates via HTTP. I just send the url to the page I want to convert to PDF and receive the pdf back. This way we have an easily scalable service that can be easily integrated with any other system/language.
- The same team maintains a docker image. So we don’t need to worry with any basic dependencies like headless chrome or fonts. Just start a container and relax.
- It is written in go, if needed, I can easily open an issue/PR or fork it.

And now you might say, all good. Just create an HTML page and we are done. I wish it would be that easy. Now it’s time to answer the questions I placed in the beginning.

**How will CSS behave in printing mode?**
**Why is the pdf so big?**
From what I experienced, there where not many sharp edges. The only thing that caught me off guard was `print-color-adjust` , it defaults to `economy` (which makes sense, to use less ink). The first pages I created were mostly text and tables, no problems at this point, until I added a couple of images and when previewing the print version, the colours were really saturated. It retrospective the solution was easy but at the time I had no clue if the problem was with gottenberg, what property I should change/add or if it was even possible. The solution was to set `print-color-adjust` to `exact` . Just be aware, that this is not free, the size of the pdf increased significantly.

**What if we have a component that cant split on a page break, it should jump in its entirety to the next page?**
**What if our cover page does not count to the page total?**
**What if the cover page does not have an header/footer?**
By default you can easily add a header and a footer to every page, the same applies to the counter. But requirements are rarely that simple. But this problems were moderately simple to solve. I disabled footers and headers and manually implemented a header and footer component, this way I have full control when they are shown and what pages count.

The big problem came with dynamically sized content. Without an image it can be hard to explain, but some components should not break (charts and content with side images) and others should (tables). Because all this components varied in the amount of info they had, I calculate the pixel height they would occupy, the vertical space I had left in the page and choose if the component should be split or not. These solution was far from perfect and I feel there should be a better. In hindsight, after exploring more properties like `page-break-before` I feel this could have solved many of my issues. Even with this in mind, one of the requirements was to have the table header always present at the top on a page break and I don’t think `page-break-*`  properties would help with that.

This feature was developed a couple months ago, so I don’t recall a lot of the issues I had but these were the lessons that stuck with me and that will apply in the next pdf I need to generate (hopefully not soon).