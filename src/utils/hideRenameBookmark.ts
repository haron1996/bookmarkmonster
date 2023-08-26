export function hideRenameBookmark() {
	const el = document.getElementById('renameBookmarkComponent') as HTMLDivElement | null;

	if (el === null) return;

	el.style.display = 'none';
}
