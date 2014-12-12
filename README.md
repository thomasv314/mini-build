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


### Why? ###

I started this project because I wanted a continuous integration server to run the test suite of a Rails
application after each commit. I played around with setting up Jenkins and Hudson and found that both had 
more features than I knew what to do with. I love Docker and wanted to learn more about the API and 
writing software in Go, so I figured this would be a good start. 


### How? ###

Add the mini-build binary to your PATH.

Run the `setup` command to create the directory structure in your user's home directory:

```
➜ mini-build setup
Setting up /home/thomas
Created directory: /home/thomas/.tmbs
Created directory: /home/thomas/.tmbs/repos
Created directory: /home/thomas/.tmbs/tests
Good to go!
```

Add a github/bitbucket repository to the mini-build config:

```
➜ mini-build add https://tommyvyo@bitbucket.org/tommyvyo/mini-build.git
Added https://tommyvyo@bitbucket.org/tommyvyo/mini-build.git to the mini-build configuration.
```




You're all set. Start mini-build and it'll listen for commit POST's from that project:

```
➜ mini-build start
Thomas' Mini Build Server - hit return to quit
 - loaded config.json
 - listening on 0.0.0.0 :59999
```