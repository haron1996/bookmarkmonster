import { newFolder } from '../stores/stores';

export function closeCreateFolder() {
	const div = document.getElementById('createFolder') as HTMLDivElement | null;

	if (div) {
		div.style.display = 'none';
	}

	newFolder.set({});
}
