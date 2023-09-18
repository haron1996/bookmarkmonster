import { redirect } from '@sveltejs/kit';
import { apiHost, session } from '../stores/stores';
import type { Session } from '../types/session';
import { page } from '$app/stores';
import { goto } from '$app/navigation';

let apiEndpoint: string;
let userSession: Partial<Session>;
let host: string = '';

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

function getHost() {
	const unsub = page.subscribe((value) => {
		host = value.url.host;
	});

	unsub();
}

export async function getUserRootFoldersAndPlainBookmarks(fetch: any): Promise<any> {
	getApiHost();

	getUserSession();

	const userRootFoldersPromise = await fetch(
		`${apiEndpoint}/authenticated/collections/getUserRootCollections`,
		{
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
		}
	);

	const userRootPlainBookmarksPromise = await fetch(
		`${apiEndpoint}/authenticated/bookmarks/getRootBookmarks`,
		{
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
		}
	);

	const responses = await Promise.all([userRootFoldersPromise, userRootPlainBookmarksPromise]);

	const f = await responses[0].json();

	let msg = f.message;

	if (msg) {
		if (msg === 'token is no longer valid') {
			getApiHost();
			window.location.replace(`${host}/signin`);
		} else {
			alert(msg);
		}

		return;
	}

	const b = await responses[1].json();

	msg = b.message;

	if (msg) {
		if (msg === 'token is no longer valid') {
			getApiHost();
			window.location.replace(`${host}/signin`);
		} else {
			alert(msg);
		}

		return;
	}

	return [f[0], b[0]];
}
