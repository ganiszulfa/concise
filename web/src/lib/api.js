import { GraphQLQuery } from '$lib/apiHelpers';

export async function GetPosts(page) {

    var query = "query ($page: Int!, $limit: Int!) {\r\n  ListPosts(isPublished: true, limit: $limit, page: $page) {\r\n    id\r\n    title\r\n    content\r\n    createdAt\r\n    updatedAt\r\n    publishedAt\r\n    slug\r\n    isPublished\r\n    isPage\r\n  }\r\n}\r\n";
    var variables = { "page": page, "limit": 10 };

    var Data = await GraphQLQuery(query, variables);

    if (Data != null) {
        return Data.ListPosts;
    }

    return [];
}


export async function GetMetadata() {

    var query = "{\r\n  ListMetadata {\r\n      key\r\n      value\r\n  }\r\n}\r\n";
    var Data = await GraphQLQuery(query, {})

    if (Data != null) {
        return Data.ListMetadata;
    }

    return [];
}

export async function GetPageTitles() {

    var query = "query ($page: Int!, $limit: Int!) {\r\n  ListPosts(isPublished: true, limit: $limit, page: $page, isPage: true) {\r\n    id\r\n    title\r\n    slug\r\n    isPublished\r\n    isPage\r\n  }\r\n}\r\n";
    var variables = { "page": 0, "limit": 10 };

    var Data = await GraphQLQuery(query, variables);

    if (Data != null) {
        return Data.ListPosts;
    }

    return [];
}

export async function GetAPost(slug) {

    var query = "query ($slug: String!) {\r\n  GetPost(slug: $slug) {\r\n    title\r\n    content\r\n    createdAt\r\n    updatedAt\r\n    publishedAt\r\n    slug\r\n    isPublished\r\n    isPage\r\n  }\r\n}\r\n"
    var variables = { "slug": slug };

    var Data = await GraphQLQuery(query, variables);

    if (Data != null) {
        return Data.GetPost;
    }

    return [];
}