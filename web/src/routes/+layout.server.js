import { GetMetadata, GetPageTitles } from '$lib/api';
import { getSiteName, getTagline, getUserName } from '$lib/metadata/helper';

export const load = async () => {

    var metadata = await GetMetadata();
    var pages = await GetPageTitles();

    return {
        siteName: getSiteName(metadata),
        siteTagline: getTagline(metadata),
        userName: getUserName(metadata),
        pages: pages,
    }
};