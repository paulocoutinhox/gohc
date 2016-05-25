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
				setTimeout(function() {
					var wr = new WebResponse().parse(response);

					if (!Util.isNullOrUndefined(success)) {
						success(wr);
					}
				}, 1500);
			},
			error: function () {
				setTimeout(function() {
					if (!Util.isNullOrUndefined(error)) {
						error();
					}
				}, 1500);
			}
		});
	};

};