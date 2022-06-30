# shm
Accessing system shared memory from Go applications
## How to use
```Go
var hostState unsafe.Pointer
...

// HostState connect to shared memory to retrieve the host state
func HostState() (unsafe.Pointer, error) {
	segment, err := shm.OpenSegment()

	if err != nil {
		return nil, err
	}

	return segment.Attach()
}

// IsCurrentHostStateActive checking current host state. Server sends data only in the active state.
func IsCurrentHostStateActive() bool {
	if hostState == nil {
		return true
	}

	if *(*uint16)(hostState) != 2 {
		return true
	}

	return false
}

...

hostState, err = HostState()
	if err != nil {
		log.Errorf("Failed to get host state shared memory: %v", err)
	}

	if IsCurrentHostStateActive() {
    ...
```
