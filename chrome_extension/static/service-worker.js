// chrome.action.onClicked.addListener(async (tab) => {
// 	// const accessToken = await chrome.cookies.getAll({
// 	// 	domain: 'localhost',
// 	// 	secure: true,
// 	// 	name: 'accessTokenCookie'
// 	// });
// 	// console.log(accessToken[0].value);
// 	// chrome.action.setPopup({
// 	// 	popup: 'index.html',
// 	// 	tabId: tab.id
// 	// });
// });

// chrome.bookmarks.getTree((tree) => {
// 	const nodes = tree[0].children;

// 	if (nodes) {
// 		for (const node of nodes) {
// 			const bookmarks = node.children;
// 			if (bookmarks) {
// 				for (const bookmark of bookmarks) {
// 					//console.log(bookmark);
// 				}
// 			}
// 		}
// 	}
// });

// chrome.runtime.onMessage.addListener(() => {});

// chrome.tabs.query({ active: true, currentWindow: true }, function (tabs) {
// 	const currentTab = tabs[0];
// 	//console.log(currentTab);
// });

// chrome.storage.local.set({ storedItem: { name: 'haron', password: 'password' } }, () => {
// 	//console.log('item stored');
// });

// chrome.storage.local.get(['storedItem'], ({ storedItem }) => {
// 	//console.log(storedItem);
// });

// chrome.tabs.onUpdated.addListener((tabId, changeInfo, tab) => {
// 	// Check for the condition you want (e.g., when the tab finishes loading)
// 	if (changeInfo.status === 'complete') {
// 		// Perform your desired actions here
// 		if (tab && tab.url) {
// 			console.log(`Tab ${tabId} was updated. URL: ${tab.url}`);
// 		}
// 	}
// });
