"github.com/joho/godotenv"
var _ = godotenv.Load()
GO dependency managment CLI tool
	- https://golang.github.io/dep/docs/installation.html
	- https://github.com/golang/dep/issues/2098
? golang setting environment variables
	- https://gobyexample.com/environment-variables
	- https://golangcode.com/get-and-set-environment-variables/
	- https://dev.to/craicoverflow/a-no-nonsense-guide-to-environment-variables-in-go-a2f
	- https://stackoverflow.com/questions/24873883/organizing-environment-variables-golang
	- http://peter.bourgon.org/go-in-production/
? golang read env file
	- https://medium.com/@felipedutratine/manage-config-in-golang-to-get-variables-from-file-and-env-variables-33d876887152
	- https://www.reddit.com/r/golang/comments/1l6ny2/loads_environment_variables_from_env_file_in_go/
	- https://github.com/joho/godotenv

docker commands / issues:

? error during connect: Get https://192.168.99.100:2376/v1.37/version: dial tcp 192.168.99.100:2376: connectex:
	docker-machine restart -> docker-machine env
	- https://github.com/docker/for-win/issues/2596
	- https://docs.docker.com/machine/reference/env/
	- https://docs.docker.com/engine/reference/commandline/rm/
? how to stop docker container
	- https://www.shellhacks.com/docker-stop-container/ --> docker stop 72ca2488b353
	!!!!- https://linuxize.com/post/how-to-remove-docker-images-containers-volumes-and-networks/
		docker system prune
		docker system prune --volumes
		docker container (volume / image) ls -a
	!!!!! - https://www.prisma.io/forum/t/solved-getting-started-on-linux-prisma-deploy-fails/5393/4
	!!!!! - https://www.prisma.io/forum/t/solved-getting-started-on-linux-prisma-deploy-fails/5393/5
	!!!!!! docker-compose down -v --rmi all --remove-orphans !!!!!!!

? docker container stop all running
	- https://coderwall.com/p/ewk0mq/stop-remove-all-docker-containers
		docker stop $(docker ps -a -q)
		docker rm $(docker ps -a -q)
	
? docker compose define env file
	!!!!!!! - https://docs.docker.com/compose/environment-variables/	

? how to install dep on windows
	- https://github.com/golang/dep/issues/2086

? golang string concatenation
	- https://www.geeksforgeeks.org/different-ways-to-concatenate-two-strings-in-golang/

? golang encrypt password
	- https://medium.com/@jcox250/password-hash-salt-using-golang-b041dc94cb72
	- https://gowebexamples.com/password-hashing/

? prisma vue js !!!!!!!!!!!!!!!!!!!!!!!!!!!
	- https://github.com/ammezie/techies/blob/master/server/.env.example !!!!!
	- https://github.com/graphql-boilerplates/vue-fullstack-graphql/blob/master/advanced !!!!!
	- https://blog.pusher.com/fullstack-graphql-app-prisma-apollo-vue/

? docker double dot
	- https://stackoverflow.com/questions/49449012/dot-and-colon-meaning/49449160

? Could not connect to server at http://localhost:4466. Please check if your server is running
	- https://www.prisma.io/forum/t/could-not-connect-to-server-at-http-localhost-4466-please-check-if-your-server-is-running/4062
	? localhost 4466 => https://github.com/prisma/prisma/issues/4021
	!!!!! - https://stackoverflow.com/questions/51334907/prisma-deploy-docker-error-could-not-connect-to-server
	docker-machine create default
	docker-machine ip default
	docker-compose up -d
	docker ps
	docker-compose logs
	- https://stackoverflow.com/questions/55299233/prisma-error-could-not-connect-to-server-at-http-localhost4466

? SQLException occurred while connecting to postgres:5432 prisma
	- https://stackoverflow.com/questions/56254196/connection-error-accessing-postgres-docker-container

? ERROR: No cluster could be found for workspace '*' and cluster 'default'
	!!!! - https://www.prisma.io/forum/t/error-no-cluster-could-be-found-for-workspace-and-cluster-default/6402/3 -> Please use 1.28.3 for now.
	- https://github.com/prisma/prisma/issues/4215

? golang convert slice of structs into slice of pointers
	- https://github.com/golang/go/issues/22791 
This:

for _, v := range a {
	b = append(b, &v)
}
and this:

for i := 0; i < len(a); i++ {
	b = append(b, &a[i])
}

? how to create error in golang
	- https://yourbasic.org/golang/create-error/


? model declarations have to be indicated with the model keyword. prisma
	- https://prisma-docs.netlify.com/docs/1.4/reference/service-configuration/prisma.yml/using-variables-nu5oith4da
? Model declarations have to be indicated with the `model` keyword vs code
	- https://github.com/prisma/prisma/blob/master/docs/1.14/04-Reference/02-Service-Configuration/03-Data-Model/01-Data-Modelling-(SDL).md
	- https://spectrum.chat/prisma/general/datamodel-prisma-file-has-errors-and-i-never-touched-the-code~e1d1a8fe-24dd-412b-bfbb-ecef6ada432e

? golang +build ignore => https://golang.org/pkg/go/build/

https://gqlgen.com/getting-started/

? go:generate => https://blog.golang.org/generate

? prisma many to many
	- https://stackoverflow.com/questions/53798509/prisma-data-modeling-has-many-and-belongs-to
	- https://www.prisma.io/forum/t/how-do-i-define-a-many-to-many-relationships-in-prisma/3129/7
	- https://dba.stackexchange.com/questions/151904/mapping-many-to-many-relationship
	- https://stackoverflow.com/questions/2923809/many-to-many-relationships-examples

FE:
? react-fullstack-graphql
	- https://blog.apollographql.com/full-stack-react-graphql-tutorial-582ac8d24e3b
	- https://www.prisma.io/docs/tutorials/bootstrapping-boilerplates/react-(fullstack)-tijghei9go#step-2:-bootstrap-your-react-fullstack-app
	- https://github.com/graphql-boilerplates/react-fullstack-graphql/tree/master/advanced/src/components
	- https://github.com/apollographql/react-apollo/blob/master/examples/hooks/client/src/NewRocketForm.tsx
? angular apollo
	- https://www.apollographql.com/docs/angular/basics/caching/
	- https://github.com/arjunyel/angular-apollo-example/tree/master/frontend/src/app
	- https://github.com/Quramy/apollo-angular-example/blob/master/src/graphql/graphql.module.ts
	- https://www.techiediaries.com/graphql-tutorial/

? docker postgres query from command line
	- https://github.com/Radu-Raicea/Dockerized-Flask/wiki/%5BDocker%5D-Access-the-PostgreSQL-command-line-terminal-through-Docker
	- https://stackoverflow.com/questions/37099564/docker-how-can-run-the-psql-command-in-the-postgres-container

? prisma Authentication token is invalid: Failed to parse token data
	- https://github.com/prisma/prisma/issues/4492 !!!
	- https://github.com/prisma/prisma/issues/3502

? graphql file upload react
	- https://github.com/jaydenseric/apollo-upload-client
	- https://blog.apollographql.com/file-uploads-with-apollo-server-2-0-5db2f3f60675
	- https://github.com/jaydenseric/apollo-upload-examples/tree/master/app/components
	- https://github.com/prisma-labs/graphql-yoga/blob/master/examples/file-upload/index.ts

? prisma file upload
	- https://www.reddit.com/r/graphql/comments/ane9rp/how_to_send_images_over_graphql/

? prisma graphql file upload
	- https://www.prisma.io/forum/t/what-is-the-best-way-to-upload-and-retrieve-images/7797/2
	- https://manticarodrigo.com/file-handling-s3-prisma-graphql-yoga/ => https://github.com/manticarodrigo/prisma-s3/blob/master/server/src/index.js
	- https://www.prisma.io/forum/t/how-to-use-prisma-with-apollo-server/4412/3
	- https://github.com/graphql-boilerplates
	- https://medium.com/@andrecoetzee153/putting-together-prisma-apollo-server-2-part-1-82f9a94e6794

? golang file upload graphql
	- https://stackoverflow.com/questions/57431027/upload-multiple-files-to-backend-go-api
	- https://github.com/smithaitufe/go-graphql-upload
	- https://medium.com/@vcomposieux/handle-file-uploads-using-a-graphql-middleware-11914ba05bfc
	- https://github.com/graphql-go/graphql/issues/141
	- https://github.com/jpascal/graphql-upload/blob/86c9aa31749a2b5e8cf63e8eafb8d3b37566f810/handler.go

? upload scalar graphql
	- https://www.prisma.io/forum/t/scalar-upload-with-prisma/7114/4
	- https://github.com/jaydenseric/graphql-upload
	- https://levelup.gitconnected.com/how-to-add-file-upload-to-your-graphql-api-34d51e341f38
	- https://moonhighway.com/how-the-upload-scalar-works
	- https://graphql-modules.com/docs/recipes/file-uploads

? go build filename
	- go build -o build/wiki wiki.go
	- https://medium.com/rungo/the-anatomy-of-functions-in-go-de56c050fe11
	- https://golang.org/cmd/go/#hdr-Modules__module_versions__and_more

? golang backtick string
	- https://stackoverflow.com/questions/30681054/what-is-the-usage-of-backtick-in-golang-structs-definition

? golang get all headers from request
	- https://stackoverflow.com/questions/47557304/how-to-obtain-all-request-headers-in-go
	- https://stackoverflow.com/questions/46021330/how-can-i-read-a-header-from-an-http-request-in-golang/46022272
	- https://stackoverflow.com/questions/12830095/setting-http-headers

? golang type assertion
	- https://yourbasic.org/golang/type-assertion-switch/

? gqlgen get headers from context
	- https://github.com/99designs/gqlgen/issues/262

? how to debug golang
	- https://scotch.io/tutorials/debugging-go-code-with-visual-studio-code

? graphql query without arguments
	- https://stackoverflow.com/questions/52769993/how-to-correctly-declare-a-graphql-query-without-parameters
	- https://stackoverflow.com/questions/44737043/not-returning-data-from-graphql-mutation

? You are calling concat on a terminating link, which will have no effect
	- https://stackoverflow.com/questions/51840201/apollo-you-are-calling-concat-on-a-terminating-link-which-will-have-no-effect
	- https://github.com/apollographql/apollo-link/issues/362

? CORS
	- https://stackoverflow.com/questions/18642828/origin-origin-is-not-allowed-by-access-control-allow-origin

? apollo response headers
	- https://stackoverflow.com/questions/47443858/apollo-link-response-headers
	- https://github.com/apollographql/apollo-client/issues/2514

? golang set headers after middleware
	- https://golangcode.com/middleware-on-handlers/

? gqlgen DateTime
	- https://github.com/99designs/gqlgen/issues/575

? how to send mutation in graphql playground
	- https://github.com/graphql/graphiql/issues/72

? regex email golang
	- https://www.golangprograms.com/regular-expression-to-validate-email-address.html

? golang i18n support
	- https://github.com/nicksnyder/go-i18n

? golang rate limiter middleware
	- https://www.alexedwards.net/blog/how-to-rate-limit-http-requests
	- https://github.com/didip/tollbooth_gin

? gqlgen rate limiting
	- https://go.libhunt.com/limiter-alternatives
	- https://github.com/ulule/limiter

	- https://www.freecodecamp.org/news/deep-dive-into-graphql-with-golang-d3e02a429ac3/
	- http://marcio.io/2015/07/handling-1-million-requests-per-minute-with-golang/

? gqlgen get request from ctx
	- https://github.com/99designs/gqlgen/issues/765

? golang gin redirect from trailing slash
	- https://github.com/gin-gonic/gin/issues/915
	- https://github.com/gin-gonic/gin/issues/482
	- https://github.com/gin-gonic/gin/issues/349

? golang slice methods
	- https://blog.golang.org/go-slices-usage-and-internals
	- https://www.geeksforgeeks.org/slices-in-golang/
	- https://stackoverflow.com/questions/22535775/how-to-get-the-last-element-of-a-slice
	!!!!! - https://github.com/golang/go/wiki/SliceTricks !!!!!!

? golang get token from authorization header bearer
	- https://stackoverflow.com/questions/39518237/how-to-extract-and-verify-token-sent-from-frontend-in-golang

? golang get domain name
	- https://stackoverflow.com/questions/16512840/get-domain-name-from-ip-address-in-go
	- https://stackoverflow.com/questions/23558425/how-do-i-get-the-local-ip-address-in-go
	- https://stackoverflow.com/questions/16512840/get-domain-name-from-ip-address-in-go

? how to set httpOnly cookie for cors
	- https://stackoverflow.com/questions/46288437/set-cookies-for-cross-origin-requests
	- https://medium.com/@_graphx/if-httponly-you-could-still-csrf-of-cors-you-can-5d7ee2c7443
	- https://stackoverflow.com/questions/36365409/setting-cookies-with-cors-requests
	- https://flaviocopes.com/golang-enable-cors/

? golang get set cookie
	- https://stackoverflow.com/questions/54275704/how-to-read-cookies-from-golang/54276297
	- https://www.golangprograms.com/get-set-and-clear-session-in-golang.html
	- https://www.socketloop.com/tutorials/golang-read-write-create-and-delete-cookie-example
	- https://stackoverflow.com/questions/12130582/setting-cookies-with-net-http
	- https://golangcode.com/add-a-http-cookie/

? golang generate hash of string
	- https://stackoverflow.com/questions/13582519/how-to-generate-hash-number-of-a-string-in-go
	- https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
	- https://gist.github.com/dopey/c69559607800d2f2f90b1b1ed4e550fb
	- https://stackoverflow.com/questions/45267125/how-to-generate-unique-random-alphanumeric-tokens-in-golang

? gin gonic cookies
	- https://stackoverflow.com/questions/27234861/correct-way-of-getting-clients-ip-addresses-from-http-request
	- https://github.com/gin-gonic/gin/issues/165
	- https://gin-gonic.com/docs/examples/cookie/

	- https://stackoverflow.com/questions/39492468/how-to-save-time-in-the-database-in-go-when-using-gorm-and-postgresql

? golang mailgun
	- thepolyglotdeveloper.com/2017/12/send-emails-mailgun-using-golang/
	- https://tutorialedge.net/golang/sending-email-using-go-and-mailgun/
	- !!!! https://medium.com/@durgaprasadbudhwani/sending-complex-html-email-with-mailgun-api-using-go-language-cfc5338a5b70
	- !!!! https://github.com/Golang-Coach/Lessons/blob/master/GoMailer/mailer.go

? golang url builder
 - https://stackoverflow.com/questions/26984420/url-builder-query-builder-in-go
 - https://stackoverflow.com/questions/23151827/how-to-get-url-in-http-request
 - https://stackoverflow.com/questions/26984420/url-builder-query-builder-in-go

? golang valid domain name
- https://stackoverflow.com/questions/48239113/how-to-check-whether-hostname-is-domain-name-in-go
- https://gist.github.com/chmike/d4126a3247a6d9a70922fc0e8b4f4013
- https://godoc.org/github.com/asaskevich/govalidator

? golang backtick string
- https://stackoverflow.com/questions/46917331/what-is-the-difference-between-backticks-double-quotes-in-golang/46925368
- https://stackoverflow.com/questions/30681054/what-is-the-usage-of-backtick-in-golang-structs-definition
- https://stackoverflow.com/questions/10858787/what-are-the-uses-for-tags-in-go

? golang tag params
- https://medium.com/golangspec/tags-in-golang-3e5db0b8ef3e

? compare two strings
- https://stackoverflow.com/questions/34383705/how-do-i-compare-strings-in-golang

? nil pointer detection
- https://stackoverflow.com/questions/20240179/nil-detection-in-go
- https://www.golangprograms.com/how-to-check-pointer-or-interface-is-nil.html

? golang string interpolation
- https://stackoverflow.com/questions/11123865/format-a-go-string-without-printing
- https://stackoverflow.com/questions/50095616/go-string-interpolation
- https://yourbasic.org/golang/fmt-printf-reference-cheat-sheet/#sprintf-format-without-printing

? apollo delete type from cache
- https://medium.com/@martinseanhunt/how-to-invalidate-cached-data-in-apollo-and-handle-updating-paginated-queries-379e4b9e4698
- https://stackoverflow.com/questions/48596265/deleting-apollo-client-cache-for-a-given-query-and-every-set-of-variables

- https://stackoverflow.com/questions/40541994/multiple-path-names-for-a-same-component-in-react-router

? golang assign value to pointer
- https://stackoverflow.com/questions/43811652/golang-assigning-a-value-to-struct-member-that-is-a-pointer
- https://medium.com/rungo/pointers-in-go-a789eafccd53
- https://www.callicoder.com/golang-pointers/

? golang after request middleware
- https://drstearns.github.io/tutorials/gomiddleware/
- https://stackoverflow.com/questions/43976539/run-middleware-after-gorilla-mux-handling