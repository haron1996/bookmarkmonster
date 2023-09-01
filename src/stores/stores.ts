import { readable, writable } from 'svelte/store';
import type { Tag } from '../types/tag';
import type { Bookmark } from '../types/bookmark';
import type { Session } from '../types/session';

//export const apiHost = readable<string>('https://api.bookmarkmonster.xyz');

export const apiHost = readable<string>('http://localhost:5000');

export const sideBarWidth = writable<number>(25);

export const session = writable<Partial<Session>>({});

export const tags = writable<Tag[]>([]);

export const bookmarks = writable<Bookmark[]>([]);

export const lastAddedBookmark = writable<Bookmark>({});

export const processingBookmark = writable<boolean>(false);

export const currentTagID = writable<string>('');

export const query = writable<string>('');

export const selectedBookmarks = writable<Bookmark[]>([]);

export const selectedTags = writable<Tag[]>([]);

export const matchedTagsFromDB = writable<Tag[]>([]);

export const bookmarkTags = writable<Tag[]>([]);

export const tagName = writable<string>('');

export const error = writable<string>('');

export const ctrlKeyActive = writable<boolean>(false);

export const deletedTag = writable<Tag>({});

export const indexOfDeletedTag = writable<number | null>();
