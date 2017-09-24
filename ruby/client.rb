this_dir = File.expand_path(File.dirname(__FILE__))
lib_dir = File.join(this_dir, 'lib')
$LOAD_PATH.unshift(lib_dir) unless $LOAD_PATH.include?(lib_dir)

require 'grpc'
require 'hello_services_pb'

def main
  stub = Hello::Greeter::Stub.new('localhost:50051', :this_channel_is_insecure)
  message = stub.say_hello(Helloworld::HelloRequest.new(name: 'Ruby Client')).message
  p "Greeting: #{message}"
end

main
