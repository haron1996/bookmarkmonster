import { readable, writable } from 'svelte/store';
import type { Tag } from '../types/tag';
import type { Bookmark } from '../types/bookmark';
import type { Session } from '../types/session';

export const apiHost = readable<string>('https://api.bookmarkmonster.xyz');

export const sideBarWidth = writable<number>(25);

export const session = writable<Partial<Session>>({});

export const tags = writable<Tag[]>([]);

export const bookmarks = writable<Bookmark[]>([]);

export const processingBookmark = writable<boolean>(false);

export const currentTagID = writable<string>('');
