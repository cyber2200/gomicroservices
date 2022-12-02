import React, { useEffect, useState, useRef } from 'react'
import { Outlet, Link } from "react-router-dom"
import { Auth } from '../../models/Auth'
import { Navigate } from 'react-router-dom'

function ProtectedLayout() {
    const effectRan = useRef(false)
    const [redirect, setRedirect] = useState(false)

    useEffect(() => {
        if (effectRan.current) return;
        const checkSession = async () => {
            const checkSessionRes = await Auth.checkSession()
            if (! checkSessionRes) {
                setRedirect(true)
            }
        }
        checkSession()
        return () => effectRan.current = true
    }, []);

    if (redirect) {
        return (
            <Navigate replace to="/" />
        )
    }

    return (
        <div>
            <Link to="/main">Home</Link> | <Link to="/main/about">About</Link> | <Link to="/main/dashboard">Dashboard</Link> | <Link to="/main/logout">Logout</Link>
            <hr />
            <Outlet />
        </div>
    );
}

export default ProtectedLayout