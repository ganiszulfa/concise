import { GetAPost } from '$lib/api';

export const load = async ({ params }) => {
	var post = await GetAPost(params.slug);
	console.log(post);
	return {
		post: post
	};
};
