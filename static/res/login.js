$(document).ready(function() {

	$('#form-alert').click(function () {
		// $(this).hide()
	});

	// $('#login-form button, #register-form button').prop('disabled', false);
	$('#login-form form, #register-form form').submit(function() {
		parseForm($(this));
		return false;
	});
});

function parseForm(oForm) {
	var oAlert = $('#form-alert')
	oAlert.hide();
	oForm.append(oAlert);

	var name = oForm.find('input[name=name]').val()
	var password = oForm.find('input[name=password]').val()

	if (!name.length) {
		showAlert('需要输入用户名', true)
		oForm.find('input[name=name]').focus()
		return;
	}

	if (!password.length) {
		showAlert('需要输入密码', true)
		oForm.find('input[name=password]').focus()
		return;
	}

	if (oForm.find('input[name=password2]').length) {
		passwordRe = oForm.find('input[name=password2]').val()
		if (password != passwordRe) {
			console.log('password diff', password, passwordRe)
			showAlert('两次输入的密码不一致', true)
			oForm.find('input[name=password2]').val('')
			oForm.find('input[name=password]').val('').focus()
			return;
		}
	}

	$.post({
		url: oForm.attr('action'),
		data: oForm.serialize(),
		dateType: 'json'
	}).always(function (data) {
		parseResult(data, oForm);
	});

	oForm.find('button[type=submit]').prop('disabled', true);
}

function parseResult(data, oForm) {

	// console.log('parseForm', data);

	oForm.find('button[type=submit]').prop('disabled', false);

	if (!data.error && !data.uid) {
		data = {
			error: '提交失败，可能是网络问题或服务器错误'
		};
	}

	if (data.error) {
		showAlert('错误：' + data.error, true)
		return;
	}

	showAlert('完成，即将跳转到首页', false)
	location.href = '/'
}

function showAlert(text, isError) {
	var oAlert = $('#form-alert');
	if (isError) {
		oAlert.removeClass('alert-success').addClass('alert-danger').show();
	} else {
		oAlert.removeClass('alert-danger').addClass('alert-success').show();
	}
	oAlert.text(text).show();
}
