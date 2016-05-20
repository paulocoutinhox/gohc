var HealthcheckList = new function () {

	this.hasHealthcheckById = function(healthcheckId) {
		var healthcheck = this.getHealthcheckById(healthcheckId);

		return (!Util.isNullOrUndefined(healthcheck));
	};

	this.addHealthcheck = function(healthcheck) {
		if (!this.hasHealthcheckById(healthcheck.token)) {
			healthcheck.lastStatus = "";
			healthcheckList.push(healthcheck);
		}
	};

	this.getHealthcheckById = function(healthcheckId) {
		for (var x = 0; x < healthcheckList.length; x++) {
			if (healthcheckList[x].token == healthcheckId) {
				return healthcheckList[x];
			}
		}

		return null;
	};

	this.getHealthcheckIndexById = function(healthcheckId) {
		for (var x = 0; x < healthcheckList.length; x++) {
			if (healthcheckList[x].token == healthcheckId) {
				return x;
			}
		}

		return null;
	};

	this.removeHealthcheckById = function(healthcheckId) {
		var healthcheckIndex = this.getHealthcheckIndexById(healthcheckId);

		if (!Util.isNullOrUndefined(healthcheckIndex)) {
			healthcheckList.splice(healthcheckIndex, 1);
		}
	};

	this.getHealthcheckLastStatusById = function(healthcheckId) {
		var healthcheck = this.getHealthcheckById(healthcheckId);

		if (!Util.isNullOrUndefined(healthcheck)) {
			return healthcheck.lastStatus;
		}

		return null;
	};

	this.updateHealthcheckLastStatusById = function(healthcheckId, status) {
		var healthcheck = this.getHealthcheckById(healthcheckId);

		if (!Util.isNullOrUndefined(healthcheck)) {
			healthcheck.lastStatus = status;
		}
	};

	this.clearList = function() {
		healthcheckList = [];
	};

	this.clearHealthcheckHtmlList = function () {
		$('.healthcheck-row').remove();
	};

	this.clearHealthcheckHtmlListItem = function () {
		$('.healthcheck-item').remove();
	};

};