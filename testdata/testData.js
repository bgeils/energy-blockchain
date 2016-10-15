'use strict';

const testData = require('./testData.json');
const logger = require('../utils/logger');
const BlockchainService = require('../blockchainServices/blockchainSrvc');

const User = require('../api/v1/user/user.model')
const blockchain = require('../blockchain/blockchain');

exports.invokeTestData = function(){

    logger.info("[TestData] Deploying Test Data")

    resetIndexes(function() {
        writeUsersToLedger(testData.users);
        writeOrdersToLedger(testData.orders);
        writeTransactionsToLedger(testData.transactions);
    })

}

function resetIndexes(cb){
    logger.info("[TestData] Resetting indexes:");

    const functionName = "reset_indexes"
    const args = [];
    const enrollmentId = "WebAppAdmin";

    BlockchainService.invoke(functionName,args,enrollmentId).then(function(result){
        logger.info("[TestData] Index reset");
        cb()
    }).catch(function(err){
        logger.error(err);
    });

}

function writeUsersToLedger(users){
    logger.info("[TestData] Number of users:", testData.users.length);

    users.forEach( function(user, idx) {
        user = new User(user.userId, user.password, user.firstName, user.lastName, user.things, user.address, user.phoneNumber, user.emailAddress );

        let userAsJson = JSON.stringify(user);

        logger.info("[TestData] Will add new user:");
        logger.info(userAsJson);

        const functionName = "add_user"
        const args = [user.userId, userAsJson];
        const enrollmentId = "WebAppAdmin";

        BlockchainService.invoke(functionName,args,enrollmentId).then(function(result){
            logger.info("[TestData] Added user: ", user.userId);
        }).catch(function(err){
            logger.error(err);
        });
    })
}

function writeOrdersToLedger(orders){
    logger.info("[TestData] Number of orders:", testData.orders);
    
    orders.forEach(function(order, idx) {
        let orderAsJson = JSON.stringify(order);
        
        logger.info("[TestData] Will add new Order:");
        logger.info(orderAsJson);
        
        const functionName = "add_order"
        const args = [order.id, orderAsJson];
        const enrollmentId = "WebAppAdmin";

        BlockchainService.invoke(functionName,args,enrollmentId).then(function(result){
            logger.info("[TestData] Added order: ", order.id);
        }).catch(function(err){
            logger.error(err);
        });
    })
}

function writeTransactionsToLedger(transactions){
    logger.info("[TestData] Number of transactions:", testData.transactions);
    
    transactions.forEach(function(transaction, idx) {
        let transactionAsJson = JSON.stringify(transaction);
        
        logger.info("[TestData] Will add new Transaction:");
        logger.info(transactionAsJson);
        
        const functionName = "add_transaction"
        const args = [transaction.id, transactionAsJson];
        const enrollmentId = "WebAppAdmin";

        BlockchainService.invoke(functionName,args,enrollmentId).then(function(result){
            logger.info("[TestData] Added transaction: ", transaction.id);
        }).catch(function(err){
            logger.error(err);
        });
    })
}

