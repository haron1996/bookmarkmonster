import { writable } from 'svelte/store';

export const alertMessage = writable<string>('');
