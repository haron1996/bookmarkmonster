import { writable } from 'svelte/store';
import type { Tag } from '../types/tag';

export const sideBarWidth = writable<number>(25);

export const tags = writable<Tag[]>([]);
