'use strict';

var express = require('express');
var controller = require('./order.controller');

var router = express.Router();

router.get('/', controller.list);
router.get('/:orderId', controller.detail);
//router.post('/', controller.add);

module.exports = router;
