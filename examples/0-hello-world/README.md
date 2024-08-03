# Hello World Go Scripting Example

Demonstrate the basic principle by writing the simplest Go script.

* Try it with `go run main.go build`.
* Other commands are `run`, `test`, and `clean`.

## Notes

There is lots of stuff wrong with this, but that's okay for now.

* It builds and runs itself, which is not useful.
	* Maybe the ability to compile the script into an executable will be useful later?
* No helpful prompts.
* Panics if no command specified.
* Only runs one process per command.
	* Can't chain commands (e.g. `run` can't run `build` first).
* `run` invokes itself recursively.
