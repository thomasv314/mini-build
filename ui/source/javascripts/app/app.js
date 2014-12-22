var tmbsApp = angular.module('tmbsApp', [
  'ngRoute'
])
.config(['$routeProvider', function($routeProvider) {
  $routeProvider
  .when('/commit/:CommitId', {
    templateUrl: 'commit-detail.html',
    controller: 'CommitDetailCtrl'
  })
  .when('/repos', {
    templateUrl: 'repo-list.html',
    controller: 'RepoListCtrl'
  })
  .otherwise({
    redirectTo: '/repos'
  });
}])
.run(function($rootScope) {

  $rootScope.tmbsURL = "http://localhost:59999/"

});

