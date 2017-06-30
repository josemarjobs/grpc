'use strict';

const PROTO_PATH = '../pb/messages.proto';

const fs = require('fs');
const grpc = require('grpc');
const serviceDef = grpc.load(PROTO_PATH);

const PORT = 9000;

const cacert = fs.readFileSync('./certs/ca.crt');
const cert = fs.readFileSync('./certs/client.crt');
const key = fs.readFileSync('./certs/client.key');
const kvpair = {
  'private_key': key,
  'cert_chain': cert
};

const creds = grpc.credentials.createSsl(cacert, key, cert);
const client = new serviceDef.EmployeeService(`localhost:${PORT}`, creds)

const option = parseInt(process.argv[2], 10);
switch (option) {
  case 1:
    sendMetadata(client);
    break;
  case 2:
    getByBadgeNumber(client);
    break;
}

function getByBadgeNumber(client) {
  client.getByBadgeNumber({badgeNumber: 7538}, (err, resp) => {
    if (err) {
      return console.log('Error: ', err);
    }
    console.log('Employee found: ', resp.employee)
  });
}

function sendMetadata(client) {
  const md = new grpc.Metadata();
  md.add('username', 'petergriffin');
  md.add('password', 'secretpass')
  
  client.getByBadgeNumber({}, md, (err) => {
    console.log(err);
  });
}