var landing_script = {
	submit_email: function() {
		console.log("submitting email...");
		$.ajax({
			method: "POST",
			url: "/",
			data: {
				email_address: "SOEMTHING",
				testing: "Testing"
			},
			success: function(response) {
				alert("Success!");
			},
			error: function(response) {
				alert("Error!");
			}
		});
	}
}
$(document).ready(function() {
	console.log("Base script loaded...");
});
