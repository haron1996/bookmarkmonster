import { writable } from 'svelte/store';

export const pageSaved = writable<boolean>(false);
