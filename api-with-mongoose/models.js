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
  author: {
    type: Schema.Types.ObjectId,
    ref: 'User'
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

postSchema.pre('validate', function(next) {
  var err = null;
  console.log('validating: ', this);
  if (this.isModified('comments') && this.comments.length > 0) {
    this.comments.forEach((value, key, list) => {
      if (!value.text || !value.author.id) {
        err = new Error('Text and author for a comment must be set.')
      }
    })
  }
  if (err) return next(err);
  next();
});

postSchema.post('save', function(doc) {
  console.log('saved: ', doc);
});

postSchema.pre('save', function(next) {
  this.updatedAt = Date.now();
  next();
});

postSchema.virtual('hasComments').get(function() {
  return this.comments.length > 0;
});

postSchema.statics.staticMethod = function() {
  return new Promise((resolve, reject) => {
    // blah blah
    resolve();
  });
}

postSchema.methods.instanceMethod = function() {
  // 'this' is the instance
}
var userSchema = new Schema({
  name: String,
  role: {
    type: String,
    enum: roles
  }
});

module.exports = (connection) => {
  return {
    Post: connection.model('Post', postSchema, 'posts'),
    User: connection.model('User', userSchema, 'users')
  }
}

