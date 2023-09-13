import { redirect } from '@sveltejs/kit';
import { apiHost, session } from '../stores/stores';
import type { Session } from '../types/session';

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

export async function getAllUserBookmarks(fetch: any): Promise<any> {
	getApiHost();
	getUserSession();
	const promise = await fetch(`${apiEndpoint}/authenticated/bookmarks/getAllUserBookmarks`, {
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

	const result = await promise.json();

	if (result.message) {
		alert(result.message);
		throw redirect(302, '/signin');
	}

	return result[0];
}
