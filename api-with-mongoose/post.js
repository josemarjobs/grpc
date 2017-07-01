const mongoose = require('mongoose');
var Schema = mongoose.Schema;

var roles = ['user', 'admin', 'staff', 'guest'];

var postSchema = new Schema({
  title: {
    type: String,
    required: true,
    trim: true,
    set: (value) => {
      return value.toUpperCase();
    },
    get: (value) => {
      return value.toLowerCase();
    }
  },
  text: {
    type: String,
    required: true,
    trim: true,
    max: 2000
  },
  followers: [Schema.Types.ObjectId],
  meta: Schema.Types.Mixed,
  comments: [{
    text: {
      type: String,
      trim: true,
      max: 2000
    },
    author: {
      id: {
        type: Schema.Types.ObjectId,
        ref: 'User'
      },
      name: String,
      role: {
        type: String,
        enum: roles
      }
    }
  }],
  viewCounter: {
    type: Number,
    validate: (value) => {
      if (value < 0) {
        return false;
      } else {
        return true
      }
    }
  },
  published: Boolean,
  createdAt: {
    type: Date,
    default: Date.now()
  },
  updatedAt: {
    type: Date,
    default: Date.now()
  }
});

module.exports = (connection) => {
  return connection.model('Post', postSchema, 'posts');
}

