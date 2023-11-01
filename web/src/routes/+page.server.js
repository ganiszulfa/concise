export const prerender = false;

import { GetPosts } from '$lib/api';

export const load = async ({ url }) => {
	var paramPage = url.searchParams.get('p');
	var page = paramPage ? parseInt(paramPage) : 1;
	var Posts = await GetPosts(page);

	return {
		posts: Posts,
		page: page
	};
};
