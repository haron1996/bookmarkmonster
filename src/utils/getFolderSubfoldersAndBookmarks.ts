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
export async function getFolderSubFoldersAndBookmarks(fetch: any, folderID: string): Promise<any> {
	getApiHost();

	getUserSession();

	const userRootFoldersPromise = await fetch(
		`${apiEndpoint}/authenticated/collections/getFolderSubfolders/${folderID}`,
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
		`${apiEndpoint}/authenticated/collections/getFolderBookmarks/${folderID}`,
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

	if (f.message) {
		return [[], []];
	}

	const b = await responses[1].json();

	if (b.message) {
		return [[], []];
	}

	return [f[0], b[0]];
}
