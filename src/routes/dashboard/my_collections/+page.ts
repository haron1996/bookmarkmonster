/** @type {import('./$types').PageLoad} */

import { browser } from '$app/environment';
import { session } from '../../../stores/stores';
import type { Bookmark } from '../../../types/bookmark';
import { getUserRootFoldersAndPlainBookmarks } from '../../../utils/getRootFoldersAndPlainBookmarks';
import type { Folder } from '../../../types/folder';
import { getUserRootFoldersAndBookmarksInTrash } from '../../../utils/getBookmarksAndFoldersInTrash';
import { getFolderSubFoldersAndBookmarks } from '../../../utils/getFolderSubfoldersAndBookmarks';
import { getFolderPath } from '../../../utils/getFolderPath';
import { redirect } from '@sveltejs/kit';

let bookmarks: Bookmark[] = [];
let folders: Folder[] = [];
let folderPath: Folder[] = [];

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
		let which: string | null = url.searchParams.get('which');
		if (which) {
			if (which === 'recycle_bin') {
				[folders, bookmarks] = await getUserRootFoldersAndBookmarksInTrash(fetch);
			}
		} else {
			let id: string | null = url.searchParams.get('id');
			if (id) {
				if (url.searchParams.get('id') === 'root') {
					[folders, bookmarks] = await getUserRootFoldersAndPlainBookmarks(fetch);
					folderPath = [];
				} else {
					[folders, bookmarks] = await getFolderSubFoldersAndBookmarks(fetch, id);
					folderPath = await getFolderPath(id, fetch);
				}
			}
		}
	}

	return { folders, streamed: { bookmarks, folderPath } };
}
