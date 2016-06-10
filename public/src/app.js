var app = angular.module('myapp', ['ui.router'])

.config(['$stateProvider', '$urlRouterProvider', function($stateProvider, $urlRouterProvider) {
	$urlRouterProvider.otherwise("/");
	$stateProvider
		.state('index', {
			url: "/",
			templateUrl: "public/src/app/controllers/main.html"
		});
}]);
