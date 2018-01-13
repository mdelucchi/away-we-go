## goNAV
![alt text](images/gopher.png "Gopher")
packages on radar for NAV dev...

**fmt:**
- implements formatted I/O with functions
- analogous to C's printf and scanf
- format 'verbs' are derived from C's but are simpler

**net/http:**
- provides HTTP client and server implementations
- Get, Head, Post, and PostForm make HTTP (or HTTPS) requests:

**io:**
- provides basic interfaces to I/O primitives
- wrap existing implementations of such primitives
- abstract the functionality, such as package os, plus some other related primitives
- wrap lower-level operations with various implementations