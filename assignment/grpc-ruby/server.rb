this_dir = File.expand_path(File.dirname(__FILE__))
proto = File.join(this_dir, 'proto')
$LOAD_PATH.unshift(proto) unless $LOAD_PATH.include?(proto)

require 'grpc'
require './proto/fizzbuzz_services_pb'

class Server < Proto::FizzBuzz::Service
  def calc_fizz_buzz(req, _nouse)
    Proto::CalcFizzBuzzReply.new(res: fizzbuzz(req.num))
  end

  private

  def fizzbuzz(num)
    return "FizzBuzz" if (num % 15).zero?
    return "Buzz" if (num % 5).zero?
    return "Fizz" if (num % 3).zero?
    num.to_s
  end
end

def main
  s = GRPC::RpcServer.new
  s.add_http2_port('0.0.0.0:50051', :this_port_is_insecure)
  s.handle(Server)
  s.run_till_terminated
end

main
