this_dir = File.expand_path(File.dirname(__FILE__))
proto = File.join(this_dir, 'proto')
$LOAD_PATH.unshift(proto) unless $LOAD_PATH.include?(proto)

require 'grpc'
require './proto/fizzbuzz_services_pb'

def main
  num = !ARGV.empty? ? ARGV[0].to_i : 0
  stub = Proto::FizzBuzz::Stub.new('localhost:50051', :this_channel_is_insecure)
  req = Proto::CalcFizzBuzzRequest.new(num: num)
  res = stub.calc_fizz_buzz(req).res
  puts "#{Time.now} #{res}"
end

main
