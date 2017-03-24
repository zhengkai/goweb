$(document).ready(function() {
	$('table.tablesorter').tablesorter();
	$('#filterBox div button').click(do_filter);
	$('#filterBox > button').click(reset_filter);
});

var filter = {
	type: {},
	nation: {},
	tier: {}
};

var filterDefault = $.extend(true, {}, filter);

function do_filter() {
	var type = $(this).data('type');
	var i = $(this).data('value') - 0;

	list = filter[type];
	if (i in list) {
		delete list[i];
	} else {
		list[i] = true;
	}
	_do_filter();
}

function reset_filter() {
	filter = $.extend(true, {}, filterDefault);
	_do_filter();
}

function _do_filter() {
	$('#filterBox button').each(function() {

		var o = $(this)
		var type = o.data('type');

		var row = filter[type];

		if ($.isEmptyObject(row)) {
			o.addClass('btn-outline-primary').removeClass('btn-primary')
			return;
		}

		if (o.data('value') in row) {
			o.addClass('btn-primary').removeClass('btn-outline-primary')
		} else {
			o.addClass('btn-outline-primary').removeClass('btn-primary')
		}
	});

	$('#tankList tr').each(function () {
		var hide = false
		var o = $(this)
		for (var k in filter) {
			var row = filter[k];
			if ($.isEmptyObject(row)) {
				continue;
			}

			var val = o.data(k);
			if (!(val in row)) {
				hide = true;
				break
			}
		}
		if (hide) {
			o.hide();
		} else {
			o.show();
		}
	});
}
