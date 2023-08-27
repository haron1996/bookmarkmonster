export function unselectBookmarkNodes() {
	const bookmarkNodes = document.querySelectorAll('.bookmark') as NodeListOf<HTMLDivElement> | null;

	if (bookmarkNodes === null) return;

	bookmarkNodes.forEach((bn) => {
		bn.classList.remove('bookmarkSelected');
	});
}
