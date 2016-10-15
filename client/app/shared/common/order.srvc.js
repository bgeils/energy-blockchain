app.service('OrdersService', ["$q", "$http", "$localStorage", function ($q, $http, $localStorage) {
    
    return {
        getAllOrders: function () {
            var deferred = $q.defer();
            
            console.log("OrdersService -- Get assigned orders based on userId: ", $localStorage.user.userId);
                        
            $http({
                method: 'GET',
                url: '/api/v1/order/'
            }).then(function success(response) {
                deferred.resolve(response.data);
            }, function error(error) {
                deferred.reject(error);
            });

            return deferred.promise;
        },
        getOrder: function () {
            var deferred = $q.defer();
            
            console.log("OrderService -- Get order with id: ", $localStorage.selectedOrder);
                        
            $http({
                method: 'GET',
                url: '/api/v1/order/'+$localStorage.selectedOrder
            }).then(function success(response) {
                deferred.resolve(response.data);
            }, function error(error) {
                deferred.reject(error);
            });

            return deferred.promise;
        }

    }

}])