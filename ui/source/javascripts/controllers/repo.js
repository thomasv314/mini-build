// Main Controller

tmbsApp.controller('RepoListCtrl', function($rootScope, $scope, $http) {

  $http.get($rootScope.tmbsURL).success(function(data) {
      $scope.Config = data;
      console.log(data)
  });
});
