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
  case 3:
    getAll(client);
    break;
  case 4:
    addPhoto(client);
    break;
}

function addPhoto(client) {
  let md = new grpc.Metadata();
  md.add('badgenumber', '7538');
  
  const call = client.addPhoto(md, (err, result) => {
    if (err) {
      return console.log('Error: ', err);
    }
    console.log('Result: ', result);
  });
  let photo = '../img.jpg';
  if (process.argv.length === 4) {
    photo = process.argv[3];
  }
  const stream = fs.createReadStream(photo);
  stream.on('data', (chunk) => {
    call.write({data: chunk});
  });

  stream.on('end', () => {
    call.end();
  })
}

function getAll(client) {
  let call = client.getAll({});
  call.on('data', (e) => {
    console.log(e.employee);
  });
  call.on('end', () => {
    console.log('Done.');
  })
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