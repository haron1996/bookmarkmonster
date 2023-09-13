import { apiHost, session } from '../stores/stores';
import type { Session } from '../types/session';
import type { Bookmark } from '../types/bookmark';

let apiEndpoint: string;
let userSession: Partial<Session>;

function getApiHost() {
	const unsub = apiHost.subscribe((value) => {
		apiEndpoint = value;
	});

	unsub();
}

function getUserSession() {
	const unsub = session.subscribe((value) => {
		userSession = value;
	});

	unsub();
}
export async function getPlainBookmarksOnly(fetch: any): Promise<Bookmark[]> {
	getApiHost();

	getUserSession();

	const response = await fetch(`${apiEndpoint}/authenticated/bookmarks/getRootBookmarks`, {
		method: 'get',
		mode: 'cors',
		cache: 'no-cache',
		credentials: 'include',
		headers: {
			'Content-Type': 'application/json',
			authorization: `Bearer${userSession.AccessToken}`
		},
		redirect: 'follow',
		referrerPolicy: 'no-referrer'
	});

	const result = await response.json();

	return result[0];
}
