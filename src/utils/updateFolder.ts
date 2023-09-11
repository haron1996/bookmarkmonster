import type { Folder } from '../types/folder';
import {
	apiHost,
	folders,
	selectedFolders,
	session,
	showUpdateFolder,
	updatingFolder
} from '../stores/stores';
import type { Session } from '../types/session';
import { removeSelectedClassFromAllDomFolders } from './removeSelectedClassFromAllDomFolders';

let apiEndpoint: string;
let userSession: Partial<Session>;
let folder: Folder[] = [];
let currentFolders: Folder[] = [];

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

function getSelectedFolder() {
	const unsub = selectedFolders.subscribe((value) => {
		folder = value;
	});

	unsub();
}

function getCurrentFolders() {
	const unsub = folders.subscribe((value) => {
		currentFolders = value;
	});

	unsub();
}

export async function updateFolder() {
	updatingFolder.set(true);

	getApiHost();

	getUserSession();

	getSelectedFolder();

	getCurrentFolders();

	const response = await fetch(`${apiEndpoint}/authenticated/collections/updateCollection`, {
		method: 'PATCH',
		mode: 'cors',
		cache: 'no-cache',
		credentials: 'include',
		headers: {
			'Content-Type': 'application/json',
			authorization: `Bearer${userSession.AccessToken}`
		},
		redirect: 'follow',
		referrerPolicy: 'no-referrer',
		body: JSON.stringify({
			name: folder[0].folder_name,
			description: folder[0].folder_description,
			id: folder[0].folder_id
		})
	});

	const result = await response.json();

	const index: number = currentFolders.findIndex((value) => {
		return value.folder_id === result[0].folder_id;
	});

	currentFolders.splice(index, 1, result[0]);

	folders.set(currentFolders);

	showUpdateFolder.set(false);

	selectedFolders.set([]);

	removeSelectedClassFromAllDomFolders();

	updatingFolder.set(false);
}
