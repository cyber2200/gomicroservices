import axios from 'axios'

export class Auth {
    static setCookie(name, value, days) {
        var expires = "";
        if (days) {
            var date = new Date();
            date.setTime(date.getTime() + (days * 24 * 60 * 60 * 1000));
            expires = "; expires=" + date.toUTCString();
        }
        document.cookie = name + "=" + (value || "") + expires + "; path=/";
    }

    static getCookie(cname) {
        let name = cname + "=";
        let decodedCookie = decodeURIComponent(document.cookie);
        let ca = decodedCookie.split(';');
        for (let i = 0; i < ca.length; i++) {
            let c = ca[i];
            while (c.charAt(0) === ' ') {
                c = c.substring(1);
            }
            if (c.indexOf(name) === 0) {
                return c.substring(name.length, c.length);
            }
        }
        return "";
    }

    static async checkSession() {
        const axiosRes = await axios.post(
            'http://localhost:3001/api/auth',
            {
                session_id: Auth.getCookie('session_id')
            }
        )
        if (axiosRes.data.res === 'OK') {
            return true
        }
        return false
    }

    static async logout() {
        await axios.post(
            'http://localhost:3001/api/logout',
            {
                session_id: Auth.getCookie('session_id')
            }
        )
    }
}