import React, { useEffect, useState, useRef } from 'react'
import axios from 'axios'
import { Navigate } from 'react-router-dom'
import { Auth } from '../models/Auth'

function Login() {
  const [email, setEmail] = useState("")
  const [password, setPassword] = useState("")
  const [formStatus, setFormStatus] = useState("")
  const [redirect, setRedirect] = useState(false)
  const effectRan = useRef(false)

  useEffect(() => {
    if (effectRan.current) return;
    const checkSession = async () => {
      const checkSessionRes = await Auth.checkSession()
      if (checkSessionRes) {
        setRedirect(true)
      }
    }
    checkSession()
    return () => effectRan.current = true
  }, []);

  const onLogin = async (e) => {
    e.preventDefault()
    try {
      setFormStatus("loading")
      const response = await axios.post(
        'http://localhost:3001/api/login',
        {
          email,
          password
        }
      );

      if (response.data.res === 'OK') {
        Auth.setCookie('session_id', response.data.session_id, 365)
        setFormStatus('Redirecting...')
        setRedirect(true)
      } else {
        setFormStatus('Email or password are woring, please try again')
      }

    } catch (err) {
      setFormStatus(err.message)
    }
  }

  if (redirect) {
    return (
      <Navigate replace to="/main" />
    )
  }

  return (
    <div className="App">
      <form onSubmit={onLogin}>
        <input type='text' placeholder='Email' onChange={(e) => setEmail(e.target.value)} />
        <input type='password' placeholder='Password' onChange={(e) => setPassword(e.target.value)} />
        <input type='submit' />
      </form>
      <div>
        {formStatus}
      </div>
    </div>
  );
}

export default Login
