import { getBookmarkTags } from './getBookmarkTags';

export async function showBookmarkDetails() {
	const el = document.getElementById('bookmarkDetails') as HTMLDivElement | null;

	if (el === null) return;

	el.style.display = 'flex';

	await getBookmarkTags();
}
