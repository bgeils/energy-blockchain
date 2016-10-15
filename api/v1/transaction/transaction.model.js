// This is an example object
// Model for Thing object
/*
    {
        "id"             :   String,
        "description"    :   String
    }

*/

function Transaction(    id,
                   OrderId,
                   Seller,
                   Buyer
               ){

    for (var key in arguments){
        if(!arguments[key]){
            throw new Error("Incorrect arguments for new Transaction.");
        }
    }

    // Attributes for Thing object
    this.id          =  id;
    this.OrderId     =  OrderId;
    this.Seller      =  Seller;
    this.Buyer       =  Buyer;
}

var method = Transaction.prototype;

// Add a method to the CaseFile prototype
method.example = function(){
    return true;
}

module.exports = Transaction;