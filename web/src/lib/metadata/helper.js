export function getSiteName(metadata) {
	return get(metadata, 'site name');
}

export function getTagline(metadata) {
	return get(metadata, 'site tagline');
}

export function getUserName(metadata) {
	return get(metadata, 'user name');
}

function get(arr, key) {
	let v = '';
	arr.every((e) => {
		if (e.key == key) {
			v = e.value;
			return false;
		}
		return true;
	});
	return v;
}
