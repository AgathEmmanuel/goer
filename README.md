# go


## Steps

> Basics

- Filenames, keywords, identifiers
- Operators, types, functions, and constants
- Pointers, structures, methods
- Maps, arrays, slices
- Go CLI
- Interface
- Error handling
- Goroutine, Channel, Buffer
- Panic, Defer, Error, Recover

> Libraries and tools

- Go dependency management tools
- Semantic versioning
- Scripts and repositories
- Go libraries
- SQL fundamentals
- GIT
- Basic authentication
- HTTP/HTTPS
- Web frameworks and routers
- Relational databases (PostgreSQL)

> Testing

- Unit testing
- Integration testing
- Behavior testing
- E2E testing

> Design patterns

- Structural
- Creational
- Behavioral
- Concurrency
- Stability

> Portfolio


- Completing online courses on Go
- Contributing to open source Go projects
- Building Go projects from scratch
- Implementing Go units in existing projects
- Coding classic algorithm problems with Go
- Completing courses on adjacent technologies (i.e. SQL)

> Interview questions


- What is a goroutine? How do you stop it?
- How do can check variable type at runtime?
- How do you format a string without printing?
- How do you concatenate strings in Go?
- What is Go 2?
- How do you initialize a struct in Go?




### Commands  

```

mkdir app1
go mod init example/app1

go get 
go get github.com/gin-gonic/gin			# to install gin package  


go get [-d] [-f] [-t] [-u] [-v] [-fix] [-insecure] [build flags] [packages]


go help get 

go list ./...  			# To list packages in your workspace, go to your workspace folder and run this command:


go list ...			# List All Packages




Packages and their Dependencies

# If you also want to see the imported packages by each package, you can try this custom format:

go list -f "{{.ImportPath}} {{.Imports}}" ./...		

# -f specifies an alternate format for the list, using the syntax of package template. The struct whose fields can be referenced can be printed by the go help list command.


# If you want to see all the dependencies recursively (dependencies of imported packages recursively), you can use this custom format:

go list -f "{{.ImportPath}} {{.Deps}}" ./...

But usually this is a long list and just the single imports ("{{.Imports}}") of each package is what you want.





Found it https://go.dev/blog/using-go-modules#removing-unused-dependencies

go mod tidy

So basically, once the package is not being imported in any package you can perform a go mod tidy and it will safely remove the unused dependencies.

And if you are vendoring the dependencies, then run the command below to make the module changes be applied in the vendor folder:

go mod vendor



>> Creating a helloword app


$ mkdir golang-helloworld
$ cd golang-helloworld


$ go mod init github.com/agathemmanuel/golang-helloworld

======== main.go
package main

import "fmt"

func main() {
	fmt.Println("Hello, world.")
}

======== 


Build and run your module.

 go build
$ go run main.go
>> Hello, world.


Formatting source code.

$ go fmt
>> main.go


Install a dependency.

$ go get -u github.com/gorilla/mux
$ go install github.com/gorilla/mux


go get installs the source for gorilla/mux to our /go/bin directory. 
The -u flag we pass is an "update" flag, which we use to grab the latest version just in case.

see how go.mod was affected by running $ cat go.mod



======== main.go

package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
	"io"
)

======== 

```



## Links
[https://learn.gopherguides.com/courses/preparing-your-environment-for-go-development/modules/setting-up-mac-linux/#slide-8](https://learn.gopherguides.com/courses/preparing-your-environment-for-go-development/modules/setting-up-mac-linux/#slide-8)  
[https://www.educative.io/blog/become-golang-developer](https://www.educative.io/blog/become-golang-developer)
[https://www.educative.io/blog/the-7-most-important-software-design-patterns](https://www.educative.io/blog/the-7-most-important-software-design-patterns)
[https://www.educative.io/blog/google-coding-interview](https://www.educative.io/blog/google-coding-interview)
[https://www.educative.io/blog/google-coding-interview-questions](https://www.educative.io/blog/google-coding-interview-questions)
[https://www.educative.io/blog/behavioral-interviews-how-to-prepare-and-ace-interview-questions](https://www.educative.io/blog/behavioral-interviews-how-to-prepare-and-ace-interview-questions)
[https://dev.to/ankit01oss/7-github-projects-to-make-you-a-better-go-developer-2nmh](https://dev.to/ankit01oss/7-github-projects-to-make-you-a-better-go-developer-2nmh)
[https://www.wolfe.id.au/2020/03/10/starting-a-go-project/](https://www.wolfe.id.au/2020/03/10/starting-a-go-project/)  
[https://medium.com/@yussufshaikh/installing-go-modules-from-github-repository-5e381cbd5683](https://medium.com/@yussufshaikh/installing-go-modules-from-github-repository-5e381cbd5683)  
[https://hackersandslackers.com/create-your-first-golang-app/](https://hackersandslackers.com/create-your-first-golang-app/)  
[https://www.eternaldev.com/blog/adding-and-removing-dependency-in-go-modules/](https://www.eternaldev.com/blog/adding-and-removing-dependency-in-go-modules/)  

