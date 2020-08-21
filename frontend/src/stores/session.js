import { writable}  from 'svelte/store';

export const { subscribe, set, update } = writable({ username: '', logged_in: false });

function updateSessionStatus() {
    fetch('/api/v1/auth/status')
        .then(resp => resp.json())
        .then(resp => {
            update(_ => resp);
        });
}

export default {
    subscribe,
    updateSessionStatus
};
