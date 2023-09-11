import { sideBarVisible } from '../stores/stores';

let sideBarState: boolean;

function getSideBarState() {
	const unsub = sideBarVisible.subscribe((value) => {
		sideBarState = value;
	});

	unsub();
}

export function toggleSideBar() {
	const sideBar = document.getElementById('sideBar') as HTMLDivElement | null;

	if (sideBar === null) return;

	getSideBarState();

	sideBarState = !sideBarState;

	sideBarVisible.set(sideBarState);
}
