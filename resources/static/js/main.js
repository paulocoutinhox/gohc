var healthcheckList = [];

function loadHealthcheckListCount() {
    Healthcheck.count(function () {
        // ?
    }, function (response) {
        if (response.success) {
            var count = response.data.count;

            $(".ph-healthcheck-count-value").html(count);

            if (count == 1) {
                $(".ph-healthcheck-count-text").html("healthcheck");
            } else {
                $(".ph-healthcheck-count-text").html("healthchecks");
            }
        } else {
            $(".ph-healthcheck-count-value").html('0');
            $(".ph-healthcheck-count-text").html("healthchecks");
        }

        autoLoadHealthcheckListCount();
    }, function (error) {
        $(".ph-healthcheck-count-value").html('0');
        $(".ph-healthcheck-count-text").html("healthchecks");
        autoLoadHealthcheckListCount();
    });
}

function autoLoadHealthcheckListCount() {
    setTimeout(function () {
        loadHealthcheckListCount();
    }, 2000);
}

$(document).ready(function () {
    loadHealthcheckListCount();
});