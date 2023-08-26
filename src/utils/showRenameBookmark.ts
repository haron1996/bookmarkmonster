export function showRenameBookmark() {
	const el = document.getElementById('renameBookmarkComponent') as HTMLDivElement | null;

	if (el === null) return;

	el.style.display = 'flex';

	const input = document.getElementById('renameBookmarkInput') as HTMLInputElement | null;

	if (input === null) return;

	input.select();
}
