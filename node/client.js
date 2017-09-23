var PROTO_PATH = __dirname + '/proto/hello.proto';

var grpc = require('grpc');
var hello_proto = grpc.load(PROTO_PATH).hello;

function main() {
  var client = new hello_proto.Greeter('localhost:50051',
                                       grpc.credentials.createInsecure());

  client.sayHello({name: "Node Client"}, function(err, response) {
    console.log('Greeting:', response.message);
  });
}

main();
