
## Cocurrency
- Cocurrency: is the notion of multiple things happening at the same time or potential for multiple processes to be in progress at the same time.
- Parallelism bs Concurrency is not the same thing

## Go Routines

- considered as lightweight thread that has sep independent execution
- can execute concurrently with other go routines
- entirely managed by go runtimes
- Main Function and Main package is the main go rounine. All go routines start from main go routines
- There is no Parent/Child relationships with GoRoutines
- Output can be non-deterministic as each go routines runs independently