export function removeSelectedClassFromAllDomFolders() {
	const folders = document.querySelectorAll('a.folder') as NodeListOf<HTMLAnchorElement> | null;

	if (folders === null) return;

	folders.forEach((f) => {
		f.classList.remove('selected');
	});
}
