# README.md

## Example using custom handlers

    swagger-codegen generate  -l go-server -i  ./swagger.yaml \
        && rm go/api_default.go && cp my_lib/swagger/release.go go/ \
        && cp my_lib/main.go main.go

        INFO  i.s.codegen.v3.AbstractGenerator - writing file /tmp/demo/./go/model_custom_error.go
        INFO  i.s.codegen.v3.AbstractGenerator - writing file /tmp/demo/./go/model_default_response.go
        INFO  i.s.codegen.v3.AbstractGenerator - writing file /tmp/demo/./go/model_release.go
        INFO  i.s.codegen.v3.AbstractGenerator - writing file /tmp/demo/./go/model_release_release.go
        INFO  i.s.codegen.v3.AbstractGenerator - writing file /tmp/demo/./go/model_release_repository.go
        INFO  i.s.codegen.v3.AbstractGenerator - writing file /tmp/demo/./go/api_default.go
        INFO  i.s.codegen.v3.AbstractGenerator - writing file /tmp/demo/./api/swagger.yaml
        INFO  i.s.codegen.v3.DefaultGenerator - writing file /tmp/demo/./Dockerfile
        INFO  i.s.codegen.v3.AbstractGenerator - writing file /tmp/demo/./main.go
        INFO  i.s.codegen.v3.AbstractGenerator - writing file /tmp/demo/./go/routers.go
        INFO  i.s.codegen.v3.AbstractGenerator - writing file /tmp/demo/./go/logger.go
        INFO  i.s.codegen.v3.AbstractGenerator - writing file /tmp/demo/./go/README.md
        INFO  i.s.codegen.v3.AbstractGenerator - writing file /tmp/demo/./.swagger-codegen-ignore
        INFO  i.s.codegen.v3.AbstractGenerator - writing file /tmp/demo/./.swagger-codegen/VERSION

## Run example

    % go run main.go
    2021/12/17 13:15:22 Server started
    2021/12/17 13:15:27 Process action published release name my first release
    2021/12/17 13:15:27 POST /release ReleasePost 173.369µs

## call example

    % curl -X 'POST' \
            'http://localhost:8080/release' \
            -H 'accept: application/json' \
            -H 'X-GitHub-Event: release' \
            -H 'X-Hub-Signature: sha1=abcd' \
            -H 'Content-Type: application/json' \
            -d '{
                    "action": "published",
                    "release": {
                            "name": "my first release",
                            "tag_name": "tag1"
                        },
                        "repository": {
                        "html_url": "https://github.com/opensvc/opensvc"
                    }
            }' \
            -v

    Note: Unnecessary use of -X or --request, POST is already inferred.
    * Trying ::1:8080...
      * Connected to localhost (::1) port 8080 (#0)
    > POST /release HTTP/1.1
    > Host: localhost:8080
    > User-Agent: curl/7.77.0
    > accept: application/json
    > X-GitHub-Event: release
    > X-Hub-Signature: sha1=abcd
    > Content-Type: application/json
    > Content-Length: 178
    >
    * Mark bundle as not supporting multiuse
      < HTTP/1.1 200 OK
      < Content-Type: application/json; charset=UTF-8
      < Date: Fri, 17 Dec 2021 12:16:45 GMT
      < Content-Length: 68
      <
      * Connection #0 to host localhost left intact


    {"message":"Process action published release name my first release"}