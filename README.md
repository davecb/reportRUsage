#getRUsage
This is as simple as possible a resource usage function, for developers.

Used like:
```
func RunSomeExampleThing(f *os.File) {
        defer reportRUsage("ExampleThing", time.Now())
```

Described in https://leaflessca.wordpress.com/2017/12/03/getting-the-resources-you-just-used/

