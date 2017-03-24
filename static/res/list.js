$(document).ready(function() {
	$('table.tablesorter').tablesorter();
	$('#filterBox div button').click(do_filter);
	$('#filterBox > button').click(reset_filter);
});

var filterDefault = {
	type: {},
	nation: {},
	tier: {}
};

var pQuery = new RegExp('^(\\d*),(\\d*),([\\da]*)');

var filter = $.extend(true, {}, filterDefault);
(function() {
	var s = window.location.search
	if (!s) {
		return
	}
	s = s.substr(1);

	var match = pQuery.exec(s);
	if (match) {
		var fill = function(s) {
			var r = {}
			for (var i = 0, j = s.length; i < j; i++) {
				var v = s[i];
				if (v == 'a') {
					v = 10;
				}
				r[v] = true;
			}
			return r;
		}
		filter.type = fill(match[1])
		filter.nation = fill(match[2])
		filter.tier = fill(match[3])
	}
})();

function do_filter() {
	var o = $(this);
	var i = o.data('value') - 0;
	var list = filter[o.data('type')];
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

var filterOrder = ['type', 'nation', 'tier'];
function buildUrl() {
	var q = [];
	var bSearch = false;
	for (var k in filterOrder) {
		var s = '';
		var row = filter[filterOrder[k]];
		if (!$.isEmptyObject(row)) {
			s = Object.keys(row).map(function(v) {
				return v > 9 ? 'a' : v;
			}).sort().join('');
			bSearch = true;
		}
		q.push(s);
	}
	var s = bSearch ? '?' + q.join(',') : '';
	if (window.location.search == s) {
		return;
	}
	window.history.pushState({}, 0, window.location.pathname + s);
}

function _do_filter() {

	buildUrl();
	//window.history.pushState({}, 0, '/');

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
