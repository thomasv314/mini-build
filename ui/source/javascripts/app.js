var tmbsApp = angular.module('tmbsApp', [
    'ngRoute'
    ]);

// Routes
tmbsApp.config(['$routeProvider', function($routeProvider) {
  $routeProvider.
  when('/repos', {
    templateUrl: 'partials/repo-list.html',
    controller: 'RepoListCtrl'
  })
  .otherwise({
    redirectTo: '/repos'
  });
}]);

tmbsApp.run(function($rootScope) {

  $rootScope.tmbsURL = "http://localhost:59999/"

});

