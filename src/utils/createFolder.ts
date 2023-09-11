import { page } from '$app/stores';
import { apiHost, creatingFolder, error, folders, newFolder, session } from '../stores/stores';
import type { Folder } from '../types/folder';
import type { Session } from '../types/session';
import { closeCreateFolder } from './closeCreateFolder';
import { sortFoldersByName } from './sortFolders';

let apiEndpoint: string;
let userSession: Partial<Session>;
let parentFolderId: string = '';
let url: URL;
let folder: Partial<Folder> = {};

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

function getParentFolderId() {
	const unsub = page.subscribe((value) => {
		url = value.url;
	});

	unsub();

	let id: string | null = url.searchParams.get('id');

	if (id) {
		if (id === 'root') {
			parentFolderId = '';
		} else {
			parentFolderId = id;
		}
	}
}

function getFolder() {
	const unsub = newFolder.subscribe((value) => {
		folder = value;
	});

	unsub();
}

export async function createFolder() {
	creatingFolder.set(true);

	getApiHost();

	getUserSession();

	getParentFolderId();

	getFolder();

	const response = await fetch(`${apiEndpoint}/authenticated/collections/createCollection`, {
		method: 'POST',
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
			name: folder.folder_name,
			parent_folder_id: parentFolderId,
			description: folder.folder_description
		})
	});

	const result = await response.json();

	if (result.message) {
		alert(result.message);
		creatingFolder.set(false);
		newFolder.set({});
		return;
	}

	folders.update((values) => [result[0], ...values]);

	// sort folders

	newFolder.set({});

	closeCreateFolder();

	creatingFolder.set(false);
}
