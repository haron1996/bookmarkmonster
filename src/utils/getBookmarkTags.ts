import { apiHost, bookmarkTags, selectedBookmarks, session } from "../stores/stores";
import type { Session } from "../types/session";

let apihost: string
let s: Partial<Session>
let bookmarkID: string

function getApiHost() {
    const unsub = apiHost.subscribe((value)=>{
        apihost = value
    })

    unsub()
}

function getSession() {
    const unsub = session.subscribe((value)=>{
        s = value
    })

    unsub()
}


function getBookmarkID() {
    const unsub = selectedBookmarks.subscribe((values) => {
        if (values[0].id) {
            bookmarkID = values[0].id
        }
    })

    unsub()
}

export async function getBookmarkTags(){
    getApiHost()

    getSession()

    getBookmarkID()

    const response = await fetch(`${apihost}/authenticated/bookmarks/bookmarkTags/${bookmarkID}`, {
        method: 'GET',
        mode: 'cors',
        cache: 'no-cache',
        credentials: 'include',
        headers: {
            'Content-Type': 'application/json',
            authorization: `Bearer${s.AccessToken}`
        },
        redirect: 'follow',
        referrerPolicy: 'no-referrer'
    });

    const result = await response.json()

    bookmarkTags.set(result[0])
}