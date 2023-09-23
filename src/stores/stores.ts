import { derived, readable, writable } from 'svelte/store';
import type { Tag } from '../types/tag';
import type { Bookmark } from '../types/bookmark';
import type { Session } from '../types/session';
import type { Folder } from '../types/folder';
import type { Screenshot } from '../types/screenshot';

export const apiHost = readable<string>('https://api.bookmarkmonster.xyz');

//export const apiHost = readable<string>('http://localhost:5000');

export const sideBarWidth = writable<number>(25);

export const hideSideBar = writable<boolean>(false);

export const session = writable<Partial<Session>>({});

export const tags = writable<Tag[]>([]);

export const bookmarks = writable<Bookmark[]>([]);

export const folders = writable<Folder[]>([]);

export const lastAddedBookmark = writable<Bookmark>({});

export const processingBookmark = writable<boolean>(false);

export const currentTagID = writable<string>('');

export const query = writable<string>('');

export const selectedBookmarks = writable<Bookmark[]>([]);

export const tagName = writable<string>('');

// alerts
export const error = writable<string>('');
export const successMessage = writable<string>('');
export const errorMessage = writable<string>('');
// end of alerts

export const ctrlKeyActive = writable<boolean>(false);

export const deletedTag = writable<Tag>({});

export const indexOfDeletedTag = writable<number | null>();

export const sideBarVisible = writable<boolean>(false);

export const newFolder = writable<Partial<Folder>>({});

export const selectedFolders = writable<Folder[]>([]);

export const creatingFolder = writable<boolean>(false);

export const folderPath = writable<Folder[]>([]);

export const loadingItems = writable<boolean>(true);

export const updatingFolder = writable<boolean>(false);

export const showUpdateFolder = writable<boolean>(false);

export const selectedItems = derived([selectedBookmarks, selectedFolders], ($allSelected) => [
	...$allSelected[0],
	...$allSelected[1]
]);

export const allItems = derived([bookmarks, folders], ($all) => [...$all[0], ...$all[1]]);

export const showUpdateBookmark = writable<boolean>(false);

export const updatingBookmark = writable<boolean>(false);

export const showAddNewBookmark = writable<boolean>(false);

export const recentUserBookmarks = writable<Bookmark[]>([]);

export const toolTipText = writable<string>('');

// move items operations
export const showMoveItemsPopup = writable<boolean>(false);
export const destinationFolder = writable<Folder>({});
export const destinationFolderPath = writable<Folder[]>([]);
export const childrenOfDestinationFolder = writable<Folder[]>([]);
export const bookmarksOfDestinationFolder = writable<Bookmark[]>([]);
export const showCreateFolderFromMoveFoldersPopup = writable<boolean>(false);

// drag and drop
export const foldersToMove = writable<Folder[]>([]);

// end of move items operations

export const breadCrumbIsOverflowing = writable<boolean>(false);
export const collapsedFolders = writable<Folder[]>([]);
export const showCollapsedFolders = writable<boolean>(false);
export const showQuickView = writable<boolean>(false);
export const showCaptureScreeshot = writable<boolean>(false);
export const screenshot = writable<Screenshot>({});
export const screenshots = writable<Screenshot[]>([]);

// tags
export const currentTag = writable<Tag>({});
export const selectedTag = writable<Tag>({});
export const tagsBarCollapsed = writable<boolean>(false);
export const selectedTags = writable<Tag[]>([]);
export const matchedTagsFromDB = writable<Tag[]>([]);
export const bookmarkTags = writable<Tag[]>([]);
export const showCreatBookmarkFromTagPage = writable<boolean>(false);
// end of tags
