## go-build-something
exploring Go :tractor:

### prep :: idioms :: the _Go_ state of mind 
- don't need abstraction layers to decouple
- go provides thinnest possible layer between the concrete (hardware) and decoupled world
- prime focus the hardware is the platform
- go is based on a real machine
- you will know how the code will run on a machine
- every decision you make in go comes at a cost
- cost vs benefit is key in designing & architecting go software
- you can write programs that require minimal amount of resources to run
- champion efficiency, quality, and simplicity

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

