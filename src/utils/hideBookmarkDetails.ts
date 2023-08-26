import { bookmarkTags } from '../stores/stores';
import { getBookmarkTags } from './getBookmarkTags';

export async function hideBookmarkDetails() {
	const el = document.getElementById('bookmarkDetails') as HTMLDivElement | null;

	if (el === null) return;

	el.style.display = 'none';

	bookmarkTags.set([]);
}
