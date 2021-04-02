#!ruby

require 'multi_json'

jsonfile = ARGV[0]

if File.exists? jsonfile
  json = IO.read jsonfile
  data = MultiJson.decode json

  endpoints = {}

  data['paths'].each do |url,methods|
    methods.each do |method,info|
      endpoint = method.upcase + ' ' + url
      puts endpoint
      endpoints[endpoint] = 1
    end
  end

  puts endpoints.length
end