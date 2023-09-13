/** @type {import('./$types').PageLoad} */

import { browser } from '$app/environment';
import { redirect } from '@sveltejs/kit';
import { session } from '../../stores/stores';
import type { Bookmark } from '../../types/bookmark';
import { getUserRecentBookmarks } from '../../utils/getUserRecentBookmarks';
let recentUserBookmarks: Bookmark[] = [];

function getUserSession(url: any) {
	if (browser) {
		const sessionString: string | null = window.localStorage.getItem('session');

		if (sessionString === null) {
			throw redirect(302, `${url.origin}/signin`);
		}

		session.set(JSON.parse(sessionString));
	}
}

export async function load({ fetch, params, url, route }: any) {
	getUserSession(url);

	if (browser) {
		recentUserBookmarks = await getUserRecentBookmarks(fetch);
	}

	return { recentUserBookmarks };
}
