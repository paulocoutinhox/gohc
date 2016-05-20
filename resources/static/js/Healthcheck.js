var Healthcheck = new function () {

	this.STATUS_SUCCESS = "success";
	this.STATUS_WARNING = "warning";
	this.STATUS_ERROR = "error";

	this.addHealthcheckToHTML = function (healthcheck) {
		var html = '' +
			'<div id="healthcheck-row-' + healthcheck.token + '" class="healthcheck-row list-group-item">' +
			'    <h4 class="list-group-item-heading">' + healthcheck.description + '</h4>' +
			'    <div class="list-group-item-text">' +
			'        <div><strong>Status:</strong> <span class="ph-healthcheck-status-' + healthcheck.token + '"></span></div>' +
			'    </div>' +
			'</div>';

		$('#healthcheck-list').append(html);
	};

	this.addHealthcheckToHTMLAsListItem = function (healthcheck) {
		var itemClass = '';
		var iconClass = '';

		if (healthcheck.status == Healthcheck.STATUS_SUCCESS) {
			itemClass = 'col-lg-2 col-md-3 col-sm-4 col-xs-12 healthcheck-item success';
			iconClass = 'icon glyphicon glyphicon-ok-sign';
		} else if (healthcheck.status == Healthcheck.STATUS_WARNING) {
			itemClass = 'col-lg-2 col-md-3 col-sm-4 col-xs-12 healthcheck-item warning';
			iconClass = 'icon glyphicon glyphicon-exclamation-sign';
		} else if (healthcheck.status == Healthcheck.STATUS_ERROR) {
			itemClass = 'col-lg-2 col-md-3 col-sm-4 col-xs-6 healthcheck-item error';
			iconClass = 'icon glyphicon glyphicon-remove-sign';
		}

		var html = '' +
			'<div id="healthcheck-row-' + healthcheck.token + '" class="' + itemClass + '">' +
			'    <div>' +
			'    <p class="' + iconClass + '"></p>' +
			'    <h5 class="title">' + healthcheck.description + '</h5>' +
			'    <h6 class="ping ph-healthcheck-ping-' + healthcheck.token + '">' + Util.msToHumanText(healthcheck.ping) + '</h6>' +
			'    </div>' +
			'</div>';

		$('#healthcheck-list').append(html);
	};

	this.clearHealthcheckList = function () {
		$('.healthcheck-row').remove();
	};

	this.list = function (preProcess, success, error) {
		if (!Util.isNullOrUndefined(preProcess)) {
			preProcess();
		}

		$.ajax({
			url: '/api/healthcheck/list',
			type: 'GET',
			dataType: 'json',
			success: function (response) {
				var wr = new WebResponse().parse(response);

				if (!Util.isNullOrUndefined(success)) {
					success(wr);
				}
			},
			error: function () {
				if (!Util.isNullOrUndefined(error)) {
					error();
				}
			}
		});
	};

	this.ping = function (token, preProcess, success, error) {
		if (!Util.isNullOrUndefined(preProcess)) {
			preProcess();
		}

		$.ajax({
			url: '/api/ping/' + token,
			type: 'GET',
			dataType: 'json',
			success: function (response) {
				var wr = new WebResponse().parse(response);

				if (!Util.isNullOrUndefined(success)) {
					success(wr);
				}
			},
			error: function () {
				if (!Util.isNullOrUndefined(error)) {
					error();
				}
			}
		});
	};

	this.count = function (preProcess, success, error) {
		if (!Util.isNullOrUndefined(preProcess)) {
			preProcess();
		}

		$.ajax({
			url: '/api/healthcheck/count',
			type: 'GET',
			dataType: 'json',
			success: function (response) {
				var wr = new WebResponse().parse(response);

				if (!Util.isNullOrUndefined(success)) {
					success(wr);
				}
			},
			error: function () {
				if (!Util.isNullOrUndefined(error)) {
					error();
				}
			}
		});
	};

};