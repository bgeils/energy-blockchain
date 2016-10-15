'use strict';

const Order = require('./order.model');
const BlockchainService = require('../../../blockchainServices/blockchainSrvc.js');
const enrollID = require('../../../utils/enrollID')

/*
    Retrieve list of all orders

    METHOD: GET
    URL : /api/v1/order
    Response:
        [{'order'}, {'order'}]
*/
exports.list = function(req, res) {
    console.log("-- Query all orders --")
    
    var userID = enrollID.getID(req);
    
    const functionName = "get_all_orders"
    const args = [userID];
    const enrollmentId = userID;
    
    BlockchainService.query(functionName,args,enrollmentId).then(function(orders){
        if (!orders) {
            res.json([]);
        } else {
            console.log("Retrieved things from the blockchain: # " + orders.length);
            res.json(orders)
        }
    }).catch(function(err){
        console.log("Error", err);
        res.sendStatus(500);   
    }); 
}

/*
    Retrieve order object

    METHOD: GET
    URL: /api/v1/order/:orderId
    Response:
        { order }
*/
exports.detail = function(req, res) {
    console.log("-- Query thing --")
    
    const functionName = "get_order"
    const args = [req.params.orderId];
    const enrollmentId = enrollID.getID(req);
    
    BlockchainService.query(functionName,args,enrollmentId).then(function(order){
        if (!order) {
            res.json([]);
        } else {
            console.log("Retrieved order from the blockchain");
            res.json(order)
        }
    }).catch(function(err){
        console.log("Error", err);
        res.sendStatus(500);   
    }); 
}

/*
    Add thing object

    METHOD: POST
    URL: /api/v1/thing/
    Response:
        {  }
*/
// exports.add = function(req, res) {
//     console.log("-- Adding thing --")
      
//     const functionName = "add_thing"
//     const args = [req.body.thingId, JSON.stringify(req.body.thing)];
//     const enrollmentId = enrollID.getID(req);
    
//     BlockchainService.invoke(functionName,args,enrollmentId).then(function(thing){
//         res.sendStatus(200);
//     }).catch(function(err){
//         console.log("Error", err);
//         res.sendStatus(500);   
//     }); 
// }

