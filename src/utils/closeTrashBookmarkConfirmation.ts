import { getBookmarkTags } from './getBookmarkTags';

export async function closeTrashBookmarkConfirmation() {
	const el = document.getElementById('trashBookmarkConfirmation') as HTMLDivElement | null;

	if (el === null) return;

	el.style.display = 'none';
}
