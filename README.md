## go-build-something
![alt text](images/gopher.png "Gopher")
gopher it...

### prep :: idioms :: mental models
- don't need abstraction layers to decouple
- go provides thinnest possible layer between the concrete (hardware) and decoupled world
- prime focus the hardware is the platform
- go is based on a real machine
- you will know how the code will run on a machine
- every decision you make in go comes at a cost
- cost vs benefit is key in designing & architecting go software
- you can write programs that require minimal amount of resources to run
- champion efficiency, quality, and simplicity

### Productivity vs Performance
- instead of sacrificing performance to be productive, go allows for both
- it's not about being faster, but "fast enough"
- working with Go is intended to be fast
- it should take at most a few seconds to build a large executable on a single computer
- go enables us to take advantage of the hardware
- fosters an excellent balance between power and expressiveness, just like C
- do want you want to do by just programming in a straightforward manner
- it's all about the REAL machine, not VIRTUAL which makes for more accurate predictions
- Go has it all really, it's just about figuring out HOW to tap into it all

_"The hardware folks will not put more cores into their hardware if the software isn’t going to use them, so, it is this balancing act of each other staring at each other, and we are hoping that Go is going to break through on the software side.” - Rick Hudson (2015)_

_"Go blows away Node in pretty much every way IMO. You get all the niceties of blocking code without actually blocking, you get (relatively) tiny binaries that you can deploy anywhere with ease and no fuss, instead of 250mb of node_modules that everyone comes up with hacks to work around." - TJ Holowaychuk (2017)_

### Kit Projects
- single Kit projects should first be established
- then create multiple app projects for the different sets of programs to be deployed together
- a standard library, used by a company or org
- packages that belong to the Kit project should have highest levels of portability
- should be usable across multiple Application projects
- provide a very specific but foundational domain of functionality
- should not be allowed to have a vendor folder
- 3rd party packages should build against the latest version of those dependences

**Sample Kit:**
```
github.com/some-repo/kit
├── CONTRIBUTORS
├── LICENSE
├── README.md
├── cfg/
├── examples/
├── log/
├── pool/
├── tcp/
├── timezone/
├── udp/
└── web/
```

### future explorations, notes, repos, projects
- expressive concurrency primitives 
- error-handling
- stlib
- 3rd party packages
- scalable serverless apps
- webhook managers
- embedding binary data
- headers
- redirects
- elasticsearch
- encoding utils
- stripe api
- analytics
- binary updates

_"...performance means nothing if your application is frail, difficult to debug, refactor and develop." - TJ Holowaychuk (2014)_

