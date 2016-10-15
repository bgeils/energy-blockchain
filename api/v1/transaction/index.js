'use strict';

var express = require('express');
var controller = require('./transaction.controller');

var router = express.Router();

router.get('/', controller.list);
router.get('/:transactionId', controller.detail);
//router.post('/', controller.add);

module.exports = router;
