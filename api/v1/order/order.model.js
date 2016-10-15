// This is an example object
// Model for Thing object
/*
    {
        "id"             :   String,
        "description"    :   String
    }

*/

function Order(    id,
                   KwhAmount,
                   PriceKwh,
                   TimeStart,
                   Duration,
                   SellerId,
                   SoldBool
               ){

    for (var key in arguments){
        if(!arguments[key]){
            throw new Error("Incorrect arguments for new Order.");
        }
    }

    // Attributes for Thing object
    this.id                  =   id;
    this.KwhAmount         =   KwhAmount;
    this.PriceKwh      =  PriceKwh;
    this.TimeStart          =  TimeStart;
    this.Duration     =  Duration;
    this.SellerId           = SellerId;
    this.SoldBool           = SoldBool;
}

var method = Order.prototype;

// Add a method to the CaseFile prototype
method.example = function(){
    return true;
}

module.exports = Order;

