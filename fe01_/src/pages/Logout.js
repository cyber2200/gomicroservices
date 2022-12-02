import React, { useEffect, useRef, useState } from 'react'
import { Navigate } from 'react-router-dom'
import { Auth } from '../models/Auth';

function Logout() {
    const effectRan = useRef(false)
    const [redirect, setRedirect] = useState(false)

    useEffect(() => {
        if (effectRan.current) return
        const logout = async () => {
            await Auth.logout()
        }
        logout()
        Auth.setCookie('session_id', '')
        setRedirect(true)
        return () => effectRan.current = true
    }, []);

    if (redirect) {
        return (
            <Navigate replace to="/" />
        )
    }

    return (
        <div className="App">
            Logout
        </div>
    );
}

export default Logout
