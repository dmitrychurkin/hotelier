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