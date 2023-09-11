export function openCreateFolder() {
	const div = document.getElementById('createFolder') as HTMLDivElement | null;

	if (div) {
		div.style.display = 'flex';
	}
}
