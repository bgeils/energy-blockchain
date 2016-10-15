'use strict';

const Transaction = require('./transaction.model');
const BlockchainService = require('../../../blockchainServices/blockchainSrvc.js');
const enrollID = require('../../../utils/enrollID')

/*
    Retrieve list of all transactions

    METHOD: GET
    URL : /api/v1/transaction
    Response:
        [{'transaction'}, {'transaction'}]
*/
exports.list = function(req, res) {
    console.log("-- Query all transactions --")
    
    var userID = enrollID.getID(req);
    
    const functionName = "get_all_transactions"
    const args = [userID];
    const enrollmentId = userID;

    BlockchainService.query(functionName,args,enrollmentId).then(function(transactions){
        if (!transactions) {
            res.json([]);
        } else {
            console.log("Retrieved things from the blockchain: # " + transactions.length);
            res.json(transactions)
        }
    }).catch(function(err){
        console.log("Error", err);
        res.sendStatus(500);   
    }); 
}

/*
    Retrieve transaction object

    METHOD: GET
    URL: /api/v1/transaction/:transactionId
    Response:
        { transaction }
*/
exports.detail = function(req, res) {
    console.log("-- Query thing --")
    
    const functionName = "get_transaction"
    const args = [req.params.transactionId];
    const enrollmentId = enrollID.getID(req);
    
    BlockchainService.query(functionName,args,enrollmentId).then(function(transaction){
        if (!transaction) {
            res.json([]);
        } else {
            console.log("Retrieved transaction from the blockchain");
            res.json(transaction)
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

