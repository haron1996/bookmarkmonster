import type { Bookmark } from '../types/bookmark';

export function sortBookmarksByTitle(b: Bookmark[]) {
	b.sort((a, b) => {
		const ba = a.title?.toLowerCase(),
			bb = b.title?.toLowerCase();

		if (ba === undefined || bb === undefined) return 0;

		if (ba < bb) {
			return -1;
		}

		if (ba > bb) {
			return 1;
		}

		return 0;
	});
}
