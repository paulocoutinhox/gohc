package domain

import (
	"fmt"
	"strconv"
	"strings"
)

func MailCreateSubject(healthcheck Healthcheck, healthcheckNotifier HealthcheckNotifier) string {
	if healthcheck.Status == HEALTHCHECK_STATUS_SUCCESS {
		return fmt.Sprintf("Healthcheck - %v - Success", healthcheck.Description)
	} else if healthcheck.Status == HEALTHCHECK_STATUS_WARNING {
		return fmt.Sprintf("Healthcheck - %v - Warning", healthcheck.Description)
	} else if healthcheck.Status == HEALTHCHECK_STATUS_ERROR {
		return fmt.Sprintf("Healthcheck - %v - Error", healthcheck.Description)
	} else if healthcheck.Status == HEALTHCHECK_STATUS_TIMEOUT {
		return fmt.Sprintf("Healthcheck - %v - Timeout", healthcheck.Description)
	}

	return ""
}

func MailCreateBody(healthcheck Healthcheck, healthcheckNotifier HealthcheckNotifier) string {
	body := `
	<table class="main" width="100%" cellpadding="0" cellspacing="0" style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; border-radius: 3px; background-color: #fff; margin: 0; border: 1px solid #e9e9e9;" bgcolor="#fff">
		<tbody>
			<tr style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; margin: 0;">
				<td class="alert alert-warning" style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 16px; vertical-align: top; color: #fff; font-weight: 500; text-align: center; border-radius: 3px 3px 0 0; background-color: [status-color]; margin: 0; padding: 20px;" align="center" bgcolor="[status-color]" valign="top">
					Healthcheck [status] alert
				</td>
			</tr>
			<tr style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; margin: 0;">
				<td class="content-wrap" style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; vertical-align: top; margin: 0; padding: 20px;" valign="top">
					<table width="100%" cellpadding="0" cellspacing="0" style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; margin: 0;">
						<tbody>
							<tr style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; margin: 0;">
								<td class="content-block" style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; vertical-align: top; margin: 0; padding: 0 0 20px;" valign="top">
									You have are receiving a new <strong style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; margin: 0;">[status]</strong> alert.
								</td>
							</tr>
							<tr style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; margin: 0;">
								<td class="content-block" style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; vertical-align: top; margin: 0; padding: 0 0 20px;" valign="top">
									Description: <strong style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; margin: 0;">[description]</srong>
								</td>
							</tr>`

	if healthcheck.Type != HEALTHCHECK_TYPE_MANUAL {
		body += `
							<tr style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; margin: 0;">
								<td class="content-block" style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; vertical-align: top; margin: 0; padding: 0 0 20px;" valign="top">
									[type-text]: <strong style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; margin: 0;">[type-value]</strong>
								</td>
							</tr>`
	}

	body += `
							<tr style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; margin: 0;">
								<td class="content-block" style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; vertical-align: top; margin: 0; padding: 0 0 20px;" valign="top">
									Status: <strong style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; margin: 0;">[status]</strong>
								</td>
							</tr>
							<tr style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; margin: 0;">
								<td class="content-block" style="font-family: 'Helvetica Neue',Helvetica,Arial,sans-serif; box-sizing: border-box; font-size: 14px; vertical-align: top; margin: 0; padding: 20px 0 20px;" valign="top">
									GoHC agent.
								</td>
							</tr>
						</tbody>
					</table>
				</td>
			</tr>
		</tbody>
	</table>
	`

	body = strings.Replace(body, "[status]", healthcheck.Status, -1)
	body = strings.Replace(body, "[description]", healthcheck.Description, -1)

	if healthcheck.Type == HEALTHCHECK_TYPE_PING {
		body = strings.Replace(body, "[type-text]", "Ping value", -1)
		body = strings.Replace(body, "[type-value]", strconv.FormatInt(healthcheck.Ping, 10), -1)
	} else if healthcheck.Type == HEALTHCHECK_TYPE_RANGE {
		body = strings.Replace(body, "[type-text]", "Range value", -1)
		body = strings.Replace(body, "[type-value]", strconv.FormatFloat(healthcheck.Range, 'f', 2, 64), -1)
	} else if healthcheck.Type == HEALTHCHECK_TYPE_MANUAL {
		body = strings.Replace(body, "[type-text]", "Manual", -1)
	}

	if healthcheck.Status == HEALTHCHECK_STATUS_SUCCESS {
		body = strings.Replace(body, "[status-color]", "#1ab394", -1)
	} else if healthcheck.Status == HEALTHCHECK_STATUS_WARNING {
		body = strings.Replace(body, "[status-color]", "#f8ac59", -1)
	} else if healthcheck.Status == HEALTHCHECK_STATUS_ERROR {
		body = strings.Replace(body, "[status-color]", "#ed5565", -1)
	} else if healthcheck.Status == HEALTHCHECK_STATUS_TIMEOUT {
		body = strings.Replace(body, "[status-color]", "#263238", -1)
	}

	return body
}
