function validate_email(field,alerttxt) {
	with (field) {
		apos = value.indexOf("@");
		dotpos = value.lastIndexOf(".");
		if (apos<1||dotpos-apos<2) {
			alert(alerttxt);
			return false;
		} else {
			return true;
		}
	}
}
function topassword(password, checkpassword) {
    //document.write("Hello World!");
    var p1 = password.value;
    var p2 = checkpassword.value;

    if (p1 == p2) {
        //alert("password right");
        return true;
    } else {
        alert("password wrong");
        return false;
    }
}
function validate_form(thisform) {
	with (thisform) {
		if (validate_email(email,"Not a valid e-mail address!")==false) {
			email.focus();
			return false;
		}
		return topassword(password, checkpassword);
	}
}

$.fn.serializeObject = function() {
	var o = {};
	var a = this.serializeArray();
	$.each(a, function(){
		if (o[this.name]){
			if (!o[this.name].push) {
				o[this.name] = [o[this.name]];
			}
			o[this.name].push(this.value || '');
		} else{
			o[this.name] = this.value || '';
		}
	});
	return o;
}

$('#registerform2').submit(function(e) {
	e.preventDefault();

	// send json  
	var v = $("#registerform2").serializeObject();
    
    $.ajax({
        type    :   'POST',
        url     :   '/api/register',
        cache   :   false,
		data: JSON.stringify(v),
        contentType: "application/json",
        processData:false,
        dataType:'json',
        success:function(response) {
            if (response.result == 1) {
                $('.registermessage').addClass('alert alert-success').text(response.message);
            } else {
                $('.registermessage').addClass('alert alert-danger').text(response.message);
            }
        }
    });
})

$('#registerform').submit(function(e) {
	e.preventDefault();

	// send json  
	var v = $("#registerform").serializeObject();
    
    $.ajax({
        type    :   'POST',
        url     :   '/api/register',
        cache   :   false,
		data: JSON.stringify(v),
        contentType: "application/json",
        processData:false,
        dataType:'json',
        success:function(response) {
            if (response.result == 1) {
                $('.registermessage').addClass('alert alert-success').text(response.message);
            } else {
                $('.registermessage').addClass('alert alert-danger').text(response.message);
            }
        }
    });
	
	// send form
	/*
    var form = $(this);
    var formdata = false;
    if(window.FormData) {
        formdata = new FormData(form[0]);
    }
    var formAction = form.attr('action');
    $.ajax({
        type    :   'POST',
        url     :   '/api/register',
        cache   :   false,
        data    :   formdata ? formdata : form.serialize(),
        contentType: false,
        processData:false,
        dataType:'json',
        success:function(response) {
            if (response.result == 1) {
                $('.registermessage').addClass('alert alert-success').text(response.message);
            } else {
                $('.registermessage').addClass('alert alert-danger').text(response.message);
            }
        }
    });
	*/
})

$('#loginform, #loginform2').submit(function(e) {
	// send json  
	var v = $("#loginform, #loginform2").serializeObject();
    
    $.ajax({
        type    :   'POST',
        url     :   '/api/login',
        cache   :   false,
		data: JSON.stringify(v),
        contentType: "application/json",
        processData:false,
        dataType:'json',
        success:function(response) {
            if (response.result == 1) {
                $('.loginmessage').addClass('alert alert-success').text(response.message);
                location.href = response.next;
            } else {
                $('.loginmessage').addClass('alert alert-danger').text(response.message);
            }
        }
    });
	// send form
	/*
	var form = $(this);
    var formdata = false;
	if(window.FormData) {
        formdata = new FormData(form[0]);
    }
    var formAction = form.attr('action');
	$.ajax({
        type    :   'POST',
        url     :   '/api/login',
        cache   :   false,
        data    :   formdata ? formdata : form.serialize(),
        contentType: false,
        processData:false,
        dataType:'json',
        success:function(response) {
            if (response.result == 1) {
                $('.loginmessage').addClass('alert alert-success').text(response.message);
                location.href = response.next;
            } else {
                $('.loginmessage').addClass('alert alert-danger').text(response.message);
            }
        }
    });
	*/
	
    e.preventDefault();
});


$('#mylogin').on('hidden.bs.modal', function () {
    $(this).find('form').trigger('reset');
    $('#loginmessage').removeClass('alert alert-danger').text('');
})

$('#post-form').submit(function(e) {
	e.preventDefault();

	// send json  
	var v = $("#post-form").serializeObject();
    
    $.ajax({
        type    :   'POST',
        url     :   '/api/post',
        cache   :   false,
		data: JSON.stringify(v),
        contentType: "application/json",
        processData:false,
        dataType:'json',
        success:function(response) {
            if (response.result == 1) {
                $('.postmessage').addClass('alert alert-success').text(response.message);
				location.href = response.next;
            } else {
                $('.postmessage').addClass('alert alert-danger').text(response.message);
            }
        }
    });
})