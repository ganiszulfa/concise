export const prerender = true;

import { GetPosts } from '$lib/api';

export const load = async ({ params }) => {

    var page = params.page ? params.page: 0;
    var Posts = await GetPosts(page);

    return {
        posts: Posts,
    }
};
