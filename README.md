# Motivation
This repo contains a somewhat curated collection of practical Go code and concepts (at least to me). I intend to expand this into something more comprehensive.
# Contexts
- docs: https://pkg.go.dev/context
- usage: carrying deadlines, cancellaton signals and other request-scoped values across API boundaries and between processes
- use cases: 
    - `Monitoring` requests across different layers of application with same request ID etc
    - `Deadlines` for requests where you want to timeout after a certain amount of time. This guards against infinite loops / spawning several child processes
- best practice: 
    - Pass context as `first argument` in a function
        ```go
        func doSomething(ctx context.Context, arg1 string, arg2 string) error {
            // do something
        }
        ```
- reference: https://www.youtube.com/watch?v=h2RdcrMLQAo

# Channels
- Unbuffered channels
    - `synchronous` communication
    - `blocking` communication
    - `deadlock` if no receiver is present
- Buffered channels
    - has a capacity
    - `asynchronous` communication
    - `non-blocking` communication
    - `deadlock` if buffer is full and no receiver is present