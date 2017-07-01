const express = require('express');
const mongoose = require('mongoose');
const bodyParser = require('body-parser');
const logger = require('morgan');

let app = express();
app.use(logger('dev'));
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({extended: true}));

var dbUri = 'mongodb://localhost:27017/api';
let connection = mongoose.createConnection(dbUri);
mongoose.Promise = global.Promise;

let Post = require('./post')(connection);

app.get('/', (req, res) => {
  res.send('ok')
});

app.get('/posts', (req, res) => {
  Post.find({}).exec()
  .then(posts => res.send(posts))
  .catch(err => res.status(400).send(err));
});

app.post('/posts', (req, res) => {
  var post = new Post(req.body);
  post.validate()
  .then(() => console.log('Valid'))
  .catch(err => console.log('Error: ', err));

  post.save()
  .then(result => res.send(result))
  .catch(err => res.status(400).send(err));
});

app.get('/posts/:id', (req, res) => {
  Post.findById(req.params.id)
  .then(post => {
    if (!post) {
      return res.sendStatus(404);
    }
    res.json(post.toJSON({getters: true}));
  })
  .catch(err => res.status(400).send(err));
});

app.put('/posts/:id', (req, res) => {
  Post.findById(req.params.id)
  .then(post => {
    if (!post) {
      throw new Error('Post not fount');
    }
    post.set(req.body);
    return post.save();
  })
  .then(result => res.json(result.toJSON()))
  .catch(err => res.status(400).send(err))
})
app.listen(3000, () => {
  console.log('Server running on port 3000');
})
