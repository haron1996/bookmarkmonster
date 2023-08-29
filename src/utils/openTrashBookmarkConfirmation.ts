export async function openTrashBookmarkConfirmation() {
	const el = document.getElementById('trashBookmarkConfirmation') as HTMLDivElement | null;

	if (el === null) return;

	el.style.display = 'flex';
}
