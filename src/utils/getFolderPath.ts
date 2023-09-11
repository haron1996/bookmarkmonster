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
export async function getFolderPath(folderID: string, fetch: any): Promise<any> {
	getUserSession();

	getApiHost();

	const response = await fetch(
		`${apiEndpoint}/authenticated/collections/getFolderPath/${folderID}`,
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

	const result = await response.json();

	return result[0];
}
