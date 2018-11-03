const path = require('path')
const caller = require('grpc-caller')

const PROTO_PATH = __dirname + '/pb/cutter.proto'
const client = caller('0.0.0.0:19003', PROTO_PATH, 'Cutter')

client.cutImage({ data: 'Bob' }, (err, res) => {
  console.log(res)
})
