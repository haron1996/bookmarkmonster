import { bookmarks, ctrlKeyActive, selectedBookmarks } from '../stores/stores';

import type { Bookmark } from '../types/bookmark';

export function selectBookmark(e: MouseEvent) {
	const target = e.currentTarget as HTMLElement;

	const bookmark = target.closest('.bookmark') as HTMLDivElement | null;

	if (bookmark === null) return;

	// let added;
	// let updated;

	// if (bookmark.dataset.added) {
	// 	added = new Date(bookmark.dataset.added);
	// }

	// if (bookmark.dataset.updated) {
	// 	updated = new Date(bookmark.dataset.updated);
	// }

	const b: Bookmark = {
		id: bookmark.dataset.id,
		title: bookmark.dataset.title,
		user_id: bookmark.dataset.userid,
		added: bookmark.dataset.added,
		bookmark: bookmark.dataset.bookmark,
		deleted: bookmark.dataset.deleted,
		favicon: bookmark.dataset.favicon,
		host: bookmark.dataset.host,
		notes: bookmark.dataset.notes,
		thumbnail: bookmark.dataset.thumbnail,
		updated: bookmark.dataset.updated
	};

	let ctrlKeyIsActive: boolean = false;

	function getCtrlKeyState() {
		const unsub = ctrlKeyActive.subscribe((value) => {
			ctrlKeyIsActive = value;
		});

		unsub();
	}

	getCtrlKeyState();

	let sB: Bookmark[] = [];

	function getSelectedBookmarks() {
		const unsub = selectedBookmarks.subscribe((values) => {
			sB = values;
		});

		unsub();
	}

	getSelectedBookmarks();

	if (ctrlKeyIsActive) {
		// check if bookmark exists in selected bookmarks
		if (
			sB.filter((value) => {
				return value.id === b.id;
			})[0] === undefined
		) {
			// if not exists? add
			selectedBookmarks.update((values) => [b, ...values]);

			// mark b as selected
			bookmark.classList.add('bookmarkSelected');
		} else {
			// if exists? remove
			sB = sB.filter((value) => {
				return value.id !== b.id;
			});

			selectedBookmarks.set(sB);
			// unselect b
			bookmark.classList.remove('bookmarkSelected');
		}
	} else {
		selectedBookmarks.set([]);
		selectedBookmarks.update((values) => [b, ...values]);
		const bookmarkNodes = document.querySelectorAll(
			'.bookmark'
		) as NodeListOf<HTMLDivElement> | null;

		if (bookmarkNodes === null) return;

		bookmarkNodes.forEach((bn) => {
			bn.classList.remove('bookmarkSelected');
		});

		bookmark.classList.add('bookmarkSelected');
	}
}
