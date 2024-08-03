# Awesome Go Scripting

Resources for writing scripts and makefiles in Go.

**Motivation**: ***Use the tools you understand best***

Writing shell scripts and makefiles can be frustrating for a few reasons:

* You often deal with them at the start of a project, but then only occasionally after that.
* If you write them infrequently, then you have to practically re-learn the syntax for each new project.
* They *feel* less powerful than your everyday programming language because it takes frequent practice to build mastery.

Writing scripts in your primary language (in this case, Go), has a number of benefits:

* Familiarity with the syntax.
* Familiarity with idiomatic techniques.
* Access to the full standard library, which opens up interesting new possibilities.

There are some downsides too:

* It moves the frustration, but might not eliminate it--teammates, like ops, may be much more familiar with bash than your programming language.
* The language tooling must be installed before the script can be run.
* Access to the full standard library, which creates a temptation to write "clever" code.

**Basic Principle**: ***Scripts are tooling***

Scripts don't add value to your product directly; instead, their primary purpose is to make you more productive and happy when you are adding real value.

* Writing a simple script should be quick and easy.
* Scripts should be very easy to understand.

## Contents

* [Awesome Go Scripting](#awesome-go-scripting)
	* [Contents](#contents)
	* [Examples](#examples)
	* [Libraries](#libraries)
	* [Tools](#tools)
		* [Improved `go run`](#improved-go-run)
	* [Resources](#resources)
		* [Blog Posts](#blog-posts)
		* [Discussions](#discussions)
		* [Videos](#videos)

## Examples

This repo contains several different examples demonstrating how to write scripts and makefiles in Go.

The examples start very simply, closely mimicking the style of shell scripts and makefiles. They increase in sophistication and include increasing bits of boilerplate to make them a bit more idiomatic.

Each example is intended to be a piece of code that you could simply copy and paste to your own repo.

* [0-hello-world](examples/0-hello-world/) - Demonstrates the basic idea.

## Libraries

* [script](https://github.com/bitfield/script) — Makes Go more shell-like.

## Tools

Programs, scripts, and other software to help with Go scripting.

## Improved `go run`

* [gorun](https://github.com/erning/gorun) — Supports shebang execution of scripts.
* [`go run` wrapper](https://gist.github.com/rkulla/bbe82f81fa1baa283a5fde2aec8fb5a9) — Implicitly convert "go file.go" to "go run file.go".

## Resources

Information about scripting in Go.

## Blog Posts

* [Scripting with Go](https://bitfieldconsulting.com/posts/scripting) - An introduction to using the `script` library.

## Discussions

* [Story: Writing Scripts with Go](https://gist.github.com/posener/73ffd326d88483df6b1cb66e8ed1e0bd) — Explores methods of making Go programs directly executable in the shell (i.e. `./main.go` instead of `go run main.go`) with a shebang.
* [StackOverflow Go Shebang Discussion](https://stackoverflow.com/questions/7707178/whats-the-appropriate-go-shebang-line)

## Videos

* [Code Club: Script](https://www.youtube.com/watch?v=6S5EqzVwpEg) - A video introduction to the `script` library.
