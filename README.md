<!--

        server:"localhost:4040"
        client:"localhost:8080"
        -----------------------------------------------------------------------------------


        --------------   .pb.go     -----------------------

        -> Request struct
        -> Request methods

        -> Response struct
        -> Response methods

        -> file

        -> client:
            - interface
            - struct
            - New constructor
            - rpc microservice methods

        -> server:
            -
 -->

<!--
#1.
    - we have to be guided through the .proto file and after it's compilation it has no use other than to get refrence

    - Now, .pb.go has all the required code which is to be implemented inorder to setup our server and client

#2.
    - both request and response has its own structs and it's own method (with some extra fields)

    - most importantly Get...() methods , those who returning data field

#3.
    - next we'll look at server so inside server and create main.go

    - aim is implement RegisterMicroserviceServer
        -- implement the server interface ( that is already defined as method below) as per Unimplemented struct below
        -- both request and response are of struct type and defined in .pb.go file only not in .proto
        -- now in main() {} , call the RegisterMicroserviceServer()
            --- need grpcServer
            --- need server struct


#4.
    - inside client main.go

    - connect to the server

    - NewMicroserviceClient constructor
 -->
