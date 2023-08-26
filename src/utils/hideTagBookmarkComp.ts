import { bookmarkTags, matchedTagsFromDB, tagName } from '../stores/stores';

export async function hideTagBookmarkComp() {
	const el = document.getElementById('tagBookmarkComponent') as HTMLDivElement | null;

	if (el === null) return;

	el.style.display = 'none';

	bookmarkTags.set([]);

	matchedTagsFromDB.set([]);

	tagName.set('');
}
