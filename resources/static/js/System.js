var System = new function () {

	this.reload = function (preProcess, success, error) {
		if (!Util.isNullOrUndefined(preProcess)) {
			preProcess();
		}

		$.ajax({
			url: '/api/system/reload',
			type: 'GET',
			dataType: 'json',
			success: function (response) {
				var wr = new WebResponse().parse(response);

				if (!Util.isNullOrUndefined(success)) {
					Util.callAfterDelay(function () {
						success(wr);
					});
				}
			},
			error: function () {
				if (!Util.isNullOrUndefined(error)) {
					Util.callAfterDelay(function () {
						error();
					});
				}
			}
		});
	};

};