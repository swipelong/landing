var main = {
	submit_email: function() {
		// Get the email address from the form.
		email_address = $("#email_input").val()
		// Validate the email address.
		if (!/^[\w._-]+[+]?[\w._-]+@[\w.-]+\.[a-zA-Z]{2,6}$/.test(email_address)) {
			// Don't continue.
			return main.show_error("Invalid email address.");
		}
		// Send it to the server.
		$.ajax({
			method: "POST",
			url: "/api",
			data: JSON.stringify({
				EmailAddress: email_address 
			}),
			success: function(srv_res) {
				main.show_success(srv_res);
			},
			error: function(srv_res) {
				console.log(srv_res);
				if (srv_res.responseText == "DUPLICATE_EMAIL_ADDRESS\n") {
					main.show_error("You've already submited that email address. <br>You can contact us directly at <a href='mailto:info@example.com'>info@example.com</a>!");
				} else {
					main.show_error("Oops... an unknown error occurred!");
				}
			}
		});
	},
	show_error: function(msg) {
			// Insert error message.
			$("#err_msg").html(msg);
			// Hide info and sccess messages.
			$("#info_msg").addClass("hidden");
			$("#success_msg").addClass("hidden");
			// Only show error.
			$("#err_msg").removeClass("hidden");
			// Don't continue.
			return;
	},
	show_success: function(msg) {
			// Hide info and sccess messages.
			$("#info_msg").addClass("hidden");
			$("#err_msg").addClass("hidden");
			// Only show error.
			$("#success_msg").removeClass("hidden");
			// Don't continue.
			return;
	}
}
$(document).ready(function() {
	console.log("Base script loaded...");
});
