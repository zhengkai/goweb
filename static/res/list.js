var sortColumn = {index: -1, asc: true}
var filterDefault = {
	type: {},
	nation: {},
	tier: {}
};

$(document).ready(function() {
	var o = $('table.tablesorter');
	o.tablesorter({
		sortInitialOrder: 'desc',
		sortMultiSortKey: 'none'
	});
	$('#filterBox div button').click(do_filter);
	$('#filterBox > button').click(reset_filter);
	initSort();
	o.bind('sortEnd', sortUpdate);
});

function initSort() {
	var hash = window.location.hash;
	if (hash.length < 2) {
		return;
	}
	hash = hash.substring(1);
	var index = parseInt(hash)|0;
	if (index < 0) {
		return;
	}
	sortColumn.index = index;
	sortColumn.asc = hash.substr(-1) === 'a';

	var o = $('table.tablesorter th[data-column=' + index + ']');
	o.click();
	if (sortColumn.asc) {
		o.click();
	}
}

function sortUpdate() {
	var l = $('table th').toArray();
	var found = false;
	for (var i in l) {
		var o = $(l[i]);
		var s = o.attr('aria-sort');
		if (s !== 'none') {
			found = true;
			sortColumn.index = i;
			sortColumn.asc = (s === 'ascending');
			break;
		}
	}
	if (!found) {
		sortColumn.index = -1;
	}
	buildUrl();
}

var pQuery = new RegExp('^(\\d*)(,|.)(\\d*)(,|.)([\\dax]*)');

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
				if (v == 'a' || v == 'x') {
					v = 10;
				}
				r[v] = true;
			}
			return r;
		}
		filter.type = fill(match[1])
		filter.nation = fill(match[3])
		filter.tier = fill(match[5])
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
				return v > 9 ? 'x' : v;
			}).sort().join('');
			bSearch = true;
		}
		q.push(s);
	}
	var s = bSearch ? '?' + q.join('.') : '';
	if (sortColumn.index >= 0) {
		s += '#' + sortColumn.index + (sortColumn.asc ? 'a' : '');
	}
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
