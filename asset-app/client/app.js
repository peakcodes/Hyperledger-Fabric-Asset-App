// SPDX-License-Identifier: Apache-2.0

'use strict';

var app = angular.module('application', []);

// Angular Controller
app.controller('appController', function($scope, appFactory){

	$("#success_holder").hide();
	$("#success_create").hide();
	$("#error_holder").hide();
	$("#error_query").hide();
	
	$scope.queryAllAsset = function(){

		appFactory.queryAllAsset(function(data){
			var array = [];
			for (var i = 0; i < data.length; i++){
				parseInt(data[i].Key);
				data[i].Record.Key = parseInt(data[i].Key);
				array.push(data[i].Record);
			}
			array.sort(function(a, b) {
			    return parseFloat(a.Key) - parseFloat(b.Key);
			});
			$scope.all_asset = array;
		});
	}

	$scope.queryAsset = function(){

		var id = $scope.asset_id;

		appFactory.queryAsset(id, function(data){
			$scope.query_asset = data;

			if ($scope.query_asset == "Could not locate asset"){
				console.log()
				$("#error_query").show();
			} else{
				$("#error_query").hide();
			}
		});
	}

	$scope.recordasset = function(){

		appFactory.recordAsset($scope.asset, function(data){
			$scope.create_asset = data;
			$("#success_create").show();
		});
	}

	$scope.changeHolder = function(){

		appFactory.changeHolder($scope.holder, function(data){
			$scope.change_holder = data;
			if ($scope.change_holder == "Error: no asset found"){
				$("#error_holder").show();
				$("#success_holder").hide();
			} else{
				$("#success_holder").show();
				$("#error_holder").hide();
			}
		});
	}

	$scope.changeCost = function(){

		appFactory.changeCost($scope.cost, function(data){
			$scope.change_cost = data;
			if ($scope.change_cost == "Error: no asset found"){
				$("#error_cost").show();
				$("#success_cost").hide();
			} else{
				$("#success_cost").show();
				$("#error_cost").hide();
			}
		});
	}

});

// Angular Factory
app.factory('appFactory', function($http){
	
	var factory = {};

    factory.queryAllAsset = function(callback){

    	$http.get('/get_all_asset/').success(function(output){
			callback(output)
		});
	}

	factory.queryAsset = function(id, callback){
    	$http.get('/get_asset/'+id).success(function(output){
			callback(output)
		});
	}

	factory.recordAsset = function(data, callback){

		data.location = data.location;

		var asset = data.id + "-" + data.location + "-" + data.item + "-" + data.holder + "-" + data.cost;

    	$http.get('/add_asset/'+asset).success(function(output){
			callback(output)
		});
	}

	factory.changeHolder = function(data, callback){

		var holder = data.id + "-" + data.name;

    	$http.get('/change_holder/'+holder).success(function(output){
			callback(output)
		});
	}

	return factory;
});


