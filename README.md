# thomas' mini build server #

This project is written in Go. It sits somewhere between being a continuous integration/build server. 

### What? ###

This is a super simple CI server. It listens for POST hooks from bitbucket/github.
When a POST comes in the build process is started. The build process uses a Dockerfile
to build a container image capable of running a test suite against your project. 

Build process workflow:

1. Checkout the latest commit
2. Build docker image which can run your tests
3. Run test suite in a new container based off the new image
4. Save the test results (pass/fail & ouput)
5. Clean up (remove docker image used to run your test)

### How? ###

Add the mini-build binary to your PATH.

Run the `setup` command to create the directory structure in your user's home directory:

```
➜  mini-build git:(master) ✗ ./mini-build 
Could not find an existing application configuration. Running setup.
Created directory: /home/thomas/.tmbs
Created directory: /home/thomas/.tmbs/repos
Created directory: /home/thomas/.tmbs/tests
Created file: /home/thomas/.tmbs/config.json

** Thomas' Mini Build Server **

  Commands:
    start   - Starts the build server
    add     - Add a repository
    help    - Display this message
```

Add a github/bitbucket repository to the mini-build config:

```
➜ mini-build add https://tommyvyo@bitbucket.org/tommyvyo/mini-build.git
Added https://tommyvyo@bitbucket.org/tommyvyo/mini-build.git to the mini-build configuration.
```

You're all set. Start mini-build and it'll listen for commit POST's from that project:

```
#!bash
➜ mini-build start
Thomas Mini Build Server - hit return to quit
 - loaded config.json
 - listening on 0.0.0.0 :59999
```

When a push notification is recieved TMBS will build a docker image based off a Dockerfile in the repository.

After the image is built the commited code is ran based off a command you specify, which then runs your test suite inside of a new container. 

After the push notification is recieved you can use TMBS to watch the progress, output, and results of each commit. 
