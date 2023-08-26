import { browser } from '$app/environment';
import { error } from '../stores/stores';

export function removeError() {
	if (browser) {
		const div = document.getElementById('error') as HTMLDivElement | null;

		if (div === null) return;

		div.classList.remove('showError');

		setTimeout(() => {
			error.set('');
		}, 1000);
	}
}
