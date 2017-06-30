'use strict';

const PROTO_PATH = '../pb/messages.proto';

const fs = require('fs');
const grpc = require('grpc');
const serviceDef = grpc.load(PROTO_PATH);

const PORT = 9000;
let employees = require('./employees');

const cacert = fs.readFileSync('./certs/ca.crt');
const cert = fs.readFileSync('./certs/server.crt');
const key = fs.readFileSync('./certs/server.key');
const kvpair = {
  'private_key': key,
  'cert_chain': cert
};

const creds = grpc.ServerCredentials.createSsl(cacert, [kvpair]);
const server = new grpc.Server();
server.addService(serviceDef.EmployeeService.service, {
  getByBadgeNumber: getByBadgeNumber,
  getAll: getAll,
  addPhoto: addPhoto,
  saveAll: saveAll,
  save: save
});
server.bind(`0.0.0.0:${PORT}`, creds);
console.log(`server running on port ${PORT}`);
server.start();

function getByBadgeNumber(call, callback) {
  const md = call.metadata.getMap();
  for (let key in md) {
    console.log(key, md[key])
  }
  
  const badgeNumber = call.request.badgeNumber;
  for (let i = 0; i<employees.length; i++) {
    if (employees[i].badgeNumber === badgeNumber) {
      return callback(null, {employee: employees[i]}); 
    }
  }
  callback(new Error('employee not found.'));
}

function getAll(call) {
  employees.forEach(employee => {
    call.write({employee});
  });

  call.end();
}

function addPhoto(call, callback) {
  const md = call.metadata.getMap()
  const badgeNumber = md['badgenumber'];
  console.log('Adding photo for: ' + badgeNumber);

  let result = new Buffer(0);

  call.on('data', (data) => {
    result = Buffer.concat([result, data.data]);
    console.log('received ' + data.data.length + ' bytes');
  });

  call.on('end', () => {
    callback(null, {isOk: true});
    console.log('Total file size: ', result.length);
  });
}

function saveAll(call, callback) {
  call.on('data', (e) => {
    employees.push(e.employee);
    call.write({employee: e.employee});
  });

  call.on('end', () => {
    employees.forEach(e => {
      console.log(e);
    })

    call.end();
  })
}

function save(call, callback) {

}
