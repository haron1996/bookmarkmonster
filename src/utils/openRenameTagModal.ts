export function openRenameTagModal() {
	const el = document.getElementById('renameTag') as HTMLDivElement | null;

	if (el === null) return;

	el.style.display = 'flex';

	el.style.backgroundColor = 'rgb(0, 0, 0, .4)';

	const input = document.getElementById('renameTagInput') as HTMLInputElement | null;

	if (input === null) return;

	input.select();
}
