strategy-text:
	go run ./behavioral-patterns/strategy/cli/main.go --output=text

strategy-image:
	go run ./behavioral-patterns/strategy/cli/main.go --output=image

command:
	go run ./behavioral-patterns/command/cli/main.go

state:
	go run ./behavioral-patterns/state/cli/main.go


goroutine:
	go run ./concurrency-introduction/goroutine/main.go

waitgroup:
	go run ./concurrency-introduction/waitgroup/main.go

callback:
	go run ./concurrency-introduction/callbacks/main.go

mutex:
	go run -race ./concurrency-introduction/mutex/cli/main.go

channel:
	go run -race ./concurrency-introduction/channels/main.go

buffered-channel:
	go run -race ./concurrency-introduction/buffered-channels/main.go

directional-channel:
	go run -race ./concurrency-introduction/directional-channels/main.go

select-statement:
	go run -race ./concurrency-introduction/select-statement/main.go

channel-range:
	go run -race ./concurrency-introduction/range/main.go


