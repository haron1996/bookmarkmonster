export function openCreateBookmark() {
	const div = document.getElementById('createBookmark') as HTMLDivElement | null;

	if (div) {
		div.style.display = 'flex';
	}
}
