---
title: "Creating A Blog Site: A Journey of Go, TailwindCSS and HTMX"
slug: creating-a-blog-site-go-tailwind-htmx
blurb: An explanation of my motivations when making this site and struggles with the technology.
date: 2025-08-04
modifieddate: 2025-11-07
tags:
  - go
  - tailwind
  - htmx
---

I am 22, born in 2002. I did not grow up in the world of Myspace pages. My first introduction to any sort of programming was C++ in the context of a VEX Robotics course at my high school. My coding adventures from then on were almost always tied to hardware, via Arduino, Unity and later Java through FRC. My first introduction to building UI's was also non html based, as I picked up flutter because the structure of creating widgets was something that I felt more familiar with, coming from a very object oriented background. All this to say I find HTML and CSS difficult. I'm also not good at creating user experiences. I very much fall into the trap of knowing how *I* would use my site / ui, but don't really consider how others would.

#### What is this for?
This site is for me to practice some frontend skills, write down my thoughts, mostly on tech related topics, and hold some of my FRC teaching materials.


#### The stack
I decided to go with Go, Templ, Tailwind, and HTMX. I chose this because I want more practice with go, in a web setting. I switched between templ and just raw html templating with go. When I started, I tried templ, thought it was too complicated so just wrote the html. Then I took another look at it a few days later and decided it wasn't that bad. It gives a little more ability for me to be dynamic and use components in an easier way, at least for me. Pretty much everyone uses tailwind at this point, for good reason, so I have the most amount of references to that, and I liked the idea of keeping the stack super light so I went with htmx. I did end up adding alpine js to handle some animations.