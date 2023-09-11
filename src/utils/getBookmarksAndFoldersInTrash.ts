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
export async function getUserRootFoldersAndBookmarksInTrash(fetch: any): Promise<any> {
	getApiHost();

	getUserSession();

	const userRootFoldersPromise = await fetch(
		`${apiEndpoint}/authenticated/collections/getUserFoldersInRecycleBin`,
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
		`${apiEndpoint}/authenticated/bookmarks/getBookmarksInTrash`,
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

	const b = await responses[1].json();

	return [f[0], b[0]];
}
