export function closeCreateBookmark() {
	const div = document.getElementById('createBookmark') as HTMLDivElement | null;

	if (div) {
		div.style.display = 'none';
	}
}
