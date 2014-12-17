require 'json'
require 'net/http'
require 'uri'

# Change to test
url = 'http://localhost:59999/push/mini-build'
filename = './json-samples/bitbucket-multiple-commits.json'

# Load & encode
file = File.read(filename)
json = JSON.parse(file).to_json.to_s
json = URI.escape(json.to_s)
res = Net::HTTP.post_form(URI(url), payload: json)
puts "[#{res}] - #{res.message} - #{res.body}"
