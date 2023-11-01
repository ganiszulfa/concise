import { redirect } from '@sveltejs/kit';

export const load = async ({ cookies }) => {
	const userPassword = cookies.get('User-Password');
	console.log(userPassword);

	if (userPassword == undefined) {
		throw redirect(302, '/admin/login');
	}

	return {};
};
