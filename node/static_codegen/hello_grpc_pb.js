// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var hello_pb = require('./hello_pb.js');

function serialize_hello_Reply(arg) {
  if (!(arg instanceof hello_pb.Reply)) {
    throw new Error('Expected argument of type hello.Reply');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_hello_Reply(buffer_arg) {
  return hello_pb.Reply.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_hello_Request(arg) {
  if (!(arg instanceof hello_pb.Request)) {
    throw new Error('Expected argument of type hello.Request');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_hello_Request(buffer_arg) {
  return hello_pb.Request.deserializeBinary(new Uint8Array(buffer_arg));
}


var GreeterService = exports.GreeterService = {
  sayHello: {
    path: '/hello.Greeter/SayHello',
    requestStream: false,
    responseStream: false,
    requestType: hello_pb.Request,
    responseType: hello_pb.Reply,
    requestSerialize: serialize_hello_Request,
    requestDeserialize: deserialize_hello_Request,
    responseSerialize: serialize_hello_Reply,
    responseDeserialize: deserialize_hello_Reply,
  },
};

exports.GreeterClient = grpc.makeGenericClientConstructor(GreeterService);
