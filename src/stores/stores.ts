import { writable } from 'svelte/store';
import type { Tag } from '../types/tag';
import type { Bookmark } from '../types/bookmark';

export const sideBarWidth = writable<number>(25);

export const tags = writable<Tag[]>([]);

export const bookmarks = writable<Bookmark[]>([]);
