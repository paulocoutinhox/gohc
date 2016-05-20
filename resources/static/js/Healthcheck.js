var Healthcheck = new function () {

	this.STATUS_SUCCESS = "success";
	this.STATUS_WARNING = "warning";
	this.STATUS_ERROR   = "error";

	this.addHealthcheckToHTML = function (healthcheck) {
		var html = '' +
			'<div id="healthcheck-row-' + healthcheck.token + '" class="healthcheck-row list-group-item">' +
			'    <h4 class="list-group-item-heading">' + healthcheck.description + '</h4>' +
			'    <div class="list-group-item-text">' +
			//'        <div><strong>Job:</strong> <span class="ph-job-id-' + job.token + '"></span></div>' +
			//'        <div><strong>Created at:</strong> <span class="ph-job-created-at-' + job.token + '"></span></div>' +
			//'        <div><strong>Started at:</strong> <span class="ph-job-started-at-' + job.token + '"></span></div>' +
			//'        <div><strong>Finished at:</strong> <span class="ph-job-finished-at-' + job.token + '"></span></div>' +
			//'        <div><strong>Duration:</strong> <span class="ph-healthcheck-duration-' + healthcheck.token + '"></span></div>' +
			'        <div><strong>Status:</strong> <span class="ph-healthcheck-status-' + healthcheck.token + '"></span></div>' +
			//'        <div><strong>Progress:</strong> <span class="ph-healthcheck-progress-' + healthcheck.token + '">' +
			'       </span></div>' +
			'    </div>' +
			'</div>';

		$('#healthcheck-list').prepend(html);
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