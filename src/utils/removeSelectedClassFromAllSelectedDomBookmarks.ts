export function removeSelectedClassFromAllDomBookmarks() {
	const bs = document.querySelectorAll('.bookmark') as NodeListOf<HTMLAnchorElement> | null;

	if (bs === null) return;

	bs.forEach((b) => {
		b.classList.remove('bookmarkSelected');
	});
}
