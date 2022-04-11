# go



## Use cases  

Concurrency  
Scalability  
Goroutines  
Light weight Containers  

---
Software architects use various inter-service messaging protocols according to their architectural requirements — some software development teams implement RESTful inter-service communications, while other teams implement asynchronous, messaging-based inter-service communications using message brokers like RabbitMQ.


Most importantly, you can use Golang for system programming, large-scale distributed systems, and highly scalable network applications and servers. It also finds use in cloud-based development, web app development, and big data or machine learning applications.



Scalability and Concurrency   


Go uses goroutines, which allows reliable and easy execution of threads and can be performed concurrently in a smooth manner. These goroutines make Go a perfect scalable programming language.  

Go can process more than 1,000 requests per second using concurrency. This feature alone makes Go superior to Node.js in terms of scalability and concurrency. It’s also worth noting that Node.js is a single-threaded asynchronous JavaScript engine.  

In Node.js single-threaded architecture, CPU-bound tasks sometimes block the event loop and slow down your program. As a result, you get a slow app and annoyed users.  


Node.js traditionally handles errors using the try-catch technique of handling exceptions, where errors are caught just when they occur, and developers can debug errors quickly and faster.  

Golang separates compile-time and runtime errors differently. This inconsistency causes confusion between developers and has led to a standard process in handling exceptions.  

However, Go developers think there will be more improvements on the language with the upcoming Go 2 version, including better error handling, error values, and generics.  



---

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
[https://www.golangprograms.com/](https://www.golangprograms.com/)  
[https://yourbasic.org/golang/for-loop/](https://yourbasic.org/golang/for-loop/)  
[https://yourbasic.org/golang/do-while-loop/](https://yourbasic.org/golang/do-while-loop/)  
[https://grpc.io/](https://grpc.io/)  
[https://developers.google.com/protocol-buffers/docs/proto3](https://developers.google.com/protocol-buffers/docs/proto3)  
[https://eng.uber.com/go-geofence-highest-query-per-second-service/](https://eng.uber.com/go-geofence-highest-query-per-second-service/)  
[Protocol Buffers Crash Course](https://youtu.be/46O73On0gyI)  
[https://github.com/protocolbuffers/protobuf/releases](https://github.com/protocolbuffers/protobuf/releases)  
[gRPC Crash Course - Modes, Examples, Pros & Cons and more](https://youtu.be/Yw4rkaTc0f8)  
[https://goswagger.io/](https://goswagger.io/)  
[Go and Node.js: a comparison – Nathan Youngman](https://youtu.be/N8Fc7RQfmdU)  
[https://blog.logrocket.com/building-microservices-go-gin/](https://blog.logrocket.com/building-microservices-go-gin/)  
[https://medium.com/@gauravsingharoy/asynchronous-programming-with-go-546b96cd50c1](https://medium.com/@gauravsingharoy/asynchronous-programming-with-go-546b96cd50c1)  
[https://mertcanarguc.medium.com/go-golang-with-async-await-a0c2e259a668](https://mertcanarguc.medium.com/go-golang-with-async-await-a0c2e259a668)  
[https://hackernoon.com/asyncawait-in-golang-an-introductory-guide-ol1e34sg](https://hackernoon.com/asyncawait-in-golang-an-introductory-guide-ol1e34sg)  
[https://levelup.gitconnected.com/use-go-channels-as-promises-and-async-await-ee62d93078ec](https://levelup.gitconnected.com/use-go-channels-as-promises-and-async-await-ee62d93078ec)  
[https://softwareengineeringdaily.com/2021/03/03/why-we-switched-from-python-to-go/](https://softwareengineeringdaily.com/2021/03/03/why-we-switched-from-python-to-go/)  
[https://getstream.io/blog/building-a-performant-api-using-go-and-cassandra/](https://getstream.io/blog/building-a-performant-api-using-go-and-cassandra/)  

[https://www.blog.duomly.com/golang-course-with-building-a-fintech-banking-app-lesson-5-bank-transactions-part-2/](https://www.blog.duomly.com/golang-course-with-building-a-fintech-banking-app-lesson-5-bank-transactions-part-2/)[https://www.freecodecamp.org/news/learn-go-by-building-11-projects/](https://www.freecodecamp.org/news/learn-go-by-building-11-projects/)  
[Learn Go Programming by Building 11 Projects – Full Course](https://youtu.be/jFfo23yIWac)  


[https://stackoverflow.com/questions/46809499/why-cant-you-put-a-variable-as-a-multidimensional-array-size-in-go](https://stackoverflow.com/questions/46809499/why-cant-you-put-a-variable-as-a-multidimensional-array-size-in-go)  
[https://linuxhint.com/golang-make-function/](https://linuxhint.com/golang-make-function/)  
[https://www.godesignpatterns.com/2014/05/arrays-vs-slices.html](https://www.godesignpatterns.com/2014/05/arrays-vs-slices.html)  
[https://go.dev/blog/slices-intro](https://go.dev/blog/slices-intro)  
[https://stackoverflow.com/questions/45423667/what-is-the-point-in-setting-a-slices-capacity](https://stackoverflow.com/questions/45423667/what-is-the-point-in-setting-a-slices-capacity)  
[https://stackoverflow.com/questions/36706843/how-to-get-the-underlying-array-of-a-slice-in-go](https://stackoverflow.com/questions/36706843/how-to-get-the-underlying-array-of-a-slice-in-go)  
[https://www.developer.com/languages/understanding-structs-in-go/](https://www.developer.com/languages/understanding-structs-in-go/)  
[https://technobeans.com/2019/01/27/golang-composite-data-types-arrays-and-slices/](https://technobeans.com/2019/01/27/golang-composite-data-types-arrays-and-slices/)  
[https://go.dev/src/runtime/slice.go](https://go.dev/src/runtime/slice.go)  
[https://go.dev/src/runtime/malloc.go](https://go.dev/src/runtime/malloc.go)  
[https://golang.hotexamples.com/examples/malloc/-/Free/golang-free-function-examples.html](https://golang.hotexamples.com/examples/malloc/-/Free/golang-free-function-examples.html)  
[https://medium.com/@ankur_anand/a-visual-guide-to-golang-memory-allocator-from-ground-up-e132258453ed](https://medium.com/@ankur_anand/a-visual-guide-to-golang-memory-allocator-from-ground-up-e132258453ed)  
[https://chris124567.github.io/2021-06-21-go-performance/](https://chris124567.github.io/2021-06-21-go-performance/)  
[https://www.codestudyblog.com/cs2112goa/1213221611.html](https://www.codestudyblog.com/cs2112goa/1213221611.html)  
[https://stackoverflow.com/questions/9320862/why-would-i-make-or-new](https://stackoverflow.com/questions/9320862/why-would-i-make-or-new)  
[https://medium.com/codex/learn-how-golang-channels-work-by-building-them-72f49ed30f2c](https://medium.com/codex/learn-how-golang-channels-work-by-building-them-72f49ed30f2c)  
[https://softwareengineering.stackexchange.com/questions/341152/type-safety-go-vs-c-pointers](https://softwareengineering.stackexchange.com/questions/341152/type-safety-go-vs-c-pointers)  
[https://www.pullrequest.com/blog/golang-vs-c-arrays/](https://www.pullrequest.com/blog/golang-vs-c-arrays/)  

