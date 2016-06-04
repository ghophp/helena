app.controller('MainController', ['$scope', '$timeout', function($scope, $timeout){
	$scope.currentText = "";
	$scope.messages = [
		{
			text: "Hello! I'm Helena. What is your name?",
			time: moment().format('LT'),
			received: true
		}
	];

	$scope.user = {
		id: "",
		name: "",
		favorites: [],
		emotion: ""
	};

	$scope.think = function() {
		return Math.random() * (2000 - 500) + 500;
	};

	$scope.welcome = function() {
		$scope.messages.push({
			text: "Nice to meet you " + $scope.user.name + " :D",
			time: moment().format('LT'),
			received: true
		});

		$timeout(function(){
			$scope.messages.push({
				text: "I heard that you like music as much as me, do you want to try something nice... First tell me your favorite, or favorites genres of music..",
				time: moment().format('LT'),
				received: true
			});
		}, $scope.think());
	};

	$scope.favorites = function() {

	};

	$scope.emotion = function() {
		//Tell me how are you feeling now..
	};

	$scope.addUserMessage = function(message, callback) {
		var current = {
			text: message,
			time: moment().format('LT'),
			received: false,
			ack: false
		};

		$scope.messages.push(current);
		$scope.currentText = "";

		$timeout(function(){
			current.ack = true;
			if (callback) {
				callback();	
			}
		}, $scope.think());
	};

	$scope.send = function() {
		if ($scope.currentText.length <= 0) {
			return;
		}

		var currentText = $scope.currentText;
		$scope.addUserMessage($scope.currentText, function(){
			if ($scope.user.name.length <= 0) {
				var name = currentText.replace(/[^\w\s]/gi,'');
				$scope.user.name = name;

				var explode = name.split(" ");
				if (explode.length > 1) {
					$scope.user.name = explode[0];
				}
				
				$scope.welcome();
				return;
			}

			if ($scope.user.favorites.length <= 0) {

				return;
			}

			if ($scope.user.emotion.length <= 0) {

				return;
			}
		});

		return false;
	};
}]);