'use strict';

// Main Controller
angular.module('tmbsApp')
.controller('RepoListCtrl', function($rootScope, $scope, $http, $location) {

  $http.get($rootScope.tmbsURL).success(function(data) {
      $scope.Config = data;
      console.log(data)
  });
  
  $scope.viewCommit = function(commit) {
    $location.path("/commit/" + commit.Id)
  };

});
