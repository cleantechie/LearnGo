<p>
bool:                    %t
int, int8 etc.:          %d
uint, uint8 etc.:        %d, %#x if printed with %#v
float32, complex64, etc: %g
string:                  %s
chan:                    %p
pointer:                 %p
</p>

Same as C
printf: prints o/p to console
sprintf: prints o/p to a string buffer
fprintf:  prints o/p to a file buffer

map[key-type]value-type
```go build``` -> generate the binary file
```go list -f '{{.Target}}' --?``` -> Navigate to the bin file 
```go install ``` -> we can run the bin without navigating to the binary file and then run the main go file
```go work edit``` edits the go.work file similarly to ```go mod edit```
```go work sync``` syncs dependencies from the workspace’s build list into each of the workspace modules.
``go work use [-r] [dir]`` adds a use directive to the go.work file for dir, if it exists, and removes the use directory if the argument directory doesn’t exist. The -r flag examines subdirectories of dir recursively.