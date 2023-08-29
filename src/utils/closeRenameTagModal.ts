import { matchedTagsFromDB, selectedTags } from '../stores/stores';

export function closeRenameTagModal() {
	const el = document.getElementById('renameTag') as HTMLDivElement | null;

	if (el === null) return;

	el.style.backgroundColor = 'rgb(0, 0, 0, 0)';

	setTimeout(() => {
		el.style.display = 'none';
	}, 500);

	matchedTagsFromDB.set([]);
}
