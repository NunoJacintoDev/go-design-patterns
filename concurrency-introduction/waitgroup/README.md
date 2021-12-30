# WaitGroups
WaitGroup comes in the synchronization package (the sync package) to help us synchronize many concurrent Goroutines. Every time we have to wait for one Goroutine to finish, we add 1 to the group, and once all of them are added, we ask the group to wait. When the Goroutine finishes, it says Done and the WaitGroup will take one from the group.

# Run example
``make waitgroups``