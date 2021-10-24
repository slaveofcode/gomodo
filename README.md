# Gomodo

Web application suite for boring developer project

> In a hurry? go ahead to the [Get Started](https://github.com/slaveofcode/gomodo#Setup)


## Background

I like to build web apps, crawl the best part of the web technologies nowadays, and wonder if all of it
could be combined into one simple starter project that can be picked up every time when I need to build an app.

I love [Vite](https://vitejs.dev) (pronounced `/vit/`) when building a web app using Vue, the simplicity and performance (faster build, HMR, etc.) make me think how if I combine that thing using Go as the backend, which I like to use [Fiber](https://gofiber.io/).

So here it is, the combination of VueJs as the frontend webapp and Go as the API provider for the backend. 

## The Idea

Web with a bunch of tooling, libraries, and configurations should have its own privilege. That's what the purpose of the `web` folder is. Go, should only be aware of the backend side, that's why there's an `app` folder. Combining those two in production mode basically only needs a compiled-ready-to-use version of the web and served by the backend for the production stage. So instead of having the web static files served via CDN or another cloud provider, we have a single binary of the web app itself, which is considerably easier to host and run.

---

## Get Started

Make sure Golang & Node already installed in your computer, with minimum version is **1.14** for Go and **12** for Node. 

1. Clone this project via
    
    $ git clone --depth=1 git@github.com:slaveofcode/gomodo.git

    or [download the whole project as ZIP](https://github.com/slaveofcode/gomodo/archive/refs/heads/main.zip)

2. Prepare the vue project via vite
 
   $ ./run.sh web:scaffold

3. Start development server for the Web
   
    $ ./run.sh web:start

4. Start backend server for API

    $ ./run.sh api:start

Now You a full have running development for both web and api

## Building the Web Static for Production

When You finished the web development and ready to publish. You will need to compile all the web sources with this command.

    $ ./run.sh web:build

## Build the Complete App

After the static resources was successfully compiled, hit this command to compile the project into a single binary.

    $ ./run.sh build

    The docker version also available via

    $ ./run.sh build:docker