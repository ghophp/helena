app.controller('MainController', ['$scope', '$timeout', '$http', function($scope, $timeout, $http){
	$scope.currentTry = 0;
	$scope.currentText = "";
	
	$scope.tries = [
		"nothing yet..",
		"nops.. nothing..",
		"that doesn't ring a bell :P",
		"what?!"
	];

	$scope.messages = [
		{
			text: "hello! I'm Helena. What is your name?",
			time: moment().format('LT'),
			received: true
		}
	];

	$scope.user = {
		id: "",
		name: "",
		favorites: [],
		emotions: []
	};

	$scope.$watchCollection('messages', function (newVal, oldVal) {
		$timeout(function(){
			var conversation = document.querySelector('.conversation-container');
			conversation.scrollTop = conversation.scrollHeight;
		}, 100);
	});

	$scope.think = function() {
		return Math.floor(Math.random() * (2000 - 500)) + 500;
	};

	$scope.try = function(tryList) {
		if ($scope.currentTry <= 0) {
			$scope.messages.push({
				text: tryList.shift(),
				time: moment().format('LT'),
				received: true
			});

			var next = function(){
				$scope.messages.push({
					text: tryList.shift(),
					time: moment().format('LT'),
					received: true
				});

				if (tryList.length > 0) {
					$timeout(next, $scope.think());					
				}
			};

			if (tryList.length > 0) {
				$timeout(next, $scope.think());					
			}

			$scope.currentTry++;
			return;
		}

		var r = Math.floor(Math.random() * ($scope.tries.length - 0)) + 0;
		$scope.messages.push({
			text: $scope.tries[r],
			time: moment().format('LT'),
			received: true
		});
	};

	$scope.welcome = function() {
		$scope.messages.push({
			text: "nice to meet you " + $scope.user.name + " :D",
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

	$scope.genres = function(genres) {
		if (genres && genres.length > 0) {
			// reset the try tick
			$scope.currentTry = 0;

			for (var x = 0; x < genres.length; x++) {
				$scope.user.favorites.push(genres[x].value);
			}
			
			$scope.messages.push({
				text: "cool! Those are the genres that you mentioned and I like also: " + $scope.user.favorites.join(", ") + ".. :)",
				time: moment().format('LT'),
				received: true
			});

			$timeout(function(){
				$scope.messages.push({
					text: "so tell me Mr. " + $scope.user.name + " how are you feeling right now?! for real.. happy, angry, depressed.. tell me how you feel..",
					time: moment().format('LT'),
					received: true
				});
			}, $scope.think());
		} else {
			$scope.try([
				"ok.. those genres must be super cool, but I don't know much about it, maybe you can be more genric, like.. instead of viking punk black metal, just sey metal.. :D"
			]);
		}
	};

	$scope.emotions = function(emotions) {
		if (emotions && emotions.length > 0) {
			// reset the try tick
			$scope.currentTry = 0;

			for (var x = 0; x < emotions.length; x++) {
				$scope.user.emotions.push(emotions[x].value);
			}
			$scope.messages.push({
				text: "thanks for sharing this with me ;) give me some time to look for something to match how you feel..",
				time: moment().format('LT'),
				received: true
			});

			// start the find process
			$scope.find();
		} else {
			$scope.try([
				"I am sorry.. your feelings are a little bit confusing for me.. it is not easy to be a machine :D",
				"anyways.. can you try to be more specific? I know about being alive, euphoric, bored, sad.. things like this ;)"
			]);
		}
	};

	$scope.find = function() {
		$http.post('/find', {
			genres: $scope.user.favorites,
			emotions: $scope.user.emotions
		}).success(function(response){

			if (response.tracks.length) {

				var tracks = response.tracks;
				var trackIds = [];
				for (var x = 0; x < tracks.length; x++) {
					trackIds.push(tracks[x].reference) 
				}

				var player = '<iframe src="https://embed.spotify.com/?uri=spotify:trackset:PREFEREDTITLE:'+
					trackIds.join(",")+
					'" frameborder="0" allowtransparency="true"></iframe>';
					
				$scope.messages.push({
					text: player,
					time: moment().format('LT'),
					received: true
				});

			} else {
				$scope.messages.push({
					text: ":/ seems that I didn't find anything that you like.. can you give another emotion maybe?",
					time: moment().format('LT'),
					received: true
				});
			}

		}).error(function(error){
			$scope.addError(error.statusText);
		});
	};

	$scope.addError = function(error) {
		$timeout(function(){
			$scope.messages.push({
				text: "ooops! I am having some problems with my wifi it seems.. can you send the message again? ("+error+")",
				time: moment().format('LT'),
				received: true
			});
		}, $scope.think());
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
				$http.jsonp('https://api.wit.ai/message?callback=JSON_CALLBACK', {
					params: {
						'q': currentText,
						'access_token' : 'MBZLWBX27FMYBXQSKSM5KZV2LL3PKOBO'
					}
				}).success(function(response){
					$scope.genres(response.entities.genre);
				}).error(function(error){
					$scope.addError(error.statusText);
				});
				return;
			}

			if ($scope.user.emotions.length <= 0) {
				$http.jsonp('https://api.wit.ai/message?callback=JSON_CALLBACK', {
					params: {
						'q': currentText,
						'access_token' : 'MBZLWBX27FMYBXQSKSM5KZV2LL3PKOBO'
					}
				}).success(function(response){
					$scope.emotions(response.entities.emotion);
				}).error(function(error){
					$scope.addError(error.statusText);
				});
				return;
			}
		});

		return false;
	};
}]);