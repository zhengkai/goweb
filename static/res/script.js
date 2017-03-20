function logout() {
	$.post({
		url: '/passport/logout.do',
	}).always(function () {
		location.href = '/';
	});
}
