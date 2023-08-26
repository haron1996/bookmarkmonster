import { getBookmarkTags } from "./getBookmarkTags"

export async function showTagBookmarkComp() {
    const el = document.getElementById('tagBookmarkComponent') as HTMLDivElement | null

    if (el === null) return

    el.style.display= 'flex'

    await getBookmarkTags()
}